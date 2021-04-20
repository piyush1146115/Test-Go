package main

import (
	"text/template"
	"os"
)

func main() {
	type Inner struct {
		A string
	}
	type Outer struct {
		Inner
	}

	type NA struct {
		O []Outer
	}

	na := NA{
		O: []Outer{
			{
				Inner: Inner{A: "1"},
			},
			{
				Inner: Inner{A: "2"},
			},
			{
				Inner: Inner{A: "3"},
			},
			{
				Inner: Inner{A: "4"},
			},
		},
	}
	tpl := template.Must(template.New("").Parse(`
{{- range $svc := .O }}
      {{ printf "%v" .A -}}
{{ end -}}
{{- range $svc := .O }}
    {{ printf "%v" .A -}}
{{ end -}}
`))
	tpl.Execute(os.Stdout, &na)
}