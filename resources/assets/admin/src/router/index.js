import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path*',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '主页', icon: 'dashboard', affix: true }
    }]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/group',
    component: Layout,
    redirect: '/group',
    hidden: false,
    children: [{
      path: 'index',
      name: 'PermissionGroups',
      component: () => import('@/views/admin/permission_groups'),
      meta: { title: '视频', icon: 'eye-open' }
    }]
  }
  // {
  //   path: '/banner',
  //   component: Layout,
  //   hidden: false,
  //   redirect: '/banner',
  //   children: [
  //     {
  //       path: 'index',
  //       name: 'BannerIndex',
  //       component: () => import('@/views/banner/index'),
  //       meta: { title: '轮播图管理', icon: 'banner', roles: ['admin'] }
  //     }
  //   ]
  // },
  // {
  //   path: '/about',
  //   component: Layout,
  //   hidden: false,
  //   redirect: '/banner',
  //   meta: {
  //     title: '关于我们',
  //     icon: 'introduction'
  //   },
  //   children: [
  //     {
  //       path: 'company',
  //       name: 'CompanyIntroduction',
  //       component: () => import('@/views/about/company_introduction'),
  //       meta: { title: '公司介绍', icon: 'introduction', roles: ['admin'] }
  //     }
  //   ]
  // },
  // {
  //   path: '/admin',
  //   component: Layout,
  //   hidden: false,
  //   redirect: '/admin',
  //   name: 'Admin',
  //   meta: {
  //     title: '系统设置',
  //     icon: 'setting'
  //   },
  //   children: [
  //     {
  //       path: 'user',
  //       name: 'User',
  //       component: () => import('@/views/admin/user'),
  //       meta: { title: '用户管理', icon: 'user', roles: ['admin'] }
  //     },
  //     {
  //       path: 'menu',
  //       name: 'Menu',
  //       component: () => import('@/views/menu/index'),
  //       meta: { title: '菜单管理', icon: 'tree-table', roles: ['admin', 'system:sysmenu:*'] }
  //     },
  //     {
  //       path: 'role',
  //       name: 'Role',
  //       component: () => import('@/views/role/index'),
  //       meta: { title: '角色管理', icon: 'peoples', roles: ['admin', 'system:sysmenu:*'] }
  //     },
  //     {
  //       path: 'login_log',
  //       name: 'LoginLog',
  //       component: () => import('@/views/admin/login_log'),
  //       meta: { title: '登录日志', icon: 'log', roles: ['admin'] }
  //     },
  //     {
  //       path: 'operation_log',
  //       name: 'OperationLog',
  //       component: () => import('@/views/admin/operation_log'),
  //       meta: { title: '操作日志', icon: 'log', roles: ['admin'] }
  //     },
  //     {
  //       path: 'operlog',
  //       name: 'operlog',
  //       component: () => import('@/views/operlog/index'),
  //       meta: { title: '操作日志2', icon: 'log', roles: ['admin'] }
  //     },
  //     {
  //       path: 'website',
  //       name: 'WebsitesSetup',
  //       component: () => import('@/views/admin/website'),
  //       meta: { title: '网站设置', icon: 'component', roles: ['admin'] }
  //     }
  //   ]
  // }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
