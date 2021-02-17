package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	rcon "github.com/forewing/csgo-rcon"
	"github.com/gin-gonic/gin"
)

var (
	client *rcon.Client
)

func main() {
	conf, err := json.MarshalIndent(flags, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Println(string(conf))
	client = rcon.New(*flags.Address, *flags.Password, time.Duration(*flags.Timeout*float64(time.Second)))

	prepareFS(*flags.Debug)
	if *flags.Debug {
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	checkPreset()

	router := gin.Default()

	// Set HTML handler
	router.SetHTMLTemplate(mustLoadTemplate())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.StaticFS("/statics", http.FS(statics))

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
	log.Fatalln(router.Run(*flags.Bind))
}

func mustLoadTemplate() *template.Template {
	t, err := template.New("").Delims("[[", "]]").ParseFS(templates, "*.html")
	if err != nil {
		log.Panicln(err)
	}
	return t
}
