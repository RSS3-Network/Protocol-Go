package schema

// this file is placed here to avoid circular import
import (
	"fmt"

	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
)

type Type interface {
	Name() string
	Tag() tag.Tag
}

func TypeString(parentTag tag.Tag, typeValue string) (Type, error) {
	switch parentTag {
	case tag.Collectible:
		return typex.CollectibleTypeString(typeValue)
	case tag.Exchange:
		return typex.ExchangeTypeString(typeValue)
	case tag.Metaverse:
		return typex.MetaverseTypeString(typeValue)
	case tag.Social:
		return typex.SocialTypeString(typeValue)
	case tag.RSS:
		return typex.RSSTypeString(typeValue)
	case tag.Transaction:
		return typex.TransactionTypeString(typeValue)
	case tag.Unknown:
		return typex.UnknownTypeString(typeValue)
	default:
		return nil, fmt.Errorf("unknown tag: %s", parentTag.String())
	}
}

func TagAndTypeString(tagValue string, typeValue string) (tag.Tag, Type, error) {
	_tag, err := tag.TagString(tagValue)
	if err != nil {
		return tag.Unknown, nil, err
	}

	_type, err := TypeString(_tag, typeValue)
	if err != nil {
		return tag.Unknown, nil, err
	}

	return _tag, _type, err
}
