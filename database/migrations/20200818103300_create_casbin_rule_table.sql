-- +migrate Up
DROP TABLE IF EXISTS `casbin_rule`;
-- Create the table in the specified schema
CREATE TABLE `casbin_rule`
(
    `p_type` varchar(100) DEFAULT NULL COMMENT 'p_type',
    `v0`     varchar(100) DEFAULT NULL COMMENT 'v0',
    `v1`     varchar(100) DEFAULT NULL COMMENT 'v1',
    `v2`     varchar(100) DEFAULT NULL COMMENT 'v2',
    `v3`     varchar(100) DEFAULT NULL COMMENT 'v3',
    `v4`     varchar(100) DEFAULT NULL COMMENT 'v4',
    `v5`     varchar(100) DEFAULT NULL COMMENT 'v5'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='casbin_rule表';

INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES ('g', '1', 'admin', NULL, NULL, NULL, NULL),
       ('p', 'admin', '/*', '|', NULL, NULL, NULL);

-- +migrate Down
DROP TABLE `casbin_rule`;