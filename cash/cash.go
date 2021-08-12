// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cash

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

// CashABI is the input ABI used to generate the binding from.
const CashABI = "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"show\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"apply_cheque\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CashFuncSigs maps the 4-byte function signature to its string representation.
var CashFuncSigs = map[string]string{
	"163dbf98": "apply_cheque(address,uint256,address,uint256,bytes)",
}

// CashBin is the compiled bytecode used for deploying new contracts.
var CashBin = "0x608060405261039e806100136000396000f3fe60806040526004361061001e5760003560e01c8063163dbf9814610023575b600080fd5b610036610031366004610239565b61004a565b604051901515815260200160405180910390f35b6040516bffffffffffffffffffffffff19606087811b821660208401526034830187905285901b16605482015260688101839052600090819060880160405160208183030381529060405280519060200120905060006100aa8285610118565b9050806001600160a01b0316886001600160a01b03161415610108576040516001600160a01b0387169086156108fc029087906000818181858888f193505050501580156100fc573d6000803e3d6000fd5b5060019250505061010f565b6000925050505b95945050505050565b6000815160411461012b57506000610217565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156101715760009350505050610217565b601b8160ff16101561018b5761018881601b61031f565b90505b8060ff16601b141580156101a357508060ff16601c14155b156101b45760009350505050610217565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610207573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b80356001600160a01b038116811461023457600080fd5b919050565b600080600080600060a0868803121561025157600080fd5b61025a8661021d565b94506020860135935061026f6040870161021d565b925060608601359150608086013567ffffffffffffffff8082111561029357600080fd5b818801915088601f8301126102a757600080fd5b8135818111156102b9576102b9610352565b604051601f8201601f19908116603f011681019083821181831017156102e1576102e1610352565b816040528281528b60208487010111156102fa57600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b600060ff821660ff84168060ff0382111561034a57634e487b7160e01b600052601160045260246000fd5b019392505050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220243a7492e5c4338e0bd0f863f8df32837d72a78c9d9c5a025cf11582403ea08864736f6c63430008070033"

