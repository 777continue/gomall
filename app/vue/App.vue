<template>
    <div id="app">
      <nav class="navbar bg-body-tertiary">
        <div class="container">
          <router-link to="/">
            <img src="/static/image/logo.jpg" alt="CloudWeGo" style="height: 3em;">
          </router-link>
  
          <ul class="navbar-nav">
            <li>
              <router-link to="/">首页</router-link>
            </li>
            <li class="dropdown">
              <a class="dropdown-toggle" data-bs-toggle="dropdown">分类</a>
              <ul class="dropdown-menu">
                <li><router-link to="/category/snacks">零食</router-link></li>
                <li><router-link to="/category/drinks">饮料</router-link></li>
              </ul>
            </li>
            <li>
              <router-link to="/about">关于</router-link>
            </li>
          </ul>
  
          <form @submit.prevent="search">
            <input v-model="searchQuery" placeholder="搜索">
            <button type="submit">搜索</button>
          </form>
  
          <div>
            <router-link v-if="!user_id" to="/sign-in">登录</router-link>
            <div v-else class="dropdown">
                <a class="dropdown-toggle" data-bs-toggle="dropdown">
                <i class="fa-solid fa-user fa-lg"></i>
                <span>你好</span>
                </a>
                <ul class="dropdown-menu">
                <li><router-link to="/orders">订单中心</router-link></li>
                <li><a @click="logout">退出</a></li>
                </ul>
            </div>
          </div>
        </div>
      </nav>
  
      <router-view></router-view>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        user_id: null,
        searchQuery: ''
      };
    },
    mounted() {
      async function fetchUserInfo() {
        try {
          const response = await fetch('/');
          if (response.ok) {
            const data = await response.json();
            this.user_id = data.user_id;
          }
        } catch (error) {
          console.error('获取用户信息失败:', error);
        }
      }
      fetchUserInfo.call(this);
    },
    methods: {
      search() {
        this.$router.push({ path: '/search', query: { q: this.searchQuery } });
      },
      async logout() {
        try {
          await fetch('/auth/logout', { method: 'POST' });
          this.user_id = null;
          this.$router.push('/sign-in');
        } catch (error) {
          console.error('登出失败:', error);
        }
      }
    }
  };
  </script>
  
  <style>
  #app {
    min-height: 100vh;
  }
  </style>