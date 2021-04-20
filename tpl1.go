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
	o := Outer{
		Inner: Inner{A : "123"},
	}
	tpl := template.Must(template.New("").Parse(`{{ .A }}`))
	tpl.Execute(os.Stdout, &o)
}