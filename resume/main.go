package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	out, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error in file creating", err)
	}
	defer out.Close()

	data := map[string]interface{}{}

	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal("error in file opening", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal("error in unmarshalling", err)
	}

	t, err := template.ParseGlob("template/*")
	if err != nil {
		log.Fatal("error in template parsing")
	}
	err = t.Execute(out, data)
	if err != nil {
		log.Fatal("error while executing", err)
	}
}
