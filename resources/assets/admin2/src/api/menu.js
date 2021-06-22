import request from '@/utils/request'

const menuApi = {
  MenuAll: '/v1/menu/all',
  MenuList: '/v1/menu/',
  AddMenu: '/v1/menu/',
  UpdateMenu: '/v1/menu/',
  ShowMenu: '/v1/menu/',
  DelMenu: '/v1/menu/'
}

/**
 * 获取菜单列表
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getMenuList (parameter) {
  return request({
    url: menuApi.MenuList,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取菜单列表
 * @param id ID
 * @returns {*}
 */
export function showMenu (id) {
  return request({
    url: menuApi.ShowMenu + id,
    method: 'get'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 新增菜单
 * parameter: {
 *
 * }
 * @param parameter
 * @returns {*}
 */
export function addMenu (parameter) {
  return request({
    url: menuApi.AddMenu,
    method: 'post',
    data: parameter,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 修改菜单
 * @param id
 * @param data
 * @returns {*}
 */
export function updateMenu (id, data) {
  return request({
    url: menuApi.UpdateMenu + id,
    method: 'put',
    data: data,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8'
    }
  })
}

/**
 * 删除菜单
 * @param id ID
 * @returns {*}
 */
export function delMenu (id) {
  return request({
    url: menuApi.DelMenu + id,
    method: 'delete'
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}

/**
 * 获取所有
 * parameter: {
 *     type: 0,
 * }
 * @param parameter
 * @returns {*}
 */
export function getMenuAll (parameter) {
  return request({
    url: menuApi.MenuAll,
    method: 'get',
    params: parameter
    // headers: {
    //   'Content-Type': 'application/json;charset=UTF-8'
    // }
  })
}
