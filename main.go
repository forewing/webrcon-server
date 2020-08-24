package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	rcon "github.com/forewing/csgo-rcon"
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

	http.HandleFunc("/api/exec", exec)
	http.HandleFunc("/api/connect", connect)

	http.Handle("/", http.FileServer(AssetFile()))

	http.ListenAndServe(*flags.Bind, logger(http.DefaultServeMux))
}

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-Ip")
		if ip == "" {
			ip = r.Header.Get("X-Forwarded-For")
		}
		if ip == "" {
			ip = r.RemoteAddr
		}
		log.Printf("%s\t%s\t%s\n", ip, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func exec(w http.ResponseWriter, req *http.Request) {
	if ok := auth(w, req); !ok {
		return
	}

	var cmd string

	if req.Method == http.MethodGet {
		data, ok := req.URL.Query()["cmd"]
		if !ok || len(data) < 1 {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		cmd = data[0]
	} else if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		data := make(map[string]string)
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		ok := false
		cmd, ok = data["cmd"]
		if !ok {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	cmd = strings.TrimSpace(cmd)
	if len(cmd) == 0 {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	message, err := client.Execute(cmd)
	message = strings.TrimSpace(message)

	fmt.Fprintln(w, message)
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func connect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, fmt.Sprintf("steam://connect/%v", *flags.Address), http.StatusTemporaryRedirect)
}

func auth(w http.ResponseWriter, req *http.Request) (ok bool) {
	defer func() {
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "", http.StatusUnauthorized)
		}
	}()

	if flags.BasicAuthDisabled {
		return true
	}

	u, p, ok := req.BasicAuth()
	if !ok {
		return false
	}
	if u == *flags.BasicAuthUsername && p == *flags.BasicAuthPassword {
		return true
	}

	return false
}
