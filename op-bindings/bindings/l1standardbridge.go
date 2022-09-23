// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// L1StandardBridgeMetaData contains all meta data concerning the L1StandardBridge contract.
var L1StandardBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHWithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeETHWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002c6938038062002c6983398101604081905262000035916200006f565b6001600160a01b031660805273420000000000000000000000000000000000001060a052600060c081905260e052600261010052620000a1565b6000602082840312156200008257600080fd5b81516001600160a01b03811681146200009a57600080fd5b9392505050565b60805160a05160c05160e05161010051612b0c6200015d6000396000611202015260006111d9015260006111b001526000818161037f015281816103fb0152818161058b015281816106d701528181610be201528181610e2f015281816113ef0152611b8f01526000818161023c0152818161054e015281816106ad0152818161070e01528181610bb801528181610c1901528181610e0501528181610e66015281816113c5015281816114260152611b530152612b0c6000f3fe6080604052600436106101125760003560e01c8063838b2520116100a55780639a2ac6d511610074578063b1a1a88211610059578063b1a1a882146103d6578063c89701a2146103e9578063e11013dd1461041d57600080fd5b80639a2ac6d5146103a3578063a9f9e675146103b657600080fd5b8063838b2520146102ea578063870876231461030a5780638f601f661461032a57806391c49bf81461037057600080fd5b80633cb747bf116100e15780633cb747bf1461022a578063540abf731461028857806354fd4d50146102a857806358a997f6146102ca57600080fd5b80630166a07a146101d157806309fc8843146101f15780631532ec34146102045780631635f5fd1461021757600080fd5b366101cc57333b156101ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101ca33333462030d4060405180602001604052806000815250610430565b005b600080fd5b3480156101dd57600080fd5b506101ca6101ec3660046123a7565b610695565b6101ca6101ff366004612458565b610ac9565b6101ca6102123660046124ab565b610ba0565b6101ca6102253660046124ab565b610ded565b34801561023657600080fd5b5061025e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561029457600080fd5b506101ca6102a336600461251e565b611190565b3480156102b457600080fd5b506102bd6111a9565b60405161027f919061260b565b3480156102d657600080fd5b506101ca6102e536600461261e565b61124c565b3480156102f657600080fd5b506101ca61030536600461251e565b6112eb565b34801561031657600080fd5b506101ca61032536600461261e565b6112fb565b34801561033657600080fd5b506103626103453660046126a1565b600260209081526000928352604080842090915290825290205481565b60405190815260200161027f565b34801561037c57600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061025e565b6101ca6103b13660046126da565b61139a565b3480156103c257600080fd5b506101ca6103d13660046123a7565b6113ad565b6101ca6103e4366004612458565b61160e565b3480156103f557600080fd5b5061025e7f000000000000000000000000000000000000000000000000000000000000000081565b6101ca61042b3660046126da565b6116aa565b8234146104e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f5374616e646172644272696467653a20616d6f756e742073656e74207769746860448201527f2063616c6c20646f6573206e6f74206d6174636820616d6f756e74206265696e60648201527f6720627269646765640000000000000000000000000000000000000000000000608482015260a4016101a2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5858460405161054492919061273d565b60405180910390a37f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b847f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b898989886040516024016105c99493929190612756565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261065c9291889060040161279f565b6000604051808303818588803b15801561067557600080fd5b505af1158015610689573d6000803e3d6000fd5b50505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107b357507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610777573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079b91906127e4565b73ffffffffffffffffffffffffffffffffffffffff16145b610865576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b61086e876116ed565b156109bc5761087d878761174f565b61092f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561099f57600080fd5b505af11580156109b3573d6000803e3d6000fd5b50505050610a3e565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546109fa908490612830565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c1683529390529190912091909155610a3e9085856117f6565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd87878787604051610ab89493929190612890565b60405180910390a450505050505050565b333b15610b58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b610b9b3333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061043092505050565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610cbe57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca691906127e4565b73ffffffffffffffffffffffffffffffffffffffff16145b610d70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631858585604051610dd1939291906128c6565b60405180910390a3610de68585858585610ded565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610f0b57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ecf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef391906127e4565b73ffffffffffffffffffffffffffffffffffffffff16145b610fbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b823414611073576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f5374616e646172644272696467653a20616d6f756e742073656e742077697468908201527f2063616c6c20646f6573206e6f74206d6174636820616d6f756e74207265717560648201527f6972656400000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8585856040516110d4939291906128c6565b60405180910390a360006110f9855a86604051806020016040528060008152506118ca565b905080611188576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101a2565b505050505050565b6111a087873388888888886118e4565b50505050505050565b60606111d47f0000000000000000000000000000000000000000000000000000000000000000611ca2565b6111fd7f0000000000000000000000000000000000000000000000000000000000000000611ca2565b6112267f0000000000000000000000000000000000000000000000000000000000000000611ca2565b604051602001611238939291906128e9565b604051602081830303815290604052905090565b333b156112db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b6111888686333388888888611ddf565b6111a08787338888888888611ddf565b333b1561138a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b61118886863333888888886118e4565b6113a73385858585611e7b565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156114cb57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561148f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b391906127e4565b73ffffffffffffffffffffffffffffffffffffffff16145b61157d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3878787876040516115f79493929190612890565b60405180910390a46111a087878787878787610695565b333b1561169d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b610b9b3333858585611e7b565b6113a73385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061043092505050565b6000611719827f1d1d8b6300000000000000000000000000000000000000000000000000000000611f27565b806117495750611749827fec4fc8e300000000000000000000000000000000000000000000000000000000611f27565b92915050565b60008273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa15801561179c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c091906127e4565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905092915050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610b9b9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611f4a565b600080600080845160208601878a8af19695505050505050565b6118ed886116ed565b15611a3b576118fc888861174f565b6119ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201869052891690639dc29fac90604401600060405180830381600087803b158015611a1e57600080fd5b505af1158015611a32573d6000803e3d6000fd5b50505050611acf565b611a5d73ffffffffffffffffffffffffffffffffffffffff8916873087612056565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b1683529290522054611a9b90859061295f565b73ffffffffffffffffffffffffffffffffffffffff808a166000908152600260209081526040808320938c16835292905220555b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf88888787604051611b499493929190612890565b60405180910390a47f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633dbb202b7f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b8a8c8b8b8b8a8a604051602401611bd39796959493929190612977565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611c669291889060040161279f565b600060405180830381600087803b158015611c8057600080fd5b505af1158015611c94573d6000803e3d6000fd5b505050505050505050505050565b606081600003611ce557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611d0f5780611cf9816129d4565b9150611d089050600a83612a3b565b9150611ce9565b60008167ffffffffffffffff811115611d2a57611d2a612a4f565b6040519080825280601f01601f191660200182016040528015611d54576020820181803683370190505b5090505b8415611dd757611d69600183612830565b9150611d76600a86612a7e565b611d8190603061295f565b60f81b818381518110611d9657611d96612a92565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611dd0600a86612a3b565b9450611d58565b949350505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039688888787604051611e599493929190612890565b60405180910390a4611e7188888888888888886118e4565b5050505050505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23348585604051611edc939291906128c6565b60405180910390a3610de68585348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061043092505050565b6000611f32836120b4565b8015611f435750611f438383612118565b9392505050565b6000611fac826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166121e79092919063ffffffff16565b805190915015610b9b5780806020019051810190611fca9190612ac1565b610b9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101a2565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113a79085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611848565b60006120e0827f01ffc9a700000000000000000000000000000000000000000000000000000000612118565b80156117495750612111827fffffffff00000000000000000000000000000000000000000000000000000000612118565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156121d0575060208210155b80156121dc5750600081115b979650505050505050565b6060611dd784846000858573ffffffffffffffffffffffffffffffffffffffff85163b612270576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101a2565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516122999190612ae3565b60006040518083038185875af1925050503d80600081146122d6576040519150601f19603f3d011682016040523d82523d6000602084013e6122db565b606091505b50915091506121dc828286606083156122f5575081611f43565b8251156123055782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a2919061260b565b73ffffffffffffffffffffffffffffffffffffffff8116811461235b57600080fd5b50565b60008083601f84011261237057600080fd5b50813567ffffffffffffffff81111561238857600080fd5b6020830191508360208285010111156123a057600080fd5b9250929050565b600080600080600080600060c0888a0312156123c257600080fd5b87356123cd81612339565b965060208801356123dd81612339565b955060408801356123ed81612339565b945060608801356123fd81612339565b93506080880135925060a088013567ffffffffffffffff81111561242057600080fd5b61242c8a828b0161235e565b989b979a50959850939692959293505050565b803563ffffffff8116811461245357600080fd5b919050565b60008060006040848603121561246d57600080fd5b6124768461243f565b9250602084013567ffffffffffffffff81111561249257600080fd5b61249e8682870161235e565b9497909650939450505050565b6000806000806000608086880312156124c357600080fd5b85356124ce81612339565b945060208601356124de81612339565b935060408601359250606086013567ffffffffffffffff81111561250157600080fd5b61250d8882890161235e565b969995985093965092949392505050565b600080600080600080600060c0888a03121561253957600080fd5b873561254481612339565b9650602088013561255481612339565b9550604088013561256481612339565b9450606088013593506125796080890161243f565b925060a088013567ffffffffffffffff81111561242057600080fd5b60005b838110156125b0578181015183820152602001612598565b838111156113a75750506000910152565b600081518084526125d9816020860160208601612595565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611f4360208301846125c1565b60008060008060008060a0878903121561263757600080fd5b863561264281612339565b9550602087013561265281612339565b9450604087013593506126676060880161243f565b9250608087013567ffffffffffffffff81111561268357600080fd5b61268f89828a0161235e565b979a9699509497509295939492505050565b600080604083850312156126b457600080fd5b82356126bf81612339565b915060208301356126cf81612339565b809150509250929050565b600080600080606085870312156126f057600080fd5b84356126fb81612339565b93506127096020860161243f565b9250604085013567ffffffffffffffff81111561272557600080fd5b6127318782880161235e565b95989497509550505050565b828152604060208201526000611dd760408301846125c1565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261279560808301846125c1565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006127ce60608301856125c1565b905063ffffffff83166040830152949350505050565b6000602082840312156127f657600080fd5b8151611f4381612339565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561284257612842612801565b500390565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152606060408201526000612795606083018486612847565b8381526040602082015260006128e0604083018486612847565b95945050505050565b600084516128fb818460208901612595565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612937816001850160208a01612595565b60019201918201528351612952816002840160208801612595565b0160020195945050505050565b6000821982111561297257612972612801565b500190565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526129c760c083018486612847565b9998505050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612a0557612a05612801565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612a4a57612a4a612a0c565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082612a8d57612a8d612a0c565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215612ad357600080fd5b81518015158114611f4357600080fd5b60008251612af5818460208701612595565b919091019291505056fea164736f6c634300080f000a",
}

