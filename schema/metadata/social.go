package metadata

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/tag"
	_type "github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*SocialPost)(nil)

type SocialPost struct {
	Handle        string   `json:"handle,omitempty"`
	Title         string   `json:"title,omitempty"`
	Summary       string   `json:"summary,omitempty"`
	Body          string   `json:"body,omitempty"`
	Media         []Media  `json:"media,omitempty"`
	ProfileID     string   `json:"profile_id,omitempty"`
	PublicationID string   `json:"publication_id,omitempty"`
	ContentURI    string   `json:"content_uri,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	AuthorURL     string   `json:"author_url,omitempty"`
	Reward        *Token   `json:"reward,omitempty"`
	Timestamp     uint64   `json:"timestamp,omitempty"`

	Target    *SocialPost `json:"target,omitempty"`
	TargetURL string      `json:"target_url,omitempty"`
}

type Media struct {
	Address  string `json:"address"`
	MimeType string `json:"mime_type"`
}

func (p SocialPost) Type() schema.Type {
	return _type.SocialPost
}

var _ Metadata = (*SocialProfileAction)(nil)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=SocialProfileAction --transform=snake --trimprefix=ActionSocialProfile --output social_profile.go --json --sql
type SocialProfileAction uint64

func (s SocialProfileAction) Type() schema.Type {
	return _type.SocialProfile
}

const (
	ActionSocialProfileCreate SocialProfileAction = iota + 1
	ActionSocialProfileRenew
	ActionSocialProfileUnwrap
	ActionSocialProfileUpdate
	ActionSocialProfileWrap
)

type SocialProfile struct {
	Action    SocialProfileAction `json:"action,omitempty"`
	ProfileID string              `json:"profile_id,omitempty"`
	Address   common.Address      `json:"address,omitempty"`
	Handle    string              `json:"handle,omitempty"`
	ImageURI  string              `json:"image_uri,omitempty"`
	Bio       string              `json:"bio,omitempty"`
	Name      string              `json:"name,omitempty"`
	Expiry    *time.Time          `json:"expiry,omitempty"`
	Key       string              `json:"key,omitempty"`
	Value     string              `json:"value,omitempty"`
}

func (f SocialProfile) Type() schema.Type {
	return _type.SocialProfile
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=SocialProxyAction --transform=snake --trimprefix=ActionSocialProxy --output social_proxy.go --json --sql
type SocialProxyAction uint64

func (s SocialProxyAction) Type() schema.Type {
	return _type.SocialProfile
}

const (
	ActionSocialProxyAppoint SocialProxyAction = iota + 1
	ActionSocialProxyRemove
)

type SocialProxy struct {
	Action       SocialProxyAction `json:"action,omitempty"`
	ProxyAddress common.Address    `json:"proxy_address"`

	Profile SocialProfile `json:"profile,omitempty"`
}

func (s SocialProxy) Tag() tag.Tag {
	return tag.Social
}

func (s SocialProxy) Type() schema.Type {
	return _type.SocialProxy
}
