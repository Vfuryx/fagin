-- +migrate Up
DROP TABLE IF EXISTS `admin_users`;
-- Create the table in the specified schema
CREATE TABLE `admin_users`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    `username`   varchar(64)         NOT NULL DEFAULT '' COMMENT '用户名',
    `nick_name`  varchar(64)         NOT NULL DEFAULT '' COMMENT '昵称',
    `phone`      varchar(64)         NOT NULL DEFAULT '' COMMENT '电话',
    `email`      varchar(64)         NOT NULL DEFAULT '' COMMENT '邮箱',
    `password`   varchar(127)        NOT NULL DEFAULT '' COMMENT '密码',
    `avatar`     varchar(255)        NOT NULL DEFAULT '' COMMENT '头像',
    `remark`     varchar(255)        NOT NULL DEFAULT '' COMMENT '备注',
    `sex`        tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '性别: 0:未知 1:男 2:女',
    `role_id`    int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '角色ID',
    PRIMARY KEY (`id`),
    KEY `idx_admin_users_deleted_at` (`deleted_at`),
    KEY `idx_admin_users_role_id` (`role_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台用户表';

INSERT INTO `admin_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `status`, `username`, `nick_name`, `phone`, `email`, `password`, `avatar`, `remark`, `sex`, `role_id`) VALUES
(1, NULL, NULL, NULL, 1, 'admin', 'admin', '', '', '$2a$10$v1Ji2oMpTD1sserUCuEfmeEUKQi4FKMonWah1fa4.xJhkxvPRvuN2', 'http://qiniu.furyx.cn/photo.jpg', '', 0, 1);

-- +migrate Down
DROP TABLE `admin_users`;