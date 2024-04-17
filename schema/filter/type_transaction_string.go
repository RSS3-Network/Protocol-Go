// Code generated by "enumer --values --type=TransactionType --transform=snake --trimprefix=TypeTransaction --output type_transaction_string.go --json --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _TransactionTypeName = "approvalbridgeburnminttransfer"

var _TransactionTypeIndex = [...]uint8{0, 8, 14, 18, 22, 30}

const _TransactionTypeLowerName = "approvalbridgeburnminttransfer"

func (i TransactionType) String() string {
	i -= 1
	if i >= TransactionType(len(_TransactionTypeIndex)-1) {
		return fmt.Sprintf("TransactionType(%d)", i+1)
	}
	return _TransactionTypeName[_TransactionTypeIndex[i]:_TransactionTypeIndex[i+1]]
}

func (TransactionType) Values() []string {
	return TransactionTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TransactionTypeNoOp() {
	var x [1]struct{}
	_ = x[TypeTransactionApproval-(1)]
	_ = x[TypeTransactionBridge-(2)]
	_ = x[TypeTransactionBurn-(3)]
	_ = x[TypeTransactionMint-(4)]
	_ = x[TypeTransactionTransfer-(5)]
}

var _TransactionTypeValues = []TransactionType{TypeTransactionApproval, TypeTransactionBridge, TypeTransactionBurn, TypeTransactionMint, TypeTransactionTransfer}

var _TransactionTypeNameToValueMap = map[string]TransactionType{
	_TransactionTypeName[0:8]:        TypeTransactionApproval,
	_TransactionTypeLowerName[0:8]:   TypeTransactionApproval,
	_TransactionTypeName[8:14]:       TypeTransactionBridge,
	_TransactionTypeLowerName[8:14]:  TypeTransactionBridge,
	_TransactionTypeName[14:18]:      TypeTransactionBurn,
	_TransactionTypeLowerName[14:18]: TypeTransactionBurn,
	_TransactionTypeName[18:22]:      TypeTransactionMint,
	_TransactionTypeLowerName[18:22]: TypeTransactionMint,
	_TransactionTypeName[22:30]:      TypeTransactionTransfer,
	_TransactionTypeLowerName[22:30]: TypeTransactionTransfer,
}

var _TransactionTypeNames = []string{
	_TransactionTypeName[0:8],
	_TransactionTypeName[8:14],
	_TransactionTypeName[14:18],
	_TransactionTypeName[18:22],
	_TransactionTypeName[22:30],
}

// TransactionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TransactionTypeString(s string) (TransactionType, error) {
	if val, ok := _TransactionTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _TransactionTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TransactionType values", s)
}

// TransactionTypeValues returns all values of the enum
func TransactionTypeValues() []TransactionType {
	return _TransactionTypeValues
}

// TransactionTypeStrings returns a slice of all String values of the enum
func TransactionTypeStrings() []string {
	strs := make([]string, len(_TransactionTypeNames))
	copy(strs, _TransactionTypeNames)
	return strs
}

// IsATransactionType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TransactionType) IsATransactionType() bool {
	for _, v := range _TransactionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TransactionType
func (i TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TransactionType
func (i *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TransactionType should be a string, got %s", data)
	}

	var err error
	*i, err = TransactionTypeString(s)
	return err
}

func (i TransactionType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *TransactionType) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of TransactionType: %[1]T(%[1]v)", value)
	}

	val, err := TransactionTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
