package dexscreener

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	baseURI = "wss://io.dexscreener.com/dex/screener/v5/pairs/h24/1"
)

type DexScreenerClient struct {
	conn     *websocket.Conn
	debug    bool
	done     chan struct{}
	handlers []func(Pair)
}

type Pair struct {
	Chain                  string      `json:"chain"`
	Protocol               string      `json:"protocol"`
	PairAddress            string      `json:"pairAddress"`
	BaseTokenName          string      `json:"baseTokenName"`
	BaseTokenSymbol        string      `json:"baseTokenSymbol"`
	BaseTokenAddress       string      `json:"baseTokenAddress"`
	Price                  string      `json:"price"`
	PriceUsd               string      `json:"priceUsd"`
	PriceChange            PriceChange `json:"priceChange"`
	Liquidity              Liquidity   `json:"liquidity"`
	Volume                 Volume      `json:"volume"`
	Fdv                    string      `json:"fdv"`
	PairCreatedAt          int64       `json:"pairCreatedAt"`
	PairCreatedAtFormatted string      `json:"pairCreatedAtFormatted"`
}

type PriceChange struct {
	H24 string `json:"h24"`
}

type Liquidity struct {
	Usd string `json:"usd"`
}

type Volume struct {
	H24 string `json:"h24"`
}

func NewDexScreenerClient(debug bool) *DexScreenerClient {
	return &DexScreenerClient{
		debug:    debug,
		done:     make(chan struct{}),
		handlers: make([]func(Pair), 0),
	}
}

func (c *DexScreenerClient) OnPair(handler func(Pair)) {
	c.handlers = append(c.handlers, handler)
}

func (c *DexScreenerClient) handleDouble(value float64) float64 {
	if value != value || value == 0 {
		return 0
	}
	return value
}

func (c *DexScreenerClient) decodeMetrics(data []byte, startPos int) (map[string]float64, int) {
	if startPos+64 > len(data) {
		return nil, startPos
	}

	metrics := make(map[string]float64)
	values := make([]float64, 8)

	reader := bytes.NewReader(data[startPos : startPos+64])
	binary.Read(reader, binary.LittleEndian, &values)

	valueMap := map[string]float64{
		"price":          values[0],
		"priceUsd":       values[1],
		"priceChangeH24": values[2],
		"liquidityUsd":   values[3],
		"volumeH24":      values[4],
		"fdv":            values[5],
		"timestamp":      values[6],
	}

	for key, value := range valueMap {
		if cleaned := c.handleDouble(value); cleaned != 0 {
			metrics[key] = cleaned
		}
	}

	return metrics, startPos + 64
}

func (c *DexScreenerClient) cleanString(s string) string {
	if s == "" {
		return ""
	}

	// Remove non-printable characters except spaces
	var cleaned strings.Builder
	for _, char := range s {
		if (char >= 32 && char < 127) || char == 9 {
			cleaned.WriteRune(char)
		}
	}

	result := cleaned.String()
	if strings.Contains(result, "@") || strings.Contains(result, "\\") {
		parts := strings.Split(result, "@")
		if len(parts) > 0 {
			result = parts[0]
		}
		parts = strings.Split(result, "\\")
		if len(parts) > 0 {
			result = parts[0]
		}
	}

	return strings.TrimSpace(result)
}

