package mapppingswrapper

import (
	"fmt"
	"github.com/modmuss50/MappingsWrapper/common"
	"github.com/modmuss50/MappingsWrapper/tiny"
	"github.com/modmuss50/MappingsWrapper/yarn"
	"testing"
)

func Test_Yarn_Setup(t *testing.T) {
	common.CheckDirs()
}

func Test_Yarn_GetVersions(t *testing.T) {
	versions, err := yarn.GetPublishedYarnVersions()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%d mc versions \n", len(versions))

	for mcVer, yarnVer := range versions {
		fmt.Printf("\t %s.%d \n", mcVer, yarnVer)
	}

}

func Test_Yarn_Intermeidarys(t *testing.T) {
	data := tiny.ReadIntermediary("18w50a")

	fmt.Printf("%d classes\n", len(data.Classes))
	fmt.Printf("%d methods\n", len(data.Methods))
	fmt.Printf("%d fields\n", len(data.Fields))

	fmt.Printf("%s -> %s\n", data.Classes[0].ObofName, data.Classes[0].DeobfName)
	fmt.Printf("%s -> %s\n", data.Methods[0].ObofName, data.Methods[0].DeobfName)
	fmt.Printf("%s -> %s\n", data.Fields[0].ObofName, data.Fields[0].DeobfName)
}
