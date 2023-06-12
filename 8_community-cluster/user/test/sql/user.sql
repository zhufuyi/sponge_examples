CREATE DATABASE IF NOT EXISTS user DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use user;


-- auto-generated definition
create table user
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    name       varchar(50)  not null comment '用户名',
    nick_name  varchar(50)  not null comment '用户昵称',
    password   varchar(100) not null comment '密码',
    email      varchar(50)  not null comment '邮件',
    phone      varchar(30)  not null comment '手机号码',
    avatar     varchar(200) null comment '头像',
    gender     tinyint      not null comment '性别，1:男，2:女，其他值:未知',
    age        tinyint      not null comment '年龄',
    birthday   varchar(30)  not null comment '出生日期',
    login_at   datetime     null comment '登录时间',
    login_ip   char(16)     null comment '登录ip',
    status     tinyint      not null comment '状态, 1:正常, 2:删除, 3:封禁'
)
    comment '用户信息';

create index user_email_index
    on user (email);

create index user_info_deleted_at_index
    on user (deleted_at);

