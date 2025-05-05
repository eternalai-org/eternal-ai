package base_new

import (
	"context"
	"core/contracts/load_balancer"
	"core/contracts/model_collection"
	"encoding/json"
	"log"
	"math/big"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/core/types"
)

type cluster struct {
	modelLoadBalancer *load_balancer.LoadBalancer
	modelCollection   *model_collection.ModelCollection
	common            port.ICommon
}

func NewCluster(common port.ICommon) (port.ICluster, error) {
	_cter := &cluster{
		common: common,
	}

	mLbc, err := load_balancer.NewLoadBalancer(
		_cter.common.GetModelLoadBalancerAddress(),
		_cter.common.GetClient(),
	)
	if err != nil {
		return nil, err
	}

	mCll, err := model_collection.NewModelCollection(
		_cter.common.GetModelCollectionAddress(),
		_cter.common.GetClient(),
	)
	if err != nil {
		return nil, err
	}

	_cter.modelLoadBalancer = mLbc
	_cter.modelCollection = mCll

	return _cter, nil
}

func (c *cluster) CreateCluster(version int, minHardware int, modelName, modelType string) (*types.Transaction, *big.Int, error) {
	ctx := context.Background()
	client := c.common.GetClient()

	auth, err := eth.CreateBindTransactionOpts(ctx, client, c.common.GetPrivateKey(), int64(c.common.GetGasLimit()))
	if err != nil {
		return nil, nil, err
	}

	clusterMetaData := model.ClusterMetaData{
		Version:     version,
		ModelName:   modelName,
		ModelType:   modelType,
		MinHardware: minHardware,
	}

	metadata, err := json.MarshalIndent(clusterMetaData, "", "\t")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	tx, err := c.modelCollection.Mint(auth, c.common.GetWalletAddres(), string(metadata))
	if err != nil {
		return nil, nil, err
	}

	err = eth.WaitForTx(client, tx.Hash())
	if err != nil {
		return tx, nil, err
	}

	tokenID, err := eth.GetTokenIDFromTx(client, tx.Hash())
	if err != nil {
		return tx, nil, err
	}

	return tx, tokenID, nil
}

func (c *cluster) CreateAGroupOfCluster(groupName string, clusterIDs []*big.Int) (*types.Transaction, error) {
	ctx := context.Background()
	client := c.common.GetClient()

	auth, err := eth.CreateBindTransactionOpts(ctx, client, c.common.GetPrivateKey(), int64(c.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	return c.modelLoadBalancer.CreateGroup(auth, groupName, clusterIDs)
}

func (c *cluster) AddClustersToGroup(groupName string, clusterIDs []*big.Int) (*types.Transaction, error) {
	ctx := context.Background()
	client := c.common.GetClient()

	auth, err := eth.CreateBindTransactionOpts(ctx, client, c.common.GetPrivateKey(), int64(c.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	return c.modelLoadBalancer.AddClustersToGroup(auth, groupName, clusterIDs)
}

func (c *cluster) RemoveClustersFromGroup(groupName string, clusterIDs []*big.Int) (*types.Transaction, error) {
	ctx := context.Background()
	client := c.common.GetClient()

	auth, err := eth.CreateBindTransactionOpts(ctx, client, c.common.GetPrivateKey(), int64(c.common.GetGasLimit()))
	if err != nil {
		return nil, err
	}

	return c.modelLoadBalancer.RemoveClustersFromGroup(auth, groupName, clusterIDs)
}
