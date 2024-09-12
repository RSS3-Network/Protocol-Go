package metadata

import (
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

var _ Metadata = (*ExchangeSwap)(nil)

type ExchangeSwap struct {
	From Token `json:"from"`
	To   Token `json:"to"`
}

func (e ExchangeSwap) Type() schema.Type {
	return typex.ExchangeSwap
}

var _ Metadata = (*ExchangeLiquidity)(nil)

type ExchangeLiquidity struct {
	Action ExchangeLiquidityAction `json:"action"`
	Tokens []Token                 `json:"tokens"`
}

func (e ExchangeLiquidity) Type() schema.Type {
	return typex.ExchangeSwap
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeLiquidityAction --transform=snake --trimprefix=ActionExchangeLiquidity --output exchange_liquidity.go --json --sql
type ExchangeLiquidityAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeLiquidityAction) Type() schema.Type {
	return typex.ExchangeLiquidity
}

const (
	ActionExchangeLiquidityAdd ExchangeLiquidityAction = iota + 1
	ActionExchangeLiquidityBorrow
	ActionExchangeLiquidityCollect
	ActionExchangeLiquidityRemove
	ActionExchangeLiquidityRepay
	ActionExchangeLiquiditySupply
	ActionExchangeLiquidityWithdraw
)

var _ Metadata = (*ExchangeStaking)(nil)

type ExchangeStaking struct {
	Action ExchangeStakingAction  `json:"action"`
	Token  Token                  `json:"token"`
	Period *ExchangeStakingPeriod `json:"period,omitempty"`
}

func (e ExchangeStaking) Type() schema.Type {
	return typex.ExchangeStaking
}

type ExchangeStakingPeriod struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeStakingAction --transform=snake --trimprefix=ActionExchangeStaking --output exchange_staking.go --json --sql
type ExchangeStakingAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeStakingAction) Type() schema.Type {
	return typex.ExchangeStaking
}

const (
	ActionExchangeStakingStake ExchangeStakingAction = iota + 1
	ActionExchangeStakingUnstake
	ActionExchangeStakingClaim
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeLoanAction --transform=snake --trimprefix=ActionExchangeLoan --output exchange_loan.go --json --sql
type ExchangeLoanAction uint64

var _ Metadata = (*ExchangeLoan)(nil)

type ExchangeLoan struct {
	Action     ExchangeLoanAction `json:"action"`
	Collateral Token              `json:"collateral"`
	Amount     *Token             `json:"amount,omitempty"`
}

func (e ExchangeLoan) Type() schema.Type {
	return typex.ExchangeLoan
}

//goland:noinspection GoMixedReceiverTypes
func (e ExchangeLoanAction) Type() schema.Type {
	return typex.ExchangeLoan
}

const (
	ActionExchangeLoanCreate ExchangeLoanAction = iota + 1
	ActionExchangeLoanRepay
	ActionExchangeLoanRefinance
	ActionExchangeLoanLiquidate
	ActionExchangeLoanSeize
)

var _ Metadata = (*ExchangeTrade)(nil)

type ExchangeTrade struct {
	Action            ExchangeTradeAction `json:"action"`
	OrderHash         string              `json:"order_hash"`
	Maker             string              `json:"maker"`
	Taker             string              `json:"taker"`
	MakerAssetID      uint64              `json:"maker_asset_id"`
	TakerAssetID      uint64              `json:"taker_asset_id"`
	MakerAmountFilled *decimal.Decimal    `json:"maker_amount_filled"`
	TakerAmountFilled *decimal.Decimal    `json:"taker_amount_filled"`
	Fee               *decimal.Decimal    `json:"fee,omitempty"`
}

func (e ExchangeTrade) Type() schema.Type {
	return typex.ExchangeTrade
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeTradeAction --transform=snake --trimprefix=ActionExchangeTrade --output exchange_trade.go --json --sql
type ExchangeTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeTradeAction) Type() schema.Type {
	return typex.ExchangeTrade
}

const (
	ActionExchangeTradeFinalized ExchangeTradeAction = iota + 1
	ActionExchangeTradeMatched
)
