package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n\tHello World From Golang on Heroku\n\n")
}
func speedHandler(w http.ResponseWriter, r *http.Request) {
	speed_test_len := r.URL.Query().Get("size")
	n, err:= strconv.Atoi(speed_test_len)
	if err != nil{
		n = 10
	}
	n *= 1024*1024

	w.Header().Set("Content-Length", strconv.Itoa(n))
	w.Header().Set("Content-Type", "application/octet-stream")
	buf := make([]byte, 4096, 4096)
	for n >= 4096 {
		_, err := w.Write(buf)
		if err != nil {
			break
		}
		n -= 4096
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/speed", speedHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server Listening on " + port)

	err := http.ListenAndServe(":"+port,  nil)
	if err != nil {
		panic("ListenAndServe")
	}
}

