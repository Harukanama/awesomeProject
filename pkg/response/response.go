package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("https://golangcode.com")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("HTTP response status", resp.StatusCode)
}