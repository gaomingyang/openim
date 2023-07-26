
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