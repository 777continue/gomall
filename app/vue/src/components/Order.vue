<template>
    <div class="row">
      <div class="card border-0" style="width: 100%;">
        <div class="card-body row">
          <div v-for="order in orders" :key="order.OrderId" class="card">
            <div class="card-body">
              <h6 class="card-subtitle mb-2 text-muted">{{ order.CreatedDate }} Order ID: {{ order.OrderId }}</h6>
              <ul class="list-group col-lg-12 col-sm-15">
                <li v-for="item in order.Items" :key="item.ProductName" class="list-group-item border-0">
                  <div class="card border-0">
                    <div class="card-body row">
                      <div class="col-3">
                        <img :src="item.Picture" style="max-width: 100px;max-height: 50px" alt="">
                      </div>
                      <div class="col-3">
                        <div class="mt-1">{{ item.ProductName }}</div>
                      </div>
                      <div class="col-2">
                        <div class="mt-1">x {{ item.Qty }}</div>
                      </div>
                      <div class="col-4">
                        <div class="mt-1">Cost: {{ item.Cost }}</div>
                      </div>
                    </div>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        orders: [], // 订单列表
      };
    },
    async mounted() {
      await this.fetchOrders();
    },
    methods: {
      async fetchOrders() {
        try {
          const response = await this.$http.get('http://localhost:8080/order');
          this.orders = response.data.orders; // 假设后端返回 { orders: 订单列表 }
        } catch (error) {
          console.error('获取订单失败:', error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* 保留原有样式 */
  .card {
    margin-bottom: 20px;
  }
  </style>