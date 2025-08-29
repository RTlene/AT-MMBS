@echo off
echo Rebuilding with JWT Authentication
echo ==================================
echo.
echo This script rebuilds the application with real JWT authentication.
echo No more 401 errors after container restart!
echo.

echo [1] Stopping existing containers...
docker-compose -f docker-compose-3307.yml down
docker-compose -f docker-compose-no-nginx.yml down

echo.
echo [2] Building new image with JWT support...
docker-compose -f docker-compose-3307.yml build --no-cache

if errorlevel 1 (
    echo.
    echo Build failed!
    pause
    exit /b 1
)

echo.
echo [3] Starting MySQL...
docker-compose -f docker-compose-3307.yml up -d mysql

echo.
echo [4] Waiting for MySQL (30 seconds)...
timeout /t 30 /nobreak

echo.
echo [5] Starting application with JWT support...
docker-compose -f docker-compose-3307.yml up -d app

echo.
echo [6] Waiting for application to start (20 seconds)...
timeout /t 20 /nobreak

echo.
echo [7] Checking status...
docker ps

echo.
echo ========================================
echo JWT Implementation Complete!
echo ========================================
echo.
echo Application: http://localhost:8080
echo Admin Panel: http://localhost:8080/admin
echo.
echo Login: admin / admin123
echo.
echo IMPORTANT CHANGES:
echo - JWT tokens are now stateless
echo - Tokens remain valid for 24 hours
echo - No more 401 errors after restart!
echo.
echo To view logs: docker logs at-mmbs-app --tail=50
echo.
pause