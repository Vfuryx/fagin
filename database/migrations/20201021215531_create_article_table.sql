-- +migrate Up
DROP TABLE IF EXISTS `articles`;
-- Create the table in the specified schema
CREATE TABLE `articles`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime                  DEFAULT NULL,
    `updated_at`  datetime                  DEFAULT NULL,
    `deleted_at`  datetime                  DEFAULT NULL,
    `title`       varchar(128)     NOT NULL COMMENT '标题',
    `content`     text             NOT NULL COMMENT '内容',
    `summary`     text             NOT NULL COMMENT '摘要',
    `author_id`   int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '作者ID',
    `category_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分类ID',
    `status`      tinyint(4)       NOT NULL DEFAULT 1 COMMENT '状态 0:隐藏 1显示',
    `post_at`     datetime         NULL     DEFAULT NULL COMMENT '发布时间',
    PRIMARY KEY (`id`),
    INDEX `index_author_id` (`author_id`) USING BTREE,
    INDEX `index_category_id` (`category_id`) USING BTREE,
    INDEX `index_status_post_at` (`status`, `post_at`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='文章表';
-- +migrate Down
DROP TABLE `articles`;