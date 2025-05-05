package models

import (
	"math/big"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type RobotProject struct {
	gorm.Model
	ProjectID     string `gorm:"unique_index"`
	ScanEnabled   bool   `gorm:"default:0"`
	TokenAddress  string `gorm:"index"`
	TokenSymbol   string
	TokenName     string
	TokenImageUrl string
	TokenSupply   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	MintHash      string
	Signature     string
	TotalBalance  numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`

	SolPrice *big.Float `gorm:"-"`
}

type RobotSaleWallet struct {
	gorm.Model
	ProjectID         string           `gorm:"unique_index:robot_sale_wallet_main_idx"`
	UserAddress       string           `gorm:"unique_index:robot_sale_wallet_main_idx"`
	SOLAddress        string           `gorm:"index"`
	SOLBalance        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SOLOnchainBalance numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	SOLMovedBalance   numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	IsSOLTransferring bool             `gorm:"default:0"`
	SOLRequestAt      *time.Time       `gorm:"index"`
	SOLScanAt         *time.Time       `gorm:"index"`
	Ranking           int              `gorm:"default:0"`
}

type RobotTokenTransfer struct {
	gorm.Model
	ProjectID       string           `gorm:"index"`
	ReceiverAddress string           `gorm:"index"`
	Amount          numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TransferAt      *time.Time       `gorm:"index"`
	TxHash          string
	Status          string
	Error           string
}
