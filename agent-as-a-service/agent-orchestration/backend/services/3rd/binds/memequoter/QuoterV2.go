// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package memequoter

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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	Fee               *big.Int
	SqrtPriceLimitX96 *big.Int
}

// QuoterV2MetaData contains all meta data concerning the QuoterV2 contract.
var QuoterV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETHArg\",\"type\":\"address\"}],\"name\":\"setWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657611c50908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80632f80bb1d146100d7578063485cc955146100d25780635b769f3c146100cd578063715018a6146100c85780638da5cb5b146100c3578063ad5c4648146100be578063bd21704a146100b9578063c45a0155146100b4578063c6a5026a146100af578063cdca1753146100aa578063f2fde38b146100a55763fa461e33146100a057600080fd5b61066a565b6105d7565b6105bd565b610598565b61056f565b610515565b610484565b61045b565b6103fa565b6103b3565b6102ba565b61027b565b634e487b7160e01b600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761010e57604052565b6100dc565b6060810190811067ffffffffffffffff82111761010e57604052565b90601f8019910116810190811067ffffffffffffffff82111761010e57604052565b67ffffffffffffffff811161010e57601f01601f191660200190565b81601f820112156101b45780359061018482610151565b92610192604051948561012f565b828452602083830101116101b457816000926020809301838601378301015290565b600080fd5b60406003198201126101b4576004359067ffffffffffffffff82116101b4576101e49160040161016d565b9060243590565b94939290916080860192865260209260808488015281518091528360a0880192019060005b81811061025e57505050858103604087015282808351928381520192019260005b8281106102445750505060609150930152565b845163ffffffff1684529381019392810192600101610231565b82516001600160a01b031684529285019291850191600101610210565b346101b4576102a561029561028f366101b9565b90611281565b90604094929451948594856101eb565b0390f35b6001600160a01b038116036101b457565b346101b45760403660031901126101b4576004356102d7816102a9565b6103276024356102e6816102a9565b6000549261030b60ff8560081c1615809581966103a5575b8115610385575b5061089e565b8361031e600160ff196000541617600055565b61036c57610901565b61032d57005b61033d61ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61038061010061ff00196000541617600055565b610901565b303b15915081610397575b5038610305565b6001915060ff161438610390565b600160ff82161091506102fe565b346101b45760203660031901126101b4576004356103d0816102a9565b6103d86107fd565b606680546001600160a01b0319166001600160a01b0392909216919091179055005b346101b457600080600319360112610458576104146107fd565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b346101b45760003660031901126101b4576033546040516001600160a01b039091168152602090f35b346101b45760003660031901126101b4576066546040516001600160a01b039091168152602090f35b60a09060031901126101b457604051906104c6826100f2565b816004356104d3816102a9565b81526024356104e1816102a9565b6020820152604435604082015260643562ffffff811681036101b4576060820152608060843591610511836102a9565b0152565b346101b45760a03660031901126101b4576102a561053a610535366104ad565b611144565b604080519485526001600160a01b03909316602085015263ffffffff9091169183019190915260608201529081906080820190565b346101b45760003660031901126101b4576065546040516001600160a01b039091168152602090f35b346101b45760a03660031901126101b4576102a561053a6105b8366104ad565b610c12565b346101b4576102a56102956105d1366101b9565b90611033565b346101b45760203660031901126101b4576004356105f4816102a9565b6105fc6107fd565b6001600160a01b038116156106165761061490610855565b005b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346101b45760603660031901126101b45760443560243560043567ffffffffffffffff83116101b4576106a26004933690850161016d565b6000906106c2828413918280156107f4575b6106bd90610a4b565b611394565b919590926106e28385896106dd60655460018060a01b031690565b611346565b50156107c95761071260e0926106fa61071e93610a68565b976001600160a01b0380871690821610979590610b13565b6001600160a01b031690565b60405196878092633850c7bd851b82525afa9283156107c4576060958394610789575b501561075a575060405192835260208301526040820152fd5b9260675480610778575b505060405192835260208301526040820152fd5b6107829114610a4b565b3880610764565b9093506107ae91925060e03d81116107bd575b6107a6818361012f565b810190610a9b565b50505050509190919238610741565b503d61079c565b610b07565b9361071260e0926107dc61071e93610a68565b976001600160a01b0381811690871610979590610b13565b508386136106b4565b6033546001600160a01b0316330361081157565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b156108a557565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b6000549161092460ff8460081c1615809481956103a5578115610385575061089e565b82610937600160ff196000541617600055565b6109d2575b61095660ff60005460081c16610951816109eb565b6109eb565b61095f33610855565b60018060a01b0390816bffffffffffffffffffffffff60a01b93168360655416176065551690606654161760665561099357565b6109a361ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1565b6109e661010061ff00196000541617600055565b61093c565b156109f257565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b156101b457565b634e487b7160e01b600052601160045260246000fd5b600160ff1b8114610a795760000390565b610a52565b51908160020b82036101b457565b519061ffff821682036101b457565b908160e09103126101b4578051610ab1816102a9565b91610abe60208301610a7e565b91610acb60408201610a8c565b91610ad860608301610a8c565b91610ae560808201610a8c565b9160a082015160ff811681036101b45760c09092015180151581036101b45790565b6040513d6000823e3d90fd5b6065546001600160a01b0393610b379391851692610b31929061157f565b906115f0565b1690565b91908260409103126101b4576020825192015190565b60005b838110610b645750506000910152565b8181015183820152602001610b54565b90602091610b8d81518092818552858086019101610b51565b601f01601f1916010190565b6001600160a01b039182168152911515602083015260408201929092529116606082015260a060808201819052610bd292910190610b74565b90565b3d15610c00573d90610be682610151565b91610bf4604051938461012f565b82523d6000602084013e565b606090565b91908203918211610a7957565b80516020820180516060840180516000969591946001600160a01b03948894859485949293859392891690891681811093610c54929062ffffff165b91610b13565b975a996040948594610d15610c6b87870151610ddd565b60808701519094906001600160a01b03165b808716610dca57508215610da457610d07610cc5610cb7610cad6401000276a45b9a5b516001600160a01b031690565b935162ffffff1690565b9b516001600160a01b031690565b89519b8c93602085019192602b936bffffffffffffffffffffffff19809360601b16845262ffffff60e81b9060e81b16601484015260601b1660178201520190565b03601f1981018a528961012f565b610d35865198899687958694630251596160e31b86523060048701610b99565b03928b165af19182610d79575b5050610d7257505050610d6a929350610d64610d5c610bd5565b925a90610c05565b91610dec565b929391929091565b9250925092565b81610d9892903d10610d9d575b610d90818361012f565b810190610b3b565b610d42565b503d610d86565b610d07610cc5610cb7610cad73fffd8963efd1fc6a506488495d951d5263988d25610c9e565b610cc5610cb7610cad610d07939a610ca0565b600160ff1b8110156101b45790565b604051633850c7bd60e01b815291939192919060e0846004816001600160a01b0389165afa9384156107c457600094610ec8575b5080805160608103610e54575050610e4481602080610e4d94518301019101610f5f565b9195909661172b565b9293929190565b604411610e9057610e74816024806004610e8c9501518301019101610eef565b60405162461bcd60e51b815291829160048301610f4e565b0390fd5b60405162461bcd60e51b815260206004820152601060248201526f2ab732bc3832b1ba32b21032b93937b960811b6044820152606490fd5b610ee191945060e03d81116107bd576107a6818361012f565b505050505090509238610e20565b6020818303126101b45780519067ffffffffffffffff82116101b4570181601f820112156101b4578051610f2281610151565b92610f30604051948561012f565b818452602082840101116101b457610bd29160208085019101610b51565b906020610bd2928181520190610b74565b908160609103126101b457805191610bd260406020840151610f80816102a9565b9301610a7e565b67ffffffffffffffff811161010e5760051b60200190565b90610fa982610f87565b610fb6604051918261012f565b8281528092610fc7601f1991610f87565b0190602036910137565b8051821015610fe55760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90601f8201809211610a7957565b6017019081601711610a7957565b91908201809211610a7957565b6000198114610a795760010190565b91906110c56110ba926000948561105161104c83611380565b610f9f565b96826110f66110f06110d46110e961106c61104c8e99611380565b9a869c8d915b62ffffff61107f89611394565b91604094919490815195611092876100f2565b6001600160a01b03918216875216602086015284015216606082015260808101899052610c12565b95929d90938c610fd1565b6001600160a01b039091169052565b6110de8d8d610fd1565b9063ffffffff169052565b8895611017565b98611024565b9560428251101560001461113a57509061111286949392611453565b93976110c59693949093919290916110f6916110f0916110d4916110e991906110ba90611072565b9897955050505050565b80516020820180516060840180516000969587956001600160a01b039594861690861681811095889586958695946111d59492939261118a9290919062ffffff16610c4e565b9860808501809b826111a2835160018060a01b031690565b161561126f575b5a9a6040968796610c7d6111c76111c28a8d0151610ddd565b610a68565b95516001600160a01b031690565b6111f5865198899687958694630251596160e31b86523060048701610b99565b03928c165af19182611254575b505061124b57505050610d6a93945061123861071261122a611222610bd5565b935a90610c05565b94516001600160a01b031690565b610dec576112466000606755565b610dec565b93509350939050565b8161126a92903d10610d9d57610d90818361012f565b611202565b61127c6040880151606755565b6111a9565b91906110c56110ba926000948561129a61104c83611380565b96826113026110f06110d46110e96112b561104c8e99611380565b9a869c8d915b62ffffff6112c889611394565b9193906040908151956112da876100f2565b6001600160a01b03918216875216602086015284015216606082015260808101899052611144565b9560428251101560001461113a57509061131e86949392611453565b93976110c5969394909391929091611302916110f0916110d4916110e991906110ba906112bb565b6001600160a01b039361135f939192610b31929061157f565b168033036101b45790565b634e487b7160e01b600052601260045260246000fd5b516013198101908111610a79576017900490565b906113a360148351101561140f565b602082015160601c9160178151106113d35760376017820151916113cb602b8251101561140f565b015160601c91565b60405162461bcd60e51b8152602060048201526014602482015273746f55696e7432345f6f75744f66426f756e647360601b6044820152606490fd5b1561141657565b60405162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b6044820152606490fd5b80516016199182820190828211610a79576114788261147181610ffb565b1015611502565b611486601761147184611009565b61149b815161149484611009565b111561153f565b601783036114b85750505050604051600081526020810160405290565b601760405194601f8416801560051b9182828901019687010193010101905b8084106114ef5750508252601f01601f191660405290565b90928351815260208091019301906114d7565b1561150957565b60405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b6044820152606490fd5b1561154657565b60405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606490fd5b9162ffffff9160006040805161159481610113565b8281526020810183905201526001600160a01b0390808216858316116115d5575b81604051956115c387610113565b16855216602084015216604082015290565b936115b5565b908160209103126101b45751610bd2816102a9565b8151602083015191926001600160a01b03928316918316828110156101b45760209362ffffff6040606494015116956040519687958694630b4c774160e11b8652600486015260248501526044840152165afa9081156107c457600091611655575090565b610bd2915060203d8111611676575b61166e818361012f565b8101906115db565b503d611664565b908160209103126101b457610bd290610a7e565b60020b9060020b9081156116b457627fffff198114600019831416610a79570590565b61136a565b9060020b9081156116b45760020b0790565b908160209103126101b4575190565b60ff1660ff039060ff8211610a7957565b91909163ffffffff80809416911601918211610a7957565b60010b617fff8114610a795760010190565b63ffffffff9081166000190191908211610a7957565b6040516334324e9f60e21b8152909392600092916020816004816001600160a01b038a165afa80156107c45761177161177d9161178393600091611bc8575b5085611691565b60020b60081d60020b90565b60010b90565b6040516334324e9f60e21b81526020816004816001600160a01b038b165afa80156107c4576117c26117cd916117d393600091611ba9575b5086611691565b6101009060020b0790565b60ff1690565b6040516334324e9f60e21b81529095906020816004816001600160a01b038c165afa80156107c45761177161177d9161181493600091611ba9575086611691565b6040516334324e9f60e21b81529093906020816004816001600160a01b038d165afa80156107c4576117c26117cd9161185693600091611b8a575b5084611691565b60405163299ce14b60e11b8152600186900b60048201529093906020816024816001600160a01b038e165afa9081156107c457600091611b6b575b50600160ff86161b1615159586611b07575b86611af6575b60405163299ce14b60e11b8152600183900b60048201526020816024816001600160a01b038f165afa9081156107c457600091611ad7575b50600160ff8b161b1615159283611a63575b83611a51575b50508460010b8160010b90808212918215611a31575b505015611a265796939293969590965b60ff60001991161b5b8360010b8760010b136119f9578360010b8760010b146119dc575b60405163299ce14b60e11b8152600188900b6004820152916020836024816001600160a01b038e165afa80156107c45761199161199e93611998926119a4966000916119ad575b5016611be7565b61ffff1690565b906116eb565b95611703565b94600019611928565b6119cf915060203d6020116119d5575b6119c7818361012f565b8101906116cb565b3861198a565b503d6119bd565b6119f36119e8866116da565b60ff60001991161c90565b16611943565b5095945095505050611a16575b611a0d5790565b610bd290611715565b90611a2090611715565b90611a06565b90969395909261191f565b14905080611a41575b388061190f565b5060ff841660ff89161115611a3a565b600290810b91900b12915038806118f9565b6040516334324e9f60e21b81529093506020816004816001600160a01b038f165afa80156107c457611a9e91600091611aa8575b50826116b9565b60020b15926118f3565b611aca915060203d602011611ad0575b611ac2818361012f565b81019061167d565b38611a97565b503d611ab8565b611af0915060203d6020116119d5576119c7818361012f565b386118e1565b95508160020b8660020b13956118a9565b6040516334324e9f60e21b81529096506020816004816001600160a01b038e165afa80156107c457611b4291600091611b4c575b50836116b9565b60020b15956118a3565b611b65915060203d602011611ad057611ac2818361012f565b38611b3b565b611b84915060203d6020116119d5576119c7818361012f565b38611891565b611ba3915060203d602011611ad057611ac2818361012f565b3861184f565b611bc2915060203d602011611ad057611ac2818361012f565b386117bb565b611be1915060203d602011611ad057611ac2818361012f565b3861176a565b806000915b611bf4575090565b9061ffff809116908114610a795760010190600019810190808211610a79571680611bec56fea26469706673582212201b23cc68907f9e734d464818991239120a21043c25109f73414959d76fa268c564736f6c63430008130033",
}

// QuoterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterV2MetaData.ABI instead.
var QuoterV2ABI = QuoterV2MetaData.ABI

// QuoterV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuoterV2MetaData.Bin instead.
var QuoterV2Bin = QuoterV2MetaData.Bin

// DeployQuoterV2 deploys a new Ethereum contract, binding an instance of QuoterV2 to it.
func DeployQuoterV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QuoterV2, error) {
	parsed, err := QuoterV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuoterV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuoterV2{QuoterV2Caller: QuoterV2Caller{contract: contract}, QuoterV2Transactor: QuoterV2Transactor{contract: contract}, QuoterV2Filterer: QuoterV2Filterer{contract: contract}}, nil
}

// QuoterV2 is an auto generated Go binding around an Ethereum contract.
type QuoterV2 struct {
	QuoterV2Caller     // Read-only binding to the contract
	QuoterV2Transactor // Write-only binding to the contract
	QuoterV2Filterer   // Log filterer for contract events
}

// QuoterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterV2Session struct {
	Contract     *QuoterV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterV2CallerSession struct {
	Contract *QuoterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QuoterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterV2TransactorSession struct {
	Contract     *QuoterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QuoterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterV2Raw struct {
	Contract *QuoterV2 // Generic contract binding to access the raw methods on
}

// QuoterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterV2CallerRaw struct {
	Contract *QuoterV2Caller // Generic read-only contract binding to access the raw methods on
}

// QuoterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterV2TransactorRaw struct {
	Contract *QuoterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoterV2 creates a new instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2(address common.Address, backend bind.ContractBackend) (*QuoterV2, error) {
	contract, err := bindQuoterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuoterV2{QuoterV2Caller: QuoterV2Caller{contract: contract}, QuoterV2Transactor: QuoterV2Transactor{contract: contract}, QuoterV2Filterer: QuoterV2Filterer{contract: contract}}, nil
}

// NewQuoterV2Caller creates a new read-only instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Caller(address common.Address, caller bind.ContractCaller) (*QuoterV2Caller, error) {
	contract, err := bindQuoterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Caller{contract: contract}, nil
}

