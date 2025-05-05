// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package load_balancer

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

// LoadBalancerMetaData contains all meta data concerning the LoadBalancer contract.
var LoadBalancerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ClusterGroupAlreadyExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ClusterGroupNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedApproval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveCluster\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveClusterGroup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSolutions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"groupId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"}],\"name\":\"ClusterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"groupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ClusterGroupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"groupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ClusterGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"groupId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"}],\"name\":\"ClusterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"groupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clusterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InferencePerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerDeauthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_clusterInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minSolutionsThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"clusterIds\",\"type\":\"uint256[]\"}],\"name\":\"addClustersToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"authorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"clusterIds\",\"type\":\"uint256[]\"}],\"name\":\"createGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"deauthorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getClusterIdsOfGroup\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getClustersGroupInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGPUManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelCollectionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPromptSchedulerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWEAITokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"rawFlag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"promptScheduler_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"modelCollection_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSolutionsThreshold_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"groupName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"clusterIds\",\"type\":\"uint256[]\"}],\"name\":\"removeClustersFromGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"removeGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minSolutionsThreshold_\",\"type\":\"uint256\"}],\"name\":\"setMinSolutionsThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"modelCollection_\",\"type\":\"address\"}],\"name\":\"updateModelCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657612fa6908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80630305ea0114612389578063267c85071461229657806328d721001461224457806328e381d114611e925780633126394214611dcb5780633a6f819014611b005780633f4ba83a14611a0e57806345125f5e146111f05780634b083ddd146111555780634c255986146110cc57806354fd4d50146110505780635c975abb1461100f5780636246b31014610e9a57806369c2f15614610e5b5780636de7c6b714610e1f578063715018a614610d815780637c612f0314610d2f5780638456cb5914610c995780638da5cb5b14610c47578063ac1250d414610bd4578063adb1816c14610b82578063d2a1f116146105e1578063d36c71e31461058f578063f2fde38b146104a5578063f3ae24151461043b5763f7013ef61461013a57600080fd5b346104365760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365761017161247a565b6024359073ffffffffffffffffffffffffffffffffffffffff9182811680910361043657604435908382168092036104365760643591848316809303610436576000549460ff8660081c161594858096610429575b8015610412575b1561038e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00968660018983161760005561035f575b50169182158015610357575b801561034f575b8015610347575b61031d5761023b60ff60005460081c16610236816128ef565b6128ef565b61024433612882565b6102756000549660ff8860081c169061025c826128ef565b610265826128ef565b60655416606555610236816128ef565b60016097557fffffffffffffffffffffffff0000000000000000000000000000000000000000928360c954161760c9558260ca54161760ca558160cb54161760cb5560cf54161760cf5560843560d0556102cb57005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b60046040517f5cb045db000000000000000000000000000000000000000000000000000000008152fd5b50831561021d565b508115610216565b50801561020f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011760005538610203565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b1580156101cd5750600160ff8816146101cd565b50600160ff8816106101c6565b600080fd5b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365773ffffffffffffffffffffffffffffffffffffffff61048761247a565b1660005260cd602052602060ff604060002054166040519015158152f35b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610436576104dc61247a565b6104e4612803565b73ffffffffffffffffffffffffffffffffffffffff81161561050b5761050990612882565b005b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602073ffffffffffffffffffffffffffffffffffffffff60ca5416604051908152f35b346104365760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365760043567ffffffffffffffff811161043657610630903690600401612633565b60243567ffffffffffffffff8111610436576106509036906004016124cb565b9190604051602081019061067c602082865161066f8187858b01612651565b8101038084520182612581565b519020928360005260cc60205261069a6003604060002001546126eb565b610b43578360005260cc60205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000815416815583519067ffffffffffffffff8211610a2457600301906106f282546126eb565b601f8111610b06575b50806020601f8211600114610a6657600091610a5b575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790555b80610788575b837fe41c83911759458108d02eee33d682c09aab4e96ff6d98687f4ed0ec44c8c5fe61078385604051918291602083526020830190612674565b0390a2005b60005b8181106107985750610749565b6107a3818385612a28565b3563ffffffff81118015610a53575b61031d578060005260ce60205260ff6040600020541661031d576107d68133612cbd565b6107e563ffffffff8216612a50565b6040516107f181612565565b60018152602081018681528260005260ce602052610842604060002092511515839060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b5180519067ffffffffffffffff8211610a245761086260018401546126eb565b601f81116109e4575b50602090601f83116001146109155791806001926109059796959460009261090a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82841b9260031b1c1916179101555b8660005260cc6020526108da816001604060002001612e5c565b867f0ad827fc9c569a07a74fc468b7163fdd515381d915a1ce973559a305bad3cfa2600080a36129fb565b61078b565b015190508b8061088e565b906001840160005260206000209160005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0851681106109cc575092600192839261090598979695837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0811610610995575b505050811b019101556108c0565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558b8080610987565b91926020600181928685015181550194019201610926565b610a14906001850160005260206000206005601f8601811c82019260208710610a1a575b601f01901c01906129e4565b8961086b565b9192508291610a08565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b5080156107b2565b905085015187610712565b60008481526020812092507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416905b818110610aee57509083600194939210610ab7575b5050811b019055610743565b8701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558780610aab565b9192602060018192868c015181550194019201610a96565b610b3390836000526020600020601f840160051c81019160208510610b39575b601f0160051c01906129e4565b866106fb565b9091508190610b26565b6040517f220c20640000000000000000000000000000000000000000000000000000000081526020600482015280610b7e6024820186612674565b0390fd5b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602073ffffffffffffffffffffffffffffffffffffffff60cf5416604051908152f35b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365760043560005260ce6020526040600020610c24600160ff835416920161273e565b90610c4360405192839215158352604060208401526040830190612674565b0390f35b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602073ffffffffffffffffffffffffffffffffffffffff60335416604051908152f35b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657610cd0612803565b610cd861297a565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060655416176065557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602073ffffffffffffffffffffffffffffffffffffffff60c95416604051908152f35b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657610db8612803565b600073ffffffffffffffffffffffffffffffffffffffff6033547fffffffffffffffffffffffff00000000000000000000000000000000000000008116603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602060d054604051908152f35b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657610e92612803565b60043560d055005b34610436576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365760043567ffffffffffffffff811161043657610eea90369060040161249d565b9091610ef4612803565b610efe8284612dd1565b928360005260cc82526001610f1881604060002001612be6565b805160005b818110610fa657877fda6239e3c1c4f60a7454566117b0db7929a39167ca9907e7fae4fbe9ed8205e4888861078389610f7a60038b8860005260cc865260406000209060008255810180546000825580610f8b575b505001612c35565b604051938385948552840191612c7e565b610f9f9160005287600020908101906129e4565b8980610f72565b8251811015610fe0578086610fdb9260051b8501015160005260ce8752610fd68560406000206000815501612c35565b6129fb565b610f1d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602060ff606554166040519015158152f35b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657610c4360405161108e81612565565b600681527f76302e302e3100000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190612674565b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365773ffffffffffffffffffffffffffffffffffffffff61111861247a565b611120612803565b16801561031d577fffffffffffffffffffffffff000000000000000000000000000000000000000060cf54161760cf55600080f35b34610436576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365760043567ffffffffffffffff8111610436576111a5903690600401612633565b6040516111c183828161066f8183019687815193849201612651565b51902060005260cc81526111db6001604060002001612be6565b90610c436040519282849384528301906126b7565b346104365760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365760043567ffffffffffffffff81116104365761123f90369060040161249d565b60243567ffffffffffffffff81116104365761125f90369060040161249d565b929091604435151560443503610436576002609754146119b057600260975561128661297a565b831561031d5761129591612dd1565b8060005260cc60205260016040600020015490811561031d578060005260cc60205261ffff6040600020541690819373ffffffffffffffffffffffffffffffffffffffff60c95416925b61ffff808216146119815761ffff60019116018461ffff82161015611979575b8260005260cc60205263ffffffff61131e826001604060002001612d99565b90549060031b1c166040517fbce2845a000000000000000000000000000000000000000000000000000000008152816004820152602081602481895afa90811561165e5760009161194a575b506113a557508561ffff8216036112df5760046040517f99765bad000000000000000000000000000000000000000000000000000000008152fd5b918793918260005260cc60205261ffff604060002091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00008254161790556024602073ffffffffffffffffffffffffffffffffffffffff60c95416604051928380927f963a027800000000000000000000000000000000000000000000000000000000825263ffffffff891660048301525afa90811561165e57600091611918575b50806117a4575b73ffffffffffffffffffffffffffffffffffffffff60cb541673ffffffffffffffffffffffffffffffffffffffff60ca5416917f617070726f766528616464726573732c75696e7432353629000000000000000060206040516114b081612565565b601881520152604051927f095ea7b3000000000000000000000000000000000000000000000000000000006020850152602484015260448301526044825281608081011067ffffffffffffffff608084011117610a24578160009291608084930160405282602083519301915af1611526612f40565b9015908115611774575b5061174a576044351561166a5773ffffffffffffffffffffffffffffffffffffffff60ca541693602060405180967f5cc6873100000000000000000000000000000000000000000000000000000000825263ffffffff8716600483015260806024830152816000816115a660848201888a612c7e565b3360448301526044351515606483015203925af194851561165e5763ffffffff6116219567ffffffffffffffff6020987fe6cc06a7b773e04156e2c577e692ac9409358de02566d0fdaff15e8604ecac5694600091611631575b50169687955b60405193849316835260408a84015233956040840191612c7e565b0390a46001609755604051908152f35b61165191508a3d8c11611657575b6116498183612581565b810190612db1565b8a611600565b503d61163f565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff60ca541693602060405180967fde1ce2bb00000000000000000000000000000000000000000000000000000000825263ffffffff8716600483015260606024830152816000816116d360648201888a612c7e565b33604483015203925af194851561165e5763ffffffff6116219567ffffffffffffffff6020987fe6cc06a7b773e04156e2c577e692ac9409358de02566d0fdaff15e8604ecac569460009161172d575b5016968795611606565b61174491508a3d8c11611657576116498183612581565b8a611723565b60046040517f39b83b3f000000000000000000000000000000000000000000000000000000008152fd5b8051801515925082611789575b505085611530565b61179c9250602080918301019101612a38565b158580611781565b73ffffffffffffffffffffffffffffffffffffffff60cb541660405180606081011067ffffffffffffffff606083011117610a245760258160607fffffffff0000000000000000000000000000000000000000000000000000000093016040528181527f7432353629000000000000000000000000000000000000000000000000000000604060208301927f7472616e7366657246726f6d28616464726573732c616464726573732c75696e84520152201690604051916020830152336024830152306044830152826064830152606482528160a081011067ffffffffffffffff60a084011117610a2457816000929160a084930160405282602083519301915af16118ae612f40565b90159081156118e8575b501561144e5760046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b80518015159250826118fd575b5050866118b8565b6119109250602080918301019101612a38565b1586806118f5565b90506020813d602011611942575b8161193360209383612581565b81010312610436575185611447565b3d9150611926565b61196c915060203d602011611972575b6119648183612581565b810190612a38565b8961136a565b503d61195a565b5060006112ff565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152fd5b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657611a45612803565b60655460ff811615611aa2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166065557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b3461043657611b0e366124fc565b80939192931561031d57611b228483612dd1565b9060005b818110611b2f57005b611b3a818387612a28565b3563ffffffff8082118015611dc3575b61031d578160005260ce908160205260ff6040600020541661031d57611b7b90611b748433612cbd565b8316612a50565b60405190611b8882612565565b60018252611b97368a896125fc565b906020830191825283600052602052611be3604060002092511515839060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b5180519067ffffffffffffffff8211610a2457611c0360018401546126eb565b90601f91828111611d86575b506020918311600114611cb7579180600192611ca797969594600092611cac575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82841b9260031b1c1916179101555b8460005260cc602052611c7c816001604060002001612e5c565b847f0ad827fc9c569a07a74fc468b7163fdd515381d915a1ce973559a305bad3cfa2600080a36129fb565b611b26565b015190508c80611c30565b906001840160005260206000209160005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085168110611d6e5750926001928392611ca798979695837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0811610611d37575b505050811b01910155611c62565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558c8080611d29565b91926020600181928685015181550194019201611cc8565b611db490600186016000526020600020600585808801821c83019360208910611dba575b01901c01906129e4565b8b611c0f565b93508293611daa565b508115611b4a565b34610436576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610436576004359067ffffffffffffffff821161043657611e1f611e81923690600401612633565b604051611e3b83828161066f8183019687815193849201612651565b51902060005260cc8152610c4360406000209161ffff835416611e6c6001611e656003870161273e565b9501612be6565b90604051958695606087526060870190612674565b9285015283820360408501526126b7565b3461043657611ea0366124fc565b929091611eab612803565b831561031d57611eba91612dd1565b9160005b818110611ec757005b611ed2818385612a28565b35908160005260ce9160209280845260ff604060002054161561031d57611ef98233612cbd565b60405190611f0682612565565b60008252604051918583019067ffffffffffffffff9284831084841117610a245760019260405260008552878201948552856000528752611f7a604060002091511515829060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083541691151516179055565b0191518051918211610a24578190611f9284546126eb565b90601f91828111612216575b50879183116001146121775760009261216c575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790555b8560005260cc83526040600020926002600185019401908260005281815260406000205494851561213b57612098907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff966120428883540183612d99565b9390546120928a8401916120568387612d99565b9091600398891b1c907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b83612d99565b905490841b1c600052848452604060002055805496871561210c57806121079801926120c48484612d99565b81939154921b1b1916905555826000525260006040812055857fc264f0c7f8facadb2a37a6d959ec894d069f522bdabe18cd0cade329ea320217600080a36129fb565b611ebe565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b602484604051907f080240290000000000000000000000000000000000000000000000000000000082526004820152fd5b015190508980611fb2565b6000858152888120937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016905b898282106122005750509084600195949392106121c9575b505050811b019055611fe4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558980806121bc565b60018596829396860151815501950193016121a4565b61223e908660005289600020600585808801821c8301938d8910611dba5701901c01906129e4565b8b611f9e565b346104365760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261043657602073ffffffffffffffffffffffffffffffffffffffff60cb5416604051908152f35b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365773ffffffffffffffffffffffffffffffffffffffff6122e261247a565b6122ea612803565b168060005260cd60205260ff6040600020541661235f578060005260cd602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b600080a2005b60046040517feacfc0ae000000000000000000000000000000000000000000000000000000008152fd5b346104365760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104365773ffffffffffffffffffffffffffffffffffffffff6123d561247a565b6123dd612803565b168060005260cd60205260ff6040600020541615612450578060005260cd60205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557f20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c600080a2005b60046040517f82b42900000000000000000000000000000000000000000000000000000000008152fd5b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361043657565b9181601f840112156104365782359167ffffffffffffffff8311610436576020838186019501011161043657565b9181601f840112156104365782359167ffffffffffffffff8311610436576020808501948460051b01011161043657565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126104365767ffffffffffffffff9160043583811161043657826125479160040161249d565b9390939260243591821161043657612561916004016124cb565b9091565b6040810190811067ffffffffffffffff821117610a2457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610a2457604052565b67ffffffffffffffff8111610a2457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192612608826125c2565b916126166040519384612581565b829481845281830111610436578281602093846000960137010152565b9080601f830112156104365781602061264e933591016125fc565b90565b60005b8381106126645750506000910152565b8181015183820152602001612654565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936126b081518092818752878088019101612651565b0116010190565b90815180825260208080930193019160005b8281106126d7575050505090565b8351855293810193928101926001016126c9565b90600182811c92168015612734575b602083101461270557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916126fa565b9060405191826000825492612752846126eb565b9081845260019485811690816000146127c1575060011461277e575b505061277c92500383612581565b565b9093915060005260209081600020936000915b8183106127a957505061277c9350820101388061276e565b85548884018501529485019487945091830191612791565b905061277c9550602093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101388061276e565b73ffffffffffffffffffffffffffffffffffffffff60335416330361282457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6033549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617603355167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b156128f657565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b60ff6065541661298657565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b8181106129ef575050565b600081556001016129e4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146119815760010190565b9190811015610fe05760051b0190565b90816020910312610436575180151581036104365790565b63ffffffff73ffffffffffffffffffffffffffffffffffffffff8060c954169060409384519384917fbce2845a000000000000000000000000000000000000000000000000000000008352169283600483015281602460209687935afa908115612bdb57600091612bbe575b5015612b9557829060ca54169160248551809481937f54d6d8f700000000000000000000000000000000000000000000000000000000835260048301525afa918215612b8a57600092612b48575b505064ffffffffff60d054911610612b1f5750565b600490517fc731db19000000000000000000000000000000000000000000000000000000008152fd5b81813d8311612b83575b612b5c8183612581565b81010312612b7f57519064ffffffffff82168203612b7c57503880612b0a565b80fd5b5080fd5b503d612b52565b83513d6000823e3d90fd5b600484517f99a0c653000000000000000000000000000000000000000000000000000000008152fd5b612bd59150843d8611611972576119648183612581565b38612abc565b85513d6000823e3d90fd5b9060405191828154918282526020928383019160005283600020936000905b828210612c1b5750505061277c92500383612581565b855484526001958601958895509381019390910190612c05565b612c3f81546126eb565b9081612c49575050565b81601f60009311600114612c5b575055565b908083918252612c7a601f60208420940160051c8401600185016129e4565b5555565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff80911691600083815260cd60205260ff60408220541615612cf4575b50505050565b60208360cf5416926024604051809581937f6352211e00000000000000000000000000000000000000000000000000000000835260048301525afa918215612d8c578192612d4e575b505016036124505738808080612cee565b9091506020813d8211612d84575b81612d6960209383612581565b81010312612b7f5751908282168203612b7c57503880612d3d565b3d9150612d5c565b50604051903d90823e3d90fd5b8054821015610fe05760005260206000200190600090565b90816020910312610436575167ffffffffffffffff811681036104365790565b604051602081019083838337612df7602082868101600083820152038084520182612581565b519020918260005260cc602052612e156003604060002001546126eb565b15612e1f57505090565b610b7e6040519283927fec91e269000000000000000000000000000000000000000000000000000000008452602060048501526024840191612c7e565b919060018301600090828252806020526040822054612f0f5784549468010000000000000000861015612ee257612ed784612ea1886001604098999a01855584612d99565b9091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b549382526020522055565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b602483604051907f346c4a0e0000000000000000000000000000000000000000000000000000000082526004820152fd5b3d15612f6b573d90612f51826125c2565b91612f5f6040519384612581565b82523d6000602084013e565b60609056fea2646970667358221220262aaab0efa924b1b723afd7bdc25095c5ca09e72ed5a08e02d3cf4ee0171b6664736f6c63430008140033",
}

