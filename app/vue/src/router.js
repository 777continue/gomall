import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '@/components/SignUp.vue'
import SignIn from '@/components/SignIn.vue'
import Home from '@/components/Home.vue'
import ManageUser from '@/components/ManageUser.vue'
import Product from '@/components/Product.vue'
import Category from './components/Category.vue'
import ManageProduct from './components/ManageProduct.vue'
import Search from './components/Search.vue'
import Order from './components/Order.vue'
import Cart from './components/Cart.vue'
import CheckOut from './components/CheckOut.vue'
import Waiting from './components/Waiting.vue'

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
            path: '/search',
            component: Search,
        },
        {
            path: '/order',
            component: Order
        },
        {
            path: '/cart',
            component: Cart
        },
        {
            path: '/checkout',
            component: CheckOut
        },
        {
            path: '/waiting',
            component: Waiting
        },
        {
            path: '/:pathMatch(.*)*',
            redirect: '/'
        }
    ]
})

export default router