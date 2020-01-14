import request from '@/utils/request'

/**
 * 获取列表
 */
export function logList(query) {
  return request({
    url: '/v1/operation/logs/',
    method: 'get',
    params: query
  })
}

/**
 * 获取详情
 */
export function show(id) {
  return request({
    url: '/v1/operation/logs/' + id,
    method: 'get'
  })
}
