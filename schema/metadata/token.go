package metadata

import (
	"github.com/shopspring/decimal"
)

type Token struct {
	Address  *string          `json:"address,omitempty"`
	ID       *decimal.Decimal `json:"id,omitempty"`
	Value    *decimal.Decimal `json:"value,omitempty"`
	Name     string           `json:"name,omitempty"`
	Symbol   string           `json:"symbol,omitempty"`
	Decimals uint8            `json:"decimals,omitempty"`
	Standard Standard         `json:"standard,omitempty"`
}
