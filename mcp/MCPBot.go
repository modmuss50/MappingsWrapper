package mcp

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/jmoiron/jsonq"
	"github.com/modmuss50/goutils"
	"sort"
)

type MCPBotExports struct {
	Versions []MCPBotVersionExports `json:"exports"`
}

type MCPBotVersionExports struct {
	MCVersion string   `json:"minecraftVersion"`
	Snapshots []string `json:"snapshots"`
	Stable    []string `json:"stable"`
}

func GetMCPBotVersions() (MCPBotExports, error) {
	versionsURL := "http://export.mcpbot.bspk.rs/versions.json"
	var botData = MCPBotExports{}
	versionsJSON, err := goutils.DownloadString(versionsURL)
	if err != nil {
		return botData, err
	}
	data := goutils.GetDataMap(versionsJSON)
	for mcVersion, element := range data {
		versionElement := jsonq.NewQuery(element)
		var versionData = MCPBotVersionExports{}
		versionData.MCVersion = mcVersion

		val, err := handleBranch("snapshot", versionElement, versionData.Snapshots)
		if err != nil {
			return botData, err
		}
		versionData.Snapshots = val
		val, err = handleBranch("stable", versionElement, versionData.Stable)
		if err != nil {
			return botData, err
		}
		versionData.Stable = val
		botData.Versions = append(botData.Versions, versionData)
	}
	return sortData(botData)
}

func handleBranch(branch string, json *jsonq.JsonQuery, target []string) ([]string, error) {
	var value []string
	value = target

	entries, err := json.ArrayOfInts(branch)
	if err != nil {
		fmt.Println(err)
		return value, err
	}

	for _, id := range entries {
		value = append(value, fmt.Sprintf("%s_%d", branch, id))
	}
	return value, nil
}

func sortData(data MCPBotExports) (MCPBotExports, error) {
	var botData = MCPBotExports{}
	vs := make([]*semver.Version, len(data.Versions))
	for i, r := range data.Versions {
		v, err := semver.NewVersion(r.MCVersion)
		if err != nil {
			return botData, err
		}

		vs[i] = v
	}
	sort.Sort(sort.Reverse(semver.Collection(vs)))
	for _, version := range vs {
		version.Original()
		data, err := GetVersionData(version.Original(), data)
		if err != nil {
			return botData, nil
		}
		botData.Versions = append(botData.Versions, data)
	}
	return botData, nil
}

func GetVersionData(version string, botData MCPBotExports) (MCPBotVersionExports, error) {
	for _, entry := range botData.Versions {
		if entry.MCVersion == version {
			return entry, nil
		}
	}
	return MCPBotVersionExports{}, errors.New("version not found")
}
