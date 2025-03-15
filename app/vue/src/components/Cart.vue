<template>
    <div class="row">
      <ul class="list-group">
        <li v-for="item in items" :key="item.id" class="list-group-item">
          <div class="card border-0">
            <div class="card-body row">
              <div class="col-4">
                <img :src="item.picture" style="max-width: 100px;max-height: 50px" alt="">
              </div>
              <div class="col-8">
                <div class="mt-1">{{ item.name }}</div>
                <div class="mt-1">Single Price: ${{ item.price }}</div>
                <div class="mt-1">Qty: {{ item.qty }}</div>
              </div>
            </div>
          </div>
        </li>
      </ul>
      <div v-if="items.length > 0" class="mt-3 mb-5">
        <div class="float-end">
          <div class="m-3 text-danger">Total: ${{ total }}</div>
          <router-link to="/checkout" class="btn btn-lg btn-success float-end">Check out</router-link>
        </div>
      </div>
      <div v-else>
        <h1 class="text-center text-danger">Your Cart is empty</h1>
        <div class="text-center"><router-link to="/">Shop Hot Sale</router-link></div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        items: [], // 购物车商品列表
        total: 0, // 总价
      };
    },
    async mounted() {
      await this.fetchCartItems();
    },
    methods: {
      async fetchCartItems() {
        try {
          const response = await this.$http.get('http://localhost:8080/cart');
          this.items = response.data.items;
          this.total = response.data.total
        } catch (error) {
            console.error('搜索失败:', error);
            alert('搜索失败，请稍后重试。');
        }
      },

    },
  };
  </script>
  
  <style scoped>
  .list-group-item {
    margin-bottom: 10px;
  }
  </style>