package common

import (
	"github.com/modmuss50/MappingsWrapper/utils"
	"path/filepath"
)

var (
	DataDir         = "data"
	SRGDataDir      = filepath.Join(DataDir, "srg")
	MCPDataDir      = filepath.Join(DataDir, "mcp")
	YarnDataDir     = filepath.Join(DataDir, "yarn")
	IntermediaryDir = filepath.Join(DataDir, "intermediary")
)

func CheckDirs() {
	utils.MakeDir(SRGDataDir)
	utils.MakeDir(MCPDataDir)
	utils.MakeDir(YarnDataDir)
	utils.MakeDir(IntermediaryDir)
}

func todos() {

	//TODO allow choosing minecraft and mapping version
	//TODO have it remember the user preferences, and default to latest mc with MCP names?

	//TODO load MCPConfig srg names 1.13 >		http://files.minecraftforge.net/maven/de/oceanlabs/mcp/mcp_config/1.13/mcp_config-1.13.zip
	//TODO load MCP Legacy < 1.13				http://files.minecraftforge.net/maven/de/oceanlabs/mcp/mcp/1.12.2/mcp-1.12.2-srg.zip

	//TODO MCP Bot export diff's
	//TODO gen idea migration file?

	//TODO methodLookup
	//TODO Access Transformer lines
	//TODO debof, obf, srg names + desc's

	//TODO fieldLookup
	//TODO Access Transformer lines
	//TODO debof, obf, srg names + desc's

	//TODO paramLookup
}
