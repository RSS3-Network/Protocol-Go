package metadata

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Token struct {
	Address        *string          `json:"address,omitempty"`
	ID             *decimal.Decimal `json:"id,omitempty"`
	Value          *decimal.Decimal `json:"value,omitempty"`
	Name           string           `json:"name,omitempty"`
	Symbol         string           `json:"symbol,omitempty"`
	URI            string           `json:"uri,omitempty"`
	Decimals       uint8            `json:"decimals,omitempty"`
	Standard       Standard         `json:"standard,omitempty"`
	ParsedImageURL string           `json:"parsed_image_url,omitempty"`
}

type NonFungibleTokenMetadata struct {
	Title        string          `json:"title,omitempty"`
	Description  string          `json:"description,omitempty"`
	ImageURL     string          `json:"image_url,omitempty"`
	ExternalURL  string          `json:"external_url,omitempty"`
	MediaURL     string          `json:"media_url,omitempty"`
	AnimationURL string          `json:"animation_url,omitempty"`
	Properties   json.RawMessage `json:"properties,omitempty"`
}
