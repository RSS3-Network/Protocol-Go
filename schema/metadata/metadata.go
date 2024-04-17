package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/tag"
	_type "github.com/rss3-network/protocol-go/schema/typex"
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
	case _type.CollectibleApproval:
		result = new(CollectibleApproval)
	case _type.CollectibleTrade:
		result = new(CollectibleTrade)
	case _type.CollectibleTransfer, _type.CollectibleMint, _type.CollectibleBurn:
		result = new(CollectibleTransfer)
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
	case _type.TransactionApproval:
		result = new(TransactionApproval)
	case _type.TransactionBridge:
		result = new(TransactionBridge)
	case _type.TransactionBurn, _type.TransactionMint, _type.TransactionTransfer:
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
	case _type.SocialComment, _type.SocialDelete, _type.SocialMint, _type.SocialPost, _type.SocialRevise, _type.SocialReward, _type.SocialShare:
		result = new(SocialPost)
	case _type.SocialProfile:
		result = new(SocialProfile)
	case _type.SocialProxy:
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
	case _type.ExchangeLiquidity:
		result = new(ExchangeLiquidity)
	case _type.ExchangeStaking:
		result = new(ExchangeStaking)
	case _type.ExchangeSwap:
		result = new(ExchangeSwap)
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
	case _type.MetaverseBurn, _type.MetaverseMint, _type.MetaverseTransfer:
		result = new(MetaverseTransfer)
	case _type.MetaverseTrade:
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
