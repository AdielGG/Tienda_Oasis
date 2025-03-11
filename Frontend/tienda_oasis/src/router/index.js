import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/login/LoginView.vue'
import AboutView from '@/views/AboutView.vue'
import ProductsView from '@/views/products/ProductsView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,

    },
    {
      path: '/products',
      name: 'products',
      component: ProductsView

    }
    
  ],
})

export default router
