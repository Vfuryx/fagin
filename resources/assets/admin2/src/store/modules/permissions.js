
import {
  getPermissionGroupList,
  showPermissionGroup,
  addPermissionGroup,
  updatePermissionGroup,
  delPermissionGroup,
  getPermissionGroupAll,
  getPermissionList,
  showPermission,
  addPermission,
  updatePermission,
  delPermission,
  getGroupPermission
} from '@/api/permissions'

const permissions = {
  state: {
    permissions: []
  },
  mutations: {
    SET_MENUS: (state, permissions) => {
      state.permissions = permissions
    }
  },
  actions: {
    // 获取列表
    getPermissionGroupList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getPermissionGroupList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增
    addPermissionGroup ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addPermissionGroup(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showPermissionGroup ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showPermissionGroup(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改
    updatePermissionGroup ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updatePermissionGroup(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delPermissionGroup ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delPermissionGroup(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 获取列表
    getPermissionGroupAll ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getPermissionGroupAll(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 获取列表
    getPermissionList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getPermissionList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增
    addPermission ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addPermission(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showPermission ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showPermission(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改
    updatePermission ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updatePermission(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delPermission ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delPermission(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 分组权限
    getGroupPermission ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getGroupPermission(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default permissions
