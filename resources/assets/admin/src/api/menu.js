import request from '@/utils/request'

// 获取菜单列表
export function listMenu(query) {
  return request({
    url: '/v1/menu/',
    method: 'get',
    params: query
  })
}

// 获取单个菜单
export function getMenu(id) {
  return request({
    url: '/v1/menu/' + id,
    method: 'get'
  })
}

// 添加菜单
export function addMenu(data) {
  return request({
    url: '/v1/menu/',
    method: 'post',
    data
  })
}

// 修改菜单
export function updateMenu(id, data) {
  return request({
    url: '/v1/menu/' + id,
    method: 'put',
    data
  })
}

// 删除菜单
export function delMenu(id) {
  return request({
    url: '/v1/menu/' + id,
    method: 'delete'
  })
}
