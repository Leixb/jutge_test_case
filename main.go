package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var code_matcher *regexp.Regexp

func main() {
	portPtr := flag.Int("port", 8080, "Port to listen to")

	flag.Parse()

	port := fmt.Sprintf(":%d", *portPtr)

	h := http.NewServeMux()

	h.HandleFunc("/", root)
	h.HandleFunc("/problems/", problems)
	h.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	code_matcher = regexp.MustCompile("^[PX][0-9]{5}_(ca|en|es)$")

	fmt.Println("Serving on: http://localhost" + port)
	if err := http.ListenAndServe(port, h); err != nil {

		log.Fatal("ListenAndServe: ", err)
	}
}