// L1StandardBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1StandardBridgeMetaData.ABI instead.
var L1StandardBridgeABI = L1StandardBridgeMetaData.ABI

// L1StandardBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1StandardBridgeMetaData.Bin instead.
var L1StandardBridgeBin = L1StandardBridgeMetaData.Bin

// DeployL1StandardBridge deploys a new Ethereum contract, binding an instance of L1StandardBridge to it.
func DeployL1StandardBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _messenger common.Address) (common.Address, *types.Transaction, *L1StandardBridge, error) {
	parsed, err := L1StandardBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1StandardBridgeBin), backend, _messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1StandardBridge{L1StandardBridgeCaller: L1StandardBridgeCaller{contract: contract}, L1StandardBridgeTransactor: L1StandardBridgeTransactor{contract: contract}, L1StandardBridgeFilterer: L1StandardBridgeFilterer{contract: contract}}, nil
}

// L1StandardBridge is an auto generated Go binding around an Ethereum contract.
type L1StandardBridge struct {
	L1StandardBridgeCaller     // Read-only binding to the contract
	L1StandardBridgeTransactor // Write-only binding to the contract
	L1StandardBridgeFilterer   // Log filterer for contract events
}

// L1StandardBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1StandardBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1StandardBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1StandardBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1StandardBridgeSession struct {
	Contract     *L1StandardBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1StandardBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1StandardBridgeCallerSession struct {
	Contract *L1StandardBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L1StandardBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1StandardBridgeTransactorSession struct {
	Contract     *L1StandardBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1StandardBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1StandardBridgeRaw struct {
	Contract *L1StandardBridge // Generic contract binding to access the raw methods on
}

// L1StandardBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1StandardBridgeCallerRaw struct {
	Contract *L1StandardBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1StandardBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1StandardBridgeTransactorRaw struct {
	Contract *L1StandardBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1StandardBridge creates a new instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridge(address common.Address, backend bind.ContractBackend) (*L1StandardBridge, error) {
	contract, err := bindL1StandardBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridge{L1StandardBridgeCaller: L1StandardBridgeCaller{contract: contract}, L1StandardBridgeTransactor: L1StandardBridgeTransactor{contract: contract}, L1StandardBridgeFilterer: L1StandardBridgeFilterer{contract: contract}}, nil
}

// NewL1StandardBridgeCaller creates a new read-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeCaller(address common.Address, caller bind.ContractCaller) (*L1StandardBridgeCaller, error) {
	contract, err := bindL1StandardBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeCaller{contract: contract}, nil
}

// NewL1StandardBridgeTransactor creates a new write-only instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1StandardBridgeTransactor, error) {
	contract, err := bindL1StandardBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeTransactor{contract: contract}, nil
}

// NewL1StandardBridgeFilterer creates a new log filterer instance of L1StandardBridge, bound to a specific deployed contract.
func NewL1StandardBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1StandardBridgeFilterer, error) {
	contract, err := bindL1StandardBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeFilterer{contract: contract}, nil
}

