@echo off
echo Docker Debug Script
echo ==================
echo.

echo [1] Checking running containers...
docker ps -a
echo.

echo [2] Checking docker-compose services...
docker-compose ps
echo.

echo [3] Checking MySQL container logs...
docker logs at-mmbs-mysql --tail=20
echo.

echo [4] Checking app container logs...
docker logs at-mmbs-app --tail=20
echo.

echo [5] Checking network...
docker network ls | findstr mmbs
echo.

echo [6] Testing MySQL connection from host...
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SELECT 1;"
if errorlevel 1 (
    echo MySQL is not ready or credentials are incorrect
) else (
    echo MySQL connection successful
)

echo.
echo [7] Checking environment variables in app container...
docker exec at-mmbs-app printenv | findstr MYSQL
echo.

pause