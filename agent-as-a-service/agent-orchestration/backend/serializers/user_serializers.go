package serializers

import (
	"encoding/json"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

type UserProfileReq struct {
	Username    string                 `json:"username"`
	Description string                 `json:"description"`
	ImageURL    string                 `json:"image_url"`
	Social      map[string]interface{} `json:"social"`
}

type UserResp struct {
	ID              uint                   `json:"id"`
	CreatedAt       time.Time              `json:"created_at"`
	NetworkID       uint64                 `json:"network_id"`
	Address         string                 `json:"address"`
	Username        string                 `json:"username"`
	Description     string                 `json:"description"`
	ImageURL        string                 `json:"image_url"`
	Social          map[string]interface{} `json:"social"`
	TwitterID       string                 `json:"twitter_id"`
	TwitterAvatar   string                 `json:"twitter_avatar"`
	TwitterUsername string                 `json:"twitter_username"`
	TwitterName     string                 `json:"twitter_name"`
}

func NewUserResp(m *models.User) *UserResp {
	if m == nil {
		return nil
	}

	social := make(map[string]interface{})
	json.Unmarshal([]byte(m.Social), &social)

	resp := &UserResp{
		ID:              m.ID,
		CreatedAt:       m.CreatedAt,
		NetworkID:       m.NetworkID,
		Address:         m.Address,
		Username:        m.Username,
		Description:     m.Description,
		ImageURL:        m.ImageURL,
		Social:          social,
		TwitterID:       m.TwitterID,
		TwitterAvatar:   m.TwitterAvatar,
		TwitterUsername: m.TwitterUsername,
		TwitterName:     m.TwitterName,
	}
	return resp
}

func NewUserRespArr(arr []*models.User) []*UserResp {
	resps := []*UserResp{}
	for _, m := range arr {
		resps = append(resps, NewUserResp(m))
	}
	return resps
}

type InfraTwitterAppInfoResp struct {
	Address       string  `json:"address"`
	EaiBalance    float64 `json:"eai_balance"`
	TotalRequest  uint64  `json:"total_request"`
	RemainRequest uint64  `json:"remain_request"`
	ETHAddress    string  `json:"eth_address"`
}

func NewInfraTwitterAppInfoResp(m *models.InfraTwitterApp) *InfraTwitterAppInfoResp {
	if m == nil {
		return nil
	}
	balance, _ := m.EaiBalance.Float64()
	return &InfraTwitterAppInfoResp{
		Address:       m.Address,
		EaiBalance:    balance,
		TotalRequest:  m.TotalRequest,
		RemainRequest: m.RemainRequest,
		ETHAddress:    m.ETHAddress,
	}
}
