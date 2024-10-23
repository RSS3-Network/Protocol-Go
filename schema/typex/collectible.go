package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=CollectibleType --transform=snake --trimprefix=Collectible --output collectible_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=CollectibleType --transform=snake --trimprefix=Collectible --indent --example=approval --output=collectible_schema.json
type CollectibleType uint64

//goland:noinspection GoMixedReceiverTypes
func (i CollectibleType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i CollectibleType) Tag() tag.Tag {
	return tag.Collectible
}

const (
	CollectibleApproval CollectibleType = iota + 1
	CollectibleBurn
	CollectibleMint
	CollectibleTrade
	CollectibleTransfer
	CollectibleAuction
)
