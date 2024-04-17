package activity

type Calldata struct {
	Raw            string `json:"raw,omitempty"`
	FunctionHash   string `json:"function_hash,omitempty"`
	ParsedFunction string `json:"parsed_function,omitempty"`
}
