-- +migrate Up
DROP TABLE IF EXISTS `admin_permission_groups`;
-- Create the table in the specified schema
CREATE TABLE `admin_permission_groups`
(
    `id`         int(10) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `name`       varchar(35)         NOT NULL DEFAULT '' COMMENT '权限分组名称',
    `type`       tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:总后台 1:商家后台 2:集团后台',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '100' COMMENT '排序，数字越大越靠前',
    PRIMARY KEY (`id`),
    KEY `admin_permission_groups_deleted_at_index` (`deleted_at`),
    KEY `idx_type` (`type`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='权限组表';

INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (1, '2021-05-26 14:26:05', '2021-05-26 14:26:05', NULL, '公共接口', 0, 999);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (2, '2021-05-26 14:26:17', '2021-05-26 14:26:17', NULL, '管理员管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (3, '2021-05-26 14:26:50', '2021-05-26 14:26:50', NULL, '总后台菜单管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (4, '2021-05-26 14:27:03', '2021-05-26 14:27:03', NULL, '商家后台菜单管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (5, '2021-05-26 14:27:15', '2021-05-26 14:27:15', NULL, '接口分组管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (6, '2021-05-26 14:27:26', '2021-05-26 14:27:26', NULL, '接口管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (7, '2021-05-26 14:27:35', '2021-05-26 14:27:35', NULL, '角色管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (8, '2021-05-26 14:30:32', '2021-05-26 14:30:32', NULL, '会员管理', 0, 100);
INSERT INTO `admin_permission_groups`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `sort`) VALUES (9, '2021-06-09 09:28:07', '2021-06-09 09:28:07', NULL, '操作日志管理', 0, 100);

-- +migrate Down
DROP TABLE `admin_permission_groups`;