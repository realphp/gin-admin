CREATE TABLE `ga_users` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	`user_name` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`password` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`nick_name` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`phone` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`role_id` INT NOT NULL DEFAULT 0,
	`is_del` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY ( `id` ) USING BTREE
	
) ENGINE = INNODB DEFAULT CHARSET = utf8;

INSERT INTO `ga_users` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'admin','',1,0,NULL,NULL);

CREATE TABLE `ga_roles` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	`name` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`parent_id` INT NOT NULL DEFAULT 0,
	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	PRIMARY KEY ( `id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE `ga_role_menus` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	`role_id` INT UNSIGNED NOT NULL DEFAULT 0,
	`menu_id` INT UNSIGNED NOT NULL DEFAULT 0,
	PRIMARY KEY ( `id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE `ga_menus` (
	`id` INT  UNSIGNED NOT NULL AUTO_INCREMENT,
	`menu_level` INT  UNSIGNED DEFAULT 0,
	`parent_id` INT  UNSIGNED DEFAULT 0,
	`path` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`name` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`hidden` TINYINT  NOT NULL DEFAULT 0,
	`component` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`title` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`icon` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`nick_name` VARCHAR ( 255 ) NOT NULL DEFAULT '',
	`sort` TINYINT NOT NULL DEFAULT 0,
	`keep_alive` TINYINT  NOT NULL DEFAULT 0,
	`default_menu` TINYINT  NOT NULL DEFAULT 0,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
	PRIMARY KEY ( `id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

INSERT INTO `ga_menus` VALUES (1,  0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', '仪表盘', 'setting', '仪表盘', 1, 0, 0,NULL,NULL);
INSERT INTO `ga_menus` VALUES (2,  0, 0, 'about', 'about', 0, 'view/about/index.vue', '关于我们', 'info', '测试菜单', 7, 0, 0,NULL,NULL);
INSERT INTO `ga_menus` VALUES (3,  0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', '超级管理员', 3, 0, 0,NULL,NULL);
INSERT INTO `ga_menus` VALUES (4,  0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', '角色管理', 1, 0, 0,NULL,NULL);
INSERT INTO `ga_menus` VALUES (5,  0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', '菜单管理', 2, 1, 0,NULL,NULL);
INSERT INTO `ga_menus` VALUES (6,  0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 'api管理', 3, 1, 0,NULL,NULL);
