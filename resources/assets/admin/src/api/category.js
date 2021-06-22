import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/v1/categories/',
    method: 'get',
    params: query
  })
}

export function create(data) {
  return request({
    url: '/v1/categories/',
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: '/v1/categories/' + id,
    method: 'put',
    data
  })
}

export function show(id) {
  return request({
    url: '/v1/categories/' + id,
    method: 'get'
  })
}

export function del(id) {
  return request({
    url: '/v1/categories/' + id,
    method: 'delete'
  })
}

// 删除banner组
export function deletes(ids) {
  return request({
    url: '/v1/categories/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}

export function all() {
  return request({
    url: '/v1/categories/all',
    method: 'post'
  })
}
