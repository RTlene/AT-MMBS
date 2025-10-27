#!/bin/bash

# AT-MMBS 简单测试脚本

BASE_URL="http://localhost:8081"
API_URL="${BASE_URL}/api"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 计数器
PASSED=0
FAILED=0

# 打印函数
pass_test() {
    echo -e "${GREEN}✅ $1${NC}"
    ((PASSED++))
}

fail_test() {
    echo -e "${RED}❌ $1${NC}"
    ((FAILED++))
}

warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

echo "=== AT-MMBS 功能测试 ==="
echo "目标: $BASE_URL"
echo ""

# 1. 测试主页
echo "1. 测试主页访问..."
if curl -s "$BASE_URL/" | grep -q "欢迎"; then
    pass_test "主页访问正常"
else
    fail_test "主页访问失败"
fi

# 2. 测试认证系统
echo ""
echo "2. 测试认证系统..."

# 测试错误登录
response=$(curl -s -X POST "$API_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"wrong","password":"wrong"}')
if echo "$response" | grep -q '"code":-1'; then
    pass_test "错误登录正确拒绝"
else
    fail_test "错误登录处理异常"
fi

# 测试正确登录
response=$(curl -s -X POST "$API_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"admin123"}')
if echo "$response" | grep -q '"token"'; then
    TOKEN=$(echo "$response" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    pass_test "管理员登录成功"
else
    fail_test "管理员登录失败"
fi

# 3. 测试API访问
echo ""
echo "3. 测试API访问..."

# 未授权访问
response=$(curl -s -w "\n%{http_code}" "$API_URL/users")
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" = "401" ]; then
    pass_test "未授权访问正确拒绝"
else
    fail_test "未授权访问未被拒绝 (HTTP $http_code)"
fi

# 授权访问
if [ ! -z "$TOKEN" ]; then
    response=$(curl -s -w "\n%{http_code}" \
        -H "Authorization: Bearer $TOKEN" \
        "$API_URL/users?page=1&size=10")
    http_code=$(echo "$response" | tail -n1)
    
    if [ "$http_code" = "200" ]; then
        pass_test "授权访问成功"
    else
        fail_test "授权访问失败 (HTTP $http_code)"
    fi
fi

# 4. 测试CRUD操作
echo ""
echo "4. 测试CRUD操作..."

if [ ! -z "$TOKEN" ]; then
    # 创建测试用户
    response=$(curl -s -X POST "$API_URL/users" \
        -H "Authorization: Bearer $TOKEN" \
        -H "Content-Type: application/json" \
        -d '{"username":"test_user_'$RANDOM'","password":"Test123","role":"user"}')
    
    if echo "$response" | grep -q '"code":0'; then
        pass_test "创建用户成功"
        USER_ID=$(echo "$response" | grep -o '"id":"[^"]*' | cut -d'"' -f4)
        
        # 删除测试用户
        if [ ! -z "$USER_ID" ]; then
            response=$(curl -s -X DELETE "$API_URL/users/$USER_ID" \
                -H "Authorization: Bearer $TOKEN")
            if echo "$response" | grep -q '"code":0'; then
                pass_test "删除用户成功"
            else
                warning "删除用户失败"
            fi
        fi
    else
        warning "创建用户失败"
    fi
fi

# 5. 安全测试
echo ""
echo "5. 安全测试..."

# SQL注入测试
response=$(curl -s -X POST "$API_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"username":"'\'' OR 1=1 --","password":"test"}')
if echo "$response" | grep -q '"code":-1'; then
    pass_test "SQL注入防护有效"
else
    fail_test "SQL注入防护可能存在问题"
fi

# 6. 性能测试
echo ""
echo "6. 简单性能测试..."
start_time=$(date +%s.%N)
curl -s "$BASE_URL/" > /dev/null
end_time=$(date +%s.%N)
response_time=$(echo "$end_time - $start_time" | bc)

if (( $(echo "$response_time < 1" | bc -l) )); then
    pass_test "响应时间: ${response_time}秒"
else
    warning "响应时间较慢: ${response_time}秒"
fi

# 生成报告
echo ""
echo "========================================"
echo "           测试报告"
echo "========================================"
echo "总测试数: $((PASSED + FAILED))"
echo -e "${GREEN}通过: $PASSED${NC}"
echo -e "${RED}失败: $FAILED${NC}"
if [ $((PASSED + FAILED)) -gt 0 ]; then
    success_rate=$(( PASSED * 100 / (PASSED + FAILED) ))
    echo "成功率: ${success_rate}%"
fi
echo "========================================"

if [ $FAILED -eq 0 ]; then
    echo -e "\n${GREEN}✅ 所有测试通过！${NC}"
    exit 0
else
    echo -e "\n${RED}❌ 有测试失败，请检查。${NC}"
    exit 1
fi