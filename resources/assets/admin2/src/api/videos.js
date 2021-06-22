import request from '@/utils/request'

export const videoApi = {
  VideoList: '/v1/videos/',
  AddVideo: '/v1/videos/',
  UpdateVideo: '/v1/videos/',
  ShowVideo: '/v1/videos/',
  DelVideo: '/v1/videos/',
  UploadVideo: '/v1/videos/upload'
}

/**
 * 获取列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getVideoList (parameter) {
  return request({
    url: videoApi.VideoList,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取所有角色
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getRoleAll (parameter) {
  return request({
    url: videoApi.RoleAll,
    method: 'post',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 展示详情
 * @param id ID
 * @returns {*}
 */
export function showVideo (id) {
  return request({
    url: videoApi.ShowVideo + id,
    method: 'get'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 新增
 * data: {
 *
 * }
 * @param data
 * @returns {*}
 */
export function addVideo (data) {
  return request({
    url: videoApi.AddVideo,
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 修改
 * @param id
 * @param data
 * @returns {*}
 */
export function updateVideo (id, data) {
  return request({
    url: videoApi.UpdateVideo + id,
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 删除
 * @param id ID
 * @returns {*}
 */
export function delVideo (id) {
  return request({
    url: videoApi.DelVideo + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取所有角色
 * data: {
 * }
 * @param data
 * @returns {*}
 */
export function uploadVideo (data) {
  return request({
    url: videoApi.UploadVideo,
    method: 'post',
    data: data,
    headers: {
      'Content-Type': 'multipart/form-data;charset=UTF-8'
    }
  })
}
