-- 数据库修复脚本
-- 修复商品表中无效的日期值

USE at_mmbs;

-- 修复create_time为0000-00-00的记录
UPDATE products 
SET create_time = CURRENT_TIMESTAMP 
WHERE create_time = '0000-00-00' OR create_time IS NULL;

-- 修复update_time为0000-00-00的记录
UPDATE products 
SET update_time = CURRENT_TIMESTAMP 
WHERE update_time = '0000-00-00' OR update_time IS NULL;

-- 确保所有商品都有有效的时间值
UPDATE products 
SET 
    create_time = COALESCE(create_time, CURRENT_TIMESTAMP),
    update_time = COALESCE(update_time, CURRENT_TIMESTAMP)
WHERE create_time IS NULL OR update_time IS NULL;

-- 显示修复结果
SELECT 
    _id,
    name,
    create_time,
    update_time
FROM products 
ORDER BY create_time DESC;
