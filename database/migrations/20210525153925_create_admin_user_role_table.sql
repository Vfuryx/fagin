-- +migrate Up
DROP TABLE IF EXISTS `admin_user_roles`;
-- Create the table in the specified schema
CREATE TABLE `admin_user_roles`
(
    `id`       int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `admin_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
    `role_id`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_admin_id_role_id` (`admin_id`, `role_id`) USING BTREE,
    KEY `idx_role_id` (`role_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='管理员角色关联表';

INSERT INTO `admin_user_roles`(`id`, `admin_id`, `role_id`) VALUES (1, 1, 1);

-- +migrate Down
DROP TABLE `admin_user_roles`;
