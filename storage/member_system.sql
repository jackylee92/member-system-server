-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-01-10 21:56:17
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `member_system`
--

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

CREATE TABLE `role` (
  `id` int(10) NOT NULL COMMENT '主键',
  `code` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '角色编号',
  `name` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `remark` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '说明',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_flag` tinyint(4) NOT NULL DEFAULT '0' COMMENT '虚拟删除',
  `typ` tinyint(2) DEFAULT '0' COMMENT '权限类型 1: 默认权限'
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='角色表';

--
-- 转存表中的数据 `role`
--

INSERT INTO `role` (`id`, `code`, `name`, `status`, `remark`, `create_time`, `update_time`, `delete_flag`, `typ`) VALUES
(1, 'default', '默认权限', 1, '默认权限', '2023-01-06 15:52:35', '2023-01-06 15:52:38', 0, 1);

-- --------------------------------------------------------

--
-- 表的结构 `user_account`
--

CREATE TABLE `user_account` (
  `id` int(10) NOT NULL,
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `account` varchar(100) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '登录名',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态0:未知 1:启用 2:禁用',
  `password` varchar(150) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '密码'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;

--
-- 转存表中的数据 `user_account`
--

INSERT INTO `user_account` (`id`, `user_id`, `account`, `create_time`, `update_time`, `delete_flag`, `status`, `password`) VALUES
(4, 1, '123123123', '2023-01-06 15:46:09', '2023-01-06 15:46:09', 0, 1, 'adsas'),
(5, 2, '111111111', '2023-01-06 16:16:36', '2023-01-06 16:16:36', 0, 1, 'adsas'),
(6, 3, '222222', '2023-01-06 16:54:36', '2023-01-06 16:54:36', 0, 1, 'adsas'),
(7, 4, '333333', '2023-01-06 17:00:40', '2023-01-06 17:00:40', 0, 1, 'adsas'),
(8, 5, '4444444', '2023-01-06 17:02:24', '2023-01-06 17:02:24', 0, 1, 'adsas'),
(9, 6, '555555', '2023-01-06 17:23:17', '2023-01-06 17:23:17', 0, 1, 'adsas'),
(10, 7, '666666', '2023-01-06 19:10:04', '2023-01-06 19:10:04', 0, 1, 'adsasljd'),
(11, 8, '123123111', '2023-01-07 00:27:32', '2023-01-07 00:27:32', 0, 1, '123123123ljd'),
(12, 9, '1231231112', '2023-01-07 00:31:47', '2023-01-07 00:31:47', 0, 1, '123123123ljd'),
(13, 10, '1231231113', '2023-01-07 00:33:08', '2023-01-07 00:33:08', 0, 1, '123123123ljd'),
(14, 11, '1231231114', '2023-01-07 00:35:30', '2023-01-07 00:35:30', 0, 1, '123123123ljd'),
(15, 12, '1231231115', '2023-01-07 00:41:54', '2023-01-07 00:41:54', 0, 1, '123123123ljd'),
(16, 13, '1231231116', '2023-01-07 00:43:30', '2023-01-07 00:43:30', 0, 1, '123123123ljd'),
(17, 14, '1231231117', '2023-01-07 00:50:57', '2023-01-07 00:50:57', 0, 1, '123123123ljd'),
(18, 15, '1231231118', '2023-01-07 01:08:45', '2023-01-07 01:08:45', 0, 1, '123123123ljd'),
(19, 16, '1231231119', '2023-01-07 01:10:07', '2023-01-07 01:10:07', 0, 1, '123123123ljd'),
(20, 17, '1231231120', '2023-01-07 01:11:11', '2023-01-07 01:11:11', 0, 1, '123123123ljd'),
(21, 18, 'jackylee92@139.com', '2023-01-07 01:11:39', '2023-01-07 01:11:39', 0, 1, '123123123ljd');

-- --------------------------------------------------------

--
-- 表的结构 `user_attr`
--

CREATE TABLE `user_attr` (
  `id` int(10) NOT NULL COMMENT '主键',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `invitation_code` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户推荐码',
  `invitation_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '推荐的用户id',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0：未知 1：可用 2：禁用',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户附属表';

-- --------------------------------------------------------

--
-- 表的结构 `user_info`
--

CREATE TABLE `user_info` (
  `id` int(10) NOT NULL,
  `username` varchar(200) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户姓名',
  `create_time` datetime DEFAULT NULL COMMENT '录入时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_flag` tinyint(2) DEFAULT '0' COMMENT '虚拟删除',
  `status` tinyint(2) DEFAULT '0',
  `remark` varchar(255) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '备注',
  `introduction` text CHARACTER SET utf8mb4 COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '介绍',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '头像'
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='用户信息';

--
-- 转存表中的数据 `user_info`
--

INSERT INTO `user_info` (`id`, `username`, `create_time`, `update_time`, `delete_flag`, `status`, `remark`, `introduction`, `avatar`) VALUES
(1, '123123123', '2023-01-06 15:46:09', '2023-01-06 15:46:09', 0, 1, NULL, '', ''),
(2, 'username-111111111', '2023-01-06 16:16:36', '2023-01-06 16:16:36', 0, 1, NULL, '', ''),
(3, 'username-222222', '2023-01-06 16:54:36', '2023-01-06 16:54:36', 0, 1, NULL, '', ''),
(4, 'username-333333', '2023-01-06 17:00:40', '2023-01-06 17:00:40', 0, 1, NULL, '', ''),
(5, 'username-4444444', '2023-01-06 17:02:24', '2023-01-06 17:02:24', 0, 1, NULL, '', ''),
(6, 'username-555555', '2023-01-06 17:23:17', '2023-01-06 17:23:17', 0, 1, NULL, '', ''),
(7, 'username-666666', '2023-01-06 19:10:04', '2023-01-06 19:10:04', 0, 1, NULL, '', ''),
(8, 'username-123123111', '2023-01-07 00:27:32', '2023-01-07 00:27:32', 0, 1, NULL, '', ''),
(9, 'username-1231231112', '2023-01-07 00:31:47', '2023-01-07 00:31:47', NULL, 1, NULL, '', ''),
(10, 'username-1231231113', '2023-01-07 00:33:08', '2023-01-07 00:33:08', NULL, 1, NULL, '', ''),
(11, 'username-1231231114', '2023-01-07 00:35:30', '2023-01-07 00:35:30', NULL, 1, NULL, '', ''),
(12, 'username-1231231115', '2023-01-07 00:41:54', '2023-01-07 00:41:54', NULL, 1, NULL, '', ''),
(13, 'username-1231231116', '2023-01-07 00:43:30', '2023-01-07 00:43:30', NULL, 1, NULL, '', ''),
(14, 'username-1231231117', '2023-01-07 00:50:57', '2023-01-07 00:50:57', 0, 1, '', '', ''),
(15, 'username-1231231118', '2023-01-07 01:08:45', '2023-01-07 01:08:45', 0, 1, '', '', ''),
(16, 'username-1231231119', '2023-01-07 01:10:07', '2023-01-07 01:10:07', 0, 1, '', 'http://', ''),
(17, 'username-1231231120', '2023-01-07 01:11:11', '2023-01-07 01:11:11', 0, 1, '', 'http://', ''),
(18, 'username-1231231120', '2023-01-07 01:11:39', '2023-01-07 01:11:39', 0, 1, '', '这家伙很拽，啥都没说！', 'http://');

-- --------------------------------------------------------

--
-- 表的结构 `user_log`
--

CREATE TABLE `user_log` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `type` tinyint(255) NOT NULL DEFAULT '0' COMMENT '类型 0:未知 1:登录 2:登出',
  `content` text COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '提交数据',
  `action` varchar(30) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '操作',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '状态 0:未知 1:正常',
  `delete_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '虚拟删除 0:未删除 1:已删除',
  `remark` varchar(255) COLLATE utf8mb4_croatian_ci NOT NULL DEFAULT '' COMMENT '一句话说明'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci;

--
-- 转存表中的数据 `user_log`
--

INSERT INTO `user_log` (`id`, `user_id`, `type`, `content`, `action`, `create_time`, `update_time`, `status`, `delete_flag`, `remark`) VALUES
(85, 6, 3, '{\"Phone\":\"555555\",\"ValidCode\":\"7398\",\"ValidCodeID\":36,\"InvitationCode\":\"asdas\",\"Password\":\"adsas\",\"SurePassword\":\"adsas\"}', 'RegisterHandle', '2023-01-06 17:23:17', '2023-01-06 17:23:17', 1, 0, '注册成功'),
(86, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2023-01-06 18:58:21', '2023-01-06 18:58:21', 1, 0, '登录失败，用户名密码为空'),
(87, 0, 1, '{\"username\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 18:59:20', '2023-01-06 18:59:20', 1, 0, '登录失败，用户名密码为空'),
(88, 0, 1, '{\"username\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 19:00:10', '2023-01-06 19:00:10', 1, 0, '登录失败，用户名密码为空'),
(89, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2023-01-06 19:01:12', '2023-01-06 19:01:12', 1, 0, '登录失败，用户名密码为空'),
(90, 0, 1, '{\"Username\":\"ljd\",\"Password\":\"123\"}', 'LoginHandle', '2023-01-06 19:03:09', '2023-01-06 19:03:09', 1, 0, '登录失败，用户名密码为空'),
(91, 0, 1, '{\"username\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 19:03:30', '2023-01-06 19:03:30', 1, 0, '登录失败，用户名密码为空'),
(92, 0, 1, '{\"username\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 19:04:00', '2023-01-06 19:04:00', 1, 0, '登录失败，用户名密码为空'),
(93, 0, 1, '{\"username\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 19:05:15', '2023-01-06 19:05:15', 1, 0, '登录失败，用户名或者密码错误'),
(94, 0, 1, '{\"account\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-06 19:06:28', '2023-01-06 19:06:28', 1, 0, '登录失败，用户名或者密码错误'),
(95, 0, 1, '{\"account\":\"555555\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:06:58', '2023-01-06 19:06:58', 1, 0, '登录失败，用户名或者密码错误'),
(96, 0, 1, '{\"account\":\"555555\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:07:06', '2023-01-06 19:07:06', 1, 0, '登录失败，用户名或者密码错误'),
(97, 7, 3, '{\"Phone\":\"666666\",\"ValidCode\":\"8080\",\"ValidCodeID\":38,\"InvitationCode\":\"asdas\",\"Password\":\"adsas\",\"SurePassword\":\"adsas\"}', 'RegisterHandle', '2023-01-06 19:10:04', '2023-01-06 19:10:04', 1, 0, '注册成功'),
(98, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:10:28', '2023-01-06 19:10:28', 1, 0, '登录失败，用户名或者密码错误'),
(99, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:11:12', '2023-01-06 19:11:12', 1, 0, '登录失败，用户名或者密码错误'),
(100, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:11:50', '2023-01-06 19:11:50', 1, 0, '用户登录，登录成功，'),
(101, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:16:52', '2023-01-06 19:16:52', 1, 0, '用户登录，登录成功，'),
(102, 7, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:18:21', '2023-01-06 19:18:21', 1, 0, '用户登录，登录成功，'),
(103, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:23:17', '2023-01-06 19:23:17', 1, 0, '登录失败，用户信息数据不存在'),
(104, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:23:19', '2023-01-06 19:23:19', 1, 0, '登录失败，用户信息数据不存在'),
(105, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:23:50', '2023-01-06 19:23:50', 1, 0, '登录失败，用户信息数据不存在'),
(106, 0, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:24:40', '2023-01-06 19:24:40', 1, 0, '登录失败，用户信息数据不存在'),
(107, 7, 1, '{\"account\":\"666666\",\"password\":\"adsas\"}', 'LoginHandle', '2023-01-06 19:26:12', '2023-01-06 19:26:12', 1, 0, '用户登录，登录成功，username-666666'),
(108, 0, 1, '{\"account\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-07 00:02:24', '2023-01-07 00:02:24', 1, 0, '登录失败，用户名或者密码错误'),
(109, 8, 3, '{\"Phone\":\"123123111\",\"ValidCode\":\"1083\",\"ValidCodeID\":43,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:27:32', '2023-01-07 00:27:32', 1, 0, '注册成功'),
(110, 9, 3, '{\"Phone\":\"1231231112\",\"ValidCode\":\"4815\",\"ValidCodeID\":44,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:31:47', '2023-01-07 00:31:47', 1, 0, '注册成功'),
(111, 10, 3, '{\"Phone\":\"1231231113\",\"ValidCode\":\"4024\",\"ValidCodeID\":45,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:33:08', '2023-01-07 00:33:08', 1, 0, '注册成功'),
(112, 11, 3, '{\"Phone\":\"1231231114\",\"ValidCode\":\"7924\",\"ValidCodeID\":46,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:35:30', '2023-01-07 00:35:30', 1, 0, '注册成功'),
(113, 0, 1, '{\"account\":\"ljd\",\"password\":\"123\"}', 'LoginHandle', '2023-01-07 00:38:46', '2023-01-07 00:38:46', 1, 0, '登录失败，用户名或者密码错误'),
(114, 0, 1, '{\"account\":\"1231231114\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 00:39:09', '2023-01-07 00:39:09', 1, 0, '登录失败，用户信息数据不存在'),
(115, 12, 3, '{\"Phone\":\"1231231115\",\"ValidCode\":\"2090\",\"ValidCodeID\":47,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:41:54', '2023-01-07 00:41:54', 1, 0, '注册成功'),
(116, 13, 3, '{\"Phone\":\"1231231116\",\"ValidCode\":\"5591\",\"ValidCodeID\":48,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:43:30', '2023-01-07 00:43:30', 1, 0, '注册成功'),
(117, 14, 3, '{\"Phone\":\"1231231117\",\"ValidCode\":\"5258\",\"ValidCodeID\":50,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 00:50:57', '2023-01-07 00:50:57', 1, 0, '注册成功'),
(118, 14, 1, '{\"account\":\"1231231117\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 00:51:21', '2023-01-07 00:51:21', 1, 0, '用户登录，登录成功，username-1231231117'),
(119, 15, 3, '{\"Phone\":\"1231231118\",\"ValidCode\":\"1655\",\"ValidCodeID\":51,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 01:08:45', '2023-01-07 01:08:45', 1, 0, '注册成功'),
(120, 15, 1, '{\"account\":\"1231231118\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 01:09:00', '2023-01-07 01:09:00', 1, 0, '用户登录，登录成功，username-1231231118'),
(121, 16, 3, '{\"Phone\":\"1231231119\",\"ValidCode\":\"7349\",\"ValidCodeID\":52,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 01:10:07', '2023-01-07 01:10:07', 1, 0, '注册成功'),
(122, 17, 3, '{\"Phone\":\"1231231120\",\"ValidCode\":\"8939\",\"ValidCodeID\":53,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 01:11:11', '2023-01-07 01:11:11', 1, 0, '注册成功'),
(123, 18, 3, '{\"Phone\":\"1231231121\",\"ValidCode\":\"5352\",\"ValidCodeID\":54,\"InvitationCode\":\"123123123\",\"Password\":\"123123123\",\"SurePassword\":\"123123123\"}', 'RegisterHandle', '2023-01-07 01:11:39', '2023-01-07 01:11:39', 1, 0, '注册成功'),
(124, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 01:11:53', '2023-01-07 01:11:53', 1, 0, '用户登录，登录成功，username-1231231121'),
(125, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 01:24:24', '2023-01-07 01:24:24', 1, 0, '用户登录，登录成功，username-1231231121'),
(126, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 01:37:55', '2023-01-07 01:37:55', 1, 0, '用户登录，登录成功，123'),
(127, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 01:40:41', '2023-01-07 01:40:41', 1, 0, '用户登录，登录成功，username-1231231120'),
(128, 0, 1, '{\"account\":\"1231231114\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 16:38:29', '2023-01-07 16:38:29', 1, 0, '登录失败，用户信息数据不存在'),
(129, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 16:38:52', '2023-01-07 16:38:52', 1, 0, '用户登录，登录成功，username-1231231120'),
(130, 0, 1, '{\"account\":\"asdfasd\",\"password\":\"asdfasdf\"}', 'LoginHandle', '2023-01-07 19:52:01', '2023-01-07 19:52:01', 1, 0, '登录失败，用户名或者密码错误'),
(131, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 19:52:41', '2023-01-07 19:52:41', 1, 0, '用户登录，登录成功，username-1231231120'),
(132, 0, 1, '{\"account\":\"1231231121\",\"password\":\"1231231121\"}', 'LoginHandle', '2023-01-07 19:55:33', '2023-01-07 19:55:33', 1, 0, '登录失败，用户名或者密码错误'),
(133, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 19:56:51', '2023-01-07 19:56:51', 1, 0, '用户登录，登录成功，username-1231231120'),
(134, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 22:54:34', '2023-01-07 22:54:34', 1, 0, '用户登录，登录成功，username-1231231120'),
(135, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:00:23', '2023-01-07 23:00:23', 1, 0, '用户登录，登录成功，username-1231231120'),
(136, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:00:26', '2023-01-07 23:00:26', 1, 0, '用户登录，登录成功，username-1231231120'),
(137, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:01:55', '2023-01-07 23:01:55', 1, 0, '用户登录，登录成功，username-1231231120'),
(138, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:02:05', '2023-01-07 23:02:05', 1, 0, '用户登录，登录成功，username-1231231120'),
(139, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:07:07', '2023-01-07 23:07:07', 1, 0, '用户登录，登录成功，username-1231231120'),
(140, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:09:48', '2023-01-07 23:09:48', 1, 0, '用户登录，登录成功，username-1231231120'),
(141, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:09:52', '2023-01-07 23:09:52', 1, 0, '用户登录，登录成功，username-1231231120'),
(142, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:12:42', '2023-01-07 23:12:42', 1, 0, '用户登录，登录成功，username-1231231120'),
(143, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:27:45', '2023-01-07 23:27:45', 1, 0, '用户登录，登录成功，username-1231231120'),
(144, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:28:39', '2023-01-07 23:28:39', 1, 0, '用户登录，登录成功，username-1231231120'),
(145, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:28:42', '2023-01-07 23:28:42', 1, 0, '用户登录，登录成功，username-1231231120'),
(146, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:29:51', '2023-01-07 23:29:51', 1, 0, '用户登录，登录成功，username-1231231120'),
(147, 18, 2, '', 'LogoutHandle', '2023-01-07 23:31:17', '2023-01-07 23:31:17', 1, 0, '用户退出登录'),
(148, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:31:27', '2023-01-07 23:31:27', 1, 0, '用户登录，登录成功，username-1231231120'),
(149, 18, 2, '', 'LogoutHandle', '2023-01-07 23:32:38', '2023-01-07 23:32:38', 1, 0, '用户退出登录'),
(150, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:32:39', '2023-01-07 23:32:39', 1, 0, '用户登录，登录成功，username-1231231120'),
(151, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:53:59', '2023-01-07 23:53:59', 1, 0, '用户登录，登录成功，username-1231231120'),
(152, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:54:21', '2023-01-07 23:54:21', 1, 0, '用户登录，登录成功，username-1231231120'),
(153, 18, 1, '{\"account\":\"1231231121\",\"password\":\"123123123\"}', 'LoginHandle', '2023-01-07 23:54:46', '2023-01-07 23:54:46', 1, 0, '用户登录，登录成功，username-1231231120');

-- --------------------------------------------------------

--
-- 表的结构 `user_role`
--

CREATE TABLE `user_role` (
  `id` int(10) NOT NULL COMMENT '主键ID',
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `role_id` int(10) DEFAULT '0' COMMENT '角色ID',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态',
  `delete_flag` tinyint(2) DEFAULT '0' COMMENT '虚拟删除',
  `remark` varchar(200) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '备注',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='用户角色关系';

--
-- 转存表中的数据 `user_role`
--

INSERT INTO `user_role` (`id`, `user_id`, `role_id`, `update_time`, `status`, `delete_flag`, `remark`, `create_time`) VALUES
(1, 0, 1, '2023-01-06 16:54:36', 0, 0, NULL, '2023-01-06 16:54:36'),
(2, 4, 1, '2023-01-06 17:00:40', 1, 0, NULL, '2023-01-06 17:00:40'),
(3, 5, 1, '2023-01-06 17:02:24', 1, 0, NULL, '2023-01-06 17:02:24'),
(4, 6, 1, '2023-01-06 17:23:17', 1, 0, NULL, '2023-01-06 17:23:17'),
(5, 7, 1, '2023-01-06 19:10:04', 1, 0, NULL, '2023-01-06 19:10:04'),
(6, 8, 1, '2023-01-07 00:27:32', 1, 0, NULL, '2023-01-07 00:27:32'),
(7, 9, 1, '2023-01-07 00:31:47', 1, 0, NULL, '2023-01-07 00:31:47'),
(8, 10, 1, '2023-01-07 00:33:08', 1, 0, NULL, '2023-01-07 00:33:08'),
(9, 11, 1, '2023-01-07 00:35:30', 1, 0, NULL, '2023-01-07 00:35:30'),
(10, 12, 1, '2023-01-07 00:41:54', 1, 0, NULL, '2023-01-07 00:41:54'),
(11, 13, 1, '2023-01-07 00:43:30', 1, 0, NULL, '2023-01-07 00:43:30'),
(12, 14, 1, '2023-01-07 00:50:57', 1, 0, '', '2023-01-07 00:50:57'),
(13, 15, 1, '2023-01-07 01:08:45', 1, 0, '', '2023-01-07 01:08:45'),
(14, 16, 1, '2023-01-07 01:10:07', 1, 0, '', '2023-01-07 01:10:07'),
(15, 17, 1, '2023-01-07 01:11:11', 1, 0, '', '2023-01-07 01:11:11'),
(16, 18, 1, '2023-01-07 01:11:39', 1, 0, '', '2023-01-07 01:11:39');

-- --------------------------------------------------------

--
-- 表的结构 `valid_code`
--

CREATE TABLE `valid_code` (
  `id` int(10) NOT NULL COMMENT '主键',
  `code` varchar(10) COLLATE utf8mb4_croatian_ci NOT NULL COMMENT '验证码',
  `status` tinyint(2) DEFAULT '0' COMMENT '状态 0：未知 1：未使用 2：已使用 3：已失效',
  `user_id` int(10) DEFAULT '0' COMMENT '用户ID',
  `email` varchar(30) COLLATE utf8mb4_croatian_ci DEFAULT '' COMMENT '邮箱',
  `phone` varchar(20) COLLATE utf8mb4_croatian_ci DEFAULT '' COMMENT '手机号',
  `device_info` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '发送验证码的设备信息',
  `msg_type` tinyint(2) DEFAULT '0' COMMENT '消息类型 0：未知 1：注册code 2：找回密码',
  `msg` varchar(255) COLLATE utf8mb4_croatian_ci DEFAULT '' COMMENT '消息内容',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `delete_flag` tinyint(2) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_croatian_ci COMMENT='验证码记录表';

--
-- 转存表中的数据 `valid_code`
--

INSERT INTO `valid_code` (`id`, `code`, `status`, `user_id`, `email`, `phone`, `device_info`, `msg_type`, `msg`, `update_time`, `expire_time`, `create_time`, `delete_flag`) VALUES
(30, '5613', 2, 0, NULL, '123123123', '', 1, '你的验证码是5613，请勿泄露。', '2023-01-06 14:16:08', '2023-02-06 14:17:07', '2023-01-06 14:16:08', 0),
(31, '7382', 1, 0, NULL, '1231231232', '', 1, '你的验证码是7382，请勿泄露。', '2023-01-06 16:13:48', '2023-01-06 16:14:48', '2023-01-06 16:13:48', 0),
(32, '4180', 2, 0, NULL, '111111111', '', 1, '你的验证码是4180，请勿泄露。', '2023-01-06 16:14:27', '2023-02-06 16:15:27', '2023-01-06 16:14:27', 0),
(33, '2491', 2, 0, NULL, '222222', '', 1, '你的验证码是2491，请勿泄露。', '2023-01-06 16:51:57', '2023-02-06 16:52:57', '2023-01-06 16:51:57', 0),
(34, '6187', 2, 0, NULL, '333333', '', 1, '你的验证码是6187，请勿泄露。', '2023-01-06 17:00:23', '2023-01-06 17:01:23', '2023-01-06 17:00:23', 0),
(35, '8971', 2, 0, NULL, '4444444', '', 1, '你的验证码是8971，请勿泄露。', '2023-01-06 17:02:24', '2023-01-06 17:03:03', '2023-01-06 17:02:04', 0),
(36, '7398', 2, 0, NULL, '555555', '', 1, '你的验证码是7398，请勿泄露。', '2023-01-06 17:23:17', '2023-01-06 17:24:02', '2023-01-06 17:23:03', 0),
(37, '3883', 1, 0, NULL, '555555', '', 1, '你的验证码是3883，请勿泄露。', '2023-01-06 18:56:04', '2023-01-06 18:57:03', '2023-01-06 18:56:04', 0),
(38, '8080', 2, 0, NULL, '666666', '', 1, '你的验证码是8080，请勿泄露。', '2023-01-06 19:10:04', '2023-01-06 19:10:47', '2023-01-06 19:09:47', 0),
(39, '5697', 1, 0, NULL, '123123111', '', 1, '你的验证码是5697，请勿泄露。', '2023-01-07 00:05:06', '2023-01-07 00:06:05', '2023-01-07 00:05:06', 0),
(40, '9028', 1, 0, NULL, '123123111', '', 1, '你的验证码是9028，请勿泄露。', '2023-01-07 00:05:12', '2023-01-07 00:06:12', '2023-01-07 00:05:12', 0),
(41, '6967', 1, 0, NULL, '123123111', '', 1, '你的验证码是6967，请勿泄露。', '2023-01-07 00:07:11', '2023-01-07 00:08:10', '2023-01-07 00:07:11', 0),
(42, '7453', 1, 0, NULL, '123123111', '', 1, '你的验证码是7453，请勿泄露。', '2023-01-07 00:07:38', '2023-01-07 00:08:38', '2023-01-07 00:07:38', 2),
(43, '1083', 2, 0, NULL, '123123111', '', 1, '你的验证码是1083，请勿泄露。', '2023-01-07 00:27:32', '2023-01-07 00:28:00', '2023-01-07 00:27:00', 0),
(44, '4815', 2, 0, NULL, '1231231112', '', 1, '你的验证码是4815，请勿泄露。', '2023-01-07 00:31:23', '2023-01-07 00:32:23', '2023-01-07 00:31:23', 0),
(45, '4024', 2, 0, NULL, '1231231113', '', 1, '你的验证码是4024，请勿泄露。', '2023-01-07 00:32:56', '2023-01-07 00:33:56', '2023-01-07 00:32:56', 0),
(46, '7924', 2, 0, NULL, '1231231114', '', 1, '你的验证码是7924，请勿泄露。', '2023-01-07 00:35:14', '2023-01-07 00:36:13', '2023-01-07 00:35:14', 0),
(47, '2090', 2, 0, NULL, '1231231115', '', 1, '你的验证码是2090，请勿泄露。', '2023-01-07 00:41:54', '2023-01-07 00:42:18', '2023-01-07 00:41:19', 0),
(48, '5591', 2, 0, NULL, '1231231116', '', 1, '你的验证码是5591，请勿泄露。', '2023-01-07 00:43:30', '2023-01-07 00:44:15', '2023-01-07 00:43:16', 0),
(49, '6359', 1, 0, '', '1231231116', '', 1, '你的验证码是6359，请勿泄露。', '2023-01-07 00:50:29', '2023-01-07 00:51:28', '2023-01-07 00:50:29', 0),
(50, '5258', 2, 0, '', '1231231117', '', 1, '你的验证码是5258，请勿泄露。', '2023-01-07 00:50:57', '2023-01-07 00:51:32', '2023-01-07 00:50:33', 0),
(51, '1655', 2, 0, '', '1231231118', '', 1, '你的验证码是1655，请勿泄露。', '2023-01-07 01:08:45', '2023-01-07 01:09:28', '2023-01-07 01:08:28', 0),
(52, '7349', 2, 0, '', '1231231119', '', 1, '你的验证码是7349，请勿泄露。', '2023-01-07 01:10:07', '2023-01-07 01:10:55', '2023-01-07 01:09:55', 0),
(53, '8939', 2, 0, '', '1231231120', '', 1, '你的验证码是8939，请勿泄露。', '2023-01-07 01:11:11', '2023-01-07 01:11:58', '2023-01-07 01:10:59', 0),
(54, '5352', 2, 0, '', '1231231121', '', 1, '你的验证码是5352，请勿泄露。', '2023-01-07 01:11:39', '2023-01-07 01:12:26', '2023-01-07 01:11:27', 0),
(55, '7484', 1, 0, 'jackylee92@139.com', '', '', 2, '你的验证码是7484，请勿泄露。', '2023-01-09 23:06:42', '2023-01-09 23:07:40', '2023-01-09 23:06:42', 0),
(56, '1112', 1, 0, 'jackylee92@139.com', '', '', 2, '你的验证码是1112，请勿泄露。', '2023-01-09 23:10:49', '2023-01-09 23:11:47', '2023-01-09 23:10:49', 0),
(57, '1106', 2, 0, 'jackylee92@139.com', '', '', 2, '你的验证码是1106，请勿泄露。', '2023-01-09 23:23:25', '2023-01-09 23:33:08', '2023-01-09 23:23:10', 0),
(58, '1415', 2, 18, 'jackylee92@139.com', '', '', 2, '你的验证码是1415，请勿泄露。', '2023-01-09 23:56:36', '2023-01-10 00:05:22', '2023-01-09 23:55:23', 0);

--
-- 转储表的索引
--

--
-- 表的索引 `role`
--
ALTER TABLE `role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_id_uindex` (`id`);

--
-- 表的索引 `user_account`
--
ALTER TABLE `user_account`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_attr`
--
ALTER TABLE `user_attr`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_info`
--
ALTER TABLE `user_info`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_log`
--
ALTER TABLE `user_log`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_role`
--
ALTER TABLE `user_role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_role_id_uindex` (`id`);

--
-- 表的索引 `valid_code`
--
ALTER TABLE `valid_code`
  ADD PRIMARY KEY (`id`),
  ADD KEY `index_code` (`code`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `role`
--
ALTER TABLE `role`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `user_account`
--
ALTER TABLE `user_account`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- 使用表AUTO_INCREMENT `user_attr`
--
ALTER TABLE `user_attr`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `user_info`
--
ALTER TABLE `user_info`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- 使用表AUTO_INCREMENT `user_log`
--
ALTER TABLE `user_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=154;

--
-- 使用表AUTO_INCREMENT `user_role`
--
ALTER TABLE `user_role`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=17;

--
-- 使用表AUTO_INCREMENT `valid_code`
--
ALTER TABLE `valid_code`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=59;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
