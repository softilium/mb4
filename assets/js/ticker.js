const App = {
    data() {
        return {
            tri: {},
            pnl: {},
            cf: {}
        }
    },
    async mounted() {
        await this.getTickerData();
    },
    methods: {

        async createCandlesChart(element, tickerName, dates, candles, volumes, deals) {

            var upColor = '#00da3c';
            var downColor = '#ec0000';

            var myChart = echarts.init(element);

            option = {
                backgroundColor: '#fff',
                animation: false,
                legend: { left: 10 },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'cross'
                    },
                    backgroundColor: 'rgba(245, 245, 245, 0.8)',
                    borderWidth: 1,
                    borderColor: '#ccc',
                    padding: 10,
                    textStyle: {
                        color: '#000'
                    },
                    position: function (pos, params, el, elRect, size) {
                        var obj = { top: 10 };
                        obj[['left', 'right'][+(pos[0] < size.viewSize[0] / 2)]] = 30;
                        return obj;
                    }
                },
                axisPointer: {
                    link: { xAxisIndex: 'all' },
                    label: {
                        backgroundColor: '#777'
                    }
                },
                toolbox: { show: false },
                brush: {
                    xAxisIndex: 'all',
                    brushLink: 'all',
                    outOfBrush: {
                        colorAlpha: 0.1
                    }
                },
                visualMap: {
                    show: false,
                    seriesIndex: 5,
                    dimension: 2,
                    pieces: [{
                        value: 1,
                        color: downColor
                    }, {
                        value: -1,
                        color: upColor
                    }]
                },
                grid: [
                    {
                        left: '5%',
                        right: '5%',
                        height: '60%'
                    },
                    {
                        left: '5%',
                        right: '5%',
                        top: '61%',
                        height: '14%'
                    },
                    {
                        left: '5%',
                        right: '5%',
                        top: '75%',
                        height: '14%'
                    }
                ],
                xAxis: [
                    {
                        type: 'category',
                        data: dates,
                        scale: true,
                        boundaryGap: false,
                        axisLine: { onZero: false },
                        splitLine: { show: false },
                        splitNumber: 20,
                        min: 'dataMin',
                        max: 'dataMax',
                        axisPointer: {
                            z: 100
                        }
                    },
                    {
                        type: 'category',
                        gridIndex: 1,
                        data: dates,
                        scale: true,
                        boundaryGap: false,
                        axisLine: { onZero: false },
                        axisTick: { show: false },
                        splitLine: { show: false },
                        axisLabel: { show: false },
                        splitNumber: 20,
                        min: 'dataMin',
                        max: 'dataMax'
                    },
                    {
                        type: 'category',
                        gridIndex: 2,
                        data: dates,
                        scale: true,
                        boundaryGap: false,
                        axisLine: { onZero: false },
                        axisTick: { show: false },
                        splitLine: { show: false },
                        axisLabel: { show: false },
                        splitNumber: 20,
                        min: 'dataMin',
                        max: 'dataMax'
                    }
                ],
                yAxis: [
                    {
                        scale: true,
                        splitArea: {
                            show: true
                        }
                    },
                    {
                        scale: true,
                        gridIndex: 1,
                        splitNumber: 2,
                        axisLabel: { show: false },
                        axisLine: { show: false },
                        axisTick: { show: false },
                        splitLine: { show: false }
                    },
                    {
                        scale: true,
                        gridIndex: 2,
                        axisLabel: { show: false },
                        axisLine: { show: false },
                        axisTick: { show: false },
                        splitLine: { show: false },
                        min: '0',
                        max: '100'
                    }
                ],
                dataZoom: [
                    {
                        type: 'inside',
                        xAxisIndex: [0, 1, 2],
                        start: 70,
                        end: 100
                    },
                    {
                        show: true,
                        xAxisIndex: [0, 1, 2],
                        type: 'slider',
                        top: '90%',
                        start: 70,
                        end: 100
                    }
                ],
                series: [
                    {
                        name: tickerName,
                        type: 'candlestick',
                        data: candles,
                        itemStyle: {
                            color: upColor,
                            color0: downColor,
                            borderColor: null,
                            borderColor0: null
                        },
                        markPoint: { data: deals }

                    },
                    {
                        name: 'Объем',
                        type: 'bar',
                        xAxisIndex: 1,
                        yAxisIndex: 1,
                        data: volumes
                    }

                ]
            };

            myChart.setOption(option);

        },
        async PnlView(element, dates, revenues, interestIncomes, ebitdas, ammortizations, interestExpenses, taxes, incomes) {

            var myChart = echarts.init(element);

            option = {
                animation: false,
                legend: { left: 10 },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: { type: 'cross' },
                },
                xAxis: [
                    {
                        type: 'category',
                        data: dates,
                        axisLabel: {
                            interval: 0,
                            rotate: 90
                        },
                    }
                ],
                yAxis: [
                    { type: 'value' }
                ],
                series: [
                    {
                        name: 'Выручка',
                        type: 'line',
                        stack: 'Доходы',
                        smooth: true,
                        emphasis: { focus: 'series' }, areaStyle: {}, symbol: 'none',
                        data: revenues
                    },
                    {
                        name: 'Финансовые доходы',
                        type: 'line',
                        stack: 'Доходы',
                        smooth: true,
                        emphasis: { focus: 'series' }, areaStyle: {}, symbol: 'none',
                        data: interestIncomes
                    },
                    {
                        name: 'Амортизация',
                        type: 'line',
                        stack: 'Расходы',
                        smooth: true,
                        emphasis: { focus: 'series' }, areaStyle: {}, symbol: 'none',
                        data: ammortizations
                    },
                    {
                        name: 'Финансовые расходы',
                        type: 'line',
                        stack: 'Расходы',
                        smooth: true,
                        emphasis: { focus: 'series' }, areaStyle: {}, symbol: 'none',
                        data: interestExpenses
                    },
                    {
                        name: 'Налог с прибыли',
                        type: 'line',
                        stack: 'Расходы',
                        smooth: true,
                        emphasis: { focus: 'series' }, areaStyle: {}, symbol: 'none',
                        data: taxes
                    },
                    {
                        name: 'Прибыль',
                        type: 'bar',
                        stack: 'Прибыль',
                        smooth: true,
                        emphasis: { focus: 'series' }, symbol: 'none',
                        data: incomes
                    },
                    {
                        name: 'EBITDA',
                        type: 'line',
                        smooth: true,
                        emphasis: { focus: 'series' }, symbol: 'none',
                        data: ebitdas
                    },

                ]
            };

            myChart.setOption(option);

        },
        async CfView(element, dates, cash, debt, equity, mcap, bookValue) {

            var myChart = echarts.init(element);

            option = {
                animation: false,
                legend: { left: 10 },
                tooltip: {
                    trigger: 'axis',
                    axisPointer: { type: 'cross' },
                },
                xAxis: [
                    {
                        type: 'category',
                        data: dates,
                        axisLabel: {
                            interval: 0,
                            rotate: 90
                        },
                    }
                ],
                yAxis: [
                    { type: 'value' }
                ],
                series: [
                    {
                        name: 'Денежные средства',
                        type: 'line',
                        emphasis: { focus: 'series' },
                        symbol: 'none',
                        smooth: true,
                        data: cash
                    },
                    {
                        name: 'Акционерный капитал',
                        type: 'line',
                        stack: 'Total',
                        emphasis: { focus: 'series' },
                        areaStyle: {},
                        symbol: 'none',
                        smooth: true,
                        data: equity
                    },
                    {
                        name: 'Обязательства',
                        type: 'line',
                        stack: 'Total',
                        emphasis: { focus: 'series' },
                        areaStyle: {},
                        symbol: 'none',
                        smooth: true,
                        data: debt
                    },
                    {
                        name: 'Балансовая стоимость',
                        type: 'line',
                        //stack: 'Total',
                        emphasis: { focus: 'series' },
                        //areaStyle: {},
                        symbol: 'none',
                        smooth: true,
                        data: bookValue
                    },
                    {
                        name: 'Капитализация',
                        type: 'line',
                        emphasis: { focus: 'series' },
                        symbol: 'none',
                        smooth: true,
                        data: mcap
                    },

                ]
            };

            myChart.setOption(option);

        },

        async getTickerData() {

            let response = await fetch(`/ticker?id=${window.tickerid}&mode=candles`, { method: "GET" });
            if (response.ok) {
                this.tri = await response.json();
                this.createCandlesChart(
                    this.$refs.candlesref,
                    window.tickerDescr,
                    this.tri.CandleDates,
                    this.tri.CandleOCLH,
                    this.tri.CandleVolumes,
                    []
                );
            } else {
                alert("Проблема при получении данных Candles");
                return;
            }

            response = await fetch(`/ticker?id=${window.tickerid}&mode=pnl`, { method: "GET" });
            if (response.ok) {
                this.pnl = await response.json();
                this.PnlView(
                    this.$refs.pnlView,
                    this.pnl.Dates,
                    this.pnl.Revenues,
                    this.pnl.InterestIncomes,
                    this.pnl.Ebitdas,
                    this.pnl.Ammortizations,
                    this.pnl.InterestExpenses,
                    this.pnl.Taxes,
                    this.pnl.Incomes
                );
            } else {
                alert("Проблема при получении данных PNL");
            }

            response = await fetch(`/ticker?id=${window.tickerid}&mode=cf`, { method: "GET" });
            if (response.ok) {
                this.cf = await response.json();
                this.CfView(
                    this.$refs.cfView,
                    this.cf.Dates,
                    this.cf.Cash,
                    this.cf.Debt,
                    this.cf.Equity,
                    this.cf.MCap,
                    this.cf.BookValue
                );
            } else {
                alert("Проблема при получении данных CF");
            }
        }
    },
};

let app = Vue.createApp(App)
applyVueFuncs(app);
app.mount('#vueapp')
