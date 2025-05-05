package factory

import (
	"context"
	"core/contracts/gpu_manager"
	"solo/chains/base_new"
	interCommon "solo/chains/common"
	"solo/config"
	"solo/internal/model"
	"solo/internal/port"
	"solo/internal/usecase"
	"solo/pkg"
	"solo/pkg/db/sqlite"
	"solo/pkg/logger"

	"go.uber.org/zap"
)

type InitObject struct {
	Miner port.IMiner
	API   port.IApi
}

func NewObject(cnf *config.Config) (*InitObject, error) {
	ctx := context.Background()
	var err error
	defer func() {
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("chainFactory",
				zap.String("chain", cnf.ChainID),
				zap.Error(err),
			)
		}
	}()

	cm, err := interCommon.NewCommon(ctx, cnf)
	if err != nil {
		return nil, err
	}

	var miner port.IMiner
	switch cnf.ChainID {
	case pkg.CHAIN_BASE:
		c, err := base_new.NewChain(ctx, cm)
		if err != nil {
			return nil, err
		}

		sthub, err := gpu_manager.NewGpuManager(
			cm.GetStakingHubAddress(), cm.GetClient(),
		)
		if err != nil {
			return nil, err
		}

		s := base_new.NewStaking(cm, sthub)
		cluster, err := base_new.NewCluster(cm)
		if err != nil {
			return nil, err
		}

		miner = usecase.NewMiner(c, s, cm, cnf, cluster)
		//api := usecase.NewAPI(c, s, cm, cnf, nil)
	}

	sqldb, err := sqlite.NewSqliteDB(cnf.NodeID)
	if err == nil {
		miner.SetDB(sqldb)
	} else {
		logger.AtLog.Logger.Error("NewObject",
			zap.String("node_id", cnf.NodeID), zap.Error(err),
		)
	}

	//migrate db
	sqldb.Connect()
	defer sqldb.Close()

	db := model.SuccessTask{}
	if _, err = sqldb.CreateTable(db.CreateDB()); err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("saveSuccessTask.CreateTable",
			zap.Error(err),
		)
		return nil, err
	}

	return &InitObject{Miner: miner}, nil
}

func NewMiner(cnf *config.Config) (*InitObject, error) {
	return NewObject(cnf)
}

func NewAPI(port int) (*InitObject, error) {
	return &InitObject{API: usecase.NewAPI(port)}, nil
}
