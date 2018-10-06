//go:generate esc -ignore=assets.go -pkg=assets -prefix=assets/ -o assets/assets.go assets/
package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/satnamram/flexkit/cmd/assets"
)

func usage() {
	println("usage: flexkit <app.js> <address>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 || os.Args[1] == "help" {
		usage()
	}

	assetFS := http.FileServer(assets.FS(false))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/app.js" {
			w.Header().Set("Content-Type", "application/javascript")
			b, err := ioutil.ReadFile(os.Args[1])
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			w.Write(b)
			return
		}
		assetFS.ServeHTTP(w, r)
	}))

	err := http.ListenAndServe(os.Args[2], nil)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
