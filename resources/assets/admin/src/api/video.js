import request from '@/utils/request'

export function getList(query) {
  return request({
    url: '/v1/video/list',
    method: 'get',
    params: query
  })
}

export function createVideo(data) {
  return request({
    url: '/v1/video/',
    method: 'post',
    data
  })
}

export function updateVideo(id, data) {
  return request({
    url: '/v1/video/' + id,
    method: 'put',
    data
  })
}

export function deleteVideo(id) {
  return request({
    url: '/v1/video/' + id,
    method: 'delete'
  })
}

// 删除角色组
export function deleteVideos(ids) {
  return request({
    url: '/v1/video/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}
