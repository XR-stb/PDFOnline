<template>
    <div class="app">
        <div class="login-box">
            <h1 class="log-title">Login</h1>
            <div class="input-box">
                <input type="text" class="username" placeholder="username">
                <input type="password" class="password" placeholder="password">
                <p class="p-btn"><button class="login-btn" type="button" onclick="login()">login</button></p>
                <p class="more">
                    <a href="SignUp.html">注册 | </a>
                    <a href="javascript:alert('放心，这里啥也没有');">忘记密码 | </a>
                    <a href="javascript:alert('这里也是空的');">更多</a>
                </p>
            </div>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style>
@import "../css/login.css";
</style>