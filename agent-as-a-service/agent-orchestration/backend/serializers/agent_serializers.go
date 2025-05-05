package serializers

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AgentTokenReq struct {
	AgentID     string `json:"agent_id"`
	TokenSymbol string `json:"token_symbol"`
	TokenName   string `json:"token_name"`
	NetworkID   uint64 `json:"network_id"`
}

type AgentTwitterInfoReq struct {
	FarcasterID         string `json:"farcaster_id"`
	FarcasterUsername   string `json:"farcaster_username"`
	TwitterClientId     string `json:"twitter_client_id"`
	TwitterClientSecret string `json:"twitter_client_secret"`
	UserPrompt          string `json:"user_prompt"`
}

type AgentExternalInfoReq struct {
	Type             string `json:"type"`
	ExternalID       string `json:"external_id"`
	ExternalUsername string `json:"external_username"`
	ExternalName     string `json:"external_name"`
}

type TwitterInfoResp struct {
	TwitterID       string `json:"twitter_id"`
	TwitterAvatar   string `json:"twitter_avatar"`
	TwitterUsername string `json:"twitter_username"`
	TwitterName     string `json:"twitter_name"`
	Description     string `json:"description"`
	ReLink          bool   `json:"re_link"`
}

type AgentSnapshotMissionInfo struct {
	ID                  uint                      `json:"id"`
	UserPrompt          string                    `json:"user_prompt"`
	Interval            int                       `json:"interval"`
	ToolSet             models.ToolsetType        `json:"tool_set"`
	AgentType           models.AgentInfoAgentType `json:"agent_type"`
	UserTwitterIDs      string                    `json:"user_twitter_ids"`
	ToolList            string                    `json:"tool_list"`
	Tokens              string                    `json:"tokens"`
	AgentBaseModel      string                    `json:"agent_base_model"`
	AgentStoreMissionID uint                      `json:"agent_store_mission_id"`
	AgentStoreID        uint                      `json:"agent_store_id"`
	MissionStoreParams  map[string]string         `json:"mission_store_params"`
	Topics              string                    `json:"topics"`
	IsTwitterSearch     bool                      `json:"is_twitter_search"`
	IsBingSearch        bool                      `json:"is_bing_search"`
	RewardAmount        numeric.BigFloat          `json:"reward_amount"`
	MinTokenHolding     numeric.BigFloat          `json:"min_token_holding"`
	RewardUser          int                       `json:"reward_user"`
	LookupInterval      int                       `json:"lookup_interval"`
}

type AgentCategoryResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewAgentCategoryResp(m *models.AgentCategory) *AgentCategoryResp {
	if m == nil {
		return nil
	}
	return &AgentCategoryResp{
		ID:   m.ID,
		Name: m.Name,
	}
}

func NewAgentCategoryRespArry(arr []*models.AgentCategory) []*AgentCategoryResp {
	resp := make([]*AgentCategoryResp, len(arr))
	for i, m := range arr {
		resp[i] = NewAgentCategoryResp(m)
	}
	return resp
}

