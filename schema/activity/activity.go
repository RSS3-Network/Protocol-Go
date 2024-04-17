package activity

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

type Activity struct {
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

type Fee struct {
	Address *string         `json:"address,omitempty"`
	Amount  decimal.Decimal `json:"amount"`
	Decimal uint            `json:"decimal"`
}

// Option is a function that can be used to modify an Activity,
// it is used in the activity builder.
type Option func(activity *Activity) error

func WithActivityPlatform(platform string) Option {
	return func(activity *Activity) error {
		activity.Platform = platform

		return nil
	}
}

func NewUnknownActivity(activity *Activity) *Activity {
	unknownActivity := &Activity{
		Type:         typex.Unknown,
		Tag:          tag.Unknown,
		Network:      activity.Network,
		ID:           activity.ID,
		Owner:        activity.Owner,
		Index:        activity.Index,
		From:         activity.From,
		To:           activity.To,
		Platform:     activity.Platform,
		Fee:          activity.Fee,
		TotalActions: activity.TotalActions,
		Direction:    activity.Direction,
		Status:       activity.Status,
		Timestamp:    activity.Timestamp,
	}

	return unknownActivity
}

type Activities []*Activity

var _ json.Unmarshaler = (*Activities)(nil)

func (f *Activities) UnmarshalJSON(bytes []byte) error {
	type ActivityAlias Activity

	type activity struct {
		*ActivityAlias
		TypeX string `json:"type"`
	}

	var temp []*activity

	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		return fmt.Errorf("unmarshal activities: %w", err)
	}

	for _, v := range temp {
		v.TotalActions = uint(len(v.Actions))

		v.Type, err = schema.TypeString(v.Tag, v.TypeX)
		if err != nil {
			return fmt.Errorf("unmarshal type: %w", err)
		}

		*f = append(*f, (*Activity)(v.ActivityAlias))
	}

	return nil
}
