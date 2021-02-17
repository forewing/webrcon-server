package main

import (
	"embed"
	"io/fs"
	"log"
	"os"
)

const (
	staticsPath   = "statics"
	templatesPath = "templates"
	presetsPath   = "presets"
)

var (
	//go:embed statics/*
	staticsEmbed embed.FS
	statics      fs.FS

	//go:embed templates/*
	templatesEmbed embed.FS
	templates      fs.FS

	//go:embed presets/*
	presetsEmbed embed.FS
	presets      fs.FS
)

func mustStripFSPrefix(sfs fs.FS, prefix string) fs.FS {
	dfs, err := fs.Sub(sfs, prefix)
	if err != nil {
		panic(err)
	}
	return dfs
}

func prepareFS(debug bool) {
	statics = mustStripFSPrefix(staticsEmbed, staticsPath)
	templates = mustStripFSPrefix(templatesEmbed, templatesPath)
	presets = mustStripFSPrefix(presetsEmbed, presetsPath)

	if debug {
		useLiveReload(&statics, staticsPath)
		useLiveReload(&templates, templatesPath)
		useLiveReload(&presets, presetsPath)
	}
}

func useLiveReload(target *fs.FS, path string) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// Cannot live reload
			return
		}
	}
	log.Printf("live reload ./%v/*", path)
	*target = os.DirFS(path)
}
