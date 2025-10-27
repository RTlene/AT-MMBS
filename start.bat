@echo off
chcp 65001 >nul
echo ========================================
echo    微信小程序商城后台管理系统
echo ========================================
echo.

echo 检查Docker是否运行...
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo 错误: 未检测到Docker，请先安装并启动Docker Desktop
    pause
    exit /b 1
)

echo 检查Docker Compose是否可用...
docker-compose --version >nul 2>&1
if %errorlevel% neq 0 (
    echo 错误: 未检测到Docker Compose，请先安装Docker Compose
    pause
    exit /b 1
)

echo.
echo 请选择操作:
echo 1. 启动服务
echo 2. 停止服务
echo 3. 重新编译并部署
echo 4. 查看服务状态
echo 5. 查看服务日志
echo 6. 退出
echo.
set /p choice=请输入选择 (1-6): 

if "%choice%"=="1" goto start
if "%choice%"=="2" goto stop
if "%choice%"=="3" goto rebuild
if "%choice%"=="4" goto status
if "%choice%"=="5" goto logs
if "%choice%"=="6" goto exit
echo 无效选择，请重新运行脚本
pause
exit /b 1

:start
echo.
echo 正在启动服务...
docker-compose up -d
if %errorlevel% neq 0 (
    echo 错误: 服务启动失败
    pause
    exit /b 1
)

echo.
echo 等待服务启动...
timeout /t 10 /nobreak >nul

echo.
echo 服务启动完成！
echo 访问地址:
echo   首页: http://localhost:8080/
echo   管理后台: http://localhost:8080/admin
echo.
echo 数据库连接信息:
echo   主机: localhost
echo   端口: 3307
echo   数据库: at_mmbs
echo   用户名: at_mmbs_user
echo   密码: at_mmbs_pass
echo.
echo 按任意键查看服务状态...
pause >nul
goto status

:stop
echo.
echo 正在停止服务...
docker-compose down
if %errorlevel% neq 0 (
    echo 错误: 服务停止失败
    pause
    exit /b 1
)
echo 服务已停止
echo.
echo 按任意键继续...
pause >nul
goto menu

:rebuild
echo.
echo 正在停止现有服务...
docker-compose down

echo.
echo 正在清理旧镜像...
docker-compose down --rmi all --volumes --remove-orphans

echo.
echo 正在重新构建镜像...
docker-compose build --no-cache
if %errorlevel% neq 0 (
    echo 错误: 镜像构建失败
    pause
    exit /b 1
)

echo.
echo 正在启动服务...
docker-compose up -d
if %errorlevel% neq 0 (
    echo 错误: 服务启动失败
    pause
    exit /b 1
)

echo.
echo 等待服务启动...
timeout /t 15 /nobreak >nul

echo.
echo 重新编译部署完成！
echo 访问地址:
echo   首页: http://localhost:8080/
echo   管理后台: http://localhost:8080/admin
echo.
echo 按任意键查看服务状态...
pause >nul
goto status

:status
echo.
echo 当前服务状态:
docker-compose ps
echo.
echo 按任意键返回主菜单...
pause >nul
goto menu

:logs
echo.
echo 正在显示服务日志...
echo 按 Ctrl+C 停止查看日志
echo.
docker-compose logs -f
echo.
echo 按任意键返回主菜单...
pause >nul
goto menu

:menu
cls
echo ========================================
echo    微信小程序商城后台管理系统
echo ========================================
echo.
echo 请选择操作:
echo 1. 启动服务
echo 2. 停止服务
echo 3. 重新编译并部署
echo 4. 查看服务状态
echo 5. 查看服务日志
echo 6. 退出
echo.
set /p choice=请输入选择 (1-6): 

if "%choice%"=="1" goto start
if "%choice%"=="2" goto stop
if "%choice%"=="3" goto rebuild
if "%choice%"=="4" goto status
if "%choice%"=="5" goto logs
if "%choice%"=="6" goto exit
echo 无效选择，请重新选择
timeout /t 2 /nobreak >nul
goto menu

:exit
echo.
echo 感谢使用！
pause
exit /b 0
