package netmapcontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0x4f, 0x23, 0xb6, 0x17, 0xe1, 0xb7, 0xa7, 0xfd, 0x3b, 0x28, 0xbc, 0xc0, 0xd1, 0x8d, 0x29, 0x36, 0xd7, 0x4f, 0x9a, 0x8f}

// Client is a wrapper over RPC-client mirroring methods of smartcontract.
type Client client.Client

// AddPeer invokes `addPeer` method of contract.
func (c *Client) AddPeer(nodeInfo []byte) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: nodeInfo}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "addPeer", args, nil)
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

// Epoch invokes `epoch` method of contract.
func (c *Client) Epoch() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "epoch", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
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

// InnerRingList invokes `innerRingList` method of contract.
func (c *Client) InnerRingList() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "innerRingList", args, nil)
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

// SnapshotByEpoch invokes `snapshotByEpoch` method of contract.
func (c *Client) SnapshotByEpoch(epoch int64) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: epoch}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "snapshotByEpoch", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Netmap invokes `netmap` method of contract.
func (c *Client) Netmap() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "netmap", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Snapshot invokes `snapshot` method of contract.
func (c *Client) Snapshot(diff int64) (interface{}, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: diff}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "snapshot", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// UpdateInnerRing invokes `updateInnerRing` method of contract.
func (c *Client) UpdateInnerRing(keys []interface{}) error {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: keys}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "updateInnerRing", args, nil)
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

// NetmapCandidates invokes `netmapCandidates` method of contract.
func (c *Client) NetmapCandidates() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "netmapCandidates", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
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

// UpdateState invokes `updateState` method of contract.
func (c *Client) UpdateState(state int64, publicKey []byte) error {
	args := make([]smartcontract.Parameter, 2)
	args[0] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: state}
	args[1] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: publicKey}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "updateState", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}
