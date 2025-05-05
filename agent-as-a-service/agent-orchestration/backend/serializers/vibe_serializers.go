package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

type AgentUserComment struct {
	ID          uint      `json:"id"`
	Comment     string    `json:"comment"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserAddress string    `json:"user_address"`
	User        *UserResp `json:"user"`
	AgentInfoID uint      `json:"agent_info_id"`
}

func NewAgentUserComment(m *models.AgentUserComment) *AgentUserComment {
	return &AgentUserComment{
		ID:          m.ID,
		Comment:     m.Comment,
		Rating:      m.Rating,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		UserAddress: m.UserAddress,
		AgentInfoID: m.AgentInfoID,
		User:        NewUserResp(m.User),
	}
}

func NewAgentUserCommentArray(arr []*models.AgentUserComment) []*AgentUserComment {
	resps := []*AgentUserComment{}
	for _, r := range arr {
		resps = append(resps, NewAgentUserComment(r))
	}
	return resps
}

type VibeTokenDeployInfoResp struct {
	NonceHex  string `json:"nonce_hex" swaggertype:"string" example:"1"`
	Name      string `json:"name" swaggertype:"string" example:"My Token"`
	Symbol    string `json:"symbol" swaggertype:"string" example:"MYT"`
	Creator   string `json:"creator" swaggertype:"string" example:"0x123...abc"`
	LowerTick int64  `json:"lower_tick" swaggertype:"integer" example:"1000"`
	UpperTick int64  `json:"upper_tick" swaggertype:"integer" example:"2000"`
	Deadline  int64  `json:"deadline" swaggertype:"integer" example:"1716153600"`
	Signature string `json:"signature" swaggertype:"string" example:"0x123...abc"`
}
