<template>
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
  </template>
  
  <script>
  export default {
    data() {
      return {
        items: []
      };
    },
    async mounted() {
      try {
        const response = await fetch('/');
        if (response.ok) {
          const data = await response.json();
          this.items = data.items;
        }
      } catch (error) {
        console.error('获取商品信息失败:', error);
      }
    }
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