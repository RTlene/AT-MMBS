#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys

def fix_admin_file():
    # 尝试不同的编码
    encodings = ['utf-8', 'latin-1', 'gbk', 'cp1252']
    content = None
    used_encoding = None
    
    # 读取文件
    for encoding in encodings:
        try:
            with open('admin.html', 'r', encoding=encoding) as f:
                content = f.read()
            used_encoding = encoding
            print(f"成功使用 {encoding} 编码读取文件")
            break
        except Exception as e:
            continue
    
    if content is None:
        print("无法读取文件，请检查文件编码")
        return False
    
    # 查找并替换第1536行附近的内容
    old_pattern = """onclick="deleteProduct('${product.id}', '${product.name || product.productName || '未知商品'}')\""""
    new_pattern = """onclick="deleteProduct('${product.id}', '${(product.name || product.productName || '未知商品').replace(/'/g, "\\\\'").replace(/"/g, '\\\\"')}')\""""
    
    if old_pattern in content:
        content = content.replace(old_pattern, new_pattern)
        
        # 写回文件
        with open('admin.html', 'w', encoding=used_encoding) as f:
            f.write(content)
        
        print("修复成功！")
        return True
    else:
        print("未找到需要修复的内容")
        # 尝试打印第1536行附近的内容
        lines = content.split('\n')
        if len(lines) > 1535:
            print(f"\n第1536行的内容：")
            print(lines[1535])
        return False

if __name__ == "__main__":
    fix_admin_file()