-- ----------------------------
-- Table structure for admin_info
-- ----------------------------
DROP TABLE IF EXISTS `admin_info`;
CREATE TABLE `admin_info`  (
   `id` int(0) NOT NULL AUTO_INCREMENT,
   `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
   `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
   `role_ids` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色ids',
   `created_at` datetime(0) DEFAULT NULL,
   `updated_at` datetime(0) DEFAULT NULL,
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
      `id` int(0) NOT NULL AUTO_INCREMENT,
      `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
      `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
      `created_at` datetime(0) DEFAULT NULL,
      `updated_at` datetime(0) DEFAULT NULL,
      `deleted_at` datetime(0) DEFAULT NULL,
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
     `id` int(0) NOT NULL AUTO_INCREMENT,
     `role_id` int(0) NOT NULL DEFAULT 0 COMMENT '角色id',
     `permission_id` int(0) NOT NULL COMMENT '权限id',
     `created_at` datetime(0) DEFAULT NULL,
     `updated_at` datetime(0) DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE INDEX `unique_index`(`role_id`, `permission_id`) USING BTREE COMMENT '唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for permission_info
-- ----------------------------
DROP TABLE IF EXISTS `permission_info`;
CREATE TABLE `permission_info`  (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限名称',
    `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路径',
    `created_at` datetime(0) DEFAULT NULL,
    `updated_at` datetime(0) DEFAULT NULL,
    `deleted_at` datetime(0) DEFAULT NULL,
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
