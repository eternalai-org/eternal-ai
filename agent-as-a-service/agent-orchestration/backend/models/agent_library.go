package models

import (
	"github.com/jinzhu/gorm"
)

type AgentLibrary struct {
	gorm.Model
	NetworkID uint64 `gorm:"unique_index:agent_library_main_idx"`
	Name      string `gorm:"unique_index:agent_library_main_idx"`
	SourceURL string
	AgentType AgentInfoAgentType
}
