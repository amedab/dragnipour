package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var out string

func init() {
	out = os.Getenv("SECRET")
}

func print(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("<pre>%s</pre>", out)))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", print)
	log.Errorf("Server crashed %s", http.ListenAndServe(":31337", r))
}
