package filter

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// Worker represents a worker name.
//
//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql
type Worker uint64

const (
	Unknown    Worker = iota // unknown
	Aave                     // aave
	Aavegotchi               // aavegotchi
	Crossbell                // crossbell
	Curve                    // curve
	ENS                      // ens
	Farcaster                // farcaster
	// Foundation serves as the catch-all worker for all Network.
	Foundation // foundation
	Highlight  // highlight
	IQWiki     // iqwiki
	KiwiStand  // kiwistand
	Lens       // lens
	Lido       // lido
	Looksrare  // looksrare
	Matters    // matters
	Mirror     // mirror
	Momoka     // momoka
	Oneinch    // 1inch
	OpenSea    // opensea
	Optimism   // optimism
	Paragraph  // paragraph
	RSS3       // rss3
	SAVM       // savm
	Stargate   // stargate
	Uniswap    // uniswap
	VSL        // vsl
)

func WorkerHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.Int {
			return data, nil
		}

		return _WorkerNameToValueMap[data.(string)], nil
	}
}