type AgentInfoResp struct {
	ID                        uint                        `json:"id"`
	CreatedAt                 time.Time                   `json:"created_at"`
	AgentCategoryID           uint                        `json:"agent_category_id"`
	AgentCategory             *AgentCategoryResp          `json:"agent_category"`
	TwitterInfoID             uint                        `json:"twitter_info_id"`
	TwitterInfo               *TwitterInfoResp            `json:"twitter_info"`
	AgentID                   string                      `json:"agent_id"`
	AgentContractID           string                      `json:"agent_contract_id"`
	AgentContractAddress      string                      `json:"agent_contract_address"`
	AgentName                 string                      `json:"agent_name"`
	NetworkID                 uint64                      `json:"network_id"`
	NetworkName               string                      `json:"network_name"`
	ETHAddress                string                      `json:"eth_address"`
	SOLAddress                string                      `json:"sol_address"`
	TipAmount                 numeric.BigFloat            `json:"tip_amount"`
	WalletBalance             numeric.BigFloat            `json:"wallet_balance"`
	Creator                   string                      `json:"creator"`
	Mentions                  uint                        `json:"mentions"`
	XFollowers                uint                        `json:"x_followers"`
	TipEthAddress             string                      `json:"tip_eth_address"`
	TipBtcAddress             string                      `json:"tip_btc_address"`
	TipSolAddress             string                      `json:"tip_sol_address"`
	IsFaucet                  bool                        `json:"is_faucet"`
	UserPrompt                string                      `json:"user_prompt"`
	AgentSnapshotMission      []*AgentSnapshotMissionInfo `json:"agent_snapshot_mission"`
	TokenName                 string                      `json:"token_name"`
	TokenSymbol               string                      `json:"token_symbol"`
	TokenAddress              string                      `json:"token_address"`
	TokenImageUrl             string                      `json:"token_image_url"`
	TokenStatus               string                      `json:"token_status"`
	TokenMode                 string                      `json:"create_token_mode"`
	TotalSupply               int64                       `json:"total_supply"`
	UsdMarketCap              float64                     `json:"usd_market_cap"`
	PriceUsd                  numeric.BigFloat            `json:"price_usd"`
	DexUrl                    string                      `json:"dex_url"`
	LatestTwitterPost         *AgentTwitterPostResp       `json:"latest_twitter_post"`
	Personality               string                      `json:"personality"`
	TmpTwitterInfo            *TwitterInfoResp            `json:"tmp_twitter_info"`
	IsClaimed                 bool                        `json:"is_claimed"`
	Meme                      *MemeResp                   `json:"meme"`
	TokenNetworkID            uint64                      `json:"token_network_id"`
	TokenNetworkName          string                      `json:"token_network_name"`
	ReplyLatestTime           *time.Time                  `json:"active_latest_time"`
	Thumbnail                 string                      `json:"thumbnail"`
	ReplyEnabled              bool                        `json:"reply_enabled"`
	AgentBaseModel            string                      `json:"agent_base_model"`
	AssistantCharacter        `json:",inline"`
	SocialInfo                []*models.SocialInfo      `json:"social_info"`
	VerifiedNftOwner          bool                      `json:"verified_nft_owner"`
	NftAddress                string                    `json:"nft_address"`
	NftTokenID                string                    `json:"nft_token_id"`
	NftTokenImage             string                    `json:"nft_token_image"`
	NftOwnerAddress           string                    `json:"nft_owner_address"`
	Status                    models.AssistantStatus    `json:"status"`
	InferenceCalls            int64                     `json:"inference_calls"`
	PromptCalls               int64                     `json:"prompt_calls"`
	TotalMintTwinFee          float64                   `json:"total_mint_twin_fee"`
	EstimateTwinDoneTimestamp *time.Time                `json:"estimate_twin_done_timestamp"`
	TokenDesc                 string                    `json:"token_desc"`
	ExternalChartUrl          string                    `json:"external_chart_url"`
	InferFee                  numeric.BigFloat          `json:"infer_fee"`
	MissionTopics             string                    `json:"mission_topics"`
	GraphData                 string                    `json:"graph_data"`
	AgentType                 models.AgentInfoAgentType `json:"agent_type"`
	ConfigData                string                    `json:"config_data"`
	SourceURL                 string                    `json:"source_url"`
	AuthenURL                 string                    `json:"authen_url"`
	AgentKBId                 uint                      `json:"kb_id"`
	DependAgents              string                    `json:"depend_agents"`
	RequiredWallet            bool                      `json:"required_wallet"`
	RequiredEnv               bool                      `json:"required_env"`
	IsOnchain                 bool                      `json:"is_onchain"`
	IsStreaming               bool                      `json:"is_streaming"`
	IsCustomUi                bool                      `json:"is_custom_ui"`
	Likes                     int64                     `json:"likes"`
	IsPublic                  bool                      `json:"is_public"`
	DockerPort                string                    `json:"docker_port"`
	RequiredInfo              string                    `json:"required_info"`
	InstalledCount            int64                     `json:"installed_count"`
	EnvExample                string                    `json:"env_example"`
	DisplayName               string                    `json:"display_name"`
	ShortDescription          string                    `json:"short_description"`
	IsForceUpdate             bool                      `json:"is_force_update"`
	CodeVersion               int                       `json:"code_version"`
	RunStatus                 string                    `json:"run_status"`
	Author                    string                    `json:"author"`
	Rating                    float64                   `json:"rating"`
	NumOfRating               int64                     `json:"num_of_rating"`
	NumOfOneStar              int64                     `json:"num_of_one_star"`
	NumOfTwoStar              int64                     `json:"num_of_two_star"`
	NumOfThreeStar            int64                     `json:"num_of_three_star"`
	NumOfFourStar             int64                     `json:"num_of_four_star"`
	NumOfFiveStar             int64                     `json:"num_of_five_star"`
	CodeLanguage              string                    `json:"code_language"`
}

type AgentTwitterPostResp struct {
	ID                uint                               `json:"id"`
	CreatedAt         time.Time                          `json:"created_at"`
	UpdatedAt         time.Time                          `json:"updated_at"`
	TwitterID         string                             `json:"twitter_id"`
	TwitterUser       *TwitterInfoResp                   `json:"twitter_user"`
	TwitterUsername   string                             `json:"twitter_username"`
	TwitterName       string                             `json:"twitter_name"`
	TwitterPostID     string                             `json:"twitter_post_id"`
	PostAt            *time.Time                         `json:"post_at"`
	Content           string                             `json:"content"`
	ReplyContent      string                             `json:"reply_content"`
	AgentInfoID       uint                               `json:"agent_info_id"`
	AgentInfo         *AgentInfoResp                     `json:"agent_info"`
	ReplyPostId       string                             `json:"reply_post_id"`
	ReplyPostAt       *time.Time                         `json:"reply_post_at"`
	ReplyPostReply    int                                `json:"reply_post_reply"`
	ReplyPostView     int                                `json:"reply_post_view"`
	ReplyPostFavorite int                                `json:"reply_post_favorite"`
	ReplyPostBookmark int                                `json:"reply_post_bookmark"`
	ReplyPostQuote    int                                `json:"reply_post_quote"`
	ReplyPostRetweet  int                                `json:"reply_post_retweet"`
	InscribeTxHash    string                             `json:"inscribe_tx_hash"`
	BitcoinTxHash     string                             `json:"bitcoin_tx_hash"`
	PostType          models.AgentSnapshotPostActionType `json:"post_type"`
}

