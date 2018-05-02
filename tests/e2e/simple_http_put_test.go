package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

// TestSimpleHttpPutNoVerbInFile tests a simple HTTP PUT call using a stub file with HTTP
// verb NOT explicitly set in filename
func TestSimpleHttpPutNoVerbInFile(t *testing.T) {
	url := fmt.Sprintf("%s/simplePUT", ts.URL)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	if value["key"] != "simplePUT.HTTP.json" {
		t.Fail()
	}
}

// TestSimpleHttpPutOneParam tests a simple HTTP PUT call using a stub file with HTTP
// verb explicitly set in filename and one parameter in stub content
func TestSimpleHttpPutOneParam(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlString, postBody)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	fmt.Printf("%v", value)
	if value["key"] != "simple.PUT.param.json" {
		t.Fail()
	}
}

// TestSimpleHttpPutTwoParams tests a simple HTTP PUT call using a stub file with HTTP
// verb explicitly set in filename and two parameters in stub content
func TestSimpleHttpPutTwoParams(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
		"param2": {"value2"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlString, postBody)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
	fmt.Printf("%v", value)
	if value["key"] != "simple.PUT.twoparams.json" {
		t.Fail()
	}
}

func TestSimpleHttpPutOneParamTooMuch(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
		"param3": {"value3"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlString, postBody)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		t.FailNow()
	}

	if res.StatusCode != 404 {
		t.FailNow()
	}

	defer res.Body.Close()
}

func TestSimpleHttpPutOneParamTooMuch2(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param3": {"value3"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlString, postBody)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		t.FailNow()
	}

	if res.StatusCode != 404 {
		t.FailNow()
	}

	defer res.Body.Close()
}

func TestSimpleHttpPutSimpleNotFound(t *testing.T) {
	urlString := fmt.Sprintf("%s/simpleNotFound", ts.URL)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", urlString, nil)
	if err != nil {
		t.Fail()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		t.FailNow()
	}

	if res.StatusCode != 404 {
		t.FailNow()
	}

	defer res.Body.Close()
}
