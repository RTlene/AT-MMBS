#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 使用latin-1编码读取文件
with open('admin.html', 'r', encoding='latin-1') as f:
    lines = f.readlines()

# 查找包含location和login的行
print("查找包含 'location' 和 'login' 的行：\n")
found = False
for i, line in enumerate(lines):
    if 'location' in line.lower() and 'login' in line.lower():
        print(f"第 {i+1} 行: {repr(line.strip())}")
        found = True

if not found:
    # 查找包含login的行
    print("\n查找包含 'login' 的行：\n")
    for i, line in enumerate(lines):
        if 'login' in line.lower():
            print(f"第 {i+1} 行: {line.strip()}")
            if i < len(lines) - 5:
                for j in range(i-2, i+3):
                    if j >= 0 and j < len(lines):
                        print(f"  {j+1}: {lines[j].strip()}")