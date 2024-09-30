CREATE DATABASE IF NOT EXISTS eshop_order DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_order.order_record
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime        null,
    updated_at    datetime        null,
    deleted_at    datetime        null,
    order_id      varchar(45)     null comment '订单id',
    user_id       bigint unsigned not null comment '用户id',
    product_id    bigint unsigned not null comment '商品id',
    product_count int unsigned    not null comment '商品数量',
    amount        int unsigned    not null comment '金额(分)',
    coupon_id     bigint unsigned not null comment '优惠券id',
    status        int(8) unsigned not null comment '订单状态: 1:待支付, 2:待发货, 3:待收货, 4:已完成, 5:已取消',
    constraint order_id
        unique (order_id)
)
    charset = utf8mb4;

create index deleted_at
    on eshop_order.order_record (deleted_at);

create index product_id
    on eshop_order.order_record (product_id);

create index user_id
    on eshop_order.order_record (user_id);
