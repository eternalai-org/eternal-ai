package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type MemeStatus string
type TransferStatus string
type NotiType string
type ChartType string

type MemeCatStatus string

const (
	MemeStatusPending          MemeStatus = "pending"
	MemeStatusNew              MemeStatus = "new"
	MemeStatusCreated          MemeStatus = "created"
	MemeStatusAddPoolLevel0    MemeStatus = "add_pool_0"
	MemeStatusAddPoolLevel1    MemeStatus = "add_pool_1"
	MemeStatusReachedMC        MemeStatus = "reached_mc"
	MemeStatusRemovePoolLelve1 MemeStatus = "remove_pool_1"
	MemeStatusAddPoolLevel2    MemeStatus = "add_pool_2"
	MemeStatusAddPoolExternal  MemeStatus = "add_pool_external"

	MemeCatStatusProcessing MemeCatStatus = "processing"
	MemeCatStatusDone       MemeCatStatus = "done"

	ChartTypeHour1 ChartType = "1h"
	ChartTypeHour4 ChartType = "4h"
	ChartTypeDay   ChartType = "1d"
	ChartTypeMin30 ChartType = "30min"
	ChartTypeMin5  ChartType = "5min"

	TransferStatusPending TransferStatus = "pending"
	TransferStatusDone    TransferStatus = "done"

	NotiTypeNewMeme     NotiType = "new_meme"
	NotiTypePrice       NotiType = "price_pump"
	NotiTypeReachMC     NotiType = "reached_mc"
	NotiTypeTradeOnNaka NotiType = "trade_on_naka"
	NotiTypeNewFollower NotiType = "new_follower"
)

type Meme struct {
	gorm.Model
	NetworkID             uint64
	OwnerAddress          string `gorm:"index"`
	OwnerID               uint   `gorm:"index"`
	Owner                 *User
	AgentInfoID           uint `gorm:"unique_index"`
	AgentInfo             *AgentInfo
	TokenAddress          string `gorm:"index"`
	TokenId               string `gorm:"index"`
	FactoryAddress        string
	TotalSuply            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Decimals              uint64
	Name                  string
	Ticker                string
	Description           string `gorm:"type:longtext collate utf8mb4_unicode_ci"`
	Err                   string `gorm:"type:text collate utf8mb4_unicode_ci"`
	Image                 string
	Twitter               string
	Telegram              string
	Website               string
	Status                MemeStatus `gorm:"index"`
	Pool                  string     `gorm:"index"`
	UniswapPool           string     `gorm:"index"`
	Token0Address         string
	Token1Address         string
	Reserve0              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Reserve1              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Supply                numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Price                 numeric.BigFloat `gorm:"index;type:decimal(36,18);default:0"`
	PriceUsd              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	PriceLast24h          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	VolumeLast24h         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalVolume           numeric.BigFloat `gorm:"index;type:decimal(36,18);default:0"`
	TimeIndex             uint64
	BaseTokenIndex        int
	ZeroForOne            bool
	BaseTokenSymbol       string           `gorm:"default:'ETH'"`
	ReplyCount            uint64           `gorm:"index;default:0"`
	LastReply             *time.Time       `gorm:"index"`
	PositionID            int64            `gorm:"default:0"`
	UniswapPositionID     int64            `gorm:"default:0"`
	Liquidity             numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	PositionLiquidity     numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TickLower             int64
	TickUpper             int64
	Tick                  int64
	AddPool1TxHash        string
	RemovePool1TxHash     string
	AddPool2TxHash        string
	AddPool2At            *time.Time `gorm:"index"`
	BurnPool2TxHash       string
	BurnPool2At           *time.Time `gorm:"index"`
	PoolFee               uint
	Weight                int `gorm:"index;default:0"`
	Shared                int
	ReqSyncAt             *time.Time
	SyncAt                *time.Time
	Fee                   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	ExternalTradeUrl      string
	NumRetries            int              `gorm:"default:0"`
	NotGraduated          bool             `gorm:"default:0"`
	Fees0CollectedBalance numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Fees1CollectedBalance numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	FeesCollectedUpdated  uint64           `gorm:"default:0"`
	PoolSwappedUpdated    uint64           `gorm:"default:0"`

	//
	Percent                float64 `gorm:"-"`
	StoreImageOnChain      bool
	MarketCap              numeric.BigFloat  `gorm:"-"`
	TotalBalance           numeric.BigFloat  `gorm:"-"`
	Holders                int               `gorm:"-"`
	LatestAgentTwitterPost *AgentTwitterPost `gorm:"-"`
}