func NewTwitterInfoResp(m *models.TwitterInfo) *TwitterInfoResp {
	if m == nil {
		return nil
	}
	resp := &TwitterInfoResp{
		TwitterID:       m.TwitterID,
		TwitterAvatar:   m.TwitterAvatar,
		TwitterUsername: m.TwitterUsername,
		TwitterName:     m.TwitterName,
		Description:     m.Description,
		ReLink:          false,
	}
	if m.RefreshError != "" && !strings.EqualFold(m.RefreshError, "ok") {
		resp.ReLink = true
	}
	return resp
}

func NewTwitterUserResp(m *models.TwitterUser) *TwitterInfoResp {
	if m == nil {
		return nil
	}
	resp := &TwitterInfoResp{
		TwitterID:       m.TwitterID,
		TwitterAvatar:   m.ProfileUrl,
		TwitterUsername: m.TwitterUsername,
		TwitterName:     m.Name,
	}
	return resp
}

func NewAgentSnapshotMissionResp(m *models.AgentSnapshotMission) *AgentSnapshotMissionInfo {
	if m == nil {
		return nil
	}
	resp := &AgentSnapshotMissionInfo{
		ID:                  m.ID,
		UserPrompt:          m.UserPrompt,
		Interval:            m.IntervalSec,
		ToolSet:             m.ToolSet,
		AgentType:           m.AgentType,
		UserTwitterIDs:      m.UserTwitterIds,
		Tokens:              m.Tokens,
		AgentBaseModel:      m.AgentBaseModel,
		ToolList:            m.ToolList,
		Topics:              m.Topics,
		IsTwitterSearch:     m.IsTwitterSearch,
		IsBingSearch:        m.IsBingSearch,
		RewardAmount:        m.RewardAmount,
		RewardUser:          m.RewardUser,
		MinTokenHolding:     m.MinTokenHolding,
		LookupInterval:      m.LookupInterval,
		AgentStoreMissionID: m.AgentStoreMissionID,
		AgentStoreID:        m.AgentStoreID,
	}

	if m.AgentStoreMissionID > 0 {
		resp.ToolList = ""
		resp.ToolSet = models.ToolsetType(fmt.Sprintf("%d", m.AgentStoreMissionID))
	}
	return resp
}

func NewKnowledgeBaseResp(m *models.KnowledgeBase) *KnowledgeBase {
	resp := &KnowledgeBase{}
	if m == nil {
		return nil
	}

	if err := utils.Copy(resp, m); err != nil {
		return nil
	}
	return resp
}

