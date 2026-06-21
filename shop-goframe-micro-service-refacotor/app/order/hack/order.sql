-- ----------------------------
-- Table structure for order_goods_info
-- ----------------------------
DROP TABLE IF EXISTS `order_goods_info`;
CREATE TABLE `order_goods_info`  (
     `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '商品维度的订单表',
     `order_id` int(0) NOT NULL DEFAULT 0 COMMENT '关联的主订单表',
     `goods_id` int(0) NOT NULL DEFAULT 0 COMMENT '商品id',
     `goods_options_id` int(0) DEFAULT 0 COMMENT '商品规格id sku id',
     `count` int(0) NOT NULL COMMENT '商品数量',
     `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
     `price` int(0) NOT NULL DEFAULT 0 COMMENT '订单金额 单位分',
     `coupon_price` int(0) NOT NULL DEFAULT 0 COMMENT '优惠券金额 单位分',
     `actual_price` int(0) NOT NULL DEFAULT 0 COMMENT '实际支付金额 单位分',
     `created_at` datetime(0) DEFAULT NULL,
     `updated_at` datetime(0) DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE
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
   `id` int(0) NOT NULL AUTO_INCREMENT,
   `number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
   `user_id` int(0) NOT NULL DEFAULT 0 COMMENT '用户id',
   `pay_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '支付方式 1微信 2支付宝 3云闪付',
   `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
   `pay_at` datetime(0) DEFAULT NULL COMMENT '支付时间',
   `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价 6待确认 7已取消',
   `consignee_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
   `consignee_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人手机号',
   `consignee_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
   `price` int(0) NOT NULL DEFAULT 0 COMMENT '订单金额 单位分',
   `coupon_price` int(0) NOT NULL DEFAULT 0 COMMENT '优惠券金额 单位分',
   `actual_price` int(0) NOT NULL DEFAULT 0 COMMENT '实际支付金额 单位分',
   `sales_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '销量同步状态：0未同步 1已同步 2同步失败 3同步中',
   `created_at` datetime(0) DEFAULT NULL,
   `updated_at` datetime(0) DEFAULT NULL,
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章（种草）表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_info
-- ----------------------------
INSERT INTO `order_info` VALUES (1, '1659231316407832000111', 1, 1, '0', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 100, 9900, 0, '2022-08-27 09:35:16', '2022-07-31 09:35:16');
INSERT INTO `order_info` VALUES (2, '1659231554317361000757', 1, 1, '0', NULL, 1, '王先生', '13269477432', '北京丰台汽车博物馆', 10000, 200, 9800, 0, '2022-07-31 09:39:14', '2022-07-31 09:39:14');
INSERT INTO `order_info` VALUES (3, '1661603467832912000516', 1, 0, '', '2022-12-13 21:52:26', 0, '', '', '', 0, 0, 0, 0, '2022-12-08 20:31:07', '2022-08-27 20:31:07');
INSERT INTO `order_info` VALUES (4, '1661603562656619000513', 1, 1, '放到快递柜就可以，不用打电话。', '2022-12-13 21:52:19', 0, '王先生', '13269477432', '北京丰台汽车博物馆', 0, 0, 0, 0, '2022-12-09 20:32:42', '2022-08-27 20:32:42');
INSERT INTO `order_info` VALUES (5, '1661604424031843000546', 1, 0, '', '2022-12-13 21:52:12', 0, '', '', '', 0, 0, 0, 0, '2022-12-10 20:47:04', '2022-08-27 20:47:04');
INSERT INTO `order_info` VALUES (6, '1661604530142913000770', 1, 1, '这是备注', '2022-12-13 21:52:05', 1, '', '', '', 100, 0, 0, 0, '2022-12-11 20:48:50', '2022-08-27 20:48:50');
INSERT INTO `order_info` VALUES (7, '166160461284091500027', 1, 1, '这是备注', '2022-12-13 21:51:58', 1, '', '', '', 100, 0, 9800, 0, '2022-12-12 20:50:50', '2022-08-27 20:50:12');
INSERT INTO `order_info` VALUES (8, '166160465089079000090', 1, 1, '这是备注', '2022-12-19 13:14:07', 1, '', '', '', 100, 0, 9800, 0, '2022-12-19 20:50:50', '2022-08-27 20:50:50');
INSERT INTO `order_info` VALUES (9, '1675756215452071000340', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, 0, '2023-02-07 15:50:15', '2023-02-07 15:50:15');
INSERT INTO `order_info` VALUES (10, '1675756478857950000699', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, 0, '2023-02-07 15:54:38', '2023-02-07 15:54:38');
INSERT INTO `order_info` VALUES (11, '1675756573337742000600', 15, 0, 'ad sint', NULL, 1, '应便更空政于八', '18130435879', '香港特别行政区三沙市望城区', 25, 31, 97, 0, '2023-02-07 15:56:13', '2023-02-07 15:56:13');
INSERT INTO `order_info` VALUES (13, '1675756783046217000276', 15, 0, 'consequat cupidatat in', NULL, 1, '在区三少', '18127326325', '广东省天津市安溪县', 87, 66, 92, 0, '2023-02-07 15:59:43', '2023-02-07 15:59:43');
INSERT INTO `order_info` VALUES (14, '1675923255725252000737', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, 0, '2023-02-09 14:14:15', '2023-02-09 14:14:15');
INSERT INTO `order_info` VALUES (15, '1675923359070046000221', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, 0, '2023-02-09 14:15:59', '2023-02-09 14:15:59');
INSERT INTO `order_info` VALUES (16, '1675925468868358000350', 15, 0, '备注', NULL, 1, '王中阳go', '18130435879', '北京市朝阳区望京SOHO', 10, 1, 9, 0, '2023-02-09 14:51:09', '2023-02-09 14:51:09');


-- ----------------------------
-- Table structure for refund_info
-- ----------------------------
DROP TABLE IF EXISTS `refund_info`;
CREATE TABLE `refund_info`  (
        `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '售后退款表',
        `number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '售后订单号',
        `order_id` int(0) NOT NULL COMMENT '订单id',
        `goods_id` int(0) NOT NULL DEFAULT 0 COMMENT '要售后的商品id\n',
        `reason` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退款原因',
        `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1待处理 2同意退款 3拒绝退款\n',
        `user_id` int(0) NOT NULL COMMENT '用户id',
        `created_at` datetime(0) DEFAULT NULL,
        `updated_at` datetime(0) DEFAULT NULL,
        `deleted_at` datetime(0) DEFAULT NULL,
        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of refund_info
-- ----------------------------
INSERT INTO `refund_info` VALUES (1, 'refund1659247832739250000428', 1, 1, '不想要了', 1, 1, '2022-07-31 14:10:32', '2022-07-31 14:10:32', NULL);

-- ----------------------------
-- Table structure for order_outbox_message
-- ----------------------------
DROP TABLE IF EXISTS `order_outbox_message`;
CREATE TABLE `order_outbox_message`  (
    `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `event_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '事件唯一id，用于幂等去重',
    `event_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '事件类型',
    `aggregate_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '聚合根id，如订单id',
    `exchange` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息投递的 exchange',
    `routing_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息投递的 routing key',
    `payload` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '消息体内容',
    `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '发送状态：0待发送 1发送中 2发送成功 3发送失败',
    `retry_count` int(0) NOT NULL DEFAULT 0 COMMENT '已重试次数',
    `next_retry_at` datetime(0) DEFAULT NULL COMMENT '下次重试时间',
    `last_error` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最近一次失败原因',
    `created_at` datetime(0) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(0) DEFAULT NULL COMMENT '更新时间',
    `sent_at` datetime(0) DEFAULT NULL COMMENT '发送成功时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `uk_event_id`(`event_id`) USING BTREE,
    INDEX `idx_status_next_retry`(`status`, `next_retry_at`) USING BTREE,
    INDEX `idx_aggregate_id`(`aggregate_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '订单事件发件箱表（事务消息）' ROW_FORMAT = Dynamic;