type MemeTradeHistory struct {
	gorm.Model
	NetworkID        uint64
	TxHash           string `gorm:"index"`
	ContractAddress  string `gorm:"index"`
	EventId          string `gorm:"unique_index:pump_trade_histories_main_idx"`
	TxAt             time.Time
	RecipientAddress string `gorm:"index"`
	RecipientUserID  uint   `gorm:"index"`
	RecipientUser    *User
	Amount0          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Amount1          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SqrtPriceX96     string
	Liquidity        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Tick             int64
	Price            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	MemeTokenAddress string           `gorm:"index"`
	TokenId          string           `gorm:"index"`
	MemeID           uint             `gorm:"index"`
	Meme             *Meme
	TokenInAddress   string
	TokenIn          *Meme            `gorm:"foreignKey:contract_address_check;AssociationForeignKey:token_in_address"`
	AmountIn         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenOutAddress  string           `gorm:"index"`
	TokenOut         *Meme            `gorm:"foreignKey:contract_address_check;AssociationForeignKey:token_out_address"`
	AmountOut        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	BaseTokenSymbol  string           `gorm:"default:'BTC'"`
	BaseTokenPrice   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	BaseAmount       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenAmount      numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsBuy            bool
}

type MemeThreads struct {
	gorm.Model
	UserID         uint `gorm:"index"`
	User           *User
	MemeID         uint `gorm:"index"`
	Meme           *Meme
	Text           string `gorm:"type:longtext collate utf8mb4_unicode_ci"`
	ImageUrl       string
	Likes          int64
	ParentThreadID uint `gorm:"index"`
	Hidden         bool `gorm:"default:0"`

	Liked bool `gorm:"-"`
}

type MemeWhiteListAddress struct {
	gorm.Model
	Address string `gorm:"index"`
}

type MemeThreadLike struct {
	gorm.Model
	UserID   uint `gorm:"unique_index:meme_thread_like_main_idx"`
	ThreadID uint `gorm:"unique_index:meme_thread_like_main_idx"`
}

type MemeFollowers struct {
	gorm.Model
	UserID       uint `gorm:"unique_index:meme_followers_main_idx"`
	User         *User
	FollowUserID uint `gorm:"unique_index:meme_followers_main_idx"`
	FollowUser   *User
}

