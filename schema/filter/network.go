package filter

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
type Network uint64

const (
	NetworkUnknown           Network = iota // unknown
	NetworkEthereum                         // ethereum
	NetworkOptimism                         // optimism
	NetworkBase                             // base
	NetworkPolygon                          // polygon
	NetworkCrossbell                        // crossbell
	NetworkArbitrum                         // arbitrum
	NetworkRSS                              // rss
	NetworkArweave                          // arweave
	NetworkFarcaster                        // farcaster
	NetworkAvalanche                        // avax
	NetworkVSL                              // vsl
	NetworkSatoshiVM                        // savm
	NetworkBTC                              // btc
	NetworkBinanceSmartChain                // binance-smart-chain
	NetworkGnosis                           // gnosis
	NetworkLinea                            // linea
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

type NetworkSource string

const (
	NetworkEthereumSource  NetworkSource = "ethereum"
	NetworkArweaveSource   NetworkSource = "arweave"
	NetworkFarcasterSource NetworkSource = "farcaster"
)

func (n Network) Source() NetworkSource {
	switch n {
	case NetworkEthereum, NetworkPolygon, NetworkOptimism, NetworkArbitrum, NetworkBase, NetworkCrossbell, NetworkAvalanche, NetworkVSL, NetworkSatoshiVM, NetworkBinanceSmartChain, NetworkGnosis, NetworkLinea:
		return NetworkEthereumSource
	case NetworkArweave:
		return NetworkArweaveSource
	case NetworkFarcaster:
		return NetworkFarcasterSource
	default:
		return ""
	}
}

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql
type EthereumChainID uint64

const (
	EthereumChainIDMainnet           EthereumChainID = 1     // ethereum
	EthereumChainIDOptimism          EthereumChainID = 10    // optimism
	EthereumChainIDPolygon           EthereumChainID = 137   // polygon
	EthereumChainIDArbitrum          EthereumChainID = 42161 // arbitrum
	EthereumChainIDBase              EthereumChainID = 8453  // base
	EthereumChainIDCrossbell         EthereumChainID = 3737  // crossbell
	EthereumChainIDAvalanche         EthereumChainID = 43114 // avax
	EthereumChainIDVSL               EthereumChainID = 12553 // vsl
	EthereumChainIDSatoshiVM         EthereumChainID = 3109  // savm
	EthereumChainIDBinanceSmartChain EthereumChainID = 56    // binance_smart_chain
	EthereumChainIDGnosis            EthereumChainID = 100   // gnosis
	EthereumChainIDLinea             EthereumChainID = 59144
)

func NetworkAndChainID(network string) (Network, EthereumChainID) {
	switch network {
	case NetworkEthereum.String():
		return NetworkEthereum, EthereumChainIDMainnet
	case NetworkOptimism.String():
		return NetworkOptimism, EthereumChainIDOptimism
	case NetworkPolygon.String():
		return NetworkPolygon, EthereumChainIDPolygon
	case NetworkArbitrum.String():
		return NetworkArbitrum, EthereumChainIDArbitrum
	case NetworkBase.String():
		return NetworkBase, EthereumChainIDBase
	case NetworkCrossbell.String():
		return NetworkCrossbell, EthereumChainIDCrossbell
	case NetworkAvalanche.String():
		return NetworkAvalanche, EthereumChainIDAvalanche
	case NetworkVSL.String():
		return NetworkVSL, EthereumChainIDVSL
	case NetworkSatoshiVM.String():
		return NetworkSatoshiVM, EthereumChainIDSatoshiVM
	case NetworkBinanceSmartChain.String():
		return NetworkBinanceSmartChain, EthereumChainIDBinanceSmartChain
	case NetworkGnosis.String():
		return NetworkGnosis, EthereumChainIDGnosis
	case NetworkLinea.String():
		return NetworkLinea, EthereumChainIDLinea
	default:
		return NetworkUnknown, 0
	}
}

func IsOptimismSuperchain(chainID uint64) bool {
	switch chainID {
	case uint64(EthereumChainIDOptimism),
		uint64(EthereumChainIDBase), uint64(EthereumChainIDVSL):
		return true
	default:
		return false
	}
}

func NetworkHookFunc() mapstructure.DecodeHookFuncType {
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
