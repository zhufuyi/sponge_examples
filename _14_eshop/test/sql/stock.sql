CREATE DATABASE IF NOT EXISTS eshop_stock DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_stock.stock
(
    id         bigint unsigned auto_increment
        primary key,
    product_id bigint unsigned  not null comment '商品id',
    stock      int(11) unsigned not null comment '库存',
    created_at datetime         null,
    updated_at datetime         null,
    deleted_at datetime         null,
    constraint product_id
        unique (product_id)
)
    charset = utf8mb4;

create index deleted_at
    on eshop_stock.stock (deleted_at);
