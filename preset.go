package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	defaultPresetPath = "csgo-default.json"
)

var (
	// path of the preset file to be served
	usePresetPath = defaultPresetPath
	// load from fs.FS if true, from disk otherwise
	usePresetEmbedFS = true
)

func checkPreset() {
	if len(*flags.Preset) > 0 {
		if _, err := presets.Open(*flags.Preset); err == nil {
			usePresetPath = *flags.Preset
			usePresetEmbedFS = true
			log.Println("Use built-in preset", usePresetPath)
			return
		}
		if _, err := os.Open(*flags.Preset); err == nil {
			usePresetPath = *flags.Preset
			usePresetEmbedFS = false
			log.Println("Use custom preset", *flags.Preset)
			return
		}
		log.Println("Error: load config preset", *flags.Preset, "failed, try default preset")
	}

	if _, err := presets.Open(defaultPresetPath); err == nil {
		usePresetPath = defaultPresetPath
		usePresetEmbedFS = true
		log.Println("Use default built-in preset", defaultPresetPath)
	} else {
		panic("Load default built-in preset failed: " + err.Error())
	}
}

func getPreset(c *gin.Context) {
	if usePresetEmbedFS {
		c.FileFromFS(usePresetPath, http.FS(presets))
		return
	}
	c.File(usePresetPath)
}
