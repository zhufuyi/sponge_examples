CREATE DATABASE IF NOT EXISTS eshop_product DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_product.product
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime     null,
    updated_at  datetime     null,
    deleted_at  datetime     null,
    name        varchar(255) null comment '商品名称',
    price       int          null comment '商品价格',
    photo       varchar(255) null comment '商品图片',
    description text         null comment '商品描述'
);

create index deleted_at
    on eshop_product.product (deleted_at);
