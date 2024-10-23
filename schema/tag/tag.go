package tag

import (
	"github.com/labstack/echo/v4"
)

// Tag represents a metadata tag.
//
//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Tag --transform=snake --trimprefix=Tag --output tag_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=Tag --transform=snake --trimprefix=Tag --indent --example=exchange --output tag_schema.json --x-go-type=tag.Tag --x-go-type-import-path=github.com/rss3-network/protocol-go/schema/tag --x-go-type-skip-optional-pointer
type Tag uint64

const (
	Unknown Tag = iota
	Collectible
	Exchange
	Governance
	Metaverse
	RSS
	Social
	Transaction
)

var _ echo.BindUnmarshaler = (*Tag)(nil)

func (t *Tag) UnmarshalParam(param string) error {
	tag, err := TagString(param)
	if err != nil {
		return err
	}

	*t = tag

	return nil
}
