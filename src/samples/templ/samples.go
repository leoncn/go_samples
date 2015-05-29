package templ

import (
	"io"
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

func (p *Person) PrintBasic(w io.Writer) error {
	t := template.New("Basic")
	if t, err := t.Parse(basic); err != nil {
		return err
	} else {
		return t.Execute(w, p)
	}
}
