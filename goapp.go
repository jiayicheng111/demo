package main 

import (
	"fmt"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go Web Hello World!")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.ListenAndServe(":8083", nil)
}
