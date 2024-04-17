package filter

import (
	"strings"

	"github.com/labstack/echo/v4"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Platform --linecomment --output platform_string.go --json --sql
type Platform int

const (
	Platform1inch      Platform = iota + 1 // 1inch
	PlatformAAVE                           // AAVE
	PlatformAavegotchi                     // Aavegotchi
	PlatformCrossbell                      // Crossbell
	PlatformCurve                          // Curve
	PlatformENS                            // ENS
	PlatformFarcaster                      // Farcaster
	PlatformHighlight                      // Highlight
	PlatformIQWiki                         // IQWiki
	PlatformKiwiStand                      // KiwiStand
	PlatformLens                           // Lens
	PlatformLido                           // Lido
	PlatformLooksRare                      // LooksRare
	PlatformMatters                        // Matters
	PlatformMirror                         // Mirror
	PlatformMomoka                         // Momoka
	PlatformOpenSea                        // OpenSea
	PlatformOptimism                       // Optimism
	PlatformParagraph                      // Paragraph
	PlatformRSS3                           // RSS3
	PlatformSAVM                           // SAVM
	PlatformStargate                       // Stargate
	PlatformUniswap                        // Uniswap
	PlatformVSL                            // VSL
)

func (p Platform) ID() string {
	return strings.ReplaceAll(strings.ToLower(p.String()), "\x20", "_")
}

var _ echo.BindUnmarshaler = (*Platform)(nil)

func (p *Platform) UnmarshalParam(param string) error {
	platform, err := PlatformString(param)
	if err != nil {
		return err
	}

	*p = platform

	return nil
}
