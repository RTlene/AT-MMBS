@echo off
echo Safe Deployment Script
echo =====================
echo This script ensures MySQL is ready before starting the application
echo.

REM Stop all services
echo [1/6] Stopping existing services...
docker-compose down --remove-orphans
echo.

REM Remove volumes to ensure clean state
echo Clean start? This will remove all data.
set /p cleanstart="Remove all volumes for fresh start? (y/N): "
if /i "%cleanstart%"=="y" (
    echo Removing volumes...
    docker-compose down -v
)

REM Build images
echo.
echo [2/6] Building images...
docker-compose build
if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

REM Start only MySQL first
echo.
echo [3/6] Starting MySQL service...
docker-compose up -d mysql
echo.

REM Wait for MySQL to be healthy
echo [4/6] Waiting for MySQL to be ready...
echo This may take up to 60 seconds...
:check_mysql
timeout /t 5 /nobreak >nul
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SELECT 1;" >nul 2>&1
if errorlevel 1 (
    echo MySQL not ready yet, waiting...
    goto check_mysql
)
echo MySQL is ready!

REM Verify database and user exist
echo.
echo [5/6] Verifying database setup...
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SHOW DATABASES;" | findstr golang_demo
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SELECT User FROM mysql.user WHERE User='mmbs_user';"

REM Start application and nginx
echo.
echo [6/6] Starting application services...
docker-compose up -d app nginx
echo.

REM Wait for app to start
echo Waiting for application to start (20 seconds)...
timeout /t 20 /nobreak

REM Check final status
echo.
echo ========================================
echo Deployment Status:
echo ========================================
docker-compose ps

echo.
echo Application Logs (last 10 lines):
docker logs at-mmbs-app --tail=10

echo.
echo ========================================
echo If successful, access:
echo - Application: http://localhost
echo - Admin Panel: http://localhost/admin
echo - Default login: admin / admin123
echo ========================================
echo.
pause