package mcpwrapper

import (
	"fmt"
	"github.com/modmuss50/goutils"
	"github.com/pkg/errors"
	"path/filepath"
)

type SRGNames struct {
	MCVersion string      `json:"minecraftVersion"`
	Fields    []SRGField  `json:"fields"`
	Methods   []SRGMethod `json:"methods"`
	Params    []SRGParm   `json:"params"`
}

type SRGField struct {
	Searge string `json:"searge"`
	Name   string `json:"name"`
	Side   string `json:"side"`
	Desc   string `json:"desc"`
}

type SRGMethod struct {
	Searge string `json:"searge"`
	Name   string `json:"name"`
	Side   string `json:"side"`
	Desc   string `json:"desc"`
}

type SRGParm struct {
	Searge string `json:"searge"`
	Name   string `json:"name"`
	Side   string `json:"side"`
}

// http://export.mcpbot.bspk.rs/mcp_snapshot/20180815-1.13/mcp_snapshot-20180815-1.13.zip

func GetSRGNames(version string) (SRGNames, error) {
	var export = SRGNames{}
	data, err := GetMCPBotVersions()
	if err != nil {
		return export, err
	}
	mcVersion, err := GetMCVersionFromExport(version, data)
	if err != nil {
		return export, err
	}
	branch, id := splitString2(version, "_")

	downloadUrl := fmt.Sprintf("http://export.mcpbot.bspk.rs/mcp_%s/%s-%s/mcp_%s-%s-%s.zip", branch, id, mcVersion, branch, id, mcVersion)

	downloadDir := filepath.Join(MCPDataDir, mcVersion, branch)
	downloadPath := filepath.Join(downloadDir, fmt.Sprintf("mcp_%s-%s-%s.zip", branch, id, mcVersion))
	extractPath := filepath.Join(downloadDir, fmt.Sprintf("mcp_%s-%s-%s", branch, id, mcVersion))

	if !fileExists(downloadPath) {
		makeDir(downloadDir)
		downloadFile(downloadUrl, downloadPath)

		makeDir(extractPath)
		extractZip(downloadPath, extractPath)
	}

	return readExport(version, data)
}

func readExport(version string, data MCPBotExports) (SRGNames, error) {
	var export = SRGNames{}
	mcVersion, err := GetMCVersionFromExport(version, data)
	if err != nil {
		return export, err
	}

	export.MCVersion = mcVersion

	branch, id := splitString2(version, "_")
	downloadDir := filepath.Join(MCPDataDir, mcVersion, branch)
	extractPath := filepath.Join(downloadDir, fmt.Sprintf("mcp_%s-%s-%s", branch, id, mcVersion))

	fieldsCsv := filepath.Join(extractPath, "fields.csv")
	methodsCsv := filepath.Join(extractPath, "methods.csv")
	paramsCsv := filepath.Join(extractPath, "params.csv")

	if !fileExists(fieldsCsv) || !fileExists(methodsCsv) || !fileExists(paramsCsv) {
		return export, errors.New("data not found")
	}

	handleFields := func(line string) {
		searge, name, side, desc := splitString4(line, ",")
		export.Fields = append(export.Fields, SRGField{Searge: searge, Name: name, Side: side, Desc: desc})
	}
	handleMethods := func(line string) {
		searge, name, side, desc := splitString4(line, ",")
		export.Methods = append(export.Methods, SRGMethod{Searge: searge, Name: name, Side: side, Desc: desc})
	}
	handleParam := func(line string) {
		searge, name, side := splitString3(line, ",")
		export.Params = append(export.Params, SRGParm{Searge: searge, Name: name, Side: side})
	}

	readLines(handleFields, fieldsCsv)
	readLines(handleMethods, methodsCsv)
	readLines(handleParam, paramsCsv)

	return export, nil
}

func GetSemiLiveNames() (SRGNames, error) {
	var export = SRGNames{}
	handleFields := func(line string) {
		searge, name, side, desc := splitString4(line, ",")
		export.Fields = append(export.Fields, SRGField{Searge: searge, Name: name, Side: side, Desc: desc})
	}
	handleMethods := func(line string) {
		searge, name, side, desc := splitString4(line, ",")
		export.Methods = append(export.Methods, SRGMethod{Searge: searge, Name: name, Side: side, Desc: desc})
	}
	handleParam := func(line string) {
		searge, name, side := splitString3(line, ",")
		export.Params = append(export.Params, SRGParm{Searge: searge, Name: name, Side: side})
	}
	err := downloadSemiLive("fields.csv", handleFields)
	if err != nil {
		return export, err
	}
	err = downloadSemiLive("methods.csv", handleMethods)
	if err != nil {
		return export, err
	}
	err = downloadSemiLive("params.csv", handleParam)
	if err != nil {
		return export, err
	}

	export.MCVersion = "semi-live" //TODO get the latest exported version?

	return export, nil
}

func downloadSemiLive(file string, handle func(line string)) error {
	downloadDir := filepath.Join(MCPDataDir, "semi-live")
	makeDir(downloadDir)
	csv := filepath.Join(downloadDir, file)
	if fileExists(csv) {
		deleteFile(file)
	}
	err := downloadFile(fmt.Sprintf("http://export.mcpbot.bspk.rs/%s", file), csv)
	if err != nil {
		return err
	}
	readLines(handle, csv)
	return nil
}

func readLines(handle func(line string), file string) {
	lines := goutils.ReadLinesFromFile(file)
	for i, line := range lines {
		if i == 0 {
			//Skip over the first line as its just the headers for the csv
			continue
		}
		handle(line)
	}
}

func GetMCVersionFromExport(version string, botData MCPBotExports) (string, error) {
	for _, versionEntry := range botData.Versions {
		for _, entry := range versionEntry.Snapshots {
			if entry == version {
				return versionEntry.MCVersion, nil
			}
		}
		for _, entry := range versionEntry.Stable {
			if entry == version {
				return versionEntry.MCVersion, nil
			}
		}
	}
	return "", nil
}
