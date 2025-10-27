@echo off
echo Testing JWT Authentication Fix
echo ==============================
echo.
echo This script will test if the JWT authentication is working correctly.
echo.

echo [1] First, let's check if the container is running...
docker ps | findstr at-mmbs-app
if errorlevel 1 (
    echo ERROR: Application container is not running!
    echo Please run deploy-no-nginx.bat first.
    pause
    exit /b 1
)

echo.
echo [2] Testing login endpoint...
echo Logging in as admin...
curl -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     -w "\n\nHTTP Status: %%{http_code}\n"

echo.
echo [3] To test authenticated requests manually:
echo.
echo 1. Open your browser and go to: http://localhost:8080/login.html
echo 2. Login with: admin / admin123
echo 3. After login, you should be redirected to the admin panel
echo 4. Check browser console (F12) for any errors
echo.
echo [4] Clear browser cache if needed:
echo - Chrome: Ctrl+Shift+Delete
echo - Firefox: Ctrl+Shift+Delete
echo - Edge: Ctrl+Shift+Delete
echo.
echo [5] Check application logs for errors:
echo docker logs at-mmbs-app --tail=50
echo.
pause