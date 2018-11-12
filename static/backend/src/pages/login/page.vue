<template>
  <div class="login-page">
    <div class="layer bg" id="login"></div>
    <div class="layer flex-center">
      <h2>GOA</h2>
      <div class="form-group">
        <el-card>
          <el-form ref="loginForm" label-position="top" :rules="rules" :model="formLogin" size="default">
            <el-form-item prop="username">
              <el-input type="text" v-model="formLogin.username" placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input type="password" v-model="formLogin.password" placeholder="密码"></el-input>
            </el-form-item>
            <el-button size="default" @click="submit" type="primary" class="button-login">登录</el-button>
          </el-form>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script>
require('particles.js')
import config from './config/default'
import { mapActions } from 'vuex'
export default {
  data () {
    return {
      // 表单
      formLogin: {
        username: '',
        password: ''
      },
      // 校验
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      }
    }
  },
  mounted () {
    particlesJS('login', config)
  },
  beforeDestroy () {
    if (pJSDom && pJSDom.length > 0) {
      pJSDom[0].pJS.fn.vendors.destroypJS()
      pJSDom = []
    }
  },
  methods: {
    ...mapActions('d2admin/account', [
      'login'
    ]),
    submit () {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.login({
            vm: this,
            username: this.formLogin.username,
            password: this.formLogin.password
          })
        } else {
          this.$message.error('请输入正确的数据')
        }
      })
    }
  }
}
</script>

<style lang="scss">
@import './style.scss';
</style>
