#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import chardet

# 先检测文件编码
with open('admin.html', 'rb') as f:
    raw_data = f.read()
    result = chardet.detect(raw_data)
    encoding = result['encoding']
    print(f"检测到文件编码: {encoding}")

# 使用检测到的编码读取文件
try:
    with open('admin.html', 'r', encoding=encoding) as f:
        content = f.read()
    
    # 替换所有登录跳转链接
    old_text = "window.location.href = '/login';"
    new_text = "window.location.href = '/login.html?logout=true';"
    
    count = content.count(old_text)
    if count > 0:
        content = content.replace(old_text, new_text)
        
        # 使用相同的编码保存文件
        with open('admin.html', 'w', encoding=encoding) as f:
            f.write(content)
        
        print(f"成功替换了 {count} 处登录跳转链接")
    else:
        print("未找到需要替换的内容")
        
except Exception as e:
    print(f"处理文件时出错: {e}")
    # 尝试使用二进制模式处理
    print("\n尝试使用二进制模式处理...")
    with open('admin.html', 'rb') as f:
        content = f.read()
    
    old_bytes = b"window.location.href = '/login';"
    new_bytes = b"window.location.href = '/login.html?logout=true';"
    
    count = content.count(old_bytes)
    if count > 0:
        content = content.replace(old_bytes, new_bytes)
        with open('admin.html', 'wb') as f:
            f.write(content)
        print(f"使用二进制模式成功替换了 {count} 处登录跳转链接")