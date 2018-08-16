package mcpwrapper

import (
	"fmt"
	"github.com/modmuss50/goutils"
	"path/filepath"
	"sort"
	"strings"
)

func getMCPConfigData(version string) (MCPData, error) {
	var data MCPData
	if !fileExists(mcpConfigSRGLocation(version)) {
		err := prepareMCPConfig(version)
		if err != nil {
			return data, err
		}
	}
	data = buildMCPConfigData(version)
	return data, nil
}

func prepareMCPConfig(version string) error {
	url := fmt.Sprintf("http://files.minecraftforge.net/maven/de/oceanlabs/mcp/mcp_config/%s/mcp_config-%s.zip", version, version)
	archivePath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp_config-%s.zip", version))
	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp_config-%s", version))

	err := downloadFile(url, archivePath)
	if err != nil {
		return err
	}

	err = extractZip(archivePath, extractedPath)
	if err != nil {
		return err
	}

	tiny_srg := filepath.Join(extractedPath, "config", "joined.tsrg")
	return convertToSRG(tiny_srg, mcpConfigSRGLocation(version))
}

func mcpConfigSRGLocation(version string) string {
	extractedPath := filepath.Join(SRGDataDir, fmt.Sprintf("mcp_config-%s", version))
	srg := filepath.Join(extractedPath, "joined.srg")
	return srg
}

func buildMCPConfigData(version string) MCPData {
	return ReadMCPData(mcpConfigSRGLocation(version))
}

func convertToSRG(tsrg string, srg string) error {
	lines := goutils.ReadLinesFromFile(tsrg)
	var outputLines []string

	class := ""
	for _, line := range lines {
		if line[0] != '	' {
			//Hey we have a class
			outputLines = append(outputLines, "CL: "+line)
			class = line
		} else {
			split := strings.Split(line[1:], " ")
			if len(split) == 2 {
				//FIELD
				classNotch, classSrg := divideString2(class)
				fieldNotch, fieldSrg := divideString2(line[1:])
				outputLines = append(outputLines, fmt.Sprintf("FD: %s/%s %s/%s", classNotch, fieldNotch, classSrg, fieldSrg))
			} else if len(split) == 3 {
				//METHOD
				classNotch, classSrg := divideString2(class)
				methodNotch, methodDesc, methodSrg := divideString3(line[1:])
				srgDesc := "(?)unknown;" //TODO remap method desc

				outputLines = append(outputLines, fmt.Sprintf("MD: %s/%s %s %s/%s %s", classNotch, methodNotch, methodDesc, classSrg, methodSrg, srgDesc))
			}
		}
	}
	sort.Strings(outputLines)
	return writeStringToFile(strings.Join(outputLines, "\n"), srg)
}
