package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("/Users/rrli/go/src/github.com/reeli/go-static-server/static/index.html")
	check(err)
	w.Header().Set("Server", "Static Server")
	w.WriteHeader(200)
	w.Write(dat)
}

func Serve() {
	fmt.Printf("Start listening port 9000...")
	http.HandleFunc("/static", handler)
	http.ListenAndServe(":9000", nil)
}
