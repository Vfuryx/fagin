-- +migrate Up
DROP TABLE IF EXISTS `company_introductions`;
-- Create the table in the specified schema
CREATE TABLE `company_introductions`
(
    `id`         int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at` datetime                     DEFAULT NULL,
    `updated_at` datetime                     DEFAULT NULL,
    `deleted_at` datetime                     DEFAULT NULL,
    `name`       varchar(32)         NOT NULL DEFAULT '' COMMENT '公司名称',
    `image`      varchar(255)        NOT NULL DEFAULT '' COMMENT '图片',
    `content`    text                NOT NULL COMMENT '内容',
    `status`     tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '状态: 0=>关闭 1=>开启',
    PRIMARY KEY (`id`),
    KEY `idx_company_introductions_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='公司简介';

INSERT INTO `company_introductions` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `image`, `content`, `status`) VALUES
(1, '2020-07-01 19:49:13', '2020-07-01 19:49:13', NULL, '公司名称', '', '', 1);

-- +migrate Down
DROP TABLE `company_introductions`;