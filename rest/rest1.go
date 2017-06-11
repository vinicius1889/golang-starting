package main

import (
"encoding/json"
"net/http"
)

type Usuario struct {
	Nome string
	Sobrenome string
}

type Message struct {
	Text string
}

func about (w http.ResponseWriter, r *http.Request) {
	m := Usuario{"Welcome to the SandovalEffect API, build v0.0.001.992, 6/22/2015 0340 UTC.","Welcome to the SandovalEffect API, build v0.0.001.992, 6/22/2015 0340 UTC."}
	b, err := json.Marshal(m)
	if err!=nil{
		panic(err)
	}

	w.Write(b)
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}


