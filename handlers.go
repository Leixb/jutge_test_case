package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func llista(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/llista_problemes.html",
	))
	files, err := ioutil.ReadDir("./problemes")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, files)
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
		t := template.Must(template.ParseFiles(
			"static/templates/layout.html",
			"static/templates/upload.html",
		))
		prog := filepath.Base(r.URL.String())
		if err := t.Execute(w, prog); err != nil {
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
	t := template.Must(template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/error.html",
	))
	if err := t.Execute(w, message); err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/root.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}
