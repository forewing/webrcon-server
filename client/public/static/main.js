Vue.component('shortcut-group', {
    props: ['shortcuts', 'name'],
    template: `
    <div class="shortcut-group">
    <h1>{{ name }}</h1>
    <shortcut-panel
        v-for="shortcut in shortcuts"
        v-bind:key="shortcut.id"
        v-bind:shortcut="shortcut"
        v-on:send-cmd="emitExec">
    </shortcut-panel>
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
        <div class="shortcut-panel">
            <button v-on:click="emitExec">
                {{ shortcut.name }}
            </button>
            <input v-if="shortcut.args" v-model="args" placeholder="args here">
        </div>
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
            "rounds": [
                { id: 1, args: true, default: "5", name: "restart", cmd: "mp_restartgame" },
                { id: 2, args: true, default: "30", name: "maxrounds", cmd: "mp_maxrounds" },
            ],
            "bots": [
                { id: 1, args: false, default: "", name: "kick bot", cmd: "bot_kick" },
                { id: 2, args: false, default: "", name: "kick ct", cmd: "bot_kick ct" },
                { id: 3, args: false, default: "", name: "kick t", cmd: "bot_kick t" },
                { id: 4, args: false, default: "", name: "add ct", cmd: "bot_add_ct" },
                { id: 5, args: false, default: "", name: "add t", cmd: "bot_add_t" },
            ],
            "cheats": [
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
