package mapppingswrapper

import (
	"fmt"
	"github.com/modmuss50/MappingsWrapper/common"
	"github.com/modmuss50/MappingsWrapper/mcp"
	"github.com/modmuss50/MappingsWrapper/utils"
	"testing"
)

func Test_MCP_Setup(t *testing.T) {
	//deleteDir(SRGDataDir)
	common.CheckDirs()
}

func Test_MCP_Prepare(t *testing.T) {
	//Delete the dir to ensure that we always test a full clone + build
	//deleteDir(filepath.Join(SRGDataDir, fmt.Sprintf("mcp-%s-config", "1.13.1")))

	data, err := mcp.GetMCPData(mcp.MCPVersion{MinecraftVersion: "1.13.1", MCPType: "mcp_config"})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("1.13.1")
	fmt.Printf("%d classes \n", len(data.Classes))
	fmt.Printf("%d Fields \n", len(data.Fields))
	fmt.Printf("%d Methods \n", len(data.Methods))
}

func Test_MCP_PrepareAll(t *testing.T) {
	versions := mcp.ReadMCPVersionsFromFile("versions.json")
	for _, version := range versions.Versions {
		data, err := mcp.GetMCPData(version)
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

func Test_MCP_GetMCPBotVersions(t *testing.T) {
	data, err := mcp.GetMCPBotVersions()
	if err != nil {
		t.Error(err)
		return
	}
	utils.PrintAsJson(data)
}

func Test_MCP_GetMCVersionFromExport(t *testing.T) {
	data, err := mcp.GetMCPBotVersions()
	if err != nil {
		t.Error(err)
		return
	}
	mcVersion, err := mcp.GetMCVersionFromExport("snapshot_20180604", data)
	if err != nil {
		t.Error(err)
		return
	}
	if mcVersion != "1.12" {
		t.Fail()
		return
	}
	mcVersion, err = mcp.GetMCVersionFromExport("stable_22", data)
	if err != nil {
		t.Error(err)
		return
	}
	if mcVersion != "1.8.9" {
		t.Fail()
		return
	}
}

func Test_MCP_DownloadExport(t *testing.T) {
	data, err := mcp.GetSRGNames("snapshot_20180815")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))

	data, err = mcp.GetSRGNames("stable_39")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))
}

func Test_MCP_GetSemiLive(t *testing.T) {
	data, err := mcp.GetSemiLiveNames()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.MCVersion)
	fmt.Printf("%d srg field names \n", len(data.Fields))
	fmt.Printf("%d srg method names \n", len(data.Methods))
	fmt.Printf("%d srg params names \n", len(data.Params))
}

func Test_MCP_MethodLook(t *testing.T) {
	mcpData, err := mcp.GetMCPData(mcp.MCPVersion{MinecraftVersion: "1.12.2", MCPType: "mcp_legacy"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := mcp.GetSRGNames("stable_39")
	if err != nil {
		t.Error(err)
		return
	}

	methods, err := mcp.LookupMethod("func_189667_a", mcpData, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	for _, method := range methods {
		fmt.Println(mcp.MethodInfoToString(method))
	}
}

func Test_MCP_MethodAccessTransformer(t *testing.T) {
	mcpData, err := mcp.GetMCPData(mcp.MCPVersion{MinecraftVersion: "1.13", MCPType: "mcp_config"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := mcp.GetSRGNames("snapshot_20180916")
	if err != nil {
		t.Error(err)
		return
	}

	methods, err := mcp.LookupMethod("func_72923_a", mcpData, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	if len(methods) == 0 {
		t.Fail()
	}

	at := mcp.MakeMethodAccessTransformer(methods[0])

	fmt.Println(at)

	if at != "public net.minecraft.world.WorldServer func_72923_a(Lnet/minecraft/entity/Entity;)V # onEntityAdded" {
		t.Fail()
	}
}

func Test_MCP_FieldLookup(t *testing.T) {
	//field_193960_m

	mcpData, err := mcp.GetMCPData(mcp.MCPVersion{MinecraftVersion: "1.13", MCPType: "mcp_config"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := mcp.GetSRGNames("snapshot_20180916")
	if err != nil {
		t.Error(err)
		return
	}

	fields, err := mcp.LookupField("field_193960_m", mcpData, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	for _, field := range fields {
		fmt.Println(mcp.FieldInfoToString(field))
	}
}

func Test_MCP_FieldAccessTransformer(t *testing.T) {
	mcpData, err := mcp.GetMCPData(mcp.MCPVersion{MinecraftVersion: "1.13", MCPType: "mcp_config"})
	if err != nil {
		t.Error(err)
		return
	}

	srgs, err := mcp.GetSRGNames("snapshot_20180916")
	if err != nil {
		t.Error(err)
		return
	}

	fields, err := mcp.LookupField("field_193022_s", mcpData, srgs)
	if err != nil {
		t.Error(err)
		return
	}

	if len(fields) == 0 {
		t.Fail()
	}

	at := mcp.MakeFieldAccessTransformer(fields[0])

	fmt.Println(at)

	if at != "public net.minecraft.client.gui.recipebook.GuiRecipeBook field_193022_s # recipeBookPage" {
		t.Fail()
	}
}

func Test_MCP_DiffGeneration(t *testing.T) {
	oldSRG, err := mcp.GetSRGNames("snapshot_20180910")
	if err != nil {
		t.Error(err)
		return
	}
	newSRG, err := mcp.GetSRGNames("snapshot_20180916")
	if err != nil {
		t.Error(err)
		return
	}

	//https://paste.modmuss50.me/view/9e89d3d9 example from old system

	diff := mcp.GenerateDiff(oldSRG, newSRG)
	diffStr := mcp.DiffToString(diff)

	fmt.Println(diffStr)
}
