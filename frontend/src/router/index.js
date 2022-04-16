import { createRouter, createWebHistory } from 'vue-router'
import Password from "../views/Password.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // This path is secret
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: Password
    }
  ]
})

export default router
