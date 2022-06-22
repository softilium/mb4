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
    }
};

let app = Vue.createApp(App)
applyVueFuncs(app);
app.mount('#vueapp')
