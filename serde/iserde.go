package serde

type ISerde interface {
	Serialize(*Operation) ([]byte, error)
	Deserialize([]byte) (*Operation, error)
}

type Operation struct {
	DataType   string
	Opr        string
	ResourceID *string
	Params     []interface{}
}

func NewOperation(dataType string, opr string, resourceID *string, params ...interface{}) *Operation {
	return &Operation{
		DataType:   dataType,
		Opr:        opr,
		ResourceID: resourceID,
		Params:     params,
	}
}
