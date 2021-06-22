import request from '@/utils/request'

// 查询列表
export function list(query) {
  return request({
    url: '/v1/permissions/',
    method: 'get',
    params: query
  })
}

// 查询详细
export function show(id) {
  return request({
    url: '/v1/permissions/' + id,
    method: 'get'
  })
}

// 新增角色
export function add(data) {
  return request({
    url: '/v1/permissions/',
    method: 'post',
    data: data
  })
}

// 修改角色
export function update(data) {
  return request({
    url: '/v1/permissions/' + data.id,
    method: 'put',
    data: data
  })
}

// 查询列表
export function gList(query) {
  return request({
    url: '/v1/permission/groups/',
    method: 'get',
    params: query
  })
}

// 查询列表
export function gAll(query) {
  return request({
    url: '/v1/permission/groups/all',
    method: 'post',
    params: query
  })
}

// 查询详细
export function gShow(id) {
  return request({
    url: '/v1/permission/groups/' + id,
    method: 'get'
  })
}

// 新增角色
export function gAdd(data) {
  return request({
    url: '/v1/permission/groups/',
    method: 'post',
    data: data
  })
}

// 修改角色
export function gUpdate(data) {
  return request({
    url: '/v1/permission/groups/' + data.id,
    method: 'put',
    data: data
  })
}

// 分组权限
export function getGroupPermission() {
  return request({
    url: '/v1/group_permissions',
    method: 'get'
  })
}

