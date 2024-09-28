create database if not exists eshop_pay DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;


create table if not exists eshop_pay.pay(
    id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    user_id bigint unsigned not null comment '用户id',
    order_id varchar(45) not null comment '订单id',
    amount int unsigned not null comment '金额(分)',
    status int(8) unsigned not null comment '支付状态，1:待支付, 2:已支付, 3:已退款, 4:已取消',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    key(deleted_at),
    key(user_id),
    UNIQUE key(order_id)
) ENGINE InnoDB DEFAULT CHARSET utf8mb4;
