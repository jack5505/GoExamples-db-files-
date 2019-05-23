package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	content := []byte("here is all of them")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(dir)
	}
	defer os.Remove(dir)

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0644); err != nil {
		log.Fatal(err)
	}
}
