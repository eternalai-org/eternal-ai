// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model_collection

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ModelCollectionMetaData contains all meta data concerning the ModelCollection contract.
var ModelCollectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerDeauthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ModelURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"NewModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"WEAITokenUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"authorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"checkModelExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"deauthorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"royaltyPortion_\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"nextModelId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateModelURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"updateWEAIToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001657613e10908161001c8239f35b600080fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c8062728e4614612b5f57806301ffc9a7146129d65780630305ea011461290f5780630387da42146114c557806306fdde031461284b578063081812fc1461280f578063095ea7b31461260457806311d7beb21461150157806317f899631461254857806318160ddd1461250b5780631950a503146102b857806319e93993146124505780631c3ff82f14610f815780631ec60b1714611c5d57806323b872dd14611c39578063267c850714611b4657806329dc4d9b14611a9b5780632a55205a146119f55780632f745c59146118eb578063376d28e6146104de5780633f4ba83a1461184f57806342842e0e1461181c5780634f6ccce714611742578063534f3b4d1461169e57806354fd4d50146115e05780635c975abb1461159e5780636352211e14611544578063637ecfc8146115015780636817c76c146114c557806370a082311461147c578063715018a6146113dc57806376d1493f146113725780638456cb59146112d357806384b0196e14610fd3578063871c15b114610f815780638da5cb5b14610f2e57806395d89b4114610e105780639fbc87131461048c578063a22cb46514610cdc578063b88d4fde14610c3d578063c87b56dd14610ae2578063d0def5211461051a578063e472ae8b146104de578063e637f98f1461048c578063e985e9c51461040a578063f2fde38b14610322578063f3ae2415146102b85763f3fef3a30361000e57346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357600080808061026b612bc9565b610273612dc9565b602435905af1610281613611565b501561028957005b60046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b600080fd5b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35773ffffffffffffffffffffffffffffffffffffffff610304612bc9565b16600052609a602052602060ff604060002054166040519015158152f35b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357610359612bc9565b610361612dc9565b73ffffffffffffffffffffffffffffffffffffffff8116156103865761001990612e49565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357610441612bc9565b610449612bec565b9073ffffffffffffffffffffffffffffffffffffffff809116600052606a60205260406000209116600052602052602060ff604060002054166040519015158152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602073ffffffffffffffffffffffffffffffffffffffff60995416604051908152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576020609754604051908152f35b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357610551612bc9565b60249067ffffffffffffffff9082358281116102b357610575903690600401612c6d565b9073ffffffffffffffffffffffffffffffffffffffff91826101ca541633141580610ac9575b610a9f576097546105ab81613a63565b6097555b6105dc81600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b156105ef576105ea90613a63565b6105af565b939463ffffffff851015610a755760985480610922575b50506040519261061584612cfb565b6000845285169283156108c55761065861065286600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b15613ad6565b61010495865496866000526020976101058952806040600020556801000000000000000081101561089757610697816106cd9360018b9401905561389b565b9091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b6106d681612eb7565b85600052610102885260406000208160005288528660406000205586600052610103885260406000205560ff610134541661081457859493926107e76107e27fe9483618ed88dacb391de5ab755452820de95aad7cca806fddd79e1768d3eb4994886108099561076c61065283600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b8960005260688d526040600020600181540190558160005260678d5260406000208a7fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055818a60007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4613641565b613134565b6107fb6107f5368387612d92565b87613b3b565b604051918291339583613a90565b0390a4604051908152f35b608487602b8a604051927f08c379a000000000000000000000000000000000000000000000000000000000845260048401528201527f4552433732315061757361626c653a20746f6b656e207472616e73666572207760448201527f68696c65207061757365640000000000000000000000000000000000000000006064820152fd5b897f4e487b710000000000000000000000000000000000000000000000000000000060005260416004526000fd5b6064876020604051917f08c379a00000000000000000000000000000000000000000000000000000000083528160048401528201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152fd5b84609b54166040516060810181811085821117610a47577fffffffff00000000000000000000000000000000000000000000000000000000916025916040528181527f7432353629000000000000000000000000000000000000000000000000000000604060208301927f7472616e7366657246726f6d28616464726573732c616464726573732c75696e8452015220166040519260208401918252338b85015230604485015260648401526064835260a083019383851090851117610897576000809493819460405251925af16109f8613611565b9015908115610a0e575b50610289578680610606565b8051801515925082610a23575b505087610a02565b81925090602091810103126102b357602001518015908115036102b3578780610a1b565b8a7f4e487b710000000000000000000000000000000000000000000000000000000060005260416004526000fd5b60046040517faa7feadc000000000000000000000000000000000000000000000000000000008152fd5b60046040517f82b42900000000000000000000000000000000000000000000000000000000008152fd5b5033600052609a60205260ff604060002054161561059b565b346102b3576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357600435610b4a610b4582600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b612f68565b6000526101988152604060002090604051918260008254610b6a81612fff565b93848452600191868382169182600014610bfc575050600114610bbd575b5050610b9692500383612d17565b6000604051610ba481612cfb565b52610bb9604051928284938452830190612c0f565b0390f35b85925060005281600020906000915b858310610be4575050610b9693508201018580610b88565b80548389018501528794508693909201918101610bcc565b91509350610b969592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201018580610b88565b346102b35760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357610c74612bc9565b610c7c612bec565b906064359060443567ffffffffffffffff83116102b357366023840112156102b35761001993610cb96107e2943690602481600401359101612d92565b92610ccc610cc784336131c0565b6130a9565b610cd78383836132c6565b6137fb565b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357610d13612bc9565b602435908115158092036102b35773ffffffffffffffffffffffffffffffffffffffff1690813314610db25733600052606a60205260406000208260005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152fd5b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576040516000606654610e5081612fff565b80845290600190818116908115610ee95750600114610e8e575b610bb984610e7a81860382612d17565b604051918291602083526020830190612c0f565b6066600090815292507f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943545b828410610ed1575050508101602001610e7a82610e6a565b80546020858701810191909152909301928101610eb9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b85019092019250610e7a9150839050610e6a565b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602073ffffffffffffffffffffffffffffffffffffffff6101ca5416604051908152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602073ffffffffffffffffffffffffffffffffffffffff609b5416604051908152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760ce5415806112c9575b1561126b5760405160d0549060008161102484612fff565b91828252600194858116908160001461122e57506001146111cd575b61104c92500382612d17565b6040519060008260d1549161106083612fff565b80835292868116908115611190575060011461112f575b61108692509492940383612d17565b6040519261109384612cfb565b600084526110e5604051937f0f0000000000000000000000000000000000000000000000000000000000000085526110d760209360e08588015260e0870190612c0f565b908582036040870152612c0f565b466060850152306080850152600060a085015283810360c085015281808651928381520195019160005b82811061111c5785870386f35b835187529581019592810192840161110f565b509060d16000527f695fb3134ad82c3b8022bc5464edd0bcc9424ef672b52245dcb6ab2374327ce390856000925b8284106111765750505090602061108692820101611077565b60209294508054838589010152019101909185859361115d565b602092506110869491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101611077565b509060d06000527fe89d44c8fd6a9bac8af33ce47f56337617d449bf7ff3956b618c646de829cbcb90846000925b8284106112145750505090602061104c92820101611040565b6020929450805483858801015201910190918484936111fb565b6020925061104c9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101611040565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b5060cf541561100c565b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35761130a612dc9565b61131261398c565b61131a61398c565b61013460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760206113d2600435600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b6040519015158152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357611413612dc9565b600073ffffffffffffffffffffffffffffffffffffffff6101ca8054907fffffffffffffffffffffffff000000000000000000000000000000000000000082169055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760206114bd6114b8612bc9565b612eb7565b604051908152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576020609854604051908152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602061ffff60995460a01c16604051908152f35b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576020611580600435612fcd565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602060ff61013454166040519015158152f35b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357604051604081019080821067ffffffffffffffff83111761166f57610bb991604052600681527f76302e302e3100000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190612c0f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760043560243567ffffffffffffffff81116102b3576117117f8a3c942991b9dbc6aa087b76b9ec1abeae3454615ece41c7da7e5b04623a096b913690600401612c6d565b9061171a612dc9565b61172e611728368484612d92565b85613b3b565b61173d60405192839283613a90565b0390a2005b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357600435610104548110156117985761178a60209161389b565b90546040519160031b1c8152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152fd5b346102b3576100196107e261183036612c9b565b906040519261183e84612cfb565b60008452610ccc610cc784336131c0565b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357611886612dc9565b61188e6139f7565b6118966139f7565b6101347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357611922612bc9565b6024359061192f81612eb7565b8210156119715773ffffffffffffffffffffffffffffffffffffffff166000526101026020526040600020906000526020526020604060002054604051908152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152fd5b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760243560995461ffff8160a01c1691828102928184041490151715611a6c5761271060409273ffffffffffffffffffffffffffffffffffffffff845193168352046020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3577fec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3602073ffffffffffffffffffffffffffffffffffffffff611b0a612bc9565b611b12612dc9565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006099541617609955604051908152a1005b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35773ffffffffffffffffffffffffffffffffffffffff611b92612bc9565b611b9a612dc9565b1680600052609a60205260ff60406000205416611c0f5780600052609a602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b600080a2005b60046040517feacfc0ae000000000000000000000000000000000000000000000000000000008152fd5b346102b357610019611c4a36612c9b565b91611c58610cc784336131c0565b6132c6565b346102b35760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760043567ffffffffffffffff81116102b357611cac903690600401612c6d565b9060243567ffffffffffffffff81116102b357611ccd903690600401612c6d565b92906064359173ffffffffffffffffffffffffffffffffffffffff831683036102b35761ffff60843516608435036102b35760c4359173ffffffffffffffffffffffffffffffffffffffff831683036102b35760005460ff8160081c161595868097612443575b801561242c575b156123a857611d8793828860017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00611d7f961617600055612379575b503691612d92565b953691612d92565b93611da260ff60005460081c16611d9d81613901565b613901565b80519067ffffffffffffffff821161166f578190611dc1606554612fff565b601f81116122cf575b50602090601f83116001146121eb576000926121e0575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176065555b835167ffffffffffffffff811161166f57611e30606654612fff565b601f811161213d575b50602094601f821160011461207d57948192939495600092612072575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176066555b611ed960ff60005460081c16611e9e81613901565b611ea781613901565b6101347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008154169055611d9d81613901565b611ee233612e49565b73ffffffffffffffffffffffffffffffffffffffff8216158015612054575b610a755760a43563ffffffff811015610a755773ffffffffffffffffffffffffffffffffffffffff926044356098556099547fffffffffffffffffffff000000000000000000000000000000000000000000008575ffff000000000000000000000000000000000000000060843560a01b16931691161717609955609755167fffffffffffffffffffffffff0000000000000000000000000000000000000000609b541617609b5573ffffffffffffffffffffffffffffffffffffffff6101ca5416600052609a602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055611fff57005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff600054166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b5073ffffffffffffffffffffffffffffffffffffffff811615611f01565b015190508580611e56565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169560666000527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943549160005b888110612125575083600195969798106120ee575b505050811b01606655611e89565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558580806120e0565b919260206001819286850151815501940192016120cb565b6066600052601f820160051c7f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540190602083106121b8575b601f0160051c7f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e9435401905b8181106121ac5750611e39565b6000815560010161219f565b7f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943549150612175565b015190508680611de1565b925060656000527f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c7906000935b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841685106122b45760019450837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081161061227d575b505050811b01606555611e14565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905586808061226f565b81810151835560209485019460019093019290910190612218565b9091506065600052601f830160051c7f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c70160208410612352575b908392915b601f820160051c7f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c70181106123435750611dca565b6000815584935060010161230e565b507f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c7612309565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011760005589611d77565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b158015611d3b5750600160ff831614611d3b565b50600160ff831610611d34565b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35760043561ffff8116908181036102b3577fb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5916020916124bc612dc9565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006099549260a01b16911617609955604051908152a1005b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357602061010454604051908152f35b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35761257f612bc9565b612587612dc9565b73ffffffffffffffffffffffffffffffffffffffff809116908115610a75577fffffffffffffffffffffffff0000000000000000000000000000000000000000907f9ab45fd23d2134d8834df8b027636cc1969ef5b9950b4f73bbdcf984cc4cc0736040609b549281519084168152856020820152a11617609b55005b346102b35760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35761263b612bc9565b6024359073ffffffffffffffffffffffffffffffffffffffff808061265f85612fcd565b1692169180831461278b57803314908115612766575b50156126e2578260005260696020526040600020827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790556126ba83612fcd565b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600080a4005b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152fd5b9050600052606a60205260406000203360005260205260ff6040600020541684612675565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f72000000000000000000000000000000000000000000000000000000000000006064820152fd5b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576020611580600435613052565b346102b35760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b357604051600060655461288b81612fff565b80845290600190818116908115610ee957506001146128b457610bb984610e7a81860382612d17565b6065600090815292507f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c75b8284106128f7575050508101602001610e7a82610e6a565b805460208587018101919091529093019281016128df565b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b35773ffffffffffffffffffffffffffffffffffffffff61295b612bc9565b612963612dc9565b1680600052609a60205260ff6040600020541615610a9f5780600052609a60205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c600080a2005b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036102b357807f2a55205a0000000000000000000000000000000000000000000000000000000060209214908115612a6b575b506040519015158152f35b7f4906490600000000000000000000000000000000000000000000000000000000811491508115612a9e575b5082612a60565b7f780e9d6300000000000000000000000000000000000000000000000000000000811491508115612ad1575b5082612a97565b7f80ac58cd00000000000000000000000000000000000000000000000000000000811491508115612b35575b8115612b0b575b5082612aca565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501482612b04565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150612afd565b346102b35760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102b3577f23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e866020600435612bbc612dc9565b80609855604051908152a1005b6004359073ffffffffffffffffffffffffffffffffffffffff821682036102b357565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036102b357565b919082519283825260005b848110612c595750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b602081830181015184830182015201612c1a565b9181601f840112156102b35782359167ffffffffffffffff83116102b357602083818601950101116102b357565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126102b35773ffffffffffffffffffffffffffffffffffffffff9060043582811681036102b3579160243590811681036102b3579060443590565b6020810190811067ffffffffffffffff82111761166f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761166f57604052565b67ffffffffffffffff811161166f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612d9e82612d58565b91612dac6040519384612d17565b8294818452818301116102b3578281602093846000960137010152565b73ffffffffffffffffffffffffffffffffffffffff6101ca54163303612deb57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6101ca90815473ffffffffffffffffffffffffffffffffffffffff80921692837fffffffffffffffffffffffff00000000000000000000000000000000000000008316179055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b73ffffffffffffffffffffffffffffffffffffffff168015612ee457600052606860205260406000205490565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152fd5b15612f6f57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152fd5b600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416612ffc811515612f68565b90565b90600182811c92168015613048575b602083101461301957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161300e565b613082610b4582600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b600052606960205273ffffffffffffffffffffffffffffffffffffffff6040600020541690565b156130b057565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206f7220617070726f766564000000000000000000000000000000000000006064820152fd5b1561313b57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608490fd5b9073ffffffffffffffffffffffffffffffffffffffff80806131e184612fcd565b16931691838314938415613214575b5083156131fe575b50505090565b61320a91929350613052565b16143880806131f8565b909350600052606a60205260406000208260005260205260ff6040600020541692386131f0565b1561324257565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152fd5b906132f9906132d484612fcd565b73ffffffffffffffffffffffffffffffffffffffff848116939092918316841461323b565b81811693841561358e57836134d55750610104805486600052610105602052806040600020556801000000000000000081101561166f57610697816133449360018a9401905561389b565b8284036134a0575b5060ff610134541661341c578161336d9161336686612fcd565b161461323b565b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60008481526069602052604081207fffffffffffffffffffffffff0000000000000000000000000000000000000000908181541690558382526068602052604082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190558482526040822060018154019055858252606760205284604083209182541617905580a4565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4552433732315061757361626c653a20746f6b656e207472616e73666572207760448201527f68696c65207061757365640000000000000000000000000000000000000000006064820152fd5b6134a990612eb7565b60406000858152610102602052818120838252602052868282205586815261010360205220553861334c565b8484036134e3575b50613344565b6134ec90612eb7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111611a6c576000908682526101039060208281526040928385205490838203613557575b50508884528383812055868452610102815282842091845252812055386134dd565b888652610102808452858720858852845285872054908a885284528587208388528452808688205586528252838520553880613535565b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b3d1561363c573d9061362282612d58565b916136306040519384612d17565b82523d6000602084013e565b606090565b909190803b156137f35760206040518092817f150b7a02000000000000000000000000000000000000000000000000000000009687825233600483015273ffffffffffffffffffffffffffffffffffffffff826136b86000998a948560248501526044840152608060648401526084830190612c0f565b0393165af19082908261378c575b5050613766576136d4613611565b80519081613761576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608490fd5b602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000161490565b909192506020813d82116137eb575b816137a860209383612d17565b810103126137e75751907fffffffff00000000000000000000000000000000000000000000000000000000821682036137e457509038806136c6565b80fd5b5080fd5b3d915061379b565b505050600190565b919290803b15613892576138729160209160405180809581947f150b7a0200000000000000000000000000000000000000000000000000000000998a845233600485015273ffffffffffffffffffffffffffffffffffffffff809a1660248501526044840152608060648401526084830190612c0f565b03916000968791165af19082908261378c575050613766576136d4613611565b50505050600190565b61010480548210156138d2576000527f4c0be60200faa20559308cb7b5a1bb3255c16cb1cab91f525b5ae7a03d02fabe0190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1561390857565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b60ff610134541661399957565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b60ff610134541615613a0557565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611a6c5760010190565b90601f836040947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe093602086528160208701528686013760008582860101520116010190565b15613add57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152fd5b919091613b6b81600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b15613d56576000908082526020916101988352604081209085519067ffffffffffffffff8211613d2957613b9f8354612fff565b601f8111613ce6575b508490601f8311600114613c2657907ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7969783613c1b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790555b604051908152a1565b015190503880613be0565b91967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08816848452868420935b818110613ccf57509160019391897ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7999a9410613c98575b505050811b019055613c12565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613c8b565b929387600181928786015181550195019301613c53565b838252858220601f840160051c810191878510613d1f575b601f0160051c01905b818110613d145750613ba8565b828155600101613d07565b9091508190613cfe565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201527f6578697374656e7420746f6b656e0000000000000000000000000000000000006064820152fdfea2646970667358221220b0e26e0c77cf961bdfecfad09eb55282a1dbbe249bd2e2b46d307502381ea27d64736f6c63430008140033",
}

// ModelCollectionABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelCollectionMetaData.ABI instead.
var ModelCollectionABI = ModelCollectionMetaData.ABI

// ModelCollectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModelCollectionMetaData.Bin instead.
var ModelCollectionBin = ModelCollectionMetaData.Bin

// DeployModelCollection deploys a new Ethereum contract, binding an instance of ModelCollection to it.
func DeployModelCollection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModelCollection, error) {
	parsed, err := ModelCollectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModelCollectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModelCollection{ModelCollectionCaller: ModelCollectionCaller{contract: contract}, ModelCollectionTransactor: ModelCollectionTransactor{contract: contract}, ModelCollectionFilterer: ModelCollectionFilterer{contract: contract}}, nil
}

// ModelCollection is an auto generated Go binding around an Ethereum contract.
type ModelCollection struct {
	ModelCollectionCaller     // Read-only binding to the contract
	ModelCollectionTransactor // Write-only binding to the contract
	ModelCollectionFilterer   // Log filterer for contract events
}

// ModelCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelCollectionSession struct {
	Contract     *ModelCollection  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCollectionCallerSession struct {
	Contract *ModelCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ModelCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelCollectionTransactorSession struct {
	Contract     *ModelCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ModelCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelCollectionRaw struct {
	Contract *ModelCollection // Generic contract binding to access the raw methods on
}

// ModelCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCollectionCallerRaw struct {
	Contract *ModelCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// ModelCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelCollectionTransactorRaw struct {
	Contract *ModelCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModelCollection creates a new instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollection(address common.Address, backend bind.ContractBackend) (*ModelCollection, error) {
	contract, err := bindModelCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModelCollection{ModelCollectionCaller: ModelCollectionCaller{contract: contract}, ModelCollectionTransactor: ModelCollectionTransactor{contract: contract}, ModelCollectionFilterer: ModelCollectionFilterer{contract: contract}}, nil
}

// NewModelCollectionCaller creates a new read-only instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionCaller(address common.Address, caller bind.ContractCaller) (*ModelCollectionCaller, error) {
	contract, err := bindModelCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionCaller{contract: contract}, nil
}

// NewModelCollectionTransactor creates a new write-only instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelCollectionTransactor, error) {
	contract, err := bindModelCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTransactor{contract: contract}, nil
}

