package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/protocol-go/schema/filter"
)

type Metadata interface {
	Type() filter.Type
}

func Unmarshal(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	switch metadataType.Tag() {
	case filter.TagCollectible:
		return unmarshalCollectibleMetadata(metadataType, data)
	case filter.TagTransaction:
		return unmarshalTransactionMetadata(metadataType, data)
	case filter.TagSocial:
		return unmarshalSocialMetadata(metadataType, data)
	case filter.TagExchange:
		return unmarshalExchangeMetadata(metadataType, data)
	case filter.TagMetaverse:
		return unmarshalMetaverseMetadata(metadataType, data)
	case filter.TagRSS:
		return unmarshalRSSMetadata(metadataType, data)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}
}

func unmarshalCollectibleMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeCollectibleTransfer, filter.TypeCollectibleMint, filter.TypeCollectibleBurn:
		result = new(CollectibleTransfer)
	case filter.TypeCollectibleApproval:
		result = new(CollectibleApproval)
	case filter.TypeCollectibleTrade:
		result = new(CollectibleTrade)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalTransactionMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeTransactionApproval:
		result = new(TransactionApproval)
	case filter.TypeTransactionTransfer, filter.TypeTransactionMint, filter.TypeTransactionBurn:
		result = new(TransactionTransfer)
	case filter.TypeTransactionBridge:
		result = new(TransactionBridge)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalSocialMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeSocialPost, filter.TypeSocialRevise, filter.TypeSocialComment, filter.TypeSocialDelete, filter.TypeSocialShare, filter.TypeSocialReward, filter.TypeSocialMint:
		result = new(SocialPost)
	case filter.TypeSocialProfile:
		result = new(SocialProfile)
	case filter.TypeSocialProxy:
		result = new(SocialProxy)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalExchangeMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeExchangeSwap:
		result = new(ExchangeSwap)
	case filter.TypeExchangeLiquidity:
		result = new(ExchangeLiquidity)
	case filter.TypeExchangeStaking:
		result = new(ExchangeStaking)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalMetaverseMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeMetaverseMint, filter.TypeMetaverseBurn, filter.TypeMetaverseTransfer:
		result = new(MetaverseTransfer)
	case filter.TypeMetaverseTrade:
		result = new(MetaverseTrade)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalRSSMetadata(_ filter.Type, data json.RawMessage) (Metadata, error) {
	result := new(RSS)

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}
