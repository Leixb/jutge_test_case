package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
)

var code_matcher *regexp.Regexp
var problem_names map[string]interface{}
var templates map[string]*template.Template

func main() {
	portPtr := flag.Int("port", 8080, "Port to listen to")

	flag.Parse()

	port := fmt.Sprintf(":%d", *portPtr)

	h := http.NewServeMux()

	h.HandleFunc("/", root)
	h.HandleFunc("/problems/", problems)
	h.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	code_matcher = regexp.MustCompile("^[PX][0-9]{5}_(ca|en|es)$")

	layouts, err := filepath.Glob("./static/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	includes, err := filepath.Glob("./static/templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	templates = make(map[string]*template.Template)
	for _, layout := range layouts {
		files := append(includes, layout)
		log.Println("making", files)
		templates[filepath.Base(layout)] =
			template.Must(template.New("").Funcs(
				template.FuncMap{
					"get_name": get_name,
				}).ParseFiles(files...))
	}

	f, err := ioutil.ReadFile("./problems.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(f, &problem_names)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("#####################################")
	log.Println("# Serving on: http://localhost" + port)
	log.Println("#####################################")
	if err := http.ListenAndServe(port, h); err != nil {

		log.Fatal("ListenAndServe: ", err)
	}
}
