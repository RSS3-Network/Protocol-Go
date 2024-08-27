package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleType --transform=snake --trimprefix=Collectible --output collectible_string.go --json --sql
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
