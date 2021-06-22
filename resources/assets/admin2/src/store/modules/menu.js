
import { getMenuList, addMenu, showMenu, updateMenu, delMenu, getMenuAll } from '@/api/menu'

const menu = {
  state: {
    menus: []
  },
  mutations: {
    SET_MENUS: (state, menus) => {
      state.menus = menus
    }
  },

  actions: {
    // 获取菜单列表
    getMenuList ({ commit }, type) {
      return new Promise((resolve, reject) => {
        getMenuList({ 'type': type }).then(response => {
          // commit('SET_TOKEN', result)
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增菜单
    addMenu ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addMenu(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示菜单详情
    showMenu ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showMenu(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改菜单
    updateMenu ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updateMenu(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除菜单
    delMenu ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delMenu(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 获取列表
    getMenuAll ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getMenuAll(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default menu
