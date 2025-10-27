@echo off
echo Testing Complete Authentication Flow
echo ===================================
echo.

echo [1] Testing login...
curl -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     -o login.json

echo.
echo Login response:
type login.json
echo.
echo.

REM Extract token from response (this is a simple approach)
echo Please copy the token from above and paste it here:
set /p token="Token: "

echo.
echo [2] Testing authenticated request with your token...
curl -X GET http://localhost:8080/api/users/profile ^
     -H "Authorization: Bearer %token%" ^
     -H "Content-Type: application/json"

echo.
echo.
echo [3] Check if this is a container restart issue:
echo - If login works but profile returns 401, the token is not being recognized
echo - This happens because tokens are stored in memory and lost on restart
echo.
echo [4] Solutions:
echo 1. Use a fresh login after each container restart
echo 2. Implement persistent token storage (Redis/Database)
echo 3. Use stateless JWT tokens instead of memory storage
echo.
pause