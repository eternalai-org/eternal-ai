package serializers

import (
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type MemeReq struct {
	Name            string           `json:"name"`
	Ticker          string           `json:"ticker"`
	Description     string           `json:"description"`
	Image           string           `json:"image"`
	Twitter         string           `json:"twitter"`
	Telegram        string           `json:"telegram"`
	Website         string           `json:"website"`
	TxHash          string           `json:"tx_hash"`
	SystemPrompt    string           `json:"system_prompt"`
	OnchainImage    bool             `json:"onchain_image"`
	TotalSuply      numeric.BigFloat `json:"total_suply"`
	Decimals        uint64           `json:"decimals"`
	Cat20Token      string           `json:"cat20_token"`
	BtcPair         bool             `json:"bt_pair"`
	FbPair          bool             `json:"fb_pair"`
	AgentInfoID     uint             `json:"agent_info_id"`
	BaseTokenSymbol string           `json:"base_token_symbol"`
	NotGraduated    bool             `json:"not_graduated"`
}

type MemeResp struct {
	ID                uint                  `json:"id"`
	NetworkID         uint64                `json:"network_id"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	OwnerAddress      string                `json:"owner_address"`
	Owner             *UserResp             `json:"owner"`
	TokenAddress      string                `json:"token_address"`
	TokenId           string                `json:"token_id"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Ticker            string                `json:"ticker"`
	Image             string                `json:"image"`
	Twitter           string                `json:"twitter"`
	Telegram          string                `json:"telegram"`
	Website           string                `json:"website"`
	TxHash            string                `json:"tx_hash"`
	Status            string                `json:"status"`
	ReplyCount        uint64                `json:"reply_count"`
	LastReply         *time.Time            `json:"last_reply"`
	Pool              string                `json:"pool"`
	UniswapPool       string                `json:"uniswap_pool"`
	Supply            numeric.BigFloat      `json:"supply"`
	Price             numeric.BigFloat      `json:"price"`
	PriceUsd          numeric.BigFloat      `json:"price_usd"`
	PriceLast24h      numeric.BigFloat      `json:"price_last24h"`
	VolumeLast24h     numeric.BigFloat      `json:"volume_last24h"`
	TotalVolume       numeric.BigFloat      `json:"total_volume"`
	BaseTokenSymbol   string                `json:"base_token_symbol"`
	Percent           float64               `json:"percent"`
	Decimals          uint64                `json:"decimals"`
	PoolFee           uint                  `json:"pool_fee"`
	MarketCap         numeric.BigFloat      `json:"market_cap"`
	TotalBalance      numeric.BigFloat      `json:"total_balance"`
	SystemPrompt      string                `json:"system_prompt"`
	Holders           int                   `json:"holders"`
	Shared            int                   `json:"shared"`
	AgentInfo         *AgentInfoResp        `json:"agent_info"`
	LatestTwitterPost *AgentTwitterPostResp `json:"latest_twitter_post"`
	DexUrl            string                `json:"dex_url"`
	TradeUrl          string                `json:"trade_url"`
	DexId             string                `json:"dex_id"`
}

func NewMemeResp(m *models.Meme) *MemeResp {
	if m == nil {
		return nil
	}
	volumeUSD := models.MulBigFloats(&m.PriceUsd.Float, &m.VolumeLast24h.Float)
	resp := &MemeResp{
		ID:              m.ID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		NetworkID:       m.NetworkID,
		OwnerAddress:    m.OwnerAddress,
		TokenAddress:    m.TokenAddress,
		TokenId:         m.TokenId,
		Description:     m.Description,
		Name:            m.Name,
		Ticker:          m.Ticker,
		Image:           m.Image,
		Twitter:         m.Twitter,
		Telegram:        m.Telegram,
		Website:         m.Website,
		Status:          string(m.Status),
		Owner:           NewUserResp(m.Owner),
		Pool:            m.Pool,
		UniswapPool:     m.UniswapPool,
		Supply:          m.Supply,
		Price:           m.Price,
		PriceUsd:        m.PriceUsd,
		PriceLast24h:    m.PriceLast24h,
		VolumeLast24h:   numeric.BigFloat{*volumeUSD},
		TotalVolume:     m.TotalVolume,
		BaseTokenSymbol: m.BaseTokenSymbol,
		ReplyCount:      m.ReplyCount,
		LastReply:       m.LastReply,
		Percent:         m.Percent,
		Decimals:        m.Decimals,
		PoolFee:         m.PoolFee,
		MarketCap:       m.MarketCap,
		TotalBalance:    m.TotalBalance,
		Holders:         m.Holders,
		Shared:          m.Shared,
	}

	if m.AgentInfo != nil {
		resp.AgentInfo = NewAgentInfoResp(m.AgentInfo)
	}
	if m.LatestAgentTwitterPost != nil {
		resp.LatestTwitterPost = NewAgentTwitterPostResp(m.LatestAgentTwitterPost)
	}
	return resp
}

func NewMemeRespWithToken(m *models.Meme) *MemeResp {
	if m == nil {
		return nil
	}
	volumeUSD := models.MulBigFloats(&m.PriceUsd.Float, &m.VolumeLast24h.Float)
	resp := &MemeResp{
		ID:              m.ID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		OwnerAddress:    m.OwnerAddress,
		TokenAddress:    m.TokenAddress,
		Description:     m.Description,
		Name:            m.Name,
		Ticker:          m.Ticker,
		Image:           m.Image,
		Twitter:         m.Twitter,
		Telegram:        m.Telegram,
		Website:         m.Website,
		Status:          string(m.Status),
		Owner:           NewUserResp(m.Owner),
		Pool:            m.Pool,
		UniswapPool:     m.UniswapPool,
		Supply:          m.Supply,
		Price:           m.Price,
		PriceUsd:        m.PriceUsd,
		PriceLast24h:    m.PriceLast24h,
		VolumeLast24h:   numeric.BigFloat{*volumeUSD},
		TotalVolume:     m.TotalVolume,
		BaseTokenSymbol: m.BaseTokenSymbol,
		ReplyCount:      m.ReplyCount,
		LastReply:       m.LastReply,
		Percent:         m.Percent,
		Decimals:        m.Decimals,
		PoolFee:         m.PoolFee,
		MarketCap:       m.MarketCap,
		TotalBalance:    m.TotalBalance,
		Holders:         m.Holders,
		Shared:          m.Shared,
		NetworkID:       m.NetworkID,
	}

	resp.DexUrl = models.GetDexUrl(m.NetworkID, m.TokenAddress)
	resp.TradeUrl = models.GetTradeUrl(m.NetworkID, m.TokenAddress, "")
	return resp
}

func NewMemeFromTokenInfoResp(m *models.AgentTokenInfo, agentInfo *models.AgentInfo) *MemeResp {
	if m == nil {
		return nil
	}
	resp := &MemeResp{
		ID:              m.ID,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		TokenAddress:    agentInfo.TokenAddress,
		Name:            agentInfo.TokenName,
		Ticker:          agentInfo.TokenSymbol,
		Pool:            m.PoolAddress,
		UniswapPool:     m.PoolAddress,
		Price:           m.Price,
		PriceUsd:        m.PriceUsd,
		PriceLast24h:    m.PriceLast24h,
		VolumeLast24h:   m.VolumeLast24h,
		TotalVolume:     m.TotalVolume,
		BaseTokenSymbol: m.BaseTokenSymbol,
		Percent:         m.PriceChange,
		MarketCap:       numeric.NewBigFloatFromFloat(big.NewFloat(m.UsdMarketCap)),
		DexId:           m.DexId,
		DexUrl:          m.DexUrl,
		Status:          string(models.MemeStatusAddPoolLevel2),
		NetworkID:       agentInfo.TokenNetworkID,
	}
	resp.DexUrl = models.GetDexUrl(agentInfo.TokenNetworkID, agentInfo.TokenAddress)
	resp.TradeUrl = models.GetTradeUrl(agentInfo.TokenNetworkID, agentInfo.TokenAddress, m.DexId)
	return resp
}

func NewMemeRespArray(arr []*models.Meme) []*MemeResp {
	resps := []*MemeResp{}
	for _, m := range arr {
		resps = append(resps, NewMemeResp(m))
	}
	return resps
}

type MemeTradeHistoryResp struct {
	TxHash           string           `json:"tx_hash"`
	TxAt             time.Time        `json:"tx_at"`
	RecipientAddress string           `json:"recipient_address"`
	RecipientUser    *UserResp        `json:"recipient_user"`
	Price            numeric.BigFloat `json:"price"`
	MemeTokenAddress string           `json:"meme_token_address"`
	MemeID           uint             `json:"meme_id"`
	Meme             *MemeResp        `json:"meme"`
	AmountIn         numeric.BigFloat `json:"amount_in"`
	AmountOut        numeric.BigFloat `json:"amount_out"`
	BaseTokenSymbol  string           `json:"base_token_symbol"`
	BaseTokenPrice   numeric.BigFloat `json:"base_token_price"`
	BaseAmount       numeric.BigFloat `json:"base_amount"`
	TokenAmount      numeric.BigFloat `json:"token_amount"`
	IsBuy            bool             `json:"is_buy"`
}

func NewTradeHistoryResp(m *models.MemeTradeHistory) *MemeTradeHistoryResp {
	if m == nil {
		return nil
	}
	resp := &MemeTradeHistoryResp{
		TxAt:             m.TxAt,
		TxHash:           m.TxHash,
		BaseTokenSymbol:  m.BaseTokenSymbol,
		BaseTokenPrice:   m.BaseTokenPrice,
		RecipientAddress: m.RecipientAddress,
		RecipientUser:    NewUserResp(m.RecipientUser),
		Price:            m.Price,
		MemeTokenAddress: m.MemeTokenAddress,
		MemeID:           m.MemeID,
		Meme:             NewMemeResp(m.Meme),
		AmountIn:         m.AmountIn,
		AmountOut:        m.AmountOut,
		BaseAmount:       m.BaseAmount,
		TokenAmount:      m.TokenAmount,
		IsBuy:            m.IsBuy,
	}
	return resp
}

func NewTradeHistoryRespArry(arr []*models.MemeTradeHistory) []*MemeTradeHistoryResp {
	resps := []*MemeTradeHistoryResp{}
	for _, m := range arr {
		resps = append(resps, NewTradeHistoryResp(m))
	}
	return resps
}

// //
type WithdrawCrossChainReq struct {
	TxHash string `json:"tx_hash"`
}

type MemeThreadReq struct {
	UserAddress    string `json:"user_address"`
	UserID         uint   `json:"user_id"`
	MemeID         uint   `json:"meme_id"`
	ThreadID       uint   `json:"thread_id"`
	Text           string `json:"text"`
	ImageUrl       string `json:"image_url"`
	ParentThreadID uint   `json:"parent_thread_id"`
}

type MemeThreadResp struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserID         uint      `json:"user_id"`
	User           *UserResp `json:"user"`
	MemeID         uint      `json:"meme_id"`
	Meme           *MemeResp `json:"meme"`
	Text           string    `json:"text"`
	ImageUrl       string    `json:"image_url"`
	Likes          int64     `json:"likes"`
	ParentThreadID uint      `json:"parent_thread_id"`
	Liked          bool      `json:"liked"`
}

