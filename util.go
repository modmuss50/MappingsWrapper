//This is just a random collection of useful stuff that doesnt really have anything to do with mcp
package mcpwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mholt/archiver"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func downloadFile(url string, file string) error {
	bytes, err := download(url)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func download(url string) ([]byte, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return nil, err2
		}
		return bodyBytes, nil
	}

	return nil, errors.New(resp.Status + " : " + url)
}

func makeDir(fileName string) {
	os.MkdirAll(fileName, os.ModePerm)
}

func getRunPath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)
	return exPath
}

func deleteDir(dir string) error {
	if !fileExists(dir) {
		return errors.New("File not found")
	}
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteFile(path string) error {
	var err = os.Remove(path)
	return err
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func extractZip(zip string, dest string) error {
	return archiver.Zip.Open(zip, dest)
}

func writeStringToFile(str string, file string) error {
	return ioutil.WriteFile(file, []byte(str), os.ModePerm)
}

func readStringFromFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func divideString2(str string) (string, string) {
	return splitString2(str, " ")
}

func splitString2(str string, sep string) (string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1]
}

func divideString3(str string) (string, string, string) {
	return splitString3(str, " ")
}

func splitString3(str string, sep string) (string, string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1], split[2]
}

func divideString4(str string) (string, string, string, string) {
	return splitString4(str, " ")
}

func splitString4(str string, sep string) (string, string, string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1], split[2], split[3]
}

func printAsJson(v interface{}) {
	fmt.Println(toJson(v))
}

func toJson(v interface{}) string {
	json, _ := json.Marshal(v)
	return string(json)
}

func splitAtLastSlash(input string) (string, string) {
	lastPos := strings.LastIndex(input, "/")
	return input[:lastPos], input[lastPos+1:]
}
