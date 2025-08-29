@echo off
echo System Diagnosis Tool
echo ====================
echo.

echo [1] Current Docker containers:
docker ps -a
echo.

echo [2] Port 3306 usage:
netstat -ano | findstr :3306
echo.

echo [3] Port 3307 usage:
netstat -ano | findstr :3307
echo.

echo [4] Port 8080 usage:
netstat -ano | findstr :8080
echo.

echo [5] Port 80 usage:
netstat -ano | findstr ":80 "
echo.

echo [6] MySQL services:
sc query | findstr -i mysql
echo.

echo [7] Docker networks:
docker network ls
echo.

echo [8] Docker volumes:
docker volume ls
echo.

echo [9] Last 20 lines of app log (if exists):
docker logs at-mmbs-app --tail=20 2>nul || echo No app container found
echo.

echo [10] Last 20 lines of MySQL log (if exists):
docker logs at-mmbs-mysql --tail=20 2>nul || echo No MySQL container found
echo.

pause