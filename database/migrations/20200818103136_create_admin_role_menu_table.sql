-- +migrate Up
DROP TABLE IF EXISTS `admin_role_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_role_menus`
(
    `id`      int unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
    `menu_id` int unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_role_id_menu_id` (`role_id`, `menu_id`) USING BTREE,
    KEY `idx_menu_id` (`menu_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台角色关联菜单表';

INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (1, 1, 1);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (2, 1, 2);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (3, 1, 3);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (4, 1, 4);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (5, 1, 7);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (6, 1, 8);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (7, 1, 9);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (8, 1, 10);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (9, 1, 13);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (10, 1, 14);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (11, 1, 15);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (12, 1, 16);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (13, 1, 17);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (14, 1, 18);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (15, 1, 19);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (16, 1, 20);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (17, 1, 21);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (18, 1, 22);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (19, 1, 23);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (20, 1, 24);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (21, 1, 25);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (22, 1, 26);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (23, 1, 27);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (24, 1, 28);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (25, 1, 29);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (26, 1, 30);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (27, 1, 31);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (28, 1, 32);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (29, 1, 33);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (30, 1, 34);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (31, 1, 35);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (32, 1, 36);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (33, 1, 37);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (34, 1, 38);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (35, 1, 39);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (36, 1, 40);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (37, 1, 41);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (38, 1, 42);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (39, 1, 43);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (40, 1, 44);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (41, 1, 45);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (42, 1, 46);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (43, 1, 47);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (44, 1, 48);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (45, 1, 49);
INSERT INTO `admin_role_menus` (`id`, `role_id`, `menu_id`)
VALUES (46, 1, 50);

-- +migrate Down
DROP TABLE `admin_role_menus`;