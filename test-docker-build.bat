@echo off
echo 测试Docker镜像构建...
echo.

REM 只构建应用镜像，不启动整个docker-compose
docker build -t at-mmbs-test .

if errorlevel 1 (
    echo.
    echo ❌ 构建失败
    echo 请检查上面的错误信息
) else (
    echo.
    echo ✅ 构建成功！
    echo 镜像名称: at-mmbs-test
    echo.
    echo 您现在可以运行 deploy.bat 来完整部署应用
)

pause