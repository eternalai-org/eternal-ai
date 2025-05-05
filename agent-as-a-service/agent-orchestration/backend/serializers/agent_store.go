package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AgentStoreReq struct {
	ID          uint                    `json:"id"`
	Type        models.AgentStoreType   `json:"type"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	AuthenUrl   string                  `json:"authen_url"`
	ApiUrl      string                  `json:"api_url"`
	Icon        string                  `json:"icon"`
	Docs        string                  `json:"docs"`
	Price       numeric.BigFloat        `json:"price"`
	Status      models.AgentStoreStatus `json:"status"`
}

type AgentStoreMissionReq struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Prompt      string           `json:"prompt"`
	Price       numeric.BigFloat `json:"price"`
	ToolList    string           `json:"tool_list"`
	Icon        string           `json:"icon"`
	NetworkID   uint64           `json:"network_id"`
	Model       string           `json:"model"`
	Status      string           `json:"status"`
}

type AuthenAgentStoreCallback struct {
	Code           string            `json:"code"`
	CallbackParams map[string]string `json:"callback_params"`
}

type AgentStoreResp struct {
	ID                 uint                     `json:"id"`
	CreatedAt          time.Time                `json:"created_at"`
	Name               string                   `json:"name"`
	Owner              string                   `json:"owner"`
	Description        string                   `json:"description"`
	AuthenUrl          string                   `json:"authen_url"`
	Icon               string                   `json:"icon"`
	Docs               string                   `json:"docs"`
	ApiUrl             string                   `json:"api_url"`
	Status             models.AgentStoreStatus  `json:"status"`
	Price              numeric.BigFloat         `json:"price"`
	NumInstall         uint                     `json:"num_install"`
	NumUsage           uint                     `json:"num_usage"`
	Type               string                   `json:"type"`
	AgentStoreMissions []*AgentStoreMissionResp `json:"agent_store_missions"`
}

type AgentStoreTryDetailResp struct {
	HistoryID         uint                   `json:"history_id"`
	DetailID          uint                   `json:"detail_id"`
	FromUser          bool                   `json:"from_user"`
	Content           string                 `json:"content"`
	AgentSnapshotPost *AgentSnapshotPostResp `json:"id"`
	CreatedAt         time.Time              `json:"created_at"`
}

func NewAgentStoreResp(m *models.AgentStore) *AgentStoreResp {
	if m == nil {
		return nil
	}
	return &AgentStoreResp{
		ID:                 m.ID,
		CreatedAt:          m.CreatedAt,
		Name:               m.Name,
		Description:        m.Description,
		AuthenUrl:          m.AuthenUrl,
		Icon:               m.Icon,
		Docs:               m.Docs,
		ApiUrl:             m.ApiUrl,
		NumInstall:         m.NumInstall,
		Status:             m.Status,
		Price:              m.Price,
		Owner:              m.OwnerAddress,
		Type:               string(m.Type),
		NumUsage:           m.NumUsage,
		AgentStoreMissions: NewAgentStoreMissionRespArray(m.AgentStoreMissions),
	}
}

func NewAgentStoreRespArray(arr []*models.AgentStore) []*AgentStoreResp {
	resps := []*AgentStoreResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreResp(r))
	}
	return resps
}

type AgentStoreMissionResp struct {
	AgentStoreID uint             `json:"agent_store_id"`
	ID           uint             `json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	UserPrompt   string           `json:"user_prompt"`
	Price        numeric.BigFloat `json:"price"`
	ToolList     string           `json:"tool_list"`
	Icon         string           `json:"icon"`
	NumUsed      uint             `json:"num_used"`
	Status       string           `json:"status"`
}

func NewAgentStoreMissionResp(m *models.AgentStoreMission) *AgentStoreMissionResp {
	if m == nil {
		return nil
	}
	return &AgentStoreMissionResp{
		AgentStoreID: m.AgentStoreID,
		ID:           m.ID,
		CreatedAt:    m.CreatedAt,
		Name:         m.Name,
		Description:  m.Description,
		UserPrompt:   m.UserPrompt,
		Price:        m.Price,
		ToolList:     m.ToolList,
		Icon:         m.Icon,
		NumUsed:      m.NumUsed,
		Status:       string(m.Status),
	}
}
func NewAgentStoreMissionRespArray(arr []*models.AgentStoreMission) []*AgentStoreMissionResp {
	resps := []*AgentStoreMissionResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreMissionResp(r))
	}
	return resps
}

func NewAgentStoreRespArrayFromInstall(arr []*models.AgentStoreInstall) []*AgentStoreResp {
	resps := []*AgentStoreResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreResp(r.AgentStore))
	}
	return resps
}

func NewAgentStoreTryDetailResp(m *models.AgentStoreTryDetail) *AgentStoreTryDetailResp {
	if m == nil {
		return nil
	}
	return &AgentStoreTryDetailResp{
		CreatedAt:         m.CreatedAt,
		HistoryID:         m.AgentStoreTryID,
		DetailID:          m.ID,
		FromUser:          m.FromUser,
		Content:           m.Content,
		AgentSnapshotPost: NewAgentSnapshotPostResp(m.AgentSnapshotPost),
	}
}

func NewAgentStoreTryDetailRespArray(arr []*models.AgentStoreTryDetail) []*AgentStoreTryDetailResp {
	resps := []*AgentStoreTryDetailResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreTryDetailResp(r))
	}
	return resps
}

// ////
type AgentUtilityTwitterReq struct {
	AgentID string `json:"agent_id"`
	Content string `json:"content"`
	TxHash  string `json:"tx_hash"`
}

type AgentUtilityTwitterResp struct {
	Message string `json:"message"`
	AuthUrl string `json:"auth_url"`
}
