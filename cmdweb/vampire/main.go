package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/devinherron/dice"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="Text">
				<button type="submit">Upload Text File</button>
			</form>
			</body></html>`
		fmt.Fprint(w, html)
	})
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("text")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()
		// TODO: Actually use this
		ext := filepath.Ext(header.Filename)
		_ = ext
	})
	test := roll()
	fmt.Println(test)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func roll() int {
	_, d6Result := dice.Roll(10, 1)
	_, d10Result := dice.Roll(6, 1)

	return d10Result - d6Result
}
