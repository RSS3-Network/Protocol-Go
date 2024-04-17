package filter

// WorkerOnNetworksMap shows the corresponding Network that a Worker is on.
// https://github.com/RSS3-Network/Node/blob/develop/deploy/config.yaml
var WorkerOnNetworksMap = map[Worker][]string{
	Foundation: {
		NetworkArbitrum.String(),
		NetworkAvalanche.String(),
		NetworkBase.String(),
		NetworkBinanceSmartChain.String(),
		NetworkCrossbell.String(),
		NetworkEthereum.String(),
		NetworkGnosis.String(),
		NetworkLinea.String(),
		NetworkOptimism.String(),
		NetworkOptimism.String(),
		NetworkPolygon.String(),
		NetworkSatoshiVM.String(),
		NetworkVSL.String(),
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
		NetworkLinea.String(),
		NetworkSatoshiVM.String(),
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
		NetworkArbitrum.String(),
		NetworkEthereum.String(),
		NetworkOptimism.String(),
		NetworkPolygon.String(),
	},
	Aave: {
		NetworkArbitrum.String(),
		NetworkAvalanche.String(),
		NetworkBase.String(),
		NetworkEthereum.String(),
		NetworkOptimism.String(),
		NetworkPolygon.String(),
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

// NetworkHasWorkersMap shows the corresponding Worker that a Network has.
var NetworkHasWorkersMap = map[Network][]string{
	NetworkArbitrum: {
		Aave.String(),
		Foundation.String(),
		Highlight.String(),
	},
	NetworkArweave: {
		Mirror.String(),
		Momoka.String(),
		Paragraph.String(),
	},
	NetworkAvalanche: {
		Aave.String(),
		Foundation.String(),
	},
	NetworkBase: {
		Aave.String(),
		Foundation.String(),
	},
	NetworkBinanceSmartChain: {
		Foundation.String(),
	},
	NetworkCrossbell: {
		Crossbell.String(),
		Foundation.String(),
	},
	NetworkEthereum: {
		Aave.String(),
		ENS.String(),
		Foundation.String(),
		Highlight.String(),
		Lido.String(),
		Looksrare.String(),
		Oneinch.String(),
		OpenSea.String(),
		Optimism.String(),
		RSS3.String(),
		Uniswap.String(),
		VSL.String(),
	},
	NetworkFarcaster: {
		Farcaster.String(),
	},
	NetworkGnosis: {
		Foundation.String(),
	},
	NetworkLinea: {
		Foundation.String(),
		Uniswap.String(),
	},
	NetworkOptimism: {
		Aave.String(),
		Foundation.String(),
		Highlight.String(),
		KiwiStand.String(),
	},
	NetworkPolygon: {
		Aave.String(),
		Aavegotchi.String(),
		Foundation.String(),
		Highlight.String(),
		IQWiki.String(),
		Lens.String(),
		Matters.String(),
	},
	NetworkSatoshiVM: {
		Foundation.String(),
		SAVM.String(),
		Uniswap.String(),
	},
	NetworkVSL: {
		Foundation.String(),
		RSS3.String(),
	},
}

// PlatformHasWorkerMap is a map of platform to worker.
var PlatformHasWorkerMap = map[Platform]string{
	Platform1inch:      Oneinch.String(),
	PlatformAAVE:       Aave.String(),
	PlatformAavegotchi: Aavegotchi.String(),
	PlatformCrossbell:  Crossbell.String(),
	PlatformENS:        ENS.String(),
	PlatformFarcaster:  Farcaster.String(),
	PlatformHighlight:  Highlight.String(),
	PlatformIQWiki:     IQWiki.String(),
	PlatformKiwiStand:  KiwiStand.String(),
	PlatformLens:       Lens.String(),
	PlatformLido:       Lido.String(),
	PlatformLooksRare:  Looksrare.String(),
	PlatformMatters:    Matters.String(),
	PlatformMirror:     Mirror.String(),
	PlatformMomoka:     Momoka.String(),
	PlatformOpenSea:    OpenSea.String(),
	PlatformOptimism:   Optimism.String(),
	PlatformParagraph:  Paragraph.String(),
	PlatformRSS3:       RSS3.String(),
	PlatformSAVM:       SAVM.String(),
	PlatformUniswap:    Uniswap.String(),
	PlatformVSL:        VSL.String(),
}

var TagHasWorkersMap = map[Tag][]string{
	TagCollectible: {
		ENS.String(),
		Highlight.String(),
		KiwiStand.String(),
		Lido.String(),
		Looksrare.String(),
		OpenSea.String(),
	},
	TagExchange: {
		Aave.String(),
		Lido.String(),
		Oneinch.String(),
		RSS3.String(),
		Uniswap.String(),
	},
	TagSocial: {
		Crossbell.String(),
		ENS.String(),
		Farcaster.String(),
		IQWiki.String(),
		Lens.String(),
		Matters.String(),
		Mirror.String(),
		Momoka.String(),
		Paragraph.String(),
	},
	TagTransaction: {
		Optimism.String(),
		SAVM.String(),
		VSL.String(),
	},
	TagMetaverse: {
		Aavegotchi.String(),
	},
}
