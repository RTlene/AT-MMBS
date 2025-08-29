@echo off
echo Stop Local MySQL Services
echo ========================
echo.
echo This script will stop any local MySQL services running on your machine
echo.

echo [1] Checking for MySQL services...
echo.

REM Check various MySQL service names
sc query MySQL >nul 2>&1
if not errorlevel 1 (
    echo Found MySQL service
    echo Stopping...
    net stop MySQL
)

sc query MySQL57 >nul 2>&1
if not errorlevel 1 (
    echo Found MySQL57 service
    echo Stopping...
    net stop MySQL57
)

sc query MySQL80 >nul 2>&1
if not errorlevel 1 (
    echo Found MySQL80 service
    echo Stopping...
    net stop MySQL80
)

sc query MariaDB >nul 2>&1
if not errorlevel 1 (
    echo Found MariaDB service
    echo Stopping...
    net stop MariaDB
)

echo.
echo [2] Checking if port 3306 is still in use...
netstat -ano | findstr :3306 | findstr LISTENING
if errorlevel 1 (
    echo.
    echo SUCCESS: Port 3306 is now free!
    echo You can now run safe-deploy.bat
) else (
    echo.
    echo WARNING: Port 3306 is still in use by another process
    echo.
    echo Finding process using port 3306...
    for /f "tokens=5" %%a in ('netstat -ano ^| findstr :3306 ^| findstr LISTENING') do (
        echo Process ID: %%a
        tasklist /FI "PID eq %%a"
        echo.
        set /p kill="Kill this process? (Y/N): "
        if /i "!kill!"=="Y" (
            taskkill /PID %%a /F
        )
    )
)

echo.
pause