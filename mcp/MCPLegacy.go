package mcp

import (
	"fmt"
	"github.com/modmuss50/MappingsWrapper/common"
	"github.com/modmuss50/MappingsWrapper/utils"
	"path/filepath"
)

func getMCPLeagcyData(version string) (common.MapppingData, error) {
	var data common.MapppingData
	if !utils.FileExists(mcpLegacySRGLocation(version)) {
		err := prepareMCPLegacy(version)
		if err != nil {
			return data, err
		}
	}
	data = buildMCPLegacyData(version)
	return data, nil
}

func prepareMCPLegacy(version string) error {
	url := fmt.Sprintf("http://files.minecraftforge.net/maven/de/oceanlabs/mcp/mcp/%s/mcp-%s-srg.zip", version, version)
	archivePath := filepath.Join(common.SRGDataDir, fmt.Sprintf("mcp-%s-srg.zip", version))
	extractedPath := filepath.Join(common.SRGDataDir, fmt.Sprintf("mcp-%s-srg", version))
	err := utils.DownloadFile(url, archivePath)
	if err != nil {
		return err
	}
	err = utils.ExtractZip(archivePath, extractedPath)
	if err != nil {
		return err
	}
	return nil
}

func buildMCPLegacyData(version string) common.MapppingData {
	return ReadMCPData(mcpLegacySRGLocation(version))
}

func mcpLegacySRGLocation(version string) string {
	extractedPath := filepath.Join(common.SRGDataDir, fmt.Sprintf("mcp-%s-srg", version))
	srg := filepath.Join(extractedPath, "joined.srg")
	return srg
}
