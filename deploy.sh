#!/bin/bash

# AT-MMBS Docker部署脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 打印函数
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查Docker和Docker Compose
check_requirements() {
    print_info "检查系统要求..."
    
    if ! command -v docker &> /dev/null; then
        print_error "Docker未安装，请先安装Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose未安装，请先安装Docker Compose"
        exit 1
    fi
    
    print_info "系统要求检查通过"
}

# 创建必要的目录
create_directories() {
    print_info "创建必要的目录..."
    mkdir -p uploads logs ssl
    chmod 755 uploads logs
    print_info "目录创建完成"
}

# 生成环境配置文件
setup_env() {
    if [ ! -f .env ]; then
        print_info "创建环境配置文件..."
        cp .env.example .env
        
        # 生成随机JWT密钥
        JWT_SECRET=$(openssl rand -base64 32)
        sed -i "s/your-secret-key-change-this-in-production/$JWT_SECRET/g" .env
        
        print_warning "已创建.env文件，请根据实际情况修改配置"
    else
        print_info "使用现有的.env文件"
    fi
}

# 构建和启动服务
deploy() {
    print_info "开始部署..."
    
    # 拉取最新镜像
    docker-compose pull
    
    # 构建应用镜像
    print_info "构建应用镜像..."
    docker-compose build
    
    # 启动服务
    print_info "启动服务..."
    docker-compose up -d
    
    # 等待服务就绪
    print_info "等待服务就绪..."
    sleep 10
    
    # 检查服务状态
    docker-compose ps
    
    print_info "部署完成！"
}

# 显示访问信息
show_info() {
    echo ""
    print_info "=== 访问信息 ==="
    echo "主页: http://localhost:8080"
    echo "API: http://localhost:8080/api/"
    echo "管理后台: http://localhost:8080/admin"
    echo ""
    echo "默认管理员账号："
    echo "用户名: admin"
    echo "密码: admin123"
    echo ""
    print_warning "请及时修改默认管理员密码！"
}

# 主函数
main() {
    print_info "AT-MMBS Docker部署脚本"
    echo ""
    
    check_requirements
    create_directories
    setup_env
    deploy
    show_info
    
    echo ""
    print_info "使用以下命令管理服务："
    echo "查看日志: docker-compose logs -f"
    echo "停止服务: docker-compose down"
    echo "重启服务: docker-compose restart"
    echo "查看状态: docker-compose ps"
}

# 执行主函数
main