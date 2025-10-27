# MySQL Connection Troubleshooting Guide

## Problem
The application fails to connect to MySQL with error:
```
dial tcp 127.0.0.1:3306: connect: connection refused
```

## Root Cause
The application is trying to connect to `127.0.0.1:3306` instead of `mysql:3306` (the MySQL container hostname).

## Solutions

### Solution 1: Use Safe Deploy Script (Recommended)
```bash
safe-deploy.bat
```
This script ensures MySQL is fully ready before starting the application.

### Solution 2: Fix MySQL Connection Step by Step

1. **Stop all services**
   ```bash
   docker-compose down -v
   ```

2. **Create .env file** (already created)
   The .env file contains proper MySQL configuration

3. **Start MySQL first**
   ```bash
   docker-compose up -d mysql
   ```

4. **Wait and test MySQL**
   ```bash
   test-mysql.bat
   ```

5. **Start application**
   ```bash
   docker-compose up -d app nginx
   ```

### Solution 3: Debug Current State
```bash
docker-debug.bat
```
This will show you what's wrong with the current deployment.

### Solution 4: Manual Fix

1. **Check if MySQL is running**
   ```bash
   docker ps | findstr mysql
   ```

2. **Check app environment variables**
   ```bash
   docker exec at-mmbs-app printenv | findstr MYSQL
   ```

3. **If MYSQL_ADDRESS is empty, restart with proper environment**
   ```bash
   docker-compose down
   docker-compose up -d
   ```

## Expected Environment Variables

The application needs these environment variables:
- `MYSQL_USERNAME=mmbs_user`
- `MYSQL_PASSWORD=mmbs_pass123`
- `MYSQL_ADDRESS=mysql:3306`
- `MYSQL_DATABASE=golang_demo`

## Testing Connection

After deployment, test the connection:
```bash
docker exec at-mmbs-app ping mysql
docker exec at-mmbs-app nslookup mysql
```

## Common Issues

1. **MySQL not ready**: The application starts before MySQL is ready
   - Solution: Use safe-deploy.bat which waits for MySQL

2. **Environment variables not passed**: Docker Compose not reading .env file
   - Solution: Ensure .env file exists in the same directory

3. **Network issues**: Containers not on the same network
   - Solution: Check docker network with `docker network ls`

4. **Wrong credentials**: MySQL user not created
   - Solution: Run test-mysql.bat to verify and create user

## Quick Fix Commands

```bash
# Complete fresh start
docker-compose down -v
safe-deploy.bat

# Or use the one-liner
docker-compose down -v && docker-compose up -d mysql && timeout /t 30 && docker-compose up -d
```