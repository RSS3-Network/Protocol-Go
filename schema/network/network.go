package network

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
type Network uint64

const (
	Unknown           Network = iota // unknown
	Arbitrum                         // arbitrum
	Arweave                          // arweave
	Avalanche                        // avax
	Base                             // base
	BinanceSmartChain                // binance-smart-chain
	Bitcoin                          // bitcoin
	Crossbell                        // crossbell
	Ethereum                         // ethereum
	Farcaster                        // farcaster
	Gnosis                           // gnosis
	Linea                            // linea
	Optimism                         // optimism
	Polygon                          // polygon
	RSS                              // rss
	SatoshiVM                        // savm
	VSL                              // vsl
)

var _ echo.BindUnmarshaler = (*Network)(nil)

func (n *Network) UnmarshalParam(param string) error {
	network, err := NetworkString(param)
	if err != nil {
		return err
	}

	*n = network

	return nil
}

type Source string

const (
	ArweaveSource   Source = "arweave"
	EthereumSource  Source = "ethereum"
	FarcasterSource Source = "farcaster"
)

func (n Network) Source() Source {
	switch n {
	case Arweave:
		return ArweaveSource
	case Ethereum, Polygon, Optimism, Arbitrum, Base, Crossbell, Avalanche, VSL, SatoshiVM, BinanceSmartChain, Gnosis, Linea:
		return EthereumSource
	case Farcaster:
		return FarcasterSource
	default:
		return ""
	}
}

func HookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.Uint64 {
			return data, nil
		}

		return _NetworkNameToValueMap[data.(string)], nil
	}
}
