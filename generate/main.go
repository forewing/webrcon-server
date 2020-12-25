// This program must be run from `../` as `go run generate/main.go`
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-bindata/go-bindata/v3"
)

const (
	output     = "./bindata.go"
	ignoreFile = `.*\.go`
)

func main() {
	generate("./statics")
	generate("./templates")
	generate("./presets")
}

func generate(path string) {
	cleanPath := filepath.Clean(path)
	config := bindata.Config{
		Package: cleanPath,
		Output:  filepath.Join(cleanPath, output),
		Prefix:  cleanPath + "/",
		Input: []bindata.InputConfig{
			bindata.InputConfig{
				Path:      cleanPath,
				Recursive: false,
			},
		},
		Ignore: []*regexp.Regexp{
			regexp.MustCompile(ignoreFile),
		},
		HttpFileSystem: true,
	}

	err := bindata.Translate(&config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bindata, %v: %v\n", path, err)
		os.Exit(1)
	}
}