// bindL1StandardBridge binds a generic wrapper to an already deployed contract.
func bindL1StandardBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1StandardBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.L1StandardBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.L1StandardBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1StandardBridge *L1StandardBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1StandardBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1StandardBridge *L1StandardBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _L1StandardBridge.Contract.Deposits(&_L1StandardBridge.CallOpts, arg0, arg1)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) L2TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "l2TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) L2TokenBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.L2TokenBridge(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Messenger() (common.Address, error) {
	return _L1StandardBridge.Contract.Messenger(&_L1StandardBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeSession) OtherBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.OtherBridge(&_L1StandardBridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L1StandardBridge *L1StandardBridgeCallerSession) OtherBridge() (common.Address, error) {
	return _L1StandardBridge.Contract.OtherBridge(&_L1StandardBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1StandardBridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeSession) Version() (string, error) {
	return _L1StandardBridge.Contract.Version(&_L1StandardBridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L1StandardBridge *L1StandardBridgeCallerSession) Version() (string, error) {
	return _L1StandardBridge.Contract.Version(&_L1StandardBridge.CallOpts)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeERC20", _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeERC20To", _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20To(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeERC20To(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeETH", _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeETH(&_L1StandardBridge.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeETH(&_L1StandardBridge.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) BridgeETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "bridgeETHTo", _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeETHTo(&_L1StandardBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.BridgeETHTo(&_L1StandardBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositERC20To(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositETH", _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETH(&_L1StandardBridge.TransactOpts, _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETH(&_L1StandardBridge.TransactOpts, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) DepositETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "depositETHTo", _to, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) DepositETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETHTo(&_L1StandardBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) DepositETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.DepositETHTo(&_L1StandardBridge.TransactOpts, _to, _minGasLimit, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeBridgeERC20", _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeERC20(&_L1StandardBridge.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeBridgeETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeBridgeETH", _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeETH(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeBridgeETH(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeERC20Withdrawal(&_L1StandardBridge.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) FinalizeETHWithdrawal(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.contract.Transact(opts, "finalizeETHWithdrawal", _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeETHWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _L1StandardBridge.Contract.FinalizeETHWithdrawal(&_L1StandardBridge.TransactOpts, _from, _to, _amount, _extraData)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1StandardBridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeSession) Receive() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Receive(&_L1StandardBridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1StandardBridge *L1StandardBridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _L1StandardBridge.Contract.Receive(&_L1StandardBridge.TransactOpts)
}

// L1StandardBridgeERC20BridgeFinalizedIterator is returned from FilterERC20BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC20BridgeFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeFinalizedIterator struct {
	Event *L1StandardBridgeERC20BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20BridgeFinalized)
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
		it.Event = new(L1StandardBridgeERC20BridgeFinalized)
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
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeFinalized is a free log retrieval operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1StandardBridgeERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20BridgeFinalizedIterator{contract: _L1StandardBridge.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeFinalized is a free log subscription operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20BridgeFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

// ParseERC20BridgeFinalized is a log parse operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20BridgeFinalized(log types.Log) (*L1StandardBridgeERC20BridgeFinalized, error) {
	event := new(L1StandardBridgeERC20BridgeFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20BridgeInitiatedIterator is returned from FilterERC20BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC20BridgeInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeInitiatedIterator struct {
	Event *L1StandardBridgeERC20BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20BridgeInitiated)
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
		it.Event = new(L1StandardBridgeERC20BridgeInitiated)
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
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeInitiated is a free log retrieval operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*L1StandardBridgeERC20BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20BridgeInitiatedIterator{contract: _L1StandardBridge.contract, event: "ERC20BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeInitiated is a free log subscription operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20BridgeInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
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

// ParseERC20BridgeInitiated is a log parse operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20BridgeInitiated(log types.Log) (*L1StandardBridgeERC20BridgeInitiated, error) {
	event := new(L1StandardBridgeERC20BridgeInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiatedIterator struct {
	Event *L1StandardBridgeERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
		it.Event = new(L1StandardBridgeERC20DepositInitiated)
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
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20DepositInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardBridgeERC20DepositInitiatedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20DepositInitiatedIterator{contract: _L1StandardBridge.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20DepositInitiated, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20DepositInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20DepositInitiated(log types.Log) (*L1StandardBridgeERC20DepositInitiated, error) {
	event := new(L1StandardBridgeERC20DepositInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalizedIterator struct {
	Event *L1StandardBridgeERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
		it.Event = new(L1StandardBridgeERC20WithdrawalFinalized)
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
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeERC20WithdrawalFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L1StandardBridgeERC20WithdrawalFinalizedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeERC20WithdrawalFinalizedIterator{contract: _L1StandardBridge.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeERC20WithdrawalFinalized, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeERC20WithdrawalFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*L1StandardBridgeERC20WithdrawalFinalized, error) {
	event := new(L1StandardBridgeERC20WithdrawalFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHBridgeFinalizedIterator is returned from FilterETHBridgeFinalized and is used to iterate over the raw logs and unpacked data for ETHBridgeFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeETHBridgeFinalizedIterator struct {
	Event *L1StandardBridgeETHBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHBridgeFinalized)
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
		it.Event = new(L1StandardBridgeETHBridgeFinalized)
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
func (it *L1StandardBridgeETHBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHBridgeFinalized represents a ETHBridgeFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeFinalized is a free log retrieval operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHBridgeFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1StandardBridgeETHBridgeFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHBridgeFinalizedIterator{contract: _L1StandardBridge.contract, event: "ETHBridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchETHBridgeFinalized is a free log subscription operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHBridgeFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHBridgeFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHBridgeFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
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

// ParseETHBridgeFinalized is a log parse operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHBridgeFinalized(log types.Log) (*L1StandardBridgeETHBridgeFinalized, error) {
	event := new(L1StandardBridgeETHBridgeFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHBridgeInitiatedIterator is returned from FilterETHBridgeInitiated and is used to iterate over the raw logs and unpacked data for ETHBridgeInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeETHBridgeInitiatedIterator struct {
	Event *L1StandardBridgeETHBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHBridgeInitiated)
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
		it.Event = new(L1StandardBridgeETHBridgeInitiated)
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
func (it *L1StandardBridgeETHBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHBridgeInitiated represents a ETHBridgeInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeInitiated is a free log retrieval operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHBridgeInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1StandardBridgeETHBridgeInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHBridgeInitiatedIterator{contract: _L1StandardBridge.contract, event: "ETHBridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchETHBridgeInitiated is a free log subscription operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHBridgeInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHBridgeInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHBridgeInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
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

// ParseETHBridgeInitiated is a log parse operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHBridgeInitiated(log types.Log) (*L1StandardBridgeETHBridgeInitiated, error) {
	event := new(L1StandardBridgeETHBridgeInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHDepositInitiatedIterator is returned from FilterETHDepositInitiated and is used to iterate over the raw logs and unpacked data for ETHDepositInitiated events raised by the L1StandardBridge contract.
type L1StandardBridgeETHDepositInitiatedIterator struct {
	Event *L1StandardBridgeETHDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHDepositInitiated)
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
		it.Event = new(L1StandardBridgeETHDepositInitiated)
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
func (it *L1StandardBridgeETHDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHDepositInitiated represents a ETHDepositInitiated event raised by the L1StandardBridge contract.
type L1StandardBridgeETHDepositInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHDepositInitiated is a free log retrieval operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHDepositInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1StandardBridgeETHDepositInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHDepositInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHDepositInitiatedIterator{contract: _L1StandardBridge.contract, event: "ETHDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchETHDepositInitiated is a free log subscription operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHDepositInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHDepositInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHDepositInitiated)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
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

// ParseETHDepositInitiated is a log parse operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHDepositInitiated(log types.Log) (*L1StandardBridgeETHDepositInitiated, error) {
	event := new(L1StandardBridgeETHDepositInitiated)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1StandardBridgeETHWithdrawalFinalizedIterator is returned from FilterETHWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ETHWithdrawalFinalized events raised by the L1StandardBridge contract.
type L1StandardBridgeETHWithdrawalFinalizedIterator struct {
	Event *L1StandardBridgeETHWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1StandardBridgeETHWithdrawalFinalized)
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
		it.Event = new(L1StandardBridgeETHWithdrawalFinalized)
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
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1StandardBridgeETHWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1StandardBridgeETHWithdrawalFinalized represents a ETHWithdrawalFinalized event raised by the L1StandardBridge contract.
type L1StandardBridgeETHWithdrawalFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHWithdrawalFinalized is a free log retrieval operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) FilterETHWithdrawalFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1StandardBridgeETHWithdrawalFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.FilterLogs(opts, "ETHWithdrawalFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1StandardBridgeETHWithdrawalFinalizedIterator{contract: _L1StandardBridge.contract, event: "ETHWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchETHWithdrawalFinalized is a free log subscription operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) WatchETHWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1StandardBridgeETHWithdrawalFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1StandardBridge.contract.WatchLogs(opts, "ETHWithdrawalFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1StandardBridgeETHWithdrawalFinalized)
				if err := _L1StandardBridge.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
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

// ParseETHWithdrawalFinalized is a log parse operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_L1StandardBridge *L1StandardBridgeFilterer) ParseETHWithdrawalFinalized(log types.Log) (*L1StandardBridgeETHWithdrawalFinalized, error) {
	event := new(L1StandardBridgeETHWithdrawalFinalized)
	if err := _L1StandardBridge.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
