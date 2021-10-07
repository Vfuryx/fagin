-- +migrate Up
DROP TABLE IF EXISTS `tag_to_article`;
-- Create the table in the specified schema
CREATE TABLE `tag_to_article`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id`     int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签ID',
    `article_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
    PRIMARY KEY (`id`),
    INDEX `index_tag_id_article_id` (`tag_id`, `article_id`) USING BTREE,
    INDEX `index_article_id_tag_id` (`article_id`, `tag_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='标签文章表';
-- +migrate Down
DROP TABLE `tag_to_article`;