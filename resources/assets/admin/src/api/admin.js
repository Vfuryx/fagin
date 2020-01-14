import request from '@/utils/request'

/**
 * 获取数据
 */
export function show() {
  return request({
    url: '/v1/user/info',
    method: 'get'
  })
}

/**
 * 更新数据
 * @param {*} id  ID
 */
export function update(id, data) {
  return request({
    url: '/v1/user/' + id,
    method: 'put',
    data
  })
}

/**
 * 更新网站设置
 * @param {*} data  data
 */
export function getInfo() {
  return request({
    url: '/v1/website/info',
    method: 'get'
  })
}

/**
 * 更新数据
 * @param {*} id  ID
 */
export function updateInfo(data) {
  return request({
    url: '/v1/website/info',
    method: 'put',
    data
  })
}

/**
 * 更新公司简介
 * @param {*} data  data
 */
export function getCompanyIntroduction() {
  return request({
    url: '/v1/company/introduction',
    method: 'get'
  })
}

/**
 * 更新数据
 * @param {*} id  ID
 */
export function updateCompanyIntroduction(data) {
  return request({
    url: '/v1/company/introduction',
    method: 'put',
    data
  })
}

/**
 * 更新数据
 * @param {*} id  ID
 */
export function uploadCompanyImage(data) {
  return request({
    url: '/v1/company/upload',
    method: 'post',
    data
  })
}