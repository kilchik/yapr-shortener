package main

import (
	"io/ioutil"
	"net/http"
)

const addr = "localhost:8080"

func main() {
	urls := make(map[string]string)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		url := string(bb)
		urls["abc"] = url

		w.WriteHeader(http.StatusCreated)
	})

	http.ListenAndServe(addr, nil)
}