// NewModelCollectionFilterer creates a new log filterer instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelCollectionFilterer, error) {
	contract, err := bindModelCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionFilterer{contract: contract}, nil
}

// bindModelCollection binds a generic wrapper to an already deployed contract.
func bindModelCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelCollectionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollection *ModelCollectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollection.Contract.ModelCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollection *ModelCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.Contract.ModelCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollection *ModelCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollection.Contract.ModelCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollection *ModelCollectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollection *ModelCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollection *ModelCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollection.Contract.contract.Transact(opts, method, params...)
}

// IsManager is a free data retrieval call binding the contract method 0x1950a503.
//
// Solidity: function _isManager(address ) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0x1950a503.
//
// Solidity: function _isManager(address ) view returns(bool)
func (_ModelCollection *ModelCollectionSession) IsManager(arg0 common.Address) (bool, error) {
	return _ModelCollection.Contract.IsManager(&_ModelCollection.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0x1950a503.
//
// Solidity: function _isManager(address ) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _ModelCollection.Contract.IsManager(&_ModelCollection.CallOpts, arg0)
}

// MintPrice is a free data retrieval call binding the contract method 0x0387da42.
//
// Solidity: function _mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x0387da42.
//
// Solidity: function _mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) MintPrice() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice(&_ModelCollection.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x0387da42.
//
// Solidity: function _mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) MintPrice() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice(&_ModelCollection.CallOpts)
}

