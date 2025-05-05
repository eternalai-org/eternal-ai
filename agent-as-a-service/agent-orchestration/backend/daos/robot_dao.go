package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstRobotSaleWallet(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.RobotSaleWallet, error) {
	var m models.RobotSaleWallet
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindRobotSaleWallet(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.RobotSaleWallet, error) {
	var ms []*models.RobotSaleWallet
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindRobotSaleWallet4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.RobotSaleWallet, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.RobotSaleWallet
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindRobotSaleWalletJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.RobotSaleWallet, error) {
	var ms []*models.RobotSaleWallet
	err := d.findJoinSelect(tx, &models.RobotSaleWallet{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstRobotSaleWalletByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.RobotSaleWallet, error) {
	var m models.RobotSaleWallet
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstRobotProject(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.RobotProject, error) {
	var m models.RobotProject
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstRobotProjectByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.RobotProject, error) {
	var m models.RobotProject
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindRobotProject(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.RobotProject, error) {
	var ms []*models.RobotProject
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) UpdatePrijectTotalBalance(tx *gorm.DB, projectID string) error {
	err := tx.Exec(`
			update robot_projects 
			join (
				select project_id, sum(sol_balance) sol_balance
				from robot_sale_wallets rsw 
				where project_id = ?
				group by project_id 
			) tmp on robot_projects.project_id = tmp.project_id
			set total_balance= tmp.sol_balance
			where robot_projects.project_id = ?
	`, projectID, projectID).Error
	return err
}

func (d *DAO) UpdateRobotProjectRanking(tx *gorm.DB, projectID string) error {
	_ = tx.Exec(`
		UPDATE robot_sale_wallets AS rsw
		JOIN (
			SELECT id,
				RANK() OVER (ORDER BY sol_balance DESC, id) AS ranking
			FROM robot_sale_wallets
			WHERE sol_balance > 0
			AND project_id = ?
		) AS tmp ON rsw.id = tmp.id
		SET rsw.ranking = tmp.ranking
		WHERE rsw.sol_balance > 0
		AND rsw.project_id = ?
	`, projectID, projectID).Error
	return nil
}
