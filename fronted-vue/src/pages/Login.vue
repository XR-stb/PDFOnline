<template>
    <div class="app">
      <div class="login-box">
        <h1 class="log-title">{{ isLogin ? 'Login' : 'Register' }}</h1>
        <div class="input-box">
          <input v-model="username" type="text" class="username" placeholder="username">
          <input v-model="password" type="password" class="password" placeholder="password">
          
          <!-- Additional fields for registration -->
          <div v-if="!isLogin">
            <input v-model="email" type="email" class="email" placeholder="email">
            <input v-model="captcha" type="text" class="captcha" placeholder="captcha">
            <button class="captcha-btn" type="button" @click="fetchCaptcha">Get Captcha</button>
          </div>
          
          <p class="p-btn"><button class="login-btn" type="button" @click="performAction">{{ isLogin ? 'Login' : 'Register' }}</button></p>
          <p class="more">
            <a :href="isLogin ? 'SignUp.html' : 'javascript:void(0)'" @click="toggleAction">{{ isLogin ? '注册' : '返回登录' }} | </a>
            <a href="javascript:alert('放心，这里啥也没有');">忘记密码 | </a>
            <a href="javascript:alert('这里也是空的');">更多</a>
          </p>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        username: '',
        password: '',
        email: '', // New field for email
        captcha: '', // New field for captcha
        isLogin: true,
      };
    },
    methods: {
      async fetchCaptcha() {
        try {
          // Make a request to fetch captcha
          const response = await axios.post('http://127.0.0.1:8080/v1/users/captcha', {
            email: this.email,
          });
  
          // Handle the response according to your needs
          alert(`Captcha fetched successfully!\nResponse: ${JSON.stringify(response.data)}`);
        } catch (error) {
          alert(`Failed to fetch captcha!\nError: ${error.message}`);
        }
      },
      async performAction() {
        try {
          if (this.isLogin === true) {
            const response = await axios.post(`http://127.0.0.1:8080/v1/users/login`, {
                username: this.username,
                password: this.password,
            });
            alert(`Operation successful!\nResponse: ${JSON.stringify(response.data)}`);
          } else {
            const response = await axios.post(`http://127.0.0.1:8080/v1/users`, {
                username: this.username,
                password: this.password,
                email: this.email,
                captcha: this.captcha, // Include captcha in the registration request
            });
            alert(`Operation successful!\nResponse: ${JSON.stringify(response.data)}`);
          }
        } catch (error) {
          alert(`Operation failed!\nError: ${error.message}`);
        }
      },
      toggleAction() {
        this.isLogin = !this.isLogin;
      },
    },
  };
  </script>
  
  <style>
  @import "../css/login.css";
  </style>
  