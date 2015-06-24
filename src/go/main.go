package main

import (
	"github.com/zenazn/goji"
	"io/ioutil"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("./assets/html/index.html")
    if err != nil {
        panic(err)
    }
	w.Write(data)
}

func main() {
	goji.Get("/", rootHandler)
	goji.Get("/*", http.FileServer(http.Dir("./assets")))
	goji.Serve()
}
