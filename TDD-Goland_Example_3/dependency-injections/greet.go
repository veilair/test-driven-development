package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingsHandler(w http.ResponseWriter, req *http.Request) {
	Greet(w, "world")
}
func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingsHandler))
}
