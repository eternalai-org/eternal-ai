package ports

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
)

type IKnowledgeUsecase interface {
	CreateAgentInfoKnowledgeBase(ctx context.Context, models []*models.AgentInfoKnowledgeBase, agentInfoId uint) ([]*models.AgentInfoKnowledgeBase, error)
	GetAgentInfoKnowledgeBaseByAgentId(context.Context, uint) (*models.AgentInfoKnowledgeBase, error)
	GetKnowledgeBaseById(context.Context, uint) (*models.KnowledgeBase, error)
	DeleteKnowledgeBaseById(context.Context, uint) error
	CreateKnowledgeBase(context.Context, *serializers.CreateKnowledgeRequest) (*serializers.KnowledgeBase, error)
	ListKnowledgeBase(context.Context, *models.ListKnowledgeBaseRequest) ([]*serializers.KnowledgeBase, error)
	WatchWalletChange(context.Context) error
	GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
	UpdateListKnowledgeBaseFile(ctx context.Context, kbId uint, files []*serializers.File) (bool, error)
	Webhook(context.Context, *models.RagHookResponse) (*models.KnowledgeBase, error)
	WebhookFile(context.Context, string, []byte, uint) (*models.KnowledgeBase, error)
	MapKnowledgeBaseByAgentIds(ctx context.Context, ids []uint) (map[uint][]*models.KnowledgeBase, error)
	GetKnowledgeBaseByKBId(context.Context, string) (*models.KnowledgeBase, error)
	GetKnowledgeBaseByKBTokenId(context.Context, string) (*models.KnowledgeBase, error)
	GetKBAgentsUsedOfSocialAgent(ctx context.Context, socialAgentId uint) ([]*models.KnowledgeBase, error)
	GetManyKnowledgeBaseByQuery(context.Context, string, string, int, int) ([]*models.KnowledgeBase, error)
	SendMessage(_ context.Context, content string, chanId int64) (int, error)
	CalcFeeByKnowledgeBaseId(ctx context.Context, kbId uint) (float64, error)
	CheckBalance(ctx context.Context, kn *models.KnowledgeBase) error
	ScanKnowledgeBaseStatusPaymentReceipt(ctx context.Context)
	ExecCrawlData(ctx context.Context, kbId uint) error
}
