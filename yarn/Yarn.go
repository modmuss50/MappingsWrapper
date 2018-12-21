package yarn

import (
	"encoding/json"
	"github.com/modmuss50/MappingsWrapper/utils"
)

//Gets the latest version of yarn from maven
//func GetYarnData() (common.MapppingData, error)  {
//	return nil, nil
//}

func GetPublishedYarnVersions() (map[string][]int, error) {

	var versions = make(map[string][]int)

	data, err := utils.DownloadString("https://maven.fabricmc.net/net/fabricmc/yarn/versions.json")
	if err != nil {
		return versions, err
	}

	err = json.Unmarshal([]byte(data), &versions)

	return versions, err
}
