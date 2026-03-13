import { defineStore } from "pinia";
import { ref } from "vue";
import type { LoginCreds, SignupCreds, User } from "@/types/auth";
import { signup, login, getMe } from "@/services/auth.service";

export const useAuthStore = defineStore("auth", () => {
    const me = ref<User>();

    async function signupAction(cred: SignupCreds) {
        try {
            const { user } = await signup(cred)
            me.value = user
        } catch (err: any) {
            console.error('failed to signup', err?.response?.data || err?.message)
        }
    }

    async function loginAction(cred:LoginCreds) {
        try {
            const { user } = await login(cred)
            me.value = user
        } catch (err: any) {
            console.error('error in loginAction', err)
            if (err.response?.status === 401) {
                throw new Error(err.response?.message)
            }
        }
    }

    async function meAction() {
        if(me.value) return;
        try {
            const { user } = await getMe()
            me.value = user
        } catch (err: any) {
            console.error('unauthorized', err?.response?.data || err?.message)
        }
    }

    return { 
        me,
        meAction,
        signupAction,
        loginAction,
    }
})