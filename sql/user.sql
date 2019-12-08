CREATE TABLE `logins` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL COMMENT'用户名',
  `password` varchar(30) NOT NULL COMMENT'密码',
  `gender` varchar(10) NOT NULL COMMENT'性别',
  `email` varchar(30) NOT NULL COMMENT'邮箱',
  `login` enum('True','False') NOT NULL DEFAULT 'False' COMMENT'登录状态',
  `role` enum('student','teacher') NOT NULL COMMENT'角色',
  `deleted_at` date,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'用户表';

CREATE TABLE `person_mappings`(
    `stu_tea_id`int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `pid` int(11) unsigned COMMENT '父id',
    `cid` int(11) unsigned COMMENT '子id',
    `deleted_at` date,
    PRIMARY KEY(`stu_tea_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'学生导师关系表';

CREATE TABLE `missions`(
    `mission_id`int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `text` text COMMENT'任务内容',
    `complete` enum('True','False') NOT NULL DEFAULT 'False' COMMENT '是否完成',
    `create_time` date NOT NULL COMMENT '创建时间',
    `complete_time` date DEFAULT NULL COMMENT '完成时间',
    `mark`int(11) unsigned NOT NULL COMMENT '分数',
    `deleted_at` date,
    PRIMARY KEY(`mission_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'任务';

CREATE TABLE `mission_mappings`(
    `mstu_tea_id` int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `mpid`int(11) unsigned COMMENT '父id',
    `mcid`int(11) unsigned COMMENT '子id',
    `mid` int(11) unsigned COMMENT '任务id',
    `deleted_at` date,
    PRIMARY KEY(`mstu_tea_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'任务从属表';

CREATE TABLE `mark_mappings`(
    `sum_mark_id` int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `stu_id`int(11) unsigned COMMENT '学生id',
    `sum_mark`int(11) unsigned COMMENT '总分',
    `deleted_at` date,
    PRIMARY KEY(`sum_mark_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'总分表';