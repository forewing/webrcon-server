package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"webrcon-server/presets"

	"github.com/gin-gonic/gin"
)

const (
	defaultPresetPath = "csgo-default.json"
)

var (
	savedPreset = []byte("{}")
)

func loadPreset() {
	if len(*flags.Preset) > 0 {
		if preset, err := presets.Asset(*flags.Preset); err == nil {
			savedPreset = preset
			log.Println("Use preset", *flags.Preset, ":", string(preset))
			return
		}
		if preset, err := ioutil.ReadFile(*flags.Preset); err == nil {
			savedPreset = preset
			log.Println("Use custom preset", *flags.Preset, ":", string(preset))
			return
		}
		log.Println("Error: load config preset", *flags.Preset, "failed, try default preset")
	}

	if preset, err := presets.Asset(defaultPresetPath); err == nil {
		savedPreset = preset
		log.Println("Use default preset:", string(preset))
	} else {
		panic("Load default preset failed: " + err.Error())
	}
}

func getPreset(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, savedPreset)
}
