-- 创建数据库
CREATE DATABASE IF NOT EXISTS at_mmbs CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE at_mmbs;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 创建商品表
CREATE TABLE IF NOT EXISTS products (
    _id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    kucun INT NOT NULL DEFAULT 0,
    tp JSON,
    sp VARCHAR(500),
    overlay_style TEXT,
    content TEXT,
    status INT DEFAULT 1 COMMENT '1:上架 0:下架',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建分类表
CREATE TABLE IF NOT EXISTS categories (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    sub_category JSON,
    status INT DEFAULT 1 COMMENT '1:启用 0:禁用',
    is_parent BOOLEAN DEFAULT FALSE,
    parent_id VARCHAR(36),
    sort INT DEFAULT 0
);

-- 创建会员表
CREATE TABLE IF NOT EXISTS members (
    _id VARCHAR(36) PRIMARY KEY,
    nick_name VARCHAR(50),
    openid VARCHAR(100) UNIQUE,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    birthday DATE,
    last_login_time TIMESTAMP,
    gender VARCHAR(10),
    phone VARCHAR(20),
    phone_verified BOOLEAN DEFAULT FALSE,
    phone_verify_time TIMESTAMP,
    region JSON,
    addresses JSON,
    referrer_id VARCHAR(36),
    direct_fans JSON,
    orders JSON,
    level_id VARCHAR(36),
    distributor_id VARCHAR(36),
    points INT DEFAULT 0,
    commission DECIMAL(10,2) DEFAULT 0,
    commission_logs JSON,
    points_logs JSON,
    auto_upgrade BOOLEAN DEFAULT TRUE
);

-- 创建会员等级表
CREATE TABLE IF NOT EXISTS member_levels (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    discount DECIMAL(3,2) DEFAULT 1.00 COMMENT '购物折扣',
    points_rate DECIMAL(3,2) DEFAULT 1.00 COMMENT '积分倍率',
    upgrade_conditions JSON COMMENT '升级条件列表',
    auto_upgrade BOOLEAN DEFAULT FALSE,
    sort INT DEFAULT 0
);

-- 创建分销等级表
CREATE TABLE IF NOT EXISTS distributor_levels (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    pickup_discount DECIMAL(3,2) DEFAULT 1.00 COMMENT '提货折扣',
    direct_rate DECIMAL(5,4) DEFAULT 0.0000 COMMENT '直接佣金比例',
    indirect_rate DECIMAL(5,4) DEFAULT 0.0000 COMMENT '间接佣金比例',
    upgrade_conditions JSON COMMENT '升级条件列表',
    auto_upgrade BOOLEAN DEFAULT FALSE,
    distributor_mode VARCHAR(20) DEFAULT 'normal' COMMENT '分销模式',
    sort INT DEFAULT 0
);

-- 创建订单表
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    order_no VARCHAR(50) NOT NULL UNIQUE,
    member_id VARCHAR(36) NOT NULL,
    products JSON NOT NULL COMMENT '订购商品列表',
    delivery_type VARCHAR(20) NOT NULL COMMENT '配送方式',
    address TEXT COMMENT '配送/门店地址',
    total_amount DECIMAL(10,2) NOT NULL COMMENT '订单金额',
    paid_amount DECIMAL(10,2) NOT NULL COMMENT '实付金额',
    payment_status INT DEFAULT 0 COMMENT '0:未支付 1:已支付 2:已退款',
    order_status INT DEFAULT 0 COMMENT '0:待确认 1:已确认 2:已发货 3:已完成 4:已取消',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 创建支付表
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL UNIQUE,
    order_no VARCHAR(50) NOT NULL,
    member_id VARCHAR(36) NOT NULL,
    payment_method VARCHAR(20) NOT NULL COMMENT '支付方式',
    amount DECIMAL(10,2) NOT NULL,
    status INT DEFAULT 0 COMMENT '0:待支付 1:支付中 2:支付成功 3:支付失败 4:已退款',
    transaction_id VARCHAR(100) COMMENT '第三方支付交易号',
    pay_time TIMESTAMP,
    refund_time TIMESTAMP,
    refund_amount DECIMAL(10,2) DEFAULT 0,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 插入一些测试数据
INSERT INTO users (id, username, password) VALUES 
('1', 'admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVEFDi');

INSERT INTO categories (id, name, is_parent, sort) VALUES 
('1', '电子产品', TRUE, 1),
('2', '服装鞋帽', TRUE, 2),
('3', '食品饮料', TRUE, 3);

INSERT INTO member_levels (id, name, description, discount, points_rate) VALUES 
('1', '普通会员', '新注册会员', 1.00, 1.00),
('2', '银卡会员', '消费满1000元', 0.95, 1.20),
('3', '金卡会员', '消费满5000元', 0.90, 1.50);

INSERT INTO distributor_levels (id, name, description, direct_rate, indirect_rate) VALUES 
('1', '普通分销商', '新加入分销商', 0.05, 0.02),
('2', '高级分销商', '销售额满10000元', 0.08, 0.03),
('3', 'VIP分销商', '销售额满50000元', 0.10, 0.04);
