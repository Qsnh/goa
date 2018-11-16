<template>
	<d2-container>
		<el-form ref="form" :model="category" label-width="150px">
			<el-tabs v-model="activeName" @tab-click="handleClick">
			    <el-tab-pane label="系统配置" name="system">
			    	<el-form-item label="网站名">
					    <el-input v-model="setting.APP_NAME" placeholder="网站名"></el-input>
					</el-form-item>
			    	<el-form-item label="备案信息">
					    <el-input v-model="setting.ICP" placeholder="备案信息"></el-input>
					</el-form-item>
					<el-form-item label="允许跨域域名">
					    <el-input v-model="setting.CORS_ORIGINAL" placeholder="允许跨域域名"></el-input>
					</el-form-item>
			    </el-tab-pane>
			    <el-tab-pane label="用户配置" name="member">
			    	<el-form-item label="默认头像">
					    <el-input v-model="setting.MEMBER_DEFAULT_AVATAR" placeholder="默认头像"></el-input>
					</el-form-item>
					<el-form-item label="新注册用户需要激活">
					    <el-switch v-model="setting.MEMBER_DEFAULT_IS_LOCK" active-text="是" inactive-text="否" :active-value="'1'" :inactive-value="'-1'"></el-switch>
					</el-form-item>
			    </el-tab-pane>
			    <el-tab-pane label="SEO配置" name="seo">
			    	<el-form-item label="首页标题">
					    <el-input v-model="setting.SEO_INDEX_TITLE" placeholder="首页标题"></el-input>
					</el-form-item>
					<el-form-item label="首页SEO关键字">
					    <el-input v-model="setting.SEO_INDEX_KEYWORDS" placeholder="首页SEO关键字"></el-input>
					</el-form-item>
					<el-form-item label="首页SEO描述">
					    <el-input v-model="setting.SEO_INDEX_DESCRIPTION" placeholder="首页SEO描述"></el-input>
					</el-form-item>
			    </el-tab-pane>
			</el-tabs>
			<el-form-item>
			    <el-button type="primary" @click="submit()">保存</el-button>
			</el-form-item>
		</el-form>
	</d2-container>
</template>

<script>
import request from '@/plugin/axios'
export default {
	created() {
		this.getSettingData()
	},
	data () {
		return {
			activeName: 'system',
			setting: {
				"APP_NAME": "",
				"CORS_ORIGINAL": "",
				"ICP": "",
				"MEMBER_DEFAULT_AVATAR": "",
				"MEMBER_DEFAULT_IS_LOCK": "",
				"SEO_INDEX_DESCRIPTION": "",
				"SEO_INDEX_KEYWORDS": "",
				"SEO_INDEX_TITLE": ""
			}
		}
	},
	methods: {
		getSettingData() {
			request.get(`/setting/data`).then(res => {
				this.setting = res
			})
		},
		submit () {
			request.put(`/setting/save`, this.setting).then(res => {

			})
		}
	}
}	
</script>