import { createRouter, createWebHashHistory } from 'vue-router'

import Login from '../pages/Login.vue'
import Home from '../pages/Home.vue'
import About from '../pages/About.vue'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/', name: 'Login', component: Login },
        { path: '/home', name: 'Home', component: Home },
        { path: '/about', name: 'About', component: About },
        {
            path: '/:catchAll(.*)',
            redirect: '/',
        },
        {
            path: '/about',
            name: 'about',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('../pages/About.vue'),
        },
    ],
})

export default router
