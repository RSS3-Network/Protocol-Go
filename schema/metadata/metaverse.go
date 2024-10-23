package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MetaverseTradeAction --transform=snake --trimprefix=ActionMetaverseTrade --output metaverse_trade.go --json --sql
type MetaverseTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t MetaverseTradeAction) Type() schema.Type {
	return typex.MetaverseTrade
}

const (
	ActionMetaverseTradeBuy MetaverseTradeAction = iota + 1
	ActionMetaverseTradeList
	ActionMetaverseTradeSell
)

var _ Metadata = (*MetaverseTransfer)(nil)

type MetaverseTransfer Token

func (m MetaverseTransfer) Type() schema.Type {
	return typex.MetaverseTransfer
}

var _ Metadata = (*MetaverseMint)(nil)

type MetaverseMint Token

func (m MetaverseMint) Type() schema.Type {
	return typex.MetaverseMint
}

var _ Metadata = (*MetaverseBurn)(nil)

type MetaverseBurn Token

func (m MetaverseBurn) Type() schema.Type {
	return typex.MetaverseBurn
}

var _ Metadata = (*MetaverseTrade)(nil)

type MetaverseTrade struct {
	Action MetaverseTradeAction `json:"action,omitempty"`

	Token
	Cost Token `json:"cost,omitempty"`
}

func (m MetaverseTrade) Type() schema.Type {
	return typex.MetaverseTrade
}
