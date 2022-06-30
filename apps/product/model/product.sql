create table `product`(
    `id` bigint unsigned not null AUTO_INCREMENT,
    `name` varchar (255) not null default '' comment '产品名称',
    `desc`  varchar (255) not null default '' comment '产品描述',
    `stock` int(10)  not null  default 0 comment '产品库存',
    `amount` int(10) not null default 0 comment '产品金额',
    `status` tinyint(3) unsigned not null default 0 comment '产品状态',
    `create_time` timestamp null  default CURRENT_TIMESTAMP ,
    'update_time' timestamp null  default current_timestamp on update current_timestamp ,
    primary key (`id`)
)engine = InnoDB default CHARSET=utf8;