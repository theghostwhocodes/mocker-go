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

// ScanMockFilesContent scans the mock file content and return an array of mocks
func ScanMockFilesContent(basePath string, dirName string, fileNames []string) ([]MockHTTP, error) {
	var results []MockHTTP
	for _, filename := range fileNames {
		content, err := ioutil.ReadFile(path.Join(basePath, dirName, filename))
		if err != nil {
			return nil, err
		}

		var jsonContent MockHTTP
		err = json.Unmarshal(content, &jsonContent)
		if err != nil {
			return nil, err
		}

		results = append(results, jsonContent)
	}

	return results, nil
}

func checkArrayEquality(array1 []string, array2 []string) bool {
	result := false
	if len(array1) != len(array2) {
		return result
	}

	for index, value := range array1 {
		if value == array2[index] {
			result = true
		} else {
			result = false
		}
	}

	return result
}

func FilterMockHeaderContent(mocks []MockHTTP, r *http.Request) (MockHTTP, error) {
	var result MockHTTP
	maxMatch := 0
	for _, mock := range mocks {
		matchCounter := 0
		for key, values := range r.Header {
			// fmt.Printf("Key %s - Value %s\n", key, value)
			mockValue, ok := mock.Request.Headers[key]
			if !ok {
				continue
			}

			if checkArrayEquality(mockValue, values) {
				matchCounter++
			}
		}

		if matchCounter > maxMatch {
			result = mock
			maxMatch = matchCounter
		}
	}

	return result, nil
}

// GetScannedMockContent return the mock content in form of a MockHTTP struct
func GetScannedMockContent(basePath string, r *http.Request) ([]MockHTTP, error) {
	urlPath := r.URL.Path[1:]
	dirName := GetDirName(urlPath)
	fileInfos, err := ioutil.ReadDir(path.Join(basePath, dirName))

	if err != nil {
		return nil, err
	}

	resourceName := GetResourceName(r.URL.Path[1:])
	mockFiles := GetMockFiles(fileInfos, resourceName, r.Method)
	results, err := ScanMockFilesContent(basePath, dirName, mockFiles)

	if err != nil {
		return nil, err
	}

	return results, nil
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
