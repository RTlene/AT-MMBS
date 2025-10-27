@echo off
echo Quick Re-login Solution for 401 Error
echo ====================================
echo.
echo This is a temporary solution for the 401 authentication error.
echo The error occurs because tokens are stored in memory and lost when container restarts.
echo.

echo [1] Opening browser to clear cache and re-login...
echo.
echo Steps to follow:
echo.
echo 1. Press Ctrl+Shift+Delete to open clear browsing data
echo 2. Select "Cached images and files" 
echo 3. Click "Clear data"
echo 4. Go to: http://localhost:8080/login.html
echo 5. Login with: admin / admin123
echo.
echo Opening login page in 5 seconds...
timeout /t 5 /nobreak >nul

start http://localhost:8080/login.html

echo.
echo [2] After logging in successfully:
echo - You should be redirected to admin panel
echo - All features should work normally
echo - You'll need to login again if container restarts
echo.
echo [3] Permanent solution:
echo - Implement real JWT tokens (stateless)
echo - Or use Redis/Database for token storage
echo - See AUTH_ISSUE_SOLUTION.md for details
echo.
pause