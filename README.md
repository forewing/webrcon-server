# WEBRCON-Server

A web based client's server side for srcds' RCON protocol.

## Deployment

1. `cp config.py.example config.py`

    Then edit ip, port, password to your own server's.

1. `docker-compose up -d`

Now it will listen on `localhost:27020` .

## Usage

`curl -X POST 127.0.0.1:27020/api/exec -H "Content-Type: application/json" -d '{"cmd":"YOUR_CMD_HERE"}' `
