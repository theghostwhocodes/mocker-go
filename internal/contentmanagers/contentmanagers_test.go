package contentmanagers

import (
	"os"
	"testing"
	"time"
)

func TestGetFileNameHttpGet(t *testing.T) {
	fileName := GetFileName("/folder/mock", "GET")
	expectedFileName := "/folder/mock.GET.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPost(t *testing.T) {
	fileName := GetFileName("/folder/mock", "POST")
	expectedFileName := "/folder/mock.POST.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPut(t *testing.T) {
	fileName := GetFileName("/folder/mock", "PUT")
	expectedFileName := "/folder/mock.PUT.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPatch(t *testing.T) {
	fileName := GetFileName("/folder/mock", "PATCH")
	expectedFileName := "/folder/mock.PATCH.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpOptions(t *testing.T) {
	fileName := GetFileName("/folder/mock", "OPTIONS")
	expectedFileName := "/folder/mock.OPTIONS.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpDelete(t *testing.T) {
	fileName := GetFileName("/folder/mock", "DELETE")
	expectedFileName := "/folder/mock.DELETE.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetDirNameSimpleDir(t *testing.T) {
	dirName := GetDirName("/folder/mock")
	expectedDirName := "/folder"
	if dirName != expectedDirName {
		t.Fail()
	}
}

func TestGetDirNameMultipleDir(t *testing.T) {
	dirName := GetDirName("/folder/subfolder/mock")
	expectedDirName := "/folder/subfolder"
	if dirName != expectedDirName {
		t.Fail()
	}
}

func TestGetDirNameRootDir(t *testing.T) {
	dirName := GetDirName("/folder")
	expectedDirName := ""
	if dirName != expectedDirName {
		t.Fail()
	}
}

func TestGetResourceNameSimpleDir(t *testing.T) {
	resourceName := GetResourceName("/folder/mock")
	expectedResourceName := "mock"
	if resourceName != expectedResourceName {
		t.Fail()
	}
}

func TestGetResourceNameMultipleDir(t *testing.T) {
	resourceName := GetResourceName("/folder/subfolder/mock")
	expectedResourceName := "mock"
	if resourceName != expectedResourceName {
		t.Fail()
	}
}

func TestGetResourceNameRootDir(t *testing.T) {
	resourceName := GetResourceName("/folder")
	expectedResourceName := "folder"
	if resourceName != expectedResourceName {
		t.Fail()
	}
}

type mockedFile struct {
	FileName string
	Dir      bool
}

func (m mockedFile) Name() string {
	return m.FileName
}

func (m mockedFile) Size() int64 {
	return 1
}

func (m mockedFile) Mode() os.FileMode {
	return 1
}

func (m mockedFile) ModTime() time.Time {
	return time.Now()
}

func (m mockedFile) IsDir() bool {
	return m.Dir
}

func (m mockedFile) Sys() interface{} {
	return nil
}

func TestGetMockFiles(t *testing.T) {
	m1 := mockedFile{FileName: "test1.GET.json", Dir: false}
	m2 := mockedFile{FileName: "test2.GET.001.json", Dir: false}
	m3 := mockedFile{FileName: "test2.GET.002.json", Dir: false}
	m4 := mockedFile{FileName: "test3.GET.json", Dir: false}
	m5 := mockedFile{FileName: "test2", Dir: true}
	m6 := mockedFile{FileName: "test4.GET.json", Dir: false}
	files := []os.FileInfo{m1, m2, m3, m4, m5, m6}

	mockFiles := GetMockFiles(files, "test2", "GET")

	if len(mockFiles) != 2 {
		t.Fail()
	}
}

func TestGetMockFilesHTTP01(t *testing.T) {
	m1 := mockedFile{FileName: "test1.GET.json", Dir: false}
	m2 := mockedFile{FileName: "test2.GET.001.json", Dir: false}
	m3 := mockedFile{FileName: "test2.GET.002.json", Dir: false}
	m4 := mockedFile{FileName: "test3.HTTP.json", Dir: false}
	m5 := mockedFile{FileName: "test2", Dir: true}
	m6 := mockedFile{FileName: "test4.GET.json", Dir: false}
	files := []os.FileInfo{m1, m2, m3, m4, m5, m6}

	mockFiles := GetMockFiles(files, "test3", "POST")

	if len(mockFiles) != 1 {
		t.Fail()
	}
}

func TestGetMockFilesHTTP02(t *testing.T) {
	m1 := mockedFile{FileName: "test1.GET.json", Dir: false}
	m2 := mockedFile{FileName: "test2.GET.json", Dir: false}
	m3 := mockedFile{FileName: "test2.HTTP.json", Dir: false}
	m4 := mockedFile{FileName: "test2.POST.json", Dir: false}
	m5 := mockedFile{FileName: "test2", Dir: true}
	m6 := mockedFile{FileName: "test4.GET.json", Dir: false}
	m7 := mockedFile{FileName: "test2.GET.001.json", Dir: false}
	files := []os.FileInfo{m1, m2, m3, m4, m5, m6, m7}

	mockFiles := GetMockFiles(files, "test2", "GET")

	if len(mockFiles) != 3 {
		t.Fail()
	}
}

func TestGetBodyContentSuccess(t *testing.T) {
	jsonMap := map[string]interface{}{
		"request": nil,
		"response": map[string]interface{}{
			"body": "This is the response body",
		},
	}

	bodyContent, err := GetBodyContent(jsonMap)
	bodyContentString := string(bodyContent)

	if bodyContentString != "\"This is the response body\"" {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestGetBodyContentResponseNil(t *testing.T) {
	jsonMap := map[string]interface{}{
		"request":  nil,
		"response": map[string]interface{}{},
	}

	bodyContent, err := GetBodyContent(jsonMap)

	if bodyContent != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestGetBodyContentNoBody(t *testing.T) {
	jsonMap := map[string]interface{}{
		"request": nil,
		"response": map[string]interface{}{
			"key": "value",
		},
	}

	bodyContent, err := GetBodyContent(jsonMap)

	if bodyContent != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}