func NewMemeThreadResp(m *models.MemeThreads) *MemeThreadResp {
	if m == nil {
		return nil
	}
	resp := &MemeThreadResp{
		MemeID:         m.MemeID,
		Meme:           NewMemeResp(m.Meme),
		ID:             m.ID,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		UserID:         m.UserID,
		User:           NewUserResp(m.User),
		Text:           m.Text,
		ImageUrl:       m.ImageUrl,
		Likes:          m.Likes,
		ParentThreadID: m.ParentThreadID,
		Liked:          m.Liked,
	}
	return resp
}

func NewMemeThreadRespArry(arr []*models.MemeThreads) []*MemeThreadResp {
	resps := []*MemeThreadResp{}
	for _, m := range arr {
		resps = append(resps, NewMemeThreadResp(m))
	}
	return resps
}

// ////
type MemeFollowersResp struct {
	UserID       uint      `json:"user_id"`
	User         *UserResp `json:"user"`
	FollowUserID uint      `json:"follow_user_id"`
	FollowUser   *UserResp `json:"follow_user"`
}

func NewMemeFollowersResp(m *models.MemeFollowers) *MemeFollowersResp {
	if m == nil {
		return nil
	}
	resp := &MemeFollowersResp{
		UserID:       m.UserID,
		User:         NewUserResp(m.User),
		FollowUserID: m.FollowUserID,
		FollowUser:   NewUserResp(m.FollowUser),
	}
	return resp
}

