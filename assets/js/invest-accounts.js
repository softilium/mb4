const App = {
    data() {
        return {
            accList: {},
            weekflow: [],
            startInvestAccountsFlow: ""
        }
    },
    async mounted() {
        await this.getAccList();
    },
    methods: {
        async getAccList() {
            let response = await fetch("/api/users/start-invest-accounts-flow", { method: "GET" });
            if (response.ok) {
                const lDate = await response.text();
                if (lDate != "")
                    this.startInvestAccountsFlow = lDate;
                else
                    this.startInvestAccountsFlow = "2017-01-01";
                response = await fetch('/api/invest-accounts', { method: 'GET' });
                if (response.ok) this.accList = await response.json();
                else alert("Проблема с получением списка брокерских счетов");
            } else {
                alert("Проблема при получении стартовой даты расчета");
            }
        },
        async startInvestAccountsFlowChanged() {
            let response = await fetch(
                `/api/users/start-invest-accounts-flow?newdate=${this.startInvestAccountsFlow}`,
                { method: 'POST' }
            );
            if (response.ok) {
                this.showContent();
            }
            else alert("Проблема при установке стартовой даты расчета");
        },
        async applyTotalYieldChart(domref) {
            var myChart = echarts.init(domref);

            if (this.weekflow.length < 1) {
                myChart.clear();
                return;
            }
            dates = [];
            yields = [];
            this.weekflow.forEach(function (el) {
                dates.push(el.eow.split('T')[0]);
                yields.push(el.totalProfit);
            });
            option = {
                animation: false,
                xAxis: {
                    name: 'Недели',
                    type: 'category',
                    data: dates,
                },
                yAxis: {
                    type: 'value'
                },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: { type: 'cross' },
                },
                series: [
                    {
                        name: 'Результат',
                        type: 'line',
                        data: yields,
                        smooth: true,
                        showSymbol: false
                    }
                ]
            };
            myChart.setOption(option);

        },
        async showContent() {
            let idsArr = [];
            this.accList.forEach(function (el) { if (el.selected) idsArr.push(el.id); }
            );
            if (idsArr.length == 0) {
                this.weekflow = [];
                this.applyTotalYieldChart(this.$refs.totalYield);
                return;
            }
            let ids = idsArr.join(',');
            let response = await fetch(`/api/invest-accounts?mode=weekflow&ids=${ids}`, { method: 'GET' });
            if (response.ok) {
                this.weekflow = await response.json();
                this.applyTotalYieldChart(this.$refs.totalYield);
            }
            else alert("Проблема с получением списка брокерских счетов");
        }
    }
};

let app = Vue.createApp(App)
applyVueFuncs(app);
app.mount('#vueapp')
