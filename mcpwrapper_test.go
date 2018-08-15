package mcpwrapper

import (
	"fmt"
	"testing"
)

func Test_setup(t *testing.T) {
	//deleteDir(DataDir)
	checkDirs()
}

func TestPrepareAll(t *testing.T) {
	versions := ReadMCPVersionsFromFile("versions.json")
	for _, version := range versions.Versions {
		data, err := GetMCPData(version)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(version.MinecraftVersion)
		fmt.Printf("%d classes \n", len(data.Classes))
		fmt.Printf("%d Fields \n", len(data.Fields))
		fmt.Printf("%d Methods \n", len(data.Methods))
	}
}

func TestGetMCPBotVersions(t *testing.T) {
	data, err := GetMCPBotVersions()
	if err != nil {
		t.Error(err)
		return
	}
	printAsJson(data)
}
