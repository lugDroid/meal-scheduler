package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"lugdroid/mealsScheduler/webapp/controller"
	"net/http"
	"os"
)

const port = ":8000"

func main() {
	templates := populateTemplates()
	controller.StartUp(templates)

	fmt.Println("Listening on http://localhost" + port)
	fmt.Println("Press Control+C to stop")

	http.ListenAndServe(port, nil)
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "../../templates"

	layout, err := template.ParseFiles(basePath + "/_layout.html")
	if err != nil {
		log.Fatal("Could not open layout template: " + err.Error())
	}

	layout, err = layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html")
	if err != nil {
		log.Fatal(("Could not parse template file: " + err.Error()))
	}

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		log.Fatal("Failed to open templates content directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			log.Fatal("Failed to open template: '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal("Failed to read content from file '" + fi.Name() + "'")
		}

		f.Close()

		template, err := layout.Clone()
		if err != nil {
			log.Fatal("Could not clone layout file " + err.Error())
		}

		_, err = template.Parse(string(content))
		if err != nil {
			log.Fatal("Failed to parse contents of '" + fi.Name() + "'")
		}

		result[fi.Name()] = template
	}

	return result
}
