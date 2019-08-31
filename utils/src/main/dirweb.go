package main

import (
	"net/http";
	"os"
)

func main() {
	port:=os.Args[1]
    http.Handle("/", http.FileServer(http.Dir("./www/")))
    http.ListenAndServe(":"+port, nil)
}