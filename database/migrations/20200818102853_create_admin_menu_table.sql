-- +migrate Up
DROP TABLE IF EXISTS `admin_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_menus`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `parent_id`  int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '父菜单id',
    `paths`      varchar(128)        NOT NULL DEFAULT '' COMMENT '菜单层级路径 0-1-2',
    `name`       varchar(32)         NOT NULL DEFAULT '' COMMENT '菜单名称',
    `title`      varchar(32)         NOT NULL DEFAULT '' COMMENT '菜单标题',
    `icon`       varchar(128)        NOT NULL DEFAULT '' COMMENT '图标',
    `type`       tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:目录 1:菜单 2:接口 3:按钮',
    `path`       varchar(128)        NOT NULL DEFAULT '' COMMENT '路由地址',
    `component`  varchar(64)         NOT NULL DEFAULT '' COMMENT '组件路径',
    `method`     varchar(16)         NOT NULL DEFAULT '' COMMENT '菜单请求方法',
    `permission` varchar(32)         NOT NULL DEFAULT '' COMMENT '权限',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `visible`    tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>关闭 1=>开启',
    `is_link`    tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否外链',
    PRIMARY KEY (`id`),
    KEY `idx_admin_menus_deleted_at` (`deleted_at`),
    KEY `idx_admin_menus_parent_id` (`parent_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台菜单表';

INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `name`, `title`, `icon`, `type`, `path`, `component`, `method`, `permission`, `sort`, `visible`, `is_link`) VALUES
(1, NULL, NULL, NULL, 0, '0-1', 'Admin', '系统管理', 'example', 0, '/admin', 'Layout', '', '', 1, 1, 0),
(2, NULL, NULL, NULL, 1, '0-1-2', 'User', '用户管理', 'user', 1, 'user', '/admin/user', '', '', 0, 1, 0),
(3, NULL, NULL, NULL, 1, '0-1-3', 'Menu', '菜单管理', 'nested', 1, 'menu', '/menu/index', '', '', 0, 1, 0),
(4, NULL, NULL, NULL, 1, '0-1-4', 'Role', '角色管理', 'peoples', 1, 'role', '/role/index', '', '', 0, 1, 0),
(5, NULL, NULL, NULL, 1, '0-1-5', 'WebsitesSetup', '网站设置', 'component', 1, 'website', '/admin/website', '', '', 0, 1, 0),
(6, NULL, NULL, NULL, 0, '0-6', 'API管理', 'API管理', 'example', 0, '/', '/', '', '', 0, 0, 0),
(7, NULL, NULL, NULL, 6, '0-6-7', 'V1', 'V1', 'example', 0, '/', '/', '', '', 0, 0, 0),
(8, NULL, NULL, NULL, 7, '0-6-7-8', '系统管理API', '系统管理API', 'example', 0, '/', '/', '', '', 0, 0, 0),
(9, NULL, NULL, NULL, 7, '0-6-7-9', '权限API', '权限API', 'example', 0, '/', '/', '', '', 0, 0, 0),
(10, NULL, NULL, NULL, 9, '0-6-7-9-10', '用户详情', '用户详情', 'example', 2, '/admin/api/v1/us/info', '', 'GET', '', 0, 0, 0),
(11, NULL, NULL, NULL, 9, '0-6-7-9-11', '用户菜单列表', '用户菜单列表', 'example', 2, '/admin/api/v1/us/menurole', '', 'GET', '', 0, 0, 0),
(12, NULL, NULL, NULL, 9, '0-6-7-9-12', '前端路由', '前端路由', 'example', 2, '/admin/api/v1/us/routes', '', 'GET', '', 0, 0, 0),
(13, NULL, NULL, NULL, 9, '0-6-7-9-13', '', '用户退出', 'example', 2, '/admin/api/v1/user/logout', '', 'POST', '', 0, 0, 0),
(14, NULL, NULL, NULL, 8, '0-6-7-8-14', '', '用户管理', 'example', 0, '/', '/', '', '', 0, 0, 0),
(15, NULL, NULL, NULL, 14, '0-6-7-8-14-15', '', '用户列表', 'example', 2, '/admin/api/v1/user/', '', 'GET', '', 0, 0, 0),
(16, NULL, NULL, NULL, 14, '0-6-7-8-14-16', '', '展示用户', 'example', 2, '/admin/api/v1/user/:id', '', 'GET', '', 0, 0, 0),
(17, NULL, NULL, NULL, 14, '0-6-7-8-14-17', '', '新增用户', 'example', 2, '/admin/api/v1/user/', '', 'POST', '', 0, 0, 0),
(18, NULL, NULL, NULL, 14, '0-6-7-8-14-18', '', '更新用户', 'example', 2, '/admin/api/v1/user/:id', '', 'PUT', '', 0, 0, 0),
(19, NULL, NULL, NULL, 14, '0-6-7-8-14-19', '', '更新用户状态', 'example', 2, '/admin/api/v1/user/:id/status/', '', 'PUT', '', 0, 0, 0),
(20, NULL, NULL, NULL, 14, '0-6-7-8-14-20', '', '用户重置密码', 'example', 2, '/admin/api/v1/user/:id/reset/', '', 'PUT', '', 0, 0, 0),
(21, NULL, NULL, NULL, 14, '0-6-7-8-14-21', '', '用户删除', 'example', 2, '/admin/api/v1/user/:id', '', 'DELETE', '', 0, 0, 0),
(22, NULL, NULL, NULL, 14, '0-6-7-8-14-22', '', '用户批量删除', 'example', 2, '/admin/api/v1/user/', '', 'DELETE', '', 0, 0, 0),
(23, NULL, NULL, NULL, 8, '0-6-7-8-23', '', '菜单管理', 'example', 0, '/', '/', '', '', 0, 0, 0),
(24, NULL, NULL, NULL, 23, '0-6-7-8-23-24', '', '菜单列表', 'example', 2, '/admin/api/v1/menu/', '', 'GET', '', 0, 0, 0),
(25, NULL, NULL, NULL, 23, '0-6-7-8-23-25', '', '菜单详情', 'example', 2, '/admin/api/v1/menu/:id', '', 'GET', '', 0, 0, 0),
(26, NULL, NULL, NULL, 23, '0-6-7-8-23-26', '', '菜单新增', 'example', 2, '/admin/api/v1/menu/', '', 'POST', '', 0, 0, 0),
(27, NULL, NULL, NULL, 23, '0-6-7-8-23-27', '', '菜单更新', 'example', 2, '/admin/api/v1/menu/:id', '', 'PUT', '', 0, 0, 0),
(28, NULL, NULL, NULL, 23, '0-6-7-8-23-28', '', '菜单删除', 'example', 2, '/admin/api/v1/menu/:id', '', 'DELETE', '', 0, 0, 0),
(29, NULL, NULL, NULL, 8, '0-6-7-8-29', '', '角色管理', 'example', 0, '/', '/', '', '', 0, 0, 0),
(30, NULL, NULL, NULL, 29, '0-6-7-8-29-30', '', '角色列表', 'example', 2, '/admin/api/v1/role/', '', 'GET', '', 0, 0, 0),
(31, NULL, NULL, NULL, 29, '0-6-7-8-29-31', '', '角色列表', 'example', 2, '/admin/api/v1/role/list', '', 'POST', '', 0, 0, 0),
(32, NULL, NULL, NULL, 29, '0-6-7-8-29-32', '', '角色展示', 'example', 2, '/admin/api/v1/role/:id', '', 'GET', '', 0, 0, 0),
(33, NULL, NULL, NULL, 29, '0-6-7-8-29-33', '', '角色新增', 'example', 2, '/admin/api/v1/role/', '', 'POST', '', 0, 0, 0),
(34, NULL, NULL, NULL, 29, '0-6-7-8-29-34', '', '角色更新', 'example', 2, '/admin/api/v1/role/:id', '', 'PUT', '', 0, 0, 0),
(35, NULL, NULL, NULL, 29, '0-6-7-8-29-35', '', '角色状态更新', 'example', 2, '/admin/api/v1/role/:id/status/', '', 'PUT', '', 0, 0, 0),
(36, NULL, NULL, NULL, 29, '0-6-7-8-29-36', '', '角色删除', 'example', 2, '/admin/api/v1/role/:id', '', 'DELETE', '', 0, 0, 0),
(37, NULL, NULL, NULL, 29, '0-6-7-8-29-37', '', '批量角色删除', 'example', 2, '/admin/api/v1/role/', '', 'DELETE', '', 0, 0, 0),
(38, NULL, NULL, NULL, 8, '0-6-7-8-38', '', '网站设置', 'example', 0, '/', '/', '', '', 0, 0, 0),
(39, NULL, NULL, NULL, 38, '0-6-7-8-38-39', '', '网站设置详情', 'example', 2, '/admin/api/v1/website/info', '', 'GET', '', 0, 0, 0),
(40, NULL, NULL, NULL, 38, '0-6-7-8-38-40', '', '网站设置更新', 'example', 2, '/admin/api/v1/website/info', '', 'PUT', '', 0, 0, 0),
(41, NULL, NULL, NULL, 38, '0-6-7-8-38-41', '', '网站设置上传', 'example', 2, '/admin/api/v1/website/upload', '', 'POST', '', 0, 0, 0),
(42, NULL, NULL, NULL, 1, '0-1-42', '', '操作日志', 'log', 1, 'operation_log', '/admin/operation_log', '', '', 0, 1, 0),
(43, NULL, NULL, NULL, 0, '0-43', '', '关于我们', 'introduction', 0, '/about', 'Layout', '', '', 20, 1, 0),
(44, NULL, NULL, NULL, 43, '0-43-44', '', '公司介绍', 'introduction', 1, 'company', '/about/company_introduction', '', '', 0, 1, 0),
(45, NULL, NULL, NULL, 0, '0-45', '', '轮播图管理', 'banner', 0, '/banner', 'Layout', '', '', 40, 1, 0),
(46, NULL, NULL, NULL, 45, '0-45-46', '', '轮播图管理', 'banner', 1, 'index', '/banner/index', '', '', 0, 1, 0),
(47, NULL, NULL, NULL, 0, '0-47', '', '视频管理', 'eye-open', 0, '/video', 'Layout', '', '', 30, 1, 0),
(48, NULL, NULL, NULL, 47, '0-47-48', '', '视频', 'eye-open', 1, 'index', '/video/index', '', '', 0, 1, 0),
(49, NULL, NULL, NULL, 7, '0-6-7-49', '', '轮播图管理API', 'example', 0, '/', '/', '', '', 0, 0, 0),
(50, NULL, NULL, NULL, 7, '0-6-7-50', '', '视频管理API', 'example', 0, '/', '/', '', '', 0, 0, 0),
(51, NULL, NULL, NULL, 7, '0-6-7-51', '', '关于我们API', 'example', 0, '/', '/', '', '', 0, 0, 0),
(52, NULL, NULL, NULL, 49, '0-6-7-49-52', '', '轮播图列表', 'example', 2, '/admin/api/v1/banner/', '', 'GET', '', 0, 0, 0),
(53, NULL, NULL, NULL, 49, '0-6-7-49-53', '', '创建轮播图', 'example', 2, '/admin/api/v1/banner/', '', 'POST', '', 0, 0, 0),
(54, NULL, NULL, NULL, 49, '0-6-7-49-54', '', '轮播图详情', 'example', 2, '/admin/api/v1/banner/:id', '', 'GET', '', 0, 0, 0),
(55, NULL, NULL, NULL, 49, '0-6-7-49-55', '', '更新轮播图', 'example', 2, '/admin/api/v1/banner/:id', '', 'PUT', '', 0, 0, 0),
(56, NULL, NULL, NULL, 49, '0-6-7-49-56', '', '上传轮播图', 'example', 2, '/admin/api/v1/banner/upload', '', 'POST', '', 0, 0, 0),
(57, NULL, NULL, NULL, 50, '0-6-7-50-57', '', '视频列表', 'example', 2, '/admin/api/v1/video/list', '', 'GET', '', 0, 0, 0),
(58, NULL, NULL, NULL, 50, '0-6-7-50-58', '', '创建视频', 'example', 2, '/admin/api/v1/video/', '', 'POST', '', 0, 0, 0),
(59, NULL, NULL, NULL, 50, '0-6-7-50-59', '', '更新视频', 'example', 2, '/admin/api/v1/video/:id', '', 'PUT', '', 0, 0, 0),
(60, NULL, NULL, NULL, 50, '0-6-7-50-60', '', '删除视频', 'example', 2, '/admin/api/v1/video/:id', '', 'DELETE', '', 0, 0, 0),
(61, NULL, NULL, NULL, 50, '0-6-7-50-61', '', '上传视频', 'example', 2, '/admin/api/v1/video/upload', '', 'POST', '', 0, 0, 0),
(62, NULL, NULL, NULL, 51, '0-6-7-51-62', '', '查看公司介绍', 'example', 2, '/admin/api/v1/company/introduction', '', 'GET', '', 0, 0, 0),
(63, NULL, NULL, NULL, 51, '0-6-7-51-63', '', '更新公司介绍', 'example', 2, '/admin/api/v1/company/introduction', '', 'PUT', '', 0, 0, 0),
(64, NULL, NULL, NULL, 51, '0-6-7-51-64', '', '上传图片', 'example', 2, '/admin/api/v1/company/upload', '', 'POST', '', 0, 0, 0),
(65, NULL, NULL, NULL, 8, '0-6-7-8-65', '', '操作日志管理', 'example', 0, '/', '/', '', '', 0, 0, 0),
(66, NULL, NULL, NULL, 65, '0-6-7-8-65-66', '', '操作日志列表', 'example', 2, '/admin/api/v1/operation/logs/', '', 'GET', '', 0, 0, 0),
(67, NULL, NULL, NULL, 65, '0-6-7-8-65-67', '', '操作日志详情', 'example', 2, '/admin/api/v1/operation/logs/:id', '', 'GET', '', 0, 0, 0),
(68, NULL, NULL, NULL, 50, '0-6-7-50-68', '', '批量删除视频', 'example', 2, '/admin/api/v1/video/', '', 'DELETE', '', 0, 0, 0);


-- +migrate Down
DROP TABLE `admin_menus`;