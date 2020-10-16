package genmake

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	ExName string
	SFiles []string
	HFiles []string
}

func Generate(path, execName string) error {
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
	t.Parse(tmpl)

	data := Config{
		ExName: execName,
		SFiles: sFiles,
		HFiles: hFiles,
	}

	return t.Execute(f, data)
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
