package integrationtests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSimpleHttpPostSingleHeader(t *testing.T) {
	url := fmt.Sprintf("%s/simple", ts.URL)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("X-Custom-Header", "Custom header")
	res, err := client.Do(req)

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
	if value["key"] != "simple.POST.singleheader.json" {
		t.Fail()
	}
}
