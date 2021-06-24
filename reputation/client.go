package reputationcontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0x67, 0x57, 0xe7, 0xff, 0x42, 0x39, 0x58, 0x9a, 0xbd, 0x7b, 0xa7, 0x65, 0x71, 0x8b, 0x80, 0xa4, 0x54, 0xd6, 0xfa, 0x91}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

// ListByEpoch invokes `listByEpoch` method of contract.
func (c *Client) ListByEpoch(epoch int64) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "listByEpoch", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Put invokes `put` method of contract.
func (c *Client) Put(epoch int64, peerID []byte, value []byte) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: peerID}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: value}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "put", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Get invokes `get` method of contract.
func (c *Client) Get(epoch int64, peerID []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: peerID}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "get", args, nil)
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

// GetByID invokes `getByID` method of contract.
func (c *Client) GetByID(id []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "getByID", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}
