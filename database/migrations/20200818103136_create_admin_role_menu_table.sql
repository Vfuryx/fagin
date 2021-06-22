-- +migrate Up
DROP TABLE IF EXISTS `admin_role_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_role_menus`
(
    `id`      int(11) unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
    `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_role_id_menu_id` (`role_id`, `menu_id`) USING BTREE,
    KEY `idx_menu_id` (`menu_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台角色关联菜单表';

INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (1, 1, 1);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (2, 1, 2);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (3, 1, 3);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (4, 1, 4);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (5, 1, 5);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (6, 1, 6);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (7, 1, 7);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (8, 1, 8);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (9, 1, 9);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (10, 1, 10);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (11, 1, 11);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (12, 1, 12);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (13, 1, 13);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (14, 1, 14);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (15, 1, 15);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (16, 1, 16);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (17, 1, 17);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (18, 1, 18);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (19, 1, 19);
INSERT INTO `admin_role_menus`(`id`, `role_id`, `menu_id`) VALUES (20, 1, 20);

-- +migrate Down
DROP TABLE `admin_role_menus`;