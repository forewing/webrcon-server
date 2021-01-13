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

func init() {
	statics = mustStripFSPrefix(staticsEmbed, staticsPath)
	templates = mustStripFSPrefix(templatesEmbed, templatesPath)
	presets = mustStripFSPrefix(presetsEmbed, presetsPath)
}

func dirExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func useLiveReload() {
	if dirExist(staticsPath) {
		log.Println("live reload ./statics/*")
		statics = os.DirFS(staticsPath)
	}
	if dirExist(templatesPath) {
		log.Println("live reload ./templates/*")
		templates = os.DirFS(templatesPath)
	}
	if dirExist(presetsPath) {
		log.Println("live reload ./presets/*")
		presets = os.DirFS(presetsPath)
	}
}
