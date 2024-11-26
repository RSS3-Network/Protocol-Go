package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
)

type Metadata interface {
	Type() schema.Type
}

func Unmarshal(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	switch metadataType.Tag() {
	case tag.Collectible:
		return unmarshalCollectibleMetadata(metadataType, data)
	case tag.Exchange:
		return unmarshalExchangeMetadata(metadataType, data)
	case tag.Governance:
		return unmarshalGovernanceMetadata(metadataType, data)
	case tag.Metaverse:
		return unmarshalMetaverseMetadata(metadataType, data)
	case tag.RSS:
		return unmarshalRSSMetadata(metadataType, data)
	case tag.Social:
		return unmarshalSocialMetadata(metadataType, data)
	case tag.Transaction:
		return unmarshalTransactionMetadata(metadataType, data)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}
}

func unmarshalCollectibleMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.CollectibleApproval:
		result = new(CollectibleApproval)
	case typex.CollectibleTrade:
		result = new(CollectibleTrade)
	case typex.CollectibleTransfer, typex.CollectibleMint, typex.CollectibleBurn:
		result = new(CollectibleTransfer)
	case typex.CollectibleAuction:
		result = new(CollectibleAuction)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalGovernanceMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.GovernanceProposal:
		result = new(GovernanceProposal)
	case typex.GovernanceVote:
		result = new(GovernanceVote)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalTransactionMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.TransactionApproval:
		result = new(TransactionApproval)
	case typex.TransactionBridge:
		result = new(TransactionBridge)
	case typex.TransactionBurn, typex.TransactionMint, typex.TransactionTransfer:
		result = new(TransactionTransfer)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalSocialMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.SocialComment, typex.SocialDelete, typex.SocialMint, typex.SocialPost, typex.SocialRevise, typex.SocialReward, typex.SocialShare, typex.SocialLike:
		result = new(SocialPost)
	case typex.SocialProfile:
		result = new(SocialProfile)
	case typex.SocialProxy:
		result = new(SocialProxy)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalExchangeMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.ExchangeLiquidity:
		result = new(ExchangeLiquidity)
	case typex.ExchangeStaking:
		result = new(ExchangeStaking)
	case typex.ExchangeSwap:
		result = new(ExchangeSwap)
	case typex.ExchangeLoan:
		result = new(ExchangeLoan)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalMetaverseMetadata(metadataType schema.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case typex.MetaverseBurn, typex.MetaverseMint, typex.MetaverseTransfer:
		result = new(MetaverseTransfer)
	case typex.MetaverseTrade:
		result = new(MetaverseTrade)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalRSSMetadata(_ schema.Type, data json.RawMessage) (Metadata, error) {
	result := new(RSS)

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}
