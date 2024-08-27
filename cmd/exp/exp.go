package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("teste").Parse("<h1>Título da página</h1>")
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
