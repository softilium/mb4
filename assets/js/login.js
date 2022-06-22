const App = {
    data() {
        return {

            userName: "",
            password: "",
            showAlert: false,
            alertText: "",

            userNameReg: "",
            passwordReg: "",
            passwordReg1: "",
            showAlertReg: false,
            alertTextReg: ""
        }
    },
    methods: {
        async login() {
            let res = await fetch(
                "/api/users/login?username=" + this.userName + "&password=" + this.password,
                { method: 'POST' }
            );
            if (res.ok)
                window.location.href = '/'
            else {
                this.showAlert = true;
                this.alertText = await res.text();
                //TODO Не выводится сообщение об ошибке
            }
        },
        async register() {
            if (this.passwordReg != this.passwordReg1) {
                this.showAlertReg = true;
                this.alertTextReg = "Пароли не совпадают";
                return;
            }
            let res = await fetch(
                `/api/users/register?username=${this.userNameReg}&password=${this.passwordReg}`,
                { method: 'POST' }
            );
            if (res.ok)
                window.location.href = '/'
            else {
                this.showAlertReg = true;
                this.alertTextReg = await res.text();
            }

        }
    }
};

Vue.createApp(App).mount('#vueapp')