func NewAgentInfoResp(m *models.AgentInfo) *AgentInfoResp {
	if m == nil {
		return nil
	}
	resp := &AgentInfoResp{
		ID:                   m.ID,
		CreatedAt:            m.CreatedAt,
		AgentCategoryID:      m.AgentCategoryID,
		AgentCategory:        NewAgentCategoryResp(m.AgentCategory),
		TwitterInfoID:        m.TwitterInfoID,
		TwitterInfo:          NewTwitterInfoResp(m.TwitterInfo),
		AgentContractID:      m.AgentContractID,
		AgentContractAddress: m.AgentContractAddress,
		AgentID:              m.AgentID,
		NetworkID:            m.NetworkID,
		NetworkName:          m.NetworkName,
		ETHAddress:           m.ETHAddress,
		SOLAddress:           m.SOLAddress,
		AgentName:            m.AgentName,
		Creator:              m.Creator,
		WalletBalance:        m.EaiBalance,
		TipEthAddress:        m.TipEthAddress,
		TipBtcAddress:        m.TipBtcAddress,
		TipSolAddress:        m.TipSolAddress,
		IsFaucet:             m.IsFaucet,
		UserPrompt:           m.UserPrompt,
		AgentSnapshotMission: []*AgentSnapshotMissionInfo{},
		TokenName:            m.TokenName,
		TokenSymbol:          m.TokenSymbol,
		TokenAddress:         m.TokenAddress,
		TokenImageUrl:        m.TokenImageUrl,
		TokenMode:            m.TokenMode,
		Personality:          m.SystemPrompt,
		TmpTwitterInfo:       NewTwitterUserResp(m.TmpTwitterInfo),
		TokenNetworkID:       m.TokenNetworkID,
		TokenNetworkName:     models.GetChainName(m.TokenNetworkID),
		ReplyLatestTime:      m.ReplyLatestTime,
		Thumbnail:            m.Thumbnail,
		TokenStatus:          m.TokenStatus,
		ReplyEnabled:         m.ReplyEnabled,
		AgentBaseModel:       m.AgentBaseModel,
		VerifiedNftOwner:     m.VerifiedNftOwner,
		NftAddress:           m.NftAddress,
		NftTokenID:           m.NftTokenID,
		NftTokenImage:        m.NftTokenImage,
		NftOwnerAddress:      m.NftOwnerAddress,
		Status:               m.Status,
		InferenceCalls:       m.InferenceCalls,
		PromptCalls:          m.PromptCalls,
		TotalMintTwinFee:     m.TotalMintTwinFee,
		TokenDesc:            m.TokenDesc,
		ExternalChartUrl:     m.ExternalChartUrl,
		InferFee:             m.InferFee,
		MissionTopics:        m.MissionTopics,
		GraphData:            m.GraphData,
		AgentType:            m.AgentType,
		SourceURL:            m.SourceUrl,
		AgentKBId:            m.AgentKBId,
		AuthenURL:            m.AuthenUrl,
		DependAgents:         m.DependAgents,
		RequiredWallet:       m.RequiredWallet,
		IsOnchain:            m.IsOnchain,
		IsStreaming:          m.IsStreaming,
		IsCustomUi:           m.IsCustomUi,
		Likes:                m.Likes,
		IsPublic:             m.IsPublic,
		DockerPort:           m.DockerPort,
		RequiredInfo:         m.RequiredInfo,
		InstalledCount:       m.InstalledCount,
		RequiredEnv:          m.RequiredEnv,
		EnvExample:           m.EnvExample,
		DisplayName:          m.DisplayName,
		ShortDescription:     m.ShortDescription,
		IsForceUpdate:        m.IsForceUpdate,
		CodeVersion:          m.CodeVersion,
		RunStatus:            m.RunStatus,
		Author:               m.Author,
		Rating:               m.Rating,
		NumOfRating:          m.NumOfRating,
		NumOfOneStar:         m.NumOfOneStar,
		NumOfTwoStar:         m.NumOfTwoStar,
		NumOfThreeStar:       m.NumOfThreeStar,
		NumOfFourStar:        m.NumOfFourStar,
		NumOfFiveStar:        m.NumOfFiveStar,
		CodeLanguage:         m.GetCodeLanguage(),
	}

	if m.NftTokenImage != "" {
		m.Thumbnail = m.NftTokenImage
	}

	if m.TokenMode == "" {
		resp.TokenMode = string(models.CreateTokenModeTypeNoToken)
	}

	resp.Bio = m.GetCharacterArrayString(m.Bio)
	resp.Lore = m.GetCharacterArrayString(m.Lore)
	resp.Knowledge = m.GetCharacterArrayString(m.Knowledge)
	resp.PostExamples = m.GetCharacterArrayString(m.PostExamples)
	resp.Topics = m.GetCharacterArrayString(m.Topics)
	resp.Adjectives = m.GetCharacterArrayString(m.Adjectives)
	resp.MessageExamples = m.GetMessageExamples()
	resp.Style = m.GetStyle()
	resp.SocialInfo = m.GetSocialInfo()

	if m.TokenAddress == "" {
		resp.TokenNetworkID = 0
		resp.TokenNetworkName = ""
	}

	if m.Creator != "0xf3b5ecf8028424443ccaf35a6d46f31ce80af709" && m.RefTweetID > 0 {
		resp.IsClaimed = true
	}

	if m.TokenInfo != nil {
		resp.TotalSupply = m.TokenInfo.TotalSupply
		resp.UsdMarketCap = m.TokenInfo.UsdMarketCap
		resp.PriceUsd = m.TokenInfo.PriceUsd
		if m.TokenInfo.PriceUsd.Cmp(big.NewFloat(0)) > 0 {
			resp.DexUrl = m.TokenInfo.DexUrl
		}
	}

	if resp.DexUrl == "" {
		resp.DexUrl = models.GetDexUrl(m.TokenNetworkID, m.TokenAddress)
	}

	if m.Meme != nil && m.Meme.Status == models.MemeStatusAddPoolLevel1 {
		resp.Meme = NewMemeRespWithToken(m.Meme)
		resp.Meme.Percent = m.MemePercent
		resp.Meme.MarketCap = m.MemeMarketCap
		resp.UsdMarketCap, _ = m.MemeMarketCap.Float64()
		resp.Meme.TradeUrl = ""
		if m.TokenNetworkID == models.SOLANA_CHAIN_ID {
			resp.Meme.Status = string(models.MemeStatusAddPoolExternal)
			resp.Meme.TokenAddress = m.TokenAddress
		}
	} else {
		if m.TokenInfo != nil && m.TokenAddress != "" && m.TokenInfo.PriceUsd.Cmp(big.NewFloat(0)) > 0 {
			resp.Meme = NewMemeFromTokenInfoResp(m.TokenInfo, m)
			if m.TokenMode == string(models.CreateTokenModeTypeLinkExisting) {
				resp.Meme.Status = string(models.MemeStatusAddPoolExternal)
				if m.TokenNetworkID == models.BSC_CHAIN_ID {
					resp.Meme.TradeUrl = fmt.Sprintf(`https://four.meme/token/%s`, m.TokenAddress)
				}
			} else if m.Meme != nil {
				resp.Meme.Supply = m.Meme.Supply
				if m.Meme.Status == models.MemeStatusAddPoolExternal {
					resp.Meme.Status = string(models.MemeStatusAddPoolExternal)
					resp.Meme.TradeUrl = m.Meme.ExternalTradeUrl
				}
			}
		} else if m.Meme != nil {
			resp.Meme = NewMemeRespWithToken(m.Meme)
			resp.Meme.Percent = m.MemePercent
			resp.Meme.MarketCap = m.MemeMarketCap
		} else if m.TokenAddress != "" && (m.TokenNetworkID == models.SOLANA_CHAIN_ID || m.TokenMode == string(models.CreateTokenModeTypeLinkExisting)) {
			resp.Meme = &MemeResp{
				Status:       string(models.MemeStatusAddPoolExternal),
				TokenAddress: m.TokenAddress,
			}

			if m.TokenNetworkID == models.SOLANA_CHAIN_ID {
				resp.Meme.TradeUrl = fmt.Sprintf("https://pump.fun/coin/%s", m.TokenAddress)
				resp.Meme.MarketCap = numeric.NewBigFloatFromString("6740")
				resp.UsdMarketCap = float64(6740)
			} else if m.TokenNetworkID == models.BSC_CHAIN_ID {
				resp.Meme.TradeUrl = fmt.Sprintf(`https://four.meme/token/%s`, m.TokenAddress)
			}
		}
	}

	if len(m.AgentSnapshotMission) > 0 {
		for _, item := range m.AgentSnapshotMission {
			resp.AgentSnapshotMission = append(resp.AgentSnapshotMission, NewAgentSnapshotMissionResp(item))
		}
	}
	switch m.AgentType {
	case models.AgentInfoAgentTypeModel:
		{
			resp.RunStatus = "local"
		}
	case models.AgentInfoAgentTypeModelOnline:
		{
			resp.RunStatus = "online"
		}
	}
	return resp
}

