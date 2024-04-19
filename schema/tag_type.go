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

// GetTypesByTag returns all Type that belong to a specific Tag.
func GetTypesByTag(tagValue tag.Tag) []Type {
	switch tagValue {
	case tag.Collectible:
		return convertToTypeSlice(typex.CollectibleTypeValues())
	case tag.Exchange:
		return convertToTypeSlice(typex.ExchangeTypeValues())
	case tag.Metaverse:
		return convertToTypeSlice(typex.MetaverseTypeValues())
	case tag.Social:
		return convertToTypeSlice(typex.SocialTypeValues())
	case tag.RSS:
		return convertToTypeSlice(typex.RSSTypeValues())
	case tag.Transaction:
		return convertToTypeSlice(typex.TransactionTypeValues())
	case tag.Unknown:
		return convertToTypeSlice(typex.UnknownTypeValues())
	default:
		return nil
	}
}

// Helper function to convert slices of specific types that implement Type into []Type.
func convertToTypeSlice[T Type](input []T) []Type {
	result := make([]Type, len(input))
	for i, v := range input {
		result[i] = v
	}

	return result
}
