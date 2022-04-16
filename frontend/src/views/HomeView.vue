<script setup>
import TheWelcome from '@/components/TheWelcome.vue'
</script>

<template>
  <main>
    This is secret
    <button @click="logout">Logout</button>
  </main>
</template>

<script>
export default {
  methods: {
    logout() {
      document.cookie = 'password=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      this.$router.replace({name: 'login'})
    }
  }
}
</script>
