package filter

// WorkerToNetworksMap is a map of worker to networks.
// https://github.com/RSS3-Network/Node/blob/develop/deploy/config.yaml
var WorkerToNetworksMap = map[Name][]string{
	Fallback: {
		NetworkEthereum.String(),
		NetworkOptimism.String(),
		NetworkPolygon.String(),
		NetworkCrossbell.String(),
		NetworkAvalanche.String(),
		NetworkBase.String(),
		NetworkOptimism.String(),
		NetworkArbitrum.String(),
		NetworkBinanceSmartChain.String(),
		NetworkVSL.String(),
		NetworkSatoshiVM.String(),
		NetworkGnosis.String(),
		NetworkLinea.String(),
	},
	Mirror: {
		NetworkArweave.String(),
	},
	Farcaster: {
		NetworkFarcaster.String(),
	},
	RSS3: {
		NetworkEthereum.String(),
		NetworkVSL.String(),
	},
	Paragraph: {
		NetworkArweave.String(),
	},
	OpenSea: {
		NetworkEthereum.String(),
	},
	Uniswap: {
		NetworkEthereum.String(),
		NetworkSatoshiVM.String(),
		NetworkLinea.String(),
	},
	Optimism: {
		NetworkEthereum.String(),
	},
	Aavegotchi: {
		NetworkPolygon.String(),
	},
	Lens: {
		NetworkPolygon.String(),
	},
	Looksrare: {
		NetworkEthereum.String(),
	},
	Matters: {
		NetworkPolygon.String(),
	},
	Momoka: {
		NetworkArweave.String(),
	},
	Highlight: {
		NetworkEthereum.String(),
		NetworkArbitrum.String(),
		NetworkPolygon.String(),
		NetworkOptimism.String(),
	},
	Aave: {
		NetworkPolygon.String(),
		NetworkEthereum.String(),
		NetworkAvalanche.String(),
		NetworkBase.String(),
		NetworkOptimism.String(),
		NetworkArbitrum.String(),
	},
	IQWiki: {
		NetworkPolygon.String(),
	},
	Lido: {
		NetworkEthereum.String(),
	},
	Crossbell: {
		NetworkCrossbell.String(),
	},
	ENS: {
		NetworkEthereum.String(),
	},
	KiwiStand: {
		NetworkEthereum.String(),
	},
	Oneinch: {
		NetworkEthereum.String(),
	},
	VSL: {
		NetworkEthereum.String(),
	},
	SAVM: {
		NetworkSatoshiVM.String(),
	},
}

// NetworkToWorkersMap is a map of network to workers.
var NetworkToWorkersMap = map[Network][]string{
	NetworkEthereum: {
		Fallback.String(),
		RSS3.String(),
		OpenSea.String(),
		Uniswap.String(),
		Optimism.String(),
		Looksrare.String(),
		Highlight.String(),
		Aave.String(),
		Lido.String(),
		ENS.String(),
		Oneinch.String(),
		VSL.String(),
	},
	NetworkArweave: {
		Mirror.String(),
		Paragraph.String(),
		Momoka.String(),
	},
	NetworkFarcaster: {
		Farcaster.String(),
	},
	NetworkPolygon: {
		Fallback.String(),
		Aavegotchi.String(),
		Lens.String(),
		Matters.String(),
		Aave.String(),
		IQWiki.String(),
		Highlight.String(),
	},
	NetworkCrossbell: {
		Fallback.String(),
		Crossbell.String(),
	},
	NetworkAvalanche: {
		Fallback.String(),
		Aave.String(),
	},
	NetworkBase: {
		Fallback.String(),
		Aave.String(),
	},
	NetworkOptimism: {
		Fallback.String(),
		Aave.String(),
		Highlight.String(),
		KiwiStand.String(),
	},
	NetworkArbitrum: {
		Fallback.String(),
		Aave.String(),
		Highlight.String(),
	},
	NetworkGnosis: {
		Fallback.String(),
	},
	NetworkLinea: {
		Fallback.String(),
		Uniswap.String(),
	},
	NetworkVSL: {
		Fallback.String(),
		RSS3.String(),
	},
	NetworkSatoshiVM: {
		Fallback.String(),
		Uniswap.String(),
		SAVM.String(),
	},
	NetworkBinanceSmartChain: {
		Fallback.String(),
	},
}

// PlatformToWorkerMap is a map of platform to worker.
var PlatformToWorkerMap = map[Platform]string{
	PlatformRSS3:       RSS3.String(),
	PlatformMirror:     Mirror.String(),
	PlatformFarcaster:  Farcaster.String(),
	PlatformParagraph:  Paragraph.String(),
	PlatformOpenSea:    OpenSea.String(),
	PlatformUniswap:    Uniswap.String(),
	PlatformOptimism:   Optimism.String(),
	PlatformAavegotchi: Aavegotchi.String(),
	PlatformLens:       Lens.String(),
	PlatformLooksRare:  Looksrare.String(),
	PlatformMatters:    Matters.String(),
	PlatformMomoka:     Momoka.String(),
	PlatformHighlight:  Highlight.String(),
	PlatformAAVE:       Aave.String(),
	PlatformIQWiki:     IQWiki.String(),
	PlatformLido:       Lido.String(),
	PlatformCrossbell:  Crossbell.String(),
	PlatformENS:        ENS.String(),
	PlatformKiwiStand:  KiwiStand.String(),
	Platform1inch:      Oneinch.String(),
	PlatformVSL:        VSL.String(),
	PlatformSAVM:       SAVM.String(),
}

var TagToWorkersMap = map[Tag][]string{
	TagTransaction: {
		Optimism.String(),
		VSL.String(),
		SAVM.String(),
	},
	TagCollectible: {
		OpenSea.String(),
		ENS.String(),
		Highlight.String(),
		Lido.String(),
		Looksrare.String(),
		KiwiStand.String(),
	},
	TagExchange: {
		RSS3.String(),
		Uniswap.String(),
		Aave.String(),
		Lido.String(),
		Oneinch.String(),
	},
	TagSocial: {
		Farcaster.String(),
		Mirror.String(),
		Lens.String(),
		Paragraph.String(),
		Crossbell.String(),
		ENS.String(),
		IQWiki.String(),
		Matters.String(),
		Momoka.String(),
	},
	TagMetaverse: {
		Aavegotchi.String(),
	},
}
