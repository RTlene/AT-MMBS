-- 修复数据库中的无效日期时间值 - 版本2
-- 修复 products 表中的 update_time 字段

USE at_mmbs;

-- 检查是否有无效的 update_time 值
SELECT COUNT(*) as invalid_update_time_count 
FROM products 
WHERE update_time = '0000-00-00 00:00:00' OR update_time IS NULL;

-- 修复 update_time 字段的无效值
UPDATE products 
SET update_time = CURRENT_TIMESTAMP 
WHERE update_time = '0000-00-00 00:00:00' OR update_time IS NULL;

-- 检查是否有无效的 create_time 值
SELECT COUNT(*) as invalid_create_time_count 
FROM products 
WHERE create_time = '0000-00-00 00:00:00' OR create_time IS NULL;

-- 修复 create_time 字段的无效值
UPDATE products 
SET create_time = CURRENT_TIMESTAMP 
WHERE create_time = '0000-00-00 00:00:00' OR create_time IS NULL;

-- 验证修复结果
SELECT 
    COUNT(*) as total_products,
    SUM(CASE WHEN update_time = '0000-00-00 00:00:00' THEN 1 ELSE 0 END) as remaining_invalid_update_time,
    SUM(CASE WHEN create_time = '0000-00-00 00:00:00' THEN 1 ELSE 0 END) as remaining_invalid_create_time
FROM products;

-- 显示修复后的商品记录
SELECT 
    _id,
    name,
    status,
    create_time,
    update_time
FROM products 
ORDER BY create_time DESC 
LIMIT 10;
