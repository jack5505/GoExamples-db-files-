package main

import (
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("yesf")
	if err != nil {
		log.Fatal(err)
	}
	err1 := ioutil.WriteFile("output", content, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
}
