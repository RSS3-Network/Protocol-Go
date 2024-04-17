package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/tag"
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

func (m MetaverseTransfer) Tag() tag.Tag {
	return tag.Metaverse
}

func (m MetaverseTransfer) Type() schema.Type {
	return typex.MetaverseTransfer
}

var _ Metadata = (*MetaverseTrade)(nil)

type MetaverseTrade struct {
	Action MetaverseTradeAction `json:"action,omitempty"`

	Token
	Cost Token `json:"cost,omitempty"`
}

func (m MetaverseTrade) Tag() tag.Tag {
	return tag.Metaverse
}

func (m MetaverseTrade) Type() schema.Type {
	return typex.MetaverseTrade
}
