
CREATE DATABASE IF NOT EXISTS creation DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use creation;


-- auto-generated definition
create table comment
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime        null,
    updated_at  datetime        null,
    deleted_at  datetime        null,
    post_id     bigint unsigned not null comment '帖子id',
    type        tinyint         not null comment '类型：0:未知, 1:文本, 2:图片, 3:视频',
    user_id     bigint unsigned not null comment '用户id',
    parent_id   bigint unsigned not null comment '父评论id',
    reply_count int unsigned    not null comment '回复数',
    like_count  int unsigned    not null comment '点赞数',
    score       tinyint         not null comment '分数',
    to_uid      bigint unsigned not null comment '给id',
    del_flag    tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '评论详情';

create index comment_info_deleted_at_index
    on comment (deleted_at);


-- auto-generated definition
create table comment_content
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime        null,
    updated_at  datetime        null,
    deleted_at  datetime        null,
    comment_id  bigint unsigned not null comment '评论id',
    content     text            not null comment '评论内容',
    device_type varchar(50)     not null comment '设备类型',
    ip          char(16)        not null comment 'ip地址'
)
    comment '评论内容';

create index comment_content_deleted_at_index
    on comment_content (deleted_at);


-- auto-generated definition
create table comment_hot
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null comment '用户id',
    comment_id bigint unsigned not null comment '评论id',
    post_id    bigint unsigned not null comment '帖子id',
    parent_id  bigint unsigned not null comment '父评论id',
    score      tinyint         not null comment '分数',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '热门评论';

create index comment_hot_deleted_at_index
    on comment_hot (deleted_at);


-- auto-generated definition
create table comment_latest
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    comment_id bigint unsigned not null comment '评论id',
    post_id    bigint unsigned not null comment '帖子id',
    parent_id  bigint unsigned not null comment '父评论id',
    user_id    bigint unsigned not null comment '用户id',
    score      tinyint         not null comment '分数',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '最新评论';

create index comment_latest_deleted_at_index
    on comment_latest (deleted_at);


-- auto-generated definition
create table post
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime        null,
    updated_at    datetime        null,
    deleted_at    datetime        null,
    post_type     tinyint         not null comment '类型：0:未知, 1:文本, 2:图片, 3:视频',
    user_id       bigint unsigned not null comment '用户id',
    title         varchar(250)    not null comment '帖子标题',
    content       text            not null comment '帖子内容',
    view_count    int unsigned    not null comment '查看帖子次数',
    like_count    int unsigned    not null comment '点赞数',
    comment_count int unsigned    not null comment '评论数',
    collect_count int unsigned    not null comment '收藏数',
    share_count   int unsigned    not null comment '分享数',
    longitude     float           not null comment '经度',
    latitude      float           not null comment '纬度',
    position      varchar(100)    not null comment '位置',
    visible       tinyint         not null comment '显示类型，0:公开,1:仅自己可见',
    del_flag      tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '帖子详情';

create index post_info_deleted_at_index
    on post (deleted_at);


-- auto-generated definition
create table post_hot
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null comment '用户id',
    post_id    bigint unsigned not null comment '帖子id',
    score      tinyint         not null comment '分数',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '热门贴子';

create index post_hot_deleted_at_index
    on post_hot (deleted_at);


-- auto-generated definition
create table post_latest
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    post_id    bigint unsigned not null comment '帖子id',
    user_id    bigint unsigned not null comment '用户id',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '最新帖子';

create index post_latest_deleted_at_index
    on post_latest (deleted_at);


-- auto-generated definition
create table user_collect
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null comment '用户id',
    post_id    bigint unsigned not null comment '贴子id'
)
    comment '用户收藏';

create index user_collect_deleted_at_index
    on user_collect (deleted_at);

create index user_collect_user_id_index
    on user_collect (user_id);


-- auto-generated definition
create table user_comment
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    comment_id bigint unsigned not null comment '评论id',
    user_id    bigint unsigned not null comment '用户id',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '用户评论';

create index user_comment_deleted_at_index
    on user_comment (deleted_at);


-- auto-generated definition
create table user_like
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    obj_type   tinyint         not null comment '点赞类型，0:未知, 1:帖子, 2:评论',
    obj_id     bigint unsigned not null comment '点赞类型对应的id',
    user_id    bigint unsigned not null comment '用户id',
    status     tinyint         not null comment '点赞状态，0:未点赞, 1:已点赞'
)
    comment '用户点赞';

create index user_like_deleted_at_index
    on user_like (deleted_at);

create index user_like_obj_id_index
    on user_like (obj_id);


-- auto-generated definition
create table user_post
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    post_id    bigint unsigned not null comment '帖子id',
    user_id    bigint unsigned not null comment '用户id',
    del_flag   tinyint         not null comment '删除方式，0:正常, 1:用户删除, 2:管理员删除'
)
    comment '最新帖子';

create index user_post_deleted_at_index
    on user_post (deleted_at);

