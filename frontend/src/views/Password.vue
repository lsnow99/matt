<template>
<input placeholder="password" type="text" v-model="password" />
<button @click="submit">Submit</button>
</template>

<script>
export default {
    data() {
        return {
            password: ''
        }
    },
    methods: {
        submit() {
            document.cookie = `password=${this.password}`
            this.$router.push({name: 'home'})
        }
    }
}
</script>