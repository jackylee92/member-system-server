#
# SQL Export
# Created by Querious (303012)
# Created: 2023-01-06 at 23:42:55
# Encoding: Unicode (UTF-8)
#


SET @ORIG_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;

SET @ORIG_UNIQUE_CHECKS = @@UNIQUE_CHECKS;
SET UNIQUE_CHECKS = 0;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = '+00:00';

SET @ORIG_SQL_MODE = @@SQL_MODE;
SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO';



DROP TABLE IF EXISTS `valid_code`;
DROP TABLE IF EXISTS `user_role`;
DROP TABLE IF EXISTS `user_log`;
DROP TABLE IF EXISTS `user_info`;
DROP TABLE IF EXISTS `user_account`;
DROP TABLE IF EXISTS `role`;


CREATE TABLE `role` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '角色编号',
  `name` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `remark` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '说明',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '虚拟删除',
  `typ` tinyint(2) DEFAULT '0' COMMENT '权限类型 1: 默认权限',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COMMENT='角色表';


CREATE TABLE `user_account` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `account` varchar(100) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '登录名',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态0:未知 1:启用 2:禁用',
  `password` varchar(150) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;


CREATE TABLE `user_info` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(200) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户姓名',
  `create_time` datetime DEFAULT NULL COMMENT '录入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_flag` tinyint(2) DEFAULT NULL COMMENT '虚拟删除',
  `status` tinyint(2) DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1 COMMENT='用户信息';


CREATE TABLE `user_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type` tinyint(255) NOT NULL DEFAULT '0' COMMENT '类型 0:未知 1:登录 2:登出',
  `content` text COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '提交数据',
  `action` varchar(30) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '操作',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0:未知 1:正常',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `remark` varchar(255) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '一句话说明',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;


CREATE TABLE `user_role` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `role_id` int(10) DEFAULT '0' COMMENT '角色ID',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态',
  `delete_flag` tinyint(2) DEFAULT '0' COMMENT '虚拟删除',
  `remark` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_role_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COMMENT='用户角色关系';


