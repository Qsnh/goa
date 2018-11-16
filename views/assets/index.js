import Vue from 'vue'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import './scss/app.scss'
import Axios from 'axios'

window.axios = Axios
Vue.use(mavonEditor)

Vue.component('editor', () => import('./components/editor.vue'))
Vue.component('captcha', () => import('./components/captcha.vue'))
Vue.component('avatar', () => import('./components/avatar.vue'))

const app = new Vue({
    el: '#app'
})