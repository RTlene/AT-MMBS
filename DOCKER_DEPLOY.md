# 商城后台系统 Docker 部署指南

## 📋 系统要求

- Docker Desktop 20.10+ 或 Docker Engine 20.10+
- Docker Compose 2.0+
- 至少 2GB 可用内存
- 至少 5GB 可用磁盘空间

## 🚀 快速部署

### 方法一：使用批处理脚本（Windows推荐）

```bash
# 双击运行
deploy.bat
```

### 方法二：使用Shell脚本（Linux/Mac推荐）

```bash
# 添加执行权限
chmod +x deploy.sh

# 运行部署脚本
./deploy.sh
```

### 方法三：手动部署

```bash
# 1. 停止现有服务
docker-compose down --remove-orphans

# 2. 构建镜像
docker-compose build --no-cache

# 3. 启动服务
docker-compose up -d

# 4. 查看服务状态
docker-compose ps
```

## 🔧 配置说明

### 环境变量

系统使用以下环境变量进行配置：

- `MYSQL_ROOT_PASSWORD`: MySQL root密码
- `MYSQL_DATABASE`: 数据库名称
- `MYSQL_USER`: 数据库用户名
- `MYSQL_PASSWORD`: 数据库密码
- `APP_PORT`: 应用端口（默认80）
- `APP_TIMEZONE`: 时区设置（默认Asia/Shanghai）

### 端口映射

- **应用服务**: 80:80
- **数据库服务**: 3306:3306

### 数据持久化

- MySQL数据存储在 `mysql_data` 卷中
- 应用静态文件包含在镜像中

## 📊 服务监控

### 健康检查

系统包含内置健康检查：

- **应用服务**: 每30秒检查HTTP响应
- **数据库服务**: 每30秒检查MySQL连接

### 查看服务状态

```bash
# 查看所有服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f app
docker-compose logs -f mysql

# 查看健康状态
docker-compose exec app wget --no-verbose --tries=1 --spider http://localhost:80/
docker-compose exec mysql mysqladmin ping -h localhost -u root -proot123456
```

## 🛠️ 常用操作

### 服务管理

```bash
# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 重启服务
docker-compose restart

# 重启特定服务
docker-compose restart app
docker-compose restart mysql
```

### 日志查看

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f app
docker-compose logs -f mysql

# 查看最近100行日志
docker-compose logs --tail=100 app
```

### 数据库操作

```bash
# 进入MySQL容器
docker-compose exec mysql mysql -u root -p

# 备份数据库
docker-compose exec mysql mysqldump -u root -proot123456 at_mmbs > backup.sql

# 恢复数据库
docker-compose exec -T mysql mysql -u root -proot123456 at_mmbs < backup.sql
```

### 镜像管理

```bash
# 查看镜像
docker images

# 清理未使用的镜像
docker system prune -f

# 重新构建镜像
docker-compose build --no-cache app
```

## 🔍 故障排除

### 常见问题

1. **端口被占用**
   ```bash
   # 检查端口占用
   netstat -ano | findstr :80
   
   # 修改docker-compose.yml中的端口映射
   ports:
     - "8080:80"  # 改为8080端口
   ```

2. **数据库连接失败**
   ```bash
   # 检查MySQL容器状态
   docker-compose ps mysql
   
   # 查看MySQL日志
   docker-compose logs mysql
   
   # 检查网络连接
   docker-compose exec app ping mysql
   ```

3. **应用启动失败**
   ```bash
   # 查看应用日志
   docker-compose logs app
   
   # 检查环境变量
   docker-compose exec app env | grep MYSQL
   ```

### 日志分析

```bash
# 实时查看错误日志
docker-compose logs -f --tail=100 app | grep -i error

# 查看特定时间段的日志
docker-compose logs --since="2024-01-01T00:00:00" app
```

## 📈 性能优化

### 资源限制

可以在 `docker-compose.yml` 中添加资源限制：

```yaml
services:
  app:
    deploy:
      resources:
        limits:
          memory: 512M
          cpus: '0.5'
        reservations:
          memory: 256M
          cpus: '0.25'
```

### 数据库优化

```yaml
services:
  mysql:
    command: >
      --default-authentication-plugin=mysql_native_password
      --character-set-server=utf8mb4
      --collation-server=utf8mb4_unicode_ci
      --innodb-buffer-pool-size=256M
      --max-connections=100
```

## 🔒 安全建议

1. **修改默认密码**: 在生产环境中修改所有默认密码
2. **限制网络访问**: 使用防火墙限制数据库端口访问
3. **定期更新**: 定期更新Docker镜像和系统依赖
4. **备份策略**: 实施定期数据库备份策略

## 📞 技术支持

如果遇到部署问题，请：

1. 检查Docker和Docker Compose版本
2. 查看服务日志和状态
3. 确认系统资源是否充足
4. 检查网络配置和防火墙设置

---

**注意**: 本系统仅供学习和测试使用，生产环境部署请根据实际需求调整配置。
