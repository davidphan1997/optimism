package crossdomain

import (
	"errors"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

var (
	abiTrue                      = common.Hash{31: 0x01}
	abiFalse                     = common.Hash{}
	errLegacyStorageSlotNotFound = errors.New("")
)

// This takes a state db and a list of withdrawals
func MigrateWithdrawals(withdrawals []*PendingWithdrawal, db vm.StateDB) error {
	for _, legacy := range withdrawals {
		legacySlot, err := legacy.StorageSlot()
		if err != nil {
			return err
		}

		legacyValue := db.GetState(predeploys.L2ToL1MessagePasserAddr, legacySlot)
		if legacyValue != abiTrue {
			return errLegacyStorageSlotNotFound
		}
		db.SetState(predeploys.L2ToL1MessagePasserAddr, legacySlot, abiFalse)

		withdrawal, err := MigrateWithdrawal(&legacy.LegacyWithdrawal)
		if err != nil {
			return err
		}

		slot, err := withdrawal.StorageSlot()
		if err != nil {
			return err
		}

		db.SetState(predeploys.L2ToL1MessagePasserAddr, slot, abiTrue)
	}
	return nil
}

// TODO(tynes): how to test this effectively?
// MigrateWithdrawal will turn a LegacyWithdrawal into a bedrock
// style Withdrawal.
func MigrateWithdrawal(withdrawal *LegacyWithdrawal) (*Withdrawal, error) {
	value := new(big.Int)

	// TODO: pass these in via args
	l1CrossDomainMessenger := common.Address{}
	l1StandardBridge := common.Address{}

	isFromL2StandardBridge := *withdrawal.Sender == predeploys.L2StandardBridgeAddr
	isToL1StandardBridge := *withdrawal.Target == l1StandardBridge

	if isFromL2StandardBridge && isToL1StandardBridge {
		abi, err := bindings.L1StandardBridgeMetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		data, err := abi.Unpack("finalizeETHWithdrawal", withdrawal.Data[4:])
		if err != nil {
			return nil, err
		}
		value = data[2].(*big.Int)
	}

	abi, err := bindings.L1CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	versionedNonce := EncodeVersionedNonce(withdrawal.Nonce, common.Big1)
	data, err := abi.Pack(
		"relayMessage",
		versionedNonce,
		withdrawal.Sender,
		withdrawal.Target,
		value,
		withdrawal.Data,
	)
	if err != nil {
		return nil, err
	}

	w := NewWithdrawal(
		withdrawal.Nonce,
		&predeploys.L2CrossDomainMessengerAddr,
		&l1CrossDomainMessenger,
		value,
		new(big.Int),
		data,
	)
	return w, nil
}
