import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '@/components/SignUp.vue'
import SignIn from '@/components/SignIn.vue'
import Home from '@/components/Home.vue'
import ManageUser from '@/components/ManageUser.vue'
import Product from '@/components/Product.vue'
import Category from './components/Category.vue'
import ManageProduct from './components/ManageProduct.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/sign-up',
            component: SignUp
        },
        {
            path: '/sign-in',
            component: SignIn
        },
        {
            path: '/manage-user',
            component: ManageUser
        },
        {
            path: '/manage-product',
            component: ManageProduct
        },
        {
            path: '/product',
            component: Product,
        },
        {
            path: '/category/:type',
            component: Category,
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/'
        }
    ]
})

export default router