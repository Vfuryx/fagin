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
