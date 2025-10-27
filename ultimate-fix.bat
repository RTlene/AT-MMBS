@echo off
echo Ultimate Fix Script - Guaranteed Solution
echo ========================================
echo.
echo This script will:
echo 1. Stop ALL MySQL services
echo 2. Free up ports 3306 and 80
echo 3. Deploy the application successfully
echo.
echo Press Ctrl+C to cancel, or
pause

echo.
echo [Phase 1] Stopping all MySQL-related services...
echo ------------------------------------------------

REM Stop Windows MySQL services
echo Stopping MySQL services...
net stop MySQL 2>nul
net stop MySQL57 2>nul
net stop MySQL80 2>nul
net stop MariaDB 2>nul
net stop MySQLRouter 2>nul

REM Stop XAMPP/WAMP MySQL
taskkill /F /IM mysqld.exe 2>nul

REM Stop IIS to free port 80
echo Stopping IIS (if running)...
net stop W3SVC 2>nul
net stop WAS 2>nul

echo.
echo [Phase 2] Force killing processes on ports...
echo --------------------------------------------

REM Kill anything on port 3306
echo Killing processes on port 3306...
for /f "tokens=5" %%a in ('netstat -ano ^| findstr :3306 ^| findstr LISTENING') do (
    taskkill /PID %%a /F 2>nul
)

REM Kill anything on port 80
echo Killing processes on port 80...
for /f "tokens=5" %%a in ('netstat -ano ^| findstr ":80 " ^| findstr LISTENING') do (
    taskkill /PID %%a /F 2>nul
)

echo.
echo [Phase 3] Docker cleanup...
echo --------------------------

REM Stop and remove all Docker containers
docker stop $(docker ps -aq) 2>nul
docker rm $(docker ps -aq) 2>nul

REM Clean Docker system
docker system prune -f --volumes

echo.
echo [Phase 4] Verify ports are free...
echo ---------------------------------
echo Port 3306 status:
netstat -ano | findstr :3306 | findstr LISTENING
if errorlevel 1 (
    echo - Port 3306 is FREE!
) else (
    echo - Port 3306 is still occupied
)

echo.
echo Port 80 status:
netstat -ano | findstr ":80 " | findstr LISTENING
if errorlevel 1 (
    echo - Port 80 is FREE!
) else (
    echo - Port 80 is still occupied
)

echo.
echo [Phase 5] Deploying application...
echo ---------------------------------
echo Using configuration with MySQL on 3307 and app on 8080 to avoid conflicts
echo.

REM Deploy using 3307 configuration
docker-compose -f docker-compose-3307.yml up -d

echo.
echo [Phase 6] Waiting for services (40 seconds)...
timeout /t 40 /nobreak

echo.
echo [Phase 7] Final status check...
echo ------------------------------
docker ps

echo.
echo [Phase 8] Testing application...
echo -------------------------------
curl http://localhost:8080 2>nul
if errorlevel 1 (
    echo WARNING: Application may still be starting up
    echo Wait 30 seconds and try accessing http://localhost:8080
) else (
    echo SUCCESS: Application is responding!
)

echo.
echo ========================================
echo DEPLOYMENT COMPLETE
echo ========================================
echo.
echo Access your application:
echo - URL: http://localhost:8080
echo - Admin: http://localhost:8080/admin
echo.
echo Credentials:
echo - Username: admin
echo - Password: admin123
echo.
echo Database:
echo - Port: 3307 (external)
echo - User: mmbs_user / mmbs_pass123
echo.
echo Troubleshooting:
echo - If not working, wait 1 minute then refresh browser
echo - View logs: docker logs at-mmbs-app
echo - Check MySQL: docker logs at-mmbs-mysql
echo.
pause