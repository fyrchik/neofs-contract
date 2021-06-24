package neofsidcontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0xac, 0x85, 0x6, 0x34, 0x4d, 0x7a, 0x4, 0x85, 0xca, 0x8a, 0xa3, 0x90, 0x43, 0x80, 0x38, 0xac, 0xf4, 0x91, 0xdd, 0xdd}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

// RemoveKey invokes `removeKey` method of contract.
func (c *Client) RemoveKey(owner []byte, keys []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: owner}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: keys}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "removeKey", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Key invokes `key` method of contract.
func (c *Client) Key(owner []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: owner}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "key", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
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

// AddKey invokes `addKey` method of contract.
func (c *Client) AddKey(owner []byte, keys []interface{}) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: owner}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: keys}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "addKey", args, nil)
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
