package serializers

import (
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type RobotSaleWalletReq struct {
	ProjectID   string `json:"project_id"`
	UserAddress string `json:"user_address"`
}

type RobotTokenTransferReq struct {
	ProjectID       string  `json:"project_id"`
	ReceiverAddress string  `json:"receiver_address"`
	Amount          float64 `json:"amount"`
}

type RobotSaleWalletResp struct {
	ProjectID   string           `json:"project_id"`
	UserAddress string           `json:"user_address"`
	SOLAddress  string           `json:"sol_address"`
	SolBalance  numeric.BigFloat `json:"sol_balance"`
	Ranking     int              `json:"ranking"`
}

func NewRobotSaleWalletResp(m *models.RobotSaleWallet) *RobotSaleWalletResp {
	return &RobotSaleWalletResp{
		ProjectID:   m.ProjectID,
		UserAddress: m.UserAddress,
		SOLAddress:  m.SOLAddress,
		SolBalance:  m.SOLBalance,
		Ranking:     m.Ranking,
	}
}

func NewRobotSaleWalletRespList(lst []*models.RobotSaleWallet) []*RobotSaleWalletResp {
	ret := []*RobotSaleWalletResp{}
	for _, m := range lst {
		ret = append(ret, NewRobotSaleWalletResp(m))
	}
	return ret
}

type RobotProjectResp struct {
	ProjectID    string           `json:"project_id"`
	TokenAddress string           `json:"token_address"`
	TokenSymbol  string           `json:"token_symbol"`
	TokenName    string           `json:"token_name"`
	TokenSupply  numeric.BigFloat `json:"token_supply"`
	TotalBalance numeric.BigFloat `json:"sol_balance"`
	SolPrice     *big.Float       `json:"sol_price"`
}

func NewRobotProjectResp(m *models.RobotProject) *RobotProjectResp {
	return &RobotProjectResp{
		ProjectID:    m.ProjectID,
		TokenAddress: m.TokenAddress,
		TokenSymbol:  m.TokenSymbol,
		TokenName:    m.TokenName,
		TokenSupply:  m.TokenSupply,
		TotalBalance: m.TotalBalance,
		SolPrice:     m.SolPrice,
	}
}

type RobotTokenTransferResp struct {
	ProjectID       string           `json:"project_id"`
	ReceiverAddress string           `json:"receiver_address"`
	Amount          numeric.BigFloat `json:"amount"`
	Status          string           `json:"status"`
	TxHash          string           `json:"tx_hash"`
	TransferAt      *time.Time       `json:"transfer_at"`
}

func NewRobotTokenTransferResp(m *models.RobotTokenTransfer) *RobotTokenTransferResp {
	return &RobotTokenTransferResp{
		ProjectID:       m.ProjectID,
		ReceiverAddress: m.ReceiverAddress,
		Amount:          m.Amount,
		Status:          m.Status,
		TxHash:          m.TxHash,
		TransferAt:      m.TransferAt,
	}
}
