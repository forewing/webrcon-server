# WEBRCON-Server

A web based control panel for srcds' RCON protocol (CS:GO).

![preview](preview.png)

## Deployment

1. `cp server/config.py.example server/config.py`

    Then edit ip, port, password to your own server's.

1. `docker-compose up -d`

Now it will listen on `localhost:27022` .

## Usage

### API

1. exec

    ```
    curl -X POST 127.0.0.1:27022/api/exec -H "Content-Type: application/json" -d '{"cmd":"YOUR_CMD_HERE"}'
    ```

    or

    ```
    curl -v -X GET '127.0.0.1:27022/api/exec?cmd=YOUR_CMD_HERE'
    ```

2. connect

    Visit `127.0.0.1:27022/api/connect`, you will be 301 redired to steam game launching shortcut.

### GUI

Browse to `http://127.0.0.1:27022/`

Please help us to add more command shortcut.

In client/public/static/main.js, edit the object

```
shortcutGroups: {
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