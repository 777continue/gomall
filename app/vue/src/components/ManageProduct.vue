<template>
  <div class="container mt-4">
    <!-- 商品管理 -->
    <div class="card mb-4">
      <div class="card-header">
        <h3>商品管理</h3>
        <div class="mb-3">
          <button class="btn btn-primary me-2" @click="listProducts">刷新商品列表</button>
          <button class="btn btn-success" data-bs-toggle="modal" data-bs-target="#addProductModal">添加商品</button>
        </div>
      </div>
      <div class="card-body">
        <div class="row">
          <div v-for="product in products" :key="product.id" class="card border-0 col-lg-4 col-md-6 col-sm-12 p-1">
            <div class="card-body row">
              <img :src="product.picture" class="col-lg-6 col-sm-12" alt="product image"
                   style="max-height: 100%; min-height: 100%;">
              <div class="col-lg-6 col-sm-12 flex-column align-self-end">
                <div class="m-2">{{ product.name }}</div>
                <div class="m-1">${{ product.price }}</div>
                <div class="m-1">库存: {{ product.stock }}</div>
                <div class="mt-2">
                  <button class="btn btn-sm btn-warning me-2" @click="openEditModal(product)">编辑</button>
                  <button class="btn btn-sm btn-danger" @click="deleteProduct(product.id)">删除</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分类管理 -->
    <div class="card">
      <div class="card-header">
        <h3>分类管理</h3>
        <div class="mb-3">
          <button class="btn btn-primary me-2" @click="listCategories">刷新分类列表</button>
          <button class="btn btn-success" data-bs-toggle="modal" data-bs-target="#addCategoryModal">添加分类</button>
        </div>
      </div>
      <div class="card-body">
        <div class="row">
          <div v-for="category in categories" :key="category.id" class="card border-0 col-lg-4 col-md-6 col-sm-12 p-1">
            <div class="card-body row">
              <div class="col-12 flex-column align-self-end">
                <div class="m-2">
                  <a href="#" @click.prevent="filterByCategory(category.id)">{{ category.name }}</a>
                </div>
                <div class="mt-2">
                  <button class="btn btn-sm btn-danger" @click="deleteCategory(category.id)">删除</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

      <!-- 添加商品模态框 -->
      <div class="modal fade" id="addProductModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">添加商品</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addProduct">
              <div class="mb-3">
                <label class="form-label">商品名称</label>
                <input type="text" class="form-control" v-model="newProduct.name" required>
              </div>
              <div class="mb-3">
                <label class="form-label">商品描述</label>
                <textarea class="form-control" v-model="newProduct.description" required></textarea>
              </div>
              <div class="mb-3">
                <label class="form-label">商品图片</label>
                <input type="file" class="form-control" @change="onFileChange($event, 'new')" accept="image/*">
              </div>
              <div class="mb-3">
                <label class="form-label">商品价格</label>
                <input type="number" class="form-control" v-model="newProduct.price" required>
              </div>
              <div class="mb-3">
                <label class="form-label">库存数量</label>
                <input type="number" class="form-control" v-model="newProduct.stock" required>
              </div>
              <button type="submit" class="btn btn-primary">添加</button>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- 编辑商品模态框 -->
    <div class="modal fade" id="editProductModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">编辑商品</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateProduct">
              <div class="mb-3">
                <label class="form-label">商品名称</label>
                <input type="text" class="form-control" v-model="editProduct.name" required>
              </div>
              <div class="mb-3">
                <label class="form-label">商品描述</label>
                <textarea class="form-control" v-model="editProduct.description" required></textarea>
              </div>
              <div class="mb-3">
                <label class="form-label">商品图片</label>
                <input type="file" class="form-control" @change="onFileChange($event, 'edit')" accept="image/*">
              </div>
              <div class="mb-3">
                <label class="form-label">商品价格</label>
                <input type="number" class="form-control" v-model="editProduct.price" required>
              </div>
              <div class="mb-3">
                <label class="form-label">库存数量</label>
                <input type="number" class="form-control" v-model="editProduct.stock" required>
              </div>
              <button type="submit" class="btn btn-primary">更新</button>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- 添加分类模态框 -->
    <div class="modal fade" id="addCategoryModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">添加分类</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addCategory">
              <div class="mb-3">
                <label class="form-label">分类名称</label>
                <input type="text" class="form-control" v-model="newCategory.name" required>
              </div>
              <button type="submit" class="btn btn-primary">添加</button>
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
      products: [],
      categories: [],
      newProduct: {
        name: '',
        description: '',
        picture: null,
        price: 0,
        stock: 0
      },
      editProduct: {
        id: null,
        name: '',
        description: '',
        picture: null,
        price: 0,
        stock: 0
      },
      newCategory: {
        name: ''
      }
    }
  },
  mounted() {
    this.checkBootstrap()
  },
  methods: {
    checkBootstrap() {
      if (typeof bootstrap === 'undefined') {
        console.error('Bootstrap未正确引入！请检查main.js中的引入配置。')
      } else {
        console.log('Bootstrap已成功引入，版本：', bootstrap.Tooltip.VERSION)
      }
    },
    async listProducts() {
      try {
        // 调用/category接口获取所有商品
        const response = await fetch('http://localhost:8080/category/all');
        if (response.ok) {
          const data = await response.json()
          this.products = data.items
        }
      } catch (error) {
        console.error('获取商品列表失败:', error);
      }
    },
    openEditModal(product) {
      this.editProduct = {
        id: product.id,
        name: product.name,
        description: product.description,
        picture: product.picture,
        price: product.price,
        stock: product.stock
      }
      // 使用Bootstrap显示模态框
      const modal = new Modal(document.getElementById('editProductModal'), {
        keyboard: false
      })
      modal.show()
    },
    async addProduct() {
      try {
        const formData = new FormData()
        formData.append('name', this.newProduct.name)
        formData.append('description', this.newProduct.description)
        formData.append('picture', this.newProduct.picture)
        formData.append('price', this.newProduct.price)
        formData.append('stock', this.newProduct.stock)

        const response = await fetch('http://localhost:8080/products', {
          method: 'POST',
          body: formData
        })
        if (response.ok) {
          await this.listProducts()
          this.resetNewProduct()
          Modal.getInstance('#addProductModal').hide()
        }
      } catch (error) {
        console.error('添加商品失败:', error)
      }
    },
    async updateProduct() {
      try {
        const formData = new FormData()
        formData.append('name', this.editProduct.name)
        formData.append('description', this.editProduct.description)
        formData.append('picture', this.editProduct.picture)
        formData.append('price', this.editProduct.price)
        formData.append('stock', this.editProduct.stock)

        const response = await fetch(`http://localhost:8080/products/${this.editProduct.id}`, {
          method: 'PUT',
          body: formData
        })
        if (response.ok) {
          await this.listProducts()
          Modal.getInstance('#editProductModal').hide()
        }
      } catch (error) {
        console.error('更新商品失败:', error)
      }
    },
    async deleteProduct(id) {
      try {
        const response = await fetch(`http://localhost:8080/products/${id}`, {
          method: 'DELETE'
        })
        if (response.ok) {
          await this.listProducts()
        }
      } catch (error) {
        console.error('删除商品失败:', error)
      }
    },
    async filterByCategory(categoryId) {
      try {
        const response = await fetch(`http://localhost:8080/category/${categoryId}`);
        if (response.ok) {
          this.products = await response.json();
        }
      } catch (error) {
        console.error('获取分类商品失败:', error);
      }
    },
    async listCategories() {
      try {
        const response = await fetch('http://localhost:8080/categories');
        if (response.ok) {
          this.categories = await response.json();
        }
      } catch (error) {
        console.error('获取分类列表失败:', error);
      }
    },
    resetNewProduct() {
      this.newProduct = {
        name: '',
        description: '',
        picture: null,
        price: 0,
        stock: 0
      }
    },
    onFileChange(event, type) {
      const file = event.target.files[0]
      if (type === 'new') {
        this.newProduct.picture = file
      } else {
        this.editProduct.picture = file
      }
    }
  }
}
</script>

<style scoped>
.card {
  margin-bottom: 20px;
}

.table {
  margin-top: 15px;
}
</style>