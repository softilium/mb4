const App = {
    data() {
        return {
            userName: "",
            password: "",
            userNameReg: "",
            passwordReg: ""
        }
    },
    methods: {
        async login() {
            let res = await fetch(
                "/api/login?username=" + this.userName + "&password=" + this.password,
                { method: 'POST' }
            );
            if (res.ok)
                window.location.href = '/'
            else
                alert("Login failed");

        },
        async register() {
            let res = await fetch(
                "/api/register?username=" + this.userNameReg + "&password=" + this.passwordReg,
                { method: 'POST' }
            );
            if (res.ok)
                window.location.href = '/'
            else
                alert("Register failed");


        }
    }
};

Vue.createApp(App).mount('#vueapp')
