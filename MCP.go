package mcpwrapper

import (
	"github.com/modmuss50/goutils"
	"strings"
)

type MCPData struct {
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

func ReadMCPData(srgFile string) MCPData {
	data := MCPData{}
	for _, line := range goutils.ReadLinesFromFile(srgFile) {
		if line[0:3] == "CL:" {
			classNotch, classSrg := divideString2(line[4:])
			data.Classes = append(data.Classes, ClassData{classNotch, classSrg})
		} else if line[0:3] == "FD:" {
			obofData, deobfData := divideString2(line[4:])
			obofClass, obofField := splitAtLastSlash(obofData)
			deobfClass, deobfField := splitAtLastSlash(deobfData)
			fieldData := FieldData{ClassData: ClassData{obofClass, deobfClass}, ObofName: obofField, DeobfName: deobfField}
			data.Fields = append(data.Fields, fieldData)
		} else if line[0:3] == "MD:" {
			obofData, obofDesc, deobfData, deobfDesc := divideString4(line[4:])
			obofClass, obofMethod := splitAtLastSlash(obofData)
			deobfClass, deobfMethod := splitAtLastSlash(deobfData)
			methodData := MethodData{ClassData: ClassData{obofClass, deobfClass}, ObofName: obofMethod, DeobfName: deobfMethod, ObofDesc: obofDesc, DeobfDesc: deobfDesc}
			data.Methods = append(data.Methods, methodData)
		}
	}
	return data
}

func splitAtLastSlash(input string) (string, string) {
	lastPos := strings.LastIndex(input, "/")
	return input[:lastPos], input[lastPos:]
}
