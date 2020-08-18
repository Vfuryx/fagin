-- +migrate Up
DROP TABLE IF EXISTS `admin_roles`;
-- Create the table in the specified schema
CREATE TABLE `admin_roles`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `name`       varchar(128)        NOT NULL DEFAULT '' COMMENT '角色名称',
    `key`        varchar(32)         NOT NULL DEFAULT '' COMMENT '角色关键字',
    `remark`     varchar(255)        NOT NULL DEFAULT '' COMMENT '角色备注',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    PRIMARY KEY (`id`),
    KEY `idx_admin_roles_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台角色表';

INSERT INTO `admin_roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `key`, `remark`, `sort`, `status`) VALUES
(1, NULL, NULL, NULL, 'admin', 'admin', 'admin', 0, 1);

-- +migrate Down
DROP TABLE `admin_roles`;