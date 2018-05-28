package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	fmt.Println("template: ", name)

	return tmpl.ExecuteTemplate(w, "base", data)
}

func llista(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./problemes")
	if err != nil {
		log.Fatal(err)
	}
	if err := renderTemplate(w, "llista_problemes.html", files); err != nil {
		log.Fatal(err)
	}
}

func problems(w http.ResponseWriter, r *http.Request) {
	prog := filepath.Base(r.URL.String())
	if prog == "problems" {
		llista(w, r)
	} else {
		upload(w, r)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		prog := filepath.Base(r.URL.String())
		fmt.Println(prog, get_name(prog))
		if !check_code(prog) {
			error_(w, r, "Invalid code")
		} else if err := renderTemplate(w, "upload.html", prog); err != nil {
			log.Fatal(err)
		}
	} else {
		r.ParseForm()
		input := r.FormValue("input")
		//prog := filepath.Base(r.FormValue("name"))
		prog := filepath.Base(r.URL.String())
		fmt.Println("prog:", prog)
		output, err := test(input, prog)
		if err != nil {
			error_(w, r, err.Error())
		} else {
			fmt.Fprintf(w, output)
		}
	}
}

func error_(w http.ResponseWriter, r *http.Request, message string) {
	if err := renderTemplate(w, "error.html", message); err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	if err := renderTemplate(w, "root.html", nil); err != nil {
		log.Fatal(err)
	}
}
