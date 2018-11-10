<template>
	<d2-container>
		<el-table class="mb-10"
		    :data="questions"
		    style="width: 100%">
		    <el-table-column
		      label="ID">
		      <template slot-scope="scope">
		        <span>{{ scope.row.id }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="分类名">
		      <template slot-scope="scope">
		        <span>{{ scope.row.Category.name }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="问题名">
		      <template slot-scope="scope">
		        <span>{{ scope.row.title }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="用户">
		      <template slot-scope="scope">
		        <p><span>{{ scope.row.User.nickname }}</span></p>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="回答数">
		      <template slot-scope="scope">
		        <span>{{ scope.row.answer_count }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column
		      label="创建时间">
		      <template slot-scope="scope">
		        <i class="el-icon-time"></i>
		        <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
		      </template>
		    </el-table-column>
		    <el-table-column label="操作">
		      <template slot-scope="scope">
		        <el-button
		          size="mini"
		          type="danger"
		          @click="handleDelete(scope.$index, scope.row)">删除</el-button>
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
		this.getQuestions()
	},
	data () {
		return {
			questions: [],
			page: 1,
			page_size: 10,
			total: 0
		}
	},
	methods: {
		getQuestions() {
			request.get('/questions').then(res => {
				res = res.data
				this.questions = res.questions
				this.page = res.page
				this.page_size = res.page_size
				this.total = res.total
			}).catch(err => {
				this.$message.warning(err)
			})
		},
		handleDelete(index, row) {
			this.$confirm('确定删除, 是否继续?', '提示', {
	          confirmButtonText: '确定',
	          cancelButtonText: '取消',
	          type: 'warning'
	        }).then(() => {
	        	request.delete(`/question/${row.id}`).then(res => {
	        		this.$message.success('删除成功')
	        		this.getQuestions()
	        	}).catch(err => {
	        		console.log(err)
	        	});
	        })
		},
		handleCurrentChange() {
			this.getQuestions()
		}
	}
}
</script>

<style lang="scss">
@import "@/assets/style/diy.scss";
</style>
