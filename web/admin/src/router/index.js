import Vue from 'vue'
import VueRouter from 'vue-router'

const Login = () => import(/* webpackChunkName: "Login" */ '../views/Login.vue')
const Admin = () => import(/* webpackChunkName: "Admin" */ '../views/Admin.vue')

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const Users = () => import(/* webpackChunkName: "Users" */ '../components/user/Users.vue')

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'viaBlog 后台管理页面'
        }
      }, {
        path: 'users',
        component: Users,
        meta: {
          title: '用户列表'
        }
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

// router 前置守卫
router.beforeEach((to, from, next) => {
  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})

export default router
