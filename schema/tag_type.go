package schema

import (
	"fmt"

	"github.com/rss3-network/protocol-go/schema/tag"
	_type "github.com/rss3-network/protocol-go/schema/typex"
)

type Type interface {
	Name() string
	Tag() tag.Tag
}

func TypeString(parentTag tag.Tag, typeValue string) (Type, error) {
	switch parentTag {
	case tag.Collectible:
		return _type.CollectibleTypeString(typeValue)
	case tag.Exchange:
		return _type.ExchangeTypeString(typeValue)
	case tag.Metaverse:
		return _type.MetaverseTypeString(typeValue)
	case tag.Social:
		return _type.SocialTypeString(typeValue)
	case tag.RSS:
		return _type.RSSTypeString(typeValue)
	case tag.Transaction:
		return _type.TransactionTypeString(typeValue)
	case tag.Unknown:
		return _type.UnknownTypeString(typeValue)
	default:
		return nil, fmt.Errorf("unknown tag: %s", parentTag.String())
	}
}
