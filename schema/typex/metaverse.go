package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=MetaverseType --transform=snake --trimprefix=Metaverse --output metaverse_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=MetaverseType --transform=snake --trimprefix=Metaverse --output=metaverse.yaml -t schema.yaml.tmpl
type MetaverseType uint64

//goland:noinspection GoMixedReceiverTypes
func (i MetaverseType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i MetaverseType) Tag() tag.Tag {
	return tag.Metaverse
}

const (
	MetaverseBurn MetaverseType = iota + 1
	MetaverseMint
	MetaverseTrade
	MetaverseTransfer
)
