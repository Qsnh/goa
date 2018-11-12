<template>
	<d2-container>
		<el-form ref="form" :model="user" label-width="100px">
			<el-form-item label="密码">
			    <el-input v-model="user.password" placeholder="密码"></el-input>
			</el-form-item>
			<el-form-item label="状态">
				<el-switch v-model="user.is_lock" :active-value="-1" :inactive-value="1" active-text="已激活" inactive-text="未激活"></el-switch>
			</el-form-item>
			<el-form-item>
			    <el-button type="primary" @click="submit()">保存</el-button>
			</el-form-item>
		</el-form>
	</d2-container>
</template>

<script>
import querystring from 'querystring'
import request from '@/plugin/axios/index'
export default {
	created() {
		this.user.id = this.$route.params.id
		this.getUser()
	},
	data () {
		return {
			user: {
				id: 0,
				password: '',
				is_lock: 1
			}
		}
	},
	methods: {
		submit() {
			request.put(`/user/${this.user.id}`, {
				password: this.user.password,
				is_lock: this.user.is_lock
			}).then(res => {
				this.$message.success('修改成功')
			})
		},
		getUser() {
			request.get(`/user/${this.user.id}`).then(res => {
				this.user = res.user
			}).catch(err => {
				this.$message.warning(err)
			})
		}
	}
}
</script>