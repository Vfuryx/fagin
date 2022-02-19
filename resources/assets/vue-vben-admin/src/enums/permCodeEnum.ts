export enum PermCodeEnum {
  // 账号管理
  AdminSystemAccountQuery = 'admin.system.account:query',
  AdminSystemAccountDetail = 'admin.system.account:detail',
  AdminSystemAccountCreate = 'admin.system.account:create',
  AdminSystemAccountUpdate = 'admin.system.account:update',
  AdminSystemAccountDelete = 'admin.system.account:delete',
  AdminSystemAccountUpdatePWD = 'admin.system.account:update-pwd',
  AdminSystemAccountUpdateStatus = 'admin.system.account:update-status',
  AdminSystemAccountLogout = 'admin.system.account:logout',
  AdminSystemAccountExists = 'admin.system.account:exists',
  AdminSystemAccountRoles = 'admin.system.account:roles',

  // 菜单管理
  AdminSystemMenuQuery = 'admin.system.menu:query',
  AdminSystemMenuDetail = 'admin.system.menu:detail',
  AdminSystemMenuCreate = 'admin.system.menu:create',
  AdminSystemMenuUpdate = 'admin.system.menu:update',
  AdminSystemMenuDelete = 'admin.system.menu:delete',
  AdminSystemMenuAll = 'admin.system.menu:all',

  // 角色管理
  AdminSystemRoleQuery = 'admin.system.role:query',
  AdminSystemRoleDetail = 'admin.system.role:detail',
  AdminSystemRoleCreate = 'admin.system.role:create',
  AdminSystemRoleUpdate = 'admin.system.role:update',
  AdminSystemRoleDelete = 'admin.system.role:delete',
  AdminSystemRoleKeyExists = 'admin.system.role:key-exists',
  AdminSystemRoleUpdateStatus = 'admin.system.role:update-status',

  // 部门管理
  AdminSystemDepartmentQuery = 'admin.system.department:query',
  AdminSystemDepartmentCreate = 'admin.system.department:create',
  AdminSystemDepartmentUpdate = 'admin.system.department:update',
  AdminSystemDepartmentDelete = 'admin.system.department:delete',

  // 操作日志
  AdminMonitorOperLogQuery = 'admin.monitor.oper-log:query',
  AdminMonitorOperLogDetail = 'admin.monitor.oper-log:detail',

  // 登录日志
  AdminMonitorLoginLogQuery = 'admin.monitor.login-log:query',
  AdminMonitorLoginLogDetail = 'admin.monitor.login-log:detail',
}
