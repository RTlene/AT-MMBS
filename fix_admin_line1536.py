#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 读取admin.html文件
with open('admin.html', 'r', encoding='utf-8') as f:
    lines = f.readlines()

# 修改第1536行（注意：Python列表索引从0开始，所以是1535）
if len(lines) > 1535:
    old_line = lines[1535]
    print(f"原始行内容:\n{old_line}")
    
    # 修复：将商品名称中的单引号转义
    new_line = old_line.replace(
        "onclick=\"deleteProduct('${product.id}', '${product.name || product.productName || '未知商品'}')\"",
        "onclick=\"deleteProduct('${product.id}', '${(product.name || product.productName || '未知商品').replace(/'/g, \"\\\\'\")}')\"" 
    )
    
    lines[1535] = new_line
    print(f"\n修改后的行内容:\n{new_line}")
    
    # 写回文件
    with open('admin.html', 'w', encoding='utf-8') as f:
        f.writelines(lines)
    
    print("\n修改完成！")
else:
    print("文件行数不足1536行")