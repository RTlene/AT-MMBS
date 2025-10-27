@echo off
echo AT-MMBS Docker Deployment Script
echo ================================
echo.

REM Check if Docker is installed
docker --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker is not installed. Please install Docker Desktop first.
    pause
    exit /b 1
)

docker-compose --version >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker Compose is not installed. Please install Docker Compose first.
    pause
    exit /b 1
)

echo [OK] Docker environment check passed
echo.

REM Stop and remove existing containers
echo Stopping existing containers...
docker-compose down --remove-orphans

REM Ask if user wants to clean old images
set /p cleanup="Do you want to clean old images? (y/N): "
if /i "%cleanup%"=="y" (
    echo Cleaning old images...
    docker system prune -f
)

REM Build new images
echo.
echo Building application image...
docker-compose build --no-cache

if errorlevel 1 (
    echo.
    echo [ERROR] Build failed, please check the error messages above
    pause
    exit /b 1
)

REM Start services
echo.
echo Starting services...
docker-compose up -d

if errorlevel 1 (
    echo.
    echo [ERROR] Failed to start services, please check the error messages above
    pause
    exit /b 1
)

REM Wait for services to start
echo.
echo Waiting for services to start (30 seconds)...
timeout /t 30 /nobreak >nul

REM Check service status
echo.
echo Checking service status...
docker-compose ps

echo.
echo ========================================
echo Deployment completed successfully!
echo ========================================
echo.
echo Application URL: http://localhost:80
echo Admin Panel: http://localhost:80/admin
echo MySQL Port: 3306
echo.
echo Useful commands:
echo   View logs: docker-compose logs -f
echo   Stop services: docker-compose down
echo   Restart services: docker-compose restart
echo   Check status: docker-compose ps
echo.
pause