package contentmanagers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

// GetFileName return the computed file name for the mock
func GetFileName(urlPath string, method string) string {
	return fmt.Sprintf("%s.%s.json", urlPath, strings.ToUpper(method))
}

// GetDirName return the computed dir name for the mock
func GetDirName(urlPath string) string {
	urlPathComponents := strings.Split(urlPath, string(os.PathSeparator))
	dirComponents := urlPathComponents[:len(urlPathComponents)-1]
	dirPath := strings.Join(dirComponents, string(os.PathSeparator))
	return dirPath
}

// GetResourceName return the computed resource name (filename) for the mock
func GetResourceName(urlPath string) string {
	urlPathComponents := strings.Split(urlPath, string(os.PathSeparator))
	resource := urlPathComponents[len(urlPathComponents)-1]
	return resource
}

// GetMockFiles return filenames of
func GetMockFiles(fileInfos []os.FileInfo, resourceName string, httpVerb string) []string {
	var files []string
	for _, element := range fileInfos {
		if !element.IsDir() {
			elementName := element.Name()
			re := fmt.Sprintf(`^%s\.(%s|HTTP)\.`, resourceName, httpVerb)
			validName := regexp.MustCompile(re)
			hasMock := validName.MatchString(elementName)
			if hasMock {
				files = append(files, elementName)
			}
		}
	}

	return files
}

// GetContent return the mock content
func GetContent(basePath string, r *http.Request) ([]byte, error) {
	urlPath := r.URL.Path[1:]
	dirName := GetDirName(urlPath)
	fileInfos, err := ioutil.ReadDir(path.Join(basePath, dirName))

	if err != nil {
		fmt.Printf("%v", err)
	}

	fileName := GetFileName(r.URL.Path[1:], r.Method)
	resourceName := GetResourceName(r.URL.Path[1:])
	GetMockFiles(fileInfos, resourceName, r.Method)
	content, err := ioutil.ReadFile(path.Join(basePath, fileName))
	return content, err
}

// GetBodyContent return the stub body content
func GetBodyContent(jsonMap map[string]interface{}) ([]byte, error) {
	response, ok := jsonMap["response"]
	responseBody := response.(map[string]interface{})
	body, ok := responseBody["body"]
	if !ok {
		return nil, errors.New("You must specify a body key")
	}
	bodyBytes, error := json.Marshal(body)
	return bodyBytes, error
}
