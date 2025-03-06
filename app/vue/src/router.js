import { createRouter, createWebHistory } from 'vue-router'
import SignUp from '@/components/SignUp.vue'
import SignIn from '@/components/SignIn.vue'
import Home from '@/components/Home.vue'
import ManageUser from '@/components/ManageUser.vue'

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
            path: '/:pathMatch(.*)*',
            redirect: '/'
        }
    ]
})

export default router