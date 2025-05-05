package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type InfraTwitterApp struct {
	gorm.Model
	Address       string
	TwitterInfoID uint
	TwitterInfo   *TwitterInfo
	EaiBalance    numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TotalRequest  uint64
	RemainRequest uint64
	ETHAddress    string `gorm:"index"`
}

type InfraTwitterTopupTx struct {
	gorm.Model
	NetworkID         uint64
	EventId           string `gorm:"unique_index"`
	InfraTwitterAppID uint   `gorm:"index"`
	InfraTwitterApp   *InfraTwitterApp
	Type              AgentEaiTopupType `gorm:"default:'deposit'"`
	DepositAddress    string
	ToAddress         string
	TxHash            string
	Amount            numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	Status            AgentEaiTopupStatus
}
