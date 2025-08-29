@echo off
echo Deployment with Alternative MySQL Port (3307)
echo =============================================
echo.
echo This script uses port 3307 for MySQL to avoid conflicts with local MySQL on port 3306
echo.

REM First, clean up any existing containers
echo [1/5] Cleaning up existing containers...
docker-compose -f docker-compose-alt-port.yml down --remove-orphans

REM Build images
echo.
echo [2/5] Building images...
docker-compose -f docker-compose-alt-port.yml build
if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

REM Start MySQL first
echo.
echo [3/5] Starting MySQL on port 3307...
docker-compose -f docker-compose-alt-port.yml up -d mysql

REM Wait for MySQL
echo.
echo [4/5] Waiting for MySQL to be ready (30 seconds)...
timeout /t 30 /nobreak

REM Start other services
echo.
echo [5/5] Starting application and nginx...
docker-compose -f docker-compose-alt-port.yml up -d app nginx

REM Check status
echo.
echo Checking deployment status...
docker-compose -f docker-compose-alt-port.yml ps

echo.
echo ========================================
echo Deployment completed!
echo ========================================
echo.
echo Access points:
echo - Application: http://localhost:8080
echo - Admin Panel: http://localhost:8080/admin  
echo - MySQL Port: 3307 (external), 3306 (internal)
echo - Default login: admin / admin123
echo.
echo Note: MySQL is accessible on port 3307 from your host machine
echo.
pause