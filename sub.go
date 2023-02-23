package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	frmtmpl := template.Must(template.ParseFiles("./view/form.html"))
// 	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			frmtmpl.Execute(w, nil)
// 			return
// 		}
// 		details := baryabols{
// 			status: true,
// 			info: []person{
// 				{
// 					name: r.FormValue("name1"),
// 					age:  r.FormValue("age1"),
// 				},
// 				{
// 					name: r.FormValue("name2"),
// 					age:  r.FormValue("age2"),
// 				},
// 			},
// 		}

// 		frmtmpl.Execute(w, details)
// 	})
// 	// if r.Method == "GET" {
// 	// 	http.ServeFile(w, r, "./view/form.html")
// 	// 	return
// 	// }
// 	// if r.Method == "POST" {
// 	// 	if err := r.ParseForm(); err != nil {
// 	// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 	// 		return
// 	// 	}
// 	// 	fmt.Fprintf(w, "Post request success \n")
// 	// 	x := person{name: r.FormValue("name"), age: r.FormValue("age")}
// 	// 	fmt.Fprintf(w, "Name: %s\n", x.name)
// 	// 	fmt.Fprintf(w, "Age: %s\n", x.age)
// 	// }
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "method not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintf(w, "hello")
// }
