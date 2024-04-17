package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	_type "github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*CollectibleTransfer)(nil)

type CollectibleTransfer Token

func (c CollectibleTransfer) Type() schema.Type {
	return _type.CollectibleTransfer
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleApprovalAction --transform=snake --trimprefix=ActionCollectibleApproval --output collectible_approval.go --json --sql
type CollectibleApprovalAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t CollectibleApprovalAction) Type() schema.Type {
	return _type.CollectibleApproval
}

const (
	ActionCollectibleApprovalApprove CollectibleApprovalAction = iota + 1
	ActionCollectibleApprovalRevoke
)

var _ Metadata = (*CollectibleApproval)(nil)

type CollectibleApproval struct {
	Action CollectibleApprovalAction `json:"action"`

	Token
}

func (c CollectibleApproval) Type() schema.Type {
	return _type.CollectibleApproval
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleTradeAction --transform=snake --trimprefix=ActionCollectibleTrade --output collectible_trade.go --json --sql
type CollectibleTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (r CollectibleTradeAction) Type() schema.Type {
	return _type.CollectibleTrade
}

const (
	ActionCollectibleTradeBuy CollectibleTradeAction = iota + 1
	ActionCollectibleTradeSell
)

var _ Metadata = (*CollectibleTradeAction)(nil)

type CollectibleTrade struct {
	Action CollectibleTradeAction `json:"action"`
	Token
	Cost *Token `json:"cost,omitempty"`
}

func (r CollectibleTrade) Type() schema.Type {
	return _type.CollectibleTrade
}
