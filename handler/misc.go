package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Misc struct {
	l *log.Logger
}

func NewMisc(l *log.Logger) *Misc {
	return &Misc{l}
}

func (m *Misc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Something went wrong.", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(rw, "You're on another page")
}
