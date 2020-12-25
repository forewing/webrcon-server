package main

import (
	"encoding/json"
	"log"
	"time"
	"webrcon-server/statics"
	"webrcon-server/templates"

	rcon "github.com/forewing/csgo-rcon"
	"github.com/gin-gonic/gin"
)

var (
	client *rcon.Client
)

//go:generate go run generate/main.go
func main() {
	conf, err := json.MarshalIndent(flags, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Println(string(conf))
	client = rcon.New(*flags.Address, *flags.Password, time.Duration(*flags.Timeout*float64(time.Second)))

	loadPreset()

	if !*flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.StaticFS("/statics", statics.AssetFile())

	router.GET("/", func(c *gin.Context) {
		c.FileFromFS("./main.html", templates.AssetFile())
	})

	router.GET("/preset.json", getPreset)

	router.GET("/api/connect", getConnect)

	routerExec := router.Group("/api/exec")
	if len(*flags.BasicAuthUsername) > 0 {
		routerExec = router.Group("/api/exec", gin.BasicAuth(gin.Accounts{
			*flags.BasicAuthUsername: *flags.BasicAuthPassword,
		}))
	}
	routerExec.GET("/", getExec)
	routerExec.POST("/", postExec)

	log.Println("Listening on", "http://"+*flags.Bind)
	router.Run(*flags.Bind)
}
