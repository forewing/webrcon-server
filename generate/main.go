// This program must be run from `../` as `go run generate/main.go`
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-bindata/go-bindata/v3"
)

const (
	output = "./bindata.go"
)

var (
	config = bindata.Config{
		Package: "main",
		Output:  output,

		Prefix: "resources/",
		Input: []bindata.InputConfig{
			bindata.InputConfig{
				Path:      filepath.Clean("resources/"),
				Recursive: false,
			},
		},

		HttpFileSystem: true,
	}
)

func main() {
	generate()
}

func generate() {
	err := bindata.Translate(&config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bindata: %v\n", err)
		os.Exit(1)
	}
}
