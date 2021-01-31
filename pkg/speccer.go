package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"
)

func main() {
	in := "templates/spec.md"
	out := "routes/spec.md"

	tmpl, err := template.ParseFiles(in)
	if err != nil {
		panic(err)
	}

	url := "https://raw.githubusercontent.com/traindown/spec/main/README.md"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	vars := struct {
		Spec string
	}{string(body)}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, vars)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(out, buf.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}
