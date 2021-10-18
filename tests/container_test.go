package tests

import (
	"crypto/sha256"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neo-go/pkg/core"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm/stackitem"
	"github.com/nspcc-dev/neo-go/pkg/wallet"
	"github.com/nspcc-dev/neofs-contract/nns"
	"github.com/stretchr/testify/require"
)

const containerPath = "../container"

const containerFee = 0_0100_0000

func deployContainerContract(t *testing.T, bc *core.Blockchain, addrNetmap, addrBalance, addrNNS util.Uint160) util.Uint160 {
	args := make([]interface{}, 6)
	args[0] = int64(0)
	args[1] = addrNetmap
	args[2] = addrBalance
	args[3] = util.Uint160{} // not needed for now
	args[4] = addrNNS
	args[5] = "neofs"
	return DeployContract(t, bc, containerPath, args)
}

func prepareContainerContract(t *testing.T, bc *core.Blockchain) (util.Uint160, util.Uint160) {
	addrNNS := DeployContract(t, bc, nnsPath, nil)

	ctrNetmap, err := ContractInfo(CommitteeAcc.Contract.ScriptHash(), netmapPath)
	require.NoError(t, err)

	ctrBalance, err := ContractInfo(CommitteeAcc.Contract.ScriptHash(), balancePath)
	require.NoError(t, err)

	ctrContainer, err := ContractInfo(CommitteeAcc.Contract.ScriptHash(), containerPath)
	require.NoError(t, err)

	deployNetmapContract(t, bc, ctrBalance.Hash, ctrContainer.Hash, "ContainerFee", int64(containerFee))
	balHash := deployBalanceContract(t, bc, ctrNetmap.Hash, ctrContainer.Hash)
	return deployContainerContract(t, bc, ctrNetmap.Hash, ctrBalance.Hash, addrNNS), balHash
}

func setContainerOwner(c []byte, owner *wallet.Account) {
	copy(c[7:], owner.Contract.ScriptHash().BytesBE())
}

func TestContainerPut(t *testing.T) {
	bc := NewChain(t)
	h, balanceHash := prepareContainerContract(t, bc)

	acc := NewAccount(t, bc)
	dummySig := make([]byte, 64)
	dummyPub := make([]byte, 33)
	dummyToken := make([]byte, 100)
	container := make([]byte, 100)
	setContainerOwner(container, acc)
	containerID := sha256.Sum256(container)

	putArgs := []interface{}{container, dummySig, dummyPub, dummyToken}
	tx := PrepareInvoke(t, bc, CommitteeAcc, h, "put", putArgs...)
	AddBlock(t, bc, tx)
	CheckFault(t, bc, tx.Hash(), "insufficient balance to create container")

	balanceMint(t, bc, acc, balanceHash, containerFee*1, []byte{})

	tx = PrepareInvoke(t, bc, acc, h, "put", putArgs...)
	AddBlock(t, bc, tx)
	CheckFault(t, bc, tx.Hash(), "alphabet witness check failed")

	tx = PrepareInvoke(t, bc, CommitteeAcc, h, "put", putArgs...)
	AddBlockCheckHalt(t, bc, tx)

	t.Run("with nice names", func(t *testing.T) {
		nnsHash := contracts[nnsPath].Hash

		balanceMint(t, bc, acc, balanceHash, containerFee*1, []byte{})

		putArgs := []interface{}{container, dummySig, dummyPub, dummyToken, "mycnt", ""}
		tx = PrepareInvoke(t, bc, CommitteeAcc, h, "putNamed", putArgs...)
		AddBlockCheckHalt(t, bc, tx)

		tx = PrepareInvoke(t, bc, acc, nnsHash, "resolve", "mycnt.neofs", int64(nns.TXT))
		CheckTestInvoke(t, bc, tx, stackitem.NewArray([]stackitem.Item{
			stackitem.NewByteArray([]byte(base58.Encode(containerID[:]))),
		}))

		tx = PrepareInvoke(t, bc, CommitteeAcc, h, "delete", containerID[:], dummySig, dummyToken)
		AddBlockCheckHalt(t, bc, tx)

		tx = PrepareInvoke(t, bc, CommitteeAcc, nnsHash, "resolve", "mycnt.neofs", int64(nns.TXT))
		CheckTestInvoke(t, bc, tx, stackitem.Null{})

		t.Run("register in advance", func(t *testing.T) {
			container[len(container)-1] = 10
			containerID = sha256.Sum256(container)

			tx = PrepareInvoke(t, bc, CommitteeAcc, nnsHash, "register",
				"second.neofs", CommitteeAcc.Contract.ScriptHash(),
				"whateveriwant@world.com", int64(0), int64(0), int64(0), int64(0))
			AddBlockCheckHalt(t, bc, tx)

			balanceMint(t, bc, acc, balanceHash, containerFee*1, []byte{})

			putArgs := []interface{}{container, dummySig, dummyPub, dummyToken, "second", "neofs"}
			tx = PrepareInvoke(t, bc, CommitteeAcc, h, "putNamed", putArgs...)
			AddBlockCheckHalt(t, bc, tx)

			tx = PrepareInvoke(t, bc, CommitteeAcc, nnsHash, "resolve", "second.neofs", int64(nns.TXT))
			CheckTestInvoke(t, bc, tx, stackitem.NewArray([]stackitem.Item{
				stackitem.NewByteArray([]byte(base58.Encode(containerID[:]))),
			}))
		})
	})
}
