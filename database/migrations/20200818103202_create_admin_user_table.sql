-- +migrate Up
DROP TABLE IF EXISTS `admin_users`;
-- Create the table in the specified schema
CREATE TABLE `admin_users`
(
    `id`            int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at`    datetime                     DEFAULT NULL,
    `updated_at`    datetime                     DEFAULT NULL,
    `deleted_at`    datetime                     DEFAULT NULL,
    `status`        tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    `username`      varchar(64)         NOT NULL DEFAULT '' COMMENT '用户名',
    `nick_name`     varchar(64)         NOT NULL DEFAULT '' COMMENT '昵称',
    `phone`         varchar(64)         NOT NULL DEFAULT '' COMMENT '电话',
    `email`         varchar(64)         NOT NULL DEFAULT '' COMMENT '邮箱',
    `password`      varchar(127)        NOT NULL DEFAULT '' COMMENT '密码',
    `avatar`        varchar(255)        NOT NULL DEFAULT '' COMMENT '头像',
    `remark`        varchar(255)        NOT NULL DEFAULT '' COMMENT '备注',
    `home_path`     varchar(255)        NOT NULL DEFAULT '' COMMENT '默认首页路径',
    `sex`           tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '性别: 0:未知 1:男 2:女',
    `login_at`      bigint(19) unsigned NOT NULL DEFAULT '0' COMMENT '总后台登录时间',
    `last_login_at` bigint(19) unsigned NOT NULL DEFAULT '0' COMMENT '总后台上次登录时间',
    `check_in_at`   bigint(19) unsigned NOT NULL DEFAULT '0' COMMENT '总后台签发时间',
    PRIMARY KEY (`id`),
    KEY `idx_admin_users_deleted_at` (`deleted_at`),
    INDEX `idx_username` (`username`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台用户表';

INSERT INTO `admin_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `status`, `username`, `nick_name`,
                           `phone`, `email`, `password`, `avatar`, `remark`, `home_path`, `sex`, `login_at`, `last_login_at`,
                           `check_in_at`)
VALUES (1, '2021-10-01 01:01:01', '2021-10-01 01:01:01', NULL, 1, 'admin', 'admin', '1380013800', 'admin@admin.com',
        '$2a$10$v1Ji2oMpTD1sserUCuEfmeEUKQi4FKMonWah1fa4.xJhkxvPRvuN2', '', '', '', 0,
        1633228500, 0, 1633228500);
-- +migrate Down
DROP TABLE `admin_users`;