Vue.component('shortcut-group', {
    props: ['shortcuts', 'name'],
    template: `
    <div class="card bg-light mb-3">
        <div class="card-header bg-middleblue">
            <h4 class="card-title my-auto">
                <span class="oi oi-puzzle-piece" title="puzzle-piece" aria-hidden="true"></span> {{ name }}
            </h4>
        </div>
        <ul class="list-group list-group-flush">
            <shortcut-panel
            v-for="shortcut in shortcuts"
            v-bind:key="shortcut.id"
            v-bind:shortcut="shortcut"
            v-on:send-cmd="emitExec">
            </shortcut-panel>
        </ul>
    </div>
    `,
    methods: {
        emitExec(cmd) {
            console.log(cmd);
            this.$emit('send-cmd', cmd)
        }
    }
})

Vue.component('shortcut-panel', {
    props: ['shortcut'],
    template: `
        <li class="list-group-item">
            <button type="button" class="btn btn-outline-info btn-block"
                v-if="!shortcut.args" v-on:click="emitExec"
            >{{ shortcut.name }}</button>
            <div class="input-group" v-if="shortcut.args">
                <div class="input-group-prepend">
                    <button class="btn btn-outline-info" type="button" id="button-submit"
                        v-on:click="emitExec">{{ shortcut.name }}</button>
                </div>
                <input type="text" class="form-control" placeholder="Command Args Here."
                    aria-label="Command Args to be passed." aria-describedby="button-submit"
                    v-model="args">
            </div>
        </li>
    `,
    data: function () {
        return {
            args: this.shortcut.default
        }
    },
    methods: {
        emitExec() {
            console.log(this.shortcut.cmd);
            var cmd = this.shortcut.cmd;
            if (this.shortcut.args) {
                cmd += ` ${this.args}`;
            }
            this.$emit('send-cmd', cmd)
        }
    }
})

var app = new Vue({
    el: '#app',
    data: {
        command: 'maps *',
        reply: '',
        shortcutGroups: {
            "Rounds": [
                { id: 1, args: true, default: "5", name: "restart", cmd: "mp_restartgame" },
                { id: 2, args: true, default: "30", name: "maxrounds", cmd: "mp_maxrounds" },
            ],
            "maps": [
                { id: 1, args: true, default: "*", name: "list maps", cmd: "maps" },
                { id: 2, args: true, default: "de_dust2", name: "change map", cmd: "map" },
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
        },
    },
    methods: {
        sendExec: function (succ, fail, cmd) {
            fetch(`/api/exec?cmd=${cmd}`)
                .then(function (res) {
                    if (res.status !== 200) {
                        fail('Fail with Status Code: ' + res.status);
                    } else {
                        res.text().then(succ);
                    }
                })
                .catch(err => console.log('Fetch Error :-S', err));
        },
        sendExecPanel: function (cmd) {
            console.log('emitted')
            this.command = cmd;
            this.sendExec(
                this.updateReply,
                this.updateReply,
                cmd);
        },
        updateReply: function (msg) {
            this.reply = msg;
        }
    }
})
