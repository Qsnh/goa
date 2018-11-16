<template>
    <div>
        <mavon-editor ref="md" :externalLink="false" :ishljs="false" @imgAdd="imgUpload" :boxShadow="false" :subfield="false" v-model="contents">
        </mavon-editor>
        <div style="display: none">
            <textarea name="description" v-model="contents"></textarea>
        </div>
    </div>
</template>

<script>
    export default {
        mounted() {
            this.$nextTick(() => {
                const mdEl = document.getElementsByClassName('auto-textarea-input');
                if (mdEl[0]) {
                    mdEl[0].blur();
                }
                window.scroll(0,0);
            });
        },
        created() {
            if (typeof window.text != 'undefined') {
                this.contents = window.text;
            }
        },
        data() {
            return {
                contents: ''
            }
        },
        methods: {
            imgUpload(pos, $file){
                let formdata = new FormData();
                formdata.append('file', $file);
                window.axios({
                    url: '/member/upload/image',
                    method: 'post',
                    data: formdata,
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    },
                    withCredentials: true
                }).then(res => {
                    let data = res.data;
                    if (data.code == 0) {
                        this.$refs.md.$img2Url(pos, data.data.path);
                    } else {
                        alert('图片上传失败');
                    }
                })
            }
        }
    }
</script>