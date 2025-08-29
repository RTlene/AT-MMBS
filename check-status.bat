@echo off
echo AT-MMBS Status Check
echo ====================
echo.

echo Checking Docker services...
docker-compose ps
echo.

echo Checking container logs (last 20 lines)...
echo ----------------------------------------
docker-compose logs --tail=20
echo.

echo Checking port usage...
netstat -an | findstr ":80 " | findstr "LISTENING"
netstat -an | findstr ":3306 " | findstr "LISTENING"
echo.

echo Testing application health...
curl -s -o nul -w "HTTP Status: %%{http_code}\n" http://localhost:80 2>nul
if errorlevel 1 (
    echo [WARNING] Cannot connect to application. It may still be starting up.
    echo Please wait a moment and try again.
) else (
    echo [OK] Application is responding
)

echo.
echo If services are running but you cannot access the application:
echo 1. Wait 30 seconds for services to fully start
echo 2. Check if port 80 is already in use by another application
echo 3. Try accessing http://localhost:8080 instead
echo 4. Run 'docker-compose logs' to see detailed error messages
echo.
pause