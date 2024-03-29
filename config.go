package main

import (
	"encoding/json"
	"flag"
	"os"

	rcon "github.com/forewing/csgo-rcon"
	"github.com/forewing/webrcon-server/version"
)

// Flags holds configs
type Flags struct {
	Address  *string  `json:",omitempty"`
	Password *string  `json:",omitempty"`
	Timeout  *float64 `json:",omitempty"`

	Preset *string `json:",omitempty"`
	Config *string `json:",omitempty"`

	PublicAddress *string `json:",omitempty"`

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

		Preset: flag.String("preset", "", "use `preset`(path), empty for default csgo config"),
		Config: flag.String("conf", "", "load configs from `file` instead of flags"),

		PublicAddress: flag.String("public-addr", "", "redirect target(public `address`) for /api/connect, empty for disabled"),

		Bind:              flag.String("bind", "0.0.0.0:8080", "webrcon-server bind `address`"),
		BasicAuthUsername: flag.String("admin-name", "", "basicauth `username` for path /api/exec"),
		BasicAuthPassword: flag.String("admin-pass", "", "basicauth `password` for path /api/exec"),

		Debug: flag.Bool("debug", false, "turn on debug mode"),
	}

	flagVersion = flag.Bool("version", false, "display versions")
)

func init() {
	flag.Parse()

	if *flagVersion {
		version.Display()
		os.Exit(0)
	}

	if len(*flags.Config) == 0 {
		return
	}
	data, err := os.ReadFile(*flags.Config)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &flags)
	if err != nil {
		panic(err)
	}
}
