/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50641
Source Host           : 127.0.0.1:33060
Source Database       : go

Target Server Type    : MYSQL
Target Server Version : 50641
File Encoding         : 65001

Date: 2018-08-16 15:49:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for questions
-- ----------------------------
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions` (
  `id` int(11) unsigned NOT NULL,
  `category_id` int(11) unsigned NOT NULL COMMENT '分类ID',
  `title` varchar(255) NOT NULL COMMENT '问题标题',
  `description` text NOT NULL COMMENT '问题描述',
  `view_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `is_ban` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '1禁止,-1未禁止',
  `created_at` int(11) unsigned NOT NULL,
  `updated_at` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
