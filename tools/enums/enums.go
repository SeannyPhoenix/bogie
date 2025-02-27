package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type raw map[string]map[string]string

func main() {
	//load enums.json

	var enums raw
	fi, err := os.Open("enums.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	// read file
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal json
	err = json.Unmarshal(fd, &enums)
	if err != nil {
		log.Fatal(err)
	}

	writeGoFile(enums)
	writeSwiftFile(enums)
}

func writeGoFile(enums raw) {
	// generate enums.go
	fo, err := os.Create("out/enums_out.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	// write package
	fo.WriteString("package enums\n\n")

	for prefix, elements := range enums {
		fo.WriteString("const (\n")
		for name, value := range elements {
			fo.WriteString("\t" + prefix + name + " = \"" + value + "\"\n")
		}
		fo.WriteString(")\n\n")
	}
}

func writeSwiftFile(enums raw) {
	fo, err := os.Create("out/enums_out.swift")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	for prefix, elements := range enums {
		fo.WriteString("enum " + prefix + ": String {\n    case\n")
		for name, value := range elements {
			fo.WriteString("    " + name + " = \"" + value + "\"\n")
		}
		fo.WriteString("}\n\n")
	}
}