// NewQuoterV2Transactor creates a new write-only instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*QuoterV2Transactor, error) {
	contract, err := bindQuoterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Transactor{contract: contract}, nil
}

// NewQuoterV2Filterer creates a new log filterer instance of QuoterV2, bound to a specific deployed contract.
func NewQuoterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*QuoterV2Filterer, error) {
	contract, err := bindQuoterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterV2Filterer{contract: contract}, nil
}

// bindQuoterV2 binds a generic wrapper to an already deployed contract.
func bindQuoterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterV2 *QuoterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterV2.Contract.QuoterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterV2 *QuoterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterV2 *QuoterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuoterV2 *QuoterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuoterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuoterV2 *QuoterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuoterV2 *QuoterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuoterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_QuoterV2 *QuoterV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_QuoterV2 *QuoterV2Session) WETH() (common.Address, error) {
	return _QuoterV2.Contract.WETH(&_QuoterV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_QuoterV2 *QuoterV2CallerSession) WETH() (common.Address, error) {
	return _QuoterV2.Contract.WETH(&_QuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2Session) Factory() (common.Address, error) {
	return _QuoterV2.Contract.Factory(&_QuoterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_QuoterV2 *QuoterV2CallerSession) Factory() (common.Address, error) {
	return _QuoterV2.Contract.Factory(&_QuoterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuoterV2 *QuoterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuoterV2 *QuoterV2Session) Owner() (common.Address, error) {
	return _QuoterV2.Contract.Owner(&_QuoterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuoterV2 *QuoterV2CallerSession) Owner() (common.Address, error) {
	return _QuoterV2.Contract.Owner(&_QuoterV2.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2Caller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _QuoterV2.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _QuoterV2.Contract.UniswapV3SwapCallback(&_QuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_QuoterV2 *QuoterV2CallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _QuoterV2.Contract.UniswapV3SwapCallback(&_QuoterV2.CallOpts, amount0Delta, amount1Delta, path)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_QuoterV2 *QuoterV2Transactor) Initialize(opts *bind.TransactOpts, _factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "initialize", _factory, _WETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_QuoterV2 *QuoterV2Session) Initialize(_factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.Initialize(&_QuoterV2.TransactOpts, _factory, _WETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _factory, address _WETH) returns()
func (_QuoterV2 *QuoterV2TransactorSession) Initialize(_factory common.Address, _WETH common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.Initialize(&_QuoterV2.TransactOpts, _factory, _WETH)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInput(&_QuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInput(&_QuoterV2.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xc6a5026a.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactInputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutput(&_QuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutput(&_QuoterV2.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Transactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2Session) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutputSingle(&_QuoterV2.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xbd21704a.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,uint24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_QuoterV2 *QuoterV2TransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _QuoterV2.Contract.QuoteExactOutputSingle(&_QuoterV2.TransactOpts, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuoterV2 *QuoterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuoterV2 *QuoterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _QuoterV2.Contract.RenounceOwnership(&_QuoterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuoterV2 *QuoterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuoterV2.Contract.RenounceOwnership(&_QuoterV2.TransactOpts)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_QuoterV2 *QuoterV2Transactor) SetWETH(opts *bind.TransactOpts, WETHArg common.Address) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "setWETH", WETHArg)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_QuoterV2 *QuoterV2Session) SetWETH(WETHArg common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.SetWETH(&_QuoterV2.TransactOpts, WETHArg)
}

// SetWETH is a paid mutator transaction binding the contract method 0x5b769f3c.
//
// Solidity: function setWETH(address WETHArg) returns()
func (_QuoterV2 *QuoterV2TransactorSession) SetWETH(WETHArg common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.SetWETH(&_QuoterV2.TransactOpts, WETHArg)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuoterV2 *QuoterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuoterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuoterV2 *QuoterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.TransferOwnership(&_QuoterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuoterV2 *QuoterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuoterV2.Contract.TransferOwnership(&_QuoterV2.TransactOpts, newOwner)
}

// QuoterV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the QuoterV2 contract.
type QuoterV2InitializedIterator struct {
	Event *QuoterV2Initialized // Event containing the contract specifics and raw log

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
func (it *QuoterV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterV2Initialized)
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
		it.Event = new(QuoterV2Initialized)
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
func (it *QuoterV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterV2Initialized represents a Initialized event raised by the QuoterV2 contract.
type QuoterV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QuoterV2 *QuoterV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*QuoterV2InitializedIterator, error) {

	logs, sub, err := _QuoterV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QuoterV2InitializedIterator{contract: _QuoterV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QuoterV2 *QuoterV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QuoterV2Initialized) (event.Subscription, error) {

	logs, sub, err := _QuoterV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterV2Initialized)
				if err := _QuoterV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_QuoterV2 *QuoterV2Filterer) ParseInitialized(log types.Log) (*QuoterV2Initialized, error) {
	event := new(QuoterV2Initialized)
	if err := _QuoterV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuoterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuoterV2 contract.
type QuoterV2OwnershipTransferredIterator struct {
	Event *QuoterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuoterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuoterV2OwnershipTransferred)
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
		it.Event = new(QuoterV2OwnershipTransferred)
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
func (it *QuoterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuoterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuoterV2OwnershipTransferred represents a OwnershipTransferred event raised by the QuoterV2 contract.
type QuoterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuoterV2 *QuoterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuoterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuoterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuoterV2OwnershipTransferredIterator{contract: _QuoterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuoterV2 *QuoterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuoterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuoterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuoterV2OwnershipTransferred)
				if err := _QuoterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuoterV2 *QuoterV2Filterer) ParseOwnershipTransferred(log types.Log) (*QuoterV2OwnershipTransferred, error) {
	event := new(QuoterV2OwnershipTransferred)
	if err := _QuoterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
