-- 设置默认字符集为utf8mb4以支持中文
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- 创建数据库并指定字符集
CREATE DATABASE IF NOT EXISTS `goods` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- CREATE DATABASE IF NOT EXISTS `admin` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `user` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `interaction` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `order` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `resource` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `banner` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
       
USE goods;
    
-- 商品主表
DROP TABLE IF EXISTS `goods_info`;
CREATE TABLE `goods_info` (
          `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
          `name` VARCHAR(200) NOT NULL COMMENT '商品名字',
          `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主图',
          `images` JSON DEFAULT NULL COMMENT '支持单图,多图',
          `price` INT NOT NULL COMMENT '价格(分)',
          `level1_category_id` INT NOT NULL COMMENT '1级分类id',
          `level2_category_id` INT NOT NULL DEFAULT 0 COMMENT '2级分类id',
          `level3_category_id` INT NOT NULL DEFAULT 0 COMMENT '3级分类id',
          `brand` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌',
          `stock` INT NOT NULL DEFAULT 0 COMMENT '库存',
          `sale` INT NOT NULL DEFAULT 0 COMMENT '销量',
          `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
          `detail_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '商品详情',
          `created_at` datetime DEFAULT NULL,
          `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
          `deleted_at` datetime DEFAULT NULL,
          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

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

USE admin;
    
-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS `admin_info`;
CREATE TABLE `admin_info`  (
           `id` int NOT NULL AUTO_INCREMENT,
           `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
           `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
           `role_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色ids',
           `created_at` datetime DEFAULT NULL,
           `updated_at` datetime DEFAULT NULL,
           `user_salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐',
           `is_admin` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否超级管理员',
           PRIMARY KEY (`id`) USING BTREE,
           UNIQUE INDEX `name_unique`(`name`) USING BTREE COMMENT '名字唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_info
-- ----------------------------
INSERT INTO `admin_info` VALUES (1, 'zhangsan', 'e91474a50e96e9e3b0c7df489b1c0a21', '2', '2022-09-25 16:40:43', '2022-11-20 11:06:01', 'e3oHjweGEc', 0);
INSERT INTO `admin_info` VALUES (3, 'wangzhongyang', '7382e435a4eb141adeabc3792d383e1c', '2', '2022-07-19 10:50:20', '2022-11-23 14:25:10', '4f8WG1bjne', 0);
INSERT INTO `admin_info` VALUES (13, '李四', '9076805c0efa82a164f0c4f2a2818851', '1', '2022-11-20 11:03:35', '2022-11-20 11:03:35', 'Io45dMSb4e', 1);
INSERT INTO `admin_info` VALUES (15, 'zhaoliu', 'd82abc6395e1c89e7837f96407cf6d5d', '2', '2022-11-20 13:45:09', '2022-11-20 13:45:49', 'aHzOD3zI7L', 0);

-- ----------------------------
-- Table structure for role_info
-- ----------------------------
DROP TABLE IF EXISTS `role_info`;
CREATE TABLE `role_info`  (
      `id` int NOT NULL AUTO_INCREMENT,
      `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
      `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
      `created_at` datetime DEFAULT NULL,
      `updated_at` datetime DEFAULT NULL,
      `deleted_at` datetime DEFAULT NULL,
      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE INDEX `unique_index`(`name`) USING BTREE COMMENT '角色昵称唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_info
-- ----------------------------
INSERT INTO `role_info` VALUES (1, '运营1', '测试', '2022-09-25 10:35:52', '2022-12-24 10:51:24', NULL);
INSERT INTO `role_info` VALUES (3, '运营', '', '2022-12-21 10:43:33', '2022-12-21 10:43:33', NULL);

-- ----------------------------
-- Table structure for role_permission_info
-- ----------------------------
DROP TABLE IF EXISTS `role_permission_info`;
CREATE TABLE `role_permission_info`  (
         `id` int NOT NULL AUTO_INCREMENT,
         `role_id` int NOT NULL DEFAULT 0 COMMENT '角色id',
         `permission_id` int NOT NULL COMMENT '权限id',
         `created_at` datetime DEFAULT NULL,
         `updated_at` datetime DEFAULT NULL,
         PRIMARY KEY (`id`) USING BTREE,
         UNIQUE INDEX `unique_index`(`role_id`, `permission_id`) USING BTREE COMMENT '唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for permission_info
-- ----------------------------
DROP TABLE IF EXISTS `permission_info`;
CREATE TABLE `permission_info`  (
        `id` int NOT NULL AUTO_INCREMENT,
        `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限名称',
        `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路径',
        `created_at` datetime DEFAULT NULL,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        PRIMARY KEY (`id`) USING BTREE,
        UNIQUE INDEX `unique_name`(`name`) USING BTREE COMMENT '名称唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission_info
-- ----------------------------
INSERT INTO `permission_info` VALUES (1, '文章1', 'admin.article.index', '2022-09-25 15:03:01', '2022-09-25 15:03:43', NULL);
INSERT INTO `permission_info` VALUES (2, '测试2', 'admin.test.index', NULL, NULL, NULL);
INSERT INTO `permission_info` VALUES (5, '商品3', 'admin/goods', '2022-12-26 19:51:44', '2022-12-26 19:52:29', NULL);
INSERT INTO `permission_info` VALUES (6, '商品2', 'admin/goods', '2022-12-26 19:52:01', '2022-12-26 19:52:01', NULL);

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
          `sex` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1男 2女',
          `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1正常 2拉黑冻结',
          `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '个性签名',
          `secret_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密保问题的答案',
          `created_at` datetime DEFAULT NULL,
          `updated_at` datetime DEFAULT NULL,
          `deleted_at` datetime DEFAULT NULL,
          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品表' ROW_FORMAT = Dynamic;

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
           `pay_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '支付时间',
           `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价',
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
