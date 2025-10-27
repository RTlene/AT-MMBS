@echo off
echo MySQL Connection Fix
echo ===================
echo.

echo [1] Stopping all services...
docker-compose down
echo.

echo [2] Starting MySQL service first...
docker-compose up -d mysql
echo.

echo [3] Waiting for MySQL to be ready (60 seconds)...
timeout /t 60 /nobreak

echo.
echo [4] Testing MySQL connection...
docker exec at-mmbs-mysql mysql -u mmbs_user -pmmbs_pass123 -D golang_demo -e "SELECT 'MySQL is ready!';"
if errorlevel 1 (
    echo MySQL is not ready yet. Waiting more...
    timeout /t 30 /nobreak
)

echo.
echo [5] Starting application service...
docker-compose up -d app
echo.

echo [6] Waiting for app to start (20 seconds)...
timeout /t 20 /nobreak

echo.
echo [7] Checking application status...
docker-compose ps
echo.

echo [8] Showing recent app logs...
docker logs at-mmbs-app --tail=30
echo.

echo If the app is still failing to connect, try:
echo 1. Run: docker-compose down -v
echo 2. Then run: quick-deploy.bat
echo.
pause