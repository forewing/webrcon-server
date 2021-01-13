package main

import (
	"embed"
	"io/fs"
	"os"
)

var (
	//go:embed statics/*
	staticsEmbed embed.FS
	statics      fs.FS

	//go:embed templates/*
	templatesEmbed embed.FS
	templates      fs.FS

	//go:embed presets/*
	presets embed.FS
)

func mustStripFSPrefix(sfs fs.FS, prefix string) fs.FS {
	dfs, err := fs.Sub(sfs, prefix)
	if err != nil {
		panic(err)
	}
	return dfs
}

func init() {
	statics = mustStripFSPrefix(staticsEmbed, "statics")
	templates = mustStripFSPrefix(templatesEmbed, "templates")
}

func useLiveReload() {
	statics = os.DirFS("statics")
	templates = os.DirFS("templates")
}
