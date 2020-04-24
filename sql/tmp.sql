CREATE TABLE `login` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) NOT NULL COMMENT'用户名',
  `password` varchar(30) NOT NULL COMMENT'密码',
  `phone_number` varchar(11) NOT NULL COMMENT'手机号',
  `gender` varchar(10) NOT NULL COMMENT'性别',
  `login` enum('True','False') NOT NULL DEFAULT 'False' COMMENT'登录状态',
  `last_time` date COMMENT '最后一次登录时间',
  `role` enum('buyer','seller','driver') NOT NULL COMMENT'角色',
  `deleted_at` date,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'用户表';

CREATE TABLE `buyer`(
    `buyer_id` int(11) unsigned NOT NULL,
    `receive_address` varchar(84) COMMENT'收货地址',
    `name`  varchar(30) COMMENT'收货人',
    `receive_number` varchar(11) COMMENT'收货号码',
    `deleted_at` date,
    PRIMARY KEY (`buyer_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'买家';

CREATE TABLE `seller`(
    `seller_id` int(11) unsigned NOT NULL,
    `deliver_address` varchar(84) COMMENT'发货地址',
    `name`  varchar(30) COMMENT'发货人',
    `deliver_number`    varchar(11) COMMENT'发货号码',
    `evaluation` int(11) unsigned DEFAULT 0 COMMENT'总体评价',
    `count` int(11) unsigned DEFAULT 0 COMMENT'评价次数',
    `comprehensive` int(11) unsigned DEFAULT 0 COMMENT'综合评价',
    `deleted_at` date,
    PRIMARY KEY (`seller_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'卖家';

CREATE TABLE `driver`(
    `driver_id` int(11) unsigned NOT NULL,
    `name` varchar(30) COMMENT'姓名',
    `identity` varchar(30) COMMENT'身份证',
    `evaluation` int(11) unsigned DEFAULT 0 COMMENT'总体评价',
    `count` int(11) unsigned DEFAULT 0 COMMENT'评价次数',
    `comprehensive` int(11) unsigned DEFAULT 0 COMMENT'综合评价',
    `deleted_at` date,
    PRIMARY KEY (`driver_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'车主';

CREATE TABLE `goods`(
    `goods_id`  int(11) unsigned NOT NULL AUTO_INCREMENT ,
    `text`  text COMMENT'商品标题',
    `seller_id` int(11) unsigned NOT NULL COMMENT'卖家id',
    `seller_username` varchar(30) NOT NULL COMMENT'卖家用户名',
    `total_num` int(11) unsigned NOT NULL COMMENT '总共数量',
    `remain_num`int(11) unsigned NOT NULL COMMENT '剩余数量',
    `price` int(11) unsigned NOT NULL COMMENT '商品单价',
    `trans_price` int(11) unsigned NOT NULL COMMENT '运输单价',
    `deliver_address` varchar(84) NOT NULL COMMENT'发货地址',
    `create_time` date NOT NULL COMMENT '创建时间',
    `detail` text COMMENT'详细信息',
    `state` int(4) unsigned NOT NULL COMMENT'状态码',
    `identification` varchar(84) NOT NULL COMMENT'唯一标识',
    `deleted_at` date,
    PRIMARY KEY (`goods_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'商品';

CREATE TABLE `seller_goods`(
    `seller_goods_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `goods_id` int(11) unsigned NOT NULL COMMENT '商品id',
    `seller_id` int(11) unsigned NOT NULL COMMENT '卖家id',
    `deleted_at` date,
    PRIMARY KEY (`seller_goods_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'卖家-商品映射表';

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
    `goods_id` int(11) unsigned NOT NULL COMMENT'商品id',
    `seller_id` int(11) unsigned NOT NULL COMMENT'卖家id',
    `buyer_id` int(11) unsigned NOT NULL COMMENT'买家id',
    `text` text COMMENT'订单标题',
    `create_time` date NOT NULL COMMENT '创建时间',
    `total_num` int(11) unsigned NOT NULL COMMENT '总共数量',
    `remain_num`int(11) unsigned NOT NULL COMMENT '剩余数量',
    `total_price` int(11) unsigned NOT NULL COMMENT '商品总价格',
    `total_trans_price` int(11) unsigned NOT NULL COMMENT '运输总价格',
    `deliver_address` varchar(84) NOT NULL COMMENT'发货地址',
    `receive_address` varchar(84) NOT NULL COMMENT'收货地址',
    `state` int(4) unsigned NOT NULL COMMENT'状态码',
    `deleted_at` date,
    PRIMARY KEY(`order_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT'订单';