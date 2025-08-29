@echo off
echo AT-MMBS 项目快速清理工具
echo ========================
echo.
echo 此工具将帮助您清理项目中的不必要文件
echo.
echo 请选择清理模式:
echo 1. 预览模式 (只显示将要删除的文件，不真正删除)
echo 2. 完全清理 (删除所有不必要的文件)
echo 3. 保留测试文件
echo 4. 保留日志文件
echo 5. 保留测试和日志文件
echo 6. 退出
echo.
set /p choice="请输入选项 (1-6): "

if "%choice%"=="1" (
    echo.
    echo 运行预览模式...
    powershell -ExecutionPolicy Bypass -File .\cleanup-project-safe.ps1 -DryRun
) else if "%choice%"=="2" (
    echo.
    echo 开始完全清理...
    powershell -ExecutionPolicy Bypass -File .\cleanup-project-safe.ps1
) else if "%choice%"=="3" (
    echo.
    echo 清理但保留测试文件...
    powershell -ExecutionPolicy Bypass -File .\cleanup-project-safe.ps1 -KeepTests
) else if "%choice%"=="4" (
    echo.
    echo 清理但保留日志文件...
    powershell -ExecutionPolicy Bypass -File .\cleanup-project-safe.ps1 -KeepLogs
) else if "%choice%"=="5" (
    echo.
    echo 清理但保留测试和日志文件...
    powershell -ExecutionPolicy Bypass -File .\cleanup-project-safe.ps1 -KeepTests -KeepLogs
) else if "%choice%"=="6" (
    echo.
    echo 退出清理工具
    exit /b 0
) else (
    echo.
    echo 无效的选项，请重新运行
)

echo.
pause