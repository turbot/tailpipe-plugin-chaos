package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
)

type StructColumns struct {
	// embed required enrichment fields
	enrichment.CommonFields

	Identifier string    `json:"identifier"`
	Timestamp  time.Time `json:"timestamp"`

	// Additional fields
	SimpleStruct        *SimpleStruct        `json:"simple_struct,omitempty"`
	ArrayStruct         *StructWithArray     `json:"array_struct,omitempty"`
	NestedStruct        *NestedStruct        `json:"nested_struct,omitempty"`
	ComplexNestedStruct *ComplexNestedStruct `json:"complex_nested_struct,omitempty"`
}

type SimpleStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type StructWithArray struct {
	Id      int64 `json:"id"`
	Numbers []int `json:"numbers"`
}

type NestedStruct struct {
	Id   int64         `json:"id"`
	Name string        `json:"name"`
	Sub  *SimpleStruct `json:"sub,omitempty"`
}

type ComplexNestedStruct struct {
	Id int64 `json:"id"`

	SubStructs []*SimpleStruct `json:"sub_structs,omitempty"`
}
