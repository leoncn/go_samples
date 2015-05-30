package templ

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

const (
	basic = `The name is {{.Name}}.
    The age is {{.Age}}.
    {{ range .Emails}}
        An email is {{.}}
    {{end}}
    
    {{with .Jobs}}
        {{range .}}
            An employer is {{.Employer}}
            and the role is {{.Role}}
        {{end}}
    {{end}}
    `

	emailExpand = `The name is {{.Name}} 
	{{ range .Emails }}
		An email addr is {{ . | emailExpand }}
	{{ end }}`

	simpleVar = `{{ $name := .Name }}
		{{ range .Emails }}
			{{$name}} has a email {{.}}
		{{ end }}`
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

func (p *Person) rendTemplate(templ string, funcs template.FuncMap, w io.Writer) error {
	t := template.New(templ)

	if funcs != nil {
		t = t.Funcs(funcs)
	}

	if t, err := t.Parse(templ); err != nil {
		return err
	} else {
		return t.Execute(w, p)
	}

}

func (p *Person) PrintBasic(w io.Writer) error {
	//return p.rendTemplate("basic", nil, w)
	t := template.New("basic")
	if t, err := t.Parse(basic); err != nil {
		return err
	} else {
		return t.Execute(w, p)
	}
}

func (p *Person) PrintEmail(w io.Writer) error {
	t := template.Must(template.New("PrtEmail").Funcs(template.FuncMap{
		"emailExpand": func(args ...interface{}) string {
			ok := false
			var s string
			if len(args) == 1 {
				s, ok = args[0].(string)
			}

			if !ok {
				s = fmt.Sprint(args)
			}

			return strings.Replace(s, "@", "at", 3)

		}}).Parse(emailExpand))

	return t.Execute(w, p)
}

func (p *Person) PrintVar(w io.Writer) error {
	t := template.Must(template.New("simpleVar").Parse(simpleVar))
	return t.Execute(w, p)
}
