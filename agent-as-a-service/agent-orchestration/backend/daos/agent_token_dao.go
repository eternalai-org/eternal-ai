package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstMemeByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.Meme, error) {
	var m models.Meme
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMeme(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.Meme, error) {
	var m models.Meme
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMeme(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.Meme, error) {
	var ms []*models.Meme
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindMeme4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.Meme, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.Meme
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.Meme{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FindMemeJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.Meme, error) {
	var ms []*models.Meme
	err := d.findJoinSelect(tx, &models.Meme{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindMemeJoin(tx *gorm.DB, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.Meme, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.Meme
	err := d.findJoin(tx, &ms, joins, filters, preloads, orders, offset, limit, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.Meme, error) {
	var m models.Meme
	if err := d.firstJoinSelect(tx, &m, selected, joins, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMemeJoinSelect4Page(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.Meme, uint, error) {
	var ms []*models.Meme
	c, err := d.findJoinSelect4PageNoCount(tx, &models.Meme{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

// //
func (d *DAO) FindMemeHistoryJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MemeTradeHistory, error) {
	var ms []*models.MemeTradeHistory
	err := d.findJoinSelect(tx, &models.MemeTradeHistory{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeHistoryByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.MemeTradeHistory, error) {
	var m models.MemeTradeHistory
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeHistory(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeTradeHistory, error) {
	var m models.MemeTradeHistory
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMemeHistory(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.MemeTradeHistory, error) {
	var ms []*models.MemeTradeHistory
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) GetMemeTradeHistoryInfo(tx *gorm.DB, pairID uint) (*models.Meme, error) {
	var rs models.Meme
	query := tx.Raw(`
		select 
			ifnull(p.price, 0) price,
			ifnull(p.tick, 0) tick,
			ifnull(p_24h.price_last24h, 0) price_last24h,
			ifnull(v_all.total_volume, 0) total_volume,
			v.*
		from (
			select f.price, f.tick
			from meme_trade_histories f  
			where f.meme_id = ?
			order by tx_at desc limit 1
		) p 
		left join (
			select cast(ifnull(sum(base_amount), 0) as decimal(36,18)) total_volume
			from meme_trade_histories f  
			where f.meme_id = ?
		) v_all on 1=1
		left join (
			select 
				cast(ifnull(sum(base_amount), 0) as decimal(36,18)) volume_last24h
			from meme_trade_histories ob 
			where 1=1
			and meme_id = ?
			and tx_at >= now() - INTERVAL 1 DAY 
		) v on 1=1
		left join (
			select f.price price_last24h
			from meme_trade_histories f  
			where f.meme_id = ?
			and tx_at <= now() - INTERVAL 1 DAY 
			order by tx_at desc limit 1
		) p_24h on 1=1
	`, pairID, pairID, pairID, pairID)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return &rs, nil
}

func (d *DAO) UpdateChartCandleDataByMemeID(tx *gorm.DB, memID uint) (bool, error) {
	if err := tx.Exec("call create_chart_candles_min30_by_key(?)", memID).Error; err != nil {
		return false, errs.NewError(err)
	}

	if err := tx.Exec("call create_chart_candles_min5_by_key(?)", memID).Error; err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

// //thread
// //
func (d *DAO) FindMemeThreadJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MemeThreads, error) {
	var ms []*models.MemeThreads
	err := d.findJoinSelect(tx, &models.MemeThreads{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeThreadByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.MemeThreads, error) {
	var m models.MemeThreads
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeThread(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeThreads, error) {
	var m models.MemeThreads
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMemeThread(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.MemeThreads, error) {
	var ms []*models.MemeThreads
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FirstMemeThreadLike(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeThreadLike, error) {
	var m models.MemeThreadLike
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// /followers
func (d *DAO) FindMemeFollowersJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MemeFollowers, error) {
	var ms []*models.MemeFollowers
	err := d.findJoinSelect(tx, &models.MemeFollowers{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeFollowersJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeFollowers, error) {
	var m models.MemeFollowers
	if err := d.firstJoinSelect(tx, &m, selected, joins, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeFollowers(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeFollowers, error) {
	var m models.MemeFollowers
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// ////////chart
func (d *DAO) UpdateChartCandleDataByPair(tx *gorm.DB, memeID uint) (bool, error) {
	if err := tx.Exec("call create_meme_chart_candles_min30_by_id(?)", memeID).Error; err != nil {
		return false, errs.NewError(err)
	}

	if err := tx.Exec("call create_meme_chart_candles_min5_by_id(?)", memeID).Error; err != nil {
		return false, errs.NewError(err)
	}

	if err := tx.Exec("call create_meme_chart_candles_hour1_by_id(?)", memeID).Error; err != nil {
		return false, errs.NewError(err)
	}

	if err := tx.Exec("call create_meme_chart_candles_hour4_by_id(?)", memeID).Error; err != nil {
		return false, errs.NewError(err)
	}

	if err := tx.Exec("call create_meme_chart_candles_day1_by_id(?)", memeID).Error; err != nil {
		return false, errs.NewError(err)
	}

	return true, nil
}

func (d *DAO) GetMemeChartCandleData30Min(tx *gorm.DB, memeID, day uint) ([]*models.ChartData, error) {
	var rs []*models.ChartData
	queryStr := `
		SELECT 
			* from meme_chart_candles_min30
		where  (pair_id = ?) and DATE_ADD(chart_time, INTERVAL ? DAY) >= now()
		ORDER BY chart_time 
	`
	query := tx.Raw(queryStr, memeID, day)
	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

func (d *DAO) GetMemeChartCandleData1Hour(tx *gorm.DB, memeID, day uint) ([]*models.ChartData, error) {
	var rs []*models.ChartData
	queryStr := `
		SELECT 
			* from meme_chart_candles_hour1
		where  (pair_id = ?) and DATE_ADD(chart_time, INTERVAL ? DAY) >= now()
		ORDER BY chart_time 
	`
	query := tx.Raw(queryStr, memeID, day)
	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

func (d *DAO) GetMemeChartCandleData4Hour(tx *gorm.DB, memeID, day uint) ([]*models.ChartData, error) {
	var rs []*models.ChartData
	queryStr := `
		SELECT 
			* from meme_chart_candles_hour4
		where  (pair_id = ?) and DATE_ADD(chart_time, INTERVAL ? DAY) >= now()
		ORDER BY chart_time 
	`
	query := tx.Raw(queryStr, memeID, day)
	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

func (d *DAO) GetMemeChartCandleData1Day(tx *gorm.DB, memeID, day uint) ([]*models.ChartData, error) {
	var rs []*models.ChartData
	queryStr := `
		SELECT 
			* from meme_chart_candles_day1
		where  (pair_id = ?) and DATE_ADD(chart_time, INTERVAL ? DAY) >= now()
		ORDER BY chart_time 
	`
	query := tx.Raw(queryStr, memeID, day)
	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

func (d *DAO) GetMemeChartCandleData5Min(tx *gorm.DB, memeID, day uint) ([]*models.ChartData, error) {
	var rs []*models.ChartData
	queryStr := `
		SELECT 
			* from meme_chart_candles_min5
		where  (pair_id = ?) and DATE_ADD(chart_time, INTERVAL ? DAY) >= now()
		ORDER BY chart_time 
	`
	query := tx.Raw(queryStr, memeID, day)
	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}

// //////

func (d *DAO) FirstMemeNotificationByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.MemeNotification, error) {
	var m models.MemeNotification
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeNotification(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeNotification, error) {
	var m models.MemeNotification
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMemeNotificationJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MemeNotification, error) {
	var ms []*models.MemeNotification
	err := d.findJoinSelect(tx, &models.MemeTradeHistory{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeWhiteListAddress(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeWhiteListAddress, error) {
	var m models.MemeWhiteListAddress
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMemeWhiteListAddress(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.MemeWhiteListAddress, error) {
	var ms []*models.MemeWhiteListAddress
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindMemeBurnJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.TokenTransfer, error) {
	var ms []*models.TokenTransfer
	err := d.findJoinSelect(tx, &models.TokenTransfer{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstMemeNotificationSeen(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, order []string) (*models.MemeNotificationSeen, error) {
	var m models.MemeNotificationSeen
	if err := d.first(tx, &m, filters, preloads, order, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeSeenByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.MemeSeen, error) {
	var m models.MemeSeen
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMemeSeen(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.MemeSeen, error) {
	var m models.MemeSeen
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindTrendingToken(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.TrendingToken, error) {
	var tokens []*models.TrendingToken
	if err := d.find(tx, &tokens, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return tokens, nil
}

func (d *DAO) FindAllMemeTokenAddress(tx *gorm.DB, networkID uint64) ([]string, error) {
	var ms []*struct {
		TokenAddress string
		Pool         string
		UniswapPool  string
	}
	err := tx.Raw(`
		select token_address, pool, uniswap_pool
		from memes
		where network_id = ?
		and token_address != ''
	`, networkID).Scan(&ms).Error
	if err != nil {
		return nil, errs.NewError(err)
	}
	addresses := []string{}
	for _, v := range ms {
		if v.TokenAddress != "" {
			addresses = append(addresses, v.TokenAddress)
		}
		if v.Pool != "" {
			addresses = append(addresses, v.Pool)
		}
		if v.UniswapPool != "" {
			addresses = append(addresses, v.UniswapPool)
		}
	}
	return addresses, nil
}
