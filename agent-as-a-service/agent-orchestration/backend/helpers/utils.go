package helpers

import (
	"bytes"
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"mime"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/gocolly/colly"
	"github.com/leekchan/accounting"
	"mvdan.cc/xurls/v2"
)

func SubStringBodyResponse(obj string, limit int) string {
	if len(obj) > limit {
		return obj[0:limit]
	}
	return obj
}

func ExtractTweetID(link string) string {
	re := regexp.MustCompile(`/status/(\d+)`)

	matches := re.FindStringSubmatch(link)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func ReplaceQuote(obj string) string {
	obj = strings.ReplaceAll(obj, "'", "''")
	obj = strings.ReplaceAll(obj, "\\", "\\\\")
	return obj
}

func SimpleAddress(hexAddress string) string {
	return fmt.Sprintf("%s...%s", hexAddress[:4], hexAddress[len(hexAddress)-4:])
}

func MergeMetaInfoURL(baseURL string, mediaURL string) string {
	if strings.HasPrefix(mediaURL, "http") || strings.HasPrefix(mediaURL, "ipfs") {
		return mediaURL
	}
	return fmt.Sprintf("%s/%s", baseURL, mediaURL)
}

func RandomSeoURLWithLength(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijklmnopqrstuvwxyz",
	)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func ExtensionsByContentType(contentType string) (string, error) {
	exts, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}
	if len(exts) == 0 {
		return "", errors.New("contentType not found")
	}
	return exts[0], nil
}

func SlackHook(slackURL, channel, content string) error {
	// slackURL := "https://hooks.slack.com/services/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"
	go func() error {
		bodyRequest, err := json.Marshal(map[string]interface{}{
			"channel":  channel,
			"username": "tc-report",
			"text":     content,
			"icon_url": "http://www.hopabot.com/img/intro-carousel/f2.png",
		})
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(bodyRequest))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return errors.New(res.Status)
		}
		return nil
	}()
	return nil
}

type BinancePrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type BinancePrice24h struct {
	Symbol             string `json:"symbol"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
}

func GetListExternalPrice(tokenSymbol string) ([]BinancePrice, error) {
	binancePriceURL := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbols=%s", tokenSymbol)
	var prices []BinancePrice

	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return prices, nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	// value, err := strconv.ParseFloat(price.Price, 32)
	// if err != nil {
	// 	log.Println("getExternalPrice", tokenSymbol, err)
	// 	return 0, nil
	// }
	return prices, nil
}

func GetListExternalPrice24h(tokenSymbol string) ([]BinancePrice24h, error) {
	binancePriceURL := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbols=%s", tokenSymbol)
	var prices []BinancePrice24h

	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return prices, nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	return prices, nil
}

func GetBinancePrice24h(tokenSymbol string) (BinancePrice24h, error) {
	binancePriceURL := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbol=%s", tokenSymbol)
	var prices BinancePrice24h

	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return prices, nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	return prices, nil
}

func GetExternalPrice(tokenSymbol string) (string, error) {
	binanceAPI := "https://api.binance.com"
	binancePriceURL := fmt.Sprintf("%v/api/v3/ticker/price?symbol=%s", binanceAPI, tokenSymbol)
	var price BinancePrice

	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return "", nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &price)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	return price.Price, nil
}

func ReadCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GenerateMD5(v string) string {
	data := []byte(v)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func SlackGameAlert(content string) error {
	slackURL := "https://hooks.slack.com/services/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"
	channel := "game-alert"
	go func() error {
		bodyRequest, err := json.Marshal(map[string]interface{}{
			"channel":  channel,
			"username": "game-alert",
			"text":     content,
			"icon_url": "http://www.hopabot.com/img/intro-carousel/f2.png",
		})
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(bodyRequest))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return errors.New(res.Status)
		}
		return nil
	}()
	return nil
}

func RandomInStrings(values []string) string {
	if len(values) == 0 {
		panic("wrong values")
	}
	if len(values) == 1 {
		return values[0]
	}
	rand.Seed(time.Now().UnixNano())
	return values[rand.Intn(len(values))]
}

func GetTimeIndex(blockNumber, txIndex, logIndex uint) uint64 {
	blockNumber = blockNumber * 1e7
	txIndex = txIndex * 1e3
	timeIndex := uint64(blockNumber) + uint64(txIndex) + uint64(logIndex)
	return timeIndex
}

func GetFileExtension(fileName string) string {
	a := strings.Split(fileName, ".")
	ext := a[len(a)-1]
	return ext
}

type NakaChainPrice struct {
	Result *struct {
		ETH   string `json:"eth,omitempty"`
		BTC   string `json:"btc,omitempty"`
		EAI   string `json:"eai,omitempty"`
		BVM   string `json:"bvm,omitempty"`
		RUNIX string `json:"runix,omitempty"`
		NAKA  string `json:"naka,omitempty"`
	} `json:"result"`
}

func GetNakaChainMarketPrice() (*NakaChainPrice, error) {
	binancePriceURL := fmt.Sprintf("https://api.nakachain.xyz/api/coin-prices")
	var prices NakaChainPrice

	var jsonErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return &prices, nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		err = json.Unmarshal(body, &jsonErr)
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	return &prices, nil
}

func SliceToStrings(from, len int, getValueAtIndex func(index int) (string, error)) ([]string, error) {
	res := []string{}
	for i := from; i < from+len; i++ {
		val, err := getValueAtIndex(i)
		if err != nil {
			return nil, err
		}
		if val != "" {
			res = append(res, val)
		}
	}
	return res, nil
}

func IsValidEthereumAddress(address string) bool {
	// Regular expression to match Ethereum addresses
	re := regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	return re.MatchString(address)
}

func AddressToBinary(address string) ([]byte, error) {
	hexStr := address[2:]
	// Convert hex string to byte slice
	addressHash, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return addressHash, nil
}

func GetDurationUnit(unit uint8) string {
	switch unit {
	case 0:
		return "days"
	case 1:
		return "weeks"
	case 3:
		return "months"
	default:
		return ""
	}
}

// //////////
type CoincodexPrice struct {
	Symbol        string  `json:"symbol"`
	LastPriceUsd  float64 `json:"last_price_usd"`
	DisplaySymbol string  `json:"display_symbol"`
}

func GetCoincodexPriceByTime(monthStr string) ([]*CoincodexPrice, error) {
	binancePriceURL := fmt.Sprintf("https://coincodex.com/api/coincodex/get_historical_snapshot/%s", monthStr) + `%2000:00/0/200?t=57577472`
	var prices = struct {
		Coins []*CoincodexPrice `json:"coins"`
	}{}

	retryTimes := 0
retry:
	retryTimes++
	if retryTimes > 2 {
		return prices.Coins, nil
	}
	resp, err := http.Get(binancePriceURL)
	if err != nil {
		log.Println(err)
		goto retry
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &prices)
	if err != nil {
		if err != nil {
			log.Println(err)
			goto retry
		}
	}
	resp.Body.Close()
	return prices.Coins, nil
}

func FormatMoney(amount *big.Float) string {
	amountF, _ := amount.Float64()
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(amountF)
}

func CrawlUrl(urlToCrawl string) string {
	var (
		err      error
		content  string
		urlToGet *url.URL
		// links    []string
	)

	// Parse URL
	if urlToGet, err = url.Parse(urlToCrawl); err != nil {
		fmt.Println(err)
		return ""
	}

	// Retrieve content of URL
	if content, err = getUrlContent(urlToGet.String()); err != nil {
		fmt.Println(err)
		return ""
	}

	// Clean up HTML entities
	content = html.UnescapeString(content)
	fmt.Println(content)
	// if links, err = parseLinks(urlToGet, content); err != nil {
	// 	log.Println(err)
	// 	return ""
	// }

	// for _, link := range links {
	// 	defer CrawlUrl(link)
	// }
	return content
}

func getUrlContent(urlToGet string) (string, error) {
	var (
		err     error
		content []byte
		resp    *http.Response
	)

	// GET content of URL
	if resp, err = http.Get(urlToGet); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != 200 {
		return "", err
	}

	// Read the body of the HTTP response
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(content), err
}

func parseLinks(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err       error
		links     []string = make([]string, 0)
		matches   [][]string
		findLinks = regexp.MustCompile("<a.*?href=\"(.*?)\"")
	)

	// Retrieve all anchor tag URLs from string
	matches = findLinks.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var linkUrl *url.URL

		// Parse the anchr tag URL
		if linkUrl, err = url.Parse(val[1]); err != nil {
			return links, err
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if linkUrl.IsAbs() {
			links = append(links, linkUrl.String())
		} else {
			links = append(links, urlToGet.Scheme+"://"+urlToGet.Host+linkUrl.String())
		}
	}

	return links, err
}

func ContentHtmlByUrl(link string) string {
	textHtml := ""
	c := colly.NewCollector()
	c.OnHTML("html body", func(e *colly.HTMLElement) { // Body / content
		textHtml, _ = e.DOM.Html()
	})
	c.Visit(link)
	return textHtml
}

func ExtractDomainFromUrl(link string) (string, error) {
	parser, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(parser.Hostname(), "www."), nil
}

func RodContentHtmlByUrl(rawUrl string) string {
	path, has := launcher.LookPath()
	if !has {
		return ""
	}
	spew.Dump(path)

	u := launcher.New().Bin(path).
		Headless(true).
		// Proxy("http://65.49.14.150:3128").
		MustLaunch()
	browser := rod.New().ControlURL(u)

	page := browser.MustConnect().MustPage(rawUrl)
	page.MustWaitLoad()

	i := 0
	for i <= 5 {
		page.MustEval(`() => window.scrollTo(0, document.body.scrollHeight)`)
		time.Sleep(1 * time.Second)
		i += 1
	}

	page.MustWaitStable()
	page.MustEval(`() => document.querySelectorAll("[crossorigin]").forEach((el) => el.removeAttribute('crossorigin'))
	`)

	htmlStr, err := page.HTML()
	if err != nil {
		return ""
	}
	htmlStr, err = MinifyHTML(htmlStr)
	return htmlStr
}

func printReport(page *rod.Page) {
	_ = page.MustElement(".ds-dex-table-row")
	// for _, row := range el.MustParents("table").First().MustElements("tr:nth-child(n+2)") {
	// 	cells := row.MustElements("td")
	// 	key := cells[0].MustProperty("textContent")
	// 	if strings.HasPrefix(key.String(), "User Agent") {
	// 		fmt.Printf("\t\t%s: %t\n\n", key, !strings.Contains(cells[1].MustProperty("textContent").String(), "HeadlessChrome/"))
	// 	} else {
	// 		fmt.Printf("\t\t%s: %s\n\n", key, cells[1].MustProperty("textContent"))
	// 	}
	// }

	page.MustScreenshot("")
}

func RodContentHtmlByUrlV2(rawUrl string) string {
	path, has := launcher.LookPath()
	if !has {
		return ""
	}
	spew.Dump(path)

	u := launcher.New().Bin(path).Headless(true).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()

	fmt.Printf("js: %x\n\n", md5.Sum([]byte(stealth.JS)))
	page := stealth.MustPage(browser)
	page.MustNavigate(rawUrl)
	// page.MustWaitLoad()
	// time.Sleep(15 * time.Second)

	printReport(page)
	// i := 0
	// for i <= 5 {
	// 	page.MustEval(`() => window.scrollTo(0, document.body.scrollHeight)`)
	// 	time.Sleep(1 * time.Second)
	// 	i += 1
	// }

	// page.MustWaitStable()
	// page.MustEval(`() => document.querySelectorAll("[crossorigin]").forEach((el) => el.removeAttribute('crossorigin'))
	// `)

	htmlStr, err := page.HTML()
	if err != nil {
		return ""
	}
	htmlStr, err = MinifyHTML(htmlStr)
	return htmlStr
}

func MinifyHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	doc.Find("*").Each(func(index int, item *goquery.Selection) {
		var str string
		str, err = item.Html()
		str = strings.TrimSpace(str)
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\t", "")
		if err == nil {
			item.SetHtml(str)
		}
	})

	htmlStr, err := doc.Html()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(htmlStr), nil
}

func ExtractLinks(content string) (string, bool) {
	rx := xurls.Relaxed()
	extractLink := rx.FindString(content)
	reg := `((https?):\/\/)?(www.)?x\.com(\/@?(\w){1,15})\/status\/[0-9]{19}`
	isTwitterPost, _ := regexp.MatchString(reg, extractLink)
	return extractLink, isTwitterPost
}

func ExtractMapInfoFromOpenAI(content string) map[string]interface{} {
	resp := map[string]interface{}{}
	content = strings.ReplaceAll(content, "`", ``)
	content = strings.ReplaceAll(content, `\n`, ``)
	content = strings.ReplaceAll(content, `\`, ``)
	_ = json.Unmarshal([]byte(content), &resp)
	return resp
}

func ExpansionStringArray(arr []string) []string {
	strs := make([]string, len(arr))
	for i, expansion := range arr {
		strs[i] = string(expansion)
	}
	return strs
}

func RandomInt(from, to int) int {
	return from + rand.Intn(to-from+1)
}

func RandomFloat(from, to float64) float64 {
	return from + rand.Float64()*(to-from)
}

func ExtractEtherAddress(content string) string {
	arrStr := strings.Split(content, " ")
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	for _, item := range arrStr {
		if re.MatchString(strings.TrimSpace(item)) {
			return strings.TrimSpace(item)
		}
	}
	return ""
}

func ExtractEVMAddress(input string) string {
	re := regexp.MustCompile(`0x[0-9a-fA-F]{40}`)
	match := re.FindString(input)
	if match != "" {
		return match
	}
	return ""
}

func GetStringValueFromMap(mapInfo map[string]interface{}, key string) string {
	if v, ok := mapInfo[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}

func GetFrequencyFromMap(mapInfo map[string]interface{}) int {
	frequency := int(2)
	switch v := mapInfo["frequency"].(type) {
	case int:
		frequency = v
		frequency = frequency * 3600
	case float64:
		{
			frequencyF := v
			frequency = int(frequencyF * 3600)
		}
	case string:
		frequency, _ = strconv.Atoi(v)
		frequency = frequency * 3600
	default:
		frequency = int(2)
		frequency = frequency * 3600
	}
	return frequency
}

func GetTokenIDFromMap(mapInfo map[string]interface{}) uint64 {
	tokenId := int(0)
	switch v := mapInfo["tokenId"].(type) {
	case int:
		tokenId = v
	case float64:
		{
			frequencyF := v
			tokenId = int(frequencyF)
		}
	case string:
		tokenId, _ = strconv.Atoi(v)
	}
	return uint64(tokenId)
}

func splitStringArray(arr []string, size int) [][]string {
	var result [][]string
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}

func GenerateTokenSymbol(tokenName string) string {
	// Remove special characters and spaces
	tokenSymbol := strings.ToLower(tokenName)
	tokenSymbol = strings.ReplaceAll(tokenSymbol, " ", "")
	tokenSymbol = strings.ReplaceAll(tokenSymbol, "-", "")
	tokenSymbol = strings.ReplaceAll(tokenSymbol, "_", "")

	// Remove any non-alphanumeric characters
	reg := regexp.MustCompile("[^a-z0-9]")
	tokenSymbol = reg.ReplaceAllString(tokenSymbol, "")

	// Take first 5 characters
	if len(tokenSymbol) > 5 {
		tokenSymbol = tokenSymbol[:5]
	}

	// If empty after cleaning, return default
	if tokenSymbol == "" {
		tokenSymbol = RandomStringWithLength(5)
	}
	fmt.Println(tokenSymbol)
	return strings.ToUpper(tokenSymbol)
}