func NewAgentInfoRespArry(arr []*models.AgentInfo) []*AgentInfoResp {
	resps := []*AgentInfoResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentInfoResp(m))
	}
	return resps
}

func NewAgentTwitterPostResp(m *models.AgentTwitterPost) *AgentTwitterPostResp {
	if m == nil {
		return nil
	}
	resp := &AgentTwitterPostResp{
		ID:                m.ID,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		TwitterID:         m.TwitterID,
		TwitterUser:       NewTwitterUserResp(m.TwitterUser),
		TwitterUsername:   m.TwitterUsername,
		TwitterName:       m.TwitterName,
		TwitterPostID:     m.TwitterPostID,
		PostAt:            m.PostAt,
		Content:           helpers.RemoveTrailingHashTag(m.Content),
		ReplyContent:      helpers.RemoveTrailingHashTag(m.ReplyContent),
		AgentInfoID:       m.AgentInfoID,
		AgentInfo:         NewAgentInfoResp(m.AgentInfo),
		ReplyPostReply:    m.ReplyPostReply,
		ReplyPostView:     m.ReplyPostView,
		ReplyPostAt:       m.ReplyPostAt,
		ReplyPostId:       m.ReplyPostId,
		ReplyPostFavorite: m.ReplyPostFavorite,
		ReplyPostBookmark: m.ReplyPostBookmark,
		ReplyPostQuote:    m.ReplyPostQuote,
		ReplyPostRetweet:  m.ReplyPostRetweet,
		InscribeTxHash:    m.InscribeTxHash,
		BitcoinTxHash:     m.BitcoinTxHash,
		PostType:          m.PostType,
	}
	return resp
}

type AgentTokenResp struct {
	ID              uint             `json:"id"`
	AgentInfoID     uint             `json:"twitter_info_id"`
	AgentInfo       *AgentInfoResp   `json:"agent_info"`
	ContractAddress string           `json:"contract_address"`
	Name            string           `json:"name"`
	Symbol          string           `json:"symbol"`
	ImageUrl        string           `json:"image_url"`
	NetworkID       uint64           `json:"network_id"`
	NetworkName     string           `json:"network_name"`
	Price           numeric.BigFloat `json:"price"`
	PriceUsd        numeric.BigFloat `json:"price_usd"`
	TotalSupply     numeric.BigFloat `json:"total_supply"`
	MarketCap       numeric.BigFloat `json:"market_cap"`
	TipAmount       numeric.BigFloat `json:"tip_amount"`
	Holders         uint             `json:"holders"`
	XFollowers      uint             `json:"x_followers"`
	BaseTokenSymbol string           `json:"base_token_symbol"`
	BaseTokenPrice  numeric.BigFloat `json:"base_token_price"`
	Percent         float64          `json:"percent"`
	TotalVolume     numeric.BigFloat `json:"total_volume"`
	Mentions        uint             `json:"mentions"`
	WalletBalance   numeric.BigFloat `json:"wallet_balance"`
	DexUrl          string           `json:"dex_url"`
}

func NewAgentTokenResp(m *models.AgentInfo) *AgentTokenResp {
	if m == nil {
		return nil
	}
	resp := &AgentTokenResp{
		ID:              m.ID,
		AgentInfoID:     m.ID,
		AgentInfo:       NewAgentInfoResp(m),
		ContractAddress: m.TokenAddress,
		Name:            m.TokenName,
		Symbol:          m.TokenSymbol,
		ImageUrl:        m.TokenImageUrl,
		TotalSupply:     m.TokenSupply,
	}

	if m.TokenInfo != nil {
		resp.Price = m.TokenInfo.Price
		resp.MarketCap = numeric.BigFloat{*models.MulBigFloats(&m.TokenInfo.PriceUsd.Float, &m.TokenSupply.Float)}
		resp.PriceUsd = m.TokenInfo.PriceUsd
		resp.BaseTokenSymbol = m.TokenInfo.BaseTokenSymbol
		resp.BaseTokenPrice = m.TokenInfo.BaseTokenPrice
		resp.TipAmount = m.TokenInfo.TipAmount
		resp.Percent = m.TokenInfo.Percent
		resp.TotalVolume = m.TokenInfo.TotalVolume
		resp.WalletBalance = m.TokenInfo.WalletBalance
		resp.DexUrl = m.TokenInfo.DexUrl
		resp.NetworkID = m.TokenInfo.NetworkID
		resp.NetworkName = m.TokenInfo.NetworkName
	}
	return resp
}

func NewAgentTokenRespArry(arr []*models.AgentInfo) []*AgentTokenResp {
	resps := []*AgentTokenResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentTokenResp(m))
	}
	return resps
}

func NewAgentTwitterPostRespArry(arr []*models.AgentTwitterPost) []*AgentTwitterPostResp {
	resps := []*AgentTwitterPostResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentTwitterPostResp(m))
	}
	return resps
}

