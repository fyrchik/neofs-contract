package auditcontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0xe2, 0x6d, 0x47, 0xbc, 0x1a, 0xf7, 0x43, 0xf8, 0x4, 0x27, 0x4, 0xa, 0x59, 0x37, 0x99, 0x51, 0x58, 0x75, 0x81, 0x5a}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

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

// List invokes `list` method of contract.
func (c *Client) List() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "list", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// ListByNode invokes `listByNode` method of contract.
func (c *Client) ListByNode(epoch int64, cid []byte, key []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: cid}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: key}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "listByNode", args, nil)
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
func (c *Client) Put(rawAuditResult []byte) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: rawAuditResult}
	
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
func (c *Client) Get(id []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	
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

// ListByCID invokes `listByCID` method of contract.
func (c *Client) ListByCID(epoch int64, cid []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: cid}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "listByCID", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}
