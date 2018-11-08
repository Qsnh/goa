import Vue from 'vue'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import './scss/app.scss'
import EditorComponent from './components/editor.vue'
import CaptchaComponent from './components/captcha.vue'
import AvatarComponent from './components/avatar.vue'
import Axios from 'axios'

window.axios = Axios
Vue.use(mavonEditor)

Vue.component('editor', EditorComponent)
Vue.component('captcha', CaptchaComponent)
Vue.component('avatar', AvatarComponent)

const app = new Vue({
    el: '#app'
})