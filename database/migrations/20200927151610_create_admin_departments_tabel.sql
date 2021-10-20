-- +migrate Up
DROP TABLE IF EXISTS `admin_departments`;
-- Create the table in the specified schema
CREATE TABLE `admin_departments`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `parent_id`  int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '父ID',
    `name`       varchar(128)        NOT NULL DEFAULT '' COMMENT '部门名称',
    `remark`     varchar(255)        NOT NULL DEFAULT '' COMMENT '部门备注',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '0' COMMENT '排序，数字越大越靠前',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台部门表';

INSERT INTO `admin_departments` (`id`, `created_at`, `updated_at`, `parent_id`, `name`, `remark`, `sort`, `status`)
VALUES (1, '2021-10-01 01:01:01', '2021-10-01 01:01:01', 0, '总部', '1', 100, 1);

-- +migrate Down
DROP TABLE `admin_departments`;