package balancecontract

import (
	"github.com/nspcc-dev/neo-go/pkg/rpc/client"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

var contractHash = util.Uint160{0x72, 0x6e, 0x41, 0x75, 0xe, 0xcb, 0x9, 0x63, 0xae, 0xdb, 0x6d, 0xe3, 0x17, 0xd1, 0xfc, 0x69, 0x70, 0x18, 0xe3, 0xda}

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

// Burn invokes `burn` method of contract.
func (c *Client) Burn(from util.Uint160, amount int64, txDetails []byte) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: from.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: txDetails}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "burn", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// BalanceOf invokes `balanceOf` method of contract.
func (c *Client) BalanceOf(account util.Uint160) (int, error) {
	args := make([]smartcontract.Parameter, 1)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: account.BytesBE()}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "balanceOf", args, nil)
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

// Transfer invokes `transfer` method of contract.
func (c *Client) Transfer(from util.Uint160, to util.Uint160, amount int64, data interface{}) (bool, error) {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: from.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: to.BytesBE()}
	args[2] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: data}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "transfer", args, nil)
	if err != nil {
		return false, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return false, err
	}

	return client.TopIntFromStack(result.Stack)
}

// Lock invokes `lock` method of contract.
func (c *Client) Lock(txDetails []byte, from util.Uint160, to util.Uint160, amount int64, until int64) error {
	args := make([]smartcontract.Parameter, 5)
	args[0] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: txDetails}
	args[1] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: from.BytesBE()}
	args[2] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: to.BytesBE()}
	args[3] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[4] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: until}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "lock", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// TransferX invokes `transferX` method of contract.
func (c *Client) TransferX(from util.Uint160, to util.Uint160, amount int64, details []byte) error {
	args := make([]smartcontract.Parameter, 4)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: from.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: to.BytesBE()}
	args[2] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[3] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: details}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "transferX", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// Mint invokes `mint` method of contract.
func (c *Client) Mint(to util.Uint160, amount int64, txDetails []byte) error {
	args := make([]smartcontract.Parameter, 3)
	args[0] = smartcontract.Parameter{Type: smartcontract.Hash160Type, Value: to.BytesBE()}
	args[1] = smartcontract.Parameter{Type: smartcontract.IntegerType, Value: amount}
	args[2] = smartcontract.Parameter{Type: smartcontract.AnyType, Value: txDetails}
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "mint", args, nil)
	if err != nil {
		return err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return err
	}

	return client.TopIntFromStack(result.Stack)
}

// CreateToken invokes `createToken` method of contract.
func (c *Client) CreateToken() (interface{}, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "createToken", args, nil)
	if err != nil {
		return nil, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return [], err
	}

	return client.TopIntFromStack(result.Stack)
}

// Decimals invokes `decimals` method of contract.
func (c *Client) Decimals() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "decimals", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
	}

	return client.TopIntFromStack(result.Stack)
}

// Symbol invokes `symbol` method of contract.
func (c *Client) Symbol() (string, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "symbol", args, nil)
	if err != nil {
		return &#34;\&#34;\&#34;&#34;, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return &#34;&#34;, err
	}

	return client.TopIntFromStack(result.Stack)
}

// TotalSupply invokes `totalSupply` method of contract.
func (c *Client) TotalSupply() (int, error) {
	args := make([]smartcontract.Parameter, 0)
	
	result, err := (*client.Client)(c).InvokeFunction(contractHash, "totalSupply", args, nil)
	if err != nil {
		return 0, err
	}

	err = client.GetInvocationError(result)
	if err != nil {
		return 0, err
	}

	return client.TopIntFromStack(result.Stack)
}
