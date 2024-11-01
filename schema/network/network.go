package network

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=Network --linecomment --output ../../openapi/enum/Network.yaml -t ../../openapi/tmpl/Network.yaml.tmpl
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
	Mastodon                         // mastodon
	Near                             // near
	Optimism                         // optimism
	Polygon                          // polygon
	RSSHub                           // rsshub
	SatoshiVM                        // savm
	VSL                              // vsl
	XLayer                           // x-layer
	Bluesky                          // bluesky
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

type Protocol string

// Open Data Protocols
const (
	ActivityPubProtocol Protocol = "activity-pub"
	ArweaveProtocol     Protocol = "arweave"
	EthereumProtocol    Protocol = "ethereum"
	FarcasterProtocol   Protocol = "farcaster"
	NearProtocol        Protocol = "near"
	RSSProtocol         Protocol = "rss"
	ATProtocol          Protocol = "atproto"
)

func (n Network) Protocol() Protocol {
	switch n {
	case Arweave:
		return ArweaveProtocol
	case Ethereum, Polygon, Optimism, Arbitrum, Base, Crossbell, Avalanche, VSL, SatoshiVM, BinanceSmartChain, Gnosis, Linea, XLayer:
		return EthereumProtocol
	case Farcaster:
		return FarcasterProtocol
	case Mastodon:
		return ActivityPubProtocol
	case Near:
		return NearProtocol
	case RSSHub:
		return RSSProtocol
	case Bluesky:
		return ATProtocol
	default:
		return ""
	}
}

func (s Protocol) Networks() []Network {
	switch s {
	case ActivityPubProtocol:
		return []Network{Mastodon}
	case ArweaveProtocol:
		return []Network{Arweave}
	case EthereumProtocol:
		return []Network{
			Ethereum,
			Polygon,
			Optimism,
			Arbitrum,
			Base,
			Crossbell,
			Avalanche,
			VSL,
			SatoshiVM,
			BinanceSmartChain,
			Gnosis,
			Linea,
			XLayer,
		}
	case FarcasterProtocol:
		return []Network{Farcaster}
	case NearProtocol:
		return []Network{Near}
	case RSSProtocol:
		return []Network{RSSHub}
	case ATProtocol:
		return []Network{Bluesky}
	default:
		return []Network{}
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
