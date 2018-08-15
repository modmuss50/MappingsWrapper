package mcpwrapper

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/jmoiron/jsonq"
	"github.com/modmuss50/goutils"
	"sort"
)

type MCPBotVersions struct {
	Versions []MCPBotVersion `json:"versions"`
}

type MCPBotVersion struct {
	MCVersion string   `json:"mcVersion"`
	Snapshots []string `json:"snapshots"`
	Stable    []string `json:"stable"`
}

func GetMCPBotVersions() (MCPBotVersions, error) {
	versionsURL := "http://export.mcpbot.bspk.rs/versions.json"
	var botData = MCPBotVersions{}
	versionsJSON, err := goutils.DownloadString(versionsURL)
	if err != nil {
		return botData, err
	}
	data := goutils.GetDataMap(versionsJSON)
	for mcVersion, element := range data {
		versionElement := jsonq.NewQuery(element)
		var versionData = MCPBotVersion{}
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

func sortData(data MCPBotVersions) (MCPBotVersions, error) {
	var botData = MCPBotVersions{}
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

func GetVersionData(version string, botData MCPBotVersions) (MCPBotVersion, error) {
	for _, entry := range botData.Versions {
		if entry.MCVersion == version {
			return entry, nil
		}
	}
	return MCPBotVersion{}, errors.New("version not found")
}
