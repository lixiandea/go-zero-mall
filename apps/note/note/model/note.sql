CREATE TABLE `note`(
    `id` int unsigned unique not null auto_increment,
    `uid` int unsigned not null default '',
    `content` varchar(10000) not null default '',
    primary key (`id`),index (`uid`)
)engine =InnoDB DEFAULT CHARSET=utf8;