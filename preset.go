package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"sigs.k8s.io/yaml"
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

func convertYAMLPreset(f *os.File) (string, error) {
	tmpFile, err := os.Create(f.Name() + ".json")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	yamlContent, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	jsonContent, err := yaml.YAMLToJSON(yamlContent)
	if err != nil {
		return "", err
	}

	_, err = tmpFile.Write(jsonContent)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

func checkPreset() {
	if len(*flags.Preset) > 0 {
		if _, err := presets.Open(*flags.Preset); err == nil {
			usePresetPath = *flags.Preset
			usePresetEmbedFS = true
			log.Println("Use built-in preset", usePresetPath)
			return
		}
		if f, err := os.Open(*flags.Preset); err == nil {
			defer f.Close()
			usePresetPath = *flags.Preset
			usePresetEmbedFS = false
			log.Println("Use custom preset", *flags.Preset)
			if strings.HasSuffix(*flags.Preset, ".yml") ||
				strings.HasSuffix(*flags.Preset, ".yaml") {
				newPath, err := convertYAMLPreset(f)
				if err == nil {
					usePresetPath = newPath
					return
				}
				log.Printf("Error: failed to convert %s to json: %v", *flags.Preset, err)
			} else {
				return
			}
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
