from flask import Flask, request, Response, abort, redirect

import srcds
import config

app = Flask(__name__)


@app.route('/api/exec', methods=['GET', 'POST'])
def api_exec():
    if request.method == 'GET':
        cmd = request.args.get("cmd")
    elif request.method == 'POST':
        cmd = request.json.get("cmd")

    if cmd:
        server = srcds.SourceRcon(config.ip, config.port, config.password)
        res = server.rcon(cmd)
        return Response(res)

    return abort(400)


@app.route('/api/connect', methods=['GET'])
def api_connect():
    return redirect(f"steam://connect/{config.ip}:{config.port}", code=301)


if __name__ == '__main__':
    app.run(debug=True, port=27020)
