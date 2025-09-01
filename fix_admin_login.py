#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 修复admin.html中的登录跳转链接

with open('admin.html', 'r', encoding='utf-8') as f:
    content = f.read()

# 替换所有的登录跳转链接
replacements = [
    ("window.location.href = '/login';", "window.location.href = '/login.html?logout=true';"),
    ("window.location.href = '/login'", "window.location.href = '/login.html?logout=true'"),
]

modified = False
for old, new in replacements:
    if old in content:
        content = content.replace(old, new)
        modified = True
        print(f"已替换: {old} -> {new}")

if modified:
    with open('admin.html', 'w', encoding='utf-8') as f:
        f.write(content)
    print("admin.html 已更新")
else:
    print("未找到需要替换的内容")

# 验证修改
with open('admin.html', 'r', encoding='utf-8') as f:
    lines = f.readlines()
    for i, line in enumerate(lines):
        if '/login' in line:
            print(f"第 {i+1} 行: {line.strip()}")