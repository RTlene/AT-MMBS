@echo off
echo 停止微信小程序商城后台管理系统...
echo.

echo 停止服务...
docker-compose down

echo.
echo 服务已停止！
echo.
echo 按任意键退出...
pause >nul
