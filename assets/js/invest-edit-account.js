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
        printCurrency(val) { return new Intl.NumberFormat('ru-RU', { style: 'decimal', minimumFractionDigits: 0, maximumFractionDigits: 0 }).format(val); },
        printPercent1(val) { return new Intl.NumberFormat('ru-RU', { style: 'decimal', minimumFractionDigits: 1, maximumFractionDigits: 1 }).format(val); },
        printDate(val) { return new Intl.DateTimeFormat('ru-RU').format(Date.parse(val)); },
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
app.mount('#vueapp')
