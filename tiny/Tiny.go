package tiny

import (
	"fmt"
	"github.com/modmuss50/MappingsWrapper/common"
	"github.com/modmuss50/MappingsWrapper/utils"
	"path/filepath"
	"strings"
)

func ReadTinyMapping(mappings string, obof string, deobf string) common.MapppingData {
	data := common.MapppingData{}
	lines := utils.ReadLinesFromFile(mappings)

	obofOffset := 0
	deobfOffset := 0

	header := strings.Fields(lines[0])

	if header[0] != "v1" {
		fmt.Println("Invalid tiny mapping file!")
		return data
	}
	for i, val := range header {
		if val == obof {
			obofOffset = i - 1
		}
		if val == deobf {
			deobfOffset = i - 1
		}
	}

	fmt.Println(obofOffset)
	fmt.Println(deobfOffset)

	//for _, line := range lines {
	//	split := strings.Fields(line)
	//	if len(split) == 0 {
	//
	//	} else if split[0] == "v1" {
	//
	//	} else if split[0] == "CLASS" {
	//
	//	} else if split[0] == "METHOD" {
	//		class := common.FindClass(split[1], data)
	//		data.Methods = append(data.Methods, common.MethodData{ClassData:class})
	//	}
	//}

	handle("CLASS", lines, func(line string, split []string) {
		data.Classes = append(data.Classes, common.ClassData{
			ObofName:  split[1+obofOffset],
			DeobfName: split[1+deobfOffset],
		})
	})

	handle("METHOD", lines, func(line string, split []string) {
		class := common.FindClass(split[1], data)
		data.Methods = append(data.Methods, common.MethodData{
			ClassData: *class,
			ObofDesc:  split[2],
			ObofName:  split[3+obofOffset],
			DeobfName: split[3+deobfOffset],
		})
	})

	handle("FIELD", lines, func(line string, split []string) {
		class := common.FindClass(split[1], data)
		data.Fields = append(data.Fields, common.FieldData{
			ClassData: *class,
			ObofType:  split[2],
			ObofName:  split[3+obofOffset],
			DeobfName: split[3+deobfOffset],
		})
	})

	return data
}

func handle(entry string, lines []string, handle lineHandler) {
	for _, line := range lines {
		split := strings.Fields(line)
		if len(split) == 0 {

		} else if split[0] == entry {
			handle(line, split)
		}
	}
}

type lineHandler func(line string, split []string)

func ReadIntermediary(mcVersion string) common.MapppingData {

	tinyFile := filepath.Join(common.IntermediaryDir, mcVersion+".tiny")

	utils.DownloadFile("https://github.com/FabricMC/intermediary/raw/master/mappings/"+mcVersion+".tiny", tinyFile)

	return ReadTinyMapping(tinyFile, "official", "intermediary")
}
