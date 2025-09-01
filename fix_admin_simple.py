#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 尝试不同的编码
encodings = ['utf-8', 'latin-1', 'iso-8859-1', 'cp1252', 'gbk', 'utf-16']

content = None
used_encoding = None

# 尝试读取文件
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
    # 如果都失败了，使用二进制模式
    print("所有编码都失败了，使用二进制模式")
    with open('admin.html', 'rb') as f:
        binary_content = f.read()
    
    # 二进制替换
    old_bytes = b"window.location.href = '/login';"
    new_bytes = b"window.location.href = '/login.html?logout=true';"
    
    count = binary_content.count(old_bytes)
    if count > 0:
        binary_content = binary_content.replace(old_bytes, new_bytes)
        with open('admin.html', 'wb') as f:
            f.write(binary_content)
        print(f"使用二进制模式成功替换了 {count} 处登录跳转链接")
    else:
        print("二进制模式下也未找到需要替换的内容")
else:
    # 文本模式替换
    old_text = "window.location.href = '/login';"
    new_text = "window.location.href = '/login.html?logout=true';"
    
    count = content.count(old_text)
    if count > 0:
        content = content.replace(old_text, new_text)
        with open('admin.html', 'w', encoding=used_encoding) as f:
            f.write(content)
        print(f"成功替换了 {count} 处登录跳转链接")
    else:
        print(f"未找到需要替换的内容")