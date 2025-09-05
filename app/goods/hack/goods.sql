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
