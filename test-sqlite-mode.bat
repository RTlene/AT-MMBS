@echo off
echo SQLite Test Mode Deployment
echo ===========================
echo.
echo This will run the application in test mode using SQLite
echo No MySQL required, no port conflicts!
echo.

REM Clean up
echo [1/4] Cleaning up old containers...
docker rm -f at-mmbs-app-sqlite 2>nul

REM Build
echo.
echo [2/4] Building application...
docker build -t at-mmbs-test-sqlite .

REM Run in SQLite mode
echo.
echo [3/4] Starting application in SQLite mode...
docker run -d ^
  --name at-mmbs-app-sqlite ^
  -p 8088:80 ^
  -e TEST_MODE=true ^
  -e SQLITE_PATH=/app/data/test.db ^
  -v %CD%/data:/app/data ^
  -v %CD%/uploads:/app/uploads ^
  -v %CD%/logs:/app/logs ^
  at-mmbs-test-sqlite

REM Wait
echo.
echo [4/4] Waiting for application to start (10 seconds)...
timeout /t 10 /nobreak

echo.
echo ========================================
echo SQLite Test Mode Deployment Complete!
echo ========================================
echo.
echo Access application at:
echo - URL: http://localhost:8088
echo - Admin: http://localhost:8088/admin
echo.
echo Login: admin / admin123
echo.
echo This is using SQLite database (no MySQL needed)
echo Data is stored in: %CD%\data\test.db
echo.
echo To stop: docker stop at-mmbs-app-sqlite
echo To view logs: docker logs at-mmbs-app-sqlite
echo.
pause