#!/bin/bash

# AT-MMBS 部署后测试脚本

set -e

# 配置
BASE_URL="${BASE_URL:-http://localhost:8080}"
API_URL="${BASE_URL}/api"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 测试结果
PASSED=0
FAILED=0

# 打印函数
print_test() {
    echo -e "\n${YELLOW}[TEST]${NC} $1"
}

print_pass() {
    echo -e "${GREEN}[PASS]${NC} $1"
    ((PASSED++))
}

print_fail() {
    echo -e "${RED}[FAIL]${NC} $1"
    ((FAILED++))
}

# 等待服务就绪
wait_for_service() {
    print_test "等待服务就绪..."
    local max_retries=30
    local retry=0
    
    while [ $retry -lt $max_retries ]; do
        if curl -s "${BASE_URL}" > /dev/null; then
            print_pass "服务已就绪"
            return 0
        fi
        sleep 2
        ((retry++))
    done
    
    print_fail "服务启动超时"
    return 1
}

# 测试主页访问
test_homepage() {
    print_test "测试主页访问"
    
    if curl -s "${BASE_URL}/" | grep -q "欢迎"; then
        print_pass "主页访问成功"
    else
        print_fail "主页访问失败"
    fi
}

# 测试API健康检查
test_api_health() {
    print_test "测试API健康检查"
    
    response=$(curl -s -w "\n%{http_code}" "${API_URL}/count")
    http_code=$(echo "$response" | tail -n1)
    
    if [ "$http_code" = "200" ] || [ "$http_code" = "401" ]; then
        print_pass "API响应正常 (HTTP $http_code)"
    else
        print_fail "API响应异常 (HTTP $http_code)"
    fi
}

# 测试登录功能
test_login() {
    print_test "测试登录功能"
    
    response=$(curl -s -X POST "${API_URL}/auth/login" \
        -H "Content-Type: application/json" \
        -d '{"username":"admin","password":"admin123"}')
    
    if echo "$response" | grep -q "token"; then
        print_pass "登录成功"
        # 提取token
        TOKEN=$(echo "$response" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
        export TOKEN
    else
        print_fail "登录失败: $response"
    fi
}

# 测试认证API
test_authenticated_api() {
    print_test "测试需要认证的API"
    
    if [ -z "$TOKEN" ]; then
        print_fail "无法测试：缺少认证token"
        return
    fi
    
    response=$(curl -s -w "\n%{http_code}" \
        -H "Authorization: Bearer $TOKEN" \
        "${API_URL}/users?page=1&size=10")
    http_code=$(echo "$response" | tail -n1)
    
    if [ "$http_code" = "200" ]; then
        print_pass "认证API访问成功"
    else
        print_fail "认证API访问失败 (HTTP $http_code)"
    fi
}

# 测试静态资源
test_static_resources() {
    print_test "测试静态资源访问"
    
    # 测试JS文件
    if curl -s "${BASE_URL}/js/admin.js" | grep -q "function"; then
        print_pass "JS文件访问正常"
    else
        print_fail "JS文件访问失败"
    fi
}

# 测试数据库连接
test_database() {
    print_test "测试数据库连接"
    
    # 通过docker-compose检查MySQL状态
    if docker-compose ps | grep -q "mysql.*Up"; then
        print_pass "MySQL服务运行正常"
    else
        print_fail "MySQL服务异常"
    fi
}

# 测试文件上传目录
test_upload_directory() {
    print_test "测试上传目录"
    
    if [ -d "uploads" ] && [ -w "uploads" ]; then
        print_pass "上传目录可写"
    else
        print_fail "上传目录不可写或不存在"
    fi
}

# 性能测试
test_performance() {
    print_test "简单性能测试"
    
    start_time=$(date +%s.%N)
    curl -s "${BASE_URL}/" > /dev/null
    end_time=$(date +%s.%N)
    
    response_time=$(echo "$end_time - $start_time" | bc)
    
    if (( $(echo "$response_time < 1" | bc -l) )); then
        print_pass "响应时间: ${response_time}秒"
    else
        print_fail "响应时间过长: ${response_time}秒"
    fi
}

# 生成测试报告
generate_report() {
    echo -e "\n========================================"
    echo -e "           测试报告"
    echo -e "========================================"
    echo -e "总测试数: $((PASSED + FAILED))"
    echo -e "${GREEN}通过: $PASSED${NC}"
    echo -e "${RED}失败: $FAILED${NC}"
    echo -e "成功率: $(( PASSED * 100 / (PASSED + FAILED) ))%"
    echo -e "========================================"
    
    if [ $FAILED -eq 0 ]; then
        echo -e "\n${GREEN}✅ 所有测试通过！系统部署成功。${NC}"
        return 0
    else
        echo -e "\n${RED}❌ 部分测试失败，请检查系统配置。${NC}"
        return 1
    fi
}

# 主函数
main() {
    echo "AT-MMBS 部署测试"
    echo "目标服务器: $BASE_URL"
    echo ""
    
    wait_for_service || exit 1
    
    test_homepage
    test_api_health
    test_login
    test_authenticated_api
    test_static_resources
    test_database
    test_upload_directory
    test_performance
    
    generate_report
}

# 执行测试
main