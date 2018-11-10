<template>
	<d2-container>
		<el-form ref="form" :model="category" label-width="100px">
			<el-form-item label="分类名">
			    <el-input v-model="category.name" placeholder="分类名"></el-input>
			</el-form-item>
			<el-form-item>
			    <el-button type="primary" @click="submit()">创建</el-button>
			</el-form-item>
		</el-form>
	</d2-container>
</template>

<script>
import querystring from 'querystring'
import request from '@/plugin/axios/index'
export default {
	data () {
		return {
			category: {
				name: ''
			}
		}
	},
	methods: {
		submit() {
			if (this.category.name == '') {
				this.$message.warning('请输入分类名')
				return
			}
			request.post('/category', "name="+this.category.name).then(res => {
				this.$message.success('创建成功')
				this.$router.push({name: 'category'})
			}).catch(err => {
				this.$message.warning(err)
			})
		}
	}
}
</script>