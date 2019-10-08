# WEBRCON-Server

A web based client's server side for srcds' RCON protocol.

## Deployment

1. `cp server/config.py.example server/config.py`

    Then edit ip, port, password to your own server's.

1. `docker-compose up -d`

Now it will listen on `localhost:27020` .

## Usage

### API

1. exec

    ```
    curl -X POST 127.0.0.1:27020/api/exec -H "Content-Type: application/json" -d '{"cmd":"YOUR_CMD_HERE"}'
    ```

    or

    ```
    curl -v -X GET '127.0.0.1:27020/api/exec?cmd=YOUR_CMD_HERE'
    ```

2. connect

    Visit `127.0.0.1:27020/api/connect`, you will be 301 redired to steam game launching shortcut.

### GUI

Browse to `http://127.0.0.1:27020/`

> You may put your own frontend files to /client/public, THIS IS A TODO FEATURE that will be added to the repo in future.