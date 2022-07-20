CREATE TABLE `user`
(
    `id`         Int          NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255) NOT NULL default '',
    `email`      VARCHAR(255) NOT NULL default '',
    `password`   VARCHAR(255) NOT NULL default '',
    `phone`      bigint       NOT NULL default 0000000000,
    `address`    VARCHAR(255) NOT NULL default '',
    `gender`     tinyint(3)   not null default 1,
    `avatar`     varchar(255) not null default '',
    `status`     varchar(255) not null default '',
    `create_time` timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp    NULL     default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    unique (`phone`),
    unique (`name`)
) engine innodb
  DEFAULT CHARSET = utf8mb4;

