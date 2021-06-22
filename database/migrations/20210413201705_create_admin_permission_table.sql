-- +migrate Up
DROP TABLE IF EXISTS `admin_permissions`;
-- Create the table in the specified schema
CREATE TABLE `admin_permissions`
(
    `id`         INT(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `type`       tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
    `gid`        int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '权限分组ID',
    `name`       varchar(35)         NOT NULL DEFAULT '' COMMENT '权限名称',
    `path`       varchar(255)        NOT NULL DEFAULT '' COMMENT '接口地址',
    `method`     varchar(16)         NOT NULL DEFAULT '' COMMENT '菜单请求方法',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>关闭 1=>开启',
    `content`    varchar(255)        NOT NULL DEFAULT '' COMMENT '接口描述',
    PRIMARY KEY (`id`),
    KEY `permissions_gid_index` (`gid`),
    KEY `permissions_path_index` (`path`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='权限表';

INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (1, '2021-05-21 17:59:08', '2021-05-30 16:07:09', NULL, 0, 2, '强制下线', '/admin/api/v1/admins/logout/:id', 'POST', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (2, '2021-05-21 17:59:08', '2021-05-30 16:07:09', NULL, 0, 7, 'KEY存在', '/admin/api/v1/role/key', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (3, '2021-05-21 17:59:08', '2021-05-30 16:07:09', NULL, 0, 2, '所有角色', '/admin/api/v1/role/all', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (4, '2021-05-21 17:59:08', '2021-05-26 14:28:59', NULL, 0, 2, '管理员列表', '/admin/api/v1/admins/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (5, '2021-05-21 17:59:08', '2021-05-26 14:28:09', NULL, 0, 1, '用户信息', '/admin/api/v1/us/info', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (6, '2021-05-21 17:59:08', '2021-05-26 14:28:19', NULL, 0, 1, '用户菜单', '/admin/api/v1/us/navs', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (7, '2021-05-21 17:59:08', '2021-05-26 14:29:06', NULL, 0, 1, '用户退出', '/admin/api/v1/user/logout', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (8, '2021-05-21 17:59:08', '2021-05-26 14:30:43', NULL, 0, 8, '用户列表', '/admin/api/v1/user/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (9, '2021-05-21 17:59:08', '2021-05-26 14:30:49', NULL, 0, 8, '展示用户', '/admin/api/v1/user/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (10, '2021-05-21 17:59:08', '2021-05-26 14:30:54', NULL, 0, 8, '新增用户', '/admin/api/v1/user/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (11, '2021-05-21 17:59:08', '2021-05-26 14:31:01', NULL, 0, 8, '更新用户', '/admin/api/v1/user/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (12, '2021-05-21 17:59:08', '2021-05-26 14:31:20', NULL, 0, 8, '更新用户状态', '/admin/api/v1/user/:id/status/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (13, '2021-05-21 17:59:08', '2021-05-26 14:31:26', NULL, 0, 8, '用户重置密码', '/admin/api/v1/user/:id/reset/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (14, '2021-05-21 17:59:08', '2021-05-26 14:31:31', NULL, 0, 8, '用户删除', '/admin/api/v1/user/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (15, '2021-05-21 17:59:08', '2021-05-26 14:31:37', NULL, 0, 8, '批量删除', '/admin/api/v1/user/', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (16, '2021-05-21 17:59:08', '2021-05-26 14:52:47', NULL, 0, 3, '菜单列表', '/admin/api/v1/menu/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (17, '2021-05-21 17:59:08', '2021-05-26 14:53:02', NULL, 0, 3, '菜单详情', '/admin/api/v1/menu/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (18, '2021-05-21 17:59:08', '2021-05-26 14:53:07', NULL, 0, 3, '菜单新增', '/admin/api/v1/menu/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (19, '2021-05-21 17:59:08', '2021-05-26 14:53:12', NULL, 0, 3, '菜单更新', '/admin/api/v1/menu/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (20, '2021-05-21 17:59:08', '2021-05-26 14:53:18', NULL, 0, 3, '菜单删除', '/admin/api/v1/menu/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (21, '2021-05-21 17:59:08', '2021-05-26 14:53:26', NULL, 0, 7, '角色列表', '/admin/api/v1/role/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (22, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '所有角色', '/admin/api/v1/role/all', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (23, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色详情', '/admin/api/v1/role/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (24, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色新增', '/admin/api/v1/role/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (25, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色更新', '/admin/api/v1/role/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (26, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '状态更新', '/admin/api/v1/role/:id/status/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (27, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色删除', '/admin/api/v1/role/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (28, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 2, '重置密码', '/admin/api/v1/admins/:id/reset/', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (29, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置详情', '/admin/api/v1/website/info', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (30, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置更新', '/admin/api/v1/website/info', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (31, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置上传', '/admin/api/v1/website/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (32, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '轮播图列表', '/admin/api/v1/banner/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (33, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '创建轮播图', '/admin/api/v1/banner/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (34, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '轮播图详情', '/admin/api/v1/banner/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (35, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新轮播图', '/admin/api/v1/banner/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (36, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传轮播图', '/admin/api/v1/banner/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (37, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '视频列表', '/admin/api/v1/video/list', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (38, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '创建视频', '/admin/api/v1/video/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (39, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新视频', '/admin/api/v1/video/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (40, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '删除视频', '/admin/api/v1/video/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (41, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传视频', '/admin/api/v1/video/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (42, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '查看公司介绍', '/admin/api/v1/company/introduction', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (43, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新公司介绍', '/admin/api/v1/company/introduction', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (44, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传图片', '/admin/api/v1/company/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (45, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '操作日志列表', '/admin/api/v1/operation/logs/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (46, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '操作日志详情', '/admin/api/v1/operation/logs/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (47, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '批量删除视频', '/admin/api/v1/video/', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (48, '2021-05-27 09:33:03', '2021-05-27 09:33:03', NULL, 0, 5, '接口分组列表', '/admin/api/v1/permission/groups/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (49, '2021-05-27 09:33:36', '2021-05-27 09:33:36', NULL, 0, 6, '接口列表', '/admin/api/v1/permissions/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (50, '2021-05-27 14:35:04', '2021-05-27 14:35:04', NULL, 0, 7, '所有权限', '/admin/api/v1/group_permissions', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (51, '2021-05-27 14:35:30', '2021-05-27 14:35:30', NULL, 0, 6, '接口详情', '/admin/api/v1/permissions/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (52, '2021-05-27 14:36:12', '2021-05-27 14:36:12', NULL, 0, 5, '接口分组详情', '/admin/api/v1/permission/groups/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (53, '2021-05-27 23:58:14', '2021-05-27 23:58:14', NULL, 0, 2, '管理员详情', '/admin/api/v1/admins/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (54, '2021-05-30 11:21:34', '2021-05-30 11:21:34', NULL, 0, 6, '所有权限组', '/admin/api/v1/permission/groups/all', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (55, '2021-05-30 11:22:40', '2021-05-30 11:22:40', NULL, 0, 7, '所有菜单', '/admin/api/v1/menu/all', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (56, '2021-05-30 11:23:40', '2021-05-30 11:23:40', NULL, 0, 5, '更新接口分组', '/admin/api/v1/permission/groups/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (57, '2021-05-30 11:24:18', '2021-05-30 11:25:08', NULL, 0, 5, '新增接口分组', '/admin/api/v1/permission/groups/', 'POST', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (58, '2021-05-30 11:25:34', '2021-05-30 11:25:39', NULL, 0, 6, '更新接口', '/admin/api/v1/permissions/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (59, '2021-05-30 15:58:14', '2021-05-30 15:58:14', NULL, 0, 2, '更新管理员', '/admin/api/v1/admins/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (60, '2021-05-30 16:03:29', '2021-05-30 16:03:29', NULL, 0, 2, '新增管理员', '/admin/api/v1/admins/', 'POST', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (61, '2021-06-01 15:02:43', '2021-06-01 15:02:43', NULL, 0, 6, '接口新增', '/admin/api/v1/permissions/', 'POST', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (62, '2021-05-30 11:25:34', '2021-05-30 11:25:39', NULL, 0, 6, '删除接口', '/admin/api/v1/permissions/:id', 'DELETE', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (63, '2021-05-30 11:25:34', '2021-05-30 11:25:39', NULL, 0, 2, '用户名存在', '/admin/api/v1/admins/username/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (64, '2021-06-09 09:29:23', '2021-06-09 09:29:23', NULL, 0, 10, '日志列表', '/admin/api/v1/operation/logs/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions`(`id`, `created_at`, `updated_at`, `deleted_at`, `type`, `gid`, `name`, `path`, `method`, `sort`, `status`, `content`) VALUES (65, '2021-06-09 09:29:58', '2021-06-09 09:29:58', NULL, 0, 10, '日志详情', '/admin/api/v1/operation/logs/:id', 'GET', 100, 1, '');

-- +migrate Down
DROP TABLE `admin_permissions`;