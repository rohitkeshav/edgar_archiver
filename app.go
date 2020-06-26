package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://my-json-server.typicode.com/typicode/demo/posts")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
