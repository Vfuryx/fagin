const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
  errorLogs: state => state.errorLog.logs,
  permission_routes: state => state.permission.routes // 动态路由
}
export default getters
