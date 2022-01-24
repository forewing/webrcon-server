package main

import (
	"flag"
	"fmt"

	"github.com/forewing/gobuild"
)

const (
	name   = "webrcon-server"
	module = "github.com/forewing/webrcon-server"
)

var (
	flagAll = flag.Bool("all", false, "build for all platforms")

	target = gobuild.Target{
		Source:      ".",
		OutputName:  name,
		OutputPath:  "./output",
		CleanOutput: true,

		ExtraFlags:   []string{"-trimpath"},
		ExtraLdFlags: "-s -w",

		VersionPath: module + "/version.Version",
		HashPath:    module + "/version.Hash",

		Compress:  gobuild.CompressRaw,
		Platforms: []gobuild.Platform{{}},
	}
)

func main() {
	flag.Parse()
	if *flagAll {
		target.OutputName = fmt.Sprintf("%s-%s-%s-%s",
			name,
			gobuild.PlaceholderVersion,
			gobuild.PlaceholderOS,
			gobuild.PlaceholderArch)
		target.Compress = gobuild.CompressZip
		target.Platforms = gobuild.PlatformCommon
	}
	err := target.Build()
	if err != nil {
		panic(err)
	}
}
