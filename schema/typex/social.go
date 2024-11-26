package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=SocialType --transform=snake --trimprefix=Social --output social_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=SocialType --transform=snake --trimprefix=Social --output=../../openapi/type/Social.yaml -t ../../openapi/tmpl/Type.yaml.tmpl
type SocialType uint64

//goland:noinspection GoMixedReceiverTypes
func (i SocialType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i SocialType) Tag() tag.Tag {
	return tag.Social
}

const (
	SocialComment SocialType = iota + 1
	SocialDelete
	SocialMint
	SocialPost
	SocialProfile
	SocialProxy
	SocialRevise
	SocialReward
	SocialShare
	SocialLike
)
