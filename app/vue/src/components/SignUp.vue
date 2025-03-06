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
                    <label for="password_confirm" class="form-label">Password Confirm <span
                            class="text-danger">*</span></label>
                    <input type="password" class="form-control" id="password_confirm" v-model="form.password_confirm"
                        required>
                </div>
                <div class="mb-3">
                    Already have an account, click here to <router-link to="/sign-in">Sign In</router-link>.
                </div>
                <button type="submit" class="btn btn-primary">Sign Up</button>
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
                password: '',
                password_confirm: ''
            }
        }
    },
    methods: {
        async handleSubmit() {
            console.log('Submitting form:', this.form); 
            if (this.form.password !== this.form.password_confirm) {
                alert('Passwords do not match');
                return;
            }

            try {
                const response = await this.$http.post('http://localhost:8080/auth/register', this.form);
                const token = response.data.token;
                localStorage.setItem('jwtToken', token);
                this.$router.push('/');
            } catch (error) {
                console.error('Registration failed:', error);
                alert('Registration failed. Please try again.');
            }
        }
    }
}
</script>