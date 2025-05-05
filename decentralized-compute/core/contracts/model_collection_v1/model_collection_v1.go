// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package model_collection_v1

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

// ModelCollectionV1MetaData contains all meta data concerning the ModelCollectionV1 contract.
var ModelCollectionV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerDeauthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"TokenModelUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"TokenURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"authorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"deauthorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nextModelId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"mintBySignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"modelAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"updateTokenModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001657614209908161001c8239f35b600080fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c8062728e461461258157806301ffc9a7146123f85780630305ea011461233157806306fdde0314612265578063081812fc14612229578063095ea7b31461201e57806311d7beb214611fdb57806318160ddd14611f9e57806318e97fd114611ef357806319e9399314611e3857806323b872dd14611e14578063267c850714611d2157806329dc4d9b14611c765780632a55205a14611bd05780632f745c5914611ac65780633f4ba83a14611a2a57806342842e0e146119f75780634f6ccce71461191d57806354fd4d501461188e5780635c975abb1461184c5780635e68842a1461172e5780636352211e146116d45780636817c76c1461169857806370a0823114611657578063715018a6146115b75780637f006226146115575780638456cb59146114b857806384b0196e1461134f5780638da5cb5b146112fc57806395d89b41146111dd5780639fbc87131461118b578063a22cb46514611057578063b88d4fde14610fb3578063c87b56dd14610e57578063d9759dd114610d1b578063e472ae8b14610cdf578063e985e9c514610c5d578063eb21af3914610be9578063f2fde38b14610b01578063f3ae241514610a97578063f3fef3a314610a13578063fa8509c8146108f95763fe9737640361000e57346108f45760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760043567ffffffffffffffff81116108f4576102479036906004016126d5565b9060243567ffffffffffffffff81116108f4576102689036906004016126d5565b92610271612654565b906084359261ffff841684036108f45760005460ff8160081c1615958680976108e7575b80156108d0575b1561084c576102e893828860017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006102e096161760005561081d575b5036916127fa565b9536916127fa565b9361030360ff60005460081c166102fe81613be5565b613be5565b80519067ffffffffffffffff8211610675578190610322606554612831565b601f8111610770575b50602090601f83116001146106af576000926106a4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176065555b835167ffffffffffffffff811161067557610391606654612831565b601f8111610610575b50602094601f821160011461055057948192939495600092610545575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176066555b61043a60ff60005460081c166103ff81613be5565b61040881613be5565b6101337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556102fe81613be5565b6104433361366f565b60443560995573ffffffffffffffffffffffffffffffffffffffff917fffffffffffffffffffff000000000000000000000000000000000000000000008375ffff0000000000000000000000000000000000000000609a549360a01b16931691161717609a5560a4356098556101c95416600052609b602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556104f057005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff600054166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b0151905038806103b7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169560666000527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943549160005b8881106105f8575083600195969798106105c1575b505050811b016066556103ea565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806105b3565b9192602060018192868501518155019401920161059e565b60666000527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e94354601f830160051c8101916020841061066b575b601f0160051c01905b81811061065f575061039a565b60008155600101610652565b9091508190610649565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b015190503880610342565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169160656000527f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c79260005b8181106107585750908460019594939210610721575b505050811b01606555610375565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610713565b929360206001819287860151815501950193016106fd565b9091506065600052601f830160051c7f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c70190602084106107f5575b90601f8493920160051c7f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c701905b8181106107e6575061032b565b600081558493506001016107d9565b7f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c791506107ab565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001661010117600055386102d8565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b15801561029c5750600160ff83161461029c565b50600160ff831610610295565b600080fd5b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45761092b6125eb565b60243567ffffffffffffffff81116108f45761094b9036906004016126d5565b9091610955612631565b73ffffffffffffffffffffffffffffffffffffffff92836101c95416331415806109fa575b6109d0575b609880548060005260976020528560406000205416156109a9576109a290612cd7565b905561097f565b50506109c89350602094609854936109c085612cd7565b6098556137d4565b604051908152f35b60046040517f82b42900000000000000000000000000000000000000000000000000000000008152fd5b5033600052609b60205260ff604060002054161561097a565b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576000808080610a4f6125eb565b610a57612e5d565b602435905af1610a65612e2d565b5015610a6d57005b60046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45773ffffffffffffffffffffffffffffffffffffffff610ae36125eb565b16600052609b602052602060ff604060002054166040519015158152f35b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457610b386125eb565b610b40612e5d565b73ffffffffffffffffffffffffffffffffffffffff811615610b65576100199061366f565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b346108f45760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457610c206125eb565b60243567ffffffffffffffff81116108f457602091610c466109c89236906004016126d5565b610c4e612631565b91610c57612654565b93612d04565b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457610c946125eb565b610c9c61260e565b9073ffffffffffffffffffffffffffffffffffffffff809116600052606a60205260406000209116600052602052602060ff604060002054166040519015158152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576020609854604051908152f35b346108f45760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457610d526125eb565b60243567ffffffffffffffff81116108f457610d729036906004016126d5565b9091610d7c612631565b610d84612654565b926084359360ff851685036108f457610db4610dbc9560c4359060a43590610daf8588888d8c612d04565b613ec8565b959095613f64565b73ffffffffffffffffffffffffffffffffffffffff94851690851690811480159190610e3d575b50610e13575b609880548060005260976020528560406000205416156109a957610e0c90612cd7565b9055610de9565b60046040517f8baa579f000000000000000000000000000000000000000000000000000000008152fd5b9050600052609b60205260ff604060002054161586610de3565b346108f4576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457600435610ebf610eba82600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b612b8f565b6000526101978152604060002090604051918260008254610edf81612831565b9384845260019186600182169182600014610f72575050600114610f33575b5050610f0c9250038361277f565b6000604051610f1a81612763565b52610f2f604051928284938452830190612677565b0390f35b85925060005281600020906000915b858310610f5a575050610f0c93508201018580610efe565b80548389018501528794508693909201918101610f42565b91509350610f0c9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201018580610efe565b346108f45760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457610fea6125eb565b610ff261260e565b906064359060443567ffffffffffffffff83116108f457366023840112156108f4576100199361102f6110529436906024816004013591016127fa565b9261104261103d843361317c565b612a9e565b61104d838383613282565b613e28565b613748565b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45761108e6125eb565b602435908115158092036108f45773ffffffffffffffffffffffffffffffffffffffff169081331461112d5733600052606a60205260406000208260005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152fd5b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602073ffffffffffffffffffffffffffffffffffffffff609a5416604051908152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760405160665460008261121e83612831565b91828252602093600190856001821691826000146112be575050600114611261575b5061124d9250038361277f565b610f2f604051928284938452830190612677565b84915060666000527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e94354906000915b8583106112a657505061124d935082010185611240565b8054838901850152879450869390920191810161128f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168582015261124d95151560051b85010192508791506112409050565b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602073ffffffffffffffffffffffffffffffffffffffff6101c95416604051908152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760cd5415806114ae575b15611450576113f1611398612884565b6113a0612969565b604051906113ad82612763565b600082526113ff6020916040519586957f0f00000000000000000000000000000000000000000000000000000000000000875260e0602088015260e0870190612677565b908582036040870152612677565b466060850152306080850152600060a085015283810360c08501526020808451928381520193019160005b82811061143957505050500390f35b83518552869550938101939281019260010161142a565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a656400000000000000000000006044820152fd5b5060ce5415611388565b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576114ef612e5d565b6114f76136dd565b6114ff6136dd565b61013360017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576004356000526097602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576115ee612e5d565b600073ffffffffffffffffffffffffffffffffffffffff6101c98054907fffffffffffffffffffffffff000000000000000000000000000000000000000082169055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760206109c86116936125eb565b612c26565b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576020609954604051908152f35b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576020611710600435612bf4565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760043573ffffffffffffffffffffffffffffffffffffffff61177d61260e565b611785612e5d565b1680156117ee5760207fa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d9183600052609782526040600020817fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055604051908152a2005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420746f6b656e206d6f64656c000000000000000000000000006044820152fd5b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602060ff61013354166040519015158152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457604051604081019080821067ffffffffffffffff83111761067557610f2f91604052600681527f76302e302e3100000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190612677565b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576004356101035481101561197357611965602091612b29565b90546040519160031b1c8152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152fd5b346108f457610019611052611a0b36612703565b9060405192611a1984612763565b6000845261104261103d843361317c565b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457611a61612e5d565b611a69613603565b611a71613603565b6101337fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457611afd6125eb565b60243590611b0a81612c26565b821015611b4c5773ffffffffffffffffffffffffffffffffffffffff166000526101016020526040600020906000526020526020604060002054604051908152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152fd5b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602435609a5461ffff8160a01c1691828102928184041490151715611c475761271060409273ffffffffffffffffffffffffffffffffffffffff845193168352046020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4577fec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3602073ffffffffffffffffffffffffffffffffffffffff611ce56125eb565b611ced612e5d565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000609a541617609a55604051908152a1005b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45773ffffffffffffffffffffffffffffffffffffffff611d6d6125eb565b611d75612e5d565b1680600052609b60205260ff60406000205416611dea5780600052609b602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b600080a2005b60046040517feacfc0ae000000000000000000000000000000000000000000000000000000008152fd5b346108f457610019611e2536612703565b91611e3361103d843361317c565b613282565b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760043561ffff8116908181036108f4577fb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd591602091611ea4612e5d565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff0000000000000000000000000000000000000000609a549260a01b16911617609a55604051908152a1005b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45760043560243567ffffffffffffffff81116108f457611f667fc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe59136906004016126d5565b611f6e612e5d565b611f82611f7c3683856127fa565b85612edd565b611f99604051928392602084526020840191612a5f565b0390a2005b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602061010354604051908152f35b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f457602061ffff609a5460a01c16604051908152f35b346108f45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576120556125eb565b6024359073ffffffffffffffffffffffffffffffffffffffff808061207985612bf4565b169216918083146121a557803314908115612180575b50156120fc578260005260696020526040600020827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790556120d483612bf4565b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600080a4005b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152fd5b9050600052606a60205260406000203360005260205260ff604060002054168461208f565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f72000000000000000000000000000000000000000000000000000000000000006064820152fd5b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576020611710600435612a08565b346108f45760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576040516065546000826122a683612831565b91828252602093600190856001821691826000146112be5750506001146122d4575061124d9250038361277f565b84915060656000527f8ff97419363ffd7000167f130ef7168fbea05faf9251824ca5043f113cc6a7c7906000915b85831061231957505061124d935082010185611240565b80548389018501528794508693909201918101612302565b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f45773ffffffffffffffffffffffffffffffffffffffff61237d6125eb565b612385612e5d565b1680600052609b60205260ff60406000205416156109d05780600052609b60205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c600080a2005b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036108f457807f2a55205a000000000000000000000000000000000000000000000000000000006020921490811561248d575b506040519015158152f35b7f49064906000000000000000000000000000000000000000000000000000000008114915081156124c0575b5082612482565b7f780e9d63000000000000000000000000000000000000000000000000000000008114915081156124f3575b50826124b9565b7f80ac58cd00000000000000000000000000000000000000000000000000000000811491508115612557575b811561252d575b50826124ec565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501482612526565b7f5b5e139f000000000000000000000000000000000000000000000000000000008114915061251f565b346108f45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108f4577f23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e8660206004356125de612e5d565b80609955604051908152a1005b6004359073ffffffffffffffffffffffffffffffffffffffff821682036108f457565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036108f457565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036108f457565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036108f457565b919082519283825260005b8481106126c15750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b602081830181015184830182015201612682565b9181601f840112156108f45782359167ffffffffffffffff83116108f457602083818601950101116108f457565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126108f45773ffffffffffffffffffffffffffffffffffffffff9060043582811681036108f4579160243590811681036108f4579060443590565b6020810190811067ffffffffffffffff82111761067557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761067557604052565b67ffffffffffffffff811161067557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612806826127c0565b91612814604051938461277f565b8294818452818301116108f4578281602093846000960137010152565b90600182811c9216801561287a575b602083101461284b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612840565b6040519060008260cf549161289883612831565b8083529260209060019081811690811561292657506001146128c5575b50506128c39250038361277f565b565b91509260cf6000527facb8d954e2cfef495862221e91bd7523613cf8808827cb33edfe4904cc51bf29936000925b82841061290e57506128c394505050810160200138806128b5565b855488850183015294850194879450928101926128f3565b9050602093506128c39592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138806128b5565b6040519060008260d0549161297d83612831565b8083529260209060019081811690811561292657506001146129a75750506128c39250038361277f565b91509260d06000527fe89d44c8fd6a9bac8af33ce47f56337617d449bf7ff3956b618c646de829cbcb936000925b8284106129f057506128c394505050810160200138806128b5565b855488850183015294850194879450928101926129d5565b612a38610eba82600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b600052606960205273ffffffffffffffffffffffffffffffffffffffff6040600020541690565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b15612aa557565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206f7220617070726f766564000000000000000000000000000000000000006064820152fd5b6101038054821015612b60576000527f02c297ab74aad0aede3a1895c857b1f2c71e6a203feb727bec95ac752998cb780190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b15612b9657565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152fd5b600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416612c23811515612b8f565b90565b73ffffffffffffffffffffffffffffffffffffffff168015612c5357600052606860205260406000205490565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611c475760010190565b909392612d7f92604051948593612d46602086019873ffffffffffffffffffffffffffffffffffffffff94858094168b526080604089015260a0880191612a5f565b93166060850152166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261277f565b519020612d8a614162565b612d926141ad565b916040519260208401927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604085015260608401524660808401523060a084015260a0835260c083019183831067ffffffffffffffff8411176106755760429360e29184604052815190207f1901000000000000000000000000000000000000000000000000000000000000855260c282015201522090565b3d15612e58573d90612e3e826127c0565b91612e4c604051938461277f565b82523d6000602084013e565b606090565b73ffffffffffffffffffffffffffffffffffffffff6101c954163303612e7f57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b919091612f0d81600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b156130f8576000908082526020916101978352604081209085519067ffffffffffffffff82116130cb57612f418354612831565b601f8111613088575b508490601f8311600114612fc857907ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7969783612fbd575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790555b604051908152a1565b015190503880612f82565b91967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08816848452868420935b81811061307157509160019391897ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7999a941061303a575b505050811b019055612fb4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061302d565b929387600181928786015181550195019301612ff5565b838252858220601f840160051c8101918785106130c1575b601f0160051c01905b8181106130b65750612f4a565b8281556001016130a9565b90915081906130a0565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201527f6578697374656e7420746f6b656e0000000000000000000000000000000000006064820152fd5b9073ffffffffffffffffffffffffffffffffffffffff808061319d84612bf4565b169316918383149384156131d0575b5083156131ba575b50505090565b6131c691929350612a08565b16143880806131b4565b909350600052606a60205260406000208260005260205260ff6040600020541692386131ac565b156131fe57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152fd5b906132b59061329084612bf4565b73ffffffffffffffffffffffffffffffffffffffff84811693909291831684146131f7565b81811693841561358057836134c75750610103805486600052610104602052806040600020556801000000000000000081101561067557613300816133369360018a94019055612b29565b9091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b828403613492575b5060ff610133541661340e578161335f9161335886612bf4565b16146131f7565b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60008481526069602052604081207fffffffffffffffffffffffff0000000000000000000000000000000000000000908181541690558382526068602052604082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190558482526040822060018154019055858252606760205284604083209182541617905580a4565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4552433732315061757361626c653a20746f6b656e207472616e73666572207760448201527f68696c65207061757365640000000000000000000000000000000000000000006064820152fd5b61349b90612c26565b60406000858152610101602052818120838252602052868282205586815261010260205220553861333e565b8484036134d5575b50613336565b6134de90612c26565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111611c47576000908682526101029060208281526040928385205490838203613549575b50508884528383812055868452610101815282842091845252812055386134cf565b888652610101808452858720858852845285872054908a885284528587208388528452808688205586528252838520553880613527565b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b60ff61013354161561361157565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b6101c990815473ffffffffffffffffffffffffffffffffffffffff80921692837fffffffffffffffffffffffff00000000000000000000000000000000000000008316179055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b60ff61013354166136ea57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b1561374f57565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608490fd5b929073ffffffffffffffffffffffffffffffffffffffff809216908115613bbb576099543410613b9157604092835161380c81612763565b60009182825287169081159788613b345761385361384d8b600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b156140fd565b6101039889548b865260209a6101048c52818a88205568010000000000000000821015613b0757613300828e92600161388e95019055612b29565b15613ad8575b60ff6101335416613a5557613949918a611052926138d861384d83600052606760205273ffffffffffffffffffffffffffffffffffffffff60406000205416151590565b85875260688c528987206001815401905581875260678c52818a8820967fffffffffffffffffffffffff0000000000000000000000000000000000000000978189825416179055887fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8180a4613c70565b61395d6139573688866127fa565b89612edd565b878252609787528385832091825416179055823b15613a52578351907f4d1c23c0000000000000000000000000000000000000000000000000000000008252876004830152808260248183885af18015613a465790889594939291613a01575b50506139f4907f3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef749495845194808652850191612a5f565b948201528033940390a390565b9091939592945067ffffffffffffffff82116130cb575084529285929091807f3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef746139bd565b508451903d90823e3d90fd5b80fd5b6084898851907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602b60248201527f4552433732315061757361626c653a20746f6b656e207472616e73666572207760448201527f68696c65207061757365640000000000000000000000000000000000000000006064820152fd5b613ae181612c26565b8385526101018a528785208186528a528a888620558a85526101028a5287852055613894565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b606487517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152fd5b60046040517f356680b7000000000000000000000000000000000000000000000000000000008152fd5b60046040517f13aa293e000000000000000000000000000000000000000000000000000000008152fd5b15613bec57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b909190803b15613e205760206040518092817f150b7a02000000000000000000000000000000000000000000000000000000009687825233600483015273ffffffffffffffffffffffffffffffffffffffff82613ce76000998a948560248501526044840152608060648401526084830190612677565b0393165af190829082613dbb575b5050613d9557613d03612e2d565b80519081613d90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608490fd5b602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000161490565b909192506020813d602011613e18575b81613dd86020938361277f565b81010312613e145751907fffffffff0000000000000000000000000000000000000000000000000000000082168203613a525750903880613cf5565b5080fd5b3d9150613dcb565b505050600190565b9290803b15613ebf57613e9f9160209173ffffffffffffffffffffffffffffffffffffffff94604051809581948293897f150b7a02000000000000000000000000000000000000000000000000000000009b8c86523360048701521660248501526044840152608060648401526084830190612677565b03916000968791165af190829082613dbb575050613d9557613d03612e2d565b50505050600190565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311613f585791608094939160ff602094604051948552168484015260408301526060820152600093849182805260015afa15613f4b57815173ffffffffffffffffffffffffffffffffffffffff811615613f45579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600390565b60058110156140ce5780613f755750565b60018103613fdb5760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152fd5b600281036140415760646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152fd5b60031461404a57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1561410457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152fd5b61416a612884565b805190811561417a576020012090565b505060cd5480156141885790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b6141b5612969565b80519081156141c5576020012090565b505060ce548015614188579056fea2646970667358221220d018b61284e76c02ae2cc8148df9a827f7b6d15a491e9a9107c720efbabe12b964736f6c63430008160033",
}

// ModelCollectionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelCollectionV1MetaData.ABI instead.
var ModelCollectionV1ABI = ModelCollectionV1MetaData.ABI

// ModelCollectionV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModelCollectionV1MetaData.Bin instead.
var ModelCollectionV1Bin = ModelCollectionV1MetaData.Bin

// DeployModelCollectionV1 deploys a new Ethereum contract, binding an instance of ModelCollectionV1 to it.
func DeployModelCollectionV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModelCollectionV1, error) {
	parsed, err := ModelCollectionV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModelCollectionV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModelCollectionV1{ModelCollectionV1Caller: ModelCollectionV1Caller{contract: contract}, ModelCollectionV1Transactor: ModelCollectionV1Transactor{contract: contract}, ModelCollectionV1Filterer: ModelCollectionV1Filterer{contract: contract}}, nil
}

// ModelCollectionV1 is an auto generated Go binding around an Ethereum contract.
type ModelCollectionV1 struct {
	ModelCollectionV1Caller     // Read-only binding to the contract
	ModelCollectionV1Transactor // Write-only binding to the contract
	ModelCollectionV1Filterer   // Log filterer for contract events
}

// ModelCollectionV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCollectionV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelCollectionV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelCollectionV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelCollectionV1Session struct {
	Contract     *ModelCollectionV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ModelCollectionV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCollectionV1CallerSession struct {
	Contract *ModelCollectionV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ModelCollectionV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelCollectionV1TransactorSession struct {
	Contract     *ModelCollectionV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ModelCollectionV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ModelCollectionV1Raw struct {
	Contract *ModelCollectionV1 // Generic contract binding to access the raw methods on
}

// ModelCollectionV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCollectionV1CallerRaw struct {
	Contract *ModelCollectionV1Caller // Generic read-only contract binding to access the raw methods on
}

// ModelCollectionV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelCollectionV1TransactorRaw struct {
	Contract *ModelCollectionV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewModelCollectionV1 creates a new instance of ModelCollectionV1, bound to a specific deployed contract.
func NewModelCollectionV1(address common.Address, backend bind.ContractBackend) (*ModelCollectionV1, error) {
	contract, err := bindModelCollectionV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1{ModelCollectionV1Caller: ModelCollectionV1Caller{contract: contract}, ModelCollectionV1Transactor: ModelCollectionV1Transactor{contract: contract}, ModelCollectionV1Filterer: ModelCollectionV1Filterer{contract: contract}}, nil
}

// NewModelCollectionV1Caller creates a new read-only instance of ModelCollectionV1, bound to a specific deployed contract.
func NewModelCollectionV1Caller(address common.Address, caller bind.ContractCaller) (*ModelCollectionV1Caller, error) {
	contract, err := bindModelCollectionV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1Caller{contract: contract}, nil
}

// NewModelCollectionV1Transactor creates a new write-only instance of ModelCollectionV1, bound to a specific deployed contract.
func NewModelCollectionV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ModelCollectionV1Transactor, error) {
	contract, err := bindModelCollectionV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1Transactor{contract: contract}, nil
}

