CREATE DATABASE IF NOT EXISTS eshop_user DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table eshop_user.user
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    name       char(50)        not null comment '用户名',
    password   char(100)       not null comment '密码',
    email      char(50)        not null comment '邮箱',
    phone      char(30)        not null comment '电话号码',
    avatar     varchar(200)    null comment '头像',
    age        tinyint         not null comment 'age',
    gender     tinyint         not null comment '性别, 1:男, 2:女, 其他值:未知',
    status     tinyint         not null comment '账号状态, 0:待验证, 1:正常, 2:封禁',
    login_at   bigint unsigned not null comment '登录时间戳',
    login_type tinyint         not null comment '登录类型, 1:web, 2:mobile, 3:desktop, 4:api',
    constraint user_email_uindex
        unique (email)
);

create index deleted_at
    on eshop_user.user (deleted_at);
