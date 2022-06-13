const App = {
    data() {
        return {
        }
    },
    methods: {
        async logout() {
            let res = await fetch(
                "/api/logout",
                { method: 'POST' }
            );
            if (res.ok)
                window.location.href = '/'
            else {
                this.showAlert = true;
                this.alertText = await res.text();
            }
        }
    }
};

Vue.createApp(App).mount('#vueapp')
