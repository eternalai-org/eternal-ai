package models

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type AgentSnapshotPostStatus string

const (
	AgentSnapshotPostStatusNew              AgentSnapshotPostStatus = "new"
	AgentSnapshotPostStatusInvalid          AgentSnapshotPostStatus = "invalid"
	AgentSnapshotPostStatusValid            AgentSnapshotPostStatus = "valid"
	AgentSnapshotPostStatusInferNew         AgentSnapshotPostStatus = "infer_new"
	AgentSnapshotPostStatusInferSubmitted   AgentSnapshotPostStatus = "infer_submitted"
	AgentSnapshotPostStatusInferError       AgentSnapshotPostStatus = "infer_error"
	AgentSnapshotPostStatusInferFailed      AgentSnapshotPostStatus = "infer_failed"
	AgentSnapshotPostStatusInferExpired     AgentSnapshotPostStatus = "infer_expired"
	AgentSnapshotPostStatusInferResolved    AgentSnapshotPostStatus = "infer_resolved"
	AgentSnapshotPostStatusInferRefund      AgentSnapshotPostStatus = "infer_refund"
	AgentSnapshotPostStatusReplied          AgentSnapshotPostStatus = "replied"
	AgentSnapshotPostStatusRepliedError     AgentSnapshotPostStatus = "replied_error"
	AgentSnapshotPostStatusRepliedCancelled AgentSnapshotPostStatus = "replied_cancelled"
	AgentSnapshotPostStatusReposted         AgentSnapshotPostStatus = "reposted"
	AgentSnapshotPostStatusRepostedError    AgentSnapshotPostStatus = "reposted_error"
)

type AgentSnapshotPost struct {
	gorm.Model
	NetworkID               uint64
	UserID                  uint `gorm:"index"`
	User                    *User
	AgentInfoID             uint `gorm:"index"`
	AgentInfo               *AgentInfo
	AgentSnapshotMissionID  uint `gorm:"index"`
	AgentSnapshotMission    *AgentSnapshotMission
	InferData               string `gorm:"type:longtext"`
	InferSnapshotHash       string
	InferTxHash             string `gorm:"index"`
	InferAt                 *time.Time
	InferNum                uint                    `gorm:"default:0"`
	InferOutputData         string                  `gorm:"type:longtext"`
	InferOutputAt           *time.Time              `gorm:"index"`
	Status                  AgentSnapshotPostStatus `gorm:"index"`
	InscribeTxHash          string
	BitcoinTxHash           string
	Fee                     numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Error                   string
	UserPrompt              string `gorm:"type:longtext"`
	HeadSystemPrompt        string `gorm:"type:longtext"`
	SystemPrompt            string `gorm:"type:longtext"`
	SystemReminder          string `gorm:"type:longtext"`
	Task                    string
	Toolset                 string
	ToolList                string `gorm:"type:longtext"`
	AgentMetaData           string `gorm:"type:longtext"`
	ResponseId              string
	AgentBaseModel          string
	ReactMaxSteps           int `gorm:"default:0"`
	AgentSnapshotPostAction []*AgentSnapshotPostAction
	OrgTweetID              string
	Token                   string
	AgentStoreMissionID     uint
	AgentStoreID            uint
	AgentStore              *AgentStore
	IsRated                 bool
	AgentStoreMissionFee    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
}

type (
	AgentSnapshotPostActionType   string
	AgentSnapshotPostActionStatus string
)

