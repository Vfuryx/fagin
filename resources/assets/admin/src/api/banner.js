import request from '@/utils/request'

export function list(query) {
  return request({
    url: '/v1/banner/',
    method: 'get',
    params: query
  })
}

export function create(data) {
  return request({
    url: '/v1/banner/',
    method: 'post',
    data
  })
}

export function update(id, data) {
  return request({
    url: '/v1/banner/' + id,
    method: 'put',
    data
  })
}

export function show(id) {
  return request({
    url: '/v1/banner/' + id,
    method: 'get'
  })
}

export function del(id) {
  return request({
    url: '/v1/banner/' + id,
    method: 'delete'
  })
}

// 删除banner组
export function delBanners(ids) {
  return request({
    url: '/v1/banner/',
    method: 'delete',
    data: {
      ids: ids
    }
  })
}
