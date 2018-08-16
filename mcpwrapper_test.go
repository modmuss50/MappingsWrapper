package mcpwrapper

import (
	"fmt"
	"testing"
)

func Test_setup(t *testing.T) {
	//deleteDir(SRGDataDir)
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

func TestGetMCVersionFromExport(t *testing.T) {
	data, err := GetMCPBotVersions()
	if err != nil {
		t.Error(err)
		return
	}
	mcVersion, err := GetMCVersionFromExport("snapshot_20180604", data)
	if err != nil {
		t.Error(err)
		return
	}
	if mcVersion != "1.12" {
		t.Fail()
		return
	}
	mcVersion, err = GetMCVersionFromExport("stable_22", data)
	if err != nil {
		t.Error(err)
		return
	}
	if mcVersion != "1.8.9" {
		t.Fail()
		return
	}
}

func TestDownloadExport(t *testing.T) {
	data, err := GetSRGNames("snapshot_20180815")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))

	data, err = GetSRGNames("stable_39")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))
}

func TestGetSemiLive(t *testing.T) {
	data, err := GetSemiLiveNames()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))
}
