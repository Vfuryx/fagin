-- +migrate Up
DROP TABLE IF EXISTS `admin_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_menus`
(
    `id`            int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at`    datetime                     DEFAULT NULL,
    `updated_at`    datetime                     DEFAULT NULL,
    `deleted_at`    datetime                     DEFAULT NULL,
    `type`          tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
    `parent_id`     int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '父菜单id',
    `paths`         varchar(128)        NOT NULL DEFAULT '' COMMENT '菜单层级路径 0-1-2',
    `name`          varchar(32)         NOT NULL DEFAULT '' COMMENT '菜单名称',
    `component`     varchar(32)         NOT NULL DEFAULT '' COMMENT '组件名称',
    `path`          varchar(128)        NOT NULL DEFAULT '' COMMENT '路由地址',
    `redirect`      varchar(128)        NOT NULL DEFAULT '' COMMENT '重定向',
    `target`        varchar(128)        NOT NULL DEFAULT '' COMMENT '目标',
    `title`         varchar(32)         NOT NULL DEFAULT '' COMMENT '菜单标题',
    `icon`          varchar(128)        NOT NULL DEFAULT '' COMMENT '图标',
    `is_show`       tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>隐藏 1=>展示',
    `is_hide_child` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>展示 1=>隐藏',
    `permission`    varchar(32)         NOT NULL DEFAULT '' COMMENT '权限',
    `sort`          int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `status`        tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1开启',
    `method`        varchar(16)         NOT NULL DEFAULT '' COMMENT '菜单请求方法',
    PRIMARY KEY (`id`),
    KEY `idx_admin_menus_deleted_at` (`deleted_at`),
    KEY `idx_admin_menus_parent_id` (`parent_id`),
    KEY `idx_admin_menus_path_method` (`method`, `path`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台菜单表';

INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (1, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-1', 'dashboard', 'RouteView', '/dashboard', '', '', '仪表盘', 'dashboard', 1, '1', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (2, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 1, '0-1-2', 'workplace', 'Workplace', '/dashboard/workplace', '', '', '工作台', 'eye', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (3, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-3', 'User', 'RouteView', '/user', '', '', '用户中心', 'contacts', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (4, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-4', 'Manage', 'RouteView', '/user/manage', '', '', '用户管理', 'usergroup-add', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (5, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-5', 'ManageUsers', 'ManageUsers', '/user/manage/users', '', '', '用户中心', 'user', 1, '', 100, 0, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (6, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-6', 'ManageAdmins', 'ManageAdmins', '/user/manage/admins', '', '', '管理员', 'smile', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (7, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-7', 'UserForm', 'UserForm', 'users/form', '', '', '用户表单', 'eye', 0, '', 100, 0, '', 1);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (8, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-8', 'Menus', 'RouteView', '/user/menu', '', '', '菜单管理', 'bars', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (9, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 8, '0-3-8-9', 'AdminMenu', 'AdminMenu', '/user/menu/admin/index', '', '', '后台菜单', 'appstore', 1, '', 100, 1, '', 1);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (10, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 8, '0-3-8-10', 'SellerMenu', 'SellerMenu', '/', '', '', '商家菜单', 'gold', 1, '', 100, 0, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (11, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-11', 'AdminPermissions', 'RouteView', '/admin/permissions', '', '', '接口管理', 'swap', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (12, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 11, '0-3-11-12', 'AdminPermissionsGroup', 'AdminPermissionsGroup', '/admin/permissions/group', '', '', '接口分组', 'menu-unfold', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (13, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 11, '0-3-11-13', 'AdminPermission', 'AdminPermission', '/admin/permissions/index', '', '', '接口列表', 'pic-right', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (14, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-14', 'AdminRole', 'RouteView', '/admin/roles', '/admin/roles/index', '', '角色管理', 'meh', 1, '', 100, 1, '', 1);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (15, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 14, '0-3-14-15', 'AdminRoleIndex', 'AdminRoleIndex', '/admin/roles/index', '', '', '角色列表', 'eye', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (16, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 14, '0-3-14-16', 'AdminRoleForm', 'AdminRoleForm', '/admin/roles/form/:id?', '', '', '角色编辑', 'eye', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (17, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-17', 'otherPage', 'PageView', '/other', '', '', '其他', '', 1, '', 0, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `permission`, `sort`, `status`, `method`, `is_hide_child`) VALUES (18, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 17, '0-17-18', 'TestIconSelect', 'TestIconSelect', '/other/icon-selector', '', '', 'icon', '', 1, '', 100, 1, '', 0);
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `is_hide_child`, `permission`, `sort`, `status`, `method`) VALUES (19, '2021-05-29 17:29:24', '2021-05-30 16:04:37', NULL, 0, 0, '0-19', 'System', 'RouteView', '/system', '', '', '系统中心', 'setting', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `parent_id`, `paths`, `name`, `component`, `path`, `redirect`, `target`, `title`, `icon`, `is_show`, `is_hide_child`, `permission`, `sort`, `status`, `method`) VALUES (20, '2021-05-29 17:31:50', '2021-05-29 17:40:29', NULL, 0, 19, '0-19-20', 'OperationLogs', 'OperationLogs', '/system/operation_logs', '', '', '操作日志', 'file-protect', 1, 0, '', 100, 1, '');

-- +migrate Down
DROP TABLE `admin_menus`;