// NextModelId0 is a free data retrieval call binding the contract method 0x376d28e6.
//
// Solidity: function _nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) NextModelId0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_nextModelId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextModelId0 is a free data retrieval call binding the contract method 0x376d28e6.
//
// Solidity: function _nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) NextModelId0() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId0(&_ModelCollection.CallOpts)
}

// NextModelId0 is a free data retrieval call binding the contract method 0x376d28e6.
//
// Solidity: function _nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) NextModelId0() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId0(&_ModelCollection.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x637ecfc8.
//
// Solidity: function _royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCaller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x637ecfc8.
//
// Solidity: function _royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionSession) RoyaltyPortion() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion(&_ModelCollection.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x637ecfc8.
//
// Solidity: function _royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyPortion() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0xe637f98f.
//
// Solidity: function _royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCaller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0xe637f98f.
//
// Solidity: function _royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionSession) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0xe637f98f.
//
// Solidity: function _royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver(&_ModelCollection.CallOpts)
}

// WEAIToken0 is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionCaller) WEAIToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken0 is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionSession) WEAIToken0() (common.Address, error) {
	return _ModelCollection.Contract.WEAIToken0(&_ModelCollection.CallOpts)
}

// WEAIToken0 is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) WEAIToken0() (common.Address, error) {
	return _ModelCollection.Contract.WEAIToken0(&_ModelCollection.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollection.Contract.BalanceOf(&_ModelCollection.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollection.Contract.BalanceOf(&_ModelCollection.CallOpts, owner)
}

// CheckModelExist is a free data retrieval call binding the contract method 0x76d1493f.
//
// Solidity: function checkModelExist(uint256 modelId) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) CheckModelExist(opts *bind.CallOpts, modelId *big.Int) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "checkModelExist", modelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckModelExist is a free data retrieval call binding the contract method 0x76d1493f.
//
// Solidity: function checkModelExist(uint256 modelId) view returns(bool)
func (_ModelCollection *ModelCollectionSession) CheckModelExist(modelId *big.Int) (bool, error) {
	return _ModelCollection.Contract.CheckModelExist(&_ModelCollection.CallOpts, modelId)
}

// CheckModelExist is a free data retrieval call binding the contract method 0x76d1493f.
//
// Solidity: function checkModelExist(uint256 modelId) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) CheckModelExist(modelId *big.Int) (bool, error) {
	return _ModelCollection.Contract.CheckModelExist(&_ModelCollection.CallOpts, modelId)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollection.Contract.Eip712Domain(&_ModelCollection.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollection.Contract.Eip712Domain(&_ModelCollection.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.GetApproved(&_ModelCollection.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.GetApproved(&_ModelCollection.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollection.Contract.IsApprovedForAll(&_ModelCollection.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollection.Contract.IsApprovedForAll(&_ModelCollection.CallOpts, owner, operator)
}

// IsManager0 is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) IsManager0(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "isManager", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager0 is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_ModelCollection *ModelCollectionSession) IsManager0(account common.Address) (bool, error) {
	return _ModelCollection.Contract.IsManager0(&_ModelCollection.CallOpts, account)
}

// IsManager0 is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) IsManager0(account common.Address) (bool, error) {
	return _ModelCollection.Contract.IsManager0(&_ModelCollection.CallOpts, account)
}

// MintPrice0 is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) MintPrice0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice0 is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) MintPrice0() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice0(&_ModelCollection.CallOpts)
}

// MintPrice0 is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) MintPrice0() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice0(&_ModelCollection.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionSession) Name() (string, error) {
	return _ModelCollection.Contract.Name(&_ModelCollection.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Name() (string, error) {
	return _ModelCollection.Contract.Name(&_ModelCollection.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) NextModelId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "nextModelId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) NextModelId() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId(&_ModelCollection.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) NextModelId() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId(&_ModelCollection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionSession) Owner() (common.Address, error) {
	return _ModelCollection.Contract.Owner(&_ModelCollection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) Owner() (common.Address, error) {
	return _ModelCollection.Contract.Owner(&_ModelCollection.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.OwnerOf(&_ModelCollection.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.OwnerOf(&_ModelCollection.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionSession) Paused() (bool, error) {
	return _ModelCollection.Contract.Paused(&_ModelCollection.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) Paused() (bool, error) {
	return _ModelCollection.Contract.Paused(&_ModelCollection.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 modelId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_ModelCollection *ModelCollectionCaller) RoyaltyInfo(opts *bind.CallOpts, modelId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyInfo", modelId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 modelId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_ModelCollection *ModelCollectionSession) RoyaltyInfo(modelId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ModelCollection.Contract.RoyaltyInfo(&_ModelCollection.CallOpts, modelId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 modelId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyInfo(modelId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ModelCollection.Contract.RoyaltyInfo(&_ModelCollection.CallOpts, modelId, salePrice)
}

// RoyaltyPortion0 is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCaller) RoyaltyPortion0(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion0 is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionSession) RoyaltyPortion0() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion0(&_ModelCollection.CallOpts)
}

// RoyaltyPortion0 is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyPortion0() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion0(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver0 is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCaller) RoyaltyReceiver0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver0 is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionSession) RoyaltyReceiver0() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver0(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver0 is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyReceiver0() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver0(&_ModelCollection.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ModelCollection.Contract.SupportsInterface(&_ModelCollection.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ModelCollection.Contract.SupportsInterface(&_ModelCollection.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionSession) Symbol() (string, error) {
	return _ModelCollection.Contract.Symbol(&_ModelCollection.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Symbol() (string, error) {
	return _ModelCollection.Contract.Symbol(&_ModelCollection.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenByIndex(&_ModelCollection.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenByIndex(&_ModelCollection.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenOfOwnerByIndex(&_ModelCollection.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenOfOwnerByIndex(&_ModelCollection.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 modelId) view returns(string)
func (_ModelCollection *ModelCollectionCaller) TokenURI(opts *bind.CallOpts, modelId *big.Int) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenURI", modelId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 modelId) view returns(string)
func (_ModelCollection *ModelCollectionSession) TokenURI(modelId *big.Int) (string, error) {
	return _ModelCollection.Contract.TokenURI(&_ModelCollection.CallOpts, modelId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 modelId) view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) TokenURI(modelId *big.Int) (string, error) {
	return _ModelCollection.Contract.TokenURI(&_ModelCollection.CallOpts, modelId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TotalSupply() (*big.Int, error) {
	return _ModelCollection.Contract.TotalSupply(&_ModelCollection.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TotalSupply() (*big.Int, error) {
	return _ModelCollection.Contract.TotalSupply(&_ModelCollection.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionSession) Version() (string, error) {
	return _ModelCollection.Contract.Version(&_ModelCollection.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Version() (string, error) {
	return _ModelCollection.Contract.Version(&_ModelCollection.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x1c3ff82f.
//
// Solidity: function wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x1c3ff82f.
//
// Solidity: function wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionSession) WEAIToken() (common.Address, error) {
	return _ModelCollection.Contract.WEAIToken(&_ModelCollection.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x1c3ff82f.
//
// Solidity: function wEAIToken() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) WEAIToken() (common.Address, error) {
	return _ModelCollection.Contract.WEAIToken(&_ModelCollection.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Approve(&_ModelCollection.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Approve(&_ModelCollection.TransactOpts, to, tokenId)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionTransactor) AuthorizeManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "authorizeManager", account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionSession) AuthorizeManager(account common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.AuthorizeManager(&_ModelCollection.TransactOpts, account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionTransactorSession) AuthorizeManager(account common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.AuthorizeManager(&_ModelCollection.TransactOpts, account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionTransactor) DeauthorizeManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "deauthorizeManager", account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionSession) DeauthorizeManager(account common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.DeauthorizeManager(&_ModelCollection.TransactOpts, account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_ModelCollection *ModelCollectionTransactorSession) DeauthorizeManager(account common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.DeauthorizeManager(&_ModelCollection.TransactOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1ec60b17.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextModelId_, address wEAIToken_) returns()
func (_ModelCollection *ModelCollectionTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextModelId_ *big.Int, wEAIToken_ common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "initialize", name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextModelId_, wEAIToken_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1ec60b17.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextModelId_, address wEAIToken_) returns()
func (_ModelCollection *ModelCollectionSession) Initialize(name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextModelId_ *big.Int, wEAIToken_ common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.Initialize(&_ModelCollection.TransactOpts, name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextModelId_, wEAIToken_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1ec60b17.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextModelId_, address wEAIToken_) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Initialize(name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextModelId_ *big.Int, wEAIToken_ common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.Initialize(&_ModelCollection.TransactOpts, name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextModelId_, wEAIToken_)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) returns(uint256)
func (_ModelCollection *ModelCollectionTransactor) Mint(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "mint", to, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) returns(uint256)
func (_ModelCollection *ModelCollectionSession) Mint(to common.Address, uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.Mint(&_ModelCollection.TransactOpts, to, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) returns(uint256)
func (_ModelCollection *ModelCollectionTransactorSession) Mint(to common.Address, uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.Mint(&_ModelCollection.TransactOpts, to, uri)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionSession) Pause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Pause(&_ModelCollection.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionTransactorSession) Pause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Pause(&_ModelCollection.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionSession) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollection.Contract.RenounceOwnership(&_ModelCollection.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollection.Contract.RenounceOwnership(&_ModelCollection.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom0(&_ModelCollection.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom0(&_ModelCollection.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.Contract.SetApprovalForAll(&_ModelCollection.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.Contract.SetApprovalForAll(&_ModelCollection.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferOwnership(&_ModelCollection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferOwnership(&_ModelCollection.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionSession) Unpause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Unpause(&_ModelCollection.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionTransactorSession) Unpause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Unpause(&_ModelCollection.TransactOpts)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 newPrice) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateMintPrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateMintPrice", newPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 newPrice) returns()
func (_ModelCollection *ModelCollectionSession) UpdateMintPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateMintPrice(&_ModelCollection.TransactOpts, newPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 newPrice) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateMintPrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateMintPrice(&_ModelCollection.TransactOpts, newPrice)
}

// UpdateModelURI is a paid mutator transaction binding the contract method 0x534f3b4d.
//
// Solidity: function updateModelURI(uint256 modelId, string uri) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateModelURI(opts *bind.TransactOpts, modelId *big.Int, uri string) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateModelURI", modelId, uri)
}

// UpdateModelURI is a paid mutator transaction binding the contract method 0x534f3b4d.
//
// Solidity: function updateModelURI(uint256 modelId, string uri) returns()
func (_ModelCollection *ModelCollectionSession) UpdateModelURI(modelId *big.Int, uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateModelURI(&_ModelCollection.TransactOpts, modelId, uri)
}

// UpdateModelURI is a paid mutator transaction binding the contract method 0x534f3b4d.
//
// Solidity: function updateModelURI(uint256 modelId, string uri) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateModelURI(modelId *big.Int, uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateModelURI(&_ModelCollection.TransactOpts, modelId, uri)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 newPortion) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, newPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateRoyaltyPortion", newPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 newPortion) returns()
func (_ModelCollection *ModelCollectionSession) UpdateRoyaltyPortion(newPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyPortion(&_ModelCollection.TransactOpts, newPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 newPortion) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateRoyaltyPortion(newPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyPortion(&_ModelCollection.TransactOpts, newPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address newReceiver) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateRoyaltyReceiver", newReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address newReceiver) returns()
func (_ModelCollection *ModelCollectionSession) UpdateRoyaltyReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyReceiver(&_ModelCollection.TransactOpts, newReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address newReceiver) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateRoyaltyReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyReceiver(&_ModelCollection.TransactOpts, newReceiver)
}

// UpdateWEAIToken is a paid mutator transaction binding the contract method 0x17f89963.
//
// Solidity: function updateWEAIToken(address newToken) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateWEAIToken(opts *bind.TransactOpts, newToken common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateWEAIToken", newToken)
}

// UpdateWEAIToken is a paid mutator transaction binding the contract method 0x17f89963.
//
// Solidity: function updateWEAIToken(address newToken) returns()
func (_ModelCollection *ModelCollectionSession) UpdateWEAIToken(newToken common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateWEAIToken(&_ModelCollection.TransactOpts, newToken)
}

// UpdateWEAIToken is a paid mutator transaction binding the contract method 0x17f89963.
//
// Solidity: function updateWEAIToken(address newToken) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateWEAIToken(newToken common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateWEAIToken(&_ModelCollection.TransactOpts, newToken)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 value) returns()
func (_ModelCollection *ModelCollectionTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "withdraw", to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 value) returns()
func (_ModelCollection *ModelCollectionSession) Withdraw(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Withdraw(&_ModelCollection.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 value) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Withdraw(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Withdraw(&_ModelCollection.TransactOpts, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionSession) Receive() (*types.Transaction, error) {
	return _ModelCollection.Contract.Receive(&_ModelCollection.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionTransactorSession) Receive() (*types.Transaction, error) {
	return _ModelCollection.Contract.Receive(&_ModelCollection.TransactOpts)
}

// ModelCollectionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ModelCollection contract.
type ModelCollectionApprovalIterator struct {
	Event *ModelCollectionApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionApproval represents a Approval event raised by the ModelCollection contract.
type ModelCollectionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ModelCollectionApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionApprovalIterator{contract: _ModelCollection.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ModelCollectionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionApproval)
				if err := _ModelCollection.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseApproval(log types.Log) (*ModelCollectionApproval, error) {
	event := new(ModelCollectionApproval)
	if err := _ModelCollection.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ModelCollection contract.
type ModelCollectionApprovalForAllIterator struct {
	Event *ModelCollectionApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionApprovalForAll represents a ApprovalForAll event raised by the ModelCollection contract.
type ModelCollectionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ModelCollectionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionApprovalForAllIterator{contract: _ModelCollection.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ModelCollectionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionApprovalForAll)
				if err := _ModelCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) ParseApprovalForAll(log types.Log) (*ModelCollectionApprovalForAll, error) {
	event := new(ModelCollectionApprovalForAll)
	if err := _ModelCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ModelCollection contract.
type ModelCollectionBatchMetadataUpdateIterator struct {
	Event *ModelCollectionBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ModelCollection contract.
type ModelCollectionBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionBatchMetadataUpdateIterator{contract: _ModelCollection.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionBatchMetadataUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseBatchMetadataUpdate(log types.Log) (*ModelCollectionBatchMetadataUpdate, error) {
	event := new(ModelCollectionBatchMetadataUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ModelCollection contract.
type ModelCollectionEIP712DomainChangedIterator struct {
	Event *ModelCollectionEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionEIP712DomainChanged represents a EIP712DomainChanged event raised by the ModelCollection contract.
type ModelCollectionEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ModelCollectionEIP712DomainChangedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionEIP712DomainChangedIterator{contract: _ModelCollection.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ModelCollectionEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionEIP712DomainChanged)
				if err := _ModelCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) ParseEIP712DomainChanged(log types.Log) (*ModelCollectionEIP712DomainChanged, error) {
	event := new(ModelCollectionEIP712DomainChanged)
	if err := _ModelCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ModelCollection contract.
type ModelCollectionInitializedIterator struct {
	Event *ModelCollectionInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionInitialized represents a Initialized event raised by the ModelCollection contract.
type ModelCollectionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) FilterInitialized(opts *bind.FilterOpts) (*ModelCollectionInitializedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionInitializedIterator{contract: _ModelCollection.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ModelCollectionInitialized) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionInitialized)
				if err := _ModelCollection.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) ParseInitialized(log types.Log) (*ModelCollectionInitialized, error) {
	event := new(ModelCollectionInitialized)
	if err := _ModelCollection.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionManagerAuthorizationIterator is returned from FilterManagerAuthorization and is used to iterate over the raw logs and unpacked data for ManagerAuthorization events raised by the ModelCollection contract.
type ModelCollectionManagerAuthorizationIterator struct {
	Event *ModelCollectionManagerAuthorization // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionManagerAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionManagerAuthorization)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionManagerAuthorization)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionManagerAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionManagerAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionManagerAuthorization represents a ManagerAuthorization event raised by the ModelCollection contract.
type ModelCollectionManagerAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAuthorization is a free log retrieval operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) FilterManagerAuthorization(opts *bind.FilterOpts, account []common.Address) (*ModelCollectionManagerAuthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionManagerAuthorizationIterator{contract: _ModelCollection.contract, event: "ManagerAuthorization", logs: logs, sub: sub}, nil
}

// WatchManagerAuthorization is a free log subscription operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) WatchManagerAuthorization(opts *bind.WatchOpts, sink chan<- *ModelCollectionManagerAuthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionManagerAuthorization)
				if err := _ModelCollection.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerAuthorization is a log parse operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) ParseManagerAuthorization(log types.Log) (*ModelCollectionManagerAuthorization, error) {
	event := new(ModelCollectionManagerAuthorization)
	if err := _ModelCollection.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionManagerDeauthorizationIterator is returned from FilterManagerDeauthorization and is used to iterate over the raw logs and unpacked data for ManagerDeauthorization events raised by the ModelCollection contract.
type ModelCollectionManagerDeauthorizationIterator struct {
	Event *ModelCollectionManagerDeauthorization // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionManagerDeauthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionManagerDeauthorization)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionManagerDeauthorization)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionManagerDeauthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionManagerDeauthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionManagerDeauthorization represents a ManagerDeauthorization event raised by the ModelCollection contract.
type ModelCollectionManagerDeauthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerDeauthorization is a free log retrieval operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) FilterManagerDeauthorization(opts *bind.FilterOpts, account []common.Address) (*ModelCollectionManagerDeauthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionManagerDeauthorizationIterator{contract: _ModelCollection.contract, event: "ManagerDeauthorization", logs: logs, sub: sub}, nil
}

// WatchManagerDeauthorization is a free log subscription operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) WatchManagerDeauthorization(opts *bind.WatchOpts, sink chan<- *ModelCollectionManagerDeauthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionManagerDeauthorization)
				if err := _ModelCollection.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerDeauthorization is a log parse operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_ModelCollection *ModelCollectionFilterer) ParseManagerDeauthorization(log types.Log) (*ModelCollectionManagerDeauthorization, error) {
	event := new(ModelCollectionManagerDeauthorization)
	if err := _ModelCollection.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ModelCollection contract.
type ModelCollectionMetadataUpdateIterator struct {
	Event *ModelCollectionMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionMetadataUpdate represents a MetadataUpdate event raised by the ModelCollection contract.
type ModelCollectionMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionMetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionMetadataUpdateIterator{contract: _ModelCollection.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionMetadataUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseMetadataUpdate(log types.Log) (*ModelCollectionMetadataUpdate, error) {
	event := new(ModelCollectionMetadataUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionMintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the ModelCollection contract.
type ModelCollectionMintPriceUpdateIterator struct {
	Event *ModelCollectionMintPriceUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionMintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionMintPriceUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionMintPriceUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionMintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionMintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionMintPriceUpdate represents a MintPriceUpdate event raised by the ModelCollection contract.
type ModelCollectionMintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*ModelCollectionMintPriceUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionMintPriceUpdateIterator{contract: _ModelCollection.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionMintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionMintPriceUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintPriceUpdate is a log parse operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) ParseMintPriceUpdate(log types.Log) (*ModelCollectionMintPriceUpdate, error) {
	event := new(ModelCollectionMintPriceUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionModelURIUpdateIterator is returned from FilterModelURIUpdate and is used to iterate over the raw logs and unpacked data for ModelURIUpdate events raised by the ModelCollection contract.
type ModelCollectionModelURIUpdateIterator struct {
	Event *ModelCollectionModelURIUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionModelURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionModelURIUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionModelURIUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionModelURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionModelURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionModelURIUpdate represents a ModelURIUpdate event raised by the ModelCollection contract.
type ModelCollectionModelURIUpdate struct {
	ModelId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelURIUpdate is a free log retrieval operation binding the contract event 0x8a3c942991b9dbc6aa087b76b9ec1abeae3454615ece41c7da7e5b04623a096b.
//
// Solidity: event ModelURIUpdate(uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) FilterModelURIUpdate(opts *bind.FilterOpts, modelId []*big.Int) (*ModelCollectionModelURIUpdateIterator, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "ModelURIUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionModelURIUpdateIterator{contract: _ModelCollection.contract, event: "ModelURIUpdate", logs: logs, sub: sub}, nil
}

// WatchModelURIUpdate is a free log subscription operation binding the contract event 0x8a3c942991b9dbc6aa087b76b9ec1abeae3454615ece41c7da7e5b04623a096b.
//
// Solidity: event ModelURIUpdate(uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) WatchModelURIUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionModelURIUpdate, modelId []*big.Int) (event.Subscription, error) {

	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "ModelURIUpdate", modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionModelURIUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "ModelURIUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelURIUpdate is a log parse operation binding the contract event 0x8a3c942991b9dbc6aa087b76b9ec1abeae3454615ece41c7da7e5b04623a096b.
//
// Solidity: event ModelURIUpdate(uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) ParseModelURIUpdate(log types.Log) (*ModelCollectionModelURIUpdate, error) {
	event := new(ModelCollectionModelURIUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "ModelURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionNewModelIterator is returned from FilterNewModel and is used to iterate over the raw logs and unpacked data for NewModel events raised by the ModelCollection contract.
type ModelCollectionNewModelIterator struct {
	Event *ModelCollectionNewModel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionNewModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionNewModel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionNewModel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionNewModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionNewModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionNewModel represents a NewModel event raised by the ModelCollection contract.
type ModelCollectionNewModel struct {
	Caller  common.Address
	Owner   common.Address
	ModelId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewModel is a free log retrieval operation binding the contract event 0xe9483618ed88dacb391de5ab755452820de95aad7cca806fddd79e1768d3eb49.
//
// Solidity: event NewModel(address indexed caller, address indexed owner, uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) FilterNewModel(opts *bind.FilterOpts, caller []common.Address, owner []common.Address, modelId []*big.Int) (*ModelCollectionNewModelIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "NewModel", callerRule, ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionNewModelIterator{contract: _ModelCollection.contract, event: "NewModel", logs: logs, sub: sub}, nil
}

// WatchNewModel is a free log subscription operation binding the contract event 0xe9483618ed88dacb391de5ab755452820de95aad7cca806fddd79e1768d3eb49.
//
// Solidity: event NewModel(address indexed caller, address indexed owner, uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) WatchNewModel(opts *bind.WatchOpts, sink chan<- *ModelCollectionNewModel, caller []common.Address, owner []common.Address, modelId []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "NewModel", callerRule, ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionNewModel)
				if err := _ModelCollection.contract.UnpackLog(event, "NewModel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewModel is a log parse operation binding the contract event 0xe9483618ed88dacb391de5ab755452820de95aad7cca806fddd79e1768d3eb49.
//
// Solidity: event NewModel(address indexed caller, address indexed owner, uint256 indexed modelId, string uri)
func (_ModelCollection *ModelCollectionFilterer) ParseNewModel(log types.Log) (*ModelCollectionNewModel, error) {
	event := new(ModelCollectionNewModel)
	if err := _ModelCollection.contract.UnpackLog(event, "NewModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ModelCollection contract.
type ModelCollectionOwnershipTransferredIterator struct {
	Event *ModelCollectionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionOwnershipTransferred represents a OwnershipTransferred event raised by the ModelCollection contract.
type ModelCollectionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModelCollectionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionOwnershipTransferredIterator{contract: _ModelCollection.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModelCollectionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionOwnershipTransferred)
				if err := _ModelCollection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) ParseOwnershipTransferred(log types.Log) (*ModelCollectionOwnershipTransferred, error) {
	event := new(ModelCollectionOwnershipTransferred)
	if err := _ModelCollection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ModelCollection contract.
type ModelCollectionPausedIterator struct {
	Event *ModelCollectionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionPaused represents a Paused event raised by the ModelCollection contract.
type ModelCollectionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) FilterPaused(opts *bind.FilterOpts) (*ModelCollectionPausedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionPausedIterator{contract: _ModelCollection.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionPaused) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionPaused)
				if err := _ModelCollection.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) ParsePaused(log types.Log) (*ModelCollectionPaused, error) {
	event := new(ModelCollectionPaused)
	if err := _ModelCollection.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionRoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the ModelCollection contract.
type ModelCollectionRoyaltyPortionUpdateIterator struct {
	Event *ModelCollectionRoyaltyPortionUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionRoyaltyPortionUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionRoyaltyPortionUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionRoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the ModelCollection contract.
type ModelCollectionRoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*ModelCollectionRoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionRoyaltyPortionUpdateIterator{contract: _ModelCollection.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionRoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionRoyaltyPortionUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyPortionUpdate is a log parse operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) ParseRoyaltyPortionUpdate(log types.Log) (*ModelCollectionRoyaltyPortionUpdate, error) {
	event := new(ModelCollectionRoyaltyPortionUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionRoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the ModelCollection contract.
type ModelCollectionRoyaltyReceiverUpdateIterator struct {
	Event *ModelCollectionRoyaltyReceiverUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionRoyaltyReceiverUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionRoyaltyReceiverUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionRoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the ModelCollection contract.
type ModelCollectionRoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*ModelCollectionRoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionRoyaltyReceiverUpdateIterator{contract: _ModelCollection.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionRoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionRoyaltyReceiverUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyReceiverUpdate is a log parse operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) ParseRoyaltyReceiverUpdate(log types.Log) (*ModelCollectionRoyaltyReceiverUpdate, error) {
	event := new(ModelCollectionRoyaltyReceiverUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ModelCollection contract.
type ModelCollectionTransferIterator struct {
	Event *ModelCollectionTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionTransfer represents a Transfer event raised by the ModelCollection contract.
type ModelCollectionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ModelCollectionTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTransferIterator{contract: _ModelCollection.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ModelCollectionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionTransfer)
				if err := _ModelCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseTransfer(log types.Log) (*ModelCollectionTransfer, error) {
	event := new(ModelCollectionTransfer)
	if err := _ModelCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ModelCollection contract.
type ModelCollectionUnpausedIterator struct {
	Event *ModelCollectionUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionUnpaused represents a Unpaused event raised by the ModelCollection contract.
type ModelCollectionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ModelCollectionUnpausedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionUnpausedIterator{contract: _ModelCollection.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionUnpaused) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionUnpaused)
				if err := _ModelCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) ParseUnpaused(log types.Log) (*ModelCollectionUnpaused, error) {
	event := new(ModelCollectionUnpaused)
	if err := _ModelCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionWEAITokenUpdateIterator is returned from FilterWEAITokenUpdate and is used to iterate over the raw logs and unpacked data for WEAITokenUpdate events raised by the ModelCollection contract.
type ModelCollectionWEAITokenUpdateIterator struct {
	Event *ModelCollectionWEAITokenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionWEAITokenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionWEAITokenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionWEAITokenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionWEAITokenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionWEAITokenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionWEAITokenUpdate represents a WEAITokenUpdate event raised by the ModelCollection contract.
type ModelCollectionWEAITokenUpdate struct {
	OldToken common.Address
	NewToken common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWEAITokenUpdate is a free log retrieval operation binding the contract event 0x9ab45fd23d2134d8834df8b027636cc1969ef5b9950b4f73bbdcf984cc4cc073.
//
// Solidity: event WEAITokenUpdate(address oldToken, address newToken)
func (_ModelCollection *ModelCollectionFilterer) FilterWEAITokenUpdate(opts *bind.FilterOpts) (*ModelCollectionWEAITokenUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "WEAITokenUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionWEAITokenUpdateIterator{contract: _ModelCollection.contract, event: "WEAITokenUpdate", logs: logs, sub: sub}, nil
}

// WatchWEAITokenUpdate is a free log subscription operation binding the contract event 0x9ab45fd23d2134d8834df8b027636cc1969ef5b9950b4f73bbdcf984cc4cc073.
//
// Solidity: event WEAITokenUpdate(address oldToken, address newToken)
func (_ModelCollection *ModelCollectionFilterer) WatchWEAITokenUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionWEAITokenUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "WEAITokenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionWEAITokenUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "WEAITokenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWEAITokenUpdate is a log parse operation binding the contract event 0x9ab45fd23d2134d8834df8b027636cc1969ef5b9950b4f73bbdcf984cc4cc073.
//
// Solidity: event WEAITokenUpdate(address oldToken, address newToken)
func (_ModelCollection *ModelCollectionFilterer) ParseWEAITokenUpdate(log types.Log) (*ModelCollectionWEAITokenUpdate, error) {
	event := new(ModelCollectionWEAITokenUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "WEAITokenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
