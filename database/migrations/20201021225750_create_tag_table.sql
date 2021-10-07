-- +migrate Up
DROP TABLE IF EXISTS `tags`;
-- Create the table in the specified schema
CREATE TABLE `tags`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `name`       varchar(32)     NOT NULL COMMENT '名称',
    `status`     tinyint(4)       NOT NULL DEFAULT 1 COMMENT '状态 0:隐藏 1显示',
    PRIMARY KEY (`id`),
    INDEX `index_name` (`name`) USING BTREE,
    INDEX `index_status` (`status`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='标签表';
-- +migrate Down
DROP TABLE `tags`;