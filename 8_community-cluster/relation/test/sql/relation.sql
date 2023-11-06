
CREATE DATABASE IF NOT EXISTS relation DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use relation;


-- auto-generated definition
create table user_following
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime        null,
    updated_at   datetime        null,
    deleted_at   datetime        null,
    user_id      bigint unsigned not null comment '用户id',
    followed_uid bigint unsigned not null comment '关注用户id',
    status       tinyint         not null comment '关注状态，0:删除, 1:正常'
)
    comment '用户粉丝';

create index user_following_deleted_at_index
    on user_following (deleted_at);

create index user_following_user_id_index
    on user_following (user_id);


-- auto-generated definition
create table user_follower
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime        null,
    updated_at   datetime        null,
    deleted_at   datetime        null,
    user_id      bigint unsigned not null comment '用户id',
    follower_uid bigint unsigned not null comment '粉丝id',
    status       tinyint         not null comment '关注状态，0:删除, 1:正常'
)
    comment '用户粉丝';

create index user_follower_deleted_at_index
    on user_follower (deleted_at);

create index user_follower_user_id_index
    on user_follower (user_id);


-- auto-generated definition
create table relation_num
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime        null,
    updated_at    datetime        null,
    deleted_at    datetime        null,
    user_id       bigint unsigned not null comment '用户id',
    following_num bigint unsigned not null comment '关注数量',
    follower_num  bigint unsigned not null comment '粉丝数量'
)
    comment '用户关注和粉丝数量';

create index relation_num_deleted_at_index
    on relation_num (deleted_at);

create index relation_num_user_id_index
    on relation_num (user_id);


