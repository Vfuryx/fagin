-- +migrate Up
DROP TABLE IF EXISTS `banners`;
-- Create the table in the specified schema
CREATE TABLE `banners`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `title`      varchar(32)         NOT NULL DEFAULT '' COMMENT '标题',
    `banner`     varchar(255)        NOT NULL DEFAULT '' COMMENT '轮播图',
    `path`       varchar(255)        NOT NULL DEFAULT '' COMMENT '路径',
    `sort`       int(10) unsigned    NOT NULL DEFAULT '100' COMMENT '排序，数字越大越靠前',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0:隐藏 1显示',
    PRIMARY KEY (`id`),
    KEY `idx_banners_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='Banner表';

-- +migrate Down
DROP TABLE `banners`;