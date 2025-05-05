package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetAllConfigs(ctx context.Context) (map[uint64]map[string]string, error) {
	resp := map[uint64]map[string]string{}
	for networkStr, network := range s.conf.Networks {
		networkID, _ := strconv.ParseUint(networkStr, 10, 64)
		m, ok := resp[networkID]
		if !ok {
			m = map[string]string{}
			resp[networkID] = m
		}
		m["agent_admin_address"] = network["agent_admin_address"]
		m["agent_contract_address"] = network["agent_contract_address"]
		m["eai_contract_address"] = network["eai_contract_address"]
		m["explorer_url"] = network["explorer_url"]
		m["uniswap_factory_contract_address"] = network["uniswap_factory_contract_address"]
		m["uniswap_position_mamanger_address"] = network["uniswap_position_mamanger_address"]
		m["weth9_contract_address"] = network["weth9_contract_address"]
	}
	return resp, nil
}

func (s *Service) GetAllConfigsExplorer(ctx context.Context) (map[uint64]string, error) {
	resp := map[uint64]string{}
	for networkStr, network := range s.conf.Networks {
		networkID, _ := strconv.ParseUint(networkStr, 10, 64)
		resp[networkID] = network["explorer_url"]
	}
	return resp, nil
}

func (s *Service) CreateMeme(ctx context.Context, address string, networkID uint64, req *serializers.MemeReq) (*models.Meme, error) {
	meme := &models.Meme{
		NetworkID:         networkID,
		OwnerAddress:      strings.ToLower(address),
		TokenAddress:      "",
		Name:              req.Name,
		Description:       req.Description,
		Ticker:            req.Ticker,
		Image:             req.Image,
		Twitter:           req.Twitter,
		Telegram:          req.Telegram,
		Website:           req.Website,
		Status:            models.MemeStatusNew,
		StoreImageOnChain: req.OnchainImage,
		TotalSuply:        numeric.NewBigFloatFromString("1000000000"),
		Supply:            numeric.NewBigFloatFromString("1000000000"),
		Decimals:          18,
		AgentInfoID:       req.AgentInfoID,
		BaseTokenSymbol:   req.BaseTokenSymbol,
		ReqSyncAt:         helpers.TimeNow(),
		SyncAt:            helpers.TimeNow(),
		TokenId:           helpers.RandomBigInt(32).String(),
		NotGraduated:      req.NotGraduated,
	}
	switch meme.NetworkID {
	case models.BASE_CHAIN_ID,
		models.ARBITRUM_CHAIN_ID,
		models.BSC_CHAIN_ID,
		models.AVALANCHE_C_CHAIN_ID,
		models.APE_CHAIN_ID,
		models.CELO_CHAIN_ID:
		{
			agentChainFee, err := s.GetAgentChainFee(
				daos.GetDBMainCtx(ctx),
				meme.NetworkID,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}
			meme.Fee = agentChainFee.TokenFee
		}
	default:
		{
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}
	agent, err := s.dao.FirstAgentInfoByID(
		daos.GetDBMainCtx(ctx),
		meme.AgentInfoID,
		map[string][]any{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agent.RefTweetID > 0 {
		meme.Fee = numeric.NewBigFloatFromString("0")
	}
	owner, err := s.GetUser(daos.GetDBMainCtx(ctx), networkID, meme.OwnerAddress, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	meme.OwnerID = owner.ID
	err = s.dao.Create(daos.GetDBMainCtx(ctx), meme)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return meme, nil
}

func (s *Service) GetListMemeReport(ctx context.Context, chainID uint64,
	address, search, status string, following bool,
	sortListStr []string, page, limit int,
) ([]*models.Meme, uint, error) {
	var resp []*models.Meme
	cacheKey := fmt.Sprintf(`CacheGetListMemeReport_%s_%s_%s_%d_%d`, address, strings.Join(sortListStr, "_"), status, page, limit)
	if search == "" {
		s.GetRedisCachedWithKey(cacheKey, &resp)
		if len(resp) > 0 {
			return resp, 0, nil
		}
	}

	sortDefault := "market_cap desc"
	if len(sortListStr) > 0 {
		sortDefault = strings.Join(sortListStr, ", ")
	}
	joinFilters := map[string][]any{}
	filters := map[string][]any{}

	if chainID > 0 {
		filters[`network_id = ?`] = []any{chainID}
	}

	selected := []string{
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), 0) percent`,
		`(memes.price_usd * memes.total_suply) market_cap`,
		`memes.*`,
	}

	if address != "" {
		if following {
			filters[`memes.owner_address in (
				select d.address
				from users u
				join meme_followers f on f.user_id = u.id
				join users d on f.follow_user_id = d.id
				where u.address = ?
			)`] = []any{strings.ToLower(address)}
		} else {
			joinFilters[`
				left join (
				SELECT
					cast(balance as decimal(36, 18)) total_balance,
					memes.id meme_id
				FROM erc20_holders
				join memes on erc20_holders.contract_address = memes.token_address
				WHERE
					erc20_holders.deleted_at IS NULL
					AND cast(balance as decimal(36, 18)) > 0.00000001
					AND erc20_holders.address = ?
			) h on memes.id = h.meme_id
			`] = []any{strings.ToLower(address)}
			filters[`memes.owner_address = ? or ifnull(h.total_balance, 0) > 0`] = []any{strings.ToLower(address)}
			selected = append(selected, `ifnull(h.total_balance, 0) total_balance`)
		}
	}

	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(memes.ticker) like ?
			or LOWER(memes.name) like ?
			or LOWER(memes.token_address) like ?
		`] = []any{search, search, search}
	}

	if status != "" {
		listStatus := strings.Split(status, ",")
		if len(listStatus) > 0 {
			filters["memes.status in (?)"] = []any{listStatus}
		}
	}
	keys, count, err := s.dao.FindMemeJoinSelect4Page(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"Owner":                    {},
			"AgentInfo":                {},
			"AgentInfo.TwitterInfo":    {},
			"AgentInfo.TmpTwitterInfo": {},
		},
		[]string{sortDefault}, page, limit,
	)

	if err != nil {
		return nil, count, errs.NewError(err)
	}

	agentInfoIds := []uint{}
	for _, v := range keys {
		agentInfoIds = append(agentInfoIds, v.AgentInfoID)
	}
	latestPost, _ := s.GetListAgentTwitterLatestPost(ctx, 0, agentInfoIds, 0, 100000)

	for _, v := range keys {
		v.LatestAgentTwitterPost = latestPost[v.ID]
	}

	if search == "" {
		err = s.SetRedisCachedWithKey(
			cacheKey,
			keys,
			30*time.Second,
		)
	}

	return keys, count, nil
}

