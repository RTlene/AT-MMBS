@echo off
echo Clean Deploy - Complete Fresh Start
echo ==================================
echo.
echo This will:
echo 1. Stop all Docker containers
echo 2. Remove all containers and volumes
echo 3. Check port conflicts
echo 4. Deploy with appropriate configuration
echo.

REM Stop all containers
echo [1/7] Stopping all containers...
docker-compose down -v
docker-compose -f docker-compose-alt-port.yml down -v
docker stop at-mmbs-mysql at-mmbs-app at-mmbs-nginx 2>nul
docker rm at-mmbs-mysql at-mmbs-app at-mmbs-nginx 2>nul

echo.
echo [2/7] Checking port 3306...
netstat -ano | findstr :3306 | findstr LISTENING
if errorlevel 1 (
    echo Port 3306 is free!
    set USE_ALT_PORT=0
) else (
    echo Port 3306 is in use!
    echo Will use port 3307 for MySQL instead.
    set USE_ALT_PORT=1
)

echo.
echo [3/7] Checking port 80...
netstat -ano | findstr ":80 " | findstr LISTENING
if not errorlevel 1 (
    echo Warning: Port 80 is in use. Application will use port 8080.
)

echo.
echo [4/7] Building Docker images...
if "%USE_ALT_PORT%"=="1" (
    docker-compose -f docker-compose-alt-port.yml build
) else (
    docker-compose build
)

if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

echo.
echo [5/7] Starting MySQL...
if "%USE_ALT_PORT%"=="1" (
    docker-compose -f docker-compose-alt-port.yml up -d mysql
    echo MySQL will be accessible on port 3307
) else (
    docker-compose up -d mysql
    echo MySQL will be accessible on port 3306
)

echo.
echo [6/7] Waiting for MySQL to start (30 seconds)...
timeout /t 30 /nobreak

echo.
echo [7/7] Starting application services...
if "%USE_ALT_PORT%"=="1" (
    docker-compose -f docker-compose-alt-port.yml up -d
) else (
    docker-compose up -d
)

echo.
echo ========================================
echo Checking deployment status...
echo ========================================
if "%USE_ALT_PORT%"=="1" (
    docker-compose -f docker-compose-alt-port.yml ps
) else (
    docker-compose ps
)

echo.
echo ========================================
echo Deployment completed!
echo ========================================
echo.
if "%USE_ALT_PORT%"=="1" (
    echo MySQL is running on port 3307
) else (
    echo MySQL is running on port 3306
)
echo.
echo Access your application at:
echo - http://localhost (if port 80 is free)
echo - http://localhost:8080 (if port 80 was busy)
echo.
echo Default login: admin / admin123
echo.
pause