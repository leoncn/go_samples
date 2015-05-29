// main.go
package main

import (
	"log"
	"os"
	"samples/templ"
)

func main() {

	job1 := templ.Job{Employer: "Monash", Role: "Honorary"}
	job2 := templ.Job{Employer: "Box Hill", Role: "Head of HE"}

	p := &templ.Person{Name: "Leon",
		Age:    50,
		Emails: []string{"mail@tr", "mail@ibm"},
		Jobs:   []*templ.Job{&job1, &job2},
	}

	if err := p.PrintBasic(os.Stdout); err != nil {
		log.Fatal(err.Error())
	}
}
