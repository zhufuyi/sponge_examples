CREATE DATABASE IF NOT EXISTS eshop_coupon DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_coupon.coupon
(
    id         bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned not null comment '用户id',
    amount     int unsigned    not null comment '优惠券金额(分)',
    status     int(8) unsigned not null comment '是否使用, 1:未使用, 2:已使用, 3:已过期',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    constraint user_id
        unique (user_id)
)
    charset = utf8mb4;

create index deleted_at
    on eshop_coupon.coupon (deleted_at);

create index user_id_2
    on eshop_coupon.coupon (user_id, status);
