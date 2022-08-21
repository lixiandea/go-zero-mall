CREATE TABLE `note` (
  `id` Int NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL default '',
  `author` Int NOT NULL default 0,
  `content` VARCHAR(65535) NOT NULL default '',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY(`author`)
) engine innodb DEFAULT CHARSET = utf8mb4;