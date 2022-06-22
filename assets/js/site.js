function applyVueFuncs(app) {

    app.config.globalProperties.$prnCurrency = function (val) {
        return new Intl.NumberFormat('ru-RU', { style: 'decimal', minimumFractionDigits: 0, maximumFractionDigits: 0 }).format(val);
    }
    ;

    app.config.globalProperties.$prnPercent1 = function (val) {
        return new Intl.NumberFormat('ru-RU', { style: 'decimal', minimumFractionDigits: 1, maximumFractionDigits: 1 }).format(val);
    };

    app.config.globalProperties.$prnDate = function (val) {
        let tokens = val.split('T');
        if (tokens.length == 2) {
            return new Intl.DateTimeFormat('ru-RU').format(Date.parse(tokens[0]))
        }
        return val;
    };
}
