<script setup lang="ts">
    import { ref, onMounted } from 'vue';
    import { useAuthStore } from '@/stores/authStore';

    const auth = useAuthStore();

    const email = ref<string>()
    const password = ref<string>()
    const error = ref<string>()

    async function login() {
        if(!email.value || !password.value) return;
        const cred = {
            email: email.value,
            password: password.value
        }
        try {
            await auth.loginAction(cred)
        } catch (err:any) {
            error.value = err.message
        } finally {
            console.log('user signed in: ', auth.me)
        }
    }

</script>

<template>
<div class="min-h-screen bg-background flex items-center justify-center">

  <div class="w-full max-w-md bg-card border border-border rounded-xl p-8 shadow-sm">

    <h1 class="text-2xl font-semibold text-center mb-6">
      Internal Ops Portal
    </h1>

    <form @submit="login" class="space-y-4">

      <div>
        <label class="block text-sm mb-1">Email</label>
        <input
        v-model="email"
          type="email"
          required="true"
          class="w-full border border-border rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-accent"
        />
      </div>

      <div>
        <label class="block text-sm mb-1">Password</label>
        <input
            v-model="password"
          type="password"
          required="true"
          class="w-full border border-border rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-accent"
        />
      </div>

      <button
        type="submit"
        class="w-full bg-accent hover:bg-accent-hover text-white py-2 rounded-lg font-medium"
      >
        Sign in
      </button>

    </form>

    <p class="text-sm text-text-secondary text-center mt-6">
      Don't have an account?
      <a class="text-accent hover:underline cursor-pointer">
        Sign up
      </a>
    </p>

  </div>

</div>
</template>