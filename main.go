package main

import (
	"fmt"
	"io"
	"net/http"
	"math/rand"
)

func callMathApi(w http.ResponseWriter, req *http.Request) {
	num := rand.Intn(100)
	resp, err := http.Get(fmt.Sprintf("https://api.mathjs.org/v4/?expr=log(%d)", num))
	if err != nil {
		fmt.Fprintf(w, "error calling mathapi: %v", err)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "error reading result from mathapi: %v", err)
		}
		fmt.Fprintf(w, "Math API result: log(%d) = %s", num, body)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/mathapi", callMathApi)

	http.ListenAndServe(":8080", nil)
}
