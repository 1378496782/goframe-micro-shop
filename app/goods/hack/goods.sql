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
  `sort` TINYINT NOT NULL DEFAULT 0 COMMENT '排序 倒序',
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
      PRIMARY KEY (`id`),
      UNIQUE KEY `uk_user_goods` (`user_id`, `goods_id`)
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
