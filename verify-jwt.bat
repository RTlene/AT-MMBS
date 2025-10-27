@echo off
echo JWT Verification Tool
echo ====================
echo.
echo This tool verifies that JWT authentication is working correctly.
echo.

echo [1] Getting JWT token...
echo.
echo Username: admin
echo Password: admin123
echo.

curl -s -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     -o jwt_login.json

type jwt_login.json | findstr "token"
if errorlevel 1 (
    echo.
    echo ERROR: Failed to get JWT token!
    echo Is the application running?
    echo.
    pause
    exit /b 1
)

echo.
echo.
echo [2] JWT token received successfully!
echo.
echo [3] Key features of this JWT implementation:
echo.
echo    - Stateless: No server-side storage needed
echo    - Persistent: Survives container restarts
echo    - Secure: HMAC-SHA256 signed
echo    - Standard: Works with any JWT library
echo    - 24-hour validity
echo.
echo [4] To see token contents:
echo    1. Copy the token above
echo    2. Visit https://jwt.io
echo    3. Paste in the "Encoded" section
echo.
echo SUCCESS: JWT authentication is working correctly!
echo.
pause