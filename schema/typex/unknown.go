package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=UnknownType --transform=snake --trimprefix=Type --output unknown_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=UnknownType --transform=snake --trimprefix=Type --output=../../openapi/type/Unknown.yaml -t ../../openapi/tmpl/Type.yaml.tmpl
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
