package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

type AgentLibraryReq struct {
	Name      string `json:"name"`
	SourceUrl string `json:"source_url"`
	AgentType int    `json:"agent_type"`
}

type AgentLibraryResp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	SourceUrl string    `json:"source_url"`
	AgentType int       `json:"agent_type"`
}

func NewAgentLibraryResp(m *models.AgentLibrary) *AgentLibraryResp {
	if m == nil {
		return nil
	}
	return &AgentLibraryResp{
		ID:        m.ID,
		CreatedAt: m.CreatedAt,
		Name:      m.Name,
		SourceUrl: m.SourceURL,
		AgentType: int(m.AgentType),
	}
}

func NewAgentLibraryRespArray(arr []*models.AgentLibrary) []*AgentLibraryResp {
	resps := []*AgentLibraryResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentLibraryResp(r))
	}
	return resps
}
