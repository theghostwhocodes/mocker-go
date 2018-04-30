package integrationtests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestSimpleHttpPost tests a simple HTTP POST call using a stub file with HTTP
// verb explicitly set in filename
func TestSimpleHttpPost(t *testing.T) {
	url := fmt.Sprintf("%s/simple", ts.URL)
	var buf io.Reader
	res, err := http.Post(url, "application/form-data", buf)
	if err != nil {
		t.Fail()
	}

	if res.StatusCode != 200 {
		t.Fail()
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}

	var content interface{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		t.Fail()
	}

	value := content.(map[string]interface{})
	if value["key"] != "simple.POST.json" {
		t.Fail()
	}
}