const (
	AgentSnapshotPostActionTypeTweet               AgentSnapshotPostActionType = "tweet"
	AgentSnapshotPostActionTypeRetweet             AgentSnapshotPostActionType = "retweet"
	AgentSnapshotPostActionTypeFollow              AgentSnapshotPostActionType = "follow"
	AgentSnapshotPostActionTypeReply               AgentSnapshotPostActionType = "reply"
	AgentSnapshotPostActionTypeReplyMulti          AgentSnapshotPostActionType = "reply_multi"
	AgentSnapshotPostActionTypeReplyMultiUnlimited AgentSnapshotPostActionType = "reply_multi_unlimited"
	AgentSnapshotPostActionTypeCreateToken         AgentSnapshotPostActionType = "create_token"
	AgentSnapshotPostActionTypeCreateAgent         AgentSnapshotPostActionType = "create_agent"
	AgentSnapshotPostActionTypeGenerateVideo       AgentSnapshotPostActionType = "generate_video"
	AgentSnapshotPostActionTypeQuoteTweet          AgentSnapshotPostActionType = "quote_tweet"
	AgentSnapshotPostActionTypeInscribeTweet       AgentSnapshotPostActionType = "inscribe_tweet"
	AgentSnapshotPostActionTypeTweetV2             AgentSnapshotPostActionType = "tweet_v2"
	AgentSnapshotPostActionTypeTweetMulti          AgentSnapshotPostActionType = "tweet_multi"
	AgentSnapshotPostActionTypeCreateLaunchpad     AgentSnapshotPostActionType = "create_launchpad"
	AgentSnapshotPostActionTypeUnknown             AgentSnapshotPostActionType = "unknown"

	AgentSnapshotPostActionTypeTradeHold     AgentSnapshotPostActionType = "hold"
	AgentSnapshotPostActionTypeTradeBuy      AgentSnapshotPostActionType = "buy"
	AgentSnapshotPostActionTypeTradeSell     AgentSnapshotPostActionType = "sell"
	AgentSnapshotPostActionTypeTradeAnalytic AgentSnapshotPostActionType = "analytic"
	AgentSnapshotPostActionTypeLaunchpadJoin AgentSnapshotPostActionType = "launchpad_join"

	AgentSnapshotPostActionStatusNew            AgentSnapshotPostActionStatus = "new"
	AgentSnapshotPostActionStatusDone           AgentSnapshotPostActionStatus = "done"
	AgentSnapshotPostActionStatusDone_          AgentSnapshotPostActionStatus = "done_"
	AgentSnapshotPostActionStatusDoneError      AgentSnapshotPostActionStatus = "done_error"
	AgentSnapshotPostActionStatusDoneDuplicated AgentSnapshotPostActionStatus = "done_duplicated"
	AgentSnapshotPostActionStatusDoneCancelled  AgentSnapshotPostActionStatus = "done_cancelled"
	AgentSnapshotPostActionStatusInvalid        AgentSnapshotPostActionStatus = "invalid"
	AgentSnapshotPostActionStatusTesting        AgentSnapshotPostActionStatus = "testing"
	AgentSnapshotPostActionStatusInscribing     AgentSnapshotPostActionStatus = "inscribing"
	AgentSnapshotPostActionStatusPaid           AgentSnapshotPostActionStatus = "paid"
)

type AgentSnapshotPostAction struct {
	gorm.Model
	NetworkID              uint64
	AgentInfoID            uint `gorm:"index"`
	AgentInfo              *AgentInfo
	AgentSnapshotMissionID uint `gorm:"index"`
	AgentSnapshotMission   *AgentSnapshotMission
	AgentSnapshotPostID    uint `gorm:"index"`
	AgentSnapshotPost      *AgentSnapshotPost
	AgentTwitterId         string
	Type                   AgentSnapshotPostActionType `gorm:"index"`
	TargetUsername         string
	TargetTwitterId        string                        `gorm:"index"`
	ConversationId         string                        `gorm:"index"`
	Tweetid                string                        `gorm:"index"`
	Content                string                        `gorm:"type:longtext"`
	Description            string                        `gorm:"type:longtext"`
	Status                 AgentSnapshotPostActionStatus `gorm:"index"`
	RefId                  string                        `gorm:"type:longtext"`
	RefIds                 string                        `gorm:"type:longtext"`
	Error                  string
	FollowerCount          uint       `gorm:"default:0"`
	IsApproved             bool       `gorm:"default:0"`
	ScheduleAt             *time.Time `gorm:"index"`
	ExecutedAt             *time.Time
	Fee                    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsMigrated             bool             `gorm:"default:0"`
	TokenName              string
	TokenSymbol            string
	TokenAddress           string
	TokenHash              string
	TokenInferID           string
	TokenImageUrl          string
	ReqRefID               string
	ToolSet                ToolsetType
	InscribeId             string
	InscribeTxHash         string
	BitcoinTxHash          string
	Price                  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RewardAmount           numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RewardUser             int
}

type (
	LuckyMoneyStatus string
)

const (
	LuckyMoneyStatusNew        LuckyMoneyStatus = "new"
	LuckyMoneyStatusProcessing LuckyMoneyStatus = "processing"
	LuckyMoneyStatusInvalid    LuckyMoneyStatus = "invalid"
	LuckyMoneyStatusDone       LuckyMoneyStatus = "done"
)

