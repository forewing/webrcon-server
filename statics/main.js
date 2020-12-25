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
        shortcutGroups: {},
    },
    mounted() {
        this.loadPreset(this.setPreset);
    },
    methods: {
        sendExec: function (succ, fail, cmd) {
            fetch(
                `./api/exec`, {
                body: JSON.stringify({
                    'cmd': cmd,
                }),
                method: 'POST',
                headers: {
                    'content-type': 'application/json',
                },
            }).then(function (res) {
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
        },
        loadPreset: function (succ) {
            fetch(`./preset.json`)
                .then(function (res) {
                    if (res.status !== 200) {
                        console.log('Fail with Status Code: ' + res.status);
                    } else {
                        res.json().then(succ);
                    }
                })
                .catch(err => console.log('Fetch preset error :-S', err));
        },
        setPreset: function (preset) {
            this.shortcutGroups = preset;
        }
    }
})
