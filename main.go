package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"os"
)

var out string

func init() {
	out = os.Getenv("SECRET")
}

func print(w ResponseWriter, req *Request) {
	w.Write("<pre>%s</pre>".format(out))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", print)
	log.Errorf("Server crashed %s", http.ListenAndServe(":31337", r).String())
}
