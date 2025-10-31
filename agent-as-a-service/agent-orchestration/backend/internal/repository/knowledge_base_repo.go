package repository

import (
	"context"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type knowledgeBaseRepo struct {
	db *gorm.DB
}

type KnowledgeBaseRepo interface {
	GetById(ctx context.Context, id uint) (*models.KnowledgeBase, error)
	DeleteById(ctx context.Context, id uint) error
	Create(ctx context.Context, model *models.KnowledgeBase) (*models.KnowledgeBase, error)
	GetByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	GetByStatuses(ctx context.Context, statuses []models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	List(ctx context.Context, req *models.ListKnowledgeBaseRequest) ([]*models.KnowledgeBase, error)
	UpdateStatus(ctx context.Context, model *models.KnowledgeBase) error
	UpdateById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
	GetByKBId(context.Context, string) (*models.KnowledgeBase, error)
	GetByKBTokenId(context.Context, string) (*models.KnowledgeBase, error)
	GetManyByQuery(ctx context.Context, query string, orderOption string, offset int, limit int) ([]*models.KnowledgeBase, error)
	GetKBAgentsUsedOfSocialAgent(ctx context.Context, socialAgentId uint) ([]*models.KnowledgeBase, error)
}

func (r *knowledgeBaseRepo) GetKBAgentsUsedOfSocialAgent(ctx context.Context, socialAgentId uint) ([]*models.KnowledgeBase, error) {
	knowledge := []*models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("AgentInfo").
		Where("kb_id <> '' and id IN (SELECT knowledge_base_id FROM agent_info_knowledge_bases WHERE agent_info_id = ?)",
			socialAgentId).
		Find(&knowledge).Error
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}

func (r *knowledgeBaseRepo) UpdateStatus(ctx context.Context, model *models.KnowledgeBase) error {
	return r.db.Model(model).Update("status", model.Status).Error
}

func (r *knowledgeBaseRepo) List(ctx context.Context, req *models.ListKnowledgeBaseRequest) ([]*models.KnowledgeBase, error) {
	knowledges := []*models.KnowledgeBase{}
	query := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles")

	if req.UserAddress != "" {
		query = query.Where("user_address = ?", req.UserAddress)
	}

	if len(req.Statuses) != 0 {
		query = query.Where("status IN (?)", req.Statuses)
	}

	if err := query.Find(&knowledges).Error; err != nil {
		return nil, err
	}
	return knowledges, nil
}

func (r *knowledgeBaseRepo) Create(ctx context.Context, model *models.KnowledgeBase) (*models.KnowledgeBase, error) {
	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (r *knowledgeBaseRepo) GetById(ctx context.Context, id uint) (*models.KnowledgeBase, error) {
	knowledge := &models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("id = ?", id).
		First(knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func (r *knowledgeBaseRepo) DeleteById(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx)
	return tx.Transaction(func(tx *gorm.DB) error {
		data := []*models.KnowledgeBaseFile{}
		if err := tx.Where("knowledge_base_id = ?", id).Delete(&data).Error; err != nil {
			return err
		}

		return tx.Delete(&models.KnowledgeBase{}, id).Error
	})
}

func (r *knowledgeBaseRepo) GetByStatuses(ctx context.Context, statuses []models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error) {
	var data []*models.KnowledgeBase
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("status IN (?)", statuses).
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *knowledgeBaseRepo) GetByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error) {
	var data []*models.KnowledgeBase
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("status = ?", status).
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *knowledgeBaseRepo) UpdateById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).Model(&models.KnowledgeBase{}).Where("id = ?", id).Updates(updatedFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *knowledgeBaseRepo) GetByKBId(ctx context.Context, kbId string) (*models.KnowledgeBase, error) {
	knowledge := &models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("AgentInfo").
		Preload("KnowledgeBaseFiles").
		Where("kb_id = ? ", kbId).
		First(knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func (r *knowledgeBaseRepo) GetByKBTokenId(ctx context.Context, kbTokenId string) (*models.KnowledgeBase, error) {
	knowledge := &models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("kb_token_id = ? ", kbTokenId).
		First(knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func (r *knowledgeBaseRepo) GetManyByQuery(ctx context.Context, query string, orderOption string, offset int, limit int) ([]*models.KnowledgeBase, error) {
	if len(query) == 0 {
		return nil, fmt.Errorf("query is empty")
	}
	if orderOption == "" {
		orderOption = "id asc"
	}
	var data []*models.KnowledgeBase
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where(query).
		Order("updated_at").
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewKnowledgeBaseRepository(db *gorm.DB) KnowledgeBaseRepo {
	return &knowledgeBaseRepo{db}
}
