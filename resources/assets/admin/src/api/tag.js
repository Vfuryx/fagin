import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/v1/tags/',
    method: 'get',
    params: query
  })
}

export function create(data) {
  return request({
    url: '/v1/tags/',
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: '/v1/tags/' + id,
    method: 'put',
    data
  })
}

export function show(id) {
  return request({
    url: '/v1/tags/' + id,
    method: 'get'
  })
}

export function del(id) {
  return request({
    url: '/v1/tags/' + id,
    method: 'delete'
  })
}

// 删除banner组
export function deletes(ids) {
  return request({
    url: '/v1/tags/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}

// 所有
export function all(data) {
  return request({
    url: '/v1/tags/all',
    method: 'post',
  })
}
