<template>
    <div>
        <mavon-editor :boxShadow="false" :subfield="false" v-model="content"/>
        <div style="display: none">
            <textarea v-model="content"></textarea>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                content: ''
            }
        }
    }
</script>