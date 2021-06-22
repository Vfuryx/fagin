import request from '@/utils/request'

export function listAuthor(query) {
  return request({
    url: '/v1/authors/',
    method: 'get',
    params: query
  })
}

export function create(data) {
  return request({
    url: '/v1/authors/',
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: '/v1/authors/' + id,
    method: 'put',
    data
  })
}

export function showAuthor(id) {
  return request({
    url: '/v1/authors/' + id,
    method: 'get'
  })
}

export function del(id) {
  return request({
    url: '/v1/authors/' + id,
    method: 'delete'
  })
}

// 删除banner组
export function deletes(ids) {
  return request({
    url: '/v1/authors/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}


export function all() {
  return request({
    url: '/v1/authors/all',
    method: 'post'
  })
}
