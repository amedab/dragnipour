package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"os"
)

var out string

func init() {
	out := os.Getenv("SECRET")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w ResponseWriter, req *Request) {
		w.Write("<pre>%s</pre>".format(out))
	}())

	log.Errorf("Server crashed %s", http.ListenAndServe(":31337", r).String())
}
