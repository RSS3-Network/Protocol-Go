// Code generated by "enumer --values --type=Network --linecomment --output network_string.go --json --yaml --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NetworkName = "unknownarbitrumarweaveavaxbasebinance-smart-chainbitcoincrossbellethereumfarcastergnosislineaoptimismpolygonrsssavmvsl"

var _NetworkIndex = [...]uint8{0, 7, 15, 22, 26, 30, 49, 56, 65, 73, 82, 88, 93, 101, 108, 111, 115, 118}

const _NetworkLowerName = "unknownarbitrumarweaveavaxbasebinance-smart-chainbitcoincrossbellethereumfarcastergnosislineaoptimismpolygonrsssavmvsl"

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
	_ = x[NetworkUnknown-(0)]
	_ = x[NetworkArbitrum-(1)]
	_ = x[NetworkArweave-(2)]
	_ = x[NetworkAvalanche-(3)]
	_ = x[NetworkBase-(4)]
	_ = x[NetworkBinanceSmartChain-(5)]
	_ = x[NetworkBitcoin-(6)]
	_ = x[NetworkCrossbell-(7)]
	_ = x[NetworkEthereum-(8)]
	_ = x[NetworkFarcaster-(9)]
	_ = x[NetworkGnosis-(10)]
	_ = x[NetworkLinea-(11)]
	_ = x[NetworkOptimism-(12)]
	_ = x[NetworkPolygon-(13)]
	_ = x[NetworkRSS-(14)]
	_ = x[NetworkSatoshiVM-(15)]
	_ = x[NetworkVSL-(16)]
}

var _NetworkValues = []Network{NetworkUnknown, NetworkArbitrum, NetworkArweave, NetworkAvalanche, NetworkBase, NetworkBinanceSmartChain, NetworkBitcoin, NetworkCrossbell, NetworkEthereum, NetworkFarcaster, NetworkGnosis, NetworkLinea, NetworkOptimism, NetworkPolygon, NetworkRSS, NetworkSatoshiVM, NetworkVSL}

var _NetworkNameToValueMap = map[string]Network{
	_NetworkName[0:7]:          NetworkUnknown,
	_NetworkLowerName[0:7]:     NetworkUnknown,
	_NetworkName[7:15]:         NetworkArbitrum,
	_NetworkLowerName[7:15]:    NetworkArbitrum,
	_NetworkName[15:22]:        NetworkArweave,
	_NetworkLowerName[15:22]:   NetworkArweave,
	_NetworkName[22:26]:        NetworkAvalanche,
	_NetworkLowerName[22:26]:   NetworkAvalanche,
	_NetworkName[26:30]:        NetworkBase,
	_NetworkLowerName[26:30]:   NetworkBase,
	_NetworkName[30:49]:        NetworkBinanceSmartChain,
	_NetworkLowerName[30:49]:   NetworkBinanceSmartChain,
	_NetworkName[49:56]:        NetworkBitcoin,
	_NetworkLowerName[49:56]:   NetworkBitcoin,
	_NetworkName[56:65]:        NetworkCrossbell,
	_NetworkLowerName[56:65]:   NetworkCrossbell,
	_NetworkName[65:73]:        NetworkEthereum,
	_NetworkLowerName[65:73]:   NetworkEthereum,
	_NetworkName[73:82]:        NetworkFarcaster,
	_NetworkLowerName[73:82]:   NetworkFarcaster,
	_NetworkName[82:88]:        NetworkGnosis,
	_NetworkLowerName[82:88]:   NetworkGnosis,
	_NetworkName[88:93]:        NetworkLinea,
	_NetworkLowerName[88:93]:   NetworkLinea,
	_NetworkName[93:101]:       NetworkOptimism,
	_NetworkLowerName[93:101]:  NetworkOptimism,
	_NetworkName[101:108]:      NetworkPolygon,
	_NetworkLowerName[101:108]: NetworkPolygon,
	_NetworkName[108:111]:      NetworkRSS,
	_NetworkLowerName[108:111]: NetworkRSS,
	_NetworkName[111:115]:      NetworkSatoshiVM,
	_NetworkLowerName[111:115]: NetworkSatoshiVM,
	_NetworkName[115:118]:      NetworkVSL,
	_NetworkLowerName[115:118]: NetworkVSL,
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
	_NetworkName[101:108],
	_NetworkName[108:111],
	_NetworkName[111:115],
	_NetworkName[115:118],
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
