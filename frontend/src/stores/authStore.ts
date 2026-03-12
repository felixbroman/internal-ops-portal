import { defineStore } from "pinia";
import api from "@/services/api";
import { ref } from "vue";
import type { LoginCreds, SignupCreds, User } from "@/types/auth";
import { signup, login, getMe } from "@/services/auth.service";

export const useAuthStore = defineStore("auth", async () => {
    const me = ref<User>();

    async function Signup(cred: SignupCreds) {
        try {
            const { user } = await signup(cred)
            me.value = user
        } catch (err: any) {
            console.error('failed to signup', err?.response?.data || err?.message)
        }
    }

    async function Login(cred:LoginCreds) {
        try {
            const { user } = await login(cred)
            me.value = user
        } catch (err: any) {
            console.error('failed to login', err?.response?.data || err?.message)
        }
    }

    async function Me() {
        try {
            const { user } = await getMe()
            me.value = user
        } catch (err: any) {
            console.error('unauthorized', err?.response?.data || err?.message)
        }
    }

    return { 
        me,
        Signup,
        Login,
    }
})