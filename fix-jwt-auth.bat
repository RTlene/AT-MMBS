@echo off
echo JWT Authentication Fix
echo =====================
echo.
echo This script will test and fix JWT authentication issues.
echo.

echo [1] Clearing browser cache...
echo Please press Ctrl+Shift+Delete in your browser and clear cache.
echo.
pause

echo.
echo [2] Testing login API...
curl -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     > login_response.txt 2>nul

echo.
echo Login response:
type login_response.txt
echo.

echo [3] Manual test steps:
echo.
echo 1. Open Chrome DevTools (F12)
echo 2. Go to Application tab
echo 3. Clear Local Storage for localhost:8080
echo 4. Go to http://localhost:8080/login.html
echo 5. Login with admin / admin123
echo 6. Check Console for errors
echo 7. Check Network tab - look for 401 errors
echo 8. Check Application > Local Storage - verify authToken exists
echo.
echo [4] If still getting 401:
echo - Check if Authorization header is sent in Network tab
echo - Format should be: Bearer [token]
echo.
pause