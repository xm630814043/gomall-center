import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/layout/Layout'

Vue.use(VueRouter)

const routes = [
  {
    path: '',
    component: Layout,
    children: [
      {
        path: '/',
        name: 'home',
        component: () => import('../views/Home.vue')
      },
      {
        path: '/cms',
        name: 'cms',
        component: () => import('../views/cms')
      },
      {
        path: '/company',
        name: 'company',
        component: () => import('../views/company')
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
