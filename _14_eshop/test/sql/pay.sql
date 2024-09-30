CREATE DATABASE IF NOT EXISTS eshop_pay DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_pay.pay
(
    id         bigint unsigned auto_increment
        primary key,
    user_id    bigint unsigned not null comment '用户id',
    order_id   varchar(45)     not null comment '订单id',
    amount     int unsigned    not null comment '金额(分)',
    status     int(8) unsigned not null comment '支付状态, 1:待支付, 2:已支付',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    constraint order_id
        unique (order_id)
)
    charset = utf8mb4;

create index deleted_at
    on eshop_pay.pay (deleted_at);

create index user_id
    on eshop_pay.pay (user_id);
