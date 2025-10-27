@echo off
echo Port Conflict Fix Script
echo =======================
echo.

echo [1] Checking what's using port 3306...
netstat -ano | findstr :3306
echo.

echo [2] Checking for local MySQL service...
sc query | findstr -i mysql
echo.

echo Choose an option:
echo 1. Stop local MySQL service (if you have MySQL installed locally)
echo 2. Use different port for Docker MySQL (recommended)
echo 3. Force kill process using port 3306
echo 4. Exit
echo.
set /p choice="Enter your choice (1-4): "

if "%choice%"=="1" (
    echo.
    echo Stopping local MySQL service...
    net stop mysql
    net stop mysql57
    net stop mysql80
    echo Done. Now run safe-deploy.bat again.
) else if "%choice%"=="2" (
    echo.
    echo Using port 3307 for Docker MySQL...
    call :use_different_port
) else if "%choice%"=="3" (
    echo.
    echo Finding PID using port 3306...
    for /f "tokens=5" %%a in ('netstat -ano ^| findstr :3306 ^| findstr LISTENING') do (
        echo Killing process %%a
        taskkill /PID %%a /F
    )
    echo Done. Now run safe-deploy.bat again.
) else (
    echo Exiting...
    exit /b 0
)

pause
exit /b 0

:use_different_port
echo Creating modified docker-compose file...
echo Please run deploy-alt-port.bat next
exit /b 0