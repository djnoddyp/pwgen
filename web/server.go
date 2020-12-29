package web

import (
	"net/http"
	"io"
	"github.com/djnoddyp/pwgen/gen"
)

func Start() {
	pwgenHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, gen.Generate() + "\n")
	}
	http.HandleFunc("/pwgen", pwgenHandler)
	http.ListenAndServe(":8080", nil)
}