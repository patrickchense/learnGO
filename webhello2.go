package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request)  {
	remPartOfURL := req.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello,"+ remPartOfURL)
}

func shouthello(w http.ResponseWriter, req *http.Request) {
	remPartOfURL := req.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "Hello,"+ strings.ToUpper(remPartOfURL))
}

func main()  {
	http.HandleFunc("/shouthello/", shouthello)
	http.HandleFunc("/hello/", hello)

	err := http.ListenAndServe("localhost:9999", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
