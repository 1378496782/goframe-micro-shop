-- 数据库 current
-- 统一文件表
CREATE TABLE `file_info` (
     `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
     `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '文件名字',
     `url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '七牛云URL',
     `uploader_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者ID',
     `created_at` datetime(0) DEFAULT NULL,
     `deleted_at` datetime(0) DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件存储表';