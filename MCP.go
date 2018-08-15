package mcpwrapper

import (
	"github.com/modmuss50/goutils"
	"encoding/json"
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

//Only use the file version in the tests, use the version hosted on github when using on a production server so you can add new versions without needing to update
func ReadMCPVersionsFromFile(file string) MCPVersionJson {
	return ReadMCPVersions(readStringFromFile(file))
}

func ReadMCPVersions(str string) MCPVersionJson {
	versionData := MCPVersionJson{}
	json.Unmarshal([]byte(str), &versionData)
	return versionData
}

type MCPVersionJson struct {
	Versions []MCPVersion `json:"versions"`
}

type MCPVersion struct {
	MinecraftVersion string `json:"mcVersion"`
	MCPType          string `json:"type"`
}

func GetMCPData(version MCPVersion) (MCPData, error) {
	var data MCPData
	if version.MCPType == "mcp_config" {
		d, err := getMCPConfigData(version.MinecraftVersion)
		if err != nil {
			return data, err
		}
		data = d
	} else if version.MCPType == "mcp_legacy" {
		d, err := getMCPLeagcyData(version.MinecraftVersion)
		if err != nil {
			return data, err
		}
		data = d
	}
	return data, nil
}
