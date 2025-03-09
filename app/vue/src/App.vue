<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
      <div class="container">
        <router-link to="/">
          <img class="navbar-brand" src="/static/image/logo.jpg" alt="CloudWeGo" style="height: 3em;">
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item">
              <router-link class="nav-link active" aria-current="page" to="/">杨镇泽的小店</router-link>
            </li>
            <li class="nav-item dropdown">
              <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                  分类
              </a>
              <ul class="dropdown-menu">
                <li><router-link class="dropdown-item" to="/category/snacks">零食</router-link></li>
                <li><router-link class="dropdown-item" to="/category/drinks">饮料</router-link></li>
              </ul>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/about">关于</router-link>
            </li>
          </ul>
          <form class="d-flex ms-auto" role="search" @submit.prevent="search">
            <input class="form-control me-2" type="search" v-model="searchQuery" placeholder="搜索" aria-label="搜索">
            <button class="btn btn-outline-success" type="submit">搜索</button>
          </form>
           <!-- 添加购物车数量显示 -->
           <router-link to="/cart" class="btn position-relative">
            <i class="fa-solid fa-cart-shopping fa-xl"></i>
            <span v-if="cartNum > 0" class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
              {{ cartNum }}
            </span>
          </router-link>
          <div>
            <router-link v-if="!user_id" class="btn btn-primary ms-3" to="/sign-in">登录</router-link>
            <div v-else class="dropdown">
              <a class="ms-3 dropdown-toggle" data-bs-toggle="dropdown">
                <i class="fa-solid fa-user fa-lg"></i>
                <span>你好</span>
              </a>
              <ul class="dropdown-menu">
                <li><router-link class="dropdown-item" to="/orders">订单中心</router-link></li>
                <li><a class="dropdown-item" @click="logout">退出</a></li>
                <li v-if="isAdmin"><router-link class="dropdown-item" to="/manage-user">用户管理</router-link></li>
                <li v-if="isAdmin"><router-link class="dropdown-item" to="/manage-product">商品管理</router-link></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div class="text-bg-primary text-center">
      <p>This Website is hosted for demo purposes only. It is not an actual shop</p>
    </div>
    <main style="min-height: calc(80vh);">
      <div class="container py-3">
        <router-view></router-view>
      </div>
    </main>
  </div>
</template>

  
  <script>
  export default {
    data() {
      return {
        user_id: localStorage.getItem('jwtToken') || null,
        searchQuery: '',
        isAdmin: localStorage.getItem('admin') || false,
        cartNum: 0, // 购物车数量
      };
    },
    async mounted() {
      console.log('Initial user_id:', this.user_id); 
      await this.fetchCartNum(); // 获取购物车数量
    },
    methods: {
      async search() {
        this.$router.push({ path: '/search', query: { q: this.searchQuery } });
      },
      async logout() {
        try {
          localStorage.removeItem('jwtToken');  
          localStorage.removeItem('admin');
          this.user_id = null;
          this.isAdmin = false;
        } catch (error) {
          console.error('登出失败:', error);
        }
      },
      async fetchCartNum() {
        try {
          const response = await this.$http.get('http://localhost:8080/cart/num');
          this.cartNum = response.data.num; 
        } catch (error) {
          console.error('获取购物车数量失败:', error);
        }
      },
    }
  };
  </script>
  
  <style>
  #app {
    min-height: 100vh;
  }
  </style>