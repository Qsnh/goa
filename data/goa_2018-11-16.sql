/*
Navicat MySQL Data Transfer

Source Server         : mysql5.6
Source Server Version : 50641
Source Host           : localhost:33060
Source Database       : goa

Target Server Type    : MYSQL
Target Server Version : 50641
File Encoding         : 65001

Date: 2018-11-16 11:33:02
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for answers
-- ----------------------------
DROP TABLE IF EXISTS `answers`;
CREATE TABLE `answers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of answers
-- ----------------------------

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '分类名',
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for questions
-- ----------------------------
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `category_id` int(11) unsigned NOT NULL COMMENT '分类ID',
  `answer_user_id` int(11) DEFAULT NULL,
  `title` varchar(255) NOT NULL COMMENT '问题标题',
  `description` text NOT NULL COMMENT '问题描述',
  `view_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `is_ban` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '1禁止,-1未禁止',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `answer_at` datetime NOT NULL,
  `answer_count` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for settings
-- ----------------------------
DROP TABLE IF EXISTS `settings`;
CREATE TABLE `settings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of settings
-- ----------------------------
INSERT INTO `settings` VALUES ('1', 'MEMBER_DEFAULT_AVATAR', '/static/avatar.png');
INSERT INTO `settings` VALUES ('2', 'MEMBER_DEFAULT_IS_LOCK', '1');
INSERT INTO `settings` VALUES ('3', 'SEO_INDEX_TITLE', 'goa是基于beego开发的在线问答系统');
INSERT INTO `settings` VALUES ('4', 'SEO_INDEX_KEYWORDS', 'goa,go,beego,npm,vuejs,env');
INSERT INTO `settings` VALUES ('5', 'SEO_INDEX_DESCRIPTION', 'goa是基于beego开发的在线问答系统');
INSERT INTO `settings` VALUES ('6', 'ICP', '暂无');
INSERT INTO `settings` VALUES ('7', 'CORS_ORIGINAL', 'http://127.0.0.1:8080');
INSERT INTO `settings` VALUES ('8', 'APP_NAME', 'GOA');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(32) NOT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `password` varchar(256) NOT NULL COMMENT '密码',
  `is_lock` tinyint(1) NOT NULL COMMENT '1锁定,-1未锁定',
  `company` varchar(255) DEFAULT NULL COMMENT '公司',
  `age` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `profession` varchar(255) DEFAULT NULL COMMENT '职业',
  `weibo` varchar(255) DEFAULT NULL COMMENT '微博',
  `wechat` varchar(255) DEFAULT NULL COMMENT '微信',
  `website` varchar(255) DEFAULT NULL COMMENT '个人主页',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;