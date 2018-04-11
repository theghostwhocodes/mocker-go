package integration_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/theghostwhocodes/mocker-go/internal/handlers"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	basePath, _ := filepath.Abs(path.Join("..", "data"))
	fmt.Printf("%v\n", basePath)
	ts = httptest.NewServer(http.HandlerFunc(handlers.HandlerFactory(basePath)))
	defer ts.Close()
	os.Exit(m.Run())
}

func TestSimpleHttpGet(t *testing.T) {
	url := fmt.Sprintf("%s/simple", ts.URL)
	res, err := http.Get(url)
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
	if value["key"] != "simple.GET.json" {
		t.Fail()
	}

}
