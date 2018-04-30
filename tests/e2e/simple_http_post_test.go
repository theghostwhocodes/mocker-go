package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

// TestSimpleHttpPostNoVerbInFile tests a simple HTTP POST call using a stub file with HTTP
// verb NOT explicitly set in filename
func TestSimpleHttpPostNoVerbInFile(t *testing.T) {
	url := fmt.Sprintf("%s/simplePOST", ts.URL)
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
	if value["key"] != "simplePOST.HTTP.json" {
		t.Fail()
	}
}

// TestSimpleHttpPostOneParam tests a simple HTTP POST call using a stub file with HTTP
// verb explicitly set in filename and one parameter in stub content
func TestSimpleHttpPostOneParam(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	res, err := http.Post(urlString, "application/x-www-form-urlencoded", postBody)
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
	if value["key"] != "simple.POST.param.json" {
		t.Fail()
	}
}

// TestSimpleHttpPostTwoParams tests a simple HTTP POST call using a stub file with HTTP
// verb explicitly set in filename and two parameters in stub content
func TestSimpleHttpPostTwoParams(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
		"param2": {"value2"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	res, err := http.Post(urlString, "application/x-www-form-urlencoded", postBody)
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
	if value["key"] != "simple.POST.twoparams.json" {
		t.Fail()
	}
}

func TestSimpleHttpPostOneParamTooMuch(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param1": {"value1"},
		"param3": {"value3"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	res, err := http.Post(urlString, "application/x-www-form-urlencoded", postBody)
	if err != nil {
		fmt.Printf("Errore %v", err)
		t.FailNow()
	}

	if res.StatusCode != 404 {
		t.FailNow()
	}

	defer res.Body.Close()
}

func TestSimpleHttpPostOneParamTooMuch2(t *testing.T) {
	urlString := fmt.Sprintf("%s/simple", ts.URL)
	form := url.Values{
		"param3": {"value3"},
	}
	postBody := bytes.NewBufferString(form.Encode())
	res, err := http.Post(urlString, "application/x-www-form-urlencoded", postBody)
	if err != nil {
		fmt.Printf("Errore %v", err)
		t.FailNow()
	}

	if res.StatusCode != 404 {
		t.FailNow()
	}

	defer res.Body.Close()
}
