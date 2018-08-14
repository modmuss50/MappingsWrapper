package mcpwrapper

import (
	"fmt"
	"path/filepath"
)

func PrepareMCPLegacy(version string) error {
	url := fmt.Sprintf("http://files.minecraftforge.net/maven/de/oceanlabs/mcp/mcp/%s/mcp-%s-srg.zip", version, version)
	archivePath := filepath.Join(DataDir, fmt.Sprintf("mcp-%s-srg.zip", version))
	extractedPath := filepath.Join(DataDir, fmt.Sprintf("mcp-%s-srg", version))
	err := downloadFile(url, archivePath)
	if err != nil {
		return err
	}
	err = extractZip(archivePath, extractedPath)
	if err != nil {
		return err
	}
	return nil
}

func GetMCPLegacyData(version string) MCPData {
	return ReadMCPData(MCPLegacySRGLocation(version))
}

func MCPLegacySRGLocation(version string) string {
	extractedPath := filepath.Join(DataDir, fmt.Sprintf("mcp-%s-srg", version))
	srg := filepath.Join(extractedPath, "joined.srg")
	return srg
}
