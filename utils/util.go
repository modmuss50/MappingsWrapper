//This is just a random collection of useful stuff that doesnt really have anything to do with mcp
package utils

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

func DownloadFile(url string, file string) error {
	bytes, err := Download(url)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Download(url string) ([]byte, error) {
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

func DownloadString(url string) (string, error) {
	bytes, err := Download(url)
	return string(bytes), err
}

func MakeDir(fileName string) {
	os.MkdirAll(fileName, os.ModePerm)
}

func GetRunPath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)
	return exPath
}

func DeleteDir(dir string) error {
	if !FileExists(dir) {
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

func DeleteFile(path string) error {
	var err = os.Remove(path)
	return err
}

func FileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func CopyFile(src string, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ExtractZip(zip string, dest string) error {
	return archiver.DefaultZip.Unarchive(zip, dest)
}

func WriteStringToFile(str string, file string) error {
	return ioutil.WriteFile(file, []byte(str), os.ModePerm)
}

func ReadStringFromFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func ReadLinesFromFile(fileName string) []string {
	return strings.Split(ReadStringFromFile(fileName), "\n")
}

func DivideString2(str string) (string, string) {
	return SplitString2(str, " ")
}

func SplitString2(str string, sep string) (string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1]
}

func DivideString3(str string) (string, string, string) {
	return SplitString3(str, " ")
}

func SplitString3(str string, sep string) (string, string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1], split[2]
}

func DivideString4(str string) (string, string, string, string) {
	return SplitString4(str, " ")
}

func SplitString4(str string, sep string) (string, string, string, string) {
	split := strings.Split(str, sep)
	return split[0], split[1], split[2], split[3]
}

func PrintAsJson(v interface{}) {
	fmt.Println(ToJson(v))
}

func ToJson(v interface{}) string {
	json, _ := json.Marshal(v)
	return string(json)
}

func SplitAtLastSlash(input string) (string, string) {
	lastPos := strings.LastIndex(input, "/")
	return input[:lastPos], input[lastPos+1:]
}
