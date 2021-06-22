import {
  getOperationLogList,
  showOperationLog,
  delOperationLog
} from '@/api/operation_log'

const operationLog = {
  state: {},
  mutations: {},
  actions: {
    // 获取列表
    getOperationLogList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getOperationLogList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showOperationLog ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showOperationLog(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delOperationLog ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delOperationLog(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default operationLog
