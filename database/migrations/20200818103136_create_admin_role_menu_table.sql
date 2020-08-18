-- +migrate Up
DROP TABLE IF EXISTS `admin_role_menus`;
-- Create the table in the specified schema
CREATE TABLE `admin_role_menus`
(
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单id',
  KEY `idx_role_id` (`role_id`),
  KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='后台角色关联菜单表';

-- +migrate Down
DROP TABLE `admin_role_menus`;