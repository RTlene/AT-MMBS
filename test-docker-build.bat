@echo off
echo 测试Docker镜像构建...
echo.
echo 清理之前的构建缓存...
docker system prune -f

echo.
echo 构建应用镜像...
docker build --no-cache -t at-mmbs-test .

if errorlevel 1 (
    echo.
    echo ❌ 构建失败
    echo 请检查上面的错误信息
    echo.
    echo 常见问题：
    echo - Go版本不匹配：已更新到Go 1.24
    echo - main函数重复定义：已解决
    echo - 依赖下载失败：检查网络连接
) else (
    echo.
    echo ✅ 构建成功！
    echo 镜像名称: at-mmbs-test
    echo.
    echo 您现在可以运行 deploy.bat 来完整部署应用
)

pause