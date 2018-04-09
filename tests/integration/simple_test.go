package integration_test

import (
	"fmt"
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
	basePath, _ := filepath.Abs(path.Join("..", "..", "test_data"))
	fmt.Printf("%v\n", basePath)
	ts = httptest.NewServer(http.HandlerFunc(handlers.HandlerFactory(basePath)))
	defer ts.Close()
	os.Exit(m.Run())
}

func TestSimpleHttpGet(t *testing.T) {
	res, err := http.Get(ts.URL)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%v\n", ts.URL)
	fmt.Printf("%v\n", res)
}
