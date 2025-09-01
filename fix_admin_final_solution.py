#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# 读取admin.html文件
try:
    with open('admin.html', 'r', encoding='utf-8') as f:
        content = f.read()
    
    # 查找并替换所有的登录跳转
    # 匹配单引号的情况
    content = content.replace("window.location.href = '/login';", "window.location.href = '/login.html?logout=true';")
    # 匹配双引号的情况
    content = content.replace('window.location.href = "/login";', 'window.location.href = "/login.html?logout=true";')
    # 匹配没有分号的情况
    content = content.replace("window.location.href = '/login'", "window.location.href = '/login.html?logout=true'")
    
    # 写回文件
    with open('admin.html', 'w', encoding='utf-8') as f:
        f.write(content)
    
    print("修改完成！")
    
    # 验证修改
    import re
    matches = re.findall(r'window\.location\.href\s*=\s*[\'"].*?login.*?[\'"];?', content)
    print(f"\n找到 {len(matches)} 处登录跳转：")
    for match in matches[:5]:  # 只显示前5个
        print(f"  {match}")
        
except Exception as e:
    print(f"错误: {e}")
    print("尝试使用其他编码...")
    
    # 尝试其他编码
    for encoding in ['utf-8', 'gbk', 'latin-1']:
        try:
            with open('admin.html', 'r', encoding=encoding) as f:
                content = f.read()
            print(f"使用 {encoding} 编码成功读取文件")
            
            # 执行替换
            content = content.replace("window.location.href = '/login';", "window.location.href = '/login.html?logout=true';")
            
            with open('admin.html', 'w', encoding=encoding) as f:
                f.write(content)
            
            print("修改完成！")
            break
        except:
            continue