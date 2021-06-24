package alphabetcontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0x48, 0x4, 0x44, 0xf6, 0x7, 0xc4, 0x77, 0x6a, 0xa7, 0x5a, 0xf, 0xe6, 0xb8, 0x3d, 0xbc, 0x70, 0xe2, 0x39, 0xa2, 0x54}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

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

// Vote invokes `vote` method of contract.
func (c *Client) Vote(epoch int64, candidates []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: candidates}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "vote", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Emit invokes `emit` method of contract.
func (c *Client) Emit() error {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "emit", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Gas invokes `gas` method of contract.
func (c *Client) Gas() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "gas", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
	}

	return client.TopIntFromStack(result.Stack)
}

// Neo invokes `neo` method of contract.
func (c *Client) Neo() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "neo", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
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

// Name invokes `name` method of contract.
func (c *Client) Name() (string, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "name", args, nil)
	if err != nil {
		return &#34;\&#34;\&#34;&#34;, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return &#34;&#34;, err
	}

	return client.TopIntFromStack(result.Stack)
}
