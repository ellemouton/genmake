package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Make struct {
	ExName string
	SFiles []string
	HFiles []string
}

func checkExt(ext string) ([]string, error) {
	pathS, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {
				files = append(files, f.Name()[:len(f.Name())-len(ext)])
			}
		}
		return nil
	})

	return files, nil
}

func main() {
	name := "prog"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(path + "/Makefile")
	if err != nil {
		log.Fatal(err)
	}

	sFiles, err := checkExt(".cpp")
	if err != nil {
		log.Fatal(err)
	}

	hFiles, err := checkExt(".h")
	if err != nil {
		log.Fatal(err)
	}

	t := template.New("make template")
	t.Parse(makeTmpl)

	data := Make{
		ExName: name,
		SFiles: sFiles,
		HFiles: hFiles,
	}

	t.Execute(f, data)
}
