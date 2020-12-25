package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"webrcon-server/statics"

	"github.com/gin-gonic/gin"
)

const (
	defaultPresetPath = "preset-csgo-default.json"
)

var (
	savedPreset = []byte("{}")
)

func loadPreset() {
	if len(*flags.Preset) > 0 {
		if preset, err := ioutil.ReadFile(*flags.Preset); err == nil {
			savedPreset = preset
			log.Println("Use preset", *flags.Preset, ":", string(preset))
			return
		}
		log.Println("Error: load preset", *flags.Preset, "failed, try default preset")
	}

	if preset, err := statics.Asset(defaultPresetPath); err == nil {
		savedPreset = preset
		log.Println("Use default preset:", string(preset))
	} else {
		panic("Load default preset failed: " + err.Error())
	}
}

func getPreset(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, savedPreset)
}
