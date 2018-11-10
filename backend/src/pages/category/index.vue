<template>
	<d2-container>
		<el-row class="mb-10">
			<el-col :span="24">
				<el-button @click="goCategoryCreatePage()" type="primary">添加分类</el-button>
			</el-col>
		</el-row>
		<el-table
		    :data="categories"
		    style="width: 100%">
		    <el-table-column
		      label="分类名">
		      <template slot-scope="scope">
		        <span>{{ scope.row.name }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="日期">
		      <template slot-scope="scope">
		        <i class="el-icon-time"></i>
		        <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column label="操作">
		      <template slot-scope="scope">
		        <el-button
		          size="mini"
		          @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
		        <el-button
		          size="mini"
		          type="danger"
		          @click="handleDelete(scope.$index, scope.row)">删除</el-button>
		      </template>
		    </el-table-column>
		  </el-table>
	</d2-container>
</template>

<script>
import request from '@/plugin/axios'
export default {
	created() {
		this.getCategories()
	},
	data () {
		return {
			categories: [],
			page: 1,
			page_size: 10
		}
	},
	methods: {
		getCategories() {
			request.get('/categories').then(res => {
				this.categories = res.categories
			}).catch(res => {
				this.$message.warning('无法获取数据')
			})
		},
		handleEdit(index, row) {
			this.$router.push({name: "category-edit", params: {id: row.id}});
		},
		handleDelete(index, row) {
			this.$confirm('确定删除, 是否继续?', '提示', {
	          confirmButtonText: '确定',
	          cancelButtonText: '取消',
	          type: 'warning'
	        }).then(() => {
	        	request.delete(`/category/${row.id}`).then(res => {
	        		this.$message.success('删除成功')
	        		this.getCategories()
	        	}).catch(err => {
	        		console.log(err)
	        	});
	        })
		},
		goCategoryCreatePage() {
			this.$router.push({name: "category-create"});
		}
	}
}
</script>

<style lang="scss">
@import "@/assets/style/diy.scss";
</style>
