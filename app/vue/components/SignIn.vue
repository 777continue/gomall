<template>
  <div class="row justify-content-center">
    <div class="col-4">
      <form @submit.prevent="handleSubmit">
        <div class="mb-3">
          <label for="email" class="form-label">Email <span class="text-danger">*</span></label>
          <input type="email" class="form-control" id="email" v-model="form.email" required>
        </div>
        <div class="mb-3">
          <label for="password" class="form-label">Password <span class="text-danger">*</span></label>
          <input type="password" class="form-control" id="password" v-model="form.password" required>
        </div>
        <div class="mb-3">
          Don't have an account, click here to <router-link to="/sign-up">Sign Up</router-link>.
        </div>
        <button type="submit" class="btn btn-primary">Sign In</button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        email: '',
        password: ''
      }
    }
  },
  mounted() {
    // 从URL参数中获取next值
    if (this.$route.query.next) {
      this.next = this.$route.query.next;
    }
  },
  methods: {
    async handleSubmit() {
      try {
        const response = await this.$http.post('/auth/login', this.form);
        const token = response.data.token;
        localStorage.setItem('JWTToken', token);
        
        const redirect = response.data.redirect || this.next;
        this.$router.push(redirect);
      } catch (error) {
          console.error('Login failed:', error);
          alert('Login failed. Please check your credentials and try again.');
      }
    }
  }
}
</script>