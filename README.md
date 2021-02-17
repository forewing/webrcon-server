# WEBRCON-Server

A web based control panel for srcds' RCON protocol (CS:GO, Minecraft, etc.).

![preview](preview.png)

## Install

### Docker

You can get the docker image at [forewing/webrcon-server](https://hub.docker.com/r/forewing/webrcon-server)

Using docker-compose is recommended, it makes connecting webrcon to game server easier

An example of webrcon + csgo compose file is at [example/docker-compose-csgo.yml](example/docker-compose-csgo.yml)

Note that the server address should be configured to game server's service name with port, in the example, it should be `csgo:27015`

### Pre-Built Binary

Download it from [GitHub Action CI](https://github.com/forewing/webrcon-server/actions?query=workflow%3ACI+is%3Asuccess), latest successful build is recommended

### Build From Source

> Require go 1.16+

Using go install

```
go install github.com/forewing/webrcon-server
```

Or get the source code and build manually

```
git clone https://github.com/forewing/webrcon-server.git
cd webrcon-server
go build
```

## Config
### Flags

```
Usage of webrcon-server:
  -addr address
        address of the server RCON, in the format of HOST:PORT (default "127.0.0.1:27015")
  -admin-name username
        basicauth username for path /api/exec
  -admin-pass password
        basicauth password for path /api/exec
  -bind address
        webrcon-server bind address (default "0.0.0.0:8080")
  -conf file
        load configs from file instead of flags
  -debug
        turn on debug mode
  -pass password
        password of the RCON
  -preset preset
        use preset(path), empty for default csgo config
  -public-addr address
        redirect target(public address) for /api/connect, empty for disabled
  -timeout timeout
        timeout(seconds) of the connection (default 1)
```

### File Config

You should pass `-conf PATH/TO/config.json` flag in order to use file config

An example config file is at [example/config.example.json](example/config.example.json)
## API

1. exec

    ```
    curl -X POST 127.0.0.1:8080/api/exec -H "Content-Type: application/json" -d '{"cmd":"YOUR_CMD_HERE"}'
    ```

    or

    ```
    curl -v -X GET '127.0.0.1:8080/api/exec?cmd=YOUR_CMD_HERE'
    ```

2. connect

    Visit `127.0.0.1:8080/api/connect`, you will be 307 redired to steam game launching shortcut if `public-addr` set to `steam://connect/SERVER_ADDRESS:PORT`.

## GUI

Browse to `http://127.0.0.1:8080/`

## Shortcut group preset

Welcome to add presets for more games!

- By default (`preset` config is empty), the server will use `presets/csgo-default.json`.

- You may use other preset listed in `presets/` by setting `preset` to corresponding file name.

- If the file name set by `preset` cannot be found in `presets/`, the server will consider it as a path to your custom preset, and load it.

- The preset must be a valid json and follow the style:

```
{
    "Rounds": [
        { id: 1, args: true, default: "5", name: "restart", cmd: "mp_restartgame" },
        { id: 2, args: true, default: "30", name: "maxrounds", cmd: "mp_maxrounds" },
    ],
    "Bots": [
        { id: 1, args: false, default: "", name: "kick bot", cmd: "bot_kick" },
        { id: 2, args: false, default: "", name: "kick ct", cmd: "bot_kick ct" },
        { id: 3, args: false, default: "", name: "kick t", cmd: "bot_kick t" },
        { id: 4, args: false, default: "", name: "add ct", cmd: "bot_add_ct" },
        { id: 5, args: false, default: "", name: "add t", cmd: "bot_add_t" },
    ],
    "Cheats": [
        { id: 1, args: false, default: "", name: "cheat on", cmd: "sv_cheats 1" },
        { id: 2, args: false, default: "", name: "cheat off", cmd: "sv_cheats 0" },
    ]
}
```
