#!/bin/bash

# 商城后台系统 Docker 部署脚本
# 作者: AI Assistant
# 日期: $(date +%Y-%m-%d)

set -e

echo "🚀 开始部署商城后台系统..."

# 检查Docker是否安装
if ! command -v docker &> /dev/null; then
    echo "❌ Docker未安装，请先安装Docker"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose未安装，请先安装Docker Compose"
    exit 1
fi

# 停止并删除现有容器
echo "🔄 停止现有容器..."
docker-compose down --remove-orphans

# 清理旧镜像（可选）
read -p "是否清理旧镜像？(y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🧹 清理旧镜像..."
    docker system prune -f
fi

# 构建新镜像
echo "🔨 构建应用镜像..."
docker-compose build --no-cache

# 启动服务
echo "🚀 启动服务..."
docker-compose up -d

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 30

# 检查服务状态
echo "📊 检查服务状态..."
docker-compose ps

# 检查健康状态
echo "🏥 检查服务健康状态..."
docker-compose exec -T app wget --no-verbose --tries=1 --spider http://localhost:80/ || echo "应用服务未就绪"
docker-compose exec -T mysql mysqladmin ping -h localhost -u root -proot123456 || echo "数据库服务未就绪"

echo "✅ 部署完成！"
echo "🌐 应用地址: http://localhost:80"
echo "🔧 管理后台: http://localhost:80/admin"
echo "🗄️  数据库端口: 3306"
echo ""
echo "📋 常用命令:"
echo "  查看日志: docker-compose logs -f"
echo "  停止服务: docker-compose down"
echo "  重启服务: docker-compose restart"
echo "  查看状态: docker-compose ps"
