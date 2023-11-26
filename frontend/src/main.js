import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';

const app = createApp(App);

// 全局配置 axios
app.config.globalProperties.$axios = axios;

app.mount('#app');
