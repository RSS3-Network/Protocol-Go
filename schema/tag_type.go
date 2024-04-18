package schema

// this file is placed here to avoid circular import
import (
	"fmt"

	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
)

// Type has a name and a tag
type Type interface {
	Name() string
	Tag() tag.Tag
}

// ParseTypeFromString returns a Type based on its string value and its parent Tag
func ParseTypeFromString(parentTag tag.Tag, typeValue string) (Type, error) {
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

// ParseTagAndTypeFromString returns a Tag and Type based on their string values
func ParseTagAndTypeFromString(tagValue string, typeValue string) (tag.Tag, Type, error) {
	_tag, err := tag.TagString(tagValue)
	if err != nil {
		return tag.Unknown, nil, err
	}

	_type, err := ParseTypeFromString(_tag, typeValue)
	if err != nil {
		return tag.Unknown, nil, err
	}

	return _tag, _type, err
}
