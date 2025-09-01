#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 读取admin.html文件
with open('admin.html', 'r', encoding='utf-8') as f:
    lines = f.readlines()

# 修改特定行
modified = False
for i in range(len(lines)):
    if "window.location.href = '/login';" in lines[i]:
        old_line = lines[i]
        lines[i] = lines[i].replace("'/login';", "'/login.html?logout=true';")
        print(f"修改第 {i+1} 行:")
        print(f"  原内容: {old_line.strip()}")
        print(f"  新内容: {lines[i].strip()}")
        modified = True

if modified:
    # 保存修改后的文件
    with open('admin.html', 'w', encoding='utf-8') as f:
        f.writelines(lines)
    print("\nadmin.html 已成功更新！")
else:
    print("未找到需要修改的内容")