import Vue from 'vue'
import Vuex from 'vuex'

import app from './modules/app'
import user from './modules/user'
import menu from './modules/menu'
import permissions from './modules/permissions'
import roles from './modules/roles'
import admins from './modules/admins'
import operationLog from './modules/operation_log'
import videos from './modules/videos'
import asyncRouter from './modules/async-router'

// default router permission control
// import permission from './modules/permission'

// dynamic router permission control (Experimental)
// import permission from './modules/async-router'
import getters from './getters'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    app,
    user,
    menu,
    permissions,
    roles,
    admins,
    operationLog,
    videos,
    asyncRouter // 改为后端获取菜单
  },
  state: {

  },
  mutations: {

  },
  actions: {

  },
  getters
})