func (c *DexScreenerClient) decodePair(data []byte) *Pair {
	pos := 0
	pair := &Pair{}

	// Skip binary prefix
	for pos < len(data) && (data[pos] == 0x00 || data[pos] == 0x0A) {
		pos++
	}

	fields := []string{"chain", "protocol", "pairAddress", "baseTokenName", "baseTokenSymbol", "baseTokenAddress"}

	for _, field := range fields {
		if pos >= len(data) {
			break
		}

		strLen := int(data[pos])
		pos++

		if strLen == 0 || strLen > 100 || pos+strLen > len(data) {
			continue
		}

		value := c.cleanString(string(data[pos : pos+strLen]))
		if value != "" {
			switch field {
			case "chain":
				pair.Chain = value
			case "protocol":
				pair.Protocol = value
			case "pairAddress":
				pair.PairAddress = value
			case "baseTokenName":
				pair.BaseTokenName = value
			case "baseTokenSymbol":
				pair.BaseTokenSymbol = value
			case "baseTokenAddress":
				pair.BaseTokenAddress = value
			}
		}
		pos += strLen
	}

	// Align to 8-byte boundary
	pos = (pos + 7) &^ 7

	metrics, pos := c.decodeMetrics(data, pos)

	if len(metrics) > 0 {
		if price, ok := metrics["price"]; ok {
			pair.Price = fmt.Sprintf("%f", price)
		}
		if priceUsd, ok := metrics["priceUsd"]; ok {
			pair.PriceUsd = fmt.Sprintf("%f", priceUsd)
		}
		if priceChange, ok := metrics["priceChangeH24"]; ok {
			pair.PriceChange = PriceChange{H24: fmt.Sprintf("%f", priceChange)}
		}
		if liquidity, ok := metrics["liquidityUsd"]; ok {
			pair.Liquidity = Liquidity{Usd: fmt.Sprintf("%f", liquidity)}
		}
		if volume, ok := metrics["volumeH24"]; ok {
			pair.Volume = Volume{H24: fmt.Sprintf("%f", volume)}
		}
		if fdv, ok := metrics["fdv"]; ok {
			pair.Fdv = fmt.Sprintf("%f", fdv)
		}
		if timestamp, ok := metrics["timestamp"]; ok && timestamp >= 0 && timestamp < 4102444800 {
			pair.PairCreatedAt = int64(timestamp)
			pair.PairCreatedAtFormatted = time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
		}
	}

	// Validate minimum required data
	if len(pair.Chain) > 0 && len(pair.Protocol) > 0 &&
		(pair.Price != "0" || pair.PriceUsd != "0" ||
			pair.Volume.H24 != "0" || pair.Liquidity.Usd != "0") {
		return pair
	}

	return nil
}

func (c *DexScreenerClient) Connect(ctx context.Context, trendingScore string) error {
	params := url.Values{}
	params.Add("rankBy[key]", trendingScore)
	params.Add("rankBy[order]", "desc")
	params.Add("filters[chainIds][0]", "solana")

	uri := fmt.Sprintf("%s?%s", baseURI, params.Encode())

	headers := http.Header{
		"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:132.0) Gecko/20100101 Firefox/132.0"},
		"Accept":          []string{"*/*"},
		"Accept-Language": []string{"en-US,en;q=0.5"},
		"Origin":          []string{"https://dexscreener.com"},
		"Pragma":          []string{"no-cache"},
		"Cache-Control":   []string{"no-cache"},
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
	}

	conn, _, err := dialer.Dial(uri, headers)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	c.conn = conn

	go c.readLoop(ctx)
	return nil
}

func (c *DexScreenerClient) readLoop(ctx context.Context) {
	defer c.conn.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.done:
			return
		default:
			_, message, err := c.conn.ReadMessage()

			fmt.Printf("message: %s\n", string(message))
			if err != nil {
				if c.debug {
					log.Printf("Error reading message: %v", err)
				}
				continue
			}

			if string(message) == "ping" {
				c.conn.WriteMessage(websocket.TextMessage, []byte("pong"))
				continue
			}

			if bytes.HasPrefix(message, []byte{0x00, 0x0A, '1', '.', '3', '.', '0', 0x0A}) {
				pairsStart := bytes.Index(message, []byte("pairs"))
				if pairsStart == -1 {
					continue
				}

				pairs := make([]Pair, 0)
				pos := pairsStart + 5

				for pos < len(message) {
					pair := c.decodePair(message[pos:])
					if pair != nil {
						pairs = append(pairs, *pair)
					}
					pos += 512
				}

				if len(pairs) > 0 {
					for _, pair := range pairs {
						for _, handler := range c.handlers {
							handler(pair)
						}
					}
				}
			}
		}
	}
}

func (c *DexScreenerClient) Close() {
	close(c.done)
	if c.conn != nil {
		c.conn.Close()
	}
}
