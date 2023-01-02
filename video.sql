/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3306
 Source Schema         : video

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 15/09/2022 13:39:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(150) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '密码',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `nickname` varchar(255) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态0:未知 1:启用 2:禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;

-- ----------------------------
-- Records of user_account
-- ----------------------------
BEGIN;
INSERT INTO `user_account` VALUES (1, 'ljd', '123ljd', '2022-09-14 17:58:33', '2022-09-14 17:58:35', 'kkk', 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for user_log
-- ----------------------------
DROP TABLE IF EXISTS `user_log`;
CREATE TABLE `user_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type` tinyint(255) NOT NULL DEFAULT '0' COMMENT '类型 0:未知 1:登录',
  `content` text COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '提交数据',
  `action` varchar(30) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '操作',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0:未知 1:正常',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `remark` varchar(255) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '一句话说明',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;

-- ----------------------------
-- Records of user_log
-- ----------------------------
BEGIN;
INSERT INTO `user_log` VALUES (1, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2022-09-14 21:59:36', '2022-09-14 21:59:36', 1, 0, '用户已登录，ljd');
INSERT INTO `user_log` VALUES (2, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2022-09-14 21:59:44', '2022-09-14 21:59:44', 1, 0, '用户已登录，ljd');
INSERT INTO `user_log` VALUES (3, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2022-09-15 11:20:33', '2022-09-15 11:20:33', 1, 0, '用户已登录，ljd');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
