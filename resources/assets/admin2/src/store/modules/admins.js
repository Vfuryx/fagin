import {
  getAdminList,
  showAdmin,
  addAdmin,
  updateAdmin,
  delAdmin,
  logoutAdmin,
  adminUsernameExist, adminResetPassword
} from '@/api/admins'
import {
  getRoleAll
} from '@/api/roles'

const admins = {
  state: {},
  mutations: {},
  actions: {
    // 获取列表
    getAdminList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getAdminList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增
    addAdmin ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addAdmin(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showAdmin ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showAdmin(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改
    updateAdmin ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updateAdmin(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delAdmin ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delAdmin(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    getRoleAll ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getRoleAll(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 强制用户下线
    logoutAdmin ({ commit }, id) {
      return new Promise((resolve, reject) => {
        logoutAdmin(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    adminUsernameExist ({ commit }, params) {
      return new Promise((resolve, reject) => {
        adminUsernameExist(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 重置密码
    adminResetPassword ({ commit }, data) {
      return new Promise((resolve, reject) => {
        adminResetPassword(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default admins
