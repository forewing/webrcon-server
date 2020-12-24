package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	rcon "github.com/forewing/csgo-rcon"
)

// Flags holds configs
type Flags struct {
	Address  *string  `json:",omitempty"`
	Password *string  `json:",omitempty"`
	Timeout  *float64 `json:",omitempty"`

	Config *string `json:",omitempty"`

	Bind              *string `json:",omitempty"`
	BasicAuthUsername *string `json:",omitempty"`
	BasicAuthPassword *string `json:",omitempty"`

	Debug *bool `json:",omitempty"`
}

var (
	flags Flags = Flags{
		Address:  flag.String("addr", rcon.DefaultAddress, "`address` of the server RCON, in the format of HOST:PORT"),
		Password: flag.String("pass", rcon.DefaultPassword, "`password` of the RCON"),
		Timeout:  flag.Float64("timeout", rcon.DefaultTimeout.Seconds(), "`timeout`(seconds) of the connection"),

		Config: flag.String("conf", "", "load configs from `file` instead of flags"),

		Bind:              flag.String("bind", "0.0.0.0:8080", "webrcon-server bind `address`"),
		BasicAuthUsername: flag.String("admin-name", "", "basicauth `username` for path /api/exec"),
		BasicAuthPassword: flag.String("admin-pass", "", "basicauth `password` for path /api/exec"),

		Debug: flag.Bool("debug", false, "turn on debug mode"),
	}
)

func init() {
	flag.Parse()
	if len(*flags.Config) == 0 {
		return
	}
	data, err := ioutil.ReadFile(*flags.Config)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &flags)
	if err != nil {
		panic(err)
	}
}