// LoadBalancerABI is the input ABI used to generate the binding from.
// Deprecated: Use LoadBalancerMetaData.ABI instead.
var LoadBalancerABI = LoadBalancerMetaData.ABI

// LoadBalancerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoadBalancerMetaData.Bin instead.
var LoadBalancerBin = LoadBalancerMetaData.Bin

// DeployLoadBalancer deploys a new Ethereum contract, binding an instance of LoadBalancer to it.
func DeployLoadBalancer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoadBalancer, error) {
	parsed, err := LoadBalancerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoadBalancerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoadBalancer{LoadBalancerCaller: LoadBalancerCaller{contract: contract}, LoadBalancerTransactor: LoadBalancerTransactor{contract: contract}, LoadBalancerFilterer: LoadBalancerFilterer{contract: contract}}, nil
}

// LoadBalancer is an auto generated Go binding around an Ethereum contract.
type LoadBalancer struct {
	LoadBalancerCaller     // Read-only binding to the contract
	LoadBalancerTransactor // Write-only binding to the contract
	LoadBalancerFilterer   // Log filterer for contract events
}

// LoadBalancerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoadBalancerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadBalancerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoadBalancerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadBalancerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoadBalancerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadBalancerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoadBalancerSession struct {
	Contract     *LoadBalancer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoadBalancerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoadBalancerCallerSession struct {
	Contract *LoadBalancerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LoadBalancerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoadBalancerTransactorSession struct {
	Contract     *LoadBalancerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LoadBalancerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoadBalancerRaw struct {
	Contract *LoadBalancer // Generic contract binding to access the raw methods on
}

// LoadBalancerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoadBalancerCallerRaw struct {
	Contract *LoadBalancerCaller // Generic read-only contract binding to access the raw methods on
}

// LoadBalancerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoadBalancerTransactorRaw struct {
	Contract *LoadBalancerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoadBalancer creates a new instance of LoadBalancer, bound to a specific deployed contract.
func NewLoadBalancer(address common.Address, backend bind.ContractBackend) (*LoadBalancer, error) {
	contract, err := bindLoadBalancer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoadBalancer{LoadBalancerCaller: LoadBalancerCaller{contract: contract}, LoadBalancerTransactor: LoadBalancerTransactor{contract: contract}, LoadBalancerFilterer: LoadBalancerFilterer{contract: contract}}, nil
}

// NewLoadBalancerCaller creates a new read-only instance of LoadBalancer, bound to a specific deployed contract.
func NewLoadBalancerCaller(address common.Address, caller bind.ContractCaller) (*LoadBalancerCaller, error) {
	contract, err := bindLoadBalancer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerCaller{contract: contract}, nil
}

// NewLoadBalancerTransactor creates a new write-only instance of LoadBalancer, bound to a specific deployed contract.
func NewLoadBalancerTransactor(address common.Address, transactor bind.ContractTransactor) (*LoadBalancerTransactor, error) {
	contract, err := bindLoadBalancer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerTransactor{contract: contract}, nil
}

// NewLoadBalancerFilterer creates a new log filterer instance of LoadBalancer, bound to a specific deployed contract.
func NewLoadBalancerFilterer(address common.Address, filterer bind.ContractFilterer) (*LoadBalancerFilterer, error) {
	contract, err := bindLoadBalancer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerFilterer{contract: contract}, nil
}

// bindLoadBalancer binds a generic wrapper to an already deployed contract.
func bindLoadBalancer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoadBalancerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoadBalancer *LoadBalancerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoadBalancer.Contract.LoadBalancerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoadBalancer *LoadBalancerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadBalancer.Contract.LoadBalancerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoadBalancer *LoadBalancerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoadBalancer.Contract.LoadBalancerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoadBalancer *LoadBalancerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoadBalancer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoadBalancer *LoadBalancerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadBalancer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoadBalancer *LoadBalancerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoadBalancer.Contract.contract.Transact(opts, method, params...)
}

