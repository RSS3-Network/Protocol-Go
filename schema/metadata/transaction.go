package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*TransactionTransfer)(nil)

type TransactionTransfer Token

func (t TransactionTransfer) Type() schema.Type {
	return typex.TransactionTransfer
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=TransactionApprovalAction --transform=snake --trimprefix=ActionTransaction --output transaction_approval.go --json --sql
type TransactionApprovalAction uint64

const (
	ActionTransactionApprove TransactionApprovalAction = iota + 1
	ActionTransactionRevoke
)

var _ Metadata = (*TransactionApproval)(nil)

type TransactionApproval struct {
	Action TransactionApprovalAction `json:"action"`

	Token
}

func (t TransactionApproval) Type() schema.Type {
	return typex.TransactionApproval
}

var _ Metadata = (*TransactionBridge)(nil)

type TransactionBridge struct {
	Action        TransactionBridgeAction `json:"action"`
	SourceNetwork network.Network         `json:"source_network"`
	TargetNetwork network.Network         `json:"target_network"`
	Token         Token                   `json:"token"`
}

func (t TransactionBridge) Type() schema.Type {
	return typex.TransactionBridge
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=TransactionBridgeAction --transform=snake --trimprefix=ActionTransactionBridge --output transaction_bridge.go --json --sql
type TransactionBridgeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t TransactionBridgeAction) Type() schema.Type {
	return typex.TransactionBridge
}

const (
	ActionTransactionBridgeDeposit TransactionBridgeAction = iota + 1
	ActionTransactionBridgeWithdraw
)