type AgentEaiTopupResp struct {
	ID             uint                       `json:"id"`
	CreatedAt      time.Time                  `json:"created_at"`
	AgentInfoID    uint                       `json:"agent_info_id"`
	AgentInfo      *AgentInfoResp             `json:"agent_info"`
	NetworkID      uint64                     `json:"network_id"`
	NetworkName    string                     `json:"network_name"`
	Type           string                     `json:"type"`
	DepositTxHash  string                     `json:"deposit_tx_hash"`
	TopupTxHash    string                     `json:"topup_tx_hash"`
	Amount         numeric.BigFloat           `json:"amount"`
	Status         models.AgentEaiTopupStatus `json:"status"`
	InscribeTxHash string                     `json:"inscribe_tx_hash"`
	DepositAddress string                     `json:"sender_address"`
	Toolset        string                     `json:"toolset"`
	ToolsetName    string                     `json:"toolset_name"`
}

func NewAgentEaiTopupResp(m *models.AgentEaiTopup) *AgentEaiTopupResp {
	if m == nil {
		return nil
	}
	resp := &AgentEaiTopupResp{
		ID:             m.ID,
		CreatedAt:      m.CreatedAt,
		AgentInfoID:    m.AgentInfoID,
		AgentInfo:      NewAgentInfoResp(m.AgentInfo),
		Type:           string(m.Type),
		NetworkName:    models.GetChainName(m.NetworkID),
		Status:         m.Status,
		NetworkID:      m.NetworkID,
		DepositTxHash:  m.DepositTxHash,
		TopupTxHash:    m.TopupTxHash,
		Amount:         m.Amount,
		InscribeTxHash: m.InscribeTxHash,
		DepositAddress: m.DepositAddress,
		Toolset:        m.Toolset,
	}
	toolsetName := "Default"
	if v, ok := models.MAP_TOOLSET_NAME[m.Toolset]; ok {
		toolsetName = v
	}
	resp.ToolsetName = toolsetName
	return resp
}

func NewAgentEaiTopupRespArry(arr []*models.AgentEaiTopup) []*AgentEaiTopupResp {
	resps := []*AgentEaiTopupResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentEaiTopupResp(m))
	}
	return resps
}

// //
type AgentSnapshotPostActionResp struct {
	ID                  uint                                 `json:"id"`
	CreatedAt           time.Time                            `json:"created_at"`
	AgentInfoID         uint                                 `json:"agent_info_id"`
	AgentSnapshotPostID uint                                 `json:"agent_snapshot_post_id"`
	AgentTwitterId      string                               `json:"agent_twitter_id"`
	Type                models.AgentSnapshotPostActionType   `json:"type"`
	TargetUsername      string                               `json:"target_username"`
	TargetTwitterId     string                               `json:"target_twitter_id"`
	Tweetid             string                               `json:"tweetid"`
	Content             string                               `json:"content"`
	Status              models.AgentSnapshotPostActionStatus `json:"status"`
	FollowerCount       uint                                 `json:"follower_count"`
	Price               numeric.BigFloat                     `json:"price"`
	InscribeTxHash      string                               `json:"inscribe_tx_hash"`
	BitcoinTxHash       string                               `json:"bitcoin_tx_hash"`
}

type AgentSnapshotPostResp struct {
	ID                      uint                           `json:"id"`
	CreatedAt               time.Time                      `json:"created_at"`
	AgentInfoID             uint                           `json:"agent_info_id"`
	AgentInfo               *AgentInfoResp                 `json:"agent_info"`
	InferData               string                         `json:"infer_data"`
	InferSnapshotHash       string                         `json:"infer_snapshot_hash"`
	InferTxHash             string                         `json:"infer_tx_hash"`
	InferAt                 *time.Time                     `json:"infer_at"`
	InferNum                uint                           `json:"infer_num"`
	InferOutputData         string                         `json:"infer_output_data"`
	InferOutputAt           *time.Time                     `json:"infer_output_at"`
	Status                  models.AgentSnapshotPostStatus `json:"status"`
	InscribeTxHash          string                         `json:"inscribe_tx_hash"`
	BitcoinTxHash           string                         `json:"bitcoin_tx_hash"`
	Fee                     numeric.BigFloat               `json:"fee"`
	AgentSnapshotPostAction []*AgentSnapshotPostActionResp `json:"agent_snapshot_post_action"`
}

func NewAgentSnapshotPostResp(m *models.AgentSnapshotPost) *AgentSnapshotPostResp {
	if m == nil {
		return nil
	}
	resp := &AgentSnapshotPostResp{
		ID:                      m.ID,
		CreatedAt:               m.CreatedAt,
		AgentInfoID:             m.AgentInfoID,
		AgentInfo:               NewAgentInfoResp(m.AgentInfo),
		InferData:               m.InferData,
		InferSnapshotHash:       m.InferSnapshotHash,
		InferTxHash:             m.InferTxHash,
		InferAt:                 m.InferAt,
		InferNum:                m.InferNum,
		InferOutputData:         m.InferOutputData,
		InferOutputAt:           m.InferOutputAt,
		Status:                  m.Status,
		InscribeTxHash:          m.InscribeTxHash,
		BitcoinTxHash:           m.BitcoinTxHash,
		Fee:                     m.Fee,
		AgentSnapshotPostAction: NewAgentSnapshotPostActionRespArry(m.AgentSnapshotPostAction),
	}
	return resp
}

