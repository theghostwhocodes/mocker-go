package contentmanagers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
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

// GetAbsoluteFileName return the full mock filename
func GetAbsoluteFileName(basePath string, fileName string) string {
	return path.Join(
		basePath,
		fileName,
	)
}

func GetMockFiles(dirPath string, urlPath string) {

}

// GetContent return the mock content
func GetContent(basePath string, r *http.Request) ([]byte, error) {
	fileName := GetFileName(r.URL.Path[1:], r.Method)
	content, err := ioutil.ReadFile(GetAbsoluteFileName(basePath, fileName))
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
