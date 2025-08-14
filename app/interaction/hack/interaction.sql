--  创建数据库   interaction

-- ----------------------------
-- Table structure for comment_info
-- ----------------------------
DROP TABLE IF EXISTS `comment_info`;
CREATE TABLE `comment_info`  (
     `id` int(0) NOT NULL AUTO_INCREMENT,
     `parent_id` int(0) NOT NULL DEFAULT 0 COMMENT '父级评论id',
     `user_id` int(0) NOT NULL DEFAULT 0,
     `object_id` int(0) NOT NULL DEFAULT 0,
     `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '评论类型：1商品 2文章',
     `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '评论内容',
     `created_at` datetime(0) DEFAULT NULL,
     `updated_at` datetime(0) DEFAULT NULL,
     `deleted_at` datetime(0) DEFAULT NULL,
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
    `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '点赞表',
    `user_id` int(0) NOT NULL,
    `type` tinyint(1) NOT NULL COMMENT '点赞类型 1商品 2文章',
    `object_id` int(0) NOT NULL DEFAULT 0 COMMENT '点赞对象id 方便后期扩展',
    `created_at` datetime(0) DEFAULT NULL,
    `updated_at` datetime(0) DEFAULT NULL,
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
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `user_id` int(0) NOT NULL DEFAULT 0 COMMENT '用户id',
    `object_id` int(0) NOT NULL DEFAULT 0 COMMENT '对象id',
    `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '收藏类型：1商品 2文章',
    `created_at` datetime(0) DEFAULT NULL,
    `updated_at` datetime(0) DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `unique_index`(`user_id`, `object_id`, `type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of collection_info
-- ----------------------------
INSERT INTO `collection_info` VALUES (3, 1, 1, 1, '2022-07-31 15:21:38', '2022-07-31 15:21:38');
INSERT INTO `collection_info` VALUES (4, 4, 4, 1, '2023-01-18 15:23:28', '2023-01-18 15:23:28');
