<template>
    <div>
        <mavon-editor ref="md" @imgAdd="imgUpload" :boxShadow="false" :subfield="false" v-model="content"></mavon-editor>
        <div style="display: none">
            <textarea name="description" v-model="content"></textarea>
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                content: ''
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