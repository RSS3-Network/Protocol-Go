package metadata

import (
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
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

const (
	ActionExchangeLoanCreate ExchangeLoanAction = iota + 1
	ActionExchangeLoanRepay
	ActionExchangeLoanRefinance
	ActionExchangeLoanLiquidate
	ActionExchangeLoanSeize
)
