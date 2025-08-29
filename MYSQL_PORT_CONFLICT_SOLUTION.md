# MySQL Port Conflict Solution

## Problem
MySQL container cannot start because port 3306 is already in use:
```
Error response from daemon: ports are not available: exposing port TCP 0.0.0.0:3306 -> 127.0.0.1:0: listen tcp 0.0.0.0:3306: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
```

## Quick Solution

### Option 1: Use Clean Deploy Script (Recommended)
This script automatically detects port conflicts and chooses the right configuration:
```bash
clean-deploy.bat
```

### Option 2: Stop Local MySQL
If you have MySQL installed locally:
```bash
stop-local-mysql.bat
```
Then run:
```bash
safe-deploy.bat
```

### Option 3: Use Alternative Port
Deploy MySQL on port 3307 instead:
```bash
deploy-alt-port.bat
```

## Manual Solutions

### Check What's Using Port 3306
```bash
netstat -ano | findstr :3306
```

### Stop Local MySQL Service
```bash
# Windows services
net stop MySQL
net stop MySQL57
net stop MySQL80
net stop MariaDB

# Or via services.msc
services.msc
# Find MySQL and stop it
```

### Kill Process Using Port 3306
```bash
# Find PID
netstat -ano | findstr :3306

# Kill process (replace 1234 with actual PID)
taskkill /PID 1234 /F
```

## Using Alternative Port Configuration

If you need to keep local MySQL running, use the alternative configuration:

1. **Deploy with alternative port**:
   ```bash
   deploy-alt-port.bat
   ```

2. **Access MySQL**:
   - From host: `localhost:3307`
   - From containers: `mysql:3306`

3. **Connect from host**:
   ```bash
   mysql -h localhost -P 3307 -u mmbs_user -p
   # Password: mmbs_pass123
   ```

## Complete Reset

If nothing works:
```bash
# Stop everything
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

# Remove volumes
docker volume prune -f

# Try clean deploy
clean-deploy.bat
```

## Verify Deployment

After successful deployment:
```bash
# Check containers
docker ps

# Check logs
docker logs at-mmbs-mysql
docker logs at-mmbs-app

# Test MySQL connection
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SELECT 1;"
```

## Access Application

- If using standard ports:
  - Application: http://localhost
  - Admin: http://localhost/admin

- If using alternative ports:
  - Application: http://localhost:8080
  - Admin: http://localhost:8080/admin
  - MySQL: localhost:3307

Default login: `admin` / `admin123`