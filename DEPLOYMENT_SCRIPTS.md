# Deployment Scripts Guide

This guide explains the new English deployment scripts that replace the Chinese versions to avoid encoding issues.

## Available Scripts

### 1. build-docker.bat
Tests Docker image build without starting services.
```bash
build-docker.bat
```

### 2. quick-deploy.bat (Recommended)
One-click deployment script that builds and starts all services.
```bash
quick-deploy.bat
```

### 3. deploy-docker.bat
Full deployment with interactive options (clean images, etc.).
```bash
deploy-docker.bat
```

### 4. check-status.bat
Check if services are running correctly.
```bash
check-status.bat
```

## Quick Start

### Step 1: Pull Latest Code
```bash
git pull origin cursor/explain-background-mode-concept-201c
```

### Step 2: Restore Files (if needed)
```bash
git checkout HEAD -- .
```

### Step 3: Deploy Application
```bash
quick-deploy.bat
```

### Step 4: Verify Deployment
```bash
check-status.bat
```

## Access Points

After successful deployment:
- Application: http://localhost:80
- Admin Panel: http://localhost:80/admin
- Default Login: admin / admin123

## Troubleshooting

### Build Fails
1. Make sure Docker Desktop is running
2. Check internet connection
3. Run `docker system prune -a` to clean cache

### Cannot Access Application
1. Wait 30 seconds after deployment
2. Check if port 80 is in use: `netstat -an | findstr :80`
3. Check logs: `docker-compose logs`

### Port Conflicts
If port 80 is already in use, edit docker-compose.yml:
```yaml
ports:
  - "${APP_PORT:-8080}:80"  # Change 8080 to another port
```

## Docker Commands

```bash
# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Restart services
docker-compose restart

# Remove everything (including data)
docker-compose down -v

# Check running containers
docker ps
```