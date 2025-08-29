@echo off
echo Testing JWT Implementation
echo =========================
echo.

echo [1] Testing login to get JWT token...
curl -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     -o jwt_response.json

echo.
echo JWT Login Response:
type jwt_response.json
echo.
echo.

echo Please copy the token from above (the long string after "token":)
set /p jwt_token="Paste JWT token here: "

echo.
echo [2] Testing token validation...
curl -X GET http://localhost:8080/api/auth/validate ^
     -H "Authorization: Bearer %jwt_token%"

echo.
echo.
echo [3] Testing authenticated API call...
curl -X GET http://localhost:8080/api/users/profile ^
     -H "Authorization: Bearer %jwt_token%"

echo.
echo.
echo [4] JWT Implementation Benefits:
echo - Token contains user info (no database lookup needed)
echo - Stateless (works after container restart)
echo - Expires after 24 hours
echo - Secure with HMAC-SHA256 signature
echo.
echo [5] To decode your JWT token, visit: https://jwt.io
echo Paste your token there to see the payload!
echo.
pause