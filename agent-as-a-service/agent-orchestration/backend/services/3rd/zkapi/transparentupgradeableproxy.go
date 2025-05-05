package zkapi

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/transparentupgradeableproxyzk"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) DeployTransparentUpgradeableProxy(prkHex string, logic common.Address, admin common.Address, data []byte) (string, string, error) {
	instanceABI, err := transparentupgradeableproxyzk.TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return "", "", err
	}
	dataBytes, err := instanceABI.Pack(
		"",
		logic,
		admin,
		data,
	)
	if err != nil {
		return "", "", err
	}
	address, txHash, err := c.DeployContract(prkHex, transparentupgradeableproxyzk.TransparentUpgradeableProxyBin, common.Bytes2Hex(dataBytes))
	if err != nil {
		return "", "", err
	}
	return address, txHash, nil
}
