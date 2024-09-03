// Code generated by "enumer --values --type=GovernanceType --transform=snake --trimprefix=Governance --output governance_string.go --json --sql"; DO NOT EDIT.

package typex

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _GovernanceTypeName = "proposalvote"

var _GovernanceTypeIndex = [...]uint8{0, 8, 12}

const _GovernanceTypeLowerName = "proposalvote"

func (i GovernanceType) String() string {
	i -= 1
	if i >= GovernanceType(len(_GovernanceTypeIndex)-1) {
		return fmt.Sprintf("GovernanceType(%d)", i+1)
	}
	return _GovernanceTypeName[_GovernanceTypeIndex[i]:_GovernanceTypeIndex[i+1]]
}

func (GovernanceType) Values() []string {
	return GovernanceTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _GovernanceTypeNoOp() {
	var x [1]struct{}
	_ = x[GovernanceProposal-(1)]
	_ = x[GovernanceVote-(2)]
}

var _GovernanceTypeValues = []GovernanceType{GovernanceProposal, GovernanceVote}

var _GovernanceTypeNameToValueMap = map[string]GovernanceType{
	_GovernanceTypeName[0:8]:       GovernanceProposal,
	_GovernanceTypeLowerName[0:8]:  GovernanceProposal,
	_GovernanceTypeName[8:12]:      GovernanceVote,
	_GovernanceTypeLowerName[8:12]: GovernanceVote,
}

var _GovernanceTypeNames = []string{
	_GovernanceTypeName[0:8],
	_GovernanceTypeName[8:12],
}

// GovernanceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func GovernanceTypeString(s string) (GovernanceType, error) {
	if val, ok := _GovernanceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _GovernanceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to GovernanceType values", s)
}

// GovernanceTypeValues returns all values of the enum
func GovernanceTypeValues() []GovernanceType {
	return _GovernanceTypeValues
}

// GovernanceTypeStrings returns a slice of all String values of the enum
func GovernanceTypeStrings() []string {
	strs := make([]string, len(_GovernanceTypeNames))
	copy(strs, _GovernanceTypeNames)
	return strs
}

// IsAGovernanceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i GovernanceType) IsAGovernanceType() bool {
	for _, v := range _GovernanceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for GovernanceType
func (i GovernanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for GovernanceType
func (i *GovernanceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("GovernanceType should be a string, got %s", data)
	}

	var err error
	*i, err = GovernanceTypeString(s)
	return err
}

func (i GovernanceType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *GovernanceType) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of GovernanceType: %[1]T(%[1]v)", value)
	}

	val, err := GovernanceTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}