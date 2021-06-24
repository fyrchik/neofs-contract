package containercontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0x20, 0x6c, 0x95, 0x47, 0x26, 0xd4, 0xe1, 0x9, 0xbc, 0x5b, 0x81, 0x31, 0x5b, 0xad, 0x70, 0x54, 0x8, 0x24, 0xd, 0x36}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

// StopContainerEstimation invokes `stopContainerEstimation` method of contract.
func (c *Client) StopContainerEstimation(epoch int64) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "stopContainerEstimation", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// ListContainerSizes invokes `listContainerSizes` method of contract.
func (c *Client) ListContainerSizes(epoch int64) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "listContainerSizes", args, nil)
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
func (c *Client) Put(container []byte, signature []byte, publicKey []byte, token []byte) error {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: container}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: signature}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: publicKey}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: token}
	
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

// GetContainerSize invokes `getContainerSize` method of contract.
func (c *Client) GetContainerSize(id []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: id}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "getContainerSize", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
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

// List invokes `list` method of contract.
func (c *Client) List(owner []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: owner}
	
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

// StartContainerEstimation invokes `startContainerEstimation` method of contract.
func (c *Client) StartContainerEstimation(epoch int64) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "startContainerEstimation", args, nil)
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

// EACL invokes `eACL` method of contract.
func (c *Client) EACL(containerID []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: containerID}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "eACL", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// NewEpoch invokes `newEpoch` method of contract.
func (c *Client) NewEpoch(epochNum int64) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epochNum}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "newEpoch", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Delete invokes `delete` method of contract.
func (c *Client) Delete(containerID []byte, signature []byte, token []byte) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: containerID}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: signature}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: token}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "delete", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// PutContainerSize invokes `putContainerSize` method of contract.
func (c *Client) PutContainerSize(epoch int64, cid []byte, usedSize int64, pubKey []byte) error {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: cid}
	args[2] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: usedSize}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: pubKey}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "putContainerSize", args, nil)
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
func (c *Client) Get(containerID []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: containerID}
	
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

// Owner invokes `owner` method of contract.
func (c *Client) Owner(containerID []byte) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: containerID}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "owner", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// SetEACL invokes `setEACL` method of contract.
func (c *Client) SetEACL(eACL []byte, signature []byte, publicKey []byte, token []byte) error {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: eACL}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: signature}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: publicKey}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: token}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "setEACL", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}
