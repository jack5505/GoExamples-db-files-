package main

import (
	"log"
	"os"
)

func main() {
	newZipFile, err := os.Create("yep")
	if err != nil {
		log.Fatal(err)
	}
	defer newZipFile.Close()

}
