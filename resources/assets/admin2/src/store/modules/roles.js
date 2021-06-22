import {
  getRoleList,
  showRole,
  addRole,
  updateRole,
  delRole,
  roleKeyExist
} from '@/api/roles'

const roles = {
  state: {
  },
  mutations: {
  },
  actions: {
    // 获取列表
    getRoleList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getRoleList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增
    addRole ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addRole(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showRole ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showRole(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改
    updateRole ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updateRole(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delRole ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delRole(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    roleKeyExist ({ commit }, params) {
      return new Promise((resolve, reject) => {
        roleKeyExist(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default roles
