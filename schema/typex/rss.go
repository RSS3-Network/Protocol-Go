package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=RSSType --transform=snake --trimprefix=RSS --output rss_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=RSSType --transform=snake --trimprefix=RSS --output=../../openapi/type/RSSType.yaml -t ../../openapi/tmpl/Type.yaml.tmpl
type RSSType uint64

//goland:noinspection GoMixedReceiverTypes
func (r RSSType) Name() string {
	return r.String()
}

//goland:noinspection GoMixedReceiverTypes
func (r RSSType) Tag() tag.Tag {
	return tag.RSS
}

const (
	RSSFeed RSSType = iota + 1
)
