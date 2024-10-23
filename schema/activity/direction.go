package activity

import "github.com/labstack/echo/v4"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Direction --transform=snake --trimprefix=Direction --output direction_string.go --json
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.3 --type=Direction --transform=snake --trimprefix=Direction  --indent --example=in --output=direction_schema.json --x-go-type=activityx.Direction --x-go-type-import-path=github.com/rss3-network/protocol-go/schema/activity --x-go-type-import-name=activityx
type Direction uint64

const (
	DirectionIn Direction = iota + 1
	DirectionOut
	DirectionSelf
)

var _ echo.BindUnmarshaler = (*Direction)(nil)

func (d *Direction) UnmarshalParam(param string) error {
	direction, err := DirectionString(param)
	if err != nil {
		return err
	}

	*d = direction

	return nil
}
