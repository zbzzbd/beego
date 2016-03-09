package progress

import (
	"bytes"
	"os"
	"path"
	"text/template"
)

var baseViewPath string
var cardPath string

func init() {
	baseViewPath, _ := os.Getwd()
	cardPath = path.Join(baseViewPath, "views/cards/")
}

func render(tpl string, data interface{}) (string, error) {
	name := path.Base(tpl)
	buff := bytes.NewBufferString("")

	t, err := template.New(name).ParseFiles(tpl)
	if err != nil {
		return "", err
	}
	if err := t.Execute(buff, data); err != nil {
		return "", err
	}
	return buff.String(), nil
}
