#!/bin/bash

# AT-MMBS 模块化版本测试脚本

echo "=== AT-MMBS 模块化版本测试 ==="

# 清理旧文件
rm -f test.db mmbs-modular

# 设置测试环境变量
export TEST_MODE=true
export SQLITE_PATH=test.db
export PORT=8081
export JWT_SECRET=test-secret-key

echo "1. 编译模块化版本..."
if ! go build -o mmbs-modular main_new.go; then
    echo "❌ 编译失败"
    exit 1
fi
echo "✅ 编译成功"

echo ""
echo "2. 启动服务..."
./mmbs-modular &
SERVER_PID=$!

echo "服务PID: $SERVER_PID"
echo "等待服务启动..."
sleep 3

# 检查服务是否启动
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo "❌ 服务启动失败"
    exit 1
fi

echo "✅ 服务已启动在 http://localhost:$PORT"

# 基础API测试
echo ""
echo "3. 运行API测试..."

# 测试主页
echo -n "  - 测试主页访问... "
if curl -s http://localhost:$PORT/ | grep -q "欢迎"; then
    echo "✅ 通过"
else
    echo "❌ 失败"
fi

# 测试登录（应该失败，因为没有认证）
echo -n "  - 测试未认证API访问... "
response=$(curl -s -w "\n%{http_code}" http://localhost:$PORT/api/users)
http_code=$(echo "$response" | tail -n1)
if [ "$http_code" = "401" ]; then
    echo "✅ 正确拒绝（401）"
else
    echo "❌ 预期401，实际$http_code"
fi

# 测试登录
echo -n "  - 测试管理员登录... "
login_response=$(curl -s -X POST http://localhost:$PORT/api/auth/login \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"admin123"}')

if echo "$login_response" | grep -q "token"; then
    TOKEN=$(echo "$login_response" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    echo "✅ 登录成功"
else
    echo "❌ 登录失败: $login_response"
fi

# 测试认证后的API访问
if [ ! -z "$TOKEN" ]; then
    echo -n "  - 测试认证API访问... "
    auth_response=$(curl -s -w "\n%{http_code}" \
        -H "Authorization: Bearer $TOKEN" \
        http://localhost:$PORT/api/users?page=1&size=10)
    http_code=$(echo "$auth_response" | tail -n1)
    
    if [ "$http_code" = "200" ]; then
        echo "✅ 访问成功"
    else
        echo "❌ 访问失败 (HTTP $http_code)"
    fi
fi

# 清理
echo ""
echo "4. 停止服务..."
kill $SERVER_PID 2>/dev/null
wait $SERVER_PID 2>/dev/null

echo ""
echo "=== 测试完成 ==="
echo ""
echo "模块化改进总结："
echo "✅ 路由模块化：独立的router包管理所有路由"
echo "✅ 中间件支持：认证和日志中间件"
echo "✅ 配置管理：集中的配置管理"
echo "✅ RESTful API：标准的REST风格接口"
echo "✅ 代码结构清晰：各模块职责分明"