/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.10.10
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.10.10:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 31/05/2021 10:53:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menus
-- ----------------------------
DROP TABLE IF EXISTS `admin_menus`;
CREATE TABLE `admin_menus` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父菜单id',
  `paths` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单层级路径 0-1-2',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `component` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件名称',
  `path` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `redirect` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '重定向',
  `target` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '目标',
  `title` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单标题',
  `icon` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `is_show` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>隐藏 1=>展示',
  `is_hide_child` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>展示 1=>隐藏',
  `permission` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1开启',
  `method` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单请求方法',
  PRIMARY KEY (`id`),
  KEY `idx_admin_menus_deleted_at` (`deleted_at`),
  KEY `idx_admin_menus_parent_id` (`parent_id`),
  KEY `idx_admin_menus_path_method` (`method`,`path`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台菜单表';

-- ----------------------------
-- Records of admin_menus
-- ----------------------------
BEGIN;
INSERT INTO `admin_menus` VALUES (1, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-1', 'dashboard', 'RouteView', '/dashboard', '', '', '仪表盘', 'dashboard', 1, 0, '1', 100, 1, '');
INSERT INTO `admin_menus` VALUES (2, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 1, '0-1-2', 'workplace', 'Workplace', '/dashboard/workplace', '', '', '工作台', 'eye', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (3, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-3', 'User', 'RouteView', '/user', '', '', '用户中心', 'contacts', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (4, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-4', 'Manage', 'RouteView', '/user/manage', '', '', '用户管理', 'usergroup-add', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (5, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-5', 'ManageUsers', 'ManageUsers', '/user/manage/users', '', '', '用户中心', 'user', 1, 0, '', 100, 0, '');
INSERT INTO `admin_menus` VALUES (6, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-6', 'ManageAdmins', 'ManageAdmins', '/user/manage/admins', '', '', '管理员', 'smile', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (7, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 4, '0-3-4-7', 'UserForm', 'UserForm', 'users/form', '', '', '用户表单', 'eye', 0, 1, '', 100, 0, '');
INSERT INTO `admin_menus` VALUES (8, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-8', 'Menus', 'RouteView', '/user/menu', '', '', '菜单管理', 'bars', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (9, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 8, '0-3-8-9', 'AdminMenu', 'AdminMenu', '/user/menu/admin/index', '', '', '后台菜单', 'appstore', 1, 1, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (10, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 8, '0-3-8-10', 'SellerMenu', 'SellerMenu', '/', '', '', '商家菜单', 'gold', 1, 0, '', 100, 0, '');
INSERT INTO `admin_menus` VALUES (11, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-11', 'AdminPermissions', 'RouteView', '/admin/permissions', '', '', '接口管理', 'swap', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (12, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 11, '0-3-11-12', 'AdminPermissionsGroup', 'AdminPermissionsGroup', '/admin/permissions/group', '', '', '接口分组', 'menu-unfold', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (13, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 11, '0-3-11-13', 'AdminPermission', 'AdminPermission', '/admin/permissions/index', '', '', '接口列表', 'pic-right', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (14, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 3, '0-3-14', 'AdminRole', 'RouteView', '/admin/roles', '/admin/roles/index', '', '角色管理', 'meh', 1, 1, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (15, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 14, '0-3-14-15', 'AdminRoleIndex', 'AdminRoleIndex', '/admin/roles/index', '', '', '角色列表', 'eye', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (16, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 14, '0-3-14-16', 'AdminRoleForm', 'AdminRoleForm', '/admin/roles/form/:id?', '', '', '角色编辑', 'eye', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (17, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 0, '0-17', 'otherPage', 'PageView', '/other', '', '', '其他', '', 1, 0, '', 0, 1, '');
INSERT INTO `admin_menus` VALUES (18, '2021-05-13 23:36:21', '2021-05-24 22:09:42', NULL, 0, 17, '0-17-18', 'TestIconSelect', 'TestIconSelect', '/other/icon-selector', '', '', 'icon', '', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (19, '2021-05-29 17:29:24', '2021-05-30 16:04:37', NULL, 0, 0, '0-19', 'System', 'RouteView', '/system', '', '', '系统中心', 'setting', 1, 0, '', 100, 1, '');
INSERT INTO `admin_menus` VALUES (20, '2021-05-29 17:31:50', '2021-05-29 17:40:29', NULL, 0, 19, '0-19-20', 'OperationLogs', 'OperationLogs', '/system/operation_logs', '', '', '操作日志', 'file-protect', 1, 0, '', 100, 1, '');
COMMIT;

-- ----------------------------
-- Table structure for admin_operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `admin_operation_logs`;
CREATE TABLE `admin_operation_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `user` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户',
  `method` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '方法',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路径',
  `ip` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `location` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '访问位置',
  `module` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作模块',
  `operation` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '操作类型',
  `input` text COLLATE utf8mb4_unicode_ci COMMENT '输入',
  `latency_time` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '耗时',
  `user_agent` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ua',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:异常 1:正常',
  PRIMARY KEY (`id`),
  KEY `idx_admin_operation_logs_deleted_at` (`deleted_at`),
  KEY `idx_path` (`path`) USING BTREE,
  KEY `idx_method` (`method`) USING BTREE,
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台操作日志表';

-- ----------------------------
-- Table structure for admin_permission_groups
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission_groups`;
CREATE TABLE `admin_permission_groups` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(35) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限分组名称',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
  `sort` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '排序，数字越大越靠前',
  PRIMARY KEY (`id`),
  KEY `admin_permission_groups_deleted_at_index` (`deleted_at`),
  KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限组表';

-- ----------------------------
-- Records of admin_permission_groups
-- ----------------------------
BEGIN;
INSERT INTO `admin_permission_groups` VALUES (1, '2021-05-26 14:26:05', '2021-05-26 14:26:05', NULL, '公共接口', 0, 999);
INSERT INTO `admin_permission_groups` VALUES (2, '2021-05-26 14:26:17', '2021-05-26 14:26:17', NULL, '管理员管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (3, '2021-05-26 14:26:50', '2021-05-26 14:26:50', NULL, '总后台菜单管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (4, '2021-05-26 14:27:03', '2021-05-26 14:27:03', NULL, '商家后台菜单管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (5, '2021-05-26 14:27:15', '2021-05-26 14:27:15', NULL, '接口分组管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (6, '2021-05-26 14:27:26', '2021-05-26 14:27:26', NULL, '接口管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (7, '2021-05-26 14:27:35', '2021-05-26 14:27:35', NULL, '角色管理', 0, 100);
INSERT INTO `admin_permission_groups` VALUES (8, '2021-05-26 14:30:32', '2021-05-26 14:30:32', NULL, '会员管理', 0, 100);
COMMIT;

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
  `gid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限分组ID',
  `name` varchar(35) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限名称',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '接口地址',
  `method` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单请求方法',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态: 0=>关闭 1=>开启',
  `content` varchar(125) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '接口描述',
  PRIMARY KEY (`id`),
  KEY `permissions_gid_index` (`gid`),
  KEY `permissions_path_index` (`path`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_permissions` VALUES (3, '2021-05-21 17:59:08', '2021-05-30 16:07:09', NULL, 0, 2, '所有角色', '/admin/api/v1/role/list', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (4, '2021-05-21 17:59:08', '2021-05-26 14:28:59', NULL, 0, 2, '管理员列表', '/admin/api/v1/admins/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (5, '2021-05-21 17:59:08', '2021-05-26 14:28:09', NULL, 0, 1, '用户信息', '/admin/api/v1/us/info', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (6, '2021-05-21 17:59:08', '2021-05-26 14:28:19', NULL, 0, 1, '用户菜单', '/admin/api/v1/us/navs', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (7, '2021-05-21 17:59:08', '2021-05-26 14:29:06', NULL, 0, 1, '用户退出', '/admin/api/v1/user/logout', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (8, '2021-05-21 17:59:08', '2021-05-26 14:30:43', NULL, 0, 8, '用户列表', '/admin/api/v1/user/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (9, '2021-05-21 17:59:08', '2021-05-26 14:30:49', NULL, 0, 8, '展示用户', '/admin/api/v1/user/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (10, '2021-05-21 17:59:08', '2021-05-26 14:30:54', NULL, 0, 8, '新增用户', '/admin/api/v1/user/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (11, '2021-05-21 17:59:08', '2021-05-26 14:31:01', NULL, 0, 8, '更新用户', '/admin/api/v1/user/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (12, '2021-05-21 17:59:08', '2021-05-26 14:31:20', NULL, 0, 8, '更新用户状态', '/admin/api/v1/user/:id/status/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (13, '2021-05-21 17:59:08', '2021-05-26 14:31:26', NULL, 0, 8, '用户重置密码', '/admin/api/v1/user/:id/reset/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (14, '2021-05-21 17:59:08', '2021-05-26 14:31:31', NULL, 0, 8, '用户删除', '/admin/api/v1/user/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (15, '2021-05-21 17:59:08', '2021-05-26 14:31:37', NULL, 0, 8, '用户批量删除', '/admin/api/v1/user/', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (16, '2021-05-21 17:59:08', '2021-05-26 14:52:47', NULL, 0, 3, '菜单列表', '/admin/api/v1/menu/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (17, '2021-05-21 17:59:08', '2021-05-26 14:53:02', NULL, 0, 3, '菜单详情', '/admin/api/v1/menu/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (18, '2021-05-21 17:59:08', '2021-05-26 14:53:07', NULL, 0, 3, '菜单新增', '/admin/api/v1/menu/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (19, '2021-05-21 17:59:08', '2021-05-26 14:53:12', NULL, 0, 3, '菜单更新', '/admin/api/v1/menu/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (20, '2021-05-21 17:59:08', '2021-05-26 14:53:18', NULL, 0, 3, '菜单删除', '/admin/api/v1/menu/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (21, '2021-05-21 17:59:08', '2021-05-26 14:53:26', NULL, 0, 7, '角色列表', '/admin/api/v1/role/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (22, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '所有角色', '/admin/api/v1/role/list', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (23, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色详情', '/admin/api/v1/role/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (24, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色新增', '/admin/api/v1/role/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (25, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色更新', '/admin/api/v1/role/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (26, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '状态更新', '/admin/api/v1/role/:id/status/', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (27, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色删除', '/admin/api/v1/role/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (28, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 7, '角色删除', '/admin/api/v1/role/', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (29, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置详情', '/admin/api/v1/website/info', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (30, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置更新', '/admin/api/v1/website/info', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (31, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '网站设置上传', '/admin/api/v1/website/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (32, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '轮播图列表', '/admin/api/v1/banner/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (33, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '创建轮播图', '/admin/api/v1/banner/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (34, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '轮播图详情', '/admin/api/v1/banner/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (35, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新轮播图', '/admin/api/v1/banner/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (36, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传轮播图', '/admin/api/v1/banner/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (37, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '视频列表', '/admin/api/v1/video/list', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (38, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '创建视频', '/admin/api/v1/video/', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (39, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新视频', '/admin/api/v1/video/:id', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (40, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '删除视频', '/admin/api/v1/video/:id', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (41, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传视频', '/admin/api/v1/video/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (42, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '查看公司介绍', '/admin/api/v1/company/introduction', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (43, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '更新公司介绍', '/admin/api/v1/company/introduction', 'PUT', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (44, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '上传图片', '/admin/api/v1/company/upload', 'POST', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (45, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '操作日志列表', '/admin/api/v1/operation/logs/', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (46, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '操作日志详情', '/admin/api/v1/operation/logs/:id', 'GET', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (47, '2021-05-21 17:59:08', '2021-05-23 12:07:32', NULL, 0, 0, '批量删除视频', '/admin/api/v1/video/', 'DELETE', 0, 0, '');
INSERT INTO `admin_permissions` VALUES (48, '2021-05-27 09:33:03', '2021-05-27 09:33:03', NULL, 0, 5, '接口分组列表', '/admin/api/v1/permission/groups/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (49, '2021-05-27 09:33:36', '2021-05-27 09:33:36', NULL, 0, 6, '接口列表', '/admin/api/v1/permissions/', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (50, '2021-05-27 14:35:04', '2021-05-27 14:35:04', NULL, 0, 7, '所有权限', '/admin/api/v1/group_permissions', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (51, '2021-05-27 14:35:30', '2021-05-27 14:35:30', NULL, 0, 6, '接口详情', '/admin/api/v1/permissions/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (52, '2021-05-27 14:36:12', '2021-05-27 14:36:12', NULL, 0, 5, '接口分组详情', '/admin/api/v1/permission/groups/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (53, '2021-05-27 23:58:14', '2021-05-27 23:58:14', NULL, 0, 2, '管理员详情', '/admin/api/v1/admins/:id', 'GET', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (54, '2021-05-30 11:21:34', '2021-05-30 11:21:34', NULL, 0, 6, '所有权限组', '/admin/api/v1/permission/groups/all', 'POST', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (55, '2021-05-30 11:22:40', '2021-05-30 11:22:40', NULL, 0, 7, '所有菜单', '/admin/api/v1/menu/all', 'POST', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (56, '2021-05-30 11:23:40', '2021-05-30 11:23:40', NULL, 0, 5, '更新接口分组', '/admin/api/v1/permission/groups/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (57, '2021-05-30 11:24:18', '2021-05-30 11:25:08', NULL, 0, 5, '新增接口分组', '/admin/api/v1/permission/groups/', 'POST', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (58, '2021-05-30 11:25:34', '2021-05-30 11:25:39', NULL, 0, 6, '更新接口', '/admin/api/v1/permissions/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (59, '2021-05-30 15:58:14', '2021-05-30 15:58:14', NULL, 0, 2, '更新管理员', '/admin/api/v1/admins/:id', 'PUT', 100, 1, '');
INSERT INTO `admin_permissions` VALUES (60, '2021-05-30 16:03:29', '2021-05-30 16:03:29', NULL, 0, 2, '新增管理员', '/admin/api/v1/admins/', 'POST', 100, 1, '');
COMMIT;

-- ----------------------------
-- Table structure for admin_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menus`;
CREATE TABLE `admin_role_menus` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_id_menu_id` (`role_id`,`menu_id`) USING BTREE,
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台角色关联菜单表';

-- ----------------------------
-- Records of admin_role_menus
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_menus` VALUES (1, 1, 1);
INSERT INTO `admin_role_menus` VALUES (2, 1, 2);
INSERT INTO `admin_role_menus` VALUES (3, 1, 3);
INSERT INTO `admin_role_menus` VALUES (4, 1, 4);
INSERT INTO `admin_role_menus` VALUES (5, 1, 5);
INSERT INTO `admin_role_menus` VALUES (6, 1, 6);
INSERT INTO `admin_role_menus` VALUES (7, 1, 7);
INSERT INTO `admin_role_menus` VALUES (8, 1, 8);
INSERT INTO `admin_role_menus` VALUES (9, 1, 9);
INSERT INTO `admin_role_menus` VALUES (10, 1, 10);
INSERT INTO `admin_role_menus` VALUES (11, 1, 11);
INSERT INTO `admin_role_menus` VALUES (12, 1, 12);
INSERT INTO `admin_role_menus` VALUES (13, 1, 13);
INSERT INTO `admin_role_menus` VALUES (14, 1, 14);
INSERT INTO `admin_role_menus` VALUES (15, 1, 15);
INSERT INTO `admin_role_menus` VALUES (16, 1, 16);
INSERT INTO `admin_role_menus` VALUES (17, 1, 17);
INSERT INTO `admin_role_menus` VALUES (18, 1, 18);
INSERT INTO `admin_role_menus` VALUES (19, 1, 19);
INSERT INTO `admin_role_menus` VALUES (20, 1, 20);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permissions`;
CREATE TABLE `admin_role_permissions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `permission_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_id_permission_id` (`role_id`,`permission_id`) USING BTREE,
  KEY `idx_permission_id` (`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台角色关联权限表';

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
  `name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `key` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色关键字',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色备注',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_type_key` (`type`,`key`) USING BTREE,
  KEY `idx_admin_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台角色表';

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles` VALUES (1, NULL, NULL, NULL, 0, 'admin', 'admin', 'admin', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for admin_user_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_roles`;
CREATE TABLE `admin_user_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_admin_id_role_id` (`admin_id`,`role_id`) USING BTREE,
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员角色关联表';

-- ----------------------------
-- Records of admin_user_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_user_roles` VALUES (1, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
  `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `phone` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(127) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sex` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '性别: 0:未知 1:男 2:女',
  PRIMARY KEY (`id`),
  KEY `idx_admin_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台用户表';

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users` VALUES (1, '2021-05-30 16:03:29', '2021-05-30 16:03:29', NULL, 1, 'admin', 'admin', '', '', '$2a$10$v1Ji2oMpTD1sserUCuEfmeEUKQi4FKMonWah1fa4.xJhkxvPRvuN2', 'http://qiniu.furyx.cn/photo.jpg', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `summary` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '摘要',
  `author_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '作者ID',
  `category_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
  `post_at` datetime DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `index_author_id` (`author_id`) USING BTREE,
  KEY `index_category_id` (`category_id`) USING BTREE,
  KEY `index_status_post_at` (`status`,`post_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

-- ----------------------------
-- Table structure for authors
-- ----------------------------
DROP TABLE IF EXISTS `authors`;
CREATE TABLE `authors` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
  PRIMARY KEY (`id`),
  KEY `index_name` (`name`) USING BTREE,
  KEY `index_status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='作者表';

-- ----------------------------
-- Table structure for banners
-- ----------------------------
DROP TABLE IF EXISTS `banners`;
CREATE TABLE `banners` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `banner` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '轮播图',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路径',
  `sort` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '排序，数字越大越靠前',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
  PRIMARY KEY (`id`),
  KEY `idx_banners_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Banner表';

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'p_type',
  `v0` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v0',
  `v1` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v1',
  `v2` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v2',
  `v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v3',
  `v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v4',
  `v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'v5',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='casbin_rule表';

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (1, 'g', '1', 'admin', NULL, NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES (2, 'p', 'admin', '/*', '|', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `sort` int(10) unsigned NOT NULL DEFAULT '100' COMMENT '排序，数字越大越靠前',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
  PRIMARY KEY (`id`),
  KEY `index_name` (`name`) USING BTREE,
  KEY `index_status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

-- ----------------------------
-- Table structure for company_introductions
-- ----------------------------
DROP TABLE IF EXISTS `company_introductions`;
CREATE TABLE `company_introductions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司名称',
  `image` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图片',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
  PRIMARY KEY (`id`),
  KEY `idx_company_introductions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公司简介';

-- ----------------------------
-- Records of company_introductions
-- ----------------------------
BEGIN;
INSERT INTO `company_introductions` VALUES (1, '2020-07-01 19:49:13', '2020-07-01 19:49:13', NULL, '公司名称', '', '', 1);
COMMIT;

-- ----------------------------
-- Table structure for gorp_migrations
-- ----------------------------
DROP TABLE IF EXISTS `gorp_migrations`;
CREATE TABLE `gorp_migrations` (
  `id` varchar(255) NOT NULL,
  `applied_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gorp_migrations
-- ----------------------------
BEGIN;
INSERT INTO `gorp_migrations` VALUES ('20200818094901_create_users_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818101116_create_videos_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818102619_create_banners_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818102721_create_admin_operation_log_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818102806_create_website_config_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818102832_create_company_introduction_table.sql', '2021-05-31 10:53:29');
INSERT INTO `gorp_migrations` VALUES ('20200818102853_create_admin_menu_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20200818103121_create_admin_role_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20200818103136_create_admin_role_menu_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20200818103202_create_admin_user_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20200818103300_create_casbin_rule_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20201021215531_create_article_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20201021225344_create_category_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20201021225750_create_tag_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20201021230200_create_tag_to_article_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20201021231232_create_author_table.sql', '2021-05-31 10:53:30');
INSERT INTO `gorp_migrations` VALUES ('20210413201623_create_admin_permission_groups_table.sql', '2021-05-31 10:53:31');
INSERT INTO `gorp_migrations` VALUES ('20210413201705_create_admin_permission_table.sql', '2021-05-31 10:53:31');
INSERT INTO `gorp_migrations` VALUES ('20210415221458_create_admin_role_permission_table.sql', '2021-05-31 10:53:31');
INSERT INTO `gorp_migrations` VALUES ('20210525153925_create_admin_user_role_table.sql', '2021-05-31 10:53:31');
COMMIT;

-- ----------------------------
-- Table structure for tag_to_article
-- ----------------------------
DROP TABLE IF EXISTS `tag_to_article`;
CREATE TABLE `tag_to_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '标签ID',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
  PRIMARY KEY (`id`),
  KEY `index_tag_id_article_id` (`tag_id`,`article_id`) USING BTREE,
  KEY `index_article_id_tag_id` (`article_id`,`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签文章表';

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
  PRIMARY KEY (`id`),
  KEY `index_name` (`name`) USING BTREE,
  KEY `index_status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(127) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ----------------------------
-- Table structure for video_infos
-- ----------------------------
DROP TABLE IF EXISTS `video_infos`;
CREATE TABLE `video_infos` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `author_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '作者',
  `title` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路径',
  `duration` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '播放时长',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '视频描述',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1:开启',
  PRIMARY KEY (`id`),
  KEY `idx_video_infos_deleted_at` (`deleted_at`),
  KEY `idx_video_infos_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='视频资源表';

-- ----------------------------
-- Table structure for website_configs
-- ----------------------------
DROP TABLE IF EXISTS `website_configs`;
CREATE TABLE `website_configs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `web_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站名称',
  `web_en_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站英文名称',
  `web_title` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站标题',
  `keywords` varchar(127) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `company_name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司名称',
  `contact_number` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `company_address` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司地址',
  `email` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `icp` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ICP备案',
  `public_security_record` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公安部备案',
  `web_logo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站logo',
  `qr_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '二维码',
  PRIMARY KEY (`id`),
  KEY `idx_website_configs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='网站配置表';

-- ----------------------------
-- Records of website_configs
-- ----------------------------
BEGIN;
INSERT INTO `website_configs` VALUES (1, NULL, NULL, NULL, '官网', 'website', 'website', '官网,website', '官网详情', '公司名称', '联系电话', '公司地址', '邮箱地址', 'ICP备案', '公安部备案', '', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
