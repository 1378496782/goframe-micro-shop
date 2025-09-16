-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_salt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐 生成密码用',
  `open_id` varchar(64) NOT NULL DEFAULT '' COMMENT '微信openid',
  `phone` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `sex` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1男 2女',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1正常 2拉黑冻结',
  `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '个性签名',
  `secret_answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密保问题的答案',
  `created_at` datetime(0) DEFAULT NULL,
  `updated_at` datetime(0) DEFAULT NULL,
  `deleted_at` datetime(0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES (1, 'lida', 'https://img1.baidu.com/it/u=2029513305,2137933177&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=472', '26bebfe4cf87cc2bd7b89c237fe42df3', 'QLAFRsKG2N','','' ,1, 1, '个性签名', '银河中学', '2022-07-28 17:19:42', '2022-07-31 19:25:01', NULL);
INSERT INTO `user_info` VALUES (2, 'wang', '', '', '','','', 1, 1, '', '', NULL, NULL, NULL);
INSERT INTO `user_info` VALUES (3, '作证表之有', 'http://dummyimage.com/100x100', '34929b5d84cf66ae73797d5b48297710', '6bZBjqX1Nk','','', 1, 92, 'incididunt Excepteur aliqua non', 'cupidatat magna', '2023-01-18 08:19:38', '2023-01-18 08:19:38', NULL);
INSERT INTO `user_info` VALUES (4, 'wzy', 'http://dummyimage.com/100x100', 'a90f083adf941cd75bde9cc371fbe00b', 'wpm0bGNBrf', '','',63, 52, 'Excepteur', 'voluptate in non ea ut', '2023-01-18 09:54:46', '2023-01-18 12:14:13', NULL);
INSERT INTO `user_info` VALUES (5, 'wangzhongyang', 'http://dummyimage.com/100x100', '82131d93ab13a1a4f9ec840a9ddbabf7', 'T0iKtv31BU','','', 1, 1, '和我一起学编程吧', '六个1', '2024-12-26 11:25:43', '2024-12-26 11:25:43', NULL);


-- ----------------------------
-- Table structure for consignee_info
-- ----------------------------
DROP TABLE IF EXISTS `consignee_info`;
CREATE TABLE `consignee_info`  (
       `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '收货地址表',
       `user_id` int(0) NOT NULL DEFAULT 0,
       `is_default` tinyint(1) NOT NULL DEFAULT 0 COMMENT '默认地址1  非默认0\n',
       `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
       `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
       `province` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
       `city` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
       `town` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '县区',
       `street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '街道乡镇',
       `detail` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址详情',
       `created_at` datetime(0) DEFAULT NULL,
       `updated_at` datetime(0) DEFAULT NULL,
       `deleted_at` datetime(0) DEFAULT NULL,
       PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of consignee_info
-- ----------------------------
INSERT INTO `consignee_info` VALUES (1, 1, 1, '王先生1', '13269477632', '北京', '北京市', '房山区', '拱辰街道', '大学城西', '2022-07-31 14:42:33', '2022-07-31 14:44:50', NULL);
