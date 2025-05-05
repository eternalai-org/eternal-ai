package evmapi

import (
	"math/big"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/agentupgradeable"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type BaseClient interface {
	WaitMined(hash string) error
	Transact(contractAddr string, prkHex string, input []byte, value *big.Int) (string, error)
	TransactionConfirmed(hash string) error
	InscribeTxs(txHashs []string) (string, error)
	SystemPromptManagerTopup(contractAddr string, prkHex string, agentId int64, amount *big.Int) (string, error)
	SystemPromptManagerMint(contractAddr string, prkHex string, to common.Address, uri string, data []byte, fee *big.Int) (string, error)
	IsContract(address string) (bool, error)
	ConvertAddressForIn(addr string) string
	ConvertAddressForOut(addr string) string
	Erc721Transfer(contractAddr string, prkHex string, toAddr string, tokenId *big.Int) (string, error)
	DeployERC20RealWorldAgent(prkHex string, name string, symbol string, amount *big.Int, recipient common.Address, minFeeToUse *big.Int, timeout uint32, tokenFee common.Address, worker common.Address) (string, string, error)
	DeployERC20UtilityAgent(prkHex string, name string, symbol string, amount *big.Int, recipient common.Address, systemPrompt string, storageInfo []byte) (string, string, error)
	DeployTransparentUpgradeableProxy(prkHex string, logic common.Address, admin common.Address, data []byte) (string, string, error)
}

func AgentUpgradeableInitializeData(agentName string, agentVersion string, codeLanguage string, pointers []agentupgradeable.IAgentCodePointer, depsAgents []common.Address, agentOwner common.Address, registrar common.Address, resolver common.Address, duration uint) ([]byte, error) {
	typeAddress, err := abi.NewType("address", "", nil)
	if err != nil {
		panic(err)
	}
	typeUint256, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(err)
	}
	arguments := abi.Arguments{
		{Type: typeAddress},
		{Type: typeAddress},
		{Type: typeUint256},
	}
	nameService, err := arguments.Pack(registrar, resolver, big.NewInt(int64(duration)))
	if err != nil {
		panic(err)
	}
	instanceABI, err := agentupgradeable.AgentUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	dataBytes, err := instanceABI.Pack(
		"initialize",
		agentName, agentVersion, codeLanguage, pointers, depsAgents, agentOwner, nameService,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
