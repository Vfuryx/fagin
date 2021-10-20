-- +migrate Up
DROP TABLE IF EXISTS `admin_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_menus`
(
    `id`             int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at`     datetime                     DEFAULT NULL,
    `updated_at`     datetime                     DEFAULT NULL,
    `deleted_at`     datetime                     DEFAULT NULL,
    `parent_id`      int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '父id',
    `paths`          varchar(128)        NOT NULL DEFAULT '' COMMENT '层级路径 0-1-2',
    `type`           tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '类型: 0:目录; 1:菜单; 2:权限;',
    `name`           varchar(32)         NOT NULL DEFAULT '' COMMENT '名称',
    `permission`     varchar(64)         NOT NULL DEFAULT '' COMMENT '权限标记',
    `component`      varchar(255)        NOT NULL DEFAULT '' COMMENT '组件名称',
    `path`           varchar(255)        NOT NULL DEFAULT '' COMMENT '路由地址',
    `method`         varchar(16)         NOT NULL DEFAULT '' COMMENT '请求方法',
    `redirect`       varchar(255)        NOT NULL DEFAULT '' COMMENT '重定向',
    `frame_src`      varchar(255)        NOT NULL DEFAULT '' COMMENT '内嵌iframe的地址',
    `current_active` varchar(255)        NOT NULL DEFAULT '' COMMENT '当前激活的菜单',
    `title`          varchar(32)         NOT NULL DEFAULT '' COMMENT '标题',
    `icon`           varchar(128)        NOT NULL DEFAULT '' COMMENT '图标',
    `is_show`        tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '展示状态: 0=>隐藏 1=>展示',
    `is_hide_child`  tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否隐藏子菜单: 0=>展示 1=>隐藏',
    `is_no_cache`    tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否无缓存: 0=>否 1=>是',
    `sort`           int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `status`         tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1开启',
    PRIMARY KEY (`id`),
    KEY `idx_deleted_at` (`deleted_at`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `unique_path_method` (`type`, `method`, `path`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台菜单表';

INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (1, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 0, '0-1', 0, 'Dashboard', '', 'LAYOUT', '/dashboard', '',
        '/dashboard/analysis', '', '仪表盘', 'ant-design:dashboard-outlined', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (2, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 1, '0-1-2', 1, 'Workbench', 'dashboard.workbench',
        '/dashboard/workbench/index', '/dashboard/workbench', '', '', '', '工作台', '', 1, 1, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (3, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 1, '0-1-3', 1, 'Analysis', 'dashboard.analysis',
        '/dashboard/analysis/index', '/dashboard/analysis', '', '', '', '分析页', '', 1, 1, 0, 101, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (4, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 0, '0-4', 0, 'System', '', 'LAYOUT', '/system', '', '',
        '', '系统管理', 'ant-design:setting-outlined', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (7, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 4, '0-4-7', 1, 'AccountManagement',
        'admin.system.account', '/system/account/index', '/system/account', '', '', '', '账号管理',
        'ant-design:team-outlined', 1, 1, 0, 101, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (8, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 4, '0-4-8', 1, 'MenuManagement', 'admin.system.menu',
        '/system/menu/index', '/system/menu', '', '', '', '菜单管理', 'ant-design:menu-outlined', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (9, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 4, '0-4-9', 1, 'RoleManagement', 'admin.system.role',
        '/system/role/index', '/system/role', '', '', '', '角色管理', 'ant-design:user-switch-outlined', 1, 0, 0, 100, 1,
        '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (10, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-10', 1, 'AccountDetail',
        'admin.system.account_detail', '/system/account/AccountDetail', '/system/account_detail/:id', '', '', '',
        '账号详情', 'ant-design:user-outlined', 0, 1, 0, 100, 1, '/system/account');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (13, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 0, '0-13', 0, 'Monitor', '', 'LAYOUT', '/monitor', '',
        '', '', '系统监控', 'ant-design:monitor-outlined', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (14, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 13, '0-13-14', 1, 'LoginLogManagement',
        'admin.monitor.login-log', '/monitor/login-log/index', '/monitor/login-log', '', '', '', '登录日志',
        'ant-design:login-outlined', 1, 1, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (15, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 13, '0-13-15', 1, 'OperLogManagement',
        'admin.monitor.oper-log', '/monitor/oper-log/index', '/monitor/oper-log', '', '', '', '操作日志',
        'ant-design:control-outlined', 1, 1, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (16, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 4, '0-4-16', 1, 'DepartmentManagement',
        'admin.system.department', '/system/department/index', '/system/department', '', '', '', '部门管理',
        'ant-design:apartment-outlined', 1, 1, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (17, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 16, '0-4-16-17', 2, '', 'admin.system.department:query',
        '', '/admin/api/departments', 'GET', '', '', '部门查询', 'ant-design:api-twotone', 0, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (18, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 16, '0-4-16-18', 2, '',
        'admin.system.department:detail', '', '/admin/api/departments/:id', 'GET', '', '', '部门详情',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (19, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 16, '0-4-16-19', 2, '',
        'admin.system.department:create', '', '/admin/api/departments', 'POST', '', '', '新增部门',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (20, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 16, '0-4-16-20', 2, '',
        'admin.system.department:update', '', '/admin/api/departments/:id', 'PUT', '', '', '更新部门',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (21, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 16, '0-4-16-21', 2, '',
        'admin.system.department:delete', '', '/admin/api/departments/:id', 'DELETE', '', '', '删除部门',
        'ant-design:api-twotone', 0, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (22, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-22', 2, '', 'admin.system.role:query', '',
        '/admin/api/roles', 'GET', '', '', '角色查询', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (23, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-23', 2, '', 'admin.system.role:detail', '',
        '/admin/api/roles/:id', 'GET', '', '', '角色详情', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (24, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-24', 2, '', 'admin.system.role:create', '',
        '/admin/api/roles', 'POST', '', '', '角色新增', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (25, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-25', 2, '', 'admin.system.role:update', '',
        '/admin/api/roles/:id', 'PUT', '', '', '角色更新', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (26, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-26', 2, '', 'admin.system.role:delete', '',
        '/admin/api/roles/:id', 'DELETE', '', '', '角色删除', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (27, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-27', 2, '', 'admin.system.role:key-exists',
        '', '/admin/api/roles/key', 'GET', '', '', '角色key是否存在', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (28, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-28', 2, '', 'admin.system.role:update-status',
        '', '/admin/api/roles/:id/status', 'PUT', '', '', '更改状态', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (29, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 9, '0-4-9-29', 2, '', 'admin.system.role:menus', '',
        '/admin/api/roles/menus', 'GET', '', '', '所有菜单', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (30, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-30', 2, '', 'admin.system.menu:query', '',
        '/admin/api/menus', 'GET', '', '', '菜单查询', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (31, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-31', 2, '', 'admin.system.menu:detail', '',
        '/admin/api/menus/:id', 'GET', '', '', '菜单详情', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (32, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-32', 2, '', 'admin.system.menu:create', '',
        '/admin/api/menus', 'POST', '', '', '新增菜单', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (33, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-33', 2, '', 'admin.system.menu:update', '',
        '/admin/api/menus/:id', 'PUT', '', '', '更新菜单', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (34, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-34', 2, '', 'admin.system.menu:delete', '',
        '/admin/api/menus/:id', 'DELETE', '', '', '删除菜单', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (35, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 8, '0-4-8-35', 2, '', 'admin.system.menu:all', '',
        '/admin/api/menus/all', 'GET', '', '', '所有菜单', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (36, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 15, '0-13-15-36', 2, '', 'admin.monitor.oper-log:query',
        '', '/admin/api/operation/logs', 'GET', '', '', '操作日志查询', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (37, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 15, '0-13-15-37', 2, '',
        'admin.monitor.oper-log:detail', '', '/admin/api/operation/logs/:id', 'GET', '', '', '操作日志详情',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (38, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 14, '0-13-14-38', 2, '',
        'admin.monitor.login-log:query', '', '/admin/api/login-logs', 'GET', '', '', '登录日志查询', 'ant-design:api-twotone',
        1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (39, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 14, '0-13-14-39', 2, '',
        'admin.monitor.login-log:detail', '', '/admin/api/login-logs/:id', 'GET', '', '', '登录日志详情',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (40, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-40', 2, '', 'admin.system.account:query', '',
        '/admin/api/accounts', 'GET', '', '', '用户查询', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (41, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-41', 2, '', 'admin.system.account:detail', '',
        '/admin/api/accounts/:id', 'GET', '', '', '用户详情', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (42, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-42', 2, '', 'admin.system.account:create', '',
        '/admin/api/accounts', 'POST', '', '', '新增用户', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (43, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-43', 2, '', 'admin.system.account:update', '',
        '/admin/api/accounts/:id', 'PUT', '', '', '更新用户', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (44, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-44', 2, '', 'admin.system.account:delete', '',
        '/admin/api/accounts/:id', 'DELETE', '', '', '删除用户', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (45, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-45', 2, '', 'admin.system.account:update-pwd',
        '', '/admin/api/accounts/:id/pwd', 'PUT', '', '', '用户更新密码', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (46, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-46', 2, '',
        'admin.system.account:update-status', '', '/admin/api/accounts/:id/status', 'PUT', '', '', '用户更新状态',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (47, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-47', 2, '', 'admin.system.account:logout', '',
        '/admin/api/accounts/:id/logout', 'POST', '', '', '强制下线', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (48, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-48', 2, '', 'admin.system.account:exists', '',
        '/admin/api/accounts/exists', 'GET', '', '', '账号名是否已存在', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (49, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-49', 2, '', 'admin.system.account:roles', '',
        '/admin/api/accounts/roles', 'GET', '', '', '角色列表', 'ant-design:api-twotone', 1, 0, 0, 100, 1, '');
INSERT INTO `admin_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `paths`, `type`, `name`,
                           `permission`, `component`, `path`, `method`, `redirect`, `frame_src`, `title`, `icon`,
                           `is_show`, `is_hide_child`, `is_no_cache`, `sort`, `status`, `current_active`)
VALUES (50, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 7, '0-4-7-50', 2, '',
        'admin.system.account:departments', '', '/admin/api/accounts/departments', 'GET', '', '', '部门列表',
        'ant-design:api-twotone', 1, 0, 0, 100, 1, '');


-- +migrate Down
DROP TABLE `admin_menus`;