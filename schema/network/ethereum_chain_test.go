package network

import (
	"testing"
)

func TestParseNetworkAndChainIDFromString(t *testing.T) {
	testCases := []struct {
		name            string
		networkString   string
		expectedNet     Network
		expectedChainID EthereumChainID
	}{
		{
			name:            "Mainnet",
			networkString:   "ethereum",
			expectedNet:     Ethereum,
			expectedChainID: EthereumChainIDMainnet,
		},
		{
			name:            "Arbitrum",
			networkString:   "arbitrum",
			expectedNet:     Arbitrum,
			expectedChainID: EthereumChainIDArbitrum,
		},
		{
			name:            "Unknown",
			networkString:   "unknown",
			expectedNet:     Unknown,
			expectedChainID: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			network, chainID := ParseNetworkAndChainIDFromString(tc.networkString)
			if network != tc.expectedNet {
				t.Errorf("Expected network %s, got %s", tc.expectedNet, network)
			}
			if chainID != tc.expectedChainID {
				t.Errorf("Expected chainID %d, got %d", tc.expectedChainID, chainID)
			}
		})
	}
}
