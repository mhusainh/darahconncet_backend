// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SertifikatDonasiMetaData contains all meta data concerning the SertifikatDonasi contract.
var SertifikatDonasiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendonor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nomorSertifikat\",\"type\":\"uint256\"}],\"name\":\"SertifikatDibuat\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"daftarSertifikat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pendonor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"namaPendonor\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tanggalDonasi\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nomorSertifikat\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"alamatPendonor\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"penerbit\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"adminVerifikator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pendonor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_namaPendonor\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nomorSertifikat\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_alamatPendonor\",\"type\":\"string\"}],\"name\":\"mintSertifikat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SertifikatDonasiABI is the input ABI used to generate the binding from.
// Deprecated: Use SertifikatDonasiMetaData.ABI instead.
var SertifikatDonasiABI = SertifikatDonasiMetaData.ABI

// SertifikatDonasi is an auto generated Go binding around an Ethereum contract.
type SertifikatDonasi struct {
	SertifikatDonasiCaller     // Read-only binding to the contract
	SertifikatDonasiTransactor // Write-only binding to the contract
	SertifikatDonasiFilterer   // Log filterer for contract events
}

// SertifikatDonasiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SertifikatDonasiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SertifikatDonasiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SertifikatDonasiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SertifikatDonasiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SertifikatDonasiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SertifikatDonasiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SertifikatDonasiSession struct {
	Contract     *SertifikatDonasi // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SertifikatDonasiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SertifikatDonasiCallerSession struct {
	Contract *SertifikatDonasiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SertifikatDonasiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SertifikatDonasiTransactorSession struct {
	Contract     *SertifikatDonasiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SertifikatDonasiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SertifikatDonasiRaw struct {
	Contract *SertifikatDonasi // Generic contract binding to access the raw methods on
}

// SertifikatDonasiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SertifikatDonasiCallerRaw struct {
	Contract *SertifikatDonasiCaller // Generic read-only contract binding to access the raw methods on
}

// SertifikatDonasiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SertifikatDonasiTransactorRaw struct {
	Contract *SertifikatDonasiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSertifikatDonasi creates a new instance of SertifikatDonasi, bound to a specific deployed contract.
func NewSertifikatDonasi(address common.Address, backend bind.ContractBackend) (*SertifikatDonasi, error) {
	contract, err := bindSertifikatDonasi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasi{SertifikatDonasiCaller: SertifikatDonasiCaller{contract: contract}, SertifikatDonasiTransactor: SertifikatDonasiTransactor{contract: contract}, SertifikatDonasiFilterer: SertifikatDonasiFilterer{contract: contract}}, nil
}

// NewSertifikatDonasiCaller creates a new read-only instance of SertifikatDonasi, bound to a specific deployed contract.
func NewSertifikatDonasiCaller(address common.Address, caller bind.ContractCaller) (*SertifikatDonasiCaller, error) {
	contract, err := bindSertifikatDonasi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasiCaller{contract: contract}, nil
}

// NewSertifikatDonasiTransactor creates a new write-only instance of SertifikatDonasi, bound to a specific deployed contract.
func NewSertifikatDonasiTransactor(address common.Address, transactor bind.ContractTransactor) (*SertifikatDonasiTransactor, error) {
	contract, err := bindSertifikatDonasi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasiTransactor{contract: contract}, nil
}

// NewSertifikatDonasiFilterer creates a new log filterer instance of SertifikatDonasi, bound to a specific deployed contract.
func NewSertifikatDonasiFilterer(address common.Address, filterer bind.ContractFilterer) (*SertifikatDonasiFilterer, error) {
	contract, err := bindSertifikatDonasi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasiFilterer{contract: contract}, nil
}

// bindSertifikatDonasi binds a generic wrapper to an already deployed contract.
func bindSertifikatDonasi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SertifikatDonasiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SertifikatDonasi *SertifikatDonasiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SertifikatDonasi.Contract.SertifikatDonasiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SertifikatDonasi *SertifikatDonasiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.SertifikatDonasiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SertifikatDonasi *SertifikatDonasiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.SertifikatDonasiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SertifikatDonasi *SertifikatDonasiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SertifikatDonasi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SertifikatDonasi *SertifikatDonasiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SertifikatDonasi *SertifikatDonasiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.contract.Transact(opts, method, params...)
}

