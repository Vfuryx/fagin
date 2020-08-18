-- +migrate Up
DROP TABLE IF EXISTS `users`;
-- Create the table in the specified schema
CREATE TABLE `users`
(
    `id`         INT(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `username`   varchar(64)         NOT NULL DEFAULT '' COMMENT '用户名',
    `email`      varchar(64)         NOT NULL DEFAULT '' COMMENT '邮箱',
    `password`   varchar(127)        NOT NULL DEFAULT '' COMMENT '密码',
    `avatar`     varchar(255)        NOT NULL DEFAULT '' COMMENT '头像',
    `phone`      varchar(20)         NOT NULL DEFAULT '' COMMENT '手机',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户表';

-- +migrate Down
DROP TABLE `users`;