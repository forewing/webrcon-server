package version

import (
	"fmt"
	"runtime"
)

const (
	VersionDefault = "dev"
	HashDefault    = "unknown"
)

var (
	Version = VersionDefault
	Hash    = HashDefault
)

// Display version
func Display() {
	fmt.Println("github.com/forewing/webrcon-server")
	fmt.Printf("version\t%v\n", Version)
	fmt.Printf("commit\t%v\n", Hash)
	fmt.Printf("runtime\t%v\n", runtime.Version())
}
