create database if not exists eshop_coupon DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;


create table if not exists eshop_coupon.coupon(
    id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    user_id bigint unsigned not null comment '用户id',
    amount int unsigned not null comment '优惠券金额(分)',
    status int(8) unsigned not null comment '使用状态，1:未使用, 2:已使用, 3:已过期',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    key(deleted_at),
    key(user_id)
) ENGINE InnoDB DEFAULT CHARSET utf8mb4;

insert into eshop_coupon.coupon(user_id, amount, status, created_at, updated_at) values(1,1000,1,now(),now());
insert into eshop_coupon.coupon(user_id, amount, status, created_at, updated_at) values(1,2000,1,now(),now());
