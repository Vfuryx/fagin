-- +migrate Up
DROP TABLE IF EXISTS `admin_login_logs`;
-- Create the table in the specified schema
CREATE TABLE `admin_login_logs`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
    `created_at` datetime            DEFAULT NULL,
    `updated_at` datetime            DEFAULT NULL,
    `deleted_at` datetime            DEFAULT NULL,
    `login_name` varchar(32)         DEFAULT '' COMMENT '登录账号',
    `ip`         varchar(46)         DEFAULT '' COMMENT '登录IP地址',
    `location`   varchar(128)        DEFAULT '' COMMENT '登录地点',
    `browser`    varchar(32)         DEFAULT '' COMMENT '浏览器类型',
    `os`         varchar(50)         DEFAULT '' COMMENT '操作系统',
    `status`     tinyint(4) unsigned DEFAULT '0' COMMENT '登录状态 0:异常 1:正常',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台登录日志表';

-- +migrate Down
DROP TABLE `admin_login_logs`;