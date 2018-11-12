<template>
  <d2-container>
    <el-row class="text-center">
      <el-col :span="8">
        <h1>今日注册</h1>
        <h4>{{count.register}}</h4>
      </el-col>
      <el-col :span="8">
        <h1>今日提问</h1>
        <h4>{{count.question}}</h4>
      </el-col>
      <el-col :span="8">
        <h1>今日回答</h1>
        <h4>{{count.answer}}</h4>
      </el-col>
    </el-row>
  </d2-container>
</template>

<script>
import request from '@/plugin/axios'
export default {
  created() {
    this.getCountData()
  },
  data () {
    return {
      count: {
        register: 0,
        question: 0,
        answer: 0
      }
    }
  },
  methods: {
    getCountData() {
      request.get('/dashboard').then(res => {
        res = res.data
        this.count.register = res.count_register
        this.count.question = res.count_question
        this.count.answer = res.count_answer
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~@/assets/style/public.scss';
.text-center {
  text-align: center;
}
</style>
