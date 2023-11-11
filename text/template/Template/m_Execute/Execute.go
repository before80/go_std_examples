package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmplText := `
{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`
	tmpl, err := template.New("titleTest").Parse(tmplText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "T2", "no data needed")
	if err != nil {
		log.Fatalf("execution failed: %s", err)
	}
}

// Output:
// TWO
