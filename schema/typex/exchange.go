package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=ExchangeType --transform=snake --trimprefix=Exchange --output exchange_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=ExchangeType --transform=snake --trimprefix=Exchange --indent --example=swap --output=exchange_schema.json
type ExchangeType uint64

//goland:noinspection GoMixedReceiverTypes
func (i ExchangeType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i ExchangeType) Tag() tag.Tag {
	return tag.Exchange
}

const (
	ExchangeLiquidity ExchangeType = iota + 1
	ExchangeStaking
	ExchangeSwap
	ExchangeLoan
)
