package serializers

type PumpOrderResponse struct {
	Status  int           `json:"status"`
	Data    PumpOrderData `json:"data"`
	Message string        `json:"message"`
}

type PumpOrderData struct {
	Orders []PumpOrder `json:"orders"`
	Total  int         `json:"total"`
}

type PumpOrder struct {
	ID           string   `json:"id"`
	WalletID     string   `json:"wallet_id"`
	Type         int      `json:"type"`
	TokenAddress string   `json:"token_address"`
	Destination  string   `json:"destination"`
	AmountIn     string   `json:"amount_in"`
	AmountOut    string   `json:"amount_out"`
	Txs          []string `json:"txs"`
	Status       int      `json:"status"`
	Error        string   `json:"error"`
	ErrorRaw     string   `json:"error_raw"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type PumpBalanceResponse struct {
	Status  int             `json:"status"`
	Data    PumpBalanceData `json:"data"`
	Message string          `json:"message"`
}

type PumpBalanceData struct {
	Amount string      `json:"amount"`
	Tokens []PumpToken `json:"tokens"`
}

type PumpToken struct {
	Mint     string `json:"mint"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Amount   string `json:"amount"`
}
