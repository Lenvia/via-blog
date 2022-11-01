import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/ant-ui'
import './assets/css/style.css'

axios.defaults.baseURL = 'http://localhost:8008/api/v1'
Vue.prototype.$http = axios


Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
