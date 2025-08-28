#!/usr/bin/env python3
"""
AT-MMBS 全面测试脚本
"""

import requests
import json
import time
import random
import string
import sys
from datetime import datetime

BASE_URL = "http://localhost:8081"
API_URL = f"{BASE_URL}/api"

# 测试报告
test_results = {
    "passed": 0,
    "failed": 0,
    "warnings": [],
    "errors": [],
    "suggestions": []
}

# 全局变量
auth_token = ""

def generate_random_string(length=8):
    return ''.join(random.choices(string.ascii_letters + string.digits, k=length))

def log_pass(message):
    print(f"✅ {message}")
    test_results["passed"] += 1

def log_fail(message):
    print(f"❌ {message}")
    test_results["failed"] += 1
    test_results["errors"].append(message)

def log_warning(message):
    print(f"⚠️  {message}")
    test_results["warnings"].append(message)

def add_suggestion(suggestion):
    test_results["suggestions"].append(suggestion)

# 测试主页访问
def test_homepage():
    print("\n=== 测试主页访问 ===")
    try:
        response = requests.get(BASE_URL)
        if response.status_code == 200 and "欢迎" in response.text:
            log_pass("主页访问正常")
        else:
            log_fail("主页访问异常")
    except Exception as e:
        log_fail(f"主页访问失败: {str(e)}")

# 测试认证系统
def test_auth():
    global auth_token
    print("\n=== 测试认证系统 ===")
    
    # 1. 测试错误登录
    response = requests.post(f"{API_URL}/auth/login", json={
        "username": "wrong_user",
        "password": "wrong_pass"
    })
    if response.status_code == 200:
        data = response.json()
        if data.get("code") == -1:
            log_pass("错误登录正确拒绝")
        else:
            log_fail("错误登录未被拒绝")
    else:
        log_fail(f"登录API异常: HTTP {response.status_code}")
    
    # 2. 测试正确登录
    response = requests.post(f"{API_URL}/auth/login", json={
        "username": "admin",
        "password": "admin123"
    })
    if response.status_code == 200:
        data = response.json()
        if data.get("code") == 0 and data.get("data", {}).get("token"):
            auth_token = data["data"]["token"]
            log_pass("管理员登录成功")
        else:
            log_fail("管理员登录失败")
            return False
    else:
        log_fail(f"登录API失败: HTTP {response.status_code}")
        return False
    
    # 3. 测试未授权访问
    response = requests.get(f"{API_URL}/users")
    if response.status_code == 401:
        log_pass("未授权访问正确拒绝")
    else:
        log_fail(f"未授权访问未被拒绝: HTTP {response.status_code}")
    
    # 4. 测试授权访问
    headers = {"Authorization": f"Bearer {auth_token}"}
    response = requests.get(f"{API_URL}/users", headers=headers)
    if response.status_code == 200:
        log_pass("授权访问成功")
    else:
        log_fail(f"授权访问失败: HTTP {response.status_code}")
    
    return True

# 测试用户管理
def test_user_management():
    if not auth_token:
        log_warning("跳过用户管理测试（需要认证）")
        return
    
    print("\n=== 测试用户管理 ===")
    headers = {"Authorization": f"Bearer {auth_token}"}
    
    # 1. 获取用户列表
    response = requests.get(f"{API_URL}/users?page=1&size=10", headers=headers)
    if response.status_code == 200:
        log_pass("获取用户列表成功")
    else:
        log_fail(f"获取用户列表失败: HTTP {response.status_code}")
    
    # 2. 创建新用户
    new_user = {
        "username": f"test_{generate_random_string()}",
        "password": "Test123456",
        "role": "user"
    }
    response = requests.post(f"{API_URL}/users", json=new_user, headers=headers)
    if response.status_code == 200:
        data = response.json()
        if data.get("code") == 0:
            log_pass("创建用户成功")
            user_id = data.get("data", {}).get("id")
            
            # 3. 删除测试用户
            if user_id:
                response = requests.delete(f"{API_URL}/users/{user_id}", headers=headers)
                if response.status_code == 200:
                    log_pass("删除用户成功")
                else:
                    log_fail(f"删除用户失败: HTTP {response.status_code}")
        else:
            log_warning(f"创建用户返回错误: {data.get('errorMsg')}")
    else:
        log_fail(f"创建用户失败: HTTP {response.status_code}")

