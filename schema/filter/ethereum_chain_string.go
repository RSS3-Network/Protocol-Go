// Code generated by "enumer --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _EthereumChainIDName = "ethereumoptimismbinance-smart-chaingnosispolygonsavmcrossbellbasevslarbitrumavaxlinea"
const _EthereumChainIDLowerName = "ethereumoptimismbinance-smart-chaingnosispolygonsavmcrossbellbasevslarbitrumavaxlinea"

var _EthereumChainIDMap = map[EthereumChainID]string{
	1:     _EthereumChainIDName[0:8],
	10:    _EthereumChainIDName[8:16],
	56:    _EthereumChainIDName[16:35],
	100:   _EthereumChainIDName[35:41],
	137:   _EthereumChainIDName[41:48],
	3109:  _EthereumChainIDName[48:52],
	3737:  _EthereumChainIDName[52:61],
	8453:  _EthereumChainIDName[61:65],
	12553: _EthereumChainIDName[65:68],
	42161: _EthereumChainIDName[68:76],
	43114: _EthereumChainIDName[76:80],
	59144: _EthereumChainIDName[80:85],
}

func (i EthereumChainID) String() string {
	if str, ok := _EthereumChainIDMap[i]; ok {
		return str
	}
	return fmt.Sprintf("EthereumChainID(%d)", i)
}

func (EthereumChainID) Values() []string {
	return EthereumChainIDStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EthereumChainIDNoOp() {
	var x [1]struct{}
	_ = x[EthereumChainIDMainnet-(1)]
	_ = x[EthereumChainIDOptimism-(10)]
	_ = x[EthereumChainIDBinanceSmartChain-(56)]
	_ = x[EthereumChainIDGnosis-(100)]
	_ = x[EthereumChainIDPolygon-(137)]
	_ = x[EthereumChainIDSatoshiVM-(3109)]
	_ = x[EthereumChainIDCrossbell-(3737)]
	_ = x[EthereumChainIDBase-(8453)]
	_ = x[EthereumChainIDVSL-(12553)]
	_ = x[EthereumChainIDArbitrum-(42161)]
	_ = x[EthereumChainIDAvalanche-(43114)]
	_ = x[EthereumChainIDLinea-(59144)]
}

var _EthereumChainIDValues = []EthereumChainID{EthereumChainIDMainnet, EthereumChainIDOptimism, EthereumChainIDBinanceSmartChain, EthereumChainIDGnosis, EthereumChainIDPolygon, EthereumChainIDSatoshiVM, EthereumChainIDCrossbell, EthereumChainIDBase, EthereumChainIDVSL, EthereumChainIDArbitrum, EthereumChainIDAvalanche, EthereumChainIDLinea}

var _EthereumChainIDNameToValueMap = map[string]EthereumChainID{
	_EthereumChainIDName[0:8]:        EthereumChainIDMainnet,
	_EthereumChainIDLowerName[0:8]:   EthereumChainIDMainnet,
	_EthereumChainIDName[8:16]:       EthereumChainIDOptimism,
	_EthereumChainIDLowerName[8:16]:  EthereumChainIDOptimism,
	_EthereumChainIDName[16:35]:      EthereumChainIDBinanceSmartChain,
	_EthereumChainIDLowerName[16:35]: EthereumChainIDBinanceSmartChain,
	_EthereumChainIDName[35:41]:      EthereumChainIDGnosis,
	_EthereumChainIDLowerName[35:41]: EthereumChainIDGnosis,
	_EthereumChainIDName[41:48]:      EthereumChainIDPolygon,
	_EthereumChainIDLowerName[41:48]: EthereumChainIDPolygon,
	_EthereumChainIDName[48:52]:      EthereumChainIDSatoshiVM,
	_EthereumChainIDLowerName[48:52]: EthereumChainIDSatoshiVM,
	_EthereumChainIDName[52:61]:      EthereumChainIDCrossbell,
	_EthereumChainIDLowerName[52:61]: EthereumChainIDCrossbell,
	_EthereumChainIDName[61:65]:      EthereumChainIDBase,
	_EthereumChainIDLowerName[61:65]: EthereumChainIDBase,
	_EthereumChainIDName[65:68]:      EthereumChainIDVSL,
	_EthereumChainIDLowerName[65:68]: EthereumChainIDVSL,
	_EthereumChainIDName[68:76]:      EthereumChainIDArbitrum,
	_EthereumChainIDLowerName[68:76]: EthereumChainIDArbitrum,
	_EthereumChainIDName[76:80]:      EthereumChainIDAvalanche,
	_EthereumChainIDLowerName[76:80]: EthereumChainIDAvalanche,
	_EthereumChainIDName[80:85]:      EthereumChainIDLinea,
	_EthereumChainIDLowerName[80:85]: EthereumChainIDLinea,
}

var _EthereumChainIDNames = []string{
	_EthereumChainIDName[0:8],
	_EthereumChainIDName[8:16],
	_EthereumChainIDName[16:35],
	_EthereumChainIDName[35:41],
	_EthereumChainIDName[41:48],
	_EthereumChainIDName[48:52],
	_EthereumChainIDName[52:61],
	_EthereumChainIDName[61:65],
	_EthereumChainIDName[65:68],
	_EthereumChainIDName[68:76],
	_EthereumChainIDName[76:80],
	_EthereumChainIDName[80:85],
}

// EthereumChainIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EthereumChainIDString(s string) (EthereumChainID, error) {
	if val, ok := _EthereumChainIDNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EthereumChainIDNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EthereumChainID values", s)
}

// EthereumChainIDValues returns all values of the enum
func EthereumChainIDValues() []EthereumChainID {
	return _EthereumChainIDValues
}

// EthereumChainIDStrings returns a slice of all String values of the enum
func EthereumChainIDStrings() []string {
	strs := make([]string, len(_EthereumChainIDNames))
	copy(strs, _EthereumChainIDNames)
	return strs
}

// IsAEthereumChainID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EthereumChainID) IsAEthereumChainID() bool {
	_, ok := _EthereumChainIDMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for EthereumChainID
func (i EthereumChainID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EthereumChainID
func (i *EthereumChainID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EthereumChainID should be a string, got %s", data)
	}

	var err error
	*i, err = EthereumChainIDString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for EthereumChainID
func (i EthereumChainID) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EthereumChainID
func (i *EthereumChainID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = EthereumChainIDString(s)
	return err
}

func (i EthereumChainID) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *EthereumChainID) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of EthereumChainID: %[1]T(%[1]v)", value)
	}

	val, err := EthereumChainIDString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
