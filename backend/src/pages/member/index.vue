<template>
	<d2-container>
		<el-table class="mb-10"
		    :data="users"
		    style="width: 100%">
		    <el-table-column
		      label="ID">
		      <template slot-scope="scope">
		        <span>{{ scope.row.id }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="邮箱">
		      <template slot-scope="scope">
		        <span>{{ scope.row.email }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="昵称">
		      <template slot-scope="scope">
		        <span>{{ scope.row.nickname }}</span>
		      </template>
		    </el-table-column>
		     <el-table-column
		      label="状态">
		      <template slot-scope="scope">
				<el-tag type="success" v-if="scope.row.is_lock == -1">已激活</el-tag>
				<el-tag type="warning" v-if="scope.row.is_lock == 1">未激活</el-tag>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="注册时间">
		      <template slot-scope="scope">
		        <i class="el-icon-time"></i>
		        <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column label="操作">
		      <template slot-scope="scope">
		        <el-button
		          size="mini"
		          type="warning"
		          @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
		      </template>
		    </el-table-column>
		  </el-table>
		  <el-row>
			<el-col :span="24">
				<el-pagination
				  background
				  layout="total, prev, pager, next"
				  @current-change="handleCurrentChange"
				  :total="total"
				  :current-page="page"
				  :page-size="page_size">
				</el-pagination>
			</el-col>
		</el-row>
	</d2-container>
</template>

<script>
import request from '@/plugin/axios'
export default {
	created() {
		this.getUsers()
	},
	data () {
		return {
			users: [],
			page: 1,
			page_size: 10,
			total: 0
		}
	},
	methods: {
		getUsers() {
			request.get('/users').then(res => {
				res = res.data
				this.users = res.users
				this.page = res.page
				this.page_size = res.page_size
				this.total = res.total
			}).catch(res => {
				this.$message.warning('无法获取数据')
			})
		},
		handleEdit(index, row) {
			this.$router.push({name: 'member-edit', params: {
				id: row.id
			}})
		},
		handleCurrentChange() {
			this.getUsers()
		}
	}
}
</script>

<style lang="scss">
@import "@/assets/style/diy.scss";
</style>
