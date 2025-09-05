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
      PRIMARY KEY (`id`) USING BTREE
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
      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of position_info
-- ----------------------------
INSERT INTO `position_info` VALUES (2, 'https://images.zsxq.com/FgdL08hVmh-40_e12vh-ifbXpGxB?e=2000966400', '测试', 'https://articles.zsxq.com/id_wd15wsegvow1.html', 0, 1, '2022-11-18 17:44:07', '2022-11-18 17:44:07', '2022-11-18 17:44:59');
