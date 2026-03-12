<script setup lang="ts">
    import { ref, onMounted } from 'vue';
    import { useAuthStore } from '@/stores/authStore';

    const auth = useAuthStore();

    const email = ref('')
    const password = ref('')
</script>

<template>

</template>