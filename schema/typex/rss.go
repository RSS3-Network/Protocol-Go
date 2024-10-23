package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=RSSType --transform=snake --trimprefix=RSS --output rss_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=RSSType --transform=snake --trimprefix=RSS --indent --example=feed --output=rss_schema.json
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