// NewModelCollectionV1Filterer creates a new log filterer instance of ModelCollectionV1, bound to a specific deployed contract.
func NewModelCollectionV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ModelCollectionV1Filterer, error) {
	contract, err := bindModelCollectionV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1Filterer{contract: contract}, nil
}

// bindModelCollectionV1 binds a generic wrapper to an already deployed contract.
func bindModelCollectionV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelCollectionV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollectionV1 *ModelCollectionV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollectionV1.Contract.ModelCollectionV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollectionV1 *ModelCollectionV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.ModelCollectionV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollectionV1 *ModelCollectionV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.ModelCollectionV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollectionV1 *ModelCollectionV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollectionV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollectionV1 *ModelCollectionV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollectionV1 *ModelCollectionV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollectionV1.Contract.BalanceOf(&_ModelCollectionV1.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollectionV1.Contract.BalanceOf(&_ModelCollectionV1.CallOpts, owner)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "eip712Domain")

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
func (_ModelCollectionV1 *ModelCollectionV1Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollectionV1.Contract.Eip712Domain(&_ModelCollectionV1.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollectionV1.Contract.Eip712Domain(&_ModelCollectionV1.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.GetApproved(&_ModelCollectionV1.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.GetApproved(&_ModelCollectionV1.CallOpts, tokenId)
}

// GetHashToSign is a free data retrieval call binding the contract method 0xeb21af39.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, address _manager) view returns(bytes32)
func (_ModelCollectionV1 *ModelCollectionV1Caller) GetHashToSign(opts *bind.CallOpts, _to common.Address, _uri string, _model common.Address, _manager common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "getHashToSign", _to, _uri, _model, _manager)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign is a free data retrieval call binding the contract method 0xeb21af39.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, address _manager) view returns(bytes32)
func (_ModelCollectionV1 *ModelCollectionV1Session) GetHashToSign(_to common.Address, _uri string, _model common.Address, _manager common.Address) ([32]byte, error) {
	return _ModelCollectionV1.Contract.GetHashToSign(&_ModelCollectionV1.CallOpts, _to, _uri, _model, _manager)
}

// GetHashToSign is a free data retrieval call binding the contract method 0xeb21af39.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, address _manager) view returns(bytes32)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) GetHashToSign(_to common.Address, _uri string, _model common.Address, _manager common.Address) ([32]byte, error) {
	return _ModelCollectionV1.Contract.GetHashToSign(&_ModelCollectionV1.CallOpts, _to, _uri, _model, _manager)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollectionV1.Contract.IsApprovedForAll(&_ModelCollectionV1.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollectionV1.Contract.IsApprovedForAll(&_ModelCollectionV1.CallOpts, owner, operator)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Caller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Session) IsManager(arg0 common.Address) (bool, error) {
	return _ModelCollectionV1.Contract.IsManager(&_ModelCollectionV1.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _ModelCollectionV1.Contract.IsManager(&_ModelCollectionV1.CallOpts, arg0)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) MintPrice() (*big.Int, error) {
	return _ModelCollectionV1.Contract.MintPrice(&_ModelCollectionV1.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) MintPrice() (*big.Int, error) {
	return _ModelCollectionV1.Contract.MintPrice(&_ModelCollectionV1.CallOpts)
}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Caller) ModelAddressOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "modelAddressOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Session) ModelAddressOf(_tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.ModelAddressOf(&_ModelCollectionV1.CallOpts, _tokenId)
}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) ModelAddressOf(_tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.ModelAddressOf(&_ModelCollectionV1.CallOpts, _tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Session) Name() (string, error) {
	return _ModelCollectionV1.Contract.Name(&_ModelCollectionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Name() (string, error) {
	return _ModelCollectionV1.Contract.Name(&_ModelCollectionV1.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) NextModelId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "nextModelId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) NextModelId() (*big.Int, error) {
	return _ModelCollectionV1.Contract.NextModelId(&_ModelCollectionV1.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) NextModelId() (*big.Int, error) {
	return _ModelCollectionV1.Contract.NextModelId(&_ModelCollectionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Session) Owner() (common.Address, error) {
	return _ModelCollectionV1.Contract.Owner(&_ModelCollectionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Owner() (common.Address, error) {
	return _ModelCollectionV1.Contract.Owner(&_ModelCollectionV1.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.OwnerOf(&_ModelCollectionV1.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollectionV1.Contract.OwnerOf(&_ModelCollectionV1.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Session) Paused() (bool, error) {
	return _ModelCollectionV1.Contract.Paused(&_ModelCollectionV1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Paused() (bool, error) {
	return _ModelCollectionV1.Contract.Paused(&_ModelCollectionV1.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ModelCollectionV1.Contract.RoyaltyInfo(&_ModelCollectionV1.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ModelCollectionV1.Contract.RoyaltyInfo(&_ModelCollectionV1.CallOpts, _tokenId, _salePrice)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollectionV1 *ModelCollectionV1Caller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollectionV1 *ModelCollectionV1Session) RoyaltyPortion() (uint16, error) {
	return _ModelCollectionV1.Contract.RoyaltyPortion(&_ModelCollectionV1.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) RoyaltyPortion() (uint16, error) {
	return _ModelCollectionV1.Contract.RoyaltyPortion(&_ModelCollectionV1.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Caller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1Session) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollectionV1.Contract.RoyaltyReceiver(&_ModelCollectionV1.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollectionV1.Contract.RoyaltyReceiver(&_ModelCollectionV1.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ModelCollectionV1.Contract.SupportsInterface(&_ModelCollectionV1.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ModelCollectionV1.Contract.SupportsInterface(&_ModelCollectionV1.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Session) Symbol() (string, error) {
	return _ModelCollectionV1.Contract.Symbol(&_ModelCollectionV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Symbol() (string, error) {
	return _ModelCollectionV1.Contract.Symbol(&_ModelCollectionV1.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollectionV1.Contract.TokenByIndex(&_ModelCollectionV1.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollectionV1.Contract.TokenByIndex(&_ModelCollectionV1.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollectionV1.Contract.TokenOfOwnerByIndex(&_ModelCollectionV1.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollectionV1.Contract.TokenOfOwnerByIndex(&_ModelCollectionV1.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ModelCollectionV1.Contract.TokenURI(&_ModelCollectionV1.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ModelCollectionV1.Contract.TokenURI(&_ModelCollectionV1.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) TotalSupply() (*big.Int, error) {
	return _ModelCollectionV1.Contract.TotalSupply(&_ModelCollectionV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) TotalSupply() (*big.Int, error) {
	return _ModelCollectionV1.Contract.TotalSupply(&_ModelCollectionV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollectionV1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollectionV1 *ModelCollectionV1Session) Version() (string, error) {
	return _ModelCollectionV1.Contract.Version(&_ModelCollectionV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollectionV1 *ModelCollectionV1CallerSession) Version() (string, error) {
	return _ModelCollectionV1.Contract.Version(&_ModelCollectionV1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Approve(&_ModelCollectionV1.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Approve(&_ModelCollectionV1.TransactOpts, to, tokenId)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) AuthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "authorizeManager", _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.AuthorizeManager(&_ModelCollectionV1.TransactOpts, _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.AuthorizeManager(&_ModelCollectionV1.TransactOpts, _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) DeauthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "deauthorizeManager", _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.DeauthorizeManager(&_ModelCollectionV1.TransactOpts, _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.DeauthorizeManager(&_ModelCollectionV1.TransactOpts, _account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "initialize", _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Initialize(&_ModelCollectionV1.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Initialize(&_ModelCollectionV1.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "mint", _to, _uri, _model)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) Mint(_to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Mint(&_ModelCollectionV1.TransactOpts, _to, _uri, _model)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Mint(_to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Mint(&_ModelCollectionV1.TransactOpts, _to, _uri, _model)
}

// MintBySignature is a paid mutator transaction binding the contract method 0xd9759dd1.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Transactor) MintBySignature(opts *bind.TransactOpts, _to common.Address, _uri string, _model common.Address, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "mintBySignature", _to, _uri, _model, _manager, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0xd9759dd1.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1Session) MintBySignature(_to common.Address, _uri string, _model common.Address, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.MintBySignature(&_ModelCollectionV1.TransactOpts, _to, _uri, _model, _manager, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0xd9759dd1.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) MintBySignature(_to common.Address, _uri string, _model common.Address, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.MintBySignature(&_ModelCollectionV1.TransactOpts, _to, _uri, _model, _manager, v, r, s)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Pause() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Pause(&_ModelCollectionV1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Pause() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Pause(&_ModelCollectionV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.RenounceOwnership(&_ModelCollectionV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.RenounceOwnership(&_ModelCollectionV1.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SafeTransferFrom(&_ModelCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SafeTransferFrom(&_ModelCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SafeTransferFrom0(&_ModelCollectionV1.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SafeTransferFrom0(&_ModelCollectionV1.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SetApprovalForAll(&_ModelCollectionV1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.SetApprovalForAll(&_ModelCollectionV1.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.TransferFrom(&_ModelCollectionV1.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.TransferFrom(&_ModelCollectionV1.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.TransferOwnership(&_ModelCollectionV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.TransferOwnership(&_ModelCollectionV1.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Unpause() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Unpause(&_ModelCollectionV1.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Unpause() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Unpause(&_ModelCollectionV1.TransactOpts)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) UpdateMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "updateMintPrice", _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateMintPrice(&_ModelCollectionV1.TransactOpts, _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateMintPrice(&_ModelCollectionV1.TransactOpts, _mintPrice)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, _royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "updateRoyaltyPortion", _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateRoyaltyPortion(&_ModelCollectionV1.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateRoyaltyPortion(&_ModelCollectionV1.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, _royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "updateRoyaltyReceiver", _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateRoyaltyReceiver(&_ModelCollectionV1.TransactOpts, _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateRoyaltyReceiver(&_ModelCollectionV1.TransactOpts, _royaltyReceiver)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) UpdateTokenModel(opts *bind.TransactOpts, _tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "updateTokenModel", _tokenId, _model)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) UpdateTokenModel(_tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateTokenModel(&_ModelCollectionV1.TransactOpts, _tokenId, _model)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) UpdateTokenModel(_tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateTokenModel(&_ModelCollectionV1.TransactOpts, _tokenId, _model)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) UpdateTokenURI(opts *bind.TransactOpts, _tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "updateTokenURI", _tokenId, _uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) UpdateTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateTokenURI(&_ModelCollectionV1.TransactOpts, _tokenId, _uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) UpdateTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.UpdateTokenURI(&_ModelCollectionV1.TransactOpts, _tokenId, _uri)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.Transact(opts, "withdraw", _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Withdraw(&_ModelCollectionV1.TransactOpts, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Withdraw(&_ModelCollectionV1.TransactOpts, _to, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollectionV1 *ModelCollectionV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollectionV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollectionV1 *ModelCollectionV1Session) Receive() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Receive(&_ModelCollectionV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollectionV1 *ModelCollectionV1TransactorSession) Receive() (*types.Transaction, error) {
	return _ModelCollectionV1.Contract.Receive(&_ModelCollectionV1.TransactOpts)
}

// ModelCollectionV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ModelCollectionV1 contract.
type ModelCollectionV1ApprovalIterator struct {
	Event *ModelCollectionV1Approval // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1Approval)
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
		it.Event = new(ModelCollectionV1Approval)
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
func (it *ModelCollectionV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1Approval represents a Approval event raised by the ModelCollectionV1 contract.
type ModelCollectionV1Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ModelCollectionV1ApprovalIterator, error) {

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

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1ApprovalIterator{contract: _ModelCollectionV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1Approval)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseApproval(log types.Log) (*ModelCollectionV1Approval, error) {
	event := new(ModelCollectionV1Approval)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ModelCollectionV1 contract.
type ModelCollectionV1ApprovalForAllIterator struct {
	Event *ModelCollectionV1ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1ApprovalForAll)
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
		it.Event = new(ModelCollectionV1ApprovalForAll)
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
func (it *ModelCollectionV1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1ApprovalForAll represents a ApprovalForAll event raised by the ModelCollectionV1 contract.
type ModelCollectionV1ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ModelCollectionV1ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1ApprovalForAllIterator{contract: _ModelCollectionV1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1ApprovalForAll)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseApprovalForAll(log types.Log) (*ModelCollectionV1ApprovalForAll, error) {
	event := new(ModelCollectionV1ApprovalForAll)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1BatchMetadataUpdateIterator struct {
	Event *ModelCollectionV1BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1BatchMetadataUpdate)
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
		it.Event = new(ModelCollectionV1BatchMetadataUpdate)
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
func (it *ModelCollectionV1BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionV1BatchMetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1BatchMetadataUpdateIterator{contract: _ModelCollectionV1.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1BatchMetadataUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseBatchMetadataUpdate(log types.Log) (*ModelCollectionV1BatchMetadataUpdate, error) {
	event := new(ModelCollectionV1BatchMetadataUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ModelCollectionV1 contract.
type ModelCollectionV1EIP712DomainChangedIterator struct {
	Event *ModelCollectionV1EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1EIP712DomainChanged)
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
		it.Event = new(ModelCollectionV1EIP712DomainChanged)
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
func (it *ModelCollectionV1EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1EIP712DomainChanged represents a EIP712DomainChanged event raised by the ModelCollectionV1 contract.
type ModelCollectionV1EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ModelCollectionV1EIP712DomainChangedIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1EIP712DomainChangedIterator{contract: _ModelCollectionV1.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1EIP712DomainChanged)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseEIP712DomainChanged(log types.Log) (*ModelCollectionV1EIP712DomainChanged, error) {
	event := new(ModelCollectionV1EIP712DomainChanged)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ModelCollectionV1 contract.
type ModelCollectionV1InitializedIterator struct {
	Event *ModelCollectionV1Initialized // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1Initialized)
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
		it.Event = new(ModelCollectionV1Initialized)
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
func (it *ModelCollectionV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1Initialized represents a Initialized event raised by the ModelCollectionV1 contract.
type ModelCollectionV1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*ModelCollectionV1InitializedIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1InitializedIterator{contract: _ModelCollectionV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1Initialized) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1Initialized)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseInitialized(log types.Log) (*ModelCollectionV1Initialized, error) {
	event := new(ModelCollectionV1Initialized)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1ManagerAuthorizationIterator is returned from FilterManagerAuthorization and is used to iterate over the raw logs and unpacked data for ManagerAuthorization events raised by the ModelCollectionV1 contract.
type ModelCollectionV1ManagerAuthorizationIterator struct {
	Event *ModelCollectionV1ManagerAuthorization // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1ManagerAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1ManagerAuthorization)
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
		it.Event = new(ModelCollectionV1ManagerAuthorization)
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
func (it *ModelCollectionV1ManagerAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1ManagerAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1ManagerAuthorization represents a ManagerAuthorization event raised by the ModelCollectionV1 contract.
type ModelCollectionV1ManagerAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAuthorization is a free log retrieval operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterManagerAuthorization(opts *bind.FilterOpts, account []common.Address) (*ModelCollectionV1ManagerAuthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1ManagerAuthorizationIterator{contract: _ModelCollectionV1.contract, event: "ManagerAuthorization", logs: logs, sub: sub}, nil
}

// WatchManagerAuthorization is a free log subscription operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchManagerAuthorization(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1ManagerAuthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1ManagerAuthorization)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseManagerAuthorization(log types.Log) (*ModelCollectionV1ManagerAuthorization, error) {
	event := new(ModelCollectionV1ManagerAuthorization)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1ManagerDeauthorizationIterator is returned from FilterManagerDeauthorization and is used to iterate over the raw logs and unpacked data for ManagerDeauthorization events raised by the ModelCollectionV1 contract.
type ModelCollectionV1ManagerDeauthorizationIterator struct {
	Event *ModelCollectionV1ManagerDeauthorization // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1ManagerDeauthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1ManagerDeauthorization)
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
		it.Event = new(ModelCollectionV1ManagerDeauthorization)
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
func (it *ModelCollectionV1ManagerDeauthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1ManagerDeauthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1ManagerDeauthorization represents a ManagerDeauthorization event raised by the ModelCollectionV1 contract.
type ModelCollectionV1ManagerDeauthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerDeauthorization is a free log retrieval operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterManagerDeauthorization(opts *bind.FilterOpts, account []common.Address) (*ModelCollectionV1ManagerDeauthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1ManagerDeauthorizationIterator{contract: _ModelCollectionV1.contract, event: "ManagerDeauthorization", logs: logs, sub: sub}, nil
}

// WatchManagerDeauthorization is a free log subscription operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchManagerDeauthorization(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1ManagerDeauthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1ManagerDeauthorization)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseManagerDeauthorization(log types.Log) (*ModelCollectionV1ManagerDeauthorization, error) {
	event := new(ModelCollectionV1ManagerDeauthorization)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1MetadataUpdateIterator struct {
	Event *ModelCollectionV1MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1MetadataUpdate)
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
		it.Event = new(ModelCollectionV1MetadataUpdate)
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
func (it *ModelCollectionV1MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1MetadataUpdate represents a MetadataUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionV1MetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1MetadataUpdateIterator{contract: _ModelCollectionV1.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1MetadataUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseMetadataUpdate(log types.Log) (*ModelCollectionV1MetadataUpdate, error) {
	event := new(ModelCollectionV1MetadataUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1MintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1MintPriceUpdateIterator struct {
	Event *ModelCollectionV1MintPriceUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1MintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1MintPriceUpdate)
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
		it.Event = new(ModelCollectionV1MintPriceUpdate)
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
func (it *ModelCollectionV1MintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1MintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1MintPriceUpdate represents a MintPriceUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1MintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*ModelCollectionV1MintPriceUpdateIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1MintPriceUpdateIterator{contract: _ModelCollectionV1.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1MintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1MintPriceUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseMintPriceUpdate(log types.Log) (*ModelCollectionV1MintPriceUpdate, error) {
	event := new(ModelCollectionV1MintPriceUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1NewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the ModelCollectionV1 contract.
type ModelCollectionV1NewTokenIterator struct {
	Event *ModelCollectionV1NewToken // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1NewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1NewToken)
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
		it.Event = new(ModelCollectionV1NewToken)
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
func (it *ModelCollectionV1NewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1NewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1NewToken represents a NewToken event raised by the ModelCollectionV1 contract.
type ModelCollectionV1NewToken struct {
	TokenId *big.Int
	Uri     string
	Model   common.Address
	Minter  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewToken is a free log retrieval operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterNewToken(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address) (*ModelCollectionV1NewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1NewTokenIterator{contract: _ModelCollectionV1.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1NewToken, tokenId []*big.Int, minter []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1NewToken)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "NewToken", log); err != nil {
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

// ParseNewToken is a log parse operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseNewToken(log types.Log) (*ModelCollectionV1NewToken, error) {
	event := new(ModelCollectionV1NewToken)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ModelCollectionV1 contract.
type ModelCollectionV1OwnershipTransferredIterator struct {
	Event *ModelCollectionV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1OwnershipTransferred)
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
		it.Event = new(ModelCollectionV1OwnershipTransferred)
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
func (it *ModelCollectionV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1OwnershipTransferred represents a OwnershipTransferred event raised by the ModelCollectionV1 contract.
type ModelCollectionV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModelCollectionV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1OwnershipTransferredIterator{contract: _ModelCollectionV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1OwnershipTransferred)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseOwnershipTransferred(log types.Log) (*ModelCollectionV1OwnershipTransferred, error) {
	event := new(ModelCollectionV1OwnershipTransferred)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ModelCollectionV1 contract.
type ModelCollectionV1PausedIterator struct {
	Event *ModelCollectionV1Paused // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1Paused)
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
		it.Event = new(ModelCollectionV1Paused)
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
func (it *ModelCollectionV1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1Paused represents a Paused event raised by the ModelCollectionV1 contract.
type ModelCollectionV1Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterPaused(opts *bind.FilterOpts) (*ModelCollectionV1PausedIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1PausedIterator{contract: _ModelCollectionV1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1Paused) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1Paused)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParsePaused(log types.Log) (*ModelCollectionV1Paused, error) {
	event := new(ModelCollectionV1Paused)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1RoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1RoyaltyPortionUpdateIterator struct {
	Event *ModelCollectionV1RoyaltyPortionUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1RoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1RoyaltyPortionUpdate)
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
		it.Event = new(ModelCollectionV1RoyaltyPortionUpdate)
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
func (it *ModelCollectionV1RoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1RoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1RoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1RoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*ModelCollectionV1RoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1RoyaltyPortionUpdateIterator{contract: _ModelCollectionV1.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1RoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1RoyaltyPortionUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseRoyaltyPortionUpdate(log types.Log) (*ModelCollectionV1RoyaltyPortionUpdate, error) {
	event := new(ModelCollectionV1RoyaltyPortionUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1RoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1RoyaltyReceiverUpdateIterator struct {
	Event *ModelCollectionV1RoyaltyReceiverUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1RoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1RoyaltyReceiverUpdate)
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
		it.Event = new(ModelCollectionV1RoyaltyReceiverUpdate)
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
func (it *ModelCollectionV1RoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1RoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1RoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1RoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*ModelCollectionV1RoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1RoyaltyReceiverUpdateIterator{contract: _ModelCollectionV1.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1RoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1RoyaltyReceiverUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseRoyaltyReceiverUpdate(log types.Log) (*ModelCollectionV1RoyaltyReceiverUpdate, error) {
	event := new(ModelCollectionV1RoyaltyReceiverUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1TokenModelUpdateIterator is returned from FilterTokenModelUpdate and is used to iterate over the raw logs and unpacked data for TokenModelUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1TokenModelUpdateIterator struct {
	Event *ModelCollectionV1TokenModelUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1TokenModelUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1TokenModelUpdate)
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
		it.Event = new(ModelCollectionV1TokenModelUpdate)
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
func (it *ModelCollectionV1TokenModelUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1TokenModelUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1TokenModelUpdate represents a TokenModelUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1TokenModelUpdate struct {
	TokenId *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenModelUpdate is a free log retrieval operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterTokenModelUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*ModelCollectionV1TokenModelUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "TokenModelUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1TokenModelUpdateIterator{contract: _ModelCollectionV1.contract, event: "TokenModelUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenModelUpdate is a free log subscription operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchTokenModelUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1TokenModelUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "TokenModelUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1TokenModelUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "TokenModelUpdate", log); err != nil {
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

// ParseTokenModelUpdate is a log parse operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseTokenModelUpdate(log types.Log) (*ModelCollectionV1TokenModelUpdate, error) {
	event := new(ModelCollectionV1TokenModelUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "TokenModelUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1TokenURIUpdateIterator is returned from FilterTokenURIUpdate and is used to iterate over the raw logs and unpacked data for TokenURIUpdate events raised by the ModelCollectionV1 contract.
type ModelCollectionV1TokenURIUpdateIterator struct {
	Event *ModelCollectionV1TokenURIUpdate // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1TokenURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1TokenURIUpdate)
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
		it.Event = new(ModelCollectionV1TokenURIUpdate)
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
func (it *ModelCollectionV1TokenURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1TokenURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1TokenURIUpdate represents a TokenURIUpdate event raised by the ModelCollectionV1 contract.
type ModelCollectionV1TokenURIUpdate struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenURIUpdate is a free log retrieval operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterTokenURIUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*ModelCollectionV1TokenURIUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "TokenURIUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1TokenURIUpdateIterator{contract: _ModelCollectionV1.contract, event: "TokenURIUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenURIUpdate is a free log subscription operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchTokenURIUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1TokenURIUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "TokenURIUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1TokenURIUpdate)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "TokenURIUpdate", log); err != nil {
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

// ParseTokenURIUpdate is a log parse operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseTokenURIUpdate(log types.Log) (*ModelCollectionV1TokenURIUpdate, error) {
	event := new(ModelCollectionV1TokenURIUpdate)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "TokenURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ModelCollectionV1 contract.
type ModelCollectionV1TransferIterator struct {
	Event *ModelCollectionV1Transfer // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1Transfer)
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
		it.Event = new(ModelCollectionV1Transfer)
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
func (it *ModelCollectionV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1Transfer represents a Transfer event raised by the ModelCollectionV1 contract.
type ModelCollectionV1Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ModelCollectionV1TransferIterator, error) {

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

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1TransferIterator{contract: _ModelCollectionV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1Transfer)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseTransfer(log types.Log) (*ModelCollectionV1Transfer, error) {
	event := new(ModelCollectionV1Transfer)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionV1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ModelCollectionV1 contract.
type ModelCollectionV1UnpausedIterator struct {
	Event *ModelCollectionV1Unpaused // Event containing the contract specifics and raw log

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
func (it *ModelCollectionV1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionV1Unpaused)
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
		it.Event = new(ModelCollectionV1Unpaused)
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
func (it *ModelCollectionV1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionV1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionV1Unpaused represents a Unpaused event raised by the ModelCollectionV1 contract.
type ModelCollectionV1Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) FilterUnpaused(opts *bind.FilterOpts) (*ModelCollectionV1UnpausedIterator, error) {

	logs, sub, err := _ModelCollectionV1.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionV1UnpausedIterator{contract: _ModelCollectionV1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollectionV1 *ModelCollectionV1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionV1Unpaused) (event.Subscription, error) {

	logs, sub, err := _ModelCollectionV1.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionV1Unpaused)
				if err := _ModelCollectionV1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ModelCollectionV1 *ModelCollectionV1Filterer) ParseUnpaused(log types.Log) (*ModelCollectionV1Unpaused, error) {
	event := new(ModelCollectionV1Unpaused)
	if err := _ModelCollectionV1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
