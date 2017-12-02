package contentManagers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func GetContent(basePath string, r *http.Request) ([]byte, error) {
	method := r.Method
	fileName := fmt.Sprintf("%s.%s.json", r.URL.Path[1:], strings.ToUpper(method))
	content, err := ioutil.ReadFile(
		path.Join(
			basePath,
			fileName,
		),
	)
	return content, err
}
