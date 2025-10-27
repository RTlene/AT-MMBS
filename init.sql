-- AT-MMBS 数据库初始化脚本
-- 该脚本会在MySQL容器首次启动时自动执行

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS golang_demo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE golang_demo;

-- 计数器表（测试用）
CREATE TABLE IF NOT EXISTS `counters` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `count` int(11) NOT NULL DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
    `id` varchar(50) NOT NULL,
    `username` varchar(50) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    `role` varchar(20) DEFAULT 'user',
    `status` tinyint(1) DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 分类表
CREATE TABLE IF NOT EXISTS `category` (
    `id` varchar(50) NOT NULL,
    `name` varchar(100) NOT NULL,
    `parent_id` varchar(50) DEFAULT '0',
    `sort_order` int(11) DEFAULT 0,
    `status` tinyint(1) DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_parent` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 商品表
CREATE TABLE IF NOT EXISTS `product` (
    `id` varchar(50) NOT NULL,
    `name` varchar(200) NOT NULL,
    `category` varchar(50),
    `price` decimal(10,2) NOT NULL,
    `kucun` int(11) DEFAULT 0,
    `content` longtext,
    `tp` json,
    `status` tinyint(1) DEFAULT 1,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_category` (`category`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 会员等级表
CREATE TABLE IF NOT EXISTS `member_level` (
    `id` varchar(50) NOT NULL,
    `name` varchar(50) NOT NULL,
    `discount` decimal(3,2) DEFAULT 1.00,
    `min_points` int(11) DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 分销商等级表
CREATE TABLE IF NOT EXISTS `distributor_level` (
    `id` varchar(50) NOT NULL,
    `name` varchar(50) NOT NULL,
    `commission_rate` decimal(5,2) DEFAULT 0.00,
    `min_sales` decimal(10,2) DEFAULT 0.00,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 会员表
CREATE TABLE IF NOT EXISTS `member` (
    `id` varchar(50) NOT NULL,
    `openid` varchar(100) UNIQUE,
    `username` varchar(50) UNIQUE,
    `password` varchar(255),
    `real_name` varchar(50),
    `phone` varchar(20),
    `email` varchar(100),
    `member_level_id` varchar(50),
    `distributor_level_id` varchar(50),
    `referrer_id` varchar(50),
    `points` int(11) DEFAULT 0,
    `balance` decimal(10,2) DEFAULT 0.00,
    `status` tinyint(1) DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_openid` (`openid`),
    INDEX `idx_member_level` (`member_level_id`),
    INDEX `idx_distributor_level` (`distributor_level_id`),
    INDEX `idx_referrer` (`referrer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 订单表
CREATE TABLE IF NOT EXISTS `order` (
    `id` varchar(50) NOT NULL,
    `order_no` varchar(32) NOT NULL UNIQUE,
    `member_id` varchar(50) NOT NULL,
    `total_amount` decimal(10,2) NOT NULL,
    `discount_amount` decimal(10,2) DEFAULT 0.00,
    `pay_amount` decimal(10,2) NOT NULL,
    `status` int(11) DEFAULT 0,
    `remark` text,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_member` (`member_id`),
    INDEX `idx_order_no` (`order_no`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 支付记录表
CREATE TABLE IF NOT EXISTS `payment` (
    `id` varchar(50) NOT NULL,
    `order_id` varchar(50) NOT NULL,
    `payment_no` varchar(64) NOT NULL,
    `payment_method` varchar(20),
    `amount` decimal(10,2) NOT NULL,
    `status` int(11) DEFAULT 0,
    `paid_at` timestamp NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_order` (`order_id`),
    INDEX `idx_payment_no` (`payment_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入默认会员等级
INSERT IGNORE INTO `member_level` (`id`, `name`, `discount`, `min_points`) VALUES 
('ml_001', '普通会员', 1.00, 0),
('ml_002', '银卡会员', 0.95, 1000),
('ml_003', '金卡会员', 0.90, 5000),
('ml_004', '钻石会员', 0.85, 10000);

-- 插入默认分销商等级
INSERT IGNORE INTO `distributor_level` (`id`, `name`, `commission_rate`, `min_sales`) VALUES 
('dl_001', '初级分销商', 5.00, 0.00),
('dl_002', '中级分销商', 8.00, 10000.00),
('dl_003', '高级分销商', 12.00, 50000.00);

-- 插入测试分类
INSERT IGNORE INTO `category` (`id`, `name`, `parent_id`, `sort_order`) VALUES 
('cat_001', '电子产品', '0', 1),
('cat_002', '服装', '0', 2),
('cat_003', '食品', '0', 3),
('cat_004', '手机', 'cat_001', 1),
('cat_005', '电脑', 'cat_001', 2);

-- 注意：默认管理员用户将由应用程序在启动时自动创建