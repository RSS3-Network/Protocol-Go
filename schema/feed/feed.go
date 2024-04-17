package feed

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	_type "github.com/rss3-network/protocol-go/schema/typex"
)

type Feed struct {
	ID           string          `json:"id"`
	Owner        string          `json:"owner,omitempty"`
	Network      network.Network `json:"network"`
	Index        uint            `json:"index"`
	From         string          `json:"from"`
	To           string          `json:"to"`
	Tag          tag.Tag         `json:"tag"`
	Type         schema.Type     `json:"type"`
	Platform     string          `json:"platform,omitempty"`
	Fee          *Fee            `json:"fee,omitempty"`
	Calldata     *Calldata       `json:"calldata,omitempty"`
	TotalActions uint            `json:"total_actions"`
	Actions      []*Action       `json:"actions"`
	Direction    Direction       `json:"direction,omitempty"`
	Status       bool            `json:"success"`
	Timestamp    uint64          `json:"timestamp"`
}

// Option is a function that can be used to modify a feed,
// it is used in the feed builder.
type Option func(feed *Feed) error

func WithFeedPlatform(platform string) Option {
	return func(feed *Feed) error {
		feed.Platform = platform

		return nil
	}
}

func NewUnknownFeed(feed *Feed) *Feed {
	unknownFeed := &Feed{
		Type:         _type.Unknown,
		Tag:          tag.Unknown,
		Network:      feed.Network,
		ID:           feed.ID,
		Owner:        feed.Owner,
		Index:        feed.Index,
		From:         feed.From,
		To:           feed.To,
		Platform:     feed.Platform,
		Fee:          feed.Fee,
		TotalActions: feed.TotalActions,
		Direction:    feed.Direction,
		Status:       feed.Status,
		Timestamp:    feed.Timestamp,
	}

	return unknownFeed
}

type Feeds []*Feed

var _ json.Unmarshaler = (*Feeds)(nil)

func (f *Feeds) UnmarshalJSON(bytes []byte) error {
	type FeedAlias Feed

	type feed struct {
		*FeedAlias
		TypeX string `json:"type"`
	}

	var temp []*feed

	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		return fmt.Errorf("unmarshal feeds: %w", err)
	}

	for _, v := range temp {
		v.TotalActions = uint(len(v.Actions))

		v.Type, err = schema.TypeString(v.Tag, v.TypeX)
		if err != nil {
			return fmt.Errorf("unmarshal type: %w", err)
		}

		*f = append(*f, (*Feed)(v.FeedAlias))
	}

	return nil
}
