
<template>
    <div class="container mt-4">
      <h2 class="mb-4">用户管理</h2>
      
      <!-- 添加用户表单 -->
      <div class="card mb-4">
        <div class="card-body">
          <h5 class="card-title">添加用户</h5>
          <form @submit.prevent="addUser">
            <div class="mb-3">
              <label for="email" class="form-label">邮箱</label>
              <input type="email" class="form-control" id="email" v-model="newUser.email" required>
            </div>
            <div class="mb-3">
              <label for="password" class="form-label">密码</label>
              <input type="password" class="form-control" id="password" v-model="newUser.password" required>
            </div>
            <div class="mb-3 form-check">
              <input type="checkbox" class="form-check-input" id="isAdmin" v-model="newUser.isAdmin">
              <label class="form-check-label" for="isAdmin">管理员</label>
            </div>
            <button type="submit" class="btn btn-primary">添加用户</button>
          </form>
        </div>
      </div>
  
      <!-- 用户列表 -->
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">用户列表</h5>
          <table class="table table-hover">
            <thead>
              <tr>
                <th scope="col">ID</th>
                <th scope="col">邮箱</th>
                <th scope="col">角色</th>
                <th scope="col">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in users" :key="user.user_id">
                <td>{{ user.user_id }}</td>
                <td>{{ user.email }}</td>
                <td>{{ user.isAdmin ? '管理员' : '普通用户' }}</td>
                <td>
                  <button class="btn btn-danger btn-sm" @click="deleteUser(user.user_id)">删除</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { Modal } from 'bootstrap';
  export default {

    data() {
      return {
        users: [],
        newUser: {
          email: '',
          password: '',
          isAdmin: false
        }
      };
    },
    async created() {
      await this.fetchUsers();
    },
    methods: {
      async fetchUsers() {
        try {
          const response = await this.$http.get('http://localhost:8080/users');
          this.users = response.data.users;
        } catch (error) {
          console.error('获取用户列表失败:', error);
        }
      },
      async addUser() {
        try {
          await this.$http.post('http://localhost:8080/users', this.newUser);
          this.newUser = { email: '', password: '', isAdmin: false };
          await this.fetchUsers();
        } catch (error) {
          console.error('添加用户失败:', error);
        }
      },
      async deleteUser(userId) {
        try {
          await this.$http.delete(`http://localhost:8080/users/${userId}`);
          await this.fetchUsers();
        } catch (error) {
          console.error('删除用户失败:', error);
        }
      }
    }
  };
  </script>
  
  <style scoped>
  .card {
    margin-bottom: 20px;
  }
  
  .table {
    margin-top: 20px;
  }
  </style>