-- 设置默认字符集为utf8mb4以支持中文
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建数据库并指定字符集
CREATE DATABASE IF NOT EXISTS `goods` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `admin` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `user` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `interaction` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `order` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `resource` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `banner` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
       
USE goods;
-- 删除已存在的 goods_info 表（如果存在）
DROP TABLE IF EXISTS `goods_info`;
CREATE TABLE `goods_info`  (
                               `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                               `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品名字',
                               `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主图',
                               `images` json NULL COMMENT '支持单图,多图',
                               `price` int NOT NULL COMMENT '价格(分)',
                               `level1_category_id` int NOT NULL DEFAULT 0 COMMENT '1级分类id',
                               `level2_category_id` int NOT NULL DEFAULT 0 COMMENT '2级分类id',
                               `level3_category_id` int NOT NULL DEFAULT 0 COMMENT '3级分类id',
                               `brand` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌',
                               `stock` int NOT NULL DEFAULT 0 COMMENT '库存',
                               `sale` int NOT NULL DEFAULT 0 COMMENT '销量',
                               `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
                               `sort` tinyint NOT NULL DEFAULT 0 COMMENT '排序 倒叙',
                               `detail_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '商品详情',
                               `enable_bargain` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否允许砍价 1是 0否 默认0',
                               `created_at` datetime NULL DEFAULT NULL,
                               `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `deleted_at` datetime NULL DEFAULT NULL,
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品表' ROW_FORMAT = Dynamic;

-- 商品详情图表
DROP TABLE IF EXISTS `goods_images`;
CREATE TABLE `goods_images` (
            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
            `url` VARCHAR(255) DEFAULT NULL COMMENT '七牛云url',
            `goods_id` INT NOT NULL COMMENT '商品ID',
            `file_id` INT NOT NULL COMMENT '文件ID（关联file_info）',
            `sort` TINYINT NOT NULL DEFAULT 0 COMMENT '排序',
            PRIMARY KEY (`id`),
            INDEX `idx_goods` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品详情图';

-- 优惠券表
DROP TABLE IF EXISTS `coupon_info`;
-- 创建 coupon_info 表
CREATE TABLE `coupon_info` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `goods_id` int NOT NULL DEFAULT '0' COMMENT '关联商品id（0表示全场通用）',
                               `name` varchar(100) NOT NULL COMMENT '优惠券名称',
                               `type` tinyint NOT NULL DEFAULT '0' COMMENT '优惠券类型：0-新人券，1-活动券，2-其他',
                               `amount` int NOT NULL DEFAULT '0' COMMENT '优惠金额（单位：分）',
                               `deadline` datetime NOT NULL COMMENT '过期时间',
                               `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               `deleted_at` datetime DEFAULT NULL COMMENT '删除时间（软删除）',
                               PRIMARY KEY (`id`),
                               KEY `idx_goods_id` (`goods_id`),
                               KEY `idx_deadline` (`deadline`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='优惠券信息表';

-- 删除已存在的 user_coupon_info 表（如果存在）
DROP TABLE IF EXISTS `user_coupon_info`;

-- 创建 user_coupon_info 表
CREATE TABLE `user_coupon_info` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
                                    `coupon_id` int NOT NULL COMMENT '优惠券id',
                                    `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待使用，1-已使用，2-已过期',
                                    `amount` int NOT NULL DEFAULT '0' COMMENT '优惠金额（单位：分）',
                                    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间（软删除）',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_user_id` (`user_id`),
                                    KEY `idx_coupon_id` (`coupon_id`),
                                    KEY `idx_status` (`status`),
                                    UNIQUE KEY `uk_user_coupon` (`user_id`, `coupon_id`) COMMENT '同一用户不能重复领取同张优惠券'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户优惠券信息表';

-- ----------------------------
-- Table structure for category_info
-- ----------------------------
DROP TABLE IF EXISTS `category_info`;
CREATE TABLE `category_info`  (
              `id` int NOT NULL AUTO_INCREMENT,
              `parent_id` int NOT NULL DEFAULT 0 COMMENT '父级id',
              `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
              `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'icon',
              `level` tinyint(1) NOT NULL DEFAULT 1 COMMENT '等级 默认1级分类',
              `sort` tinyint(1) NOT NULL DEFAULT 1,
              `created_at` datetime DEFAULT NULL,
              `updated_at` datetime DEFAULT NULL,
              `deleted_at` datetime DEFAULT NULL,
              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '轮播图表\n' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category_info
-- ----------------------------
INSERT INTO `category_info` VALUES (1, 0, '家用电器', '', 1, 1, NULL, NULL, NULL);
INSERT INTO `category_info` VALUES (2, 1, '电视', '', 2, 1, NULL, NULL, NULL);
INSERT INTO `category_info` VALUES (3, 2, '全面屏电视', '', 3, 1, NULL, NULL, NULL);
INSERT INTO `category_info` VALUES (4, 2, '教育电视', '', 3, 1, NULL, NULL, NULL);
INSERT INTO `category_info` VALUES (5, 1, '智慧屏电视', '', 3, 1, NULL, NULL, NULL);
INSERT INTO `category_info` VALUES (6, 0, '手机/数码', '', 1, 2, '2022-07-27 15:07:31', '2022-07-27 15:08:57', NULL);
INSERT INTO `category_info` VALUES (7, 66, '111', 'http://dummyimage.com/400x400', 62, 26, '2022-07-27 15:08:41', '2023-01-13 21:25:55', NULL);
INSERT INTO `category_info` VALUES (8, 9, '理收每从最想', 'http://dummyimage.com/400x400', 68, 99, '2023-01-13 21:17:33', '2023-01-13 21:17:33', '2023-01-13 21:19:07');

-- ----------------------------
-- Table structure for cart_info
-- ----------------------------
DROP TABLE IF EXISTS `cart_info`;
CREATE TABLE `cart_info`  (
      `id` int NOT NULL AUTO_INCREMENT COMMENT '购物车表',
      `user_id` int NOT NULL DEFAULT 0,
      `goods_id` int NOT NULL DEFAULT 0 COMMENT '商品id',
      `count` int NOT NULL COMMENT '商品数量',
      `created_at` datetime DEFAULT NULL,
      `updated_at` datetime DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- 删除已存在的 coupon_info 表（如果存在）
DROP TABLE IF EXISTS `coupon_info`;

-- 创建 coupon_info 表
CREATE TABLE `coupon_info` (
       `id` int NOT NULL AUTO_INCREMENT,
       `goods_id` int NOT NULL DEFAULT '0' COMMENT '关联商品id（0表示全场通用）',
       `name` varchar(100) NOT NULL COMMENT '优惠券名称',
       `type` tinyint NOT NULL DEFAULT '0' COMMENT '优惠券类型：0-新人券，1-活动券，2-其他',
       `amount` int NOT NULL DEFAULT '0' COMMENT '优惠金额（单位：分）',
       `deadline` datetime NOT NULL COMMENT '过期时间',
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
       `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
       `deleted_at` datetime DEFAULT NULL COMMENT '删除时间（软删除）',
       PRIMARY KEY (`id`),
       KEY `idx_goods_id` (`goods_id`),
       KEY `idx_deadline` (`deadline`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='优惠券信息表';

-- 删除已存在的 user_coupon_info 表（如果存在）
DROP TABLE IF EXISTS `user_coupon_info`;

-- 创建 user_coupon_info 表
CREATE TABLE `user_coupon_info` (
        `id` int NOT NULL AUTO_INCREMENT,
        `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
        `coupon_id` int NOT NULL COMMENT '优惠券id',
        `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态：0-待使用，1-已使用，2-已过期',
        `amount` int NOT NULL DEFAULT '0' COMMENT '优惠金额（单位：分）',
        `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间（软删除）',
        PRIMARY KEY (`id`),
        KEY `idx_user_id` (`user_id`),
        KEY `idx_coupon_id` (`coupon_id`),
        KEY `idx_status` (`status`),
        UNIQUE KEY `uk_user_coupon` (`user_id`, `coupon_id`) COMMENT '同一用户不能重复领取同张优惠券'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户优惠券信息表';

USE user;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
          `id` int NOT NULL AUTO_INCREMENT,
          `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
          `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
          `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
          `user_salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐 生成密码用',
          `open_id` varchar(64) NOT NULL DEFAULT '' COMMENT '微信openid',
          `user_salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐 生成密码用',
          `sex` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1男 2女',
          `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1正常 2拉黑冻结',
          `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '个性签名',
          `secret_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密保问题的答案',
          `created_at` datetime DEFAULT NULL,
          `updated_at` datetime DEFAULT NULL,
          `deleted_at` datetime DEFAULT NULL,
          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (1, 'lida', 'https://img1.baidu.com/it/u=2029513305,2137933177&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=472', '26bebfe4cf87cc2bd7b89c237fe42df3', 'QLAFRsKG2N', 1, 1, '个性签名', '银河中学', '2022-07-28 17:19:42', '2022-07-31 19:25:01', NULL);
INSERT INTO `user_info` VALUES (2, 'wang', '', '', '', 1, 1, '', '', NULL, NULL, NULL);
INSERT INTO `user_info` VALUES (3, '作证表之有', 'http://dummyimage.com/100x100', '34929b5d84cf66ae73797d5b48297710', '6bZBjqX1Nk', 1, 92, 'incididunt Excepteur aliqua non', 'cupidatat magna', '2023-01-18 08:19:38', '2023-01-18 08:19:38', NULL);
INSERT INTO `user_info` VALUES (4, 'wzy', 'http://dummyimage.com/100x100', 'a90f083adf941cd75bde9cc371fbe00b', 'wpm0bGNBrf', 63, 52, 'Excepteur', 'voluptate in non ea ut', '2023-01-18 09:54:46', '2023-01-18 12:14:13', NULL);
INSERT INTO `user_info` VALUES (5, 'wangzhongyang', 'http://dummyimage.com/100x100', '82131d93ab13a1a4f9ec840a9ddbabf7', 'T0iKtv31BU', 1, 1, '和我一起学编程吧', '六个1', '2024-12-26 11:25:43', '2024-12-26 11:25:43', NULL);


-- ----------------------------
-- Table structure for consignee_info
-- ----------------------------
DROP TABLE IF EXISTS `consignee_info`;
CREATE TABLE `consignee_info`  (
           `id` int NOT NULL AUTO_INCREMENT COMMENT '收货地址表',
           `user_id` int NOT NULL DEFAULT 0,
           `is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '默认地址1  非默认0\n',
           `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
           `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
           `province` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
           `city` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
           `town` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
           `street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '街道乡镇',
           `detail` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址详情',
           `created_at` datetime DEFAULT NULL,
           `updated_at` datetime DEFAULT NULL,
           `deleted_at` datetime DEFAULT NULL,
           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of consignee_info
-- ----------------------------
INSERT INTO `consignee_info` VALUES (1, 1, 1, '王先生1', '13269477632', '北京', '北京市', '房山区', '拱辰街道', '大学城西', '2022-07-31 14:42:33', '2022-07-31 14:44:50', NULL);

USE interaction;
    
--  创建数据库   interaction

-- ----------------------------
-- Table structure for comment_info
-- ----------------------------
DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info`  (
         `id` int NOT NULL AUTO_INCREMENT,
         `parent_id` int NOT NULL DEFAULT 0 COMMENT '父级评论id',
         `user_id` int NOT NULL DEFAULT 0,
         `object_id` int NOT NULL DEFAULT 0,
         `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '评论类型：1商品 2文章',
         `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论内容',
         `created_at` datetime DEFAULT NULL,
         `updated_at` datetime DEFAULT NULL,
         `deleted_at` datetime DEFAULT NULL,
         PRIMARY KEY (`id`) USING BTREE,
         UNIQUE INDEX `unique_index`(`user_id`, `object_id`, `type`, `content`, `parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment_info
-- ----------------------------
INSERT INTO `comment_info` VALUES (4, 0, 1, 1, 2, '好评 下次还会买', '2022-07-31 17:23:48', '2022-07-31 17:23:48', NULL);
INSERT INTO `comment_info` VALUES (5, 0, 1, 1, 2, '来个评论', '2022-07-31 17:24:10', '2022-07-31 17:24:10', NULL);
INSERT INTO `comment_info` VALUES (7, 5, 1, 1, 2, '来个评论', '2022-07-31 17:24:59', '2022-07-31 17:24:59', NULL);
INSERT INTO `comment_info` VALUES (10, 1, 4, 1, 1, 'labore', '2023-01-19 14:25:24', '2023-01-19 14:25:24', NULL);
INSERT INTO `comment_info` VALUES (11, 1, 4, 1, 1, 'xxxxx', '2023-01-19 14:26:50', '2023-01-19 14:26:50', NULL);


-- ----------------------------
-- Table structure for praise_info
-- ----------------------------
DROP TABLE IF EXISTS `praise_info`;
CREATE TABLE `praise_info`  (
        `id` int NOT NULL AUTO_INCREMENT COMMENT '点赞表',
        `user_id` int NOT NULL,
        `type` tinyint(1) NOT NULL COMMENT '点赞类型 1商品 2文章',
        `object_id` int NOT NULL DEFAULT 0 COMMENT '点赞对象id 方便后期扩展',
        `created_at` datetime DEFAULT NULL,
        `updated_at` datetime DEFAULT NULL,
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `unique_index`(`user_id`, `type`, `object_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of praise_info
-- ----------------------------
INSERT INTO `praise_info` VALUES (8, 4, 1, 1, '2023-01-19 12:18:07', '2023-01-19 12:18:07');


-- ----------------------------
-- Table structure for collection_info
-- ----------------------------
DROP TABLE IF EXISTS `collection_info`;
CREATE TABLE `collection_info`  (
        `id` int NOT NULL AUTO_INCREMENT,
        `user_id` int NOT NULL DEFAULT 0 COMMENT '用户id',
        `object_id` int NOT NULL DEFAULT 0 COMMENT '对象id',
        `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏类型：1商品 2文章',
        `created_at` datetime DEFAULT NULL,
        `updated_at` datetime DEFAULT NULL,
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `unique_index`(`user_id`, `object_id`, `type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of collection_info
-- ----------------------------
INSERT INTO `collection_info` VALUES (3, 1, 1, 1, '2022-07-31 15:21:38', '2022-07-31 15:21:38');
INSERT INTO `collection_info` VALUES (4, 4, 4, 1, '2023-01-18 15:23:28', '2023-01-18 15:23:28');


USE `order`;
    
-- ----------------------------
-- Table structure for order_goods_info
-- ----------------------------
DROP TABLE IF EXISTS `order_goods_info`;
CREATE TABLE `order_goods_info`  (
         `id` int NOT NULL AUTO_INCREMENT COMMENT '商品维度的订单表',
         `order_id` int NOT NULL DEFAULT 0 COMMENT '关联的主订单表',
         `goods_id` int NOT NULL DEFAULT 0 COMMENT '商品id',
         `goods_options_id` int DEFAULT 0 COMMENT '商品规格id sku id',
         `count` int NOT NULL COMMENT '商品数量',
         `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
         `price` int NOT NULL DEFAULT 0 COMMENT '订单金额 单位分',
         `coupon_price` int NOT NULL DEFAULT 0 COMMENT '优惠券金额 单位分',
         `actual_price` int NOT NULL DEFAULT 0 COMMENT '实际支付金额 单位分',
         `created_at` datetime DEFAULT NULL,
         `updated_at` datetime DEFAULT NULL,
         PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章（种草）表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_goods_info
-- ----------------------------
INSERT INTO `order_goods_info` VALUES (1, 1, 1, 0, 1, '', 100, 10, 90, NULL, NULL);
INSERT INTO `order_goods_info` VALUES (2, 8, 1, 0, 1, '', 0, 0, 0, '2022-08-27 20:50:50', '2022-08-27 20:50:50');
INSERT INTO `order_goods_info` VALUES (3, 8, 2, 0, 3, '', 0, 0, 0, '2022-08-27 20:50:50', '2022-08-27 20:50:50');
INSERT INTO `order_goods_info` VALUES (21, 13, 1, 5, 1, 'laboris consectetur in minim', 74, 67, 4, '2023-02-07 15:59:43', '2023-02-07 15:59:43');
INSERT INTO `order_goods_info` VALUES (22, 13, 1, 6, 2, 'ut amet laboris laborum dolore', 69, 89, 80, '2023-02-07 15:59:43', '2023-02-07 15:59:43');
INSERT INTO `order_goods_info` VALUES (23, 15, 1, 5, 1, '无', 10, 1, 9, '2023-02-09 14:15:59', '2023-02-09 14:15:59');
INSERT INTO `order_goods_info` VALUES (24, 16, 1, 5, 1, '无', 10, 1, 9, '2023-02-09 14:51:09', '2023-02-09 14:51:09');

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info`  (
           `id` int NOT NULL AUTO_INCREMENT,
           `number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
           `user_id` int NOT NULL DEFAULT 0 COMMENT '用户id',
           `pay_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '支付方式 1微信 2支付宝 3云闪付',
           `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
           `pay_at` datetime DEFAULT NULL COMMENT '支付时间',
           `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消',
           `consignee_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
           `consignee_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人手机号',
           `consignee_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
           `price` int NOT NULL DEFAULT 0 COMMENT '订单金额 单位分',
           `coupon_price` int NOT NULL DEFAULT 0 COMMENT '优惠券金额 单位分',
           `actual_price` int NOT NULL DEFAULT 0 COMMENT '实际支付金额 单位分',
           `created_at` datetime DEFAULT NULL,
           `updated_at` datetime DEFAULT NULL,
           PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章（种草）表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_info
-- ----------------------------
INSERT INTO `order_info` VALUES (1, '1659231316407832000111', 1, 1, '0', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 100, 9900, '2022-08-27 09:35:16', '2022-07-31 09:35:16');
INSERT INTO `order_info` VALUES (2, '1659231554317361000757', 1, 1, '0', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 200, 9800, '2022-07-31 09:39:14', '2022-07-31 09:39:14');
INSERT INTO `order_info` VALUES (3, '1661603467832912000516', 1, 0, '', '2022-12-13 21:52:26', 0, '', '', '', 0, 0, 0, '2022-12-08 20:31:07', '2022-08-27 20:31:07');
INSERT INTO `order_info` VALUES (4, '1661603562656619000513', 1, 1, '放到快递柜就可以，不用打电话。', '2022-12-13 21:52:19', 0, '王先生', '13269477432', '北京丰台汽车博物馆', 0, 0, 0, '2022-12-09 20:32:42', '2022-08-27 20:32:42');
INSERT INTO `order_info` VALUES (5, '1661604424031843000546', 1, 0, '', '2022-12-13 21:52:12', 0, '', '', '', 0, 0, 0, '2022-12-10 20:47:04', '2022-08-27 20:47:04');
INSERT INTO `order_info` VALUES (6, '1661604530142913000770', 1, 1, '这是备注', '2022-12-13 21:52:05', 1, '', '', '', 100, 0, 0, '2022-12-11 20:48:50', '2022-08-27 20:48:50');
INSERT INTO `order_info` VALUES (7, '166160461284091500027', 1, 1, '这是备注', '2022-12-13 21:51:58', 1, '', '', '', 100, 0, 9800, '2022-12-12 20:50:50', '2022-08-27 20:50:12');
INSERT INTO `order_info` VALUES (8, '166160465089079000090', 1, 1, '这是备注', '2022-12-19 13:14:07', 1, '', '', '', 100, 0, 9800, '2022-12-19 20:50:50', '2022-08-27 20:50:50');
INSERT INTO `order_info` VALUES (9, '1675756215452071000340', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, '2023-02-07 15:50:15', '2023-02-07 15:50:15');
INSERT INTO `order_info` VALUES (10, '1675756478857950000699', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, '2023-02-07 15:54:38', '2023-02-07 15:54:38');
INSERT INTO `order_info` VALUES (11, '1675756573337742000600', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, '2023-02-07 15:56:13', '2023-02-07 15:56:13');
INSERT INTO `order_info` VALUES (13, '1675756783046217000276', 15, 0, 'consequat cupidatat in', NULL, 1, '在区三少', '18127326325', '广东省天津市安溪县', 87, 66, 92, '2023-02-07 15:59:43', '2023-02-07 15:59:43');
INSERT INTO `order_info` VALUES (14, '1675923255725252000737', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, '2023-02-09 14:14:15', '2023-02-09 14:14:15');
INSERT INTO `order_info` VALUES (15, '1675923359070046000221', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, '2023-02-09 14:15:59', '2023-02-09 14:15:59');
INSERT INTO `order_info` VALUES (16, '1675925468868358000350', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, '2023-02-09 14:51:09', '2023-02-09 14:51:09');


-- ----------------------------
-- Table structure for refund_info
-- ----------------------------
DROP TABLE IF EXISTS `refund_info`;
CREATE TABLE `refund_info`  (
        `id` int NOT NULL AUTO_INCREMENT COMMENT '售后退款表',
        `number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '售后订单号',
        `order_id` int NOT NULL COMMENT '订单id',
        `goods_id` int NOT NULL DEFAULT 0 COMMENT '要售后的商品id\n',
        `reason` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退款原因',
        `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1待处理 2同意退款 3拒绝退款\n',
        `user_id` int NOT NULL COMMENT '用户id',
        `created_at` datetime DEFAULT NULL,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of refund_info
-- ----------------------------
INSERT INTO `refund_info` VALUES (1, 'refund1659247832739250000428', 1, 1, '不想要了', 1, 1, '2022-07-31 14:10:32', '2022-07-31 14:10:32', NULL);

USE resource;

-- 数据库 resource
DROP TABLE IF EXISTS `file_info`;
-- 统一文件表
CREATE TABLE `file_info` (
     `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
     `name` VARCHAR(255) NOT NULL COMMENT '文件名字',
     `url` VARCHAR(255) NOT NULL COMMENT '七牛云URL',
     `uploader_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者ID（根据uploader_type区分是用户ID还是管理员ID）',
     `uploader_type` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者类型：1-H5用户，2-管理员',
     `file_type` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件类型：1-图片，2-视频，3-其他',
     `created_at` datetime(0) DEFAULT NULL,
     `deleted_at` datetime(0) DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件存储表';

USE banner;

-- 创建数据库  banner

-- ----------------------------
-- Table structure for rotation_info
-- ----------------------------
DROP TABLE IF EXISTS `rotation_info`;
CREATE TABLE `rotation_info`  (
          `id` int NOT NULL AUTO_INCREMENT,
          `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '轮播图片',
          `link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转链接',
          `sort` tinyint(1) NOT NULL DEFAULT 0 COMMENT '排序字段',
          `created_at` datetime DEFAULT NULL,
          `updated_at` datetime DEFAULT NULL,
          `deleted_at` datetime DEFAULT NULL,
          PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '轮播图表\n' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rotation_info
-- ----------------------------
INSERT INTO `rotation_info` VALUES (1, '111', '11', 10, '2022-07-19 04:53:01', '2022-07-19 04:59:24', NULL);

-- ----------------------------
-- Table structure for position_info
-- ----------------------------
DROP TABLE IF EXISTS `position_info`;
CREATE TABLE `position_info`  (
          `id` int NOT NULL AUTO_INCREMENT,
          `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片链接',
          `goods_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
          `link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转链接',
          `sort` tinyint(1) NOT NULL DEFAULT 0 COMMENT '排序',
          `goods_id` int NOT NULL DEFAULT 0 COMMENT '商品id',
          `created_at` datetime DEFAULT NULL,
          `updated_at` datetime DEFAULT NULL,
          `deleted_at` datetime DEFAULT NULL,
          PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of position_info
-- ----------------------------
INSERT INTO `position_info` VALUES (2, 'https://images.zsxq.com/FgdL08hVmh-40_e12vh-ifbXpGxB?e=2000966400', '测试', 'https://articles.zsxq.com/id_wd15wsegvow1.html', 0, 1, '2022-11-18 17:44:07', '2022-11-18 17:44:07', '2022-11-18 17:44:59');

/*
下面是管理后台需要的数据表
*/
USE admin;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
                                `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                                `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('g', 'u_1', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_43', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '54', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '55', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '56', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '57', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '58', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '17', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '18', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_5', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_31', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_6', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_16', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_10', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '114', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '115', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '116', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '117', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '118', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '119', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '33', 'All', '', '', '');

-- ----------------------------
-- Table structure for demo_data_auth
-- ----------------------------
DROP TABLE IF EXISTS `demo_data_auth`;
CREATE TABLE `demo_data_auth`  (
                                   `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
                                   `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标题',
                                   `created_by` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
                                   `updated_by` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改人',
                                   `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                   `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                                   `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '数据权限测试' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_data_auth
-- ----------------------------
INSERT INTO `demo_data_auth` VALUES (3, '测试01', 31, 31, '2022-03-03 10:15:11', '2023-02-07 22:38:05', NULL);
INSERT INTO `demo_data_auth` VALUES (4, '测试02', 16, 31, '2022-03-03 10:36:52', '2023-02-07 18:27:39', NULL);
INSERT INTO `demo_data_auth` VALUES (5, '测试03', 22, 31, '2022-03-03 10:37:47', '2023-02-07 18:28:49', NULL);
INSERT INTO `demo_data_auth` VALUES (6, '测试04', 28, 31, '2022-03-03 10:37:53', '2023-02-07 18:27:43', NULL);
INSERT INTO `demo_data_auth` VALUES (7, '测试05', 20, 0, '2022-03-03 10:37:58', '2022-03-03 10:37:58', NULL);
INSERT INTO `demo_data_auth` VALUES (8, '测试06', 32, 31, '2022-03-03 10:38:05', '2023-02-07 18:29:44', NULL);
INSERT INTO `demo_data_auth` VALUES (9, '测试07', 32, 31, '2022-03-03 10:38:12', '2023-02-07 22:38:08', NULL);
INSERT INTO `demo_data_auth` VALUES (10, '测试08', 31, 31, '2022-03-03 10:38:18', '2023-02-07 22:38:10', NULL);

-- ----------------------------
-- Table structure for demo_gen
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen`;
CREATE TABLE `demo_gen`  (
                             `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `demo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
                             `demo_age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
                             `classes` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级',
                             `demo_born` datetime NULL DEFAULT NULL COMMENT '出生年月',
                             `demo_gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                             `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
                             `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                             `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                             `demo_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
                             `demo_cate` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
                             `demo_thumb` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '头像',
                             `demo_photo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '相册',
                             `demo_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '个人描述',
                             `demo_file` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '相关附件',
                             `classes_two` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级二',
                             `cate_trees` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '树型单选',
                             `cate_trees_two` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '树形多选',
                             `options` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '其他选项',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成测试表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of demo_gen
-- ----------------------------
INSERT INTO `demo_gen` VALUES (4, '张三', 50, '6', '2022-12-20 11:23:36', 1, '2021-08-21 12:07:19', '2024-05-07 14:16:32', NULL, 3100, 31, 1, '0,1', 'upload_file/2022-11-11/co9copop81co0gysbz.jpg', '[{\"name\":\"74595d93db72bc45e5f5161ca35f5995.jpg\",\"url\":\"upload_file/2022-12-20/cp6bmkqcyfpc30zh52.jpg\",\"fileType\":\"image/jpeg\",\"size\":62462},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-20/cp6bmkqcnncsyckbae.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-20/cp6bmkqcnncsx40dak.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"t01b97b4bd97190d33a.jpg\",\"url\":\"upload_file/2024-05-07/d136y46kwjfwzi7c4q.jpg\",\"fileType\":\"image/jpeg\",\"size\":9859}]', '<p style=\"text-indent: 2em;\">快乐就好可厉害了夸奖哈啥是利空打击阿松大快乐就好可厉害了夸奖哈啥是利空打击阿松大快乐就好可厉害了夸奖哈啥是利空打击阿松大快乐就好可厉害了夸奖哈啥是利空打击阿松大快乐就好可厉害了夸奖哈啥是利空打击阿松大快乐就好可厉害了夸奖哈啥是利空打击阿松大</p><p><img src=\"http://localhost:8808/upload_file/2022-12-20/cp6bmsehgh207ehhry.jpg\" style=\"width: 400px; height: 229px;\" width=\"400\" height=\"229\" border=\"0\" vspace=\"0\"/></p><p><img src=\"http://localhost:8808/upload_file/2022-12-20/cp6bmserxxy0cg4u72.jpg\" style=\"width: 400px; height: 250px;\" width=\"400\" height=\"250\" border=\"0\" vspace=\"0\"/></p><p><br/></p>', '[{\"name\":\"1.xlsx\",\"url\":\"http://localhost:8808//pub_upload/2021-08-21/cdow7mg24tu4f5yuid.xlsx\",\"fileType\":\"\",\"size\":0},{\"name\":\"楚雄市数据楚雄应用统计表.xls\",\"url\":\"http://localhost:8808//pub_upload/2021-08-21/cdow7rjjzk0wtpk74t.xls\",\"fileType\":\"\",\"size\":0},{\"name\":\"数据云南指标.xlsx\",\"url\":\"upload_file/2023-08-22/cuyvs3r442wsh1dq2a.xlsx\",\"fileType\":\"application/vnd.DEMO_WECHAT_OPEN_ID.spreadsheetml.sheet\",\"size\":11618}]', '4,6,5', 4, '6,5', '[{\"key\":\"name\",\"value\":\"张三\"},{\"key\":\"age\",\"value\":\"18\"}]');
INSERT INTO `demo_gen` VALUES (5, '里斯', 30, '5', '2022-12-06 03:03:04', 2, '2022-12-14 16:19:21', '2023-05-28 14:36:42', NULL, 31, 31, 0, '0', 'upload_file/2022-12-20/cp6bni4ojf30sroz55.jpg', '[{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-20/cp6bnmhefzswddxrys.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-20/cp6bnmhefzswthixpp.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079}]', '<p>客户机阿斯利康到家啦可是建档立卡</p>', '[{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-20/cp69z5ryzgews7ovj7.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-20/cp6amc8ulju033lhyp.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947},{\"name\":\"a333.jpg\",\"url\":\"upload_file/2022-12-20/cp6amc8ulju0j9yq3d.jpg\",\"fileType\":\"image/jpeg\",\"size\":18381}]', '4', 5, '1,5', NULL);
INSERT INTO `demo_gen` VALUES (6, '王五', 56, '4', '2022-12-05 15:05:02', 1, '2022-12-15 08:56:00', '2023-09-19 15:11:39', NULL, 31, 31, 1, '1', 'upload_file/2022-12-20/cp6bnyl49axcbedg2y.jpg', '[{\"name\":\"2.jpg\",\"url\":\"upload_file/2022-12-20/cp6bo15yvl9gvzlkok.jpg\",\"fileType\":\"image/jpeg\",\"size\":18618},{\"name\":\"t01b3a5a18109dea24a.jpg\",\"url\":\"upload_file/2022-12-20/cp6bo15yvl9gffm2ar.jpg\",\"fileType\":\"image/jpeg\",\"size\":36537}]', '<p>个人描述信息</p>', '[]', '5', 0, '', '[{\"key\":\"\",\"value\":\"\"}]');
INSERT INTO `demo_gen` VALUES (7, '赵四', 53, '4', '2022-12-15 10:12:24', 1, '2022-12-15 10:13:53', '2022-12-20 11:25:05', NULL, 31, 31, 1, '1', 'upload_file/2022-12-20/cp6bo5wte6mo0ahvhl.jpg', '[{\"name\":\"t01698c1bc3af22a34b.jpg\",\"url\":\"upload_file/2022-12-20/cp6boael3pk4yhzjlm.jpg\",\"fileType\":\"image/jpeg\",\"size\":21834},{\"name\":\"1.jpg\",\"url\":\"upload_file/2022-12-20/cp6boaelehg0h6gvub.jpg\",\"fileType\":\"image/jpeg\",\"size\":13610}]', '<p>富文本内容</p>', '[]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (8, '刘涛', 18, '4', '2022-12-15 10:47:23', 1, '2022-12-15 10:49:47', '2022-12-20 11:26:15', NULL, 31, 31, 1, '0', 'upload_file/2022-12-20/cp6boein2zq4uuywuh.jpg', '[{\"name\":\"t01b97b4bd97190d33a.jpg\",\"url\":\"upload_file/2022-12-20/cp6boiffcarkavs1tp.jpg\",\"fileType\":\"image/jpeg\",\"size\":9859}]', '<p>富文本</p><p><br/></p><video id=\"tmpVideo0\" class=\"edui-video-video\" controls=\"\" preload=\"none\" width=\"532\" height=\"318\" src=\"http://localhost:8808/upload_file/2022-12-20/cp6boza7n3ykyp0vdr.mp4\" data-setup=\"{}\" style=\"width: 532px; height: 318px;\"><source src=\"http://localhost:8808/upload_file/2022-12-20/cp6boza7n3ykyp0vdr.mp4\" type=\"video/mp4\"/></video><p><br/></p><p><br/></p><p><br/></p>', '[{\"name\":\"74595d93db72bc45e5f5161ca35f5995.jpg\",\"url\":\"upload_file/2022-12-15/cp21ptkim0ow0z3iys.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"a333.jpg\",\"url\":\"upload_file/2022-12-15/cp21qys54nxwgxzjxg.jpg\",\"fileType\":\"\",\"size\":0}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (9, '刘涛', 18, '4', '2022-12-15 10:47:23', 1, '2022-12-15 10:52:38', '2022-12-15 10:52:38', '2022-12-20 10:02:37', 31, 0, 1, '', 'upload_file/2022-12-15/cp21qt8kek30kh1fej.jpg', '[{\"name\":\"65d3d7ad866394bf86309af1bbba11a3.jpeg\",\"url\":\"http://localhost:8808/upload_file/2022-12-15/cp21plv1zpawhwy53b.jpeg\",\"path\":\"upload_file/2022-12-15/cp21plv1zpawhwy53b.jpeg\"},{\"name\":\"a333.jpg\",\"url\":\"http://localhost:8808/upload_file/2022-12-15/cp21quzjof9wg9cm28.jpg\",\"path\":\"upload_file/2022-12-15/cp21quzjof9wg9cm28.jpg\"}]', '<p>富文本内容</p>', '[{\"name\":\"74595d93db72bc45e5f5161ca35f5995.jpg\",\"url\":\"http://localhost:8808/upload_file/2022-12-15/cp21ptkim0ow0z3iys.jpg\",\"path\":\"upload_file/2022-12-15/cp21ptkim0ow0z3iys.jpg\"},{\"name\":\"a333.jpg\",\"url\":\"http://localhost:8808/upload_file/2022-12-15/cp21qys54nxwgxzjxg.jpg\",\"path\":\"upload_file/2022-12-15/cp21qys54nxwgxzjxg.jpg\"}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (10, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 11:06:29', '2022-12-15 11:06:29', '2022-12-20 10:02:28', 31, 0, 1, '', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejse0u2xs.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>流量卡圣诞节拉克丝的拉克斯基的</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (11, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 11:33:49', '2022-12-15 16:42:28', '2022-12-20 10:02:28', 31, 31, 1, '1', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"t01b97b4bd97190d33a.jpg\",\"url\":\"upload_file/2022-12-15/cp29adtc1nyklqvr9a.jpg\",\"fileType\":\"image/jpeg\",\"size\":9859},{\"name\":\"t01b69a50d3f1ec115b.jpg\",\"url\":\"upload_file/2022-12-15/cp29af4z2178lzkdd3.jpg\",\"fileType\":\"image/jpeg\",\"size\":19040}]', '<p>aaaaaaaaaaaaaaaaaaaaaaa</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (12, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 16:56:16', '2022-12-15 16:56:16', '2022-12-20 10:02:21', 31, 0, 1, '1', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejse0u2xs.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>流量卡圣诞节拉克丝的拉克斯基的</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (13, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 16:56:17', '2022-12-18 14:16:56', '2022-12-20 10:02:15', 31, 31, 1, '1,0', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"t01b97b4bd97190d33a.jpg\",\"url\":\"upload_file/2022-12-15/cp2a2juhxyq0ehrr0s.jpg\",\"fileType\":\"image/jpeg\",\"size\":9859}]', '<p>流量卡圣诞节拉克丝的拉克54656斯基的</p><p><img src=\"http://localhost:8808/upload_file/2022-12-16/cp2vz25diy08kqgzav.jpg\" alt=\"5b5fd982ce018.jpg\"/></p><p><video id=\"tmpVideo0\" class=\"edui-video-video\" controls=\"\" preload=\"none\" width=\"558\" height=\"297\" src=\"http://localhost:8808/upload_file/2022-12-16/cp3614o8pbc4h2rfas.mp4\" data-setup=\"{}\" style=\"width: 558px; height: 297px;\"><source src=\"http://localhost:8808/upload_file/2022-12-16/cp3614o8pbc4h2rfas.mp4\" type=\"video/mp4\"/></video><br/></p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"2.jpg\",\"url\":\"upload_file/2022-12-15/cp2ajn91urfchlqi1e.jpg\",\"fileType\":\"image/jpeg\",\"size\":18618}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (14, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 16:56:18', '2022-12-15 16:56:18', '2022-12-20 10:02:07', 31, 0, 1, '1', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejse0u2xs.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>流量卡圣诞节拉克丝的拉克斯基的</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (15, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 16:56:19', '2022-12-15 16:56:19', '2022-12-20 10:02:07', 31, 0, 1, '1', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejse0u2xs.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>流量卡圣诞节拉克丝的拉克斯基的</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (16, '刘淼', 25, '4', '2022-12-15 11:03:17', 1, '2022-12-15 16:56:19', '2022-12-15 16:56:19', '2022-12-20 10:02:07', 31, 0, 1, '1', 'upload_file/2022-12-15/cp222yx2yjhshflfgx.jpeg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyjdhcwwgowa.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejsbp173g.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2022-12-15/cp2230vyuejse0u2xs.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>流量卡圣诞节拉克丝的拉克斯基的</p>', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzc58mucz.jpg\",\"fileType\":\"image/jpeg\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzch4dgpt.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"upload_file/2022-12-15/cp2237cm8pzcaun4c4.png\",\"fileType\":\"image/png\",\"size\":105997}]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (17, '大头', 26, '4', '2023-01-06 00:00:00', 1, '2023-01-06 18:06:50', '2023-01-06 18:06:50', NULL, 31, 0, 1, '0', 'upload_file/2023-01-06/cpl0u24gle003jcdpj.jpg', '[{\"name\":\"74595d93db72bc45e5f5161ca35f5995.jpg\",\"url\":\"upload_file/2023-01-06/cpl0u37j3ag4bdc0bi.jpg\",\"fileType\":\"image/jpeg\",\"size\":62462}]', '<p>啊实打实的</p>', '[]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (18, '小刀', 25, '0', '2023-02-17 00:00:00', 2, '2023-02-24 22:36:33', '2023-02-24 22:36:43', NULL, 31, 31, 1, '1,0', 'upload_file/2023-02-24/cqqv9xf8lo6ojn5fjw.jpg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2023-02-24/cqqv9zqenao4bofgej.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2023-02-24/cqqv9zqgp9egz2m0ta.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2023-02-24/cqqv9zqg3sn0m3t1ee.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947}]', '<p>啊实打实大苏打<img src=\"http://localhost:8808/upload_file/2023-02-24/cqqva32mef40kw5tu5.jpg\" alt=\"a333.jpg\"/></p>', '[]', '5', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (19, '啊啊', 0, '', NULL, 1, '2023-05-11 23:02:54', '2023-05-11 23:02:54', NULL, 31, 0, 1, '0', '', '[]', '<p>啊实打实的</p>', '[]', '4', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (20, '赵云', 18, '4', '2023-05-21 00:00:00', 1, '2023-05-21 23:12:29', '2023-05-21 23:12:29', NULL, 31, 0, 1, '0', '', '[]', '', '[]', '', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (21, '张飞', 25, '5', '2023-05-21 00:00:00', 1, '2023-05-21 23:12:47', '2023-05-21 23:12:47', NULL, 31, 0, 1, '0', '', '[]', '', '[]', '', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (22, '关羽', 33, '6', '2023-05-21 23:12:58', 2, '2023-05-21 23:13:03', '2023-05-21 23:13:03', NULL, 31, 0, 0, '1', '', '[]', '', '[]', '', 0, NULL, NULL);
INSERT INTO `demo_gen` VALUES (23, 'aaa', 546, '5', NULL, 1, '2023-05-23 22:04:23', '2023-05-23 22:27:52', NULL, 31, 31, 1, '0,1', '', '[]', '', '[]', '5,6,4', 0, NULL, NULL);

-- ----------------------------
-- Table structure for demo_gen_class
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_class`;
CREATE TABLE `demo_gen_class`  (
                                   `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
                                   `class_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成关联测试表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of demo_gen_class
-- ----------------------------
INSERT INTO `demo_gen_class` VALUES (4, '分类一');
INSERT INTO `demo_gen_class` VALUES (5, '分类二');
INSERT INTO `demo_gen_class` VALUES (6, '分类三');
INSERT INTO `demo_gen_class` VALUES (7, '分类四');

-- ----------------------------
-- Table structure for demo_gen_other
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_other`;
CREATE TABLE `demo_gen_other`  (
                                   `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                   `info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
                                   `img` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '单图',
                                   `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '多图',
                                   `file` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '单文件',
                                   `files` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '多文件',
                                   `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '特殊字段测试' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of demo_gen_other
-- ----------------------------
INSERT INTO `demo_gen_other` VALUES (9, '<p>khash喀什打开就是的埃里克森</p>\n\n<p>的拉卡市的案例咯技术的</p>\n\n<p><img alt=\"\" src=\"http://localhost:8200/pub_upload/2021-08-12/cdhcvg0ow40ofpcy2k.jpg\" style=\"height:286px; width:757px\" /></p>\n', 'pub_upload/2021-08-12/cdhcvhxq38u8xbjcfd.jpg', '[{\"name\":\"5b5fd982ce018.jpg\",\"status\":\"success\",\"uid\":1628753354094,\"url\":\"pub_upload/2021-08-12/cdhcvjhlcdvcolnz7t.jpg\"},{\"name\":\"111.jpg\",\"status\":\"success\",\"uid\":1628753354095,\"url\":\"pub_upload/2021-08-12/cdhcvlav6u3oiwnybe.jpg\"}]', '[{\"name\":\"5b0fac3e20268.jpg\",\"status\":\"success\",\"uid\":1628753358537,\"url\":\"pub_upload/2021-08-12/cdhcvncc4gikxrqnh1.jpg\"}]', '[{\"name\":\"5b0fac3e20268.jpg\",\"status\":\"success\",\"uid\":1628753365149,\"url\":\"pub_upload/2021-08-12/cdhcvow3zq4g9zdqze.jpg\"},{\"name\":\"5b5fd982ce018.jpg\",\"status\":\"success\",\"uid\":1628753365150,\"url\":\"pub_upload/2021-08-12/cdhcvqdnyl9opkyu9p.jpg\"}]', 'asdasdasd');

-- ----------------------------
-- Table structure for demo_gen_tree
-- ----------------------------
DROP TABLE IF EXISTS `demo_gen_tree`;
CREATE TABLE `demo_gen_tree`  (
                                  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
                                  `demo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
                                  `demo_age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
                                  `classes` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '班级',
                                  `demo_born` datetime NULL DEFAULT NULL COMMENT '出生年月',
                                  `demo_gender` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别',
                                  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                                  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                                  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
                                  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                                  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                                  `demo_status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态',
                                  `demo_cate` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
                                  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联人',
                                  `user_ids` json NULL COMMENT '关联人信息',
                                  `depart_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联部门',
                                  `depart_ids` json NULL COMMENT '关联部门信息',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成树形结构测试表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of demo_gen_tree
-- ----------------------------
INSERT INTO `demo_gen_tree` VALUES (1, 0, '张三', 20, '5', '2021-08-23 00:00:00', 1, '2021-08-04 11:54:21', '2024-11-18 09:50:10', NULL, 31, 31, 0, '0', 1, '[2, 1]', 103, '[103, 104]');
INSERT INTO `demo_gen_tree` VALUES (2, 1, '李四', 28, '5', '2021-08-24 00:00:00', 1, '2021-08-04 11:54:38', '2024-11-18 09:50:53', NULL, 31, 31, 0, '0', 3, '[3, 4]', 104, '[104, 105]');
INSERT INTO `demo_gen_tree` VALUES (3, 0, '王五', 63, '5', '2021-08-17 00:00:00', 1, '2021-08-04 11:54:56', '2025-08-14 15:17:00', NULL, 31, 3, 1, '0', 0, '[]', 0, '[]');
INSERT INTO `demo_gen_tree` VALUES (4, 3, '小小', 65, '4', '2021-08-24 00:00:00', 1, '2021-08-04 15:51:55', '2024-10-12 09:31:39', NULL, 31, 31, 0, '0', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (5, 3, '麻花', 23, '5', NULL, 2, '2021-08-04 15:52:13', '2022-12-20 11:33:26', NULL, 31, 31, 1, '0', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (6, 2, '赵六', 232, '4', '2021-08-21 00:00:00', 0, '2021-08-21 12:11:53', '2024-10-08 09:06:19', NULL, 31, 31, 1, '1', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (7, 6, '啊啊啊', 45, '4', '2023-05-10 00:00:00', 1, '2023-05-09 22:30:14', '2023-05-09 22:30:14', '2023-05-09 22:30:18', 31, 0, 1, '1', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (8, 0, '666', 0, '4', NULL, 0, '2023-11-01 23:31:22', '2024-10-08 09:06:26', NULL, 31, 31, 1, '', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (9, 0, '11222', 0, '', NULL, 0, '2023-11-01 23:32:22', '2023-11-02 09:01:03', '2023-11-02 10:37:32', 31, 31, 1, '', 0, NULL, 0, NULL);
INSERT INTO `demo_gen_tree` VALUES (10, 0, '22', 0, '', NULL, 0, '2023-11-01 23:32:31', '2023-11-01 23:32:31', '2023-11-02 09:00:52', 31, 0, 1, '', 0, NULL, 0, NULL);

-- ----------------------------
-- Table structure for demo_snow_id
-- ----------------------------
DROP TABLE IF EXISTS `demo_snow_id`;
CREATE TABLE `demo_snow_id`  (
                                 `id` bigint(20) UNSIGNED NOT NULL COMMENT 'ID',
                                 `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
                                 `age` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '年龄',
                                 `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '照片',
                                 `photos` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '相册',
                                 `files` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文件',
                                 `photos2` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '相册二',
                                 `files2` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文件二',
                                 `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联人',
                                 `user_ids` json NULL COMMENT '关联人信息',
                                 `depart_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联部门',
                                 `depart_ids` json NULL COMMENT '关联部门信息',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '雪花ID测试' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_snow_id
-- ----------------------------
INSERT INTO `demo_snow_id` VALUES (709793997605830657, '张三', 50, 'upload_file/2023-12-14/cxnws7ui0e1cn8prhe.jpg', '[{\"name\":\"t01b97b4bd97190d33a.jpg\",\"url\":\"upload_file/2024-04-08/d0ef78job784jgz2xy.jpg\",\"fileType\":\"image/jpeg\",\"size\":9859},{\"name\":\"333.jpg\",\"url\":\"upload_file/2024-04-08/d0ef8eobw28owj4gnm.jpg\",\"fileType\":\"image/jpeg\",\"size\":4659}]', '[{\"name\":\"5b0fac3e20268.jpg\",\"url\":\"\",\"fileType\":\"\",\"size\":137811},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"\",\"fileType\":\"\",\"size\":84889},{\"name\":\"5b5589ebf0af8.png\",\"url\":\"\",\"fileType\":\"\",\"size\":105997}]', '[{\"name\":\"a222.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4s2szk5tkwpg.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"2.jpg\",\"url\":\"upload_file/2024-11-01/d5anf4es0jj0osvshv.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4o7f5wltzywq.jpg\",\"fileType\":\"\",\"size\":0}]', '[{\"name\":\"a333.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4ub5w0zlzu2d.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"65d3d7ad866394bf86309af1bbba11a3.jpeg\",\"url\":\"upload_file/2024-11-01/d5aqb3pt8oi8ps7hog.jpeg\",\"fileType\":\"\",\"size\":0}]', 2, '[1, 2]', 103, '[103, 105]');
INSERT INTO `demo_snow_id` VALUES (709794007588274177, '李四', 40, '', '[]', '[]', '[]', '[]', 1, '[3, 4, 6]', 103, '[104, 105]');
INSERT INTO `demo_snow_id` VALUES (709798138776387585, '王五', 25, '', NULL, NULL, NULL, NULL, 0, NULL, 0, NULL);
INSERT INTO `demo_snow_id` VALUES (721249471278612481, '赵六', 56, '', NULL, NULL, NULL, NULL, 0, NULL, 0, NULL);
INSERT INTO `demo_snow_id` VALUES (721249471278678017, '田七', 38, '', NULL, NULL, NULL, NULL, 0, NULL, 0, NULL);
INSERT INTO `demo_snow_id` VALUES (721250110037557249, '毛八', 19, '', '[]', '[]', '[]', '[]', 2, '[1, 2, 3]', 0, '[]');
INSERT INTO `demo_snow_id` VALUES (738028591207415809, '阿木', 30, 'upload_file/2024-04-08/d0ednni611n0vjmhxv.jpg', '[{\"name\":\"t01d4a5295ac54f8d00.jpg\",\"url\":\"upload_file/2024-04-08/d0ednob5bv5k8oikbp.jpg\",\"fileType\":\"image/jpeg\",\"size\":32926}]', '[]', '[]', '[]', 2, '[]', 0, '[]');
INSERT INTO `demo_snow_id` VALUES (768480019478478849, '测试', 30, 'upload_file/2024-10-30/d58y6h4ub5w0zlzu2d.jpg', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4o7f5wltzywq.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a222.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4s2szk5tkwpg.jpg\",\"fileType\":\"image/jpeg\",\"size\":25947},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4s2szkdevuyk.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"65d3d7ad866394bf86309af1bbba11a3.jpeg\",\"url\":\"upload_file/2024-11-01/d5aqb3pt8oi8ps7hog.jpeg\",\"fileType\":\"image/jpeg\",\"size\":102522}]', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4o7f5wltzywq.jpg\",\"fileType\":\"image/jpeg\",\"size\":87448},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4s2szkdevuyk.jpg\",\"fileType\":\"image/jpeg\",\"size\":104079},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2024-11-04/d5d1763mgdg0ift6an.jpg\",\"fileType\":\"image/jpeg\",\"size\":84889}]', '[{\"name\":\"a333.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4ub5w0zlzu2d.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"5b5fd982ce018.jpg\",\"url\":\"upload_file/2024-11-04/d5d1763mgdg0ift6an.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"65d3d7ad866394bf86309af1bbba11a3.jpeg\",\"url\":\"upload_file/2024-11-01/d5aqb3pt8oi8ps7hog.jpeg\",\"fileType\":\"\",\"size\":0}]', '[{\"name\":\"00300281425_2fbbb18b.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4o7f5wltzywq.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"2.jpg\",\"url\":\"upload_file/2024-11-01/d5anf4es0jj0osvshv.jpg\",\"fileType\":\"\",\"size\":0},{\"name\":\"a111.jpg\",\"url\":\"upload_file/2024-10-30/d58y6h4s2szkdevuyk.jpg\",\"fileType\":\"\",\"size\":0}]', 4, '[1, 2]', 104, '[103, 104]');

-- ----------------------------
-- Table structure for sys_attachment
-- ----------------------------
DROP TABLE IF EXISTS `sys_attachment`;
CREATE TABLE `sys_attachment`  (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文件ID',
                                   `app_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '应用ID',
                                   `drive` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传驱动',
                                   `name` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件原始名',
                                   `kind` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传类型',
                                   `mime_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '扩展类型',
                                   `path` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '本地路径',
                                   `size` bigint(20) NULL DEFAULT 0 COMMENT '文件大小',
                                   `ext` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展名',
                                   `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'md5校验码',
                                   `created_by` bigint(20) NULL DEFAULT 0 COMMENT '上传人ID',
                                   `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
                                   `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                   `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   INDEX `md5`(`md5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_attachment
-- ----------------------------
INSERT INTO `sys_attachment` VALUES (2, 'system', '0', 'a222.jpg', 'image', 'image/jpeg', 'upload_file/2024-11-18/d5p41p66q7fctdzicz.jpg', 25947, 'jpg', 'f4378e34be8f7a6c311ad068b3498405', 31, 1, '2024-11-18 15:02:39', '2024-11-18 15:02:39');
INSERT INTO `sys_attachment` VALUES (3, 'system', '0', 'a111.jpg', 'image', 'application/octet-stream', 'upload_file/2024-11-18/d5p41ssy6gg4beecej.jpg', 104079, 'jpg', '11c14ce4809ff610e44b1da373500703', 31, 1, '2024-11-18 15:02:47', '2024-11-18 15:02:47');
INSERT INTO `sys_attachment` VALUES (4, 'system', '0', 'a333.jpg', 'image', 'image/jpeg', 'upload_file/2025-08-13/dc14bhjdc1rkyibhei.jpg', 18381, 'jpg', 'b1f4daf376347c18666c4670325cebd7', 31, 1, '2025-08-13 15:23:16', '2025-08-13 15:23:16');

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
                                  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
                                  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
                                  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
                                  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
                                  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件',
                                  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                                  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
                                  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
                                  `is_hide` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
                                  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
                                  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
                                  `is_link` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
                                  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
                                  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
                                  `is_iframe` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否内嵌iframe',
                                  `is_cached` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存',
                                  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由重定向地址',
                                  `is_affix` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否固定',
                                  `link_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '链接地址',
                                  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                                  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `name`(`name`) USING BTREE,
                                  INDEX `pid`(`pid`) USING BTREE,
                                  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 182 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'api/v1/system/auth', '权限管理', 'ele-Stamp', '', '', 0, 30, 0, '/system/auth', 'layout/routerView/parent', 0, '', 0, 0, 1, '0', 0, '', '2022-03-24 15:03:37', '2022-04-14 16:29:19');
INSERT INTO `sys_auth_rule` VALUES (2, 1, 'api/v1/system/auth/menuList', '菜单管理', 'ele-Calendar', '', '', 1, 0, 0, '/system/auth/menuList', 'system/menu/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 17:24:13', '2022-03-29 10:54:49');
INSERT INTO `sys_auth_rule` VALUES (3, 2, 'api/v1/system/menu/add', '添加菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 16:48:43', '2022-03-29 17:05:19');
INSERT INTO `sys_auth_rule` VALUES (4, 2, 'api/v1/system/menu/update', '修改菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 17:04:25', '2022-03-29 18:11:36');
INSERT INTO `sys_auth_rule` VALUES (10, 1, 'api/v1/system/role/list', '角色管理', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0, '/system/auth/roleList', 'system/role/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 18:15:03', '2022-03-30 10:25:34');
INSERT INTO `sys_auth_rule` VALUES (11, 2, 'api/v1/system/menu/delete', '删除菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:10', '2022-04-06 14:49:17');
INSERT INTO `sys_auth_rule` VALUES (12, 10, 'api/v1/system/role/add', '添加角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:46', '2022-04-06 14:49:46');
INSERT INTO `sys_auth_rule` VALUES (13, 10, '/api/v1/system/role/edit', '修改角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:08', '2022-04-06 14:50:08');
INSERT INTO `sys_auth_rule` VALUES (14, 10, '/api/v1/system/role/delete', '删除角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:22', '2022-04-06 14:50:22');
INSERT INTO `sys_auth_rule` VALUES (15, 1, 'api/v1/system/dept/list', '部门管理', 'iconfont icon-siweidaotu', '', '', 1, 0, 0, '/system/auth/deptList', 'system/dept/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:52:23', '2022-04-07 22:59:20');
INSERT INTO `sys_auth_rule` VALUES (16, 17, 'aliyun', '阿里云-iframe', 'iconfont icon-diannao1', '', '', 1, 0, 0, '/demo/outLink/aliyun', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'https://www.aliyun.com/daily-act/ecs/activity_selection?spm=5176.8789780.J_3965641470.5.568845b58KHj51', '2022-04-06 17:26:29', '2022-04-07 15:27:17');
INSERT INTO `sys_auth_rule` VALUES (17, 0, 'outLink', '外链测试', 'iconfont icon-zhongduancanshu', '', '', 0, 20, 0, '/demo/outLink', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 15:20:51', '2022-04-14 16:29:07');
INSERT INTO `sys_auth_rule` VALUES (18, 17, 'tenyun', '腾讯云-外链', 'iconfont icon-shouye_dongtaihui', '', '', 1, 0, 0, '/demo/outLink/tenyun', 'layout/routerView/link', 1, '', 0, 0, 1, '', 0, 'https://cloud.tencent.com/act/new?cps_key=DEMO_CPS_KEY&fromSource=gwzcw.3775555.3775555.3775555&utm_id=gwzcw.3775555.3775555.3775555&utm_medium=cpc', '2022-04-07 15:23:52', '2023-09-28 10:27:41');
INSERT INTO `sys_auth_rule` VALUES (19, 15, 'api/v1/system/dept/add', '添加部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:56:39', '2022-04-07 22:56:39');
INSERT INTO `sys_auth_rule` VALUES (20, 15, 'api/v1/system/dept/edit', '修改部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:00', '2022-04-07 22:57:00');
INSERT INTO `sys_auth_rule` VALUES (21, 15, 'api/v1/system/dept/delete', '删除部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:30', '2022-04-07 22:57:30');
INSERT INTO `sys_auth_rule` VALUES (22, 1, 'api/v1/system/post/list', '岗位管理', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0, '/system/auth/postList', 'system/post/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:58:46', '2022-04-09 14:26:15');
INSERT INTO `sys_auth_rule` VALUES (23, 22, 'api/v1/system/post/add', '添加岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:14:49', '2022-04-09 14:14:49');
INSERT INTO `sys_auth_rule` VALUES (24, 22, 'api/v1/system/post/edit', '修改岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:25', '2022-04-09 14:15:25');
INSERT INTO `sys_auth_rule` VALUES (25, 22, 'api/v1/system/post/delete', '删除岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:47', '2022-04-09 14:15:47');
INSERT INTO `sys_auth_rule` VALUES (26, 1, 'api/v1/system/user/list', '用户管理', 'ele-User', '', '', 1, 0, 0, '/system/auth/user/list', 'system/user/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:19:10', '2022-04-09 14:19:58');
INSERT INTO `sys_auth_rule` VALUES (27, 0, 'api/v1/system/dict', '系统配置', 'iconfont icon-shuxingtu', '', '', 0, 40, 0, '/system/dict', 'layout/routerView/parent', 0, '', 0, 0, 1, '654', 0, '', '2022-04-14 16:28:51', '2022-04-18 14:40:56');
INSERT INTO `sys_auth_rule` VALUES (28, 27, 'api/v1/system/dict/type/list', '字典管理', 'iconfont icon-crew_feature', '', '', 1, 0, 0, '/system/dict/type/list', 'system/dict/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:32:10', '2022-04-16 17:02:50');
INSERT INTO `sys_auth_rule` VALUES (29, 27, 'api/v1/system/dict/dataList', '字典数据管理', 'iconfont icon-putong', '', '', 1, 0, 1, '/system/dict/data/list/:dictType', 'system/dict/dataList', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 12:04:17', '2022-04-18 14:58:43');
INSERT INTO `sys_auth_rule` VALUES (30, 27, 'api/v1/system/config/list', '参数管理', 'ele-Cherry', '', '', 1, 0, 0, '/system/config/list', 'system/config/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 21:05:20', '2022-04-18 21:13:19');
INSERT INTO `sys_auth_rule` VALUES (31, 0, 'api/v1/system/monitor', '系统监控', 'iconfont icon-xuanzeqi', '', '', 0, 30, 0, '/system/monitor', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:40:19', '2022-04-19 10:44:38');
INSERT INTO `sys_auth_rule` VALUES (32, 31, 'api/v1/system/monitor/server', '服务监控', 'iconfont icon-shuju', '', '', 1, 0, 0, '/system/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:43:32', '2022-04-19 10:44:47');
INSERT INTO `sys_auth_rule` VALUES (33, 35, 'api/swagger', 'api文档', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/system/swagger', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'http://localhost:8808/swagger', '2022-04-21 09:23:43', '2022-11-29 17:10:35');
INSERT INTO `sys_auth_rule` VALUES (34, 31, 'api/v1/system/loginLog/list', '登录日志', 'ele-Finished', '', '', 1, 0, 0, '/system/monitor/loginLog', 'system/monitor/loginLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-28 09:59:47', '2022-04-28 09:59:47');
INSERT INTO `sys_auth_rule` VALUES (35, 0, 'api/v1/system/tools', '系统工具', 'iconfont icon-zujian', '', '', 0, 25, 0, '/system/tools', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:29:08', '2022-10-26 10:11:25');
INSERT INTO `sys_auth_rule` VALUES (36, 35, 'api/v1/system/tools/gen/tableList', '代码生成', 'iconfont icon-step', '', '', 1, 0, 0, '/system/tools/gen', 'system/tools/gen/index', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:31:08', '2022-10-31 10:17:23');
INSERT INTO `sys_auth_rule` VALUES (37, 36, 'api/v1/system/tools/gen/columnList', '代码生成配置', 'ele-Edit', '', '', 1, 0, 1, '/system/tools/gen/edit', 'system/tools/gen/component/edit', 0, '', 0, 0, 1, '', 0, '', '2022-10-31 10:11:12', '2022-10-31 10:19:19');
INSERT INTO `sys_auth_rule` VALUES (38, 31, 'api/v1/system/operLog/list', '操作日志', 'iconfont icon-bolangnengshiyanchang', '', '', 1, 0, 0, '/system/monitor/operLog', 'system/monitor/operLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-12-23 16:19:05', '2022-12-23 16:21:50');
INSERT INTO `sys_auth_rule` VALUES (39, 31, 'api/v1/system/online/list', '在线用户', 'iconfont icon-skin', '', '', 1, 0, 0, '/system/monitor/userOnlie', 'system/monitor/userOnline/index', 0, '', 0, 0, 1, '', 0, '', '2023-01-11 15:48:06', '2023-01-11 17:02:39');
INSERT INTO `sys_auth_rule` VALUES (53, 35, 'api/v1/system/sysJob/list', '定时任务', 'fa fa-superpowers', '', '', 1, 0, 0, '/system/sysJob/list', 'system/sysJob/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-01-12 17:51:27');
INSERT INTO `sys_auth_rule` VALUES (54, 53, 'api/v1/system/sysJob/get', '定时任务查询', '', '', '定时任务查询', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (55, 53, 'api/v1/system/sysJob/add', '定时任务添加', '', '', '定时任务添加', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (56, 53, 'api/v1/system/sysJob/edit', '定时任务修改', '', '', '定时任务修改', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (57, 53, 'api/v1/system/sysJob/delete', '定时任务删除', '', '', '定时任务删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (58, 53, 'api/v1/system/sysJob/run', '执行一次', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2023-01-12 18:20:13', '2023-01-12 18:20:13');
INSERT INTO `sys_auth_rule` VALUES (59, 0, 'api/v1/system/sysNotice', '通知公告', 'iconfont icon-fuwenbenkuang', '', '', 0, 0, 0, '/system/sysNotice', 'layout/routerView/parent', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-11-09 15:40:55');
INSERT INTO `sys_auth_rule` VALUES (60, 59, 'api/v1/system/sysNotice/list', '通知公告管理', 'ele-Fold', '', '', 1, 0, 0, '/system/sysNotice/list', 'system/sysNotice/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-11-09 15:41:13');
INSERT INTO `sys_auth_rule` VALUES (61, 60, 'api/v1/system/sysNotice/get', '通知公告查询', '', '', '通知公告查询', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (62, 60, 'api/v1/system/sysNotice/add', '通知公告添加', '', '', '通知公告添加', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (63, 60, 'api/v1/system/sysNotice/edit', '通知公告修改', '', '', '通知公告修改', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (64, 60, 'api/v1/system/sysNotice/delete', '通知公告删除', '', '', '通知公告删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (65, 59, 'api/v1/system/sysNotice/show', '通知公告展示', 'iconfont icon-tongzhi', '', '', 0, 0, 0, '/system/sysNotice/show', 'system/sysNotice/show/index', 0, '', 0, 0, 1, '', 0, '', '2023-12-25 10:34:32', '2024-01-03 10:09:12');
INSERT INTO `sys_auth_rule` VALUES (114, 10, 'api/v1/system/role/dataScope', '授权数据权限', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:18:23', '2024-04-11 11:19:17');
INSERT INTO `sys_auth_rule` VALUES (115, 10, 'api/v1/system/role/setRoleUser', '授权用户权限', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:18:59', '2024-04-11 11:18:59');
INSERT INTO `sys_auth_rule` VALUES (116, 26, 'api/v1/system/user/add', '新增用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:16:14', '2024-04-11 11:16:14');
INSERT INTO `sys_auth_rule` VALUES (117, 26, 'api/v1/system/user/edit', '修改用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:16:45', '2024-04-11 11:19:25');
INSERT INTO `sys_auth_rule` VALUES (118, 26, 'api/v1/system/user/getEdit', '用户信息', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:17:21', '2024-04-11 11:19:30');
INSERT INTO `sys_auth_rule` VALUES (119, 26, 'api/v1/system/user/delete', '删除用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:17:39', '2024-04-11 11:19:34');
INSERT INTO `sys_auth_rule` VALUES (141, 26, 'api/v1/system/user/all', '管理所有', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-09-13 16:57:13', '2024-09-13 16:57:13');
INSERT INTO `sys_auth_rule` VALUES (142, 15, 'api/v1/system/dept/all', '管理所有', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-09-13 16:57:49', '2024-09-13 16:57:49');
INSERT INTO `sys_auth_rule` VALUES (143, 35, 'api/v1/system/sysAttachment/list', '附件管理', 'ele-Folder', '', '', 1, 0, 0, '/system/sysAttachment/list', 'system/sysAttachment/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2024-11-04 14:39:02');
INSERT INTO `sys_auth_rule` VALUES (144, 143, 'api/v1/system/sysAttachment/delete', '附件管理删除', '', '', '附件管理删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
                               `config_id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
                               `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
                               `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
                               `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
                               `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
                               `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
                               `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
                               `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                               `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                               PRIMARY KEY (`config_id`) USING BTREE,
                               UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '文件上传大小限制', NULL, '2021-07-06 14:57:35');
INSERT INTO `sys_config` VALUES (2, '文件上传-文件类型', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png,mp4', 1, 31, 31, '文件上传后缀类型限制', NULL, '2022-12-16 09:52:45');
INSERT INTO `sys_config` VALUES (3, '图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '图片上传后缀类型限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (4, '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '图片上传大小限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (11, '静态资源', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
                             `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
                             `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
                             `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '祖级列表',
                             `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
                             `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
                             `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
                             `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
                             `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                             `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '部门状态（0正常 1停用）',
                             `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
                             `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                             `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                             PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 204 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', '奇讯科技', 0, '[1,2,3]', '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2024-01-29 16:00:28', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '深圳总公司', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', '长沙分公司', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '研发部门', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', '市场部门', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2021-11-04 09:16:38', NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', '测试部门', 3, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', '财务部门', 4, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', '运维部门', 5, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', '市场部门', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', '财务部门', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (200, 100, '', '大数据', 1, '', '18888888888', 'liou@qq.com', 0, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:57', NULL);
INSERT INTO `sys_dept` VALUES (201, 100, '', '开发', 1, NULL, '18888888888', 'li@qq.com', 0, 31, NULL, '2021-07-13 15:56:52', '2022-04-07 22:35:21', NULL);
INSERT INTO `sys_dept` VALUES (202, 108, '', '外勤', 1, NULL, '18888888888', 'aa@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (203, 108, '', '行政', 0, '', '18888888888', 'aa@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:47', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
                                  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
                                  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
                                  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典标签',
                                  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典键值',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
                                  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
                                  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
                                  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（1正常 0停用）',
                                  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
                                  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                                  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                                  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 123 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, '备注信息', '2022-04-18 16:46:22', '2025-06-06 16:08:23');
INSERT INTO `sys_dict_data` VALUES (2, 0, '女', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, '备注信息', NULL, '2025-06-06 16:08:23');
INSERT INTO `sys_dict_data` VALUES (3, 0, '保密', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, '备注信息', NULL, '2025-06-06 16:08:23');
INSERT INTO `sys_dict_data` VALUES (24, 0, '频道页', '1', 'cms_category_type', '', '', 0, 1, 31, 31, '作为频道页，不可作为栏目发布文章，可添加下级分类', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (25, 0, '发布栏目', '2', 'cms_category_type', '', '', 0, 1, 31, 31, '作为发布栏目，可添加文章', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (26, 0, '跳转栏目', '3', 'cms_category_type', '', '', 0, 1, 31, 31, '不直接发布内容，用于跳转页面', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (27, 0, '单页栏目', '4', 'cms_category_type', '', '', 0, 1, 31, 31, '单页面模式，分类直接显示为文章', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (28, 0, '正常', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '暂停', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '默认', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '系统', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, '成功', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, '2025-06-09 08:59:59');
INSERT INTO `sys_dict_data` VALUES (33, 0, '失败', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, '2025-06-09 08:59:59');
INSERT INTO `sys_dict_data` VALUES (34, 0, '成功', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (35, 0, '失败', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '重复执行', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, '执行一次', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '显示', '0', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (39, 0, '隐藏', '1', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (40, 0, '正常', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (41, 0, '停用', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (49, 0, '是', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (50, 0, '否', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (51, 0, '已发布', '1', 'cms_article_pub_type', '', '', 1, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (54, 0, '未发布', '0', 'cms_article_pub_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (55, 0, '置顶', '1', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (56, 0, '推荐', '2', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (57, 0, '普通文章', '0', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (58, 0, '跳转链接', '1', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (59, 0, 'cms模型', '6', 'cms_cate_models', '', '', 0, 1, 1, 1, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (62, 0, '系统后台', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, '2025-06-09 11:04:59');
INSERT INTO `sys_dict_data` VALUES (63, 0, '政务工作', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, '2025-06-09 11:04:59');
INSERT INTO `sys_dict_data` VALUES (64, 0, '幻灯', '3', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (65, 0, '[work]测试业务表', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, '2025-06-09 11:02:08');
INSERT INTO `sys_dict_data` VALUES (66, 0, '回退修改', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:16');
INSERT INTO `sys_dict_data` VALUES (67, 0, '保存中', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:16');
INSERT INTO `sys_dict_data` VALUES (68, 0, '流程中', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:16');
INSERT INTO `sys_dict_data` VALUES (69, 0, '审批通过', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, '2025-06-09 11:04:16');
INSERT INTO `sys_dict_data` VALUES (70, 2, '发布栏目', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (71, 3, '跳转栏目', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (72, 4, '单页栏目', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (73, 2, '置顶', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (74, 3, '幻灯', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (75, 4, '推荐', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (76, 1, '一般', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (77, 1, '频道页', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (78, 0, '普通', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:27');
INSERT INTO `sys_dict_data` VALUES (79, 0, '加急', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:27');
INSERT INTO `sys_dict_data` VALUES (80, 0, '紧急', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:04:27');
INSERT INTO `sys_dict_data` VALUES (81, 0, '特急', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2025-06-09 11:04:27');
INSERT INTO `sys_dict_data` VALUES (82, 0, '频道页', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (83, 0, '发布栏目', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (84, 0, '跳转栏目', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (85, 0, '单页栏目', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (87, 0, '[cms]文章表', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, '2025-06-09 11:02:08');
INSERT INTO `sys_dict_data` VALUES (91, 0, '测试一下', '666', 'cms_article_type', '', '', 0, 1, 31, 0, '', '2021-08-03 17:04:12', '2021-08-03 17:04:12');
INSERT INTO `sys_dict_data` VALUES (92, 0, '缓存测试222', '33333', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:16:45', '2021-08-03 17:19:41');
INSERT INTO `sys_dict_data` VALUES (93, 0, '缓存测试222', '11111', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:26:14', '2021-08-03 17:26:26');
INSERT INTO `sys_dict_data` VALUES (99, 0, '不显示', 'none', 'cms_nav_position', '', '', 1, 1, 22, 0, '', '2021-08-31 15:37:35', '2021-08-31 15:37:35');
INSERT INTO `sys_dict_data` VALUES (100, 0, '顶部导航', 'top', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:37:57', '2021-08-31 15:37:57');
INSERT INTO `sys_dict_data` VALUES (101, 0, '底部导航', 'bottom', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:38:08', '2021-08-31 15:38:08');
INSERT INTO `sys_dict_data` VALUES (102, 0, '读取', 'GET', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:10', '2022-12-23 19:03:02');
INSERT INTO `sys_dict_data` VALUES (103, 0, '新增', 'POST', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:22', '2022-12-23 19:03:10');
INSERT INTO `sys_dict_data` VALUES (104, 0, '修改', 'PUT', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:32', '2022-12-23 19:03:19');
INSERT INTO `sys_dict_data` VALUES (105, 0, '删除', 'DELETE', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:44', '2022-12-23 19:03:27');
INSERT INTO `sys_dict_data` VALUES (106, 0, '无标签', '0', 'notice_tag', '', '', 0, 1, 31, 31, '', '2023-12-28 15:48:45', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (107, 0, '提醒', '1', 'notice_tag', '', '', 0, 1, 31, 31, '', '2023-12-28 15:48:54', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (108, 0, '一般', '2', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:35', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (109, 0, '次要', '3', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:44', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (110, 0, '重要', '4', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:53', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (111, 0, '紧急', '5', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:53:01', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_data` VALUES (112, 0, '本地上传', '0', 'sys_upload_drive', '', '', 0, 1, 31, 0, '', '2024-10-23 14:37:27', '2024-10-23 14:37:27');
INSERT INTO `sys_dict_data` VALUES (113, 0, '腾讯云', '1', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:37:38', '2024-10-23 14:38:05');
INSERT INTO `sys_dict_data` VALUES (114, 0, '七牛云', '2', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:37:52', '2024-10-23 14:38:43');
INSERT INTO `sys_dict_data` VALUES (115, 0, '阿里云', '3', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:38:11', '2024-10-23 14:38:49');
INSERT INTO `sys_dict_data` VALUES (116, 0, '图片', 'image', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:18', '2024-10-23 14:54:18');
INSERT INTO `sys_dict_data` VALUES (117, 0, '文档', 'doc', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:32', '2024-10-23 14:54:32');
INSERT INTO `sys_dict_data` VALUES (118, 0, '音频', 'audio', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:55', '2024-10-23 14:54:55');
INSERT INTO `sys_dict_data` VALUES (119, 0, '视频', 'video', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:09', '2024-10-23 14:55:09');
INSERT INTO `sys_dict_data` VALUES (120, 0, '压缩包', 'zip', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:22', '2024-10-23 14:55:22');
INSERT INTO `sys_dict_data` VALUES (121, 0, '其它', 'other', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:40', '2024-10-23 14:55:40');
INSERT INTO `sys_dict_data` VALUES (123, 0, '允许砍价', '1', 'shop_bargain_switch', '', '', 0, 1, 31, 0, '', '2025-09-22 16:28:42', '2025-09-22 16:28:42');
INSERT INTO `sys_dict_data` VALUES (124, 0, '禁止砍价', '0', 'shop_bargain_switch', '', '', 1, 1, 31, 0, '', '2025-09-22 16:29:00', '2025-09-22 16:29:00');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
                                  `dict_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
                                  `pid` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
                                  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典名称',
                                  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
                                  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
                                  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
                                  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
                                  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                                  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
                                  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
                                  PRIMARY KEY (`dict_id`) USING BTREE,
                                  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, 54, '用户性别', 'sys_user_sex', 1, 31, 31, '用于选择用户性别', NULL, '2025-06-06 16:08:23');
INSERT INTO `sys_dict_type` VALUES (2, 55, '分类类型', 'cms_category_type', 1, 31, 3, '文章分类类型', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_type` VALUES (3, 54, '任务状态', 'sys_job_status', 1, 31, 31, '任务状态列表', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, 54, '任务分组', 'sys_job_group', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (14, 54, '管理员登录状态', 'admin_login_status', 1, 31, 31, '', NULL, '2025-06-09 08:59:59');
INSERT INTO `sys_dict_type` VALUES (15, 54, '操作日志状态', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (16, 54, '任务策略', 'sys_job_policy', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (17, 54, '菜单状态', 'sys_show_hide', 1, 31, 0, '菜单状态', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (18, 54, '系统开关', 'sys_normal_disable', 1, 31, 31, '系统开关', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (24, 54, '系统内置', 'sys_yes_no', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (25, 55, '文章发布状态', 'cms_article_pub_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (26, 55, '文章附加状态', 'cms_article_attr', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (27, 55, '文章类型', 'cms_article_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (28, 55, '文章栏目模型分类', 'cms_cate_models', 1, 1, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (30, 54, '菜单模块类型', 'menu_module_type', 1, 2, 31, '', NULL, '2025-06-09 11:04:59');
INSERT INTO `sys_dict_type` VALUES (31, 56, '工作流程类型', 'flow_type', 1, 2, 31, '', NULL, '2025-06-09 11:02:08');
INSERT INTO `sys_dict_type` VALUES (32, 56, '工作流程审批状态', 'flow_status', 1, 31, 31, '工作流程审批状态', NULL, '2025-06-09 11:04:16');
INSERT INTO `sys_dict_type` VALUES (33, 54, '博客分类类型', 'sys_blog_type', 1, 31, 31, '博客分类中的标志', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (34, 54, '博客日志标志', 'sys_log_sign', 1, 31, 0, '博客日志管理中的标志数据字典', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (35, 56, '工作流紧急状态', 'flow_level', 1, 31, 31, '', NULL, '2025-06-09 11:04:27');
INSERT INTO `sys_dict_type` VALUES (49, 55, 'CMS栏目导航位置', 'cms_nav_position', 1, 22, 0, '', '2021-08-31 15:37:04', '2021-08-31 15:37:04');
INSERT INTO `sys_dict_type` VALUES (50, 54, '操作日志类型', 'sys_oper_log_type', 1, 31, 0, '', '2022-12-21 11:55:02', '2022-12-21 11:55:02');
INSERT INTO `sys_dict_type` VALUES (51, 54, '系统公告标签', 'notice_tag', 1, 31, 31, '', '2023-12-28 15:48:03', '2025-06-09 11:04:35');
INSERT INTO `sys_dict_type` VALUES (52, 54, '附件上传驱动', 'sys_upload_drive', 1, 31, 31, '', '2024-10-23 14:36:17', '2024-10-23 14:36:30');
INSERT INTO `sys_dict_type` VALUES (53, 54, '上传文件类型', 'sys_upload_file_type', 1, 31, 0, '', '2024-10-23 14:53:50', '2024-10-23 14:53:50');
INSERT INTO `sys_dict_type` VALUES (54, 0, '系统字典', 'sys', 1, 31, 31, '', '2025-06-06 15:57:36', '2025-06-09 09:14:45');
INSERT INTO `sys_dict_type` VALUES (55, 0, 'cms字典', 'cms', 1, 31, 0, '', '2025-06-09 08:50:41', '2025-06-09 08:50:41');
INSERT INTO `sys_dict_type` VALUES (56, 0, '工作流字典', 'flow', 1, 31, 31, '', '2025-06-09 11:01:37', '2025-06-09 11:01:59');
INSERT INTO `sys_dict_type` VALUES (57, 0, '商城', 'shop', 1, 31, 0, '', '2025-09-22 16:25:07', '2025-09-22 16:25:07');
INSERT INTO `sys_dict_type` VALUES (58, 57, '商品砍价', 'shop_bargain_switch', 1, 31, 0, '', '2025-09-22 16:25:50', '2025-09-22 16:25:50');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
                            `job_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务ID',
                            `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
                            `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数',
                            `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
                            `invoke_target` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
                            `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
                            `misfire_policy` tinyint(4) NULL DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
                            `concurrent` tinyint(4) NULL DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
                            `status` tinyint(4) NULL DEFAULT 0 COMMENT '状态（0正常 1暂停）',
                            `created_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
                            `updated_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
                            `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注信息',
                            `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                            `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                            PRIMARY KEY (`job_id`) USING BTREE,
                            UNIQUE INDEX `invoke_target`(`invoke_target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '测试任务1', '', 'DEFAULT', 'test1', '* * * * * ?', 1, 0, 1, 1, 31, '', '2021-07-16 16:01:59', '2023-05-29 17:06:22');
INSERT INTO `sys_job` VALUES (2, '测试任务2', 'hello|gfast', 'DEFAULT', 'test2', '* * * * * ?', 1, 0, 1, 1, 31, '备注', '2021-07-16 17:15:09', '2021-07-16 17:15:09');
INSERT INTO `sys_job` VALUES (8, '在线用户定时更新', '', 'DEFAULT', 'checkUserOnline', '5 */10 * * * ?', 1, 0, 0, 2, 1, '', '2021-07-19 08:57:24', '2021-07-19 08:57:24');

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
                                `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
                                `target_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '方法名',
                                `created_at` datetime NULL DEFAULT NULL COMMENT '执行日期',
                                `result` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '执行结果',
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 663 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------
INSERT INTO `sys_job_log` VALUES (1, 'checkUserOnline', '2024-04-12 17:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (2, 'checkUserOnline', '2024-04-12 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (3, 'checkUserOnline', '2024-04-12 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (4, 'checkUserOnline', '2024-04-12 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (5, 'checkUserOnline', '2024-04-12 18:00:06', '在线用户定时更新，执行成功');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
                                  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
                                  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
                                  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
                                  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
                                  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
                                  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
                                  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
                                  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
                                  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
                                  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
                                  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
                               `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
                               `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
                               `type` bigint(20) NOT NULL COMMENT '类型',
                               `tag` int(11) NULL DEFAULT NULL COMMENT '标签',
                               `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
                               `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
                               `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
                               `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态',
                               `created_by` bigint(20) NULL DEFAULT NULL COMMENT '发送人',
                               `updated_by` bigint(20) NULL DEFAULT 0 COMMENT '修改人',
                               `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                               `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                               `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                               `receiver` json NULL COMMENT '接收者（私信）',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知公告' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
INSERT INTO `sys_notice` VALUES (1, '测试001', 1, 2, '<p>666666666</p>', '', 0, 1, 31, 31, '2024-01-02 17:46:59', '2024-05-06 17:46:42', NULL, NULL);
INSERT INTO `sys_notice` VALUES (2, '测试私信', 2, 3, '<p>888888888888</p>', '', 0, 1, 31, 31, '2024-01-02 17:47:36', '2025-06-12 08:59:51', NULL, '[2, 3, 31]');
INSERT INTO `sys_notice` VALUES (6, 'demo', 2, 2, '<p>888</p>', '', 0, 1, 31, 0, '2025-06-12 09:00:32', '2025-06-12 09:00:32', NULL, '[31]');

-- ----------------------------
-- Table structure for sys_notice_read
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice_read`;
CREATE TABLE `sys_notice_read`  (
                                    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
                                    `notice_id` bigint(20) NOT NULL COMMENT '信息id',
                                    `user_id` bigint(20) NOT NULL COMMENT '用户id',
                                    `clicks` int(11) NULL DEFAULT NULL COMMENT '点击次数',
                                    `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                    `created_at` datetime NULL DEFAULT NULL COMMENT '阅读时间',
                                    PRIMARY KEY (`id`) USING BTREE,
                                    UNIQUE INDEX `notice_id`(`notice_id`, `user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '已读记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice_read
-- ----------------------------
INSERT INTO `sys_notice_read` VALUES (1, 1, 31, 0, '2024-01-02 17:47:05', '2024-01-02 17:47:05');
INSERT INTO `sys_notice_read` VALUES (2, 2, 31, 0, '2025-06-12 09:45:15', '2025-06-12 09:45:15');

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
                                 `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
                                 `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
                                 `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
                                 `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
                                 `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
                                 `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
                                 `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
                                 `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
                                 `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
                                 `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
                                 `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
                                 `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
                                 `error_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误消息',
                                 `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
                                 PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (1, '', 0, '/api/v1/system/operLog/clear', 'DELETE', 1, 'demo', '财务部门', '/api/v1/system/operLog/clear', '::1', '内网IP', '{}', NULL, '2025-08-15 11:08:52');
INSERT INTO `sys_oper_log` VALUES (2, '操作日志', 0, '/api/v1/system/operLog/list', 'GET', 1, 'demo', '财务部门', '/api/v1/system/operLog/list?pageNum=1&pageSize=10', '::1', '内网IP', '{\"pageNum\":\"1\",\"pageSize\":\"10\"}', NULL, '2025-08-15 11:08:52');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
                             `post_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
                             `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码',
                             `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
                             `post_sort` int(4) NOT NULL COMMENT '显示顺序',
                             `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
                             `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
                             `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
                             `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
                             `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                             PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位信息表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, 1, '', 0, 0, '2021-07-11 11:32:58', NULL, NULL);
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理', 2, 1, '', 0, 0, '2021-07-12 11:01:26', NULL, NULL);
INSERT INTO `sys_post` VALUES (3, 'hr', '人力资源', 3, 1, '', 0, 31, '2021-07-12 11:01:30', '2022-09-16 16:48:18', NULL);
INSERT INTO `sys_post` VALUES (4, 'user', '普通员工', 4, 0, '普通员工', 0, 31, '2021-07-12 11:01:33', '2022-04-08 15:32:23', NULL);
INSERT INTO `sys_post` VALUES (5, 'it', 'IT部', 5, 1, '信息部', 31, 31, '2021-07-12 11:09:42', '2022-04-09 12:59:12', NULL);
INSERT INTO `sys_post` VALUES (6, '1111', '1111', 0, 1, '11111', 31, 0, '2022-04-08 15:32:44', '2022-04-08 15:32:44', '2022-04-08 15:51:24');
INSERT INTO `sys_post` VALUES (7, '222', '2222', 0, 1, '22222', 31, 0, '2022-04-08 15:32:55', '2022-04-08 15:32:55', '2022-04-08 15:51:24');
INSERT INTO `sys_post` VALUES (8, '33333', '3333', 0, 0, '33333', 31, 0, '2022-04-08 15:33:01', '2022-04-08 15:33:01', '2022-04-08 15:51:40');
INSERT INTO `sys_post` VALUES (9, '222', '111', 0, 1, '2313213', 31, 0, '2022-04-08 15:52:53', '2022-04-08 15:52:53', '2022-04-08 15:52:56');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
                             `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
                             `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
                             `list_order` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
                             `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
                             `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                             `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                             `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加人',
                             `effectiveTime` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '角色有效日期',
                             PRIMARY KEY (`id`) USING BTREE,
                             INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 0, 1, 0, '超级管理员', '备注', 2, '2022-04-01 11:38:39', '2023-09-28 10:27:55', 0, NULL);
INSERT INTO `sys_role` VALUES (2, 1, 1, 0, '普通管理员', '备注', 5, '2022-04-01 11:38:39', '2025-08-14 14:52:32', 0, '{\"effectiveType\":0,\"weekDay\":[1,2,3,4,5],\"dayRange\":[\"2024-04-12 08:00:00\",\"2024-04-12 18:00:00\"],\"dateRange\":null}');
INSERT INTO `sys_role` VALUES (3, 0, 1, 0, '站点管理员', '站点管理人员', 3, '2022-04-01 11:38:39', '2022-04-01 11:38:39', 0, NULL);
INSERT INTO `sys_role` VALUES (4, 5, 1, 0, '初级管理员', '初级管理员', 3, '2022-04-01 11:38:39', '2024-03-18 10:16:15', 0, '{\"effectiveType\":0,\"weekDay\":null,\"dayRange\":null,\"dateRange\":null}');
INSERT INTO `sys_role` VALUES (5, 0, 1, 0, '高级管理员', '高级管理员', 2, '2022-04-01 11:38:39', '2022-04-01 11:38:39', 0, NULL);
INSERT INTO `sys_role` VALUES (8, 0, 1, 0, '区级管理员', '', 2, '2022-04-01 11:38:39', '2022-04-06 09:53:40', 0, NULL);
INSERT INTO `sys_role` VALUES (9, 0, 1, 0, '测试', '', 3, '2023-04-22 12:39:13', '2023-09-28 15:48:56', 3, NULL);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
                                  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
                                  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
                                  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (1, 101);
INSERT INTO `sys_role_dept` VALUES (1, 103);
INSERT INTO `sys_role_dept` VALUES (1, 104);
INSERT INTO `sys_role_dept` VALUES (1, 105);
INSERT INTO `sys_role_dept` VALUES (1, 106);
INSERT INTO `sys_role_dept` VALUES (1, 107);
INSERT INTO `sys_role_dept` VALUES (5, 103);
INSERT INTO `sys_role_dept` VALUES (5, 104);
INSERT INTO `sys_role_dept` VALUES (5, 105);
INSERT INTO `sys_role_dept` VALUES (8, 105);
INSERT INTO `sys_role_dept` VALUES (8, 106);

-- ----------------------------
-- Table structure for sys_role_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_scope`;
CREATE TABLE `sys_role_scope`  (
                                   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                   `role_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
                                   `menu_id` int(11) NOT NULL COMMENT 'api接口id',
                                   `data_scope` int(11) NOT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
                                   `dept_ids` json NULL COMMENT '扩展数据',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   UNIQUE INDEX `role_id`(`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 99 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色数据权限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_scope
-- ----------------------------
INSERT INTO `sys_role_scope` VALUES (80, 2, 59, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (81, 2, 60, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (82, 2, 61, 5, '[101, 103, 104, 105, 106, 107]');
INSERT INTO `sys_role_scope` VALUES (83, 2, 62, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (84, 2, 63, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (85, 2, 64, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (86, 2, 65, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (87, 2, 120, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (88, 2, 121, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (89, 2, 122, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (90, 2, 123, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (91, 2, 124, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (92, 2, 125, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (93, 2, 26, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (94, 2, 116, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (95, 2, 117, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (96, 2, 118, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (97, 2, 119, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (98, 2, 141, 4, '[]');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
                             `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
                             `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
                             `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
                             `birthday` int(11) NOT NULL DEFAULT 0 COMMENT '生日',
                             `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
                             `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
                             `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
                             `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
                             `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
                             `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
                             `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
                             `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
                             `is_admin` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
                             `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
                             `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
                             `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
                             `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
                             `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
                             `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
                             `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                             `open_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信open id',
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
                             UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
                             INDEX `user_nickname`(`user_nickname`) USING BTREE,
                             INDEX `open_id`(`open_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '13578342363', '超级管理员', 0, 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '::1', '2023-10-31 11:22:06', '2021-06-22 17:58:00', '2023-04-22 14:39:18', NULL, '');
INSERT INTO `sys_user` VALUES (2, 'yixiaohu', '13699885599', '奈斯', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'yxh@qq.com', 1, 'upload_file/2022-11-04/co3e5ljknns8jhlp8s.jpg', 102, '备注', 1, '', '', '::1', '2022-11-04 09:54:56', '2021-06-22 17:58:00', '2022-11-04 17:54:56', NULL, '');
INSERT INTO `sys_user` VALUES (3, 'zs', '16399669855', '张三', 0, '41e3778c20338f4d7d6cc886fd3b2a52', 'redoHIj524', 1, 'zs@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg', 101, '', 1, '', '', '::1', '2025-08-14 11:43:40', '2021-06-22 17:58:00', '2024-09-14 09:08:10', NULL, '');
INSERT INTO `sys_user` VALUES (4, 'qlgl', '13758596696', '测试c', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'qlgl@qq.com', 0, '', 102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-11-03 15:44:20', NULL, '');
INSERT INTO `sys_user` VALUES (5, 'test', '13845696696', '测试2', 0, '5f79ce5a3804f28e5c2c1162aaa66c86', 'OsSecLdLvK', 1, '123@qq.com', 0, '', 101, '', 0, '', '', '::1', '2023-10-26 14:59:59', '2021-06-22 17:58:00', '2024-07-22 11:35:51', NULL, '');
INSERT INTO `sys_user` VALUES (6, '18999998889', '13755866654', '刘大大', 0, '5df78d20315a5af61f45d20f72c184fc', 'lC6OoXDCbM', 1, '1223@qq.com', 0, '', 103, '', 1, '', '', '[::1]', '2022-02-25 14:29:22', '2021-06-22 17:58:00', '2022-11-03 17:05:07', NULL, '');
INSERT INTO `sys_user` VALUES (7, 'zmm', '13788566696', '张明明', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11123@qq.com', 0, '', 104, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:42', NULL, '');
INSERT INTO `sys_user` VALUES (8, 'lxx', '13756566696', '李小小', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123333@qq.com', 0, '', 101, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2024-07-22 11:35:44', NULL, '');
INSERT INTO `sys_user` VALUES (10, 'xmm', '13588999969', '小秘密', 0, '2de2a8df703bfc634cfda2cb2f6a59be', 'Frz7LJY7SE', 1, '696@qq.com', 0, '', 102, '', 1, '', '', '::1', '2024-09-14 10:02:18', '2021-06-22 17:58:00', '2024-09-14 09:16:27', NULL, '');
INSERT INTO `sys_user` VALUES (14, 'cd_19', '13699888899', '看金利科技', 0, '1169d5fe4119fd4277a95f02d7036171', '7paigEoedh', 1, '', 0, '', 102, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2022-04-12 18:13:22', '2024-07-22 11:35:58', '');
INSERT INTO `sys_user` VALUES (15, 'lmm', '13587754545', '刘敏敏', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@coc.com', 0, '', 102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2024-09-14 09:20:21', NULL, '');
INSERT INTO `sys_user` VALUES (16, 'ldn', '13899658874', '李大牛', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@ll.con', 0, '', 102, '', 1, '', '', '::1', '2024-05-07 10:09:35', '2021-06-22 17:58:00', '2024-05-07 10:09:58', NULL, '');
INSERT INTO `sys_user` VALUES (20, 'dbc', '13877555566', '大百词', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (22, 'yxfmlbb', '15969423326', '大数据部门测试', 0, '66f89b40ee4a10aabaf70c15756429ea', 'mvd2OtUe8f', 1, 'yxh6691@qq.com', 0, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-09-29/cem20k3fdciosy7nwo.jpeg', 200, '', 1, '2222233', '1222', '[::1]', '2021-10-28 11:36:07', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (23, 'wangming', '18687460581', '王明', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '::1', '2023-11-03 16:03:56', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, 'DEMO_WECHAT_OPEN_ID');
INSERT INTO `sys_user` VALUES (24, 'zhk', '13699885591', '综合科', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '192.168.0.146', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (28, 'demo3', '18699888855', '测试账号1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123132@qq.com', 0, '', 109, '', 1, '', '', '192.168.0.229', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (31, 'demo', '15334455789', '李四', 0, '6dd68eea81e0fca319add0bd58c3fdf6', '46PvWe1Sl7', 1, '123@qq.com', 2, 'upload_file/2024-07-29/d31lsrb6o3rgmrdiih.jpg', 109, '3', 1, '云南省曲靖市22223', '生活变的再糟糕，也不妨碍我变得更好', '::1', '2025-08-15 10:46:14', '2021-06-22 17:58:00', '2024-07-29 08:41:22', NULL, '');
INSERT INTO `sys_user` VALUES (32, 'demo100', '18699888859', '测试账号1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '[::1]', '2021-11-24 18:01:21', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (33, 'demo110', '18699888853', '测试账号1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (34, 'yxfmlbb2', '15969423327', '研发部门测试', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '1111@qqq.com', 1, '', 103, '', 0, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (35, 'wk666', '18888888888', 'wk', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '396861557@qq.com', 1, '', 100, '', 1, '', '', '[::1]', '2021-12-09 14:52:37', '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (36, 'zxd', '13699885565', '张晓东', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'zxk@qq.com', 1, '', 201, '666', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', NULL, '');
INSERT INTO `sys_user` VALUES (37, 'yxfmlbb3', '13513513511', '张三', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '111@qq.com', 0, '', 204, '', 1, '', '', '[::1]', '2021-07-26 14:49:25', '2021-06-22 17:58:00', '2021-07-26 14:49:18', NULL, '');
INSERT INTO `sys_user` VALUES (38, 'test_user', '18888888880', 'test', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11@qq.com', 1, '', 200, '111', 0, '', '', '', NULL, '2021-06-22 17:58:00', '2021-07-12 22:05:29', NULL, '');
INSERT INTO `sys_user` VALUES (39, 'asan', '18687460555', '阿三', 0, '2354837137115700e2adf870ac113dcf', 'drdDvbtYZW', 1, '456654@qq.com', 1, '', 201, '666666', 1, '', '', '', NULL, '2021-07-12 17:21:43', '2021-07-12 21:13:31', '2021-07-12 22:00:44', '');
INSERT INTO `sys_user` VALUES (40, 'asi', '13655888888', '啊四', 0, 'fbb755b35d48759dad47bb1540249fd1', '9dfUstcxrz', 1, '5464@qq.com', 1, '', 201, 'adsaasd', 1, '', '', '', NULL, '2021-07-12 17:46:27', '2021-07-12 21:29:41', '2021-07-12 22:00:44', '');
INSERT INTO `sys_user` VALUES (41, 'awu', '13578556546', '阿五', 0, '3b36a96afa0dfd66aa915e0816e0e9f6', '9gHRa9ho4U', 0, '132321@qq.com', 1, '', 201, 'asdasdasd', 1, '', '', '', NULL, '2021-07-12 17:54:31', '2021-07-12 21:46:34', '2021-07-12 21:59:56', '');
INSERT INTO `sys_user` VALUES (42, 'demo01', '13699888556', '测试01222', 0, '048dc94116558fb40920f3553ecd5fe8', 'KiVrfzKJQx', 1, '456@qq.com', 2, '', 109, '测试用户', 1, '', '', '', NULL, '2022-04-12 16:15:23', '2022-04-12 17:54:49', NULL, '');
INSERT INTO `sys_user` VALUES (43, 'lx', '13855668855', '李希', 0, 'a1e3fe5983811d344950c04902ee6b76', 'Ezcs0vshdf', 1, '', 0, '', 101, '', 1, '', '', '::1', '2023-06-14 11:03:46', '2023-06-14 11:01:47', '2023-06-14 11:01:47', NULL, '');

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
                                    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                                    `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
                                    `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
                                    `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
                                    `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
                                    `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
                                    `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
                                    `os` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
                                    PRIMARY KEY (`id`) USING BTREE,
                                    UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (1, '8594e9fea2a11eb43afb4664276e7ff6', 'DEMO_TOKEN_REDACTED', '2025-08-15 10:46:14', 'demo', '::1', 'Chrome', 'Windows 10');

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
                                  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
                                  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
                                  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 2);
INSERT INTO `sys_user_post` VALUES (1, 3);
INSERT INTO `sys_user_post` VALUES (2, 1);
INSERT INTO `sys_user_post` VALUES (2, 2);
INSERT INTO `sys_user_post` VALUES (3, 2);
INSERT INTO `sys_user_post` VALUES (4, 1);
INSERT INTO `sys_user_post` VALUES (5, 2);
INSERT INTO `sys_user_post` VALUES (10, 1);
INSERT INTO `sys_user_post` VALUES (10, 2);
INSERT INTO `sys_user_post` VALUES (10, 3);
INSERT INTO `sys_user_post` VALUES (10, 4);
INSERT INTO `sys_user_post` VALUES (10, 5);
INSERT INTO `sys_user_post` VALUES (15, 4);
INSERT INTO `sys_user_post` VALUES (16, 2);
INSERT INTO `sys_user_post` VALUES (22, 1);
INSERT INTO `sys_user_post` VALUES (22, 2);
INSERT INTO `sys_user_post` VALUES (31, 2);
INSERT INTO `sys_user_post` VALUES (34, 1);
INSERT INTO `sys_user_post` VALUES (35, 2);
INSERT INTO `sys_user_post` VALUES (35, 3);
INSERT INTO `sys_user_post` VALUES (36, 1);
INSERT INTO `sys_user_post` VALUES (37, 3);
INSERT INTO `sys_user_post` VALUES (38, 2);
INSERT INTO `sys_user_post` VALUES (38, 3);
INSERT INTO `sys_user_post` VALUES (42, 2);
INSERT INTO `sys_user_post` VALUES (42, 3);

-- ----------------------------
-- Table structure for tools_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table`;
CREATE TABLE `tools_gen_table`  (
                                    `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
                                    `group` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '数据库group',
                                    `table_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表名称',
                                    `table_comment` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '表描述',
                                    `class_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '实体类名称',
                                    `tpl_category` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
                                    `package_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成包路径',
                                    `module_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成模块名',
                                    `business_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成业务名',
                                    `function_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能名',
                                    `function_author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
                                    `options` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '其它生成选项',
                                    `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
                                    `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
                                    `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '备注',
                                    `overwrite` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否覆盖原有文件',
                                    `sort_column` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '排序字段名',
                                    `sort_type` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'asc' COMMENT '排序方式 (asc顺序 desc倒序)',
                                    `show_detail` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否有查看详情功能',
                                    `excel_port` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否有导出excel功能',
                                    `use_snow_id` bit(1) NOT NULL DEFAULT b'0' COMMENT '主键是否雪花ID',
                                    `use_virtual` bit(1) NOT NULL DEFAULT b'0' COMMENT '树表是否使用虚拟表',
                                    `excel_imp` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否导入excel',
                                    `overwrite_info` json NULL COMMENT '生成覆盖的文件',
                                    `menu_pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级菜单ID',
                                    PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 113 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of tools_gen_table
-- ----------------------------
INSERT INTO `tools_gen_table` VALUES (91, 'default', 'demo_gen', '代码生成测试表', 'DemoGen', 'crud', 'internal/app/demo/n1/n2', 'n2', 'demo_gen', '代码生成测试', 'gfast', '', '2022-11-01 17:27:43', '2025-08-14 17:42:42', '', b'1', 'id', 'asc', b'1', b'1', b'0', b'0', b'1', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": false}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (92, 'default', 'demo_gen_class', '代码生成关联测试表', 'DemoGenClass', 'crud', 'internal/app/demo/n1/n2', 'n2', 'demo_gen_class', '分类信息', 'gfast', '', '2022-11-03 06:36:57', '2025-08-14 11:59:23', '分类', b'1', 'id', 'asc', b'1', b'1', b'0', b'0', b'1', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": false}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (93, 'default', 'demo_gen_tree', '代码生成树形结构测试表', 'DemoGenTree', 'tree', 'internal/app/demo/n1/n2', 'n2', 'demo_gen_tree', '代码生成树形结构测试', 'gfast', '{\"treeCode\":\"id\",\"treeName\":\"demoName\",\"treeParentCode\":\"parentId\"}', '2022-11-29 15:11:34', '2025-08-14 11:59:32', '', b'1', 'id', 'asc', b'1', b'0', b'0', b'1', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": false}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (94, 'default', 'demo_data_auth', '数据权限测试', 'DemoDataAuth', 'crud', 'internal/app/demo', 'demo', 'demo_data_auth', '数据权限测试', 'gfast', '', '2023-02-12 11:18:42', '2024-03-18 10:18:00', '', b'1', 'id', 'asc', b'0', b'0', b'0', b'0', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": true}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (95, 'default', 'demo_snow_id', '雪花ID测试', 'DemoSnowId', 'crud', 'internal/app/demo', 'demo', 'demo_snow_id', '雪花ID测试', 'gfast', '', '2023-09-19 15:34:46', '2024-12-19 16:57:58', '', b'1', 'id', 'asc', b'1', b'1', b'1', b'0', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": true}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (96, 'default', 'demo_city_code', '省市区县和天气预报编码', 'DemoCityCode', 'tree', 'internal/app/demo', 'demo', 'demo_city_code', '省市区县', 'gfast', '{\"treeCode\":\"id\",\"treeName\":\"name\",\"treeParentCode\":\"pid\"}', '2023-11-02 10:38:51', '2024-04-07 09:30:57', '', b'1', 'id', 'asc', b'1', b'0', b'0', b'1', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": true}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (97, 'default', 'product_category', '商品分类表', 'ProductCategory', 'tree', 'internal/app/demo', 'demo', 'product_category', '商品分类表', 'gfast', '{\"treeCode\":\"categoryId\",\"treeName\":\"categoryName\",\"treeParentCode\":\"parentId\"}', '2024-04-07 09:08:30', '2024-04-07 10:08:26', '', b'1', 'category_id', 'asc', b'0', b'0', b'0', b'1', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": true}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);
INSERT INTO `tools_gen_table` VALUES (98, 'default', 'demo_gen_other', '特殊字段测试', 'DemoGenOther', 'crud', 'internal/app/system', 'system', 'demo_gen_other', '特殊字段测试', 'gfast', '', '2024-07-22 15:26:47', '2024-07-22 15:26:47', '', b'0', 'id', 'asc', b'0', b'0', b'0', b'0', b'0', 'null', 0);
INSERT INTO `tools_gen_table` VALUES (112, 'default', 'lm_config', '大模型相关配置数据表', 'LmConfig', 'crud', 'internal/app/llvm', 'llvm', 'lm_config', '大模型相关配置数据表', 'gfast', '', '2025-04-28 15:15:01', '2025-04-28 15:31:04', '', b'1', 'id', 'asc', b'1', b'0', b'0', b'0', b'0', '[{\"key\": \"api\", \"value\": true}, {\"key\": \"controller\", \"value\": true}, {\"key\": \"dao\", \"value\": true}, {\"key\": \"dao_internal\", \"value\": true}, {\"key\": \"logic\", \"value\": true}, {\"key\": \"model\", \"value\": true}, {\"key\": \"model_do\", \"value\": true}, {\"key\": \"model_entity\", \"value\": true}, {\"key\": \"router\", \"value\": true}, {\"key\": \"router_func\", \"value\": true}, {\"key\": \"service\", \"value\": true}, {\"key\": \"sql\", \"value\": true}, {\"key\": \"tsApi\", \"value\": true}, {\"key\": \"tsModel\", \"value\": true}, {\"key\": \"vue\", \"value\": true}, {\"key\": \"vueDetail\", \"value\": true}, {\"key\": \"vueEdit\", \"value\": true}]', 0);

-- ----------------------------
-- Table structure for tools_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table_column`;
CREATE TABLE `tools_gen_table_column`  (
                                           `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
                                           `table_id` bigint(20) NULL DEFAULT NULL COMMENT '归属表编号',
                                           `column_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '列名称',
                                           `column_comment` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '列描述',
                                           `column_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '列类型',
                                           `go_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
                                           `ts_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'TS类型',
                                           `go_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
                                           `html_field` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
                                           `is_pk` bit(1) NULL DEFAULT b'0' COMMENT '是否主键（1是）',
                                           `is_increment` bit(1) NULL DEFAULT b'0' COMMENT '是否自增（1是）',
                                           `is_required` bit(1) NULL DEFAULT b'0' COMMENT '是否必填（1是）',
                                           `is_edit` bit(1) NULL DEFAULT b'0' COMMENT '是否编辑字段（1是）',
                                           `is_list` bit(1) NULL DEFAULT b'1' COMMENT '是否列表字段（1是）',
                                           `is_detail` bit(1) NULL DEFAULT b'1' COMMENT '是否详情字段',
                                           `is_query` bit(1) NULL DEFAULT b'0' COMMENT '是否查询字段（1是）',
                                           `sort_order_edit` int(11) NULL DEFAULT 999 COMMENT '插入编辑显示顺序',
                                           `sort_order_list` int(11) NULL DEFAULT 999 COMMENT '列表显示顺序',
                                           `sort_order_detail` int(11) NULL DEFAULT 999 COMMENT '详情显示顺序',
                                           `sort_order_query` int(11) NULL DEFAULT 999 COMMENT '查询显示顺序',
                                           `query_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
                                           `html_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
                                           `dict_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
                                           `link_table_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表名',
                                           `link_table_class` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表类名',
                                           `link_table_module_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表模块名',
                                           `link_table_business_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表业务名',
                                           `link_table_package` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表包名',
                                           `link_label_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表键名',
                                           `link_label_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表字段值',
                                           `col_span` int(11) NULL DEFAULT 12 COMMENT '详情页占列数',
                                           `row_span` int(11) NULL DEFAULT 1 COMMENT '详情页占行数',
                                           `is_row_start` bit(1) NULL DEFAULT b'0' COMMENT '详情页为行首',
                                           `min_width` int(11) NULL DEFAULT 100 COMMENT '表格最小宽度',
                                           `is_fixed` bit(1) NULL DEFAULT b'0' COMMENT '是否表格列左固定',
                                           `is_overflow_tooltip` bit(1) NULL DEFAULT b'0' COMMENT '是否过长自动隐藏',
                                           `is_cascade` bit(1) NULL DEFAULT b'0' COMMENT '是否级联查询',
                                           `parent_column_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上级字段名',
                                           `cascade_column_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '级联查询字段',
                                           PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1133 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tools_gen_table_column
-- ----------------------------
INSERT INTO `tools_gen_table_column` VALUES (944, 91, 'id', 'ID', 'int(11) unsigned', 'uint', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'0', b'0', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (945, 91, 'demo_name', '姓名', 'varchar(20)', 'string', 'string', 'DemoName', 'demoName', b'0', b'0', b'1', b'1', b'1', b'0', b'1', 2, 2, 2, 2, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (946, 91, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'number', 'DemoAge', 'demoAge', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (947, 91, 'classes', '班级', 'varchar(30)', 'string', 'string', 'Classes', 'classes', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'EQ', 'radio', '', 'demo_gen_class', 'DemoGenClass', 'n2', 'demo_gen_class', 'internal/app/demo/n1/n2', 'id', 'class_name', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (948, 91, 'demo_born', '出生年月', 'datetime', 'Time', 'string', 'DemoBorn', 'demoBorn', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 5, 5, 5, 5, 'BETWEEN', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (949, 91, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'number', 'DemoGender', 'demoGender', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 6, 6, 6, 6, 'EQ', 'radio', 'sys_user_sex', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (950, 91, 'created_at', '创建日期', 'datetime', 'Time', 'string', 'CreatedAt', 'createdAt', b'0', b'0', b'0', b'0', b'1', b'1', b'1', 7, 7, 7, 7, 'BETWEEN', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (951, 91, 'updated_at', '修改日期', 'datetime', 'Time', 'string', 'UpdatedAt', 'updatedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 8, 8, 8, 8, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (952, 91, 'deleted_at', '删除日期', 'datetime', 'Time', 'string', 'DeletedAt', 'deletedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 9, 9, 9, 9, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (953, 91, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'number', 'CreatedBy', 'createdBy', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 10, 10, 10, 10, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (954, 91, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'number', 'UpdatedBy', 'updatedBy', b'0', b'0', b'0', b'0', b'0', b'1', b'0', 11, 11, 11, 11, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (955, 91, 'demo_status', '状态', 'tinyint(4)', 'int', 'number', 'DemoStatus', 'demoStatus', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 12, 12, 12, 12, 'EQ', 'radio', 'sys_normal_disable', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (956, 91, 'demo_cate', '分类', 'varchar(30)', 'string', 'string', 'DemoCate', 'demoCate', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 13, 13, 13, 13, 'EQ', 'checkbox', 'cms_article_type', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (957, 91, 'demo_thumb', '头像', 'text', 'string', 'string', 'DemoThumb', 'demoThumb', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 14, 14, 14, 14, 'EQ', 'imagefile', '', '', '', '', '', '', '', '', 5, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (958, 91, 'demo_photo', '相册', 'text', 'string', 'string', 'DemoPhoto', 'demoPhoto', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 15, 15, 15, 15, 'EQ', 'images', '', '', '', '', '', '', '', '', 5, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (959, 91, 'demo_info', '个人描述', 'text', 'string', 'string', 'DemoInfo', 'demoInfo', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 16, 16, 16, 16, 'EQ', 'richtext', '', '', '', '', '', '', '', '', 5, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (960, 91, 'demo_file', '相关附件', 'text', 'string', 'string', 'DemoFile', 'demoFile', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 17, 17, 17, 17, 'EQ', 'files', '', '', '', '', '', '', '', '', 5, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (961, 92, 'id', '分类id', 'int(10) unsigned', 'uint', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'1', b'0', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (962, 92, 'class_name', '分类名', 'varchar(30)', 'string', 'string', 'ClassName', 'className', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (963, 93, 'id', '', 'int(11) unsigned', 'uint', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'0', b'0', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (964, 93, 'parent_id', '父级ID', 'int(10) unsigned', 'uint', 'number', 'ParentId', 'parentId', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 2, 2, 2, 2, 'EQ', 'treeSelect', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (965, 93, 'demo_name', '姓名', 'varchar(20)', 'string', 'string', 'DemoName', 'demoName', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 3, 3, 3, 3, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (966, 93, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'number', 'DemoAge', 'demoAge', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (967, 93, 'classes', '班级', 'varchar(30)', 'string', 'string', 'Classes', 'classes', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 5, 5, 5, 5, 'EQ', 'select', '', 'demo_gen_class', 'DemoGenClass', 'n2', 'demo_gen_class', 'internal/app/demo/n1/n2', 'id', 'class_name', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (968, 93, 'demo_born', '出生年月', 'datetime', 'Time', 'string', 'DemoBorn', 'demoBorn', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 6, 6, 6, 6, 'EQ', 'date', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (969, 93, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'number', 'DemoGender', 'demoGender', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 7, 7, 7, 7, 'EQ', 'radio', 'sys_user_sex', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (970, 93, 'created_at', '创建日期', 'datetime', 'Time', 'string', 'CreatedAt', 'createdAt', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 8, 8, 8, 8, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 200, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (971, 93, 'updated_at', '修改日期', 'datetime', 'Time', 'string', 'UpdatedAt', 'updatedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 9, 9, 9, 9, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (972, 93, 'deleted_at', '删除日期', 'datetime', 'Time', 'string', 'DeletedAt', 'deletedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 10, 10, 10, 10, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (973, 93, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'number', 'CreatedBy', 'createdBy', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 11, 11, 11, 11, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (974, 93, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'number', 'UpdatedBy', 'updatedBy', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 12, 12, 12, 12, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (975, 93, 'demo_status', '状态', 'tinyint(4)', 'int', 'number', 'DemoStatus', 'demoStatus', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 13, 13, 13, 13, 'EQ', 'radio', 'sys_normal_disable', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (976, 93, 'demo_cate', '分类', 'varchar(30)', 'string', 'string', 'DemoCate', 'demoCate', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 14, 14, 14, 14, 'EQ', 'select', 'sys_oper_log_status', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (977, 91, 'classes_two', '班级2', 'varchar(30)', 'string', 'string', 'ClassesTwo', 'classesTwo', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'EQ', 'checkbox', '', 'demo_gen_class', 'DemoGenClass', 'n2', 'demo_gen_class', 'internal/app/demo/n1/n2', 'id', 'class_name', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (978, 94, 'id', '', 'int(11) unsigned', 'uint', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'1', b'1', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (979, 94, 'title', '标题', 'varchar(255)', 'string', 'string', 'Title', 'title', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (980, 94, 'created_by', '创建人', 'int(10) unsigned', 'uint', 'number', 'CreatedBy', 'createdBy', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (981, 94, 'updated_by', '修改人', 'int(10) unsigned', 'uint', 'number', 'UpdatedBy', 'updatedBy', b'0', b'0', b'0', b'0', b'0', b'1', b'0', 4, 4, 4, 4, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (982, 94, 'created_at', '创建时间', 'datetime', 'Time', 'string', 'CreatedAt', 'createdAt', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 5, 5, 5, 5, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (983, 94, 'updated_at', '修改时间', 'datetime', 'Time', 'string', 'UpdatedAt', 'updatedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 6, 6, 6, 6, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (984, 94, 'deleted_at', '删除时间', 'datetime', 'Time', 'string', 'DeletedAt', 'deletedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 7, 7, 7, 7, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (994, 91, 'cate_trees', '树型结构', 'varchar(10)', 'uint', 'number', 'CateTrees', 'cateTrees', b'0', b'0', b'0', b'1', b'0', b'1', b'1', 19, 19, 19, 19, 'EQ', 'treeSelect', '', 'demo_gen_tree', 'DemoGenTree', 'n2', 'demo_gen_tree', 'internal/app/demo/n1/n2', 'id', 'demo_name', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (995, 91, 'cate_trees_two', '树形多选', 'varchar(255)', 'string', 'string', 'CateTreesTwo', 'cateTreesTwo', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 20, 20, 20, 20, 'EQ', 'treeSelects', '', 'demo_gen_tree', 'DemoGenTree', 'n2', 'demo_gen_tree', 'internal/app/demo/n1/n2', 'id', 'demo_name', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (996, 91, 'options', '其他选项', 'text', 'string', 'string', 'Options', 'options', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 21, 21, 21, 21, 'EQ', 'keyValue', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (997, 95, 'id', 'ID', 'bigint(20) unsigned', 'uint64', 'number', 'Id', 'id', b'1', b'0', b'1', b'0', b'1', b'1', b'1', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (998, 95, 'name', '姓名', 'varchar(30)', 'string', 'string', 'Name', 'name', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (999, 95, 'age', '年龄', 'int(10) unsigned', 'uint', 'number', 'Age', 'age', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1000, 96, 'id', '城市ID', 'varchar(255)', 'string', 'string', 'Id', 'id', b'1', b'0', b'1', b'0', b'1', b'1', b'0', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1001, 96, 'pid', '城市父ID', 'varchar(255)', 'string', 'string', 'Pid', 'pid', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 2, 2, 2, 2, 'EQ', 'treeSelect', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1002, 96, 'deep', '级别', 'varchar(255)', 'string', 'string', 'Deep', 'deep', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1003, 96, 'name', '城市名称', 'varchar(255)', 'string', 'string', 'Name', 'name', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1004, 96, 'pinyin_prefix', '城市拼音头', 'varchar(255)', 'string', 'string', 'PinyinPrefix', 'pinyinPrefix', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 5, 5, 5, 5, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1005, 96, 'pinyin', '城市拼音', 'varchar(255)', 'string', 'string', 'Pinyin', 'pinyin', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 6, 6, 6, 6, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1006, 96, 'ext_id', '完整ID', 'varchar(255)', 'string', 'string', 'ExtId', 'extId', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 7, 7, 7, 7, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1007, 96, 'ext_name', '城市全称', 'varchar(255)', 'string', 'string', 'ExtName', 'extName', b'0', b'0', b'1', b'1', b'1', b'1', b'0', 8, 8, 8, 8, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1008, 96, 'weathercode', '天气预报的编码', 'varchar(255)', 'string', 'string', 'Weathercode', 'weathercode', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 9, 9, 9, 9, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 100, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1009, 97, 'category_id', '分类ID', 'bigint(10) unsigned', 'uint64', 'number', 'CategoryId', 'categoryId', b'1', b'1', b'1', b'0', b'1', b'1', b'0', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1010, 97, 'category_name', '分类名称', 'varchar(11)', 'string', 'string', 'CategoryName', 'categoryName', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'LIKE', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1011, 97, 'parent_id', '父分类ID', 'bigint(10) unsigned', 'uint64', 'number', 'ParentId', 'parentId', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 3, 3, 3, 3, 'EQ', 'treeSelect', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1012, 97, 'category_pic', '分类图标', 'varchar(255)', 'string', 'string', 'CategoryPic', 'categoryPic', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 4, 4, 4, 4, 'EQ', 'imagefile', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1013, 97, 'category_big_pic', '分类图片', 'varchar(255)', 'string', 'string', 'CategoryBigPic', 'categoryBigPic', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 5, 5, 5, 5, 'EQ', 'imagefile', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1014, 97, 'category_code', '分类编码', 'varchar(11)', 'string', 'string', 'CategoryCode', 'categoryCode', b'0', b'0', b'1', b'1', b'1', b'1', b'0', 6, 6, 6, 6, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1015, 97, 'category_level', '分类层级', 'tinyint(4)', 'int', 'number', 'CategoryLevel', 'categoryLevel', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 7, 7, 7, 7, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1016, 97, 'category_status', '分类状态', 'tinyint(4)', 'int', 'number', 'CategoryStatus', 'categoryStatus', b'0', b'0', b'1', b'1', b'0', b'1', b'0', 8, 8, 8, 8, 'EQ', 'radio', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1017, 97, 'is_show', '是否推荐', 'tinyint(4)', 'int', 'number', 'IsShow', 'isShow', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 9, 9, 9, 9, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1018, 97, 'sort', '排序方式(数字越小越靠前)', 'int(11) unsigned', 'uint', 'number', 'Sort', 'sort', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 10, 10, 10, 10, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1019, 97, 'app_id', '应用id', 'int(11) unsigned', 'uint', 'number', 'AppId', 'appId', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 11, 11, 11, 11, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1020, 97, 'created_at', '', 'datetime', 'Time', 'string', 'CreatedAt', 'createdAt', b'0', b'0', b'0', b'0', b'1', b'1', b'0', 12, 12, 12, 12, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1021, 97, 'updated_at', '', 'datetime', 'Time', 'string', 'UpdatedAt', 'updatedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 13, 13, 13, 13, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1022, 97, 'deleted_at', '', 'datetime', 'Time', 'string', 'DeletedAt', 'deletedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 14, 14, 14, 14, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1023, 98, 'id', 'ID', 'int(11) unsigned', 'uint', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'1', b'1', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1024, 98, 'info', '内容', 'text', 'string', 'string', 'Info', 'info', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1025, 98, 'img', '单图', 'text', 'string', 'string', 'Img', 'img', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1026, 98, 'imgs', '多图', 'text', 'string', 'string', 'Imgs', 'imgs', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1027, 98, 'file', '单文件', 'text', 'string', 'string', 'File', 'file', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 5, 5, 5, 5, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1028, 98, 'files', '多文件', 'text', 'string', 'string', 'Files', 'files', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 6, 6, 6, 6, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1029, 98, 'remark', '描述', 'varchar(255)', 'string', 'string', 'Remark', 'remark', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 7, 7, 7, 7, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1030, 95, 'thumb', '照片', 'varchar(255)', 'string', 'string', 'Thumb', 'thumb', b'0', b'0', b'0', b'1', b'1', b'1', b'0', 4, 4, 4, 4, 'EQ', 'imagefile', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1031, 95, 'photos', '相册', 'text', 'string', 'string', 'Photos', 'photos', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 5, 5, 5, 5, 'EQ', 'images', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1032, 95, 'files', '文件', 'text', 'string', 'string', 'Files', 'files', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 6, 6, 6, 6, 'EQ', 'files', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1033, 95, 'photos2', '相册二', 'text', 'string', 'string', 'Photos2', 'photos2', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 7, 7, 7, 7, 'EQ', 'imageSelector', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1034, 95, 'files2', '文件二', 'text', 'string', 'string', 'Files2', 'files2', b'0', b'0', b'0', b'1', b'0', b'1', b'0', 8, 8, 8, 8, 'EQ', 'fileSelector', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1035, 95, 'user_id', '关联人', 'bigint(20) unsigned', 'uint64', 'number', 'UserId', 'userId', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 9, 9, 9, 9, 'EQ', 'userSelectorSingle', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1036, 95, 'user_ids', '关联人信息', 'json', 'string', 'number[]', 'UserIds', 'userIds', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 10, 10, 10, 10, 'EQ', 'userSelectorMultiple', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1037, 95, 'depart_id', '关联部门', 'bigint(20) unsigned', 'uint64', 'number', 'DepartId', 'departId', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 11, 11, 11, 11, 'EQ', 'deptSelectorSingle', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1038, 95, 'depart_ids', '关联部门信息', 'json', 'string', 'number[]', 'DepartIds', 'departIds', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 12, 12, 12, 12, 'EQ', 'deptSelectorMultiple', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1039, 93, 'user_id', '关联人', 'bigint(20) unsigned', 'uint64', 'number', 'UserId', 'userId', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 15, 15, 15, 15, 'EQ', 'userSelectorSingle', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1040, 93, 'user_ids', '关联人信息', 'json', '', '', 'UserIds', 'userIds', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 16, 16, 16, 16, 'EQ', 'userSelectorMultiple', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1041, 93, 'depart_id', '关联部门', 'bigint(20) unsigned', 'uint64', 'number', 'DepartId', 'departId', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 17, 17, 17, 17, 'EQ', 'deptSelectorSingle', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1042, 93, 'depart_ids', '关联部门信息', 'json', '', '', 'DepartIds', 'departIds', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 18, 18, 18, 18, 'EQ', 'deptSelectorMultiple', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1124, 112, 'id', '主键ID', 'int(11)', 'int', 'number', 'Id', 'id', b'1', b'1', b'1', b'0', b'1', b'1', b'1', 1, 1, 1, 1, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1125, 112, 'model_name', '模型名称', 'varchar(255)', 'string', 'string', 'ModelName', 'modelName', b'0', b'0', b'1', b'1', b'1', b'1', b'1', 2, 2, 2, 2, 'EQ', 'select', 'lm_model_name', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1126, 112, 'api_key', 'API密钥', 'varchar(255)', 'string', 'string', 'ApiKey', 'apiKey', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 3, 3, 3, 3, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1127, 112, 'base_url', '基础URL', 'varchar(255)', 'string', 'string', 'BaseUrl', 'baseUrl', b'0', b'0', b'0', b'1', b'1', b'1', b'1', 4, 4, 4, 4, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1128, 112, 'created_at', '创建时间', 'timestamp', 'Time', 'string', 'CreatedAt', 'createdAt', b'0', b'0', b'0', b'0', b'1', b'1', b'1', 5, 5, 5, 5, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1129, 112, 'updated_at', '更新时间', 'timestamp', 'Time', 'string', 'UpdatedAt', 'updatedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 6, 6, 6, 6, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1130, 112, 'deleted_at', '删除时间', 'timestamp', 'Time', 'string', 'DeletedAt', 'deletedAt', b'0', b'0', b'0', b'0', b'0', b'0', b'0', 7, 7, 7, 7, 'EQ', 'datetime', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1131, 112, 'created_by', '创建者ID', 'int(11)', 'int', 'number', 'CreatedBy', 'createdBy', b'0', b'0', b'0', b'0', b'1', b'1', b'1', 8, 8, 8, 8, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');
INSERT INTO `tools_gen_table_column` VALUES (1132, 112, 'updated_by', '更新者ID', 'int(11)', 'int', 'number', 'UpdatedBy', 'updatedBy', b'0', b'0', b'0', b'0', b'0', b'1', b'0', 9, 9, 9, 9, 'EQ', 'input', '', '', '', '', '', '', '', '', 1, 1, b'0', 150, b'0', b'0', b'0', '', '');

SET FOREIGN_KEY_CHECKS = 1;