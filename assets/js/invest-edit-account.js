const App = {
    data() {
        return {
            // stubs for initial rendering
            obj: {
                edges: {
                    Valuations: []
                }
            },
            selectedVal: {
                id: null,
                RecDate: null,
                Value: null
            }
        }
    },
    async mounted() {
        await this.getList();
    },
    methods: {
        async setNewSelectedVal() {
            this.selectedVal.id = null;
            this.selectedVal.Value = 0;
            this.selectedVal.RecDate = '2022-01-01';
        },
        async setSelectedVal(item) {
            this.selectedVal.id = item.id;
            this.selectedVal.Value = item.Value;
            this.selectedVal.RecDate = item.RecDate.split('T')[0]; // strip time from date
        },
        async saveSelectedVal() {
            let res = null;
            this.selectedVal.RecDate += "T03:00:00+03:00"; // restore full time from date
            if (this.selectedVal.id == null) {
                res = await fetch(`/api/invest-account-valuations?id=${this.selectedVal.id}&owner=${window.accid}`,
                    { method: 'POST', body: JSON.stringify(this.selectedVal) }
                );
            } else {
                res = await fetch(`/api/invest-account-valuations?id=${this.selectedVal.id}`,
                    { method: 'PUT', body: JSON.stringify(this.selectedVal) }
                );
            }
            if (res.ok) this.getList(); else alert("Ошибка при сохранении");
        },
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
