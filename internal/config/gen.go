// Copyright 2020 The pdfcpu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
)

var debug = flag.Bool("debug", false, "")

func main() {
	flag.Parse()

	// Generate config.go
	{
		w := &bytes.Buffer{}
		w.WriteString(header)
		writeConfigBytes(w)
		finish(w, "config.go")
	}

}

const header = `// generated by "go run gen.go". DO NOT EDIT.

package config
`

func writeConfigBytes(w *bytes.Buffer) {
	bb, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	varDec := `// ConfigFileBytes is a byteslice representing config.yml.
	var ConfigFileBytes = []byte{
	`
	w.WriteString(varDec)
	for _, b := range bb {
		w.WriteString(fmt.Sprintf("%d,", int(b)))
	}
	w.WriteString("}\n\n")
}

func finish(w *bytes.Buffer, filename string) {
	if *debug {
		os.Stdout.Write(w.Bytes())
		return
	}
	out, err := format.Source(w.Bytes())
	if err != nil {
		log.Fatalf("format.Source: %v", err)
	}
	if err := ioutil.WriteFile(filename, out, 0660); err != nil {
		log.Fatalf("ioutil.WriteFile: %v", err)
	}
}
