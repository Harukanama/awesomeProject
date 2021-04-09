package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/health HTTP/1.1", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	http.ListenAndServe(":8000", nil)

	resp, err := http.Get("https://golangcode.com")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("HTTP response status", resp.StatusCode)

}
