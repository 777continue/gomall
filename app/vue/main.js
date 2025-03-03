import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import http from './http'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'

const app = createApp(App)
app.use(router)
app.config.globalProperties.$http = http
app.mount('#app')