// ClusterInfo is a free data retrieval call binding the contract method 0xac1250d4.
//
// Solidity: function _clusterInfo(uint256 ) view returns(bool isRegistered, string groupName)
func (_LoadBalancer *LoadBalancerCaller) ClusterInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsRegistered bool
	GroupName    string
}, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "_clusterInfo", arg0)

	outstruct := new(struct {
		IsRegistered bool
		GroupName    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRegistered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.GroupName = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ClusterInfo is a free data retrieval call binding the contract method 0xac1250d4.
//
// Solidity: function _clusterInfo(uint256 ) view returns(bool isRegistered, string groupName)
func (_LoadBalancer *LoadBalancerSession) ClusterInfo(arg0 *big.Int) (struct {
	IsRegistered bool
	GroupName    string
}, error) {
	return _LoadBalancer.Contract.ClusterInfo(&_LoadBalancer.CallOpts, arg0)
}

// ClusterInfo is a free data retrieval call binding the contract method 0xac1250d4.
//
// Solidity: function _clusterInfo(uint256 ) view returns(bool isRegistered, string groupName)
func (_LoadBalancer *LoadBalancerCallerSession) ClusterInfo(arg0 *big.Int) (struct {
	IsRegistered bool
	GroupName    string
}, error) {
	return _LoadBalancer.Contract.ClusterInfo(&_LoadBalancer.CallOpts, arg0)
}

// MinSolutionsThreshold is a free data retrieval call binding the contract method 0x6de7c6b7.
//
// Solidity: function _minSolutionsThreshold() view returns(uint256)
func (_LoadBalancer *LoadBalancerCaller) MinSolutionsThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "_minSolutionsThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSolutionsThreshold is a free data retrieval call binding the contract method 0x6de7c6b7.
//
// Solidity: function _minSolutionsThreshold() view returns(uint256)
func (_LoadBalancer *LoadBalancerSession) MinSolutionsThreshold() (*big.Int, error) {
	return _LoadBalancer.Contract.MinSolutionsThreshold(&_LoadBalancer.CallOpts)
}

// MinSolutionsThreshold is a free data retrieval call binding the contract method 0x6de7c6b7.
//
// Solidity: function _minSolutionsThreshold() view returns(uint256)
func (_LoadBalancer *LoadBalancerCallerSession) MinSolutionsThreshold() (*big.Int, error) {
	return _LoadBalancer.Contract.MinSolutionsThreshold(&_LoadBalancer.CallOpts)
}

// GetClusterIdsOfGroup is a free data retrieval call binding the contract method 0x4b083ddd.
//
// Solidity: function getClusterIdsOfGroup(string name) view returns(uint256[])
func (_LoadBalancer *LoadBalancerCaller) GetClusterIdsOfGroup(opts *bind.CallOpts, name string) ([]*big.Int, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getClusterIdsOfGroup", name)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetClusterIdsOfGroup is a free data retrieval call binding the contract method 0x4b083ddd.
//
// Solidity: function getClusterIdsOfGroup(string name) view returns(uint256[])
func (_LoadBalancer *LoadBalancerSession) GetClusterIdsOfGroup(name string) ([]*big.Int, error) {
	return _LoadBalancer.Contract.GetClusterIdsOfGroup(&_LoadBalancer.CallOpts, name)
}

// GetClusterIdsOfGroup is a free data retrieval call binding the contract method 0x4b083ddd.
//
// Solidity: function getClusterIdsOfGroup(string name) view returns(uint256[])
func (_LoadBalancer *LoadBalancerCallerSession) GetClusterIdsOfGroup(name string) ([]*big.Int, error) {
	return _LoadBalancer.Contract.GetClusterIdsOfGroup(&_LoadBalancer.CallOpts, name)
}

// GetClustersGroupInfo is a free data retrieval call binding the contract method 0x31263942.
//
// Solidity: function getClustersGroupInfo(string name) view returns(string, uint16, uint256[])
func (_LoadBalancer *LoadBalancerCaller) GetClustersGroupInfo(opts *bind.CallOpts, name string) (string, uint16, []*big.Int, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getClustersGroupInfo", name)

	if err != nil {
		return *new(string), *new(uint16), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// GetClustersGroupInfo is a free data retrieval call binding the contract method 0x31263942.
//
// Solidity: function getClustersGroupInfo(string name) view returns(string, uint16, uint256[])
func (_LoadBalancer *LoadBalancerSession) GetClustersGroupInfo(name string) (string, uint16, []*big.Int, error) {
	return _LoadBalancer.Contract.GetClustersGroupInfo(&_LoadBalancer.CallOpts, name)
}

// GetClustersGroupInfo is a free data retrieval call binding the contract method 0x31263942.
//
// Solidity: function getClustersGroupInfo(string name) view returns(string, uint16, uint256[])
func (_LoadBalancer *LoadBalancerCallerSession) GetClustersGroupInfo(name string) (string, uint16, []*big.Int, error) {
	return _LoadBalancer.Contract.GetClustersGroupInfo(&_LoadBalancer.CallOpts, name)
}

// GetGPUManagerAddress is a free data retrieval call binding the contract method 0x7c612f03.
//
// Solidity: function getGPUManagerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCaller) GetGPUManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getGPUManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGPUManagerAddress is a free data retrieval call binding the contract method 0x7c612f03.
//
// Solidity: function getGPUManagerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerSession) GetGPUManagerAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetGPUManagerAddress(&_LoadBalancer.CallOpts)
}

// GetGPUManagerAddress is a free data retrieval call binding the contract method 0x7c612f03.
//
// Solidity: function getGPUManagerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCallerSession) GetGPUManagerAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetGPUManagerAddress(&_LoadBalancer.CallOpts)
}

// GetModelCollectionAddress is a free data retrieval call binding the contract method 0xadb1816c.
//
// Solidity: function getModelCollectionAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCaller) GetModelCollectionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getModelCollectionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetModelCollectionAddress is a free data retrieval call binding the contract method 0xadb1816c.
//
// Solidity: function getModelCollectionAddress() view returns(address)
func (_LoadBalancer *LoadBalancerSession) GetModelCollectionAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetModelCollectionAddress(&_LoadBalancer.CallOpts)
}

// GetModelCollectionAddress is a free data retrieval call binding the contract method 0xadb1816c.
//
// Solidity: function getModelCollectionAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCallerSession) GetModelCollectionAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetModelCollectionAddress(&_LoadBalancer.CallOpts)
}

// GetPromptSchedulerAddress is a free data retrieval call binding the contract method 0xd36c71e3.
//
// Solidity: function getPromptSchedulerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCaller) GetPromptSchedulerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getPromptSchedulerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPromptSchedulerAddress is a free data retrieval call binding the contract method 0xd36c71e3.
//
// Solidity: function getPromptSchedulerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerSession) GetPromptSchedulerAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetPromptSchedulerAddress(&_LoadBalancer.CallOpts)
}

