package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := Routes()

	fmt.Println("Server running localhost:8888")
	http.ListenAndServe(":8888", r)
}
