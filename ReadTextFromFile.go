package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//This read string from file
	content, err := ioutil.ReadFile("yesf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
