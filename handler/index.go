package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger
}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func (i *Index) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("./handler/test.html")

	if err != nil {
		http.Error(rw, "Something went wrong.", http.StatusBadRequest)

		i.l.Println(err.Error())
		return
	}

	fmt.Fprintln(rw, string(f))
}