func NewAgentSnapshotPostRespArry(arr []*models.AgentSnapshotPost) []*AgentSnapshotPostResp {
	resps := []*AgentSnapshotPostResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentSnapshotPostResp(m))
	}
	return resps
}

func NewAgentSnapshotPostActionResp(m *models.AgentSnapshotPostAction) *AgentSnapshotPostActionResp {
	if m == nil {
		return nil
	}
	resp := &AgentSnapshotPostActionResp{
		ID:                  m.ID,
		CreatedAt:           m.CreatedAt,
		AgentInfoID:         m.AgentInfoID,
		AgentSnapshotPostID: m.AgentSnapshotPostID,
		AgentTwitterId:      m.AgentTwitterId,
		Type:                m.Type,
		TargetUsername:      m.TargetUsername,
		TargetTwitterId:     m.TargetTwitterId,
		Tweetid:             m.Tweetid,
		Content:             m.Content,
		Status:              m.Status,
		FollowerCount:       m.FollowerCount,
		Price:               m.Price,
		InscribeTxHash:      m.InscribeTxHash,
		BitcoinTxHash:       m.BitcoinTxHash,
	}
	return resp
}

func NewAgentSnapshotPostActionRespArry(arr []*models.AgentSnapshotPostAction) []*AgentSnapshotPostActionResp {
	resps := []*AgentSnapshotPostActionResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentSnapshotPostActionResp(m))
	}
	return resps
}

type AdminAgentActionReq struct {
	ChainID         uint64                             `json:"chain_id"`
	AgentContractId string                             `json:"agent_contract_id"`
	ActionType      models.AgentSnapshotPostActionType `json:"action_type"`
	IsTesting       bool                               `json:"is_testing"`
	RefID           string                             `json:"ref_id"`
	MissionID       uint                               `json:"mission_id"`
	ConversationId  string                             `json:"conversation_id"`
	InscribeTxHash  string                             `json:"inscribe_tx_hash"`
	BitcoinTxHash   string                             `json:"bitcoin_tx_hash"`
	ActionInput     struct {
		Content        string           `json:"content"`
		TweetId        string           `json:"tweet_id"`
		Twid           string           `json:"twid"`
		TargetUsername string           `json:"target_username"`
		Comment        string           `json:"comment"`
		Name           string           `json:"name"`
		Symbol         string           `json:"symbol"`
		Description    string           `json:"description"`
		ImageUrl       string           `json:"image_url"`
		Price          numeric.BigFloat `json:"price"`
	} `json:"action_input"`
}

type CreateAgentWalletActionReq struct {
	Username    string      `json:"username"`
	ActionType  string      `json:"action_type"`
	ActionInput interface{} `json:"action_input"`
}

