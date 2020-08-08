Vue.component('shortcut-group', {
    props: ['shortcuts', 'name'],
    template: "#shortcut-group",
    methods: {
        emitExec(cmd) {
            this.$emit('send-cmd', cmd)
        }
    }
})

Vue.component('shortcut-panel', {
    props: ['shortcut'],
    template: "#shortcut",
    data: function () {
        return {
            args: this.shortcut.default
        }
    },
    methods: {
        emitExec() {
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
                { id: 1, args: true, default: "5", name: "Restart Game", cmd: "mp_restartgame" },
                { id: 2, args: true, default: "30", name: "Max Rounds", cmd: "mp_maxrounds" },
                { id: 3, args: true, default: "999", name: "Warmup Time", cmd: "mp_warmuptime" },
                { id: 4, args: false, default: "", name: "Infinite Time", cmd: "mp_roundtime_defuse 60; mp_roundtime_hostage 60; mp_roundtime 60;mp_restartgame 1" },
            ],
            "Maps": [
                { id: 1, args: true, default: "*", name: "List Maps", cmd: "maps" },
                { id: 2, args: true, default: "de_dust2", name: "Change Map", cmd: "map" },
            ],
            "Users": [
                { id: 1, args: false, default: "", name: "List Users", cmd: "users" },
                { id: 2, args: true, default: "", name: "User Info", cmd: "user" },
                { id: 3, args: true, default: "", name: "Kick User", cmd: "kick" },
                { id: 4, args: true, default: "", name: "Kill User", cmd: "kill" },
            ],
            "Bots": [
                { id: 1, args: false, default: "", name: "Kick All Bot", cmd: "bot_kick" },
                { id: 2, args: false, default: "", name: "Kick CT", cmd: "bot_kick ct" },
                { id: 3, args: false, default: "", name: "Kick T", cmd: "bot_kick t" },
                { id: 4, args: false, default: "", name: "Add CT", cmd: "bot_add_ct" },
                { id: 5, args: false, default: "", name: "Add T", cmd: "bot_add_t" },
            ],
            "Modes": [
                { id: 1, args: true, default: "de_dust2", name: "Classic Casual", cmd: "game_mode 0; game_type 0; changelevel" },
                { id: 2, args: true, default: "de_dust2", name: "Classic Competitive", cmd: "game_mode 1; game_type 0; changelevel" },
                { id: 3, args: true, default: "ar_baggage", name: "Arms Race", cmd: "game_mode 0; game_type 1; changelevel" },
                { id: 4, args: true, default: "de_dust2", name: "Demolition", cmd: "game_mode 1; game_type 1; changelevel" },
                { id: 5, args: true, default: "de_dust2", name: "Deathmatch", cmd: "game_mode 2; game_type 1; changelevel" },
                { id: 6, args: true, default: "dz_sirocco", name: "Dangerzone", cmd: "game_type 6; game_mode 0; changelevel" },
            ],
            "Cheats": [
                { id: 1, args: false, default: "", name: "Cheat On", cmd: "sv_cheats 1" },
                { id: 2, args: false, default: "", name: "Cheat Off", cmd: "sv_cheats 0" },
                { id: 3, args: true, default: "1.0", name: "Time Scale", cmd: "host_timescale" },
            ]
        },
    },
    methods: {
        sendExec: function (succ, fail, cmd) {
            fetch(`./api/exec?cmd=${encodeURIComponent(cmd)}`)
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