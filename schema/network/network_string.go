// Code generated by "enumer --values --type=Network --linecomment --output network_string.go --json --yaml --sql"; DO NOT EDIT.

package network

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NetworkName = "unknownarbitrumarweaveavaxbasebinance-smart-chainbitcoincrossbellethereumfarcastergnosislineamastodonnearoptimismpolygonrsshubsavmvslx-layer"

var _NetworkIndex = [...]uint8{0, 7, 15, 22, 26, 30, 49, 56, 65, 73, 82, 88, 93, 101, 105, 113, 120, 126, 130, 133, 140}

const _NetworkLowerName = "unknownarbitrumarweaveavaxbasebinance-smart-chainbitcoincrossbellethereumfarcastergnosislineamastodonnearoptimismpolygonrsshubsavmvslx-layer"

func (i Network) String() string {
	if i >= Network(len(_NetworkIndex)-1) {
		return fmt.Sprintf("Network(%d)", i)
	}
	return _NetworkName[_NetworkIndex[i]:_NetworkIndex[i+1]]
}

func (Network) Values() []string {
	return NetworkStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NetworkNoOp() {
	var x [1]struct{}
	_ = x[Unknown-(0)]
	_ = x[Arbitrum-(1)]
	_ = x[Arweave-(2)]
	_ = x[Avalanche-(3)]
	_ = x[Base-(4)]
	_ = x[BinanceSmartChain-(5)]
	_ = x[Bitcoin-(6)]
	_ = x[Crossbell-(7)]
	_ = x[Ethereum-(8)]
	_ = x[Farcaster-(9)]
	_ = x[Gnosis-(10)]
	_ = x[Linea-(11)]
	_ = x[Mastodon-(12)]
	_ = x[Near-(13)]
	_ = x[Optimism-(14)]
	_ = x[Polygon-(15)]
	_ = x[RSSHub-(16)]
	_ = x[SatoshiVM-(17)]
	_ = x[VSL-(18)]
	_ = x[XLayer-(19)]
}

var _NetworkValues = []Network{Unknown, Arbitrum, Arweave, Avalanche, Base, BinanceSmartChain, Bitcoin, Crossbell, Ethereum, Farcaster, Gnosis, Linea, Mastodon, Near, Optimism, Polygon, RSSHub, SatoshiVM, VSL, XLayer}

var _NetworkNameToValueMap = map[string]Network{
	_NetworkName[0:7]:          Unknown,
	_NetworkLowerName[0:7]:     Unknown,
	_NetworkName[7:15]:         Arbitrum,
	_NetworkLowerName[7:15]:    Arbitrum,
	_NetworkName[15:22]:        Arweave,
	_NetworkLowerName[15:22]:   Arweave,
	_NetworkName[22:26]:        Avalanche,
	_NetworkLowerName[22:26]:   Avalanche,
	_NetworkName[26:30]:        Base,
	_NetworkLowerName[26:30]:   Base,
	_NetworkName[30:49]:        BinanceSmartChain,
	_NetworkLowerName[30:49]:   BinanceSmartChain,
	_NetworkName[49:56]:        Bitcoin,
	_NetworkLowerName[49:56]:   Bitcoin,
	_NetworkName[56:65]:        Crossbell,
	_NetworkLowerName[56:65]:   Crossbell,
	_NetworkName[65:73]:        Ethereum,
	_NetworkLowerName[65:73]:   Ethereum,
	_NetworkName[73:82]:        Farcaster,
	_NetworkLowerName[73:82]:   Farcaster,
	_NetworkName[82:88]:        Gnosis,
	_NetworkLowerName[82:88]:   Gnosis,
	_NetworkName[88:93]:        Linea,
	_NetworkLowerName[88:93]:   Linea,
	_NetworkName[93:101]:       Mastodon,
	_NetworkLowerName[93:101]:  Mastodon,
	_NetworkName[101:105]:      Near,
	_NetworkLowerName[101:105]: Near,
	_NetworkName[105:113]:      Optimism,
	_NetworkLowerName[105:113]: Optimism,
	_NetworkName[113:120]:      Polygon,
	_NetworkLowerName[113:120]: Polygon,
	_NetworkName[120:126]:      RSSHub,
	_NetworkLowerName[120:126]: RSSHub,
	_NetworkName[126:130]:      SatoshiVM,
	_NetworkLowerName[126:130]: SatoshiVM,
	_NetworkName[130:133]:      VSL,
	_NetworkLowerName[130:133]: VSL,
	_NetworkName[133:140]:      XLayer,
	_NetworkLowerName[133:140]: XLayer,
}

var _NetworkNames = []string{
	_NetworkName[0:7],
	_NetworkName[7:15],
	_NetworkName[15:22],
	_NetworkName[22:26],
	_NetworkName[26:30],
	_NetworkName[30:49],
	_NetworkName[49:56],
	_NetworkName[56:65],
	_NetworkName[65:73],
	_NetworkName[73:82],
	_NetworkName[82:88],
	_NetworkName[88:93],
	_NetworkName[93:101],
	_NetworkName[101:105],
	_NetworkName[105:113],
	_NetworkName[113:120],
	_NetworkName[120:126],
	_NetworkName[126:130],
	_NetworkName[130:133],
	_NetworkName[133:140],
}

// NetworkString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NetworkString(s string) (Network, error) {
	if val, ok := _NetworkNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NetworkNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Network values", s)
}

// NetworkValues returns all values of the enum
func NetworkValues() []Network {
	return _NetworkValues
}

// NetworkStrings returns a slice of all String values of the enum
func NetworkStrings() []string {
	strs := make([]string, len(_NetworkNames))
	copy(strs, _NetworkNames)
	return strs
}

// IsANetwork returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Network) IsANetwork() bool {
	for _, v := range _NetworkValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Network
func (i Network) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Network
func (i *Network) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Network should be a string, got %s", data)
	}

	var err error
	*i, err = NetworkString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Network
func (i Network) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Network
func (i *Network) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NetworkString(s)
	return err
}

func (i Network) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Network) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of Network: %[1]T(%[1]v)", value)
	}

	val, err := NetworkString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
