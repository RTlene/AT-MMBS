@echo off
echo Deploy Without Nginx (Port 80 Conflict Solution)
echo ================================================
echo.
echo This deployment:
echo - Runs MySQL on port 3306 (or 3307 if conflict)
echo - Runs the application on port 8080 (no nginx)
echo - Fixes JWT token authentication issues
echo.

REM Stop existing containers
echo [1/6] Stopping existing containers...
docker-compose down
docker-compose -f docker-compose-no-nginx.yml down
docker-compose -f docker-compose-alt-port.yml down

REM Check MySQL port
echo.
echo [2/6] Checking MySQL port availability...
netstat -ano | findstr :3306 | findstr LISTENING >nul
if errorlevel 1 (
    echo MySQL port 3306 is available
    set COMPOSE_FILE=docker-compose-no-nginx.yml
    set MYSQL_PORT_INFO=3306
) else (
    echo MySQL port 3306 is occupied, will use 3307
    REM Create a modified version with port 3307
    copy docker-compose-no-nginx.yml docker-compose-no-nginx-alt.yml >nul
    powershell -Command "(gc docker-compose-no-nginx-alt.yml) -replace '3306:3306', '3307:3306' | Out-File -encoding ASCII docker-compose-no-nginx-alt.yml"
    set COMPOSE_FILE=docker-compose-no-nginx-alt.yml
    set MYSQL_PORT_INFO=3307
)

REM Build images
echo.
echo [3/6] Building application image...
docker-compose -f %COMPOSE_FILE% build
if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

REM Start MySQL
echo.
echo [4/6] Starting MySQL on port %MYSQL_PORT_INFO%...
docker-compose -f %COMPOSE_FILE% up -d mysql

REM Wait for MySQL
echo.
echo [5/6] Waiting for MySQL to be ready (30 seconds)...
timeout /t 30 /nobreak

REM Start application
echo.
echo [6/6] Starting application on port 8080...
docker-compose -f %COMPOSE_FILE% up -d app

REM Show status
echo.
echo ========================================
echo Deployment Status:
echo ========================================
docker-compose -f %COMPOSE_FILE% ps

echo.
echo ========================================
echo DEPLOYMENT COMPLETED!
echo ========================================
echo.
echo Application Access:
echo ------------------
echo URL: http://localhost:8080
echo Admin Panel: http://localhost:8080/admin
echo.
echo Login Credentials:
echo -----------------
echo Username: admin
echo Password: admin123
echo.
echo Database Access:
echo ---------------
echo Host: localhost
echo Port: %MYSQL_PORT_INFO%
echo Username: mmbs_user
echo Password: mmbs_pass123
echo Database: golang_demo
echo.
echo Important: JWT authentication has been fixed.
echo Clear your browser cache if you still see 401 errors.
echo.
echo To view logs: docker-compose -f %COMPOSE_FILE% logs -f
echo To stop: docker-compose -f %COMPOSE_FILE% down
echo.
pause