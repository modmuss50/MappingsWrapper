package common

type MapppingData struct {
	Classes []ClassData  `json:"classes"`
	Fields  []FieldData  `json:"fields"`
	Methods []MethodData `json:"methods"`
}

type ClassData struct {
	ObofName  string `json:"obofName"`
	DeobfName string `json:"deobfName"`
}

type FieldData struct {
	ClassData ClassData `json:"classData"`
	ObofType  string    `json:"obofType"`
	ObofName  string    `json:"obofName"`
	DeobfName string    `json:"deobfName"`
}

type MethodData struct {
	ClassData ClassData `json:"classData"`
	ObofName  string    `json:"obofName"`
	DeobfName string    `json:"deobfName"`
	ObofDesc  string    `json:"obofDesc"`
	DeobfDesc string    `json:"deobfDesc"`
}

func FindClass(obofName string, data MapppingData) *ClassData {
	for _, class := range data.Classes {
		if class.ObofName == obofName {
			return &class
		}
	}
	return nil
}
