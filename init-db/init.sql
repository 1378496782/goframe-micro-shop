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

/*
下面是管理后台需要的数据表
*/

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
-- Table structure for demo_city_code
-- ----------------------------
DROP TABLE IF EXISTS `demo_city_code`;
CREATE TABLE `demo_city_code`  (
                                   `id` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '城市ID',
                                   `pid` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '城市父ID',
                                   `deep` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '级别',
                                   `name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '城市名称',
                                   `pinyin_prefix` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '城市拼音头',
                                   `pinyin` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '城市拼音',
                                   `ext_id` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '完整ID',
                                   `ext_name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '城市全称',
                                   `weathercode` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '天气预报的编码',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   UNIQUE INDEX `id，name,code`(`id`, `name`, `weathercode`) USING BTREE COMMENT '这三个字段并列一起，必须是唯一的'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '省市区县和天气预报编码' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_city_code
-- ----------------------------
INSERT INTO `demo_city_code` VALUES ('11', '0', '0', '北京', 'b', 'bei jing', '110000000000', '北京市', '101010100');
INSERT INTO `demo_city_code` VALUES ('1101', '11', '1', '北京', 'b', 'bei jing', '110100000000', '北京市', '101010100');
INSERT INTO `demo_city_code` VALUES ('110101', '1101', '2', '东城', 'd', 'dong cheng', '110101000000', '东城区', '101010100');
INSERT INTO `demo_city_code` VALUES ('110102', '1101', '2', '西城', 'x', 'xi cheng', '110102000000', '西城区', '101010100');
INSERT INTO `demo_city_code` VALUES ('110105', '1101', '2', '朝阳', 'c', 'chao yang', '110105000000', '朝阳区', '101010300');
INSERT INTO `demo_city_code` VALUES ('110106', '1101', '2', '丰台', 'f', 'feng tai', '110106000000', '丰台区', '101010900');
INSERT INTO `demo_city_code` VALUES ('110107', '1101', '2', '石景山', 's', 'shi jing shan', '110107000000', '石景山区', '101011000');
INSERT INTO `demo_city_code` VALUES ('110108', '1101', '2', '海淀', 'h', 'hai dian', '110108000000', '海淀区', '101010200');
INSERT INTO `demo_city_code` VALUES ('110109', '1101', '2', '门头沟', 'm', 'men tou gou', '110109000000', '门头沟区', '101011400');
INSERT INTO `demo_city_code` VALUES ('110111', '1101', '2', '房山', 'f', 'fang shan', '110111000000', '房山区', '101011200');
INSERT INTO `demo_city_code` VALUES ('110112', '1101', '2', '通州', 't', 'tong zhou', '110112000000', '通州区', '101010600');
INSERT INTO `demo_city_code` VALUES ('110113', '1101', '2', '顺义', 's', 'shun yi', '110113000000', '顺义区', '101010400');
INSERT INTO `demo_city_code` VALUES ('110114', '1101', '2', '昌平', 'c', 'chang ping', '110114000000', '昌平区', '101010700');
INSERT INTO `demo_city_code` VALUES ('110115', '1101', '2', '大兴', 'd', 'da xing', '110115000000', '大兴区', '101011100');
INSERT INTO `demo_city_code` VALUES ('110116', '1101', '2', '怀柔', 'h', 'huai rou', '110116000000', '怀柔区', '101010500');
INSERT INTO `demo_city_code` VALUES ('110117', '1101', '2', '平谷', 'p', 'ping gu', '110117000000', '平谷区', '101011500');
INSERT INTO `demo_city_code` VALUES ('110118', '1101', '2', '密云', 'm', 'mi yun', '110118000000', '密云区', '101011300');
INSERT INTO `demo_city_code` VALUES ('110119', '1101', '2', '延庆', 'y', 'yan qing', '110119000000', '延庆区', '101010800');
INSERT INTO `demo_city_code` VALUES ('12', '0', '0', '天津', 't', 'tian jin', '120000000000', '天津市', '101030100');
INSERT INTO `demo_city_code` VALUES ('1201', '12', '1', '天津', 't', 'tian jin', '120100000000', '天津市', '101030100');
INSERT INTO `demo_city_code` VALUES ('120101', '1201', '2', '和平', 'h', 'he ping', '120101000000', '和平区', '101281204');
INSERT INTO `demo_city_code` VALUES ('120102', '1201', '2', '河东', 'h', 'he dong', '120102000000', '河东区', '101030100');
INSERT INTO `demo_city_code` VALUES ('120103', '1201', '2', '河西', 'h', 'he xi', '120103000000', '河西区', '101030100');
INSERT INTO `demo_city_code` VALUES ('120104', '1201', '2', '南开', 'n', 'nan kai', '120104000000', '南开区', '101030100');
INSERT INTO `demo_city_code` VALUES ('120105', '1201', '2', '河北', 'h', 'he bei', '120105000000', '河北区', '101030100');
INSERT INTO `demo_city_code` VALUES ('120106', '1201', '2', '红桥', 'h', 'hong qiao', '120106000000', '红桥区', '101030100');
INSERT INTO `demo_city_code` VALUES ('120110', '1201', '2', '东丽', 'd', 'dong li', '120110000000', '东丽区', '101030400');
INSERT INTO `demo_city_code` VALUES ('120111', '1201', '2', '西青', 'x', 'xi qing', '120111000000', '西青区', '101030500');
INSERT INTO `demo_city_code` VALUES ('120112', '1201', '2', '津南', 'j', 'jin nan', '120112000000', '津南区', '101031000');
INSERT INTO `demo_city_code` VALUES ('120113', '1201', '2', '北辰', 'b', 'bei chen', '120113000000', '北辰区', '101030600');
INSERT INTO `demo_city_code` VALUES ('120114', '1201', '2', '武清', 'w', 'wu qing', '120114000000', '武清区', '101030200');
INSERT INTO `demo_city_code` VALUES ('120115', '1201', '2', '宝坻', 'b', 'bao di', '120115000000', '宝坻区', '101030300');
INSERT INTO `demo_city_code` VALUES ('120116', '1201', '2', '滨海新区', 'b', 'bin hai xin qu', '120116000000', '滨海新区', '101031100');
INSERT INTO `demo_city_code` VALUES ('120117', '1201', '2', '宁河', 'n', 'ning he', '120117000000', '宁河区', '101030700');
INSERT INTO `demo_city_code` VALUES ('120118', '1201', '2', '静海', 'j', 'jing hai', '120118000000', '静海区', '101030900');
INSERT INTO `demo_city_code` VALUES ('120119', '1201', '2', '蓟州', 'j', 'ji zhou', '120119000000', '蓟州区', '101031400');
INSERT INTO `demo_city_code` VALUES ('13', '0', '0', '河北', 'h', 'he bei', '130000000000', '河北省', NULL);
INSERT INTO `demo_city_code` VALUES ('1301', '13', '1', '石家庄', 's', 'shi jia zhuang', '130100000000', '石家庄市', '101090101');
INSERT INTO `demo_city_code` VALUES ('130102', '1301', '2', '长安', 'c', 'chang an', '130102000000', '长安区', '101110102');
INSERT INTO `demo_city_code` VALUES ('130104', '1301', '2', '桥西', 'q', 'qiao xi', '130104000000', '桥西区', NULL);
INSERT INTO `demo_city_code` VALUES ('130105', '1301', '2', '新华', 'x', 'xin hua', '130105000000', '新华区', NULL);
INSERT INTO `demo_city_code` VALUES ('130107', '1301', '2', '井陉矿区', 'j', 'jing xing kuang qu', '130107000000', '井陉矿区', NULL);
INSERT INTO `demo_city_code` VALUES ('130108', '1301', '2', '裕华', 'y', 'yu hua', '130108000000', '裕华区', NULL);
INSERT INTO `demo_city_code` VALUES ('130109', '1301', '2', '藁城', 'g', 'gao cheng', '130109000000', '藁城区', '101090115');
INSERT INTO `demo_city_code` VALUES ('130110', '1301', '2', '鹿泉', 'l', 'lu quan', '130110000000', '鹿泉区', '101090118');
INSERT INTO `demo_city_code` VALUES ('130111', '1301', '2', '栾城', 'l', 'luan cheng', '130111000000', '栾城区', '101090104');
INSERT INTO `demo_city_code` VALUES ('130121', '1301', '2', '井陉', 'j', 'jing xing', '130121000000', '井陉县', '101090102');
INSERT INTO `demo_city_code` VALUES ('130123', '1301', '2', '正定', 'z', 'zheng ding', '130123000000', '正定县', '101090103');
INSERT INTO `demo_city_code` VALUES ('130125', '1301', '2', '行唐', 'x', 'xing tang', '130125000000', '行唐县', '101090105');
INSERT INTO `demo_city_code` VALUES ('130126', '1301', '2', '灵寿', 'l', 'ling shou', '130126000000', '灵寿县', '101090106');
INSERT INTO `demo_city_code` VALUES ('130127', '1301', '2', '高邑', 'g', 'gao yi', '130127000000', '高邑县', '101090107');
INSERT INTO `demo_city_code` VALUES ('130128', '1301', '2', '深泽', 's', 'shen ze', '130128000000', '深泽县', '101090108');
INSERT INTO `demo_city_code` VALUES ('130129', '1301', '2', '赞皇', 'z', 'zan huang', '130129000000', '赞皇县', '101090109');
INSERT INTO `demo_city_code` VALUES ('130130', '1301', '2', '无极', 'w', 'wu ji', '130130000000', '无极县', '101090110');
INSERT INTO `demo_city_code` VALUES ('130131', '1301', '2', '平山', 'p', 'ping shan', '130131000000', '平山县', '101090111');
INSERT INTO `demo_city_code` VALUES ('130132', '1301', '2', '元氏', 'y', 'yuan shi', '130132000000', '元氏县', '101090112');
INSERT INTO `demo_city_code` VALUES ('130133', '1301', '2', '赵县', 'z', 'zhao xian', '130133000000', '赵县', '101090113');
INSERT INTO `demo_city_code` VALUES ('130181', '1301', '2', '辛集', 'x', 'xin ji', '130181000000', '辛集市', '101090114');
INSERT INTO `demo_city_code` VALUES ('130183', '1301', '2', '晋州', 'j', 'jin zhou', '130183000000', '晋州市', '101090116');
INSERT INTO `demo_city_code` VALUES ('130184', '1301', '2', '新乐', 'x', 'xin le', '130184000000', '新乐市', '101090117');
INSERT INTO `demo_city_code` VALUES ('1302', '13', '1', '唐山', 't', 'tang shan', '130200000000', '唐山市', '101090501');
INSERT INTO `demo_city_code` VALUES ('130202', '1302', '2', '路南', 'l', 'lu nan', '130202000000', '路南区', NULL);
INSERT INTO `demo_city_code` VALUES ('130203', '1302', '2', '路北', 'l', 'lu bei', '130203000000', '路北区', NULL);
INSERT INTO `demo_city_code` VALUES ('130204', '1302', '2', '古冶', 'g', 'gu ye', '130204000000', '古冶区', NULL);
INSERT INTO `demo_city_code` VALUES ('130205', '1302', '2', '开平', 'k', 'kai ping', '130205000000', '开平区', '101281103');
INSERT INTO `demo_city_code` VALUES ('130207', '1302', '2', '丰南', 'f', 'feng nan', '130207000000', '丰南区', '101090502');
INSERT INTO `demo_city_code` VALUES ('130208', '1302', '2', '丰润', 'f', 'feng run', '130208000000', '丰润区', '101090503');
INSERT INTO `demo_city_code` VALUES ('130209', '1302', '2', '曹妃甸', 'c', 'cao fei dian', '130209000000', '曹妃甸区', '101090512');
INSERT INTO `demo_city_code` VALUES ('130224', '1302', '2', '滦南', 'l', 'luan nan', '130224000000', '滦南县', '101090505');
INSERT INTO `demo_city_code` VALUES ('130225', '1302', '2', '乐亭', 'l', 'lao ting', '130225000000', '乐亭县', '101090506');
INSERT INTO `demo_city_code` VALUES ('130227', '1302', '2', '迁西', 'q', 'qian xi', '130227000000', '迁西县', '101090507');
INSERT INTO `demo_city_code` VALUES ('130229', '1302', '2', '玉田', 'y', 'yu tian', '130229000000', '玉田县', '101090508');
INSERT INTO `demo_city_code` VALUES ('130281', '1302', '2', '遵化', 'z', 'zun hua', '130281000000', '遵化市', '101090510');
INSERT INTO `demo_city_code` VALUES ('130283', '1302', '2', '迁安', 'q', 'qian an', '130283000000', '迁安市', '101090511');
INSERT INTO `demo_city_code` VALUES ('130284', '1302', '2', '滦州', 'l', 'luan zhou', '130284000000', '滦州市', NULL);
INSERT INTO `demo_city_code` VALUES ('1303', '13', '1', '秦皇岛', 'q', 'qin huang dao', '130300000000', '秦皇岛市', '101091101');
INSERT INTO `demo_city_code` VALUES ('130302', '1303', '2', '海港', 'h', 'hai gang', '130302000000', '海港区', NULL);
INSERT INTO `demo_city_code` VALUES ('130303', '1303', '2', '山海关', 's', 'shan hai guan', '130303000000', '山海关区', NULL);
INSERT INTO `demo_city_code` VALUES ('130304', '1303', '2', '北戴河', 'b', 'bei dai he', '130304000000', '北戴河区', '101091106');
INSERT INTO `demo_city_code` VALUES ('130306', '1303', '2', '抚宁', 'f', 'fu ning', '130306000000', '抚宁区', '101091104');
INSERT INTO `demo_city_code` VALUES ('130321', '1303', '2', '青龙', 'q', 'qing long', '130321000000', '青龙满族自治县', '101091102');
INSERT INTO `demo_city_code` VALUES ('130322', '1303', '2', '昌黎', 'c', 'chang li', '130322000000', '昌黎县', '101091103');
INSERT INTO `demo_city_code` VALUES ('130324', '1303', '2', '卢龙', 'l', 'lu long', '130324000000', '卢龙县', '101091105');
INSERT INTO `demo_city_code` VALUES ('1304', '13', '1', '邯郸', 'h', 'han dan', '130400000000', '邯郸市', '101091001');
INSERT INTO `demo_city_code` VALUES ('130402', '1304', '2', '邯山', 'h', 'han shan', '130402000000', '邯山区', '101091017');
INSERT INTO `demo_city_code` VALUES ('130403', '1304', '2', '丛台', 'c', 'cong tai', '130403000000', '丛台区', '101091018');
INSERT INTO `demo_city_code` VALUES ('130404', '1304', '2', '复兴', 'f', 'fu xing', '130404000000', '复兴区', '101091019');
INSERT INTO `demo_city_code` VALUES ('130406', '1304', '2', '峰峰矿区', 'f', 'feng feng kuang qu', '130406000000', '峰峰矿区', '101091001017');
INSERT INTO `demo_city_code` VALUES ('130407', '1304', '2', '肥乡', 'f', 'fei xiang', '130407000000', '肥乡区', '101091008');
INSERT INTO `demo_city_code` VALUES ('130408', '1304', '2', '永年', 'y', 'yong nian', '130408000000', '永年区', '101091009');
INSERT INTO `demo_city_code` VALUES ('130423', '1304', '2', '临漳', 'l', 'lin zhang', '130423000000', '临漳县', '101091003');
INSERT INTO `demo_city_code` VALUES ('130424', '1304', '2', '成安', 'c', 'cheng an', '130424000000', '成安县', '101091004');
INSERT INTO `demo_city_code` VALUES ('130425', '1304', '2', '大名', 'd', 'da ming', '130425000000', '大名县', '101091005');
INSERT INTO `demo_city_code` VALUES ('130426', '1304', '2', '涉县', 's', 'she xian', '130426000000', '涉县', '101091006');
INSERT INTO `demo_city_code` VALUES ('130427', '1304', '2', '磁县', 'c', 'ci xian', '130427000000', '磁县', '101091007');
INSERT INTO `demo_city_code` VALUES ('130430', '1304', '2', '邱县', 'q', 'qiu xian', '130430000000', '邱县', '101091010');
INSERT INTO `demo_city_code` VALUES ('130431', '1304', '2', '鸡泽', 'j', 'ji ze', '130431000000', '鸡泽县', '101091011');
INSERT INTO `demo_city_code` VALUES ('130432', '1304', '2', '广平', 'g', 'guang ping', '130432000000', '广平县', '101091012');
INSERT INTO `demo_city_code` VALUES ('130433', '1304', '2', '馆陶', 'g', 'guan tao', '130433000000', '馆陶县', '101091013');
INSERT INTO `demo_city_code` VALUES ('130434', '1304', '2', '魏县', 'w', 'wei xian', '130434000000', '魏县', '101091014');
INSERT INTO `demo_city_code` VALUES ('130435', '1304', '2', '曲周', 'q', 'qu zhou', '130435000000', '曲周县', '101091015');
INSERT INTO `demo_city_code` VALUES ('130481', '1304', '2', '武安', 'w', 'wu an', '130481000000', '武安市', '101091016');
INSERT INTO `demo_city_code` VALUES ('1305', '13', '1', '邢台', 'x', 'xing tai', '130500000000', '邢台市', '101090901');
INSERT INTO `demo_city_code` VALUES ('130502', '1305', '2', '襄都', 'x', 'xiang du', '130502000000', '襄都区', '101090504');
INSERT INTO `demo_city_code` VALUES ('130503', '1305', '2', '信都', 'x', 'xin du', '130503000000', '信都区', '101090919');
INSERT INTO `demo_city_code` VALUES ('130505', '1305', '2', '任泽', 'r', 'ren ze', '130505000000', '任泽区', '101090918');
INSERT INTO `demo_city_code` VALUES ('130506', '1305', '2', '南和', 'n', 'nan he', '130506000000', '南和区', '101090907');
INSERT INTO `demo_city_code` VALUES ('130522', '1305', '2', '临城', 'l', 'lin cheng', '130522000000', '临城县', '101090902');
INSERT INTO `demo_city_code` VALUES ('130523', '1305', '2', '内丘', 'n', 'nei qiu', '130523000000', '内丘县', '101090904');
INSERT INTO `demo_city_code` VALUES ('130524', '1305', '2', '柏乡', 'b', 'bai xiang', '130524000000', '柏乡县', '101090905');
INSERT INTO `demo_city_code` VALUES ('130525', '1305', '2', '隆尧', 'l', 'long yao', '130525000000', '隆尧县', '101090906');
INSERT INTO `demo_city_code` VALUES ('130528', '1305', '2', '宁晋', 'n', 'ning jin', '130528000000', '宁晋县', '101090908');
INSERT INTO `demo_city_code` VALUES ('130529', '1305', '2', '巨鹿', 'j', 'ju lu', '130529000000', '巨鹿县', '101090909');
INSERT INTO `demo_city_code` VALUES ('130530', '1305', '2', '新河', 'x', 'xin he', '130530000000', '新河县', '101090910');
INSERT INTO `demo_city_code` VALUES ('130531', '1305', '2', '广宗', 'g', 'guang zong', '130531000000', '广宗县', '101090911');
INSERT INTO `demo_city_code` VALUES ('130532', '1305', '2', '平乡', 'p', 'ping xiang', '130532000000', '平乡县', '101090912');
INSERT INTO `demo_city_code` VALUES ('130533', '1305', '2', '威县', 'w', 'wei xian', '130533000000', '威县', '101090913');
INSERT INTO `demo_city_code` VALUES ('130534', '1305', '2', '清河', 'q', 'qing he', '130534000000', '清河县', '101090914');
INSERT INTO `demo_city_code` VALUES ('130535', '1305', '2', '临西', 'l', 'lin xi', '130535000000', '临西县', '101090915');
INSERT INTO `demo_city_code` VALUES ('130581', '1305', '2', '南宫', 'n', 'nan gong', '130581000000', '南宫市', '101090916');
INSERT INTO `demo_city_code` VALUES ('130582', '1305', '2', '沙河', 's', 'sha he', '130582000000', '沙河市', '101090917');
INSERT INTO `demo_city_code` VALUES ('1306', '13', '1', '保定', 'b', 'bao ding', '130600000000', '保定市', '101090201');
INSERT INTO `demo_city_code` VALUES ('130602', '1306', '2', '竞秀', 'j', 'jing xiu', '130602000000', '竞秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('130606', '1306', '2', '莲池', 'l', 'lian chi', '130606000000', '莲池区', NULL);
INSERT INTO `demo_city_code` VALUES ('130607', '1306', '2', '满城', 'm', 'man cheng', '130607000000', '满城区', '101090202');
INSERT INTO `demo_city_code` VALUES ('130608', '1306', '2', '清苑', 'q', 'qing yuan', '130608000000', '清苑区', '101090224');
INSERT INTO `demo_city_code` VALUES ('130609', '1306', '2', '徐水', 'x', 'xu shui', '130609000000', '徐水区', '101090204');
INSERT INTO `demo_city_code` VALUES ('130623', '1306', '2', '涞水', 'l', 'lai shui', '130623000000', '涞水县', '101090222');
INSERT INTO `demo_city_code` VALUES ('130624', '1306', '2', '阜平', 'f', 'fu ping', '130624000000', '阜平县', '101090203');
INSERT INTO `demo_city_code` VALUES ('130626', '1306', '2', '定兴', 'd', 'ding xing', '130626000000', '定兴县', '101090223');
INSERT INTO `demo_city_code` VALUES ('130627', '1306', '2', '唐县', 't', 'tang xian', '130627000000', '唐县', '101090205');
INSERT INTO `demo_city_code` VALUES ('130628', '1306', '2', '高阳', 'g', 'gao yang', '130628000000', '高阳县', '101090206');
INSERT INTO `demo_city_code` VALUES ('130629', '1306', '2', '容城', 'r', 'rong cheng', '130629000000', '容城县', '101090207');
INSERT INTO `demo_city_code` VALUES ('130630', '1306', '2', '涞源', 'l', 'lai yuan', '130630000000', '涞源县', '101090209');
INSERT INTO `demo_city_code` VALUES ('130631', '1306', '2', '望都', 'w', 'wang du', '130631000000', '望都县', '101090210');
INSERT INTO `demo_city_code` VALUES ('130632', '1306', '2', '安新', 'a', 'an xin', '130632000000', '安新县', '101090211');
INSERT INTO `demo_city_code` VALUES ('130633', '1306', '2', '易县', 'y', 'yi xian', '130633000000', '易县', '101090212');
INSERT INTO `demo_city_code` VALUES ('130634', '1306', '2', '曲阳', 'q', 'qu yang', '130634000000', '曲阳县', '101090214');
INSERT INTO `demo_city_code` VALUES ('130635', '1306', '2', '蠡县', 'l', 'li xian', '130635000000', '蠡县', '101090215');
INSERT INTO `demo_city_code` VALUES ('130636', '1306', '2', '顺平', 's', 'shun ping', '130636000000', '顺平县', '101090216');
INSERT INTO `demo_city_code` VALUES ('130637', '1306', '2', '博野', 'b', 'bo ye', '130637000000', '博野县', '101090225');
INSERT INTO `demo_city_code` VALUES ('130638', '1306', '2', '雄县', 'x', 'xiong xian', '130638000000', '雄县', '101090217');
INSERT INTO `demo_city_code` VALUES ('130681', '1306', '2', '涿州', 'z', 'zhuo zhou', '130681000000', '涿州市', '101090218');
INSERT INTO `demo_city_code` VALUES ('130682', '1306', '2', '定州', 'd', 'ding zhou', '130682000000', '定州市', '101090219');
INSERT INTO `demo_city_code` VALUES ('130683', '1306', '2', '安国', 'a', 'an guo', '130683000000', '安国市', '101090220');
INSERT INTO `demo_city_code` VALUES ('130684', '1306', '2', '高碑店', 'g', 'gao bei dian', '130684000000', '高碑店市', '101090221');
INSERT INTO `demo_city_code` VALUES ('1307', '13', '1', '张家口', 'z', 'zhang jia kou', '130700000000', '张家口市', '101090301');
INSERT INTO `demo_city_code` VALUES ('130702', '1307', '2', '桥东', 'q', 'qiao dong', '130702000000', '桥东区', NULL);
INSERT INTO `demo_city_code` VALUES ('130703', '1307', '2', '桥西', 'q', 'qiao xi', '130703000000', '桥西区', NULL);
INSERT INTO `demo_city_code` VALUES ('130705', '1307', '2', '宣化', 'x', 'xuan hua', '130705000000', '宣化区', '101090302');
INSERT INTO `demo_city_code` VALUES ('130706', '1307', '2', '下花园', 'x', 'xia hua yuan', '130706000000', '下花园区', NULL);
INSERT INTO `demo_city_code` VALUES ('130708', '1307', '2', '万全', 'w', 'wan quan', '130708000000', '万全区', '101090310');
INSERT INTO `demo_city_code` VALUES ('130709', '1307', '2', '崇礼', 'c', 'chong li', '130709000000', '崇礼区', '101090314');
INSERT INTO `demo_city_code` VALUES ('130722', '1307', '2', '张北', 'z', 'zhang bei', '130722000000', '张北县', '101090303');
INSERT INTO `demo_city_code` VALUES ('130723', '1307', '2', '康保', 'k', 'kang bao', '130723000000', '康保县', '101090304');
INSERT INTO `demo_city_code` VALUES ('130724', '1307', '2', '沽源', 'g', 'gu yuan', '130724000000', '沽源县', '101090305');
INSERT INTO `demo_city_code` VALUES ('130725', '1307', '2', '尚义', 's', 'shang yi', '130725000000', '尚义县', '101090306');
INSERT INTO `demo_city_code` VALUES ('130726', '1307', '2', '蔚县', 'y', 'yu xian', '130726000000', '蔚县', '101090307');
INSERT INTO `demo_city_code` VALUES ('130727', '1307', '2', '阳原', 'y', 'yang yuan', '130727000000', '阳原县', '101090308');
INSERT INTO `demo_city_code` VALUES ('130728', '1307', '2', '怀安', 'h', 'huai an', '130728000000', '怀安县', '101090309');
INSERT INTO `demo_city_code` VALUES ('130730', '1307', '2', '怀来', 'h', 'huai lai', '130730000000', '怀来县', '101090311');
INSERT INTO `demo_city_code` VALUES ('130731', '1307', '2', '涿鹿', 'z', 'zhuo lu', '130731000000', '涿鹿县', '101090312');
INSERT INTO `demo_city_code` VALUES ('130732', '1307', '2', '赤城', 'c', 'chi cheng', '130732000000', '赤城县', '101090313');
INSERT INTO `demo_city_code` VALUES ('1308', '13', '1', '承德', 'c', 'cheng de', '130800000000', '承德市', '101090402');
INSERT INTO `demo_city_code` VALUES ('130802', '1308', '2', '双桥', 's', 'shuang qiao', '130802000000', '双桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('130803', '1308', '2', '双滦', 's', 'shuang luan', '130803000000', '双滦区', NULL);
INSERT INTO `demo_city_code` VALUES ('130804', '1308', '2', '鹰手营子矿区', 'y', 'ying shou ying zi kuang qu', '130804000000', '鹰手营子矿区', NULL);
INSERT INTO `demo_city_code` VALUES ('130821', '1308', '2', '承德县', 'c', 'cheng de xian', '130821000000', '承德县', '101090403');
INSERT INTO `demo_city_code` VALUES ('130822', '1308', '2', '兴隆', 'x', 'xing long', '130822000000', '兴隆县', '101090404');
INSERT INTO `demo_city_code` VALUES ('130824', '1308', '2', '滦平', 'l', 'luan ping', '130824000000', '滦平县', '101090406');
INSERT INTO `demo_city_code` VALUES ('130825', '1308', '2', '隆化', 'l', 'long hua', '130825000000', '隆化县', '101090407');
INSERT INTO `demo_city_code` VALUES ('130826', '1308', '2', '丰宁', 'f', 'feng ning', '130826000000', '丰宁满族自治县', '101090408');
INSERT INTO `demo_city_code` VALUES ('130827', '1308', '2', '宽城', 'k', 'kuan cheng', '130827000000', '宽城满族自治县', '101090409');
INSERT INTO `demo_city_code` VALUES ('130828', '1308', '2', '围场', 'w', 'wei chang', '130828000000', '围场满族蒙古族自治县', '101090410');
INSERT INTO `demo_city_code` VALUES ('130881', '1308', '2', '平泉', 'p', 'ping quan', '130881000000', '平泉市', '101090405');
INSERT INTO `demo_city_code` VALUES ('1309', '13', '1', '沧州', 'c', 'cang zhou', '130900000000', '沧州市', '101090701');
INSERT INTO `demo_city_code` VALUES ('130902', '1309', '2', '新华', 'x', 'xin hua', '130902000000', '新华区', NULL);
INSERT INTO `demo_city_code` VALUES ('130903', '1309', '2', '运河', 'y', 'yun he', '130903000000', '运河区', NULL);
INSERT INTO `demo_city_code` VALUES ('130921', '1309', '2', '沧县', 'c', 'cang xian', '130921000000', '沧县', '101090716');
INSERT INTO `demo_city_code` VALUES ('130922', '1309', '2', '青县', 'q', 'qing xian', '130922000000', '青县', '101090702');
INSERT INTO `demo_city_code` VALUES ('130923', '1309', '2', '东光', 'd', 'dong guang', '130923000000', '东光县', '101090703');
INSERT INTO `demo_city_code` VALUES ('130924', '1309', '2', '海兴', 'h', 'hai xing', '130924000000', '海兴县', '101090704');
INSERT INTO `demo_city_code` VALUES ('130925', '1309', '2', '盐山', 'y', 'yan shan', '130925000000', '盐山县', '101090705');
INSERT INTO `demo_city_code` VALUES ('130926', '1309', '2', '肃宁', 's', 'su ning', '130926000000', '肃宁县', '101090706');
INSERT INTO `demo_city_code` VALUES ('130927', '1309', '2', '南皮', 'n', 'nan pi', '130927000000', '南皮县', '101090707');
INSERT INTO `demo_city_code` VALUES ('130928', '1309', '2', '吴桥', 'w', 'wu qiao', '130928000000', '吴桥县', '101090708');
INSERT INTO `demo_city_code` VALUES ('130929', '1309', '2', '献县', 'x', 'xian xian', '130929000000', '献县', '101090709');
INSERT INTO `demo_city_code` VALUES ('130930', '1309', '2', '孟村', 'm', 'meng cun', '130930000000', '孟村回族自治县', '101090710');
INSERT INTO `demo_city_code` VALUES ('130981', '1309', '2', '泊头', 'b', 'bo tou', '130981000000', '泊头市', '101090711');
INSERT INTO `demo_city_code` VALUES ('130982', '1309', '2', '任丘', 'r', 'ren qiu', '130982000000', '任丘市', '101090712');
INSERT INTO `demo_city_code` VALUES ('130983', '1309', '2', '黄骅', 'h', 'huang hua', '130983000000', '黄骅市', '101090713');
INSERT INTO `demo_city_code` VALUES ('130984', '1309', '2', '河间', 'h', 'he jian', '130984000000', '河间市', '101090714');
INSERT INTO `demo_city_code` VALUES ('1310', '13', '1', '廊坊', 'l', 'lang fang', '131000000000', '廊坊市', '101090601');
INSERT INTO `demo_city_code` VALUES ('131002', '1310', '2', '安次', 'a', 'an ci', '131002000000', '安次区', NULL);
INSERT INTO `demo_city_code` VALUES ('131003', '1310', '2', '广阳', 'g', 'guang yang', '131003000000', '广阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('131022', '1310', '2', '固安', 'g', 'gu an', '131022000000', '固安县', '101090602');
INSERT INTO `demo_city_code` VALUES ('131023', '1310', '2', '永清', 'y', 'yong qing', '131023000000', '永清县', '101090603');
INSERT INTO `demo_city_code` VALUES ('131024', '1310', '2', '香河', 'x', 'xiang he', '131024000000', '香河县', '101090604');
INSERT INTO `demo_city_code` VALUES ('131025', '1310', '2', '大城', 'd', 'da cheng', '131025000000', '大城县', '101090605');
INSERT INTO `demo_city_code` VALUES ('131026', '1310', '2', '文安', 'w', 'wen an', '131026000000', '文安县', '101090606');
INSERT INTO `demo_city_code` VALUES ('131028', '1310', '2', '大厂', 'd', 'da chang', '131028000000', '大厂回族自治县', '101090607');
INSERT INTO `demo_city_code` VALUES ('131081', '1310', '2', '霸州', 'b', 'ba zhou', '131081000000', '霸州市', '101090608');
INSERT INTO `demo_city_code` VALUES ('131082', '1310', '2', '三河', 's', 'san he', '131082000000', '三河市', '101090609');
INSERT INTO `demo_city_code` VALUES ('1311', '13', '1', '衡水', 'h', 'heng shui', '131100000000', '衡水市', '101090801');
INSERT INTO `demo_city_code` VALUES ('131102', '1311', '2', '桃城', 't', 'tao cheng', '131102000000', '桃城区', NULL);
INSERT INTO `demo_city_code` VALUES ('131103', '1311', '2', '冀州', 'j', 'ji zhou', '131103000000', '冀州区', '101090810');
INSERT INTO `demo_city_code` VALUES ('131121', '1311', '2', '枣强', 'z', 'zao qiang', '131121000000', '枣强县', '101090802');
INSERT INTO `demo_city_code` VALUES ('131122', '1311', '2', '武邑', 'w', 'wu yi', '131122000000', '武邑县', '101090803');
INSERT INTO `demo_city_code` VALUES ('131123', '1311', '2', '武强', 'w', 'wu qiang', '131123000000', '武强县', '101090804');
INSERT INTO `demo_city_code` VALUES ('131124', '1311', '2', '饶阳', 'r', 'rao yang', '131124000000', '饶阳县', '101090805');
INSERT INTO `demo_city_code` VALUES ('131125', '1311', '2', '安平', 'a', 'an ping', '131125000000', '安平县', '101090806');
INSERT INTO `demo_city_code` VALUES ('131126', '1311', '2', '故城', 'g', 'gu cheng', '131126000000', '故城县', '101090807');
INSERT INTO `demo_city_code` VALUES ('131127', '1311', '2', '景县', 'j', 'jing xian', '131127000000', '景县', '101090808');
INSERT INTO `demo_city_code` VALUES ('131128', '1311', '2', '阜城', 'f', 'fu cheng', '131128000000', '阜城县', '101090809');
INSERT INTO `demo_city_code` VALUES ('131182', '1311', '2', '深州', 's', 'shen zhou', '131182000000', '深州市', '101090811');
INSERT INTO `demo_city_code` VALUES ('14', '0', '0', '山西', 's', 'shan xi', '140000000000', '山西省', NULL);
INSERT INTO `demo_city_code` VALUES ('1401', '14', '1', '太原', 't', 'tai yuan', '140100000000', '太原市', '101100101');
INSERT INTO `demo_city_code` VALUES ('140105', '1401', '2', '小店', 'x', 'xiao dian', '140105000000', '小店区', NULL);
INSERT INTO `demo_city_code` VALUES ('140106', '1401', '2', '迎泽', 'y', 'ying ze', '140106000000', '迎泽区', NULL);
INSERT INTO `demo_city_code` VALUES ('140107', '1401', '2', '杏花岭', 'x', 'xing hua ling', '140107000000', '杏花岭区', NULL);
INSERT INTO `demo_city_code` VALUES ('140108', '1401', '2', '尖草坪', 'j', 'jian cao ping', '140108000000', '尖草坪区', NULL);
INSERT INTO `demo_city_code` VALUES ('140109', '1401', '2', '万柏林', 'w', 'wan bai lin', '140109000000', '万柏林区', NULL);
INSERT INTO `demo_city_code` VALUES ('140110', '1401', '2', '晋源', 'j', 'jin yuan', '140110000000', '晋源区', NULL);
INSERT INTO `demo_city_code` VALUES ('140121', '1401', '2', '清徐', 'q', 'qing xu', '140121000000', '清徐县', '101100102');
INSERT INTO `demo_city_code` VALUES ('140122', '1401', '2', '阳曲', 'y', 'yang qu', '140122000000', '阳曲县', '101100103');
INSERT INTO `demo_city_code` VALUES ('140123', '1401', '2', '娄烦', 'l', 'lou fan', '140123000000', '娄烦县', '101100104');
INSERT INTO `demo_city_code` VALUES ('140181', '1401', '2', '古交', 'g', 'gu jiao', '140181000000', '古交市', '101100105');
INSERT INTO `demo_city_code` VALUES ('1402', '14', '1', '大同', 'd', 'da tong', '140200000000', '大同市', '101100201');
INSERT INTO `demo_city_code` VALUES ('140212', '1402', '2', '新荣', 'x', 'xin rong', '140212000000', '新荣区', NULL);
INSERT INTO `demo_city_code` VALUES ('140213', '1402', '2', '平城', 'p', 'ping cheng', '140213000000', '平城区', NULL);
INSERT INTO `demo_city_code` VALUES ('140214', '1402', '2', '云冈', 'y', 'yun gang', '140214000000', '云冈区', NULL);
INSERT INTO `demo_city_code` VALUES ('140215', '1402', '2', '云州', 'y', 'yun zhou', '140215000000', '云州区', NULL);
INSERT INTO `demo_city_code` VALUES ('140221', '1402', '2', '阳高', 'y', 'yang gao', '140221000000', '阳高县', '101100202');
INSERT INTO `demo_city_code` VALUES ('140222', '1402', '2', '天镇', 't', 'tian zhen', '140222000000', '天镇县', '101100204');
INSERT INTO `demo_city_code` VALUES ('140223', '1402', '2', '广灵', 'g', 'guang ling', '140223000000', '广灵县', '101100205');
INSERT INTO `demo_city_code` VALUES ('140224', '1402', '2', '灵丘', 'l', 'ling qiu', '140224000000', '灵丘县', '101100206');
INSERT INTO `demo_city_code` VALUES ('140225', '1402', '2', '浑源', 'h', 'hun yuan', '140225000000', '浑源县', '101100207');
INSERT INTO `demo_city_code` VALUES ('140226', '1402', '2', '左云', 'z', 'zuo yun', '140226000000', '左云县', '101100208');
INSERT INTO `demo_city_code` VALUES ('1403', '14', '1', '阳泉', 'y', 'yang quan', '140300000000', '阳泉市', '101100301');
INSERT INTO `demo_city_code` VALUES ('140302', '1403', '2', '城区', 'c', 'cheng qu', '140302000000', '城区', NULL);
INSERT INTO `demo_city_code` VALUES ('140303', '1403', '2', '矿区', 'k', 'kuang qu', '140303000000', '矿区', NULL);
INSERT INTO `demo_city_code` VALUES ('140311', '1403', '2', '郊区', 'j', 'jiao qu', '140311000000', '郊区', NULL);
INSERT INTO `demo_city_code` VALUES ('140321', '1403', '2', '平定', 'p', 'ping ding', '140321000000', '平定县', '101100303');
INSERT INTO `demo_city_code` VALUES ('140322', '1403', '2', '盂县', 'y', 'yu xian', '140322000000', '盂县', '101100302');
INSERT INTO `demo_city_code` VALUES ('1404', '14', '1', '长治', 'c', 'chang zhi', '140400000000', '长治市', '101100501');
INSERT INTO `demo_city_code` VALUES ('140403', '1404', '2', '潞州', 'l', 'lu zhou', '140403000000', '潞州区', NULL);
INSERT INTO `demo_city_code` VALUES ('140404', '1404', '2', '上党', 's', 'shang dang', '140404000000', '上党区', NULL);
INSERT INTO `demo_city_code` VALUES ('140405', '1404', '2', '屯留', 't', 'tun liu', '140405000000', '屯留区', '101100503');
INSERT INTO `demo_city_code` VALUES ('140406', '1404', '2', '潞城', 'l', 'lu cheng', '140406000000', '潞城区', '101100504');
INSERT INTO `demo_city_code` VALUES ('140423', '1404', '2', '襄垣', 'x', 'xiang yuan', '140423000000', '襄垣县', '101100505');
INSERT INTO `demo_city_code` VALUES ('140425', '1404', '2', '平顺', 'p', 'ping shun', '140425000000', '平顺县', '101100506');
INSERT INTO `demo_city_code` VALUES ('140426', '1404', '2', '黎城', 'l', 'li cheng', '140426000000', '黎城县', '101100502');
INSERT INTO `demo_city_code` VALUES ('140427', '1404', '2', '壶关', 'h', 'hu guan', '140427000000', '壶关县', '101100511');
INSERT INTO `demo_city_code` VALUES ('140428', '1404', '2', '长子', 'z', 'zhang zi', '140428000000', '长子县', '101100509');
INSERT INTO `demo_city_code` VALUES ('140429', '1404', '2', '武乡', 'w', 'wu xiang', '140429000000', '武乡县', '101100507');
INSERT INTO `demo_city_code` VALUES ('140430', '1404', '2', '沁县', 'q', 'qin xian', '140430000000', '沁县', '101100508');
INSERT INTO `demo_city_code` VALUES ('140431', '1404', '2', '沁源', 'q', 'qin yuan', '140431000000', '沁源县', '101100510');
INSERT INTO `demo_city_code` VALUES ('1405', '14', '1', '晋城', 'j', 'jin cheng', '140500000000', '晋城市', '101100601');
INSERT INTO `demo_city_code` VALUES ('140502', '1405', '2', '城区', 'c', 'cheng qu', '140502000000', '城区', NULL);
INSERT INTO `demo_city_code` VALUES ('140521', '1405', '2', '沁水', 'q', 'qin shui', '140521000000', '沁水县', '101100602');
INSERT INTO `demo_city_code` VALUES ('140522', '1405', '2', '阳城', 'y', 'yang cheng', '140522000000', '阳城县', '101100603');
INSERT INTO `demo_city_code` VALUES ('140524', '1405', '2', '陵川', 'l', 'ling chuan', '140524000000', '陵川县', '101100604');
INSERT INTO `demo_city_code` VALUES ('140525', '1405', '2', '泽州', 'z', 'ze zhou', '140525000000', '泽州县', '101100606');
INSERT INTO `demo_city_code` VALUES ('140581', '1405', '2', '高平', 'g', 'gao ping', '140581000000', '高平市', '101100605');
INSERT INTO `demo_city_code` VALUES ('1406', '14', '1', '朔州', 's', 'shuo zhou', '140600000000', '朔州市', '101100901');
INSERT INTO `demo_city_code` VALUES ('140602', '1406', '2', '朔城', 's', 'shuo cheng', '140602000000', '朔城区', NULL);
INSERT INTO `demo_city_code` VALUES ('140603', '1406', '2', '平鲁', 'p', 'ping lu', '140603000000', '平鲁区', '101100902');
INSERT INTO `demo_city_code` VALUES ('140621', '1406', '2', '山阴', 's', 'shan yin', '140621000000', '山阴县', '101100903');
INSERT INTO `demo_city_code` VALUES ('140622', '1406', '2', '应县', 'y', 'ying xian', '140622000000', '应县', '101100905');
INSERT INTO `demo_city_code` VALUES ('140623', '1406', '2', '右玉', 'y', 'you yu', '140623000000', '右玉县', '101100904');
INSERT INTO `demo_city_code` VALUES ('140681', '1406', '2', '怀仁', 'h', 'huai ren', '140681000000', '怀仁市', '101100906');
INSERT INTO `demo_city_code` VALUES ('1407', '14', '1', '晋中', 'j', 'jin zhong', '140700000000', '晋中市', '101100401');
INSERT INTO `demo_city_code` VALUES ('140702', '1407', '2', '榆次', 'y', 'yu ci', '140702000000', '榆次区', '101100402');
INSERT INTO `demo_city_code` VALUES ('140703', '1407', '2', '太谷', 't', 'tai gu', '140703000000', '太谷区', '101100408');
INSERT INTO `demo_city_code` VALUES ('140721', '1407', '2', '榆社', 'y', 'yu she', '140721000000', '榆社县', '101100403');
INSERT INTO `demo_city_code` VALUES ('140722', '1407', '2', '左权', 'z', 'zuo quan', '140722000000', '左权县', '101100404');
INSERT INTO `demo_city_code` VALUES ('140723', '1407', '2', '和顺', 'h', 'he shun', '140723000000', '和顺县', '101100405');
INSERT INTO `demo_city_code` VALUES ('140724', '1407', '2', '昔阳', 'x', 'xi yang', '140724000000', '昔阳县', '101100406');
INSERT INTO `demo_city_code` VALUES ('140725', '1407', '2', '寿阳', 's', 'shou yang', '140725000000', '寿阳县', '101100407');
INSERT INTO `demo_city_code` VALUES ('140727', '1407', '2', '祁县', 'q', 'qi xian', '140727000000', '祁县', '101100409');
INSERT INTO `demo_city_code` VALUES ('140728', '1407', '2', '平遥', 'p', 'ping yao', '140728000000', '平遥县', '101100410');
INSERT INTO `demo_city_code` VALUES ('140729', '1407', '2', '灵石', 'l', 'ling shi', '140729000000', '灵石县', '101100411');
INSERT INTO `demo_city_code` VALUES ('140781', '1407', '2', '介休', 'j', 'jie xiu', '140781000000', '介休市', '101100412');
INSERT INTO `demo_city_code` VALUES ('1408', '14', '1', '运城', 'y', 'yun cheng', '140800000000', '运城市', '101100801');
INSERT INTO `demo_city_code` VALUES ('140802', '1408', '2', '盐湖', 'y', 'yan hu', '140802000000', '盐湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('140821', '1408', '2', '临猗', 'l', 'lin yi', '140821000000', '临猗县', '101100802');
INSERT INTO `demo_city_code` VALUES ('140822', '1408', '2', '万荣', 'w', 'wan rong', '140822000000', '万荣县', '101100804');
INSERT INTO `demo_city_code` VALUES ('140823', '1408', '2', '闻喜', 'w', 'wen xi', '140823000000', '闻喜县', '101100808');
INSERT INTO `demo_city_code` VALUES ('140824', '1408', '2', '稷山', 'j', 'ji shan', '140824000000', '稷山县', '101100803');
INSERT INTO `demo_city_code` VALUES ('140825', '1408', '2', '新绛', 'x', 'xin jiang', '140825000000', '新绛县', '101100806');
INSERT INTO `demo_city_code` VALUES ('140826', '1408', '2', '绛县', 'j', 'jiang xian', '140826000000', '绛县', '101100807');
INSERT INTO `demo_city_code` VALUES ('140827', '1408', '2', '垣曲', 'y', 'yuan qu', '140827000000', '垣曲县', '101100809');
INSERT INTO `demo_city_code` VALUES ('140828', '1408', '2', '夏县', 'x', 'xia xian', '140828000000', '夏县', '101100812');
INSERT INTO `demo_city_code` VALUES ('140829', '1408', '2', '平陆', 'p', 'ping lu', '140829000000', '平陆县', '101100813');
INSERT INTO `demo_city_code` VALUES ('140830', '1408', '2', '芮城', 'r', 'rui cheng', '140830000000', '芮城县', '101100811');
INSERT INTO `demo_city_code` VALUES ('140881', '1408', '2', '永济', 'y', 'yong ji', '140881000000', '永济市', '101100810');
INSERT INTO `demo_city_code` VALUES ('140882', '1408', '2', '河津', 'h', 'he jin', '140882000000', '河津市', '101100805');
INSERT INTO `demo_city_code` VALUES ('1409', '14', '1', '忻州', 'x', 'xin zhou', '140900000000', '忻州市', '101101001');
INSERT INTO `demo_city_code` VALUES ('140902', '1409', '2', '忻府', 'x', 'xin fu', '140902000000', '忻府区', NULL);
INSERT INTO `demo_city_code` VALUES ('140921', '1409', '2', '定襄', 'd', 'ding xiang', '140921000000', '定襄县', '101101002');
INSERT INTO `demo_city_code` VALUES ('140922', '1409', '2', '五台', 'w', 'wu tai', '140922000000', '五台县', NULL);
INSERT INTO `demo_city_code` VALUES ('140923', '1409', '2', '代县', 'd', 'dai xian', '140923000000', '代县', '101101008');
INSERT INTO `demo_city_code` VALUES ('140924', '1409', '2', '繁峙', 'f', 'fan shi', '140924000000', '繁峙县', '101101009');
INSERT INTO `demo_city_code` VALUES ('140925', '1409', '2', '宁武', 'n', 'ning wu', '140925000000', '宁武县', '101101007');
INSERT INTO `demo_city_code` VALUES ('140926', '1409', '2', '静乐', 'j', 'jing le', '140926000000', '静乐县', '101101012');
INSERT INTO `demo_city_code` VALUES ('140927', '1409', '2', '神池', 's', 'shen chi', '140927000000', '神池县', '101101006');
INSERT INTO `demo_city_code` VALUES ('140928', '1409', '2', '五寨', 'w', 'wu zhai', '140928000000', '五寨县', '101101014');
INSERT INTO `demo_city_code` VALUES ('140929', '1409', '2', '岢岚', 'k', 'ke lan', '140929000000', '岢岚县', '101101013');
INSERT INTO `demo_city_code` VALUES ('140930', '1409', '2', '河曲', 'h', 'he qu', '140930000000', '河曲县', '101101004');
INSERT INTO `demo_city_code` VALUES ('140931', '1409', '2', '保德', 'b', 'bao de', '140931000000', '保德县', '101101011');
INSERT INTO `demo_city_code` VALUES ('140932', '1409', '2', '偏关', 'p', 'pian guan', '140932000000', '偏关县', '101101005');
INSERT INTO `demo_city_code` VALUES ('140981', '1409', '2', '原平', 'y', 'yuan ping', '140981000000', '原平市', '101101015');
INSERT INTO `demo_city_code` VALUES ('1410', '14', '1', '临汾', 'l', 'lin fen', '141000000000', '临汾市', '101100701');
INSERT INTO `demo_city_code` VALUES ('141002', '1410', '2', '尧都', 'y', 'yao du', '141002000000', '尧都区', NULL);
INSERT INTO `demo_city_code` VALUES ('141021', '1410', '2', '曲沃', 'q', 'qu wo', '141021000000', '曲沃县', '101100702');
INSERT INTO `demo_city_code` VALUES ('141022', '1410', '2', '翼城', 'y', 'yi cheng', '141022000000', '翼城县', '101100713');
INSERT INTO `demo_city_code` VALUES ('141023', '1410', '2', '襄汾', 'x', 'xiang fen', '141023000000', '襄汾县', '101100707');
INSERT INTO `demo_city_code` VALUES ('141024', '1410', '2', '洪洞', 'h', 'hong tong', '141024000000', '洪洞县', '101100710');
INSERT INTO `demo_city_code` VALUES ('141025', '1410', '2', '古县', 'g', 'gu xian', '141025000000', '古县', '101100717');
INSERT INTO `demo_city_code` VALUES ('141026', '1410', '2', '安泽', 'a', 'an ze', '141026000000', '安泽县', '101100716');
INSERT INTO `demo_city_code` VALUES ('141027', '1410', '2', '浮山', 'f', 'fu shan', '141027000000', '浮山县', '101100715');
INSERT INTO `demo_city_code` VALUES ('141028', '1410', '2', '吉县', 'j', 'ji xian', '141028000000', '吉县', '101100706');
INSERT INTO `demo_city_code` VALUES ('141029', '1410', '2', '乡宁', 'x', 'xiang ning', '141029000000', '乡宁县', '101100712');
INSERT INTO `demo_city_code` VALUES ('141030', '1410', '2', '大宁', 'd', 'da ning', '141030000000', '大宁县', '101100705');
INSERT INTO `demo_city_code` VALUES ('141031', '1410', '2', '隰县', 'x', 'xi xian', '141031000000', '隰县', '101100704');
INSERT INTO `demo_city_code` VALUES ('141032', '1410', '2', '永和', 'y', 'yong he', '141032000000', '永和县', '101100703');
INSERT INTO `demo_city_code` VALUES ('141033', '1410', '2', '蒲县', 'p', 'pu xian', '141033000000', '蒲县', '101100708');
INSERT INTO `demo_city_code` VALUES ('141034', '1410', '2', '汾西', 'f', 'fen xi', '141034000000', '汾西县', '101100709');
INSERT INTO `demo_city_code` VALUES ('141081', '1410', '2', '侯马', 'h', 'hou ma', '141081000000', '侯马市', '101100714');
INSERT INTO `demo_city_code` VALUES ('141082', '1410', '2', '霍州', 'h', 'huo zhou', '141082000000', '霍州市', '101100711');
INSERT INTO `demo_city_code` VALUES ('1411', '14', '1', '吕梁', 'l', 'lv liang', '141100000000', '吕梁市', '101101100');
INSERT INTO `demo_city_code` VALUES ('141102', '1411', '2', '离石', 'l', 'li shi', '141102000000', '离石区', '101101101');
INSERT INTO `demo_city_code` VALUES ('141121', '1411', '2', '文水', 'w', 'wen shui', '141121000000', '文水县', '101101112');
INSERT INTO `demo_city_code` VALUES ('141122', '1411', '2', '交城', 'j', 'jiao cheng', '141122000000', '交城县', '101101113');
INSERT INTO `demo_city_code` VALUES ('141123', '1411', '2', '兴县', 'x', 'xing xian', '141123000000', '兴县', '101101103');
INSERT INTO `demo_city_code` VALUES ('141124', '1411', '2', '临县', 'l', 'lin xian', '141124000000', '临县', '101101102');
INSERT INTO `demo_city_code` VALUES ('141125', '1411', '2', '柳林', 'l', 'liu lin', '141125000000', '柳林县', '101101105');
INSERT INTO `demo_city_code` VALUES ('141126', '1411', '2', '石楼', 's', 'shi lou', '141126000000', '石楼县', '101101106');
INSERT INTO `demo_city_code` VALUES ('141127', '1411', '2', '岚县', 'l', 'lan xian', '141127000000', '岚县', '101101104');
INSERT INTO `demo_city_code` VALUES ('141128', '1411', '2', '方山', 'f', 'fang shan', '141128000000', '方山县', '101101107');
INSERT INTO `demo_city_code` VALUES ('141129', '1411', '2', '中阳', 'z', 'zhong yang', '141129000000', '中阳县', '101101109');
INSERT INTO `demo_city_code` VALUES ('141130', '1411', '2', '交口', 'j', 'jiao kou', '141130000000', '交口县', '101101108');
INSERT INTO `demo_city_code` VALUES ('141181', '1411', '2', '孝义', 'x', 'xiao yi', '141181000000', '孝义市', '101101110');
INSERT INTO `demo_city_code` VALUES ('141182', '1411', '2', '汾阳', 'f', 'fen yang', '141182000000', '汾阳市', '101101111');
INSERT INTO `demo_city_code` VALUES ('15', '0', '0', '内蒙古', 'n', 'nei meng gu', '150000000000', '内蒙古自治区', NULL);
INSERT INTO `demo_city_code` VALUES ('1501', '15', '1', '呼和浩特', 'h', 'hu he hao te', '150100000000', '呼和浩特市', '101080101');
INSERT INTO `demo_city_code` VALUES ('150102', '1501', '2', '新城', 'x', 'xin cheng', '150102000000', '新城区', NULL);
INSERT INTO `demo_city_code` VALUES ('150103', '1501', '2', '回民', 'h', 'hui min', '150103000000', '回民区', NULL);
INSERT INTO `demo_city_code` VALUES ('150104', '1501', '2', '玉泉', 'y', 'yu quan', '150104000000', '玉泉区', NULL);
INSERT INTO `demo_city_code` VALUES ('150105', '1501', '2', '赛罕', 's', 'sai han', '150105000000', '赛罕区', NULL);
INSERT INTO `demo_city_code` VALUES ('150121', '1501', '2', '土默特左旗', 't', 'tu mo te zuo qi', '150121000000', '土默特左旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150122', '1501', '2', '托克托', 't', 'tuo ke tuo', '150122000000', '托克托县', NULL);
INSERT INTO `demo_city_code` VALUES ('150123', '1501', '2', '和林格尔', 'h', 'he lin ge er', '150123000000', '和林格尔县', NULL);
INSERT INTO `demo_city_code` VALUES ('150124', '1501', '2', '清水河', 'q', 'qing shui he', '150124000000', '清水河县', '101080105');
INSERT INTO `demo_city_code` VALUES ('150125', '1501', '2', '武川', 'w', 'wu chuan', '150125000000', '武川县', '101080107');
INSERT INTO `demo_city_code` VALUES ('1502', '15', '1', '包头', 'b', 'bao tou', '150200000000', '包头市', '101080201');
INSERT INTO `demo_city_code` VALUES ('150202', '1502', '2', '东河', 'd', 'dong he', '150202000000', '东河区', NULL);
INSERT INTO `demo_city_code` VALUES ('150203', '1502', '2', '昆都仑', 'k', 'kun du lun', '150203000000', '昆都仑区', NULL);
INSERT INTO `demo_city_code` VALUES ('150204', '1502', '2', '青山', 'q', 'qing shan', '150204000000', '青山区', NULL);
INSERT INTO `demo_city_code` VALUES ('150205', '1502', '2', '石拐', 's', 'shi guai', '150205000000', '石拐区', NULL);
INSERT INTO `demo_city_code` VALUES ('150206', '1502', '2', '白云鄂博矿区', 'b', 'bai yun e bo kuang qu', '150206000000', '白云鄂博矿区', NULL);
INSERT INTO `demo_city_code` VALUES ('150207', '1502', '2', '九原', 'j', 'jiu yuan', '150207000000', '九原区', NULL);
INSERT INTO `demo_city_code` VALUES ('150221', '1502', '2', '土默特右旗', 't', 'tu mo te you qi', '150221000000', '土默特右旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150222', '1502', '2', '固阳', 'g', 'gu yang', '150222000000', '固阳县', '101080205');
INSERT INTO `demo_city_code` VALUES ('150223', '1502', '2', '达尔罕茂明安联合旗', 'd', 'da er han mao ming an lian he qi', '150223000000', '达尔罕茂明安联合旗', NULL);
INSERT INTO `demo_city_code` VALUES ('1503', '15', '1', '乌海', 'w', 'wu hai', '150300000000', '乌海市', '101080301');
INSERT INTO `demo_city_code` VALUES ('150302', '1503', '2', '海勃湾', 'h', 'hai bo wan', '150302000000', '海勃湾区', NULL);
INSERT INTO `demo_city_code` VALUES ('150303', '1503', '2', '海南', 'h', 'hai nan', '150303000000', '海南区', '101150401');
INSERT INTO `demo_city_code` VALUES ('150304', '1503', '2', '乌达', 'w', 'wu da', '150304000000', '乌达区', NULL);
INSERT INTO `demo_city_code` VALUES ('1504', '15', '1', '赤峰', 'c', 'chi feng', '150400000000', '赤峰市', '101080601');
INSERT INTO `demo_city_code` VALUES ('150402', '1504', '2', '红山', 'h', 'hong shan', '150402000000', '红山区', NULL);
INSERT INTO `demo_city_code` VALUES ('150403', '1504', '2', '元宝山', 'y', 'yuan bao shan', '150403000000', '元宝山区', NULL);
INSERT INTO `demo_city_code` VALUES ('150404', '1504', '2', '松山', 's', 'song shan', '150404000000', '松山区', NULL);
INSERT INTO `demo_city_code` VALUES ('150421', '1504', '2', '阿鲁科尔沁旗', 'a', 'a lu ke er qin qi', '150421000000', '阿鲁科尔沁旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150422', '1504', '2', '巴林左旗', 'b', 'ba lin zuo qi', '150422000000', '巴林左旗', '101080605');
INSERT INTO `demo_city_code` VALUES ('150423', '1504', '2', '巴林右旗', 'b', 'ba lin you qi', '150423000000', '巴林右旗', '101080606');
INSERT INTO `demo_city_code` VALUES ('150424', '1504', '2', '林西', 'l', 'lin xi', '150424000000', '林西县', '101080607');
INSERT INTO `demo_city_code` VALUES ('150425', '1504', '2', '克什克腾旗', 'k', 'ke shi ke teng qi', '150425000000', '克什克腾旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150426', '1504', '2', '翁牛特旗', 'w', 'weng niu te qi', '150426000000', '翁牛特旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150428', '1504', '2', '喀喇沁旗', 'k', 'ka la qin qi', '150428000000', '喀喇沁旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150429', '1504', '2', '宁城', 'n', 'ning cheng', '150429000000', '宁城县', '101080613');
INSERT INTO `demo_city_code` VALUES ('150430', '1504', '2', '敖汉旗', 'a', 'ao han qi', '150430000000', '敖汉旗', NULL);
INSERT INTO `demo_city_code` VALUES ('1505', '15', '1', '通辽', 't', 'tong liao', '150500000000', '通辽市', '101080501');
INSERT INTO `demo_city_code` VALUES ('150502', '1505', '2', '科尔沁', 'k', 'ke er qin', '150502000000', '科尔沁区', NULL);
INSERT INTO `demo_city_code` VALUES ('150521', '1505', '2', '科尔沁左翼中旗', 'k', 'ke er qin zuo yi zhong qi', '150521000000', '科尔沁左翼中旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150522', '1505', '2', '科尔沁左翼后旗', 'k', 'ke er qin zuo yi hou qi', '150522000000', '科尔沁左翼后旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150523', '1505', '2', '开鲁', 'k', 'kai lu', '150523000000', '开鲁县', '101080506');
INSERT INTO `demo_city_code` VALUES ('150524', '1505', '2', '库伦旗', 'k', 'ku lun qi', '150524000000', '库伦旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150525', '1505', '2', '奈曼旗', 'n', 'nai man qi', '150525000000', '奈曼旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150526', '1505', '2', '扎鲁特旗', 'z', 'zha lu te qi', '150526000000', '扎鲁特旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150581', '1505', '2', '霍林郭勒', 'h', 'huo lin guo le', '150581000000', '霍林郭勒市', '101081108');
INSERT INTO `demo_city_code` VALUES ('1506', '15', '1', '鄂尔多斯', 'e', 'e er duo si', '150600000000', '鄂尔多斯市', '101080701');
INSERT INTO `demo_city_code` VALUES ('150602', '1506', '2', '东胜', 'd', 'dong sheng', '150602000000', '东胜区', '101080713');
INSERT INTO `demo_city_code` VALUES ('150603', '1506', '2', '康巴什', 'k', 'kang ba shi', '150603000000', '康巴什区', NULL);
INSERT INTO `demo_city_code` VALUES ('150621', '1506', '2', '达拉特旗', 'd', 'da la te qi', '150621000000', '达拉特旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150622', '1506', '2', '准格尔旗', 'z', 'zhun ge er qi', '150622000000', '准格尔旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150623', '1506', '2', '鄂托克前旗', 'e', 'e tuo ke qian qi', '150623000000', '鄂托克前旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150624', '1506', '2', '鄂托克旗', 'e', 'e tuo ke qi', '150624000000', '鄂托克旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150625', '1506', '2', '杭锦旗', 'h', 'hang jin qi', '150625000000', '杭锦旗', '101080709');
INSERT INTO `demo_city_code` VALUES ('150626', '1506', '2', '乌审旗', 'w', 'wu shen qi', '150626000000', '乌审旗', '101080710');
INSERT INTO `demo_city_code` VALUES ('150627', '1506', '2', '伊金霍洛旗', 'y', 'yi jin huo luo qi', '150627000000', '伊金霍洛旗', NULL);
INSERT INTO `demo_city_code` VALUES ('1507', '15', '1', '呼伦贝尔', 'h', 'hu lun bei er', '150700000000', '呼伦贝尔市', NULL);
INSERT INTO `demo_city_code` VALUES ('150702', '1507', '2', '海拉尔', 'h', 'hai la er', '150702000000', '海拉尔区', '101081001');
INSERT INTO `demo_city_code` VALUES ('150703', '1507', '2', '扎赉诺尔区', 'z', 'zha lai nuo er qu', '150703000000', '扎赉诺尔区', NULL);
INSERT INTO `demo_city_code` VALUES ('150721', '1507', '2', '阿荣旗', 'a', 'a rong qi', '150721000000', '阿荣旗', '101081003');
INSERT INTO `demo_city_code` VALUES ('150722', '1507', '2', '莫力达瓦', 'm', 'mo li da wa', '150722000000', '莫力达瓦达斡尔族自治旗', '101081004');
INSERT INTO `demo_city_code` VALUES ('150723', '1507', '2', '鄂伦春自治旗', 'e', 'e lun chun zi zhi qi', '150723000000', '鄂伦春自治旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150724', '1507', '2', '鄂温克族自治旗', 'e', 'e wen ke zu zi zhi qi', '150724000000', '鄂温克族自治旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150725', '1507', '2', '陈巴尔虎旗', 'c', 'chen ba er hu qi', '150725000000', '陈巴尔虎旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150726', '1507', '2', '新巴尔虎左旗', 'x', 'xin ba er hu zuo qi', '150726000000', '新巴尔虎左旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150727', '1507', '2', '新巴尔虎右旗', 'x', 'xin ba er hu you qi', '150727000000', '新巴尔虎右旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150781', '1507', '2', '满洲里', 'm', 'man zhou li', '150781000000', '满洲里市', '101081010');
INSERT INTO `demo_city_code` VALUES ('150782', '1507', '2', '牙克石', 'y', 'ya ke shi', '150782000000', '牙克石市', '101081011');
INSERT INTO `demo_city_code` VALUES ('150783', '1507', '2', '扎兰屯', 'z', 'zha lan tun', '150783000000', '扎兰屯市', '101081012');
INSERT INTO `demo_city_code` VALUES ('150784', '1507', '2', '额尔古纳', 'e', 'e er gu na', '150784000000', '额尔古纳市', '101081014');
INSERT INTO `demo_city_code` VALUES ('150785', '1507', '2', '根河', 'g', 'gen he', '150785000000', '根河市', '101081015');
INSERT INTO `demo_city_code` VALUES ('1508', '15', '1', '巴彦淖尔', 'b', 'ba yan nao er', '150800000000', '巴彦淖尔市', NULL);
INSERT INTO `demo_city_code` VALUES ('150802', '1508', '2', '临河', 'l', 'lin he', '150802000000', '临河区', '101080801');
INSERT INTO `demo_city_code` VALUES ('150821', '1508', '2', '五原', 'w', 'wu yuan', '150821000000', '五原县', '101080802');
INSERT INTO `demo_city_code` VALUES ('150822', '1508', '2', '磴口', 'd', 'deng kou', '150822000000', '磴口县', '101080803');
INSERT INTO `demo_city_code` VALUES ('150823', '1508', '2', '乌拉特前旗', 'w', 'wu la te qian qi', '150823000000', '乌拉特前旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150824', '1508', '2', '乌拉特中旗', 'w', 'wu la te zhong qi', '150824000000', '乌拉特中旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150825', '1508', '2', '乌拉特后旗', 'w', 'wu la te hou qi', '150825000000', '乌拉特后旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150826', '1508', '2', '杭锦后旗', 'h', 'hang jin hou qi', '150826000000', '杭锦后旗', '101080810');
INSERT INTO `demo_city_code` VALUES ('1509', '15', '1', '乌兰察布', 'w', 'wu lan cha bu', '150900000000', '乌兰察布市', NULL);
INSERT INTO `demo_city_code` VALUES ('150902', '1509', '2', '集宁', 'j', 'ji ning', '150902000000', '集宁区', '101080401');
INSERT INTO `demo_city_code` VALUES ('150921', '1509', '2', '卓资', 'z', 'zhuo zi', '150921000000', '卓资县', '101080402');
INSERT INTO `demo_city_code` VALUES ('150922', '1509', '2', '化德', 'h', 'hua de', '150922000000', '化德县', '101080403');
INSERT INTO `demo_city_code` VALUES ('150923', '1509', '2', '商都', 's', 'shang du', '150923000000', '商都县', '101080404');
INSERT INTO `demo_city_code` VALUES ('150924', '1509', '2', '兴和', 'x', 'xing he', '150924000000', '兴和县', '101080406');
INSERT INTO `demo_city_code` VALUES ('150925', '1509', '2', '凉城', 'l', 'liang cheng', '150925000000', '凉城县', '101080407');
INSERT INTO `demo_city_code` VALUES ('150926', '1509', '2', '察哈尔右翼前旗', 'c', 'cha ha er you yi qian qi', '150926000000', '察哈尔右翼前旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150927', '1509', '2', '察哈尔右翼中旗', 'c', 'cha ha er you yi zhong qi', '150927000000', '察哈尔右翼中旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150928', '1509', '2', '察哈尔右翼后旗', 'c', 'cha ha er you yi hou qi', '150928000000', '察哈尔右翼后旗', NULL);
INSERT INTO `demo_city_code` VALUES ('150929', '1509', '2', '四子王旗', 's', 'si zi wang qi', '150929000000', '四子王旗', '101080411');
INSERT INTO `demo_city_code` VALUES ('150981', '1509', '2', '丰镇', 'f', 'feng zhen', '150981000000', '丰镇市', '101080412');
INSERT INTO `demo_city_code` VALUES ('1522', '15', '1', '兴安', 'x', 'xing an', '152200000000', '兴安盟', '101300506');
INSERT INTO `demo_city_code` VALUES ('152201', '1522', '2', '乌兰浩特', 'w', 'wu lan hao te', '152201000000', '乌兰浩特市', '101081101');
INSERT INTO `demo_city_code` VALUES ('152202', '1522', '2', '阿尔山', 'a', 'a er shan', '152202000000', '阿尔山市', '101081102');
INSERT INTO `demo_city_code` VALUES ('152221', '1522', '2', '科尔沁右翼前旗', 'k', 'ke er qin you yi qian qi', '152221000000', '科尔沁右翼前旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152222', '1522', '2', '科尔沁右翼中旗', 'k', 'ke er qin you yi zhong qi', '152222000000', '科尔沁右翼中旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152223', '1522', '2', '扎赉特旗', 'z', 'zha lai te qi', '152223000000', '扎赉特旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152224', '1522', '2', '突泉', 't', 'tu quan', '152224000000', '突泉县', '101081107');
INSERT INTO `demo_city_code` VALUES ('1525', '15', '1', '锡林郭勒', 'x', 'xi lin guo le', '152500000000', '锡林郭勒盟', NULL);
INSERT INTO `demo_city_code` VALUES ('152501', '1525', '2', '二连浩特', 'e', 'er lian hao te', '152501000000', '二连浩特市', '101080903');
INSERT INTO `demo_city_code` VALUES ('152502', '1525', '2', '锡林浩特', 'x', 'xi lin hao te', '152502000000', '锡林浩特市', '101080901');
INSERT INTO `demo_city_code` VALUES ('152522', '1525', '2', '阿巴嘎旗', 'a', 'a ba ga qi', '152522000000', '阿巴嘎旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152523', '1525', '2', '苏尼特左旗', 's', 'su ni te zuo qi', '152523000000', '苏尼特左旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152524', '1525', '2', '苏尼特右旗', 's', 'su ni te you qi', '152524000000', '苏尼特右旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152525', '1525', '2', '东乌珠穆沁旗', 'd', 'dong wu zhu mu qin qi', '152525000000', '东乌珠穆沁旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152526', '1525', '2', '西乌珠穆沁旗', 'x', 'xi wu zhu mu qin qi', '152526000000', '西乌珠穆沁旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152527', '1525', '2', '太仆寺旗', 't', 'tai pu si qi', '152527000000', '太仆寺旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152528', '1525', '2', '镶黄旗', 'x', 'xiang huang qi', '152528000000', '镶黄旗', '101080912');
INSERT INTO `demo_city_code` VALUES ('152529', '1525', '2', '正镶白旗', 'z', 'zheng xiang bai qi', '152529000000', '正镶白旗', '101080913');
INSERT INTO `demo_city_code` VALUES ('152530', '1525', '2', '正蓝旗', 'z', 'zheng lan qi', '152530000000', '正蓝旗', '101080914');
INSERT INTO `demo_city_code` VALUES ('152531', '1525', '2', '多伦', 'd', 'duo lun', '152531000000', '多伦县', '101080915');
INSERT INTO `demo_city_code` VALUES ('1529', '15', '1', '阿拉善', 'a', 'a la shan', '152900000000', '阿拉善盟', NULL);
INSERT INTO `demo_city_code` VALUES ('152921', '1529', '2', '阿拉善左旗', 'a', 'a la shan zuo qi', '152921000000', '阿拉善左旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152922', '1529', '2', '阿拉善右旗', 'a', 'a la shan you qi', '152922000000', '阿拉善右旗', NULL);
INSERT INTO `demo_city_code` VALUES ('152923', '1529', '2', '额济纳旗', 'e', 'e ji na qi', '152923000000', '额济纳旗', NULL);
INSERT INTO `demo_city_code` VALUES ('21', '0', '0', '辽宁', 'l', 'liao ning', '210000000000', '辽宁省', NULL);
INSERT INTO `demo_city_code` VALUES ('2101', '21', '1', '沈阳', 's', 'shen yang', '210100000000', '沈阳市', '101070101');
INSERT INTO `demo_city_code` VALUES ('210102', '2101', '2', '和平', 'h', 'he ping', '210102000000', '和平区', '101281204');
INSERT INTO `demo_city_code` VALUES ('210103', '2101', '2', '沈河', 's', 'shen he', '210103000000', '沈河区', NULL);
INSERT INTO `demo_city_code` VALUES ('210104', '2101', '2', '大东', 'd', 'da dong', '210104000000', '大东区', NULL);
INSERT INTO `demo_city_code` VALUES ('210105', '2101', '2', '皇姑', 'h', 'huang gu', '210105000000', '皇姑区', NULL);
INSERT INTO `demo_city_code` VALUES ('210106', '2101', '2', '铁西', 't', 'tie xi', '210106000000', '铁西区', NULL);
INSERT INTO `demo_city_code` VALUES ('210111', '2101', '2', '苏家屯', 's', 'su jia tun', '210111000000', '苏家屯区', NULL);
INSERT INTO `demo_city_code` VALUES ('210112', '2101', '2', '浑南', 'h', 'hun nan', '210112000000', '浑南区', NULL);
INSERT INTO `demo_city_code` VALUES ('210113', '2101', '2', '沈北新区', 's', 'shen bei xin qu', '210113000000', '沈北新区', NULL);
INSERT INTO `demo_city_code` VALUES ('210114', '2101', '2', '于洪', 'y', 'yu hong', '210114000000', '于洪区', NULL);
INSERT INTO `demo_city_code` VALUES ('210115', '2101', '2', '辽中', 'l', 'liao zhong', '210115000000', '辽中区', '101070103');
INSERT INTO `demo_city_code` VALUES ('210123', '2101', '2', '康平', 'k', 'kang ping', '210123000000', '康平县', '101070104');
INSERT INTO `demo_city_code` VALUES ('210124', '2101', '2', '法库', 'f', 'fa ku', '210124000000', '法库县', '101070105');
INSERT INTO `demo_city_code` VALUES ('210181', '2101', '2', '新民', 'x', 'xin min', '210181000000', '新民市', '101070106');
INSERT INTO `demo_city_code` VALUES ('2102', '21', '1', '大连', 'd', 'da lian', '210200000000', '大连市', '101070201');
INSERT INTO `demo_city_code` VALUES ('210202', '2102', '2', '中山', 'z', 'zhong shan', '210202000000', '中山区', '101281701');
INSERT INTO `demo_city_code` VALUES ('210203', '2102', '2', '西岗', 'x', 'xi gang', '210203000000', '西岗区', NULL);
INSERT INTO `demo_city_code` VALUES ('210204', '2102', '2', '沙河口', 's', 'sha he kou', '210204000000', '沙河口区', NULL);
INSERT INTO `demo_city_code` VALUES ('210211', '2102', '2', '甘井子', 'g', 'gan jing zi', '210211000000', '甘井子区', NULL);
INSERT INTO `demo_city_code` VALUES ('210212', '2102', '2', '旅顺口', 'l', 'lv shun kou', '210212000000', '旅顺口区', NULL);
INSERT INTO `demo_city_code` VALUES ('210213', '2102', '2', '金州', 'j', 'jin zhou', '210213000000', '金州区', '101070203');
INSERT INTO `demo_city_code` VALUES ('210214', '2102', '2', '普兰店', 'p', 'pu lan dian', '210214000000', '普兰店区', '101070204');
INSERT INTO `demo_city_code` VALUES ('210224', '2102', '2', '长海', 'c', 'chang hai', '210224000000', '长海县', '101070206');
INSERT INTO `demo_city_code` VALUES ('210281', '2102', '2', '瓦房店', 'w', 'wa fang dian', '210281000000', '瓦房店市', '101070202');
INSERT INTO `demo_city_code` VALUES ('210283', '2102', '2', '庄河', 'z', 'zhuang he', '210283000000', '庄河市', '101070207');
INSERT INTO `demo_city_code` VALUES ('2103', '21', '1', '鞍山', 'a', 'an shan', '210300000000', '鞍山市', '101070301');
INSERT INTO `demo_city_code` VALUES ('210302', '2103', '2', '铁东', 't', 'tie dong', '210302000000', '铁东区', NULL);
INSERT INTO `demo_city_code` VALUES ('210303', '2103', '2', '铁西', 't', 'tie xi', '210303000000', '铁西区', NULL);
INSERT INTO `demo_city_code` VALUES ('210304', '2103', '2', '立山', 'l', 'li shan', '210304000000', '立山区', NULL);
INSERT INTO `demo_city_code` VALUES ('210311', '2103', '2', '千山', 'q', 'qian shan', '210311000000', '千山区', NULL);
INSERT INTO `demo_city_code` VALUES ('210321', '2103', '2', '台安', 't', 'tai an', '210321000000', '台安县', '101070302');
INSERT INTO `demo_city_code` VALUES ('210323', '2103', '2', '岫岩', 'x', 'xiu yan', '210323000000', '岫岩满族自治县', '101070303');
INSERT INTO `demo_city_code` VALUES ('210381', '2103', '2', '海城', 'h', 'hai cheng', '210381000000', '海城市', '101070304');
INSERT INTO `demo_city_code` VALUES ('2104', '21', '1', '抚顺', 'f', 'fu shun', '210400000000', '抚顺市', '101070401');
INSERT INTO `demo_city_code` VALUES ('210402', '2104', '2', '新抚', 'x', 'xin fu', '210402000000', '新抚区', NULL);
INSERT INTO `demo_city_code` VALUES ('210403', '2104', '2', '东洲', 'd', 'dong zhou', '210403000000', '东洲区', NULL);
INSERT INTO `demo_city_code` VALUES ('210404', '2104', '2', '望花', 'w', 'wang hua', '210404000000', '望花区', NULL);
INSERT INTO `demo_city_code` VALUES ('210411', '2104', '2', '顺城', 's', 'shun cheng', '210411000000', '顺城区', NULL);
INSERT INTO `demo_city_code` VALUES ('210421', '2104', '2', '抚顺县', 'f', 'fu shun xian', '210421000000', '抚顺县', NULL);
INSERT INTO `demo_city_code` VALUES ('210422', '2104', '2', '新宾', 'x', 'xin bin', '210422000000', '新宾满族自治县', '101070402');
INSERT INTO `demo_city_code` VALUES ('210423', '2104', '2', '清原', 'q', 'qing yuan', '210423000000', '清原满族自治县', '101070403');
INSERT INTO `demo_city_code` VALUES ('2105', '21', '1', '本溪', 'b', 'ben xi', '210500000000', '本溪市', '101070501');
INSERT INTO `demo_city_code` VALUES ('210502', '2105', '2', '平山', 'p', 'ping shan', '210502000000', '平山区', '101090111');
INSERT INTO `demo_city_code` VALUES ('210503', '2105', '2', '溪湖', 'x', 'xi hu', '210503000000', '溪湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('210504', '2105', '2', '明山', 'm', 'ming shan', '210504000000', '明山区', NULL);
INSERT INTO `demo_city_code` VALUES ('210505', '2105', '2', '南芬', 'n', 'nan fen', '210505000000', '南芬区', NULL);
INSERT INTO `demo_city_code` VALUES ('210521', '2105', '2', '本溪满族自治县', 'b', 'ben xi man zu zi zhi xian', '210521000000', '本溪满族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('210522', '2105', '2', '桓仁', 'h', 'huan ren', '210522000000', '桓仁满族自治县', '101070504');
INSERT INTO `demo_city_code` VALUES ('2106', '21', '1', '丹东', 'd', 'dan dong', '210600000000', '丹东市', '101070601');
INSERT INTO `demo_city_code` VALUES ('210602', '2106', '2', '元宝', 'y', 'yuan bao', '210602000000', '元宝区', NULL);
INSERT INTO `demo_city_code` VALUES ('210603', '2106', '2', '振兴', 'z', 'zhen xing', '210603000000', '振兴区', NULL);
INSERT INTO `demo_city_code` VALUES ('210604', '2106', '2', '振安', 'z', 'zhen an', '210604000000', '振安区', NULL);
INSERT INTO `demo_city_code` VALUES ('210624', '2106', '2', '宽甸', 'k', 'kuan dian', '210624000000', '宽甸满族自治县', '101070603');
INSERT INTO `demo_city_code` VALUES ('210681', '2106', '2', '东港', 'd', 'dong gang', '210681000000', '东港市', '101070604');
INSERT INTO `demo_city_code` VALUES ('210682', '2106', '2', '凤城', 'f', 'feng cheng', '210682000000', '凤城市', '101070602');
INSERT INTO `demo_city_code` VALUES ('2107', '21', '1', '锦州', 'j', 'jin zhou', '210700000000', '锦州市', '101070701');
INSERT INTO `demo_city_code` VALUES ('210702', '2107', '2', '古塔', 'g', 'gu ta', '210702000000', '古塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('210703', '2107', '2', '凌河', 'l', 'ling he', '210703000000', '凌河区', NULL);
INSERT INTO `demo_city_code` VALUES ('210711', '2107', '2', '太和', 't', 'tai he', '210711000000', '太和区', '101220806');
INSERT INTO `demo_city_code` VALUES ('210726', '2107', '2', '黑山', 'h', 'hei shan', '210726000000', '黑山县', '101070705');
INSERT INTO `demo_city_code` VALUES ('210727', '2107', '2', '义县', 'y', 'yi xian', '210727000000', '义县', '101070704');
INSERT INTO `demo_city_code` VALUES ('210781', '2107', '2', '凌海', 'l', 'ling hai', '210781000000', '凌海市', '101070702');
INSERT INTO `demo_city_code` VALUES ('210782', '2107', '2', '北镇', 'b', 'bei zhen', '210782000000', '北镇市', '101070706');
INSERT INTO `demo_city_code` VALUES ('2108', '21', '1', '营口', 'y', 'ying kou', '210800000000', '营口市', '101070801');
INSERT INTO `demo_city_code` VALUES ('210802', '2108', '2', '站前', 'z', 'zhan qian', '210802000000', '站前区', NULL);
INSERT INTO `demo_city_code` VALUES ('210803', '2108', '2', '西市', 'x', 'xi shi', '210803000000', '西市区', NULL);
INSERT INTO `demo_city_code` VALUES ('210804', '2108', '2', '鲅鱼圈', 'b', 'ba yu quan', '210804000000', '鲅鱼圈区', NULL);
INSERT INTO `demo_city_code` VALUES ('210811', '2108', '2', '老边', 'l', 'lao bian', '210811000000', '老边区', NULL);
INSERT INTO `demo_city_code` VALUES ('210881', '2108', '2', '盖州', 'g', 'gai zhou', '210881000000', '盖州市', '101070803');
INSERT INTO `demo_city_code` VALUES ('210882', '2108', '2', '大石桥', 'd', 'da shi qiao', '210882000000', '大石桥市', '101070802');
INSERT INTO `demo_city_code` VALUES ('2109', '21', '1', '阜新', 'f', 'fu xin', '210900000000', '阜新市', '101070901');
INSERT INTO `demo_city_code` VALUES ('210902', '2109', '2', '海州', 'h', 'hai zhou', '210902000000', '海州区', NULL);
INSERT INTO `demo_city_code` VALUES ('210903', '2109', '2', '新邱', 'x', 'xin qiu', '210903000000', '新邱区', NULL);
INSERT INTO `demo_city_code` VALUES ('210904', '2109', '2', '太平', 't', 'tai ping', '210904000000', '太平区', NULL);
INSERT INTO `demo_city_code` VALUES ('210905', '2109', '2', '清河门', 'q', 'qing he men', '210905000000', '清河门区', NULL);
INSERT INTO `demo_city_code` VALUES ('210911', '2109', '2', '细河', 'x', 'xi he', '210911000000', '细河区', NULL);
INSERT INTO `demo_city_code` VALUES ('210921', '2109', '2', '阜新蒙古族自治县', 'f', 'fu xin meng gu zu zi zhi xian', '210921000000', '阜新蒙古族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('210922', '2109', '2', '彰武', 'z', 'zhang wu', '210922000000', '彰武县', '101070902');
INSERT INTO `demo_city_code` VALUES ('2110', '21', '1', '辽阳', 'l', 'liao yang', '211000000000', '辽阳市', '101071001');
INSERT INTO `demo_city_code` VALUES ('211002', '2110', '2', '白塔', 'b', 'bai ta', '211002000000', '白塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('211003', '2110', '2', '文圣', 'w', 'wen sheng', '211003000000', '文圣区', NULL);
INSERT INTO `demo_city_code` VALUES ('211004', '2110', '2', '宏伟', 'h', 'hong wei', '211004000000', '宏伟区', NULL);
INSERT INTO `demo_city_code` VALUES ('211005', '2110', '2', '弓长岭', 'g', 'gong chang ling', '211005000000', '弓长岭区', '101071004');
INSERT INTO `demo_city_code` VALUES ('211011', '2110', '2', '太子河', 't', 'tai zi he', '211011000000', '太子河区', NULL);
INSERT INTO `demo_city_code` VALUES ('211021', '2110', '2', '辽阳县', 'l', 'liao yang xian', '211021000000', '辽阳县', '101071002');
INSERT INTO `demo_city_code` VALUES ('211081', '2110', '2', '灯塔', 'd', 'deng ta', '211081000000', '灯塔市', '101071003');
INSERT INTO `demo_city_code` VALUES ('2111', '21', '1', '盘锦', 'p', 'pan jin', '211100000000', '盘锦市', '101071301');
INSERT INTO `demo_city_code` VALUES ('211102', '2111', '2', '双台子', 's', 'shuang tai zi', '211102000000', '双台子区', NULL);
INSERT INTO `demo_city_code` VALUES ('211103', '2111', '2', '兴隆台', 'x', 'xing long tai', '211103000000', '兴隆台区', NULL);
INSERT INTO `demo_city_code` VALUES ('211104', '2111', '2', '大洼', 'd', 'da wa', '211104000000', '大洼区', '101071302');
INSERT INTO `demo_city_code` VALUES ('211122', '2111', '2', '盘山', 'p', 'pan shan', '211122000000', '盘山县', '101071303');
INSERT INTO `demo_city_code` VALUES ('2112', '21', '1', '铁岭', 't', 'tie ling', '211200000000', '铁岭市', '101071101');
INSERT INTO `demo_city_code` VALUES ('211202', '2112', '2', '银州', 'y', 'yin zhou', '211202000000', '银州区', NULL);
INSERT INTO `demo_city_code` VALUES ('211204', '2112', '2', '清河', 'q', 'qing he', '211204000000', '清河区', '101090914');
INSERT INTO `demo_city_code` VALUES ('211221', '2112', '2', '铁岭县', 't', 'tie ling xian', '211221000000', '铁岭县', NULL);
INSERT INTO `demo_city_code` VALUES ('211223', '2112', '2', '西丰', 'x', 'xi feng', '211223000000', '西丰县', '101071104');
INSERT INTO `demo_city_code` VALUES ('211224', '2112', '2', '昌图', 'c', 'chang tu', '211224000000', '昌图县', '101071103');
INSERT INTO `demo_city_code` VALUES ('211281', '2112', '2', '调兵山', 'd', 'diao bing shan', '211281000000', '调兵山市', '101071105');
INSERT INTO `demo_city_code` VALUES ('211282', '2112', '2', '开原', 'k', 'kai yuan', '211282000000', '开原市', '101071102');
INSERT INTO `demo_city_code` VALUES ('2113', '21', '1', '朝阳', 'c', 'chao yang', '211300000000', '朝阳市', '101071201');
INSERT INTO `demo_city_code` VALUES ('211302', '2113', '2', '双塔', 's', 'shuang ta', '211302000000', '双塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('211303', '2113', '2', '龙城', 'l', 'long cheng', '211303000000', '龙城区', NULL);
INSERT INTO `demo_city_code` VALUES ('211321', '2113', '2', '朝阳县', 'c', 'chao yang xian', '211321000000', '朝阳县', NULL);
INSERT INTO `demo_city_code` VALUES ('211322', '2113', '2', '建平', 'j', 'jian ping', '211322000000', '建平县', NULL);
INSERT INTO `demo_city_code` VALUES ('211324', '2113', '2', '喀喇沁左翼', 'k', 'ka la qin zuo yi', '211324000000', '喀喇沁左翼蒙古族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('211381', '2113', '2', '北票', 'b', 'bei piao', '211381000000', '北票市', '101071205');
INSERT INTO `demo_city_code` VALUES ('211382', '2113', '2', '凌源', 'l', 'ling yuan', '211382000000', '凌源市', '101071203');
INSERT INTO `demo_city_code` VALUES ('2114', '21', '1', '葫芦岛', 'h', 'hu lu dao', '211400000000', '葫芦岛市', '101071401');
INSERT INTO `demo_city_code` VALUES ('211402', '2114', '2', '连山', 'l', 'lian shan', '211402000000', '连山区', '101281304');
INSERT INTO `demo_city_code` VALUES ('211403', '2114', '2', '龙港', 'l', 'long gang', '211403000000', '龙港区', NULL);
INSERT INTO `demo_city_code` VALUES ('211404', '2114', '2', '南票', 'n', 'nan piao', '211404000000', '南票区', NULL);
INSERT INTO `demo_city_code` VALUES ('211421', '2114', '2', '绥中', 's', 'sui zhong', '211421000000', '绥中县', '101071403');
INSERT INTO `demo_city_code` VALUES ('211422', '2114', '2', '建昌', 'j', 'jian chang', '211422000000', '建昌县', '101071402');
INSERT INTO `demo_city_code` VALUES ('211481', '2114', '2', '兴城', 'x', 'xing cheng', '211481000000', '兴城市', '101071404');
INSERT INTO `demo_city_code` VALUES ('22', '0', '0', '吉林', 'j', 'ji lin', '220000000000', '吉林省', '101060201');
INSERT INTO `demo_city_code` VALUES ('2201', '22', '1', '长春', 'c', 'chang chun', '220100000000', '长春市', '101060101');
INSERT INTO `demo_city_code` VALUES ('220102', '2201', '2', '南关', 'n', 'nan guan', '220102000000', '南关区', NULL);
INSERT INTO `demo_city_code` VALUES ('220103', '2201', '2', '宽城', 'k', 'kuan cheng', '220103000000', '宽城区', '101090409');
INSERT INTO `demo_city_code` VALUES ('220104', '2201', '2', '朝阳', 'c', 'chao yang', '220104000000', '朝阳区', '101010300');
INSERT INTO `demo_city_code` VALUES ('220105', '2201', '2', '二道', 'e', 'er dao', '220105000000', '二道区', NULL);
INSERT INTO `demo_city_code` VALUES ('220106', '2201', '2', '绿园', 'l', 'lv yuan', '220106000000', '绿园区', NULL);
INSERT INTO `demo_city_code` VALUES ('220112', '2201', '2', '双阳', 's', 'shuang yang', '220112000000', '双阳区', '101060106');
INSERT INTO `demo_city_code` VALUES ('220113', '2201', '2', '九台', 'j', 'jiu tai', '220113000000', '九台区', '101060104');
INSERT INTO `demo_city_code` VALUES ('220122', '2201', '2', '农安', 'n', 'nong an', '220122000000', '农安县', '101060102');
INSERT INTO `demo_city_code` VALUES ('220182', '2201', '2', '榆树', 'y', 'yu shu', '220182000000', '榆树市', '101060105');
INSERT INTO `demo_city_code` VALUES ('220183', '2201', '2', '德惠', 'd', 'de hui', '220183000000', '德惠市', '101060103');
INSERT INTO `demo_city_code` VALUES ('220184', '2201', '2', '公主岭', 'g', 'gong zhu ling', '220184000000', '公主岭市', '101060404');
INSERT INTO `demo_city_code` VALUES ('2202', '22', '1', '吉林市', 'j', 'ji lin shi', '220200000000', '吉林市', NULL);
INSERT INTO `demo_city_code` VALUES ('220202', '2202', '2', '昌邑', 'c', 'chang yi', '220202000000', '昌邑区', '101120606');
INSERT INTO `demo_city_code` VALUES ('220203', '2202', '2', '龙潭', 'l', 'long tan', '220203000000', '龙潭区', NULL);
INSERT INTO `demo_city_code` VALUES ('220204', '2202', '2', '船营', 'c', 'chuan ying', '220204000000', '船营区', NULL);
INSERT INTO `demo_city_code` VALUES ('220211', '2202', '2', '丰满', 'f', 'feng man', '220211000000', '丰满区', NULL);
INSERT INTO `demo_city_code` VALUES ('220221', '2202', '2', '永吉', 'y', 'yong ji', '220221000000', '永吉县', '101060203');
INSERT INTO `demo_city_code` VALUES ('220281', '2202', '2', '蛟河', 'j', 'jiao he', '220281000000', '蛟河市', '101060204');
INSERT INTO `demo_city_code` VALUES ('220282', '2202', '2', '桦甸', 'h', 'hua dian', '220282000000', '桦甸市', '101060206');
INSERT INTO `demo_city_code` VALUES ('220283', '2202', '2', '舒兰', 's', 'shu lan', '220283000000', '舒兰市', '101060202');
INSERT INTO `demo_city_code` VALUES ('220284', '2202', '2', '磐石', 'p', 'pan shi', '220284000000', '磐石市', '101060205');
INSERT INTO `demo_city_code` VALUES ('2203', '22', '1', '四平', 's', 'si ping', '220300000000', '四平市', '101060401');
INSERT INTO `demo_city_code` VALUES ('220302', '2203', '2', '铁西', 't', 'tie xi', '220302000000', '铁西区', NULL);
INSERT INTO `demo_city_code` VALUES ('220303', '2203', '2', '铁东', 't', 'tie dong', '220303000000', '铁东区', NULL);
INSERT INTO `demo_city_code` VALUES ('220322', '2203', '2', '梨树', 'l', 'li shu', '220322000000', '梨树县', '101060403');
INSERT INTO `demo_city_code` VALUES ('220323', '2203', '2', '伊通', 'y', 'yi tong', '220323000000', '伊通满族自治县', '101060405');
INSERT INTO `demo_city_code` VALUES ('220382', '2203', '2', '双辽', 's', 'shuang liao', '220382000000', '双辽市', '101060402');
INSERT INTO `demo_city_code` VALUES ('2204', '22', '1', '辽源', 'l', 'liao yuan', '220400000000', '辽源市', '101060701');
INSERT INTO `demo_city_code` VALUES ('220402', '2204', '2', '龙山', 'l', 'long shan', '220402000000', '龙山区', '101251507');
INSERT INTO `demo_city_code` VALUES ('220403', '2204', '2', '西安', 'x', 'xi an', '220403000000', '西安区', '101110101');
INSERT INTO `demo_city_code` VALUES ('220421', '2204', '2', '东丰', 'd', 'dong feng', '220421000000', '东丰县', '101060702');
INSERT INTO `demo_city_code` VALUES ('220422', '2204', '2', '东辽', 'd', 'dong liao', '220422000000', '东辽县', '101060703');
INSERT INTO `demo_city_code` VALUES ('2205', '22', '1', '通化', 't', 'tong hua', '220500000000', '通化市', '101060501');
INSERT INTO `demo_city_code` VALUES ('220502', '2205', '2', '东昌', 'd', 'dong chang', '220502000000', '东昌区', NULL);
INSERT INTO `demo_city_code` VALUES ('220503', '2205', '2', '二道江', 'e', 'er dao jiang', '220503000000', '二道江区', NULL);
INSERT INTO `demo_city_code` VALUES ('220521', '2205', '2', '通化县', 't', 'tong hua xian', '220521000000', '通化县', '101060506');
INSERT INTO `demo_city_code` VALUES ('220523', '2205', '2', '辉南', 'h', 'hui nan', '220523000000', '辉南县', '101060504');
INSERT INTO `demo_city_code` VALUES ('220524', '2205', '2', '柳河', 'l', 'liu he', '220524000000', '柳河县', '101060503');
INSERT INTO `demo_city_code` VALUES ('220581', '2205', '2', '梅河口', 'm', 'mei he kou', '220581000000', '梅河口市', '101060502');
INSERT INTO `demo_city_code` VALUES ('220582', '2205', '2', '集安', 'j', 'ji an', '220582000000', '集安市', '101060505');
INSERT INTO `demo_city_code` VALUES ('2206', '22', '1', '白山', 'b', 'bai shan', '220600000000', '白山市', '101060901');
INSERT INTO `demo_city_code` VALUES ('220602', '2206', '2', '浑江', 'h', 'hun jiang', '220602000000', '浑江区', NULL);
INSERT INTO `demo_city_code` VALUES ('220605', '2206', '2', '江源', 'j', 'jiang yuan', '220605000000', '江源区', '101060907');
INSERT INTO `demo_city_code` VALUES ('220621', '2206', '2', '抚松', 'f', 'fu song', '220621000000', '抚松县', '101060906');
INSERT INTO `demo_city_code` VALUES ('220622', '2206', '2', '靖宇', 'j', 'jing yu', '220622000000', '靖宇县', '101060902');
INSERT INTO `demo_city_code` VALUES ('220623', '2206', '2', '长白', 'c', 'chang bai', '220623000000', '长白朝鲜族自治县', '101060905');
INSERT INTO `demo_city_code` VALUES ('220681', '2206', '2', '临江', 'l', 'lin jiang', '220681000000', '临江市', '101060903');
INSERT INTO `demo_city_code` VALUES ('2207', '22', '1', '松原', 's', 'song yuan', '220700000000', '松原市', '101060801');
INSERT INTO `demo_city_code` VALUES ('220702', '2207', '2', '宁江', 'n', 'ning jiang', '220702000000', '宁江区', NULL);
INSERT INTO `demo_city_code` VALUES ('220721', '2207', '2', '前郭尔罗斯', 'q', 'qian guo er luo si', '220721000000', '前郭尔罗斯蒙古族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('220722', '2207', '2', '长岭', 'c', 'chang ling', '220722000000', '长岭县', '101060804');
INSERT INTO `demo_city_code` VALUES ('220723', '2207', '2', '乾安', 'q', 'qian an', '220723000000', '乾安县', '101060802');
INSERT INTO `demo_city_code` VALUES ('220781', '2207', '2', '扶余', 'f', 'fu yu', '220781000000', '扶余市', '101060805');
INSERT INTO `demo_city_code` VALUES ('2208', '22', '1', '白城', 'b', 'bai cheng', '220800000000', '白城市', '101060601');
INSERT INTO `demo_city_code` VALUES ('220802', '2208', '2', '洮北', 't', 'tao bei', '220802000000', '洮北区', NULL);
INSERT INTO `demo_city_code` VALUES ('220821', '2208', '2', '镇赉', 'z', 'zhen lai', '220821000000', '镇赉县', '101060604');
INSERT INTO `demo_city_code` VALUES ('220822', '2208', '2', '通榆', 't', 'tong yu', '220822000000', '通榆县', '101060605');
INSERT INTO `demo_city_code` VALUES ('220881', '2208', '2', '洮南', 't', 'tao nan', '220881000000', '洮南市', '101060602');
INSERT INTO `demo_city_code` VALUES ('220882', '2208', '2', '大安', 'd', 'da an', '220882000000', '大安市', '101060603');
INSERT INTO `demo_city_code` VALUES ('2224', '22', '1', '延边', 'y', 'yan bian', '222400000000', '延边朝鲜族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('222401', '2224', '2', '延吉', 'y', 'yan ji', '222401000000', '延吉市', '101060301');
INSERT INTO `demo_city_code` VALUES ('222402', '2224', '2', '图们', 't', 'tu men', '222402000000', '图们市', '101060309');
INSERT INTO `demo_city_code` VALUES ('222403', '2224', '2', '敦化', 'd', 'dun hua', '222403000000', '敦化市', '101060302');
INSERT INTO `demo_city_code` VALUES ('222404', '2224', '2', '珲春', 'h', 'hun chun', '222404000000', '珲春市', '101060308');
INSERT INTO `demo_city_code` VALUES ('222405', '2224', '2', '龙井', 'l', 'long jing', '222405000000', '龙井市', '101060307');
INSERT INTO `demo_city_code` VALUES ('222406', '2224', '2', '和龙', 'h', 'he long', '222406000000', '和龙市', '101060305');
INSERT INTO `demo_city_code` VALUES ('222424', '2224', '2', '汪清', 'w', 'wang qing', '222424000000', '汪清县', '101060304');
INSERT INTO `demo_city_code` VALUES ('222426', '2224', '2', '安图', 'a', 'an tu', '222426000000', '安图县', '101060303');
INSERT INTO `demo_city_code` VALUES ('23', '0', '0', '黑龙江', 'h', 'hei long jiang', '230000000000', '黑龙江省', NULL);
INSERT INTO `demo_city_code` VALUES ('2301', '23', '1', '哈尔滨', 'h', 'ha er bin', '230100000000', '哈尔滨市', '101050101');
INSERT INTO `demo_city_code` VALUES ('230102', '2301', '2', '道里', 'd', 'dao li', '230102000000', '道里区', NULL);
INSERT INTO `demo_city_code` VALUES ('230103', '2301', '2', '南岗', 'n', 'nan gang', '230103000000', '南岗区', NULL);
INSERT INTO `demo_city_code` VALUES ('230104', '2301', '2', '道外', 'd', 'dao wai', '230104000000', '道外区', NULL);
INSERT INTO `demo_city_code` VALUES ('230108', '2301', '2', '平房', 'p', 'ping fang', '230108000000', '平房区', NULL);
INSERT INTO `demo_city_code` VALUES ('230109', '2301', '2', '松北', 's', 'song bei', '230109000000', '松北区', NULL);
INSERT INTO `demo_city_code` VALUES ('230110', '2301', '2', '香坊', 'x', 'xiang fang', '230110000000', '香坊区', NULL);
INSERT INTO `demo_city_code` VALUES ('230111', '2301', '2', '呼兰', 'h', 'hu lan', '230111000000', '呼兰区', '101050103');
INSERT INTO `demo_city_code` VALUES ('230112', '2301', '2', '阿城', 'a', 'a cheng', '230112000000', '阿城区', '101050104');
INSERT INTO `demo_city_code` VALUES ('230113', '2301', '2', '双城', 's', 'shuang cheng', '230113000000', '双城区', '101050102');
INSERT INTO `demo_city_code` VALUES ('230123', '2301', '2', '依兰', 'y', 'yi lan', '230123000000', '依兰县', '101050106');
INSERT INTO `demo_city_code` VALUES ('230124', '2301', '2', '方正', 'f', 'fang zheng', '230124000000', '方正县', '101050109');
INSERT INTO `demo_city_code` VALUES ('230125', '2301', '2', '宾县', 'b', 'bin xian', '230125000000', '宾县', '101050105');
INSERT INTO `demo_city_code` VALUES ('230126', '2301', '2', '巴彦', 'b', 'ba yan', '230126000000', '巴彦县', '101050107');
INSERT INTO `demo_city_code` VALUES ('230127', '2301', '2', '木兰', 'm', 'mu lan', '230127000000', '木兰县', '101050113');
INSERT INTO `demo_city_code` VALUES ('230128', '2301', '2', '通河', 't', 'tong he', '230128000000', '通河县', '101050108');
INSERT INTO `demo_city_code` VALUES ('230129', '2301', '2', '延寿', 'y', 'yan shou', '230129000000', '延寿县', '101050110');
INSERT INTO `demo_city_code` VALUES ('230183', '2301', '2', '尚志', 's', 'shang zhi', '230183000000', '尚志市', '101050111');
INSERT INTO `demo_city_code` VALUES ('230184', '2301', '2', '五常', 'w', 'wu chang', '230184000000', '五常市', '101050112');
INSERT INTO `demo_city_code` VALUES ('2302', '23', '1', '齐齐哈尔', 'q', 'qi qi ha er', '230200000000', '齐齐哈尔市', '101050201');
INSERT INTO `demo_city_code` VALUES ('230202', '2302', '2', '龙沙', 'l', 'long sha', '230202000000', '龙沙区', NULL);
INSERT INTO `demo_city_code` VALUES ('230203', '2302', '2', '建华', 'j', 'jian hua', '230203000000', '建华区', NULL);
INSERT INTO `demo_city_code` VALUES ('230204', '2302', '2', '铁锋', 't', 'tie feng', '230204000000', '铁锋区', NULL);
INSERT INTO `demo_city_code` VALUES ('230205', '2302', '2', '昂昂溪', 'a', 'ang ang xi', '230205000000', '昂昂溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('230206', '2302', '2', '富拉尔基区', 'f', 'fu la er ji qu', '230206000000', '富拉尔基区', NULL);
INSERT INTO `demo_city_code` VALUES ('230207', '2302', '2', '碾子山', 'n', 'nian zi shan', '230207000000', '碾子山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230208', '2302', '2', '梅里斯', 'm', 'mei li si', '230208000000', '梅里斯达斡尔族区', NULL);
INSERT INTO `demo_city_code` VALUES ('230221', '2302', '2', '龙江', 'l', 'long jiang', '230221000000', '龙江县', '101050203');
INSERT INTO `demo_city_code` VALUES ('230223', '2302', '2', '依安', 'y', 'yi an', '230223000000', '依安县', '101050206');
INSERT INTO `demo_city_code` VALUES ('230224', '2302', '2', '泰来', 't', 'tai lai', '230224000000', '泰来县', '101050210');
INSERT INTO `demo_city_code` VALUES ('230225', '2302', '2', '甘南', 'g', 'gan nan', '230225000000', '甘南县', '101050204');
INSERT INTO `demo_city_code` VALUES ('230227', '2302', '2', '富裕', 'f', 'fu yu', '230227000000', '富裕县', '101050205');
INSERT INTO `demo_city_code` VALUES ('230229', '2302', '2', '克山', 'k', 'ke shan', '230229000000', '克山县', '101050208');
INSERT INTO `demo_city_code` VALUES ('230230', '2302', '2', '克东', 'k', 'ke dong', '230230000000', '克东县', '101050209');
INSERT INTO `demo_city_code` VALUES ('230231', '2302', '2', '拜泉', 'b', 'bai quan', '230231000000', '拜泉县', '101050207');
INSERT INTO `demo_city_code` VALUES ('230281', '2302', '2', '讷河', 'n', 'ne he', '230281000000', '讷河市', '101050202');
INSERT INTO `demo_city_code` VALUES ('2303', '23', '1', '鸡西', 'j', 'ji xi', '230300000000', '鸡西市', '101051101');
INSERT INTO `demo_city_code` VALUES ('230302', '2303', '2', '鸡冠', 'j', 'ji guan', '230302000000', '鸡冠区', NULL);
INSERT INTO `demo_city_code` VALUES ('230303', '2303', '2', '恒山', 'h', 'heng shan', '230303000000', '恒山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230304', '2303', '2', '滴道', 'd', 'di dao', '230304000000', '滴道区', NULL);
INSERT INTO `demo_city_code` VALUES ('230305', '2303', '2', '梨树', 'l', 'li shu', '230305000000', '梨树区', '101060403');
INSERT INTO `demo_city_code` VALUES ('230306', '2303', '2', '城子河', 'c', 'cheng zi he', '230306000000', '城子河区', NULL);
INSERT INTO `demo_city_code` VALUES ('230307', '2303', '2', '麻山', 'm', 'ma shan', '230307000000', '麻山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230321', '2303', '2', '鸡东', 'j', 'ji dong', '230321000000', '鸡东县', '101051104');
INSERT INTO `demo_city_code` VALUES ('230381', '2303', '2', '虎林', 'h', 'hu lin', '230381000000', '虎林市', '101051102');
INSERT INTO `demo_city_code` VALUES ('230382', '2303', '2', '密山', 'm', 'mi shan', '230382000000', '密山市', '101051103');
INSERT INTO `demo_city_code` VALUES ('2304', '23', '1', '鹤岗', 'h', 'he gang', '230400000000', '鹤岗市', '101051201');
INSERT INTO `demo_city_code` VALUES ('230402', '2304', '2', '向阳', 'x', 'xiang yang', '230402000000', '向阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('230403', '2304', '2', '工农', 'g', 'gong nong', '230403000000', '工农区', NULL);
INSERT INTO `demo_city_code` VALUES ('230404', '2304', '2', '南山', 'n', 'nan shan', '230404000000', '南山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230405', '2304', '2', '兴安', 'x', 'xing an', '230405000000', '兴安区', '101300506');
INSERT INTO `demo_city_code` VALUES ('230406', '2304', '2', '东山', 'd', 'dong shan', '230406000000', '东山区', '101230608');
INSERT INTO `demo_city_code` VALUES ('230407', '2304', '2', '兴山', 'x', 'xing shan', '230407000000', '兴山区', '101200904');
INSERT INTO `demo_city_code` VALUES ('230421', '2304', '2', '萝北', 'l', 'luo bei', '230421000000', '萝北县', '101051203');
INSERT INTO `demo_city_code` VALUES ('230422', '2304', '2', '绥滨', 's', 'sui bin', '230422000000', '绥滨县', '101051202');
INSERT INTO `demo_city_code` VALUES ('2305', '23', '1', '双鸭山', 's', 'shuang ya shan', '230500000000', '双鸭山市', '101051301');
INSERT INTO `demo_city_code` VALUES ('230502', '2305', '2', '尖山', 'j', 'jian shan', '230502000000', '尖山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230503', '2305', '2', '岭东', 'l', 'ling dong', '230503000000', '岭东区', NULL);
INSERT INTO `demo_city_code` VALUES ('230505', '2305', '2', '四方台', 's', 'si fang tai', '230505000000', '四方台区', NULL);
INSERT INTO `demo_city_code` VALUES ('230506', '2305', '2', '宝山', 'b', 'bao shan', '230506000000', '宝山区', '101020300');
INSERT INTO `demo_city_code` VALUES ('230521', '2305', '2', '集贤', 'j', 'ji xian', '230521000000', '集贤县', '101051302');
INSERT INTO `demo_city_code` VALUES ('230522', '2305', '2', '友谊', 'y', 'you yi', '230522000000', '友谊县', '101051305');
INSERT INTO `demo_city_code` VALUES ('230523', '2305', '2', '宝清', 'b', 'bao qing', '230523000000', '宝清县', '101051303');
INSERT INTO `demo_city_code` VALUES ('230524', '2305', '2', '饶河', 'r', 'rao he', '230524000000', '饶河县', '101051304');
INSERT INTO `demo_city_code` VALUES ('2306', '23', '1', '大庆', 'd', 'da qing', '230600000000', '大庆市', '101050901');
INSERT INTO `demo_city_code` VALUES ('230602', '2306', '2', '萨尔图', 's', 'sa er tu', '230602000000', '萨尔图区', NULL);
INSERT INTO `demo_city_code` VALUES ('230603', '2306', '2', '龙凤', 'l', 'long feng', '230603000000', '龙凤区', NULL);
INSERT INTO `demo_city_code` VALUES ('230604', '2306', '2', '让胡路', 'r', 'rang hu lu', '230604000000', '让胡路区', NULL);
INSERT INTO `demo_city_code` VALUES ('230605', '2306', '2', '红岗', 'h', 'hong gang', '230605000000', '红岗区', NULL);
INSERT INTO `demo_city_code` VALUES ('230606', '2306', '2', '大同', 'd', 'da tong', '230606000000', '大同区', '101100201');
INSERT INTO `demo_city_code` VALUES ('230621', '2306', '2', '肇州', 'z', 'zhao zhou', '230621000000', '肇州县', '101050903');
INSERT INTO `demo_city_code` VALUES ('230622', '2306', '2', '肇源', 'z', 'zhao yuan', '230622000000', '肇源县', '101050904');
INSERT INTO `demo_city_code` VALUES ('230623', '2306', '2', '林甸', 'l', 'lin dian', '230623000000', '林甸县', '101050902');
INSERT INTO `demo_city_code` VALUES ('230624', '2306', '2', '杜尔伯特', 'd', 'du er bo te', '230624000000', '杜尔伯特蒙古族自治县', '101050905');
INSERT INTO `demo_city_code` VALUES ('2307', '23', '1', '伊春', 'y', 'yi chun', '230700000000', '伊春市', '101050801');
INSERT INTO `demo_city_code` VALUES ('230717', '2307', '2', '伊美', 'y', 'yi mei', '230717000000', '伊美区', NULL);
INSERT INTO `demo_city_code` VALUES ('230718', '2307', '2', '乌翠', 'w', 'wu cui', '230718000000', '乌翠区', NULL);
INSERT INTO `demo_city_code` VALUES ('230719', '2307', '2', '友好', 'y', 'you hao', '230719000000', '友好区', NULL);
INSERT INTO `demo_city_code` VALUES ('230722', '2307', '2', '嘉荫', 'j', 'jia yin', '230722000000', '嘉荫县', '101050805');
INSERT INTO `demo_city_code` VALUES ('230723', '2307', '2', '汤旺', 't', 'tang wang', '230723000000', '汤旺县', NULL);
INSERT INTO `demo_city_code` VALUES ('230724', '2307', '2', '丰林', 'f', 'feng lin', '230724000000', '丰林县', NULL);
INSERT INTO `demo_city_code` VALUES ('230725', '2307', '2', '大箐山', 'd', 'da qing shan', '230725000000', '大箐山县', NULL);
INSERT INTO `demo_city_code` VALUES ('230726', '2307', '2', '南岔', 'n', 'nan cha', '230726000000', '南岔县', NULL);
INSERT INTO `demo_city_code` VALUES ('230751', '2307', '2', '金林', 'j', 'jin lin', '230751000000', '金林区', NULL);
INSERT INTO `demo_city_code` VALUES ('230781', '2307', '2', '铁力', 't', 'tie li', '230781000000', '铁力市', '101050804');
INSERT INTO `demo_city_code` VALUES ('2308', '23', '1', '佳木斯', 'j', 'jia mu si', '230800000000', '佳木斯市', '101050401');
INSERT INTO `demo_city_code` VALUES ('230803', '2308', '2', '向阳', 'x', 'xiang yang', '230803000000', '向阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('230804', '2308', '2', '前进', 'q', 'qian jin', '230804000000', '前进区', NULL);
INSERT INTO `demo_city_code` VALUES ('230805', '2308', '2', '东风', 'd', 'dong feng', '230805000000', '东风区', NULL);
INSERT INTO `demo_city_code` VALUES ('230811', '2308', '2', '郊区', 'j', 'jiao qu', '230811000000', '郊区', NULL);
INSERT INTO `demo_city_code` VALUES ('230822', '2308', '2', '桦南', 'h', 'hua nan', '230822000000', '桦南县', '101050405');
INSERT INTO `demo_city_code` VALUES ('230826', '2308', '2', '桦川', 'h', 'hua chuan', '230826000000', '桦川县', '101050404');
INSERT INTO `demo_city_code` VALUES ('230828', '2308', '2', '汤原', 't', 'tang yuan', '230828000000', '汤原县', '101050402');
INSERT INTO `demo_city_code` VALUES ('230881', '2308', '2', '同江', 't', 'tong jiang', '230881000000', '同江市', '101050406');
INSERT INTO `demo_city_code` VALUES ('230882', '2308', '2', '富锦', 'f', 'fu jin', '230882000000', '富锦市', '101050407');
INSERT INTO `demo_city_code` VALUES ('230883', '2308', '2', '抚远', 'f', 'fu yuan', '230883000000', '抚远市', '101050403');
INSERT INTO `demo_city_code` VALUES ('2309', '23', '1', '七台河', 'q', 'qi tai he', '230900000000', '七台河市', '101051002');
INSERT INTO `demo_city_code` VALUES ('230902', '2309', '2', '新兴', 'x', 'xin xing', '230902000000', '新兴区', '101281403');
INSERT INTO `demo_city_code` VALUES ('230903', '2309', '2', '桃山', 't', 'tao shan', '230903000000', '桃山区', NULL);
INSERT INTO `demo_city_code` VALUES ('230904', '2309', '2', '茄子河', 'q', 'qie zi he', '230904000000', '茄子河区', NULL);
INSERT INTO `demo_city_code` VALUES ('230921', '2309', '2', '勃利', 'b', 'bo li', '230921000000', '勃利县', '101051003');
INSERT INTO `demo_city_code` VALUES ('2310', '23', '1', '牡丹江', 'm', 'mu dan jiang', '231000000000', '牡丹江市', '101050301');
INSERT INTO `demo_city_code` VALUES ('231002', '2310', '2', '东安', 'd', 'dong an', '231002000000', '东安区', '101251403');
INSERT INTO `demo_city_code` VALUES ('231003', '2310', '2', '阳明', 'y', 'yang ming', '231003000000', '阳明区', NULL);
INSERT INTO `demo_city_code` VALUES ('231004', '2310', '2', '爱民', 'a', 'ai min', '231004000000', '爱民区', NULL);
INSERT INTO `demo_city_code` VALUES ('231005', '2310', '2', '西安', 'x', 'xi an', '231005000000', '西安区', '101110101');
INSERT INTO `demo_city_code` VALUES ('231025', '2310', '2', '林口', 'l', 'lin kou', '231025000000', '林口县', '101050304');
INSERT INTO `demo_city_code` VALUES ('231081', '2310', '2', '绥芬河', 's', 'sui fen he', '231081000000', '绥芬河市', '101050305');
INSERT INTO `demo_city_code` VALUES ('231083', '2310', '2', '海林', 'h', 'hai lin', '231083000000', '海林市', '101050302');
INSERT INTO `demo_city_code` VALUES ('231084', '2310', '2', '宁安', 'n', 'ning an', '231084000000', '宁安市', '101050306');
INSERT INTO `demo_city_code` VALUES ('231085', '2310', '2', '穆棱', 'm', 'mu ling', '231085000000', '穆棱市', '101050303');
INSERT INTO `demo_city_code` VALUES ('231086', '2310', '2', '东宁', 'd', 'dong ning', '231086000000', '东宁市', '101050307');
INSERT INTO `demo_city_code` VALUES ('2311', '23', '1', '黑河', 'h', 'hei he', '231100000000', '黑河市', '101050601');
INSERT INTO `demo_city_code` VALUES ('231102', '2311', '2', '爱辉', 'a', 'ai hui', '231102000000', '爱辉区', NULL);
INSERT INTO `demo_city_code` VALUES ('231123', '2311', '2', '逊克', 'x', 'xun ke', '231123000000', '逊克县', '101050604');
INSERT INTO `demo_city_code` VALUES ('231124', '2311', '2', '孙吴', 's', 'sun wu', '231124000000', '孙吴县', '101050603');
INSERT INTO `demo_city_code` VALUES ('231181', '2311', '2', '北安', 'b', 'bei an', '231181000000', '北安市', '101050606');
INSERT INTO `demo_city_code` VALUES ('231182', '2311', '2', '五大连池', 'w', 'wu da lian chi', '231182000000', '五大连池市', '101050605');
INSERT INTO `demo_city_code` VALUES ('231183', '2311', '2', '嫩江', 'n', 'nen jiang', '231183000000', '嫩江市', '101050602');
INSERT INTO `demo_city_code` VALUES ('2312', '23', '1', '绥化', 's', 'sui hua', '231200000000', '绥化市', '101050501');
INSERT INTO `demo_city_code` VALUES ('231202', '2312', '2', '北林', 'b', 'bei lin', '231202000000', '北林区', NULL);
INSERT INTO `demo_city_code` VALUES ('231221', '2312', '2', '望奎', 'w', 'wang kui', '231221000000', '望奎县', '101050506');
INSERT INTO `demo_city_code` VALUES ('231222', '2312', '2', '兰西', 'l', 'lan xi', '231222000000', '兰西县', '101050507');
INSERT INTO `demo_city_code` VALUES ('231223', '2312', '2', '青冈', 'q', 'qing gang', '231223000000', '青冈县', '101050508');
INSERT INTO `demo_city_code` VALUES ('231224', '2312', '2', '庆安', 'q', 'qing an', '231224000000', '庆安县', '101050509');
INSERT INTO `demo_city_code` VALUES ('231225', '2312', '2', '明水', 'm', 'ming shui', '231225000000', '明水县', '101050505');
INSERT INTO `demo_city_code` VALUES ('231226', '2312', '2', '绥棱', 's', 'sui ling', '231226000000', '绥棱县', '101050510');
INSERT INTO `demo_city_code` VALUES ('231281', '2312', '2', '安达', 'a', 'an da', '231281000000', '安达市', '101050503');
INSERT INTO `demo_city_code` VALUES ('231282', '2312', '2', '肇东', 'z', 'zhao dong', '231282000000', '肇东市', '101050502');
INSERT INTO `demo_city_code` VALUES ('231283', '2312', '2', '海伦', 'h', 'hai lun', '231283000000', '海伦市', '101050504');
INSERT INTO `demo_city_code` VALUES ('2327', '23', '1', '大兴安岭', 'd', 'da xing an ling', '232700000000', '大兴安岭地区', '101050701');
INSERT INTO `demo_city_code` VALUES ('232701', '2327', '2', '漠河', 'm', 'mo he', '232701000000', '漠河市', '101050703');
INSERT INTO `demo_city_code` VALUES ('232718', '2327', '2', '加格达奇区', 'j', 'jia ge da qi qu', '232718000000', '加格达奇区', NULL);
INSERT INTO `demo_city_code` VALUES ('232721', '2327', '2', '呼玛', 'h', 'hu ma', '232721000000', '呼玛县', '101050704');
INSERT INTO `demo_city_code` VALUES ('232722', '2327', '2', '塔河', 't', 'ta he', '232722000000', '塔河县', '101050702');
INSERT INTO `demo_city_code` VALUES ('31', '0', '0', '上海', 's', 'shang hai', '310000000000', '上海市', '101020100');
INSERT INTO `demo_city_code` VALUES ('3101', '31', '1', '上海', 's', 'shang hai', '310100000000', '上海市', '101020100');
INSERT INTO `demo_city_code` VALUES ('310101', '3101', '2', '黄浦', 'h', 'huang pu', '310101000000', '黄浦区', NULL);
INSERT INTO `demo_city_code` VALUES ('310104', '3101', '2', '徐汇', 'x', 'xu hui', '310104000000', '徐汇区', NULL);
INSERT INTO `demo_city_code` VALUES ('310105', '3101', '2', '长宁', 'c', 'chang ning', '310105000000', '长宁区', '101271106');
INSERT INTO `demo_city_code` VALUES ('310106', '3101', '2', '静安', 'j', 'jing an', '310106000000', '静安区', NULL);
INSERT INTO `demo_city_code` VALUES ('310107', '3101', '2', '普陀', 'p', 'pu tuo', '310107000000', '普陀区', '101211105');
INSERT INTO `demo_city_code` VALUES ('310109', '3101', '2', '虹口', 'h', 'hong kou', '310109000000', '虹口区', NULL);
INSERT INTO `demo_city_code` VALUES ('310110', '3101', '2', '杨浦', 'y', 'yang pu', '310110000000', '杨浦区', NULL);
INSERT INTO `demo_city_code` VALUES ('310112', '3101', '2', '闵行', 'm', 'min hang', '310112000000', '闵行区', '101020200');
INSERT INTO `demo_city_code` VALUES ('310113', '3101', '2', '宝山', 'b', 'bao shan', '310113000000', '宝山区', '101020300');
INSERT INTO `demo_city_code` VALUES ('310114', '3101', '2', '嘉定', 'j', 'jia ding', '310114000000', '嘉定区', '101020500');
INSERT INTO `demo_city_code` VALUES ('310115', '3101', '2', '浦东新区', 'p', 'pu dong xin qu', '310115000000', '浦东新区', NULL);
INSERT INTO `demo_city_code` VALUES ('310116', '3101', '2', '金山', 'j', 'jin shan', '310116000000', '金山区', '101020700');
INSERT INTO `demo_city_code` VALUES ('310117', '3101', '2', '松江', 's', 'song jiang', '310117000000', '松江区', '101020900');
INSERT INTO `demo_city_code` VALUES ('310118', '3101', '2', '青浦', 'q', 'qing pu', '310118000000', '青浦区', '101020800');
INSERT INTO `demo_city_code` VALUES ('310120', '3101', '2', '奉贤', 'f', 'feng xian', '310120000000', '奉贤区', '101021000');
INSERT INTO `demo_city_code` VALUES ('310151', '3101', '2', '崇明', 'c', 'chong ming', '310151000000', '崇明区', '101021100');
INSERT INTO `demo_city_code` VALUES ('32', '0', '0', '江苏', 'j', 'jiang su', '320000000000', '江苏省', NULL);
INSERT INTO `demo_city_code` VALUES ('3201', '32', '1', '南京', 'n', 'nan jing', '320100000000', '南京市', '101190101');
INSERT INTO `demo_city_code` VALUES ('320102', '3201', '2', '玄武', 'x', 'xuan wu', '320102000000', '玄武区', NULL);
INSERT INTO `demo_city_code` VALUES ('320104', '3201', '2', '秦淮', 'q', 'qin huai', '320104000000', '秦淮区', NULL);
INSERT INTO `demo_city_code` VALUES ('320105', '3201', '2', '建邺', 'j', 'jian ye', '320105000000', '建邺区', NULL);
INSERT INTO `demo_city_code` VALUES ('320106', '3201', '2', '鼓楼', 'g', 'gu lou', '320106000000', '鼓楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('320111', '3201', '2', '浦口', 'p', 'pu kou', '320111000000', '浦口区', '101190107');
INSERT INTO `demo_city_code` VALUES ('320113', '3201', '2', '栖霞', 'q', 'qi xia', '320113000000', '栖霞区', '101120507');
INSERT INTO `demo_city_code` VALUES ('320114', '3201', '2', '雨花台', 'y', 'yu hua tai', '320114000000', '雨花台区', NULL);
INSERT INTO `demo_city_code` VALUES ('320115', '3201', '2', '江宁', 'j', 'jiang ning', '320115000000', '江宁区', '101190104');
INSERT INTO `demo_city_code` VALUES ('320116', '3201', '2', '六合', 'l', 'lu he', '320116000000', '六合区', '101190105');
INSERT INTO `demo_city_code` VALUES ('320117', '3201', '2', '溧水', 'l', 'li shui', '320117000000', '溧水区', '101190102');
INSERT INTO `demo_city_code` VALUES ('320118', '3201', '2', '高淳', 'g', 'gao chun', '320118000000', '高淳区', '101190103');
INSERT INTO `demo_city_code` VALUES ('3202', '32', '1', '无锡', 'w', 'wu xi', '320200000000', '无锡市', '101190201');
INSERT INTO `demo_city_code` VALUES ('320205', '3202', '2', '锡山', 'x', 'xi shan', '320205000000', '锡山区', '101190204');
INSERT INTO `demo_city_code` VALUES ('320206', '3202', '2', '惠山', 'h', 'hui shan', '320206000000', '惠山区', NULL);
INSERT INTO `demo_city_code` VALUES ('320211', '3202', '2', '滨湖', 'b', 'bin hu', '320211000000', '滨湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('320213', '3202', '2', '梁溪', 'l', 'liang xi', '320213000000', '梁溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('320214', '3202', '2', '新吴', 'x', 'xin wu', '320214000000', '新吴区', NULL);
INSERT INTO `demo_city_code` VALUES ('320281', '3202', '2', '江阴', 'j', 'jiang yin', '320281000000', '江阴市', '101190202');
INSERT INTO `demo_city_code` VALUES ('320282', '3202', '2', '宜兴', 'y', 'yi xing', '320282000000', '宜兴市', '101190203');
INSERT INTO `demo_city_code` VALUES ('3203', '32', '1', '徐州', 'x', 'xu zhou', '320300000000', '徐州市', '101190801');
INSERT INTO `demo_city_code` VALUES ('320302', '3203', '2', '鼓楼', 'g', 'gu lou', '320302000000', '鼓楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('320303', '3203', '2', '云龙', 'y', 'yun long', '320303000000', '云龙区', '101290202');
INSERT INTO `demo_city_code` VALUES ('320305', '3203', '2', '贾汪', 'j', 'jia wang', '320305000000', '贾汪区', NULL);
INSERT INTO `demo_city_code` VALUES ('320311', '3203', '2', '泉山', 'q', 'quan shan', '320311000000', '泉山区', NULL);
INSERT INTO `demo_city_code` VALUES ('320312', '3203', '2', '铜山', 't', 'tong shan', '320312000000', '铜山区', '101190802');
INSERT INTO `demo_city_code` VALUES ('320321', '3203', '2', '丰县', 'f', 'feng xian', '320321000000', '丰县', '101190803');
INSERT INTO `demo_city_code` VALUES ('320322', '3203', '2', '沛县', 'p', 'pei xian', '320322000000', '沛县', '101190804');
INSERT INTO `demo_city_code` VALUES ('320324', '3203', '2', '睢宁', 's', 'sui ning', '320324000000', '睢宁县', '101190806');
INSERT INTO `demo_city_code` VALUES ('320381', '3203', '2', '新沂', 'x', 'xin yi', '320381000000', '新沂市', '101190807');
INSERT INTO `demo_city_code` VALUES ('320382', '3203', '2', '邳州', 'p', 'pi zhou', '320382000000', '邳州市', '101190805');
INSERT INTO `demo_city_code` VALUES ('3204', '32', '1', '常州', 'c', 'chang zhou', '320400000000', '常州市', '101191101');
INSERT INTO `demo_city_code` VALUES ('320402', '3204', '2', '天宁', 't', 'tian ning', '320402000000', '天宁区', NULL);
INSERT INTO `demo_city_code` VALUES ('320404', '3204', '2', '钟楼', 'z', 'zhong lou', '320404000000', '钟楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('320411', '3204', '2', '新北', 'x', 'xin bei', '320411000000', '新北区', NULL);
INSERT INTO `demo_city_code` VALUES ('320412', '3204', '2', '武进', 'w', 'wu jin', '320412000000', '武进区', '101191104');
INSERT INTO `demo_city_code` VALUES ('320413', '3204', '2', '金坛', 'j', 'jin tan', '320413000000', '金坛区', '101191103');
INSERT INTO `demo_city_code` VALUES ('320481', '3204', '2', '溧阳', 'l', 'li yang', '320481000000', '溧阳市', '101191102');
INSERT INTO `demo_city_code` VALUES ('3205', '32', '1', '苏州', 's', 'su zhou', '320500000000', '苏州市', '101190401');
INSERT INTO `demo_city_code` VALUES ('320505', '3205', '2', '虎丘', 'h', 'hu qiu', '320505000000', '虎丘区', NULL);
INSERT INTO `demo_city_code` VALUES ('320506', '3205', '2', '吴中', 'w', 'wu zhong', '320506000000', '吴中区', '101190405');
INSERT INTO `demo_city_code` VALUES ('320507', '3205', '2', '相城', 'x', 'xiang cheng', '320507000000', '相城区', NULL);
INSERT INTO `demo_city_code` VALUES ('320508', '3205', '2', '姑苏', 'g', 'gu su', '320508000000', '姑苏区', NULL);
INSERT INTO `demo_city_code` VALUES ('320509', '3205', '2', '吴江', 'w', 'wu jiang', '320509000000', '吴江区', '101190407');
INSERT INTO `demo_city_code` VALUES ('320581', '3205', '2', '常熟', 'c', 'chang shu', '320581000000', '常熟市', '101190402');
INSERT INTO `demo_city_code` VALUES ('320582', '3205', '2', '张家港', 'z', 'zhang jia gang', '320582000000', '张家港市', '101190403');
INSERT INTO `demo_city_code` VALUES ('320583', '3205', '2', '昆山', 'k', 'kun shan', '320583000000', '昆山市', '101190404');
INSERT INTO `demo_city_code` VALUES ('320585', '3205', '2', '太仓', 't', 'tai cang', '320585000000', '太仓市', '101190408');
INSERT INTO `demo_city_code` VALUES ('3206', '32', '1', '南通', 'n', 'nan tong', '320600000000', '南通市', '101190501');
INSERT INTO `demo_city_code` VALUES ('320612', '3206', '2', '通州', 't', 'tong zhou', '320612000000', '通州区', '101010600');
INSERT INTO `demo_city_code` VALUES ('320613', '3206', '2', '崇川', 'c', 'chong chuan', '320613000000', '崇川区', NULL);
INSERT INTO `demo_city_code` VALUES ('320614', '3206', '2', '海门', 'h', 'hai men', '320614000000', '海门区', '101190508');
INSERT INTO `demo_city_code` VALUES ('320623', '3206', '2', '如东', 'r', 'ru dong', '320623000000', '如东县', '101190504');
INSERT INTO `demo_city_code` VALUES ('320681', '3206', '2', '启东', 'q', 'qi dong', '320681000000', '启东市', '101190507');
INSERT INTO `demo_city_code` VALUES ('320682', '3206', '2', '如皋', 'r', 'ru gao', '320682000000', '如皋市', '101190503');
INSERT INTO `demo_city_code` VALUES ('320685', '3206', '2', '海安', 'h', 'hai an', '320685000000', '海安市', '101190502');
INSERT INTO `demo_city_code` VALUES ('3207', '32', '1', '连云港', 'l', 'lian yun gang', '320700000000', '连云港市', '101191001');
INSERT INTO `demo_city_code` VALUES ('320703', '3207', '2', '连云', 'l', 'lian yun', '320703000000', '连云区', NULL);
INSERT INTO `demo_city_code` VALUES ('320706', '3207', '2', '海州', 'h', 'hai zhou', '320706000000', '海州区', NULL);
INSERT INTO `demo_city_code` VALUES ('320707', '3207', '2', '赣榆', 'g', 'gan yu', '320707000000', '赣榆区', '101191003');
INSERT INTO `demo_city_code` VALUES ('320722', '3207', '2', '东海', 'd', 'dong hai', '320722000000', '东海县', '101191002');
INSERT INTO `demo_city_code` VALUES ('320723', '3207', '2', '灌云', 'g', 'guan yun', '320723000000', '灌云县', '101191004');
INSERT INTO `demo_city_code` VALUES ('320724', '3207', '2', '灌南', 'g', 'guan nan', '320724000000', '灌南县', '101191005');
INSERT INTO `demo_city_code` VALUES ('3208', '32', '1', '淮安', 'h', 'huai an', '320800000000', '淮安市', '101190901');
INSERT INTO `demo_city_code` VALUES ('320803', '3208', '2', '淮安区', 'h', 'huai an qu', '320803000000', '淮安区', '101190908');
INSERT INTO `demo_city_code` VALUES ('320804', '3208', '2', '淮阴', 'h', 'huai yin', '320804000000', '淮阴区', NULL);
INSERT INTO `demo_city_code` VALUES ('320812', '3208', '2', '清江浦', 'q', 'qing jiang pu', '320812000000', '清江浦区', NULL);
INSERT INTO `demo_city_code` VALUES ('320813', '3208', '2', '洪泽', 'h', 'hong ze', '320813000000', '洪泽区', '101190904');
INSERT INTO `demo_city_code` VALUES ('320826', '3208', '2', '涟水', 'l', 'lian shui', '320826000000', '涟水县', '101190905');
INSERT INTO `demo_city_code` VALUES ('320830', '3208', '2', '盱眙', 'x', 'xu yi', '320830000000', '盱眙县', '101190903');
INSERT INTO `demo_city_code` VALUES ('320831', '3208', '2', '金湖', 'j', 'jin hu', '320831000000', '金湖县', '101190902');
INSERT INTO `demo_city_code` VALUES ('3209', '32', '1', '盐城', 'y', 'yan cheng', '320900000000', '盐城市', '101190701');
INSERT INTO `demo_city_code` VALUES ('320902', '3209', '2', '亭湖', 't', 'ting hu', '320902000000', '亭湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('320903', '3209', '2', '盐都', 'y', 'yan du', '320903000000', '盐都区', '101190709');
INSERT INTO `demo_city_code` VALUES ('320904', '3209', '2', '大丰', 'd', 'da feng', '320904000000', '大丰区', '101190708');
INSERT INTO `demo_city_code` VALUES ('320921', '3209', '2', '响水', 'x', 'xiang shui', '320921000000', '响水县', '101190702');
INSERT INTO `demo_city_code` VALUES ('320922', '3209', '2', '滨海', 'b', 'bin hai', '320922000000', '滨海县', '101190703');
INSERT INTO `demo_city_code` VALUES ('320923', '3209', '2', '阜宁', 'f', 'fu ning', '320923000000', '阜宁县', '101190704');
INSERT INTO `demo_city_code` VALUES ('320924', '3209', '2', '射阳', 's', 'she yang', '320924000000', '射阳县', '101190705');
INSERT INTO `demo_city_code` VALUES ('320925', '3209', '2', '建湖', 'j', 'jian hu', '320925000000', '建湖县', '101190706');
INSERT INTO `demo_city_code` VALUES ('320981', '3209', '2', '东台', 'd', 'dong tai', '320981000000', '东台市', '101190707');
INSERT INTO `demo_city_code` VALUES ('3210', '32', '1', '扬州', 'y', 'yang zhou', '321000000000', '扬州市', '101190601');
INSERT INTO `demo_city_code` VALUES ('321002', '3210', '2', '广陵', 'g', 'guang ling', '321002000000', '广陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('321003', '3210', '2', '邗江', 'h', 'han jiang', '321003000000', '邗江区', '101190606');
INSERT INTO `demo_city_code` VALUES ('321012', '3210', '2', '江都', 'j', 'jiang du', '321012000000', '江都区', '101190605');
INSERT INTO `demo_city_code` VALUES ('321023', '3210', '2', '宝应', 'b', 'bao ying', '321023000000', '宝应县', '101190602');
INSERT INTO `demo_city_code` VALUES ('321081', '3210', '2', '仪征', 'y', 'yi zheng', '321081000000', '仪征市', '101190603');
INSERT INTO `demo_city_code` VALUES ('321084', '3210', '2', '高邮', 'g', 'gao you', '321084000000', '高邮市', '101190604');
INSERT INTO `demo_city_code` VALUES ('3211', '32', '1', '镇江', 'z', 'zhen jiang', '321100000000', '镇江市', '101190301');
INSERT INTO `demo_city_code` VALUES ('321102', '3211', '2', '京口', 'j', 'jing kou', '321102000000', '京口区', NULL);
INSERT INTO `demo_city_code` VALUES ('321111', '3211', '2', '润州', 'r', 'run zhou', '321111000000', '润州区', NULL);
INSERT INTO `demo_city_code` VALUES ('321112', '3211', '2', '丹徒', 'd', 'dan tu', '321112000000', '丹徒区', '101190305');
INSERT INTO `demo_city_code` VALUES ('321181', '3211', '2', '丹阳', 'd', 'dan yang', '321181000000', '丹阳市', '101190302');
INSERT INTO `demo_city_code` VALUES ('321182', '3211', '2', '扬中', 'y', 'yang zhong', '321182000000', '扬中市', '101190303');
INSERT INTO `demo_city_code` VALUES ('321183', '3211', '2', '句容', 'j', 'ju rong', '321183000000', '句容市', '101190304');
INSERT INTO `demo_city_code` VALUES ('3212', '32', '1', '泰州', 't', 'tai zhou', '321200000000', '泰州市', '101191201');
INSERT INTO `demo_city_code` VALUES ('321202', '3212', '2', '海陵', 'h', 'hai ling', '321202000000', '海陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('321203', '3212', '2', '高港', 'g', 'gao gang', '321203000000', '高港区', NULL);
INSERT INTO `demo_city_code` VALUES ('321204', '3212', '2', '姜堰', 'j', 'jiang yan', '321204000000', '姜堰区', '101191204');
INSERT INTO `demo_city_code` VALUES ('321281', '3212', '2', '兴化', 'x', 'xing hua', '321281000000', '兴化市', '101191202');
INSERT INTO `demo_city_code` VALUES ('321282', '3212', '2', '靖江', 'j', 'jing jiang', '321282000000', '靖江市', '101191205');
INSERT INTO `demo_city_code` VALUES ('321283', '3212', '2', '泰兴', 't', 'tai xing', '321283000000', '泰兴市', '101191203');
INSERT INTO `demo_city_code` VALUES ('3213', '32', '1', '宿迁', 's', 'su qian', '321300000000', '宿迁市', '101191301');
INSERT INTO `demo_city_code` VALUES ('321302', '3213', '2', '宿城', 's', 'su cheng', '321302000000', '宿城区', NULL);
INSERT INTO `demo_city_code` VALUES ('321311', '3213', '2', '宿豫', 's', 'su yu', '321311000000', '宿豫区', '101191305');
INSERT INTO `demo_city_code` VALUES ('321322', '3213', '2', '沭阳', 's', 'shu yang', '321322000000', '沭阳县', '101191302');
INSERT INTO `demo_city_code` VALUES ('321323', '3213', '2', '泗阳', 's', 'si yang', '321323000000', '泗阳县', '101191303');
INSERT INTO `demo_city_code` VALUES ('321324', '3213', '2', '泗洪', 's', 'si hong', '321324000000', '泗洪县', '101191304');
INSERT INTO `demo_city_code` VALUES ('33', '0', '0', '浙江', 'z', 'zhe jiang', '330000000000', '浙江省', NULL);
INSERT INTO `demo_city_code` VALUES ('3301', '33', '1', '杭州', 'h', 'hang zhou', '330100000000', '杭州市', '101210101');
INSERT INTO `demo_city_code` VALUES ('330102', '3301', '2', '上城', 's', 'shang cheng', '330102000000', '上城区', NULL);
INSERT INTO `demo_city_code` VALUES ('330105', '3301', '2', '拱墅', 'g', 'gong shu', '330105000000', '拱墅区', NULL);
INSERT INTO `demo_city_code` VALUES ('330106', '3301', '2', '西湖', 'x', 'xi hu', '330106000000', '西湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('330108', '3301', '2', '滨江', 'b', 'bin jiang', '330108000000', '滨江区', NULL);
INSERT INTO `demo_city_code` VALUES ('330109', '3301', '2', '萧山', 'x', 'xiao shan', '330109000000', '萧山区', '101210102');
INSERT INTO `demo_city_code` VALUES ('330110', '3301', '2', '余杭', 'y', 'yu hang', '330110000000', '余杭区', '101210106');
INSERT INTO `demo_city_code` VALUES ('330111', '3301', '2', '富阳', 'f', 'fu yang', '330111000000', '富阳区', '101210108');
INSERT INTO `demo_city_code` VALUES ('330112', '3301', '2', '临安', 'l', 'lin an', '330112000000', '临安区', '101210107');
INSERT INTO `demo_city_code` VALUES ('330113', '3301', '2', '临平', 'l', 'lin ping', '330113000000', '临平区', NULL);
INSERT INTO `demo_city_code` VALUES ('330114', '3301', '2', '钱塘', 'q', 'qian tang', '330114000000', '钱塘区', NULL);
INSERT INTO `demo_city_code` VALUES ('330122', '3301', '2', '桐庐', 't', 'tong lu', '330122000000', '桐庐县', '101210103');
INSERT INTO `demo_city_code` VALUES ('330127', '3301', '2', '淳安', 'c', 'chun an', '330127000000', '淳安县', '101210104');
INSERT INTO `demo_city_code` VALUES ('330182', '3301', '2', '建德', 'j', 'jian de', '330182000000', '建德市', '101210105');
INSERT INTO `demo_city_code` VALUES ('3302', '33', '1', '宁波', 'n', 'ning bo', '330200000000', '宁波市', '101210401');
INSERT INTO `demo_city_code` VALUES ('330203', '3302', '2', '海曙', 'h', 'hai shu', '330203000000', '海曙区', NULL);
INSERT INTO `demo_city_code` VALUES ('330205', '3302', '2', '江北', 'j', 'jiang bei', '330205000000', '江北区', NULL);
INSERT INTO `demo_city_code` VALUES ('330206', '3302', '2', '北仑', 'b', 'bei lun', '330206000000', '北仑区', '101210410');
INSERT INTO `demo_city_code` VALUES ('330211', '3302', '2', '镇海', 'z', 'zhen hai', '330211000000', '镇海区', '101210412');
INSERT INTO `demo_city_code` VALUES ('330212', '3302', '2', '鄞州', 'y', 'yin zhou', '330212000000', '鄞州区', '101210411');
INSERT INTO `demo_city_code` VALUES ('330213', '3302', '2', '奉化', 'f', 'feng hua', '330213000000', '奉化区', '101210405');
INSERT INTO `demo_city_code` VALUES ('330225', '3302', '2', '象山', 'x', 'xiang shan', '330225000000', '象山县', '101210406');
INSERT INTO `demo_city_code` VALUES ('330226', '3302', '2', '宁海', 'n', 'ning hai', '330226000000', '宁海县', '101210408');
INSERT INTO `demo_city_code` VALUES ('330281', '3302', '2', '余姚', 'y', 'yu yao', '330281000000', '余姚市', '101210404');
INSERT INTO `demo_city_code` VALUES ('330282', '3302', '2', '慈溪', 'c', 'ci xi', '330282000000', '慈溪市', '101210403');
INSERT INTO `demo_city_code` VALUES ('3303', '33', '1', '温州', 'w', 'wen zhou', '330300000000', '温州市', '101210701');
INSERT INTO `demo_city_code` VALUES ('330302', '3303', '2', '鹿城', 'l', 'lu cheng', '330302000000', '鹿城区', NULL);
INSERT INTO `demo_city_code` VALUES ('330303', '3303', '2', '龙湾', 'l', 'long wan', '330303000000', '龙湾区', NULL);
INSERT INTO `demo_city_code` VALUES ('330304', '3303', '2', '瓯海', 'o', 'ou hai', '330304000000', '瓯海区', NULL);
INSERT INTO `demo_city_code` VALUES ('330305', '3303', '2', '洞头', 'd', 'dong tou', '330305000000', '洞头区', '101210706');
INSERT INTO `demo_city_code` VALUES ('330324', '3303', '2', '永嘉', 'y', 'yong jia', '330324000000', '永嘉县', '101210708');
INSERT INTO `demo_city_code` VALUES ('330326', '3303', '2', '平阳', 'p', 'ping yang', '330326000000', '平阳县', '101210704');
INSERT INTO `demo_city_code` VALUES ('330327', '3303', '2', '苍南', 'c', 'cang nan', '330327000000', '苍南县', '101210709');
INSERT INTO `demo_city_code` VALUES ('330328', '3303', '2', '文成', 'w', 'wen cheng', '330328000000', '文成县', '101210703');
INSERT INTO `demo_city_code` VALUES ('330329', '3303', '2', '泰顺', 't', 'tai shun', '330329000000', '泰顺县', '101210702');
INSERT INTO `demo_city_code` VALUES ('330381', '3303', '2', '瑞安', 'r', 'rui an', '330381000000', '瑞安市', '101210705');
INSERT INTO `demo_city_code` VALUES ('330382', '3303', '2', '乐清', 'y', 'yue qing', '330382000000', '乐清市', '101210707');
INSERT INTO `demo_city_code` VALUES ('330383', '3303', '2', '龙港', 'l', 'long gang', '330383000000', '龙港市', NULL);
INSERT INTO `demo_city_code` VALUES ('3304', '33', '1', '嘉兴', 'j', 'jia xing', '330400000000', '嘉兴市', '101210301');
INSERT INTO `demo_city_code` VALUES ('330402', '3304', '2', '南湖', 'n', 'nan hu', '330402000000', '南湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('330411', '3304', '2', '秀洲', 'x', 'xiu zhou', '330411000000', '秀洲区', NULL);
INSERT INTO `demo_city_code` VALUES ('330421', '3304', '2', '嘉善', 'j', 'jia shan', '330421000000', '嘉善县', '101210302');
INSERT INTO `demo_city_code` VALUES ('330424', '3304', '2', '海盐', 'h', 'hai yan', '330424000000', '海盐县', '101210306');
INSERT INTO `demo_city_code` VALUES ('330481', '3304', '2', '海宁', 'h', 'hai ning', '330481000000', '海宁市', '101210303');
INSERT INTO `demo_city_code` VALUES ('330482', '3304', '2', '平湖', 'p', 'ping hu', '330482000000', '平湖市', '101210305');
INSERT INTO `demo_city_code` VALUES ('330483', '3304', '2', '桐乡', 't', 'tong xiang', '330483000000', '桐乡市', '101210304');
INSERT INTO `demo_city_code` VALUES ('3305', '33', '1', '湖州', 'h', 'hu zhou', '330500000000', '湖州市', '101210201');
INSERT INTO `demo_city_code` VALUES ('330502', '3305', '2', '吴兴', 'w', 'wu xing', '330502000000', '吴兴区', NULL);
INSERT INTO `demo_city_code` VALUES ('330503', '3305', '2', '南浔', 'n', 'nan xun', '330503000000', '南浔区', NULL);
INSERT INTO `demo_city_code` VALUES ('330521', '3305', '2', '德清', 'd', 'de qing', '330521000000', '德清县', '101210204');
INSERT INTO `demo_city_code` VALUES ('330522', '3305', '2', '长兴', 'c', 'chang xing', '330522000000', '长兴县', '101210202');
INSERT INTO `demo_city_code` VALUES ('330523', '3305', '2', '安吉', 'a', 'an ji', '330523000000', '安吉县', '101210203');
INSERT INTO `demo_city_code` VALUES ('3306', '33', '1', '绍兴', 's', 'shao xing', '330600000000', '绍兴市', '101210501');
INSERT INTO `demo_city_code` VALUES ('330602', '3306', '2', '越城', 'y', 'yue cheng', '330602000000', '越城区', NULL);
INSERT INTO `demo_city_code` VALUES ('330603', '3306', '2', '柯桥', 'k', 'ke qiao', '330603000000', '柯桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('330604', '3306', '2', '上虞', 's', 'shang yu', '330604000000', '上虞区', '101210503');
INSERT INTO `demo_city_code` VALUES ('330624', '3306', '2', '新昌', 'x', 'xin chang', '330624000000', '新昌县', '101210504');
INSERT INTO `demo_city_code` VALUES ('330681', '3306', '2', '诸暨', 'z', 'zhu ji', '330681000000', '诸暨市', '101210502');
INSERT INTO `demo_city_code` VALUES ('330683', '3306', '2', '嵊州', 's', 'sheng zhou', '330683000000', '嵊州市', '101210505');
INSERT INTO `demo_city_code` VALUES ('3307', '33', '1', '金华', 'j', 'jin hua', '330700000000', '金华市', '101210901');
INSERT INTO `demo_city_code` VALUES ('330702', '3307', '2', '婺城', 'w', 'wu cheng', '330702000000', '婺城区', NULL);
INSERT INTO `demo_city_code` VALUES ('330703', '3307', '2', '金东', 'j', 'jin dong', '330703000000', '金东区', NULL);
INSERT INTO `demo_city_code` VALUES ('330723', '3307', '2', '武义', 'w', 'wu yi', '330723000000', '武义县', '101210906');
INSERT INTO `demo_city_code` VALUES ('330726', '3307', '2', '浦江', 'p', 'pu jiang', '330726000000', '浦江县', '101210902');
INSERT INTO `demo_city_code` VALUES ('330727', '3307', '2', '磐安', 'p', 'pan an', '330727000000', '磐安县', '101210908');
INSERT INTO `demo_city_code` VALUES ('330781', '3307', '2', '兰溪', 'l', 'lan xi', '330781000000', '兰溪市', '101210903');
INSERT INTO `demo_city_code` VALUES ('330782', '3307', '2', '义乌', 'y', 'yi wu', '330782000000', '义乌市', '101210904');
INSERT INTO `demo_city_code` VALUES ('330783', '3307', '2', '东阳', 'd', 'dong yang', '330783000000', '东阳市', '101210905');
INSERT INTO `demo_city_code` VALUES ('330784', '3307', '2', '永康', 'y', 'yong kang', '330784000000', '永康市', '101210907');
INSERT INTO `demo_city_code` VALUES ('3308', '33', '1', '衢州', 'q', 'qu zhou', '330800000000', '衢州市', '101211001');
INSERT INTO `demo_city_code` VALUES ('330802', '3308', '2', '柯城', 'k', 'ke cheng', '330802000000', '柯城区', NULL);
INSERT INTO `demo_city_code` VALUES ('330803', '3308', '2', '衢江', 'q', 'qu jiang', '330803000000', '衢江区', '101211006');
INSERT INTO `demo_city_code` VALUES ('330822', '3308', '2', '常山', 'c', 'chang shan', '330822000000', '常山县', '101211002');
INSERT INTO `demo_city_code` VALUES ('330824', '3308', '2', '开化', 'k', 'kai hua', '330824000000', '开化县', '101211003');
INSERT INTO `demo_city_code` VALUES ('330825', '3308', '2', '龙游', 'l', 'long you', '330825000000', '龙游县', '101211004');
INSERT INTO `demo_city_code` VALUES ('330881', '3308', '2', '江山', 'j', 'jiang shan', '330881000000', '江山市', '101211005');
INSERT INTO `demo_city_code` VALUES ('3309', '33', '1', '舟山', 'z', 'zhou shan', '330900000000', '舟山市', '101211101');
INSERT INTO `demo_city_code` VALUES ('330902', '3309', '2', '定海', 'd', 'ding hai', '330902000000', '定海区', '101211106');
INSERT INTO `demo_city_code` VALUES ('330903', '3309', '2', '普陀', 'p', 'pu tuo', '330903000000', '普陀区', '101211105');
INSERT INTO `demo_city_code` VALUES ('330921', '3309', '2', '岱山', 'd', 'dai shan', '330921000000', '岱山县', '101211104');
INSERT INTO `demo_city_code` VALUES ('330922', '3309', '2', '嵊泗', 's', 'sheng si', '330922000000', '嵊泗县', '101211102');
INSERT INTO `demo_city_code` VALUES ('3310', '33', '1', '台州', 't', 'tai zhou', '331000000000', '台州市', '101210601');
INSERT INTO `demo_city_code` VALUES ('331002', '3310', '2', '椒江', 'j', 'jiao jiang', '331002000000', '椒江区', '101210611');
INSERT INTO `demo_city_code` VALUES ('331003', '3310', '2', '黄岩', 'h', 'huang yan', '331003000000', '黄岩区', '101210612');
INSERT INTO `demo_city_code` VALUES ('331004', '3310', '2', '路桥', 'l', 'lu qiao', '331004000000', '路桥区', '101210613');
INSERT INTO `demo_city_code` VALUES ('331022', '3310', '2', '三门', 's', 'san men', '331022000000', '三门县', '101210604');
INSERT INTO `demo_city_code` VALUES ('331023', '3310', '2', '天台', 't', 'tian tai', '331023000000', '天台县', '101210605');
INSERT INTO `demo_city_code` VALUES ('331024', '3310', '2', '仙居', 'x', 'xian ju', '331024000000', '仙居县', '101210606');
INSERT INTO `demo_city_code` VALUES ('331081', '3310', '2', '温岭', 'w', 'wen ling', '331081000000', '温岭市', '101210607');
INSERT INTO `demo_city_code` VALUES ('331082', '3310', '2', '临海', 'l', 'lin hai', '331082000000', '临海市', '101210610');
INSERT INTO `demo_city_code` VALUES ('331083', '3310', '2', '玉环', 'y', 'yu huan', '331083000000', '玉环市', '101210603');
INSERT INTO `demo_city_code` VALUES ('3311', '33', '1', '丽水', 'l', 'li shui', '331100000000', '丽水市', '101210801');
INSERT INTO `demo_city_code` VALUES ('331102', '3311', '2', '莲都', 'l', 'lian du', '331102000000', '莲都区', NULL);
INSERT INTO `demo_city_code` VALUES ('331121', '3311', '2', '青田', 'q', 'qing tian', '331121000000', '青田县', '101210805');
INSERT INTO `demo_city_code` VALUES ('331122', '3311', '2', '缙云', 'j', 'jin yun', '331122000000', '缙云县', '101210804');
INSERT INTO `demo_city_code` VALUES ('331123', '3311', '2', '遂昌', 's', 'sui chang', '331123000000', '遂昌县', '101210802');
INSERT INTO `demo_city_code` VALUES ('331124', '3311', '2', '松阳', 's', 'song yang', '331124000000', '松阳县', '101210808');
INSERT INTO `demo_city_code` VALUES ('331125', '3311', '2', '云和', 'y', 'yun he', '331125000000', '云和县', '101210806');
INSERT INTO `demo_city_code` VALUES ('331126', '3311', '2', '庆元', 'q', 'qing yuan', '331126000000', '庆元县', '101210807');
INSERT INTO `demo_city_code` VALUES ('331127', '3311', '2', '景宁', 'j', 'jing ning', '331127000000', '景宁畲族自治县', '101210809');
INSERT INTO `demo_city_code` VALUES ('331181', '3311', '2', '龙泉', 'l', 'long quan', '331181000000', '龙泉市', '101210803');
INSERT INTO `demo_city_code` VALUES ('34', '0', '0', '安徽', 'a', 'an hui', '340000000000', '安徽省', NULL);
INSERT INTO `demo_city_code` VALUES ('3401', '34', '1', '合肥', 'h', 'he fei', '340100000000', '合肥市', '101220101');
INSERT INTO `demo_city_code` VALUES ('340102', '3401', '2', '瑶海', 'y', 'yao hai', '340102000000', '瑶海区', NULL);
INSERT INTO `demo_city_code` VALUES ('340103', '3401', '2', '庐阳', 'l', 'lu yang', '340103000000', '庐阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('340104', '3401', '2', '蜀山', 's', 'shu shan', '340104000000', '蜀山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340111', '3401', '2', '包河', 'b', 'bao he', '340111000000', '包河区', NULL);
INSERT INTO `demo_city_code` VALUES ('340121', '3401', '2', '长丰', 'c', 'chang feng', '340121000000', '长丰县', '101220102');
INSERT INTO `demo_city_code` VALUES ('340122', '3401', '2', '肥东', 'f', 'fei dong', '340122000000', '肥东县', '101220103');
INSERT INTO `demo_city_code` VALUES ('340123', '3401', '2', '肥西', 'f', 'fei xi', '340123000000', '肥西县', '101220104');
INSERT INTO `demo_city_code` VALUES ('340124', '3401', '2', '庐江', 'l', 'lu jiang', '340124000000', '庐江县', '101221602');
INSERT INTO `demo_city_code` VALUES ('340181', '3401', '2', '巢湖', 'c', 'chao hu', '340181000000', '巢湖市', '101221601');
INSERT INTO `demo_city_code` VALUES ('3402', '34', '1', '芜湖', 'w', 'wu hu', '340200000000', '芜湖市', '101220301');
INSERT INTO `demo_city_code` VALUES ('340202', '3402', '2', '镜湖', 'j', 'jing hu', '340202000000', '镜湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('340207', '3402', '2', '鸠江', 'j', 'jiu jiang', '340207000000', '鸠江区', NULL);
INSERT INTO `demo_city_code` VALUES ('340209', '3402', '2', '弋江', 'y', 'yi jiang', '340209000000', '弋江区', NULL);
INSERT INTO `demo_city_code` VALUES ('340210', '3402', '2', '湾沚', 'w', 'wan zhi', '340210000000', '湾沚区', NULL);
INSERT INTO `demo_city_code` VALUES ('340212', '3402', '2', '繁昌', 'f', 'fan chang', '340212000000', '繁昌区', '101220302');
INSERT INTO `demo_city_code` VALUES ('340223', '3402', '2', '南陵', 'n', 'nan ling', '340223000000', '南陵县', '101220304');
INSERT INTO `demo_city_code` VALUES ('340281', '3402', '2', '无为', 'w', 'wu wei', '340281000000', '无为市', '101221603');
INSERT INTO `demo_city_code` VALUES ('3403', '34', '1', '蚌埠', 'b', 'beng bu', '340300000000', '蚌埠市', '101220201');
INSERT INTO `demo_city_code` VALUES ('340302', '3403', '2', '龙子湖', 'l', 'long zi hu', '340302000000', '龙子湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('340303', '3403', '2', '蚌山', 'b', 'beng shan', '340303000000', '蚌山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340304', '3403', '2', '禹会', 'y', 'yu hui', '340304000000', '禹会区', NULL);
INSERT INTO `demo_city_code` VALUES ('340311', '3403', '2', '淮上', 'h', 'huai shang', '340311000000', '淮上区', NULL);
INSERT INTO `demo_city_code` VALUES ('340321', '3403', '2', '怀远', 'h', 'huai yuan', '340321000000', '怀远县', '101220202');
INSERT INTO `demo_city_code` VALUES ('340322', '3403', '2', '五河', 'w', 'wu he', '340322000000', '五河县', '101220204');
INSERT INTO `demo_city_code` VALUES ('340323', '3403', '2', '固镇', 'g', 'gu zhen', '340323000000', '固镇县', '101220203');
INSERT INTO `demo_city_code` VALUES ('3404', '34', '1', '淮南', 'h', 'huai nan', '340400000000', '淮南市', '101220401');
INSERT INTO `demo_city_code` VALUES ('340402', '3404', '2', '大通', 'd', 'da tong', '340402000000', '大通区', '101150102');
INSERT INTO `demo_city_code` VALUES ('340403', '3404', '2', '田家庵', 't', 'tian jia an', '340403000000', '田家庵区', NULL);
INSERT INTO `demo_city_code` VALUES ('340404', '3404', '2', '谢家集', 'x', 'xie jia ji', '340404000000', '谢家集区', NULL);
INSERT INTO `demo_city_code` VALUES ('340405', '3404', '2', '八公山', 'b', 'ba gong shan', '340405000000', '八公山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340406', '3404', '2', '潘集', 'p', 'pan ji', '340406000000', '潘集区', '101220403');
INSERT INTO `demo_city_code` VALUES ('340421', '3404', '2', '凤台', 'f', 'feng tai', '340421000000', '凤台县', '101220402');
INSERT INTO `demo_city_code` VALUES ('340422', '3404', '2', '寿县', 's', 'shou xian', '340422000000', '寿县', '101221503');
INSERT INTO `demo_city_code` VALUES ('3405', '34', '1', '马鞍山', 'm', 'ma an shan', '340500000000', '马鞍山市', '101220501');
INSERT INTO `demo_city_code` VALUES ('340503', '3405', '2', '花山', 'h', 'hua shan', '340503000000', '花山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340504', '3405', '2', '雨山', 'y', 'yu shan', '340504000000', '雨山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340506', '3405', '2', '博望', 'b', 'bo wang', '340506000000', '博望区', NULL);
INSERT INTO `demo_city_code` VALUES ('340521', '3405', '2', '当涂', 'd', 'dang tu', '340521000000', '当涂县', '101220502');
INSERT INTO `demo_city_code` VALUES ('340522', '3405', '2', '含山', 'h', 'han shan', '340522000000', '含山县', '101221604');
INSERT INTO `demo_city_code` VALUES ('340523', '3405', '2', '和县', 'h', 'he xian', '340523000000', '和县', '101221605');
INSERT INTO `demo_city_code` VALUES ('3406', '34', '1', '淮北', 'h', 'huai bei', '340600000000', '淮北市', '101221201');
INSERT INTO `demo_city_code` VALUES ('340602', '3406', '2', '杜集', 'd', 'du ji', '340602000000', '杜集区', NULL);
INSERT INTO `demo_city_code` VALUES ('340603', '3406', '2', '相山', 'x', 'xiang shan', '340603000000', '相山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340604', '3406', '2', '烈山', 'l', 'lie shan', '340604000000', '烈山区', NULL);
INSERT INTO `demo_city_code` VALUES ('340621', '3406', '2', '濉溪', 's', 'sui xi', '340621000000', '濉溪县', '101221202');
INSERT INTO `demo_city_code` VALUES ('3407', '34', '1', '铜陵', 't', 'tong ling', '340700000000', '铜陵市', '101221301');
INSERT INTO `demo_city_code` VALUES ('340705', '3407', '2', '铜官', 't', 'tong guan', '340705000000', '铜官区', NULL);
INSERT INTO `demo_city_code` VALUES ('340706', '3407', '2', '义安', 'y', 'yi an', '340706000000', '义安区', NULL);
INSERT INTO `demo_city_code` VALUES ('340711', '3407', '2', '郊区', 'j', 'jiao qu', '340711000000', '郊区', NULL);
INSERT INTO `demo_city_code` VALUES ('340722', '3407', '2', '枞阳', 'z', 'zong yang', '340722000000', '枞阳县', '101220602');
INSERT INTO `demo_city_code` VALUES ('3408', '34', '1', '安庆', 'a', 'an qing', '340800000000', '安庆市', '101220601');
INSERT INTO `demo_city_code` VALUES ('340802', '3408', '2', '迎江', 'y', 'ying jiang', '340802000000', '迎江区', NULL);
INSERT INTO `demo_city_code` VALUES ('340803', '3408', '2', '大观', 'd', 'da guan', '340803000000', '大观区', NULL);
INSERT INTO `demo_city_code` VALUES ('340811', '3408', '2', '宜秀', 'y', 'yi xiu', '340811000000', '宜秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('340822', '3408', '2', '怀宁', 'h', 'huai ning', '340822000000', '怀宁县', '101220605');
INSERT INTO `demo_city_code` VALUES ('340825', '3408', '2', '太湖', 't', 'tai hu', '340825000000', '太湖县', '101220603');
INSERT INTO `demo_city_code` VALUES ('340826', '3408', '2', '宿松', 's', 'su song', '340826000000', '宿松县', '101220606');
INSERT INTO `demo_city_code` VALUES ('340827', '3408', '2', '望江', 'w', 'wang jiang', '340827000000', '望江县', '101220607');
INSERT INTO `demo_city_code` VALUES ('340828', '3408', '2', '岳西', 'y', 'yue xi', '340828000000', '岳西县', '101220608');
INSERT INTO `demo_city_code` VALUES ('340881', '3408', '2', '桐城', 't', 'tong cheng', '340881000000', '桐城市', '101220609');
INSERT INTO `demo_city_code` VALUES ('340882', '3408', '2', '潜山', 'q', 'qian shan', '340882000000', '潜山市', '101220604');
INSERT INTO `demo_city_code` VALUES ('3410', '34', '1', '黄山', 'h', 'huang shan', '341000000000', '黄山市', NULL);
INSERT INTO `demo_city_code` VALUES ('341002', '3410', '2', '屯溪', 't', 'tun xi', '341002000000', '屯溪区', '101221003');
INSERT INTO `demo_city_code` VALUES ('341003', '3410', '2', '黄山区', 'h', 'huang shan qu', '341003000000', '黄山区', '101221002');
INSERT INTO `demo_city_code` VALUES ('341004', '3410', '2', '徽州', 'h', 'hui zhou', '341004000000', '徽州区', NULL);
INSERT INTO `demo_city_code` VALUES ('341021', '3410', '2', '歙县', 's', 'she xian', '341021000000', '歙县', '101221006');
INSERT INTO `demo_city_code` VALUES ('341022', '3410', '2', '休宁', 'x', 'xiu ning', '341022000000', '休宁县', '101221007');
INSERT INTO `demo_city_code` VALUES ('341023', '3410', '2', '黟县', 'y', 'yi xian', '341023000000', '黟县', '101221005');
INSERT INTO `demo_city_code` VALUES ('341024', '3410', '2', '祁门', 'q', 'qi men', '341024000000', '祁门县', '101221004');
INSERT INTO `demo_city_code` VALUES ('3411', '34', '1', '滁州', 'c', 'chu zhou', '341100000000', '滁州市', '101221101');
INSERT INTO `demo_city_code` VALUES ('341102', '3411', '2', '琅琊', 'l', 'lang ya', '341102000000', '琅琊区', NULL);
INSERT INTO `demo_city_code` VALUES ('341103', '3411', '2', '南谯', 'n', 'nan qiao', '341103000000', '南谯区', NULL);
INSERT INTO `demo_city_code` VALUES ('341122', '3411', '2', '来安', 'l', 'lai an', '341122000000', '来安县', '101221106');
INSERT INTO `demo_city_code` VALUES ('341124', '3411', '2', '全椒', 'q', 'quan jiao', '341124000000', '全椒县', '101221105');
INSERT INTO `demo_city_code` VALUES ('341125', '3411', '2', '定远', 'd', 'ding yuan', '341125000000', '定远县', '101221104');
INSERT INTO `demo_city_code` VALUES ('341126', '3411', '2', '凤阳', 'f', 'feng yang', '341126000000', '凤阳县', '101221102');
INSERT INTO `demo_city_code` VALUES ('341181', '3411', '2', '天长', 't', 'tian chang', '341181000000', '天长市', '101221107');
INSERT INTO `demo_city_code` VALUES ('341182', '3411', '2', '明光', 'm', 'ming guang', '341182000000', '明光市', '101221103');
INSERT INTO `demo_city_code` VALUES ('3412', '34', '1', '阜阳', 'f', 'fu yang', '341200000000', '阜阳市', '101220801');
INSERT INTO `demo_city_code` VALUES ('341202', '3412', '2', '颍州', 'y', 'ying zhou', '341202000000', '颍州区', NULL);
INSERT INTO `demo_city_code` VALUES ('341203', '3412', '2', '颍东', 'y', 'ying dong', '341203000000', '颍东区', NULL);
INSERT INTO `demo_city_code` VALUES ('341204', '3412', '2', '颍泉', 'y', 'ying quan', '341204000000', '颍泉区', NULL);
INSERT INTO `demo_city_code` VALUES ('341221', '3412', '2', '临泉', 'l', 'lin quan', '341221000000', '临泉县', '101220804');
INSERT INTO `demo_city_code` VALUES ('341222', '3412', '2', '太和', 't', 'tai he', '341222000000', '太和县', '101220806');
INSERT INTO `demo_city_code` VALUES ('341225', '3412', '2', '阜南', 'f', 'fu nan', '341225000000', '阜南县', '101220802');
INSERT INTO `demo_city_code` VALUES ('341226', '3412', '2', '颍上', 'y', 'ying shang', '341226000000', '颍上县', '101220803');
INSERT INTO `demo_city_code` VALUES ('341282', '3412', '2', '界首', 'j', 'jie shou', '341282000000', '界首市', '101220805');
INSERT INTO `demo_city_code` VALUES ('3413', '34', '1', '宿州', 's', 'su zhou', '341300000000', '宿州市', '101220701');
INSERT INTO `demo_city_code` VALUES ('341302', '3413', '2', '埇桥', 'y', 'yong qiao', '341302000000', '埇桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('341321', '3413', '2', '砀山', 'd', 'dang shan', '341321000000', '砀山县', '101220702');
INSERT INTO `demo_city_code` VALUES ('341322', '3413', '2', '萧县', 'x', 'xiao xian', '341322000000', '萧县', '101220705');
INSERT INTO `demo_city_code` VALUES ('341323', '3413', '2', '灵璧', 'l', 'ling bi', '341323000000', '灵璧县', '101220703');
INSERT INTO `demo_city_code` VALUES ('341324', '3413', '2', '泗县', 's', 'si xian', '341324000000', '泗县', '101220704');
INSERT INTO `demo_city_code` VALUES ('3415', '34', '1', '六安', 'l', 'lu an', '341500000000', '六安市', '101221501');
INSERT INTO `demo_city_code` VALUES ('341502', '3415', '2', '金安', 'j', 'jin an', '341502000000', '金安区', NULL);
INSERT INTO `demo_city_code` VALUES ('341503', '3415', '2', '裕安', 'y', 'yu an', '341503000000', '裕安区', NULL);
INSERT INTO `demo_city_code` VALUES ('341504', '3415', '2', '叶集', 'y', 'ye ji', '341504000000', '叶集区', NULL);
INSERT INTO `demo_city_code` VALUES ('341522', '3415', '2', '霍邱', 'h', 'huo qiu', '341522000000', '霍邱县', '101221502');
INSERT INTO `demo_city_code` VALUES ('341523', '3415', '2', '舒城', 's', 'shu cheng', '341523000000', '舒城县', '101221507');
INSERT INTO `demo_city_code` VALUES ('341524', '3415', '2', '金寨', 'j', 'jin zhai', '341524000000', '金寨县', '101221505');
INSERT INTO `demo_city_code` VALUES ('341525', '3415', '2', '霍山', 'h', 'huo shan', '341525000000', '霍山县', '101221506');
INSERT INTO `demo_city_code` VALUES ('3416', '34', '1', '亳州', 'b', 'bo zhou', '341600000000', '亳州市', '101220901');
INSERT INTO `demo_city_code` VALUES ('341602', '3416', '2', '谯城', 'q', 'qiao cheng', '341602000000', '谯城区', NULL);
INSERT INTO `demo_city_code` VALUES ('341621', '3416', '2', '涡阳', 'g', 'guo yang', '341621000000', '涡阳县', '101220902');
INSERT INTO `demo_city_code` VALUES ('341622', '3416', '2', '蒙城', 'm', 'meng cheng', '341622000000', '蒙城县', '101220904');
INSERT INTO `demo_city_code` VALUES ('341623', '3416', '2', '利辛', 'l', 'li xin', '341623000000', '利辛县', '101220903');
INSERT INTO `demo_city_code` VALUES ('3417', '34', '1', '池州', 'c', 'chi zhou', '341700000000', '池州市', '101221701');
INSERT INTO `demo_city_code` VALUES ('341702', '3417', '2', '贵池', 'g', 'gui chi', '341702000000', '贵池区', NULL);
INSERT INTO `demo_city_code` VALUES ('341721', '3417', '2', '东至', 'd', 'dong zhi', '341721000000', '东至县', '101221702');
INSERT INTO `demo_city_code` VALUES ('341722', '3417', '2', '石台', 's', 'shi tai', '341722000000', '石台县', '101221705');
INSERT INTO `demo_city_code` VALUES ('341723', '3417', '2', '青阳', 'q', 'qing yang', '341723000000', '青阳县', '101221703');
INSERT INTO `demo_city_code` VALUES ('3418', '34', '1', '宣城', 'x', 'xuan cheng', '341800000000', '宣城市', '101221401');
INSERT INTO `demo_city_code` VALUES ('341802', '3418', '2', '宣州', 'x', 'xuan zhou', '341802000000', '宣州区', NULL);
INSERT INTO `demo_city_code` VALUES ('341821', '3418', '2', '郎溪', 'l', 'lang xi', '341821000000', '郎溪县', '101221407');
INSERT INTO `demo_city_code` VALUES ('341823', '3418', '2', '泾县', 'j', 'jing xian', '341823000000', '泾县', '101221402');
INSERT INTO `demo_city_code` VALUES ('341824', '3418', '2', '绩溪', 'j', 'ji xi', '341824000000', '绩溪县', '101221405');
INSERT INTO `demo_city_code` VALUES ('341825', '3418', '2', '旌德', 'j', 'jing de', '341825000000', '旌德县', '101221403');
INSERT INTO `demo_city_code` VALUES ('341881', '3418', '2', '宁国', 'n', 'ning guo', '341881000000', '宁国市', '101221404');
INSERT INTO `demo_city_code` VALUES ('341882', '3418', '2', '广德', 'g', 'guang de', '341882000000', '广德市', '101221406');
INSERT INTO `demo_city_code` VALUES ('35', '0', '0', '福建', 'f', 'fu jian', '350000000000', '福建省', NULL);
INSERT INTO `demo_city_code` VALUES ('3501', '35', '1', '福州', 'f', 'fu zhou', '350100000000', '福州市', '101230101');
INSERT INTO `demo_city_code` VALUES ('350102', '3501', '2', '鼓楼', 'g', 'gu lou', '350102000000', '鼓楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('350103', '3501', '2', '台江', 't', 'tai jiang', '350103000000', '台江区', '101260510');
INSERT INTO `demo_city_code` VALUES ('350104', '3501', '2', '仓山', 'c', 'cang shan', '350104000000', '仓山区', NULL);
INSERT INTO `demo_city_code` VALUES ('350105', '3501', '2', '马尾', 'm', 'ma wei', '350105000000', '马尾区', NULL);
INSERT INTO `demo_city_code` VALUES ('350111', '3501', '2', '晋安', 'j', 'jin an', '350111000000', '晋安区', NULL);
INSERT INTO `demo_city_code` VALUES ('350112', '3501', '2', '长乐', 'c', 'chang le', '350112000000', '长乐区', '101230110');
INSERT INTO `demo_city_code` VALUES ('350121', '3501', '2', '闽侯', 'm', 'min hou', '350121000000', '闽侯县', '101230103');
INSERT INTO `demo_city_code` VALUES ('350122', '3501', '2', '连江', 'l', 'lian jiang', '350122000000', '连江县', '101230105');
INSERT INTO `demo_city_code` VALUES ('350123', '3501', '2', '罗源', 'l', 'luo yuan', '350123000000', '罗源县', '101230104');
INSERT INTO `demo_city_code` VALUES ('350124', '3501', '2', '闽清', 'm', 'min qing', '350124000000', '闽清县', '101230102');
INSERT INTO `demo_city_code` VALUES ('350125', '3501', '2', '永泰', 'y', 'yong tai', '350125000000', '永泰县', '101230107');
INSERT INTO `demo_city_code` VALUES ('350128', '3501', '2', '平潭', 'p', 'ping tan', '350128000000', '平潭县', '101230108');
INSERT INTO `demo_city_code` VALUES ('350181', '3501', '2', '福清', 'f', 'fu qing', '350181000000', '福清市', '101230111');
INSERT INTO `demo_city_code` VALUES ('3502', '35', '1', '厦门', 'x', 'xia men', '350200000000', '厦门市', '101230201');
INSERT INTO `demo_city_code` VALUES ('350203', '3502', '2', '思明', 's', 'si ming', '350203000000', '思明区', NULL);
INSERT INTO `demo_city_code` VALUES ('350205', '3502', '2', '海沧', 'h', 'hai cang', '350205000000', '海沧区', NULL);
INSERT INTO `demo_city_code` VALUES ('350206', '3502', '2', '湖里', 'h', 'hu li', '350206000000', '湖里区', NULL);
INSERT INTO `demo_city_code` VALUES ('350211', '3502', '2', '集美', 'j', 'ji mei', '350211000000', '集美区', NULL);
INSERT INTO `demo_city_code` VALUES ('350212', '3502', '2', '同安', 't', 'tong an', '350212000000', '同安区', '101230202');
INSERT INTO `demo_city_code` VALUES ('350213', '3502', '2', '翔安', 'x', 'xiang an', '350213000000', '翔安区', NULL);
INSERT INTO `demo_city_code` VALUES ('3503', '35', '1', '莆田', 'p', 'pu tian', '350300000000', '莆田市', '101230401');
INSERT INTO `demo_city_code` VALUES ('350302', '3503', '2', '城厢', 'c', 'cheng xiang', '350302000000', '城厢区', '101230407');
INSERT INTO `demo_city_code` VALUES ('350303', '3503', '2', '涵江', 'h', 'han jiang', '350303000000', '涵江区', '101230404');
INSERT INTO `demo_city_code` VALUES ('350304', '3503', '2', '荔城', 'l', 'li cheng', '350304000000', '荔城区', '101230406');
INSERT INTO `demo_city_code` VALUES ('350305', '3503', '2', '秀屿', 'x', 'xiu yu', '350305000000', '秀屿区', '101230405');
INSERT INTO `demo_city_code` VALUES ('350322', '3503', '2', '仙游', 'x', 'xian you', '350322000000', '仙游县', '101230402');
INSERT INTO `demo_city_code` VALUES ('3504', '35', '1', '三明', 's', 'san ming', '350400000000', '三明市', '101230801');
INSERT INTO `demo_city_code` VALUES ('350404', '3504', '2', '三元', 's', 'san yuan', '350404000000', '三元区', NULL);
INSERT INTO `demo_city_code` VALUES ('350405', '3504', '2', '沙县', 's', 'sha xian', '350405000000', '沙县区', '101230808');
INSERT INTO `demo_city_code` VALUES ('350421', '3504', '2', '明溪', 'm', 'ming xi', '350421000000', '明溪县', '101230807');
INSERT INTO `demo_city_code` VALUES ('350423', '3504', '2', '清流', 'q', 'qing liu', '350423000000', '清流县', '101230803');
INSERT INTO `demo_city_code` VALUES ('350424', '3504', '2', '宁化', 'n', 'ning hua', '350424000000', '宁化县', '101230802');
INSERT INTO `demo_city_code` VALUES ('350425', '3504', '2', '大田', 'd', 'da tian', '350425000000', '大田县', '101230811');
INSERT INTO `demo_city_code` VALUES ('350426', '3504', '2', '尤溪', 'y', 'you xi', '350426000000', '尤溪县', '101230809');
INSERT INTO `demo_city_code` VALUES ('350428', '3504', '2', '将乐', 'j', 'jiang le', '350428000000', '将乐县', '101230805');
INSERT INTO `demo_city_code` VALUES ('350429', '3504', '2', '泰宁', 't', 'tai ning', '350429000000', '泰宁县', '101230804');
INSERT INTO `demo_city_code` VALUES ('350430', '3504', '2', '建宁', 'j', 'jian ning', '350430000000', '建宁县', '101230806');
INSERT INTO `demo_city_code` VALUES ('350481', '3504', '2', '永安', 'y', 'yong an', '350481000000', '永安市', '101230810');
INSERT INTO `demo_city_code` VALUES ('3505', '35', '1', '泉州', 'q', 'quan zhou', '350500000000', '泉州市', '101230501');
INSERT INTO `demo_city_code` VALUES ('350502', '3505', '2', '鲤城', 'l', 'li cheng', '350502000000', '鲤城区', NULL);
INSERT INTO `demo_city_code` VALUES ('350503', '3505', '2', '丰泽', 'f', 'feng ze', '350503000000', '丰泽区', NULL);
INSERT INTO `demo_city_code` VALUES ('350504', '3505', '2', '洛江', 'l', 'luo jiang', '350504000000', '洛江区', NULL);
INSERT INTO `demo_city_code` VALUES ('350505', '3505', '2', '泉港', 'q', 'quan gang', '350505000000', '泉港区', NULL);
INSERT INTO `demo_city_code` VALUES ('350521', '3505', '2', '惠安', 'h', 'hui an', '350521000000', '惠安县', '101230508');
INSERT INTO `demo_city_code` VALUES ('350524', '3505', '2', '安溪', 'a', 'an xi', '350524000000', '安溪县', '101230502');
INSERT INTO `demo_city_code` VALUES ('350525', '3505', '2', '永春', 'y', 'yong chun', '350525000000', '永春县', '101230504');
INSERT INTO `demo_city_code` VALUES ('350526', '3505', '2', '德化', 'd', 'de hua', '350526000000', '德化县', '101230505');
INSERT INTO `demo_city_code` VALUES ('350527', '3505', '2', '金门', 'j', 'jin men', '350527000000', '金门县', NULL);
INSERT INTO `demo_city_code` VALUES ('350581', '3505', '2', '石狮', 's', 'shi shi', '350581000000', '石狮市', '101230510');
INSERT INTO `demo_city_code` VALUES ('350582', '3505', '2', '晋江', 'j', 'jin jiang', '350582000000', '晋江市', '101230509');
INSERT INTO `demo_city_code` VALUES ('350583', '3505', '2', '南安', 'n', 'nan an', '350583000000', '南安市', '101230506');
INSERT INTO `demo_city_code` VALUES ('3506', '35', '1', '漳州', 'z', 'zhang zhou', '350600000000', '漳州市', '101230601');
INSERT INTO `demo_city_code` VALUES ('350602', '3506', '2', '芗城', 'x', 'xiang cheng', '350602000000', '芗城区', NULL);
INSERT INTO `demo_city_code` VALUES ('350603', '3506', '2', '龙文', 'l', 'long wen', '350603000000', '龙文区', NULL);
INSERT INTO `demo_city_code` VALUES ('350604', '3506', '2', '龙海', 'l', 'long hai', '350604000000', '龙海区', '101230605');
INSERT INTO `demo_city_code` VALUES ('350605', '3506', '2', '长泰', 'c', 'chang tai', '350605000000', '长泰区', '101230602');
INSERT INTO `demo_city_code` VALUES ('350622', '3506', '2', '云霄', 'y', 'yun xiao', '350622000000', '云霄县', '101230609');
INSERT INTO `demo_city_code` VALUES ('350623', '3506', '2', '漳浦', 'z', 'zhang pu', '350623000000', '漳浦县', '101230606');
INSERT INTO `demo_city_code` VALUES ('350624', '3506', '2', '诏安', 'z', 'zhao an', '350624000000', '诏安县', '101230607');
INSERT INTO `demo_city_code` VALUES ('350626', '3506', '2', '东山', 'd', 'dong shan', '350626000000', '东山县', '101230608');
INSERT INTO `demo_city_code` VALUES ('350627', '3506', '2', '南靖', 'n', 'nan jing', '350627000000', '南靖县', '101230603');
INSERT INTO `demo_city_code` VALUES ('350628', '3506', '2', '平和', 'p', 'ping he', '350628000000', '平和县', '101230604');
INSERT INTO `demo_city_code` VALUES ('350629', '3506', '2', '华安', 'h', 'hua an', '350629000000', '华安县', '101230610');
INSERT INTO `demo_city_code` VALUES ('3507', '35', '1', '南平', 'n', 'nan ping', '350700000000', '南平市', '101230901');
INSERT INTO `demo_city_code` VALUES ('350702', '3507', '2', '延平', 'y', 'yan ping', '350702000000', '延平区', NULL);
INSERT INTO `demo_city_code` VALUES ('350703', '3507', '2', '建阳', 'j', 'jian yang', '350703000000', '建阳区', '101230907');
INSERT INTO `demo_city_code` VALUES ('350721', '3507', '2', '顺昌', 's', 'shun chang', '350721000000', '顺昌县', '101230902');
INSERT INTO `demo_city_code` VALUES ('350722', '3507', '2', '浦城', 'p', 'pu cheng', '350722000000', '浦城县', '101230906');
INSERT INTO `demo_city_code` VALUES ('350723', '3507', '2', '光泽', 'g', 'guang ze', '350723000000', '光泽县', '101230903');
INSERT INTO `demo_city_code` VALUES ('350724', '3507', '2', '松溪', 's', 'song xi', '350724000000', '松溪县', '101230908');
INSERT INTO `demo_city_code` VALUES ('350725', '3507', '2', '政和', 'z', 'zheng he', '350725000000', '政和县', '101230909');
INSERT INTO `demo_city_code` VALUES ('350781', '3507', '2', '邵武', 's', 'shao wu', '350781000000', '邵武市', '101230904');
INSERT INTO `demo_city_code` VALUES ('350782', '3507', '2', '武夷山', 'w', 'wu yi shan', '350782000000', '武夷山市', '101230905');
INSERT INTO `demo_city_code` VALUES ('350783', '3507', '2', '建瓯', 'j', 'jian ou', '350783000000', '建瓯市', '101230910');
INSERT INTO `demo_city_code` VALUES ('3508', '35', '1', '龙岩', 'l', 'long yan', '350800000000', '龙岩市', '101230701');
INSERT INTO `demo_city_code` VALUES ('350802', '3508', '2', '新罗', 'x', 'xin luo', '350802000000', '新罗区', NULL);
INSERT INTO `demo_city_code` VALUES ('350803', '3508', '2', '永定', 'y', 'yong ding', '350803000000', '永定区', '101230706');
INSERT INTO `demo_city_code` VALUES ('350821', '3508', '2', '长汀', 'c', 'chang ting', '350821000000', '长汀县', '101230702');
INSERT INTO `demo_city_code` VALUES ('350823', '3508', '2', '上杭', 's', 'shang hang', '350823000000', '上杭县', '101230705');
INSERT INTO `demo_city_code` VALUES ('350824', '3508', '2', '武平', 'w', 'wu ping', '350824000000', '武平县', '101230704');
INSERT INTO `demo_city_code` VALUES ('350825', '3508', '2', '连城', 'l', 'lian cheng', '350825000000', '连城县', '101230703');
INSERT INTO `demo_city_code` VALUES ('350881', '3508', '2', '漳平', 'z', 'zhang ping', '350881000000', '漳平市', '101230707');
INSERT INTO `demo_city_code` VALUES ('3509', '35', '1', '宁德', 'n', 'ning de', '350900000000', '宁德市', '101230301');
INSERT INTO `demo_city_code` VALUES ('350902', '3509', '2', '蕉城', 'j', 'jiao cheng', '350902000000', '蕉城区', NULL);
INSERT INTO `demo_city_code` VALUES ('350921', '3509', '2', '霞浦', 'x', 'xia pu', '350921000000', '霞浦县', '101230303');
INSERT INTO `demo_city_code` VALUES ('350922', '3509', '2', '古田', 'g', 'gu tian', '350922000000', '古田县', '101230302');
INSERT INTO `demo_city_code` VALUES ('350923', '3509', '2', '屏南', 'p', 'ping nan', '350923000000', '屏南县', '101230309');
INSERT INTO `demo_city_code` VALUES ('350924', '3509', '2', '寿宁', 's', 'shou ning', '350924000000', '寿宁县', '101230304');
INSERT INTO `demo_city_code` VALUES ('350925', '3509', '2', '周宁', 'z', 'zhou ning', '350925000000', '周宁县', '101230305');
INSERT INTO `demo_city_code` VALUES ('350926', '3509', '2', '柘荣', 'z', 'zhe rong', '350926000000', '柘荣县', '101230307');
INSERT INTO `demo_city_code` VALUES ('350981', '3509', '2', '福安', 'f', 'fu an', '350981000000', '福安市', '101230306');
INSERT INTO `demo_city_code` VALUES ('350982', '3509', '2', '福鼎', 'f', 'fu ding', '350982000000', '福鼎市', '101230308');
INSERT INTO `demo_city_code` VALUES ('36', '0', '0', '江西', 'j', 'jiang xi', '360000000000', '江西省', NULL);
INSERT INTO `demo_city_code` VALUES ('3601', '36', '1', '南昌', 'n', 'nan chang', '360100000000', '南昌市', '101240101');
INSERT INTO `demo_city_code` VALUES ('360102', '3601', '2', '东湖', 'd', 'dong hu', '360102000000', '东湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('360103', '3601', '2', '西湖', 'x', 'xi hu', '360103000000', '西湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('360104', '3601', '2', '青云谱', 'q', 'qing yun pu', '360104000000', '青云谱区', NULL);
INSERT INTO `demo_city_code` VALUES ('360111', '3601', '2', '青山湖', 'q', 'qing shan hu', '360111000000', '青山湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('360112', '3601', '2', '新建', 'x', 'xin jian', '360112000000', '新建区', '101240102');
INSERT INTO `demo_city_code` VALUES ('360113', '3601', '2', '红谷滩', 'h', 'hong gu tan', '360113000000', '红谷滩区', NULL);
INSERT INTO `demo_city_code` VALUES ('360121', '3601', '2', '南昌县', 'n', 'nan chang xian', '360121000000', '南昌县', '101240103');
INSERT INTO `demo_city_code` VALUES ('360123', '3601', '2', '安义', 'a', 'an yi', '360123000000', '安义县', '101240104');
INSERT INTO `demo_city_code` VALUES ('360124', '3601', '2', '进贤', 'j', 'jin xian', '360124000000', '进贤县', '101240105');
INSERT INTO `demo_city_code` VALUES ('3602', '36', '1', '景德镇', 'j', 'jing de zhen', '360200000000', '景德镇市', '101240801');
INSERT INTO `demo_city_code` VALUES ('360202', '3602', '2', '昌江', 'c', 'chang jiang', '360202000000', '昌江区', '101310206');
INSERT INTO `demo_city_code` VALUES ('360203', '3602', '2', '珠山', 'z', 'zhu shan', '360203000000', '珠山区', NULL);
INSERT INTO `demo_city_code` VALUES ('360222', '3602', '2', '浮梁', 'f', 'fu liang', '360222000000', '浮梁县', '101240803');
INSERT INTO `demo_city_code` VALUES ('360281', '3602', '2', '乐平', 'l', 'le ping', '360281000000', '乐平市', '101240802');
INSERT INTO `demo_city_code` VALUES ('3603', '36', '1', '萍乡', 'p', 'ping xiang', '360300000000', '萍乡市', '101240901');
INSERT INTO `demo_city_code` VALUES ('360302', '3603', '2', '安源', 'a', 'an yuan', '360302000000', '安源区', '101240904');
INSERT INTO `demo_city_code` VALUES ('360313', '3603', '2', '湘东', 'x', 'xiang dong', '360313000000', '湘东区', '101240906');
INSERT INTO `demo_city_code` VALUES ('360321', '3603', '2', '莲花', 'l', 'lian hua', '360321000000', '莲花县', '101240902');
INSERT INTO `demo_city_code` VALUES ('360322', '3603', '2', '上栗', 's', 'shang li', '360322000000', '上栗县', '101240903');
INSERT INTO `demo_city_code` VALUES ('360323', '3603', '2', '芦溪', 'l', 'lu xi', '360323000000', '芦溪县', '101240905');
INSERT INTO `demo_city_code` VALUES ('3604', '36', '1', '九江', 'j', 'jiu jiang', '360400000000', '九江市', '101240201');
INSERT INTO `demo_city_code` VALUES ('360402', '3604', '2', '濂溪', 'l', 'lian xi', '360402000000', '濂溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('360403', '3604', '2', '浔阳', 'x', 'xun yang', '360403000000', '浔阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('360404', '3604', '2', '柴桑', 'c', 'chai sang', '360404000000', '柴桑区', NULL);
INSERT INTO `demo_city_code` VALUES ('360423', '3604', '2', '武宁', 'w', 'wu ning', '360423000000', '武宁县', '101240204');
INSERT INTO `demo_city_code` VALUES ('360424', '3604', '2', '修水', 'x', 'xiu shui', '360424000000', '修水县', '101240212');
INSERT INTO `demo_city_code` VALUES ('360425', '3604', '2', '永修', 'y', 'yong xiu', '360425000000', '永修县', '101240206');
INSERT INTO `demo_city_code` VALUES ('360426', '3604', '2', '德安', 'd', 'de an', '360426000000', '德安县', '101240205');
INSERT INTO `demo_city_code` VALUES ('360428', '3604', '2', '都昌', 'd', 'du chang', '360428000000', '都昌县', '101240210');
INSERT INTO `demo_city_code` VALUES ('360429', '3604', '2', '湖口', 'h', 'hu kou', '360429000000', '湖口县', '101240207');
INSERT INTO `demo_city_code` VALUES ('360430', '3604', '2', '彭泽', 'p', 'peng ze', '360430000000', '彭泽县', '101240208');
INSERT INTO `demo_city_code` VALUES ('360481', '3604', '2', '瑞昌', 'r', 'rui chang', '360481000000', '瑞昌市', '101240202');
INSERT INTO `demo_city_code` VALUES ('360482', '3604', '2', '共青城', 'g', 'gong qing cheng', '360482000000', '共青城市', NULL);
INSERT INTO `demo_city_code` VALUES ('360483', '3604', '2', '庐山', 'l', 'lu shan', '360483000000', '庐山市', '101240203');
INSERT INTO `demo_city_code` VALUES ('3605', '36', '1', '新余', 'x', 'xin yu', '360500000000', '新余市', '101241001');
INSERT INTO `demo_city_code` VALUES ('360502', '3605', '2', '渝水', 'y', 'yu shui', '360502000000', '渝水区', NULL);
INSERT INTO `demo_city_code` VALUES ('360521', '3605', '2', '分宜', 'f', 'fen yi', '360521000000', '分宜县', '101241002');
INSERT INTO `demo_city_code` VALUES ('3606', '36', '1', '鹰潭', 'y', 'ying tan', '360600000000', '鹰潭市', '101241101');
INSERT INTO `demo_city_code` VALUES ('360602', '3606', '2', '月湖', 'y', 'yue hu', '360602000000', '月湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('360603', '3606', '2', '余江', 'y', 'yu jiang', '360603000000', '余江区', '101241102');
INSERT INTO `demo_city_code` VALUES ('360681', '3606', '2', '贵溪', 'g', 'gui xi', '360681000000', '贵溪市', '101241103');
INSERT INTO `demo_city_code` VALUES ('3607', '36', '1', '赣州', 'g', 'gan zhou', '360700000000', '赣州市', '101240701');
INSERT INTO `demo_city_code` VALUES ('360702', '3607', '2', '章贡', 'z', 'zhang gong', '360702000000', '章贡区', NULL);
INSERT INTO `demo_city_code` VALUES ('360703', '3607', '2', '南康', 'n', 'nan kang', '360703000000', '南康区', '101240704');
INSERT INTO `demo_city_code` VALUES ('360704', '3607', '2', '赣县', 'g', 'gan xian', '360704000000', '赣县区', '101240718');
INSERT INTO `demo_city_code` VALUES ('360722', '3607', '2', '信丰', 'x', 'xin feng', '360722000000', '信丰县', '101240706');
INSERT INTO `demo_city_code` VALUES ('360723', '3607', '2', '大余', 'd', 'da yu', '360723000000', '大余县', '101240705');
INSERT INTO `demo_city_code` VALUES ('360724', '3607', '2', '上犹', 's', 'shang you', '360724000000', '上犹县', '101240703');
INSERT INTO `demo_city_code` VALUES ('360725', '3607', '2', '崇义', 'c', 'chong yi', '360725000000', '崇义县', '101240702');
INSERT INTO `demo_city_code` VALUES ('360726', '3607', '2', '安远', 'a', 'an yuan', '360726000000', '安远县', '101240712');
INSERT INTO `demo_city_code` VALUES ('360728', '3607', '2', '定南', 'd', 'ding nan', '360728000000', '定南县', '101240715');
INSERT INTO `demo_city_code` VALUES ('360729', '3607', '2', '全南', 'q', 'quan nan', '360729000000', '全南县', '101240713');
INSERT INTO `demo_city_code` VALUES ('360730', '3607', '2', '宁都', 'n', 'ning du', '360730000000', '宁都县', '101240707');
INSERT INTO `demo_city_code` VALUES ('360731', '3607', '2', '于都', 'y', 'yu du', '360731000000', '于都县', '101240710');
INSERT INTO `demo_city_code` VALUES ('360732', '3607', '2', '兴国', 'x', 'xing guo', '360732000000', '兴国县', '101240717');
INSERT INTO `demo_city_code` VALUES ('360733', '3607', '2', '会昌', 'h', 'hui chang', '360733000000', '会昌县', '101240711');
INSERT INTO `demo_city_code` VALUES ('360734', '3607', '2', '寻乌', 'x', 'xun wu', '360734000000', '寻乌县', '101240716');
INSERT INTO `demo_city_code` VALUES ('360735', '3607', '2', '石城', 's', 'shi cheng', '360735000000', '石城县', '101240708');
INSERT INTO `demo_city_code` VALUES ('360781', '3607', '2', '瑞金', 'r', 'rui jin', '360781000000', '瑞金市', '101240709');
INSERT INTO `demo_city_code` VALUES ('360783', '3607', '2', '龙南', 'l', 'long nan', '360783000000', '龙南市', '101240714');
INSERT INTO `demo_city_code` VALUES ('3608', '36', '1', '吉安', 'j', 'ji an', '360800000000', '吉安市', '101240601');
INSERT INTO `demo_city_code` VALUES ('360802', '3608', '2', '吉州', 'j', 'ji zhou', '360802000000', '吉州区', NULL);
INSERT INTO `demo_city_code` VALUES ('360803', '3608', '2', '青原', 'q', 'qing yuan', '360803000000', '青原区', NULL);
INSERT INTO `demo_city_code` VALUES ('360821', '3608', '2', '吉安县', 'j', 'ji an xian', '360821000000', '吉安县', '101240602');
INSERT INTO `demo_city_code` VALUES ('360822', '3608', '2', '吉水', 'j', 'ji shui', '360822000000', '吉水县', '101240603');
INSERT INTO `demo_city_code` VALUES ('360823', '3608', '2', '峡江', 'x', 'xia jiang', '360823000000', '峡江县', '101240605');
INSERT INTO `demo_city_code` VALUES ('360824', '3608', '2', '新干', 'x', 'xin gan', '360824000000', '新干县', '101240604');
INSERT INTO `demo_city_code` VALUES ('360825', '3608', '2', '永丰', 'y', 'yong feng', '360825000000', '永丰县', '101240606');
INSERT INTO `demo_city_code` VALUES ('360826', '3608', '2', '泰和', 't', 'tai he', '360826000000', '泰和县', '101240611');
INSERT INTO `demo_city_code` VALUES ('360827', '3608', '2', '遂川', 's', 'sui chuan', '360827000000', '遂川县', '101240610');
INSERT INTO `demo_city_code` VALUES ('360828', '3608', '2', '万安', 'w', 'wan an', '360828000000', '万安县', '101240609');
INSERT INTO `demo_city_code` VALUES ('360829', '3608', '2', '安福', 'a', 'an fu', '360829000000', '安福县', '101240612');
INSERT INTO `demo_city_code` VALUES ('360830', '3608', '2', '永新', 'y', 'yong xin', '360830000000', '永新县', '101240607');
INSERT INTO `demo_city_code` VALUES ('360881', '3608', '2', '井冈山', 'j', 'jing gang shan', '360881000000', '井冈山市', '101240608');
INSERT INTO `demo_city_code` VALUES ('3609', '36', '1', '宜春', 'y', 'yi chun', '360900000000', '宜春市', '101240501');
INSERT INTO `demo_city_code` VALUES ('360902', '3609', '2', '袁州', 'y', 'yuan zhou', '360902000000', '袁州区', NULL);
INSERT INTO `demo_city_code` VALUES ('360921', '3609', '2', '奉新', 'f', 'feng xin', '360921000000', '奉新县', '101240507');
INSERT INTO `demo_city_code` VALUES ('360922', '3609', '2', '万载', 'w', 'wan zai', '360922000000', '万载县', '101240504');
INSERT INTO `demo_city_code` VALUES ('360923', '3609', '2', '上高', 's', 'shang gao', '360923000000', '上高县', '101240505');
INSERT INTO `demo_city_code` VALUES ('360924', '3609', '2', '宜丰', 'y', 'yi feng', '360924000000', '宜丰县', '101240503');
INSERT INTO `demo_city_code` VALUES ('360925', '3609', '2', '靖安', 'j', 'jing an', '360925000000', '靖安县', '101240506');
INSERT INTO `demo_city_code` VALUES ('360926', '3609', '2', '铜鼓', 't', 'tong gu', '360926000000', '铜鼓县', '101240502');
INSERT INTO `demo_city_code` VALUES ('360981', '3609', '2', '丰城', 'f', 'feng cheng', '360981000000', '丰城市', '101240510');
INSERT INTO `demo_city_code` VALUES ('360982', '3609', '2', '樟树', 'z', 'zhang shu', '360982000000', '樟树市', '101240509');
INSERT INTO `demo_city_code` VALUES ('360983', '3609', '2', '高安', 'g', 'gao an', '360983000000', '高安市', '101240508');
INSERT INTO `demo_city_code` VALUES ('3610', '36', '1', '抚州', 'f', 'fu zhou', '361000000000', '抚州市', '101240401');
INSERT INTO `demo_city_code` VALUES ('361002', '3610', '2', '临川', 'l', 'lin chuan', '361002000000', '临川区', NULL);
INSERT INTO `demo_city_code` VALUES ('361003', '3610', '2', '东乡', 'd', 'dong xiang', '361003000000', '东乡区', '101161106');
INSERT INTO `demo_city_code` VALUES ('361021', '3610', '2', '南城', 'n', 'nan cheng', '361021000000', '南城县', '101240408');
INSERT INTO `demo_city_code` VALUES ('361022', '3610', '2', '黎川', 'l', 'li chuan', '361022000000', '黎川县', '101240410');
INSERT INTO `demo_city_code` VALUES ('361023', '3610', '2', '南丰', 'n', 'nan feng', '361023000000', '南丰县', '101240409');
INSERT INTO `demo_city_code` VALUES ('361024', '3610', '2', '崇仁', 'c', 'chong ren', '361024000000', '崇仁县', '101240404');
INSERT INTO `demo_city_code` VALUES ('361025', '3610', '2', '乐安', 'l', 'le an', '361025000000', '乐安县', '101240403');
INSERT INTO `demo_city_code` VALUES ('361026', '3610', '2', '宜黄', 'y', 'yi huang', '361026000000', '宜黄县', '101240407');
INSERT INTO `demo_city_code` VALUES ('361027', '3610', '2', '金溪', 'j', 'jin xi', '361027000000', '金溪县', '101240405');
INSERT INTO `demo_city_code` VALUES ('361028', '3610', '2', '资溪', 'z', 'zi xi', '361028000000', '资溪县', '101240406');
INSERT INTO `demo_city_code` VALUES ('361030', '3610', '2', '广昌', 'g', 'guang chang', '361030000000', '广昌县', '101240402');
INSERT INTO `demo_city_code` VALUES ('3611', '36', '1', '上饶', 's', 'shang rao', '361100000000', '上饶市', '101240301');
INSERT INTO `demo_city_code` VALUES ('361102', '3611', '2', '信州', 'x', 'xin zhou', '361102000000', '信州区', NULL);
INSERT INTO `demo_city_code` VALUES ('361103', '3611', '2', '广丰', 'g', 'guang feng', '361103000000', '广丰区', '101240313');
INSERT INTO `demo_city_code` VALUES ('361104', '3611', '2', '广信', 'g', 'guang xin', '361104000000', '广信区', NULL);
INSERT INTO `demo_city_code` VALUES ('361123', '3611', '2', '玉山', 'y', 'yu shan', '361123000000', '玉山县', '101240312');
INSERT INTO `demo_city_code` VALUES ('361124', '3611', '2', '铅山', 'y', 'yan shan', '361124000000', '铅山县', '101240311');
INSERT INTO `demo_city_code` VALUES ('361125', '3611', '2', '横峰', 'h', 'heng feng', '361125000000', '横峰县', '101240310');
INSERT INTO `demo_city_code` VALUES ('361126', '3611', '2', '弋阳', 'y', 'yi yang', '361126000000', '弋阳县', '101240309');
INSERT INTO `demo_city_code` VALUES ('361127', '3611', '2', '余干', 'y', 'yu gan', '361127000000', '余干县', '101240305');
INSERT INTO `demo_city_code` VALUES ('361128', '3611', '2', '鄱阳', 'p', 'po yang', '361128000000', '鄱阳县', '101240302');
INSERT INTO `demo_city_code` VALUES ('361129', '3611', '2', '万年', 'w', 'wan nian', '361129000000', '万年县', '101240306');
INSERT INTO `demo_city_code` VALUES ('361130', '3611', '2', '婺源', 'w', 'wu yuan', '361130000000', '婺源县', '101240303');
INSERT INTO `demo_city_code` VALUES ('361181', '3611', '2', '德兴', 'd', 'de xing', '361181000000', '德兴市', '101240307');
INSERT INTO `demo_city_code` VALUES ('37', '0', '0', '山东', 's', 'shan dong', '370000000000', '山东省', NULL);
INSERT INTO `demo_city_code` VALUES ('3701', '37', '1', '济南', 'j', 'ji nan', '370100000000', '济南市', '101120101');
INSERT INTO `demo_city_code` VALUES ('370102', '3701', '2', '历下', 'l', 'li xia', '370102000000', '历下区', NULL);
INSERT INTO `demo_city_code` VALUES ('370103', '3701', '2', '市中', 's', 'shi zhong', '370103000000', '市中区', NULL);
INSERT INTO `demo_city_code` VALUES ('370104', '3701', '2', '槐荫', 'h', 'huai yin', '370104000000', '槐荫区', NULL);
INSERT INTO `demo_city_code` VALUES ('370105', '3701', '2', '天桥', 't', 'tian qiao', '370105000000', '天桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('370112', '3701', '2', '历城', 'l', 'li cheng', '370112000000', '历城区', NULL);
INSERT INTO `demo_city_code` VALUES ('370113', '3701', '2', '长清', 'c', 'chang qing', '370113000000', '长清区', '101120102');
INSERT INTO `demo_city_code` VALUES ('370114', '3701', '2', '章丘', 'z', 'zhang qiu', '370114000000', '章丘区', '101120104');
INSERT INTO `demo_city_code` VALUES ('370115', '3701', '2', '济阳', 'j', 'ji yang', '370115000000', '济阳区', '101120106');
INSERT INTO `demo_city_code` VALUES ('370116', '3701', '2', '莱芜', 'l', 'lai wu', '370116000000', '莱芜区', '101121601');
INSERT INTO `demo_city_code` VALUES ('370117', '3701', '2', '钢城', 'g', 'gang cheng', '370117000000', '钢城区', NULL);
INSERT INTO `demo_city_code` VALUES ('370124', '3701', '2', '平阴', 'p', 'ping yin', '370124000000', '平阴县', '101120105');
INSERT INTO `demo_city_code` VALUES ('370126', '3701', '2', '商河', 's', 'shang he', '370126000000', '商河县', '101120103');
INSERT INTO `demo_city_code` VALUES ('3702', '37', '1', '青岛', 'q', 'qing dao', '370200000000', '青岛市', '101120201');
INSERT INTO `demo_city_code` VALUES ('370202', '3702', '2', '市南', 's', 'shi nan', '370202000000', '市南区', NULL);
INSERT INTO `demo_city_code` VALUES ('370203', '3702', '2', '市北', 's', 'shi bei', '370203000000', '市北区', NULL);
INSERT INTO `demo_city_code` VALUES ('370211', '3702', '2', '黄岛', 'h', 'huang dao', '370211000000', '黄岛区', NULL);
INSERT INTO `demo_city_code` VALUES ('370212', '3702', '2', '崂山', 'l', 'lao shan', '370212000000', '崂山区', '101120202');
INSERT INTO `demo_city_code` VALUES ('370213', '3702', '2', '李沧', 'l', 'li cang', '370213000000', '李沧区', NULL);
INSERT INTO `demo_city_code` VALUES ('370214', '3702', '2', '城阳', 'c', 'cheng yang', '370214000000', '城阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('370215', '3702', '2', '即墨', 'j', 'ji mo', '370215000000', '即墨区', '101120204');
INSERT INTO `demo_city_code` VALUES ('370281', '3702', '2', '胶州', 'j', 'jiao zhou', '370281000000', '胶州市', '101120205');
INSERT INTO `demo_city_code` VALUES ('370283', '3702', '2', '平度', 'p', 'ping du', '370283000000', '平度市', '101120208');
INSERT INTO `demo_city_code` VALUES ('370285', '3702', '2', '莱西', 'l', 'lai xi', '370285000000', '莱西市', '101120207');
INSERT INTO `demo_city_code` VALUES ('3703', '37', '1', '淄博', 'z', 'zi bo', '370300000000', '淄博市', '101120301');
INSERT INTO `demo_city_code` VALUES ('370302', '3703', '2', '淄川', 'z', 'zi chuan', '370302000000', '淄川区', '101120302');
INSERT INTO `demo_city_code` VALUES ('370303', '3703', '2', '张店', 'z', 'zhang dian', '370303000000', '张店区', NULL);
INSERT INTO `demo_city_code` VALUES ('370304', '3703', '2', '博山', 'b', 'bo shan', '370304000000', '博山区', '101120303');
INSERT INTO `demo_city_code` VALUES ('370305', '3703', '2', '临淄', 'l', 'lin zi', '370305000000', '临淄区', '101120308');
INSERT INTO `demo_city_code` VALUES ('370306', '3703', '2', '周村', 'z', 'zhou cun', '370306000000', '周村区', '101120305');
INSERT INTO `demo_city_code` VALUES ('370321', '3703', '2', '桓台', 'h', 'huan tai', '370321000000', '桓台县', '101120307');
INSERT INTO `demo_city_code` VALUES ('370322', '3703', '2', '高青', 'g', 'gao qing', '370322000000', '高青县', '101120304');
INSERT INTO `demo_city_code` VALUES ('370323', '3703', '2', '沂源', 'y', 'yi yuan', '370323000000', '沂源县', '101120306');
INSERT INTO `demo_city_code` VALUES ('3704', '37', '1', '枣庄', 'z', 'zao zhuang', '370400000000', '枣庄市', '101121401');
INSERT INTO `demo_city_code` VALUES ('370402', '3704', '2', '市中', 's', 'shi zhong', '370402000000', '市中区', NULL);
INSERT INTO `demo_city_code` VALUES ('370403', '3704', '2', '薛城', 'x', 'xue cheng', '370403000000', '薛城区', '101121402');
INSERT INTO `demo_city_code` VALUES ('370404', '3704', '2', '峄城', 'y', 'yi cheng', '370404000000', '峄城区', '101121403');
INSERT INTO `demo_city_code` VALUES ('370405', '3704', '2', '台儿庄', 't', 'tai er zhuang', '370405000000', '台儿庄区', '101121404');
INSERT INTO `demo_city_code` VALUES ('370406', '3704', '2', '山亭', 's', 'shan ting', '370406000000', '山亭区', NULL);
INSERT INTO `demo_city_code` VALUES ('370481', '3704', '2', '滕州', 't', 'teng zhou', '370481000000', '滕州市', '101121405');
INSERT INTO `demo_city_code` VALUES ('3705', '37', '1', '东营', 'd', 'dong ying', '370500000000', '东营市', '101121201');
INSERT INTO `demo_city_code` VALUES ('370502', '3705', '2', '东营区', 'd', 'dong ying qu', '370502000000', '东营区', NULL);
INSERT INTO `demo_city_code` VALUES ('370503', '3705', '2', '河口', 'h', 'he kou', '370503000000', '河口区', '101121202');
INSERT INTO `demo_city_code` VALUES ('370505', '3705', '2', '垦利', 'k', 'ken li', '370505000000', '垦利区', '101121203');
INSERT INTO `demo_city_code` VALUES ('370522', '3705', '2', '利津', 'l', 'li jin', '370522000000', '利津县', '101121204');
INSERT INTO `demo_city_code` VALUES ('370523', '3705', '2', '广饶', 'g', 'guang rao', '370523000000', '广饶县', '101121205');
INSERT INTO `demo_city_code` VALUES ('3706', '37', '1', '烟台', 'y', 'yan tai', '370600000000', '烟台市', '101120501');
INSERT INTO `demo_city_code` VALUES ('370602', '3706', '2', '芝罘', 'z', 'zhi fu', '370602000000', '芝罘区', NULL);
INSERT INTO `demo_city_code` VALUES ('370611', '3706', '2', '福山', 'f', 'fu shan', '370611000000', '福山区', '101120508');
INSERT INTO `demo_city_code` VALUES ('370612', '3706', '2', '牟平', 'm', 'mu ping', '370612000000', '牟平区', '101120509');
INSERT INTO `demo_city_code` VALUES ('370613', '3706', '2', '莱山', 'l', 'lai shan', '370613000000', '莱山区', NULL);
INSERT INTO `demo_city_code` VALUES ('370614', '3706', '2', '蓬莱', 'p', 'peng lai', '370614000000', '蓬莱区', '101120504');
INSERT INTO `demo_city_code` VALUES ('370681', '3706', '2', '龙口', 'l', 'long kou', '370681000000', '龙口市', '101120505');
INSERT INTO `demo_city_code` VALUES ('370682', '3706', '2', '莱阳', 'l', 'lai yang', '370682000000', '莱阳市', '101120510');
INSERT INTO `demo_city_code` VALUES ('370683', '3706', '2', '莱州', 'l', 'lai zhou', '370683000000', '莱州市', '101120502');
INSERT INTO `demo_city_code` VALUES ('370685', '3706', '2', '招远', 'z', 'zhao yuan', '370685000000', '招远市', '101120506');
INSERT INTO `demo_city_code` VALUES ('370686', '3706', '2', '栖霞', 'q', 'qi xia', '370686000000', '栖霞市', '101120507');
INSERT INTO `demo_city_code` VALUES ('370687', '3706', '2', '海阳', 'h', 'hai yang', '370687000000', '海阳市', '101120511');
INSERT INTO `demo_city_code` VALUES ('3707', '37', '1', '潍坊', 'w', 'wei fang', '370700000000', '潍坊市', '101120601');
INSERT INTO `demo_city_code` VALUES ('370702', '3707', '2', '潍城', 'w', 'wei cheng', '370702000000', '潍城区', NULL);
INSERT INTO `demo_city_code` VALUES ('370703', '3707', '2', '寒亭', 'h', 'han ting', '370703000000', '寒亭区', NULL);
INSERT INTO `demo_city_code` VALUES ('370704', '3707', '2', '坊子', 'f', 'fang zi', '370704000000', '坊子区', NULL);
INSERT INTO `demo_city_code` VALUES ('370705', '3707', '2', '奎文', 'k', 'kui wen', '370705000000', '奎文区', NULL);
INSERT INTO `demo_city_code` VALUES ('370724', '3707', '2', '临朐', 'l', 'lin qu', '370724000000', '临朐县', '101120604');
INSERT INTO `demo_city_code` VALUES ('370725', '3707', '2', '昌乐', 'c', 'chang le', '370725000000', '昌乐县', '101120605');
INSERT INTO `demo_city_code` VALUES ('370781', '3707', '2', '青州', 'q', 'qing zhou', '370781000000', '青州市', '101120602');
INSERT INTO `demo_city_code` VALUES ('370782', '3707', '2', '诸城', 'z', 'zhu cheng', '370782000000', '诸城市', '101120609');
INSERT INTO `demo_city_code` VALUES ('370783', '3707', '2', '寿光', 's', 'shou guang', '370783000000', '寿光市', '101120603');
INSERT INTO `demo_city_code` VALUES ('370784', '3707', '2', '安丘', 'a', 'an qiu', '370784000000', '安丘市', '101120607');
INSERT INTO `demo_city_code` VALUES ('370785', '3707', '2', '高密', 'g', 'gao mi', '370785000000', '高密市', '101120608');
INSERT INTO `demo_city_code` VALUES ('370786', '3707', '2', '昌邑', 'c', 'chang yi', '370786000000', '昌邑市', '101120606');
INSERT INTO `demo_city_code` VALUES ('3708', '37', '1', '济宁', 'j', 'ji ning', '370800000000', '济宁市', '101120701');
INSERT INTO `demo_city_code` VALUES ('370811', '3708', '2', '任城', 'r', 'ren cheng', '370811000000', '任城区', NULL);
INSERT INTO `demo_city_code` VALUES ('370812', '3708', '2', '兖州', 'y', 'yan zhou', '370812000000', '兖州区', '101120705');
INSERT INTO `demo_city_code` VALUES ('370826', '3708', '2', '微山', 'w', 'wei shan', '370826000000', '微山县', '101120703');
INSERT INTO `demo_city_code` VALUES ('370827', '3708', '2', '鱼台', 'y', 'yu tai', '370827000000', '鱼台县', '101120704');
INSERT INTO `demo_city_code` VALUES ('370828', '3708', '2', '金乡', 'j', 'jin xiang', '370828000000', '金乡县', '101120706');
INSERT INTO `demo_city_code` VALUES ('370829', '3708', '2', '嘉祥', 'j', 'jia xiang', '370829000000', '嘉祥县', '101120702');
INSERT INTO `demo_city_code` VALUES ('370830', '3708', '2', '汶上', 'w', 'wen shang', '370830000000', '汶上县', '101120707');
INSERT INTO `demo_city_code` VALUES ('370831', '3708', '2', '泗水', 's', 'si shui', '370831000000', '泗水县', '101120708');
INSERT INTO `demo_city_code` VALUES ('370832', '3708', '2', '梁山', 'l', 'liang shan', '370832000000', '梁山县', '101120709');
INSERT INTO `demo_city_code` VALUES ('370881', '3708', '2', '曲阜', 'q', 'qu fu', '370881000000', '曲阜市', '101120710');
INSERT INTO `demo_city_code` VALUES ('370883', '3708', '2', '邹城', 'z', 'zou cheng', '370883000000', '邹城市', '101120711');
INSERT INTO `demo_city_code` VALUES ('3709', '37', '1', '泰安', 't', 'tai an', '370900000000', '泰安市', '101120801');
INSERT INTO `demo_city_code` VALUES ('370902', '3709', '2', '泰山', 't', 'tai shan', '370902000000', '泰山区', NULL);
INSERT INTO `demo_city_code` VALUES ('370911', '3709', '2', '岱岳', 'd', 'dai yue', '370911000000', '岱岳区', NULL);
INSERT INTO `demo_city_code` VALUES ('370921', '3709', '2', '宁阳', 'n', 'ning yang', '370921000000', '宁阳县', '101120806');
INSERT INTO `demo_city_code` VALUES ('370923', '3709', '2', '东平', 'd', 'dong ping', '370923000000', '东平县', '101120805');
INSERT INTO `demo_city_code` VALUES ('370982', '3709', '2', '新泰', 'x', 'xin tai', '370982000000', '新泰市', '101120802');
INSERT INTO `demo_city_code` VALUES ('370983', '3709', '2', '肥城', 'f', 'fei cheng', '370983000000', '肥城市', '101120804');
INSERT INTO `demo_city_code` VALUES ('3710', '37', '1', '威海', 'w', 'wei hai', '371000000000', '威海市', '101121301');
INSERT INTO `demo_city_code` VALUES ('371002', '3710', '2', '环翠', 'h', 'huan cui', '371002000000', '环翠区', NULL);
INSERT INTO `demo_city_code` VALUES ('371003', '3710', '2', '文登', 'w', 'wen deng', '371003000000', '文登区', '101121302');
INSERT INTO `demo_city_code` VALUES ('371082', '3710', '2', '荣成', 'r', 'rong cheng', '371082000000', '荣成市', '101121303');
INSERT INTO `demo_city_code` VALUES ('371083', '3710', '2', '乳山', 'r', 'ru shan', '371083000000', '乳山市', '101121304');
INSERT INTO `demo_city_code` VALUES ('3711', '37', '1', '日照', 'r', 'ri zhao', '371100000000', '日照市', '101121501');
INSERT INTO `demo_city_code` VALUES ('371102', '3711', '2', '东港', 'd', 'dong gang', '371102000000', '东港区', '101070604');
INSERT INTO `demo_city_code` VALUES ('371103', '3711', '2', '岚山', 'l', 'lan shan', '371103000000', '岚山区', NULL);
INSERT INTO `demo_city_code` VALUES ('371121', '3711', '2', '五莲', 'w', 'wu lian', '371121000000', '五莲县', '101121502');
INSERT INTO `demo_city_code` VALUES ('371122', '3711', '2', '莒县', 'j', 'ju xian', '371122000000', '莒县', '101121503');
INSERT INTO `demo_city_code` VALUES ('3713', '37', '1', '临沂', 'l', 'lin yi', '371300000000', '临沂市', '101120901');
INSERT INTO `demo_city_code` VALUES ('371302', '3713', '2', '兰山', 'l', 'lan shan', '371302000000', '兰山区', NULL);
INSERT INTO `demo_city_code` VALUES ('371311', '3713', '2', '罗庄', 'l', 'luo zhuang', '371311000000', '罗庄区', NULL);
INSERT INTO `demo_city_code` VALUES ('371312', '3713', '2', '河东', 'h', 'he dong', '371312000000', '河东区', NULL);
INSERT INTO `demo_city_code` VALUES ('371321', '3713', '2', '沂南', 'y', 'yi nan', '371321000000', '沂南县', '101120903');
INSERT INTO `demo_city_code` VALUES ('371322', '3713', '2', '郯城', 't', 'tan cheng', '371322000000', '郯城县', '101120906');
INSERT INTO `demo_city_code` VALUES ('371323', '3713', '2', '沂水', 'y', 'yi shui', '371323000000', '沂水县', '101120910');
INSERT INTO `demo_city_code` VALUES ('371324', '3713', '2', '兰陵', 'l', 'lan ling', '371324000000', '兰陵县', NULL);
INSERT INTO `demo_city_code` VALUES ('371325', '3713', '2', '费县', 'f', 'fei xian', '371325000000', '费县', '101120909');
INSERT INTO `demo_city_code` VALUES ('371326', '3713', '2', '平邑', 'p', 'ping yi', '371326000000', '平邑县', '101120908');
INSERT INTO `demo_city_code` VALUES ('371327', '3713', '2', '莒南', 'j', 'ju nan', '371327000000', '莒南县', '101120902');
INSERT INTO `demo_city_code` VALUES ('371328', '3713', '2', '蒙阴', 'm', 'meng yin', '371328000000', '蒙阴县', '101120907');
INSERT INTO `demo_city_code` VALUES ('371329', '3713', '2', '临沭', 'l', 'lin shu', '371329000000', '临沭县', '101120905');
INSERT INTO `demo_city_code` VALUES ('3714', '37', '1', '德州', 'd', 'de zhou', '371400000000', '德州市', '101120401');
INSERT INTO `demo_city_code` VALUES ('371402', '3714', '2', '德城', 'd', 'de cheng', '371402000000', '德城区', NULL);
INSERT INTO `demo_city_code` VALUES ('371403', '3714', '2', '陵城', 'l', 'ling cheng', '371403000000', '陵城区', NULL);
INSERT INTO `demo_city_code` VALUES ('371422', '3714', '2', '宁津', 'n', 'ning jin', '371422000000', '宁津县', '101120409');
INSERT INTO `demo_city_code` VALUES ('371423', '3714', '2', '庆云', 'q', 'qing yun', '371423000000', '庆云县', '101120407');
INSERT INTO `demo_city_code` VALUES ('371424', '3714', '2', '临邑', 'l', 'lin yi', '371424000000', '临邑县', '101120403');
INSERT INTO `demo_city_code` VALUES ('371425', '3714', '2', '齐河', 'q', 'qi he', '371425000000', '齐河县', '101120405');
INSERT INTO `demo_city_code` VALUES ('371426', '3714', '2', '平原', 'p', 'ping yuan', '371426000000', '平原县', '101120408');
INSERT INTO `demo_city_code` VALUES ('371427', '3714', '2', '夏津', 'x', 'xia jin', '371427000000', '夏津县', '101120410');
INSERT INTO `demo_city_code` VALUES ('371428', '3714', '2', '武城', 'w', 'wu cheng', '371428000000', '武城县', '101120402');
INSERT INTO `demo_city_code` VALUES ('371481', '3714', '2', '乐陵', 'l', 'le ling', '371481000000', '乐陵市', '101120406');
INSERT INTO `demo_city_code` VALUES ('371482', '3714', '2', '禹城', 'y', 'yu cheng', '371482000000', '禹城市', '101120411');
INSERT INTO `demo_city_code` VALUES ('3715', '37', '1', '聊城', 'l', 'liao cheng', '371500000000', '聊城市', '101121701');
INSERT INTO `demo_city_code` VALUES ('371502', '3715', '2', '东昌府', 'd', 'dong chang fu', '371502000000', '东昌府区', NULL);
INSERT INTO `demo_city_code` VALUES ('371503', '3715', '2', '茌平', 'c', 'chi ping', '371503000000', '茌平区', '101121705');
INSERT INTO `demo_city_code` VALUES ('371521', '3715', '2', '阳谷', 'y', 'yang gu', '371521000000', '阳谷县', '101121703');
INSERT INTO `demo_city_code` VALUES ('371522', '3715', '2', '莘县', 's', 'shen xian', '371522000000', '莘县', '101121709');
INSERT INTO `demo_city_code` VALUES ('371524', '3715', '2', '东阿', 'd', 'dong e', '371524000000', '东阿县', '101121706');
INSERT INTO `demo_city_code` VALUES ('371525', '3715', '2', '冠县', 'g', 'guan xian', '371525000000', '冠县', '101121702');
INSERT INTO `demo_city_code` VALUES ('371526', '3715', '2', '高唐', 'g', 'gao tang', '371526000000', '高唐县', '101121704');
INSERT INTO `demo_city_code` VALUES ('371581', '3715', '2', '临清', 'l', 'lin qing', '371581000000', '临清市', '101121707');
INSERT INTO `demo_city_code` VALUES ('3716', '37', '1', '滨州', 'b', 'bin zhou', '371600000000', '滨州市', '101121101');
INSERT INTO `demo_city_code` VALUES ('371602', '3716', '2', '滨城', 'b', 'bin cheng', '371602000000', '滨城区', NULL);
INSERT INTO `demo_city_code` VALUES ('371603', '3716', '2', '沾化', 'z', 'zhan hua', '371603000000', '沾化区', '101121106');
INSERT INTO `demo_city_code` VALUES ('371621', '3716', '2', '惠民', 'h', 'hui min', '371621000000', '惠民县', '101121105');
INSERT INTO `demo_city_code` VALUES ('371622', '3716', '2', '阳信', 'y', 'yang xin', '371622000000', '阳信县', '101121104');
INSERT INTO `demo_city_code` VALUES ('371623', '3716', '2', '无棣', 'w', 'wu di', '371623000000', '无棣县', '101121103');
INSERT INTO `demo_city_code` VALUES ('371625', '3716', '2', '博兴', 'b', 'bo xing', '371625000000', '博兴县', '101121102');
INSERT INTO `demo_city_code` VALUES ('371681', '3716', '2', '邹平', 'z', 'zou ping', '371681000000', '邹平市', '101121107');
INSERT INTO `demo_city_code` VALUES ('3717', '37', '1', '菏泽', 'h', 'he ze', '371700000000', '菏泽市', '101121001');
INSERT INTO `demo_city_code` VALUES ('371702', '3717', '2', '牡丹', 'm', 'mu dan', '371702000000', '牡丹区', NULL);
INSERT INTO `demo_city_code` VALUES ('371703', '3717', '2', '定陶', 'd', 'ding tao', '371703000000', '定陶区', '101121005');
INSERT INTO `demo_city_code` VALUES ('371721', '3717', '2', '曹县', 'c', 'cao xian', '371721000000', '曹县', '101121007');
INSERT INTO `demo_city_code` VALUES ('371722', '3717', '2', '单县', 's', 'shan xian', '371722000000', '单县', '101121009');
INSERT INTO `demo_city_code` VALUES ('371723', '3717', '2', '成武', 'c', 'cheng wu', '371723000000', '成武县', '101121008');
INSERT INTO `demo_city_code` VALUES ('371724', '3717', '2', '巨野', 'j', 'ju ye', '371724000000', '巨野县', '101121006');
INSERT INTO `demo_city_code` VALUES ('371725', '3717', '2', '郓城', 'y', 'yun cheng', '371725000000', '郓城县', '101121003');
INSERT INTO `demo_city_code` VALUES ('371726', '3717', '2', '鄄城', 'j', 'juan cheng', '371726000000', '鄄城县', '101121002');
INSERT INTO `demo_city_code` VALUES ('371728', '3717', '2', '东明', 'd', 'dong ming', '371728000000', '东明县', '101121004');
INSERT INTO `demo_city_code` VALUES ('41', '0', '0', '河南', 'h', 'he nan', '410000000000', '河南省', '101080706');
INSERT INTO `demo_city_code` VALUES ('4101', '41', '1', '郑州', 'z', 'zheng zhou', '410100000000', '郑州市', '101180101');
INSERT INTO `demo_city_code` VALUES ('410102', '4101', '2', '中原', 'z', 'zhong yuan', '410102000000', '中原区', NULL);
INSERT INTO `demo_city_code` VALUES ('410103', '4101', '2', '二七', 'e', 'er qi', '410103000000', '二七区', NULL);
INSERT INTO `demo_city_code` VALUES ('410104', '4101', '2', '管城', 'g', 'guan cheng', '410104000000', '管城回族区', NULL);
INSERT INTO `demo_city_code` VALUES ('410105', '4101', '2', '金水', 'j', 'jin shui', '410105000000', '金水区', NULL);
INSERT INTO `demo_city_code` VALUES ('410106', '4101', '2', '上街', 's', 'shang jie', '410106000000', '上街区', '101180108');
INSERT INTO `demo_city_code` VALUES ('410108', '4101', '2', '惠济', 'h', 'hui ji', '410108000000', '惠济区', NULL);
INSERT INTO `demo_city_code` VALUES ('410122', '4101', '2', '中牟', 'z', 'zhong mu', '410122000000', '中牟县', '101180107');
INSERT INTO `demo_city_code` VALUES ('410181', '4101', '2', '巩义', 'g', 'gong yi', '410181000000', '巩义市', '101180102');
INSERT INTO `demo_city_code` VALUES ('410182', '4101', '2', '荥阳', 'x', 'xing yang', '410182000000', '荥阳市', '101180103');
INSERT INTO `demo_city_code` VALUES ('410183', '4101', '2', '新密', 'x', 'xin mi', '410183000000', '新密市', '101180105');
INSERT INTO `demo_city_code` VALUES ('410184', '4101', '2', '新郑', 'x', 'xin zheng', '410184000000', '新郑市', '101180106');
INSERT INTO `demo_city_code` VALUES ('410185', '4101', '2', '登封', 'd', 'deng feng', '410185000000', '登封市', '101180104');
INSERT INTO `demo_city_code` VALUES ('4102', '41', '1', '开封', 'k', 'kai feng', '410200000000', '开封市', '101180801');
INSERT INTO `demo_city_code` VALUES ('410202', '4102', '2', '龙亭', 'l', 'long ting', '410202000000', '龙亭区', NULL);
INSERT INTO `demo_city_code` VALUES ('410203', '4102', '2', '顺河', 's', 'shun he', '410203000000', '顺河回族区', NULL);
INSERT INTO `demo_city_code` VALUES ('410204', '4102', '2', '鼓楼', 'g', 'gu lou', '410204000000', '鼓楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('410205', '4102', '2', '禹王台', 'y', 'yu wang tai', '410205000000', '禹王台区', NULL);
INSERT INTO `demo_city_code` VALUES ('410212', '4102', '2', '祥符', 'x', 'xiang fu', '410212000000', '祥符区', NULL);
INSERT INTO `demo_city_code` VALUES ('410221', '4102', '2', '杞县', 'q', 'qi xian', '410221000000', '杞县', '101180802');
INSERT INTO `demo_city_code` VALUES ('410222', '4102', '2', '通许', 't', 'tong xu', '410222000000', '通许县', '101180804');
INSERT INTO `demo_city_code` VALUES ('410223', '4102', '2', '尉氏', 'w', 'wei shi', '410223000000', '尉氏县', '101180803');
INSERT INTO `demo_city_code` VALUES ('410225', '4102', '2', '兰考', 'l', 'lan kao', '410225000000', '兰考县', '101180805');
INSERT INTO `demo_city_code` VALUES ('4103', '41', '1', '洛阳', 'l', 'luo yang', '410300000000', '洛阳市', '101180901');
INSERT INTO `demo_city_code` VALUES ('410302', '4103', '2', '老城', 'l', 'lao cheng', '410302000000', '老城区', NULL);
INSERT INTO `demo_city_code` VALUES ('410303', '4103', '2', '西工', 'x', 'xi gong', '410303000000', '西工区', NULL);
INSERT INTO `demo_city_code` VALUES ('410304', '4103', '2', '瀍河', 'c', 'chan he', '410304000000', '瀍河回族区', NULL);
INSERT INTO `demo_city_code` VALUES ('410305', '4103', '2', '涧西', 'j', 'jian xi', '410305000000', '涧西区', NULL);
INSERT INTO `demo_city_code` VALUES ('410307', '4103', '2', '偃师', 'y', 'yan shi', '410307000000', '偃师区', '101180908');
INSERT INTO `demo_city_code` VALUES ('410308', '4103', '2', '孟津', 'm', 'meng jin', '410308000000', '孟津区', '101180903');
INSERT INTO `demo_city_code` VALUES ('410311', '4103', '2', '洛龙', 'l', 'luo long', '410311000000', '洛龙区', NULL);
INSERT INTO `demo_city_code` VALUES ('410323', '4103', '2', '新安', 'x', 'xin an', '410323000000', '新安县', '101180902');
INSERT INTO `demo_city_code` VALUES ('410324', '4103', '2', '栾川', 'l', 'luan chuan', '410324000000', '栾川县', '101180909');
INSERT INTO `demo_city_code` VALUES ('410325', '4103', '2', '嵩县', 's', 'song xian', '410325000000', '嵩县', '101180907');
INSERT INTO `demo_city_code` VALUES ('410326', '4103', '2', '汝阳', 'r', 'ru yang', '410326000000', '汝阳县', '101180910');
INSERT INTO `demo_city_code` VALUES ('410327', '4103', '2', '宜阳', 'y', 'yi yang', '410327000000', '宜阳县', '101180904');
INSERT INTO `demo_city_code` VALUES ('410328', '4103', '2', '洛宁', 'l', 'luo ning', '410328000000', '洛宁县', '101180905');
INSERT INTO `demo_city_code` VALUES ('410329', '4103', '2', '伊川', 'y', 'yi chuan', '410329000000', '伊川县', '101180906');
INSERT INTO `demo_city_code` VALUES ('4104', '41', '1', '平顶山', 'p', 'ping ding shan', '410400000000', '平顶山市', '101180501');
INSERT INTO `demo_city_code` VALUES ('410402', '4104', '2', '新华', 'x', 'xin hua', '410402000000', '新华区', NULL);
INSERT INTO `demo_city_code` VALUES ('410403', '4104', '2', '卫东', 'w', 'wei dong', '410403000000', '卫东区', NULL);
INSERT INTO `demo_city_code` VALUES ('410404', '4104', '2', '石龙', 's', 'shi long', '410404000000', '石龙区', '101180508');
INSERT INTO `demo_city_code` VALUES ('410411', '4104', '2', '湛河', 'z', 'zhan he', '410411000000', '湛河区', NULL);
INSERT INTO `demo_city_code` VALUES ('410421', '4104', '2', '宝丰', 'b', 'bao feng', '410421000000', '宝丰县', '101180503');
INSERT INTO `demo_city_code` VALUES ('410422', '4104', '2', '叶县', 'y', 'ye xian', '410422000000', '叶县', '101180505');
INSERT INTO `demo_city_code` VALUES ('410423', '4104', '2', '鲁山', 'l', 'lu shan', '410423000000', '鲁山县', '101180507');
INSERT INTO `demo_city_code` VALUES ('410425', '4104', '2', '郏县', 'j', 'jia xian', '410425000000', '郏县', '101180502');
INSERT INTO `demo_city_code` VALUES ('410481', '4104', '2', '舞钢', 'w', 'wu gang', '410481000000', '舞钢市', '101180506');
INSERT INTO `demo_city_code` VALUES ('410482', '4104', '2', '汝州', 'r', 'ru zhou', '410482000000', '汝州市', '101180504');
INSERT INTO `demo_city_code` VALUES ('4105', '41', '1', '安阳', 'a', 'an yang', '410500000000', '安阳市', '101180201');
INSERT INTO `demo_city_code` VALUES ('410502', '4105', '2', '文峰', 'w', 'wen feng', '410502000000', '文峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('410503', '4105', '2', '北关', 'b', 'bei guan', '410503000000', '北关区', NULL);
INSERT INTO `demo_city_code` VALUES ('410505', '4105', '2', '殷都', 'y', 'yin du', '410505000000', '殷都区', NULL);
INSERT INTO `demo_city_code` VALUES ('410506', '4105', '2', '龙安', 'l', 'long an', '410506000000', '龙安区', NULL);
INSERT INTO `demo_city_code` VALUES ('410522', '4105', '2', '安阳县', 'a', 'an yang xian', '410522000000', '安阳县', NULL);
INSERT INTO `demo_city_code` VALUES ('410523', '4105', '2', '汤阴', 't', 'tang yin', '410523000000', '汤阴县', '101180202');
INSERT INTO `demo_city_code` VALUES ('410526', '4105', '2', '滑县', 'h', 'hua xian', '410526000000', '滑县', '101180203');
INSERT INTO `demo_city_code` VALUES ('410527', '4105', '2', '内黄', 'n', 'nei huang', '410527000000', '内黄县', '101180204');
INSERT INTO `demo_city_code` VALUES ('410581', '4105', '2', '林州', 'l', 'lin zhou', '410581000000', '林州市', '101180205');
INSERT INTO `demo_city_code` VALUES ('4106', '41', '1', '鹤壁', 'h', 'he bi', '410600000000', '鹤壁市', '101181201');
INSERT INTO `demo_city_code` VALUES ('410602', '4106', '2', '鹤山', 'h', 'he shan', '410602000000', '鹤山区', '101281108');
INSERT INTO `demo_city_code` VALUES ('410603', '4106', '2', '山城', 's', 'shan cheng', '410603000000', '山城区', NULL);
INSERT INTO `demo_city_code` VALUES ('410611', '4106', '2', '淇滨', 'q', 'qi bin', '410611000000', '淇滨区', NULL);
INSERT INTO `demo_city_code` VALUES ('410621', '4106', '2', '浚县', 'x', 'xun xian', '410621000000', '浚县', '101181202');
INSERT INTO `demo_city_code` VALUES ('410622', '4106', '2', '淇县', 'q', 'qi xian', '410622000000', '淇县', '101181203');
INSERT INTO `demo_city_code` VALUES ('4107', '41', '1', '新乡', 'x', 'xin xiang', '410700000000', '新乡市', '101180301');
INSERT INTO `demo_city_code` VALUES ('410702', '4107', '2', '红旗', 'h', 'hong qi', '410702000000', '红旗区', NULL);
INSERT INTO `demo_city_code` VALUES ('410703', '4107', '2', '卫滨', 'w', 'wei bin', '410703000000', '卫滨区', NULL);
INSERT INTO `demo_city_code` VALUES ('410704', '4107', '2', '凤泉', 'f', 'feng quan', '410704000000', '凤泉区', NULL);
INSERT INTO `demo_city_code` VALUES ('410711', '4107', '2', '牧野', 'm', 'mu ye', '410711000000', '牧野区', NULL);
INSERT INTO `demo_city_code` VALUES ('410721', '4107', '2', '新乡县', 'x', 'xin xiang xian', '410721000000', '新乡县', NULL);
INSERT INTO `demo_city_code` VALUES ('410724', '4107', '2', '获嘉', 'h', 'huo jia', '410724000000', '获嘉县', '101180302');
INSERT INTO `demo_city_code` VALUES ('410725', '4107', '2', '原阳', 'y', 'yuan yang', '410725000000', '原阳县', '101180303');
INSERT INTO `demo_city_code` VALUES ('410726', '4107', '2', '延津', 'y', 'yan jin', '410726000000', '延津县', '101180306');
INSERT INTO `demo_city_code` VALUES ('410727', '4107', '2', '封丘', 'f', 'feng qiu', '410727000000', '封丘县', '101180307');
INSERT INTO `demo_city_code` VALUES ('410781', '4107', '2', '卫辉', 'w', 'wei hui', '410781000000', '卫辉市', '101180305');
INSERT INTO `demo_city_code` VALUES ('410782', '4107', '2', '辉县', 'h', 'hui xian', '410782000000', '辉县市', '101180304');
INSERT INTO `demo_city_code` VALUES ('410783', '4107', '2', '长垣', 'c', 'chang yuan', '410783000000', '长垣市', '101180308');
INSERT INTO `demo_city_code` VALUES ('4108', '41', '1', '焦作', 'j', 'jiao zuo', '410800000000', '焦作市', '101181101');
INSERT INTO `demo_city_code` VALUES ('410802', '4108', '2', '解放', 'j', 'jie fang', '410802000000', '解放区', NULL);
INSERT INTO `demo_city_code` VALUES ('410803', '4108', '2', '中站', 'z', 'zhong zhan', '410803000000', '中站区', NULL);
INSERT INTO `demo_city_code` VALUES ('410804', '4108', '2', '马村', 'm', 'ma cun', '410804000000', '马村区', NULL);
INSERT INTO `demo_city_code` VALUES ('410811', '4108', '2', '山阳', 's', 'shan yang', '410811000000', '山阳区', '101110608');
INSERT INTO `demo_city_code` VALUES ('410821', '4108', '2', '修武', 'x', 'xiu wu', '410821000000', '修武县', '101181102');
INSERT INTO `demo_city_code` VALUES ('410822', '4108', '2', '博爱', 'b', 'bo ai', '410822000000', '博爱县', '101181106');
INSERT INTO `demo_city_code` VALUES ('410823', '4108', '2', '武陟', 'w', 'wu zhi', '410823000000', '武陟县', '101181103');
INSERT INTO `demo_city_code` VALUES ('410825', '4108', '2', '温县', 'w', 'wen xian', '410825000000', '温县', '101181107');
INSERT INTO `demo_city_code` VALUES ('410882', '4108', '2', '沁阳', 'q', 'qin yang', '410882000000', '沁阳市', '101181104');
INSERT INTO `demo_city_code` VALUES ('410883', '4108', '2', '孟州', 'm', 'meng zhou', '410883000000', '孟州市', '101181108');
INSERT INTO `demo_city_code` VALUES ('4109', '41', '1', '濮阳', 'p', 'pu yang', '410900000000', '濮阳市', '101181301');
INSERT INTO `demo_city_code` VALUES ('410902', '4109', '2', '华龙', 'h', 'hua long', '410902000000', '华龙区', NULL);
INSERT INTO `demo_city_code` VALUES ('410922', '4109', '2', '清丰', 'q', 'qing feng', '410922000000', '清丰县', '101181304');
INSERT INTO `demo_city_code` VALUES ('410923', '4109', '2', '南乐', 'n', 'nan le', '410923000000', '南乐县', '101181303');
INSERT INTO `demo_city_code` VALUES ('410926', '4109', '2', '范县', 'f', 'fan xian', '410926000000', '范县', '101181305');
INSERT INTO `demo_city_code` VALUES ('410927', '4109', '2', '台前', 't', 'tai qian', '410927000000', '台前县', '101181302');
INSERT INTO `demo_city_code` VALUES ('410928', '4109', '2', '濮阳县', 'p', 'pu yang xian', '410928000000', '濮阳县', NULL);
INSERT INTO `demo_city_code` VALUES ('4110', '41', '1', '许昌', 'x', 'xu chang', '411000000000', '许昌市', '101180401');
INSERT INTO `demo_city_code` VALUES ('411002', '4110', '2', '魏都', 'w', 'wei du', '411002000000', '魏都区', NULL);
INSERT INTO `demo_city_code` VALUES ('411003', '4110', '2', '建安', 'j', 'jian an', '411003000000', '建安区', NULL);
INSERT INTO `demo_city_code` VALUES ('411024', '4110', '2', '鄢陵', 'y', 'yan ling', '411024000000', '鄢陵县', '101180402');
INSERT INTO `demo_city_code` VALUES ('411025', '4110', '2', '襄城', 'x', 'xiang cheng', '411025000000', '襄城县', '101180403');
INSERT INTO `demo_city_code` VALUES ('411081', '4110', '2', '禹州', 'y', 'yu zhou', '411081000000', '禹州市', '101180405');
INSERT INTO `demo_city_code` VALUES ('411082', '4110', '2', '长葛', 'c', 'chang ge', '411082000000', '长葛市', '101180404');
INSERT INTO `demo_city_code` VALUES ('4111', '41', '1', '漯河', 'l', 'luo he', '411100000000', '漯河市', '101181501');
INSERT INTO `demo_city_code` VALUES ('411102', '4111', '2', '源汇', 'y', 'yuan hui', '411102000000', '源汇区', NULL);
INSERT INTO `demo_city_code` VALUES ('411103', '4111', '2', '郾城', 'y', 'yan cheng', '411103000000', '郾城区', NULL);
INSERT INTO `demo_city_code` VALUES ('411104', '4111', '2', '召陵', 's', 'shao ling', '411104000000', '召陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('411121', '4111', '2', '舞阳', 'w', 'wu yang', '411121000000', '舞阳县', '101181503');
INSERT INTO `demo_city_code` VALUES ('411122', '4111', '2', '临颍', 'l', 'lin ying', '411122000000', '临颍县', '101181502');
INSERT INTO `demo_city_code` VALUES ('4112', '41', '1', '三门峡', 's', 'san men xia', '411200000000', '三门峡市', '101181701');
INSERT INTO `demo_city_code` VALUES ('411202', '4112', '2', '湖滨', 'h', 'hu bin', '411202000000', '湖滨区', NULL);
INSERT INTO `demo_city_code` VALUES ('411203', '4112', '2', '陕州', 's', 'shan zhou', '411203000000', '陕州区', NULL);
INSERT INTO `demo_city_code` VALUES ('411221', '4112', '2', '渑池', 'm', 'mian chi', '411221000000', '渑池县', '101181703');
INSERT INTO `demo_city_code` VALUES ('411224', '4112', '2', '卢氏', 'l', 'lu shi', '411224000000', '卢氏县', '101181704');
INSERT INTO `demo_city_code` VALUES ('411281', '4112', '2', '义马', 'y', 'yi ma', '411281000000', '义马市', '101181705');
INSERT INTO `demo_city_code` VALUES ('411282', '4112', '2', '灵宝', 'l', 'ling bao', '411282000000', '灵宝市', '101181702');
INSERT INTO `demo_city_code` VALUES ('4113', '41', '1', '南阳', 'n', 'nan yang', '411300000000', '南阳市', '101180701');
INSERT INTO `demo_city_code` VALUES ('411302', '4113', '2', '宛城', 'w', 'wan cheng', '411302000000', '宛城区', NULL);
INSERT INTO `demo_city_code` VALUES ('411303', '4113', '2', '卧龙', 'w', 'wo long', '411303000000', '卧龙区', NULL);
INSERT INTO `demo_city_code` VALUES ('411321', '4113', '2', '南召', 'n', 'nan zhao', '411321000000', '南召县', '101180702');
INSERT INTO `demo_city_code` VALUES ('411322', '4113', '2', '方城', 'f', 'fang cheng', '411322000000', '方城县', '101180703');
INSERT INTO `demo_city_code` VALUES ('411323', '4113', '2', '西峡', 'x', 'xi xia', '411323000000', '西峡县', '101180705');
INSERT INTO `demo_city_code` VALUES ('411324', '4113', '2', '镇平', 'z', 'zhen ping', '411324000000', '镇平县', '101180707');
INSERT INTO `demo_city_code` VALUES ('411325', '4113', '2', '内乡', 'n', 'nei xiang', '411325000000', '内乡县', '101180706');
INSERT INTO `demo_city_code` VALUES ('411326', '4113', '2', '淅川', 'x', 'xi chuan', '411326000000', '淅川县', '101180708');
INSERT INTO `demo_city_code` VALUES ('411327', '4113', '2', '社旗', 's', 'she qi', '411327000000', '社旗县', '101180704');
INSERT INTO `demo_city_code` VALUES ('411328', '4113', '2', '唐河', 't', 'tang he', '411328000000', '唐河县', '101180710');
INSERT INTO `demo_city_code` VALUES ('411329', '4113', '2', '新野', 'x', 'xin ye', '411329000000', '新野县', '101180709');
INSERT INTO `demo_city_code` VALUES ('411330', '4113', '2', '桐柏', 't', 'tong bai', '411330000000', '桐柏县', '101180712');
INSERT INTO `demo_city_code` VALUES ('411381', '4113', '2', '邓州', 'd', 'deng zhou', '411381000000', '邓州市', '101180711');
INSERT INTO `demo_city_code` VALUES ('4114', '41', '1', '商丘', 's', 'shang qiu', '411400000000', '商丘市', '101181001');
INSERT INTO `demo_city_code` VALUES ('411402', '4114', '2', '梁园', 'l', 'liang yuan', '411402000000', '梁园区', NULL);
INSERT INTO `demo_city_code` VALUES ('411403', '4114', '2', '睢阳', 's', 'sui yang', '411403000000', '睢阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('411421', '4114', '2', '民权', 'm', 'min quan', '411421000000', '民权县', '101181004');
INSERT INTO `demo_city_code` VALUES ('411422', '4114', '2', '睢县', 's', 'sui xian', '411422000000', '睢县', '101181003');
INSERT INTO `demo_city_code` VALUES ('411423', '4114', '2', '宁陵', 'n', 'ning ling', '411423000000', '宁陵县', '101181007');
INSERT INTO `demo_city_code` VALUES ('411424', '4114', '2', '柘城', 'z', 'zhe cheng', '411424000000', '柘城县', '101181006');
INSERT INTO `demo_city_code` VALUES ('411425', '4114', '2', '虞城', 'y', 'yu cheng', '411425000000', '虞城县', '101181005');
INSERT INTO `demo_city_code` VALUES ('411426', '4114', '2', '夏邑', 'x', 'xia yi', '411426000000', '夏邑县', '101181008');
INSERT INTO `demo_city_code` VALUES ('411481', '4114', '2', '永城', 'y', 'yong cheng', '411481000000', '永城市', '101181009');
INSERT INTO `demo_city_code` VALUES ('4115', '41', '1', '信阳', 'x', 'xin yang', '411500000000', '信阳市', '101180601');
INSERT INTO `demo_city_code` VALUES ('411502', '4115', '2', '浉河', 's', 'shi he', '411502000000', '浉河区', NULL);
INSERT INTO `demo_city_code` VALUES ('411503', '4115', '2', '平桥', 'p', 'ping qiao', '411503000000', '平桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('411521', '4115', '2', '罗山', 'l', 'luo shan', '411521000000', '罗山县', '101180603');
INSERT INTO `demo_city_code` VALUES ('411522', '4115', '2', '光山', 'g', 'guang shan', '411522000000', '光山县', '101180604');
INSERT INTO `demo_city_code` VALUES ('411523', '4115', '2', '新县', 'x', 'xin xian', '411523000000', '新县', '101180605');
INSERT INTO `demo_city_code` VALUES ('411524', '4115', '2', '商城', 's', 'shang cheng', '411524000000', '商城县', '101180609');
INSERT INTO `demo_city_code` VALUES ('411525', '4115', '2', '固始', 'g', 'gu shi', '411525000000', '固始县', '101180608');
INSERT INTO `demo_city_code` VALUES ('411526', '4115', '2', '潢川', 'h', 'huang chuan', '411526000000', '潢川县', '101180607');
INSERT INTO `demo_city_code` VALUES ('411527', '4115', '2', '淮滨', 'h', 'huai bin', '411527000000', '淮滨县', '101180606');
INSERT INTO `demo_city_code` VALUES ('411528', '4115', '2', '息县', 'x', 'xi xian', '411528000000', '息县', '101180602');
INSERT INTO `demo_city_code` VALUES ('4116', '41', '1', '周口', 'z', 'zhou kou', '411600000000', '周口市', '101181401');
INSERT INTO `demo_city_code` VALUES ('411602', '4116', '2', '川汇', 'c', 'chuan hui', '411602000000', '川汇区', NULL);
INSERT INTO `demo_city_code` VALUES ('411603', '4116', '2', '淮阳', 'h', 'huai yang', '411603000000', '淮阳区', '101181404');
INSERT INTO `demo_city_code` VALUES ('411621', '4116', '2', '扶沟', 'f', 'fu gou', '411621000000', '扶沟县', '101181402');
INSERT INTO `demo_city_code` VALUES ('411622', '4116', '2', '西华', 'x', 'xi hua', '411622000000', '西华县', '101181405');
INSERT INTO `demo_city_code` VALUES ('411623', '4116', '2', '商水', 's', 'shang shui', '411623000000', '商水县', '101181406');
INSERT INTO `demo_city_code` VALUES ('411624', '4116', '2', '沈丘', 's', 'shen qiu', '411624000000', '沈丘县', '101181410');
INSERT INTO `demo_city_code` VALUES ('411625', '4116', '2', '郸城', 'd', 'dan cheng', '411625000000', '郸城县', '101181408');
INSERT INTO `demo_city_code` VALUES ('411627', '4116', '2', '太康', 't', 'tai kang', '411627000000', '太康县', '101181403');
INSERT INTO `demo_city_code` VALUES ('411628', '4116', '2', '鹿邑', 'l', 'lu yi', '411628000000', '鹿邑县', '101181409');
INSERT INTO `demo_city_code` VALUES ('411681', '4116', '2', '项城', 'x', 'xiang cheng', '411681000000', '项城市', '101181407');
INSERT INTO `demo_city_code` VALUES ('4117', '41', '1', '驻马店', 'z', 'zhu ma dian', '411700000000', '驻马店市', '101181601');
INSERT INTO `demo_city_code` VALUES ('411702', '4117', '2', '驿城', 'y', 'yi cheng', '411702000000', '驿城区', NULL);
INSERT INTO `demo_city_code` VALUES ('411721', '4117', '2', '西平', 'x', 'xi ping', '411721000000', '西平县', '101181602');
INSERT INTO `demo_city_code` VALUES ('411722', '4117', '2', '上蔡', 's', 'shang cai', '411722000000', '上蔡县', '101181604');
INSERT INTO `demo_city_code` VALUES ('411723', '4117', '2', '平舆', 'p', 'ping yu', '411723000000', '平舆县', '101181607');
INSERT INTO `demo_city_code` VALUES ('411724', '4117', '2', '正阳', 'z', 'zheng yang', '411724000000', '正阳县', '101181610');
INSERT INTO `demo_city_code` VALUES ('411725', '4117', '2', '确山', 'q', 'que shan', '411725000000', '确山县', '101181609');
INSERT INTO `demo_city_code` VALUES ('411726', '4117', '2', '泌阳', 'b', 'bi yang', '411726000000', '泌阳县', '101181606');
INSERT INTO `demo_city_code` VALUES ('411727', '4117', '2', '汝南', 'r', 'ru nan', '411727000000', '汝南县', '101181605');
INSERT INTO `demo_city_code` VALUES ('411728', '4117', '2', '遂平', 's', 'sui ping', '411728000000', '遂平县', '101181603');
INSERT INTO `demo_city_code` VALUES ('411729', '4117', '2', '新蔡', 'x', 'xin cai', '411729000000', '新蔡县', '101181608');
INSERT INTO `demo_city_code` VALUES ('419001', '41', '1', '济源', 'j', 'ji yuan', '419001000000', '济源市', '101181801');
INSERT INTO `demo_city_code` VALUES ('419001000', '419001', '2', '济源', 'j', 'ji yuan', '419001000000', '济源市', '101181801');
INSERT INTO `demo_city_code` VALUES ('42', '0', '0', '湖北', 'h', 'hu bei', '420000000000', '湖北省', NULL);
INSERT INTO `demo_city_code` VALUES ('4201', '42', '1', '武汉', 'w', 'wu han', '420100000000', '武汉市', '101200101');
INSERT INTO `demo_city_code` VALUES ('420102', '4201', '2', '江岸', 'j', 'jiang an', '420102000000', '江岸区', NULL);
INSERT INTO `demo_city_code` VALUES ('420103', '4201', '2', '江汉', 'j', 'jiang han', '420103000000', '江汉区', NULL);
INSERT INTO `demo_city_code` VALUES ('420104', '4201', '2', '硚口', 'q', 'qiao kou', '420104000000', '硚口区', NULL);
INSERT INTO `demo_city_code` VALUES ('420105', '4201', '2', '汉阳', 'h', 'han yang', '420105000000', '汉阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('420106', '4201', '2', '武昌', 'w', 'wu chang', '420106000000', '武昌区', NULL);
INSERT INTO `demo_city_code` VALUES ('420107', '4201', '2', '青山', 'q', 'qing shan', '420107000000', '青山区', NULL);
INSERT INTO `demo_city_code` VALUES ('420111', '4201', '2', '洪山', 'h', 'hong shan', '420111000000', '洪山区', NULL);
INSERT INTO `demo_city_code` VALUES ('420112', '4201', '2', '东西湖', 'd', 'dong xi hu', '420112000000', '东西湖区', '101200106');
INSERT INTO `demo_city_code` VALUES ('420113', '4201', '2', '汉南', 'h', 'han nan', '420113000000', '汉南区', NULL);
INSERT INTO `demo_city_code` VALUES ('420114', '4201', '2', '蔡甸', 'c', 'cai dian', '420114000000', '蔡甸区', '101200102');
INSERT INTO `demo_city_code` VALUES ('420115', '4201', '2', '江夏', 'j', 'jiang xia', '420115000000', '江夏区', '101200105');
INSERT INTO `demo_city_code` VALUES ('420116', '4201', '2', '黄陂', 'h', 'huang pi', '420116000000', '黄陂区', '101200103');
INSERT INTO `demo_city_code` VALUES ('420117', '4201', '2', '新洲', 'x', 'xin zhou', '420117000000', '新洲区', '101200104');
INSERT INTO `demo_city_code` VALUES ('4202', '42', '1', '黄石', 'h', 'huang shi', '420200000000', '黄石市', '101200601');
INSERT INTO `demo_city_code` VALUES ('420202', '4202', '2', '黄石港', 'h', 'huang shi gang', '420202000000', '黄石港区', NULL);
INSERT INTO `demo_city_code` VALUES ('420203', '4202', '2', '西塞山', 'x', 'xi sai shan', '420203000000', '西塞山区', '101200606');
INSERT INTO `demo_city_code` VALUES ('420204', '4202', '2', '下陆', 'x', 'xia lu', '420204000000', '下陆区', '101200605');
INSERT INTO `demo_city_code` VALUES ('420205', '4202', '2', '铁山', 't', 'tie shan', '420205000000', '铁山区', '101200604');
INSERT INTO `demo_city_code` VALUES ('420222', '4202', '2', '阳新', 'y', 'yang xin', '420222000000', '阳新县', '101200603');
INSERT INTO `demo_city_code` VALUES ('420281', '4202', '2', '大冶', 'd', 'da ye', '420281000000', '大冶市', '101200602');
INSERT INTO `demo_city_code` VALUES ('4203', '42', '1', '十堰', 's', 'shi yan', '420300000000', '十堰市', '101201101');
INSERT INTO `demo_city_code` VALUES ('420302', '4203', '2', '茅箭', 'm', 'mao jian', '420302000000', '茅箭区', '101201108');
INSERT INTO `demo_city_code` VALUES ('420303', '4203', '2', '张湾', 'z', 'zhang wan', '420303000000', '张湾区', '101201109');
INSERT INTO `demo_city_code` VALUES ('420304', '4203', '2', '郧阳', 'y', 'yun yang', '420304000000', '郧阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('420322', '4203', '2', '郧西', 'y', 'yun xi', '420322000000', '郧西县', '101201103');
INSERT INTO `demo_city_code` VALUES ('420323', '4203', '2', '竹山', 'z', 'zhu shan', '420323000000', '竹山县', '101201105');
INSERT INTO `demo_city_code` VALUES ('420324', '4203', '2', '竹溪', 'z', 'zhu xi', '420324000000', '竹溪县', '101201102');
INSERT INTO `demo_city_code` VALUES ('420325', '4203', '2', '房县', 'f', 'fang xian', '420325000000', '房县', '101201106');
INSERT INTO `demo_city_code` VALUES ('420381', '4203', '2', '丹江口', 'd', 'dan jiang kou', '420381000000', '丹江口市', '101201107');
INSERT INTO `demo_city_code` VALUES ('4205', '42', '1', '宜昌', 'y', 'yi chang', '420500000000', '宜昌市', '101200901');
INSERT INTO `demo_city_code` VALUES ('420502', '4205', '2', '西陵', 'x', 'xi ling', '420502000000', '西陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('420503', '4205', '2', '伍家岗', 'w', 'wu jia gang', '420503000000', '伍家岗区', NULL);
INSERT INTO `demo_city_code` VALUES ('420504', '4205', '2', '点军', 'd', 'dian jun', '420504000000', '点军区', NULL);
INSERT INTO `demo_city_code` VALUES ('420505', '4205', '2', '猇亭', 'x', 'xiao ting', '420505000000', '猇亭区', NULL);
INSERT INTO `demo_city_code` VALUES ('420506', '4205', '2', '夷陵', 'y', 'yi ling', '420506000000', '夷陵区', '101200912');
INSERT INTO `demo_city_code` VALUES ('420525', '4205', '2', '远安', 'y', 'yuan an', '420525000000', '远安县', '101200902');
INSERT INTO `demo_city_code` VALUES ('420526', '4205', '2', '兴山', 'x', 'xing shan', '420526000000', '兴山县', '101200904');
INSERT INTO `demo_city_code` VALUES ('420527', '4205', '2', '秭归', 'z', 'zi gui', '420527000000', '秭归县', '101200903');
INSERT INTO `demo_city_code` VALUES ('420528', '4205', '2', '长阳', 'c', 'chang yang', '420528000000', '长阳土家族自治县', '101200908');
INSERT INTO `demo_city_code` VALUES ('420529', '4205', '2', '五峰', 'w', 'wu feng', '420529000000', '五峰土家族自治县', '101200906');
INSERT INTO `demo_city_code` VALUES ('420581', '4205', '2', '宜都', 'y', 'yi du', '420581000000', '宜都市', '101200909');
INSERT INTO `demo_city_code` VALUES ('420582', '4205', '2', '当阳', 'd', 'dang yang', '420582000000', '当阳市', '101200907');
INSERT INTO `demo_city_code` VALUES ('420583', '4205', '2', '枝江', 'z', 'zhi jiang', '420583000000', '枝江市', '101200910');
INSERT INTO `demo_city_code` VALUES ('4206', '42', '1', '襄阳', 'x', 'xiang yang', '420600000000', '襄阳市', '101200201');
INSERT INTO `demo_city_code` VALUES ('420602', '4206', '2', '襄城', 'x', 'xiang cheng', '420602000000', '襄城区', '101180403');
INSERT INTO `demo_city_code` VALUES ('420606', '4206', '2', '樊城', 'f', 'fan cheng', '420606000000', '樊城区', NULL);
INSERT INTO `demo_city_code` VALUES ('420607', '4206', '2', '襄州', 'x', 'xiang zhou', '420607000000', '襄州区', '101200202');
INSERT INTO `demo_city_code` VALUES ('420624', '4206', '2', '南漳', 'n', 'nan zhang', '420624000000', '南漳县', '101200204');
INSERT INTO `demo_city_code` VALUES ('420625', '4206', '2', '谷城', 'g', 'gu cheng', '420625000000', '谷城县', '101200207');
INSERT INTO `demo_city_code` VALUES ('420626', '4206', '2', '保康', 'b', 'bao kang', '420626000000', '保康县', '101200203');
INSERT INTO `demo_city_code` VALUES ('420682', '4206', '2', '老河口', 'l', 'lao he kou', '420682000000', '老河口市', '101200206');
INSERT INTO `demo_city_code` VALUES ('420683', '4206', '2', '枣阳', 'z', 'zao yang', '420683000000', '枣阳市', '101200208');
INSERT INTO `demo_city_code` VALUES ('420684', '4206', '2', '宜城', 'y', 'yi cheng', '420684000000', '宜城市', '101200205');
INSERT INTO `demo_city_code` VALUES ('4207', '42', '1', '鄂州', 'e', 'e zhou', '420700000000', '鄂州市', '101200301');
INSERT INTO `demo_city_code` VALUES ('420702', '4207', '2', '梁子湖', 'l', 'liang zi hu', '420702000000', '梁子湖区', '101200302');
INSERT INTO `demo_city_code` VALUES ('420703', '4207', '2', '华容', 'h', 'hua rong', '420703000000', '华容区', '101251002');
INSERT INTO `demo_city_code` VALUES ('420704', '4207', '2', '鄂城', 'e', 'e cheng', '420704000000', '鄂城区', NULL);
INSERT INTO `demo_city_code` VALUES ('4208', '42', '1', '荆门', 'j', 'jing men', '420800000000', '荆门市', '101201401');
INSERT INTO `demo_city_code` VALUES ('420802', '4208', '2', '东宝', 'd', 'dong bao', '420802000000', '东宝区', NULL);
INSERT INTO `demo_city_code` VALUES ('420804', '4208', '2', '掇刀', 'd', 'duo dao', '420804000000', '掇刀区', '101201404');
INSERT INTO `demo_city_code` VALUES ('420822', '4208', '2', '沙洋', 's', 'sha yang', '420822000000', '沙洋县', '101201405');
INSERT INTO `demo_city_code` VALUES ('420881', '4208', '2', '钟祥', 'z', 'zhong xiang', '420881000000', '钟祥市', '101201402');
INSERT INTO `demo_city_code` VALUES ('420882', '4208', '2', '京山', 'j', 'jing shan', '420882000000', '京山市', '101201403');
INSERT INTO `demo_city_code` VALUES ('4209', '42', '1', '孝感', 'x', 'xiao gan', '420900000000', '孝感市', '101200401');
INSERT INTO `demo_city_code` VALUES ('420902', '4209', '2', '孝南', 'x', 'xiao nan', '420902000000', '孝南区', NULL);
INSERT INTO `demo_city_code` VALUES ('420921', '4209', '2', '孝昌', 'x', 'xiao chang', '420921000000', '孝昌县', '101200407');
INSERT INTO `demo_city_code` VALUES ('420922', '4209', '2', '大悟', 'd', 'da wu', '420922000000', '大悟县', '101200404');
INSERT INTO `demo_city_code` VALUES ('420923', '4209', '2', '云梦', 'y', 'yun meng', '420923000000', '云梦县', '101200403');
INSERT INTO `demo_city_code` VALUES ('420981', '4209', '2', '应城', 'y', 'ying cheng', '420981000000', '应城市', '101200405');
INSERT INTO `demo_city_code` VALUES ('420982', '4209', '2', '安陆', 'a', 'an lu', '420982000000', '安陆市', '101200402');
INSERT INTO `demo_city_code` VALUES ('420984', '4209', '2', '汉川', 'h', 'han chuan', '420984000000', '汉川市', '101200406');
INSERT INTO `demo_city_code` VALUES ('4210', '42', '1', '荆州', 'j', 'jing zhou', '421000000000', '荆州市', '101200801');
INSERT INTO `demo_city_code` VALUES ('421002', '4210', '2', '沙市', 's', 'sha shi', '421002000000', '沙市区', '101201406');
INSERT INTO `demo_city_code` VALUES ('421003', '4210', '2', '荆州区', 'j', 'jing zhou qu', '421003000000', '荆州区', NULL);
INSERT INTO `demo_city_code` VALUES ('421022', '4210', '2', '公安', 'g', 'gong an', '421022000000', '公安县', '101200803');
INSERT INTO `demo_city_code` VALUES ('421024', '4210', '2', '江陵', 'j', 'jiang ling', '421024000000', '江陵县', '101200802');
INSERT INTO `demo_city_code` VALUES ('421081', '4210', '2', '石首', 's', 'shi shou', '421081000000', '石首市', '101200804');
INSERT INTO `demo_city_code` VALUES ('421083', '4210', '2', '洪湖', 'h', 'hong hu', '421083000000', '洪湖市', '101200806');
INSERT INTO `demo_city_code` VALUES ('421087', '4210', '2', '松滋', 's', 'song zi', '421087000000', '松滋市', '101200807');
INSERT INTO `demo_city_code` VALUES ('421088', '4210', '2', '监利', 'j', 'jian li', '421088000000', '监利市', '101200805');
INSERT INTO `demo_city_code` VALUES ('4211', '42', '1', '黄冈', 'h', 'huang gang', '421100000000', '黄冈市', '101200501');
INSERT INTO `demo_city_code` VALUES ('421102', '4211', '2', '黄州', 'h', 'huang zhou', '421102000000', '黄州区', NULL);
INSERT INTO `demo_city_code` VALUES ('421121', '4211', '2', '团风', 't', 'tuan feng', '421121000000', '团风县', '101200510');
INSERT INTO `demo_city_code` VALUES ('421122', '4211', '2', '红安', 'h', 'hong an', '421122000000', '红安县', '101200502');
INSERT INTO `demo_city_code` VALUES ('421123', '4211', '2', '罗田', 'l', 'luo tian', '421123000000', '罗田县', '101200504');
INSERT INTO `demo_city_code` VALUES ('421124', '4211', '2', '英山', 'y', 'ying shan', '421124000000', '英山县', '101200505');
INSERT INTO `demo_city_code` VALUES ('421125', '4211', '2', '浠水', 'x', 'xi shui', '421125000000', '浠水县', '101200506');
INSERT INTO `demo_city_code` VALUES ('421126', '4211', '2', '蕲春', 'q', 'qi chun', '421126000000', '蕲春县', '101200507');
INSERT INTO `demo_city_code` VALUES ('421127', '4211', '2', '黄梅', 'h', 'huang mei', '421127000000', '黄梅县', '101200508');
INSERT INTO `demo_city_code` VALUES ('421181', '4211', '2', '麻城', 'm', 'ma cheng', '421181000000', '麻城市', '101200503');
INSERT INTO `demo_city_code` VALUES ('421182', '4211', '2', '武穴', 'w', 'wu xue', '421182000000', '武穴市', '101200509');
INSERT INTO `demo_city_code` VALUES ('4212', '42', '1', '咸宁', 'x', 'xian ning', '421200000000', '咸宁市', '101200701');
INSERT INTO `demo_city_code` VALUES ('421202', '4212', '2', '咸安', 'x', 'xian an', '421202000000', '咸安区', NULL);
INSERT INTO `demo_city_code` VALUES ('421221', '4212', '2', '嘉鱼', 'j', 'jia yu', '421221000000', '嘉鱼县', '101200703');
INSERT INTO `demo_city_code` VALUES ('421222', '4212', '2', '通城', 't', 'tong cheng', '421222000000', '通城县', '101200705');
INSERT INTO `demo_city_code` VALUES ('421223', '4212', '2', '崇阳', 'c', 'chong yang', '421223000000', '崇阳县', '101200704');
INSERT INTO `demo_city_code` VALUES ('421224', '4212', '2', '通山', 't', 'tong shan', '421224000000', '通山县', '101200706');
INSERT INTO `demo_city_code` VALUES ('421281', '4212', '2', '赤壁', 'c', 'chi bi', '421281000000', '赤壁市', '101200702');
INSERT INTO `demo_city_code` VALUES ('4213', '42', '1', '随州', 's', 'sui zhou', '421300000000', '随州市', '101201301');
INSERT INTO `demo_city_code` VALUES ('421303', '4213', '2', '曾都', 'z', 'zeng du', '421303000000', '曾都区', NULL);
INSERT INTO `demo_city_code` VALUES ('421321', '4213', '2', '随县', 's', 'sui xian', '421321000000', '随县', NULL);
INSERT INTO `demo_city_code` VALUES ('421381', '4213', '2', '广水', 'g', 'guang shui', '421381000000', '广水市', '101201302');
INSERT INTO `demo_city_code` VALUES ('4228', '42', '1', '恩施', 'e', 'en shi', '422800000000', '恩施土家族苗族自治州', '101201001');
INSERT INTO `demo_city_code` VALUES ('422801', '4228', '2', '恩施市', 'e', 'en shi shi', '422801000000', '恩施市', NULL);
INSERT INTO `demo_city_code` VALUES ('422802', '4228', '2', '利川', 'l', 'li chuan', '422802000000', '利川市', '101201002');
INSERT INTO `demo_city_code` VALUES ('422822', '4228', '2', '建始', 'j', 'jian shi', '422822000000', '建始县', '101201003');
INSERT INTO `demo_city_code` VALUES ('422823', '4228', '2', '巴东', 'b', 'ba dong', '422823000000', '巴东县', '101201008');
INSERT INTO `demo_city_code` VALUES ('422825', '4228', '2', '宣恩', 'x', 'xuan en', '422825000000', '宣恩县', '101201005');
INSERT INTO `demo_city_code` VALUES ('422826', '4228', '2', '咸丰', 'x', 'xian feng', '422826000000', '咸丰县', '101201004');
INSERT INTO `demo_city_code` VALUES ('422827', '4228', '2', '来凤', 'l', 'lai feng', '422827000000', '来凤县', '101201007');
INSERT INTO `demo_city_code` VALUES ('422828', '4228', '2', '鹤峰', 'h', 'he feng', '422828000000', '鹤峰县', '101201006');
INSERT INTO `demo_city_code` VALUES ('429004', '42', '1', '仙桃', 'x', 'xian tao', '429004000000', '仙桃市', '101201601');
INSERT INTO `demo_city_code` VALUES ('429004000', '429004', '2', '仙桃', 'x', 'xian tao', '429004000000', '仙桃市', '101201601');
INSERT INTO `demo_city_code` VALUES ('429005', '42', '1', '潜江', 'q', 'qian jiang', '429005000000', '潜江市', '101201701');
INSERT INTO `demo_city_code` VALUES ('429005000', '429005', '2', '潜江', 'q', 'qian jiang', '429005000000', '潜江市', '101201701');
INSERT INTO `demo_city_code` VALUES ('429006', '42', '1', '天门', 't', 'tian men', '429006000000', '天门市', '101201501');
INSERT INTO `demo_city_code` VALUES ('429006000', '429006', '2', '天门', 't', 'tian men', '429006000000', '天门市', '101201501');
INSERT INTO `demo_city_code` VALUES ('429021', '42', '1', '神农架', 's', 'shen nong jia', '429021000000', '神农架林区', '101201201');
INSERT INTO `demo_city_code` VALUES ('429021000', '429021', '2', '神农架林区', 's', 'shen nong jia lin qu', '429021000000', '神农架林区', NULL);
INSERT INTO `demo_city_code` VALUES ('43', '0', '0', '湖南', 'h', 'hu nan', '430000000000', '湖南省', NULL);
INSERT INTO `demo_city_code` VALUES ('4301', '43', '1', '长沙', 'c', 'chang sha', '430100000000', '长沙市', '101250101');
INSERT INTO `demo_city_code` VALUES ('430102', '4301', '2', '芙蓉', 'f', 'fu rong', '430102000000', '芙蓉区', NULL);
INSERT INTO `demo_city_code` VALUES ('430103', '4301', '2', '天心', 't', 'tian xin', '430103000000', '天心区', NULL);
INSERT INTO `demo_city_code` VALUES ('430104', '4301', '2', '岳麓', 'y', 'yue lu', '430104000000', '岳麓区', NULL);
INSERT INTO `demo_city_code` VALUES ('430105', '4301', '2', '开福', 'k', 'kai fu', '430105000000', '开福区', NULL);
INSERT INTO `demo_city_code` VALUES ('430111', '4301', '2', '雨花', 'y', 'yu hua', '430111000000', '雨花区', NULL);
INSERT INTO `demo_city_code` VALUES ('430112', '4301', '2', '望城', 'w', 'wang cheng', '430112000000', '望城区', '101250105');
INSERT INTO `demo_city_code` VALUES ('430121', '4301', '2', '长沙县', 'c', 'chang sha xian', '430121000000', '长沙县', NULL);
INSERT INTO `demo_city_code` VALUES ('430181', '4301', '2', '浏阳', 'l', 'liu yang', '430181000000', '浏阳市', '101250103');
INSERT INTO `demo_city_code` VALUES ('430182', '4301', '2', '宁乡', 'n', 'ning xiang', '430182000000', '宁乡市', '101250102');
INSERT INTO `demo_city_code` VALUES ('4302', '43', '1', '株洲', 'z', 'zhu zhou', '430200000000', '株洲市', '101250301');
INSERT INTO `demo_city_code` VALUES ('430202', '4302', '2', '荷塘', 'h', 'he tang', '430202000000', '荷塘区', NULL);
INSERT INTO `demo_city_code` VALUES ('430203', '4302', '2', '芦淞', 'l', 'lu song', '430203000000', '芦淞区', NULL);
INSERT INTO `demo_city_code` VALUES ('430204', '4302', '2', '石峰', 's', 'shi feng', '430204000000', '石峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('430211', '4302', '2', '天元', 't', 'tian yuan', '430211000000', '天元区', NULL);
INSERT INTO `demo_city_code` VALUES ('430212', '4302', '2', '渌口', 'l', 'lu kou', '430212000000', '渌口区', NULL);
INSERT INTO `demo_city_code` VALUES ('430223', '4302', '2', '攸县', 'y', 'you xian', '430223000000', '攸县', '101250302');
INSERT INTO `demo_city_code` VALUES ('430224', '4302', '2', '茶陵', 'c', 'cha ling', '430224000000', '茶陵县', '101250305');
INSERT INTO `demo_city_code` VALUES ('430225', '4302', '2', '炎陵', 'y', 'yan ling', '430225000000', '炎陵县', '101250306');
INSERT INTO `demo_city_code` VALUES ('430281', '4302', '2', '醴陵', 'l', 'li ling', '430281000000', '醴陵市', '101250303');
INSERT INTO `demo_city_code` VALUES ('4303', '43', '1', '湘潭', 'x', 'xiang tan', '430300000000', '湘潭市', '101250201');
INSERT INTO `demo_city_code` VALUES ('430302', '4303', '2', '雨湖', 'y', 'yu hu', '430302000000', '雨湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('430304', '4303', '2', '岳塘', 'y', 'yue tang', '430304000000', '岳塘区', NULL);
INSERT INTO `demo_city_code` VALUES ('430321', '4303', '2', '湘潭县', 'x', 'xiang tan xian', '430321000000', '湘潭县', NULL);
INSERT INTO `demo_city_code` VALUES ('430381', '4303', '2', '湘乡', 'x', 'xiang xiang', '430381000000', '湘乡市', '101250203');
INSERT INTO `demo_city_code` VALUES ('430382', '4303', '2', '韶山', 's', 'shao shan', '430382000000', '韶山市', '101250202');
INSERT INTO `demo_city_code` VALUES ('4304', '43', '1', '衡阳', 'h', 'heng yang', '430400000000', '衡阳市', '101250401');
INSERT INTO `demo_city_code` VALUES ('430405', '4304', '2', '珠晖', 'z', 'zhu hui', '430405000000', '珠晖区', NULL);
INSERT INTO `demo_city_code` VALUES ('430406', '4304', '2', '雁峰', 'y', 'yan feng', '430406000000', '雁峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('430407', '4304', '2', '石鼓', 's', 'shi gu', '430407000000', '石鼓区', NULL);
INSERT INTO `demo_city_code` VALUES ('430408', '4304', '2', '蒸湘', 'z', 'zheng xiang', '430408000000', '蒸湘区', NULL);
INSERT INTO `demo_city_code` VALUES ('430412', '4304', '2', '南岳', 'n', 'nan yue', '430412000000', '南岳区', '101250409');
INSERT INTO `demo_city_code` VALUES ('430421', '4304', '2', '衡阳县', 'h', 'heng yang xian', '430421000000', '衡阳县', '101250405');
INSERT INTO `demo_city_code` VALUES ('430422', '4304', '2', '衡南', 'h', 'heng nan', '430422000000', '衡南县', '101250407');
INSERT INTO `demo_city_code` VALUES ('430423', '4304', '2', '衡山', 'h', 'heng shan', '430423000000', '衡山县', '101250402');
INSERT INTO `demo_city_code` VALUES ('430424', '4304', '2', '衡东', 'h', 'heng dong', '430424000000', '衡东县', '101250403');
INSERT INTO `demo_city_code` VALUES ('430426', '4304', '2', '祁东', 'q', 'qi dong', '430426000000', '祁东县', '101250404');
INSERT INTO `demo_city_code` VALUES ('430481', '4304', '2', '耒阳', 'l', 'lei yang', '430481000000', '耒阳市', '101250408');
INSERT INTO `demo_city_code` VALUES ('430482', '4304', '2', '常宁', 'c', 'chang ning', '430482000000', '常宁市', '101250406');
INSERT INTO `demo_city_code` VALUES ('4305', '43', '1', '邵阳', 's', 'shao yang', '430500000000', '邵阳市', '101250901');
INSERT INTO `demo_city_code` VALUES ('430502', '4305', '2', '双清', 's', 'shuang qing', '430502000000', '双清区', NULL);
INSERT INTO `demo_city_code` VALUES ('430503', '4305', '2', '大祥', 'd', 'da xiang', '430503000000', '大祥区', NULL);
INSERT INTO `demo_city_code` VALUES ('430511', '4305', '2', '北塔', 'b', 'bei ta', '430511000000', '北塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('430522', '4305', '2', '新邵', 'x', 'xin shao', '430522000000', '新邵县', '101250904');
INSERT INTO `demo_city_code` VALUES ('430523', '4305', '2', '邵阳县', 's', 'shao yang xian', '430523000000', '邵阳县', '101250910');
INSERT INTO `demo_city_code` VALUES ('430524', '4305', '2', '隆回', 'l', 'long hui', '430524000000', '隆回县', '101250902');
INSERT INTO `demo_city_code` VALUES ('430525', '4305', '2', '洞口', 'd', 'dong kou', '430525000000', '洞口县', '101250903');
INSERT INTO `demo_city_code` VALUES ('430527', '4305', '2', '绥宁', 's', 'sui ning', '430527000000', '绥宁县', '101250906');
INSERT INTO `demo_city_code` VALUES ('430528', '4305', '2', '新宁', 'x', 'xin ning', '430528000000', '新宁县', '101250907');
INSERT INTO `demo_city_code` VALUES ('430529', '4305', '2', '城步', 'c', 'cheng bu', '430529000000', '城步苗族自治县', '101250909');
INSERT INTO `demo_city_code` VALUES ('430581', '4305', '2', '武冈', 'w', 'wu gang', '430581000000', '武冈市', '101250908');
INSERT INTO `demo_city_code` VALUES ('430582', '4305', '2', '邵东', 's', 'shao dong', '430582000000', '邵东市', '101250905');
INSERT INTO `demo_city_code` VALUES ('4306', '43', '1', '岳阳', 'y', 'yue yang', '430600000000', '岳阳市', '101251001');
INSERT INTO `demo_city_code` VALUES ('430602', '4306', '2', '岳阳楼', 'y', 'yue yang lou', '430602000000', '岳阳楼区', NULL);
INSERT INTO `demo_city_code` VALUES ('430603', '4306', '2', '云溪', 'y', 'yun xi', '430603000000', '云溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('430611', '4306', '2', '君山', 'j', 'jun shan', '430611000000', '君山区', NULL);
INSERT INTO `demo_city_code` VALUES ('430621', '4306', '2', '岳阳县', 'y', 'yue yang xian', '430621000000', '岳阳县', NULL);
INSERT INTO `demo_city_code` VALUES ('430623', '4306', '2', '华容', 'h', 'hua rong', '430623000000', '华容县', '101251002');
INSERT INTO `demo_city_code` VALUES ('430624', '4306', '2', '湘阴', 'x', 'xiang yin', '430624000000', '湘阴县', '101251003');
INSERT INTO `demo_city_code` VALUES ('430626', '4306', '2', '平江', 'p', 'ping jiang', '430626000000', '平江县', '101251005');
INSERT INTO `demo_city_code` VALUES ('430681', '4306', '2', '汨罗', 'm', 'mi luo', '430681000000', '汨罗市', '101251004');
INSERT INTO `demo_city_code` VALUES ('430682', '4306', '2', '临湘', 'l', 'lin xiang', '430682000000', '临湘市', '101251006');
INSERT INTO `demo_city_code` VALUES ('4307', '43', '1', '常德', 'c', 'chang de', '430700000000', '常德市', '101250601');
INSERT INTO `demo_city_code` VALUES ('430702', '4307', '2', '武陵', 'w', 'wu ling', '430702000000', '武陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('430703', '4307', '2', '鼎城', 'd', 'ding cheng', '430703000000', '鼎城区', NULL);
INSERT INTO `demo_city_code` VALUES ('430721', '4307', '2', '安乡', 'a', 'an xiang', '430721000000', '安乡县', '101250602');
INSERT INTO `demo_city_code` VALUES ('430722', '4307', '2', '汉寿', 'h', 'han shou', '430722000000', '汉寿县', '101250604');
INSERT INTO `demo_city_code` VALUES ('430723', '4307', '2', '澧县', 'l', 'li xian', '430723000000', '澧县', '101250605');
INSERT INTO `demo_city_code` VALUES ('430724', '4307', '2', '临澧', 'l', 'lin li', '430724000000', '临澧县', '101250606');
INSERT INTO `demo_city_code` VALUES ('430725', '4307', '2', '桃源', 't', 'tao yuan', '430725000000', '桃源县', '101250603');
INSERT INTO `demo_city_code` VALUES ('430726', '4307', '2', '石门', 's', 'shi men', '430726000000', '石门县', '101250607');
INSERT INTO `demo_city_code` VALUES ('430781', '4307', '2', '津市', 'j', 'jin shi', '430781000000', '津市市', '101250608');
INSERT INTO `demo_city_code` VALUES ('4308', '43', '1', '张家界', 'z', 'zhang jia jie', '430800000000', '张家界市', '101251101');
INSERT INTO `demo_city_code` VALUES ('430802', '4308', '2', '永定', 'y', 'yong ding', '430802000000', '永定区', '101230706');
INSERT INTO `demo_city_code` VALUES ('430811', '4308', '2', '武陵源', 'w', 'wu ling yuan', '430811000000', '武陵源区', '101251104');
INSERT INTO `demo_city_code` VALUES ('430821', '4308', '2', '慈利', 'c', 'ci li', '430821000000', '慈利县', '101251103');
INSERT INTO `demo_city_code` VALUES ('430822', '4308', '2', '桑植', 's', 'sang zhi', '430822000000', '桑植县', '101251102');
INSERT INTO `demo_city_code` VALUES ('4309', '43', '1', '益阳', 'y', 'yi yang', '430900000000', '益阳市', '101250700');
INSERT INTO `demo_city_code` VALUES ('430902', '4309', '2', '资阳', 'z', 'zi yang', '430902000000', '资阳区', '101271301');
INSERT INTO `demo_city_code` VALUES ('430903', '4309', '2', '赫山', 'h', 'he shan', '430903000000', '赫山区', NULL);
INSERT INTO `demo_city_code` VALUES ('430921', '4309', '2', '南县', 'n', 'nan xian', '430921000000', '南县', '101250702');
INSERT INTO `demo_city_code` VALUES ('430922', '4309', '2', '桃江', 't', 'tao jiang', '430922000000', '桃江县', '101250703');
INSERT INTO `demo_city_code` VALUES ('430923', '4309', '2', '安化', 'a', 'an hua', '430923000000', '安化县', '101250704');
INSERT INTO `demo_city_code` VALUES ('430981', '4309', '2', '沅江', 'y', 'yuan jiang', '430981000000', '沅江市', '101250705');
INSERT INTO `demo_city_code` VALUES ('4310', '43', '1', '郴州', 'c', 'chen zhou', '431000000000', '郴州市', '101250501');
INSERT INTO `demo_city_code` VALUES ('431002', '4310', '2', '北湖', 'b', 'bei hu', '431002000000', '北湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('431003', '4310', '2', '苏仙', 's', 'su xian', '431003000000', '苏仙区', '101250512');
INSERT INTO `demo_city_code` VALUES ('431021', '4310', '2', '桂阳', 'g', 'gui yang', '431021000000', '桂阳县', '101250502');
INSERT INTO `demo_city_code` VALUES ('431022', '4310', '2', '宜章', 'y', 'yi zhang', '431022000000', '宜章县', '101250504');
INSERT INTO `demo_city_code` VALUES ('431023', '4310', '2', '永兴', 'y', 'yong xing', '431023000000', '永兴县', '101250510');
INSERT INTO `demo_city_code` VALUES ('431024', '4310', '2', '嘉禾', 'j', 'jia he', '431024000000', '嘉禾县', '101250503');
INSERT INTO `demo_city_code` VALUES ('431025', '4310', '2', '临武', 'l', 'lin wu', '431025000000', '临武县', '101250505');
INSERT INTO `demo_city_code` VALUES ('431026', '4310', '2', '汝城', 'r', 'ru cheng', '431026000000', '汝城县', '101250508');
INSERT INTO `demo_city_code` VALUES ('431027', '4310', '2', '桂东', 'g', 'gui dong', '431027000000', '桂东县', '101250511');
INSERT INTO `demo_city_code` VALUES ('431028', '4310', '2', '安仁', 'a', 'an ren', '431028000000', '安仁县', '101250509');
INSERT INTO `demo_city_code` VALUES ('431081', '4310', '2', '资兴', 'z', 'zi xing', '431081000000', '资兴市', '101250507');
INSERT INTO `demo_city_code` VALUES ('4311', '43', '1', '永州', 'y', 'yong zhou', '431100000000', '永州市', '101251401');
INSERT INTO `demo_city_code` VALUES ('431102', '4311', '2', '零陵', 'l', 'ling ling', '431102000000', '零陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('431103', '4311', '2', '冷水滩', 'l', 'leng shui tan', '431103000000', '冷水滩区', '101251411');
INSERT INTO `demo_city_code` VALUES ('431122', '4311', '2', '东安', 'd', 'dong an', '431122000000', '东安县', '101251403');
INSERT INTO `demo_city_code` VALUES ('431123', '4311', '2', '双牌', 's', 'shuang pai', '431123000000', '双牌县', '101251404');
INSERT INTO `demo_city_code` VALUES ('431124', '4311', '2', '道县', 'd', 'dao xian', '431124000000', '道县', '101251405');
INSERT INTO `demo_city_code` VALUES ('431125', '4311', '2', '江永', 'j', 'jiang yong', '431125000000', '江永县', '101251407');
INSERT INTO `demo_city_code` VALUES ('431126', '4311', '2', '宁远', 'n', 'ning yuan', '431126000000', '宁远县', '101251406');
INSERT INTO `demo_city_code` VALUES ('431127', '4311', '2', '蓝山', 'l', 'lan shan', '431127000000', '蓝山县', '101251408');
INSERT INTO `demo_city_code` VALUES ('431128', '4311', '2', '新田', 'x', 'xin tian', '431128000000', '新田县', '101251409');
INSERT INTO `demo_city_code` VALUES ('431129', '4311', '2', '江华', 'j', 'jiang hua', '431129000000', '江华瑶族自治县', '101251410');
INSERT INTO `demo_city_code` VALUES ('431181', '4311', '2', '祁阳', 'q', 'qi yang', '431181000000', '祁阳市', '101251402');
INSERT INTO `demo_city_code` VALUES ('4312', '43', '1', '怀化', 'h', 'huai hua', '431200000000', '怀化市', '101251201');
INSERT INTO `demo_city_code` VALUES ('431202', '4312', '2', '鹤城', 'h', 'he cheng', '431202000000', '鹤城区', NULL);
INSERT INTO `demo_city_code` VALUES ('431221', '4312', '2', '中方', 'z', 'zhong fang', '431221000000', '中方县', '101251212');
INSERT INTO `demo_city_code` VALUES ('431222', '4312', '2', '沅陵', 'y', 'yuan ling', '431222000000', '沅陵县', '101251203');
INSERT INTO `demo_city_code` VALUES ('431223', '4312', '2', '辰溪', 'c', 'chen xi', '431223000000', '辰溪县', '101251204');
INSERT INTO `demo_city_code` VALUES ('431224', '4312', '2', '溆浦', 'x', 'xu pu', '431224000000', '溆浦县', '101251211');
INSERT INTO `demo_city_code` VALUES ('431225', '4312', '2', '会同', 'h', 'hui tong', '431225000000', '会同县', '101251206');
INSERT INTO `demo_city_code` VALUES ('431226', '4312', '2', '麻阳', 'm', 'ma yang', '431226000000', '麻阳苗族自治县', '101251208');
INSERT INTO `demo_city_code` VALUES ('431227', '4312', '2', '新晃', 'x', 'xin huang', '431227000000', '新晃侗族自治县', '101251209');
INSERT INTO `demo_city_code` VALUES ('431228', '4312', '2', '芷江', 'z', 'zhi jiang', '431228000000', '芷江侗族自治县', '101251210');
INSERT INTO `demo_city_code` VALUES ('431229', '4312', '2', '靖州', 'j', 'jing zhou', '431229000000', '靖州苗族侗族自治县', '101251205');
INSERT INTO `demo_city_code` VALUES ('431230', '4312', '2', '通道', 't', 'tong dao', '431230000000', '通道侗族自治县', '101251207');
INSERT INTO `demo_city_code` VALUES ('431281', '4312', '2', '洪江', 'h', 'hong jiang', '431281000000', '洪江市', '101251213');
INSERT INTO `demo_city_code` VALUES ('4313', '43', '1', '娄底', 'l', 'lou di', '431300000000', '娄底市', '101250801');
INSERT INTO `demo_city_code` VALUES ('431302', '4313', '2', '娄星', 'l', 'lou xing', '431302000000', '娄星区', NULL);
INSERT INTO `demo_city_code` VALUES ('431321', '4313', '2', '双峰', 's', 'shuang feng', '431321000000', '双峰县', '101250802');
INSERT INTO `demo_city_code` VALUES ('431322', '4313', '2', '新化', 'x', 'xin hua', '431322000000', '新化县', '101250805');
INSERT INTO `demo_city_code` VALUES ('431381', '4313', '2', '冷水江', 'l', 'leng shui jiang', '431381000000', '冷水江市', '101250803');
INSERT INTO `demo_city_code` VALUES ('431382', '4313', '2', '涟源', 'l', 'lian yuan', '431382000000', '涟源市', '101250806');
INSERT INTO `demo_city_code` VALUES ('4331', '43', '1', '湘西', 'x', 'xiang xi', '433100000000', '湘西土家族苗族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('433101', '4331', '2', '吉首', 'j', 'ji shou', '433101000000', '吉首市', '101251501');
INSERT INTO `demo_city_code` VALUES ('433122', '4331', '2', '泸溪', 'l', 'lu xi', '433122000000', '泸溪县', '101251506');
INSERT INTO `demo_city_code` VALUES ('433123', '4331', '2', '凤凰', 'f', 'feng huang', '433123000000', '凤凰县', '101251505');
INSERT INTO `demo_city_code` VALUES ('433124', '4331', '2', '花垣', 'h', 'hua yuan', '433124000000', '花垣县', '101251508');
INSERT INTO `demo_city_code` VALUES ('433125', '4331', '2', '保靖', 'b', 'bao jing', '433125000000', '保靖县', '101251502');
INSERT INTO `demo_city_code` VALUES ('433126', '4331', '2', '古丈', 'g', 'gu zhang', '433126000000', '古丈县', '101251504');
INSERT INTO `demo_city_code` VALUES ('433127', '4331', '2', '永顺', 'y', 'yong shun', '433127000000', '永顺县', '101251503');
INSERT INTO `demo_city_code` VALUES ('433130', '4331', '2', '龙山', 'l', 'long shan', '433130000000', '龙山县', '101251507');
INSERT INTO `demo_city_code` VALUES ('44', '0', '0', '广东', 'g', 'guang dong', '440000000000', '广东省', NULL);
INSERT INTO `demo_city_code` VALUES ('4401', '44', '1', '广州', 'g', 'guang zhou', '440100000000', '广州市', '101280101');
INSERT INTO `demo_city_code` VALUES ('440103', '4401', '2', '荔湾', 'l', 'li wan', '440103000000', '荔湾区', NULL);
INSERT INTO `demo_city_code` VALUES ('440104', '4401', '2', '越秀', 'y', 'yue xiu', '440104000000', '越秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('440105', '4401', '2', '海珠', 'h', 'hai zhu', '440105000000', '海珠区', NULL);
INSERT INTO `demo_city_code` VALUES ('440106', '4401', '2', '天河', 't', 'tian he', '440106000000', '天河区', NULL);
INSERT INTO `demo_city_code` VALUES ('440111', '4401', '2', '白云', 'b', 'bai yun', '440111000000', '白云区', '101260102');
INSERT INTO `demo_city_code` VALUES ('440112', '4401', '2', '黄埔', 'h', 'huang pu', '440112000000', '黄埔区', NULL);
INSERT INTO `demo_city_code` VALUES ('440113', '4401', '2', '番禺', 'p', 'pan yu', '440113000000', '番禺区', '101280102');
INSERT INTO `demo_city_code` VALUES ('440114', '4401', '2', '花都', 'h', 'hua du', '440114000000', '花都区', '101280105');
INSERT INTO `demo_city_code` VALUES ('440115', '4401', '2', '南沙', 'n', 'nan sha', '440115000000', '南沙区', '101310220');
INSERT INTO `demo_city_code` VALUES ('440117', '4401', '2', '从化', 'c', 'cong hua', '440117000000', '从化区', '101280103');
INSERT INTO `demo_city_code` VALUES ('440118', '4401', '2', '增城', 'z', 'zeng cheng', '440118000000', '增城区', '101280104');
INSERT INTO `demo_city_code` VALUES ('4402', '44', '1', '韶关', 's', 'shao guan', '440200000000', '韶关市', '101280201');
INSERT INTO `demo_city_code` VALUES ('440203', '4402', '2', '武江', 'w', 'wu jiang', '440203000000', '武江区', '101280211');
INSERT INTO `demo_city_code` VALUES ('440204', '4402', '2', '浈江', 'z', 'zhen jiang', '440204000000', '浈江区', '101280210');
INSERT INTO `demo_city_code` VALUES ('440205', '4402', '2', '曲江', 'q', 'qu jiang', '440205000000', '曲江区', '101280209');
INSERT INTO `demo_city_code` VALUES ('440222', '4402', '2', '始兴', 's', 'shi xing', '440222000000', '始兴县', '101280203');
INSERT INTO `demo_city_code` VALUES ('440224', '4402', '2', '仁化', 'r', 'ren hua', '440224000000', '仁化县', '101280206');
INSERT INTO `demo_city_code` VALUES ('440229', '4402', '2', '翁源', 'w', 'weng yuan', '440229000000', '翁源县', '101280204');
INSERT INTO `demo_city_code` VALUES ('440232', '4402', '2', '乳源', 'r', 'ru yuan', '440232000000', '乳源瑶族自治县', '101280202');
INSERT INTO `demo_city_code` VALUES ('440233', '4402', '2', '新丰', 'x', 'xin feng', '440233000000', '新丰县', '101280208');
INSERT INTO `demo_city_code` VALUES ('440281', '4402', '2', '乐昌', 'l', 'le chang', '440281000000', '乐昌市', '101280205');
INSERT INTO `demo_city_code` VALUES ('440282', '4402', '2', '南雄', 'n', 'nan xiong', '440282000000', '南雄市', '101280207');
INSERT INTO `demo_city_code` VALUES ('4403', '44', '1', '深圳', 's', 'shen zhen', '440300000000', '深圳市', '101280601');
INSERT INTO `demo_city_code` VALUES ('440303', '4403', '2', '罗湖', 'l', 'luo hu', '440303000000', '罗湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('440304', '4403', '2', '福田', 'f', 'fu tian', '440304000000', '福田区', NULL);
INSERT INTO `demo_city_code` VALUES ('440305', '4403', '2', '南山', 'n', 'nan shan', '440305000000', '南山区', NULL);
INSERT INTO `demo_city_code` VALUES ('440306', '4403', '2', '宝安', 'b', 'bao an', '440306000000', '宝安区', NULL);
INSERT INTO `demo_city_code` VALUES ('440307', '4403', '2', '龙岗', 'l', 'long gang', '440307000000', '龙岗区', NULL);
INSERT INTO `demo_city_code` VALUES ('440308', '4403', '2', '盐田', 'y', 'yan tian', '440308000000', '盐田区', NULL);
INSERT INTO `demo_city_code` VALUES ('440309', '4403', '2', '龙华', 'l', 'long hua', '440309000000', '龙华区', NULL);
INSERT INTO `demo_city_code` VALUES ('440310', '4403', '2', '坪山', 'p', 'ping shan', '440310000000', '坪山区', NULL);
INSERT INTO `demo_city_code` VALUES ('440311', '4403', '2', '光明', 'g', 'guang ming', '440311000000', '光明区', NULL);
INSERT INTO `demo_city_code` VALUES ('4404', '44', '1', '珠海', 'z', 'zhu hai', '440400000000', '珠海市', '101280701');
INSERT INTO `demo_city_code` VALUES ('440402', '4404', '2', '香洲', 'x', 'xiang zhou', '440402000000', '香洲区', NULL);
INSERT INTO `demo_city_code` VALUES ('440403', '4404', '2', '斗门', 'd', 'dou men', '440403000000', '斗门区', '101280702');
INSERT INTO `demo_city_code` VALUES ('440404', '4404', '2', '金湾', 'j', 'jin wan', '440404000000', '金湾区', '101280703');
INSERT INTO `demo_city_code` VALUES ('4405', '44', '1', '汕头', 's', 'shan tou', '440500000000', '汕头市', '101280501');
INSERT INTO `demo_city_code` VALUES ('440507', '4405', '2', '龙湖', 'l', 'long hu', '440507000000', '龙湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('440511', '4405', '2', '金平', 'j', 'jin ping', '440511000000', '金平区', '101290312');
INSERT INTO `demo_city_code` VALUES ('440512', '4405', '2', '濠江', 'h', 'hao jiang', '440512000000', '濠江区', NULL);
INSERT INTO `demo_city_code` VALUES ('440513', '4405', '2', '潮阳', 'c', 'chao yang', '440513000000', '潮阳区', '101280502');
INSERT INTO `demo_city_code` VALUES ('440514', '4405', '2', '潮南', 'c', 'chao nan', '440514000000', '潮南区', NULL);
INSERT INTO `demo_city_code` VALUES ('440515', '4405', '2', '澄海', 'c', 'cheng hai', '440515000000', '澄海区', '101280503');
INSERT INTO `demo_city_code` VALUES ('440523', '4405', '2', '南澳', 'n', 'nan ao', '440523000000', '南澳县', '101280504');
INSERT INTO `demo_city_code` VALUES ('4406', '44', '1', '佛山', 'f', 'fo shan', '440600000000', '佛山市', '101280800');
INSERT INTO `demo_city_code` VALUES ('440604', '4406', '2', '禅城', 'c', 'chan cheng', '440604000000', '禅城区', NULL);
INSERT INTO `demo_city_code` VALUES ('440605', '4406', '2', '南海', 'n', 'nan hai', '440605000000', '南海区', '101280803');
INSERT INTO `demo_city_code` VALUES ('440606', '4406', '2', '顺德', 's', 'shun de', '440606000000', '顺德区', '101280801');
INSERT INTO `demo_city_code` VALUES ('440607', '4406', '2', '三水', 's', 'san shui', '440607000000', '三水区', '101280802');
INSERT INTO `demo_city_code` VALUES ('440608', '4406', '2', '高明', 'g', 'gao ming', '440608000000', '高明区', '101280804');
INSERT INTO `demo_city_code` VALUES ('4407', '44', '1', '江门', 'j', 'jiang men', '440700000000', '江门市', '101281101');
INSERT INTO `demo_city_code` VALUES ('440703', '4407', '2', '蓬江', 'p', 'peng jiang', '440703000000', '蓬江区', '101281107');
INSERT INTO `demo_city_code` VALUES ('440704', '4407', '2', '江海', 'j', 'jiang hai', '440704000000', '江海区', '101281109');
INSERT INTO `demo_city_code` VALUES ('440705', '4407', '2', '新会', 'x', 'xin hui', '440705000000', '新会区', '101281104');
INSERT INTO `demo_city_code` VALUES ('440781', '4407', '2', '台山', 't', 'tai shan', '440781000000', '台山市', '101281106');
INSERT INTO `demo_city_code` VALUES ('440783', '4407', '2', '开平', 'k', 'kai ping', '440783000000', '开平市', '101281103');
INSERT INTO `demo_city_code` VALUES ('440784', '4407', '2', '鹤山', 'h', 'he shan', '440784000000', '鹤山市', '101281108');
INSERT INTO `demo_city_code` VALUES ('440785', '4407', '2', '恩平', 'e', 'en ping', '440785000000', '恩平市', '101281105');
INSERT INTO `demo_city_code` VALUES ('4408', '44', '1', '湛江', 'z', 'zhan jiang', '440800000000', '湛江市', '101281001');
INSERT INTO `demo_city_code` VALUES ('440802', '4408', '2', '赤坎', 'c', 'chi kan', '440802000000', '赤坎区', '101281006');
INSERT INTO `demo_city_code` VALUES ('440803', '4408', '2', '霞山', 'x', 'xia shan', '440803000000', '霞山区', '101281009');
INSERT INTO `demo_city_code` VALUES ('440804', '4408', '2', '坡头', 'p', 'po tou', '440804000000', '坡头区', '101281008');
INSERT INTO `demo_city_code` VALUES ('440811', '4408', '2', '麻章', 'm', 'ma zhang', '440811000000', '麻章区', '101281010');
INSERT INTO `demo_city_code` VALUES ('440823', '4408', '2', '遂溪', 's', 'sui xi', '440823000000', '遂溪县', '101281007');
INSERT INTO `demo_city_code` VALUES ('440825', '4408', '2', '徐闻', 'x', 'xu wen', '440825000000', '徐闻县', '101281004');
INSERT INTO `demo_city_code` VALUES ('440881', '4408', '2', '廉江', 'l', 'lian jiang', '440881000000', '廉江市', '101281005');
INSERT INTO `demo_city_code` VALUES ('440882', '4408', '2', '雷州', 'l', 'lei zhou', '440882000000', '雷州市', '101281003');
INSERT INTO `demo_city_code` VALUES ('440883', '4408', '2', '吴川', 'w', 'wu chuan', '440883000000', '吴川市', '101281002');
INSERT INTO `demo_city_code` VALUES ('4409', '44', '1', '茂名', 'm', 'mao ming', '440900000000', '茂名市', '101282001');
INSERT INTO `demo_city_code` VALUES ('440902', '4409', '2', '茂南', 'm', 'mao nan', '440902000000', '茂南区', NULL);
INSERT INTO `demo_city_code` VALUES ('440904', '4409', '2', '电白', 'd', 'dian bai', '440904000000', '电白区', '101282004');
INSERT INTO `demo_city_code` VALUES ('440981', '4409', '2', '高州', 'g', 'gao zhou', '440981000000', '高州市', '101282002');
INSERT INTO `demo_city_code` VALUES ('440982', '4409', '2', '化州', 'h', 'hua zhou', '440982000000', '化州市', '101282003');
INSERT INTO `demo_city_code` VALUES ('440983', '4409', '2', '信宜', 'x', 'xin yi', '440983000000', '信宜市', '101282005');
INSERT INTO `demo_city_code` VALUES ('4412', '44', '1', '肇庆', 'z', 'zhao qing', '441200000000', '肇庆市', '101280901');
INSERT INTO `demo_city_code` VALUES ('441202', '4412', '2', '端州', 'd', 'duan zhou', '441202000000', '端州区', NULL);
INSERT INTO `demo_city_code` VALUES ('441203', '4412', '2', '鼎湖', 'd', 'ding hu', '441203000000', '鼎湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('441204', '4412', '2', '高要', 'g', 'gao yao', '441204000000', '高要区', '101280908');
INSERT INTO `demo_city_code` VALUES ('441223', '4412', '2', '广宁', 'g', 'guang ning', '441223000000', '广宁县', '101280902');
INSERT INTO `demo_city_code` VALUES ('441224', '4412', '2', '怀集', 'h', 'huai ji', '441224000000', '怀集县', '101280906');
INSERT INTO `demo_city_code` VALUES ('441225', '4412', '2', '封开', 'f', 'feng kai', '441225000000', '封开县', '101280907');
INSERT INTO `demo_city_code` VALUES ('441226', '4412', '2', '德庆', 'd', 'de qing', '441226000000', '德庆县', '101280905');
INSERT INTO `demo_city_code` VALUES ('441284', '4412', '2', '四会', 's', 'si hui', '441284000000', '四会市', '101280903');
INSERT INTO `demo_city_code` VALUES ('4413', '44', '1', '惠州', 'h', 'hui zhou', '441300000000', '惠州市', '101280301');
INSERT INTO `demo_city_code` VALUES ('441302', '4413', '2', '惠城', 'h', 'hui cheng', '441302000000', '惠城区', NULL);
INSERT INTO `demo_city_code` VALUES ('441303', '4413', '2', '惠阳', 'h', 'hui yang', '441303000000', '惠阳区', '101280303');
INSERT INTO `demo_city_code` VALUES ('441322', '4413', '2', '博罗', 'b', 'bo luo', '441322000000', '博罗县', '101280302');
INSERT INTO `demo_city_code` VALUES ('441323', '4413', '2', '惠东', 'h', 'hui dong', '441323000000', '惠东县', '101280304');
INSERT INTO `demo_city_code` VALUES ('441324', '4413', '2', '龙门', 'l', 'long men', '441324000000', '龙门县', '101280305');
INSERT INTO `demo_city_code` VALUES ('4414', '44', '1', '梅州', 'm', 'mei zhou', '441400000000', '梅州市', '101280401');
INSERT INTO `demo_city_code` VALUES ('441402', '4414', '2', '梅江', 'm', 'mei jiang', '441402000000', '梅江区', NULL);
INSERT INTO `demo_city_code` VALUES ('441403', '4414', '2', '梅县', 'm', 'mei xian', '441403000000', '梅县区', '101280409');
INSERT INTO `demo_city_code` VALUES ('441422', '4414', '2', '大埔', 'd', 'da bu', '441422000000', '大埔县', '101280404');
INSERT INTO `demo_city_code` VALUES ('441423', '4414', '2', '丰顺', 'f', 'feng shun', '441423000000', '丰顺县', '101280406');
INSERT INTO `demo_city_code` VALUES ('441424', '4414', '2', '五华', 'w', 'wu hua', '441424000000', '五华县', '101280408');
INSERT INTO `demo_city_code` VALUES ('441426', '4414', '2', '平远', 'p', 'ping yuan', '441426000000', '平远县', '101280407');
INSERT INTO `demo_city_code` VALUES ('441427', '4414', '2', '蕉岭', 'j', 'jiao ling', '441427000000', '蕉岭县', '101280403');
INSERT INTO `demo_city_code` VALUES ('441481', '4414', '2', '兴宁', 'x', 'xing ning', '441481000000', '兴宁市', '101280402');
INSERT INTO `demo_city_code` VALUES ('4415', '44', '1', '汕尾', 's', 'shan wei', '441500000000', '汕尾市', '101282101');
INSERT INTO `demo_city_code` VALUES ('441502', '4415', '2', '城区', 'c', 'cheng qu', '441502000000', '城区', NULL);
INSERT INTO `demo_city_code` VALUES ('441521', '4415', '2', '海丰', 'h', 'hai feng', '441521000000', '海丰县', '101282102');
INSERT INTO `demo_city_code` VALUES ('441523', '4415', '2', '陆河', 'l', 'lu he', '441523000000', '陆河县', '101282104');
INSERT INTO `demo_city_code` VALUES ('441581', '4415', '2', '陆丰', 'l', 'lu feng', '441581000000', '陆丰市', '101282103');
INSERT INTO `demo_city_code` VALUES ('4416', '44', '1', '河源', 'h', 'he yuan', '441600000000', '河源市', '101281201');
INSERT INTO `demo_city_code` VALUES ('441602', '4416', '2', '源城', 'y', 'yuan cheng', '441602000000', '源城区', NULL);
INSERT INTO `demo_city_code` VALUES ('441621', '4416', '2', '紫金', 'z', 'zi jin', '441621000000', '紫金县', '101281202');
INSERT INTO `demo_city_code` VALUES ('441622', '4416', '2', '龙川', 'l', 'long chuan', '441622000000', '龙川县', '101281205');
INSERT INTO `demo_city_code` VALUES ('441623', '4416', '2', '连平', 'l', 'lian ping', '441623000000', '连平县', '101281203');
INSERT INTO `demo_city_code` VALUES ('441624', '4416', '2', '和平', 'h', 'he ping', '441624000000', '和平县', '101281204');
INSERT INTO `demo_city_code` VALUES ('441625', '4416', '2', '东源', 'd', 'dong yuan', '441625000000', '东源县', '101281206');
INSERT INTO `demo_city_code` VALUES ('4417', '44', '1', '阳江', 'y', 'yang jiang', '441700000000', '阳江市', '101281801');
INSERT INTO `demo_city_code` VALUES ('441702', '4417', '2', '江城', 'j', 'jiang cheng', '441702000000', '江城区', '101290907');
INSERT INTO `demo_city_code` VALUES ('441704', '4417', '2', '阳东', 'y', 'yang dong', '441704000000', '阳东区', '101281803');
INSERT INTO `demo_city_code` VALUES ('441721', '4417', '2', '阳西', 'y', 'yang xi', '441721000000', '阳西县', '101281804');
INSERT INTO `demo_city_code` VALUES ('441781', '4417', '2', '阳春', 'y', 'yang chun', '441781000000', '阳春市', '101281802');
INSERT INTO `demo_city_code` VALUES ('4418', '44', '1', '清远', 'q', 'qing yuan', '441800000000', '清远市', '101281301');
INSERT INTO `demo_city_code` VALUES ('441802', '4418', '2', '清城', 'q', 'qing cheng', '441802000000', '清城区', NULL);
INSERT INTO `demo_city_code` VALUES ('441803', '4418', '2', '清新区', 'q', 'qing xin qu', '441803000000', '清新区', NULL);
INSERT INTO `demo_city_code` VALUES ('441821', '4418', '2', '佛冈', 'f', 'fo gang', '441821000000', '佛冈县', '101281306');
INSERT INTO `demo_city_code` VALUES ('441823', '4418', '2', '阳山', 'y', 'yang shan', '441823000000', '阳山县', '101281305');
INSERT INTO `demo_city_code` VALUES ('441825', '4418', '2', '连山', 'l', 'lian shan', '441825000000', '连山壮族瑶族自治县', '101281304');
INSERT INTO `demo_city_code` VALUES ('441826', '4418', '2', '连南', 'l', 'lian nan', '441826000000', '连南瑶族自治县', '101281302');
INSERT INTO `demo_city_code` VALUES ('441881', '4418', '2', '英德', 'y', 'ying de', '441881000000', '英德市', '101281307');
INSERT INTO `demo_city_code` VALUES ('441882', '4418', '2', '连州', 'l', 'lian zhou', '441882000000', '连州市', '101281303');
INSERT INTO `demo_city_code` VALUES ('4419', '44', '1', '东莞', 'd', 'dong guan', '441900000000', '东莞市', '101281601');
INSERT INTO `demo_city_code` VALUES ('441900', '4419', '2', '东莞', 'd', 'dong guan', '441900000000', '东莞市', '101281601');
INSERT INTO `demo_city_code` VALUES ('4420', '44', '1', '中山', 'z', 'zhong shan', '442000000000', '中山市', '101281701');
INSERT INTO `demo_city_code` VALUES ('442000', '4420', '2', '中山', 'z', 'zhong shan', '442000000000', '中山市', '101281701');
INSERT INTO `demo_city_code` VALUES ('4451', '44', '1', '潮州', 'c', 'chao zhou', '445100000000', '潮州市', '101281501');
INSERT INTO `demo_city_code` VALUES ('445102', '4451', '2', '湘桥', 'x', 'xiang qiao', '445102000000', '湘桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('445103', '4451', '2', '潮安', 'c', 'chao an', '445103000000', '潮安区', '101281503');
INSERT INTO `demo_city_code` VALUES ('445122', '4451', '2', '饶平', 'r', 'rao ping', '445122000000', '饶平县', '101281502');
INSERT INTO `demo_city_code` VALUES ('4452', '44', '1', '揭阳', 'j', 'jie yang', '445200000000', '揭阳市', '101281901');
INSERT INTO `demo_city_code` VALUES ('445202', '4452', '2', '榕城', 'r', 'rong cheng', '445202000000', '榕城区', NULL);
INSERT INTO `demo_city_code` VALUES ('445203', '4452', '2', '揭东', 'j', 'jie dong', '445203000000', '揭东区', '101281905');
INSERT INTO `demo_city_code` VALUES ('445222', '4452', '2', '揭西', 'j', 'jie xi', '445222000000', '揭西县', '101281902');
INSERT INTO `demo_city_code` VALUES ('445224', '4452', '2', '惠来', 'h', 'hui lai', '445224000000', '惠来县', '101281904');
INSERT INTO `demo_city_code` VALUES ('445281', '4452', '2', '普宁', 'p', 'pu ning', '445281000000', '普宁市', '101281903');
INSERT INTO `demo_city_code` VALUES ('4453', '44', '1', '云浮', 'y', 'yun fu', '445300000000', '云浮市', '101281401');
INSERT INTO `demo_city_code` VALUES ('445302', '4453', '2', '云城', 'y', 'yun cheng', '445302000000', '云城区', NULL);
INSERT INTO `demo_city_code` VALUES ('445303', '4453', '2', '云安', 'y', 'yun an', '445303000000', '云安区', '101281406');
INSERT INTO `demo_city_code` VALUES ('445321', '4453', '2', '新兴', 'x', 'xin xing', '445321000000', '新兴县', '101281403');
INSERT INTO `demo_city_code` VALUES ('445322', '4453', '2', '郁南', 'y', 'yu nan', '445322000000', '郁南县', '101281404');
INSERT INTO `demo_city_code` VALUES ('445381', '4453', '2', '罗定', 'l', 'luo ding', '445381000000', '罗定市', '101281402');
INSERT INTO `demo_city_code` VALUES ('45', '0', '0', '广西', 'g', 'guang xi', '450000000000', '广西壮族自治区', NULL);
INSERT INTO `demo_city_code` VALUES ('4501', '45', '1', '南宁', 'n', 'nan ning', '450100000000', '南宁市', '101300101');
INSERT INTO `demo_city_code` VALUES ('450102', '4501', '2', '兴宁', 'x', 'xing ning', '450102000000', '兴宁区', '101280402');
INSERT INTO `demo_city_code` VALUES ('450103', '4501', '2', '青秀', 'q', 'qing xiu', '450103000000', '青秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('450105', '4501', '2', '江南', 'j', 'jiang nan', '450105000000', '江南区', NULL);
INSERT INTO `demo_city_code` VALUES ('450107', '4501', '2', '西乡塘', 'x', 'xi xiang tang', '450107000000', '西乡塘区', NULL);
INSERT INTO `demo_city_code` VALUES ('450108', '4501', '2', '良庆', 'l', 'liang qing', '450108000000', '良庆区', NULL);
INSERT INTO `demo_city_code` VALUES ('450109', '4501', '2', '邕宁', 'y', 'yong ning', '450109000000', '邕宁区', '101300103');
INSERT INTO `demo_city_code` VALUES ('450110', '4501', '2', '武鸣', 'w', 'wu ming', '450110000000', '武鸣区', '101300108');
INSERT INTO `demo_city_code` VALUES ('450123', '4501', '2', '隆安', 'l', 'long an', '450123000000', '隆安县', '101300105');
INSERT INTO `demo_city_code` VALUES ('450124', '4501', '2', '马山', 'm', 'ma shan', '450124000000', '马山县', '101300106');
INSERT INTO `demo_city_code` VALUES ('450125', '4501', '2', '上林', 's', 'shang lin', '450125000000', '上林县', '101300107');
INSERT INTO `demo_city_code` VALUES ('450126', '4501', '2', '宾阳', 'b', 'bin yang', '450126000000', '宾阳县', '101300109');
INSERT INTO `demo_city_code` VALUES ('450181', '4501', '2', '横州', 'h', 'heng zhou', '450181000000', '横州市', NULL);
INSERT INTO `demo_city_code` VALUES ('4502', '45', '1', '柳州', 'l', 'liu zhou', '450200000000', '柳州市', '101300301');
INSERT INTO `demo_city_code` VALUES ('450202', '4502', '2', '城中', 'c', 'cheng zhong', '450202000000', '城中区', NULL);
INSERT INTO `demo_city_code` VALUES ('450203', '4502', '2', '鱼峰', 'y', 'yu feng', '450203000000', '鱼峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('450204', '4502', '2', '柳南', 'l', 'liu nan', '450204000000', '柳南区', NULL);
INSERT INTO `demo_city_code` VALUES ('450205', '4502', '2', '柳北', 'l', 'liu bei', '450205000000', '柳北区', NULL);
INSERT INTO `demo_city_code` VALUES ('450206', '4502', '2', '柳江', 'l', 'liu jiang', '450206000000', '柳江区', '101300305');
INSERT INTO `demo_city_code` VALUES ('450222', '4502', '2', '柳城', 'l', 'liu cheng', '450222000000', '柳城县', '101300302');
INSERT INTO `demo_city_code` VALUES ('450223', '4502', '2', '鹿寨', 'l', 'lu zhai', '450223000000', '鹿寨县', '101300304');
INSERT INTO `demo_city_code` VALUES ('450224', '4502', '2', '融安', 'r', 'rong an', '450224000000', '融安县', '101300306');
INSERT INTO `demo_city_code` VALUES ('450225', '4502', '2', '融水', 'r', 'rong shui', '450225000000', '融水苗族自治县', '101300307');
INSERT INTO `demo_city_code` VALUES ('450226', '4502', '2', '三江', 's', 'san jiang', '450226000000', '三江侗族自治县', '101300308');
INSERT INTO `demo_city_code` VALUES ('4503', '45', '1', '桂林', 'g', 'gui lin', '450300000000', '桂林市', '101300501');
INSERT INTO `demo_city_code` VALUES ('450302', '4503', '2', '秀峰', 'x', 'xiu feng', '450302000000', '秀峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('450303', '4503', '2', '叠彩', 'd', 'die cai', '450303000000', '叠彩区', NULL);
INSERT INTO `demo_city_code` VALUES ('450304', '4503', '2', '象山', 'x', 'xiang shan', '450304000000', '象山区', '101210406');
INSERT INTO `demo_city_code` VALUES ('450305', '4503', '2', '七星', 'q', 'qi xing', '450305000000', '七星区', NULL);
INSERT INTO `demo_city_code` VALUES ('450311', '4503', '2', '雁山', 'y', 'yan shan', '450311000000', '雁山区', NULL);
INSERT INTO `demo_city_code` VALUES ('450312', '4503', '2', '临桂', 'l', 'lin gui', '450312000000', '临桂区', '101300505');
INSERT INTO `demo_city_code` VALUES ('450321', '4503', '2', '阳朔', 'y', 'yang shuo', '450321000000', '阳朔县', '101300510');
INSERT INTO `demo_city_code` VALUES ('450323', '4503', '2', '灵川', 'l', 'ling chuan', '450323000000', '灵川县', '101300507');
INSERT INTO `demo_city_code` VALUES ('450324', '4503', '2', '全州', 'q', 'quan zhou', '450324000000', '全州县', '101300508');
INSERT INTO `demo_city_code` VALUES ('450325', '4503', '2', '兴安', 'x', 'xing an', '450325000000', '兴安县', '101300506');
INSERT INTO `demo_city_code` VALUES ('450326', '4503', '2', '永福', 'y', 'yong fu', '450326000000', '永福县', '101300504');
INSERT INTO `demo_city_code` VALUES ('450327', '4503', '2', '灌阳', 'g', 'guan yang', '450327000000', '灌阳县', '101300509');
INSERT INTO `demo_city_code` VALUES ('450328', '4503', '2', '龙胜', 'l', 'long sheng', '450328000000', '龙胜各族自治县', '101300503');
INSERT INTO `demo_city_code` VALUES ('450329', '4503', '2', '资源', 'z', 'zi yuan', '450329000000', '资源县', '101300514');
INSERT INTO `demo_city_code` VALUES ('450330', '4503', '2', '平乐', 'p', 'ping le', '450330000000', '平乐县', '101300512');
INSERT INTO `demo_city_code` VALUES ('450332', '4503', '2', '恭城', 'g', 'gong cheng', '450332000000', '恭城瑶族自治县', '101300511');
INSERT INTO `demo_city_code` VALUES ('450381', '4503', '2', '荔浦', 'l', 'li pu', '450381000000', '荔浦市', '101300513');
INSERT INTO `demo_city_code` VALUES ('4504', '45', '1', '梧州', 'w', 'wu zhou', '450400000000', '梧州市', '101300601');
INSERT INTO `demo_city_code` VALUES ('450403', '4504', '2', '万秀', 'w', 'wan xiu', '450403000000', '万秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('450405', '4504', '2', '长洲', 'c', 'chang zhou', '450405000000', '长洲区', NULL);
INSERT INTO `demo_city_code` VALUES ('450406', '4504', '2', '龙圩', 'l', 'long xu', '450406000000', '龙圩区', NULL);
INSERT INTO `demo_city_code` VALUES ('450421', '4504', '2', '苍梧', 'c', 'cang wu', '450421000000', '苍梧县', '101300604');
INSERT INTO `demo_city_code` VALUES ('450422', '4504', '2', '藤县', 't', 'teng xian', '450422000000', '藤县', '101300602');
INSERT INTO `demo_city_code` VALUES ('450423', '4504', '2', '蒙山', 'm', 'meng shan', '450423000000', '蒙山县', '101300605');
INSERT INTO `demo_city_code` VALUES ('450481', '4504', '2', '岑溪', 'c', 'cen xi', '450481000000', '岑溪市', '101300606');
INSERT INTO `demo_city_code` VALUES ('4505', '45', '1', '北海', 'b', 'bei hai', '450500000000', '北海市', '101301301');
INSERT INTO `demo_city_code` VALUES ('450502', '4505', '2', '海城', 'h', 'hai cheng', '450502000000', '海城区', '101070304');
INSERT INTO `demo_city_code` VALUES ('450503', '4505', '2', '银海', 'y', 'yin hai', '450503000000', '银海区', NULL);
INSERT INTO `demo_city_code` VALUES ('450512', '4505', '2', '铁山港', 't', 'tie shan gang', '450512000000', '铁山港区', NULL);
INSERT INTO `demo_city_code` VALUES ('450521', '4505', '2', '合浦', 'h', 'he pu', '450521000000', '合浦县', '101301302');
INSERT INTO `demo_city_code` VALUES ('4506', '45', '1', '防城港', 'f', 'fang cheng gang', '450600000000', '防城港市', '101301401');
INSERT INTO `demo_city_code` VALUES ('450602', '4506', '2', '港口', 'g', 'gang kou', '450602000000', '港口区', NULL);
INSERT INTO `demo_city_code` VALUES ('450603', '4506', '2', '防城', 'f', 'fang cheng', '450603000000', '防城区', '101301405');
INSERT INTO `demo_city_code` VALUES ('450621', '4506', '2', '上思', 's', 'shang si', '450621000000', '上思县', '101301402');
INSERT INTO `demo_city_code` VALUES ('450681', '4506', '2', '东兴', 'd', 'dong xing', '450681000000', '东兴市', '101271202');
INSERT INTO `demo_city_code` VALUES ('4507', '45', '1', '钦州', 'q', 'qin zhou', '450700000000', '钦州市', '101301101');
INSERT INTO `demo_city_code` VALUES ('450702', '4507', '2', '钦南', 'q', 'qin nan', '450702000000', '钦南区', NULL);
INSERT INTO `demo_city_code` VALUES ('450703', '4507', '2', '钦北', 'q', 'qin bei', '450703000000', '钦北区', NULL);
INSERT INTO `demo_city_code` VALUES ('450721', '4507', '2', '灵山', 'l', 'ling shan', '450721000000', '灵山县', '101301103');
INSERT INTO `demo_city_code` VALUES ('450722', '4507', '2', '浦北', 'p', 'pu bei', '450722000000', '浦北县', '101301102');
INSERT INTO `demo_city_code` VALUES ('4508', '45', '1', '贵港', 'g', 'gui gang', '450800000000', '贵港市', '101300801');
INSERT INTO `demo_city_code` VALUES ('450802', '4508', '2', '港北', 'g', 'gang bei', '450802000000', '港北区', NULL);
INSERT INTO `demo_city_code` VALUES ('450803', '4508', '2', '港南', 'g', 'gang nan', '450803000000', '港南区', NULL);
INSERT INTO `demo_city_code` VALUES ('450804', '4508', '2', '覃塘', 'q', 'qin tang', '450804000000', '覃塘区', NULL);
INSERT INTO `demo_city_code` VALUES ('450821', '4508', '2', '平南', 'p', 'ping nan', '450821000000', '平南县', '101300803');
INSERT INTO `demo_city_code` VALUES ('450881', '4508', '2', '桂平', 'g', 'gui ping', '450881000000', '桂平市', '101300802');
INSERT INTO `demo_city_code` VALUES ('4509', '45', '1', '玉林', 'y', 'yu lin', '450900000000', '玉林市', '101300901');
INSERT INTO `demo_city_code` VALUES ('450902', '4509', '2', '玉州', 'y', 'yu zhou', '450902000000', '玉州区', NULL);
INSERT INTO `demo_city_code` VALUES ('450903', '4509', '2', '福绵', 'f', 'fu mian', '450903000000', '福绵区', NULL);
INSERT INTO `demo_city_code` VALUES ('450921', '4509', '2', '容县', 'r', 'rong xian', '450921000000', '容县', '101300904');
INSERT INTO `demo_city_code` VALUES ('450922', '4509', '2', '陆川', 'l', 'lu chuan', '450922000000', '陆川县', '101300905');
INSERT INTO `demo_city_code` VALUES ('450923', '4509', '2', '博白', 'b', 'bo bai', '450923000000', '博白县', '101300902');
INSERT INTO `demo_city_code` VALUES ('450924', '4509', '2', '兴业', 'x', 'xing ye', '450924000000', '兴业县', '101300906');
INSERT INTO `demo_city_code` VALUES ('450981', '4509', '2', '北流', 'b', 'bei liu', '450981000000', '北流市', '101300903');
INSERT INTO `demo_city_code` VALUES ('4510', '45', '1', '百色', 'b', 'bai se', '451000000000', '百色市', '101301001');
INSERT INTO `demo_city_code` VALUES ('451002', '4510', '2', '右江', 'y', 'you jiang', '451002000000', '右江区', NULL);
INSERT INTO `demo_city_code` VALUES ('451003', '4510', '2', '田阳', 't', 'tian yang', '451003000000', '田阳区', '101301003');
INSERT INTO `demo_city_code` VALUES ('451022', '4510', '2', '田东', 't', 'tian dong', '451022000000', '田东县', '101301006');
INSERT INTO `demo_city_code` VALUES ('451024', '4510', '2', '德保', 'd', 'de bao', '451024000000', '德保县', '101301004');
INSERT INTO `demo_city_code` VALUES ('451026', '4510', '2', '那坡', 'n', 'na po', '451026000000', '那坡县', '101301002');
INSERT INTO `demo_city_code` VALUES ('451027', '4510', '2', '凌云', 'l', 'ling yun', '451027000000', '凌云县', '101301011');
INSERT INTO `demo_city_code` VALUES ('451028', '4510', '2', '乐业', 'l', 'le ye', '451028000000', '乐业县', '101301010');
INSERT INTO `demo_city_code` VALUES ('451029', '4510', '2', '田林', 't', 'tian lin', '451029000000', '田林县', '101301012');
INSERT INTO `demo_city_code` VALUES ('451030', '4510', '2', '西林', 'x', 'xi lin', '451030000000', '西林县', '101301009');
INSERT INTO `demo_city_code` VALUES ('451031', '4510', '2', '隆林', 'l', 'long lin', '451031000000', '隆林各族自治县', '101301008');
INSERT INTO `demo_city_code` VALUES ('451081', '4510', '2', '靖西', 'j', 'jing xi', '451081000000', '靖西市', '101301005');
INSERT INTO `demo_city_code` VALUES ('451082', '4510', '2', '平果', 'p', 'ping guo', '451082000000', '平果市', '101301007');
INSERT INTO `demo_city_code` VALUES ('4511', '45', '1', '贺州', 'h', 'he zhou', '451100000000', '贺州市', '101300701');
INSERT INTO `demo_city_code` VALUES ('451102', '4511', '2', '八步', 'b', 'ba bu', '451102000000', '八步区', NULL);
INSERT INTO `demo_city_code` VALUES ('451103', '4511', '2', '平桂', 'p', 'ping gui', '451103000000', '平桂区', NULL);
INSERT INTO `demo_city_code` VALUES ('451121', '4511', '2', '昭平', 'z', 'zhao ping', '451121000000', '昭平县', '101300702');
INSERT INTO `demo_city_code` VALUES ('451122', '4511', '2', '钟山', 'z', 'zhong shan', '451122000000', '钟山县', '101300704');
INSERT INTO `demo_city_code` VALUES ('451123', '4511', '2', '富川', 'f', 'fu chuan', '451123000000', '富川瑶族自治县', '101300703');
INSERT INTO `demo_city_code` VALUES ('4512', '45', '1', '河池', 'h', 'he chi', '451200000000', '河池市', '101301201');
INSERT INTO `demo_city_code` VALUES ('451202', '4512', '2', '金城江', 'j', 'jin cheng jiang', '451202000000', '金城江区', NULL);
INSERT INTO `demo_city_code` VALUES ('451203', '4512', '2', '宜州', 'y', 'yi zhou', '451203000000', '宜州区', '101301207');
INSERT INTO `demo_city_code` VALUES ('451221', '4512', '2', '南丹', 'n', 'nan dan', '451221000000', '南丹县', '101301209');
INSERT INTO `demo_city_code` VALUES ('451222', '4512', '2', '天峨', 't', 'tian e', '451222000000', '天峨县', '101301202');
INSERT INTO `demo_city_code` VALUES ('451223', '4512', '2', '凤山', 'f', 'feng shan', '451223000000', '凤山县', '101301208');
INSERT INTO `demo_city_code` VALUES ('451224', '4512', '2', '东兰', 'd', 'dong lan', '451224000000', '东兰县', '101301203');
INSERT INTO `demo_city_code` VALUES ('451225', '4512', '2', '罗城', 'l', 'luo cheng', '451225000000', '罗城仫佬族自治县', '101301206');
INSERT INTO `demo_city_code` VALUES ('451226', '4512', '2', '环江', 'h', 'huan jiang', '451226000000', '环江毛南族自治县', '101301205');
INSERT INTO `demo_city_code` VALUES ('451227', '4512', '2', '巴马', 'b', 'ba ma', '451227000000', '巴马瑶族自治县', '101301204');
INSERT INTO `demo_city_code` VALUES ('451228', '4512', '2', '都安', 'd', 'du an', '451228000000', '都安瑶族自治县', '101301210');
INSERT INTO `demo_city_code` VALUES ('451229', '4512', '2', '大化', 'd', 'da hua', '451229000000', '大化瑶族自治县', '101301211');
INSERT INTO `demo_city_code` VALUES ('4513', '45', '1', '来宾', 'l', 'lai bin', '451300000000', '来宾市', '101300401');
INSERT INTO `demo_city_code` VALUES ('451302', '4513', '2', '兴宾', 'x', 'xing bin', '451302000000', '兴宾区', NULL);
INSERT INTO `demo_city_code` VALUES ('451321', '4513', '2', '忻城', 'x', 'xin cheng', '451321000000', '忻城县', '101300402');
INSERT INTO `demo_city_code` VALUES ('451322', '4513', '2', '象州', 'x', 'xiang zhou', '451322000000', '象州县', '101300404');
INSERT INTO `demo_city_code` VALUES ('451323', '4513', '2', '武宣', 'w', 'wu xuan', '451323000000', '武宣县', '101300405');
INSERT INTO `demo_city_code` VALUES ('451324', '4513', '2', '金秀', 'j', 'jin xiu', '451324000000', '金秀瑶族自治县', '101300403');
INSERT INTO `demo_city_code` VALUES ('451381', '4513', '2', '合山', 'h', 'he shan', '451381000000', '合山市', '101300406');
INSERT INTO `demo_city_code` VALUES ('4514', '45', '1', '崇左', 'c', 'chong zuo', '451400000000', '崇左市', '101300201');
INSERT INTO `demo_city_code` VALUES ('451402', '4514', '2', '江州', 'j', 'jiang zhou', '451402000000', '江州区', NULL);
INSERT INTO `demo_city_code` VALUES ('451421', '4514', '2', '扶绥', 'f', 'fu sui', '451421000000', '扶绥县', '101300206');
INSERT INTO `demo_city_code` VALUES ('451422', '4514', '2', '宁明', 'n', 'ning ming', '451422000000', '宁明县', '101300207');
INSERT INTO `demo_city_code` VALUES ('451423', '4514', '2', '龙州', 'l', 'long zhou', '451423000000', '龙州县', '101300203');
INSERT INTO `demo_city_code` VALUES ('451424', '4514', '2', '大新', 'd', 'da xin', '451424000000', '大新县', '101300205');
INSERT INTO `demo_city_code` VALUES ('451425', '4514', '2', '天等', 't', 'tian deng', '451425000000', '天等县', '101300202');
INSERT INTO `demo_city_code` VALUES ('451481', '4514', '2', '凭祥', 'p', 'ping xiang', '451481000000', '凭祥市', '101300204');
INSERT INTO `demo_city_code` VALUES ('46', '0', '0', '海南', 'h', 'hai nan', '460000000000', '海南省', '101150401');
INSERT INTO `demo_city_code` VALUES ('4601', '46', '1', '海口', 'h', 'hai kou', '460100000000', '海口市', '101310101');
INSERT INTO `demo_city_code` VALUES ('460105', '4601', '2', '秀英', 'x', 'xiu ying', '460105000000', '秀英区', NULL);
INSERT INTO `demo_city_code` VALUES ('460106', '4601', '2', '龙华', 'l', 'long hua', '460106000000', '龙华区', NULL);
INSERT INTO `demo_city_code` VALUES ('460107', '4601', '2', '琼山', 'q', 'qiong shan', '460107000000', '琼山区', NULL);
INSERT INTO `demo_city_code` VALUES ('460108', '4601', '2', '美兰', 'm', 'mei lan', '460108000000', '美兰区', NULL);
INSERT INTO `demo_city_code` VALUES ('4602', '46', '1', '三亚', 's', 'san ya', '460200000000', '三亚市', '101310201');
INSERT INTO `demo_city_code` VALUES ('460202', '4602', '2', '海棠', 'h', 'hai tang', '460202000000', '海棠区', NULL);
INSERT INTO `demo_city_code` VALUES ('460203', '4602', '2', '吉阳', 'j', 'ji yang', '460203000000', '吉阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('460204', '4602', '2', '天涯', 't', 'tian ya', '460204000000', '天涯区', NULL);
INSERT INTO `demo_city_code` VALUES ('460205', '4602', '2', '崖州', 'y', 'ya zhou', '460205000000', '崖州区', NULL);
INSERT INTO `demo_city_code` VALUES ('4603', '46', '1', '三沙', 's', 'san sha', '460300000000', '三沙市', NULL);
INSERT INTO `demo_city_code` VALUES ('460301', '4603', '2', '西沙', 'x', 'xi sha', '460301000000', '西沙区', '101310217');
INSERT INTO `demo_city_code` VALUES ('460302', '4603', '2', '南沙', 'n', 'nan sha', '460302000000', '南沙区', '101310220');
INSERT INTO `demo_city_code` VALUES ('4604', '46', '1', '儋州', 'd', 'dan zhou', '460400000000', '儋州市', '101310205');
INSERT INTO `demo_city_code` VALUES ('460400', '4604', '2', '儋州', 'd', 'dan zhou', '460400000000', '儋州市', '101310205');
INSERT INTO `demo_city_code` VALUES ('469001', '46', '1', '五指山', 'w', 'wu zhi shan', '469001000000', '五指山市', '101310222');
INSERT INTO `demo_city_code` VALUES ('469001000', '469001', '2', '五指山', 'w', 'wu zhi shan', '469001000000', '五指山市', '101310222');
INSERT INTO `demo_city_code` VALUES ('469002', '46', '1', '琼海', 'q', 'qiong hai', '469002000000', '琼海市', '101310211');
INSERT INTO `demo_city_code` VALUES ('469002000', '469002', '2', '琼海', 'q', 'qiong hai', '469002000000', '琼海市', '101310211');
INSERT INTO `demo_city_code` VALUES ('469005', '46', '1', '文昌', 'w', 'wen chang', '469005000000', '文昌市', '101310212');
INSERT INTO `demo_city_code` VALUES ('469005000', '469005', '2', '文昌', 'w', 'wen chang', '469005000000', '文昌市', '101310212');
INSERT INTO `demo_city_code` VALUES ('469006', '46', '1', '万宁', 'w', 'wan ning', '469006000000', '万宁市', '101310215');
INSERT INTO `demo_city_code` VALUES ('469006000', '469006', '2', '万宁', 'w', 'wan ning', '469006000000', '万宁市', '101310215');
INSERT INTO `demo_city_code` VALUES ('469007', '46', '1', '东方', 'd', 'dong fang', '469007000000', '东方市', '101310202');
INSERT INTO `demo_city_code` VALUES ('469007000', '469007', '2', '东方', 'd', 'dong fang', '469007000000', '东方市', '101310202');
INSERT INTO `demo_city_code` VALUES ('469021', '46', '1', '定安', 'd', 'ding an', '469021000000', '定安县', '101310209');
INSERT INTO `demo_city_code` VALUES ('469021000', '469021', '2', '定安', 'd', 'ding an', '469021000000', '定安县', '101310209');
INSERT INTO `demo_city_code` VALUES ('469022', '46', '1', '屯昌', 't', 'tun chang', '469022000000', '屯昌县', '101310210');
INSERT INTO `demo_city_code` VALUES ('469022000', '469022', '2', '屯昌', 't', 'tun chang', '469022000000', '屯昌县', '101310210');
INSERT INTO `demo_city_code` VALUES ('469023', '46', '1', '澄迈', 'c', 'cheng mai', '469023000000', '澄迈县', '101310204');
INSERT INTO `demo_city_code` VALUES ('469023000', '469023', '2', '澄迈', 'c', 'cheng mai', '469023000000', '澄迈县', '101310204');
INSERT INTO `demo_city_code` VALUES ('469024', '46', '1', '临高', 'l', 'lin gao', '469024000000', '临高县', '101310203');
INSERT INTO `demo_city_code` VALUES ('469024000', '469024', '2', '临高', 'l', 'lin gao', '469024000000', '临高县', '101310203');
INSERT INTO `demo_city_code` VALUES ('469025', '46', '1', '白沙', 'b', 'bai sha', '469025000000', '白沙黎族自治县', '101310207');
INSERT INTO `demo_city_code` VALUES ('469025000', '469025', '2', '白沙', 'b', 'bai sha', '469025000000', '白沙黎族自治县', '101310207');
INSERT INTO `demo_city_code` VALUES ('469026', '46', '1', '昌江', 'c', 'chang jiang', '469026000000', '昌江黎族自治县', '101310206');
INSERT INTO `demo_city_code` VALUES ('469026000', '469026', '2', '昌江', 'c', 'chang jiang', '469026000000', '昌江黎族自治县', '101310206');
INSERT INTO `demo_city_code` VALUES ('469027', '46', '1', '乐东', 'l', 'le dong', '469027000000', '乐东黎族自治县', '101310221');
INSERT INTO `demo_city_code` VALUES ('469027000', '469027', '2', '乐东', 'l', 'le dong', '469027000000', '乐东黎族自治县', '101310221');
INSERT INTO `demo_city_code` VALUES ('469028', '46', '1', '陵水', 'l', 'ling shui', '469028000000', '陵水黎族自治县', '101310216');
INSERT INTO `demo_city_code` VALUES ('469028000', '469028', '2', '陵水', 'l', 'ling shui', '469028000000', '陵水黎族自治县', '101310216');
INSERT INTO `demo_city_code` VALUES ('469029', '46', '1', '保亭', 'b', 'bao ting', '469029000000', '保亭黎族苗族自治县', '101310214');
INSERT INTO `demo_city_code` VALUES ('469029000', '469029', '2', '保亭', 'b', 'bao ting', '469029000000', '保亭黎族苗族自治县', '101310214');
INSERT INTO `demo_city_code` VALUES ('469030', '46', '1', '琼中', 'q', 'qiong zhong', '469030000000', '琼中黎族苗族自治县', '101310208');
INSERT INTO `demo_city_code` VALUES ('469030000', '469030', '2', '琼中', 'q', 'qiong zhong', '469030000000', '琼中黎族苗族自治县', '101310208');
INSERT INTO `demo_city_code` VALUES ('50', '0', '0', '重庆', 'c', 'chong qing', '500000000000', '重庆市', '101040100');
INSERT INTO `demo_city_code` VALUES ('5001', '50', '1', '重庆城区', 'c', 'chong qing cheng qu', '500100000000', '重庆城区', NULL);
INSERT INTO `demo_city_code` VALUES ('500101', '5001', '2', '万州', 'w', 'wan zhou', '500101000000', '万州区', '101041300');
INSERT INTO `demo_city_code` VALUES ('500102', '5001', '2', '涪陵', 'f', 'fu ling', '500102000000', '涪陵区', '101041400');
INSERT INTO `demo_city_code` VALUES ('500103', '5001', '2', '渝中', 'y', 'yu zhong', '500103000000', '渝中区', NULL);
INSERT INTO `demo_city_code` VALUES ('500104', '5001', '2', '大渡口', 'd', 'da du kou', '500104000000', '大渡口区', NULL);
INSERT INTO `demo_city_code` VALUES ('500105', '5001', '2', '江北', 'j', 'jiang bei', '500105000000', '江北区', NULL);
INSERT INTO `demo_city_code` VALUES ('500106', '5001', '2', '沙坪坝', 's', 'sha ping ba', '500106000000', '沙坪坝区', NULL);
INSERT INTO `demo_city_code` VALUES ('500107', '5001', '2', '九龙坡', 'j', 'jiu long po', '500107000000', '九龙坡区', NULL);
INSERT INTO `demo_city_code` VALUES ('500108', '5001', '2', '南岸', 'n', 'nan an', '500108000000', '南岸区', NULL);
INSERT INTO `demo_city_code` VALUES ('500109', '5001', '2', '北碚', 'b', 'bei bei', '500109000000', '北碚区', '101040800');
INSERT INTO `demo_city_code` VALUES ('500110', '5001', '2', '綦江', 'q', 'qi jiang', '500110000000', '綦江区', '101043300');
INSERT INTO `demo_city_code` VALUES ('500111', '5001', '2', '大足', 'd', 'da zu', '500111000000', '大足区', '101042600');
INSERT INTO `demo_city_code` VALUES ('500112', '5001', '2', '渝北', 'y', 'yu bei', '500112000000', '渝北区', '101040700');
INSERT INTO `demo_city_code` VALUES ('500113', '5001', '2', '巴南', 'b', 'ba nan', '500113000000', '巴南区', '101040900');
INSERT INTO `demo_city_code` VALUES ('500114', '5001', '2', '黔江', 'q', 'qian jiang', '500114000000', '黔江区', '101041100');
INSERT INTO `demo_city_code` VALUES ('500115', '5001', '2', '长寿', 'c', 'chang shou', '500115000000', '长寿区', '101041000');
INSERT INTO `demo_city_code` VALUES ('500116', '5001', '2', '江津', 'j', 'jiang jin', '500116000000', '江津区', '101040500');
INSERT INTO `demo_city_code` VALUES ('500117', '5001', '2', '合川', 'h', 'he chuan', '500117000000', '合川区', '101040300');
INSERT INTO `demo_city_code` VALUES ('500118', '5001', '2', '永川', 'y', 'yong chuan', '500118000000', '永川区', '101040200');
INSERT INTO `demo_city_code` VALUES ('500119', '5001', '2', '南川', 'n', 'nan chuan', '500119000000', '南川区', '101040400');
INSERT INTO `demo_city_code` VALUES ('500120', '5001', '2', '璧山', 'b', 'bi shan', '500120000000', '璧山区', '101042900');
INSERT INTO `demo_city_code` VALUES ('500151', '5001', '2', '铜梁', 't', 'tong liang', '500151000000', '铜梁区', '101042800');
INSERT INTO `demo_city_code` VALUES ('500152', '5001', '2', '潼南', 't', 'tong nan', '500152000000', '潼南区', '101042100');
INSERT INTO `demo_city_code` VALUES ('500153', '5001', '2', '荣昌', 'r', 'rong chang', '500153000000', '荣昌区', '101042700');
INSERT INTO `demo_city_code` VALUES ('500154', '5001', '2', '开州', 'k', 'kai zhou', '500154000000', '开州区', NULL);
INSERT INTO `demo_city_code` VALUES ('500155', '5001', '2', '梁平', 'l', 'liang ping', '500155000000', '梁平区', '101042300');
INSERT INTO `demo_city_code` VALUES ('500156', '5001', '2', '武隆', 'w', 'wu long', '500156000000', '武隆区', '101043100');
INSERT INTO `demo_city_code` VALUES ('5002', '50', '1', '重庆郊县', 'c', 'chong qing jiao xian', '500200000000', '重庆郊县', NULL);
INSERT INTO `demo_city_code` VALUES ('500229', '5002', '2', '城口', 'c', 'cheng kou', '500229000000', '城口县', '101041600');
INSERT INTO `demo_city_code` VALUES ('500230', '5002', '2', '丰都', 'f', 'feng du', '500230000000', '丰都县', '101043000');
INSERT INTO `demo_city_code` VALUES ('500231', '5002', '2', '垫江', 'd', 'dian jiang', '500231000000', '垫江县', '101042200');
INSERT INTO `demo_city_code` VALUES ('500233', '5002', '2', '忠县', 'z', 'zhong xian', '500233000000', '忠县', '101042400');
INSERT INTO `demo_city_code` VALUES ('500235', '5002', '2', '云阳', 'y', 'yun yang', '500235000000', '云阳县', '101041700');
INSERT INTO `demo_city_code` VALUES ('500236', '5002', '2', '奉节', 'f', 'feng jie', '500236000000', '奉节县', '101041900');
INSERT INTO `demo_city_code` VALUES ('500237', '5002', '2', '巫山', 'w', 'wu shan', '500237000000', '巫山县', '101042000');
INSERT INTO `demo_city_code` VALUES ('500238', '5002', '2', '巫溪', 'w', 'wu xi', '500238000000', '巫溪县', '101041800');
INSERT INTO `demo_city_code` VALUES ('500240', '5002', '2', '石柱', 's', 'shi zhu', '500240000000', '石柱土家族自治县', '101042500');
INSERT INTO `demo_city_code` VALUES ('500241', '5002', '2', '秀山', 'x', 'xiu shan', '500241000000', '秀山土家族苗族自治县', '101043600');
INSERT INTO `demo_city_code` VALUES ('500242', '5002', '2', '酉阳', 'y', 'you yang', '500242000000', '酉阳土家族苗族自治县', '101043400');
INSERT INTO `demo_city_code` VALUES ('500243', '5002', '2', '彭水', 'p', 'peng shui', '500243000000', '彭水苗族土家族自治县', '101043200');
INSERT INTO `demo_city_code` VALUES ('51', '0', '0', '四川', 's', 'si chuan', '510000000000', '四川省', NULL);
INSERT INTO `demo_city_code` VALUES ('5101', '51', '1', '成都', 'c', 'cheng du', '510100000000', '成都市', '101270101');
INSERT INTO `demo_city_code` VALUES ('510104', '5101', '2', '锦江', 'j', 'jin jiang', '510104000000', '锦江区', NULL);
INSERT INTO `demo_city_code` VALUES ('510105', '5101', '2', '青羊', 'q', 'qing yang', '510105000000', '青羊区', NULL);
INSERT INTO `demo_city_code` VALUES ('510106', '5101', '2', '金牛', 'j', 'jin niu', '510106000000', '金牛区', NULL);
INSERT INTO `demo_city_code` VALUES ('510107', '5101', '2', '武侯', 'w', 'wu hou', '510107000000', '武侯区', NULL);
INSERT INTO `demo_city_code` VALUES ('510108', '5101', '2', '成华', 'c', 'cheng hua', '510108000000', '成华区', NULL);
INSERT INTO `demo_city_code` VALUES ('510112', '5101', '2', '龙泉驿', 'l', 'long quan yi', '510112000000', '龙泉驿区', '101270102');
INSERT INTO `demo_city_code` VALUES ('510113', '5101', '2', '青白江', 'q', 'qing bai jiang', '510113000000', '青白江区', NULL);
INSERT INTO `demo_city_code` VALUES ('510114', '5101', '2', '新都', 'x', 'xin du', '510114000000', '新都区', '101270103');
INSERT INTO `demo_city_code` VALUES ('510115', '5101', '2', '温江', 'w', 'wen jiang', '510115000000', '温江区', '101270104');
INSERT INTO `demo_city_code` VALUES ('510116', '5101', '2', '双流', 's', 'shuang liu', '510116000000', '双流区', '101270106');
INSERT INTO `demo_city_code` VALUES ('510117', '5101', '2', '郫都', 'p', 'pi du', '510117000000', '郫都区', NULL);
INSERT INTO `demo_city_code` VALUES ('510118', '5101', '2', '新津', 'x', 'xin jin', '510118000000', '新津区', '101270110');
INSERT INTO `demo_city_code` VALUES ('510121', '5101', '2', '金堂', 'j', 'jin tang', '510121000000', '金堂县', '101270105');
INSERT INTO `demo_city_code` VALUES ('510129', '5101', '2', '大邑', 'd', 'da yi', '510129000000', '大邑县', '101270108');
INSERT INTO `demo_city_code` VALUES ('510131', '5101', '2', '蒲江', 'p', 'pu jiang', '510131000000', '蒲江县', '101270109');
INSERT INTO `demo_city_code` VALUES ('510181', '5101', '2', '都江堰', 'd', 'du jiang yan', '510181000000', '都江堰市', '101270111');
INSERT INTO `demo_city_code` VALUES ('510182', '5101', '2', '彭州', 'p', 'peng zhou', '510182000000', '彭州市', '101270112');
INSERT INTO `demo_city_code` VALUES ('510183', '5101', '2', '邛崃', 'q', 'qiong lai', '510183000000', '邛崃市', '101270113');
INSERT INTO `demo_city_code` VALUES ('510184', '5101', '2', '崇州', 'c', 'chong zhou', '510184000000', '崇州市', '101270114');
INSERT INTO `demo_city_code` VALUES ('510185', '5101', '2', '简阳', 'j', 'jian yang', '510185000000', '简阳市', '101271304');
INSERT INTO `demo_city_code` VALUES ('5103', '51', '1', '自贡', 'z', 'zi gong', '510300000000', '自贡市', '101270301');
INSERT INTO `demo_city_code` VALUES ('510302', '5103', '2', '自流井', 'z', 'zi liu jing', '510302000000', '自流井区', NULL);
INSERT INTO `demo_city_code` VALUES ('510303', '5103', '2', '贡井', 'g', 'gong jing', '510303000000', '贡井区', NULL);
INSERT INTO `demo_city_code` VALUES ('510304', '5103', '2', '大安', 'd', 'da an', '510304000000', '大安区', '101060603');
INSERT INTO `demo_city_code` VALUES ('510311', '5103', '2', '沿滩', 'y', 'yan tan', '510311000000', '沿滩区', NULL);
INSERT INTO `demo_city_code` VALUES ('510321', '5103', '2', '荣县', 'r', 'rong xian', '510321000000', '荣县', '101270303');
INSERT INTO `demo_city_code` VALUES ('510322', '5103', '2', '富顺', 'f', 'fu shun', '510322000000', '富顺县', '101270302');
INSERT INTO `demo_city_code` VALUES ('5104', '51', '1', '攀枝花', 'p', 'pan zhi hua', '510400000000', '攀枝花市', '101270201');
INSERT INTO `demo_city_code` VALUES ('510402', '5104', '2', '东区', 'd', 'dong qu', '510402000000', '东区', NULL);
INSERT INTO `demo_city_code` VALUES ('510403', '5104', '2', '西区', 'x', 'xi qu', '510403000000', '西区', NULL);
INSERT INTO `demo_city_code` VALUES ('510411', '5104', '2', '仁和', 'r', 'ren he', '510411000000', '仁和区', '101270202');
INSERT INTO `demo_city_code` VALUES ('510421', '5104', '2', '米易', 'm', 'mi yi', '510421000000', '米易县', '101270203');
INSERT INTO `demo_city_code` VALUES ('510422', '5104', '2', '盐边', 'y', 'yan bian', '510422000000', '盐边县', '101270204');
INSERT INTO `demo_city_code` VALUES ('5105', '51', '1', '泸州', 'l', 'lu zhou', '510500000000', '泸州市', '101271001');
INSERT INTO `demo_city_code` VALUES ('510502', '5105', '2', '江阳', 'j', 'jiang yang', '510502000000', '江阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('510503', '5105', '2', '纳溪', 'n', 'na xi', '510503000000', '纳溪区', '101271007');
INSERT INTO `demo_city_code` VALUES ('510504', '5105', '2', '龙马潭', 'l', 'long ma tan', '510504000000', '龙马潭区', NULL);
INSERT INTO `demo_city_code` VALUES ('510521', '5105', '2', '泸县', 'l', 'lu xian', '510521000000', '泸县', '101271003');
INSERT INTO `demo_city_code` VALUES ('510522', '5105', '2', '合江', 'h', 'he jiang', '510522000000', '合江县', '101271004');
INSERT INTO `demo_city_code` VALUES ('510524', '5105', '2', '叙永', 'x', 'xu yong', '510524000000', '叙永县', '101271005');
INSERT INTO `demo_city_code` VALUES ('510525', '5105', '2', '古蔺', 'g', 'gu lin', '510525000000', '古蔺县', '101271006');
INSERT INTO `demo_city_code` VALUES ('5106', '51', '1', '德阳', 'd', 'de yang', '510600000000', '德阳市', '101272001');
INSERT INTO `demo_city_code` VALUES ('510603', '5106', '2', '旌阳', 'j', 'jing yang', '510603000000', '旌阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('510604', '5106', '2', '罗江', 'l', 'luo jiang', '510604000000', '罗江区', '101272006');
INSERT INTO `demo_city_code` VALUES ('510623', '5106', '2', '中江', 'z', 'zhong jiang', '510623000000', '中江县', '101272002');
INSERT INTO `demo_city_code` VALUES ('510681', '5106', '2', '广汉', 'g', 'guang han', '510681000000', '广汉市', '101272003');
INSERT INTO `demo_city_code` VALUES ('510682', '5106', '2', '什邡', 's', 'shi fang', '510682000000', '什邡市', '101272004');
INSERT INTO `demo_city_code` VALUES ('510683', '5106', '2', '绵竹', 'm', 'mian zhu', '510683000000', '绵竹市', '101272005');
INSERT INTO `demo_city_code` VALUES ('5107', '51', '1', '绵阳', 'm', 'mian yang', '510700000000', '绵阳市', '101270401');
INSERT INTO `demo_city_code` VALUES ('510703', '5107', '2', '涪城', 'f', 'fu cheng', '510703000000', '涪城区', NULL);
INSERT INTO `demo_city_code` VALUES ('510704', '5107', '2', '游仙', 'y', 'you xian', '510704000000', '游仙区', NULL);
INSERT INTO `demo_city_code` VALUES ('510705', '5107', '2', '安州', 'a', 'an zhou', '510705000000', '安州区', NULL);
INSERT INTO `demo_city_code` VALUES ('510722', '5107', '2', '三台', 's', 'san tai', '510722000000', '三台县', '101270402');
INSERT INTO `demo_city_code` VALUES ('510723', '5107', '2', '盐亭', 'y', 'yan ting', '510723000000', '盐亭县', '101270403');
INSERT INTO `demo_city_code` VALUES ('510725', '5107', '2', '梓潼', 'z', 'zi tong', '510725000000', '梓潼县', '101270405');
INSERT INTO `demo_city_code` VALUES ('510726', '5107', '2', '北川', 'b', 'bei chuan', '510726000000', '北川羌族自治县', '101270406');
INSERT INTO `demo_city_code` VALUES ('510727', '5107', '2', '平武', 'p', 'ping wu', '510727000000', '平武县', '101270407');
INSERT INTO `demo_city_code` VALUES ('510781', '5107', '2', '江油', 'j', 'jiang you', '510781000000', '江油市', '101270408');
INSERT INTO `demo_city_code` VALUES ('5108', '51', '1', '广元', 'g', 'guang yuan', '510800000000', '广元市', '101272101');
INSERT INTO `demo_city_code` VALUES ('510802', '5108', '2', '利州', 'l', 'li zhou', '510802000000', '利州区', NULL);
INSERT INTO `demo_city_code` VALUES ('510811', '5108', '2', '昭化', 'z', 'zhao hua', '510811000000', '昭化区', NULL);
INSERT INTO `demo_city_code` VALUES ('510812', '5108', '2', '朝天', 'c', 'chao tian', '510812000000', '朝天区', NULL);
INSERT INTO `demo_city_code` VALUES ('510821', '5108', '2', '旺苍', 'w', 'wang cang', '510821000000', '旺苍县', '101272102');
INSERT INTO `demo_city_code` VALUES ('510822', '5108', '2', '青川', 'q', 'qing chuan', '510822000000', '青川县', '101272103');
INSERT INTO `demo_city_code` VALUES ('510823', '5108', '2', '剑阁', 'j', 'jian ge', '510823000000', '剑阁县', '101272104');
INSERT INTO `demo_city_code` VALUES ('510824', '5108', '2', '苍溪', 'c', 'cang xi', '510824000000', '苍溪县', '101272105');
INSERT INTO `demo_city_code` VALUES ('5109', '51', '1', '遂宁', 's', 'sui ning', '510900000000', '遂宁市', '101270701');
INSERT INTO `demo_city_code` VALUES ('510903', '5109', '2', '船山', 'c', 'chuan shan', '510903000000', '船山区', NULL);
INSERT INTO `demo_city_code` VALUES ('510904', '5109', '2', '安居', 'a', 'an ju', '510904000000', '安居区', NULL);
INSERT INTO `demo_city_code` VALUES ('510921', '5109', '2', '蓬溪', 'p', 'peng xi', '510921000000', '蓬溪县', '101270702');
INSERT INTO `demo_city_code` VALUES ('510923', '5109', '2', '大英', 'd', 'da ying', '510923000000', '大英县', NULL);
INSERT INTO `demo_city_code` VALUES ('510981', '5109', '2', '射洪', 's', 'she hong', '510981000000', '射洪市', '101270703');
INSERT INTO `demo_city_code` VALUES ('5110', '51', '1', '内江', 'n', 'nei jiang', '511000000000', '内江市', '101271201');
INSERT INTO `demo_city_code` VALUES ('511002', '5110', '2', '市中', 's', 'shi zhong', '511002000000', '市中区', NULL);
INSERT INTO `demo_city_code` VALUES ('511011', '5110', '2', '东兴', 'd', 'dong xing', '511011000000', '东兴区', '101271202');
INSERT INTO `demo_city_code` VALUES ('511024', '5110', '2', '威远', 'w', 'wei yuan', '511024000000', '威远县', '101271203');
INSERT INTO `demo_city_code` VALUES ('511025', '5110', '2', '资中', 'z', 'zi zhong', '511025000000', '资中县', '101271204');
INSERT INTO `demo_city_code` VALUES ('511083', '5110', '2', '隆昌', 'l', 'long chang', '511083000000', '隆昌市', '101271205');
INSERT INTO `demo_city_code` VALUES ('5111', '51', '1', '乐山', 'l', 'le shan', '511100000000', '乐山市', '101271401');
INSERT INTO `demo_city_code` VALUES ('511102', '5111', '2', '市中', 's', 'shi zhong', '511102000000', '市中区', NULL);
INSERT INTO `demo_city_code` VALUES ('511111', '5111', '2', '沙湾', 's', 'sha wan', '511111000000', '沙湾区', '101131107');
INSERT INTO `demo_city_code` VALUES ('511112', '5111', '2', '五通桥', 'w', 'wu tong qiao', '511112000000', '五通桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('511113', '5111', '2', '金口河', 'j', 'jin kou he', '511113000000', '金口河区', NULL);
INSERT INTO `demo_city_code` VALUES ('511123', '5111', '2', '犍为', 'q', 'qian wei', '511123000000', '犍为县', '101271402');
INSERT INTO `demo_city_code` VALUES ('511124', '5111', '2', '井研', 'j', 'jing yan', '511124000000', '井研县', '101271403');
INSERT INTO `demo_city_code` VALUES ('511126', '5111', '2', '夹江', 'j', 'jia jiang', '511126000000', '夹江县', '101271404');
INSERT INTO `demo_city_code` VALUES ('511129', '5111', '2', '沐川', 'm', 'mu chuan', '511129000000', '沐川县', '101271405');
INSERT INTO `demo_city_code` VALUES ('511132', '5111', '2', '峨边', 'e', 'e bian', '511132000000', '峨边彝族自治县', '101271406');
INSERT INTO `demo_city_code` VALUES ('511133', '5111', '2', '马边', 'm', 'ma bian', '511133000000', '马边彝族自治县', '101271407');
INSERT INTO `demo_city_code` VALUES ('511181', '5111', '2', '峨眉山', 'e', 'e mei shan', '511181000000', '峨眉山市', '101271409');
INSERT INTO `demo_city_code` VALUES ('5113', '51', '1', '南充', 'n', 'nan chong', '511300000000', '南充市', '101270501');
INSERT INTO `demo_city_code` VALUES ('511302', '5113', '2', '顺庆', 's', 'shun qing', '511302000000', '顺庆区', NULL);
INSERT INTO `demo_city_code` VALUES ('511303', '5113', '2', '高坪', 'g', 'gao ping', '511303000000', '高坪区', NULL);
INSERT INTO `demo_city_code` VALUES ('511304', '5113', '2', '嘉陵', 'j', 'jia ling', '511304000000', '嘉陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('511321', '5113', '2', '南部', 'n', 'nan bu', '511321000000', '南部县', '101270502');
INSERT INTO `demo_city_code` VALUES ('511322', '5113', '2', '营山', 'y', 'ying shan', '511322000000', '营山县', '101270503');
INSERT INTO `demo_city_code` VALUES ('511323', '5113', '2', '蓬安', 'p', 'peng an', '511323000000', '蓬安县', '101270504');
INSERT INTO `demo_city_code` VALUES ('511324', '5113', '2', '仪陇', 'y', 'yi long', '511324000000', '仪陇县', '101270505');
INSERT INTO `demo_city_code` VALUES ('511325', '5113', '2', '西充', 'x', 'xi chong', '511325000000', '西充县', '101270506');
INSERT INTO `demo_city_code` VALUES ('511381', '5113', '2', '阆中', 'l', 'lang zhong', '511381000000', '阆中市', '101270507');
INSERT INTO `demo_city_code` VALUES ('5114', '51', '1', '眉山', 'm', 'mei shan', '511400000000', '眉山市', '101271501');
INSERT INTO `demo_city_code` VALUES ('511402', '5114', '2', '东坡', 'd', 'dong po', '511402000000', '东坡区', NULL);
INSERT INTO `demo_city_code` VALUES ('511403', '5114', '2', '彭山', 'p', 'peng shan', '511403000000', '彭山区', '101271503');
INSERT INTO `demo_city_code` VALUES ('511421', '5114', '2', '仁寿', 'r', 'ren shou', '511421000000', '仁寿县', '101271502');
INSERT INTO `demo_city_code` VALUES ('511423', '5114', '2', '洪雅', 'h', 'hong ya', '511423000000', '洪雅县', '101271504');
INSERT INTO `demo_city_code` VALUES ('511424', '5114', '2', '丹棱', 'd', 'dan ling', '511424000000', '丹棱县', '101271505');
INSERT INTO `demo_city_code` VALUES ('511425', '5114', '2', '青神', 'q', 'qing shen', '511425000000', '青神县', '101271506');
INSERT INTO `demo_city_code` VALUES ('5115', '51', '1', '宜宾', 'y', 'yi bin', '511500000000', '宜宾市', '101271101');
INSERT INTO `demo_city_code` VALUES ('511502', '5115', '2', '翠屏', 'c', 'cui ping', '511502000000', '翠屏区', NULL);
INSERT INTO `demo_city_code` VALUES ('511503', '5115', '2', '南溪', 'n', 'nan xi', '511503000000', '南溪区', '101271104');
INSERT INTO `demo_city_code` VALUES ('511504', '5115', '2', '叙州', 'x', 'xu zhou', '511504000000', '叙州区', NULL);
INSERT INTO `demo_city_code` VALUES ('511523', '5115', '2', '江安', 'j', 'jiang an', '511523000000', '江安县', '101271105');
INSERT INTO `demo_city_code` VALUES ('511524', '5115', '2', '长宁', 'c', 'chang ning', '511524000000', '长宁县', '101271106');
INSERT INTO `demo_city_code` VALUES ('511525', '5115', '2', '高县', 'g', 'gao xian', '511525000000', '高县', '101271107');
INSERT INTO `demo_city_code` VALUES ('511526', '5115', '2', '珙县', 'g', 'gong xian', '511526000000', '珙县', '101271108');
INSERT INTO `demo_city_code` VALUES ('511527', '5115', '2', '筠连', 'j', 'jun lian', '511527000000', '筠连县', '101271109');
INSERT INTO `demo_city_code` VALUES ('511528', '5115', '2', '兴文', 'x', 'xing wen', '511528000000', '兴文县', '101271110');
INSERT INTO `demo_city_code` VALUES ('511529', '5115', '2', '屏山', 'p', 'ping shan', '511529000000', '屏山县', '101271111');
INSERT INTO `demo_city_code` VALUES ('5116', '51', '1', '广安', 'g', 'guang an', '511600000000', '广安市', '101270801');
INSERT INTO `demo_city_code` VALUES ('511602', '5116', '2', '广安区', 'g', 'guang an qu', '511602000000', '广安区', NULL);
INSERT INTO `demo_city_code` VALUES ('511603', '5116', '2', '前锋', 'q', 'qian feng', '511603000000', '前锋区', NULL);
INSERT INTO `demo_city_code` VALUES ('511621', '5116', '2', '岳池', 'y', 'yue chi', '511621000000', '岳池县', '101270802');
INSERT INTO `demo_city_code` VALUES ('511622', '5116', '2', '武胜', 'w', 'wu sheng', '511622000000', '武胜县', '101270803');
INSERT INTO `demo_city_code` VALUES ('511623', '5116', '2', '邻水', 'l', 'lin shui', '511623000000', '邻水县', '101270804');
INSERT INTO `demo_city_code` VALUES ('511681', '5116', '2', '华蓥', 'h', 'hua ying', '511681000000', '华蓥市', '101270805');
INSERT INTO `demo_city_code` VALUES ('5117', '51', '1', '达州', 'd', 'da zhou', '511700000000', '达州市', '101270601');
INSERT INTO `demo_city_code` VALUES ('511702', '5117', '2', '通川', 't', 'tong chuan', '511702000000', '通川区', '101270607');
INSERT INTO `demo_city_code` VALUES ('511703', '5117', '2', '达川', 'd', 'da chuan', '511703000000', '达川区', NULL);
INSERT INTO `demo_city_code` VALUES ('511722', '5117', '2', '宣汉', 'x', 'xuan han', '511722000000', '宣汉县', '101270602');
INSERT INTO `demo_city_code` VALUES ('511723', '5117', '2', '开江', 'k', 'kai jiang', '511723000000', '开江县', '101270603');
INSERT INTO `demo_city_code` VALUES ('511724', '5117', '2', '大竹', 'd', 'da zhu', '511724000000', '大竹县', '101270604');
INSERT INTO `demo_city_code` VALUES ('511725', '5117', '2', '渠县', 'q', 'qu xian', '511725000000', '渠县', '101270605');
INSERT INTO `demo_city_code` VALUES ('511781', '5117', '2', '万源', 'w', 'wan yuan', '511781000000', '万源市', '101270606');
INSERT INTO `demo_city_code` VALUES ('5118', '51', '1', '雅安', 'y', 'ya an', '511800000000', '雅安市', '101271701');
INSERT INTO `demo_city_code` VALUES ('511802', '5118', '2', '雨城', 'y', 'yu cheng', '511802000000', '雨城区', NULL);
INSERT INTO `demo_city_code` VALUES ('511803', '5118', '2', '名山', 'm', 'ming shan', '511803000000', '名山区', '101271702');
INSERT INTO `demo_city_code` VALUES ('511822', '5118', '2', '荥经', 'y', 'ying jing', '511822000000', '荥经县', '101271703');
INSERT INTO `demo_city_code` VALUES ('511823', '5118', '2', '汉源', 'h', 'han yuan', '511823000000', '汉源县', '101271704');
INSERT INTO `demo_city_code` VALUES ('511824', '5118', '2', '石棉', 's', 'shi mian', '511824000000', '石棉县', '101271705');
INSERT INTO `demo_city_code` VALUES ('511825', '5118', '2', '天全', 't', 'tian quan', '511825000000', '天全县', '101271706');
INSERT INTO `demo_city_code` VALUES ('511826', '5118', '2', '芦山', 'l', 'lu shan', '511826000000', '芦山县', '101271707');
INSERT INTO `demo_city_code` VALUES ('511827', '5118', '2', '宝兴', 'b', 'bao xing', '511827000000', '宝兴县', '101271708');
INSERT INTO `demo_city_code` VALUES ('5119', '51', '1', '巴中', 'b', 'ba zhong', '511900000000', '巴中市', '101270901');
INSERT INTO `demo_city_code` VALUES ('511902', '5119', '2', '巴州', 'b', 'ba zhou', '511902000000', '巴州区', NULL);
INSERT INTO `demo_city_code` VALUES ('511903', '5119', '2', '恩阳', 'e', 'en yang', '511903000000', '恩阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('511921', '5119', '2', '通江', 't', 'tong jiang', '511921000000', '通江县', '101270902');
INSERT INTO `demo_city_code` VALUES ('511922', '5119', '2', '南江', 'n', 'nan jiang', '511922000000', '南江县', '101270903');
INSERT INTO `demo_city_code` VALUES ('511923', '5119', '2', '平昌', 'p', 'ping chang', '511923000000', '平昌县', '101270904');
INSERT INTO `demo_city_code` VALUES ('5120', '51', '1', '资阳', 'z', 'zi yang', '512000000000', '资阳市', '101271301');
INSERT INTO `demo_city_code` VALUES ('512002', '5120', '2', '雁江', 'y', 'yan jiang', '512002000000', '雁江区', NULL);
INSERT INTO `demo_city_code` VALUES ('512021', '5120', '2', '安岳', 'a', 'an yue', '512021000000', '安岳县', '101271302');
INSERT INTO `demo_city_code` VALUES ('512022', '5120', '2', '乐至', 'l', 'le zhi', '512022000000', '乐至县', '101271303');
INSERT INTO `demo_city_code` VALUES ('5132', '51', '1', '阿坝', 'a', 'a ba', '513200000000', '阿坝藏族羌族自治州', '101271901');
INSERT INTO `demo_city_code` VALUES ('513201', '5132', '2', '马尔康', 'm', 'ma er kang', '513201000000', '马尔康市', '101271910');
INSERT INTO `demo_city_code` VALUES ('513221', '5132', '2', '汶川', 'w', 'wen chuan', '513221000000', '汶川县', '101271902');
INSERT INTO `demo_city_code` VALUES ('513222', '5132', '2', '理县', 'l', 'li xian', '513222000000', '理县', '101271903');
INSERT INTO `demo_city_code` VALUES ('513223', '5132', '2', '茂县', 'm', 'mao xian', '513223000000', '茂县', '101271904');
INSERT INTO `demo_city_code` VALUES ('513224', '5132', '2', '松潘', 's', 'song pan', '513224000000', '松潘县', '101271905');
INSERT INTO `demo_city_code` VALUES ('513225', '5132', '2', '九寨沟', 'j', 'jiu zhai gou', '513225000000', '九寨沟县', '101271906');
INSERT INTO `demo_city_code` VALUES ('513226', '5132', '2', '金川', 'j', 'jin chuan', '513226000000', '金川县', '101271907');
INSERT INTO `demo_city_code` VALUES ('513227', '5132', '2', '小金', 'x', 'xiao jin', '513227000000', '小金县', '101271908');
INSERT INTO `demo_city_code` VALUES ('513228', '5132', '2', '黑水', 'h', 'hei shui', '513228000000', '黑水县', '101271909');
INSERT INTO `demo_city_code` VALUES ('513230', '5132', '2', '壤塘', 'r', 'rang tang', '513230000000', '壤塘县', '101271911');
INSERT INTO `demo_city_code` VALUES ('513231', '5132', '2', '阿坝县', 'a', 'a ba xian', '513231000000', '阿坝县', NULL);
INSERT INTO `demo_city_code` VALUES ('513232', '5132', '2', '若尔盖', 'r', 'ruo er gai', '513232000000', '若尔盖县', '101271912');
INSERT INTO `demo_city_code` VALUES ('513233', '5132', '2', '红原', 'h', 'hong yuan', '513233000000', '红原县', '101271913');
INSERT INTO `demo_city_code` VALUES ('5133', '51', '1', '甘孜', 'g', 'gan zi', '513300000000', '甘孜藏族自治州', '101271801');
INSERT INTO `demo_city_code` VALUES ('513301', '5133', '2', '康定', 'k', 'kang ding', '513301000000', '康定市', '101271802');
INSERT INTO `demo_city_code` VALUES ('513322', '5133', '2', '泸定', 'l', 'lu ding', '513322000000', '泸定县', '101271803');
INSERT INTO `demo_city_code` VALUES ('513323', '5133', '2', '丹巴', 'd', 'dan ba', '513323000000', '丹巴县', '101271804');
INSERT INTO `demo_city_code` VALUES ('513324', '5133', '2', '九龙', 'j', 'jiu long', '513324000000', '九龙县', '101271805');
INSERT INTO `demo_city_code` VALUES ('513325', '5133', '2', '雅江', 'y', 'ya jiang', '513325000000', '雅江县', '101271806');
INSERT INTO `demo_city_code` VALUES ('513326', '5133', '2', '道孚', 'd', 'dao fu', '513326000000', '道孚县', '101271807');
INSERT INTO `demo_city_code` VALUES ('513327', '5133', '2', '炉霍', 'l', 'lu huo', '513327000000', '炉霍县', '101271808');
INSERT INTO `demo_city_code` VALUES ('513328', '5133', '2', '甘孜县', 'g', 'gan zi xian', '513328000000', '甘孜县', NULL);
INSERT INTO `demo_city_code` VALUES ('513329', '5133', '2', '新龙', 'x', 'xin long', '513329000000', '新龙县', '101271809');
INSERT INTO `demo_city_code` VALUES ('513330', '5133', '2', '德格', 'd', 'de ge', '513330000000', '德格县', '101271810');
INSERT INTO `demo_city_code` VALUES ('513331', '5133', '2', '白玉', 'b', 'bai yu', '513331000000', '白玉县', '101271811');
INSERT INTO `demo_city_code` VALUES ('513332', '5133', '2', '石渠', 's', 'shi qu', '513332000000', '石渠县', '101271812');
INSERT INTO `demo_city_code` VALUES ('513333', '5133', '2', '色达', 's', 'se da', '513333000000', '色达县', '101271813');
INSERT INTO `demo_city_code` VALUES ('513334', '5133', '2', '理塘', 'l', 'li tang', '513334000000', '理塘县', '101271814');
INSERT INTO `demo_city_code` VALUES ('513335', '5133', '2', '巴塘', 'b', 'ba tang', '513335000000', '巴塘县', '101271815');
INSERT INTO `demo_city_code` VALUES ('513336', '5133', '2', '乡城', 'x', 'xiang cheng', '513336000000', '乡城县', '101271816');
INSERT INTO `demo_city_code` VALUES ('513337', '5133', '2', '稻城', 'd', 'dao cheng', '513337000000', '稻城县', '101271817');
INSERT INTO `demo_city_code` VALUES ('513338', '5133', '2', '得荣', 'd', 'de rong', '513338000000', '得荣县', '101271818');
INSERT INTO `demo_city_code` VALUES ('5134', '51', '1', '凉山', 'l', 'liang shan', '513400000000', '凉山彝族自治州', '101271601');
INSERT INTO `demo_city_code` VALUES ('513401', '5134', '2', '西昌', 'x', 'xi chang', '513401000000', '西昌市', '101271610');
INSERT INTO `demo_city_code` VALUES ('513402', '5134', '2', '会理', 'h', 'hui li', '513402000000', '会理市', '101271606');
INSERT INTO `demo_city_code` VALUES ('513422', '5134', '2', '木里', 'm', 'mu li', '513422000000', '木里藏族自治县', '101271603');
INSERT INTO `demo_city_code` VALUES ('513423', '5134', '2', '盐源', 'y', 'yan yuan', '513423000000', '盐源县', '101271604');
INSERT INTO `demo_city_code` VALUES ('513424', '5134', '2', '德昌', 'd', 'de chang', '513424000000', '德昌县', '101271605');
INSERT INTO `demo_city_code` VALUES ('513426', '5134', '2', '会东', 'h', 'hui dong', '513426000000', '会东县', '101271607');
INSERT INTO `demo_city_code` VALUES ('513427', '5134', '2', '宁南', 'n', 'ning nan', '513427000000', '宁南县', '101271608');
INSERT INTO `demo_city_code` VALUES ('513428', '5134', '2', '普格', 'p', 'pu ge', '513428000000', '普格县', '101271609');
INSERT INTO `demo_city_code` VALUES ('513429', '5134', '2', '布拖', 'b', 'bu tuo', '513429000000', '布拖县', '101271619');
INSERT INTO `demo_city_code` VALUES ('513430', '5134', '2', '金阳', 'j', 'jin yang', '513430000000', '金阳县', '101271611');
INSERT INTO `demo_city_code` VALUES ('513431', '5134', '2', '昭觉', 'z', 'zhao jue', '513431000000', '昭觉县', '101271612');
INSERT INTO `demo_city_code` VALUES ('513432', '5134', '2', '喜德', 'x', 'xi de', '513432000000', '喜德县', '101271613');
INSERT INTO `demo_city_code` VALUES ('513433', '5134', '2', '冕宁', 'm', 'mian ning', '513433000000', '冕宁县', '101271614');
INSERT INTO `demo_city_code` VALUES ('513434', '5134', '2', '越西', 'y', 'yue xi', '513434000000', '越西县', '101271615');
INSERT INTO `demo_city_code` VALUES ('513435', '5134', '2', '甘洛', 'g', 'gan luo', '513435000000', '甘洛县', '101271616');
INSERT INTO `demo_city_code` VALUES ('513436', '5134', '2', '美姑', 'm', 'mei gu', '513436000000', '美姑县', '101271618');
INSERT INTO `demo_city_code` VALUES ('513437', '5134', '2', '雷波', 'l', 'lei bo', '513437000000', '雷波县', '101271617');
INSERT INTO `demo_city_code` VALUES ('52', '0', '0', '贵州', 'g', 'gui zhou', '520000000000', '贵州省', NULL);
INSERT INTO `demo_city_code` VALUES ('5201', '52', '1', '贵阳', 'g', 'gui yang', '520100000000', '贵阳市', '101260101');
INSERT INTO `demo_city_code` VALUES ('520102', '5201', '2', '南明', 'n', 'nan ming', '520102000000', '南明区', '101260111');
INSERT INTO `demo_city_code` VALUES ('520103', '5201', '2', '云岩', 'y', 'yun yan', '520103000000', '云岩区', '101260110');
INSERT INTO `demo_city_code` VALUES ('520111', '5201', '2', '花溪', 'h', 'hua xi', '520111000000', '花溪区', '101260103');
INSERT INTO `demo_city_code` VALUES ('520112', '5201', '2', '乌当', 'w', 'wu dang', '520112000000', '乌当区', '101260104');
INSERT INTO `demo_city_code` VALUES ('520113', '5201', '2', '白云', 'b', 'bai yun', '520113000000', '白云区', '101260102');
INSERT INTO `demo_city_code` VALUES ('520115', '5201', '2', '观山湖', 'g', 'guan shan hu', '520115000000', '观山湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('520121', '5201', '2', '开阳', 'k', 'kai yang', '520121000000', '开阳县', '101260106');
INSERT INTO `demo_city_code` VALUES ('520122', '5201', '2', '息烽', 'x', 'xi feng', '520122000000', '息烽县', '101260105');
INSERT INTO `demo_city_code` VALUES ('520123', '5201', '2', '修文', 'x', 'xiu wen', '520123000000', '修文县', '101260107');
INSERT INTO `demo_city_code` VALUES ('520181', '5201', '2', '清镇', 'q', 'qing zhen', '520181000000', '清镇市', '101260108');
INSERT INTO `demo_city_code` VALUES ('5202', '52', '1', '六盘水', 'l', 'liu pan shui', '520200000000', '六盘水市', NULL);
INSERT INTO `demo_city_code` VALUES ('520201', '5202', '2', '钟山', 'z', 'zhong shan', '520201000000', '钟山区', '101300704');
INSERT INTO `demo_city_code` VALUES ('520203', '5202', '2', '六枝特', 'l', 'liu zhi te', '520203000000', '六枝特区', NULL);
INSERT INTO `demo_city_code` VALUES ('520204', '5202', '2', '水城', 's', 'shui cheng', '520204000000', '水城区', '101260801');
INSERT INTO `demo_city_code` VALUES ('520281', '5202', '2', '盘州', 'p', 'pan zhou', '520281000000', '盘州市', NULL);
INSERT INTO `demo_city_code` VALUES ('5203', '52', '1', '遵义', 'z', 'zun yi', '520300000000', '遵义市', '101260201');
INSERT INTO `demo_city_code` VALUES ('520302', '5203', '2', '红花岗', 'h', 'hong hua gang', '520302000000', '红花岗区', '101260215');
INSERT INTO `demo_city_code` VALUES ('520303', '5203', '2', '汇川', 'h', 'hui chuan', '520303000000', '汇川区', '101260214');
INSERT INTO `demo_city_code` VALUES ('520304', '5203', '2', '播州', 'b', 'bo zhou', '520304000000', '播州区', NULL);
INSERT INTO `demo_city_code` VALUES ('520322', '5203', '2', '桐梓', 't', 'tong zi', '520322000000', '桐梓县', '101260207');
INSERT INTO `demo_city_code` VALUES ('520323', '5203', '2', '绥阳', 's', 'sui yang', '520323000000', '绥阳县', '101260204');
INSERT INTO `demo_city_code` VALUES ('520324', '5203', '2', '正安', 'z', 'zheng an', '520324000000', '正安县', '101260211');
INSERT INTO `demo_city_code` VALUES ('520325', '5203', '2', '道真', 'd', 'dao zhen', '520325000000', '道真仡佬族苗族自治县', '101260210');
INSERT INTO `demo_city_code` VALUES ('520326', '5203', '2', '务川', 'w', 'wu chuan', '520326000000', '务川仡佬族苗族自治县', '101260212');
INSERT INTO `demo_city_code` VALUES ('520327', '5203', '2', '凤冈', 'f', 'feng gang', '520327000000', '凤冈县', '101260206');
INSERT INTO `demo_city_code` VALUES ('520328', '5203', '2', '湄潭', 'm', 'mei tan', '520328000000', '湄潭县', '101260205');
INSERT INTO `demo_city_code` VALUES ('520329', '5203', '2', '余庆', 'y', 'yu qing', '520329000000', '余庆县', '101260213');
INSERT INTO `demo_city_code` VALUES ('520330', '5203', '2', '习水', 'x', 'xi shui', '520330000000', '习水县', '101260209');
INSERT INTO `demo_city_code` VALUES ('520381', '5203', '2', '赤水', 'c', 'chi shui', '520381000000', '赤水市', '101260208');
INSERT INTO `demo_city_code` VALUES ('520382', '5203', '2', '仁怀', 'r', 'ren huai', '520382000000', '仁怀市', '101260203');
INSERT INTO `demo_city_code` VALUES ('5204', '52', '1', '安顺', 'a', 'an shun', '520400000000', '安顺市', '101260301');
INSERT INTO `demo_city_code` VALUES ('520402', '5204', '2', '西秀', 'x', 'xi xiu', '520402000000', '西秀区', NULL);
INSERT INTO `demo_city_code` VALUES ('520403', '5204', '2', '平坝', 'p', 'ping ba', '520403000000', '平坝区', '101260304');
INSERT INTO `demo_city_code` VALUES ('520422', '5204', '2', '普定', 'p', 'pu ding', '520422000000', '普定县', '101260302');
INSERT INTO `demo_city_code` VALUES ('520423', '5204', '2', '镇宁', 'z', 'zhen ning', '520423000000', '镇宁布依族苗族自治县', '101260303');
INSERT INTO `demo_city_code` VALUES ('520424', '5204', '2', '关岭', 'g', 'guan ling', '520424000000', '关岭布依族苗族自治县', '101260306');
INSERT INTO `demo_city_code` VALUES ('520425', '5204', '2', '紫云', 'z', 'zi yun', '520425000000', '紫云苗族布依族自治县', '101260305');
INSERT INTO `demo_city_code` VALUES ('5205', '52', '1', '毕节', 'b', 'bi jie', '520500000000', '毕节市', '101260701');
INSERT INTO `demo_city_code` VALUES ('520502', '5205', '2', '七星关', 'q', 'qi xing guan', '520502000000', '七星关区', NULL);
INSERT INTO `demo_city_code` VALUES ('520521', '5205', '2', '大方', 'd', 'da fang', '520521000000', '大方县', '101260705');
INSERT INTO `demo_city_code` VALUES ('520523', '5205', '2', '金沙', 'j', 'jin sha', '520523000000', '金沙县', '101260703');
INSERT INTO `demo_city_code` VALUES ('520524', '5205', '2', '织金', 'z', 'zhi jin', '520524000000', '织金县', '101260707');
INSERT INTO `demo_city_code` VALUES ('520525', '5205', '2', '纳雍', 'n', 'na yong', '520525000000', '纳雍县', '101260706');
INSERT INTO `demo_city_code` VALUES ('520526', '5205', '2', '威宁', 'w', 'wei ning', '520526000000', '威宁彝族回族苗族自治县', '101260704');
INSERT INTO `demo_city_code` VALUES ('520527', '5205', '2', '赫章', 'h', 'he zhang', '520527000000', '赫章县', '101260702');
INSERT INTO `demo_city_code` VALUES ('520581', '5205', '2', '黔西', 'q', 'qian xi', '520581000000', '黔西市', '101260708');
INSERT INTO `demo_city_code` VALUES ('5206', '52', '1', '铜仁', 't', 'tong ren', '520600000000', '铜仁市', '101260601');
INSERT INTO `demo_city_code` VALUES ('520602', '5206', '2', '碧江', 'b', 'bi jiang', '520602000000', '碧江区', NULL);
INSERT INTO `demo_city_code` VALUES ('520603', '5206', '2', '万山', 'w', 'wan shan', '520603000000', '万山区', '101260604');
INSERT INTO `demo_city_code` VALUES ('520621', '5206', '2', '江口', 'j', 'jiang kou', '520621000000', '江口县', '101260602');
INSERT INTO `demo_city_code` VALUES ('520622', '5206', '2', '玉屏', 'y', 'yu ping', '520622000000', '玉屏侗族自治县', '101260603');
INSERT INTO `demo_city_code` VALUES ('520623', '5206', '2', '石阡', 's', 'shi qian', '520623000000', '石阡县', '101260608');
INSERT INTO `demo_city_code` VALUES ('520624', '5206', '2', '思南', 's', 'si nan', '520624000000', '思南县', '101260605');
INSERT INTO `demo_city_code` VALUES ('520625', '5206', '2', '印江', 'y', 'yin jiang', '520625000000', '印江土家族苗族自治县', '101260607');
INSERT INTO `demo_city_code` VALUES ('520626', '5206', '2', '德江', 'd', 'de jiang', '520626000000', '德江县', '101260610');
INSERT INTO `demo_city_code` VALUES ('520627', '5206', '2', '沿河', 'y', 'yan he', '520627000000', '沿河土家族自治县', '101260609');
INSERT INTO `demo_city_code` VALUES ('520628', '5206', '2', '松桃', 's', 'song tao', '520628000000', '松桃苗族自治县', '101260611');
INSERT INTO `demo_city_code` VALUES ('5223', '52', '1', '黔西南', 'q', 'qian xi nan', '522300000000', '黔西南布依族苗族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('522301', '5223', '2', '兴义', 'x', 'xing yi', '522301000000', '兴义市', '101260901');
INSERT INTO `demo_city_code` VALUES ('522302', '5223', '2', '兴仁', 'x', 'xing ren', '522302000000', '兴仁市', '101260903');
INSERT INTO `demo_city_code` VALUES ('522323', '5223', '2', '普安', 'p', 'pu an', '522323000000', '普安县', '101260909');
INSERT INTO `demo_city_code` VALUES ('522324', '5223', '2', '晴隆', 'q', 'qing long', '522324000000', '晴隆县', '101260902');
INSERT INTO `demo_city_code` VALUES ('522325', '5223', '2', '贞丰', 'z', 'zhen feng', '522325000000', '贞丰县', '101260904');
INSERT INTO `demo_city_code` VALUES ('522326', '5223', '2', '望谟', 'w', 'wang mo', '522326000000', '望谟县', '101260905');
INSERT INTO `demo_city_code` VALUES ('522327', '5223', '2', '册亨', 'c', 'ce heng', '522327000000', '册亨县', '101260908');
INSERT INTO `demo_city_code` VALUES ('522328', '5223', '2', '安龙', 'a', 'an long', '522328000000', '安龙县', '101260907');
INSERT INTO `demo_city_code` VALUES ('5226', '52', '1', '黔东南', 'q', 'qian dong nan', '522600000000', '黔东南苗族侗族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('522601', '5226', '2', '凯里', 'k', 'kai li', '522601000000', '凯里市', '101260501');
INSERT INTO `demo_city_code` VALUES ('522622', '5226', '2', '黄平', 'h', 'huang ping', '522622000000', '黄平县', '101260505');
INSERT INTO `demo_city_code` VALUES ('522623', '5226', '2', '施秉', 's', 'shi bing', '522623000000', '施秉县', '101260503');
INSERT INTO `demo_city_code` VALUES ('522624', '5226', '2', '三穗', 's', 'san sui', '522624000000', '三穗县', '101260509');
INSERT INTO `demo_city_code` VALUES ('522625', '5226', '2', '镇远', 'z', 'zhen yuan', '522625000000', '镇远县', '101260504');
INSERT INTO `demo_city_code` VALUES ('522626', '5226', '2', '岑巩', 'c', 'cen gong', '522626000000', '岑巩县', '101260502');
INSERT INTO `demo_city_code` VALUES ('522627', '5226', '2', '天柱', 't', 'tian zhu', '522627000000', '天柱县', '101260514');
INSERT INTO `demo_city_code` VALUES ('522628', '5226', '2', '锦屏', 'j', 'jin ping', '522628000000', '锦屏县', '101260515');
INSERT INTO `demo_city_code` VALUES ('522629', '5226', '2', '剑河', 'j', 'jian he', '522629000000', '剑河县', '101260511');
INSERT INTO `demo_city_code` VALUES ('522630', '5226', '2', '台江', 't', 'tai jiang', '522630000000', '台江县', '101260510');
INSERT INTO `demo_city_code` VALUES ('522631', '5226', '2', '黎平', 'l', 'li ping', '522631000000', '黎平县', '101260513');
INSERT INTO `demo_city_code` VALUES ('522632', '5226', '2', '榕江', 'r', 'rong jiang', '522632000000', '榕江县', '101260516');
INSERT INTO `demo_city_code` VALUES ('522633', '5226', '2', '从江', 'c', 'cong jiang', '522633000000', '从江县', '101260517');
INSERT INTO `demo_city_code` VALUES ('522634', '5226', '2', '雷山', 'l', 'lei shan', '522634000000', '雷山县', '101260512');
INSERT INTO `demo_city_code` VALUES ('522635', '5226', '2', '麻江', 'm', 'ma jiang', '522635000000', '麻江县', '101260507');
INSERT INTO `demo_city_code` VALUES ('522636', '5226', '2', '丹寨', 'd', 'dan zhai', '522636000000', '丹寨县', '101260508');
INSERT INTO `demo_city_code` VALUES ('5227', '52', '1', '黔南', 'q', 'qian nan', '522700000000', '黔南布依族苗族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('522701', '5227', '2', '都匀', 'd', 'du yun', '522701000000', '都匀市', '101260401');
INSERT INTO `demo_city_code` VALUES ('522702', '5227', '2', '福泉', 'f', 'fu quan', '522702000000', '福泉市', '101260405');
INSERT INTO `demo_city_code` VALUES ('522722', '5227', '2', '荔波', 'l', 'li bo', '522722000000', '荔波县', '101260412');
INSERT INTO `demo_city_code` VALUES ('522723', '5227', '2', '贵定', 'g', 'gui ding', '522723000000', '贵定县', '101260402');
INSERT INTO `demo_city_code` VALUES ('522725', '5227', '2', '瓮安', 'w', 'weng an', '522725000000', '瓮安县', '101260403');
INSERT INTO `demo_city_code` VALUES ('522726', '5227', '2', '独山', 'd', 'du shan', '522726000000', '独山县', '101260410');
INSERT INTO `demo_city_code` VALUES ('522727', '5227', '2', '平塘', 'p', 'ping tang', '522727000000', '平塘县', '101260409');
INSERT INTO `demo_city_code` VALUES ('522728', '5227', '2', '罗甸', 'l', 'luo dian', '522728000000', '罗甸县', '101260408');
INSERT INTO `demo_city_code` VALUES ('522729', '5227', '2', '长顺', 'c', 'chang shun', '522729000000', '长顺县', '101260404');
INSERT INTO `demo_city_code` VALUES ('522730', '5227', '2', '龙里', 'l', 'long li', '522730000000', '龙里县', '101260407');
INSERT INTO `demo_city_code` VALUES ('522731', '5227', '2', '惠水', 'h', 'hui shui', '522731000000', '惠水县', '101260406');
INSERT INTO `demo_city_code` VALUES ('522732', '5227', '2', '三都', 's', 'san du', '522732000000', '三都水族自治县', '101260411');
INSERT INTO `demo_city_code` VALUES ('53', '0', '0', '云南', 'y', 'yun nan', '530000000000', '云南省', NULL);
INSERT INTO `demo_city_code` VALUES ('5301', '53', '1', '昆明', 'k', 'kun ming', '530100000000', '昆明市', '101290101');
INSERT INTO `demo_city_code` VALUES ('530102', '5301', '2', '五华', 'w', 'wu hua', '530102000000', '五华区', '101280408');
INSERT INTO `demo_city_code` VALUES ('530103', '5301', '2', '盘龙', 'p', 'pan long', '530103000000', '盘龙区', NULL);
INSERT INTO `demo_city_code` VALUES ('530111', '5301', '2', '官渡', 'g', 'guan du', '530111000000', '官渡区', NULL);
INSERT INTO `demo_city_code` VALUES ('530112', '5301', '2', '西山', 'x', 'xi shan', '530112000000', '西山区', NULL);
INSERT INTO `demo_city_code` VALUES ('530113', '5301', '2', '东川', 'd', 'dong chuan', '530113000000', '东川区', '101290103');
INSERT INTO `demo_city_code` VALUES ('530114', '5301', '2', '呈贡', 'c', 'cheng gong', '530114000000', '呈贡区', '101290108');
INSERT INTO `demo_city_code` VALUES ('530115', '5301', '2', '晋宁', 'j', 'jin ning', '530115000000', '晋宁区', '101290105');
INSERT INTO `demo_city_code` VALUES ('530124', '5301', '2', '富民', 'f', 'fu min', '530124000000', '富民县', '101290109');
INSERT INTO `demo_city_code` VALUES ('530125', '5301', '2', '宜良', 'y', 'yi liang', '530125000000', '宜良县', '101290106');
INSERT INTO `demo_city_code` VALUES ('530126', '5301', '2', '石林', 's', 'shi lin', '530126000000', '石林彝族自治县', '101290107');
INSERT INTO `demo_city_code` VALUES ('530127', '5301', '2', '嵩明', 's', 'song ming', '530127000000', '嵩明县', '101290110');
INSERT INTO `demo_city_code` VALUES ('530128', '5301', '2', '禄劝', 'l', 'lu quan', '530128000000', '禄劝彝族苗族自治县', '101290111');
INSERT INTO `demo_city_code` VALUES ('530129', '5301', '2', '寻甸', 'x', 'xun dian', '530129000000', '寻甸回族彝族自治县', '101290104');
INSERT INTO `demo_city_code` VALUES ('530181', '5301', '2', '安宁', 'a', 'an ning', '530181000000', '安宁市', '101290112');
INSERT INTO `demo_city_code` VALUES ('5303', '53', '1', '曲靖', 'q', 'qu jing', '530300000000', '曲靖市', '101290401');
INSERT INTO `demo_city_code` VALUES ('530302', '5303', '2', '麒麟', 'q', 'qi lin', '530302000000', '麒麟区', NULL);
INSERT INTO `demo_city_code` VALUES ('530303', '5303', '2', '沾益', 'z', 'zhan yi', '530303000000', '沾益区', '101290402');
INSERT INTO `demo_city_code` VALUES ('530304', '5303', '2', '马龙', 'm', 'ma long', '530304000000', '马龙区', '101290405');
INSERT INTO `demo_city_code` VALUES ('530322', '5303', '2', '陆良', 'l', 'lu liang', '530322000000', '陆良县', '101290403');
INSERT INTO `demo_city_code` VALUES ('530323', '5303', '2', '师宗', 's', 'shi zong', '530323000000', '师宗县', '101290406');
INSERT INTO `demo_city_code` VALUES ('530324', '5303', '2', '罗平', 'l', 'luo ping', '530324000000', '罗平县', '101290407');
INSERT INTO `demo_city_code` VALUES ('530325', '5303', '2', '富源', 'f', 'fu yuan', '530325000000', '富源县', '101290404');
INSERT INTO `demo_city_code` VALUES ('530326', '5303', '2', '会泽', 'h', 'hui ze', '530326000000', '会泽县', '101290408');
INSERT INTO `demo_city_code` VALUES ('530381', '5303', '2', '宣威', 'x', 'xuan wei', '530381000000', '宣威市', '101290409');
INSERT INTO `demo_city_code` VALUES ('5304', '53', '1', '玉溪', 'y', 'yu xi', '530400000000', '玉溪市', '101290701');
INSERT INTO `demo_city_code` VALUES ('530402', '5304', '2', '红塔', 'h', 'hong ta', '530402000000', '红塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('530403', '5304', '2', '江川', 'j', 'jiang chuan', '530403000000', '江川区', '101290703');
INSERT INTO `demo_city_code` VALUES ('530423', '5304', '2', '通海', 't', 'tong hai', '530423000000', '通海县', '101290704');
INSERT INTO `demo_city_code` VALUES ('530424', '5304', '2', '华宁', 'h', 'hua ning', '530424000000', '华宁县', '101290705');
INSERT INTO `demo_city_code` VALUES ('530425', '5304', '2', '易门', 'y', 'yi men', '530425000000', '易门县', '101290707');
INSERT INTO `demo_city_code` VALUES ('530426', '5304', '2', '峨山', 'e', 'e shan', '530426000000', '峨山彝族自治县', '101290708');
INSERT INTO `demo_city_code` VALUES ('530427', '5304', '2', '新平', 'x', 'xin ping', '530427000000', '新平彝族傣族自治县', '101290706');
INSERT INTO `demo_city_code` VALUES ('530428', '5304', '2', '元江', 'y', 'yuan jiang', '530428000000', '元江哈尼族彝族傣族自治县', '101290709');
INSERT INTO `demo_city_code` VALUES ('530481', '5304', '2', '澄江', 'c', 'cheng jiang', '530481000000', '澄江市', '101290702');
INSERT INTO `demo_city_code` VALUES ('5305', '53', '1', '保山', 'b', 'bao shan', '530500000000', '保山市', '101290501');
INSERT INTO `demo_city_code` VALUES ('530502', '5305', '2', '隆阳', 'l', 'long yang', '530502000000', '隆阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('530521', '5305', '2', '施甸', 's', 'shi dian', '530521000000', '施甸县', '101290504');
INSERT INTO `demo_city_code` VALUES ('530523', '5305', '2', '龙陵', 'l', 'long ling', '530523000000', '龙陵县', '101290503');
INSERT INTO `demo_city_code` VALUES ('530524', '5305', '2', '昌宁', 'c', 'chang ning', '530524000000', '昌宁县', '101290505');
INSERT INTO `demo_city_code` VALUES ('530581', '5305', '2', '腾冲', 't', 'teng chong', '530581000000', '腾冲市', '101290506');
INSERT INTO `demo_city_code` VALUES ('5306', '53', '1', '昭通', 'z', 'zhao tong', '530600000000', '昭通市', '101291001');
INSERT INTO `demo_city_code` VALUES ('530602', '5306', '2', '昭阳', 'z', 'zhao yang', '530602000000', '昭阳区', NULL);
INSERT INTO `demo_city_code` VALUES ('530621', '5306', '2', '鲁甸', 'l', 'lu dian', '530621000000', '鲁甸县', '101291002');
INSERT INTO `demo_city_code` VALUES ('530622', '5306', '2', '巧家', 'q', 'qiao jia', '530622000000', '巧家县', '101291006');
INSERT INTO `demo_city_code` VALUES ('530623', '5306', '2', '盐津', 'y', 'yan jin', '530623000000', '盐津县', '101291009');
INSERT INTO `demo_city_code` VALUES ('530624', '5306', '2', '大关', 'd', 'da guan', '530624000000', '大关县', '101291010');
INSERT INTO `demo_city_code` VALUES ('530625', '5306', '2', '永善', 'y', 'yong shan', '530625000000', '永善县', '101291008');
INSERT INTO `demo_city_code` VALUES ('530626', '5306', '2', '绥江', 's', 'sui jiang', '530626000000', '绥江县', '101291007');
INSERT INTO `demo_city_code` VALUES ('530627', '5306', '2', '镇雄', 'z', 'zhen xiong', '530627000000', '镇雄县', '101291004');
INSERT INTO `demo_city_code` VALUES ('530628', '5306', '2', '彝良', 'y', 'yi liang', '530628000000', '彝良县', '101291003');
INSERT INTO `demo_city_code` VALUES ('530629', '5306', '2', '威信', 'w', 'wei xin', '530629000000', '威信县', '101291005');
INSERT INTO `demo_city_code` VALUES ('530681', '5306', '2', '水富', 's', 'shui fu', '530681000000', '水富市', '101291011');
INSERT INTO `demo_city_code` VALUES ('5307', '53', '1', '丽江', 'l', 'li jiang', '530700000000', '丽江市', '101291401');
INSERT INTO `demo_city_code` VALUES ('530702', '5307', '2', '古城', 'g', 'gu cheng', '530702000000', '古城区', NULL);
INSERT INTO `demo_city_code` VALUES ('530721', '5307', '2', '玉龙', 'y', 'yu long', '530721000000', '玉龙纳西族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('530722', '5307', '2', '永胜', 'y', 'yong sheng', '530722000000', '永胜县', '101291402');
INSERT INTO `demo_city_code` VALUES ('530723', '5307', '2', '华坪', 'h', 'hua ping', '530723000000', '华坪县', '101291403');
INSERT INTO `demo_city_code` VALUES ('530724', '5307', '2', '宁蒗', 'n', 'ning lang', '530724000000', '宁蒗彝族自治县', '101291404');
INSERT INTO `demo_city_code` VALUES ('5308', '53', '1', '普洱', 'p', 'pu er', '530800000000', '普洱市', '101290901');
INSERT INTO `demo_city_code` VALUES ('530802', '5308', '2', '思茅', 's', 'si mao', '530802000000', '思茅区', NULL);
INSERT INTO `demo_city_code` VALUES ('530821', '5308', '2', '宁洱', 'n', 'ning er', '530821000000', '宁洱哈尼族彝族自治县', '101290912');
INSERT INTO `demo_city_code` VALUES ('530822', '5308', '2', '墨江', 'm', 'mo jiang', '530822000000', '墨江哈尼族自治县', '101290906');
INSERT INTO `demo_city_code` VALUES ('530823', '5308', '2', '景东', 'j', 'jing dong', '530823000000', '景东彝族自治县', '101290903');
INSERT INTO `demo_city_code` VALUES ('530824', '5308', '2', '景谷', 'j', 'jing gu', '530824000000', '景谷傣族彝族自治县', '101290902');
INSERT INTO `demo_city_code` VALUES ('530825', '5308', '2', '镇沅', 'z', 'zhen yuan', '530825000000', '镇沅彝族哈尼族拉祜族自治县', '101290911');
INSERT INTO `demo_city_code` VALUES ('530826', '5308', '2', '江城', 'j', 'jiang cheng', '530826000000', '江城哈尼族彝族自治县', '101290907');
INSERT INTO `demo_city_code` VALUES ('530827', '5308', '2', '孟连', 'm', 'meng lian', '530827000000', '孟连傣族拉祜族佤族自治县', '101290908');
INSERT INTO `demo_city_code` VALUES ('530828', '5308', '2', '澜沧', 'l', 'lan cang', '530828000000', '澜沧拉祜族自治县', '101290904');
INSERT INTO `demo_city_code` VALUES ('530829', '5308', '2', '西盟', 'x', 'xi meng', '530829000000', '西盟佤族自治县', '101290909');
INSERT INTO `demo_city_code` VALUES ('5309', '53', '1', '临沧', 'l', 'lin cang', '530900000000', '临沧市', '101291101');
INSERT INTO `demo_city_code` VALUES ('530902', '5309', '2', '临翔', 'l', 'lin xiang', '530902000000', '临翔区', NULL);
INSERT INTO `demo_city_code` VALUES ('530921', '5309', '2', '凤庆', 'f', 'feng qing', '530921000000', '凤庆县', '101291105');
INSERT INTO `demo_city_code` VALUES ('530922', '5309', '2', '云县', 'y', 'yun xian', '530922000000', '云县', '101291107');
INSERT INTO `demo_city_code` VALUES ('530923', '5309', '2', '永德', 'y', 'yong de', '530923000000', '永德县', '101291106');
INSERT INTO `demo_city_code` VALUES ('530924', '5309', '2', '镇康', 'z', 'zhen kang', '530924000000', '镇康县', '101291108');
INSERT INTO `demo_city_code` VALUES ('530925', '5309', '2', '双江', 's', 'shuang jiang', '530925000000', '双江拉祜族佤族布朗族傣族自治县', '101291104');
INSERT INTO `demo_city_code` VALUES ('530926', '5309', '2', '耿马', 'g', 'geng ma', '530926000000', '耿马傣族佤族自治县', '101291103');
INSERT INTO `demo_city_code` VALUES ('530927', '5309', '2', '沧源', 'c', 'cang yuan', '530927000000', '沧源佤族自治县', '101291102');
INSERT INTO `demo_city_code` VALUES ('5323', '53', '1', '楚雄', 'c', 'chu xiong', '532300000000', '楚雄彝族自治州', '101290801');
INSERT INTO `demo_city_code` VALUES ('532301', '5323', '2', '楚雄市', 'c', 'chu xiong shi', '532301000000', '楚雄市', NULL);
INSERT INTO `demo_city_code` VALUES ('532302', '5323', '2', '禄丰', 'l', 'lu feng', '532302000000', '禄丰市', '101290808');
INSERT INTO `demo_city_code` VALUES ('532322', '5323', '2', '双柏', 's', 'shuang bai', '532322000000', '双柏县', '101290809');
INSERT INTO `demo_city_code` VALUES ('532323', '5323', '2', '牟定', 'm', 'mou ding', '532323000000', '牟定县', '101290805');
INSERT INTO `demo_city_code` VALUES ('532324', '5323', '2', '南华', 'n', 'nan hua', '532324000000', '南华县', '101290806');
INSERT INTO `demo_city_code` VALUES ('532325', '5323', '2', '姚安', 'y', 'yao an', '532325000000', '姚安县', '101290804');
INSERT INTO `demo_city_code` VALUES ('532326', '5323', '2', '大姚', 'd', 'da yao', '532326000000', '大姚县', '101290802');
INSERT INTO `demo_city_code` VALUES ('532327', '5323', '2', '永仁', 'y', 'yong ren', '532327000000', '永仁县', '101290810');
INSERT INTO `demo_city_code` VALUES ('532328', '5323', '2', '元谋', 'y', 'yuan mou', '532328000000', '元谋县', '101290803');
INSERT INTO `demo_city_code` VALUES ('532329', '5323', '2', '武定', 'w', 'wu ding', '532329000000', '武定县', '101290807');
INSERT INTO `demo_city_code` VALUES ('5325', '53', '1', '红河', 'h', 'hong he', '532500000000', '红河哈尼族彝族自治州', '101290301');
INSERT INTO `demo_city_code` VALUES ('532501', '5325', '2', '个旧', 'g', 'ge jiu', '532501000000', '个旧市', '101290308');
INSERT INTO `demo_city_code` VALUES ('532502', '5325', '2', '开远', 'k', 'kai yuan', '532502000000', '开远市', '101290307');
INSERT INTO `demo_city_code` VALUES ('532503', '5325', '2', '蒙自', 'm', 'meng zi', '532503000000', '蒙自市', '101290309');
INSERT INTO `demo_city_code` VALUES ('532504', '5325', '2', '弥勒', 'm', 'mi le', '532504000000', '弥勒市', '101290304');
INSERT INTO `demo_city_code` VALUES ('532523', '5325', '2', '屏边', 'p', 'ping bian', '532523000000', '屏边苗族自治县', '101290310');
INSERT INTO `demo_city_code` VALUES ('532524', '5325', '2', '建水', 'j', 'jian shui', '532524000000', '建水县', '101290303');
INSERT INTO `demo_city_code` VALUES ('532525', '5325', '2', '石屏', 's', 'shi ping', '532525000000', '石屏县', '101290302');
INSERT INTO `demo_city_code` VALUES ('532527', '5325', '2', '泸西', 'l', 'lu xi', '532527000000', '泸西县', '101290311');
INSERT INTO `demo_city_code` VALUES ('532528', '5325', '2', '元阳', 'y', 'yuan yang', '532528000000', '元阳县', '101290305');
INSERT INTO `demo_city_code` VALUES ('532529', '5325', '2', '红河县', 'h', 'hong he xian', '532529000000', '红河县', NULL);
INSERT INTO `demo_city_code` VALUES ('532530', '5325', '2', '金平', 'j', 'jin ping', '532530000000', '金平苗族瑶族傣族自治县', '101290312');
INSERT INTO `demo_city_code` VALUES ('532531', '5325', '2', '绿春', 'l', 'lv chun', '532531000000', '绿春县', '101290306');
INSERT INTO `demo_city_code` VALUES ('532532', '5325', '2', '河口', 'h', 'he kou', '532532000000', '河口瑶族自治县', '101121202');
INSERT INTO `demo_city_code` VALUES ('5326', '53', '1', '文山', 'w', 'wen shan', '532600000000', '文山壮族苗族自治州', '101290601');
INSERT INTO `demo_city_code` VALUES ('532601', '5326', '2', '文山市', 'w', 'wen shan shi', '532601000000', '文山市', NULL);
INSERT INTO `demo_city_code` VALUES ('532622', '5326', '2', '砚山', 'y', 'yan shan', '532622000000', '砚山县', '101290605');
INSERT INTO `demo_city_code` VALUES ('532623', '5326', '2', '西畴', 'x', 'xi chou', '532623000000', '西畴县', '101290602');
INSERT INTO `demo_city_code` VALUES ('532624', '5326', '2', '麻栗坡', 'm', 'ma li po', '532624000000', '麻栗坡县', '101290604');
INSERT INTO `demo_city_code` VALUES ('532625', '5326', '2', '马关', 'm', 'ma guan', '532625000000', '马关县', '101290603');
INSERT INTO `demo_city_code` VALUES ('532626', '5326', '2', '丘北', 'q', 'qiu bei', '532626000000', '丘北县', '101290606');
INSERT INTO `demo_city_code` VALUES ('532627', '5326', '2', '广南', 'g', 'guang nan', '532627000000', '广南县', '101290607');
INSERT INTO `demo_city_code` VALUES ('532628', '5326', '2', '富宁', 'f', 'fu ning', '532628000000', '富宁县', '101290608');
INSERT INTO `demo_city_code` VALUES ('5328', '53', '1', '西双版纳', 'x', 'xi shuang ban na', '532800000000', '西双版纳傣族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('532801', '5328', '2', '景洪', 'j', 'jing hong', '532801000000', '景洪市', '101291601');
INSERT INTO `demo_city_code` VALUES ('532822', '5328', '2', '勐海', 'm', 'meng hai', '532822000000', '勐海县', '101291603');
INSERT INTO `demo_city_code` VALUES ('532823', '5328', '2', '勐腊', 'm', 'meng la', '532823000000', '勐腊县', '101291605');
INSERT INTO `demo_city_code` VALUES ('5329', '53', '1', '大理', 'd', 'da li', '532900000000', '大理白族自治州', '101290201');
INSERT INTO `demo_city_code` VALUES ('532901', '5329', '2', '大理市', 'd', 'da li shi', '532901000000', '大理市', NULL);
INSERT INTO `demo_city_code` VALUES ('532922', '5329', '2', '漾濞', 'y', 'yang bi', '532922000000', '漾濞彝族自治县', '101290203');
INSERT INTO `demo_city_code` VALUES ('532923', '5329', '2', '祥云', 'x', 'xiang yun', '532923000000', '祥云县', '101290207');
INSERT INTO `demo_city_code` VALUES ('532924', '5329', '2', '宾川', 'b', 'bin chuan', '532924000000', '宾川县', '101290205');
INSERT INTO `demo_city_code` VALUES ('532925', '5329', '2', '弥渡', 'm', 'mi du', '532925000000', '弥渡县', '101290206');
INSERT INTO `demo_city_code` VALUES ('532926', '5329', '2', '南涧', 'n', 'nan jian', '532926000000', '南涧彝族自治县', '101290212');
INSERT INTO `demo_city_code` VALUES ('532927', '5329', '2', '巍山', 'w', 'wei shan', '532927000000', '巍山彝族回族自治县', '101290208');
INSERT INTO `demo_city_code` VALUES ('532928', '5329', '2', '永平', 'y', 'yong ping', '532928000000', '永平县', '101290204');
INSERT INTO `demo_city_code` VALUES ('532929', '5329', '2', '云龙', 'y', 'yun long', '532929000000', '云龙县', '101290202');
INSERT INTO `demo_city_code` VALUES ('532930', '5329', '2', '洱源', 'e', 'er yuan', '532930000000', '洱源县', '101290210');
INSERT INTO `demo_city_code` VALUES ('532931', '5329', '2', '剑川', 'j', 'jian chuan', '532931000000', '剑川县', '101290209');
INSERT INTO `demo_city_code` VALUES ('532932', '5329', '2', '鹤庆', 'h', 'he qing', '532932000000', '鹤庆县', '101290211');
INSERT INTO `demo_city_code` VALUES ('5331', '53', '1', '德宏', 'd', 'de hong', '533100000000', '德宏傣族景颇族自治州', '101291501');
INSERT INTO `demo_city_code` VALUES ('533102', '5331', '2', '瑞丽', 'r', 'rui li', '533102000000', '瑞丽市', '101291506');
INSERT INTO `demo_city_code` VALUES ('533103', '5331', '2', '芒市', 'm', 'mang shi', '533103000000', '芒市', NULL);
INSERT INTO `demo_city_code` VALUES ('533122', '5331', '2', '梁河', 'l', 'liang he', '533122000000', '梁河县', '101291507');
INSERT INTO `demo_city_code` VALUES ('533123', '5331', '2', '盈江', 'y', 'ying jiang', '533123000000', '盈江县', '101291504');
INSERT INTO `demo_city_code` VALUES ('533124', '5331', '2', '陇川', 'l', 'long chuan', '533124000000', '陇川县', '101291503');
INSERT INTO `demo_city_code` VALUES ('5333', '53', '1', '怒江', 'n', 'nu jiang', '533300000000', '怒江傈僳族自治州', '101291201');
INSERT INTO `demo_city_code` VALUES ('533301', '5333', '2', '泸水', 'l', 'lu shui', '533301000000', '泸水市', '101291205');
INSERT INTO `demo_city_code` VALUES ('533323', '5333', '2', '福贡', 'f', 'fu gong', '533323000000', '福贡县', '101291203');
INSERT INTO `demo_city_code` VALUES ('533324', '5333', '2', '贡山', 'g', 'gong shan', '533324000000', '贡山独龙族怒族自治县', '101291207');
INSERT INTO `demo_city_code` VALUES ('533325', '5333', '2', '兰坪', 'l', 'lan ping', '533325000000', '兰坪白族普米族自治县', '101291204');
INSERT INTO `demo_city_code` VALUES ('5334', '53', '1', '迪庆', 'd', 'di qing', '533400000000', '迪庆藏族自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('533401', '5334', '2', '香格里拉', 'x', 'xiang ge li la', '533401000000', '香格里拉市', '101291301');
INSERT INTO `demo_city_code` VALUES ('533422', '5334', '2', '德钦', 'd', 'de qin', '533422000000', '德钦县', '101291302');
INSERT INTO `demo_city_code` VALUES ('533423', '5334', '2', '维西', 'w', 'wei xi', '533423000000', '维西傈僳族自治县', '101291303');
INSERT INTO `demo_city_code` VALUES ('54', '0', '0', '西藏', 'x', 'xi zang', '540000000000', '西藏自治区', NULL);
INSERT INTO `demo_city_code` VALUES ('5401', '54', '1', '拉萨', 'l', 'la sa', '540100000000', '拉萨市', '101140101');
INSERT INTO `demo_city_code` VALUES ('540102', '5401', '2', '城关', 'c', 'cheng guan', '540102000000', '城关区', NULL);
INSERT INTO `demo_city_code` VALUES ('540103', '5401', '2', '堆龙德庆区', 'd', 'dui long de qing qu', '540103000000', '堆龙德庆区', NULL);
INSERT INTO `demo_city_code` VALUES ('540104', '5401', '2', '达孜', 'd', 'da zi', '540104000000', '达孜区', '101140107');
INSERT INTO `demo_city_code` VALUES ('540121', '5401', '2', '林周', 'l', 'lin zhou', '540121000000', '林周县', '101140104');
INSERT INTO `demo_city_code` VALUES ('540122', '5401', '2', '当雄', 'd', 'dang xiong', '540122000000', '当雄县', '101140102');
INSERT INTO `demo_city_code` VALUES ('540123', '5401', '2', '尼木', 'n', 'ni mu', '540123000000', '尼木县', '101140103');
INSERT INTO `demo_city_code` VALUES ('540124', '5401', '2', '曲水', 'q', 'qu shui', '540124000000', '曲水县', '101140106');
INSERT INTO `demo_city_code` VALUES ('540127', '5401', '2', '墨竹工卡', 'm', 'mo zhu gong ka', '540127000000', '墨竹工卡县', '101140108');
INSERT INTO `demo_city_code` VALUES ('5402', '54', '1', '日喀则', 'r', 'ri ka ze', '540200000000', '日喀则市', '101140201');
INSERT INTO `demo_city_code` VALUES ('540202', '5402', '2', '桑珠孜', 's', 'sang zhu zi', '540202000000', '桑珠孜区', NULL);
INSERT INTO `demo_city_code` VALUES ('540221', '5402', '2', '南木林', 'n', 'nan mu lin', '540221000000', '南木林县', '101140203');
INSERT INTO `demo_city_code` VALUES ('540222', '5402', '2', '江孜', 'j', 'jiang zi', '540222000000', '江孜县', '101140206');
INSERT INTO `demo_city_code` VALUES ('540223', '5402', '2', '定日', 'd', 'ding ri', '540223000000', '定日县', '101140205');
INSERT INTO `demo_city_code` VALUES ('540224', '5402', '2', '萨迦', 's', 'sa jia', '540224000000', '萨迦县', '101140213');
INSERT INTO `demo_city_code` VALUES ('540225', '5402', '2', '拉孜', 'l', 'la zi', '540225000000', '拉孜县', '101140202');
INSERT INTO `demo_city_code` VALUES ('540226', '5402', '2', '昂仁', 'a', 'ang ren', '540226000000', '昂仁县', '101140211');
INSERT INTO `demo_city_code` VALUES ('540227', '5402', '2', '谢通门', 'x', 'xie tong men', '540227000000', '谢通门县', '101140214');
INSERT INTO `demo_city_code` VALUES ('540228', '5402', '2', '白朗', 'b', 'bai lang', '540228000000', '白朗县', '101140217');
INSERT INTO `demo_city_code` VALUES ('540229', '5402', '2', '仁布', 'r', 'ren bu', '540229000000', '仁布县', '101140220');
INSERT INTO `demo_city_code` VALUES ('540230', '5402', '2', '康马', 'k', 'kang ma', '540230000000', '康马县', '101140219');
INSERT INTO `demo_city_code` VALUES ('540231', '5402', '2', '定结', 'd', 'ding jie', '540231000000', '定结县', '101140212');
INSERT INTO `demo_city_code` VALUES ('540232', '5402', '2', '仲巴', 'z', 'zhong ba', '540232000000', '仲巴县', '101140208');
INSERT INTO `demo_city_code` VALUES ('540233', '5402', '2', '亚东', 'y', 'ya dong', '540233000000', '亚东县', '101140218');
INSERT INTO `demo_city_code` VALUES ('540234', '5402', '2', '吉隆', 'j', 'ji long', '540234000000', '吉隆县', '101140210');
INSERT INTO `demo_city_code` VALUES ('540235', '5402', '2', '聂拉木', 'n', 'nie la mu', '540235000000', '聂拉木县', '101140204');
INSERT INTO `demo_city_code` VALUES ('540236', '5402', '2', '萨嘎', 's', 'sa ga', '540236000000', '萨嘎县', '101140209');
INSERT INTO `demo_city_code` VALUES ('540237', '5402', '2', '岗巴', 'g', 'gang ba', '540237000000', '岗巴县', '101140216');
INSERT INTO `demo_city_code` VALUES ('5403', '54', '1', '昌都', 'c', 'chang du', '540300000000', '昌都市', '101140501');
INSERT INTO `demo_city_code` VALUES ('540302', '5403', '2', '卡若', 'k', 'ka ruo', '540302000000', '卡若区', NULL);
INSERT INTO `demo_city_code` VALUES ('540321', '5403', '2', '江达', 'j', 'jiang da', '540321000000', '江达县', '101140509');
INSERT INTO `demo_city_code` VALUES ('540322', '5403', '2', '贡觉', 'g', 'gong jue', '540322000000', '贡觉县', '101140511');
INSERT INTO `demo_city_code` VALUES ('540323', '5403', '2', '类乌齐', 'l', 'lei wu qi', '540323000000', '类乌齐县', '101140507');
INSERT INTO `demo_city_code` VALUES ('540324', '5403', '2', '丁青', 'd', 'ding qing', '540324000000', '丁青县', '101140502');
INSERT INTO `demo_city_code` VALUES ('540325', '5403', '2', '察雅', 'c', 'cha ya', '540325000000', '察雅县', '101140510');
INSERT INTO `demo_city_code` VALUES ('540326', '5403', '2', '八宿', 'b', 'ba su', '540326000000', '八宿县', '101140508');
INSERT INTO `demo_city_code` VALUES ('540327', '5403', '2', '左贡', 'z', 'zuo gong', '540327000000', '左贡县', '101140505');
INSERT INTO `demo_city_code` VALUES ('540328', '5403', '2', '芒康', 'm', 'mang kang', '540328000000', '芒康县', '101140506');
INSERT INTO `demo_city_code` VALUES ('540329', '5403', '2', '洛隆', 'l', 'luo long', '540329000000', '洛隆县', '101140504');
INSERT INTO `demo_city_code` VALUES ('540330', '5403', '2', '边坝', 'b', 'bian ba', '540330000000', '边坝县', '101140503');
INSERT INTO `demo_city_code` VALUES ('5404', '54', '1', '林芝', 'l', 'lin zhi', '540400000000', '林芝市', '101140401');
INSERT INTO `demo_city_code` VALUES ('540402', '5404', '2', '巴宜', 'b', 'ba yi', '540402000000', '巴宜区', NULL);
INSERT INTO `demo_city_code` VALUES ('540421', '5404', '2', '工布江达', 'g', 'gong bu jiang da', '540421000000', '工布江达县', '101140405');
INSERT INTO `demo_city_code` VALUES ('540422', '5404', '2', '米林', 'm', 'mi lin', '540422000000', '米林市', '101140403');
INSERT INTO `demo_city_code` VALUES ('540423', '5404', '2', '墨脱', 'm', 'mo tuo', '540423000000', '墨脱县', '101140407');
INSERT INTO `demo_city_code` VALUES ('540424', '5404', '2', '波密', 'b', 'bo mi', '540424000000', '波密县', '101140402');
INSERT INTO `demo_city_code` VALUES ('540425', '5404', '2', '察隅', 'c', 'cha yu', '540425000000', '察隅县', '101140404');
INSERT INTO `demo_city_code` VALUES ('540426', '5404', '2', '朗县', 'l', 'lang xian', '540426000000', '朗县', '101140406');
INSERT INTO `demo_city_code` VALUES ('5405', '54', '1', '山南', 's', 'shan nan', '540500000000', '山南市', '101140301');
INSERT INTO `demo_city_code` VALUES ('540502', '5405', '2', '乃东', 'n', 'nai dong', '540502000000', '乃东区', '101140309');
INSERT INTO `demo_city_code` VALUES ('540521', '5405', '2', '扎囊', 'z', 'zha nang', '540521000000', '扎囊县', '101140303');
INSERT INTO `demo_city_code` VALUES ('540522', '5405', '2', '贡嘎', 'g', 'gong ga', '540522000000', '贡嘎县', '101140302');
INSERT INTO `demo_city_code` VALUES ('540523', '5405', '2', '桑日', 's', 'sang ri', '540523000000', '桑日县', '101140310');
INSERT INTO `demo_city_code` VALUES ('540524', '5405', '2', '琼结', 'q', 'qiong jie', '540524000000', '琼结县', '101140313');
INSERT INTO `demo_city_code` VALUES ('540525', '5405', '2', '曲松', 'q', 'qu song', '540525000000', '曲松县', '101140314');
INSERT INTO `demo_city_code` VALUES ('540526', '5405', '2', '措美', 'c', 'cuo mei', '540526000000', '措美县', '101140312');
INSERT INTO `demo_city_code` VALUES ('540527', '5405', '2', '洛扎', 'l', 'luo zha', '540527000000', '洛扎县', '101140311');
INSERT INTO `demo_city_code` VALUES ('540528', '5405', '2', '加查', 'j', 'jia cha', '540528000000', '加查县', '101140304');
INSERT INTO `demo_city_code` VALUES ('540529', '5405', '2', '隆子', 'l', 'long zi', '540529000000', '隆子县', '101140307');
INSERT INTO `demo_city_code` VALUES ('540530', '5405', '2', '错那', 'c', 'cuo na', '540530000000', '错那市', '101140306');
INSERT INTO `demo_city_code` VALUES ('540531', '5405', '2', '浪卡子', 'l', 'lang ka zi', '540531000000', '浪卡子县', '101140305');
INSERT INTO `demo_city_code` VALUES ('5406', '54', '1', '那曲', 'n', 'na qu', '540600000000', '那曲市', '101140601');
INSERT INTO `demo_city_code` VALUES ('540602', '5406', '2', '色尼', 's', 'se ni', '540602000000', '色尼区', NULL);
INSERT INTO `demo_city_code` VALUES ('540621', '5406', '2', '嘉黎', 'j', 'jia li', '540621000000', '嘉黎县', '101140603');
INSERT INTO `demo_city_code` VALUES ('540622', '5406', '2', '比如', 'b', 'bi ru', '540622000000', '比如县', '101140609');
INSERT INTO `demo_city_code` VALUES ('540623', '5406', '2', '聂荣', 'n', 'nie rong', '540623000000', '聂荣县', '101140607');
INSERT INTO `demo_city_code` VALUES ('540624', '5406', '2', '安多', 'a', 'an duo', '540624000000', '安多县', '101140605');
INSERT INTO `demo_city_code` VALUES ('540625', '5406', '2', '申扎', 's', 'shen zha', '540625000000', '申扎县', '101140703');
INSERT INTO `demo_city_code` VALUES ('540626', '5406', '2', '索县', 's', 'suo xian', '540626000000', '索县', '101140606');
INSERT INTO `demo_city_code` VALUES ('540627', '5406', '2', '班戈', 'b', 'ban ge', '540627000000', '班戈县', '101140604');
INSERT INTO `demo_city_code` VALUES ('540628', '5406', '2', '巴青', 'b', 'ba qing', '540628000000', '巴青县', '101140608');
INSERT INTO `demo_city_code` VALUES ('540629', '5406', '2', '尼玛', 'n', 'ni ma', '540629000000', '尼玛县', '101140602');
INSERT INTO `demo_city_code` VALUES ('540630', '5406', '2', '双湖', 's', 'shuang hu', '540630000000', '双湖县', '101140610');
INSERT INTO `demo_city_code` VALUES ('5425', '54', '1', '阿里', 'a', 'a li', '542500000000', '阿里地区', '101140701');
INSERT INTO `demo_city_code` VALUES ('542521', '5425', '2', '普兰', 'p', 'pu lan', '542521000000', '普兰县', '101140705');
INSERT INTO `demo_city_code` VALUES ('542522', '5425', '2', '札达', 'z', 'zha da', '542522000000', '札达县', '101140706');
INSERT INTO `demo_city_code` VALUES ('542523', '5425', '2', '噶尔', 'g', 'ga er', '542523000000', '噶尔县', '101140707');
INSERT INTO `demo_city_code` VALUES ('542524', '5425', '2', '日土', 'r', 'ri tu', '542524000000', '日土县', '101140708');
INSERT INTO `demo_city_code` VALUES ('542525', '5425', '2', '革吉', 'g', 'ge ji', '542525000000', '革吉县', '101140709');
INSERT INTO `demo_city_code` VALUES ('542526', '5425', '2', '改则', 'g', 'gai ze', '542526000000', '改则县', '101140702');
INSERT INTO `demo_city_code` VALUES ('542527', '5425', '2', '措勤', 'c', 'cuo qin', '542527000000', '措勤县', '101140710');
INSERT INTO `demo_city_code` VALUES ('61', '0', '0', '陕西', 's', 'shan xi', '610000000000', '陕西省', NULL);
INSERT INTO `demo_city_code` VALUES ('6101', '61', '1', '西安', 'x', 'xi an', '610100000000', '西安市', '101110101');
INSERT INTO `demo_city_code` VALUES ('610102', '6101', '2', '新城', 'x', 'xin cheng', '610102000000', '新城区', NULL);
INSERT INTO `demo_city_code` VALUES ('610103', '6101', '2', '碑林', 'b', 'bei lin', '610103000000', '碑林区', NULL);
INSERT INTO `demo_city_code` VALUES ('610104', '6101', '2', '莲湖', 'l', 'lian hu', '610104000000', '莲湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('610111', '6101', '2', '灞桥', 'b', 'ba qiao', '610111000000', '灞桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('610112', '6101', '2', '未央', 'w', 'wei yang', '610112000000', '未央区', NULL);
INSERT INTO `demo_city_code` VALUES ('610113', '6101', '2', '雁塔', 'y', 'yan ta', '610113000000', '雁塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('610114', '6101', '2', '阎良', 'y', 'yan liang', '610114000000', '阎良区', NULL);
INSERT INTO `demo_city_code` VALUES ('610115', '6101', '2', '临潼', 'l', 'lin tong', '610115000000', '临潼区', '101110103');
INSERT INTO `demo_city_code` VALUES ('610116', '6101', '2', '长安', 'c', 'chang an', '610116000000', '长安区', '101110102');
INSERT INTO `demo_city_code` VALUES ('610117', '6101', '2', '高陵', 'g', 'gao ling', '610117000000', '高陵区', '101110107');
INSERT INTO `demo_city_code` VALUES ('610118', '6101', '2', '鄠邑', 'h', 'hu yi', '610118000000', '鄠邑区', NULL);
INSERT INTO `demo_city_code` VALUES ('610122', '6101', '2', '蓝田', 'l', 'lan tian', '610122000000', '蓝田县', '101110104');
INSERT INTO `demo_city_code` VALUES ('610124', '6101', '2', '周至', 'z', 'zhou zhi', '610124000000', '周至县', '101110105');
INSERT INTO `demo_city_code` VALUES ('6102', '61', '1', '铜川', 't', 'tong chuan', '610200000000', '铜川市', '101111001');
INSERT INTO `demo_city_code` VALUES ('610202', '6102', '2', '王益', 'w', 'wang yi', '610202000000', '王益区', NULL);
INSERT INTO `demo_city_code` VALUES ('610203', '6102', '2', '印台', 'y', 'yin tai', '610203000000', '印台区', NULL);
INSERT INTO `demo_city_code` VALUES ('610204', '6102', '2', '耀州', 'y', 'yao zhou', '610204000000', '耀州区', '101111004');
INSERT INTO `demo_city_code` VALUES ('610222', '6102', '2', '宜君', 'y', 'yi jun', '610222000000', '宜君县', '101111003');
INSERT INTO `demo_city_code` VALUES ('6103', '61', '1', '宝鸡', 'b', 'bao ji', '610300000000', '宝鸡市', '101110901');
INSERT INTO `demo_city_code` VALUES ('610302', '6103', '2', '渭滨', 'w', 'wei bin', '610302000000', '渭滨区', NULL);
INSERT INTO `demo_city_code` VALUES ('610303', '6103', '2', '金台', 'j', 'jin tai', '610303000000', '金台区', NULL);
INSERT INTO `demo_city_code` VALUES ('610304', '6103', '2', '陈仓', 'c', 'chen cang', '610304000000', '陈仓区', '101110912');
INSERT INTO `demo_city_code` VALUES ('610305', '6103', '2', '凤翔', 'f', 'feng xiang', '610305000000', '凤翔区', '101110906');
INSERT INTO `demo_city_code` VALUES ('610323', '6103', '2', '岐山', 'q', 'qi shan', '610323000000', '岐山县', '101110905');
INSERT INTO `demo_city_code` VALUES ('610324', '6103', '2', '扶风', 'f', 'fu feng', '610324000000', '扶风县', '101110907');
INSERT INTO `demo_city_code` VALUES ('610326', '6103', '2', '眉县', 'm', 'mei xian', '610326000000', '眉县', '101110908');
INSERT INTO `demo_city_code` VALUES ('610327', '6103', '2', '陇县', 'l', 'long xian', '610327000000', '陇县', '101110911');
INSERT INTO `demo_city_code` VALUES ('610328', '6103', '2', '千阳', 'q', 'qian yang', '610328000000', '千阳县', '101110903');
INSERT INTO `demo_city_code` VALUES ('610329', '6103', '2', '麟游', 'l', 'lin you', '610329000000', '麟游县', '101110904');
INSERT INTO `demo_city_code` VALUES ('610330', '6103', '2', '凤县', 'f', 'feng xian', '610330000000', '凤县', '101110910');
INSERT INTO `demo_city_code` VALUES ('610331', '6103', '2', '太白', 't', 'tai bai', '610331000000', '太白县', '101110909');
INSERT INTO `demo_city_code` VALUES ('6104', '61', '1', '咸阳', 'x', 'xian yang', '610400000000', '咸阳市', '101110200');
INSERT INTO `demo_city_code` VALUES ('610402', '6104', '2', '秦都', 'q', 'qin du', '610402000000', '秦都区', NULL);
INSERT INTO `demo_city_code` VALUES ('610403', '6104', '2', '杨陵', 'y', 'yang ling', '610403000000', '杨陵区', NULL);
INSERT INTO `demo_city_code` VALUES ('610404', '6104', '2', '渭城', 'w', 'wei cheng', '610404000000', '渭城区', NULL);
INSERT INTO `demo_city_code` VALUES ('610422', '6104', '2', '三原', 's', 'san yuan', '610422000000', '三原县', '101110201');
INSERT INTO `demo_city_code` VALUES ('610423', '6104', '2', '泾阳', 'j', 'jing yang', '610423000000', '泾阳县', '101110205');
INSERT INTO `demo_city_code` VALUES ('610424', '6104', '2', '乾县', 'q', 'qian xian', '610424000000', '乾县', '101110207');
INSERT INTO `demo_city_code` VALUES ('610425', '6104', '2', '礼泉', 'l', 'li quan', '610425000000', '礼泉县', '101110202');
INSERT INTO `demo_city_code` VALUES ('610426', '6104', '2', '永寿', 'y', 'yong shou', '610426000000', '永寿县', '101110203');
INSERT INTO `demo_city_code` VALUES ('610428', '6104', '2', '长武', 'c', 'chang wu', '610428000000', '长武县', '101110209');
INSERT INTO `demo_city_code` VALUES ('610429', '6104', '2', '旬邑', 'x', 'xun yi', '610429000000', '旬邑县', '101110210');
INSERT INTO `demo_city_code` VALUES ('610430', '6104', '2', '淳化', 'c', 'chun hua', '610430000000', '淳化县', '101110204');
INSERT INTO `demo_city_code` VALUES ('610431', '6104', '2', '武功', 'w', 'wu gong', '610431000000', '武功县', '101110206');
INSERT INTO `demo_city_code` VALUES ('610481', '6104', '2', '兴平', 'x', 'xing ping', '610481000000', '兴平市', '101110211');
INSERT INTO `demo_city_code` VALUES ('610482', '6104', '2', '彬州', 'b', 'bin zhou', '610482000000', '彬州市', NULL);
INSERT INTO `demo_city_code` VALUES ('6105', '61', '1', '渭南', 'w', 'wei nan', '610500000000', '渭南市', '101110501');
INSERT INTO `demo_city_code` VALUES ('610502', '6105', '2', '临渭', 'l', 'lin wei', '610502000000', '临渭区', NULL);
INSERT INTO `demo_city_code` VALUES ('610503', '6105', '2', '华州', 'h', 'hua zhou', '610503000000', '华州区', NULL);
INSERT INTO `demo_city_code` VALUES ('610522', '6105', '2', '潼关', 't', 'tong guan', '610522000000', '潼关县', '101110503');
INSERT INTO `demo_city_code` VALUES ('610523', '6105', '2', '大荔', 'd', 'da li', '610523000000', '大荔县', '101110504');
INSERT INTO `demo_city_code` VALUES ('610524', '6105', '2', '合阳', 'h', 'he yang', '610524000000', '合阳县', '101110509');
INSERT INTO `demo_city_code` VALUES ('610525', '6105', '2', '澄城', 'c', 'cheng cheng', '610525000000', '澄城县', '101110508');
INSERT INTO `demo_city_code` VALUES ('610526', '6105', '2', '蒲城', 'p', 'pu cheng', '610526000000', '蒲城县', '101110507');
INSERT INTO `demo_city_code` VALUES ('610527', '6105', '2', '白水', 'b', 'bai shui', '610527000000', '白水县', '101110505');
INSERT INTO `demo_city_code` VALUES ('610528', '6105', '2', '富平', 'f', 'fu ping', '610528000000', '富平县', '101110506');
INSERT INTO `demo_city_code` VALUES ('610581', '6105', '2', '韩城', 'h', 'han cheng', '610581000000', '韩城市', '101110510');
INSERT INTO `demo_city_code` VALUES ('610582', '6105', '2', '华阴', 'h', 'hua yin', '610582000000', '华阴市', '101110511');
INSERT INTO `demo_city_code` VALUES ('6106', '61', '1', '延安', 'y', 'yan an', '610600000000', '延安市', '101110300');
INSERT INTO `demo_city_code` VALUES ('610602', '6106', '2', '宝塔', 'b', 'bao ta', '610602000000', '宝塔区', NULL);
INSERT INTO `demo_city_code` VALUES ('610603', '6106', '2', '安塞', 'a', 'an sai', '610603000000', '安塞区', '101110307');
INSERT INTO `demo_city_code` VALUES ('610621', '6106', '2', '延长', 'y', 'yan chang', '610621000000', '延长县', '101110301');
INSERT INTO `demo_city_code` VALUES ('610622', '6106', '2', '延川', 'y', 'yan chuan', '610622000000', '延川县', '101110302');
INSERT INTO `demo_city_code` VALUES ('610625', '6106', '2', '志丹', 'z', 'zhi dan', '610625000000', '志丹县', '101110306');
INSERT INTO `demo_city_code` VALUES ('610626', '6106', '2', '吴起', 'w', 'wu qi', '610626000000', '吴起县', '101110312');
INSERT INTO `demo_city_code` VALUES ('610627', '6106', '2', '甘泉', 'g', 'gan quan', '610627000000', '甘泉县', '101110308');
INSERT INTO `demo_city_code` VALUES ('610628', '6106', '2', '富县', 'f', 'fu xian', '610628000000', '富县', '101110305');
INSERT INTO `demo_city_code` VALUES ('610629', '6106', '2', '洛川', 'l', 'luo chuan', '610629000000', '洛川县', '101110309');
INSERT INTO `demo_city_code` VALUES ('610630', '6106', '2', '宜川', 'y', 'yi chuan', '610630000000', '宜川县', '101110304');
INSERT INTO `demo_city_code` VALUES ('610631', '6106', '2', '黄龙', 'h', 'huang long', '610631000000', '黄龙县', '101110311');
INSERT INTO `demo_city_code` VALUES ('610632', '6106', '2', '黄陵', 'h', 'huang ling', '610632000000', '黄陵县', '101110310');
INSERT INTO `demo_city_code` VALUES ('610681', '6106', '2', '子长', 'z', 'zi chang', '610681000000', '子长市', '101110303');
INSERT INTO `demo_city_code` VALUES ('6107', '61', '1', '汉中', 'h', 'han zhong', '610700000000', '汉中市', '101110801');
INSERT INTO `demo_city_code` VALUES ('610702', '6107', '2', '汉台', 'h', 'han tai', '610702000000', '汉台区', NULL);
INSERT INTO `demo_city_code` VALUES ('610703', '6107', '2', '南郑', 'n', 'nan zheng', '610703000000', '南郑区', '101110810');
INSERT INTO `demo_city_code` VALUES ('610722', '6107', '2', '城固', 'c', 'cheng gu', '610722000000', '城固县', '101110806');
INSERT INTO `demo_city_code` VALUES ('610723', '6107', '2', '洋县', 'y', 'yang xian', '610723000000', '洋县', '101110805');
INSERT INTO `demo_city_code` VALUES ('610724', '6107', '2', '西乡', 'x', 'xi xiang', '610724000000', '西乡县', '101110807');
INSERT INTO `demo_city_code` VALUES ('610725', '6107', '2', '勉县', 'm', 'mian xian', '610725000000', '勉县', '101110803');
INSERT INTO `demo_city_code` VALUES ('610726', '6107', '2', '宁强', 'n', 'ning qiang', '610726000000', '宁强县', '101110809');
INSERT INTO `demo_city_code` VALUES ('610727', '6107', '2', '略阳', 'l', 'lue yang', '610727000000', '略阳县', '101110802');
INSERT INTO `demo_city_code` VALUES ('610728', '6107', '2', '镇巴', 'z', 'zhen ba', '610728000000', '镇巴县', '101110811');
INSERT INTO `demo_city_code` VALUES ('610729', '6107', '2', '留坝', 'l', 'liu ba', '610729000000', '留坝县', '101110804');
INSERT INTO `demo_city_code` VALUES ('610730', '6107', '2', '佛坪', 'f', 'fo ping', '610730000000', '佛坪县', '101110808');
INSERT INTO `demo_city_code` VALUES ('6108', '61', '1', '榆林', 'y', 'yu lin', '610800000000', '榆林市', '101110401');
INSERT INTO `demo_city_code` VALUES ('610802', '6108', '2', '榆阳', 'y', 'yu yang', '610802000000', '榆阳区', '101110413');
INSERT INTO `demo_city_code` VALUES ('610803', '6108', '2', '横山', 'h', 'heng shan', '610803000000', '横山区', '101110407');
INSERT INTO `demo_city_code` VALUES ('610822', '6108', '2', '府谷', 'f', 'fu gu', '610822000000', '府谷县', '101110402');
INSERT INTO `demo_city_code` VALUES ('610824', '6108', '2', '靖边', 'j', 'jing bian', '610824000000', '靖边县', '101110406');
INSERT INTO `demo_city_code` VALUES ('610825', '6108', '2', '定边', 'd', 'ding bian', '610825000000', '定边县', '101110405');
INSERT INTO `demo_city_code` VALUES ('610826', '6108', '2', '绥德', 's', 'sui de', '610826000000', '绥德县', '101110410');
INSERT INTO `demo_city_code` VALUES ('610827', '6108', '2', '米脂', 'm', 'mi zhi', '610827000000', '米脂县', '101110408');
INSERT INTO `demo_city_code` VALUES ('610828', '6108', '2', '佳县', 'j', 'jia xian', '610828000000', '佳县', '101110404');
INSERT INTO `demo_city_code` VALUES ('610829', '6108', '2', '吴堡', 'w', 'wu bu', '610829000000', '吴堡县', '101110411');
INSERT INTO `demo_city_code` VALUES ('610830', '6108', '2', '清涧', 'q', 'qing jian', '610830000000', '清涧县', '101110412');
INSERT INTO `demo_city_code` VALUES ('610831', '6108', '2', '子洲', 'z', 'zi zhou', '610831000000', '子洲县', '101110409');
INSERT INTO `demo_city_code` VALUES ('610881', '6108', '2', '神木', 's', 'shen mu', '610881000000', '神木市', '101110403');
INSERT INTO `demo_city_code` VALUES ('6109', '61', '1', '安康', 'a', 'an kang', '610900000000', '安康市', '101110701');
INSERT INTO `demo_city_code` VALUES ('610902', '6109', '2', '汉滨', 'h', 'han bin', '610902000000', '汉滨区', NULL);
INSERT INTO `demo_city_code` VALUES ('610921', '6109', '2', '汉阴', 'h', 'han yin', '610921000000', '汉阴县', '101110704');
INSERT INTO `demo_city_code` VALUES ('610922', '6109', '2', '石泉', 's', 'shi quan', '610922000000', '石泉县', '101110703');
INSERT INTO `demo_city_code` VALUES ('610923', '6109', '2', '宁陕', 'n', 'ning shan', '610923000000', '宁陕县', '101110710');
INSERT INTO `demo_city_code` VALUES ('610924', '6109', '2', '紫阳', 'z', 'zi yang', '610924000000', '紫阳县', '101110702');
INSERT INTO `demo_city_code` VALUES ('610925', '6109', '2', '岚皋', 'l', 'lan gao', '610925000000', '岚皋县', '101110706');
INSERT INTO `demo_city_code` VALUES ('610926', '6109', '2', '平利', 'p', 'ping li', '610926000000', '平利县', '101110707');
INSERT INTO `demo_city_code` VALUES ('610927', '6109', '2', '镇坪', 'z', 'zhen ping', '610927000000', '镇坪县', '101110709');
INSERT INTO `demo_city_code` VALUES ('610929', '6109', '2', '白河', 'b', 'bai he', '610929000000', '白河县', '101110708');
INSERT INTO `demo_city_code` VALUES ('610981', '6109', '2', '旬阳', 'x', 'xun yang', '610981000000', '旬阳市', '101110705');
INSERT INTO `demo_city_code` VALUES ('6110', '61', '1', '商洛', 's', 'shang luo', '611000000000', '商洛市', '101110601');
INSERT INTO `demo_city_code` VALUES ('611002', '6110', '2', '商州', 's', 'shang zhou', '611002000000', '商州区', '101110604');
INSERT INTO `demo_city_code` VALUES ('611021', '6110', '2', '洛南', 'l', 'luo nan', '611021000000', '洛南县', '101110602');
INSERT INTO `demo_city_code` VALUES ('611022', '6110', '2', '丹凤', 'd', 'dan feng', '611022000000', '丹凤县', '101110606');
INSERT INTO `demo_city_code` VALUES ('611023', '6110', '2', '商南', 's', 'shang nan', '611023000000', '商南县', '101110607');
INSERT INTO `demo_city_code` VALUES ('611024', '6110', '2', '山阳', 's', 'shan yang', '611024000000', '山阳县', '101110608');
INSERT INTO `demo_city_code` VALUES ('611025', '6110', '2', '镇安', 'z', 'zhen an', '611025000000', '镇安县', '101110605');
INSERT INTO `demo_city_code` VALUES ('611026', '6110', '2', '柞水', 'z', 'zha shui', '611026000000', '柞水县', '101110603');
INSERT INTO `demo_city_code` VALUES ('62', '0', '0', '甘肃', 'g', 'gan su', '620000000000', '甘肃省', NULL);
INSERT INTO `demo_city_code` VALUES ('6201', '62', '1', '兰州', 'l', 'lan zhou', '620100000000', '兰州市', '101160101');
INSERT INTO `demo_city_code` VALUES ('620102', '6201', '2', '城关', 'c', 'cheng guan', '620102000000', '城关区', NULL);
INSERT INTO `demo_city_code` VALUES ('620103', '6201', '2', '七里河', 'q', 'qi li he', '620103000000', '七里河区', NULL);
INSERT INTO `demo_city_code` VALUES ('620104', '6201', '2', '西固', 'x', 'xi gu', '620104000000', '西固区', NULL);
INSERT INTO `demo_city_code` VALUES ('620105', '6201', '2', '安宁', 'a', 'an ning', '620105000000', '安宁区', '101290112');
INSERT INTO `demo_city_code` VALUES ('620111', '6201', '2', '红古', 'h', 'hong gu', '620111000000', '红古区', NULL);
INSERT INTO `demo_city_code` VALUES ('620121', '6201', '2', '永登', 'y', 'yong deng', '620121000000', '永登县', '101160103');
INSERT INTO `demo_city_code` VALUES ('620122', '6201', '2', '皋兰', 'g', 'gao lan', '620122000000', '皋兰县', '101160102');
INSERT INTO `demo_city_code` VALUES ('620123', '6201', '2', '榆中', 'y', 'yu zhong', '620123000000', '榆中县', '101160104');
INSERT INTO `demo_city_code` VALUES ('6202', '62', '1', '嘉峪关', 'j', 'jia yu guan', '620200000000', '嘉峪关市', '101161401');
INSERT INTO `demo_city_code` VALUES ('620200', '6202', '2', '嘉峪关', 'j', 'jia yu guan', '620200000000', '嘉峪关市', '101161401');
INSERT INTO `demo_city_code` VALUES ('6203', '62', '1', '金昌', 'j', 'jin chang', '620300000000', '金昌市', '101160601');
INSERT INTO `demo_city_code` VALUES ('620302', '6203', '2', '金川', 'j', 'jin chuan', '620302000000', '金川区', '101271907');
INSERT INTO `demo_city_code` VALUES ('620321', '6203', '2', '永昌', 'y', 'yong chang', '620321000000', '永昌县', '101160602');
INSERT INTO `demo_city_code` VALUES ('6204', '62', '1', '白银', 'b', 'bai yin', '620400000000', '白银市', '101161301');
INSERT INTO `demo_city_code` VALUES ('620402', '6204', '2', '白银区', 'b', 'bai yin qu', '620402000000', '白银区', NULL);
INSERT INTO `demo_city_code` VALUES ('620403', '6204', '2', '平川', 'p', 'ping chuan', '620403000000', '平川区', '101161304');
INSERT INTO `demo_city_code` VALUES ('620421', '6204', '2', '靖远', 'j', 'jing yuan', '620421000000', '靖远县', '101161302');
INSERT INTO `demo_city_code` VALUES ('620422', '6204', '2', '会宁', 'h', 'hui ning', '620422000000', '会宁县', '101161303');
INSERT INTO `demo_city_code` VALUES ('620423', '6204', '2', '景泰', 'j', 'jing tai', '620423000000', '景泰县', '101161305');
INSERT INTO `demo_city_code` VALUES ('6205', '62', '1', '天水', 't', 'tian shui', '620500000000', '天水市', '101160901');
INSERT INTO `demo_city_code` VALUES ('620502', '6205', '2', '秦州', 'q', 'qin zhou', '620502000000', '秦州区', NULL);
INSERT INTO `demo_city_code` VALUES ('620503', '6205', '2', '麦积', 'm', 'mai ji', '620503000000', '麦积区', '101160908');
INSERT INTO `demo_city_code` VALUES ('620521', '6205', '2', '清水', 'q', 'qing shui', '620521000000', '清水县', '101160903');
INSERT INTO `demo_city_code` VALUES ('620522', '6205', '2', '秦安', 'q', 'qin an', '620522000000', '秦安县', '101160904');
INSERT INTO `demo_city_code` VALUES ('620523', '6205', '2', '甘谷', 'g', 'gan gu', '620523000000', '甘谷县', '101160905');
INSERT INTO `demo_city_code` VALUES ('620524', '6205', '2', '武山', 'w', 'wu shan', '620524000000', '武山县', '101160906');
INSERT INTO `demo_city_code` VALUES ('620525', '6205', '2', '张家川', 'z', 'zhang jia chuan', '620525000000', '张家川回族自治县', '101160907');
INSERT INTO `demo_city_code` VALUES ('6206', '62', '1', '武威', 'w', 'wu wei', '620600000000', '武威市', '101160501');
INSERT INTO `demo_city_code` VALUES ('620602', '6206', '2', '凉州', 'l', 'liang zhou', '620602000000', '凉州区', NULL);
INSERT INTO `demo_city_code` VALUES ('620621', '6206', '2', '民勤', 'm', 'min qin', '620621000000', '民勤县', '101160502');
INSERT INTO `demo_city_code` VALUES ('620622', '6206', '2', '古浪', 'g', 'gu lang', '620622000000', '古浪县', '101160503');
INSERT INTO `demo_city_code` VALUES ('620623', '6206', '2', '天祝', 't', 'tian zhu', '620623000000', '天祝藏族自治县', '101160505');
INSERT INTO `demo_city_code` VALUES ('6207', '62', '1', '张掖', 'z', 'zhang ye', '620700000000', '张掖市', '101160701');
INSERT INTO `demo_city_code` VALUES ('620702', '6207', '2', '甘州', 'g', 'gan zhou', '620702000000', '甘州区', NULL);
INSERT INTO `demo_city_code` VALUES ('620721', '6207', '2', '肃南', 's', 'su nan', '620721000000', '肃南裕固族自治县', '101160702');
INSERT INTO `demo_city_code` VALUES ('620722', '6207', '2', '民乐', 'm', 'min le', '620722000000', '民乐县', '101160703');
INSERT INTO `demo_city_code` VALUES ('620723', '6207', '2', '临泽', 'l', 'lin ze', '620723000000', '临泽县', '101160704');
INSERT INTO `demo_city_code` VALUES ('620724', '6207', '2', '高台', 'g', 'gao tai', '620724000000', '高台县', '101160705');
INSERT INTO `demo_city_code` VALUES ('620725', '6207', '2', '山丹', 's', 'shan dan', '620725000000', '山丹县', '101160706');
INSERT INTO `demo_city_code` VALUES ('6208', '62', '1', '平凉', 'p', 'ping liang', '620800000000', '平凉市', '101160301');
INSERT INTO `demo_city_code` VALUES ('620802', '6208', '2', '崆峒', 'k', 'kong tong', '620802000000', '崆峒区', '101160308');
INSERT INTO `demo_city_code` VALUES ('620821', '6208', '2', '泾川', 'j', 'jing chuan', '620821000000', '泾川县', '101160302');
INSERT INTO `demo_city_code` VALUES ('620822', '6208', '2', '灵台', 'l', 'ling tai', '620822000000', '灵台县', '101160303');
INSERT INTO `demo_city_code` VALUES ('620823', '6208', '2', '崇信', 'c', 'chong xin', '620823000000', '崇信县', '101160304');
INSERT INTO `demo_city_code` VALUES ('620825', '6208', '2', '庄浪', 'z', 'zhuang lang', '620825000000', '庄浪县', '101160306');
INSERT INTO `demo_city_code` VALUES ('620826', '6208', '2', '静宁', 'j', 'jing ning', '620826000000', '静宁县', '101160307');
INSERT INTO `demo_city_code` VALUES ('620881', '6208', '2', '华亭', 'h', 'hua ting', '620881000000', '华亭市', '101160305');
INSERT INTO `demo_city_code` VALUES ('6209', '62', '1', '酒泉', 'j', 'jiu quan', '620900000000', '酒泉市', '101160801');
INSERT INTO `demo_city_code` VALUES ('620902', '6209', '2', '肃州', 's', 'su zhou', '620902000000', '肃州区', NULL);
INSERT INTO `demo_city_code` VALUES ('620921', '6209', '2', '金塔', 'j', 'jin ta', '620921000000', '金塔县', '101160803');
INSERT INTO `demo_city_code` VALUES ('620922', '6209', '2', '瓜州', 'g', 'gua zhou', '620922000000', '瓜州县', '101160805');
INSERT INTO `demo_city_code` VALUES ('620923', '6209', '2', '肃北', 's', 'su bei', '620923000000', '肃北蒙古族自治县', '101160806');
INSERT INTO `demo_city_code` VALUES ('620924', '6209', '2', '阿克塞', 'a', 'a ke sai', '620924000000', '阿克塞哈萨克族自治县', '101160804');
INSERT INTO `demo_city_code` VALUES ('620981', '6209', '2', '玉门', 'y', 'yu men', '620981000000', '玉门市', '101160807');
INSERT INTO `demo_city_code` VALUES ('620982', '6209', '2', '敦煌', 'd', 'dun huang', '620982000000', '敦煌市', '101160808');
INSERT INTO `demo_city_code` VALUES ('6210', '62', '1', '庆阳', 'q', 'qing yang', '621000000000', '庆阳市', NULL);
INSERT INTO `demo_city_code` VALUES ('621002', '6210', '2', '西峰', 'x', 'xi feng', '621002000000', '西峰区', '101160401');
INSERT INTO `demo_city_code` VALUES ('621021', '6210', '2', '庆城', 'q', 'qing cheng', '621021000000', '庆城县', '101160409');
INSERT INTO `demo_city_code` VALUES ('621022', '6210', '2', '环县', 'h', 'huan xian', '621022000000', '环县', '101160403');
INSERT INTO `demo_city_code` VALUES ('621023', '6210', '2', '华池', 'h', 'hua chi', '621023000000', '华池县', '101160404');
INSERT INTO `demo_city_code` VALUES ('621024', '6210', '2', '合水', 'h', 'he shui', '621024000000', '合水县', '101160405');
INSERT INTO `demo_city_code` VALUES ('621025', '6210', '2', '正宁', 'z', 'zheng ning', '621025000000', '正宁县', '101160406');
INSERT INTO `demo_city_code` VALUES ('621026', '6210', '2', '宁县', 'n', 'ning xian', '621026000000', '宁县', '101160407');
INSERT INTO `demo_city_code` VALUES ('621027', '6210', '2', '镇原', 'z', 'zhen yuan', '621027000000', '镇原县', '101160408');
INSERT INTO `demo_city_code` VALUES ('6211', '62', '1', '定西', 'd', 'ding xi', '621100000000', '定西市', '101160201');
INSERT INTO `demo_city_code` VALUES ('621102', '6211', '2', '安定', 'a', 'an ding', '621102000000', '安定区', '101160208');
INSERT INTO `demo_city_code` VALUES ('621121', '6211', '2', '通渭', 't', 'tong wei', '621121000000', '通渭县', '101160202');
INSERT INTO `demo_city_code` VALUES ('621122', '6211', '2', '陇西', 'l', 'long xi', '621122000000', '陇西县', '101160203');
INSERT INTO `demo_city_code` VALUES ('621123', '6211', '2', '渭源', 'w', 'wei yuan', '621123000000', '渭源县', '101160204');
INSERT INTO `demo_city_code` VALUES ('621124', '6211', '2', '临洮', 'l', 'lin tao', '621124000000', '临洮县', '101160205');
INSERT INTO `demo_city_code` VALUES ('621125', '6211', '2', '漳县', 'z', 'zhang xian', '621125000000', '漳县', '101160206');
INSERT INTO `demo_city_code` VALUES ('621126', '6211', '2', '岷县', 'm', 'min xian', '621126000000', '岷县', '101160207');
INSERT INTO `demo_city_code` VALUES ('6212', '62', '1', '陇南', 'l', 'long nan', '621200000000', '陇南市', NULL);
INSERT INTO `demo_city_code` VALUES ('621202', '6212', '2', '武都', 'w', 'wu du', '621202000000', '武都区', '101161001');
INSERT INTO `demo_city_code` VALUES ('621221', '6212', '2', '成县', 'c', 'cheng xian', '621221000000', '成县', '101161002');
INSERT INTO `demo_city_code` VALUES ('621222', '6212', '2', '文县', 'w', 'wen xian', '621222000000', '文县', '101161003');
INSERT INTO `demo_city_code` VALUES ('621223', '6212', '2', '宕昌', 't', 'tan chang', '621223000000', '宕昌县', '101161004');
INSERT INTO `demo_city_code` VALUES ('621224', '6212', '2', '康县', 'k', 'kang xian', '621224000000', '康县', '101161005');
INSERT INTO `demo_city_code` VALUES ('621225', '6212', '2', '西和', 'x', 'xi he', '621225000000', '西和县', '101161006');
INSERT INTO `demo_city_code` VALUES ('621226', '6212', '2', '礼县', 'l', 'li xian', '621226000000', '礼县', '101161007');
INSERT INTO `demo_city_code` VALUES ('621227', '6212', '2', '徽县', 'h', 'hui xian', '621227000000', '徽县', '101161008');
INSERT INTO `demo_city_code` VALUES ('621228', '6212', '2', '两当', 'l', 'liang dang', '621228000000', '两当县', '101161009');
INSERT INTO `demo_city_code` VALUES ('6229', '62', '1', '临夏', 'l', 'lin xia', '622900000000', '临夏回族自治州', '101161101');
INSERT INTO `demo_city_code` VALUES ('622901', '6229', '2', '临夏市', 'l', 'lin xia shi', '622901000000', '临夏市', NULL);
INSERT INTO `demo_city_code` VALUES ('622921', '6229', '2', '临夏县', 'l', 'lin xia xian', '622921000000', '临夏县', NULL);
INSERT INTO `demo_city_code` VALUES ('622922', '6229', '2', '康乐', 'k', 'kang le', '622922000000', '康乐县', '101161102');
INSERT INTO `demo_city_code` VALUES ('622923', '6229', '2', '永靖', 'y', 'yong jing', '622923000000', '永靖县', '101161103');
INSERT INTO `demo_city_code` VALUES ('622924', '6229', '2', '广河', 'g', 'guang he', '622924000000', '广河县', '101161104');
INSERT INTO `demo_city_code` VALUES ('622925', '6229', '2', '和政', 'h', 'he zheng', '622925000000', '和政县', '101161105');
INSERT INTO `demo_city_code` VALUES ('622926', '6229', '2', '东乡族自治县', 'd', 'dong xiang zu zi zhi xian', '622926000000', '东乡族自治县', NULL);
INSERT INTO `demo_city_code` VALUES ('622927', '6229', '2', '积石山', 'j', 'ji shi shan', '622927000000', '积石山保安族东乡族撒拉族自治县', '101161107');
INSERT INTO `demo_city_code` VALUES ('6230', '62', '1', '甘南', 'g', 'gan nan', '623000000000', '甘南藏族自治州', '101050204');
INSERT INTO `demo_city_code` VALUES ('623001', '6230', '2', '合作', 'h', 'he zuo', '623001000000', '合作市', '101161201');
INSERT INTO `demo_city_code` VALUES ('623021', '6230', '2', '临潭', 'l', 'lin tan', '623021000000', '临潭县', '101161202');
INSERT INTO `demo_city_code` VALUES ('623022', '6230', '2', '卓尼', 'z', 'zhuo ni', '623022000000', '卓尼县', '101161203');
INSERT INTO `demo_city_code` VALUES ('623023', '6230', '2', '舟曲', 'z', 'zhou qu', '623023000000', '舟曲县', '101161204');
INSERT INTO `demo_city_code` VALUES ('623024', '6230', '2', '迭部', 'd', 'die bu', '623024000000', '迭部县', '101161205');
INSERT INTO `demo_city_code` VALUES ('623025', '6230', '2', '玛曲', 'm', 'ma qu', '623025000000', '玛曲县', '101161206');
INSERT INTO `demo_city_code` VALUES ('623026', '6230', '2', '碌曲', 'l', 'lu qu', '623026000000', '碌曲县', '101161207');
INSERT INTO `demo_city_code` VALUES ('623027', '6230', '2', '夏河', 'x', 'xia he', '623027000000', '夏河县', '101161208');
INSERT INTO `demo_city_code` VALUES ('63', '0', '0', '青海', 'q', 'qing hai', '630000000000', '青海省', NULL);
INSERT INTO `demo_city_code` VALUES ('6301', '63', '1', '西宁', 'x', 'xi ning', '630100000000', '西宁市', '101150101');
INSERT INTO `demo_city_code` VALUES ('630102', '6301', '2', '城东', 'c', 'cheng dong', '630102000000', '城东区', NULL);
INSERT INTO `demo_city_code` VALUES ('630103', '6301', '2', '城中', 'c', 'cheng zhong', '630103000000', '城中区', NULL);
INSERT INTO `demo_city_code` VALUES ('630104', '6301', '2', '城西', 'c', 'cheng xi', '630104000000', '城西区', NULL);
INSERT INTO `demo_city_code` VALUES ('630105', '6301', '2', '城北', 'c', 'cheng bei', '630105000000', '城北区', NULL);
INSERT INTO `demo_city_code` VALUES ('630106', '6301', '2', '湟中', 'h', 'huang zhong', '630106000000', '湟中区', '101150104');
INSERT INTO `demo_city_code` VALUES ('630121', '6301', '2', '大通', 'd', 'da tong', '630121000000', '大通回族土族自治县', '101150102');
INSERT INTO `demo_city_code` VALUES ('630123', '6301', '2', '湟源', 'h', 'huang yuan', '630123000000', '湟源县', '101150103');
INSERT INTO `demo_city_code` VALUES ('6302', '63', '1', '海东', 'h', 'hai dong', '630200000000', '海东市', '101150201');
INSERT INTO `demo_city_code` VALUES ('630202', '6302', '2', '乐都', 'l', 'le du', '630202000000', '乐都区', '101150202');
INSERT INTO `demo_city_code` VALUES ('630203', '6302', '2', '平安', 'p', 'ping an', '630203000000', '平安区', '101150208');
INSERT INTO `demo_city_code` VALUES ('630222', '6302', '2', '民和', 'm', 'min he', '630222000000', '民和回族土族自治县', '101150203');
INSERT INTO `demo_city_code` VALUES ('630223', '6302', '2', '互助', 'h', 'hu zhu', '630223000000', '互助土族自治县', '101150204');
INSERT INTO `demo_city_code` VALUES ('630224', '6302', '2', '化隆', 'h', 'hua long', '630224000000', '化隆回族自治县', '101150205');
INSERT INTO `demo_city_code` VALUES ('630225', '6302', '2', '循化', 'x', 'xun hua', '630225000000', '循化撒拉族自治县', '101150206');
INSERT INTO `demo_city_code` VALUES ('6322', '63', '1', '海北', 'h', 'hai bei', '632200000000', '海北藏族自治州', '101150801');
INSERT INTO `demo_city_code` VALUES ('632221', '6322', '2', '门源', 'm', 'men yuan', '632221000000', '门源回族自治县', '101150802');
INSERT INTO `demo_city_code` VALUES ('632222', '6322', '2', '祁连', 'q', 'qi lian', '632222000000', '祁连县', '101150803');
INSERT INTO `demo_city_code` VALUES ('632223', '6322', '2', '海晏', 'h', 'hai yan', '632223000000', '海晏县', '101150804');
INSERT INTO `demo_city_code` VALUES ('632224', '6322', '2', '刚察', 'g', 'gang cha', '632224000000', '刚察县', '101150806');
INSERT INTO `demo_city_code` VALUES ('6323', '63', '1', '黄南', 'h', 'huang nan', '632300000000', '黄南藏族自治州', '101150301');
INSERT INTO `demo_city_code` VALUES ('632301', '6323', '2', '同仁', 't', 'tong ren', '632301000000', '同仁市', '101150305');
INSERT INTO `demo_city_code` VALUES ('632322', '6323', '2', '尖扎', 'j', 'jian zha', '632322000000', '尖扎县', '101150302');
INSERT INTO `demo_city_code` VALUES ('632323', '6323', '2', '泽库', 'z', 'ze ku', '632323000000', '泽库县', '101150303');
INSERT INTO `demo_city_code` VALUES ('632324', '6323', '2', '河南', 'h', 'he nan', '632324000000', '河南蒙古族自治县', '101080706');
INSERT INTO `demo_city_code` VALUES ('6325', '63', '1', '海南', 'h', 'hai nan', '632500000000', '海南藏族自治州', '101150401');
INSERT INTO `demo_city_code` VALUES ('632521', '6325', '2', '共和', 'g', 'gong he', '632521000000', '共和县', '101150409');
INSERT INTO `demo_city_code` VALUES ('632522', '6325', '2', '同德', 't', 'tong de', '632522000000', '同德县', '101150408');
INSERT INTO `demo_city_code` VALUES ('632523', '6325', '2', '贵德', 'g', 'gui de', '632523000000', '贵德县', '101150404');
INSERT INTO `demo_city_code` VALUES ('632524', '6325', '2', '兴海', 'x', 'xing hai', '632524000000', '兴海县', '101150406');
INSERT INTO `demo_city_code` VALUES ('632525', '6325', '2', '贵南', 'g', 'gui nan', '632525000000', '贵南县', '101150407');
INSERT INTO `demo_city_code` VALUES ('6326', '63', '1', '果洛', 'g', 'guo luo', '632600000000', '果洛藏族自治州', '101150501');
INSERT INTO `demo_city_code` VALUES ('632621', '6326', '2', '玛沁', 'm', 'ma qin', '632621000000', '玛沁县', '101150508');
INSERT INTO `demo_city_code` VALUES ('632622', '6326', '2', '班玛', 'b', 'ban ma', '632622000000', '班玛县', '101150502');
INSERT INTO `demo_city_code` VALUES ('632623', '6326', '2', '甘德', 'g', 'gan de', '632623000000', '甘德县', '101150503');
INSERT INTO `demo_city_code` VALUES ('632624', '6326', '2', '达日', 'd', 'da ri', '632624000000', '达日县', '101150504');
INSERT INTO `demo_city_code` VALUES ('632625', '6326', '2', '久治', 'j', 'jiu zhi', '632625000000', '久治县', '101150505');
INSERT INTO `demo_city_code` VALUES ('632626', '6326', '2', '玛多', 'm', 'ma duo', '632626000000', '玛多县', '101150506');
INSERT INTO `demo_city_code` VALUES ('6327', '63', '1', '玉树', 'y', 'yu shu', '632700000000', '玉树藏族自治州', '101150601');
INSERT INTO `demo_city_code` VALUES ('632701', '6327', '2', '玉树市', 'y', 'yu shu shi', '632701000000', '玉树市', NULL);
INSERT INTO `demo_city_code` VALUES ('632722', '6327', '2', '杂多', 'z', 'za duo', '632722000000', '杂多县', '101150604');
INSERT INTO `demo_city_code` VALUES ('632723', '6327', '2', '称多', 'c', 'cheng duo', '632723000000', '称多县', '101150602');
INSERT INTO `demo_city_code` VALUES ('632724', '6327', '2', '治多', 'z', 'zhi duo', '632724000000', '治多县', '101150603');
INSERT INTO `demo_city_code` VALUES ('632725', '6327', '2', '囊谦', 'n', 'nang qian', '632725000000', '囊谦县', '101150605');
INSERT INTO `demo_city_code` VALUES ('632726', '6327', '2', '曲麻莱', 'q', 'qu ma lai', '632726000000', '曲麻莱县', '101150606');
INSERT INTO `demo_city_code` VALUES ('6328', '63', '1', '海西', 'h', 'hai xi', '632800000000', '海西蒙古族藏族自治州', '101150701');
INSERT INTO `demo_city_code` VALUES ('632801', '6328', '2', '格尔木', 'g', 'ge er mu', '632801000000', '格尔木市', '101150901');
INSERT INTO `demo_city_code` VALUES ('632802', '6328', '2', '德令哈', 'd', 'de ling ha', '632802000000', '德令哈市', '101150716');
INSERT INTO `demo_city_code` VALUES ('632803', '6328', '2', '茫崖', 'm', 'mang ya', '632803000000', '茫崖市', '101150712');
INSERT INTO `demo_city_code` VALUES ('632821', '6328', '2', '乌兰', 'w', 'wu lan', '632821000000', '乌兰县', '101150709');
INSERT INTO `demo_city_code` VALUES ('632822', '6328', '2', '都兰', 'd', 'du lan', '632822000000', '都兰县', '101150902');
INSERT INTO `demo_city_code` VALUES ('632823', '6328', '2', '天峻', 't', 'tian jun', '632823000000', '天峻县', '101150708');
INSERT INTO `demo_city_code` VALUES ('632825', '6328', '2', '大柴旦行政委员会', 'd', 'da chai dan xing zheng wei yuan hui', '632825000000', '大柴旦行政委员会', NULL);
INSERT INTO `demo_city_code` VALUES ('64', '0', '0', '宁夏', 'n', 'ning xia', '640000000000', '宁夏回族自治区', NULL);
INSERT INTO `demo_city_code` VALUES ('6401', '64', '1', '银川', 'y', 'yin chuan', '640100000000', '银川市', '101170101');
INSERT INTO `demo_city_code` VALUES ('640104', '6401', '2', '兴庆', 'x', 'xing qing', '640104000000', '兴庆区', NULL);
INSERT INTO `demo_city_code` VALUES ('640105', '6401', '2', '西夏', 'x', 'xi xia', '640105000000', '西夏区', NULL);
INSERT INTO `demo_city_code` VALUES ('640106', '6401', '2', '金凤', 'j', 'jin feng', '640106000000', '金凤区', NULL);
INSERT INTO `demo_city_code` VALUES ('640121', '6401', '2', '永宁', 'y', 'yong ning', '640121000000', '永宁县', '101170102');
INSERT INTO `demo_city_code` VALUES ('640122', '6401', '2', '贺兰', 'h', 'he lan', '640122000000', '贺兰县', '101170104');
INSERT INTO `demo_city_code` VALUES ('640181', '6401', '2', '灵武', 'l', 'ling wu', '640181000000', '灵武市', '101170103');
INSERT INTO `demo_city_code` VALUES ('6402', '64', '1', '石嘴山', 's', 'shi zui shan', '640200000000', '石嘴山市', '101170201');
INSERT INTO `demo_city_code` VALUES ('640202', '6402', '2', '大武口', 'd', 'da wu kou', '640202000000', '大武口区', NULL);
INSERT INTO `demo_city_code` VALUES ('640205', '6402', '2', '惠农', 'h', 'hui nong', '640205000000', '惠农区', '101170202');
INSERT INTO `demo_city_code` VALUES ('640221', '6402', '2', '平罗', 'p', 'ping luo', '640221000000', '平罗县', '101170203');
INSERT INTO `demo_city_code` VALUES ('6403', '64', '1', '吴忠', 'w', 'wu zhong', '640300000000', '吴忠市', '101170301');
INSERT INTO `demo_city_code` VALUES ('640302', '6403', '2', '利通', 'l', 'li tong', '640302000000', '利通区', NULL);
INSERT INTO `demo_city_code` VALUES ('640303', '6403', '2', '红寺堡', 'h', 'hong si bu', '640303000000', '红寺堡区', NULL);
INSERT INTO `demo_city_code` VALUES ('640323', '6403', '2', '盐池', 'y', 'yan chi', '640323000000', '盐池县', '101170303');
INSERT INTO `demo_city_code` VALUES ('640324', '6403', '2', '同心', 't', 'tong xin', '640324000000', '同心县', '101170302');
INSERT INTO `demo_city_code` VALUES ('640381', '6403', '2', '青铜峡', 'q', 'qing tong xia', '640381000000', '青铜峡市', '101170306');
INSERT INTO `demo_city_code` VALUES ('6404', '64', '1', '固原', 'g', 'gu yuan', '640400000000', '固原市', '101170401');
INSERT INTO `demo_city_code` VALUES ('640402', '6404', '2', '原州', 'y', 'yuan zhou', '640402000000', '原州区', NULL);
INSERT INTO `demo_city_code` VALUES ('640422', '6404', '2', '西吉', 'x', 'xi ji', '640422000000', '西吉县', '101170402');
INSERT INTO `demo_city_code` VALUES ('640423', '6404', '2', '隆德', 'l', 'long de', '640423000000', '隆德县', '101170403');
INSERT INTO `demo_city_code` VALUES ('640424', '6404', '2', '泾源', 'j', 'jing yuan', '640424000000', '泾源县', '101170404');
INSERT INTO `demo_city_code` VALUES ('640425', '6404', '2', '彭阳', 'p', 'peng yang', '640425000000', '彭阳县', '101170406');
INSERT INTO `demo_city_code` VALUES ('6405', '64', '1', '中卫', 'z', 'zhong wei', '640500000000', '中卫市', '101170501');
INSERT INTO `demo_city_code` VALUES ('640502', '6405', '2', '沙坡头', 's', 'sha po tou', '640502000000', '沙坡头区', NULL);
INSERT INTO `demo_city_code` VALUES ('640521', '6405', '2', '中宁', 'z', 'zhong ning', '640521000000', '中宁县', '101170502');
INSERT INTO `demo_city_code` VALUES ('640522', '6405', '2', '海原', 'h', 'hai yuan', '640522000000', '海原县', '101170504');
INSERT INTO `demo_city_code` VALUES ('65', '0', '0', '新疆', 'x', 'xin jiang', '650000000000', '新疆维吾尔自治区', NULL);
INSERT INTO `demo_city_code` VALUES ('6501', '65', '1', '乌鲁木齐', 'w', 'wu lu mu qi', '650100000000', '乌鲁木齐市', '101130101');
INSERT INTO `demo_city_code` VALUES ('650102', '6501', '2', '天山', 't', 'tian shan', '650102000000', '天山区', NULL);
INSERT INTO `demo_city_code` VALUES ('650103', '6501', '2', '沙依巴克区', 's', 'sha yi ba ke qu', '650103000000', '沙依巴克区', NULL);
INSERT INTO `demo_city_code` VALUES ('650104', '6501', '2', '新市', 'x', 'xin shi', '650104000000', '新市区', NULL);
INSERT INTO `demo_city_code` VALUES ('650105', '6501', '2', '水磨沟', 's', 'shui mo gou', '650105000000', '水磨沟区', NULL);
INSERT INTO `demo_city_code` VALUES ('650106', '6501', '2', '头屯河', 't', 'tou tun he', '650106000000', '头屯河区', NULL);
INSERT INTO `demo_city_code` VALUES ('650107', '6501', '2', '达坂城', 'd', 'da ban cheng', '650107000000', '达坂城区', '101130105');
INSERT INTO `demo_city_code` VALUES ('650109', '6501', '2', '米东', 'm', 'mi dong', '650109000000', '米东区', NULL);
INSERT INTO `demo_city_code` VALUES ('650121', '6501', '2', '乌鲁木齐县', 'w', 'wu lu mu qi xian', '650121000000', '乌鲁木齐县', NULL);
INSERT INTO `demo_city_code` VALUES ('6502', '65', '1', '克拉玛依', 'k', 'ke la ma yi', '650200000000', '克拉玛依市', '101130201');
INSERT INTO `demo_city_code` VALUES ('650202', '6502', '2', '独山子', 'd', 'du shan zi', '650202000000', '独山子区', NULL);
INSERT INTO `demo_city_code` VALUES ('650203', '6502', '2', '克拉玛依区', 'k', 'ke la ma yi qu', '650203000000', '克拉玛依区', NULL);
INSERT INTO `demo_city_code` VALUES ('650204', '6502', '2', '白碱滩', 'b', 'bai jian tan', '650204000000', '白碱滩区', '101130203');
INSERT INTO `demo_city_code` VALUES ('650205', '6502', '2', '乌尔禾', 'w', 'wu er he', '650205000000', '乌尔禾区', '101130202');
INSERT INTO `demo_city_code` VALUES ('6504', '65', '1', '吐鲁番', 't', 'tu lu fan', '650400000000', '吐鲁番市', '101130501');
INSERT INTO `demo_city_code` VALUES ('650402', '6504', '2', '高昌', 'g', 'gao chang', '650402000000', '高昌区', NULL);
INSERT INTO `demo_city_code` VALUES ('650421', '6504', '2', '鄯善', 's', 'shan shan', '650421000000', '鄯善县', '101130504');
INSERT INTO `demo_city_code` VALUES ('650422', '6504', '2', '托克逊', 't', 'tuo ke xun', '650422000000', '托克逊县', '101130502');
INSERT INTO `demo_city_code` VALUES ('6505', '65', '1', '哈密', 'h', 'ha mi', '650500000000', '哈密市', '101131201');
INSERT INTO `demo_city_code` VALUES ('650502', '6505', '2', '伊州', 'y', 'yi zhou', '650502000000', '伊州区', NULL);
INSERT INTO `demo_city_code` VALUES ('650521', '6505', '2', '巴里坤', 'b', 'ba li kun', '650521000000', '巴里坤哈萨克自治县', '101131203');
INSERT INTO `demo_city_code` VALUES ('650522', '6505', '2', '伊吾', 'y', 'yi wu', '650522000000', '伊吾县', '101131204');
INSERT INTO `demo_city_code` VALUES ('6523', '65', '1', '昌吉', 'c', 'chang ji', '652300000000', '昌吉回族自治州', '101130401');
INSERT INTO `demo_city_code` VALUES ('652301', '6523', '2', '昌吉市', 'c', 'chang ji shi', '652301000000', '昌吉市', NULL);
INSERT INTO `demo_city_code` VALUES ('652302', '6523', '2', '阜康', 'f', 'fu kang', '652302000000', '阜康市', '101130404');
INSERT INTO `demo_city_code` VALUES ('652323', '6523', '2', '呼图壁', 'h', 'hu tu bi', '652323000000', '呼图壁县', '101130402');
INSERT INTO `demo_city_code` VALUES ('652324', '6523', '2', '玛纳斯', 'm', 'ma na si', '652324000000', '玛纳斯县', '101130407');
INSERT INTO `demo_city_code` VALUES ('652325', '6523', '2', '奇台', 'q', 'qi tai', '652325000000', '奇台县', '101130406');
INSERT INTO `demo_city_code` VALUES ('652327', '6523', '2', '吉木萨尔', 'j', 'ji mu sa er', '652327000000', '吉木萨尔县', '101130405');
INSERT INTO `demo_city_code` VALUES ('652328', '6523', '2', '木垒', 'm', 'mu lei', '652328000000', '木垒哈萨克自治县', '101130408');
INSERT INTO `demo_city_code` VALUES ('6527', '65', '1', '博尔塔拉', 'b', 'bo er ta la', '652700000000', '博尔塔拉蒙古自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('652701', '6527', '2', '博乐', 'b', 'bo le', '652701000000', '博乐市', '101131601');
INSERT INTO `demo_city_code` VALUES ('652702', '6527', '2', '阿拉山口', 'a', 'a la shan kou', '652702000000', '阿拉山口市', '101131606');
INSERT INTO `demo_city_code` VALUES ('652722', '6527', '2', '精河', 'j', 'jing he', '652722000000', '精河县', '101131603');
INSERT INTO `demo_city_code` VALUES ('652723', '6527', '2', '温泉', 'w', 'wen quan', '652723000000', '温泉县', '101131602');
INSERT INTO `demo_city_code` VALUES ('6528', '65', '1', '巴音郭楞', 'b', 'ba yin guo leng', '652800000000', '巴音郭楞蒙古自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('652801', '6528', '2', '库尔勒', 'k', 'ku er le', '652801000000', '库尔勒市', '101130601');
INSERT INTO `demo_city_code` VALUES ('652822', '6528', '2', '轮台', 'l', 'lun tai', '652822000000', '轮台县', '101130602');
INSERT INTO `demo_city_code` VALUES ('652823', '6528', '2', '尉犁', 'y', 'yu li', '652823000000', '尉犁县', '101130603');
INSERT INTO `demo_city_code` VALUES ('652824', '6528', '2', '若羌', 'r', 'ruo qiang', '652824000000', '若羌县', '101130604');
INSERT INTO `demo_city_code` VALUES ('652825', '6528', '2', '且末', 'q', 'qie mo', '652825000000', '且末县', '101130605');
INSERT INTO `demo_city_code` VALUES ('652826', '6528', '2', '焉耆', 'y', 'yan qi', '652826000000', '焉耆回族自治县', '101130607');
INSERT INTO `demo_city_code` VALUES ('652827', '6528', '2', '和静', 'h', 'he jing', '652827000000', '和静县', '101130606');
INSERT INTO `demo_city_code` VALUES ('652828', '6528', '2', '和硕', 'h', 'he shuo', '652828000000', '和硕县', '101130608');
INSERT INTO `demo_city_code` VALUES ('652829', '6528', '2', '博湖', 'b', 'bo hu', '652829000000', '博湖县', '101130612');
INSERT INTO `demo_city_code` VALUES ('6529', '65', '1', '阿克苏', 'a', 'a ke su', '652900000000', '阿克苏地区', '101130801');
INSERT INTO `demo_city_code` VALUES ('652901', '6529', '2', '阿克苏市', 'a', 'a ke su shi', '652901000000', '阿克苏市', NULL);
INSERT INTO `demo_city_code` VALUES ('652902', '6529', '2', '库车', 'k', 'ku che', '652902000000', '库车市', '101130807');
INSERT INTO `demo_city_code` VALUES ('652922', '6529', '2', '温宿', 'w', 'wen su', '652922000000', '温宿县', '101130803');
INSERT INTO `demo_city_code` VALUES ('652924', '6529', '2', '沙雅', 's', 'sha ya', '652924000000', '沙雅县', '101130806');
INSERT INTO `demo_city_code` VALUES ('652925', '6529', '2', '新和', 'x', 'xin he', '652925000000', '新和县', '101130805');
INSERT INTO `demo_city_code` VALUES ('652926', '6529', '2', '拜城', 'b', 'bai cheng', '652926000000', '拜城县', '101130804');
INSERT INTO `demo_city_code` VALUES ('652927', '6529', '2', '乌什', 'w', 'wu shi', '652927000000', '乌什县', '101130802');
INSERT INTO `demo_city_code` VALUES ('652928', '6529', '2', '阿瓦提', 'a', 'a wa ti', '652928000000', '阿瓦提县', '101130809');
INSERT INTO `demo_city_code` VALUES ('652929', '6529', '2', '柯坪', 'k', 'ke ping', '652929000000', '柯坪县', '101130808');
INSERT INTO `demo_city_code` VALUES ('6530', '65', '1', '克孜勒苏', 'k', 'ke zi le su', '653000000000', '克孜勒苏柯尔克孜自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('653001', '6530', '2', '阿图什', 'a', 'a tu shi', '653001000000', '阿图什市', '101131501');
INSERT INTO `demo_city_code` VALUES ('653022', '6530', '2', '阿克陶', 'a', 'a ke tao', '653022000000', '阿克陶县', '101131503');
INSERT INTO `demo_city_code` VALUES ('653023', '6530', '2', '阿合奇', 'a', 'a he qi', '653023000000', '阿合奇县', '101131504');
INSERT INTO `demo_city_code` VALUES ('653024', '6530', '2', '乌恰', 'w', 'wu qia', '653024000000', '乌恰县', '101131502');
INSERT INTO `demo_city_code` VALUES ('6531', '65', '1', '喀什', 'k', 'ka shi', '653100000000', '喀什地区', '101130901');
INSERT INTO `demo_city_code` VALUES ('653101', '6531', '2', '喀什市', 'k', 'ka shi shi', '653101000000', '喀什市', NULL);
INSERT INTO `demo_city_code` VALUES ('653121', '6531', '2', '疏附', 's', 'shu fu', '653121000000', '疏附县', '101130911');
INSERT INTO `demo_city_code` VALUES ('653122', '6531', '2', '疏勒', 's', 'shu le', '653122000000', '疏勒县', '101130912');
INSERT INTO `demo_city_code` VALUES ('653123', '6531', '2', '英吉沙', 'y', 'ying ji sha', '653123000000', '英吉沙县', '101130902');
INSERT INTO `demo_city_code` VALUES ('653124', '6531', '2', '泽普', 'z', 'ze pu', '653124000000', '泽普县', '101130907');
INSERT INTO `demo_city_code` VALUES ('653125', '6531', '2', '莎车', 's', 'sha che', '653125000000', '莎车县', '101130905');
INSERT INTO `demo_city_code` VALUES ('653126', '6531', '2', '叶城', 'y', 'ye cheng', '653126000000', '叶城县', '101130906');
INSERT INTO `demo_city_code` VALUES ('653127', '6531', '2', '麦盖提', 'm', 'mai gai ti', '653127000000', '麦盖提县', '101130904');
INSERT INTO `demo_city_code` VALUES ('653128', '6531', '2', '岳普湖', 'y', 'yue pu hu', '653128000000', '岳普湖县', '101130909');
INSERT INTO `demo_city_code` VALUES ('653129', '6531', '2', '伽师', 'j', 'jia shi', '653129000000', '伽师县', '101130910');
INSERT INTO `demo_city_code` VALUES ('653130', '6531', '2', '巴楚', 'b', 'ba chu', '653130000000', '巴楚县', '101130908');
INSERT INTO `demo_city_code` VALUES ('653131', '6531', '2', '塔什库尔干', 't', 'ta shi ku er gan', '653131000000', '塔什库尔干塔吉克自治县', '101130903');
INSERT INTO `demo_city_code` VALUES ('6532', '65', '1', '和田', 'h', 'he tian', '653200000000', '和田地区', '101131301');
INSERT INTO `demo_city_code` VALUES ('653201', '6532', '2', '和田市', 'h', 'he tian shi', '653201000000', '和田市', NULL);
INSERT INTO `demo_city_code` VALUES ('653221', '6532', '2', '和田县', 'h', 'he tian xian', '653221000000', '和田县', NULL);
INSERT INTO `demo_city_code` VALUES ('653222', '6532', '2', '墨玉', 'm', 'mo yu', '653222000000', '墨玉县', '101131304');
INSERT INTO `demo_city_code` VALUES ('653223', '6532', '2', '皮山', 'p', 'pi shan', '653223000000', '皮山县', '101131302');
INSERT INTO `demo_city_code` VALUES ('653224', '6532', '2', '洛浦', 'l', 'luo pu', '653224000000', '洛浦县', '101131305');
INSERT INTO `demo_city_code` VALUES ('653225', '6532', '2', '策勒', 'c', 'ce le', '653225000000', '策勒县', '101131303');
INSERT INTO `demo_city_code` VALUES ('653226', '6532', '2', '于田', 'y', 'yu tian', '653226000000', '于田县', '101131307');
INSERT INTO `demo_city_code` VALUES ('653227', '6532', '2', '民丰', 'm', 'min feng', '653227000000', '民丰县', '101131306');
INSERT INTO `demo_city_code` VALUES ('6540', '65', '1', '伊犁', 'y', 'yi li', '654000000000', '伊犁哈萨克自治州', NULL);
INSERT INTO `demo_city_code` VALUES ('654002', '6540', '2', '伊宁市', 'y', 'yi ning shi', '654002000000', '伊宁市', NULL);
INSERT INTO `demo_city_code` VALUES ('654003', '6540', '2', '奎屯', 'k', 'kui tun', '654003000000', '奎屯市', '101131011');
INSERT INTO `demo_city_code` VALUES ('654004', '6540', '2', '霍尔果斯', 'h', 'huo er guo si', '654004000000', '霍尔果斯市', '101131010');
INSERT INTO `demo_city_code` VALUES ('654021', '6540', '2', '伊宁县', 'y', 'yi ning xian', '654021000000', '伊宁县', '101131004');
INSERT INTO `demo_city_code` VALUES ('654022', '6540', '2', '察布查尔', 'c', 'cha bu cha er', '654022000000', '察布查尔锡伯自治县', '101131002');
INSERT INTO `demo_city_code` VALUES ('654023', '6540', '2', '霍城', 'h', 'huo cheng', '654023000000', '霍城县', '101131009');
INSERT INTO `demo_city_code` VALUES ('654024', '6540', '2', '巩留', 'g', 'gong liu', '654024000000', '巩留县', '101131005');
INSERT INTO `demo_city_code` VALUES ('654025', '6540', '2', '新源', 'x', 'xin yuan', '654025000000', '新源县', '101131006');
INSERT INTO `demo_city_code` VALUES ('654026', '6540', '2', '昭苏', 'z', 'zhao su', '654026000000', '昭苏县', '101131007');
INSERT INTO `demo_city_code` VALUES ('654027', '6540', '2', '特克斯', 't', 'te ke si', '654027000000', '特克斯县', '101131008');
INSERT INTO `demo_city_code` VALUES ('654028', '6540', '2', '尼勒克', 'n', 'ni le ke', '654028000000', '尼勒克县', '101131003');
INSERT INTO `demo_city_code` VALUES ('6542', '65', '1', '塔城', 't', 'ta cheng', '654200000000', '塔城地区', '101131101');
INSERT INTO `demo_city_code` VALUES ('654201', '6542', '2', '塔城市', 't', 'ta cheng shi', '654201000000', '塔城市', NULL);
INSERT INTO `demo_city_code` VALUES ('654202', '6542', '2', '乌苏', 'w', 'wu su', '654202000000', '乌苏市', '101131106');
INSERT INTO `demo_city_code` VALUES ('654203', '6542', '2', '沙湾', 's', 'sha wan', '654203000000', '沙湾市', '101131107');
INSERT INTO `demo_city_code` VALUES ('654221', '6542', '2', '额敏', 'e', 'e min', '654221000000', '额敏县', '101131103');
INSERT INTO `demo_city_code` VALUES ('654224', '6542', '2', '托里', 't', 'tuo li', '654224000000', '托里县', '101131105');
INSERT INTO `demo_city_code` VALUES ('654225', '6542', '2', '裕民', 'y', 'yu min', '654225000000', '裕民县', '101131102');
INSERT INTO `demo_city_code` VALUES ('654226', '6542', '2', '和布克赛尔', 'h', 'he bu ke sai er', '654226000000', '和布克赛尔蒙古自治县', '101131104');
INSERT INTO `demo_city_code` VALUES ('6543', '65', '1', '阿勒泰', 'a', 'a le tai', '654300000000', '阿勒泰地区', '101131401');
INSERT INTO `demo_city_code` VALUES ('654301', '6543', '2', '阿勒泰市', 'a', 'a le tai shi', '654301000000', '阿勒泰市', NULL);
INSERT INTO `demo_city_code` VALUES ('654321', '6543', '2', '布尔津', 'b', 'bu er jin', '654321000000', '布尔津县', '101131406');
INSERT INTO `demo_city_code` VALUES ('654322', '6543', '2', '富蕴', 'f', 'fu yun', '654322000000', '富蕴县', '101131408');
INSERT INTO `demo_city_code` VALUES ('654323', '6543', '2', '福海', 'f', 'fu hai', '654323000000', '福海县', '101131407');
INSERT INTO `demo_city_code` VALUES ('654324', '6543', '2', '哈巴河', 'h', 'ha ba he', '654324000000', '哈巴河县', '101131402');
INSERT INTO `demo_city_code` VALUES ('654325', '6543', '2', '青河', 'q', 'qing he', '654325000000', '青河县', '101131409');
INSERT INTO `demo_city_code` VALUES ('654326', '6543', '2', '吉木乃', 'j', 'ji mu nai', '654326000000', '吉木乃县', '101131405');
INSERT INTO `demo_city_code` VALUES ('659001', '65', '1', '石河子', 's', 'shi he zi', '659001000000', '石河子市', '101130301');
INSERT INTO `demo_city_code` VALUES ('659001000', '659001', '2', '石河子', 's', 'shi he zi', '659001000000', '石河子市', '101130301');
INSERT INTO `demo_city_code` VALUES ('659002', '65', '1', '阿拉尔', 'a', 'a la er', '659002000000', '阿拉尔市', '101130701');
INSERT INTO `demo_city_code` VALUES ('659002000', '659002', '2', '阿拉尔', 'a', 'a la er', '659002000000', '阿拉尔市', '101130701');
INSERT INTO `demo_city_code` VALUES ('659003', '65', '1', '图木舒克', 't', 'tu mu shu ke', '659003000000', '图木舒克市', NULL);
INSERT INTO `demo_city_code` VALUES ('659003000', '659003', '2', '图木舒克', 't', 'tu mu shu ke', '659003000000', '图木舒克市', NULL);
INSERT INTO `demo_city_code` VALUES ('659004', '65', '1', '五家渠', 'w', 'wu jia qu', '659004000000', '五家渠市', NULL);
INSERT INTO `demo_city_code` VALUES ('659004000', '659004', '2', '五家渠', 'w', 'wu jia qu', '659004000000', '五家渠市', NULL);
INSERT INTO `demo_city_code` VALUES ('659005', '65', '1', '北屯', 'b', 'bei tun', '659005000000', '北屯市', NULL);
INSERT INTO `demo_city_code` VALUES ('659005000', '659005', '2', '北屯', 'b', 'bei tun', '659005000000', '北屯市', NULL);
INSERT INTO `demo_city_code` VALUES ('659006', '65', '1', '铁门关', 't', 'tie men guan', '659006000000', '铁门关市', NULL);
INSERT INTO `demo_city_code` VALUES ('659006000', '659006', '2', '铁门关', 't', 'tie men guan', '659006000000', '铁门关市', NULL);
INSERT INTO `demo_city_code` VALUES ('659007', '65', '1', '双河', 's', 'shuang he', '659007000000', '双河市', NULL);
INSERT INTO `demo_city_code` VALUES ('659007000', '659007', '2', '双河', 's', 'shuang he', '659007000000', '双河市', NULL);
INSERT INTO `demo_city_code` VALUES ('659008', '65', '1', '可克达拉', 'k', 'ke ke da la', '659008000000', '可克达拉市', NULL);
INSERT INTO `demo_city_code` VALUES ('659008000', '659008', '2', '可克达拉', 'k', 'ke ke da la', '659008000000', '可克达拉市', NULL);
INSERT INTO `demo_city_code` VALUES ('659009', '65', '1', '昆玉', 'k', 'kun yu', '659009000000', '昆玉市', NULL);
INSERT INTO `demo_city_code` VALUES ('659009000', '659009', '2', '昆玉', 'k', 'kun yu', '659009000000', '昆玉市', NULL);
INSERT INTO `demo_city_code` VALUES ('659010', '65', '1', '胡杨河', 'h', 'hu yang he', '659010000000', '胡杨河市', NULL);
INSERT INTO `demo_city_code` VALUES ('659010000', '659010', '2', '胡杨河', 'h', 'hu yang he', '659010000000', '胡杨河市', NULL);
INSERT INTO `demo_city_code` VALUES ('659011', '65', '1', '新星', 'x', 'xin xing', '659011000000', '新星市', NULL);
INSERT INTO `demo_city_code` VALUES ('659011000', '659011', '2', '新星', 'x', 'xin xing', '659011000000', '新星市', NULL);
INSERT INTO `demo_city_code` VALUES ('71', '0', '0', '台湾', '~3', 'tai wan', '710000000000', '台湾省', NULL);
INSERT INTO `demo_city_code` VALUES ('7101', '71', '1', '台北', 't', 'tai bei', '710100000000', '台北市', '101340101');
INSERT INTO `demo_city_code` VALUES ('710101', '7101', '2', '中正', 'z', 'zhong zheng', '710101000000', '中正区', NULL);
INSERT INTO `demo_city_code` VALUES ('710102', '7101', '2', '大同', 'd', 'da tong', '710102000000', '大同区', '101100201');
INSERT INTO `demo_city_code` VALUES ('710103', '7101', '2', '中山', 'z', 'zhong shan', '710103000000', '中山区', '101281701');
INSERT INTO `demo_city_code` VALUES ('710104', '7101', '2', '松山', 's', 'song shan', '710104000000', '松山区', NULL);
INSERT INTO `demo_city_code` VALUES ('710105', '7101', '2', '大安', 'd', 'da an', '710105000000', '大安区', '101060603');
INSERT INTO `demo_city_code` VALUES ('710106', '7101', '2', '万华', 'w', 'wan hua', '710106000000', '万华区', NULL);
INSERT INTO `demo_city_code` VALUES ('710107', '7101', '2', '信义', 'x', 'xin yi', '710107000000', '信义区', NULL);
INSERT INTO `demo_city_code` VALUES ('710108', '7101', '2', '士林', 's', 'shi lin', '710108000000', '士林区', NULL);
INSERT INTO `demo_city_code` VALUES ('710109', '7101', '2', '北投', 'b', 'bei tou', '710109000000', '北投区', NULL);
INSERT INTO `demo_city_code` VALUES ('710110', '7101', '2', '内湖', 'n', 'nei hu', '710110000000', '内湖区', NULL);
INSERT INTO `demo_city_code` VALUES ('710111', '7101', '2', '南港', 'n', 'nan gang', '710111000000', '南港区', NULL);
INSERT INTO `demo_city_code` VALUES ('710112', '7101', '2', '文山', 'w', 'wen shan', '710112000000', '文山区', '101290601');
INSERT INTO `demo_city_code` VALUES ('7102', '71', '1', '高雄', 'g', 'gao xiong', '710200000000', '高雄市', '101340201');
INSERT INTO `demo_city_code` VALUES ('710201', '7102', '2', '新兴', 'x', 'xin xing', '710201000000', '新兴区', '101281403');
INSERT INTO `demo_city_code` VALUES ('710202', '7102', '2', '前金', 'q', 'qian jin', '710202000000', '前金区', NULL);
INSERT INTO `demo_city_code` VALUES ('710203', '7102', '2', '苓雅', 'l', 'ling ya', '710203000000', '苓雅区', NULL);
INSERT INTO `demo_city_code` VALUES ('710204', '7102', '2', '盐埕', 'y', 'yan cheng', '710204000000', '盐埕区', NULL);
INSERT INTO `demo_city_code` VALUES ('710205', '7102', '2', '鼓山', 'g', 'gu shan', '710205000000', '鼓山区', NULL);
INSERT INTO `demo_city_code` VALUES ('710206', '7102', '2', '旗津', 'q', 'qi jin', '710206000000', '旗津区', NULL);
INSERT INTO `demo_city_code` VALUES ('710207', '7102', '2', '前镇', 'q', 'qian zhen', '710207000000', '前镇区', NULL);
INSERT INTO `demo_city_code` VALUES ('710208', '7102', '2', '三民', 's', 'san min', '710208000000', '三民区', NULL);
INSERT INTO `demo_city_code` VALUES ('710209', '7102', '2', '左营', 'z', 'zuo ying', '710209000000', '左营区', NULL);
INSERT INTO `demo_city_code` VALUES ('710210', '7102', '2', '楠梓', 'n', 'nan zi', '710210000000', '楠梓区', NULL);
INSERT INTO `demo_city_code` VALUES ('710211', '7102', '2', '小港', 'x', 'xiao gang', '710211000000', '小港区', NULL);
INSERT INTO `demo_city_code` VALUES ('710242', '7102', '2', '仁武', 'r', 'ren wu', '710242000000', '仁武区', NULL);
INSERT INTO `demo_city_code` VALUES ('710243', '7102', '2', '大社', 'd', 'da she', '710243000000', '大社区', NULL);
INSERT INTO `demo_city_code` VALUES ('710244', '7102', '2', '冈山', 'g', 'gang shan', '710244000000', '冈山区', NULL);
INSERT INTO `demo_city_code` VALUES ('710245', '7102', '2', '路竹', 'l', 'lu zhu', '710245000000', '路竹区', NULL);
INSERT INTO `demo_city_code` VALUES ('710246', '7102', '2', '阿莲', 'a', 'a lian', '710246000000', '阿莲区', NULL);
INSERT INTO `demo_city_code` VALUES ('710247', '7102', '2', '田寮', 't', 'tian liao', '710247000000', '田寮区', NULL);
INSERT INTO `demo_city_code` VALUES ('710248', '7102', '2', '燕巢', 'y', 'yan chao', '710248000000', '燕巢区', NULL);
INSERT INTO `demo_city_code` VALUES ('710249', '7102', '2', '桥头', 'q', 'qiao tou', '710249000000', '桥头区', NULL);
INSERT INTO `demo_city_code` VALUES ('710250', '7102', '2', '梓官', 'z', 'zi guan', '710250000000', '梓官区', NULL);
INSERT INTO `demo_city_code` VALUES ('710251', '7102', '2', '弥陀', 'm', 'mi tuo', '710251000000', '弥陀区', NULL);
INSERT INTO `demo_city_code` VALUES ('710252', '7102', '2', '永安', 'y', 'yong an', '710252000000', '永安区', '101230810');
INSERT INTO `demo_city_code` VALUES ('710253', '7102', '2', '湖内', 'h', 'hu nei', '710253000000', '湖内区', NULL);
INSERT INTO `demo_city_code` VALUES ('710254', '7102', '2', '凤山', 'f', 'feng shan', '710254000000', '凤山区', '101301208');
INSERT INTO `demo_city_code` VALUES ('710255', '7102', '2', '大寮', 'd', 'da liao', '710255000000', '大寮区', NULL);
INSERT INTO `demo_city_code` VALUES ('710256', '7102', '2', '林园', 'l', 'lin yuan', '710256000000', '林园区', NULL);
INSERT INTO `demo_city_code` VALUES ('710257', '7102', '2', '鸟松', 'n', 'niao song', '710257000000', '鸟松区', NULL);
INSERT INTO `demo_city_code` VALUES ('710258', '7102', '2', '大树', 'd', 'da shu', '710258000000', '大树区', NULL);
INSERT INTO `demo_city_code` VALUES ('710259', '7102', '2', '旗山', 'q', 'qi shan', '710259000000', '旗山区', NULL);
INSERT INTO `demo_city_code` VALUES ('710260', '7102', '2', '美浓', 'm', 'mei nong', '710260000000', '美浓区', NULL);
INSERT INTO `demo_city_code` VALUES ('710261', '7102', '2', '六龟', 'l', 'liu gui', '710261000000', '六龟区', NULL);
INSERT INTO `demo_city_code` VALUES ('710262', '7102', '2', '内门', 'n', 'nei men', '710262000000', '内门区', NULL);
INSERT INTO `demo_city_code` VALUES ('710263', '7102', '2', '杉林', 's', 'shan lin', '710263000000', '杉林区', NULL);
INSERT INTO `demo_city_code` VALUES ('710264', '7102', '2', '甲仙', 'j', 'jia xian', '710264000000', '甲仙区', NULL);
INSERT INTO `demo_city_code` VALUES ('710265', '7102', '2', '桃源', 't', 'tao yuan', '710265000000', '桃源区', '101250603');
INSERT INTO `demo_city_code` VALUES ('710266', '7102', '2', '那玛夏', 'n', 'na ma xia', '710266000000', '那玛夏区', NULL);
INSERT INTO `demo_city_code` VALUES ('710267', '7102', '2', '茂林', 'm', 'mao lin', '710267000000', '茂林区', NULL);
INSERT INTO `demo_city_code` VALUES ('710268', '7102', '2', '茄萣', 'q', 'qie ding', '710268000000', '茄萣区', NULL);
INSERT INTO `demo_city_code` VALUES ('7103', '71', '1', '台南', 't', 'tai nan', '710300000000', '台南市', '101340203');
INSERT INTO `demo_city_code` VALUES ('710301', '7103', '2', '中西', 'z', 'zhong xi', '710301000000', '中西区', NULL);
INSERT INTO `demo_city_code` VALUES ('710302', '7103', '2', '东区', 'd', 'dong qu', '710302000000', '东区', NULL);
INSERT INTO `demo_city_code` VALUES ('710303', '7103', '2', '南区', 'n', 'nan qu', '710303000000', '南区', NULL);
INSERT INTO `demo_city_code` VALUES ('710304', '7103', '2', '北区', 'b', 'bei qu', '710304000000', '北区', NULL);
INSERT INTO `demo_city_code` VALUES ('710305', '7103', '2', '安平', 'a', 'an ping', '710305000000', '安平区', '101090806');
INSERT INTO `demo_city_code` VALUES ('710306', '7103', '2', '安南', 'a', 'an nan', '710306000000', '安南区', NULL);
INSERT INTO `demo_city_code` VALUES ('710339', '7103', '2', '永康', 'y', 'yong kang', '710339000000', '永康区', '101210907');
INSERT INTO `demo_city_code` VALUES ('710340', '7103', '2', '归仁', 'g', 'gui ren', '710340000000', '归仁区', NULL);
INSERT INTO `demo_city_code` VALUES ('710341', '7103', '2', '新化', 'x', 'xin hua', '710341000000', '新化区', '101250805');
INSERT INTO `demo_city_code` VALUES ('710342', '7103', '2', '左镇', 'z', 'zuo zhen', '710342000000', '左镇区', NULL);
INSERT INTO `demo_city_code` VALUES ('710343', '7103', '2', '玉井', 'y', 'yu jing', '710343000000', '玉井区', NULL);
INSERT INTO `demo_city_code` VALUES ('710344', '7103', '2', '楠西', 'n', 'nan xi', '710344000000', '楠西区', NULL);
INSERT INTO `demo_city_code` VALUES ('710345', '7103', '2', '南化', 'n', 'nan hua', '710345000000', '南化区', NULL);
INSERT INTO `demo_city_code` VALUES ('710346', '7103', '2', '仁德', 'r', 'ren de', '710346000000', '仁德区', NULL);
INSERT INTO `demo_city_code` VALUES ('710347', '7103', '2', '关庙', 'g', 'guan miao', '710347000000', '关庙区', NULL);
INSERT INTO `demo_city_code` VALUES ('710348', '7103', '2', '龙崎', 'l', 'long qi', '710348000000', '龙崎区', NULL);
INSERT INTO `demo_city_code` VALUES ('710349', '7103', '2', '官田', 'g', 'guan tian', '710349000000', '官田区', NULL);
INSERT INTO `demo_city_code` VALUES ('710350', '7103', '2', '麻豆', 'm', 'ma dou', '710350000000', '麻豆区', NULL);
INSERT INTO `demo_city_code` VALUES ('710351', '7103', '2', '佳里', 'j', 'jia li', '710351000000', '佳里区', NULL);
INSERT INTO `demo_city_code` VALUES ('710352', '7103', '2', '西港', 'x', 'xi gang', '710352000000', '西港区', NULL);
INSERT INTO `demo_city_code` VALUES ('710353', '7103', '2', '七股', 'q', 'qi gu', '710353000000', '七股区', NULL);
INSERT INTO `demo_city_code` VALUES ('710354', '7103', '2', '将军', 'j', 'jiang jun', '710354000000', '将军区', NULL);
INSERT INTO `demo_city_code` VALUES ('710355', '7103', '2', '学甲', 'x', 'xue jia', '710355000000', '学甲区', NULL);
INSERT INTO `demo_city_code` VALUES ('710356', '7103', '2', '北门', 'b', 'bei men', '710356000000', '北门区', NULL);
INSERT INTO `demo_city_code` VALUES ('710357', '7103', '2', '新营', 'x', 'xin ying', '710357000000', '新营区', NULL);
INSERT INTO `demo_city_code` VALUES ('710358', '7103', '2', '后壁', 'h', 'hou bi', '710358000000', '后壁区', NULL);
INSERT INTO `demo_city_code` VALUES ('710359', '7103', '2', '白河', 'b', 'bai he', '710359000000', '白河区', '101110708');
INSERT INTO `demo_city_code` VALUES ('710360', '7103', '2', '东山', 'd', 'dong shan', '710360000000', '东山区', '101230608');
INSERT INTO `demo_city_code` VALUES ('710361', '7103', '2', '六甲', 'l', 'liu jia', '710361000000', '六甲区', NULL);
INSERT INTO `demo_city_code` VALUES ('710362', '7103', '2', '下营', 'x', 'xia ying', '710362000000', '下营区', NULL);
INSERT INTO `demo_city_code` VALUES ('710363', '7103', '2', '柳营', 'l', 'liu ying', '710363000000', '柳营区', NULL);
INSERT INTO `demo_city_code` VALUES ('710364', '7103', '2', '盐水', 'y', 'yan shui', '710364000000', '盐水区', NULL);
INSERT INTO `demo_city_code` VALUES ('710365', '7103', '2', '善化', 's', 'shan hua', '710365000000', '善化区', NULL);
INSERT INTO `demo_city_code` VALUES ('710366', '7103', '2', '大内', 'd', 'da nei', '710366000000', '大内区', NULL);
INSERT INTO `demo_city_code` VALUES ('710367', '7103', '2', '山上', 's', 'shan shang', '710367000000', '山上区', NULL);
INSERT INTO `demo_city_code` VALUES ('710368', '7103', '2', '新市', 'x', 'xin shi', '710368000000', '新市区', NULL);
INSERT INTO `demo_city_code` VALUES ('710369', '7103', '2', '安定', 'a', 'an ding', '710369000000', '安定区', '101160208');
INSERT INTO `demo_city_code` VALUES ('7104', '71', '1', '台中', 't', 'tai zhong', '710400000000', '台中市', '101340401');
INSERT INTO `demo_city_code` VALUES ('710401', '7104', '2', '中区', 'z', 'zhong qu', '710401000000', '中区', NULL);
INSERT INTO `demo_city_code` VALUES ('710402', '7104', '2', '东区', 'd', 'dong qu', '710402000000', '东区', NULL);
INSERT INTO `demo_city_code` VALUES ('710403', '7104', '2', '南区', 'n', 'nan qu', '710403000000', '南区', NULL);
INSERT INTO `demo_city_code` VALUES ('710404', '7104', '2', '西区', 'x', 'xi qu', '710404000000', '西区', NULL);
INSERT INTO `demo_city_code` VALUES ('710405', '7104', '2', '北区', 'b', 'bei qu', '710405000000', '北区', NULL);
INSERT INTO `demo_city_code` VALUES ('710406', '7104', '2', '北屯', 'b', 'bei tun', '710406000000', '北屯区', NULL);
INSERT INTO `demo_city_code` VALUES ('710407', '7104', '2', '西屯', 'x', 'xi tun', '710407000000', '西屯区', NULL);
INSERT INTO `demo_city_code` VALUES ('710408', '7104', '2', '南屯', 'n', 'nan tun', '710408000000', '南屯区', NULL);
INSERT INTO `demo_city_code` VALUES ('710431', '7104', '2', '太平', 't', 'tai ping', '710431000000', '太平区', NULL);
INSERT INTO `demo_city_code` VALUES ('710432', '7104', '2', '大里', 'd', 'da li', '710432000000', '大里区', NULL);
INSERT INTO `demo_city_code` VALUES ('710433', '7104', '2', '雾峰', 'w', 'wu feng', '710433000000', '雾峰区', NULL);
INSERT INTO `demo_city_code` VALUES ('710434', '7104', '2', '乌日', 'w', 'wu ri', '710434000000', '乌日区', NULL);
INSERT INTO `demo_city_code` VALUES ('710435', '7104', '2', '丰原', 'f', 'feng yuan', '710435000000', '丰原区', NULL);
INSERT INTO `demo_city_code` VALUES ('710436', '7104', '2', '后里', 'h', 'hou li', '710436000000', '后里区', NULL);
INSERT INTO `demo_city_code` VALUES ('710437', '7104', '2', '石冈', 's', 'shi gang', '710437000000', '石冈区', NULL);
INSERT INTO `demo_city_code` VALUES ('710438', '7104', '2', '东势', 'd', 'dong shi', '710438000000', '东势区', NULL);
INSERT INTO `demo_city_code` VALUES ('710439', '7104', '2', '和平', 'h', 'he ping', '710439000000', '和平区', '101281204');
INSERT INTO `demo_city_code` VALUES ('710440', '7104', '2', '新社', 'x', 'xin she', '710440000000', '新社区', NULL);
INSERT INTO `demo_city_code` VALUES ('710441', '7104', '2', '潭子', 't', 'tan zi', '710441000000', '潭子区', NULL);
INSERT INTO `demo_city_code` VALUES ('710442', '7104', '2', '大雅', 'd', 'da ya', '710442000000', '大雅区', NULL);
INSERT INTO `demo_city_code` VALUES ('710443', '7104', '2', '神冈', 's', 'shen gang', '710443000000', '神冈区', NULL);
INSERT INTO `demo_city_code` VALUES ('710444', '7104', '2', '大肚', 'd', 'da du', '710444000000', '大肚区', NULL);
INSERT INTO `demo_city_code` VALUES ('710445', '7104', '2', '沙鹿', 's', 'sha lu', '710445000000', '沙鹿区', NULL);
INSERT INTO `demo_city_code` VALUES ('710446', '7104', '2', '龙井', 'l', 'long jing', '710446000000', '龙井区', '101060307');
INSERT INTO `demo_city_code` VALUES ('710447', '7104', '2', '梧栖', 'w', 'wu qi', '710447000000', '梧栖区', NULL);
INSERT INTO `demo_city_code` VALUES ('710448', '7104', '2', '清水', 'q', 'qing shui', '710448000000', '清水区', '101160903');
INSERT INTO `demo_city_code` VALUES ('710449', '7104', '2', '大甲', 'd', 'da jia', '710449000000', '大甲区', NULL);
INSERT INTO `demo_city_code` VALUES ('710450', '7104', '2', '外埔', 'w', 'wai bu', '710450000000', '外埔区', NULL);
INSERT INTO `demo_city_code` VALUES ('710451', '7104', '2', '大安', 'd', 'da an', '710451000000', '大安区', '101060603');
INSERT INTO `demo_city_code` VALUES ('7106', '71', '1', '南投', 'n', 'nan tou', '710600000000', '南投县', '101340404');
INSERT INTO `demo_city_code` VALUES ('710614', '7106', '2', '南投市', 'n', 'nan tou shi', '710614000000', '南投市', NULL);
INSERT INTO `demo_city_code` VALUES ('710615', '7106', '2', '中寮', 'z', 'zhong liao', '710615000000', '中寮乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710616', '7106', '2', '草屯', 'c', 'cao tun', '710616000000', '草屯镇', NULL);
INSERT INTO `demo_city_code` VALUES ('710617', '7106', '2', '国姓', 'g', 'guo xing', '710617000000', '国姓乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710618', '7106', '2', '埔里', 'b', 'bu li', '710618000000', '埔里镇', NULL);
INSERT INTO `demo_city_code` VALUES ('710619', '7106', '2', '仁爱', 'r', 'ren ai', '710619000000', '仁爱乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710620', '7106', '2', '名间', 'm', 'ming jian', '710620000000', '名间乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710621', '7106', '2', '集集', 'j', 'ji ji', '710621000000', '集集镇', NULL);
INSERT INTO `demo_city_code` VALUES ('710622', '7106', '2', '水里', 's', 'shui li', '710622000000', '水里乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710623', '7106', '2', '鱼池', 'y', 'yu chi', '710623000000', '鱼池乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710624', '7106', '2', '信义', 'x', 'xin yi', '710624000000', '信义乡', NULL);
INSERT INTO `demo_city_code` VALUES ('710625', '7106', '2', '竹山', 'z', 'zhu shan', '710625000000', '竹山镇', '101201105');
INSERT INTO `demo_city_code` VALUES ('710626', '7106', '2', '鹿谷', 'l', 'lu gu', '710626000000', '鹿谷乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7107', '71', '1', '基隆', 'j', 'ji long', '710700000000', '基隆市', NULL);
INSERT INTO `demo_city_code` VALUES ('710701', '7107', '2', '仁爱', 'r', 'ren ai', '710701000000', '仁爱区', NULL);
INSERT INTO `demo_city_code` VALUES ('710702', '7107', '2', '信义', 'x', 'xin yi', '710702000000', '信义区', NULL);
INSERT INTO `demo_city_code` VALUES ('710703', '7107', '2', '中正', 'z', 'zhong zheng', '710703000000', '中正区', NULL);
INSERT INTO `demo_city_code` VALUES ('710704', '7107', '2', '中山', 'z', 'zhong shan', '710704000000', '中山区', '101281701');
INSERT INTO `demo_city_code` VALUES ('710705', '7107', '2', '安乐', 'a', 'an le', '710705000000', '安乐区', NULL);
INSERT INTO `demo_city_code` VALUES ('710706', '7107', '2', '暖暖', 'n', 'nuan nuan', '710706000000', '暖暖区', NULL);
INSERT INTO `demo_city_code` VALUES ('710707', '7107', '2', '七堵', 'q', 'qi du', '710707000000', '七堵区', NULL);
INSERT INTO `demo_city_code` VALUES ('7108', '71', '1', '新竹市', 'x', 'xin zhu shi', '710800000000', '新竹市', NULL);
INSERT INTO `demo_city_code` VALUES ('710801', '7108', '2', '东区', 'd', 'dong qu', '710801000000', '东区', NULL);
INSERT INTO `demo_city_code` VALUES ('710802', '7108', '2', '北区', 'b', 'bei qu', '710802000000', '北区', NULL);
INSERT INTO `demo_city_code` VALUES ('710803', '7108', '2', '香山', 'x', 'xiang shan', '710803000000', '香山区', NULL);
INSERT INTO `demo_city_code` VALUES ('7109', '71', '1', '嘉义市', 'j', 'jia yi shi', '710900000000', '嘉义市', NULL);
INSERT INTO `demo_city_code` VALUES ('710901', '7109', '2', '东区', 'd', 'dong qu', '710901000000', '东区', NULL);
INSERT INTO `demo_city_code` VALUES ('710902', '7109', '2', '西区', 'x', 'xi qu', '710902000000', '西区', NULL);
INSERT INTO `demo_city_code` VALUES ('7111', '71', '1', '新北', 'x', 'xin bei', '711100000000', '新北市', NULL);
INSERT INTO `demo_city_code` VALUES ('711130', '7111', '2', '万里', 'w', 'wan li', '711130000000', '万里区', NULL);
INSERT INTO `demo_city_code` VALUES ('711131', '7111', '2', '金山', 'j', 'jin shan', '711131000000', '金山区', '101020700');
INSERT INTO `demo_city_code` VALUES ('711132', '7111', '2', '板桥', 'b', 'ban qiao', '711132000000', '板桥区', NULL);
INSERT INTO `demo_city_code` VALUES ('711133', '7111', '2', '汐止', 'x', 'xi zhi', '711133000000', '汐止区', NULL);
INSERT INTO `demo_city_code` VALUES ('711134', '7111', '2', '深坑', 's', 'shen keng', '711134000000', '深坑区', NULL);
INSERT INTO `demo_city_code` VALUES ('711135', '7111', '2', '石碇', 's', 'shi ding', '711135000000', '石碇区', NULL);
INSERT INTO `demo_city_code` VALUES ('711136', '7111', '2', '瑞芳', 'r', 'rui fang', '711136000000', '瑞芳区', NULL);
INSERT INTO `demo_city_code` VALUES ('711137', '7111', '2', '平溪', 'p', 'ping xi', '711137000000', '平溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('711138', '7111', '2', '双溪', 's', 'shuang xi', '711138000000', '双溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('711139', '7111', '2', '贡寮', 'g', 'gong liao', '711139000000', '贡寮区', NULL);
INSERT INTO `demo_city_code` VALUES ('711140', '7111', '2', '新店', 'x', 'xin dian', '711140000000', '新店区', NULL);
INSERT INTO `demo_city_code` VALUES ('711141', '7111', '2', '坪林', 'p', 'ping lin', '711141000000', '坪林区', NULL);
INSERT INTO `demo_city_code` VALUES ('711142', '7111', '2', '乌来', 'w', 'wu lai', '711142000000', '乌来区', NULL);
INSERT INTO `demo_city_code` VALUES ('711143', '7111', '2', '永和', 'y', 'yong he', '711143000000', '永和区', '101100703');
INSERT INTO `demo_city_code` VALUES ('711144', '7111', '2', '中和', 'z', 'zhong he', '711144000000', '中和区', NULL);
INSERT INTO `demo_city_code` VALUES ('711145', '7111', '2', '土城', 't', 'tu cheng', '711145000000', '土城区', NULL);
INSERT INTO `demo_city_code` VALUES ('711146', '7111', '2', '三峡', 's', 'san xia', '711146000000', '三峡区', '101200911');
INSERT INTO `demo_city_code` VALUES ('711147', '7111', '2', '树林', 's', 'shu lin', '711147000000', '树林区', NULL);
INSERT INTO `demo_city_code` VALUES ('711148', '7111', '2', '莺歌', 'y', 'ying ge', '711148000000', '莺歌区', NULL);
INSERT INTO `demo_city_code` VALUES ('711149', '7111', '2', '三重', 's', 'san chong', '711149000000', '三重区', NULL);
INSERT INTO `demo_city_code` VALUES ('711150', '7111', '2', '新庄', 'x', 'xin zhuang', '711150000000', '新庄区', NULL);
INSERT INTO `demo_city_code` VALUES ('711151', '7111', '2', '泰山', 't', 'tai shan', '711151000000', '泰山区', NULL);
INSERT INTO `demo_city_code` VALUES ('711152', '7111', '2', '林口', 'l', 'lin kou', '711152000000', '林口区', '101050304');
INSERT INTO `demo_city_code` VALUES ('711153', '7111', '2', '芦洲', 'l', 'lu zhou', '711153000000', '芦洲区', NULL);
INSERT INTO `demo_city_code` VALUES ('711154', '7111', '2', '五股', 'w', 'wu gu', '711154000000', '五股区', NULL);
INSERT INTO `demo_city_code` VALUES ('711155', '7111', '2', '八里', 'b', 'ba li', '711155000000', '八里区', NULL);
INSERT INTO `demo_city_code` VALUES ('711156', '7111', '2', '淡水', 'd', 'dan shui', '711156000000', '淡水区', NULL);
INSERT INTO `demo_city_code` VALUES ('711157', '7111', '2', '三芝', 's', 'san zhi', '711157000000', '三芝区', NULL);
INSERT INTO `demo_city_code` VALUES ('711158', '7111', '2', '石门', 's', 'shi men', '711158000000', '石门区', '101250607');
INSERT INTO `demo_city_code` VALUES ('7112', '71', '1', '宜兰', 'y', 'yi lan', '711200000000', '宜兰县', '101340104');
INSERT INTO `demo_city_code` VALUES ('711214', '7112', '2', '宜兰市', 'y', 'yi lan shi', '711214000000', '宜兰市', NULL);
INSERT INTO `demo_city_code` VALUES ('711215', '7112', '2', '头城', 't', 'tou cheng', '711215000000', '头城镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711216', '7112', '2', '礁溪', 'j', 'jiao xi', '711216000000', '礁溪乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711217', '7112', '2', '壮围', 'z', 'zhuang wei', '711217000000', '壮围乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711218', '7112', '2', '员山', 'y', 'yuan shan', '711218000000', '员山乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711219', '7112', '2', '罗东', 'l', 'luo dong', '711219000000', '罗东镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711220', '7112', '2', '三星', 's', 'san xing', '711220000000', '三星乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711221', '7112', '2', '大同', 'd', 'da tong', '711221000000', '大同乡', '101100201');
INSERT INTO `demo_city_code` VALUES ('711222', '7112', '2', '五结', 'w', 'wu jie', '711222000000', '五结乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711223', '7112', '2', '冬山', 'd', 'dong shan', '711223000000', '冬山乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711224', '7112', '2', '苏澳', 's', 'su ao', '711224000000', '苏澳镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711225', '7112', '2', '南澳', 'n', 'nan ao', '711225000000', '南澳乡', '101280504');
INSERT INTO `demo_city_code` VALUES ('7113', '71', '1', '新竹县', 'x', 'xin zhu xian', '711300000000', '新竹县', NULL);
INSERT INTO `demo_city_code` VALUES ('711314', '7113', '2', '竹北', 'z', 'zhu bei', '711314000000', '竹北市', NULL);
INSERT INTO `demo_city_code` VALUES ('711315', '7113', '2', '湖口', 'h', 'hu kou', '711315000000', '湖口乡', '101240207');
INSERT INTO `demo_city_code` VALUES ('711316', '7113', '2', '新丰', 'x', 'xin feng', '711316000000', '新丰乡', '101280208');
INSERT INTO `demo_city_code` VALUES ('711317', '7113', '2', '新埔', 'x', 'xin bu', '711317000000', '新埔镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711318', '7113', '2', '关西', 'g', 'guan xi', '711318000000', '关西镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711319', '7113', '2', '芎林', 'x', 'xiong lin', '711319000000', '芎林乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711320', '7113', '2', '宝山', 'b', 'bao shan', '711320000000', '宝山乡', '101020300');
INSERT INTO `demo_city_code` VALUES ('711321', '7113', '2', '竹东', 'z', 'zhu dong', '711321000000', '竹东镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711322', '7113', '2', '五峰', 'w', 'wu feng', '711322000000', '五峰乡', '101200906');
INSERT INTO `demo_city_code` VALUES ('711323', '7113', '2', '横山', 'h', 'heng shan', '711323000000', '横山乡', '101110407');
INSERT INTO `demo_city_code` VALUES ('711324', '7113', '2', '尖石', 'j', 'jian shi', '711324000000', '尖石乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711325', '7113', '2', '北埔', 'b', 'bei bu', '711325000000', '北埔乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711326', '7113', '2', '峨眉', 'e', 'e mei', '711326000000', '峨眉乡', '101271408');
INSERT INTO `demo_city_code` VALUES ('7114', '71', '1', '桃园', 't', 'tao yuan', '711400000000', '桃园市', '101340102');
INSERT INTO `demo_city_code` VALUES ('711414', '7114', '2', '中坜', 'z', 'zhong li', '711414000000', '中坜区', NULL);
INSERT INTO `demo_city_code` VALUES ('711415', '7114', '2', '平镇', 'p', 'ping zhen', '711415000000', '平镇区', NULL);
INSERT INTO `demo_city_code` VALUES ('711416', '7114', '2', '龙潭', 'l', 'long tan', '711416000000', '龙潭区', NULL);
INSERT INTO `demo_city_code` VALUES ('711417', '7114', '2', '杨梅', 'y', 'yang mei', '711417000000', '杨梅区', NULL);
INSERT INTO `demo_city_code` VALUES ('711418', '7114', '2', '新屋', 'x', 'xin wu', '711418000000', '新屋区', NULL);
INSERT INTO `demo_city_code` VALUES ('711419', '7114', '2', '观音', 'g', 'guan yin', '711419000000', '观音区', NULL);
INSERT INTO `demo_city_code` VALUES ('711420', '7114', '2', '桃园区', 't', 'tao yuan qu', '711420000000', '桃园区', NULL);
INSERT INTO `demo_city_code` VALUES ('711421', '7114', '2', '龟山', 'g', 'gui shan', '711421000000', '龟山区', NULL);
INSERT INTO `demo_city_code` VALUES ('711422', '7114', '2', '八德', 'b', 'ba de', '711422000000', '八德区', NULL);
INSERT INTO `demo_city_code` VALUES ('711423', '7114', '2', '大溪', 'd', 'da xi', '711423000000', '大溪区', NULL);
INSERT INTO `demo_city_code` VALUES ('711424', '7114', '2', '复兴', 'f', 'fu xing', '711424000000', '复兴区', NULL);
INSERT INTO `demo_city_code` VALUES ('711425', '7114', '2', '大园', 'd', 'da yuan', '711425000000', '大园区', NULL);
INSERT INTO `demo_city_code` VALUES ('711426', '7114', '2', '芦竹', 'l', 'lu zhu', '711426000000', '芦竹区', NULL);
INSERT INTO `demo_city_code` VALUES ('7115', '71', '1', '苗栗', 'm', 'miao li', '711500000000', '苗栗县', '101340402');
INSERT INTO `demo_city_code` VALUES ('711519', '7115', '2', '竹南', 'z', 'zhu nan', '711519000000', '竹南镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711520', '7115', '2', '头份', 't', 'tou fen', '711520000000', '头份市', NULL);
INSERT INTO `demo_city_code` VALUES ('711521', '7115', '2', '三湾', 's', 'san wan', '711521000000', '三湾乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711522', '7115', '2', '南庄', 'n', 'nan zhuang', '711522000000', '南庄乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711523', '7115', '2', '狮潭', 's', 'shi tan', '711523000000', '狮潭乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711524', '7115', '2', '后龙', 'h', 'hou long', '711524000000', '后龙镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711525', '7115', '2', '通霄', 't', 'tong xiao', '711525000000', '通霄镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711526', '7115', '2', '苑里', 'y', 'yuan li', '711526000000', '苑里镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711527', '7115', '2', '苗栗市', 'm', 'miao li shi', '711527000000', '苗栗市', NULL);
INSERT INTO `demo_city_code` VALUES ('711528', '7115', '2', '造桥', 'z', 'zao qiao', '711528000000', '造桥乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711529', '7115', '2', '头屋', 't', 'tou wu', '711529000000', '头屋乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711530', '7115', '2', '公馆', 'g', 'gong guan', '711530000000', '公馆乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711531', '7115', '2', '大湖', 'd', 'da hu', '711531000000', '大湖乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711532', '7115', '2', '泰安', 't', 'tai an', '711532000000', '泰安乡', '101120801');
INSERT INTO `demo_city_code` VALUES ('711533', '7115', '2', '铜锣', 't', 'tong luo', '711533000000', '铜锣乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711534', '7115', '2', '三义', 's', 'san yi', '711534000000', '三义乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711535', '7115', '2', '西湖', 'x', 'xi hu', '711535000000', '西湖乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711536', '7115', '2', '卓兰', 'z', 'zhuo lan', '711536000000', '卓兰镇', NULL);
INSERT INTO `demo_city_code` VALUES ('7117', '71', '1', '彰化', 'z', 'zhang hua', '711700000000', '彰化县', '101340403');
INSERT INTO `demo_city_code` VALUES ('711727', '7117', '2', '彰化市', 'z', 'zhang hua shi', '711727000000', '彰化市', NULL);
INSERT INTO `demo_city_code` VALUES ('711728', '7117', '2', '芬园', 'f', 'fen yuan', '711728000000', '芬园乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711729', '7117', '2', '花坛', 'h', 'hua tan', '711729000000', '花坛乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711730', '7117', '2', '秀水', 'x', 'xiu shui', '711730000000', '秀水乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711731', '7117', '2', '鹿港', 'l', 'lu gang', '711731000000', '鹿港镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711732', '7117', '2', '福兴', 'f', 'fu xing', '711732000000', '福兴乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711733', '7117', '2', '线西', 'x', 'xian xi', '711733000000', '线西乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711734', '7117', '2', '和美', 'h', 'he mei', '711734000000', '和美镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711735', '7117', '2', '伸港', 's', 'shen gang', '711735000000', '伸港乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711736', '7117', '2', '员林', 'y', 'yuan lin', '711736000000', '员林市', NULL);
INSERT INTO `demo_city_code` VALUES ('711737', '7117', '2', '社头', 's', 'she tou', '711737000000', '社头乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711738', '7117', '2', '永靖', 'y', 'yong jing', '711738000000', '永靖乡', '101161103');
INSERT INTO `demo_city_code` VALUES ('711739', '7117', '2', '埔心', 'b', 'bu xin', '711739000000', '埔心乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711740', '7117', '2', '溪湖', 'x', 'xi hu', '711740000000', '溪湖镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711741', '7117', '2', '大村', 'd', 'da cun', '711741000000', '大村乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711742', '7117', '2', '埔盐', 'b', 'bu yan', '711742000000', '埔盐乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711743', '7117', '2', '田中', 't', 'tian zhong', '711743000000', '田中镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711744', '7117', '2', '北斗', 'b', 'bei dou', '711744000000', '北斗镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711745', '7117', '2', '田尾', 't', 'tian wei', '711745000000', '田尾乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711746', '7117', '2', '埤头', 'p', 'pi tou', '711746000000', '埤头乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711747', '7117', '2', '溪州', 'x', 'xi zhou', '711747000000', '溪州乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711748', '7117', '2', '竹塘', 'z', 'zhu tang', '711748000000', '竹塘乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711749', '7117', '2', '二林', 'e', 'er lin', '711749000000', '二林镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711750', '7117', '2', '大城', 'd', 'da cheng', '711750000000', '大城乡', '101090605');
INSERT INTO `demo_city_code` VALUES ('711751', '7117', '2', '芳苑', 'f', 'fang yuan', '711751000000', '芳苑乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711752', '7117', '2', '二水', 'e', 'er shui', '711752000000', '二水乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7119', '71', '1', '嘉义县', 'j', 'jia yi xian', '711900000000', '嘉义县', NULL);
INSERT INTO `demo_city_code` VALUES ('711919', '7119', '2', '番路', 'f', 'fan lu', '711919000000', '番路乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711920', '7119', '2', '梅山', 'm', 'mei shan', '711920000000', '梅山乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711921', '7119', '2', '竹崎', 'z', 'zhu qi', '711921000000', '竹崎乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711922', '7119', '2', '阿里山', 'a', 'a li shan', '711922000000', '阿里山乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711923', '7119', '2', '中埔', 'z', 'zhong bu', '711923000000', '中埔乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711924', '7119', '2', '大埔', 'd', 'da bu', '711924000000', '大埔乡', '101280404');
INSERT INTO `demo_city_code` VALUES ('711925', '7119', '2', '水上', 's', 'shui shang', '711925000000', '水上乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711926', '7119', '2', '鹿草', 'l', 'lu cao', '711926000000', '鹿草乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711927', '7119', '2', '太保', 't', 'tai bao', '711927000000', '太保市', NULL);
INSERT INTO `demo_city_code` VALUES ('711928', '7119', '2', '朴子', 'p', 'piao zi', '711928000000', '朴子市', NULL);
INSERT INTO `demo_city_code` VALUES ('711929', '7119', '2', '东石', 'd', 'dong shi', '711929000000', '东石乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711930', '7119', '2', '六脚', 'l', 'liu jiao', '711930000000', '六脚乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711931', '7119', '2', '新港', 'x', 'xin gang', '711931000000', '新港乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711932', '7119', '2', '民雄', 'm', 'min xiong', '711932000000', '民雄乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711933', '7119', '2', '大林', 'd', 'da lin', '711933000000', '大林镇', NULL);
INSERT INTO `demo_city_code` VALUES ('711934', '7119', '2', '溪口', 'x', 'xi kou', '711934000000', '溪口乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711935', '7119', '2', '义竹', 'y', 'yi zhu', '711935000000', '义竹乡', NULL);
INSERT INTO `demo_city_code` VALUES ('711936', '7119', '2', '布袋', 'b', 'bu dai', '711936000000', '布袋镇', NULL);
INSERT INTO `demo_city_code` VALUES ('7121', '71', '1', '云林', 'y', 'yun lin', '712100000000', '云林县', '101340406');
INSERT INTO `demo_city_code` VALUES ('712121', '7121', '2', '斗南', 'd', 'dou nan', '712121000000', '斗南镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712122', '7121', '2', '大埤', 'd', 'da pi', '712122000000', '大埤乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712123', '7121', '2', '虎尾', 'h', 'hu wei', '712123000000', '虎尾镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712124', '7121', '2', '土库', 't', 'tu ku', '712124000000', '土库镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712125', '7121', '2', '褒忠', 'b', 'bao zhong', '712125000000', '褒忠乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712126', '7121', '2', '东势', 'd', 'dong shi', '712126000000', '东势乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712127', '7121', '2', '台西', 't', 'tai xi', '712127000000', '台西乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712128', '7121', '2', '仑背', 'l', 'lun bei', '712128000000', '仑背乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712129', '7121', '2', '麦寮', 'm', 'mai liao', '712129000000', '麦寮乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712130', '7121', '2', '斗六', 'd', 'dou liu', '712130000000', '斗六市', NULL);
INSERT INTO `demo_city_code` VALUES ('712131', '7121', '2', '林内', 'l', 'lin nei', '712131000000', '林内乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712132', '7121', '2', '古坑', 'g', 'gu keng', '712132000000', '古坑乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712133', '7121', '2', '莿桐', 'c', 'ci tong', '712133000000', '莿桐乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712134', '7121', '2', '西螺', 'x', 'xi luo', '712134000000', '西螺镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712135', '7121', '2', '二仑', 'e', 'er lun', '712135000000', '二仑乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712136', '7121', '2', '北港', 'b', 'bei gang', '712136000000', '北港镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712137', '7121', '2', '水林', 's', 'shui lin', '712137000000', '水林乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712138', '7121', '2', '口湖', 'k', 'kou hu', '712138000000', '口湖乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712139', '7121', '2', '四湖', 's', 'si hu', '712139000000', '四湖乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712140', '7121', '2', '元长', 'y', 'yuan chang', '712140000000', '元长乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7124', '71', '1', '屏东', 'p', 'ping dong', '712400000000', '屏东县', '101340205');
INSERT INTO `demo_city_code` VALUES ('712434', '7124', '2', '屏东市', 'p', 'ping dong shi', '712434000000', '屏东市', NULL);
INSERT INTO `demo_city_code` VALUES ('712435', '7124', '2', '三地门', 's', 'san di men', '712435000000', '三地门乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712436', '7124', '2', '雾台', 'w', 'wu tai', '712436000000', '雾台乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712437', '7124', '2', '玛家', 'm', 'ma jia', '712437000000', '玛家乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712438', '7124', '2', '九如', 'j', 'jiu ru', '712438000000', '九如乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712439', '7124', '2', '里港', 'l', 'li gang', '712439000000', '里港乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712440', '7124', '2', '高树', 'g', 'gao shu', '712440000000', '高树乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712441', '7124', '2', '盐埔', 'y', 'yan bu', '712441000000', '盐埔乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712442', '7124', '2', '长治', 'c', 'chang zhi', '712442000000', '长治乡', '101100501');
INSERT INTO `demo_city_code` VALUES ('712443', '7124', '2', '麟洛', 'l', 'lin luo', '712443000000', '麟洛乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712444', '7124', '2', '竹田', 'z', 'zhu tian', '712444000000', '竹田乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712445', '7124', '2', '内埔', 'n', 'nei bu', '712445000000', '内埔乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712446', '7124', '2', '万丹', 'w', 'wan dan', '712446000000', '万丹乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712447', '7124', '2', '潮州', 'c', 'chao zhou', '712447000000', '潮州镇', '101281501');
INSERT INTO `demo_city_code` VALUES ('712448', '7124', '2', '泰武', 't', 'tai wu', '712448000000', '泰武乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712449', '7124', '2', '来义', 'l', 'lai yi', '712449000000', '来义乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712450', '7124', '2', '万峦', 'w', 'wan luan', '712450000000', '万峦乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712451', '7124', '2', '崁顶', 'k', 'kan ding', '712451000000', '崁顶乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712452', '7124', '2', '新埤', 'x', 'xin pi', '712452000000', '新埤乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712453', '7124', '2', '南州', 'n', 'nan zhou', '712453000000', '南州乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712454', '7124', '2', '林边', 'l', 'lin bian', '712454000000', '林边乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712455', '7124', '2', '东港', 'd', 'dong gang', '712455000000', '东港镇', '101070604');
INSERT INTO `demo_city_code` VALUES ('712456', '7124', '2', '琉球', 'l', 'liu qiu', '712456000000', '琉球乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712457', '7124', '2', '佳冬', 'j', 'jia dong', '712457000000', '佳冬乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712458', '7124', '2', '新园', 'x', 'xin yuan', '712458000000', '新园乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712459', '7124', '2', '枋寮', 'f', 'fang liao', '712459000000', '枋寮乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712460', '7124', '2', '枋山', 'f', 'fang shan', '712460000000', '枋山乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712461', '7124', '2', '春日', 'c', 'chun ri', '712461000000', '春日乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712462', '7124', '2', '狮子', 's', 'shi zi', '712462000000', '狮子乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712463', '7124', '2', '车城', 'c', 'che cheng', '712463000000', '车城乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712464', '7124', '2', '牡丹', 'm', 'mu dan', '712464000000', '牡丹乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712465', '7124', '2', '恒春', 'h', 'heng chun', '712465000000', '恒春镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712466', '7124', '2', '满州', 'm', 'man zhou', '712466000000', '满州乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7125', '71', '1', '台东', 't', 'tai dong', '712500000000', '台东县', '101340204');
INSERT INTO `demo_city_code` VALUES ('712517', '7125', '2', '台东市', 't', 'tai dong shi', '712517000000', '台东市', NULL);
INSERT INTO `demo_city_code` VALUES ('712518', '7125', '2', '绿岛', 'l', 'lv dao', '712518000000', '绿岛乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712519', '7125', '2', '兰屿', 'l', 'lan yu', '712519000000', '兰屿乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712520', '7125', '2', '延平', 'y', 'yan ping', '712520000000', '延平乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712521', '7125', '2', '卑南', 'b', 'bei nan', '712521000000', '卑南乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712522', '7125', '2', '鹿野', 'l', 'lu ye', '712522000000', '鹿野乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712523', '7125', '2', '关山', 'g', 'guan shan', '712523000000', '关山镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712524', '7125', '2', '海端', 'h', 'hai duan', '712524000000', '海端乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712525', '7125', '2', '池上', 'c', 'chi shang', '712525000000', '池上乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712526', '7125', '2', '东河', 'd', 'dong he', '712526000000', '东河乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712527', '7125', '2', '成功', 'c', 'cheng gong', '712527000000', '成功镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712528', '7125', '2', '长滨', 'c', 'chang bin', '712528000000', '长滨乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712529', '7125', '2', '金峰', 'j', 'jin feng', '712529000000', '金峰乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712530', '7125', '2', '大武', 'd', 'da wu', '712530000000', '大武乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712531', '7125', '2', '达仁', 'd', 'da ren', '712531000000', '达仁乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712532', '7125', '2', '太麻里', 't', 'tai ma li', '712532000000', '太麻里乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7126', '71', '1', '花莲', 'h', 'hua lian', '712600000000', '花莲县', '101340405');
INSERT INTO `demo_city_code` VALUES ('712615', '7126', '2', '花莲市', 'h', 'hua lian shi', '712615000000', '花莲市', NULL);
INSERT INTO `demo_city_code` VALUES ('712616', '7126', '2', '新城', 'x', 'xin cheng', '712616000000', '新城乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712618', '7126', '2', '秀林', 'x', 'xiu lin', '712618000000', '秀林乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712619', '7126', '2', '吉安', 'j', 'ji an', '712619000000', '吉安乡', '101240601');
INSERT INTO `demo_city_code` VALUES ('712620', '7126', '2', '寿丰', 's', 'shou feng', '712620000000', '寿丰乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712621', '7126', '2', '凤林', 'f', 'feng lin', '712621000000', '凤林镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712622', '7126', '2', '光复', 'g', 'guang fu', '712622000000', '光复乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712623', '7126', '2', '丰滨', 'f', 'feng bin', '712623000000', '丰滨乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712624', '7126', '2', '瑞穗', 'r', 'rui sui', '712624000000', '瑞穗乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712625', '7126', '2', '万荣', 'w', 'wan rong', '712625000000', '万荣乡', '101100804');
INSERT INTO `demo_city_code` VALUES ('712626', '7126', '2', '玉里', 'y', 'yu li', '712626000000', '玉里镇', NULL);
INSERT INTO `demo_city_code` VALUES ('712627', '7126', '2', '卓溪', 'z', 'zhuo xi', '712627000000', '卓溪乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712628', '7126', '2', '富里', 'f', 'fu li', '712628000000', '富里乡', NULL);
INSERT INTO `demo_city_code` VALUES ('7127', '71', '1', '澎湖', 'p', 'peng hu', '712700000000', '澎湖县', NULL);
INSERT INTO `demo_city_code` VALUES ('712707', '7127', '2', '马公', 'm', 'ma gong', '712707000000', '马公市', NULL);
INSERT INTO `demo_city_code` VALUES ('712708', '7127', '2', '西屿', 'x', 'xi yu', '712708000000', '西屿乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712709', '7127', '2', '望安', 'w', 'wang an', '712709000000', '望安乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712710', '7127', '2', '七美', 'q', 'qi mei', '712710000000', '七美乡', NULL);
INSERT INTO `demo_city_code` VALUES ('712711', '7127', '2', '白沙', 'b', 'bai sha', '712711000000', '白沙乡', '101310207');
INSERT INTO `demo_city_code` VALUES ('712712', '7127', '2', '湖西', 'h', 'hu xi', '712712000000', '湖西乡', NULL);
INSERT INTO `demo_city_code` VALUES ('81', '0', '0', '香港', '~1', 'xiang gang', '810000000000', '香港特别行政区', '101320101');
INSERT INTO `demo_city_code` VALUES ('8100', '81', '1', '香港', 'x', 'xiang gang', '810000000000', '香港特别行政区', '101320101');
INSERT INTO `demo_city_code` VALUES ('810000', '8100', '2', '香港', 'x', 'xiang gang', '810000000000', '香港特别行政区', '101320101');
INSERT INTO `demo_city_code` VALUES ('82', '0', '0', '澳门', '~2', 'ao men', '820000000000', '澳门特别行政区', '101330101');
INSERT INTO `demo_city_code` VALUES ('8200', '82', '1', '澳门', 'a', 'ao men', '820000000000', '澳门特别行政区', '101330101');
INSERT INTO `demo_city_code` VALUES ('820000', '8200', '2', '澳门', 'a', 'ao men', '820000000000', '澳门特别行政区', '101330101');
INSERT INTO `demo_city_code` VALUES ('91', '0', '0', '国外', '~4', 'guo wai', '0', '国外', NULL);
INSERT INTO `demo_city_code` VALUES ('9100', '91', '1', '国外', 'g', 'guo wai', '0', '国外', NULL);
INSERT INTO `demo_city_code` VALUES ('910000', '9100', '2', '国外', 'g', 'guo wai', '0', '国外', NULL);

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
INSERT INTO `sys_job_log` VALUES (6, 'checkUserOnline', '2024-04-12 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (7, 'checkUserOnline', '2024-04-30 15:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (8, 'checkUserOnline', '2024-04-30 16:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (9, 'checkUserOnline', '2024-05-06 09:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (10, 'checkUserOnline', '2024-05-06 09:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (11, 'checkUserOnline', '2024-05-06 09:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (12, 'checkUserOnline', '2024-05-06 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (13, 'checkUserOnline', '2024-05-06 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (14, 'checkUserOnline', '2024-05-06 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (15, 'checkUserOnline', '2024-05-06 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (16, 'checkUserOnline', '2024-05-06 10:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (17, 'checkUserOnline', '2024-05-06 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (18, 'checkUserOnline', '2024-05-06 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (19, 'checkUserOnline', '2024-05-06 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (20, 'checkUserOnline', '2024-05-06 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (21, 'checkUserOnline', '2024-05-06 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (22, 'checkUserOnline', '2024-05-06 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (23, 'checkUserOnline', '2024-05-06 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (24, 'checkUserOnline', '2024-05-06 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (25, 'checkUserOnline', '2024-05-06 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (26, 'checkUserOnline', '2024-05-06 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (27, 'checkUserOnline', '2024-05-06 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (28, 'checkUserOnline', '2024-05-06 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (29, 'checkUserOnline', '2024-05-06 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (30, 'checkUserOnline', '2024-05-06 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (31, 'checkUserOnline', '2024-05-06 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (32, 'checkUserOnline', '2024-05-06 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (33, 'checkUserOnline', '2024-05-06 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (34, 'checkUserOnline', '2024-05-06 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (35, 'checkUserOnline', '2024-05-06 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (36, 'checkUserOnline', '2024-05-06 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (37, 'checkUserOnline', '2024-05-06 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (38, 'checkUserOnline', '2024-05-06 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (39, 'checkUserOnline', '2024-05-06 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (40, 'checkUserOnline', '2024-05-06 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (41, 'checkUserOnline', '2024-05-06 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (42, 'checkUserOnline', '2024-05-06 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (43, 'checkUserOnline', '2024-05-06 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (44, 'checkUserOnline', '2024-05-06 14:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (45, 'checkUserOnline', '2024-05-06 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (46, 'checkUserOnline', '2024-05-06 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (47, 'checkUserOnline', '2024-05-06 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (48, 'checkUserOnline', '2024-05-06 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (49, 'checkUserOnline', '2024-05-06 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (50, 'checkUserOnline', '2024-05-06 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (51, 'checkUserOnline', '2024-05-06 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (52, 'checkUserOnline', '2024-05-06 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (53, 'checkUserOnline', '2024-05-06 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (54, 'checkUserOnline', '2024-05-06 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (55, 'checkUserOnline', '2024-05-06 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (56, 'checkUserOnline', '2024-05-06 16:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (57, 'checkUserOnline', '2024-05-06 17:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (58, 'checkUserOnline', '2024-05-06 17:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (59, 'checkUserOnline', '2024-05-06 17:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (60, 'checkUserOnline', '2024-05-06 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (61, 'checkUserOnline', '2024-05-06 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (62, 'checkUserOnline', '2024-05-06 17:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (63, 'checkUserOnline', '2024-05-06 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (64, 'checkUserOnline', '2024-05-06 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (65, 'checkUserOnline', '2024-05-07 08:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (66, 'checkUserOnline', '2024-05-07 09:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (67, 'checkUserOnline', '2024-05-07 09:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (68, 'checkUserOnline', '2024-05-07 09:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (69, 'checkUserOnline', '2024-05-07 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (70, 'checkUserOnline', '2024-05-07 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (71, 'checkUserOnline', '2024-05-07 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (72, 'checkUserOnline', '2024-05-07 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (73, 'checkUserOnline', '2024-05-07 10:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (74, 'checkUserOnline', '2024-05-07 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (75, 'checkUserOnline', '2024-05-07 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (76, 'checkUserOnline', '2024-05-07 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (77, 'checkUserOnline', '2024-05-07 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (78, 'checkUserOnline', '2024-05-07 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (79, 'checkUserOnline', '2024-05-07 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (80, 'checkUserOnline', '2024-05-07 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (81, 'checkUserOnline', '2024-05-07 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (82, 'checkUserOnline', '2024-05-07 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (83, 'checkUserOnline', '2024-05-07 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (84, 'checkUserOnline', '2024-05-07 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (85, 'checkUserOnline', '2024-05-07 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (86, 'checkUserOnline', '2024-05-07 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (87, 'checkUserOnline', '2024-05-07 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (88, 'checkUserOnline', '2024-05-07 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (89, 'checkUserOnline', '2024-05-07 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (90, 'checkUserOnline', '2024-05-07 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (91, 'checkUserOnline', '2024-05-07 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (92, 'checkUserOnline', '2024-05-07 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (93, 'checkUserOnline', '2024-05-07 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (94, 'checkUserOnline', '2024-05-07 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (95, 'checkUserOnline', '2024-05-07 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (96, 'checkUserOnline', '2024-05-07 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (97, 'checkUserOnline', '2024-05-07 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (98, 'checkUserOnline', '2024-05-07 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (99, 'checkUserOnline', '2024-05-07 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (100, 'checkUserOnline', '2024-05-07 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (101, 'checkUserOnline', '2024-05-07 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (102, 'checkUserOnline', '2024-05-07 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (103, 'checkUserOnline', '2024-05-07 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (104, 'checkUserOnline', '2024-05-07 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (105, 'checkUserOnline', '2024-07-09 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (106, 'checkUserOnline', '2024-07-09 15:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (107, 'checkUserOnline', '2024-07-09 15:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (108, 'checkUserOnline', '2024-07-09 16:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (109, 'checkUserOnline', '2024-07-09 16:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (110, 'checkUserOnline', '2024-07-09 16:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (111, 'checkUserOnline', '2024-07-09 16:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (112, 'checkUserOnline', '2024-07-19 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (113, 'checkUserOnline', '2024-07-19 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (114, 'checkUserOnline', '2024-07-19 17:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (115, 'checkUserOnline', '2024-07-19 18:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (116, 'checkUserOnline', '2024-07-19 18:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (117, 'checkUserOnline', '2024-07-19 18:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (118, 'checkUserOnline', '2024-07-22 09:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (119, 'checkUserOnline', '2024-07-22 09:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (120, 'checkUserOnline', '2024-07-22 09:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (121, 'checkUserOnline', '2024-07-22 09:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (122, 'checkUserOnline', '2024-07-22 09:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (123, 'checkUserOnline', '2024-07-22 10:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (124, 'checkUserOnline', '2024-07-22 10:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (125, 'checkUserOnline', '2024-07-22 10:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (126, 'checkUserOnline', '2024-07-22 10:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (127, 'checkUserOnline', '2024-07-22 10:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (128, 'checkUserOnline', '2024-07-22 10:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (129, 'checkUserOnline', '2024-07-22 11:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (130, 'checkUserOnline', '2024-07-22 11:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (131, 'checkUserOnline', '2024-07-22 12:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (132, 'checkUserOnline', '2024-07-22 12:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (133, 'checkUserOnline', '2024-07-22 12:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (134, 'checkUserOnline', '2024-07-22 12:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (135, 'checkUserOnline', '2024-07-22 12:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (136, 'checkUserOnline', '2024-07-22 12:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (137, 'checkUserOnline', '2024-07-22 13:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (138, 'checkUserOnline', '2024-07-22 13:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (139, 'checkUserOnline', '2024-07-22 13:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (140, 'checkUserOnline', '2024-07-22 13:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (141, 'checkUserOnline', '2024-07-22 13:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (142, 'checkUserOnline', '2024-07-22 13:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (143, 'checkUserOnline', '2024-07-22 14:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (144, 'checkUserOnline', '2024-07-22 14:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (145, 'checkUserOnline', '2024-07-22 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (146, 'checkUserOnline', '2024-07-22 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (147, 'checkUserOnline', '2024-07-22 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (148, 'checkUserOnline', '2024-07-22 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (149, 'checkUserOnline', '2024-07-22 15:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (150, 'checkUserOnline', '2024-07-22 15:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (151, 'checkUserOnline', '2024-07-22 15:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (152, 'checkUserOnline', '2024-07-22 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (153, 'checkUserOnline', '2024-07-22 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (154, 'checkUserOnline', '2024-07-22 16:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (155, 'checkUserOnline', '2024-07-22 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (156, 'checkUserOnline', '2024-07-22 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (157, 'checkUserOnline', '2024-07-22 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (158, 'checkUserOnline', '2024-07-22 16:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (159, 'checkUserOnline', '2024-07-22 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (160, 'checkUserOnline', '2024-07-22 17:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (161, 'checkUserOnline', '2024-07-22 17:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (162, 'checkUserOnline', '2024-07-22 17:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (163, 'checkUserOnline', '2024-07-22 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (164, 'checkUserOnline', '2024-07-22 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (165, 'checkUserOnline', '2024-07-22 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (166, 'checkUserOnline', '2024-07-22 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (167, 'checkUserOnline', '2024-07-23 08:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (168, 'checkUserOnline', '2024-07-23 08:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (169, 'checkUserOnline', '2024-07-23 08:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (170, 'checkUserOnline', '2024-07-23 09:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (171, 'checkUserOnline', '2024-07-23 09:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (172, 'checkUserOnline', '2024-07-23 09:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (173, 'checkUserOnline', '2024-07-23 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (174, 'checkUserOnline', '2024-07-23 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (175, 'checkUserOnline', '2024-07-23 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (176, 'checkUserOnline', '2024-07-23 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (177, 'checkUserOnline', '2024-07-23 10:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (178, 'checkUserOnline', '2024-07-23 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (179, 'checkUserOnline', '2024-07-23 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (180, 'checkUserOnline', '2024-07-23 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (181, 'checkUserOnline', '2024-07-23 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (182, 'checkUserOnline', '2024-07-23 11:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (183, 'checkUserOnline', '2024-07-23 11:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (184, 'checkUserOnline', '2024-07-23 11:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (185, 'checkUserOnline', '2024-07-23 11:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (186, 'checkUserOnline', '2024-07-23 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (187, 'checkUserOnline', '2024-07-23 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (188, 'checkUserOnline', '2024-07-23 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (189, 'checkUserOnline', '2024-07-23 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (190, 'checkUserOnline', '2024-07-23 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (191, 'checkUserOnline', '2024-07-23 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (192, 'checkUserOnline', '2024-07-23 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (193, 'checkUserOnline', '2024-07-23 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (194, 'checkUserOnline', '2024-07-23 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (195, 'checkUserOnline', '2024-07-23 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (196, 'checkUserOnline', '2024-07-23 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (197, 'checkUserOnline', '2024-07-23 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (198, 'checkUserOnline', '2024-07-23 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (199, 'checkUserOnline', '2024-07-23 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (200, 'checkUserOnline', '2024-07-23 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (201, 'checkUserOnline', '2024-07-23 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (202, 'checkUserOnline', '2024-07-23 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (203, 'checkUserOnline', '2024-07-23 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (204, 'checkUserOnline', '2024-07-23 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (205, 'checkUserOnline', '2024-07-23 14:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (206, 'checkUserOnline', '2024-07-23 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (207, 'checkUserOnline', '2024-07-23 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (208, 'checkUserOnline', '2024-07-23 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (209, 'checkUserOnline', '2024-07-23 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (210, 'checkUserOnline', '2024-07-23 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (211, 'checkUserOnline', '2024-07-23 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (212, 'checkUserOnline', '2024-07-23 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (213, 'checkUserOnline', '2024-07-23 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (214, 'checkUserOnline', '2024-07-23 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (215, 'checkUserOnline', '2024-07-23 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (216, 'checkUserOnline', '2024-07-23 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (217, 'checkUserOnline', '2024-07-23 16:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (218, 'checkUserOnline', '2024-07-23 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (219, 'checkUserOnline', '2024-07-23 17:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (220, 'checkUserOnline', '2024-07-23 17:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (221, 'checkUserOnline', '2024-07-23 17:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (222, 'checkUserOnline', '2024-07-23 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (223, 'checkUserOnline', '2024-07-23 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (224, 'checkUserOnline', '2024-07-23 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (225, 'checkUserOnline', '2024-07-27 11:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (226, 'checkUserOnline', '2024-07-27 11:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (227, 'checkUserOnline', '2024-07-29 08:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (228, 'checkUserOnline', '2024-09-13 17:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (229, 'checkUserOnline', '2024-09-13 17:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (230, 'checkUserOnline', '2024-09-13 17:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (231, 'checkUserOnline', '2024-09-13 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (232, 'checkUserOnline', '2024-09-13 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (233, 'checkUserOnline', '2024-09-13 17:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (234, 'checkUserOnline', '2024-09-13 18:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (235, 'checkUserOnline', '2024-09-13 18:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (236, 'checkUserOnline', '2024-09-13 18:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (237, 'checkUserOnline', '2024-09-13 18:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (238, 'checkUserOnline', '2024-09-13 18:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (239, 'checkUserOnline', '2024-09-13 18:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (240, 'checkUserOnline', '2024-09-13 19:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (241, 'checkUserOnline', '2024-09-13 19:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (242, 'checkUserOnline', '2024-09-13 19:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (243, 'checkUserOnline', '2024-09-13 19:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (244, 'checkUserOnline', '2024-09-13 19:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (245, 'checkUserOnline', '2024-09-13 19:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (246, 'checkUserOnline', '2024-09-13 20:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (247, 'checkUserOnline', '2024-09-13 20:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (248, 'checkUserOnline', '2024-09-13 20:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (249, 'checkUserOnline', '2024-09-13 20:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (250, 'checkUserOnline', '2024-09-13 20:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (251, 'checkUserOnline', '2024-09-13 20:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (252, 'checkUserOnline', '2024-09-13 21:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (253, 'checkUserOnline', '2024-09-13 21:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (254, 'checkUserOnline', '2024-09-13 21:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (255, 'checkUserOnline', '2024-09-13 21:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (256, 'checkUserOnline', '2024-09-13 21:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (257, 'checkUserOnline', '2024-09-13 21:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (258, 'checkUserOnline', '2024-09-13 22:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (259, 'checkUserOnline', '2024-09-13 22:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (260, 'checkUserOnline', '2024-09-13 22:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (261, 'checkUserOnline', '2024-09-14 08:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (262, 'checkUserOnline', '2024-09-14 09:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (263, 'checkUserOnline', '2024-09-14 09:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (264, 'checkUserOnline', '2024-09-14 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (265, 'checkUserOnline', '2024-09-14 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (266, 'checkUserOnline', '2024-09-14 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (267, 'checkUserOnline', '2024-09-14 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (268, 'checkUserOnline', '2024-09-14 10:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (269, 'checkUserOnline', '2024-09-14 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (270, 'checkUserOnline', '2024-09-14 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (271, 'checkUserOnline', '2024-09-14 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (272, 'checkUserOnline', '2024-09-14 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (273, 'checkUserOnline', '2024-09-14 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (274, 'checkUserOnline', '2024-09-14 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (275, 'checkUserOnline', '2024-09-14 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (276, 'checkUserOnline', '2024-09-14 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (277, 'checkUserOnline', '2024-09-14 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (278, 'checkUserOnline', '2024-09-30 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (279, 'checkUserOnline', '2024-10-08 16:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (280, 'checkUserOnline', '2024-10-08 16:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (281, 'checkUserOnline', '2024-10-08 16:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (282, 'checkUserOnline', '2024-10-08 17:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (283, 'checkUserOnline', '2024-10-22 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (284, 'checkUserOnline', '2024-10-22 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (285, 'checkUserOnline', '2024-10-22 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (286, 'checkUserOnline', '2024-10-22 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (287, 'checkUserOnline', '2024-10-22 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (288, 'checkUserOnline', '2024-10-22 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (289, 'checkUserOnline', '2024-10-22 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (290, 'checkUserOnline', '2024-10-22 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (291, 'checkUserOnline', '2024-10-22 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (292, 'checkUserOnline', '2024-10-22 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (293, 'checkUserOnline', '2024-10-22 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (294, 'checkUserOnline', '2024-10-22 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (295, 'checkUserOnline', '2024-10-22 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (296, 'checkUserOnline', '2024-10-22 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (297, 'checkUserOnline', '2024-10-22 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (298, 'checkUserOnline', '2024-10-22 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (299, 'checkUserOnline', '2024-10-22 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (300, 'checkUserOnline', '2024-10-22 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (301, 'checkUserOnline', '2024-10-22 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (302, 'checkUserOnline', '2024-10-22 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (303, 'checkUserOnline', '2024-10-22 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (304, 'checkUserOnline', '2024-10-22 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (305, 'checkUserOnline', '2024-10-22 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (306, 'checkUserOnline', '2024-10-22 14:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (307, 'checkUserOnline', '2024-10-22 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (308, 'checkUserOnline', '2024-10-22 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (309, 'checkUserOnline', '2024-11-18 11:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (310, 'checkUserOnline', '2024-11-18 12:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (311, 'checkUserOnline', '2024-11-18 12:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (312, 'checkUserOnline', '2024-11-18 12:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (313, 'checkUserOnline', '2024-11-18 12:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (314, 'checkUserOnline', '2024-11-18 12:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (315, 'checkUserOnline', '2024-11-18 12:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (316, 'checkUserOnline', '2024-11-18 13:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (317, 'checkUserOnline', '2024-11-18 13:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (318, 'checkUserOnline', '2024-11-18 13:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (319, 'checkUserOnline', '2024-11-18 13:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (320, 'checkUserOnline', '2024-11-18 13:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (321, 'checkUserOnline', '2024-11-18 13:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (322, 'checkUserOnline', '2024-11-18 14:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (323, 'checkUserOnline', '2024-11-18 14:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (324, 'checkUserOnline', '2024-11-18 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (325, 'checkUserOnline', '2024-11-18 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (326, 'checkUserOnline', '2024-11-18 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (327, 'checkUserOnline', '2024-11-18 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (328, 'checkUserOnline', '2024-11-18 15:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (329, 'checkUserOnline', '2024-11-18 15:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (330, 'checkUserOnline', '2024-11-18 15:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (331, 'checkUserOnline', '2024-11-18 15:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (332, 'checkUserOnline', '2024-11-18 15:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (333, 'checkUserOnline', '2024-11-18 15:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (334, 'checkUserOnline', '2024-11-18 16:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (335, 'checkUserOnline', '2024-12-19 16:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (336, 'checkUserOnline', '2024-12-19 16:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (337, 'checkUserOnline', '2024-12-19 16:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (338, 'checkUserOnline', '2024-12-19 17:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (339, 'checkUserOnline', '2025-03-11 17:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (340, 'checkUserOnline', '2025-03-11 17:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (341, 'checkUserOnline', '2025-03-11 17:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (342, 'checkUserOnline', '2025-03-11 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (343, 'checkUserOnline', '2025-04-25 10:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (344, 'checkUserOnline', '2025-04-25 10:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (345, 'checkUserOnline', '2025-04-25 10:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (346, 'checkUserOnline', '2025-04-25 11:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (347, 'checkUserOnline', '2025-04-25 11:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (348, 'checkUserOnline', '2025-04-25 11:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (349, 'checkUserOnline', '2025-04-25 11:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (350, 'checkUserOnline', '2025-04-25 11:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (351, 'checkUserOnline', '2025-04-25 11:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (352, 'checkUserOnline', '2025-04-25 12:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (353, 'checkUserOnline', '2025-04-25 12:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (354, 'checkUserOnline', '2025-04-25 12:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (355, 'checkUserOnline', '2025-04-25 12:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (356, 'checkUserOnline', '2025-04-25 12:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (357, 'checkUserOnline', '2025-04-25 12:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (358, 'checkUserOnline', '2025-04-25 13:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (359, 'checkUserOnline', '2025-04-25 13:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (360, 'checkUserOnline', '2025-04-25 13:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (361, 'checkUserOnline', '2025-04-25 13:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (362, 'checkUserOnline', '2025-04-25 13:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (363, 'checkUserOnline', '2025-04-25 13:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (364, 'checkUserOnline', '2025-04-25 14:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (365, 'checkUserOnline', '2025-04-25 14:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (366, 'checkUserOnline', '2025-04-25 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (367, 'checkUserOnline', '2025-04-25 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (368, 'checkUserOnline', '2025-04-25 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (369, 'checkUserOnline', '2025-04-25 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (370, 'checkUserOnline', '2025-04-25 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (371, 'checkUserOnline', '2025-04-25 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (372, 'checkUserOnline', '2025-04-25 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (373, 'checkUserOnline', '2025-04-25 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (374, 'checkUserOnline', '2025-04-25 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (375, 'checkUserOnline', '2025-04-25 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (376, 'checkUserOnline', '2025-04-25 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (377, 'checkUserOnline', '2025-04-25 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (378, 'checkUserOnline', '2025-04-25 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (379, 'checkUserOnline', '2025-04-25 16:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (380, 'checkUserOnline', '2025-04-25 16:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (381, 'checkUserOnline', '2025-04-25 16:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (382, 'checkUserOnline', '2025-04-25 17:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (383, 'checkUserOnline', '2025-04-25 17:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (384, 'checkUserOnline', '2025-04-25 17:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (385, 'checkUserOnline', '2025-04-25 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (386, 'checkUserOnline', '2025-04-25 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (387, 'checkUserOnline', '2025-04-25 17:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (388, 'checkUserOnline', '2025-04-25 18:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (389, 'checkUserOnline', '2025-04-25 18:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (390, 'checkUserOnline', '2025-04-27 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (391, 'checkUserOnline', '2025-04-27 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (392, 'checkUserOnline', '2025-04-27 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (393, 'checkUserOnline', '2025-04-27 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (394, 'checkUserOnline', '2025-04-27 10:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (395, 'checkUserOnline', '2025-04-27 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (396, 'checkUserOnline', '2025-04-27 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (397, 'checkUserOnline', '2025-04-27 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (398, 'checkUserOnline', '2025-04-27 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (399, 'checkUserOnline', '2025-04-27 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (400, 'checkUserOnline', '2025-04-27 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (401, 'checkUserOnline', '2025-04-27 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (402, 'checkUserOnline', '2025-04-27 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (403, 'checkUserOnline', '2025-04-27 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (404, 'checkUserOnline', '2025-04-27 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (405, 'checkUserOnline', '2025-04-27 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (406, 'checkUserOnline', '2025-04-27 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (407, 'checkUserOnline', '2025-04-27 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (408, 'checkUserOnline', '2025-04-27 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (409, 'checkUserOnline', '2025-04-27 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (410, 'checkUserOnline', '2025-04-27 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (411, 'checkUserOnline', '2025-04-27 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (412, 'checkUserOnline', '2025-04-27 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (413, 'checkUserOnline', '2025-04-27 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (414, 'checkUserOnline', '2025-04-27 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (415, 'checkUserOnline', '2025-04-27 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (416, 'checkUserOnline', '2025-04-27 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (417, 'checkUserOnline', '2025-04-27 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (418, 'checkUserOnline', '2025-04-27 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (419, 'checkUserOnline', '2025-04-27 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (420, 'checkUserOnline', '2025-04-27 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (421, 'checkUserOnline', '2025-04-27 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (422, 'checkUserOnline', '2025-04-27 14:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (423, 'checkUserOnline', '2025-04-27 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (424, 'checkUserOnline', '2025-04-27 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (425, 'checkUserOnline', '2025-04-27 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (426, 'checkUserOnline', '2025-04-27 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (427, 'checkUserOnline', '2025-04-27 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (428, 'checkUserOnline', '2025-04-27 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (429, 'checkUserOnline', '2025-04-27 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (430, 'checkUserOnline', '2025-04-27 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (431, 'checkUserOnline', '2025-04-27 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (432, 'checkUserOnline', '2025-04-27 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (433, 'checkUserOnline', '2025-04-27 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (434, 'checkUserOnline', '2025-04-27 16:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (435, 'checkUserOnline', '2025-04-27 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (436, 'checkUserOnline', '2025-04-27 17:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (437, 'checkUserOnline', '2025-04-27 17:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (438, 'checkUserOnline', '2025-04-27 17:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (439, 'checkUserOnline', '2025-04-27 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (440, 'checkUserOnline', '2025-04-27 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (441, 'checkUserOnline', '2025-04-27 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (442, 'checkUserOnline', '2025-04-27 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (443, 'checkUserOnline', '2025-04-28 08:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (444, 'checkUserOnline', '2025-04-28 09:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (445, 'checkUserOnline', '2025-04-28 09:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (446, 'checkUserOnline', '2025-04-28 09:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (447, 'checkUserOnline', '2025-04-28 09:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (448, 'checkUserOnline', '2025-04-28 09:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (449, 'checkUserOnline', '2025-04-28 09:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (450, 'checkUserOnline', '2025-04-28 10:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (451, 'checkUserOnline', '2025-04-28 10:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (452, 'checkUserOnline', '2025-04-28 10:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (453, 'checkUserOnline', '2025-04-28 10:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (454, 'checkUserOnline', '2025-04-28 10:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (455, 'checkUserOnline', '2025-04-28 10:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (456, 'checkUserOnline', '2025-04-28 11:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (457, 'checkUserOnline', '2025-04-28 11:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (458, 'checkUserOnline', '2025-04-28 11:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (459, 'checkUserOnline', '2025-04-28 11:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (460, 'checkUserOnline', '2025-04-28 11:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (461, 'checkUserOnline', '2025-04-28 11:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (462, 'checkUserOnline', '2025-04-28 12:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (463, 'checkUserOnline', '2025-04-28 12:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (464, 'checkUserOnline', '2025-04-28 12:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (465, 'checkUserOnline', '2025-04-28 12:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (466, 'checkUserOnline', '2025-04-28 12:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (467, 'checkUserOnline', '2025-04-28 12:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (468, 'checkUserOnline', '2025-04-28 13:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (469, 'checkUserOnline', '2025-04-28 13:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (470, 'checkUserOnline', '2025-04-28 13:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (471, 'checkUserOnline', '2025-04-28 13:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (472, 'checkUserOnline', '2025-04-28 13:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (473, 'checkUserOnline', '2025-04-28 13:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (474, 'checkUserOnline', '2025-04-28 14:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (475, 'checkUserOnline', '2025-04-28 14:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (476, 'checkUserOnline', '2025-04-28 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (477, 'checkUserOnline', '2025-04-28 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (478, 'checkUserOnline', '2025-04-28 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (479, 'checkUserOnline', '2025-04-28 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (480, 'checkUserOnline', '2025-04-28 15:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (481, 'checkUserOnline', '2025-04-28 15:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (482, 'checkUserOnline', '2025-04-28 15:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (483, 'checkUserOnline', '2025-04-28 15:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (484, 'checkUserOnline', '2025-05-22 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (485, 'checkUserOnline', '2025-05-22 16:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (486, 'checkUserOnline', '2025-05-22 16:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (487, 'checkUserOnline', '2025-05-22 16:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (488, 'checkUserOnline', '2025-05-22 16:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (489, 'checkUserOnline', '2025-05-22 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (490, 'checkUserOnline', '2025-05-22 17:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (491, 'checkUserOnline', '2025-05-22 17:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (492, 'checkUserOnline', '2025-05-22 17:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (493, 'checkUserOnline', '2025-05-22 17:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (494, 'checkUserOnline', '2025-05-22 17:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (495, 'checkUserOnline', '2025-05-22 18:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (496, 'checkUserOnline', '2025-05-22 18:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (497, 'checkUserOnline', '2025-05-30 10:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (498, 'checkUserOnline', '2025-05-30 10:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (499, 'checkUserOnline', '2025-05-30 11:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (500, 'checkUserOnline', '2025-05-30 11:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (501, 'checkUserOnline', '2025-05-30 11:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (502, 'checkUserOnline', '2025-05-30 11:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (503, 'checkUserOnline', '2025-05-30 11:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (504, 'checkUserOnline', '2025-05-30 11:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (505, 'checkUserOnline', '2025-05-30 12:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (506, 'checkUserOnline', '2025-05-30 12:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (507, 'checkUserOnline', '2025-05-30 12:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (508, 'checkUserOnline', '2025-05-30 12:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (509, 'checkUserOnline', '2025-05-30 12:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (510, 'checkUserOnline', '2025-05-30 12:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (511, 'checkUserOnline', '2025-05-30 13:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (512, 'checkUserOnline', '2025-05-30 13:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (513, 'checkUserOnline', '2025-05-30 13:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (514, 'checkUserOnline', '2025-05-30 13:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (515, 'checkUserOnline', '2025-05-30 13:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (516, 'checkUserOnline', '2025-05-30 13:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (517, 'checkUserOnline', '2025-05-30 14:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (518, 'checkUserOnline', '2025-05-30 14:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (519, 'checkUserOnline', '2025-05-30 14:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (520, 'checkUserOnline', '2025-05-30 14:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (521, 'checkUserOnline', '2025-05-30 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (522, 'checkUserOnline', '2025-05-30 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (523, 'checkUserOnline', '2025-05-30 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (524, 'checkUserOnline', '2025-05-30 15:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (525, 'checkUserOnline', '2025-05-30 15:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (526, 'checkUserOnline', '2025-05-30 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (527, 'checkUserOnline', '2025-05-30 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (528, 'checkUserOnline', '2025-05-30 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (529, 'checkUserOnline', '2025-05-30 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (530, 'checkUserOnline', '2025-05-30 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (531, 'checkUserOnline', '2025-05-30 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (532, 'checkUserOnline', '2025-05-30 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (533, 'checkUserOnline', '2025-06-10 15:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (534, 'checkUserOnline', '2025-06-11 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (535, 'checkUserOnline', '2025-06-11 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (536, 'checkUserOnline', '2025-06-12 09:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (537, 'checkUserOnline', '2025-06-12 09:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (538, 'checkUserOnline', '2025-06-12 09:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (539, 'checkUserOnline', '2025-06-12 09:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (540, 'checkUserOnline', '2025-06-12 10:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (541, 'checkUserOnline', '2025-06-12 10:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (542, 'checkUserOnline', '2025-06-12 10:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (543, 'checkUserOnline', '2025-06-12 10:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (544, 'checkUserOnline', '2025-06-12 10:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (545, 'checkUserOnline', '2025-06-12 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (546, 'checkUserOnline', '2025-06-12 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (547, 'checkUserOnline', '2025-06-12 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (548, 'checkUserOnline', '2025-06-12 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (549, 'checkUserOnline', '2025-06-12 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (550, 'checkUserOnline', '2025-06-12 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (551, 'checkUserOnline', '2025-06-12 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (552, 'checkUserOnline', '2025-06-12 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (553, 'checkUserOnline', '2025-06-12 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (554, 'checkUserOnline', '2025-06-12 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (555, 'checkUserOnline', '2025-06-12 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (556, 'checkUserOnline', '2025-06-12 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (557, 'checkUserOnline', '2025-06-12 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (558, 'checkUserOnline', '2025-06-12 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (559, 'checkUserOnline', '2025-06-12 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (560, 'checkUserOnline', '2025-06-12 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (561, 'checkUserOnline', '2025-06-12 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (562, 'checkUserOnline', '2025-06-12 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (563, 'checkUserOnline', '2025-06-12 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (564, 'checkUserOnline', '2025-06-12 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (565, 'checkUserOnline', '2025-06-12 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (566, 'checkUserOnline', '2025-06-12 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (567, 'checkUserOnline', '2025-06-12 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (568, 'checkUserOnline', '2025-06-12 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (569, 'checkUserOnline', '2025-06-12 14:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (570, 'checkUserOnline', '2025-06-12 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (571, 'checkUserOnline', '2025-06-12 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (572, 'checkUserOnline', '2025-06-12 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (573, 'checkUserOnline', '2025-06-12 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (574, 'checkUserOnline', '2025-06-12 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (575, 'checkUserOnline', '2025-06-12 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (576, 'checkUserOnline', '2025-06-12 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (577, 'checkUserOnline', '2025-08-13 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (578, 'checkUserOnline', '2025-08-13 11:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (579, 'checkUserOnline', '2025-08-13 11:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (580, 'checkUserOnline', '2025-08-13 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (581, 'checkUserOnline', '2025-08-13 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (582, 'checkUserOnline', '2025-08-13 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (583, 'checkUserOnline', '2025-08-13 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (584, 'checkUserOnline', '2025-08-13 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (585, 'checkUserOnline', '2025-08-13 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (586, 'checkUserOnline', '2025-08-13 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (587, 'checkUserOnline', '2025-08-13 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (588, 'checkUserOnline', '2025-08-13 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (589, 'checkUserOnline', '2025-08-13 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (590, 'checkUserOnline', '2025-08-13 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (591, 'checkUserOnline', '2025-08-13 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (592, 'checkUserOnline', '2025-08-13 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (593, 'checkUserOnline', '2025-08-13 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (594, 'checkUserOnline', '2025-08-13 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (595, 'checkUserOnline', '2025-08-13 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (596, 'checkUserOnline', '2025-08-13 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (597, 'checkUserOnline', '2025-08-13 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (598, 'checkUserOnline', '2025-08-13 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (599, 'checkUserOnline', '2025-08-13 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (600, 'checkUserOnline', '2025-08-13 14:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (601, 'checkUserOnline', '2025-08-13 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (602, 'checkUserOnline', '2025-08-13 15:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (603, 'checkUserOnline', '2025-08-13 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (604, 'checkUserOnline', '2025-08-13 15:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (605, 'checkUserOnline', '2025-08-13 15:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (606, 'checkUserOnline', '2025-08-13 15:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (607, 'checkUserOnline', '2025-08-13 15:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (608, 'checkUserOnline', '2025-08-13 16:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (609, 'checkUserOnline', '2025-08-13 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (610, 'checkUserOnline', '2025-08-13 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (611, 'checkUserOnline', '2025-08-13 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (612, 'checkUserOnline', '2025-08-13 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (613, 'checkUserOnline', '2025-08-13 16:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (614, 'checkUserOnline', '2025-08-13 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (615, 'checkUserOnline', '2025-08-13 17:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (616, 'checkUserOnline', '2025-08-13 17:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (617, 'checkUserOnline', '2025-08-13 17:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (618, 'checkUserOnline', '2025-08-13 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (619, 'checkUserOnline', '2025-08-13 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (620, 'checkUserOnline', '2025-08-13 18:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (621, 'checkUserOnline', '2025-08-13 18:10:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (622, 'checkUserOnline', '2025-08-13 18:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (623, 'checkUserOnline', '2025-08-14 11:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (624, 'checkUserOnline', '2025-08-14 11:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (625, 'checkUserOnline', '2025-08-14 11:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (626, 'checkUserOnline', '2025-08-14 11:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (627, 'checkUserOnline', '2025-08-14 12:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (628, 'checkUserOnline', '2025-08-14 12:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (629, 'checkUserOnline', '2025-08-14 12:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (630, 'checkUserOnline', '2025-08-14 12:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (631, 'checkUserOnline', '2025-08-14 12:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (632, 'checkUserOnline', '2025-08-14 12:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (633, 'checkUserOnline', '2025-08-14 13:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (634, 'checkUserOnline', '2025-08-14 13:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (635, 'checkUserOnline', '2025-08-14 13:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (636, 'checkUserOnline', '2025-08-14 13:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (637, 'checkUserOnline', '2025-08-14 13:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (638, 'checkUserOnline', '2025-08-14 13:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (639, 'checkUserOnline', '2025-08-14 14:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (640, 'checkUserOnline', '2025-08-14 14:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (641, 'checkUserOnline', '2025-08-14 14:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (642, 'checkUserOnline', '2025-08-14 14:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (643, 'checkUserOnline', '2025-08-14 14:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (644, 'checkUserOnline', '2025-08-14 14:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (645, 'checkUserOnline', '2025-08-14 15:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (646, 'checkUserOnline', '2025-08-14 15:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (647, 'checkUserOnline', '2025-08-14 15:30:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (648, 'checkUserOnline', '2025-08-14 15:40:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (649, 'checkUserOnline', '2025-08-14 15:50:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (650, 'checkUserOnline', '2025-08-14 16:00:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (651, 'checkUserOnline', '2025-08-14 16:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (652, 'checkUserOnline', '2025-08-14 16:20:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (653, 'checkUserOnline', '2025-08-14 16:30:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (654, 'checkUserOnline', '2025-08-14 16:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (655, 'checkUserOnline', '2025-08-14 17:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (656, 'checkUserOnline', '2025-08-14 17:40:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (657, 'checkUserOnline', '2025-08-14 17:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (658, 'checkUserOnline', '2025-08-14 18:00:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (659, 'checkUserOnline', '2025-08-14 18:10:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (660, 'checkUserOnline', '2025-08-14 18:20:05', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (661, 'checkUserOnline', '2025-08-15 10:50:06', '在线用户定时更新，执行成功');
INSERT INTO `sys_job_log` VALUES (662, 'checkUserOnline', '2025-08-15 11:00:06', '在线用户定时更新，执行成功');

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