// GetPromptSchedulerAddress is a free data retrieval call binding the contract method 0xd36c71e3.
//
// Solidity: function getPromptSchedulerAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCallerSession) GetPromptSchedulerAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetPromptSchedulerAddress(&_LoadBalancer.CallOpts)
}

// GetWEAITokenAddress is a free data retrieval call binding the contract method 0x28d72100.
//
// Solidity: function getWEAITokenAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCaller) GetWEAITokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "getWEAITokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWEAITokenAddress is a free data retrieval call binding the contract method 0x28d72100.
//
// Solidity: function getWEAITokenAddress() view returns(address)
func (_LoadBalancer *LoadBalancerSession) GetWEAITokenAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetWEAITokenAddress(&_LoadBalancer.CallOpts)
}

// GetWEAITokenAddress is a free data retrieval call binding the contract method 0x28d72100.
//
// Solidity: function getWEAITokenAddress() view returns(address)
func (_LoadBalancer *LoadBalancerCallerSession) GetWEAITokenAddress() (common.Address, error) {
	return _LoadBalancer.Contract.GetWEAITokenAddress(&_LoadBalancer.CallOpts)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_LoadBalancer *LoadBalancerCaller) IsManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "isManager", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_LoadBalancer *LoadBalancerSession) IsManager(account common.Address) (bool, error) {
	return _LoadBalancer.Contract.IsManager(&_LoadBalancer.CallOpts, account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address account) view returns(bool)
func (_LoadBalancer *LoadBalancerCallerSession) IsManager(account common.Address) (bool, error) {
	return _LoadBalancer.Contract.IsManager(&_LoadBalancer.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoadBalancer *LoadBalancerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoadBalancer *LoadBalancerSession) Owner() (common.Address, error) {
	return _LoadBalancer.Contract.Owner(&_LoadBalancer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoadBalancer *LoadBalancerCallerSession) Owner() (common.Address, error) {
	return _LoadBalancer.Contract.Owner(&_LoadBalancer.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoadBalancer *LoadBalancerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoadBalancer *LoadBalancerSession) Paused() (bool, error) {
	return _LoadBalancer.Contract.Paused(&_LoadBalancer.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoadBalancer *LoadBalancerCallerSession) Paused() (bool, error) {
	return _LoadBalancer.Contract.Paused(&_LoadBalancer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoadBalancer *LoadBalancerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LoadBalancer.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoadBalancer *LoadBalancerSession) Version() (string, error) {
	return _LoadBalancer.Contract.Version(&_LoadBalancer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoadBalancer *LoadBalancerCallerSession) Version() (string, error) {
	return _LoadBalancer.Contract.Version(&_LoadBalancer.CallOpts)
}

// AddClustersToGroup is a paid mutator transaction binding the contract method 0x3a6f8190.
//
// Solidity: function addClustersToGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactor) AddClustersToGroup(opts *bind.TransactOpts, groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "addClustersToGroup", groupName, clusterIds)
}

// AddClustersToGroup is a paid mutator transaction binding the contract method 0x3a6f8190.
//
// Solidity: function addClustersToGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerSession) AddClustersToGroup(groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.AddClustersToGroup(&_LoadBalancer.TransactOpts, groupName, clusterIds)
}

// AddClustersToGroup is a paid mutator transaction binding the contract method 0x3a6f8190.
//
// Solidity: function addClustersToGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) AddClustersToGroup(groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.AddClustersToGroup(&_LoadBalancer.TransactOpts, groupName, clusterIds)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerTransactor) AuthorizeManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "authorizeManager", account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerSession) AuthorizeManager(account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.AuthorizeManager(&_LoadBalancer.TransactOpts, account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) AuthorizeManager(account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.AuthorizeManager(&_LoadBalancer.TransactOpts, account)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xd2a1f116.
//
// Solidity: function createGroup(string name, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactor) CreateGroup(opts *bind.TransactOpts, name string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "createGroup", name, clusterIds)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xd2a1f116.
//
// Solidity: function createGroup(string name, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerSession) CreateGroup(name string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.CreateGroup(&_LoadBalancer.TransactOpts, name, clusterIds)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xd2a1f116.
//
// Solidity: function createGroup(string name, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) CreateGroup(name string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.CreateGroup(&_LoadBalancer.TransactOpts, name, clusterIds)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerTransactor) DeauthorizeManager(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "deauthorizeManager", account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerSession) DeauthorizeManager(account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.DeauthorizeManager(&_LoadBalancer.TransactOpts, account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address account) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) DeauthorizeManager(account common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.DeauthorizeManager(&_LoadBalancer.TransactOpts, account)
}

// Infer is a paid mutator transaction binding the contract method 0x45125f5e.
//
// Solidity: function infer(string groupName, bytes data, bool rawFlag) returns(uint256)
func (_LoadBalancer *LoadBalancerTransactor) Infer(opts *bind.TransactOpts, groupName string, data []byte, rawFlag bool) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "infer", groupName, data, rawFlag)
}

// Infer is a paid mutator transaction binding the contract method 0x45125f5e.
//
// Solidity: function infer(string groupName, bytes data, bool rawFlag) returns(uint256)
func (_LoadBalancer *LoadBalancerSession) Infer(groupName string, data []byte, rawFlag bool) (*types.Transaction, error) {
	return _LoadBalancer.Contract.Infer(&_LoadBalancer.TransactOpts, groupName, data, rawFlag)
}

// Infer is a paid mutator transaction binding the contract method 0x45125f5e.
//
// Solidity: function infer(string groupName, bytes data, bool rawFlag) returns(uint256)
func (_LoadBalancer *LoadBalancerTransactorSession) Infer(groupName string, data []byte, rawFlag bool) (*types.Transaction, error) {
	return _LoadBalancer.Contract.Infer(&_LoadBalancer.TransactOpts, groupName, data, rawFlag)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address gpuManager_, address promptScheduler_, address wEAIToken_, address modelCollection_, uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerTransactor) Initialize(opts *bind.TransactOpts, gpuManager_ common.Address, promptScheduler_ common.Address, wEAIToken_ common.Address, modelCollection_ common.Address, minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "initialize", gpuManager_, promptScheduler_, wEAIToken_, modelCollection_, minSolutionsThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address gpuManager_, address promptScheduler_, address wEAIToken_, address modelCollection_, uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerSession) Initialize(gpuManager_ common.Address, promptScheduler_ common.Address, wEAIToken_ common.Address, modelCollection_ common.Address, minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.Initialize(&_LoadBalancer.TransactOpts, gpuManager_, promptScheduler_, wEAIToken_, modelCollection_, minSolutionsThreshold_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address gpuManager_, address promptScheduler_, address wEAIToken_, address modelCollection_, uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) Initialize(gpuManager_ common.Address, promptScheduler_ common.Address, wEAIToken_ common.Address, modelCollection_ common.Address, minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.Initialize(&_LoadBalancer.TransactOpts, gpuManager_, promptScheduler_, wEAIToken_, modelCollection_, minSolutionsThreshold_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LoadBalancer *LoadBalancerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LoadBalancer *LoadBalancerSession) Pause() (*types.Transaction, error) {
	return _LoadBalancer.Contract.Pause(&_LoadBalancer.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LoadBalancer *LoadBalancerTransactorSession) Pause() (*types.Transaction, error) {
	return _LoadBalancer.Contract.Pause(&_LoadBalancer.TransactOpts)
}

// RemoveClustersFromGroup is a paid mutator transaction binding the contract method 0x28e381d1.
//
// Solidity: function removeClustersFromGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactor) RemoveClustersFromGroup(opts *bind.TransactOpts, groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "removeClustersFromGroup", groupName, clusterIds)
}

// RemoveClustersFromGroup is a paid mutator transaction binding the contract method 0x28e381d1.
//
// Solidity: function removeClustersFromGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerSession) RemoveClustersFromGroup(groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.RemoveClustersFromGroup(&_LoadBalancer.TransactOpts, groupName, clusterIds)
}

// RemoveClustersFromGroup is a paid mutator transaction binding the contract method 0x28e381d1.
//
// Solidity: function removeClustersFromGroup(string groupName, uint256[] clusterIds) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) RemoveClustersFromGroup(groupName string, clusterIds []*big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.RemoveClustersFromGroup(&_LoadBalancer.TransactOpts, groupName, clusterIds)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x6246b310.
//
// Solidity: function removeGroup(string name) returns()
func (_LoadBalancer *LoadBalancerTransactor) RemoveGroup(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "removeGroup", name)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x6246b310.
//
// Solidity: function removeGroup(string name) returns()
func (_LoadBalancer *LoadBalancerSession) RemoveGroup(name string) (*types.Transaction, error) {
	return _LoadBalancer.Contract.RemoveGroup(&_LoadBalancer.TransactOpts, name)
}

// RemoveGroup is a paid mutator transaction binding the contract method 0x6246b310.
//
// Solidity: function removeGroup(string name) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) RemoveGroup(name string) (*types.Transaction, error) {
	return _LoadBalancer.Contract.RemoveGroup(&_LoadBalancer.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoadBalancer *LoadBalancerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoadBalancer *LoadBalancerSession) RenounceOwnership() (*types.Transaction, error) {
	return _LoadBalancer.Contract.RenounceOwnership(&_LoadBalancer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoadBalancer *LoadBalancerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LoadBalancer.Contract.RenounceOwnership(&_LoadBalancer.TransactOpts)
}

// SetMinSolutionsThreshold is a paid mutator transaction binding the contract method 0x69c2f156.
//
// Solidity: function setMinSolutionsThreshold(uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerTransactor) SetMinSolutionsThreshold(opts *bind.TransactOpts, minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "setMinSolutionsThreshold", minSolutionsThreshold_)
}

// SetMinSolutionsThreshold is a paid mutator transaction binding the contract method 0x69c2f156.
//
// Solidity: function setMinSolutionsThreshold(uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerSession) SetMinSolutionsThreshold(minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.SetMinSolutionsThreshold(&_LoadBalancer.TransactOpts, minSolutionsThreshold_)
}

// SetMinSolutionsThreshold is a paid mutator transaction binding the contract method 0x69c2f156.
//
// Solidity: function setMinSolutionsThreshold(uint256 minSolutionsThreshold_) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) SetMinSolutionsThreshold(minSolutionsThreshold_ *big.Int) (*types.Transaction, error) {
	return _LoadBalancer.Contract.SetMinSolutionsThreshold(&_LoadBalancer.TransactOpts, minSolutionsThreshold_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoadBalancer *LoadBalancerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoadBalancer *LoadBalancerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.TransferOwnership(&_LoadBalancer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.TransferOwnership(&_LoadBalancer.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LoadBalancer *LoadBalancerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LoadBalancer *LoadBalancerSession) Unpause() (*types.Transaction, error) {
	return _LoadBalancer.Contract.Unpause(&_LoadBalancer.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LoadBalancer *LoadBalancerTransactorSession) Unpause() (*types.Transaction, error) {
	return _LoadBalancer.Contract.Unpause(&_LoadBalancer.TransactOpts)
}

// UpdateModelCollection is a paid mutator transaction binding the contract method 0x4c255986.
//
// Solidity: function updateModelCollection(address modelCollection_) returns()
func (_LoadBalancer *LoadBalancerTransactor) UpdateModelCollection(opts *bind.TransactOpts, modelCollection_ common.Address) (*types.Transaction, error) {
	return _LoadBalancer.contract.Transact(opts, "updateModelCollection", modelCollection_)
}

// UpdateModelCollection is a paid mutator transaction binding the contract method 0x4c255986.
//
// Solidity: function updateModelCollection(address modelCollection_) returns()
func (_LoadBalancer *LoadBalancerSession) UpdateModelCollection(modelCollection_ common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.UpdateModelCollection(&_LoadBalancer.TransactOpts, modelCollection_)
}

// UpdateModelCollection is a paid mutator transaction binding the contract method 0x4c255986.
//
// Solidity: function updateModelCollection(address modelCollection_) returns()
func (_LoadBalancer *LoadBalancerTransactorSession) UpdateModelCollection(modelCollection_ common.Address) (*types.Transaction, error) {
	return _LoadBalancer.Contract.UpdateModelCollection(&_LoadBalancer.TransactOpts, modelCollection_)
}

// LoadBalancerClusterAddedIterator is returned from FilterClusterAdded and is used to iterate over the raw logs and unpacked data for ClusterAdded events raised by the LoadBalancer contract.
type LoadBalancerClusterAddedIterator struct {
	Event *LoadBalancerClusterAdded // Event containing the contract specifics and raw log

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
func (it *LoadBalancerClusterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerClusterAdded)
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
		it.Event = new(LoadBalancerClusterAdded)
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
func (it *LoadBalancerClusterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerClusterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerClusterAdded represents a ClusterAdded event raised by the LoadBalancer contract.
type LoadBalancerClusterAdded struct {
	GroupId   [32]byte
	ClusterId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClusterAdded is a free log retrieval operation binding the contract event 0x0ad827fc9c569a07a74fc468b7163fdd515381d915a1ce973559a305bad3cfa2.
//
// Solidity: event ClusterAdded(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) FilterClusterAdded(opts *bind.FilterOpts, groupId [][32]byte, clusterId []*big.Int) (*LoadBalancerClusterAddedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ClusterAdded", groupIdRule, clusterIdRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerClusterAddedIterator{contract: _LoadBalancer.contract, event: "ClusterAdded", logs: logs, sub: sub}, nil
}

// WatchClusterAdded is a free log subscription operation binding the contract event 0x0ad827fc9c569a07a74fc468b7163fdd515381d915a1ce973559a305bad3cfa2.
//
// Solidity: event ClusterAdded(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) WatchClusterAdded(opts *bind.WatchOpts, sink chan<- *LoadBalancerClusterAdded, groupId [][32]byte, clusterId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ClusterAdded", groupIdRule, clusterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerClusterAdded)
				if err := _LoadBalancer.contract.UnpackLog(event, "ClusterAdded", log); err != nil {
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

// ParseClusterAdded is a log parse operation binding the contract event 0x0ad827fc9c569a07a74fc468b7163fdd515381d915a1ce973559a305bad3cfa2.
//
// Solidity: event ClusterAdded(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) ParseClusterAdded(log types.Log) (*LoadBalancerClusterAdded, error) {
	event := new(LoadBalancerClusterAdded)
	if err := _LoadBalancer.contract.UnpackLog(event, "ClusterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerClusterGroupCreatedIterator is returned from FilterClusterGroupCreated and is used to iterate over the raw logs and unpacked data for ClusterGroupCreated events raised by the LoadBalancer contract.
type LoadBalancerClusterGroupCreatedIterator struct {
	Event *LoadBalancerClusterGroupCreated // Event containing the contract specifics and raw log

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
func (it *LoadBalancerClusterGroupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerClusterGroupCreated)
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
		it.Event = new(LoadBalancerClusterGroupCreated)
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
func (it *LoadBalancerClusterGroupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerClusterGroupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerClusterGroupCreated represents a ClusterGroupCreated event raised by the LoadBalancer contract.
type LoadBalancerClusterGroupCreated struct {
	GroupId [32]byte
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClusterGroupCreated is a free log retrieval operation binding the contract event 0xe41c83911759458108d02eee33d682c09aab4e96ff6d98687f4ed0ec44c8c5fe.
//
// Solidity: event ClusterGroupCreated(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) FilterClusterGroupCreated(opts *bind.FilterOpts, groupId [][32]byte) (*LoadBalancerClusterGroupCreatedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ClusterGroupCreated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerClusterGroupCreatedIterator{contract: _LoadBalancer.contract, event: "ClusterGroupCreated", logs: logs, sub: sub}, nil
}

// WatchClusterGroupCreated is a free log subscription operation binding the contract event 0xe41c83911759458108d02eee33d682c09aab4e96ff6d98687f4ed0ec44c8c5fe.
//
// Solidity: event ClusterGroupCreated(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) WatchClusterGroupCreated(opts *bind.WatchOpts, sink chan<- *LoadBalancerClusterGroupCreated, groupId [][32]byte) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ClusterGroupCreated", groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerClusterGroupCreated)
				if err := _LoadBalancer.contract.UnpackLog(event, "ClusterGroupCreated", log); err != nil {
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

// ParseClusterGroupCreated is a log parse operation binding the contract event 0xe41c83911759458108d02eee33d682c09aab4e96ff6d98687f4ed0ec44c8c5fe.
//
// Solidity: event ClusterGroupCreated(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) ParseClusterGroupCreated(log types.Log) (*LoadBalancerClusterGroupCreated, error) {
	event := new(LoadBalancerClusterGroupCreated)
	if err := _LoadBalancer.contract.UnpackLog(event, "ClusterGroupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerClusterGroupRemovedIterator is returned from FilterClusterGroupRemoved and is used to iterate over the raw logs and unpacked data for ClusterGroupRemoved events raised by the LoadBalancer contract.
type LoadBalancerClusterGroupRemovedIterator struct {
	Event *LoadBalancerClusterGroupRemoved // Event containing the contract specifics and raw log

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
func (it *LoadBalancerClusterGroupRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerClusterGroupRemoved)
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
		it.Event = new(LoadBalancerClusterGroupRemoved)
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
func (it *LoadBalancerClusterGroupRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerClusterGroupRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerClusterGroupRemoved represents a ClusterGroupRemoved event raised by the LoadBalancer contract.
type LoadBalancerClusterGroupRemoved struct {
	GroupId [32]byte
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClusterGroupRemoved is a free log retrieval operation binding the contract event 0xda6239e3c1c4f60a7454566117b0db7929a39167ca9907e7fae4fbe9ed8205e4.
//
// Solidity: event ClusterGroupRemoved(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) FilterClusterGroupRemoved(opts *bind.FilterOpts, groupId [][32]byte) (*LoadBalancerClusterGroupRemovedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ClusterGroupRemoved", groupIdRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerClusterGroupRemovedIterator{contract: _LoadBalancer.contract, event: "ClusterGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchClusterGroupRemoved is a free log subscription operation binding the contract event 0xda6239e3c1c4f60a7454566117b0db7929a39167ca9907e7fae4fbe9ed8205e4.
//
// Solidity: event ClusterGroupRemoved(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) WatchClusterGroupRemoved(opts *bind.WatchOpts, sink chan<- *LoadBalancerClusterGroupRemoved, groupId [][32]byte) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ClusterGroupRemoved", groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerClusterGroupRemoved)
				if err := _LoadBalancer.contract.UnpackLog(event, "ClusterGroupRemoved", log); err != nil {
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

// ParseClusterGroupRemoved is a log parse operation binding the contract event 0xda6239e3c1c4f60a7454566117b0db7929a39167ca9907e7fae4fbe9ed8205e4.
//
// Solidity: event ClusterGroupRemoved(bytes32 indexed groupId, string name)
func (_LoadBalancer *LoadBalancerFilterer) ParseClusterGroupRemoved(log types.Log) (*LoadBalancerClusterGroupRemoved, error) {
	event := new(LoadBalancerClusterGroupRemoved)
	if err := _LoadBalancer.contract.UnpackLog(event, "ClusterGroupRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerClusterRemovedIterator is returned from FilterClusterRemoved and is used to iterate over the raw logs and unpacked data for ClusterRemoved events raised by the LoadBalancer contract.
type LoadBalancerClusterRemovedIterator struct {
	Event *LoadBalancerClusterRemoved // Event containing the contract specifics and raw log

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
func (it *LoadBalancerClusterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerClusterRemoved)
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
		it.Event = new(LoadBalancerClusterRemoved)
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
func (it *LoadBalancerClusterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerClusterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerClusterRemoved represents a ClusterRemoved event raised by the LoadBalancer contract.
type LoadBalancerClusterRemoved struct {
	GroupId   [32]byte
	ClusterId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClusterRemoved is a free log retrieval operation binding the contract event 0xc264f0c7f8facadb2a37a6d959ec894d069f522bdabe18cd0cade329ea320217.
//
// Solidity: event ClusterRemoved(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) FilterClusterRemoved(opts *bind.FilterOpts, groupId [][32]byte, clusterId []*big.Int) (*LoadBalancerClusterRemovedIterator, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ClusterRemoved", groupIdRule, clusterIdRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerClusterRemovedIterator{contract: _LoadBalancer.contract, event: "ClusterRemoved", logs: logs, sub: sub}, nil
}

// WatchClusterRemoved is a free log subscription operation binding the contract event 0xc264f0c7f8facadb2a37a6d959ec894d069f522bdabe18cd0cade329ea320217.
//
// Solidity: event ClusterRemoved(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) WatchClusterRemoved(opts *bind.WatchOpts, sink chan<- *LoadBalancerClusterRemoved, groupId [][32]byte, clusterId []*big.Int) (event.Subscription, error) {

	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}
	var clusterIdRule []interface{}
	for _, clusterIdItem := range clusterId {
		clusterIdRule = append(clusterIdRule, clusterIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ClusterRemoved", groupIdRule, clusterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerClusterRemoved)
				if err := _LoadBalancer.contract.UnpackLog(event, "ClusterRemoved", log); err != nil {
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

// ParseClusterRemoved is a log parse operation binding the contract event 0xc264f0c7f8facadb2a37a6d959ec894d069f522bdabe18cd0cade329ea320217.
//
// Solidity: event ClusterRemoved(bytes32 indexed groupId, uint256 indexed clusterId)
func (_LoadBalancer *LoadBalancerFilterer) ParseClusterRemoved(log types.Log) (*LoadBalancerClusterRemoved, error) {
	event := new(LoadBalancerClusterRemoved)
	if err := _LoadBalancer.contract.UnpackLog(event, "ClusterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerInferencePerformedIterator is returned from FilterInferencePerformed and is used to iterate over the raw logs and unpacked data for InferencePerformed events raised by the LoadBalancer contract.
type LoadBalancerInferencePerformedIterator struct {
	Event *LoadBalancerInferencePerformed // Event containing the contract specifics and raw log

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
func (it *LoadBalancerInferencePerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerInferencePerformed)
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
		it.Event = new(LoadBalancerInferencePerformed)
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
func (it *LoadBalancerInferencePerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerInferencePerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerInferencePerformed represents a InferencePerformed event raised by the LoadBalancer contract.
type LoadBalancerInferencePerformed struct {
	Caller      common.Address
	InferenceId *big.Int
	GroupId     [32]byte
	ClusterId   *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferencePerformed is a free log retrieval operation binding the contract event 0xe6cc06a7b773e04156e2c577e692ac9409358de02566d0fdaff15e8604ecac56.
//
// Solidity: event InferencePerformed(address indexed caller, uint256 indexed inferenceId, bytes32 indexed groupId, uint256 clusterId, bytes data)
func (_LoadBalancer *LoadBalancerFilterer) FilterInferencePerformed(opts *bind.FilterOpts, caller []common.Address, inferenceId []*big.Int, groupId [][32]byte) (*LoadBalancerInferencePerformedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "InferencePerformed", callerRule, inferenceIdRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerInferencePerformedIterator{contract: _LoadBalancer.contract, event: "InferencePerformed", logs: logs, sub: sub}, nil
}

// WatchInferencePerformed is a free log subscription operation binding the contract event 0xe6cc06a7b773e04156e2c577e692ac9409358de02566d0fdaff15e8604ecac56.
//
// Solidity: event InferencePerformed(address indexed caller, uint256 indexed inferenceId, bytes32 indexed groupId, uint256 clusterId, bytes data)
func (_LoadBalancer *LoadBalancerFilterer) WatchInferencePerformed(opts *bind.WatchOpts, sink chan<- *LoadBalancerInferencePerformed, caller []common.Address, inferenceId []*big.Int, groupId [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var groupIdRule []interface{}
	for _, groupIdItem := range groupId {
		groupIdRule = append(groupIdRule, groupIdItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "InferencePerformed", callerRule, inferenceIdRule, groupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerInferencePerformed)
				if err := _LoadBalancer.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
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

// ParseInferencePerformed is a log parse operation binding the contract event 0xe6cc06a7b773e04156e2c577e692ac9409358de02566d0fdaff15e8604ecac56.
//
// Solidity: event InferencePerformed(address indexed caller, uint256 indexed inferenceId, bytes32 indexed groupId, uint256 clusterId, bytes data)
func (_LoadBalancer *LoadBalancerFilterer) ParseInferencePerformed(log types.Log) (*LoadBalancerInferencePerformed, error) {
	event := new(LoadBalancerInferencePerformed)
	if err := _LoadBalancer.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LoadBalancer contract.
type LoadBalancerInitializedIterator struct {
	Event *LoadBalancerInitialized // Event containing the contract specifics and raw log

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
func (it *LoadBalancerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerInitialized)
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
		it.Event = new(LoadBalancerInitialized)
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
func (it *LoadBalancerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerInitialized represents a Initialized event raised by the LoadBalancer contract.
type LoadBalancerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LoadBalancer *LoadBalancerFilterer) FilterInitialized(opts *bind.FilterOpts) (*LoadBalancerInitializedIterator, error) {

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LoadBalancerInitializedIterator{contract: _LoadBalancer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LoadBalancer *LoadBalancerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LoadBalancerInitialized) (event.Subscription, error) {

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerInitialized)
				if err := _LoadBalancer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParseInitialized(log types.Log) (*LoadBalancerInitialized, error) {
	event := new(LoadBalancerInitialized)
	if err := _LoadBalancer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerManagerAuthorizationIterator is returned from FilterManagerAuthorization and is used to iterate over the raw logs and unpacked data for ManagerAuthorization events raised by the LoadBalancer contract.
type LoadBalancerManagerAuthorizationIterator struct {
	Event *LoadBalancerManagerAuthorization // Event containing the contract specifics and raw log

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
func (it *LoadBalancerManagerAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerManagerAuthorization)
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
		it.Event = new(LoadBalancerManagerAuthorization)
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
func (it *LoadBalancerManagerAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerManagerAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerManagerAuthorization represents a ManagerAuthorization event raised by the LoadBalancer contract.
type LoadBalancerManagerAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAuthorization is a free log retrieval operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_LoadBalancer *LoadBalancerFilterer) FilterManagerAuthorization(opts *bind.FilterOpts, account []common.Address) (*LoadBalancerManagerAuthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerManagerAuthorizationIterator{contract: _LoadBalancer.contract, event: "ManagerAuthorization", logs: logs, sub: sub}, nil
}

// WatchManagerAuthorization is a free log subscription operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_LoadBalancer *LoadBalancerFilterer) WatchManagerAuthorization(opts *bind.WatchOpts, sink chan<- *LoadBalancerManagerAuthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerManagerAuthorization)
				if err := _LoadBalancer.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParseManagerAuthorization(log types.Log) (*LoadBalancerManagerAuthorization, error) {
	event := new(LoadBalancerManagerAuthorization)
	if err := _LoadBalancer.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerManagerDeauthorizationIterator is returned from FilterManagerDeauthorization and is used to iterate over the raw logs and unpacked data for ManagerDeauthorization events raised by the LoadBalancer contract.
type LoadBalancerManagerDeauthorizationIterator struct {
	Event *LoadBalancerManagerDeauthorization // Event containing the contract specifics and raw log

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
func (it *LoadBalancerManagerDeauthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerManagerDeauthorization)
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
		it.Event = new(LoadBalancerManagerDeauthorization)
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
func (it *LoadBalancerManagerDeauthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerManagerDeauthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerManagerDeauthorization represents a ManagerDeauthorization event raised by the LoadBalancer contract.
type LoadBalancerManagerDeauthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerDeauthorization is a free log retrieval operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_LoadBalancer *LoadBalancerFilterer) FilterManagerDeauthorization(opts *bind.FilterOpts, account []common.Address) (*LoadBalancerManagerDeauthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerManagerDeauthorizationIterator{contract: _LoadBalancer.contract, event: "ManagerDeauthorization", logs: logs, sub: sub}, nil
}

// WatchManagerDeauthorization is a free log subscription operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_LoadBalancer *LoadBalancerFilterer) WatchManagerDeauthorization(opts *bind.WatchOpts, sink chan<- *LoadBalancerManagerDeauthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerManagerDeauthorization)
				if err := _LoadBalancer.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParseManagerDeauthorization(log types.Log) (*LoadBalancerManagerDeauthorization, error) {
	event := new(LoadBalancerManagerDeauthorization)
	if err := _LoadBalancer.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoadBalancer contract.
type LoadBalancerOwnershipTransferredIterator struct {
	Event *LoadBalancerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoadBalancerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerOwnershipTransferred)
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
		it.Event = new(LoadBalancerOwnershipTransferred)
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
func (it *LoadBalancerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerOwnershipTransferred represents a OwnershipTransferred event raised by the LoadBalancer contract.
type LoadBalancerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoadBalancer *LoadBalancerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoadBalancerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoadBalancerOwnershipTransferredIterator{contract: _LoadBalancer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoadBalancer *LoadBalancerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoadBalancerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerOwnershipTransferred)
				if err := _LoadBalancer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParseOwnershipTransferred(log types.Log) (*LoadBalancerOwnershipTransferred, error) {
	event := new(LoadBalancerOwnershipTransferred)
	if err := _LoadBalancer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LoadBalancer contract.
type LoadBalancerPausedIterator struct {
	Event *LoadBalancerPaused // Event containing the contract specifics and raw log

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
func (it *LoadBalancerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerPaused)
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
		it.Event = new(LoadBalancerPaused)
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
func (it *LoadBalancerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerPaused represents a Paused event raised by the LoadBalancer contract.
type LoadBalancerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LoadBalancer *LoadBalancerFilterer) FilterPaused(opts *bind.FilterOpts) (*LoadBalancerPausedIterator, error) {

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LoadBalancerPausedIterator{contract: _LoadBalancer.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LoadBalancer *LoadBalancerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LoadBalancerPaused) (event.Subscription, error) {

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerPaused)
				if err := _LoadBalancer.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParsePaused(log types.Log) (*LoadBalancerPaused, error) {
	event := new(LoadBalancerPaused)
	if err := _LoadBalancer.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoadBalancerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LoadBalancer contract.
type LoadBalancerUnpausedIterator struct {
	Event *LoadBalancerUnpaused // Event containing the contract specifics and raw log

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
func (it *LoadBalancerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadBalancerUnpaused)
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
		it.Event = new(LoadBalancerUnpaused)
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
func (it *LoadBalancerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadBalancerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadBalancerUnpaused represents a Unpaused event raised by the LoadBalancer contract.
type LoadBalancerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LoadBalancer *LoadBalancerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LoadBalancerUnpausedIterator, error) {

	logs, sub, err := _LoadBalancer.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LoadBalancerUnpausedIterator{contract: _LoadBalancer.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LoadBalancer *LoadBalancerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LoadBalancerUnpaused) (event.Subscription, error) {

	logs, sub, err := _LoadBalancer.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadBalancerUnpaused)
				if err := _LoadBalancer.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LoadBalancer *LoadBalancerFilterer) ParseUnpaused(log types.Log) (*LoadBalancerUnpaused, error) {
	event := new(LoadBalancerUnpaused)
	if err := _LoadBalancer.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
