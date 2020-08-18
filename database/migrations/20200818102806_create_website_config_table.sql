-- +migrate Up
DROP TABLE IF EXISTS `website_configs`;
-- Create the table in the specified schema
CREATE TABLE `website_configs`
(
    `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`             datetime                  DEFAULT NULL,
    `updated_at`             datetime                  DEFAULT NULL,
    `deleted_at`             datetime                  DEFAULT NULL,
    `web_name`               varchar(32)      NOT NULL DEFAULT '' COMMENT '网站名称',
    `web_en_name`            varchar(32)      NOT NULL DEFAULT '' COMMENT '网站英文名称',
    `web_title`              varchar(32)      NOT NULL DEFAULT '' COMMENT '网站标题',
    `keywords`               varchar(127)     NOT NULL DEFAULT '' COMMENT '关键词',
    `description`            varchar(255)     NOT NULL DEFAULT '' COMMENT '描述',
    `company_name`           varchar(32)      NOT NULL DEFAULT '' COMMENT '公司名称',
    `contact_number`         varchar(16)      NOT NULL DEFAULT '' COMMENT '联系电话',
    `company_address`        varchar(32)      NOT NULL DEFAULT '' COMMENT '公司地址',
    `email`                  varchar(32)      NOT NULL DEFAULT '' COMMENT '邮箱地址',
    `icp`                    varchar(32)      NOT NULL DEFAULT '' COMMENT 'ICP备案',
    `public_security_record` varchar(32)      NOT NULL DEFAULT '' COMMENT '公安部备案',
    `web_logo`               varchar(255)     NOT NULL DEFAULT '' COMMENT '网站logo',
    `qr_code`                varchar(255)     NOT NULL DEFAULT '' COMMENT '二维码',
    PRIMARY KEY (`id`),
    KEY `idx_website_configs_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='网站配置表';


INSERT INTO `website_configs` (`id`, `created_at`, `updated_at`, `deleted_at`, `web_name`, `web_en_name`, `web_title`, `keywords`, `description`, `company_name`, `contact_number`, `company_address`, `email`, `icp`, `public_security_record`, `web_logo`, `qr_code`) VALUES
(1, NULL, NULL, NULL, '官网', 'website', 'website', '官网,website', '官网详情', '公司名称', '联系电话', '公司地址', '邮箱地址', 'ICP备案', '公安部备案', '', '');


-- +migrate Down
DROP TABLE `website_configs`;