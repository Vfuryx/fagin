-- +migrate Up
DROP TABLE IF EXISTS `video_infos`;
-- Create the table in the specified schema
CREATE TABLE `video_infos`
(
    `id`          int(11) unsigned    NOT NULL AUTO_INCREMENT,
    `created_at`  datetime                     DEFAULT NULL,
    `updated_at`  datetime                     DEFAULT NULL,
    `deleted_at`  datetime                     DEFAULT NULL,
    `author_id`   int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '作者',
    `title`       varchar(60)         NOT NULL DEFAULT '' COMMENT '标题',
    `path`        varchar(255)        NOT NULL DEFAULT '' COMMENT '路径',
    `duration`    varchar(32)         NOT NULL DEFAULT '' COMMENT '播放时长',
    `description` varchar(255)        NOT NULL DEFAULT '' COMMENT '视频描述',
    `status`      tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:关闭 1:开启',
    PRIMARY KEY (`id`),
    KEY `idx_video_infos_deleted_at` (`deleted_at`),
    KEY `idx_video_infos_author_id` (`author_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='视频资源表';

-- +migrate Down
DROP TABLE `video_infos`;