const App = {
    data() {
        return {
            obj: {}
        }
    },
    async mounted() {
        await this.getList();
    },
    methods: {
        async getList() {
            let response = await fetch(`/api/invest-accounts?id=${window.accid}`, { method: 'GET' });
            if (response.ok) {
                this.obj = await response.json();
            }
            else alert("Проблема с получением брокерского счета");
        },
        async deleteValuation(id) {
            if (confirm("Удалить указанную строку оценки?")) {

                let response = await fetch(`/api/invest-account-valuations?id=${id}`, { method: 'DELETE' });
                if (response.ok) {
                    await this.getList();
                } else {
                    alert("Ошибка удаления строки оценки")
                }
            }
        },
        async setDescr() {
            let response = await fetch(
                `/api/invest-accounts?id=${window.accid}&newdescr=${this.obj.Descr}`,
                { method: 'PUT' }
            );
            if (!response.ok) alert("Проблема с установкой наименования брокерского счета");
        }

    }
};

let app = Vue.createApp(App)
applyVueFuncs(app);
app.mount('#vueapp')
