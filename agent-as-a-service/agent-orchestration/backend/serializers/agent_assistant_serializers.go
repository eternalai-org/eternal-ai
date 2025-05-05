package serializers

import (
	"encoding/json"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AssistantCharacter struct {
	Bio             []string `json:"bio"`
	Lore            []string `json:"lore"`
	Knowledge       []string `json:"knowledge"`
	MessageExamples [][]struct {
		User    string `json:"user"`
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"messageExamples"`
	PostExamples []string            `json:"postExamples"`
	Topics       []string            `json:"topics"`
	Style        map[string][]string `json:"style"`
	Adjectives   []string            `json:"adjectives"`
}

type AssistantsReq struct {
	ID                   uint                       `json:"id"`
	CategoryID           uint                       `json:"category_id"`
	AgentID              string                     `json:"agent_id"`
	AgentName            string                     `json:"agent_name"`
	Creator              string                     `json:"creator"`
	AgentContractAddress string                     `json:"agent_contract_address"`
	AgentContractID      string                     `json:"agent_contract_id"`
	Minter               string                     `json:"minter"`
	MetaData             string                     `json:"meta_data"`
	ChainID              uint64                     `json:"chain_id"`
	Thumbnail            string                     `json:"thumbnail"`
	CreateTokenMode      models.CreateTokenModeType `json:"create_token_mode"`
	Status               models.AssistantStatus     `json:"status"`
	UserPrompt           string                     `json:"user_prompt"`
	SystemContent        string                     `json:"system_content"`
	Ticker               string                     `json:"ticker"`
	TokenName            string                     `json:"token_name"`
	TokenAddress         string                     `json:"token_address"`
	TokenImageUrl        string                     `json:"token_image_url"`
	TokenDesc            string                     `json:"token_desc"`
	AssistantCharacter   `json:",inline"`
	TokenChainId         string                    `json:"token_chain_id"`
	SocialInfo           []*models.SocialInfo      `json:"social_info"`
	AgentType            models.AgentInfoAgentType `json:"agent_type"`
	ConfigData           string                    `json:"config_data"`

	VerifiedNFTOwner   bool   `json:"verified_nft_owner"`
	NFTAddress         string `json:"nft_address"`
	NFTTokenID         string `json:"nft_token_id"`
	NFTTokenImage      string `json:"nft_token_image"`
	NFTOwnerAddress    string `json:"nft_owner_address"`
	NFTSignature       string `json:"nft_signature"`
	NFTSignMessage     string `json:"nft_sign_message"`
	NFTDelegateAddress string `json:"nft_delegate_address"`
	NFTPublicKey       string `json:"nft_public_key"`

	AgentBaseModel         string                  `json:"agent_base_model"`
	TwinTwitterUsernames   string                  `json:"twin_twitter_usernames"`
	MissionTopics          string                  `json:"mission_topics"`
	CreateKnowledgeRequest *CreateKnowledgeRequest `json:"create_knowledge_request"`
	KbIds                  []uint                  `json:"kb_ids"`

	SourceUrl string `json:"source_url"`
	AuthenUrl string `json:"authen_url"`

	MinFeeToUse numeric.BigFloat `json:"min_fee_to_use"`
	Worker      string           `json:"worker"`

	DependAgents     string  `json:"depend_agents"`
	RequiredWallet   *bool   `json:"required_wallet"`
	RequiredEnv      *bool   `json:"required_env"`
	EnvExample       string  `json:"env_example"`
	IsOnchain        *bool   `json:"is_onchain"`
	IsStreaming      *bool   `json:"is_streaming"`
	IsCustomUi       *bool   `json:"is_custom_ui"`
	RequiredInfo     *string `json:"required_info"`
	InferFee         *string `json:"infer_fee"`
	DisplayName      string  `json:"display_name"`
	ShortDescription string  `json:"short_description"`
	IsForceUpdate    *bool   `json:"is_force_update"`
	Author           string  `json:"author"`
}

func (m *AssistantsReq) GetAssistantCharacter(character interface{}) string {
	urlsJson, _ := json.Marshal(character)
	if string(urlsJson) == "null" {
		return ""
	}
	return string(urlsJson)
}

type UpdateAgentAssistantInContractRequest struct {
	ID                      uint   `json:"id"`
	AgentID                 string `json:"agent_id"`
	HashName                string `json:"hash_name"`
	HashSystemPrompt        string `json:"hash_system_prompt"`
	SignatureName           string `json:"signature_name"`
	SignatureSystemPrompt   string `json:"signature_system_prompt"`
	RandomNonceName         string `json:"random_nonce_name"`
	RandomNonceSystemPrompt string `json:"random_nonce_system_prompt"`
}

type UpdateAgentAssistantInContractResponse struct {
	TxUpdateName         string `json:"tx_update_name"`
	TxUpdateSystemPrompt string `json:"tx_update_system_prompt"`
}

type UpdateTwinStatusRequest struct {
	AgentID              string  `json:"agent_id"`
	TwinStatus           string  `json:"twin_status"`
	KnowledgeBaseID      string  `json:"knowledge_base_id"`
	SystemPrompt         string  `json:"system_prompt"`
	TwinTrainingProgress float64 `json:"twin_training_progress"`
	TwinTrainingMessage  string  `json:"twin_training_message"`
}

// ///////////////
type Assistants struct {
	ID                   string    `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	InferFee             string    `bson:"infer_fee" json:"infer_fee"`
	AgentName            string    `bson:"agent_name" json:"agent_name"`
	Creator              string    `bson:"creator" json:"creator"`
	Minter               string    `bson:"minter" json:"minter"`
	ContractAgentID      string    `bson:"contract_agent_id" json:"contract_agent_id"`
	MetaData             string    `bson:"meta_data" json:"meta_data"`
	ChainID              uint64    `bson:"chain_id" json:"chain_id"`
	AgentContractAddress string    `bson:"agent_contract_address" json:"agent_contract_address"`
	TxHash               string    `bson:"tx_hash" json:"tx_hash"`
	Uri                  string    `bson:"uri" json:"uri"`
	Thumbnail            string    `bson:"thumbnail" json:"thumbnail"`

	TokenChainId    uint64                     `json:"token_chain_id" `
	TokenImage      string                     `json:"token_image" bson:"token_image"`
	Ticker          string                     `json:"ticker" bson:"ticker"`
	TokenName       string                     `json:"token_name" bson:"token_name"`
	TokenAddress    string                     `bson:"token_address" json:"token_address"`
	CreateTokenMode models.CreateTokenModeType `bson:"create_token_mode" json:"create_token_mode"`

	Status models.AssistantStatus `bson:"status" json:"status"`

	SystemReminder string `bson:"system_reminder" json:"system_reminder"`
	AgentBaseModel string `bson:"open_ai_assistant_model" json:"agent_base_model"`

	SystemContent string `bson:"system_content" json:"system_content"`

	TwitterID       string `bson:"twitter_id" json:"twitter_id"`
	TwitterName     string `bson:"twitter_name" json:"twitter_name"`
	TwitterUserName string `bson:"twitter_username" json:"twitter_username"`
	TwitterAvatar   string `bson:"twitter_avatar" json:"twitter_avatar"`

	AssistantCharacter `bson:",inline"`

	VerifiedNFTOwner   bool   `json:"verified_nft_owner"`
	NFTAddress         string `json:"nft_address"`
	NFTTokenID         string `json:"nft_token_id"`
	NFTTokenImage      string `json:"nft_token_image"`
	NFTOwnerAddress    string `json:"nft_owner_address"`
	NFTSignature       string `json:"nft_signature" bson:"nft_signature"`
	NFTSignMessage     string `json:"nft_sign_message" bson:"nft_sign_message"`
	NFTDelegateAddress string `json:"nft_delegate_address" bson:"nft_delegate_address"`
	NFTPublicKey       string `json:"nft_public_key" bson:"nft_public_key"`

	SocialInfo []*models.SocialInfo `json:"social_info" bson:"social_info"`

	TwinTwitterUsernames string  `json:"twin_twitter_usernames"`
	TwinStatus           string  `json:"twin_status"`
	TwinTrainingProgress float64 `json:"twin_training_progress"`
	TwinTrainingMessage  string  `json:"twin_training_message"`

	GraphData string `json:"graph_data"`
}

type AssistantResp struct {
	Assistants    `json:",inline"`
	AgentInfo     *AgentInfoResp `json:"agent_info"`
	KnowledgeBase *KnowledgeBase `json:"knowledge_base"`
}

func NewAssistantResp(m *models.AgentInfo) *AssistantResp {
	if m == nil {
		return nil
	}
	resp := &AssistantResp{}
	resp.ID = m.AgentID
	resp.CreatedAt = m.CreatedAt
	resp.InferFee = m.InferFee.String()
	resp.AgentName = m.AgentName
	resp.Creator = m.Creator
	resp.Minter = m.Minter
	resp.ContractAgentID = m.AgentContractID
	resp.AgentContractAddress = m.AgentContractAddress
	resp.ChainID = m.NetworkID
	resp.TxHash = m.MintHash
	resp.Uri = m.Uri
	resp.Thumbnail = m.Thumbnail
	resp.TokenAddress = m.TokenAddress
	resp.TokenChainId = m.TokenNetworkID
	resp.TokenImage = m.TokenImageUrl
	resp.TokenName = m.TokenName
	resp.Ticker = m.TokenSymbol
	resp.AgentBaseModel = m.AgentBaseModel
	resp.Bio = m.GetCharacterArrayString(m.Bio)
	resp.Lore = m.GetCharacterArrayString(m.Lore)
	resp.Knowledge = m.GetCharacterArrayString(m.Knowledge)
	resp.PostExamples = m.GetCharacterArrayString(m.PostExamples)
	resp.Topics = m.GetCharacterArrayString(m.Topics)
	resp.Adjectives = m.GetCharacterArrayString(m.Adjectives)
	resp.MessageExamples = m.GetMessageExamples()
	resp.Style = m.GetStyle()
	resp.SocialInfo = m.GetSocialInfo()
	resp.Status = m.Status
	resp.SystemReminder = m.SystemReminder
	resp.MetaData = m.MetaData
	resp.SystemContent = m.SystemPrompt
	resp.VerifiedNFTOwner = m.VerifiedNftOwner
	resp.NFTAddress = m.NftAddress
	resp.NFTTokenID = m.NftTokenID
	resp.NFTTokenImage = m.NftTokenImage
	resp.NFTOwnerAddress = m.NftOwnerAddress
	resp.NFTSignature = m.NftSignature
	resp.NFTSignMessage = m.NftSignMessage
	resp.NFTDelegateAddress = m.NftDelegateAddress
	resp.NFTPublicKey = m.NftPublicKey
	resp.CreateTokenMode = models.CreateTokenModeType(m.TokenMode)
	resp.TwinTwitterUsernames = m.TwinTwitterUsernames
	resp.TwinStatus = string(m.TwinStatus)
	resp.TwinTrainingProgress = m.TwinTrainingProgress
	resp.TwinTrainingMessage = m.TwinTrainingMessage
	resp.GraphData = m.GraphData

	if m.TwitterInfo != nil {
		resp.TwitterID = m.TwitterInfo.TwitterID
		resp.TwitterName = m.TwitterInfo.TwitterName
		resp.TwitterUserName = m.TwitterInfo.TwitterUsername
		resp.TwitterAvatar = m.TwitterInfo.TwitterAvatar
	}

	if m.TokenMode == "" {
		resp.CreateTokenMode = models.CreateTokenModeTypeNoToken
	}

	resp.AgentInfo = NewAgentInfoResp(m)
	resp.KnowledgeBase = NewKnowledgeBaseResp(m.KnowledgeBase)
	return resp
}

func NewAssistantRespArry(arr []*models.AgentInfo) []*AssistantResp {
	resps := []*AssistantResp{}
	for _, m := range arr {
		resps = append(resps, NewAssistantResp(m))
	}
	return resps
}

type DataUploadToLightHouse struct {
	Content string `json:"content"`
}

type StudioReq struct {
	GraphData string `json:"graph_data"`
}

type AgentActionReq struct {
	Action          string   `json:"action"`
	Ids             []uint   `json:"ids"`
	ContractAddress []string `json:"contract_address"`
}

type AgentCommentReq struct {
	Comment string  `json:"comment"`
	Rating  float64 `json:"rating"`
}
