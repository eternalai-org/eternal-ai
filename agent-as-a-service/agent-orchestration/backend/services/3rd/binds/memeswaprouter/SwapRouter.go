// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package memeswaprouter

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// SwapRouterMetaData contains all meta data concerning the SwapRouter contract.
var SwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETHArg\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETHWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001c576000196067556124a990816100228239f35b600080fdfe60806040526004361015610023575b361561001957600080fd5b610021611bc0565b005b60003560e01c806312210e8a146101835780632135ac891461017e578063414bf389146101795780634659a49414610174578063485cc9551461016f5780635b769f3c1461016a578063715018a6146101655780638da5cb5b14610160578063a4a78f0c1461015b578063ac9650d814610156578063ad5c464814610151578063c04b8d591461014c578063c2e3140a14610147578063c45a015514610142578063db3e21981461013d578063df2ab5bb14610138578063e0e189a014610133578063e16d9ce51461012e578063f28c049814610129578063f2fde38b14610124578063f3995c671461011f5763fa461e330361000e57610f86565b610f6e565b610edd565b610e2b565b610d46565b610c57565b610bbf565b610b14565b610aeb565b610a72565b6109e0565b610898565b610792565b61066d565b610644565b6105e3565b61059c565b6104a3565b61048b565b610349565b6101c6565b610198565b600091031261019357565b600080fd5b600036600319011261019357476101ab57005b610021473361240d565b6001600160a01b0381160361019357565b6080366003190112610193576024356101de816101b5565b606435906044356101ee836101b5565b8015158061032c575b610200906111d1565b60665461022390610217906001600160a01b031681565b6001600160a01b031690565b6040516370a0823160e01b8152306004820152909390602081602481885afa9081156102f9576000916102fe575b50610260600435821015611c13565b8061026757005b843b1561019357604051632e1a7d4d60e01b815260048101829052946000908690602490829084905af19485156102f9576102b36102ca946102bb92610021986102e0575b5083611dfd565b612710900490565b9182806102d0575b5050611df0565b9061240d565b6102d99161240d565b38826102c3565b806102ed6102f3926108d7565b80610188565b386102ac565b611362565b61031f915060203d8111610325575b6103178183610940565b810190611c04565b38610251565b503d61030d565b5060648111156101f7565b61010090600319011261019357600490565b6101003660031901126101935761044e61036236610337565b6103726080820135421115611239565b61043e60c06104336060840135610388816101b5565b60e0850135610396816101b5565b8535916103a2836101b5565b6104146103b16040890161127b565b9361040660208a01356103c3816101b5565b604051968793602085019192602b936bffffffffffffffffffffffff19809360601b16845262ffffff60e81b9060e81b16601484015260601b1660178201520190565b03601f198101855284610940565b60405192610421846108ef565b835233602084015260a08701356114a3565b92013582101561128b565b6040519081529081906020820190565b0390f35b60c09060031901126101935760043561046a816101b5565b90602435906044359060643560ff8116810361019357906084359060a43590565b61002161049736610452565b94939093929192611ec2565b34610193576040366003190112610193576004356104c0816101b5565b6105106024356104cf816101b5565b600054926104f460ff8560081c16158095819661058e575b811561056e575b5061116e565b83610507600160ff196000541617600055565b61055557611a76565b61051657005b61052661ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61056961010061ff00196000541617600055565b611a76565b303b15915081610580575b50386104ee565b6001915060ff161438610579565b600160ff82161091506104e7565b34610193576020366003190112610193576004356105b9816101b5565b6105c16110cd565b606680546001600160a01b0319166001600160a01b0392909216919091179055005b3461019357600080600319360112610641576105fd6110cd565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b34610193576000366003190112610193576033546040516001600160a01b039091168152602090f35b61067636610452565b604051636eb1769f60e11b81523360048201523060248201529094919391906020816044816001600160a01b038b165afa9081156102f9576000916106ca575b50600019116106c157005b61002195611ec2565b6106e2915060203d8111610325576103178183610940565b386106b6565b60005b8381106106fb5750506000910152565b81810151838201526020016106eb565b90602091610724815180928185528580860191016106e8565b601f01601f1916010190565b602080820190808352835180925260408301928160408460051b8301019501936000915b8483106107645750505050505090565b9091929394958480610782600193603f198682030187528a5161070b565b9801930193019194939290610754565b602036600319011261019357600480356001600160401b03918282116101935736602383011215610193578181013592831161019357602490818301928236918660051b010111610193576107e684611926565b9360005b8181106107ff576040518061044e8882610730565b60008061080d838589611995565b6040939161081f8551809381936119b5565b0390305af49061082d6119c3565b9182901561085c57505090610857916108468289611a62565b526108518188611a62565b50611970565b6107ea565b8683879260448251106101935782610894938561087f94015183010191016119f3565b925162461bcd60e51b81529283928301611a51565b0390fd5b34610193576000366003190112610193576066546040516001600160a01b039091168152602090f35b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116108ea57604052565b6108c1565b604081019081106001600160401b038211176108ea57604052565b60a081019081106001600160401b038211176108ea57604052565b606081019081106001600160401b038211176108ea57604052565b90601f801991011681019081106001600160401b038211176108ea57604052565b6040519061096e826108ef565b565b6001600160401b0381116108ea57601f01601f191660200190565b92919261099782610970565b916109a56040519384610940565b829481845281830111610193578281602093846000960137010152565b9080601f83011215610193578160206109dd9335910161098b565b90565b600319602036820112610193576004356001600160401b03918282116101935760a090823603011261019357604051610a188161090a565b816004013592831161019357608461043e92610a3d61044e95600436918401016109c2565b83526024810135610a4d816101b5565b602084015260448101356040840152606481013560608401520135608082015261158e565b610a7b36610452565b604051636eb1769f60e11b81523360048201523060248201529094919391906020816044816001600160a01b038b165afa80156102f9578291600091610acd575b5010610ac457005b61002195611e4c565b610ae5915060203d8111610325576103178183610940565b38610abc565b34610193576000366003190112610193576065546040516001600160a01b039091168152602090f35b6101003660031901126101935761044e610b2d36610337565b610b3d6080820135421115611239565b610bb460c0610ba96060840135610b53816101b5565b60e0850135610b61816101b5565b602086013591610b70836101b5565b610b8e610b7f6040890161127b565b9361040689356103c3816101b5565b610b96610961565b92835233602084015260a087013561175a565b92013582111561189c565b60001960675561043e565b606036600319011261019357600435610bd7816101b5565b604435610be3816101b5565b6040516370a0823160e01b8152306004820152906020826024816001600160a01b0387165afa9182156102f957600092610c37575b50610c27602435831015611c53565b81610c2e57005b6100219261235b565b610c5091925060203d8111610325576103178183610940565b9038610c18565b60a036600319011261019357600435610c6f816101b5565b60443590610c7c826101b5565b60843591606435610c8c846101b5565b80151580610d3b575b610c9e906111d1565b6040516370a0823160e01b8152306004820152936020856024816001600160a01b0388165afa9485156102f957600095610d1b575b50610ce2602435861015611c53565b84610ce957005b84610cfd6102b3610d0b9461002198611dfd565b918280610d11575050611df0565b9161235b565b6102d9918761235b565b610d3491955060203d8111610325576103178183610940565b9338610cd3565b506064811115610c95565b604036600319011261019357602435610d5e816101b5565b606654610d7590610217906001600160a01b031681565b6040516370a0823160e01b81523060048201529091602082602481865afa9182156102f957600092610e0b575b50610db1600435831015611c13565b81610db857005b823b1561019357604051632e1a7d4d60e01b815260048101839052926000908490602490829084905af19283156102f95761002193610df8575b5061240d565b806102ed610e05926108d7565b38610df2565b610e2491925060203d8111610325576103178183610940565b9038610da2565b600319602036820112610193576004356001600160401b0381116101935760a08160040192823603011261019357610eb3610ea161044e93610e736044850135421115611239565b610e8b602485013591610e85836101b5565b806118dd565b929060405193610e9a856108ef565b369161098b565b82523360208301526064840135611638565b50610ec7608460675492013582111561189c565b6000196067556040519081529081906020820190565b3461019357602036600319011261019357600435610efa816101b5565b610f026110cd565b6001600160a01b03811615610f1a5761002190611125565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b610021610f7a36610452565b94939093929192611e4c565b34610193576060366003190112610193576044356024356004356001600160401b038084116101935736602385011215610193578360040135908111610193578301602401923684116101935761101f610ff96000956024878613948580156110c4575b610ff3906111d1565b016111d8565b936110048551611f73565b606554929591949192859087906001600160a01b0316611f33565b50156110ae57506001600160a01b03818116908316105b1561105c57506020929092015161105992906001600160a01b03165b3391611cac565b80f35b905061106c835160429051101590565b1561108e57509061108a9161108182516120fa565b82523390611638565b5080f35b6110526020611059946110a085606755565b01516001600160a01b031690565b92506001600160a01b0382811690821610611036565b50888813610fea565b6033546001600160a01b031633036110e157565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b1561117557565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b1561019357565b906020828203126101935781356001600160401b03928382116101935701604081830312610193576040519261120d846108ef565b8135908111610193576020926112249183016109c2565b83520135611231816101b5565b602082015290565b1561124057565b60405162461bcd60e51b8152602060048201526013602482015272151c985b9cd858dd1a5bdb881d1bdbc81bdb19606a1b6044820152606490fd5b3562ffffff811681036101935790565b1561129257565b60405162461bcd60e51b8152602060048201526013602482015272151bdbc81b1a5d1d1b19481c9958d95a5d9959606a1b6044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b6020815260406112fe8351826020850152606084019061070b565b6020909301516001600160a01b031691015290565b9190826040910312610193576020825192015190565b6001600160a01b039182168152911515602083015260408201929092529116606082015260a0608082018190526109dd9291019061070b565b6040513d6000823e3d90fd5b600160ff1b811461137f5760000390565b6112cd565b61021792916040916001600160a01b038083161561149b575b60006113c86113c26113af8851611f73565b909a9195808c16908716109a8b96611566565b93611557565b93828214611477576114196113ef6113fd6401000276a4995b8951928391602083016112e3565b03601f198101835282610940565b8751630251596160e31b81529889978896879560048701611329565b03925af19081156102f9576109dd926000918293611445575b501561143e575061136e565b905061136e565b909250611469915060403d8111611470575b6114618183610940565b810190611313565b9138611432565b503d611457565b6114196113ef6113fd73fffd8963efd1fc6a506488495d951d5263988d25996113e1565b30925061139d565b6114e593926040926001600160a01b039291610217916000918582161561154f575b6114eb6114d28951611f73565b89829d939892168a8916109c8d98611566565b95611557565b95811615831461153f575082821461151c576114196113ef6113fd6401000276a45b998951928391602083016112e3565b6114196113ef6113fd73fffd8963efd1fc6a506488495d951d5263988d2561150d565b6113ef6113fd61141992996113e1565b3091506114c5565b600160ff1b8110156101935790565b6065546001600160a01b039361158a939185169261158492906121a6565b90612217565b1690565b61159e6040820151421115611239565b6000335b8251906115ec60428351101591606086019384518460001461161c576115e76115cb3093612032565b936115d4610961565b9485526001600160a01b03166020850152565b611384565b80925260001461160957503061160283516120fa565b83526115a2565b6109dd915060809092015182101561128b565b60208801516115e7906115cb906001600160a01b031693612032565b919291906001600160a01b039081811615611753575b6040906116716102176116618851611f73565b9190968082169088161096611566565b83600061168561168088611557565b61136e565b9382821461172f576116c76113ef6116ab6401000276a49c8951928391602083016112e3565b8751630251596160e31b81529b8c978896879560048701611329565b03925af19081156102f957600094859261170c575b50156116f857906116ef61096e9261136e565b935b93146111d1565b929061170661096e9261136e565b936116f1565b909450611727915060403d8111611470576114618183610940565b9093386116dc565b6116c76113ef6116ab73fffd8963efd1fc6a506488495d951d5263988d259c6113e1565b503061164e565b919392919061179a906040906001600160a01b039081811615611895575b61021760006117878851611f73565b8287168288161098899590939092611566565b926117a761168089611557565b948b16159a8b83146118855750828214611862576117f66113ef6117da6401000276a45b9b8951928391602083016112e3565b8751630251596160e31b81529a8b978896879560048701611329565b03925af19081156102f957600093849261183f575b501561182f5761181a9061136e565b915b93611825575050565b61096e91146111d1565b916118399061136e565b9161181c565b90935061185a915060403d8111611470576114618183610940565b90923861180b565b6117f66113ef6117da73fffd8963efd1fc6a506488495d951d5263988d256117cb565b6113ef6117da6117f6929b6113e1565b5030611778565b156118a357565b60405162461bcd60e51b8152602060048201526012602482015271151bdbc81b5d58da081c995c5d595cdd195960721b6044820152606490fd5b903590601e198136030182121561019357018035906001600160401b0382116101935760200191813603831361019357565b6001600160401b0381116108ea5760051b60200190565b906119308261190f565b61193d6040519182610940565b828152809261194e601f199161190f565b019060005b82811061195f57505050565b806060602080938501015201611953565b600019811461137f5760010190565b634e487b7160e01b600052603260045260246000fd5b908210156119b0576119ac9160051b8101906118dd565b9091565b61197f565b908092918237016000815290565b3d156119ee573d906119d482610970565b916119e26040519384610940565b82523d6000602084013e565b606090565b602081830312610193578051906001600160401b038211610193570181601f82011215610193578051611a2581610970565b92611a336040519485610940565b81845260208284010111610193576109dd91602080850191016106e8565b9060206109dd92818152019061070b565b80518210156119b05760209160051b010190565b60005491611a9960ff8460081c16158094819561058e57811561056e575061116e565b82611aac600160ff196000541617600055565b611b47575b611acb60ff60005460081c16611ac681611b60565b611b60565b611ad433611125565b60018060a01b0390816bffffffffffffffffffffffff60a01b931683606554161760655516906066541617606655611b0857565b611b1861ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1565b611b5b61010061ff00196000541617600055565b611ab1565b15611b6757565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6066546001600160a01b03163303611bd457565b60405162461bcd60e51b815260206004820152600860248201526709cdee840ae8aa8960c31b6044820152606490fd5b90816020910312610193575190565b15611c1a57565b60405162461bcd60e51b8152602060048201526011602482015270092dce6eaccccd2c6d2cadce840ae8aa89607b1b6044820152606490fd5b15611c5a57565b60405162461bcd60e51b815260206004820152601260248201527124b739bab33334b1b4b2b73a103a37b5b2b760711b6044820152606490fd5b90816020910312610193575180151581036101935790565b6066549093929190611cc6906001600160a01b0316610217565b6001600160a01b03908582161480611de6575b15611dcb57505060665491925090611cfb90610217906001600160a01b031681565b803b156101935760008391600460405180968193630d0e30db60e41b83525af19182156102f957611d7d93602093611db8575b50606654611d4690610217906001600160a01b031681565b60405163a9059cbb60e01b81526001600160a01b039092166004830152602482019290925292839190829060009082906044820190565b03925af180156102f957611d8e5750565b611dae9060203d8111611db1575b611da68183610940565b810190611c94565b50565b503d611d9c565b806102ed611dc5926108d7565b38611d2e565b81163003611ddd575061096e9261235b565b61096e936122a4565b5083471015611cd9565b9190820391821161137f57565b600092918115918215611e14575b50501561019357565b908092945081029081049081831485171561137f5793611e3657143880611e0b565b634e487b7160e01b600052601260045260246000fd5b92946001600160a01b0390931693919291843b156101935760009460e493869260ff604051998a98899763d505accf60e01b89523360048a01523060248a01526044890152606488015216608486015260a485015260c48401525af180156102f957611eb55750565b806102ed61096e926108d7565b92946001600160a01b0390931693919291843b156101935760009461010493869260ff604051998a9889976323f2ebc360e21b89523360048a01523060248a015260448901526064880152600160848801521660a486015260c485015260e48401525af180156102f957611eb55750565b6001600160a01b0393611f4c93919261158492906121a6565b168033036101935790565b90601f820180921161137f57565b601701908160171161137f57565b90611f82601483511015611fee565b602082015160601c916017815110611fb2576037601782015191611faa602b82511015611fee565b015160601c91565b60405162461bcd60e51b8152602060048201526014602482015273746f55696e7432345f6f75744f66426f756e647360601b6044820152606490fd5b15611ff557565b60405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b6044820152606490fd5b612040602b825110156120ba565b60405190600b8083019101603683015b80831061206a575050602b8252601f01601f191660405290565b9091825181526020809101920190612050565b1561208457565b60405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b6044820152606490fd5b156120c157565b60405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606490fd5b8051601619918282019082821161137f5761211f8261211881611f57565b101561207d565b61212d601761211884611f65565b612142815161213b84611f65565b11156120ba565b8161215c5750505050604051600081526020810160405290565b601760405194601f8416801560051b9182828901019687010193010101905b8084106121935750508252601f01601f191660405290565b909283518152602080910193019061217b565b9162ffffff916000604080516121bb81610925565b8281526020810183905201526001600160a01b0390808216858316116121fc575b81604051956121ea87610925565b16855216602084015216604082015290565b936121dc565b9081602091031261019357516109dd816101b5565b8151602083015191926001600160a01b03928316918316828110156101935760209362ffffff6040606494015116956040519687958694630b4c774160e11b8652600486015260248501526044840152165afa9081156102f95760009161227c575090565b6109dd915060203d811161229d575b6122958183610940565b810190612202565b503d61228b565b9091600080949381946040519160208301946323b872dd60e01b865260018060a01b0380921660248501521660448301526064820152606481526122e78161090a565b51925af16122f36119c3565b8161232c575b501561230157565b60405162461bcd60e51b815260206004820152600360248201526229aa2360e91b6044820152606490fd5b8051801592508215612341575b5050386122f9565b6123549250602080918301019101611c94565b3880612339565b60405163a9059cbb60e01b602082019081526001600160a01b0390931660248201526044810193909352600092839290839061239a81606481016113ef565b51925af16123a66119c3565b816123de575b50156123b457565b60405162461bcd60e51b815260206004820152600260248201526114d560f21b6044820152606490fd5b80518015925082156123f3575b5050386123ac565b6124069250602080918301019101611c94565b38806123eb565b60405160208101908082106001600160401b038311176108ea576000938493848094938194604052525af16124406119c3565b501561244857565b60405162461bcd60e51b815260206004820152600360248201526253544560e81b6044820152606490fdfea2646970667358221220fedc74cc9a8b56d98385c190d3c76a16fe5e399e1979676f679dbb23e6fc875c64736f6c63430008130033",
}

// SwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapRouterMetaData.ABI instead.
var SwapRouterABI = SwapRouterMetaData.ABI

// SwapRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapRouterMetaData.Bin instead.
var SwapRouterBin = SwapRouterMetaData.Bin

// DeploySwapRouter deploys a new Ethereum contract, binding an instance of SwapRouter to it.
func DeploySwapRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapRouter, error) {
	parsed, err := SwapRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapRouter{SwapRouterCaller: SwapRouterCaller{contract: contract}, SwapRouterTransactor: SwapRouterTransactor{contract: contract}, SwapRouterFilterer: SwapRouterFilterer{contract: contract}}, nil
}

// SwapRouter is an auto generated Go binding around an Ethereum contract.
type SwapRouter struct {
	SwapRouterCaller     // Read-only binding to the contract
	SwapRouterTransactor // Write-only binding to the contract
	SwapRouterFilterer   // Log filterer for contract events
}

// SwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapRouterSession struct {
	Contract     *SwapRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapRouterCallerSession struct {
	Contract *SwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapRouterTransactorSession struct {
	Contract     *SwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRouterRaw struct {
	Contract *SwapRouter // Generic contract binding to access the raw methods on
}

// SwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapRouterCallerRaw struct {
	Contract *SwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapRouterTransactorRaw struct {
	Contract *SwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapRouter creates a new instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouter(address common.Address, backend bind.ContractBackend) (*SwapRouter, error) {
	contract, err := bindSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRouter{SwapRouterCaller: SwapRouterCaller{contract: contract}, SwapRouterTransactor: SwapRouterTransactor{contract: contract}, SwapRouterFilterer: SwapRouterFilterer{contract: contract}}, nil
}

// NewSwapRouterCaller creates a new read-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*SwapRouterCaller, error) {
	contract, err := bindSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterCaller{contract: contract}, nil
}

// NewSwapRouterTransactor creates a new write-only instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapRouterTransactor, error) {
	contract, err := bindSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRouterTransactor{contract: contract}, nil
}

// NewSwapRouterFilterer creates a new log filterer instance of SwapRouter, bound to a specific deployed contract.
func NewSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapRouterFilterer, error) {
	contract, err := bindSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapRouterFilterer{contract: contract}, nil
}

