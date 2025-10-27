@echo off
echo Testing MySQL Connection
echo =======================
echo.

echo [1] Checking if MySQL container is running...
docker ps | findstr mysql
if errorlevel 1 (
    echo MySQL container is not running!
    echo Please run: docker-compose up -d mysql
    pause
    exit /b 1
)

echo.
echo [2] Testing root connection...
docker exec -it at-mmbs-mysql mysql -u root -proot123456 -e "SELECT VERSION();"
if errorlevel 1 (
    echo Failed to connect as root
) else (
    echo Root connection OK
)

echo.
echo [3] Checking if database exists...
docker exec -it at-mmbs-mysql mysql -u root -proot123456 -e "SHOW DATABASES;" | findstr golang_demo
if errorlevel 1 (
    echo Database 'golang_demo' not found!
) else (
    echo Database 'golang_demo' exists
)

echo.
echo [4] Testing application user...
docker exec -it at-mmbs-mysql mysql -u mmbs_user -pmmbs_pass123 -D golang_demo -e "SELECT 'User connection OK';"
if errorlevel 1 (
    echo Failed to connect as mmbs_user
    echo Creating user...
    docker exec -it at-mmbs-mysql mysql -u root -proot123456 -e "CREATE USER IF NOT EXISTS 'mmbs_user'@'%%' IDENTIFIED BY 'mmbs_pass123'; GRANT ALL ON golang_demo.* TO 'mmbs_user'@'%%'; FLUSH PRIVILEGES;"
) else (
    echo User connection OK
)

echo.
echo [5] Checking tables...
docker exec -it at-mmbs-mysql mysql -u mmbs_user -pmmbs_pass123 -D golang_demo -e "SHOW TABLES;"

echo.
echo [6] Testing from app network...
docker run --rm --network mmbs-network mysql:8.0 mysql -h mysql -u mmbs_user -pmmbs_pass123 -D golang_demo -e "SELECT 'Network connection OK';"

echo.
pause