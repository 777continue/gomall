<template>
  <div>
    <div class="row">
      <div class="card border-0" style="width: 100%;">
        <div class="card-body row">
          <div id="productPicture" class="carousel slide col-lg-6 col-sm-12">
            <div class="carousel-indicators">
              <button 
                v-for="(_, index) in 3" 
                :key="index"
                type="button" 
                data-bs-target="#productPicture" 
                :data-bs-slide-to="index"
                :class="{ active: index === 0 }"
                :aria-label="`Slide ${index + 1}`"
              ></button>
            </div>
            <div class="carousel-inner">
              <div 
                v-for="(_, index) in 3" 
                :key="index"
                class="carousel-item" 
                :class="{ active: index === 0 }"
              >
                <img :src="item.picture" class="d-block w-100" alt="Product Image">
              </div>
            </div>
            <button class="carousel-control-prev" type="button" data-bs-target="#productPicture" data-bs-slide="prev">
              <span class="carousel-control-prev-icon" aria-hidden="true"></span>
              <span class="visually-hidden">Previous</span>
            </button>
            <button class="carousel-control-next" type="button" data-bs-target="#productPicture" data-bs-slide="next">
              <span class="carousel-control-next-icon" aria-hidden="true"></span>
              <span class="visually-hidden">Next</span>
            </button>
          </div>
          <div class="col-lg-1"></div>
          <div class="col-lg-5 col-sm-12 flex-column align-self-center">
            <form @submit.prevent="addToCart">
              <h5 class="card-title">{{ item.name }}</h5>
              <p class="card-text">{{ item.description }}</p>
              <p class="card-text">${{ item.price }}</p>
              <input type="hidden" :value="item.id" name="productId">
              <label for="productNum">数量：</label>
              <input 
                type="number" 
                class="form-control mt-3" 
                id="productNum" 
                name="productNum" 
                v-model="quantity"
                min="1"
              />
              <button type="submit" class="btn btn-primary mt-3">Add to Cart</button>
            </form>
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
      item: {},
      quantity: 1,
      loading: true,
      error: null
    }
  },
  async created() {
    try {
      const productId = this.$route.query.id
      const response = await fetch(`http://localhost:8080/product?id=${productId}`)
      if (response.ok) {
          const data = await response.json();
          this.item = data.item;
        }else {
        throw new Error('Failed to fetch product data')
      }
    } catch (error) {
      console.error('获取产品信息失败:', error)
      this.error = error.message
    } finally {
      this.loading = false
    }
  },
  methods: {
    async addToCart() {
      const formData = new FormData();
      formData.append('productId', this.item.id);
      formData.append('productNum', this.quantity);

      const response = await this.$http.post('http://localhost:8080/cart', formData );
      if (response) {
          this.$router.push('/cart');
      } else {
          throw new Error('Failed to add item to cart');
      }
    },  
  },
};
</script>
<style scoped>
/* 可以添加自定义样式 */
</style>