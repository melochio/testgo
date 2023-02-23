package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gookit/goutil/dump"
)

type Person struct {
	Name string
	Age  string
}
type Baryabols struct {
	Status bool
	Info   []Person
}

func main() {
	fileServer := http.FileServer(http.Dir("./view"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	frmtmpl := template.Must(template.ParseFiles("./view/form.html"))
	http.HandleFunc("/forms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("nil was executed")
			frmtmpl.Execute(w, nil)
			return
		}
		fmt.Println("nil was not executed")
		details := Baryabols{
			Status: true,
			Info: []Person{
				{
					Name: r.FormValue("name1"),
					Age:  r.FormValue("age1"),
				},
				{
					Name: r.FormValue("name2"),
					Age:  r.FormValue("age2"),
				},
			},
		}
		dump.P(frmtmpl.Execute(w, details))
		if err := frmtmpl.Execute(w, details); err != nil {
			fmt.Println(err)
		}
	})
	// http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("asd")
	}
}
