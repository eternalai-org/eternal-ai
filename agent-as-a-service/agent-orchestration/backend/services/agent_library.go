package services

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetListAgentLibrary(ctx context.Context, agentType int, networkID uint64) ([]*models.AgentLibrary, error) {
	filter := map[string][]interface{}{
		"network_id = ? ": {networkID},
	}
	if agentType > 0 {
		filter["agent_type = ? "] = []interface{}{agentType}
	}
	res, err := s.dao.FindAgentLibrary(
		daos.GetDBMainCtx(ctx),
		filter,
		map[string][]interface{}{},
		[]string{"id desc"},
		0,
		1000,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) SaveAgentLibrary(ctx context.Context, networkID uint64, req *serializers.AgentLibraryReq) (*models.AgentLibrary, error) {
	var agentLibrary *models.AgentLibrary
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {

			agentLibrary := &models.AgentLibrary{
				NetworkID: networkID,
				Name:      req.Name,
				SourceURL: req.SourceUrl,
				AgentType: models.AgentInfoAgentType(req.AgentType),
			}
			return s.dao.Create(tx, agentLibrary)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return agentLibrary, nil
}
