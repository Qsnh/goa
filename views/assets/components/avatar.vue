<template>
    <div class="row">
        <div class="col-sm text-center">
            <form action="" method="post">
                <input type="hidden" name="_xsrf" v-model="xsrf">
                <input type="hidden" name="avatar" v-model="avatar">
                <p class="mb-3">
                    <img class="img-thumbnail" v-if="avatar" width="120" height="120" :src="avatar">
                </p>
                <div class="btn-group mb-3" role="group">
                    <button type="button" class="btn btn-primary" id="pick-avatar">选择jpg,gif,png格式的图片</button>
                    <button type="submit" class="btn btn-warning">保存</button>
                </div>
            </form>
            <avatar-cropper
                    @uploaded="handleUploaded"
                    trigger="#pick-avatar"
                    upload-url="/member/upload/image" />
        </div>
    </div>
</template>

<script>
    import AvatarCropper from "vue-avatar-cropper"

    export default {
        created() {
            this.avatar = this.userAvatar
        },
        props: [
            'userAvatar', 'xsrf'
        ],
        components: {
            AvatarCropper
        },
        data() {
            return {
                avatar: undefined
            }
        },
        methods: {
            handleUploaded(resp) {
                this.avatar = resp.data.path;
            }
        }
    }
</script>