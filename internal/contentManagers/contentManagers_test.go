package contentManagers_test

import (
	"path"
	"testing"

	"github.com/theghostwhocodes/mocker-go/internal/contentManagers"
)

func TestGetFileNameHttpGet(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "GET")
	expectedFileName := "/folder/mock.GET.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPost(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "POST")
	expectedFileName := "/folder/mock.POST.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPut(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "PUT")
	expectedFileName := "/folder/mock.PUT.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpPatch(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "PATCH")
	expectedFileName := "/folder/mock.PATCH.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpOptions(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "OPTIONS")
	expectedFileName := "/folder/mock.OPTIONS.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetFileNameHttpDelete(t *testing.T) {
	fileName := contentManagers.GetFileName("/folder/mock", "DELETE")
	expectedFileName := "/folder/mock.DELETE.json"
	if fileName != expectedFileName {
		t.Fail()
	}
}

func TestGetAbsoluteFileNameHttpGet(t *testing.T) {
	fileName := "/folder/mock.GET.json"
	basePath := "/var/www/"
	absoluteFileName := contentManagers.GetAbsoluteFileName(basePath, fileName)
	expectedAbsoluteFileName := path.Join(
		basePath,
		fileName,
	)
	if absoluteFileName != expectedAbsoluteFileName {
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

	bodyContent, err := contentManagers.GetBodyContent(jsonMap)
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

	bodyContent, err := contentManagers.GetBodyContent(jsonMap)

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

	bodyContent, err := contentManagers.GetBodyContent(jsonMap)

	if bodyContent != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}