// bindSwapRouter binds a generic wrapper to an already deployed contract.
func bindSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.SwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.SwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRouter *SwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRouter *SwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRouter *SwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapRouter *SwapRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapRouter *SwapRouterSession) WETH() (common.Address, error) {
	return _SwapRouter.Contract.WETH(&_SwapRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_SwapRouter *SwapRouterCallerSession) WETH() (common.Address, error) {
	return _SwapRouter.Contract.WETH(&_SwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterSession) Factory() (common.Address, error) {
	return _SwapRouter.Contract.Factory(&_SwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapRouter *SwapRouterCallerSession) Factory() (common.Address, error) {
	return _SwapRouter.Contract.Factory(&_SwapRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRouter *SwapRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRouter *SwapRouterSession) Owner() (common.Address, error) {
	return _SwapRouter.Contract.Owner(&_SwapRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRouter *SwapRouterCallerSession) Owner() (common.Address, error) {
	return _SwapRouter.Contract.Owner(&_SwapRouter.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInput(&_SwapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInput(&_SwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SwapRouter *SwapRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactInputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutput(&_SwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutput(&_SwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutputSingle(&_SwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SwapRouter *SwapRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SwapRouter.Contract.ExactOutputSingle(&_SwapRouter.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_SwapRouter *SwapRouterTransactor) Initialize(opts *bind.TransactOpts, _factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "initialize", _factory, _WETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_SwapRouter *SwapRouterSession) Initialize(_factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.Initialize(&_SwapRouter.TransactOpts, _factory, _WETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_SwapRouter *SwapRouterTransactorSession) Initialize(_factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.Initialize(&_SwapRouter.TransactOpts, _factory, _WETH)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.Multicall(&_SwapRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SwapRouter *SwapRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.Multicall(&_SwapRouter.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterSession) RefundETH() (*types.Transaction, error) {
	return _SwapRouter.Contract.RefundETH(&_SwapRouter.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SwapRouter *SwapRouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _SwapRouter.Contract.RefundETH(&_SwapRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRouter *SwapRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRouter *SwapRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapRouter.Contract.RenounceOwnership(&_SwapRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRouter *SwapRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapRouter.Contract.RenounceOwnership(&_SwapRouter.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermit(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermit(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowed(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowed(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowedIfNecessary(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitAllowedIfNecessary(&_SwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitIfNecessary(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.SelfPermitIfNecessary(&_SwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_SwapRouter *SwapRouterTransactor) SetWETH(opts *bind.TransactOpts, WETHArg common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "setWETH", WETHArg)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_SwapRouter *SwapRouterSession) SetWETH(WETHArg common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SetWETH(&_SwapRouter.TransactOpts, WETHArg)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_SwapRouter *SwapRouterTransactorSession) SetWETH(WETHArg common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SetWETH(&_SwapRouter.TransactOpts, WETHArg)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepToken(&_SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepToken(&_SwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepTokenWithFee(&_SwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.SweepTokenWithFee(&_SwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRouter *SwapRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRouter *SwapRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.TransferOwnership(&_SwapRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRouter *SwapRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.TransferOwnership(&_SwapRouter.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.UniswapV3SwapCallback(&_SwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SwapRouter *SwapRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SwapRouter.Contract.UniswapV3SwapCallback(&_SwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0xe16d9ce5.
//
// Solidity: function unwrapWETH(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) UnwrapWETH(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "unwrapWETH", amountMinimum, recipient)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0xe16d9ce5.
//
// Solidity: function unwrapWETH(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterSession) UnwrapWETH(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH(&_SwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH is a paid mutator transaction binding the contract method 0xe16d9ce5.
//
// Solidity: function unwrapWETH(uint256 amountMinimum, address recipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) UnwrapWETH(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETH(&_SwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETHWithFee is a paid mutator transaction binding the contract method 0x2135ac89.
//
// Solidity: function unwrapWETHWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactor) UnwrapWETHWithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.contract.Transact(opts, "unwrapWETHWithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETHWithFee is a paid mutator transaction binding the contract method 0x2135ac89.
//
// Solidity: function unwrapWETHWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterSession) UnwrapWETHWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETHWithFee(&_SwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETHWithFee is a paid mutator transaction binding the contract method 0x2135ac89.
//
// Solidity: function unwrapWETHWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SwapRouter *SwapRouterTransactorSession) UnwrapWETHWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SwapRouter.Contract.UnwrapWETHWithFee(&_SwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterSession) Receive() (*types.Transaction, error) {
	return _SwapRouter.Contract.Receive(&_SwapRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapRouter *SwapRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapRouter.Contract.Receive(&_SwapRouter.TransactOpts)
}

// SwapRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SwapRouter contract.
type SwapRouterInitializedIterator struct {
	Event *SwapRouterInitialized // Event containing the contract specifics and raw log

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
func (it *SwapRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRouterInitialized)
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
		it.Event = new(SwapRouterInitialized)
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
func (it *SwapRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRouterInitialized represents a Initialized event raised by the SwapRouter contract.
type SwapRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SwapRouter *SwapRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*SwapRouterInitializedIterator, error) {

	logs, sub, err := _SwapRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SwapRouterInitializedIterator{contract: _SwapRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SwapRouter *SwapRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SwapRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _SwapRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRouterInitialized)
				if err := _SwapRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SwapRouter *SwapRouterFilterer) ParseInitialized(log types.Log) (*SwapRouterInitialized, error) {
	event := new(SwapRouterInitialized)
	if err := _SwapRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapRouter contract.
type SwapRouterOwnershipTransferredIterator struct {
	Event *SwapRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRouterOwnershipTransferred)
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
		it.Event = new(SwapRouterOwnershipTransferred)
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
func (it *SwapRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRouterOwnershipTransferred represents a OwnershipTransferred event raised by the SwapRouter contract.
type SwapRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapRouter *SwapRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapRouterOwnershipTransferredIterator{contract: _SwapRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapRouter *SwapRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRouterOwnershipTransferred)
				if err := _SwapRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapRouter *SwapRouterFilterer) ParseOwnershipTransferred(log types.Log) (*SwapRouterOwnershipTransferred, error) {
	event := new(SwapRouterOwnershipTransferred)
	if err := _SwapRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
