import Vue from 'vue'
import VueRouter from 'vue-router'

const Login = () => import(/* webpackChunkName: "Login" */ '../views/Login.vue')
const Admin = () => import(/* webpackChunkName: "Admin" */ '../views/Admin.vue')

// 页面路由组件
const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const Users = () => import(/* webpackChunkName: "Users" */ '../components/user/Users.vue')
const Articles = () => import(/* webpackChunkName: "Articles" */ '../components/article/Articles.vue')
const Categories = () => import(/* webpackChunkName: "Categories" */ '../components/category/Categories.vue')
const AddArt = () => import(/* webpackChunkName: "AddArt" */ '../components/article/AddArt.vue')
const Profile = () => import(/* webpackChunkName: "UserList" */ '../components/user/Profile.vue')
const Comments = () => import(/* webpackChunkName: "UserList" */ '../components/comment/Comments.vue')

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
      }, {
        path: 'articles',
        component: Articles,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'addart',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'categories',
        component: Categories,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'profile',
        component: Profile,
        meta: {
          title: '个人设置'
        }
      },
      {
        path: 'comments',
        component: Comments,
        meta: {
          title: '评论管理'
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
