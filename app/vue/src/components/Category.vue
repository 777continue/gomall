<template>
    <div class="row">
      <div v-for="item in items" :key="item.id" class="card border-0 col-lg-4 col-md-6 col-sm-12 p-1">
        <router-link :to="`/product?id=${item.id}`" class="btn">
          <div class="card-body row">
            <img :src="item.picture" class="col-lg-6 col-sm-12" alt="product image"
                 style="max-height: 100%; min-height: 100%;">
            <div class="col-lg-6 col-sm-12 flex-column align-self-end">
              <div class="m-2">{{ item.name }}</div>
              <div class="m-1">${{ item.price }}</div>
            </div>
          </div>
        </router-link>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        items: [],
        categoryType: ''
      }
    },
    watch: {
      '$route.params.type': {
        immediate: true,
        handler(newType) {
          this.fetchCategoryData(newType)
        }
      }
    },
    methods: {
      async fetchCategoryData(type) {
        try {
          const response = await fetch(`http://localhost:8080/category/${type}`)
          if (response.ok) {
            const data = await response.json()
            this.items = data.items
          }
        } catch (error) {
          console.error('获取分类商品信息失败:', error)
        }
      }
    }
  }
  </script>
  
  <style scoped>
  .card {
    transition: transform 0.3s ease;
  }
  
  .card:hover {
    transform: scale(1.05);
  }
  
  img {
    object-fit: cover;
  }
  </style>