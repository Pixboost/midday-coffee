package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DIR: %s\n", wd)

	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(wd + "/dist/web"))))
}