type MemeTokenHolder struct {
	gorm.Model
	ContractAddress string           `gorm:"unique_index:token_holder_main_uidx"`
	Address         string           `gorm:"unique_index:token_holder_main_uidx"`
	Balance         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`

	UserName            string           `gorm:"-"`
	ImageURL            string           `gorm:"-"`
	MemeName            string           `gorm:"-"`
	MemeTicker          string           `gorm:"-"`
	MemeImage           string           `gorm:"-"`
	MemePrice           numeric.BigFloat `gorm:"-"`
	MemePriceUsd        numeric.BigFloat `gorm:"-"`
	MemeBaseTokenSymbol string           `gorm:"-"`
}

type MemeNotification struct {
	gorm.Model
	EventId    string `gorm:"unique_index:meme_notification_main_idx"`
	UserID     uint   `gorm:"index;default:0"`
	User       *User
	MemeID     uint `gorm:"index;default:0"`
	Meme       *Meme
	FollowerID uint `gorm:"index;default:0"`
	Follower   *User
	NotiType   NotiType
	Value      string `gorm:"type:longtext"`
	Seen       bool   `gorm:"index;default:0"`
}

type MemeNotificationSeen struct {
	gorm.Model
	NotificationID uint `gorm:"unique_index:meme_notification_seen_main_idx"`
	UserID         uint `gorm:"unique_index:meme_notification_seen_main_idx"`
	Seen           bool `gorm:"index;default:0"`
}

type ChartData struct {
	PairID      string           `json:"pair_id"`
	AvgPrice    numeric.BigFloat `json:"avg_price"`
	MinPrice    numeric.BigFloat `json:"min_price"`
	MaxPrice    numeric.BigFloat `json:"max_price"`
	BaseVolume  numeric.BigFloat `json:"base_volume"`
	TokenVolume numeric.BigFloat `json:"token_volume"`
	OpenPrice   numeric.BigFloat `json:"open_price"`
	ClosePrice  numeric.BigFloat `json:"close_price"`
	ChartTime   *time.Time       `json:"chart_time"`
}

type MemeSeen struct {
	gorm.Model
	UserID      uint       `gorm:"unique_index:meme_seen_main_idx"`
	UserAddress string     `gorm:"index"`
	MemeID      uint       `gorm:"unique_index:meme_seen_main_idx"`
	MemeAddress string     `gorm:"index"`
	SeenTime    *time.Time `gorm:"index"`
}

type Cat20TransferTransaction struct {
	gorm.Model
	TokenId         string
	ReceiverAddress string
	SendAmount      numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SenderAddress   string           `gorm:"index"`
	TxHash          string           `gorm:"index"`
	TxAt            *time.Time
	Error           string
}

type TokenTransfer struct {
	gorm.Model
	NetworkID       uint64
	EventId         string     `gorm:"unique_index"`
	ContractAddress string     `gorm:"index"`
	TransactionAt   *time.Time `gorm:"index"`
	TxHash          string
	From            string           `gorm:"index"`
	To              string           `gorm:"index"`
	Value           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`

	Name            string `gorm:"-"`
	Ticker          string `gorm:"-"`
	Image           string `gorm:"-"`
	TwitterID       string `gorm:"-"`
	TwitterName     string `gorm:"-"`
	TwitterUsername string `gorm:"-"`
	TwitterAvatar   string `gorm:"-"`
	UserTwitterID   string `gorm:"-"`
	UserName        string `gorm:"-"`
	ImageURL        string `gorm:"-"`
}

type Bot struct {
	gorm.Model
	NetworkID uint64
	Address   string
}

type BotOrderType string
type BotOrderStatus string

const (
	BotOrderTypeBuyAfterCreated BotOrderType = "buy_after_created"
	BotOrderTypeBuy             BotOrderType = "buy"
	BotOrderTypeSell            BotOrderType = "sell"

	BotOrderStatusNew        BotOrderStatus = "new"
	BotOrderStatusDepositing BotOrderStatus = "depositing"
	BotOrderStatusDone       BotOrderStatus = "done"
	BotOrderStatusPoolled    BotOrderStatus = "poolled"
	BotOrderStatusError      BotOrderStatus = "error"
)

type BotOrder struct {
	gorm.Model
	NetworkID   uint64
	Address     string
	Type        BotOrderType
	ScheduledAt *time.Time
	MemeID      uint
	TokenIn     string
	TokenOut    string
	AmountIn    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	AmountOut   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Price       numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	ParentPrice numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	UniFee      uint             `gorm:"default:0"`
	DepositHash string
	TxHash      string
	PoolledHash string
	Status      BotOrderStatus
	ParentID    uint
	Error       string
}

type MemeFeesCollected struct {
	gorm.Model
	EventId string `gorm:"unique_index"`
	MemeID  uint   `gorm:"index"`
	TxHash  string
	Amount0 numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Amount1 numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}
