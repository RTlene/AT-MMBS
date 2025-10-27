@echo off
echo Deploy with MySQL on Port 3307
echo ==============================
echo.
echo This will deploy MySQL on port 3307 to avoid conflicts.
echo.

REM Clean everything first
echo [1/5] Cleaning up old containers...
docker-compose down -v 2>nul
docker-compose -f docker-compose-no-nginx.yml down -v 2>nul
docker-compose -f docker-compose-alt-port.yml down -v 2>nul
docker-compose -f docker-compose-3307.yml down -v 2>nul
docker stop at-mmbs-mysql at-mmbs-app 2>nul
docker rm at-mmbs-mysql at-mmbs-app 2>nul

REM Build
echo.
echo [2/5] Building application...
docker-compose -f docker-compose-3307.yml build

REM Start MySQL
echo.
echo [3/5] Starting MySQL on port 3307...
docker-compose -f docker-compose-3307.yml up -d mysql

REM Wait
echo.
echo [4/5] Waiting for MySQL (40 seconds)...
timeout /t 40 /nobreak

REM Start app
echo.
echo [5/5] Starting application...
docker-compose -f docker-compose-3307.yml up -d app

REM Check status
echo.
echo Checking status...
docker ps

echo.
echo ========================================
echo Deployment should be complete!
echo ========================================
echo.
echo Access:
echo - Application: http://localhost:8080
echo - Admin Panel: http://localhost:8080/admin
echo - MySQL Port: 3307
echo.
echo Login: admin / admin123
echo.
echo Commands:
echo - View app logs: docker logs at-mmbs-app
echo - View MySQL logs: docker logs at-mmbs-mysql
echo - Stop all: docker-compose -f docker-compose-3307.yml down
echo.
pause