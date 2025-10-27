@echo off
echo AT-MMBS Project Cleanup Tool
echo =============================
echo.
echo This tool will help you clean unnecessary files from your project
echo.
echo Please choose cleanup mode:
echo 1. Preview mode (show files to be deleted, don't actually delete)
echo 2. Full cleanup (delete all unnecessary files)
echo 3. Keep test files
echo 4. Keep log files
echo 5. Keep both test and log files
echo 6. Exit
echo.
set /p choice="Enter your choice (1-6): "

if "%choice%"=="1" (
    echo.
    echo Running preview mode...
    powershell -ExecutionPolicy Bypass -File .\cleanup.ps1 -DryRun
) else if "%choice%"=="2" (
    echo.
    echo Starting full cleanup...
    powershell -ExecutionPolicy Bypass -File .\cleanup.ps1
) else if "%choice%"=="3" (
    echo.
    echo Cleaning but keeping test files...
    powershell -ExecutionPolicy Bypass -File .\cleanup.ps1 -KeepTests
) else if "%choice%"=="4" (
    echo.
    echo Cleaning but keeping log files...
    powershell -ExecutionPolicy Bypass -File .\cleanup.ps1 -KeepLogs
) else if "%choice%"=="5" (
    echo.
    echo Cleaning but keeping test and log files...
    powershell -ExecutionPolicy Bypass -File .\cleanup.ps1 -KeepTests -KeepLogs
) else if "%choice%"=="6" (
    echo.
    echo Exiting cleanup tool
    exit /b 0
) else (
    echo.
    echo Invalid option, please run again
)

echo.
pause