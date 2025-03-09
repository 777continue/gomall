<template>
    <div>
      <!-- 搜索条件显示 -->
      <div v-if="q" class="d-inline-block border border-success rounded-4 p-1 ps-4 pe-4">
        <div>
          {{ q }}
          <a href="/" class="ps-4">
            <button type="button" class="btn btn-sm btn-danger" aria-label="Close">X</button>
          </a>
        </div>
      </div>
  
    <!-- 商品列表 -->
    <div class="row">
      <div v-for="item in items" :key="item.id" class="card col-xl-3 col-lg-4 col-md-6 col-sm-12 p-5 border-0">
        <router-link :to="`/product?id=${item.id}`" class="btn">
          <img :src="item.picture" class="card-img-top" alt="product image">
          <div class="card-body">
            <p class="card-text">{{ item.name }}</p>
            <h5 class="card-title">{{ item.price }}</h5>
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>
  
  <script>
  export default {
    data() {
      return {
        q: this.$route.query.q || '', // 从路由参数中获取搜索关键词
        items: [], // 商品列表
      };
    },
    async mounted() {
      if (this.q) {
        await this.fetchSearchResults();
      }
    },
    methods: {
      async fetchSearchResults() {
        try {
          const response = await this.$http.get('http://localhost:8080/search', {
            params: { q: this.q }
          });
          this.items = response.data.items; // 假设后端返回的是商品列表
        } catch (error) {
          console.error('搜索失败:', error);
          alert('搜索失败，请稍后重试。');
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .card {
    transition: transform 0.3s ease;
  }
  
  .card:hover {
    transform: scale(1.05);
  }
  
  .card-img-top {
    height: 200px;
    object-fit: cover;
  }
  </style>