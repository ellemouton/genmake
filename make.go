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

	sFiles, err := checkExt(path, ".cpp")
	if err != nil {
		log.Fatal(err)
	}

	hFiles, err := checkExt(path, ".h")
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

func checkExt(path, ext string) ([]string, error) {
	var files []string
	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {
				files = append(files, f.Name()[:len(f.Name())-len(ext)])
			}
		}
		return nil
	})

	return files, nil
}
