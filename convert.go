package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

func Convert(w io.Writer, s map[string]interface{}, name string) error {
	for i, r := range s {
		switch t := r.(type) {
		case map[string]interface{}:
			Convert(w, t, i)
		}
	}

	tpname := getTypeName(name)
	fmt.Fprintln(w, "//", tpname, "implements json structure for", name)
	fmt.Fprintln(w, "type", tpname, "struct {")

	for i, r := range s {
		jsoncomment := fmt.Sprintf("`json:\"%s,omitempty\"`", i)
		switch t := r.(type) {
		case map[string]interface{}:
			fmt.Fprintln(w, "  ", getTitleName(i), "*"+getTypeName(i), jsoncomment)
		case string:
			_, timeparseerror := time.Parse(time.RFC3339, t)
			switch {
			case timeparseerror == nil:
				fmt.Fprintln(w, "  ", getTitleName(i), "time.Time", jsoncomment)
			default:
				fmt.Fprintln(w, "  ", getTitleName(i), "string", jsoncomment)
			}
		case int64:
			fmt.Fprintln(w, "  ", getTitleName(i), "int", jsoncomment)
		case float64:
			if t == float64(int(t)) {
				fmt.Fprintln(w, "  ", getTitleName(i), "int", jsoncomment)
			} else {
				fmt.Fprintln(w, "  ", getTitleName(i), "float64", jsoncomment)
			}
		case bool:
			fmt.Fprintln(w, "  ", getTitleName(i), "bool", jsoncomment)
		case nil:

		}
	}

	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	return nil
}

func getTitleName(s string) string {
	toupper := true

	b := &bytes.Buffer{}

	for _, x := range s {
		if x == '_' {
			toupper = true
			continue
		}
		if toupper == true {
			fmt.Fprint(b, strings.ToUpper(string(x)))
			toupper = false
		} else {
			fmt.Fprint(b, string(x))
		}
	}

	return b.String()
}

func getTypeName(s string) string {
	return strings.ToUpper(s)
}
