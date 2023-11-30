create database if not exists eshop_order DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;


create table if not exists eshop_order.order_record(
    id bigint unsigned  PRIMARY KEY AUTO_INCREMENT,
    order_id varchar(45) comment '订单id',
    user_id bigint unsigned not null comment '用户id',
    product_id bigint unsigned not null comment '商品id',
    amount int unsigned not null comment '金额(分)',
    product_count int unsigned not null comment '商品数量',
    coupon_id bigint unsigned not null comment '优惠券id',
    status int(8) unsigned not null comment '订单状态，1:待支付, 2:待发货, 3:待收货, 4:已完成, 5:已取消',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    key(deleted_at),
    key(product_id),
    key(user_id),
    UNIQUE key(order_id)
) ENGINE InnoDB DEFAULT CHARSET utf8mb4;
