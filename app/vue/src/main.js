import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import http from './http'
import 'bootstrap/dist/css/bootstrap.min.css'
import * as bootstrap from 'bootstrap'

// 检验Bootstrap是否成功引入
if (typeof bootstrap === 'undefined') {
    console.error('Bootstrap未正确引入!请检查引入路径和版本。')
} else {
    console.log('Bootstrap已成功引入,版本:', bootstrap.Tooltip.VERSION)
}


const app = createApp(App)
app.use(router)
app.config.globalProperties.$http = http
app.config.globalProperties.$bootstrap = bootstrap
app.mount('#app')