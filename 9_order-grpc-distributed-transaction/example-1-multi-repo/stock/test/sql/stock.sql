create database if not exists eshop_stock DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;


create table if not exists eshop_stock.stock(
    id bigint unsigned PRIMARY KEY AUTO_INCREMENT,
    product_id bigint unsigned not null comment '商品id',
    stock int(11) unsigned not null comment '库存数量',
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    key(deleted_at),
    unique key(product_id)
) ENGINE InnoDB DEFAULT CHARSET utf8mb4;

insert into eshop_stock.stock(product_id, stock, created_at, updated_at) values(1,100,now(),now());
insert into eshop_stock.stock(product_id, stock, created_at, updated_at) values(2,200,now(),now());
