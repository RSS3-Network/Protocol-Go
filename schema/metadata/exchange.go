package metadata

import (
	"time"

	"github.com/rss3-network/protocol-go/schema"
	_type "github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*ExchangeSwap)(nil)

type ExchangeSwap struct {
	From Token `json:"from"`
	To   Token `json:"to"`
}

func (e ExchangeSwap) Type() schema.Type {
	return _type.ExchangeSwap
}

var _ Metadata = (*ExchangeLiquidity)(nil)

type ExchangeLiquidity struct {
	Action ExchangeLiquidityAction `json:"action"`
	Tokens []Token                 `json:"tokens"`
}

func (e ExchangeLiquidity) Type() schema.Type {
	return _type.ExchangeSwap
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeLiquidityAction --transform=snake --trimprefix=ActionExchangeLiquidity --output exchange_liquidity.go --json --sql
type ExchangeLiquidityAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeLiquidityAction) Type() schema.Type {
	return _type.ExchangeLiquidity
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
	return _type.ExchangeStaking
}

type ExchangeStakingPeriod struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeStakingAction --transform=snake --trimprefix=ActionExchangeStaking --output exchange_staking.go --json --sql
type ExchangeStakingAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeStakingAction) Type() schema.Type {
	return _type.ExchangeStaking
}

const (
	ActionExchangeStakingStake ExchangeStakingAction = iota + 1
	ActionExchangeStakingUnstake
	ActionExchangeStakingClaim
)
