package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://golang.org1")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
