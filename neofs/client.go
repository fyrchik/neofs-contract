package neofscontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0xe1, 0x7c, 0xc5, 0xcc, 0x6e, 0x2, 0xc1, 0x86, 0x33, 0x8a, 0x25, 0xbb, 0xb7, 0xc6, 0xb5, 0x2c, 0x5e, 0x99, 0x68, 0x1c}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

// Migrate invokes `migrate` method of contract.
func (c *Client) Migrate(script []byte, manifest []byte, data interface{}) (bool, error) {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: script}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: manifest}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: data}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "migrate", args, nil)
	if err != nil {
		return false, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return false, err
	}

	return client.TopIntFromStack(result.Stack)
}

// OnNEP17Payment invokes `onNEP17Payment` method of contract.
func (c *Client) OnNEP17Payment(from util.Uint160, amount int64, data interface{}) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: from.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: data}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "onNEP17Payment", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// AlphabetAddress invokes `alphabetAddress` method of contract.
func (c *Client) AlphabetAddress() (util.Uint160, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "alphabetAddress", args, nil)
	if err != nil {
		return util.Uint160{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0000000000000000000000000000000000000000, err
	}

	return client.TopIntFromStack(result.Stack)
}

// SetConfig invokes `setConfig` method of contract.
func (c *Client) SetConfig(id []byte, key []byte, val []byte) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: key}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: val}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "setConfig", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Cheque invokes `cheque` method of contract.
func (c *Client) Cheque(id []byte, user util.Uint160, amount int64, lockAcc []byte) error {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	args[1] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: user.BytesBE()}
	args[2] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: lockAcc}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "cheque", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// AlphabetUpdate invokes `alphabetUpdate` method of contract.
func (c *Client) AlphabetUpdate(id []byte, args []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: args}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "alphabetUpdate", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Config invokes `config` method of contract.
func (c *Client) Config(key []byte) (&lt;nil&gt;, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: key}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "config", args, nil)
	if err != nil {
		return &lt;nil&gt;, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return , err
	}

	return client.TopIntFromStack(result.Stack)
}

// InnerRingCandidateRemove invokes `innerRingCandidateRemove` method of contract.
func (c *Client) InnerRingCandidateRemove(key []byte) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: key}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "innerRingCandidateRemove", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Version invokes `version` method of contract.
func (c *Client) Version() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "version", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
	}

	return client.TopIntFromStack(result.Stack)
}

// InnerRingCandidateAdd invokes `innerRingCandidateAdd` method of contract.
func (c *Client) InnerRingCandidateAdd(key []byte) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: key}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "innerRingCandidateAdd", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// InnerRingCandidates invokes `innerRingCandidates` method of contract.
func (c *Client) InnerRingCandidates() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "innerRingCandidates", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// InitConfig invokes `initConfig` method of contract.
func (c *Client) InitConfig(args []interface{}) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: args}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "initConfig", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Withdraw invokes `withdraw` method of contract.
func (c *Client) Withdraw(user util.Uint160, amount int64) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: user.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "withdraw", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// AlphabetList invokes `alphabetList` method of contract.
func (c *Client) AlphabetList() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "alphabetList", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Unbind invokes `unbind` method of contract.
func (c *Client) Unbind(user []byte, keys []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: user}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: keys}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "unbind", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// ListConfig invokes `listConfig` method of contract.
func (c *Client) ListConfig() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "listConfig", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Bind invokes `bind` method of contract.
func (c *Client) Bind(user []byte, keys []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: user}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: keys}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "bind", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}