// //////Feed API
func (s *Service) GetFeedMemeReport(ctx context.Context, address, search, sortType string, page, limit int) ([]*models.Meme, uint, error) {
	joinFilters := map[string][]any{}
	filters := map[string][]any{}
	sortNum, _ := strconv.Atoi(sortType)
	sortDefault := `seen_time asc, weight desc, percent desc`
	if sortNum > 0 {
		sortDefault = `seen_time asc, weight asc, percent asc`
	}

	selected := []string{
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), 0) percent`,
		`(memes.price_usd * memes.total_suply) market_cap`,
		`memes.*`,
	}

	if address != "" {
		joinFilters[`left join meme_seens on meme_seens.meme_id = memes.id and meme_seens.user_address = ?`] = []any{strings.ToLower(address)}
		selected = append(selected, `ifnull(DATE_FORMAT(meme_seens.seen_time, '%Y-%m-%d'), '2000-01-01') seen_time`)
	} else {
		selected = append(selected, `'2000-01-01' seen_time`)
	}

	if search != "" {
		search = fmt.Sprintf("%%%s%%", strings.ToLower(search))
		filters[`
			LOWER(memes.ticker) like ?
			or LOWER(memes.name) like ?
			or LOWER(memes.token_address) like ?
			or LOWER(memes.description) like ?
		`] = []any{search, search, search, search}
	}

	keys, count, err := s.dao.FindMemeJoinSelect4Page(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"Owner": []any{},
		},
		[]string{sortDefault}, page, limit,
	)

	if err != nil {
		return nil, count, errs.NewError(err)
	}

	return keys, count, nil
}

func (s *Service) GetMemeDetail(ctx context.Context, address, memAddress string) (string, error) {
	var resp string
	cacheKey := fmt.Sprintf(`CacheMemeDetail_%s`, strings.ToLower(memAddress))
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheMemeDetail(daos.GetDBMainCtx(ctx), memAddress)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}
	go s.CheckedSeenMeme(ctx, address, memAddress)

	return resp, nil
}

func (s *Service) CheckedSeenMeme(ctx context.Context, address, memAddress string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			filters := map[string][]any{
				"memes.token_address = ?": {strings.ToLower(memAddress)},
			}

			meme, err := s.dao.FirstMeme(tx,
				filters,
				map[string][]any{},
				false,
			)

			if err != nil {
				return errs.NewError(err)
			}

			user, err := s.GetUser(tx, meme.NetworkID, strings.ToLower(address), false)
			if err != nil {
				return errs.NewError(err)
			}

			if meme != nil && user != nil {
				memeSeen, err := s.dao.FirstMemeSeen(tx,
					map[string][]any{
						"user_id = ?": {user.ID},
						"meme_id = ?": {meme.ID},
					},
					map[string][]any{}, false)

				if err != nil {
					return errs.NewError(err)
				}
				if memeSeen == nil {
					memeSeen = &models.MemeSeen{
						UserID:      user.ID,
						UserAddress: strings.ToLower(user.Address),
						MemeID:      meme.ID,
						MemeAddress: strings.ToLower(meme.TokenAddress),
						SeenTime:    helpers.TimeNow(),
					}
					err = s.dao.Create(tx, memeSeen)
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					memeSeen, _ = s.dao.FirstMemeSeenByID(tx, memeSeen.ID, map[string][]any{}, true)
					memeSeen.SeenTime = helpers.TimeNow()
					err = s.dao.Save(tx, memeSeen)
					if err != nil {
						return errs.NewError(err)
					}
				}
			}

			return nil
		},
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil

}

func (s *Service) CacheMemeDetail(tx *gorm.DB, memAddress string) error {
	filters := map[string][]any{
		"memes.token_address = ?": {strings.ToLower(memAddress)},
	}
	joinFilters := map[string][]any{}
	joinFilters[`
		left join (
			SELECT
				count(DISTINCT erc20_holders.address) holders,
				memes.id meme_id
			FROM erc20_holders
			join memes on erc20_holders.contract_address = memes.token_address
			WHERE 1=1
				AND erc20_holders.deleted_at IS NULL
				AND cast(balance as decimal(36, 18)) > 0.00000001
				AND erc20_holders.contract_address = ?
			group by memes.id
		) h on memes.id = h.meme_id
	`] = []any{strings.ToLower(memAddress)}
	selected := []string{
		`ifnull((cast(( memes.price - memes.price_last24h) / memes.price_last24h * 100 as decimal(20, 2))), 0) percent`,
		"memes.*",
		"ifnull(h.holders, 0) holders",
		`(memes.price_usd * memes.total_suply) market_cap`,
	}
	key, err := s.dao.FirstMemeJoinSelect(tx,
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"Owner":                    {},
			"AgentInfo":                {},
			"AgentInfo.TwitterInfo":    {},
			"AgentInfo.TmpTwitterInfo": {},
		},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}
	cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewMemeResp(key)})
	if err != nil {
		errs.NewError(err)
	}
	err = s.SetRedisCachedWithKey(
		fmt.Sprintf(`CacheMemeDetail_%s`, strings.ToLower(memAddress)),
		string(cacheData),
		30*time.Second,
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) GetMemeTradeHistory(ctx context.Context, userAddress, tokenAddress string, page, limit int) ([]*models.MemeTradeHistory, uint, error) {
	filters := map[string][]any{}
	preloads := map[string][]any{}
	joinFilters := map[string][]any{
		`join memes on meme_trade_histories.meme_id = memes.id`: {},
	}

	if userAddress != "" {
		filters["meme_trade_histories.user_address = ?"] = []any{strings.ToLower(userAddress)}
		preloads[`Meme`] = []any{}
	}

	if tokenAddress != "" {
		filters["meme_trade_histories.meme_token_address = ? or memes.agent_info_id = ?"] = []any{strings.ToLower(tokenAddress), tokenAddress}
		preloads[`RecipientUser`] = []any{}
	}

	selected := []string{
		"meme_trade_histories.*",
	}

	keys, err := s.dao.FindMemeHistoryJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		preloads,
		[]string{"meme_trade_histories.tx_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) GetMemeTradeHistoryLatest(ctx context.Context, tokenAddress string) (string, error) {
	var resp string
	cacheKey := fmt.Sprintf(`CacheMemeTradeHistoryLatest_%s`, strings.ToLower(tokenAddress))
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheMemeTradeHistoryLatest(daos.GetDBMainCtx(ctx), tokenAddress)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) CacheMemeTradeHistoryLatest(tx *gorm.DB, tokenAddress string) error {
	filters := map[string][]any{}
	preloads := map[string][]any{}

	if tokenAddress != "" {
		filters["meme_token_address = ?"] = []any{strings.ToLower(tokenAddress)}
		preloads[`RecipientUser`] = []any{}
	}

	joinFilters := map[string][]any{}

	selected := []string{
		"meme_trade_histories.*",
	}

	keys, err := s.dao.FindMemeHistoryJoinSelect(tx,
		selected,
		joinFilters,
		filters,
		preloads,
		[]string{"tx_at desc"}, 1, 20,
	)

	if err != nil {
		errs.NewError(err)
	}

	cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewTradeHistoryRespArry(keys)})
	if err != nil {
		errs.NewError(err)
	}

	err = s.SetRedisCachedWithKey(
		fmt.Sprintf(`CacheMemeTradeHistoryLatest_%s`, strings.ToLower(tokenAddress)),
		string(cacheData),
		1*time.Hour,
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

func (s *Service) GetTokenHolders(ctx context.Context, tokenAddress string) (string, error) {
	var resp string
	cacheKey := fmt.Sprintf(`CacheMemeHolders1_%s`, strings.ToLower(tokenAddress))
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheMemeHolders(daos.GetDBMainCtx(ctx), tokenAddress)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) CacheMemeHolders(tx *gorm.DB, tokenAddress string) error {
	filters := map[string][]any{
		`cast(erc20_holders.balance as decimal(36, 18)) > 0.0000000001`: {},
	}

	joinFilters := map[string][]any{
		`join agent_infos on agent_infos.token_address = erc20_holders.contract_address and agent_infos.token_address is not null and agent_infos.token_address != ""`: {},
	}

	if tokenAddress != "" {
		filters["erc20_holders.contract_address = ? or agent_infos.id=?"] = []any{strings.ToLower(tokenAddress), tokenAddress}
		filters["erc20_holders.address != ?"] = []any{strings.ToLower("0x92a57bb5d9d61214cf4b7a85d3f74fef10e1ff37")}
	}

	selected := []string{
		"erc20_holders.*",
		"cast(erc20_holders.balance as decimal(36, 18)) total_balance",
	}

	keys, err := s.dao.FindErc20HolderJoinSelect(tx,
		selected,
		joinFilters,
		filters,
		map[string][]any{},
		[]string{"total_balance desc"}, 1, 1000,
	)

	if err != nil {
		return errs.NewError(err)
	}

	cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewTokenHolderRespArray(keys)})
	if err != nil {
		errs.NewError(err)
	}

	err = s.SetRedisCachedWithKey(
		fmt.Sprintf(`CacheMemeHolders1_%s`, strings.ToLower(tokenAddress)),
		string(cacheData),
		1*time.Hour,
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

func (s *Service) GetTokenHolding(ctx context.Context, userAddress string, page, limit int) ([]*models.Erc20Holder, uint, error) {
	filters := map[string][]any{
		`cast(balance as decimal(36, 18)) > 0.00000001`: {},
	}

	if userAddress != "" {
		filters["erc20_holders.address = ?"] = []any{strings.ToLower(userAddress)}
	}

	joinFilters := map[string][]any{
		`
			join memes on erc20_holders.contract_address = memes.token_address
		`: {},
	}

	selected := []string{
		"cast(erc20_holders.balance as decimal(36, 18)) total_balance",
		"cast(erc20_holders.balance as decimal(36, 18)) balance",
		"erc20_holders.contract_address",
		"erc20_holders.address",
		`memes.name meme_name, memes.ticker meme_ticker, memes.image meme_image,
		memes.price meme_price, memes.price_usd meme_price_usd, memes.base_token_symbol meme_base_token_symbol`,
		"memes.owner_address address, memes.token_address",
	}

	keys, err := s.dao.FindErc20HolderJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{},
		[]string{"total_balance desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

// /////Users
func (s *Service) GetMemeUserProfile(ctx context.Context, networkID uint64, address string) (*models.User, error) {
	user, err := s.GeUserByAddress(ctx, networkID, address)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if user == nil {
		return nil, errs.NewError(errs.ErrUserNotFound)
	}

	filters := map[string][]any{
		"address = ?":    {user.Address},
		"network_id = ?": {networkID},
	}

	joinFilters := map[string][]any{
		`
		left join (
			select ifnull(count(*), 0) total_noti, ifnull(sum(case when seen = 1 then 1 end), 0) total_seen
			from users u
			left join meme_notifications on u.id= meme_notifications.user_id or meme_notifications.user_id = 0
			where u.address = ?
		) all_noti on 1=1
		left join (
			select count(*) total_seen
			from users u
			join meme_notification_seens on u.id= meme_notification_seens.user_id
			where u.address = ?
		) seen on 1=1
		`: {user.Address, user.Address},
	}

	selected := []string{
		"users.*",
		"all_noti.total_noti, (seen.total_seen + all_noti.total_seen) total_seen",
		"all_noti.total_noti - (seen.total_seen + all_noti.total_seen) total_unseen",
	}

	keys, err := s.dao.FirstUserJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{},
		false,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return keys, nil
}

func (s *Service) GetListFollowers(ctx context.Context, userAddress string, page, limit int) ([]*models.MemeFollowers, uint, error) {
	filters := map[string][]any{}
	joinFilters := map[string][]any{
		`join users on users.id=meme_followers.follow_user_id and users.address = ?`: {strings.ToLower(userAddress)},
	}

	selected := []string{
		"meme_followers.*",
	}

	keys, err := s.dao.FindMemeFollowersJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{
			"User": []any{},
		},
		[]string{"created_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) GetMemeChartCandleData(ctx context.Context, tokenAddress string, day uint, chartType string) (string, error) {
	var resp string
	if chartType == "" {
		chartType = string(models.ChartTypeMin30)
	}

	meme, err := s.dao.FirstMeme(daos.GetDBMainCtx(ctx),
		map[string][]any{
			"token_address = ?": {strings.ToLower(tokenAddress)},
		},
		map[string][]any{},
		false,
	)
	if err != nil {
		return resp, errs.NewError(err)
	}

	cacheKey := fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, meme.ID, chartType)
	if meme != nil {
		err := s.GetRedisCachedWithKey(cacheKey, &resp)
		if err != nil {
			s.CacheMemeCandleDataChart(daos.GetDBMainCtx(ctx), meme.ID)
		}

		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) CacheMemeCandleDataChart(tx *gorm.DB, memeID uint) error {
	_, err := s.dao.UpdateChartCandleDataByPair(tx, memeID)
	if err != nil {
		return errs.NewError(err)
	}
	//chart 5min
	{
		chartData, err := s.dao.GetMemeChartCandleData5Min(
			tx, memeID, 90)
		if err != nil {
			return errs.NewError(err)
		}
		cacheData, err := json.Marshal(&serializers.Resp{Result: chartData})
		if err != nil {
			errs.NewError(err)
		}
		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, memeID, models.ChartTypeMin5),
			string(cacheData),
			1*time.Hour,
		)
		if err != nil {
			return errs.NewError(err)
		}
	}

	//chart 30min
	{
		chartData, err := s.dao.GetMemeChartCandleData30Min(
			tx, memeID, 90)
		if err != nil {
			return errs.NewError(err)
		}
		cacheData, err := json.Marshal(&serializers.Resp{Result: chartData})
		if err != nil {
			errs.NewError(err)
		}

		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, memeID, models.ChartTypeMin30),
			string(cacheData),
			1*time.Hour,
		)
		if err != nil {
			return errs.NewError(err)
		}
	}

	//chart 1hour
	{
		chartData, err := s.dao.GetMemeChartCandleData1Hour(
			tx, memeID, 90)
		if err != nil {
			return errs.NewError(err)
		}
		cacheData, err := json.Marshal(&serializers.Resp{Result: chartData})
		if err != nil {
			errs.NewError(err)
		}

		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, memeID, models.ChartTypeHour1),
			string(cacheData),
			1*time.Hour,
		)
		if err != nil {
			return errs.NewError(err)
		}
	}

	//chart 4hour
	{
		chartData, err := s.dao.GetMemeChartCandleData4Hour(
			tx, memeID, 90)
		if err != nil {
			return errs.NewError(err)
		}
		cacheData, err := json.Marshal(&serializers.Resp{Result: chartData})
		if err != nil {
			errs.NewError(err)
		}

		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, memeID, models.ChartTypeHour4),
			string(cacheData),
			1*time.Hour,
		)
		if err != nil {
			return errs.NewError(err)
		}
	}

	//chart 1day
	{
		chartData, err := s.dao.GetMemeChartCandleData1Day(
			tx, memeID, 90)
		if err != nil {
			return errs.NewError(err)
		}
		cacheData, err := json.Marshal(&serializers.Resp{Result: chartData})
		if err != nil {
			errs.NewError(err)
		}

		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`MemeGetChartCandleData_%d_%s`, memeID, models.ChartTypeDay),
			string(cacheData),
			1*time.Hour,
		)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) MemeSnapshotTokenHolder(ctx context.Context, tokenAddress string) error {
	filters := map[string][]any{
		`cast(balance as decimal(36, 18)) > 0`: {},
	}

	if tokenAddress != "" {
		filters["erc20_holders.contract_address = ?"] = []any{strings.ToLower(tokenAddress)}
	}

	joinFilters := map[string][]any{}

	selected := []string{
		"erc20_holders.*",
	}

	keys, err := s.dao.FindTokenHolderJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{},
		[]string{}, 1, 10000,
	)

	if err != nil {
		return errs.NewError(err)
	}

	for _, key := range keys {
		holder := &models.MemeTokenHolder{
			ContractAddress: key.ContractAddress,
			Address:         key.Address,
			Balance:         numeric.NewBigFloatFromString(key.Balance),
		}
		err = s.dao.Create(daos.GetDBMainCtx(ctx), holder)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) GetMemeWhiteListAddress(ctx context.Context) ([]string, error) {
	keys, err := s.dao.FindMemeWhiteListAddress(daos.GetDBMainCtx(ctx),
		map[string][]any{},
		map[string][]any{},
		[]string{},
		0, 1000,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	resp := []string{}
	for _, item := range keys {
		resp = append(resp, item.Address)
	}

	return resp, nil
}

func (s *Service) GetMemeBurnHistory(ctx context.Context, networkID uint64, userAddress, tokenAddress string, page, limit int) ([]*models.TokenTransfer, uint, error) {
	burnAddress := s.conf.GetConfigKeyString(networkID, "meme_burn_address")
	filters := map[string][]any{
		"token_transfers.to = ?": {strings.ToLower(burnAddress)},
	}

	if userAddress != "" {
		filters["token_transfers.from = ?"] = []any{strings.ToLower(userAddress)}
	}

	joinFilters := map[string][]any{
		`
			join users on users.address = token_transfers.from
			join memes on memes.token_address = token_transfers.contract_address
		`: {},
	}

	selected := []string{
		`
		token_transfers.*,
		users.twitter_name,
		users.twitter_username,
		users.twitter_avatar,
		users.user_twitter_id,
		users.user_name,
		users.image_url,
		users.user_twitter_id,
		memes.name,
		memes.ticker,
		memes.image,
		memes.token_address
		`,
	}

	keys, err := s.dao.FindMemeBurnJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]any{},
		[]string{"transaction_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) ShareMeme(ctx context.Context, address, memAddress string) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			filters := map[string][]any{
				"memes.token_address = ?": {strings.ToLower(memAddress)},
			}

			meme, err := s.dao.FirstMeme(tx,
				filters,
				map[string][]any{},
				false,
			)

			if err != nil {
				return errs.NewError(err)
			}

			if meme != nil {
				err = daos.GetDBMainCtx(ctx).
					Model(meme).
					Updates(
						map[string]any{
							"shared": gorm.Expr("shared + 1"),
						},
					).Error
				if err != nil {
					return errs.NewError(err)
				}
			}

			return nil
		},
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil

}

func (s *Service) VibeTokenGetDeployInfo(ctx context.Context, address string, agentID uint) (*serializers.VibeTokenDeployInfoResp, error) {
	agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), agentID, map[string][]any{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if !strings.EqualFold(address, agentInfo.Creator) {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if !agentInfo.IsVibeAgent() {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if agentInfo.TokenNetworkID == 0 {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if agentInfo.TokenSymbol == "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if agentInfo.TokenName == "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	vibeToken, err := s.dao.FirstMeme(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"agent_info_id = ?": {agentID},
		},
		map[string][]any{},
		false,
	)
	if err != nil {
		return nil, err
	}
	if vibeToken == nil {
		_, err = s.CreateMeme(
			ctx, agentInfo.Creator,
			agentInfo.TokenNetworkID,
			&serializers.MemeReq{
				Name:            agentInfo.TokenName,
				Ticker:          agentInfo.TokenSymbol,
				Description:     agentInfo.TokenDesc,
				Image:           agentInfo.TokenImageUrl,
				Twitter:         fmt.Sprintf("https://x.com/%s", agentInfo.TwitterUsername),
				AgentInfoID:     agentInfo.ID,
				BaseTokenSymbol: string(models.BaseTokenSymbolEAI),
			})
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	vibeToken, err = s.dao.FirstMeme(
		daos.GetDBMainCtx(ctx),
		map[string][]any{
			"agent_info_id = ?": {agentID},
		},
		map[string][]any{},
		false,
	)
	if err != nil {
		return nil, err
	}
	if vibeToken == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	if vibeToken.TokenAddress != "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	baseTokenPrice := s.GetTokenMarketPrice(daos.GetDBMainCtx(ctx), vibeToken.BaseTokenSymbol)
	lowerPrice, _ := models.QuoBigFloats(
		big.NewFloat(models.LOWER_PRICE_USD),
		big.NewFloat(models.TOKEN_SUPPLY),
		baseTokenPrice,
	).Float64()
	upperPrice, _ := models.QuoBigFloats(
		big.NewFloat(models.UPPER_PRICE_USD),
		big.NewFloat(models.TOKEN_SUPPLY),
		baseTokenPrice,
	).Float64()
	lowerTick := models.PriceToTick(lowerPrice, 1200)
	upperTick := models.PriceToTick(upperPrice, 1200)
	// calculate lower and upper tick base on EAI price
	deadline := time.Now().Add(time.Minute * 10).Unix()
	signature, err := s.GetEthereumClient(ctx, vibeToken.NetworkID).GetSignatureForDeployToken(
		s.GetAddressPrk(s.conf.GetConfigKeyString(uint64(vibeToken.NetworkID), "vibe_token_factory_admin")),
		helpers.HexToAddress(s.conf.GetConfigKeyString(uint64(vibeToken.NetworkID), "vibe_token_factory_address")),
		vibeToken.NetworkID,
		helpers.HexToBytes32(agentInfo.AgentID),
		vibeToken.Name,
		vibeToken.Ticker,
		helpers.HexToAddress(agentInfo.Creator),
		big.NewInt(int64(lowerTick)),
		big.NewInt(int64(upperTick)),
		big.NewInt(deadline),
	)
	if err != nil {
		return nil, err
	}
	return &serializers.VibeTokenDeployInfoResp{
		NonceHex:  common.HexToHash(agentInfo.AgentID).Hex(),
		Name:      vibeToken.Name,
		Symbol:    vibeToken.Ticker,
		Creator:   agentInfo.Creator,
		LowerTick: lowerTick,
		UpperTick: upperTick,
		Deadline:  deadline,
		Signature: signature,
	}, nil
}
