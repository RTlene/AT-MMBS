@echo off
echo Deployment Verification Script
echo ==============================
echo.

echo [1] Checking Docker containers...
echo --------------------------------
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
echo.

echo [2] Testing application health...
echo --------------------------------
curl -s http://localhost:8080 -o nul -w "Application (port 8080): HTTP %%{http_code}\n" 2>nul || echo Application (port 8080): Not responding
curl -s http://localhost -o nul -w "Application (port 80): HTTP %%{http_code}\n" 2>nul || echo Application (port 80): Not responding
echo.

echo [3] Testing login API...
echo --------------------------------
echo Attempting login as admin...
curl -X POST http://localhost:8080/api/auth/login ^
     -H "Content-Type: application/json" ^
     -d "{\"username\":\"admin\",\"password\":\"admin123\"}" ^
     -w "\nHTTP Status: %%{http_code}\n" 2>nul

echo.
echo [4] Quick Links:
echo --------------------------------
echo If deployment is successful, access:
echo.
echo Port 8080 deployment (no nginx):
echo   - Application: http://localhost:8080
echo   - Admin Panel: http://localhost:8080/admin
echo.
echo Port 80 deployment (with nginx):
echo   - Application: http://localhost
echo   - Admin Panel: http://localhost/admin
echo.
echo [5] Troubleshooting Commands:
echo --------------------------------
echo View logs: docker logs at-mmbs-app --tail=50
echo Check MySQL: docker logs at-mmbs-mysql --tail=20
echo Restart app: docker restart at-mmbs-app
echo.
pause