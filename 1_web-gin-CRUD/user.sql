CREATE DATABASE IF NOT EXISTS school DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use school;

create table teacher
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime     null,
    updated_at  datetime     null,
    deleted_at  datetime     null,
    name        varchar(50)  not null comment '用户名',
    password    varchar(100) not null comment '密码',
    email       varchar(50)  not null comment '邮件',
    phone       varchar(30)  not null comment '手机号码',
    avatar      varchar(200) null comment '头像',
    gender      tinyint      not null comment '性别，1:男，2:女，其他值:未知',
    age         tinyint      not null comment '年龄',
    birthday    varchar(30)  not null comment '出生日期',
    school_name varchar(50)  not null comment '学校名称',
    college     varchar(50)  not null comment '学院',
    title       varchar(10)  not null comment '职称',
    profile     text         not null comment '个人简介'
)
    comment '老师';

create index teacher_deleted_at_index
    on teacher (deleted_at);


create table course
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime    null,
    updated_at datetime    null,
    deleted_at datetime    null,
    code       varchar(10) not null comment '课程代号',
    name       varchar(50) not null comment '课程名称',
    credit     tinyint     not null comment '学分',
    college    varchar(50) not null comment '学院',
    semester   varchar(20) not null comment '学期',
    time       varchar(30) not null comment '上课时间',
    place      varchar(30) not null comment '上课地点'
)
    comment '课程';

create index course_deleted_at_index
    on course (deleted_at);


create table teach
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime    null,
    updated_at   datetime    null,
    deleted_at   datetime    null,
    teacher_id   bigint      not null comment '老师id',
    teacher_name varchar(50) not null comment '老师名称',
    course_id    bigint      not null comment '课程id',
    course_name  varchar(50) not null comment '课程名称',
    score        char(5)     not null comment '学生评价教学质量，5个等级：A,B,C,D,E'
)
    comment '老师课程';

create index teach_course_id_index
    on teach (course_id);

create index teach_deleted_at_index
    on teach (deleted_at);

create index teach_teacher_id_index
    on teach (teacher_id);

