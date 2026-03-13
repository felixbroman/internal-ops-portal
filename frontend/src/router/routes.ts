import Login from "@/views/auth/Login.vue";

export const routes = [
    {
        path: '/login',
        name: 'login',
        component: Login,
        meta: {
            title: 'IOP-login'
        }
    }
]