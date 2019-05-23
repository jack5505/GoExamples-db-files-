package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {

	buf, e := os.Create("yep.zip")
	if e != nil {
		log.Fatal(e)
	}
	w := zip.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "hello world"},
		{"gopher.txt", "helafafea salom"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err1 := f.Write([]byte(file.Body))
		if err1 != nil {
			log.Fatal(err1)
		}

	}
	err3 := w.Close()
	if err3 != nil {
		log.Fatal(err3)
	}

}
