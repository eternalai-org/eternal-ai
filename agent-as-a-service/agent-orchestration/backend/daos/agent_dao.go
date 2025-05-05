package daos

import (
	"strconv"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

// ////
func (d *DAO) FindTwitterInfo(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.TwitterInfo, error) {
	var ms []*models.TwitterInfo
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FirstTwitterInfoByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.TwitterInfo, error) {
	var m models.TwitterInfo
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstTwitterInfo(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.TwitterInfo, error) {
	var m models.TwitterInfo
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstUserTwitterPost(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.UserTwitterPost, error) {
	var m models.UserTwitterPost
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindUserTwitterPost(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.UserTwitterPost, error) {
	var ms []*models.UserTwitterPost
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindAgentTokenJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentTokenInfo, error) {
	var ms []*models.AgentTokenInfo
	err := d.findJoinSelect(tx, &models.AgentTokenInfo{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstAgentToken(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.AgentTokenInfo, error) {
	var m models.AgentTokenInfo
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentTokenByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.AgentTokenInfo, error) {
	var m models.AgentTokenInfo
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentTradeHistoryByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.AgentTradeHistory, error) {
	var m models.AgentTradeHistory
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentTradeHistory(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.AgentTradeHistory, error) {
	var m models.AgentTradeHistory
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) GetAgentTradeHistoryInfo(tx *gorm.DB, agentTokenID uint, networkID uint64, tokenAddress string) (*models.AgentTokenInfo, error) {
	var rs models.AgentTokenInfo
	query := tx.Raw(`
		select 
			ifnull(p.price, 0) price,
			ifnull(p.tick, 0) tick,
			ifnull(p_24h.price_last24h, 0) price_last24h,
			ifnull(v_all.total_volume, 0) total_volume,
			ifnull(v.volume_last24h, 0) volume_last24h,
			ifnull(holders.holders, 0) holders
		from (
			select f.price, f.tick
			from agent_trade_histories f  
			where f.agent_token_id = ?
			order by tx_at desc limit 1
		) p,
		(
			select cast(ifnull(sum(base_amount), 0) as decimal(36,18)) total_volume
			from agent_trade_histories f  
			where f.agent_token_id = ?
		) v_all,
		(
			select 
				cast(ifnull(sum(base_amount), 0) as decimal(36,18)) volume_last24h
			from agent_trade_histories ob 
			where 1=1
			and agent_token_id = ?
			and tx_at >= now() - INTERVAL 1 DAY 
		) v,
		(
			select f.price price_last24h
			from agent_trade_histories f  
			where f.agent_token_id = ?
			and tx_at <= now() - INTERVAL 1 DAY 
			order by tx_at desc limit 1
		) p_24h,
		(
			select count(distinct address) holders
			from erc20_holders
			where network_id = ?
			and contract_address = ?
			and cast(balance as decimal(36, 18)) >= 0.0000001
		) holders
	`, agentTokenID, agentTokenID, agentTokenID, agentTokenID, networkID, tokenAddress)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return &rs, nil
}

func (d *DAO) GetAgentSocialInfo(tx *gorm.DB, agentID uint) (*models.AgentTokenInfo, error) {
	var rs models.AgentTokenInfo
	query := tx.Raw(`
		select 
			ifnull(mentions.mentions, 0) mentions,
			ifnull(tip.amount, 0) tip_amount
		from (
			select count(1) mentions
			from agent_twitter_posts atp 
			where agent_info_id= ?
		) mentions,
		(
			select sum(amount) amount
			from agent_eai_topups 
			where agent_info_id= ?
			and status='done'
		) tip
	`, agentID, agentID)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return &rs, nil
}

func (d *DAO) GetAgentSummaryReport(tx *gorm.DB, hiddenNetworkId string) ([]*models.AgentInfo, error) {
	var rs []*models.AgentInfo
	hideNetwork := []int64{43338, 222672, 0}
	if hiddenNetworkId != "" {
		hiddenNetworkIdArry := strings.Split(hiddenNetworkId, ",")
		for _, s := range hiddenNetworkIdArry {
			i64, _ := strconv.ParseInt(s, 10, 64)
			hideNetwork = append(hideNetwork, i64)
		}
	}
	query := tx.Raw(`
		select network_id, network_name,  count(1) counts
		from agent_infos
		where 1=1
		and network_id not in (?)
		and deleted_at is null
		group by network_id, network_name
		order by counts desc
	`, hideNetwork)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

func (d *DAO) FirstAbilityLuckyMoney(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.AbilityLuckyMoney, error) {
	var m models.AbilityLuckyMoney
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindAbilityLuckyMoney(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.AbilityLuckyMoney, error) {
	var ms []*models.AbilityLuckyMoney
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FirstAbilityLuckyMoneyByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.AbilityLuckyMoney, error) {
	var m models.AbilityLuckyMoney
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentReactionHistory(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.AgentReactionHistory, error) {
	var m models.AgentReactionHistory
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
