-- 数据库 resource
-- 统一文件表
CREATE TABLE `file_info` (
     `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
     `name` VARCHAR(255) NOT NULL COMMENT '文件名字',
     `url` VARCHAR(255) NOT NULL COMMENT '七牛云URL',
     `uploader_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者ID（根据uploader_type区分是用户ID还是管理员ID）',
     `uploader_type` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传者类型：1-H5用户，2-管理员',
     `file_type` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件类型：1-图片，2-视频，3-其他',
     `created_at` datetime(0) DEFAULT NULL,
     `deleted_at` datetime(0) DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件存储表';