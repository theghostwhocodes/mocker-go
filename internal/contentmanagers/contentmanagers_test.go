package contentmanagers

import (
	"testing"
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
