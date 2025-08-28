@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

echo 🚀 开始部署商城后台系统...

REM 检查Docker是否安装
docker --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker未安装，请先安装Docker Desktop
    pause
    exit /b 1
)

docker-compose --version >nul 2>&1
if errorlevel 1 (
    echo ❌ Docker Compose未安装，请先安装Docker Compose
    pause
    exit /b 1
)

echo ✅ Docker环境检查通过

REM 停止并删除现有容器
echo 🔄 停止现有容器...
docker-compose down --remove-orphans

REM 清理旧镜像（可选）
set /p "cleanup=是否清理旧镜像？(y/N): "
if /i "!cleanup!"=="y" (
    echo 🧹 清理旧镜像...
    docker system prune -f
)

REM 构建新镜像
echo 🔨 构建应用镜像...
docker-compose build --no-cache

if errorlevel 1 (
    echo ❌ 构建失败，请检查错误信息
    pause
    exit /b 1
)

REM 启动服务
echo 🚀 启动服务...
docker-compose up -d

if errorlevel 1 (
    echo ❌ 启动失败，请检查错误信息
    pause
    exit /b 1
)

REM 等待服务启动
echo ⏳ 等待服务启动...
timeout /t 30 /nobreak >nul

REM 检查服务状态
echo 📊 检查服务状态...
docker-compose ps

echo.
echo ✅ 部署完成！
echo 🌐 应用地址: http://localhost:80
echo 🔧 管理后台: http://localhost:80/admin
echo 🗄️  数据库端口: 3306
echo.
echo 📋 常用命令:
echo   查看日志: docker-compose logs -f
echo   停止服务: docker-compose down
echo   重启服务: docker-compose restart
echo   查看状态: docker-compose ps
echo.
pause
