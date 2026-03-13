<script setup lang="ts">
import { RouterView } from 'vue-router';
import { ref, onMounted } from 'vue';
import { useAuthStore } from './stores/authStore';
import router from './router';

const auth = useAuthStore();

onMounted(() => {

})
</script>

<template>
  <main class="flex flex-col box-border items-center justify-center  w-screen min-h-screen">
     <RouterView />
   </main>
</template>
