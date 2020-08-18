-- +migrate Up
DROP TABLE IF EXISTS `admin_operation_logs`;
-- Create the table in the specified schema
CREATE TABLE `admin_operation_logs`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`   datetime                     DEFAULT NULL,
    `updated_at`   datetime                     DEFAULT NULL,
    `deleted_at`   datetime                     DEFAULT NULL,
    `user`         varchar(32)         NOT NULL DEFAULT '' COMMENT '用户',
    `method`       varchar(8)          NOT NULL DEFAULT '' COMMENT '方法',
    `path`         varchar(255)        NOT NULL DEFAULT '' COMMENT '路径',
    `ip`           varchar(16)         NOT NULL DEFAULT '' COMMENT 'IP',
    `location`     varchar(128)        NOT NULL DEFAULT '' COMMENT '访问位置',
    `module`       varchar(32)         NOT NULL DEFAULT '' COMMENT '操作模块',
    `operation`    varchar(32)         NOT NULL DEFAULT '' COMMENT '操作类型',
    `input`        text COMMENT '输入',
    `latency_time` varchar(128)        NOT NULL DEFAULT '' COMMENT '耗时',
    `user_agent`   varchar(255)        NOT NULL DEFAULT '' COMMENT 'ua',
    `status`       tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:异常 1:正常',
    PRIMARY KEY (`id`),
    KEY `idx_admin_operation_logs_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台操作日志表';

-- +migrate Down
DROP TABLE `admin_operation_logs`;