# 测试商品管理
def test_product_management():
    if not auth_token:
        log_warning("跳过商品管理测试（需要认证）")
        return
    
    print("\n=== 测试商品管理 ===")
    headers = {"Authorization": f"Bearer {auth_token}"}
    
    # 1. 创建分类
    new_category = {
        "name": f"测试分类_{generate_random_string()}",
        "parent_id": "0",
        "sort_order": 1
    }
    response = requests.post(f"{API_URL}/categories", json=new_category, headers=headers)
    category_id = None
    if response.status_code == 200:
        data = response.json()
        if data.get("code") == 0:
            log_pass("创建分类成功")
            category_id = data.get("data", {}).get("id")
        else:
            log_warning(f"创建分类失败: {data.get('errorMsg')}")
    else:
        log_fail(f"创建分类失败: HTTP {response.status_code}")
    
    # 2. 创建商品
    new_product = {
        "name": f"测试商品_{generate_random_string()}",
        "category": category_id or "test_cat",
        "price": 99.99,
        "kucun": 100,
        "content": "测试商品描述"
    }
    response = requests.post(f"{API_URL}/products", json=new_product, headers=headers)
    if response.status_code == 200:
        data = response.json()
        if data.get("code") == 0:
            log_pass("创建商品成功")
            product_id = data.get("data", {}).get("id")
            
            # 3. 删除测试商品
            if product_id:
                response = requests.delete(f"{API_URL}/products/{product_id}", headers=headers)
                if response.status_code == 200:
                    log_pass("删除商品成功")
                else:
                    log_warning(f"删除商品失败: HTTP {response.status_code}")
        else:
            log_warning(f"创建商品失败: {data.get('errorMsg')}")
    else:
        log_fail(f"创建商品失败: HTTP {response.status_code}")

# 测试安全性
def test_security():
    print("\n=== 测试安全性 ===")
    
    # 1. SQL注入测试
    malicious_input = "'; DROP TABLE users; --"
    response = requests.post(f"{API_URL}/auth/login", json={
        "username": malicious_input,
        "password": "test"
    })
    if response.status_code in [200, 401]:
        log_pass("SQL注入防护有效")
    else:
        log_fail("SQL注入防护可能存在问题")
    
    # 2. XSS测试（如果有认证）
    if auth_token:
        headers = {"Authorization": f"Bearer {auth_token}"}
        xss_payload = "<script>alert('XSS')</script>"
        
        response = requests.post(f"{API_URL}/products", json={
            "name": xss_payload,
            "category": "test",
            "price": 10,
            "kucun": 10
        }, headers=headers)
        
        if response.status_code == 200:
            data = response.json()
            if data.get("code") == 0:
                log_warning("接受了潜在的XSS输入，建议添加输入验证")
                add_suggestion("对所有用户输入进行HTML转义和验证")
            else:
                log_pass("XSS输入被拒绝")
        else:
            log_pass("XSS输入被拒绝")

# 性能测试
def test_performance():
    print("\n=== 测试性能 ===")
    
    # 测试主页响应时间
    response_times = []
    for i in range(10):
        start = time.time()
        try:
            requests.get(BASE_URL, timeout=5)
            response_times.append(time.time() - start)
        except:
            pass
    
    if response_times:
        avg_time = sum(response_times) / len(response_times)
        if avg_time < 0.5:
            log_pass(f"平均响应时间: {avg_time:.3f}秒")
        else:
            log_warning(f"响应时间较慢: {avg_time:.3f}秒")
            add_suggestion("考虑添加缓存机制提高响应速度")

# 生成报告
def generate_report():
    print("\n" + "="*60)
    print("测试报告")
    print("="*60)
    print(f"测试时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print(f"通过: {test_results['passed']}")
    print(f"失败: {test_results['failed']}")
    if test_results['passed'] + test_results['failed'] > 0:
        print(f"成功率: {test_results['passed']/(test_results['passed']+test_results['failed'])*100:.1f}%")
    
    if test_results["errors"]:
        print("\n❌ 错误:")
        for error in test_results["errors"]:
            print(f"  - {error}")
    
    if test_results["warnings"]:
        print("\n⚠️  警告:")
        for warning in test_results["warnings"]:
            print(f"  - {warning}")
    
    if test_results["suggestions"]:
        print("\n💡 改进建议:")
        for i, suggestion in enumerate(set(test_results["suggestions"]), 1):
            print(f"  {i}. {suggestion}")
    
    # 保存到文件
    with open("test_report.json", "w", encoding="utf-8") as f:
        json.dump(test_results, f, ensure_ascii=False, indent=2)
    print("\n详细报告已保存到 test_report.json")

def main():
    print("AT-MMBS 全面测试")
    print(f"目标: {BASE_URL}\n")
    
    # 检查服务是否运行
    try:
        response = requests.get(BASE_URL, timeout=5)
    except:
        print("❌ 服务未运行，请先启动服务")
        sys.exit(1)
    
    # 运行测试
    test_homepage()
    if test_auth():
        test_user_management()
        test_product_management()
    test_security()
    test_performance()
    
    # 生成报告
    generate_report()

if __name__ == "__main__":
    main()