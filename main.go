package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type msg struct {
	Name string `json:"NAME"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/status", statusHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"STATUS\": \"OK\"}")
}

func handler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil || len(body) <= 0 {
		fmt.Fprintf(w, "{ \"HELLO\": \"WORLD\"}")
		return
	}

	var input msg
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(body, &input)

	response := fmt.Sprint("{ \"HELLO\": \"", input.Name, "\"}")
	fmt.Fprintf(w, response)
}
