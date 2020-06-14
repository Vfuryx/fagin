import request from '@/utils/request'

// 查询角色列表
export function listRole(query) {
  return request({
    url: '/v1/role/',
    method: 'get',
    params: query
  })
}

// 查询角色详细
export function getRole(roleId) {
  return request({
    url: '/v1/role/' + roleId,
    method: 'get'
  })
}

// 新增角色
export function addRole(data) {
  return request({
    url: '/v1/role/',
    method: 'post',
    data: data
  })
}

// 修改角色
export function updateRole(data) {
  return request({
    url: '/v1/role/' + data.id,
    method: 'put',
    data: data
  })
}

// 角色数据权限
export function dataScope(data) {
  return request({
    url: '/v1/roledatascope',
    method: 'put',
    data: data
  })
}

// 角色状态修改
export function changeRoleStatus(id, status) {
  const data = {
    status
  }
  return request({
    url: '/v1/role/' + id + '/status/',
    method: 'put',
    data: data
  })
}

// 删除角色
export function delRole(roleId) {
  return request({
    url: '/v1/role/' + roleId,
    method: 'delete'
  })
}
// 删除角色组
export function delRoles(roleId) {
  return request({
    url: '/v1/role/',
    method: 'delete',
    data: {
      ids: roleId
    }
  })
}

export function getListrole(id) {
  return request({
    url: '/v1/menu/role/' + id,
    method: 'get'
  })
}

export function getRoutes() {
  return request({
    url: '/v1/us/menurole',
    method: 'get'
  })
}

export function getMenuNames() {
  return request({
    url: '/v1/menuids',
    method: 'get'
  })
}
