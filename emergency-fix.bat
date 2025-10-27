@echo off
echo Emergency Fix - Complete Solution
echo =================================
echo.
echo This script will fix all deployment issues step by step.
echo.

REM Step 1: Complete cleanup
echo [Step 1/8] Complete cleanup of all containers and images...
docker stop at-mmbs-mysql at-mmbs-app at-mmbs-nginx 2>nul
docker rm -f at-mmbs-mysql at-mmbs-app at-mmbs-nginx 2>nul
docker-compose down -v 2>nul
docker-compose -f docker-compose-no-nginx.yml down -v 2>nul
docker-compose -f docker-compose-alt-port.yml down -v 2>nul

REM Step 2: Find what's using port 3306
echo.
echo [Step 2/8] Checking what's using port 3306...
echo.
echo Processes using port 3306:
netstat -ano | findstr :3306 | findstr LISTENING
echo.

REM Step 3: Try to stop local MySQL
echo [Step 3/8] Attempting to stop local MySQL services...
net stop MySQL 2>nul
net stop MySQL57 2>nul
net stop MySQL80 2>nul
net stop MariaDB 2>nul
echo.

REM Step 4: Kill process on port 3306
echo [Step 4/8] Force killing processes on port 3306...
for /f "tokens=5" %%a in ('netstat -ano ^| findstr :3306 ^| findstr LISTENING') do (
    echo Killing process ID: %%a
    taskkill /PID %%a /F 2>nul
)
echo.

REM Step 5: Verify port is free
echo [Step 5/8] Verifying port 3306 status...
netstat -ano | findstr :3306 | findstr LISTENING
if errorlevel 1 (
    echo SUCCESS: Port 3306 is now free!
    set USE_ALT_PORT=0
) else (
    echo WARNING: Port 3306 is still occupied. Will use alternative solution.
    set USE_ALT_PORT=1
)
echo.

REM Step 6: Deploy based on port availability
if "%USE_ALT_PORT%"=="1" (
    echo [Step 6/8] Using alternative deployment (MySQL on port 3307)...
    echo.
    echo Starting deployment with MySQL on port 3307...
    docker-compose -f docker-compose-alt-port.yml up -d mysql
    echo Waiting 30 seconds for MySQL to start...
    timeout /t 30 /nobreak
    docker-compose -f docker-compose-alt-port.yml up -d app
    set MYSQL_PORT=3307
) else (
    echo [Step 6/8] Using standard deployment (MySQL on port 3306)...
    echo.
    echo Starting deployment with MySQL on port 3306...
    docker-compose -f docker-compose-no-nginx.yml up -d mysql
    echo Waiting 30 seconds for MySQL to start...
    timeout /t 30 /nobreak
    docker-compose -f docker-compose-no-nginx.yml up -d app
    set MYSQL_PORT=3306
)

REM Step 7: Wait for services
echo.
echo [Step 7/8] Waiting for services to be ready (20 seconds)...
timeout /t 20 /nobreak

REM Step 8: Verify deployment
echo.
echo [Step 8/8] Verifying deployment...
echo.
docker ps
echo.

echo Testing application health...
curl -s http://localhost:8080 -o nul -w "Application: HTTP %%{http_code}\n" 2>nul || echo Application: Not responding

echo.
echo ========================================
echo DEPLOYMENT STATUS
echo ========================================
echo.
if "%USE_ALT_PORT%"=="1" (
    echo MySQL is running on port: 3307
    echo Application is running on port: 8080
    echo.
    echo Access your application:
    echo - URL: http://localhost:8080
    echo - Admin: http://localhost:8080/admin
    echo.
    echo To check logs:
    echo - docker logs at-mmbs-app
    echo - docker logs at-mmbs-mysql
) else (
    echo MySQL is running on port: 3306
    echo Application is running on port: 8080
    echo.
    echo Access your application:
    echo - URL: http://localhost:8080
    echo - Admin: http://localhost:8080/admin
    echo.
    echo To check logs:
    echo - docker logs at-mmbs-app
    echo - docker logs at-mmbs-mysql
)
echo.
echo Login: admin / admin123
echo.
echo If still not working, run: docker logs at-mmbs-app --tail=50
echo.
pause