// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package codexam

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICodexAMABI is the input ABI used to generate the binding from.
const ICodexAMABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_company\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_website\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_socialAccounts\",\"type\":\"string[]\"}],\"name\":\"editAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_targetAddr\",\"type\":\"address\"}],\"name\":\"follow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_company\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_website\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_socialAccounts\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_followingNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_followerNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getAccountName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"getBatchAccountById\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"getBatchAccountName\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"_names\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getFollower\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getFollowing\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avatar\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_bio\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_company\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_website\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_socialAccounts\",\"type\":\"string[]\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_targetAddr\",\"type\":\"address\"}],\"name\":\"unfollow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}],"

// ICodexAM is an auto generated Go binding around an Ethereum contract.
type ICodexAM struct {
	ICodexAMCaller     // Read-only binding to the contract
	ICodexAMTransactor // Write-only binding to the contract
	ICodexAMFilterer   // Log filterer for contract events
}

// ICodexAMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICodexAMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICodexAMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICodexAMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICodexAMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICodexAMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICodexAMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICodexAMSession struct {
	Contract     *ICodexAM         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICodexAMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICodexAMCallerSession struct {
	Contract *ICodexAMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ICodexAMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICodexAMTransactorSession struct {
	Contract     *ICodexAMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICodexAMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICodexAMRaw struct {
	Contract *ICodexAM // Generic contract binding to access the raw methods on
}

// ICodexAMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICodexAMCallerRaw struct {
	Contract *ICodexAMCaller // Generic read-only contract binding to access the raw methods on
}

// ICodexAMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICodexAMTransactorRaw struct {
	Contract *ICodexAMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICodexAM creates a new instance of ICodexAM, bound to a specific deployed contract.
func NewICodexAM(address common.Address, backend bind.ContractBackend) (*ICodexAM, error) {
	contract, err := bindICodexAM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICodexAM{ICodexAMCaller: ICodexAMCaller{contract: contract}, ICodexAMTransactor: ICodexAMTransactor{contract: contract}, ICodexAMFilterer: ICodexAMFilterer{contract: contract}}, nil
}

// NewICodexAMCaller creates a new read-only instance of ICodexAM, bound to a specific deployed contract.
func NewICodexAMCaller(address common.Address, caller bind.ContractCaller) (*ICodexAMCaller, error) {
	contract, err := bindICodexAM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICodexAMCaller{contract: contract}, nil
}

// NewICodexAMTransactor creates a new write-only instance of ICodexAM, bound to a specific deployed contract.
func NewICodexAMTransactor(address common.Address, transactor bind.ContractTransactor) (*ICodexAMTransactor, error) {
	contract, err := bindICodexAM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICodexAMTransactor{contract: contract}, nil
}

// NewICodexAMFilterer creates a new log filterer instance of ICodexAM, bound to a specific deployed contract.
func NewICodexAMFilterer(address common.Address, filterer bind.ContractFilterer) (*ICodexAMFilterer, error) {
	contract, err := bindICodexAM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICodexAMFilterer{contract: contract}, nil
}

// bindICodexAM binds a generic wrapper to an already deployed contract.
func bindICodexAM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICodexAMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICodexAM *ICodexAMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICodexAM.Contract.ICodexAMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICodexAM *ICodexAMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICodexAM.Contract.ICodexAMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICodexAM *ICodexAMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICodexAM.Contract.ICodexAMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICodexAM *ICodexAMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICodexAM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICodexAM *ICodexAMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICodexAM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICodexAM *ICodexAMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICodexAM.Contract.contract.Transact(opts, method, params...)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(uint256 _id, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts, uint256 _followingNumber, uint256 _followerNumber)
func (_ICodexAM *ICodexAMCaller) GetAccountDetails(opts *bind.CallOpts, _account common.Address) (struct {
	Id              *big.Int
	Name            string
	Avatar          string
	Bio             string
	Company         string
	Location        string
	Website         string
	SocialAccounts  []string
	FollowingNumber *big.Int
	FollowerNumber  *big.Int
}, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getAccountDetails", _account)

	outstruct := new(struct {
		Id              *big.Int
		Name            string
		Avatar          string
		Bio             string
		Company         string
		Location        string
		Website         string
		SocialAccounts  []string
		FollowingNumber *big.Int
		FollowerNumber  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Avatar = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Bio = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Company = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Location = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Website = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.SocialAccounts = *abi.ConvertType(out[7], new([]string)).(*[]string)
	outstruct.FollowingNumber = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FollowerNumber = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(uint256 _id, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts, uint256 _followingNumber, uint256 _followerNumber)
func (_ICodexAM *ICodexAMSession) GetAccountDetails(_account common.Address) (struct {
	Id              *big.Int
	Name            string
	Avatar          string
	Bio             string
	Company         string
	Location        string
	Website         string
	SocialAccounts  []string
	FollowingNumber *big.Int
	FollowerNumber  *big.Int
}, error) {
	return _ICodexAM.Contract.GetAccountDetails(&_ICodexAM.CallOpts, _account)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(address _account) view returns(uint256 _id, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts, uint256 _followingNumber, uint256 _followerNumber)
func (_ICodexAM *ICodexAMCallerSession) GetAccountDetails(_account common.Address) (struct {
	Id              *big.Int
	Name            string
	Avatar          string
	Bio             string
	Company         string
	Location        string
	Website         string
	SocialAccounts  []string
	FollowingNumber *big.Int
	FollowerNumber  *big.Int
}, error) {
	return _ICodexAM.Contract.GetAccountDetails(&_ICodexAM.CallOpts, _account)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address _account) view returns(uint256)
func (_ICodexAM *ICodexAMCaller) GetAccountId(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getAccountId", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address _account) view returns(uint256)
func (_ICodexAM *ICodexAMSession) GetAccountId(_account common.Address) (*big.Int, error) {
	return _ICodexAM.Contract.GetAccountId(&_ICodexAM.CallOpts, _account)
}

// GetAccountId is a free data retrieval call binding the contract method 0xe0b490f7.
//
// Solidity: function getAccountId(address _account) view returns(uint256)
func (_ICodexAM *ICodexAMCallerSession) GetAccountId(_account common.Address) (*big.Int, error) {
	return _ICodexAM.Contract.GetAccountId(&_ICodexAM.CallOpts, _account)
}

// GetAccountName is a free data retrieval call binding the contract method 0xd989c8db.
//
// Solidity: function getAccountName(address _account) view returns(string)
func (_ICodexAM *ICodexAMCaller) GetAccountName(opts *bind.CallOpts, _account common.Address) (string, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getAccountName", _account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAccountName is a free data retrieval call binding the contract method 0xd989c8db.
//
// Solidity: function getAccountName(address _account) view returns(string)
func (_ICodexAM *ICodexAMSession) GetAccountName(_account common.Address) (string, error) {
	return _ICodexAM.Contract.GetAccountName(&_ICodexAM.CallOpts, _account)
}

// GetAccountName is a free data retrieval call binding the contract method 0xd989c8db.
//
// Solidity: function getAccountName(address _account) view returns(string)
func (_ICodexAM *ICodexAMCallerSession) GetAccountName(_account common.Address) (string, error) {
	return _ICodexAM.Contract.GetAccountName(&_ICodexAM.CallOpts, _account)
}

// GetBatchAccountById is a free data retrieval call binding the contract method 0x81f02589.
//
// Solidity: function getBatchAccountById(uint256[] _ids) view returns(address[] _accounts)
func (_ICodexAM *ICodexAMCaller) GetBatchAccountById(opts *bind.CallOpts, _ids []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getBatchAccountById", _ids)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchAccountById is a free data retrieval call binding the contract method 0x81f02589.
//
// Solidity: function getBatchAccountById(uint256[] _ids) view returns(address[] _accounts)
func (_ICodexAM *ICodexAMSession) GetBatchAccountById(_ids []*big.Int) ([]common.Address, error) {
	return _ICodexAM.Contract.GetBatchAccountById(&_ICodexAM.CallOpts, _ids)
}

// GetBatchAccountById is a free data retrieval call binding the contract method 0x81f02589.
//
// Solidity: function getBatchAccountById(uint256[] _ids) view returns(address[] _accounts)
func (_ICodexAM *ICodexAMCallerSession) GetBatchAccountById(_ids []*big.Int) ([]common.Address, error) {
	return _ICodexAM.Contract.GetBatchAccountById(&_ICodexAM.CallOpts, _ids)
}

// GetBatchAccountName is a free data retrieval call binding the contract method 0xaba48a7a.
//
// Solidity: function getBatchAccountName(address[] _accounts) view returns(string[] _names)
func (_ICodexAM *ICodexAMCaller) GetBatchAccountName(opts *bind.CallOpts, _accounts []common.Address) ([]string, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getBatchAccountName", _accounts)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetBatchAccountName is a free data retrieval call binding the contract method 0xaba48a7a.
//
// Solidity: function getBatchAccountName(address[] _accounts) view returns(string[] _names)
func (_ICodexAM *ICodexAMSession) GetBatchAccountName(_accounts []common.Address) ([]string, error) {
	return _ICodexAM.Contract.GetBatchAccountName(&_ICodexAM.CallOpts, _accounts)
}

// GetBatchAccountName is a free data retrieval call binding the contract method 0xaba48a7a.
//
// Solidity: function getBatchAccountName(address[] _accounts) view returns(string[] _names)
func (_ICodexAM *ICodexAMCallerSession) GetBatchAccountName(_accounts []common.Address) ([]string, error) {
	return _ICodexAM.Contract.GetBatchAccountName(&_ICodexAM.CallOpts, _accounts)
}

// GetFollower is a free data retrieval call binding the contract method 0xfb97abcc.
//
// Solidity: function getFollower(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMCaller) GetFollower(opts *bind.CallOpts, _account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getFollower", _account, offset, limit)

	outstruct := new(struct {
		Ids         []*big.Int
		TotalLength *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLength = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFollower is a free data retrieval call binding the contract method 0xfb97abcc.
//
// Solidity: function getFollower(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMSession) GetFollower(_account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	return _ICodexAM.Contract.GetFollower(&_ICodexAM.CallOpts, _account, offset, limit)
}

// GetFollower is a free data retrieval call binding the contract method 0xfb97abcc.
//
// Solidity: function getFollower(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMCallerSession) GetFollower(_account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	return _ICodexAM.Contract.GetFollower(&_ICodexAM.CallOpts, _account, offset, limit)
}

// GetFollowing is a free data retrieval call binding the contract method 0x7475676b.
//
// Solidity: function getFollowing(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMCaller) GetFollowing(opts *bind.CallOpts, _account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	var out []interface{}
	err := _ICodexAM.contract.Call(opts, &out, "getFollowing", _account, offset, limit)

	outstruct := new(struct {
		Ids         []*big.Int
		TotalLength *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalLength = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFollowing is a free data retrieval call binding the contract method 0x7475676b.
//
// Solidity: function getFollowing(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMSession) GetFollowing(_account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	return _ICodexAM.Contract.GetFollowing(&_ICodexAM.CallOpts, _account, offset, limit)
}

// GetFollowing is a free data retrieval call binding the contract method 0x7475676b.
//
// Solidity: function getFollowing(address _account, uint256 offset, uint256 limit) view returns(uint256[] _ids, uint256 _totalLength)
func (_ICodexAM *ICodexAMCallerSession) GetFollowing(_account common.Address, offset *big.Int, limit *big.Int) (struct {
	Ids         []*big.Int
	TotalLength *big.Int
}, error) {
	return _ICodexAM.Contract.GetFollowing(&_ICodexAM.CallOpts, _account, offset, limit)
}

// EditAccount is a paid mutator transaction binding the contract method 0x00bcc529.
//
// Solidity: function editAccount(string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMTransactor) EditAccount(opts *bind.TransactOpts, _name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.contract.Transact(opts, "editAccount", _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// EditAccount is a paid mutator transaction binding the contract method 0x00bcc529.
//
// Solidity: function editAccount(string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMSession) EditAccount(_name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.Contract.EditAccount(&_ICodexAM.TransactOpts, _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// EditAccount is a paid mutator transaction binding the contract method 0x00bcc529.
//
// Solidity: function editAccount(string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMTransactorSession) EditAccount(_name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.Contract.EditAccount(&_ICodexAM.TransactOpts, _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMTransactor) Follow(opts *bind.TransactOpts, _targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.contract.Transact(opts, "follow", _targetAddr)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMSession) Follow(_targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.Contract.Follow(&_ICodexAM.TransactOpts, _targetAddr)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMTransactorSession) Follow(_targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.Contract.Follow(&_ICodexAM.TransactOpts, _targetAddr)
}

// Register is a paid mutator transaction binding the contract method 0x9a1b2fe9.
//
// Solidity: function register(address _account, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMTransactor) Register(opts *bind.TransactOpts, _account common.Address, _name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.contract.Transact(opts, "register", _account, _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// Register is a paid mutator transaction binding the contract method 0x9a1b2fe9.
//
// Solidity: function register(address _account, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMSession) Register(_account common.Address, _name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.Contract.Register(&_ICodexAM.TransactOpts, _account, _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// Register is a paid mutator transaction binding the contract method 0x9a1b2fe9.
//
// Solidity: function register(address _account, string _name, string _avatar, string _bio, string _company, string _location, string _website, string[] _socialAccounts) returns(bool)
func (_ICodexAM *ICodexAMTransactorSession) Register(_account common.Address, _name string, _avatar string, _bio string, _company string, _location string, _website string, _socialAccounts []string) (*types.Transaction, error) {
	return _ICodexAM.Contract.Register(&_ICodexAM.TransactOpts, _account, _name, _avatar, _bio, _company, _location, _website, _socialAccounts)
}

// Unfollow is a paid mutator transaction binding the contract method 0x015a4ead.
//
// Solidity: function unfollow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMTransactor) Unfollow(opts *bind.TransactOpts, _targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.contract.Transact(opts, "unfollow", _targetAddr)
}

// Unfollow is a paid mutator transaction binding the contract method 0x015a4ead.
//
// Solidity: function unfollow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMSession) Unfollow(_targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.Contract.Unfollow(&_ICodexAM.TransactOpts, _targetAddr)
}

// Unfollow is a paid mutator transaction binding the contract method 0x015a4ead.
//
// Solidity: function unfollow(address _targetAddr) returns(bool)
func (_ICodexAM *ICodexAMTransactorSession) Unfollow(_targetAddr common.Address) (*types.Transaction, error) {
	return _ICodexAM.Contract.Unfollow(&_ICodexAM.TransactOpts, _targetAddr)
}
