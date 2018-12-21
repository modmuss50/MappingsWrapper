package mcp

import (
	"errors"
	"fmt"
	"github.com/modmuss50/MappingsWrapper/common"
	"strings"
)

type MethodInfo struct {
	Mcp       string           `json:"mcp"`
	Searge    string           `json:"searge"`
	Name      string           `json:"name"`
	Side      string           `json:"side"`
	Desc      string           `json:"desc"`
	ClassData common.ClassData `json:"classData"`
	ObofDesc  string           `json:"obofDesc"`
	DeobfDesc string           `json:"deobfDesc"`
}

func LookupMethod(name string, data common.MapppingData, srg SRGNames) ([]MethodInfo, error) {
	for _, method := range srg.Methods {
		if method.Searge == name {
			return createMethodInfo(method, findMethod(method, data))
		}
	}
	return []MethodInfo{}, errors.New("failed to find method " + name)
}

//finds the full method data based off an srg method
func findMethod(method SRGMethod, data common.MapppingData) []common.MethodData {
	var methods []common.MethodData
	for _, mcpMethod := range data.Methods {
		if mcpMethod.DeobfName == method.Searge {
			methods = append(methods, mcpMethod)
		}
	}
	return methods
}

func createMethodInfo(methodSrg SRGMethod, methodData []common.MethodData) ([]MethodInfo, error) {
	if len(methodData) == 0 {
		return []MethodInfo{}, errors.New("Failed to find method data for " + methodSrg.Searge + methodSrg.Desc)
	}
	var methods []MethodInfo
	for _, method := range methodData {
		info := MethodInfo{Mcp: methodSrg.Name, Searge: methodSrg.Searge, Name: method.ObofName, Side: methodSrg.Side, Desc: methodSrg.Desc, ClassData: method.ClassData, ObofDesc: method.ObofDesc, DeobfDesc: method.DeobfDesc}
		methods = append(methods, info)
	}
	return methods, nil
}

func MethodInfoToString(info MethodInfo) string {
	return fmt.Sprintf("mcp: `%s` srg: `%s` notch: `%s` (owner `%s` `%s`)", info.Mcp, info.Searge, info.Name, info.ClassData.DeobfName, info.ClassData.ObofName)
}

//public net.minecraft.village.VillageCollection func_75549_c()V # removeAnnihilatedVillages
func MakeMethodAccessTransformer(info MethodInfo) string {
	return fmt.Sprintf("public %s %s%s # %s", strings.Replace(info.ClassData.DeobfName, "/", ".", -1), info.Searge, info.DeobfDesc, info.Mcp)
}

type FieldInfo struct {
	Mcp       string           `json:"mcp"`
	Searge    string           `json:"searge"`
	Name      string           `json:"name"`
	Side      string           `json:"side"`
	Desc      string           `json:"desc"`
	ClassData common.ClassData `json:"classData"`
}

func LookupField(name string, data common.MapppingData, srg SRGNames) ([]FieldInfo, error) {
	for _, field := range srg.Fields {
		if field.Searge == name {
			return createFieldInfo(field, findField(field, data))
		}
	}
	return []FieldInfo{}, errors.New("failed to find field " + name)
}

//finds the full method data based off an srg method
func findField(method SRGField, data common.MapppingData) []common.FieldData {
	var fields []common.FieldData
	for _, mcpField := range data.Fields {
		if mcpField.DeobfName == method.Searge {
			fields = append(fields, mcpField)
		}
	}
	return fields
}

func createFieldInfo(fieldSrg SRGField, fieldData []common.FieldData) ([]FieldInfo, error) {
	if len(fieldData) == 0 {
		return []FieldInfo{}, errors.New("Failed to find field data for " + fieldSrg.Searge + fieldSrg.Desc)
	}
	var fields []FieldInfo
	for _, field := range fieldData {
		info := FieldInfo{Mcp: fieldSrg.Name, Searge: fieldSrg.Searge, Name: field.ObofName, Side: fieldSrg.Side, Desc: fieldSrg.Desc, ClassData: field.ClassData}
		fields = append(fields, info)
	}
	return fields, nil
}

func FieldInfoToString(info FieldInfo) string {
	return fmt.Sprintf("mcp: `%s` srg: `%s` notch: `%s` (owner `%s` `%s`)", info.Mcp, info.Searge, info.Name, info.ClassData.DeobfName, info.ClassData.ObofName)
}

//public net.minecraft.client.gui.recipebook.GuiRecipeBook field_193960_m # toggleRecipesBtn
func MakeFieldAccessTransformer(info FieldInfo) string {
	return fmt.Sprintf("public %s %s # %s", strings.Replace(info.ClassData.DeobfName, "/", ".", -1), info.Searge, info.Mcp)
}

func FindMethodSRGName(name string, srgs SRGNames) (SRGMethod, error) {
	for _, method := range srgs.Methods {
		if method.Searge == name {
			return method, nil
		}
	}
	return SRGMethod{}, errors.New("failed to find method in names")
}

func FindFieldSRGName(name string, srgs SRGNames) (SRGField, error) {
	for _, field := range srgs.Fields {
		if field.Searge == name {
			return field, nil
		}
	}
	return SRGField{}, errors.New("failed to find field in names")
}

func FindParamSRGName(name string, srgs SRGNames) (SRGParam, error) {
	for _, param := range srgs.Params {
		if param.Searge == name {
			return param, nil
		}
	}
	return SRGParam{}, errors.New("failed to find param in names")
}
