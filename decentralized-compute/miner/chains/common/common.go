package common

import (
	"context"
	"fmt"
	"math/big"
	"solo/pkg"

	"core/contracts/erc20"
	"solo/config"
	"solo/internal/port"
	"solo/pkg/eth"
	"solo/pkg/logger"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type common struct {
	cnf          *config.Config
	client       *ethclient.Client
	privateKey   string
	address      *ethCommon.Address
	gasLimit     uint64
	modelAddress string

	stakingHubAddress      string
	erc20contractAddress   string
	workerHubAddress       string
	loadBalancerAddress    string
	modelCollectionAddress string
	erc20contract          *erc20.Erc20
}

func (c *common) GetWalletAddres() ethCommon.Address {
	return *c.address
}

func (c *common) GetModelAddress() string {
	return c.modelAddress
}

func (c *common) GetWorkerHubAddress() ethCommon.Address {
	return ethCommon.HexToAddress(c.workerHubAddress)
}

func (c *common) GetClient() *ethclient.Client {
	return c.client
}

func (c *common) GetPrivateKey() string {
	return c.privateKey
}

func (c *common) GetStakingHubAddress() ethCommon.Address {
	return ethCommon.HexToAddress(c.stakingHubAddress)
}

func (c *common) GetErc20contractAddress() ethCommon.Address {
	return ethCommon.HexToAddress(c.erc20contractAddress)
}

func (c *common) GetGasLimit() uint64 {
	return c.gasLimit
}

func (c *common) GetErc20contract() *erc20.Erc20 {
	return c.erc20contract
}

func (c *common) CurrentBlock() uint64 {
	bln, err := c.client.BlockNumber(context.Background())
	if err != nil {
		return uint64(0)
	}
	return bln
}

func (c *common) FromBlock(block uint64) uint64 {
	if block != 0 {
		return block
	}

	cblock := c.CurrentBlock()
	cblock = cblock - 10
	return cblock
}

func (c *common) ToBlock() uint64 {
	return c.CurrentBlock()
}

func (c *common) connect(rpc string) error {
	ethClient, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	c.client = ethClient
	return nil
}

func (c *common) account(cnf config.Config) error {
	c.privateKey = cnf.Account
	_, address, err := eth.GetAccountInfo(c.privateKey)
	if err != nil {
		return err
	}

	c.address = address

	return nil
}

func (b *common) connectContractErc20() error {
	erc20Hub, err := erc20.NewErc20(ethCommon.HexToAddress(b.erc20contractAddress), b.client)
	if err != nil {
		return err
	}

	b.erc20contract = erc20Hub
	return nil
}

func NewCommon(ctx context.Context, cnf *config.Config) (port.ICommon, error) {
	var err error
	defer func() {
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("chainFactory",
				zap.String("chain", cnf.ChainID),
				zap.Error(err),
			)
		}
	}()

	c := &common{cnf: cnf}
	if err = c.connect(cnf.Rpc); err != nil {
		return nil, err
	}

	if cnf.Account == "" {
		return nil, fmt.Errorf("ACCOUNT_PRIV is empty. Use %s to set it", pkg.COMMAND_CONFIG)
	}

	if err = c.account(*cnf); err != nil {
		return nil, err
	}

	c.workerHubAddress = cnf.WorkerHubAddress
	c.stakingHubAddress = cnf.StakingHubAddress
	c.erc20contractAddress = cnf.Erc20Address
	c.loadBalancerAddress = cnf.ModelLoadBalancerAddress
	c.modelCollectionAddress = cnf.ModelCollectionAddress

	if err = c.connectContractErc20(); err != nil {
		return nil, err
	}

	return c, nil
}

func (b *common) GetConfig() *config.Config {
	return b.cnf
}

func (b *common) Erc20Balance() (*big.Int, error) {
	return b.GetErc20contract().BalanceOf(nil, b.GetWalletAddres())
}

func (b *common) GetModelCollectionAddress() ethCommon.Address {
	return ethCommon.HexToAddress(b.modelCollectionAddress)
}

func (b *common) GetModelLoadBalancerAddress() ethCommon.Address {
	return ethCommon.HexToAddress(b.loadBalancerAddress)
}