type AbilityLuckyMoney struct {
	gorm.Model
	NetworkID                 uint64
	AgentInfoID               uint `gorm:"index"`
	AgentInfo                 *AgentInfo
	AgentSnapshotMissionID    uint `gorm:"index"`
	AgentSnapshotMission      *AgentSnapshotMission
	AgentSnapshotPostID       uint `gorm:"index"`
	AgentSnapshotPost         *AgentSnapshotPost
	AgentSnapshotPostActionID uint `gorm:"index"`
	AgentSnapshotPostAction   *AgentSnapshotPostAction
	RewardAmount              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TokenBalance              numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status                    LuckyMoneyStatus
	Error                     string
	TweetID                   string     `gorm:"unique_index"`
	TweetAt                   *time.Time `gorm:"index"`
	TwitterID                 string
	TwitterUsername           string
	TwitterName               string
	Content                   string `gorm:"type:longtext"`
	UserAddress               string
	RefTweetID                string
	TxHash                    string
}

type ToolsetType string

const (
	ToolsetTypeDefault          ToolsetType = "default"
	ToolsetTypeReplyMentions    ToolsetType = "reply_mentions"
	ToolsetTypeReplyNonMentions ToolsetType = "reply_non_mentions"
	ToolsetTypeShadowReply      ToolsetType = "shadow_reply"
	ToolsetTypeFollow           ToolsetType = "follow"
	ToolsetTypePost             ToolsetType = "post"
	ToolsetTypePostSearchV2     ToolsetType = "post_search_v2"
	ToolsetTypePostSearchV3     ToolsetType = "post_search_v3"

	ToolsetTypeIssueToken              ToolsetType = "issue_token"
	ToolsetTypeInscribeTweet           ToolsetType = "inscribe_tweet"
	ToolsetTypeInscribeTweetNews       ToolsetType = "tweet_news"
	ToolsetTypeTrading                 ToolsetType = "trading"
	ToolsetTypeTradeNews               ToolsetType = "trade_news"
	ToolsetTypeTradeAnalytics          ToolsetType = "trade_analytics"
	ToolsetTypeTradeAnalyticsOnTwitter ToolsetType = "trade_analytics_twitter"
	ToolsetTypeTradeAnalyticsMentions  ToolsetType = "trade_analytics_mentions"
	ToolsetTypeLuckyMoneys             ToolsetType = "lucky_moneys"
	ToolsetTypeLaunchpadJoin           ToolsetType = "launchpad_join"
	ToolsetTypeMissionStore            ToolsetType = "mission_store"

	ToolsetTypeReplyMentionsFarcaster    ToolsetType = "reply_mentions_farcaster"
	ToolsetTypeReplyNonMentionsFarcaster ToolsetType = "reply_non_mentions_farcaster"
	ToolsetTypePostFarcaster             ToolsetType = "post_farcaster"

	ToolsetTypeTavily ToolsetType = "tavily"
)

type AgentSnapshotMission struct {
	gorm.Model
	NetworkID           uint64
	AgentInfoID         uint `gorm:"index"`
	AgentInfo           *AgentInfo
	UserPrompt          string `gorm:"type:longtext"`
	IntervalSec         int    `gorm:"default:0"`
	LookupInterval      int    `gorm:"default:0"`
	Enabled             bool   `gorm:"default:0"`
	ReplyEnabled        bool   `gorm:"default:0"`
	IsTesting           bool   `gorm:"default:0"`
	ToolSet             ToolsetType
	AgentType           AgentInfoAgentType `gorm:"default:0"`
	InferAt             *time.Time
	SkipThough          bool   `gorm:"default:0"`
	ToolList            string `gorm:"type:longtext"`
	UserTwitterIds      string `gorm:"type:longtext"`
	TeleChatID          string
	Tokens              string `gorm:"type:longtext"`
	ReactMaxSteps       int    `gorm:"default:0"`
	NotDelay            bool   `gorm:"default:0"`
	AgentBaseModel      string
	Topics              string           `gorm:"type:longtext"`
	IsTwitterSearch     bool             `gorm:"default:0"`
	IsBingSearch        bool             `gorm:"default:0"`
	RewardAmount        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	RewardUser          int
	MinTokenHolding     numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	LaunchpadMemberID   uint
	AgentStoreID        uint
	AgentStore          *AgentStore
	AgentStoreMissionID uint
	AgentStoreMission   *AgentStoreMission
}

