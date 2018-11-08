<template>
    <img @click="change" class="rounded" :src="src" width="240" height="80">
</template>
<script>
    export default {
        data () {
            return {
                src: '/captcha'
            }
        },
        methods: {
            change() {
                this.src = '/captcha?' + Math.random()
            }
        }
    }
</script>