package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	infile, err := os.Open("./testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	defer infile.Close()

	outfile, err := os.Create("./testdata/output.go")
	if err != nil {
		t.Fatal(err)
	}
	defer outfile.Close()

	s := make(map[string]interface{})

	err = json.NewDecoder(infile).Decode(&s)

	if err != nil {
		t.Fatal(err)
	}

	err = Convert(outfile, s, "test")
	if err != nil {
		t.Fatal(err)
	}
}