// DeployCash deploys a new Ethereum contract, binding an instance of Cash to it.
func DeployCash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cash, error) {
	parsed, err := abi.JSON(strings.NewReader(CashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// Cash is an auto generated Go binding around an Ethereum contract.
type Cash struct {
	CashCaller     // Read-only binding to the contract
	CashTransactor // Write-only binding to the contract
	CashFilterer   // Log filterer for contract events
}

// CashCaller is an auto generated read-only Go binding around an Ethereum contract.
type CashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CashSession struct {
	Contract     *Cash             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CashCallerSession struct {
	Contract *CashCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CashTransactorSession struct {
	Contract     *CashTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CashRaw is an auto generated low-level Go binding around an Ethereum contract.
type CashRaw struct {
	Contract *Cash // Generic contract binding to access the raw methods on
}

// CashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CashCallerRaw struct {
	Contract *CashCaller // Generic read-only contract binding to access the raw methods on
}

// CashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CashTransactorRaw struct {
	Contract *CashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCash creates a new instance of Cash, bound to a specific deployed contract.
func NewCash(address common.Address, backend bind.ContractBackend) (*Cash, error) {
	contract, err := bindCash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cash{CashCaller: CashCaller{contract: contract}, CashTransactor: CashTransactor{contract: contract}, CashFilterer: CashFilterer{contract: contract}}, nil
}

// NewCashCaller creates a new read-only instance of Cash, bound to a specific deployed contract.
func NewCashCaller(address common.Address, caller bind.ContractCaller) (*CashCaller, error) {
	contract, err := bindCash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CashCaller{contract: contract}, nil
}

// NewCashTransactor creates a new write-only instance of Cash, bound to a specific deployed contract.
func NewCashTransactor(address common.Address, transactor bind.ContractTransactor) (*CashTransactor, error) {
	contract, err := bindCash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CashTransactor{contract: contract}, nil
}

// NewCashFilterer creates a new log filterer instance of Cash, bound to a specific deployed contract.
func NewCashFilterer(address common.Address, filterer bind.ContractFilterer) (*CashFilterer, error) {
	contract, err := bindCash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CashFilterer{contract: contract}, nil
}

// bindCash binds a generic wrapper to an already deployed contract.
func bindCash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.CashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.CashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cash *CashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cash *CashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cash *CashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cash.Contract.contract.Transact(opts, method, params...)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashTransactor) ApplyCheque(opts *bind.TransactOpts, userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.contract.Transact(opts, "apply_cheque", userAddr, nonce, stAddr, payAmount, sign)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashSession) ApplyCheque(userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, userAddr, nonce, stAddr, payAmount, sign)
}

// ApplyCheque is a paid mutator transaction binding the contract method 0x163dbf98.
//
// Solidity: function apply_cheque(address userAddr, uint256 nonce, address stAddr, uint256 payAmount, bytes sign) payable returns(bool)
func (_Cash *CashTransactorSession) ApplyCheque(userAddr common.Address, nonce *big.Int, stAddr common.Address, payAmount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Cash.Contract.ApplyCheque(&_Cash.TransactOpts, userAddr, nonce, stAddr, payAmount, sign)
}

// CashShowIterator is returned from FilterShow and is used to iterate over the raw logs and unpacked data for Show events raised by the Cash contract.
type CashShowIterator struct {
	Event *CashShow // Event containing the contract specifics and raw log

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
func (it *CashShowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CashShow)
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
		it.Event = new(CashShow)
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
func (it *CashShowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CashShowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CashShow represents a Show event raised by the Cash contract.
type CashShow struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterShow is a free log retrieval operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) FilterShow(opts *bind.FilterOpts) (*CashShowIterator, error) {

	logs, sub, err := _Cash.contract.FilterLogs(opts, "show")
	if err != nil {
		return nil, err
	}
	return &CashShowIterator{contract: _Cash.contract, event: "show", logs: logs, sub: sub}, nil
}

// WatchShow is a free log subscription operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) WatchShow(opts *bind.WatchOpts, sink chan<- *CashShow) (event.Subscription, error) {

	logs, sub, err := _Cash.contract.WatchLogs(opts, "show")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CashShow)
				if err := _Cash.contract.UnpackLog(event, "show", log); err != nil {
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

// ParseShow is a log parse operation binding the contract event 0xfb687b8805689fd7d97b0f174d316f660a67e3dc6bfb412c484bef58f01f18e6.
//
// Solidity: event show(bytes arg0)
func (_Cash *CashFilterer) ParseShow(log types.Log) (*CashShow, error) {
	event := new(CashShow)
	if err := _Cash.contract.UnpackLog(event, "show", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecoverABI is the input ABI used to generate the binding from.
const RecoverABI = "[]"

// RecoverBin is the compiled bytecode used for deploying new contracts.
var RecoverBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122008c2e92f83a37e13d63ea20a5f4d22e5ace01e54733a6573342e2c4ef4c4ebfa64736f6c63430008070033"

// DeployRecover deploys a new Ethereum contract, binding an instance of Recover to it.
func DeployRecover(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recover, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RecoverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// Recover is an auto generated Go binding around an Ethereum contract.
type Recover struct {
	RecoverCaller     // Read-only binding to the contract
	RecoverTransactor // Write-only binding to the contract
	RecoverFilterer   // Log filterer for contract events
}

// RecoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoverSession struct {
	Contract     *Recover          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoverCallerSession struct {
	Contract *RecoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoverTransactorSession struct {
	Contract     *RecoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoverRaw struct {
	Contract *Recover // Generic contract binding to access the raw methods on
}

// RecoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoverCallerRaw struct {
	Contract *RecoverCaller // Generic read-only contract binding to access the raw methods on
}

// RecoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoverTransactorRaw struct {
	Contract *RecoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecover creates a new instance of Recover, bound to a specific deployed contract.
func NewRecover(address common.Address, backend bind.ContractBackend) (*Recover, error) {
	contract, err := bindRecover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// NewRecoverCaller creates a new read-only instance of Recover, bound to a specific deployed contract.
func NewRecoverCaller(address common.Address, caller bind.ContractCaller) (*RecoverCaller, error) {
	contract, err := bindRecover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverCaller{contract: contract}, nil
}

// NewRecoverTransactor creates a new write-only instance of Recover, bound to a specific deployed contract.
func NewRecoverTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoverTransactor, error) {
	contract, err := bindRecover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverTransactor{contract: contract}, nil
}

// NewRecoverFilterer creates a new log filterer instance of Recover, bound to a specific deployed contract.
func NewRecoverFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoverFilterer, error) {
	contract, err := bindRecover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoverFilterer{contract: contract}, nil
}

// bindRecover binds a generic wrapper to an already deployed contract.
func bindRecover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.RecoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transact(opts, method, params...)
}
