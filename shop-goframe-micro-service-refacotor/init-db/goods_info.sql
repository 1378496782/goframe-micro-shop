/*
 Navicat MySQL Dump SQL

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : goods

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 22/09/2025 16:56:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for goods_info
-- ----------------------------
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

-- ----------------------------
-- Records of goods_info
-- ----------------------------
INSERT INTO `goods_info` VALUES (1, '微服务电商', 'upload_file/2025-09-08/dcn3r4rs6v9caaf474.jpg', '[{\"url\": \"upload_file/2025-09-08/dcn3r4rs6v9caaf474.jpg\", \"name\": \"无敌浩浩+王铭浩+好好少年.jpg\", \"size\": 397117, \"fileType\": \"image/jpeg\"}]', 111, 1, 1, 2, '', 0, 0, '', 1, '<p>2131313</p>', 0, '2025-09-08 11:35:31', '2025-09-18 14:47:13', NULL);
INSERT INTO `goods_info` VALUES (2, '闪送小程序', '111', '[{\"url\": \"upload_file/2025-09-22/dcz6t7d884ncyuqijj.png\", \"name\": \"8e701935f10a693f44d2183e71bead06.png\", \"size\": 536013, \"fileType\": \"image/png\"}]', 1, 1, 1, 1, '', 99, 9, '课程 AI', 1, '1', 1, NULL, '2025-09-22 16:50:53', NULL);
INSERT INTO `goods_info` VALUES (3, '训练营', 'upload_file/2025-09-22/dcz6t7d884ncyuqijj.png', '[{\"url\": \"upload_file/2025-09-22/dcz6t91t93ok6dgjly.jpg\", \"name\": \"2e90605d988f2d5a98bc1bf5c913a920.jpg\", \"size\": 99467, \"fileType\": \"image/jpeg\"}, {\"url\": \"upload_file/2025-09-22/dcz6tbbqz110pzkwrt.jpg\", \"name\": \"pexels-mikhail-nilov-7988210.jpg\", \"size\": 2124987, \"fileType\": \"image/jpeg\"}]', 99999, 0, 0, 0, '', 99, 99, '课程', 1, '11', 1, '2025-09-22 16:32:46', '2025-09-22 16:32:46', NULL);

SET FOREIGN_KEY_CHECKS = 1;
