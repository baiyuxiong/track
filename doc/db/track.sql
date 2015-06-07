-- phpMyAdmin SQL Dump
-- version 4.2.2
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: 2015-06-08 06:35:43
-- 服务器版本： 5.6.24
-- PHP Version: 5.5.24

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `track`
--

-- --------------------------------------------------------

--
-- 表的结构 `company`
--

CREATE TABLE IF NOT EXISTS `company` (
`id` int(11) NOT NULL,
  `owner_id` int(11) NOT NULL COMMENT '所有者',
  `name` varchar(128) NOT NULL,
  `info` text NOT NULL,
  `logo` varchar(256) NOT NULL,
  `phone` varchar(12) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `company`
--

INSERT INTO `company` (`id`, `owner_id`, `name`, `info`, `logo`, `phone`, `address`, `created_at`, `updated_at`) VALUES
(1, 2, 'track cmpa', '', 'abcdef', '', '', '2015-06-07 21:22:01', '2015-06-07 21:22:01');

-- --------------------------------------------------------

--
-- 表的结构 `company_users`
--

CREATE TABLE IF NOT EXISTS `company_users` (
  `company_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL COMMENT '0 未审核 1已审核 2 删除',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='单位成员列表';

-- --------------------------------------------------------

--
-- 表的结构 `news`
--

CREATE TABLE IF NOT EXISTS `news` (
`id` int(10) unsigned NOT NULL,
  `company_id` int(11) NOT NULL COMMENT '企业ID',
  `project_id` int(11) NOT NULL COMMENT '项目ID',
  `owner_id` int(11) NOT NULL COMMENT '作者',
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `content` text COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `project`
--

CREATE TABLE IF NOT EXISTS `project` (
`id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `owner_id` int(11) NOT NULL,
  `name` varchar(256) NOT NULL,
  `info` text NOT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `sms_code`
--

CREATE TABLE IF NOT EXISTS `sms_code` (
  `username` char(11) NOT NULL,
  `code` char(6) NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='短信验证码';

--
-- 转存表中的数据 `sms_code`
--

INSERT INTO `sms_code` (`username`, `code`, `updated_at`) VALUES
('13456789012', '021558', '2015-06-07 21:22:01');

-- --------------------------------------------------------

--
-- 表的结构 `task`
--

CREATE TABLE IF NOT EXISTS `task` (
`id` int(11) NOT NULL,
  `company_id` int(11) NOT NULL,
  `project_id` int(11) NOT NULL,
  `owner_id` int(11) NOT NULL,
  `priority` tinyint(4) NOT NULL COMMENT '优先级 1低 2中 3高',
  `status` tinyint(4) NOT NULL COMMENT '状态 0初始 1处理中 2 结束',
  `name` varchar(256) NOT NULL,
  `info` text NOT NULL,
  `deadline` datetime NOT NULL COMMENT '截止日期',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='任务' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `task_chat_log`
--

CREATE TABLE IF NOT EXISTS `task_chat_log` (
  `task_id` bigint(11) NOT NULL,
  `log` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='聊天日志';

-- --------------------------------------------------------

--
-- 表的结构 `task_transfer`
--

CREATE TABLE IF NOT EXISTS `task_transfer` (
`id` bigint(20) NOT NULL,
  `company_id` int(11) NOT NULL,
  `project_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `in_charge_id` int(11) NOT NULL COMMENT '当前谁负责',
  `assign_to` int(11) DEFAULT NULL COMMENT '指派给谁',
  `progress` tinyint(4) NOT NULL COMMENT '进度百分比',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='任务流转' AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE IF NOT EXISTS `users` (
`id` int(11) unsigned NOT NULL,
  `ip_address` varchar(15) NOT NULL,
  `username` char(11) NOT NULL,
  `password` varchar(80) NOT NULL,
  `salt` char(16) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  `token` varchar(32) DEFAULT NULL,
  `is_activited` tinyint(4) NOT NULL COMMENT '1 已激活',
  `activated_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `ip_address`, `username`, `password`, `salt`, `email`, `token`, `is_activited`, `activated_at`, `created_at`, `updated_at`) VALUES
(1, '127.0.0.1', '13456789012', '53d4d4a9c648d4d614476f193ee03e8c', 'ihbsEJ8csJm1Ckgh', '', 'bdb7c99ffde63231dca586da4563604a', 1, '2015-06-07 21:22:01', '2015-06-07 21:22:01', '2015-06-07 21:22:01');

-- --------------------------------------------------------

--
-- 表的结构 `user_profiles`
--

CREATE TABLE IF NOT EXISTS `user_profiles` (
  `user_id` int(11) NOT NULL,
  `gender` tinyint(4) NOT NULL COMMENT '1 男 2 女',
  `name` varchar(32) NOT NULL,
  `avatar` varchar(256) NOT NULL,
  `avatar_thumb1` varchar(256) NOT NULL,
  `avatar_thumb2` varchar(256) NOT NULL,
  `avatar_thumb3` varchar(256) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `user_profiles`
--

INSERT INTO `user_profiles` (`user_id`, `gender`, `name`, `avatar`, `avatar_thumb1`, `avatar_thumb2`, `avatar_thumb3`, `created_at`, `updated_at`) VALUES
(1, 1, 'test', '', '', '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(2, 1, 'test2', '', '', '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(3, 1, 'test3', '', '', '', '', '0000-00-00 00:00:00', '0000-00-00 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `company`
--
ALTER TABLE `company`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `company_users`
--
ALTER TABLE `company_users`
 ADD UNIQUE KEY `company_id` (`company_id`,`user_id`);

--
-- Indexes for table `news`
--
ALTER TABLE `news`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `project`
--
ALTER TABLE `project`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sms_code`
--
ALTER TABLE `sms_code`
 ADD PRIMARY KEY (`username`), ADD UNIQUE KEY `user_id` (`username`);

--
-- Indexes for table `task`
--
ALTER TABLE `task`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `task_transfer`
--
ALTER TABLE `task_transfer`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_profiles`
--
ALTER TABLE `user_profiles`
 ADD PRIMARY KEY (`user_id`), ADD UNIQUE KEY `user_id` (`user_id`), ADD KEY `uid` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `company`
--
ALTER TABLE `company`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `news`
--
ALTER TABLE `news`
MODIFY `id` int(10) unsigned NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `project`
--
ALTER TABLE `project`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `task`
--
ALTER TABLE `task`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `task_transfer`
--
ALTER TABLE `task_transfer`
MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
MODIFY `id` int(11) unsigned NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
