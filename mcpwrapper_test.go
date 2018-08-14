package mcpwrapper

import (
	"encoding/json"
	"fmt"
	"testing"
)

//Cleans out all the data dirs for the tests
func Test_setup(t *testing.T) {
	deleteDir(DataDir)
	checkDirs()
}

func TestPrepareMCPLegacy(t *testing.T) {
	version := "1.12.2"
	err := PrepareMCPLegacy(version)
	if err != nil {
		t.Error(err)
	}
}

func TestPrepareMCPConfig(t *testing.T) {
	version := "1.13"
	err := PrepareMCPConfig(version)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMCPConfigData(t *testing.T) {
	data := GetMCPConfigData("1.13")
	fmt.Println("1.13:")
	fmt.Printf("%d classes \n", len(data.Classes))
	fmt.Printf("%d Fields \n", len(data.Fields))
	fmt.Printf("%d Methods \n", len(data.Methods))
}

func TestGetMCPLegacyData(t *testing.T) {
	data := GetMCPLegacyData("1.12.2")
	fmt.Println("1.12.2:")
	fmt.Printf("%d classes \n", len(data.Classes))
	fmt.Printf("%d Fields \n", len(data.Fields))
	fmt.Printf("%d Methods \n", len(data.Methods))
}

func printAsJson(v interface{}) {
	json, _ := json.Marshal(v)
	fmt.Println(string(json))
}
