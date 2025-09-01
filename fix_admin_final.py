#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import re

# 读取admin.html文件
with open('admin.html', 'r', encoding='utf-8') as f:
    content = f.read()

# 使用正则表达式替换所有的登录跳转
# 匹配 window.location.href = '/login' 后面可能有分号
pattern = r"window\.location\.href\s*=\s*['\"]\/login['\"]"
replacement = "window.location.href = '/login.html?logout=true'"

# 执行替换
new_content = re.sub(pattern, replacement, content)

# 计算替换次数
count = len(re.findall(pattern, content))

if count > 0:
    # 保存修改后的文件
    with open('admin.html', 'w', encoding='utf-8') as f:
        f.write(new_content)
    print(f"成功替换了 {count} 处登录跳转链接")
    
    # 显示替换后的结果
    lines = new_content.split('\n')
    for i, line in enumerate(lines):
        if 'login.html?logout=true' in line:
            print(f"第 {i+1} 行: {line.strip()}")
else:
    print("未找到需要替换的内容")
    # 显示包含login的行，用于调试
    lines = content.split('\n')
    for i, line in enumerate(lines):
        if 'window.location.href' in line and 'login' in line:
            print(f"第 {i+1} 行: {repr(line)}")