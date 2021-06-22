import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/v1/articles/',
    method: 'get',
    params: query
  })
}

export function create(data) {
  return request({
    url: '/v1/articles/',
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: '/v1/articles/' + id,
    method: 'put',
    data
  })
}

export function show(id) {
  return request({
    url: '/v1/articles/' + id,
    method: 'get'
  })
}

export function del(id) {
  return request({
    url: '/v1/articles/' + id,
    method: 'delete'
  })
}

// 删除banner组
export function deletes(ids) {
  return request({
    url: '/v1/articles/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}
