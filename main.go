package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type bldfunc func(io.Writer, map[string]interface{}, string) error

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file given")
		return
	}
	filename := os.Args[1]

	output_filename := strings.Replace(filename, ".json", ".go", 1)
	if len(os.Args) > 2 {
		output_filename = os.Args[2]
	}

	package_name := "default"

	wd, err := os.Getwd()
	if err == nil {
		package_name = filepath.Base(wd)

		if m, _ := filepath.Glob(wd + "/main.go"); len(m) > 0 {
			package_name = "main"
		}

	}

	var fnc bldfunc

	content := make(map[string]interface{})

	infile, err := os.Open(filename)
	defer infile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	outfile, err := os.Create(output_filename)
	defer outfile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	printheader(outfile, package_name)

	err = json.NewDecoder(infile).Decode(&content)
	if err != nil {
		log.Fatal(err)
		return
	}

	fnc = Convert

	name := "config"
	b := filepath.Base(filename)
	name = strings.Replace(b, ".json", "", -1)

	if fnc != nil {
		fnc(outfile, content, name)
	}
}

func printheader(outfile io.Writer, pn string) {
	fmt.Fprintln(outfile, "package", pn)
	fmt.Fprintln(outfile)
}
