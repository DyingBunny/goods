CREATE TABLE `login` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL COMMENT'用户名',
  `password` varchar(30) NOT NULL COMMENT'密码',
  `phone_number` varchar(11) NOT NULL COMMENT'手机号',
  `gender` varchar(10) NOT NULL COMMENT'性别',
  `email` varchar(30) NOT NULL COMMENT'邮箱',
  `login` enum('True','False') NOT NULL DEFAULT 'False' COMMENT'登录状态',
  `role` enum('buyer','seller','owner') NOT NULL COMMENT'角色',
  `role_id` int(11) unsigned COMMENT'具体用户id',
  `deleted_at` date,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'用户表';

CREATE TABLE `driver`(
    `driver_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `evaluation` int(11) unsigned DEFAULT 0 COMMENT'总体评价',
    `count` int(11) unsigned DEFAULT 0 COMMENT'评价次数',
    `address_id` int(11) unsigned NOT NULL COMMENT'位置编号',
    `deleted_at` date,
    PRIMARY KEY (`driver_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'骑手';

CREATE TABLE `receive_mapping`(
    `user_address_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `pid` int(11) unsigned COMMENT '父id',
    `cid` int(11) unsigned COMMENT '子id',
    `deleted_at` date,
    PRIMARY KEY (`user_location_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'收货地表';

CREATE TABLE `deliver_mapping`(
    `user_address_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `pid` int(11) unsigned COMMENT '父id',
    `cid` int(11) unsigned COMMENT '子id',
    `deleted_at` date,
    PRIMARY KEY (`user_address_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'发货地表';

CREATE TABLE `address`(
    `address_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(84) NOT NULL COMMENT'位置',
    `lng` decimal(10,7) NOT NULL COMMENT'经度',
    `lat` decimal(10,7) NOT NULL COMMENT'纬度',
    `deleted_at` date,
    PRIMARY KEY (`address_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'地址表';

CREATE TABLE `order`(
    `order_id`int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `text` text COMMENT'订单标题',
    `complete` enum('True','False') NOT NULL DEFAULT 'False' COMMENT '是否完成',
    `create_time` date NOT NULL COMMENT '创建时间',
    `complete_time` date DEFAULT NULL COMMENT '完成时间',
    `price` int(11) unsigned NOT NULL COMMENT '价格',
    `distance` decimal(10,2) unsigned NOT NULL COMMENT '距离',
    `deleted_at` date,
    PRIMARY KEY(`order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'订单';