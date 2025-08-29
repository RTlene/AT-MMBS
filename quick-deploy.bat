@echo off
echo AT-MMBS Quick Deploy
echo ====================
echo.
echo This script will quickly deploy your application using Docker.
echo.

REM Check Docker
docker --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker is not installed. Please install Docker Desktop.
    echo Download from: https://www.docker.com/products/docker-desktop
    pause
    exit /b 1
)

echo [1/4] Stopping old containers...
docker-compose down --remove-orphans >nul 2>&1

echo [2/4] Building application...
docker-compose build

if errorlevel 1 (
    echo.
    echo [ERROR] Build failed. Common solutions:
    echo - Check if Docker Desktop is running
    echo - Make sure you have internet connection
    echo - Try running: docker system prune -a
    pause
    exit /b 1
)

echo [3/4] Starting services...
docker-compose up -d

if errorlevel 1 (
    echo.
    echo [ERROR] Failed to start services
    pause
    exit /b 1
)

echo [4/4] Waiting for services to be ready...
timeout /t 10 /nobreak >nul

echo.
echo ========================================
echo DEPLOYMENT SUCCESSFUL!
echo ========================================
echo.
echo Your application is now running at:
echo - Main App: http://localhost:80
echo - Admin Panel: http://localhost:80/admin
echo - Default login: admin / admin123
echo.
echo To stop the application, run: docker-compose down
echo To view logs, run: docker-compose logs -f
echo.
pause