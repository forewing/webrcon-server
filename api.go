package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ipFilter = regexp.MustCompile(
		`rcon from \"[0-9.:]+\":`,
	)
	ipFilterReplace = "rcon from \"***\":"
)

func getConnect(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("steam://connect/%v", *flags.Address))
}

func getExec(c *gin.Context) {
	cmd := c.DefaultQuery("cmd", "")
	msg, err := exec(cmd)
	if err != nil {
		c.String(http.StatusBadRequest, "%v\n%v", msg, err)
		return
	}

	c.String(http.StatusOK, msg)
}

type execModel struct {
	Cmd string `json:"cmd" binding:"required"`
}

func postExec(c *gin.Context) {
	var postData execModel
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	cmd := postData.Cmd

	msg, err := exec(cmd)
	if err != nil {
		c.String(http.StatusBadRequest, "%v\n%v", msg, err)
		return
	}
	c.String(http.StatusOK, msg)
}

func exec(cmd string) (string, error) {
	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		return "Error, empty command", nil
	}

	log.Println("cmd: ", cmd)
	msg, err := client.Execute(cmd)
	msg = strings.TrimSpace(msg)
	log.Println("msg: ", msg)

	if err != nil {
		log.Println("err: ", err)
	}

	msg = ipFilter.ReplaceAllString(msg, ipFilterReplace)

	return msg, err
}
