package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=UnknownType --transform=snake --trimprefix=Type --output unknown_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=UnknownType --transform=snake --trimprefix=Type --indent --example=unknown --output=unknown_schema.json
type UnknownType uint64

//goland:noinspection GoMixedReceiverTypes
func (i UnknownType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i UnknownType) Tag() tag.Tag {
	return tag.Unknown
}

const (
	Unknown UnknownType = iota
)
