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

func TestMethodLook(t *testing.T) {
	mcp, err := GetMCPData(MCPVersion{MinecraftVersion: "1.12.2", MCPType: "mcp_legacy"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := GetSRGNames("stable_39")
	if err != nil {
		t.Error(err)
		return
	}

	methods, err := LookupMethod("func_189667_a", mcp, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	for _, method := range methods {
		fmt.Println(MethodInfoToString(method))
	}
}

func TestMethodAccessTransformer(t *testing.T) {
	mcp, err := GetMCPData(MCPVersion{MinecraftVersion: "1.13", MCPType: "mcp_config"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := GetSRGNames("snapshot_20180916")
	if err != nil {
		t.Error(err)
		return
	}

	methods, err := LookupMethod("func_72923_a", mcp, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	if len(methods) == 0 {
		t.Fail()
	}

	at := MakeMethodAccessTransformer(methods[0])

	fmt.Println(at)

	if at != "public net.minecraft.world.WorldServer func_72923_a(Lnet/minecraft/entity/Entity;)V # onEntityAdded" {
		t.Fail()
	}
}
