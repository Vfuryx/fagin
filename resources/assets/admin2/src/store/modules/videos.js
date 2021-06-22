import {
  getVideoList,
  showVideo,
  addVideo,
  updateVideo,
  delVideo, uploadVideo
} from '@/api/videos'

const admins = {
  state: {},
  mutations: {},
  actions: {
    // 获取列表
    getVideoList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getVideoList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 新增
    addVideo ({ commit }, data) {
      return new Promise((resolve, reject) => {
        addVideo(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 展示详情
    showVideo ({ commit }, id) {
      return new Promise((resolve, reject) => {
        showVideo(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 修改
    updateVideo ({ commit }, data) {
      return new Promise((resolve, reject) => {
        updateVideo(data.id, data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 删除
    delVideo ({ commit }, id) {
      return new Promise((resolve, reject) => {
        delVideo(id).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 上传
    uploadVideo ({ commit }, data) {
      return new Promise((resolve, reject) => {
        uploadVideo(data).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default admins