func NewMemeFollowersRespArray(arr []*models.MemeFollowers) []*MemeFollowersResp {
	resps := []*MemeFollowersResp{}
	for _, m := range arr {
		resps = append(resps, NewMemeFollowersResp(m))
	}
	return resps
}

// ///
type MemeNotificationResp struct {
	ID         uint            `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	MemeID     uint            `json:"meme_id"`
	Meme       *MemeResp       `json:"meme"`
	FollowerID uint            `json:"follower_id"`
	Follower   *UserResp       `json:"follower"`
	NotiType   models.NotiType `json:"noti_type"`
	Value      string          `json:"value"`
	Seen       bool            `json:"seen"`
}

func NewMemeNotificationResp(m *models.MemeNotification) *MemeNotificationResp {
	if m == nil {
		return nil
	}
	resp := &MemeNotificationResp{
		ID:         m.ID,
		CreatedAt:  m.CreatedAt,
		MemeID:     m.MemeID,
		Meme:       NewMemeResp(m.Meme),
		FollowerID: m.FollowerID,
		Follower:   NewUserResp(m.Follower),
		NotiType:   m.NotiType,
		Value:      m.Value,
		Seen:       m.Seen,
	}
	return resp
}

func NewMemeNotificationRespArry(arr []*models.MemeNotification) []*MemeNotificationResp {
	resps := []*MemeNotificationResp{}
	for _, m := range arr {
		resps = append(resps, NewMemeNotificationResp(m))
	}
	return resps
}

type MemeBurnHistoryResp struct {
	TxHash          string           `json:"tx_hash"`
	TxAt            time.Time        `json:"tx_at"`
	Value           numeric.BigFloat `json:"value"`
	ContractAddress string           `json:"contract_address"`
	Name            string           `json:"token_name"`
	Ticker          string           `json:"token_ticker"`
	Image           string           `json:"token_image"`
	TwitterName     string           `json:"twitter_name"`
	TwitterUsername string           `json:"twitter_username"`
	TwitterAvatar   string           `json:"twitter_avatar"`
	UserTwitterID   string           `json:"user_twitter_id"`
	UserAddress     string           `json:"user_address"`
	UserName        string           `json:"user_name"`
	ImageURL        string           `json:"user_image_url"`
}

func NewMemeBurnHistoryResp(m *models.TokenTransfer) *MemeBurnHistoryResp {
	if m == nil {
		return nil
	}
	resp := &MemeBurnHistoryResp{
		TxAt:            *m.TransactionAt,
		TxHash:          m.TxHash,
		Value:           m.Value,
		ContractAddress: m.ContractAddress,
		Name:            m.Name,
		Ticker:          m.Ticker,
		Image:           m.Image,
		TwitterName:     m.TwitterName,
		TwitterUsername: m.TwitterUsername,
		TwitterAvatar:   m.TwitterAvatar,
		UserTwitterID:   m.UserTwitterID,
		UserAddress:     m.From,
		UserName:        m.UserName,
		ImageURL:        m.ImageURL,
	}
	return resp
}

func NewMemeBurnHistoryRespArry(arr []*models.TokenTransfer) []*MemeBurnHistoryResp {
	resps := []*MemeBurnHistoryResp{}
	for _, m := range arr {
		resps = append(resps, NewMemeBurnHistoryResp(m))
	}
	return resps
}

type Cat20TransferTransactionResp struct {
	TokenId         string           `json:"token_id"`
	ReceiverAddress string           `json:"receiver_address"`
	SendAmount      numeric.BigFloat `json:"send_amount"`
	SenderAddress   string           `json:"sender_address"`
	TxHash          string           `json:"tx_hash"`
	TxAt            *time.Time       `json:"tx_at"`
	Error           string           `json:"error"`
}

func NewCat20TransferTransactionResp(m *models.Cat20TransferTransaction) *Cat20TransferTransactionResp {
	if m == nil {
		return nil
	}
	resp := &Cat20TransferTransactionResp{
		TokenId:         m.TokenId,
		ReceiverAddress: m.ReceiverAddress,
		SendAmount:      m.SendAmount,
		SenderAddress:   m.SenderAddress,
		TxHash:          m.TxHash,
		TxAt:            m.TxAt,
		Error:           m.Error,
	}
	return resp
}

type TransferCatReq struct {
	TokenId         string `json:"token_id"`
	ReceiverAddress string `json:"receiver_address"`
	SendAmount      string `json:"send_amount"`
	SenderAddress   string `json:"sender_address"`
	WifPrivateKey   string `json:"wif_private_key"`
}

type UserVideoResp struct {
	ID                 uint                  `json:"id"`
	CreatedAt          time.Time             `json:"created_at"`
	TokenName          string                `json:"token_name"`
	TokenSymbol        string                `json:"token_symbol"`
	TokenAddress       string                `json:"token_address"`
	TokenImageUrl      string                `json:"token_image_url"`
	TokenDesc          string                `json:"token_desc"`
	RequestorAddress   string                `json:"requestor_address"`
	OwnerTwitterID     string                `json:"owner_twitter_id"`
	AgentTwitterPostID uint                  `json:"agent_twitter_post_id"`
	AgentTwitterPost   *AgentTwitterPostResp `json:"agent_twitter_post"`
}

func NewUserVideoResp(m *models.ClankerVideoToken) *UserVideoResp {
	if m == nil {
		return nil
	}
	return &UserVideoResp{
		ID:                 m.ID,
		CreatedAt:          m.CreatedAt,
		TokenName:          m.TokenName,
		TokenSymbol:        m.TokenSymbol,
		TokenAddress:       m.TokenAddress,
		TokenImageUrl:      m.TokenImageUrl,
		TokenDesc:          m.TokenDesc,
		RequestorAddress:   m.RequestorAddress,
		OwnerTwitterID:     m.OwnerTwitterID,
		AgentTwitterPostID: m.AgentTwitterPostID,
		AgentTwitterPost:   NewAgentTwitterPostResp(m.AgentTwitterPost),
	}
}

func NewUserVideoRespArray(arr []*models.ClankerVideoToken) []*UserVideoResp {
	resps := []*UserVideoResp{}
	for _, r := range arr {
		resps = append(resps, NewUserVideoResp(r))
	}
	return resps
}

type TrendingTokenResp struct {
	ChainId        string     `json:"chain_id"`
	TokenAddress   string     `json:"token_address"`
	Name           string     `json:"name"`
	Symbol         string     `json:"symbol"`
	Decimals       int        `json:"decimals"`
	Logo           string     `json:"logo"`
	UsdPrice       float64    `json:"usd_price"`
	MarketCap      float64    `json:"market_cap"`
	LiquidityUsd   float64    `json:"liquidity_usd"`
	Holders        int        `json:"holders"`
	MintAt         *time.Time `json:"mint_at"`
	Volume1h       float64    `json:"volume1h"`
	Transactions1h int        `json:"transactions1h"`
	Buyers1h       int        `json:"buyers1h"`
}

func NewTrendingTokenResp(m *models.TrendingToken) *TrendingTokenResp {
	if m == nil {
		return nil
	}
	return &TrendingTokenResp{
		TokenAddress:   m.TokenAddress,
		Name:           m.Name,
		Symbol:         m.Symbol,
		Decimals:       m.Decimals,
		Logo:           m.Logo,
		UsdPrice:       m.UsdPrice,
		MarketCap:      m.MarketCap,
		LiquidityUsd:   m.LiquidityUsd,
		Holders:        m.Holders,
		MintAt:         m.MintAt,
		Volume1h:       m.Volume1h,
		Transactions1h: m.Transactions1h,
		Buyers1h:       m.Buyers1h,
		ChainId:        m.ChainId,
	}
}

func NewTrendingTokenRespArray(arr []*models.TrendingToken) []*TrendingTokenResp {
	resps := []*TrendingTokenResp{}
	for _, m := range arr {
		resps = append(resps, NewTrendingTokenResp(m))
	}
	return resps
}

type UserVideoInfoResp struct {
	TwitterID   string            `json:"twitter_id"`
	Address     string            `json:"address"`
	UserAddress string            `json:"user_address"`
	WalletType  models.WalletType `json:"wallet_type"`
}

func NewUserVideoInfoResp(m *models.PrivyWallet) *UserVideoInfoResp {
	if m == nil {
		return nil
	}
	return &UserVideoInfoResp{
		Address:     m.Address,
		TwitterID:   m.TwitterID,
		UserAddress: m.UserAddress,
		WalletType:  m.WalletType,
	}
}
