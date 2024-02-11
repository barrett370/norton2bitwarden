package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/barrett370/norton2bitwarden/formats"
)

func main() {
	passwords, err := formats.DecodeNortonExport("./test.csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", passwords)

	out, err := os.Create("out.json")
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
