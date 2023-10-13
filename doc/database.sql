-- 用户表users,存储用户信息，用于注册和登录
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密的密码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态1正常 0禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

CREATE TABLE `groups` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `group_name` varchar(255)  default '' COMMENT 'group name',
    `group_info` varchar(1024) default '' comment 'group info',
    `created_at` datetime default CURRENT_TIMESTAMP,
    `updated_at` datetime default null ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment "group";

CREATE TABLE `group_members` (
      `id` int unsigned NOT NULL AUTO_INCREMENT,
      `group_id` int  default 0 COMMENT 'group id',
      `user_id` int default 0 comment 'member ids',
      `role` enum('admin','member'),
      `created_by` int default 0 comment "creator id",
      `created_at` datetime default CURRENT_TIMESTAMP,
      `updated_at` datetime default null ON UPDATE CURRENT_TIMESTAMP,
      PRIMARY KEY (`id`),
      unique key `idx_group_user` (`group_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment "group_members";

create table `user_groups` (
    `id` int unsigned not null auto_increment,
    `user_id` int unsigned not null default 0 comment "user id",
    `group_id` int unsigned not null default 0 comment "group id",
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default null on update current_timestamp,
    primary key (`id`),
    unique key `idx_user_group` (`user_id`,`group_id`)
) engine = innodb default charset=utf8mb4 collate=utf8mb4_unicode_ci comment "user's groups";

create table `group_join_requests` (
    `id` int unsigned not null auto_increment,
    `user_id` int unsigned not null default 0 comment "user id",
    `group_id` int unsigned not null default 0 comment "group id",
    `message` varchar(1024) not null default '' comment "optional message",
    `created_at` datetime default current_timestamp,
    primary key (`id`),
    key `idx_user_group` (`user_id`,`group_id`)
) engine = innodb default charset=utf8mb4 collate = utf8mb4_unicode_ci comment "group join requests";