CREATE TABLE `video` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `description` text NOT NULL,
    `url` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL default CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    `author` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX(`author`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;