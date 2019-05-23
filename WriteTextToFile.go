package main

import (
	"io/ioutil"
	"log"
)

func main() {
	//write into file
	message := []byte("here is all of us")
	err := ioutil.WriteFile("yesf", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
