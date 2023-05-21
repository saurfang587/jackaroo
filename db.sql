CREATE
DATABASE IF NOT EXISTS `jackaroo`;
USE
`jackaroo`;
DEOP
TABLE IF EXISTS `job`;
CREATE TABLE job
(
    uuid          bigint auto_increment comment '主键id' primary key,
    id            bigint null comment '职位id',
    company       varchar(256) not null comment '所属公司',
    title         varchar(256) not null comment '职位名称',
    job_category  varchar(256) null comment '职位类型',
    job_type_name varchar(256) null comment '招聘类型',
    job_detail    longtext null comment '职位细节',
    job_location  longtext null comment '工作地点',
    push_time     varchar(256) null comment '职位发布时间',
    fetch_time    varchar(256) not null comment '爬取时间'
)ENGINE=InnoDB
DEFAULT CHARSET = utf8mb4 COMMENT = '职位信息表';

