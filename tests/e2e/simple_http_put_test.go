package integrationtests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestSimpleHttpPut tests a simple HTTP PUT call using a stub file with HTTP
// verb explicitly set in filename
func TestSimpleHttpPut(t *testing.T) {
	url := fmt.Sprintf("%s/simple", ts.URL)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Custom-Type", "application/form-data")
	res, err := client.Do(req)

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
	if value["key"] != "simple.PUT.json" {
		t.Fail()
	}
}
