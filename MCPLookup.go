package mcpwrapper

import (
	"errors"
	"fmt"
	"strings"
)

type MethodInfo struct {
	Mcp       string    `json:"mcp"`
	Searge    string    `json:"searge"`
	Name      string    `json:"name"`
	Side      string    `json:"side"`
	Desc      string    `json:"desc"`
	ClassData ClassData `json:"classData"`
	ObofDesc  string    `json:"obofDesc"`
	DeobfDesc string    `json:"deobfDesc"`
}

func LookupMethod(name string, data MCPData, srg SRGNames) ([]MethodInfo, error) {
	for _, method := range srg.Methods {
		if method.Searge == name {
			return createMethodInfo(method, findMethod(method, data))
		}
	}
	return []MethodInfo{}, errors.New("failed to find method " + name)
}

//finds the full method data based off an srg method
func findMethod(method SRGMethod, data MCPData) []MethodData {
	var methods []MethodData
	for _, mcpMethod := range data.Methods {
		if mcpMethod.DeobfName == method.Searge {
			methods = append(methods, mcpMethod)
		}
	}
	return methods
}

func createMethodInfo(methodSrg SRGMethod, methodData []MethodData) ([]MethodInfo, error) {
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
	return fmt.Sprintf("mcp: `%s` srg: `%s` notch: `%s` (class `%s` `%s`)", info.Mcp, info.Searge, info.Name, info.ClassData.DeobfName, info.ClassData.ObofName)
}

//public net.minecraft.village.VillageCollection func_75549_c()V # removeAnnihilatedVillages
func MakeMethodAccessTransformer(info MethodInfo) string {
	//TODO use deobof desc
	return fmt.Sprintf("public %s %s%s # %s", strings.Replace(info.ClassData.DeobfName, "/", ".", -1), info.Searge, info.ObofDesc, info.Mcp)
}
