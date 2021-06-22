-- +migrate Up
DROP TABLE IF EXISTS `admin_role_permissions`;
-- Create the table in the specified schema
CREATE TABLE `admin_role_permissions`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `role_id`       int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
    `permission_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_role_id_permission_id` (`role_id`, `permission_id`) USING BTREE,
    KEY `idx_permission_id` (`permission_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台角色关联权限表';

-- +migrate Down
DROP TABLE `admin_role_permissions`;