-- 商品主表
DROP TABLE IF EXISTS `goods_info`;
CREATE TABLE `goods_info` (
      `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
      `name` VARCHAR(200) NOT NULL DEFAULT '',
      `price` INT NOT NULL DEFAULT 1 COMMENT '价格(分)',
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
    `goods_id` INT NOT NULL COMMENT '商品ID',
    `file_id` INT NOT NULL COMMENT '文件ID（关联file_info）',
    `sort` TINYINT NOT NULL DEFAULT 0 COMMENT '排序',
    PRIMARY KEY (`id`),
    INDEX `idx_goods` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品详情图';