package mcp

import (
	"encoding/json"
	"github.com/modmuss50/MappingsWrapper/common"
	"github.com/modmuss50/MappingsWrapper/utils"
	"github.com/modmuss50/goutils"
)

func ReadMCPData(srgFile string) common.MapppingData {
	data := common.MapppingData{}
	for _, line := range goutils.ReadLinesFromFile(srgFile) {
		if line[0:3] == "CL:" {
			classNotch, classSrg := utils.DivideString2(line[4:])
			data.Classes = append(data.Classes, common.ClassData{ObofName: classNotch, DeobfName: classSrg})
		} else if line[0:3] == "FD:" {
			obofData, deobfData := utils.DivideString2(line[4:])
			obofClass, obofField := utils.SplitAtLastSlash(obofData)
			deobfClass, deobfField := utils.SplitAtLastSlash(deobfData)
			fieldData := common.FieldData{ClassData: common.ClassData{obofClass, deobfClass}, ObofName: obofField, DeobfName: deobfField}
			data.Fields = append(data.Fields, fieldData)
		} else if line[0:3] == "MD:" {
			obofData, obofDesc, deobfData, deobfDesc := utils.DivideString4(line[4:])
			obofClass, obofMethod := utils.SplitAtLastSlash(obofData)
			deobfClass, deobfMethod := utils.SplitAtLastSlash(deobfData)
			methodData := common.MethodData{ClassData: common.ClassData{obofClass, deobfClass}, ObofName: obofMethod, DeobfName: deobfMethod, ObofDesc: obofDesc, DeobfDesc: deobfDesc}
			data.Methods = append(data.Methods, methodData)
		}
	}
	return data
}

//Only use the file version in the tests, use the version hosted on github when using on a production server so you can add new versions without needing to update
func ReadMCPVersionsFromFile(file string) MCPVersionJson {
	return ReadMCPVersions(utils.ReadStringFromFile(file))
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

func GetMCPData(version MCPVersion) (common.MapppingData, error) {
	var data common.MapppingData
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
