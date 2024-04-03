// Code generated by "enumer --values --type=Name --linecomment --output worker_string.go --json --yaml --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NameName = "unknownfallbackmirrorfarcasterrss3paragraphopenseauniswapoptimismaavegotchilenslooksraremattersmomokahighlightaaveiqwikilidocrossbellenskiwistand1inchvslsavmstargatecurve"

var _NameIndex = [...]uint8{0, 7, 15, 21, 30, 34, 43, 50, 57, 65, 75, 79, 88, 95, 101, 110, 114, 120, 124, 133, 136, 145, 150, 153, 157, 165, 170}

const _NameLowerName = "unknownfallbackmirrorfarcasterrss3paragraphopenseauniswapoptimismaavegotchilenslooksraremattersmomokahighlightaaveiqwikilidocrossbellenskiwistand1inchvslsavmstargatecurve"

func (i Name) String() string {
	if i < 0 || i >= Name(len(_NameIndex)-1) {
		return fmt.Sprintf("Name(%d)", i)
	}
	return _NameName[_NameIndex[i]:_NameIndex[i+1]]
}

func (Name) Values() []string {
	return NameStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NameNoOp() {
	var x [1]struct{}
	_ = x[Unknown-(0)]
	_ = x[Fallback-(1)]
	_ = x[Mirror-(2)]
	_ = x[Farcaster-(3)]
	_ = x[RSS3-(4)]
	_ = x[Paragraph-(5)]
	_ = x[OpenSea-(6)]
	_ = x[Uniswap-(7)]
	_ = x[Optimism-(8)]
	_ = x[Aavegotchi-(9)]
	_ = x[Lens-(10)]
	_ = x[Looksrare-(11)]
	_ = x[Matters-(12)]
	_ = x[Momoka-(13)]
	_ = x[Highlight-(14)]
	_ = x[Aave-(15)]
	_ = x[IQWiki-(16)]
	_ = x[Lido-(17)]
	_ = x[Crossbell-(18)]
	_ = x[ENS-(19)]
	_ = x[KiwiStand-(20)]
	_ = x[Oneinch-(21)]
	_ = x[VSL-(22)]
	_ = x[SAVM-(23)]
	_ = x[Stargate-(24)]
	_ = x[Curve-(25)]
}

var _NameValues = []Name{Unknown, Fallback, Mirror, Farcaster, RSS3, Paragraph, OpenSea, Uniswap, Optimism, Aavegotchi, Lens, Looksrare, Matters, Momoka, Highlight, Aave, IQWiki, Lido, Crossbell, ENS, KiwiStand, Oneinch, VSL, SAVM, Stargate, Curve}

var _NameNameToValueMap = map[string]Name{
	_NameName[0:7]:          Unknown,
	_NameLowerName[0:7]:     Unknown,
	_NameName[7:15]:         Fallback,
	_NameLowerName[7:15]:    Fallback,
	_NameName[15:21]:        Mirror,
	_NameLowerName[15:21]:   Mirror,
	_NameName[21:30]:        Farcaster,
	_NameLowerName[21:30]:   Farcaster,
	_NameName[30:34]:        RSS3,
	_NameLowerName[30:34]:   RSS3,
	_NameName[34:43]:        Paragraph,
	_NameLowerName[34:43]:   Paragraph,
	_NameName[43:50]:        OpenSea,
	_NameLowerName[43:50]:   OpenSea,
	_NameName[50:57]:        Uniswap,
	_NameLowerName[50:57]:   Uniswap,
	_NameName[57:65]:        Optimism,
	_NameLowerName[57:65]:   Optimism,
	_NameName[65:75]:        Aavegotchi,
	_NameLowerName[65:75]:   Aavegotchi,
	_NameName[75:79]:        Lens,
	_NameLowerName[75:79]:   Lens,
	_NameName[79:88]:        Looksrare,
	_NameLowerName[79:88]:   Looksrare,
	_NameName[88:95]:        Matters,
	_NameLowerName[88:95]:   Matters,
	_NameName[95:101]:       Momoka,
	_NameLowerName[95:101]:  Momoka,
	_NameName[101:110]:      Highlight,
	_NameLowerName[101:110]: Highlight,
	_NameName[110:114]:      Aave,
	_NameLowerName[110:114]: Aave,
	_NameName[114:120]:      IQWiki,
	_NameLowerName[114:120]: IQWiki,
	_NameName[120:124]:      Lido,
	_NameLowerName[120:124]: Lido,
	_NameName[124:133]:      Crossbell,
	_NameLowerName[124:133]: Crossbell,
	_NameName[133:136]:      ENS,
	_NameLowerName[133:136]: ENS,
	_NameName[136:145]:      KiwiStand,
	_NameLowerName[136:145]: KiwiStand,
	_NameName[145:150]:      Oneinch,
	_NameLowerName[145:150]: Oneinch,
	_NameName[150:153]:      VSL,
	_NameLowerName[150:153]: VSL,
	_NameName[153:157]:      SAVM,
	_NameLowerName[153:157]: SAVM,
	_NameName[157:165]:      Stargate,
	_NameLowerName[157:165]: Stargate,
	_NameName[165:170]:      Curve,
	_NameLowerName[165:170]: Curve,
}

var _NameNames = []string{
	_NameName[0:7],
	_NameName[7:15],
	_NameName[15:21],
	_NameName[21:30],
	_NameName[30:34],
	_NameName[34:43],
	_NameName[43:50],
	_NameName[50:57],
	_NameName[57:65],
	_NameName[65:75],
	_NameName[75:79],
	_NameName[79:88],
	_NameName[88:95],
	_NameName[95:101],
	_NameName[101:110],
	_NameName[110:114],
	_NameName[114:120],
	_NameName[120:124],
	_NameName[124:133],
	_NameName[133:136],
	_NameName[136:145],
	_NameName[145:150],
	_NameName[150:153],
	_NameName[153:157],
	_NameName[157:165],
	_NameName[165:170],
}

// NameString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NameString(s string) (Name, error) {
	if val, ok := _NameNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NameNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Name values", s)
}

// NameValues returns all values of the enum
func NameValues() []Name {
	return _NameValues
}

// NameStrings returns a slice of all String values of the enum
func NameStrings() []string {
	strs := make([]string, len(_NameNames))
	copy(strs, _NameNames)
	return strs
}

// IsAName returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Name) IsAName() bool {
	for _, v := range _NameValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Name
func (i Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Name
func (i *Name) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Name should be a string, got %s", data)
	}

	var err error
	*i, err = NameString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Name
func (i Name) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Name
func (i *Name) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NameString(s)
	return err
}

func (i Name) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Name) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Name: %[1]T(%[1]v)", value)
	}

	val, err := NameString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
