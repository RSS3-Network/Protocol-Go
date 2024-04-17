package network

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql
type EthereumChainID uint64

const (
	EthereumChainIDArbitrum          EthereumChainID = 42161 // arbitrum
	EthereumChainIDAvalanche         EthereumChainID = 43114 // avax
	EthereumChainIDBase              EthereumChainID = 8453  // base
	EthereumChainIDBinanceSmartChain EthereumChainID = 56    // binance-smart-chain
	EthereumChainIDCrossbell         EthereumChainID = 3737  // crossbell
	EthereumChainIDGnosis            EthereumChainID = 100   // gnosis
	EthereumChainIDLinea             EthereumChainID = 59144 // linea
	EthereumChainIDMainnet           EthereumChainID = 1     // ethereum
	EthereumChainIDOptimism          EthereumChainID = 10    // optimism
	EthereumChainIDPolygon           EthereumChainID = 137   // polygon
	EthereumChainIDSatoshiVM         EthereumChainID = 3109  // savm
	EthereumChainIDVSL               EthereumChainID = 12553 // vsl
)

func NameAndChainID(network string) (Network, EthereumChainID) {
	switch network {
	case Arbitrum.String():
		return Arbitrum, EthereumChainIDArbitrum
	case Avalanche.String():
		return Avalanche, EthereumChainIDAvalanche
	case Base.String():
		return Base, EthereumChainIDBase
	case BinanceSmartChain.String():
		return BinanceSmartChain, EthereumChainIDBinanceSmartChain
	case Crossbell.String():
		return Crossbell, EthereumChainIDCrossbell
	case Ethereum.String():
		return Ethereum, EthereumChainIDMainnet
	case Gnosis.String():
		return Gnosis, EthereumChainIDGnosis
	case Linea.String():
		return Linea, EthereumChainIDLinea
	case Optimism.String():
		return Optimism, EthereumChainIDOptimism
	case Polygon.String():
		return Polygon, EthereumChainIDPolygon
	case SatoshiVM.String():
		return SatoshiVM, EthereumChainIDSatoshiVM
	case VSL.String():
		return VSL, EthereumChainIDVSL
	default:
		return Unknown, 0
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
