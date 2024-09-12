package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*CollectibleTransfer)(nil)

type CollectibleTransfer Token

func (c CollectibleTransfer) Type() schema.Type {
	return typex.CollectibleTransfer
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleApprovalAction --transform=snake --trimprefix=ActionCollectibleApproval --output collectible_approval.go --json --sql
type CollectibleApprovalAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t CollectibleApprovalAction) Type() schema.Type {
	return typex.CollectibleApproval
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
	return typex.CollectibleApproval
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleTradeAction --transform=snake --trimprefix=ActionCollectibleTrade --output collectible_trade.go --json --sql
type CollectibleTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (r CollectibleTradeAction) Type() schema.Type {
	return typex.CollectibleTrade
}

const (
	ActionCollectibleTradeBuy CollectibleTradeAction = iota + 1
	ActionCollectibleTradeSell
	ActionCollectibleTradeOffer
	ActionCollectibleTradeSet
	ActionCollectibleTradeCreate
	ActionCollectibleTradeFinalize
)

var _ Metadata = (*CollectibleTradeAction)(nil)

type CollectibleTrade struct {
	Action CollectibleTradeAction `json:"action"`
	Token
	Cost *Token `json:"cost,omitempty"`
}

func (r CollectibleTrade) Type() schema.Type {
	return typex.CollectibleTrade
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleAuctionAction --transform=snake --trimprefix=ActionCollectibleAuction --output collectible_auction.go --json --sql
type CollectibleAuctionAction uint64

var _ Metadata = (*CollectibleAuction)(nil)

type CollectibleAuction struct {
	Action CollectibleAuctionAction `json:"action"`
	Token
	Cost *Token `json:"cost,omitempty"`
}

func (r CollectibleAuction) Type() schema.Type {
	return typex.CollectibleAuction
}

//goland:noinspection GoMixedReceiverTypes
func (r CollectibleAuctionAction) Type() schema.Type {
	return typex.CollectibleAuction
}

const (
	ActionCollectibleAuctionCreate CollectibleAuctionAction = iota + 1
	ActionCollectibleAuctionBid
	ActionCollectibleAuctionFinalize
	ActionCollectibleAuctionCancel
	ActionCollectibleAuctionUpdate
	ActionCollectibleAuctionInvalidate
)
