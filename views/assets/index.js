import Vue from 'vue'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import './scss/app.scss'
import EditorComponent from './components/editor'
import Axios from 'axios'

window.axios = Axios
Vue.use(mavonEditor)

Vue.component('editor', EditorComponent)

const app = new Vue({
    el: '#app'
})