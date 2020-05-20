import Vue from 'vue' // 导入vue插件
import VueRouter from 'vue-router' // 导入vue-router
import Layout from '@/layout/Layout' // 导入Layout组件
import Home from '../views/Home.vue'

// import Users from '../views/company/users'
// import Checked from '../views/company/checked'
// import Unchecked from '../views/company/unchecked'
import Company from '../views/company/index'
import Lists from '../views/company/lists'
import Updates from '../views/company/updates'
import Checklist from '../views/company/checklist'

import Cms from '../views/cms/index'
import Stores from '../views/cms/stores'
import Restore from '../views/cms/restore'
import Reproduces from '../views/cms/reproduces'
import Create from '../views/cms/create'

//Vue.use(plugin, arguments)就是执行一个plugin函数，或执行plugin的install方法进行插件注册(如果plugin是一个函数，
// 则执行；若是一个插件，则执行plugin的install方法...)；并向plugin或其install方法传入Vue对象作为第一个参数；如果有多个参数，use的其它参数作为plugin或install的其它参数。
Vue.use(VueRouter)

const routes = [
  {
    path: '',
    component: Layout,
    children: [                                         //二级路由
      {
        path: '/',// 匹配路由的根路径
        name: 'home',
        component:Home
      },
      {
        path: '/cms',// 匹配路由的根路径
        name: 'cms',
        component: Cms,
        children:[
            {
              name:'reproduces',
              path:'reproduces',
              component:Reproduces
            },
          {
            name:'restore',
            path:'restore',
            component:Restore
          },
          {
            name:'stores',
            path:'stores',
            component:Stores
          },
          {
            name:'create',
            path:'create',
            component:Create
          }
        ]
      },
      {
        path: '/company',// 匹配路由的根路径
        name: 'company',
        component:Company,
        children: [
          {
            name: 'updates',
            path: 'updates',
            component: Updates
          },
          {
            name: 'lists',
            path: 'lists',
            component: Lists
          },
          {
            name:'checklist',
            path:'checklist',
            component:Checklist
          },
        ]
      },
      {
        path: '/goods',
        name: 'goods',
        component: () => import('../views/goods')
      },
      {
        path: '/goods/property',
        name: 'property',
        component: () => import('../views/goods/property')
      }
    ]
  }
]
const router = new VueRouter({
  routes
})
export default router
