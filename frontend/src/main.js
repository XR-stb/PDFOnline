import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 创建应用实例
const app = createApp(App);

// 全局配置 axios
app.config.globalProperties.$axios = axios;

// 使用 Element Plus
app.use(ElementPlus);

// 挂载应用
app.mount('#app');