// DaftarSertifikat is a free data retrieval call binding the contract method 0xf55f6ba9.
//
// Solidity: function daftarSertifikat(uint256 ) view returns(uint256 id, address pendonor, string namaPendonor, uint256 tanggalDonasi, uint256 nomorSertifikat, string alamatPendonor, string penerbit, address adminVerifikator)
func (_SertifikatDonasi *SertifikatDonasiCaller) DaftarSertifikat(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id               *big.Int
	Pendonor         common.Address
	NamaPendonor     string
	TanggalDonasi    *big.Int
	NomorSertifikat  *big.Int
	AlamatPendonor   string
	Penerbit         string
	AdminVerifikator common.Address
}, error) {
	var out []interface{}
	err := _SertifikatDonasi.contract.Call(opts, &out, "daftarSertifikat", arg0)

	outstruct := new(struct {
		Id               *big.Int
		Pendonor         common.Address
		NamaPendonor     string
		TanggalDonasi    *big.Int
		NomorSertifikat  *big.Int
		AlamatPendonor   string
		Penerbit         string
		AdminVerifikator common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pendonor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NamaPendonor = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.TanggalDonasi = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NomorSertifikat = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AlamatPendonor = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Penerbit = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.AdminVerifikator = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DaftarSertifikat is a free data retrieval call binding the contract method 0xf55f6ba9.
//
// Solidity: function daftarSertifikat(uint256 ) view returns(uint256 id, address pendonor, string namaPendonor, uint256 tanggalDonasi, uint256 nomorSertifikat, string alamatPendonor, string penerbit, address adminVerifikator)
func (_SertifikatDonasi *SertifikatDonasiSession) DaftarSertifikat(arg0 *big.Int) (struct {
	Id               *big.Int
	Pendonor         common.Address
	NamaPendonor     string
	TanggalDonasi    *big.Int
	NomorSertifikat  *big.Int
	AlamatPendonor   string
	Penerbit         string
	AdminVerifikator common.Address
}, error) {
	return _SertifikatDonasi.Contract.DaftarSertifikat(&_SertifikatDonasi.CallOpts, arg0)
}

// DaftarSertifikat is a free data retrieval call binding the contract method 0xf55f6ba9.
//
// Solidity: function daftarSertifikat(uint256 ) view returns(uint256 id, address pendonor, string namaPendonor, uint256 tanggalDonasi, uint256 nomorSertifikat, string alamatPendonor, string penerbit, address adminVerifikator)
func (_SertifikatDonasi *SertifikatDonasiCallerSession) DaftarSertifikat(arg0 *big.Int) (struct {
	Id               *big.Int
	Pendonor         common.Address
	NamaPendonor     string
	TanggalDonasi    *big.Int
	NomorSertifikat  *big.Int
	AlamatPendonor   string
	Penerbit         string
	AdminVerifikator common.Address
}, error) {
	return _SertifikatDonasi.Contract.DaftarSertifikat(&_SertifikatDonasi.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SertifikatDonasi *SertifikatDonasiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SertifikatDonasi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SertifikatDonasi *SertifikatDonasiSession) Owner() (common.Address, error) {
	return _SertifikatDonasi.Contract.Owner(&_SertifikatDonasi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SertifikatDonasi *SertifikatDonasiCallerSession) Owner() (common.Address, error) {
	return _SertifikatDonasi.Contract.Owner(&_SertifikatDonasi.CallOpts)
}

// MintSertifikat is a paid mutator transaction binding the contract method 0x0bacf3a0.
//
// Solidity: function mintSertifikat(address _pendonor, string _namaPendonor, uint256 _nomorSertifikat, string _alamatPendonor) returns()
func (_SertifikatDonasi *SertifikatDonasiTransactor) MintSertifikat(opts *bind.TransactOpts, _pendonor common.Address, _namaPendonor string, _nomorSertifikat *big.Int, _alamatPendonor string) (*types.Transaction, error) {
	return _SertifikatDonasi.contract.Transact(opts, "mintSertifikat", _pendonor, _namaPendonor, _nomorSertifikat, _alamatPendonor)
}

// MintSertifikat is a paid mutator transaction binding the contract method 0x0bacf3a0.
//
// Solidity: function mintSertifikat(address _pendonor, string _namaPendonor, uint256 _nomorSertifikat, string _alamatPendonor) returns()
func (_SertifikatDonasi *SertifikatDonasiSession) MintSertifikat(_pendonor common.Address, _namaPendonor string, _nomorSertifikat *big.Int, _alamatPendonor string) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.MintSertifikat(&_SertifikatDonasi.TransactOpts, _pendonor, _namaPendonor, _nomorSertifikat, _alamatPendonor)
}

// MintSertifikat is a paid mutator transaction binding the contract method 0x0bacf3a0.
//
// Solidity: function mintSertifikat(address _pendonor, string _namaPendonor, uint256 _nomorSertifikat, string _alamatPendonor) returns()
func (_SertifikatDonasi *SertifikatDonasiTransactorSession) MintSertifikat(_pendonor common.Address, _namaPendonor string, _nomorSertifikat *big.Int, _alamatPendonor string) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.MintSertifikat(&_SertifikatDonasi.TransactOpts, _pendonor, _namaPendonor, _nomorSertifikat, _alamatPendonor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SertifikatDonasi *SertifikatDonasiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SertifikatDonasi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SertifikatDonasi *SertifikatDonasiSession) RenounceOwnership() (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.RenounceOwnership(&_SertifikatDonasi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SertifikatDonasi *SertifikatDonasiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.RenounceOwnership(&_SertifikatDonasi.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SertifikatDonasi *SertifikatDonasiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SertifikatDonasi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SertifikatDonasi *SertifikatDonasiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.TransferOwnership(&_SertifikatDonasi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SertifikatDonasi *SertifikatDonasiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SertifikatDonasi.Contract.TransferOwnership(&_SertifikatDonasi.TransactOpts, newOwner)
}

// SertifikatDonasiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SertifikatDonasi contract.
type SertifikatDonasiOwnershipTransferredIterator struct {
	Event *SertifikatDonasiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SertifikatDonasiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SertifikatDonasiOwnershipTransferred)
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
		it.Event = new(SertifikatDonasiOwnershipTransferred)
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
func (it *SertifikatDonasiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SertifikatDonasiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SertifikatDonasiOwnershipTransferred represents a OwnershipTransferred event raised by the SertifikatDonasi contract.
type SertifikatDonasiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SertifikatDonasi *SertifikatDonasiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SertifikatDonasiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SertifikatDonasi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasiOwnershipTransferredIterator{contract: _SertifikatDonasi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SertifikatDonasi *SertifikatDonasiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SertifikatDonasiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SertifikatDonasi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SertifikatDonasiOwnershipTransferred)
				if err := _SertifikatDonasi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SertifikatDonasi *SertifikatDonasiFilterer) ParseOwnershipTransferred(log types.Log) (*SertifikatDonasiOwnershipTransferred, error) {
	event := new(SertifikatDonasiOwnershipTransferred)
	if err := _SertifikatDonasi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SertifikatDonasiSertifikatDibuatIterator is returned from FilterSertifikatDibuat and is used to iterate over the raw logs and unpacked data for SertifikatDibuat events raised by the SertifikatDonasi contract.
type SertifikatDonasiSertifikatDibuatIterator struct {
	Event *SertifikatDonasiSertifikatDibuat // Event containing the contract specifics and raw log

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
func (it *SertifikatDonasiSertifikatDibuatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SertifikatDonasiSertifikatDibuat)
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
		it.Event = new(SertifikatDonasiSertifikatDibuat)
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
func (it *SertifikatDonasiSertifikatDibuatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SertifikatDonasiSertifikatDibuatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SertifikatDonasiSertifikatDibuat represents a SertifikatDibuat event raised by the SertifikatDonasi contract.
type SertifikatDonasiSertifikatDibuat struct {
	Id              *big.Int
	Pendonor        common.Address
	NomorSertifikat *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSertifikatDibuat is a free log retrieval operation binding the contract event 0x2dd29160cc9010dcd6255f2e140c0c97e43efef714a91f2ed5606e13a8de8bd4.
//
// Solidity: event SertifikatDibuat(uint256 indexed id, address indexed pendonor, uint256 nomorSertifikat)
func (_SertifikatDonasi *SertifikatDonasiFilterer) FilterSertifikatDibuat(opts *bind.FilterOpts, id []*big.Int, pendonor []common.Address) (*SertifikatDonasiSertifikatDibuatIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pendonorRule []interface{}
	for _, pendonorItem := range pendonor {
		pendonorRule = append(pendonorRule, pendonorItem)
	}

	logs, sub, err := _SertifikatDonasi.contract.FilterLogs(opts, "SertifikatDibuat", idRule, pendonorRule)
	if err != nil {
		return nil, err
	}
	return &SertifikatDonasiSertifikatDibuatIterator{contract: _SertifikatDonasi.contract, event: "SertifikatDibuat", logs: logs, sub: sub}, nil
}

// WatchSertifikatDibuat is a free log subscription operation binding the contract event 0x2dd29160cc9010dcd6255f2e140c0c97e43efef714a91f2ed5606e13a8de8bd4.
//
// Solidity: event SertifikatDibuat(uint256 indexed id, address indexed pendonor, uint256 nomorSertifikat)
func (_SertifikatDonasi *SertifikatDonasiFilterer) WatchSertifikatDibuat(opts *bind.WatchOpts, sink chan<- *SertifikatDonasiSertifikatDibuat, id []*big.Int, pendonor []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var pendonorRule []interface{}
	for _, pendonorItem := range pendonor {
		pendonorRule = append(pendonorRule, pendonorItem)
	}

	logs, sub, err := _SertifikatDonasi.contract.WatchLogs(opts, "SertifikatDibuat", idRule, pendonorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SertifikatDonasiSertifikatDibuat)
				if err := _SertifikatDonasi.contract.UnpackLog(event, "SertifikatDibuat", log); err != nil {
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

// ParseSertifikatDibuat is a log parse operation binding the contract event 0x2dd29160cc9010dcd6255f2e140c0c97e43efef714a91f2ed5606e13a8de8bd4.
//
// Solidity: event SertifikatDibuat(uint256 indexed id, address indexed pendonor, uint256 nomorSertifikat)
func (_SertifikatDonasi *SertifikatDonasiFilterer) ParseSertifikatDibuat(log types.Log) (*SertifikatDonasiSertifikatDibuat, error) {
	event := new(SertifikatDonasiSertifikatDibuat)
	if err := _SertifikatDonasi.contract.UnpackLog(event, "SertifikatDibuat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