type TeleMsgStatus string

const (
	TeleMsgStatusNew TeleMsgStatus = "new"
)

type AgentTeleMsg struct {
	gorm.Model
	MessageID              string `gorm:"unique_index"`
	MessageDate            *time.Time
	Content                string `gorm:"type:longtext"`
	ChatID                 string
	ChatUsername           string
	AgentInfoID            uint `gorm:"index"`
	AgentInfo              *AgentInfo
	AgentSnapshotMissionID uint `gorm:"index"`
	AgentSnapshotMission   *AgentSnapshotMission
	Status                 TeleMsgStatus
}

type PlatformType string

const (
	PlatformTypeTwitter   PlatformType = "twitter"
	PlatformTypeFarcaster PlatformType = "farcaster"
)

type AgentSnapshotMissionConfigs struct {
	gorm.Model
	NetworkID   uint64
	ToolSet     ToolsetType
	ToolSetName string
	Platform    PlatformType
	IsTesting   bool `gorm:"default:0"`
}

type AgentSnapshotMissionResp struct {
	ID           uint               `json:"id"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	NetworkID    uint64             `json:"network_id"`
	AgentInfoID  uint               `json:"agent_info_id"`
	AgentInfo    *AgentInfo         `json:"agent_info"`
	UserPrompt   string             `json:"user_prompt"`
	IntervalSec  int                `json:"interval_sec"`
	Enabled      bool               `json:"enabled"`
	ReplyEnabled bool               `json:"reply_enabled"`
	IsTesting    bool               `json:"is_testing"`
	ToolSet      ToolsetType        `json:"tool_set"`
	AgentType    AgentInfoAgentType `json:"agent_type"`
	InferAt      *time.Time         `json:"infer_at"`
	SkipThough   bool               `json:"skip_though"`
}

type AgentSnapshotPostActionResp struct {
	ID                   uint                          `json:"id"`
	ExecutedAt           *time.Time                    `json:"executed_at"`
	ToolSet              string                        `json:"tool_set"`
	Type                 string                        `json:"type"`
	Tweetid              string                        `json:"tweetid"`
	TargetUsername       string                        `json:"target_username"`
	PostTweetid          string                        `json:"post_tweetid"`
	Content              string                        `json:"content"`
	UserPrompt           string                        `json:"user_prompt"`
	InferInputData       string                        `json:"infer_input_data"`
	InferOutputData      string                        `json:"infer_output_data"`
	Description          string                        `json:"description"`
	TargetTwitterId      string                        `json:"target_twitter_id"`
	TokenInferID         string                        `json:"token_infer_id"`
	TokenImageUrl        string                        `json:"token_image_url"`
	AgentSnapshotMission *AgentSnapshotMissionResp     `json:"agent_snapshot_mission"`
	Status               AgentSnapshotPostActionStatus `json:"status"`
}

type ParamWakeupRequest struct {
	QuoteUsername string `json:"quote_username"`
	ReactMaxSteps int    `json:"react_max_steps"`
}

type WakeupRequestMetadata struct {
	TwitterId       string             `json:"twitter_id"`
	TwitterUsername string             `json:"twitter_username"`
	AgentContractId string             `json:"agent_contract_id"`
	ChainId         string             `json:"chain_id"`
	SystemReminder  string             `json:"system_reminder"`
	Params          ParamWakeupRequest `json:"params"`
	RefID           string             `json:"ref_id"`
	KnowledgeBaseId string             `json:"knowledge_base_id"`
}

type AgentWakeupKnowledgeBase struct {
	KbId    string `json:"kb_id"`
	ChainId string `json:"chain_id"`
}

type AgentMetadataRequest struct {
	TokenInfo struct {
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
		Chain   string `json:"chain"`
	} `json:"token_info"`
	KbAgents []AgentWakeupKnowledgeBase `json:"kb_agents"`
}

type CallWakeupRequest struct {
	Toolkit       []interface{}         `json:"toolkit"`
	Prompt        string                `json:"prompt"`
	Task          string                `json:"task"`
	Toolset       string                `json:"toolset"`
	MetaData      WakeupRequestMetadata `json:"meta_data"`
	SystemPrompt  string                `json:"system_prompt"`
	Model         string                `json:"model"`
	AgentMetaData AgentMetadataRequest  `json:"agent_meta_data"`
	ToolList      string                `json:"tool_list"`
}