type AdminCreatePumpfunMemeReq struct {
	RefID       string  `json:"ref_id"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	ImageBase64 string  `json:"image_base64"`
}

type AdminTradePumpfunMemeReq struct {
	RefID  string  `json:"ref_id"`
	Action string  `json:"action"`
	Mint   string  `json:"mint"`
	Amount float64 `json:"amount"`
}

type SolanaTokenBalanceResp struct {
	IsNative bool    `json:"is_native"`
	Mint     string  `json:"mint"`
	Amount   float64 `json:"amount"`
}

type WalletActionTradeResp struct {
	CreatedAt time.Time        `json:"created_at"`
	Mint      string           `json:"mint"`
	Side      string           `json:"side"`
	AmountIn  numeric.BigFloat `json:"amount_in"`
	AmountOut numeric.BigFloat `json:"amount_out"`
	TxHash    string           `json:"tx_hash"`
}

type PumpFunTradeResp struct {
	Signature   string  `json:"signature"`
	Mint        string  `json:"mint"`
	SolAmount   float64 `json:"sol_amount"`
	TokenAmount float64 `json:"token_amount"`
	IsBuy       bool    `json:"is_buy"`
	Timestamp   int64   `json:"timestamp"`
}

type DataChartResp struct {
	Time  time.Time `json:"time"`
	Price float64   `json:"price"`
}

// //coinbase
type AdminTradeBaseMemeReq struct {
	FromAssetId string `json:"from_asset_id"`
	ToAssetId   string `json:"to_asset_id"`
	Amount      string `json:"amount"`
}

type AgentSnapshotMissionConfigsResp struct {
	ID          uint               `json:"id"`
	NetworkID   uint64             `json:"network_id"`
	ToolSet     models.ToolsetType `json:"tool_set"`
	ToolSetName string             `json:"tool_set_name"`
}

func NewAgentSnapshotMissionConfigsResp(m *models.AgentSnapshotMissionConfigs) *AgentSnapshotMissionConfigsResp {
	if m == nil {
		return nil
	}
	resp := &AgentSnapshotMissionConfigsResp{
		ID:          m.ID,
		NetworkID:   m.NetworkID,
		ToolSet:     m.ToolSet,
		ToolSetName: m.ToolSetName,
	}
	return resp
}

func NewAgentSnapshotMissionConfigsRespArry(arr []*models.AgentSnapshotMissionConfigs) []*AgentSnapshotMissionConfigsResp {
	resps := []*AgentSnapshotMissionConfigsResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentSnapshotMissionConfigsResp(m))
	}
	return resps
}

type AgentChatMessageReq struct {
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type AgentTradeTokenResp struct {
	NetworkID    uint64 `json:"network_id"`
	NetworkName  string `json:"network_name"`
	TokenSymbol  string `json:"token_symbol"`
	TokenName    string `json:"token_name"`
	TokenAddress string `json:"token_address"`
}

func NewAgentTradeTokenResp(m *models.AgentTradeToken) *AgentTradeTokenResp {
	if m == nil {
		return nil
	}
	resp := &AgentTradeTokenResp{
		NetworkID:    m.NetworkID,
		NetworkName:  m.NetworkName,
		TokenSymbol:  m.TokenSymbol,
		TokenName:    m.TokenName,
		TokenAddress: m.TokenAddress,
	}
	return resp
}

func NewAgentTradeTokenRespArry(arr []*models.AgentTradeToken) []*AgentTradeTokenResp {
	resps := []*AgentTradeTokenResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentTradeTokenResp(m))
	}
	return resps
}

type AdminTweetReq struct {
	Text string `json:"text"`
}

type AdminAgentActionByRefReq struct {
	ActionType models.AgentSnapshotPostActionType `json:"action_type"`
	Reason     string                             `json:"reason"`
}

type TwitterTweetLikedResp struct {
	LikedUserID      string    `json:"liked_user_id"`
	TweetID          string    `json:"tweet_id"`
	TwitterID        string    `json:"twitter_id"`
	LikeCount        int       `json:"like_count"`
	RetweetCount     int       `json:"retweet_count"`
	ReplyCount       int       `json:"reply_count"`
	QuoteCount       int       `json:"quote_count"`
	ImpressionCount  int       `json:"impression_count"`
	FullText         string    `json:"full_text"`
	PostedAt         time.Time `json:"posted_at"`
	InReplyToUserID  string    `json:"in_reply_to_user_id"`
	InReplyToTweetID string    `json:"in_reply_to_tweet_id"`
	IsReply          bool      `json:"is_reply"`
	OriginalText     string    `json:"original_text"`

	// Retweet-related fields
	IsRetweet     bool   `json:"is_retweet"`
	RepostTweetID string `json:"repost_tweet_id"`
	RepostText    string `json:"repost_text"`

	// Quote-related fields
	IsQuote      bool   `json:"is_quote"`
	QuoteTweetID string `json:"quote_tweet_id"`
	QuoteText    string `json:"quote_text"`
}

func NewTwitterTweetLikedResp(m *models.TwitterTweetLiked) *TwitterTweetLikedResp {
	if m == nil {
		return nil
	}
	resp := &TwitterTweetLikedResp{
		LikedUserID:      m.LikedUserID,
		TweetID:          m.TweetID,
		TwitterID:        m.TwitterID,
		LikeCount:        m.LikeCount,
		RetweetCount:     m.RetweetCount,
		ReplyCount:       m.ReplyCount,
		QuoteCount:       m.QuoteCount,
		ImpressionCount:  m.ImpressionCount,
		FullText:         m.FullText,
		PostedAt:         m.PostedAt,
		InReplyToUserID:  m.InReplyToUserID,
		InReplyToTweetID: m.InReplyToTweetID,
		IsReply:          m.IsReply,
		OriginalText:     m.OriginalText,
		IsRetweet:        m.IsRetweet,
		RepostTweetID:    m.RepostTweetID,
		RepostText:       m.RepostText,
		IsQuote:          m.IsQuote,
		QuoteTweetID:     m.QuoteTweetID,
		QuoteText:        m.QuoteText,
	}
	return resp
}

func NewTwitterTweetLikedRespArr(arr []*models.TwitterTweetLiked) []*TwitterTweetLikedResp {
	resps := []*TwitterTweetLikedResp{}
	for _, m := range arr {
		resps = append(resps, NewTwitterTweetLikedResp(m))
	}
	return resps
}

type AgentReportResp struct {
	NetworkID   uint64 `json:"network_id"`
	NetworkName string `json:"network_name"`
	Counts      int64  `json:"counts"`
}

func NewAgentReportResp(m *models.AgentInfo) *AgentReportResp {
	if m == nil {
		return nil
	}

	resp := &AgentReportResp{
		NetworkID:   m.NetworkID,
		NetworkName: m.NetworkName,
		Counts:      m.Counts,
	}
	return resp
}

func NewAgentReportRespArr(arr []*models.AgentInfo) []*AgentReportResp {
	resps := []*AgentReportResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentReportResp(m))
	}
	return resps
}

type AgentChainFeeResp struct {
	NetworkID uint64           `json:"network_id"`
	InferFee  numeric.BigFloat `json:"infer_fee"`
	MintFee   numeric.BigFloat `json:"mint_fee"`
	TokenFee  numeric.BigFloat `json:"token_fee"`
}

func NewAgentChainFeeResp(m *models.AgentChainFee) *AgentChainFeeResp {
	if m == nil {
		return nil
	}
	resp := &AgentChainFeeResp{
		NetworkID: m.NetworkID,
		InferFee:  m.InferFee,
		MintFee:   m.MintFee,
		TokenFee:  m.TokenFee,
	}
	return resp
}

func NewAgentChainFeeRespArr(arr []*models.AgentChainFee) []*AgentChainFeeResp {
	resps := []*AgentChainFeeResp{}
	for _, m := range arr {
		resps = append(resps, NewAgentChainFeeResp(m))
	}
	return resps
}
