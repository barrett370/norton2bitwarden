package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/barrett370/norton2bitwarden/formats"
)

var (
	inputFile  string
	outputFile string
)

func main() {
	flag.StringVar(&inputFile, "input", "input.csv", "norton password export")
	flag.StringVar(&outputFile, "output", "output.json", "file to write bitwarden json to")
	flag.Parse()
	passwords, err := formats.DecodeNortonExport(inputFile)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	bf, err := formats.BitwardenFileFromNorton(passwords)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(out).Encode(bf)

	if err != nil {
		panic(err)
	}
}