CREATE TABLE `valid_code` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(10) COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '验证码',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态 0：未知 1：未使用 2：已使用 3：已失效',
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `email` varchar(30) COLLATE utf8mb4_croatian_ci DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) COLLATE utf8mb4_croatian_ci DEFAULT NULL COMMENT '手机号',
  `msg_type` tinyint(2) DEFAULT '0' COMMENT '消息类型 0：未知 1：注册code',
  `msg` varchar(255) COLLATE utf8mb4_croatian_ci DEFAULT NULL COMMENT '消息内容',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `delete_flag` tinyint(2) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `index_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci COMMENT='验证码记录表';




LOCK TABLES `role` WRITE;
INSERT INTO `role` (`id`, `code`, `name`, `status`, `remark`, `create_time`, `update_time`, `delete_flag`, `typ`) VALUES 
	(1,'default','默认权限',1,'默认权限','2023-01-06 15:52:35','2023-01-06 15:52:38',0,1);
UNLOCK TABLES;


LOCK TABLES `user_account` WRITE;
INSERT INTO `user_account` (`id`, `user_id`, `account`, `create_time`, `update_time`, `delete_flag`, `status`, `password`) VALUES 
	(4,1,'123123123','2023-01-06 15:46:09','2023-01-06 15:46:09',0,1,'adsas'),
	(5,2,'111111111','2023-01-06 16:16:36','2023-01-06 16:16:36',0,1,'adsas'),
	(6,3,'222222','2023-01-06 16:54:36','2023-01-06 16:54:36',0,1,'adsas'),
	(7,4,'333333','2023-01-06 17:00:40','2023-01-06 17:00:40',0,1,'adsas'),
	(8,5,'4444444','2023-01-06 17:02:24','2023-01-06 17:02:24',0,1,'adsas'),
	(9,6,'555555','2023-01-06 17:23:17','2023-01-06 17:23:17',0,1,'adsas'),
	(10,7,'666666','2023-01-06 19:10:04','2023-01-06 19:10:04',0,1,'adsasljd');
UNLOCK TABLES;


LOCK TABLES `user_info` WRITE;
INSERT INTO `user_info` (`id`, `username`, `create_time`, `update_time`, `delete_flag`, `status`, `remark`) VALUES 
	(1,'123123123','2023-01-06 15:46:09','2023-01-06 15:46:09',0,1,NULL),
	(2,'username-111111111','2023-01-06 16:16:36','2023-01-06 16:16:36',0,1,NULL),
	(3,'username-222222','2023-01-06 16:54:36','2023-01-06 16:54:36',0,1,NULL),
	(4,'username-333333','2023-01-06 17:00:40','2023-01-06 17:00:40',0,1,NULL),
	(5,'username-4444444','2023-01-06 17:02:24','2023-01-06 17:02:24',0,1,NULL),
	(6,'username-555555','2023-01-06 17:23:17','2023-01-06 17:23:17',0,1,NULL),
	(7,'username-666666','2023-01-06 19:10:04','2023-01-06 19:10:04',0,1,NULL);
UNLOCK TABLES;


LOCK TABLES `user_log` WRITE;
INSERT INTO `user_log` (`id`, `user_id`, `type`, `content`, `action`, `create_time`, `update_time`, `status`, `delete_flag`, `remark`) VALUES 
	(85,6,3,'{"Phone":"555555","ValidCode":"7398","ValidCodeID":36,"InvitationCode":"asdas","Password":"adsas","SurePassword":"adsas"}','RegisterHandle','2023-01-06 17:23:17','2023-01-06 17:23:17',1,0,'注册成功'),
	(86,0,1,'{"Username":"ljd","Password":"123"}','LoginHandle','2023-01-06 18:58:21','2023-01-06 18:58:21',1,0,'登录失败，用户名密码为空'),
	(87,0,1,'{"username":"ljd","password":"123"}','LoginHandle','2023-01-06 18:59:20','2023-01-06 18:59:20',1,0,'登录失败，用户名密码为空'),
	(88,0,1,'{"username":"ljd","password":"123"}','LoginHandle','2023-01-06 19:00:10','2023-01-06 19:00:10',1,0,'登录失败，用户名密码为空'),
	(89,0,1,'{"Username":"ljd","Password":"123"}','LoginHandle','2023-01-06 19:01:12','2023-01-06 19:01:12',1,0,'登录失败，用户名密码为空'),
	(90,0,1,'{"Username":"ljd","Password":"123"}','LoginHandle','2023-01-06 19:03:09','2023-01-06 19:03:09',1,0,'登录失败，用户名密码为空'),
	(91,0,1,'{"username":"ljd","password":"123"}','LoginHandle','2023-01-06 19:03:30','2023-01-06 19:03:30',1,0,'登录失败，用户名密码为空'),
	(92,0,1,'{"username":"ljd","password":"123"}','LoginHandle','2023-01-06 19:04:00','2023-01-06 19:04:00',1,0,'登录失败，用户名密码为空'),
	(93,0,1,'{"username":"ljd","password":"123"}','LoginHandle','2023-01-06 19:05:15','2023-01-06 19:05:15',1,0,'登录失败，用户名或者密码错误'),
	(94,0,1,'{"account":"ljd","password":"123"}','LoginHandle','2023-01-06 19:06:28','2023-01-06 19:06:28',1,0,'登录失败，用户名或者密码错误'),
	(95,0,1,'{"account":"555555","password":"adsas"}','LoginHandle','2023-01-06 19:06:58','2023-01-06 19:06:58',1,0,'登录失败，用户名或者密码错误'),
	(96,0,1,'{"account":"555555","password":"adsas"}','LoginHandle','2023-01-06 19:07:06','2023-01-06 19:07:06',1,0,'登录失败，用户名或者密码错误'),
	(97,7,3,'{"Phone":"666666","ValidCode":"8080","ValidCodeID":38,"InvitationCode":"asdas","Password":"adsas","SurePassword":"adsas"}','RegisterHandle','2023-01-06 19:10:04','2023-01-06 19:10:04',1,0,'注册成功'),
	(98,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:10:28','2023-01-06 19:10:28',1,0,'登录失败，用户名或者密码错误'),
	(99,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:11:12','2023-01-06 19:11:12',1,0,'登录失败，用户名或者密码错误'),
	(100,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:11:50','2023-01-06 19:11:50',1,0,'用户登录，登录成功，'),
	(101,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:16:52','2023-01-06 19:16:52',1,0,'用户登录，登录成功，'),
	(102,7,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:18:21','2023-01-06 19:18:21',1,0,'用户登录，登录成功，'),
	(103,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:23:17','2023-01-06 19:23:17',1,0,'登录失败，用户信息数据不存在'),
	(104,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:23:19','2023-01-06 19:23:19',1,0,'登录失败，用户信息数据不存在'),
	(105,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:23:50','2023-01-06 19:23:50',1,0,'登录失败，用户信息数据不存在'),
	(106,0,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:24:40','2023-01-06 19:24:40',1,0,'登录失败，用户信息数据不存在'),
	(107,7,1,'{"account":"666666","password":"adsas"}','LoginHandle','2023-01-06 19:26:12','2023-01-06 19:26:12',1,0,'用户登录，登录成功，username-666666');
UNLOCK TABLES;


LOCK TABLES `user_role` WRITE;
INSERT INTO `user_role` (`id`, `user_id`, `role_id`, `update_time`, `status`, `delete_flag`, `remark`, `create_time`) VALUES 
	(1,0,1,'2023-01-06 16:54:36',0,0,NULL,'2023-01-06 16:54:36'),
	(2,4,1,'2023-01-06 17:00:40',1,0,NULL,'2023-01-06 17:00:40'),
	(3,5,1,'2023-01-06 17:02:24',1,0,NULL,'2023-01-06 17:02:24'),
	(4,6,1,'2023-01-06 17:23:17',1,0,NULL,'2023-01-06 17:23:17'),
	(5,7,1,'2023-01-06 19:10:04',1,0,NULL,'2023-01-06 19:10:04');
UNLOCK TABLES;


LOCK TABLES `valid_code` WRITE;
INSERT INTO `valid_code` (`id`, `code`, `status`, `user_id`, `email`, `phone`, `msg_type`, `msg`, `update_time`, `expire_time`, `create_time`, `delete_flag`) VALUES 
	(30,'5613',2,0,NULL,'123123123',1,'你的验证码是5613，请勿泄露。','2023-01-06 14:16:08','2023-02-06 14:17:07','2023-01-06 14:16:08',0),
	(31,'7382',1,0,NULL,'1231231232',1,'你的验证码是7382，请勿泄露。','2023-01-06 16:13:48','2023-01-06 16:14:48','2023-01-06 16:13:48',0),
	(32,'4180',2,0,NULL,'111111111',1,'你的验证码是4180，请勿泄露。','2023-01-06 16:14:27','2023-02-06 16:15:27','2023-01-06 16:14:27',0),
	(33,'2491',2,0,NULL,'222222',1,'你的验证码是2491，请勿泄露。','2023-01-06 16:51:57','2023-02-06 16:52:57','2023-01-06 16:51:57',0),
	(34,'6187',2,0,NULL,'333333',1,'你的验证码是6187，请勿泄露。','2023-01-06 17:00:23','2023-01-06 17:01:23','2023-01-06 17:00:23',0),
	(35,'8971',2,0,NULL,'4444444',1,'你的验证码是8971，请勿泄露。','2023-01-06 17:02:24','2023-01-06 17:03:03','2023-01-06 17:02:04',0),
	(36,'7398',2,0,NULL,'555555',1,'你的验证码是7398，请勿泄露。','2023-01-06 17:23:17','2023-01-06 17:24:02','2023-01-06 17:23:03',0),
	(37,'3883',1,0,NULL,'555555',1,'你的验证码是3883，请勿泄露。','2023-01-06 18:56:04','2023-01-06 18:57:03','2023-01-06 18:56:04',0),
	(38,'8080',2,0,NULL,'666666',1,'你的验证码是8080，请勿泄露。','2023-01-06 19:10:04','2023-01-06 19:10:47','2023-01-06 19:09:47',0);
UNLOCK TABLES;






SET FOREIGN_KEY_CHECKS = @ORIG_FOREIGN_KEY_CHECKS;

SET UNIQUE_CHECKS = @ORIG_UNIQUE_CHECKS;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = @ORIG_TIME_ZONE;

SET SQL_MODE = @ORIG_SQL_MODE;



# Export Finished: 2023-01-06 at 23:42:55

