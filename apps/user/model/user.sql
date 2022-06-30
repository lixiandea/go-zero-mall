CREATE TABLE `user`(
    `id` bigint unsigned not null auto_increment,
    `name` varchar(255) not null default 'DEFAULT_USER' COMMENT '用户姓名',
    `gender` tinyint(3) unsigned not null default '0' COMMENT '用户性别',
    `mobile` varchar(256) not null default '' comment '用户手机号码',
    `password` varchar(255) not null default '' comment '用户密码',
    `create_time` timestamp null  default CURRENT_TIMESTAMP,
    `update_time` timestamp null default CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
    primary key (`id`),
    unique key `idx_mobile_unique` (`mobile`)
)engine =InnoDB DEFAULT CHARSET=utf8;