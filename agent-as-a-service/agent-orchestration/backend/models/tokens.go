package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type Erc20Holder struct {
	gorm.Model
	NetworkID           uint64           `gorm:"unique_index:token_holder_main_uidx"`
	ContractAddress     string           `gorm:"unique_index:token_holder_main_uidx"`
	Address             string           `gorm:"unique_index:token_holder_main_uidx"`
	Balance             string           `gorm:"default:0"`
	UserName            string           `gorm:"-"`
	ImageURL            string           `gorm:"-"`
	MemeName            string           `gorm:"-"`
	MemeTicker          string           `gorm:"-"`
	MemeImage           string           `gorm:"-"`
	MemePrice           numeric.BigFloat `gorm:"-"`
	MemePriceUsd        numeric.BigFloat `gorm:"-"`
	MemeBaseTokenSymbol string           `gorm:"-"`
}

type Erc721Holder struct {
	gorm.Model
	NetworkID       uint64 `gorm:"unique_index:nft_holders_main_idx"`
	ContractAddress string `gorm:"unique_index:nft_holders_main_idx"`
	TokenID         uint   `gorm:"unique_index:nft_holders_main_idx"`
	OwnerAddress    string `gorm:"index"`
}
type Erc1155Holder struct {
	gorm.Model
	NetworkID       uint64 `gorm:"unique_index:nft_holders_main_idx"`
	ContractAddress string `gorm:"unique_index:nft_holders_main_idx"`
	TokenID         uint   `gorm:"unique_index:nft_holders_main_idx"`
	Address         string `gorm:"unique_index:nft_holders_main_idx"`
	Balance         string `gorm:"default:0"`
}

type TrendingToken struct {
	gorm.Model
	ChainId      string `gorm:"unique_index:trending_token_main_idx"`
	TokenAddress string `gorm:"unique_index:trending_token_main_idx"`
	Name         string
	Symbol       string `gorm:"index"`
	Decimals     int
	Logo         string
	UsdPrice     float64
	MarketCap    float64 `gorm:"index"`
	LiquidityUsd float64
	Holders      int `gorm:"index"`
	MintAt       *time.Time

	// Time-based metrics (1h, 4h, 12h, 24h)
	PriceChange1h  float64
	PriceChange4h  float64
	PriceChange12h float64
	PriceChange24h float64

	Volume1h  float64
	Volume4h  float64
	Volume12h float64
	Volume24h float64 `gorm:"index"`

	Transactions1h  int
	Transactions4h  int
	Transactions12h int
	Transactions24h int `gorm:"index"`

	BuyTransactions1h  int
	BuyTransactions4h  int
	BuyTransactions12h int
	BuyTransactions24h int

	SellTransactions1h  int
	SellTransactions4h  int
	SellTransactions12h int
	SellTransactions24h int

	Buyers1h  int
	Buyers4h  int
	Buyers12h int
	Buyers24h int `gorm:"index"`

	Sellers1h  int
	Sellers4h  int
	Sellers12h int
	Sellers24h int `gorm:"index"`
}
