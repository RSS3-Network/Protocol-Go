package network

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=EthereumChainID --linecomment --output ../../openapi/enum/EthereumChain.yaml -t ../../openapi/tmpl/Network.yaml.tmpl
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
	EthereumChainIDXLayer            EthereumChainID = 196   // x-layer
)

// ParseNetworkAndChainIDFromString returns the Network and its corresponding EthereumChainID based on a string network name.
func ParseNetworkAndChainIDFromString(networkString string) (Network, EthereumChainID) {
	network, err := NetworkString(networkString)
	if err != nil {
		return Unknown, 0
	}

	chainID, err := EthereumChainIDString(networkString)
	if err == nil {
		return network, chainID
	}

	return Unknown, 0
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
