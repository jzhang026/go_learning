package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path ", r.URL.Path)
	fmt.Fprintf(w, "hello %s! | From URL: %s", r.URL.Path[1:], r.URL.Path)
}

func main() {
	http.HandleFunc("/",handler)
	if err := http.ListenAndServe("0.0.0.0:8083", nil); err != nil {
		log.Fatal("failed to start server")
	} else {
		log.Println("server listening on 8083")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("input: %s \n", string(data))
	input := &Input{}
	_ = json.Unmarshal(data, input)

	output, _ := json.Marshal(Output{
		Msg: "hello " + input.Msg,
	})

	fmt.Println(string(output))
	fmt.Fprintf(w, "%s", string(output))
}

type Input struct {
	Msg string
}

type Output struct {
	Msg string
}
