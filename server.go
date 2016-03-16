package main

import (
	"fmt"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n\tHello World From Golang on Heroku")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server Listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic("ListenAndServe")
	}
}

