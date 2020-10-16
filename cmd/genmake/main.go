package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ellemouton/genmake"
)

func main() {
	name := "prog"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	err = genmake.Generate(path, name)
	if err != nil {
		log.Fatal(err)
	}
}
