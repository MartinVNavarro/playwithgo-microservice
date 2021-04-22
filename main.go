package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/misc", Misc)

	http.ListenAndServe(":9090", nil)
}

func Index(rw http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something went wrong.", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(rw, "Hello World")
}

func Misc(rw http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something went wrong.", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(rw, "You're on another page")
}
