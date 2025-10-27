# AT-MMBS 部署指南

## 目录
- [快速部署](#快速部署)
- [环境要求](#环境要求)
- [详细部署步骤](#详细部署步骤)
- [配置说明](#配置说明)
- [常见问题](#常见问题)
- [维护指南](#维护指南)

## 快速部署

### 一键部署（推荐）
```bash
# 克隆仓库
git clone https://github.com/yourusername/AT-MMBS.git
cd AT-MMBS

# 执行部署脚本
chmod +x deploy.sh
./deploy.sh
```

### 手动部署
```bash
# 1. 复制环境配置
cp .env.example .env

# 2. 修改配置（按需）
nano .env

# 3. 启动服务
docker-compose up -d

# 4. 查看日志
docker-compose logs -f
```

## 环境要求

### 系统要求
- Linux/macOS/Windows (支持Docker)
- Docker 20.10+
- Docker Compose 1.29+
- 2GB+ RAM
- 10GB+ 磁盘空间

### 端口占用
默认使用以下端口（可在.env中修改）：
- 8080: 应用主端口
- 3306: MySQL端口
- 80: Nginx端口（可选）

## 详细部署步骤

### 1. 准备工作
```bash
# 安装Docker（Ubuntu示例）
sudo apt update
sudo apt install docker.io docker-compose -y

# 启动Docker服务
sudo systemctl start docker
sudo systemctl enable docker

# 添加当前用户到docker组（避免使用sudo）
sudo usermod -aG docker $USER
```

### 2. 获取代码
```bash
git clone https://github.com/yourusername/AT-MMBS.git
cd AT-MMBS
```

### 3. 配置环境
```bash
# 复制配置文件
cp .env.example .env

# 编辑配置（重要）
vim .env
```

### 4. 启动服务
```bash
# 使用docker-compose启动
docker-compose up -d

# 或使用部署脚本
./deploy.sh
```

### 5. 验证部署
```bash
# 检查服务状态
docker-compose ps

# 查看日志
docker-compose logs -f app

# 测试访问
curl http://localhost:8080
```

## 配置说明

### 环境变量(.env)
```bash
# MySQL配置
MYSQL_ROOT_PASSWORD=root123456    # MySQL root密码
MYSQL_DATABASE=golang_demo         # 数据库名
MYSQL_USER=mmbs_user              # 数据库用户
MYSQL_PASSWORD=mmbs_pass123       # 数据库密码
MYSQL_PORT=3306                   # MySQL端口

# 应用配置
APP_PORT=8080                     # 应用访问端口
JWT_SECRET=your-secret-key        # JWT密钥（必须修改）

# Nginx配置
NGINX_PORT=80                     # HTTP端口
NGINX_SSL_PORT=443               # HTTPS端口
```

### docker-compose.yml
主要服务：
- `mysql`: 数据库服务
- `app`: Go应用服务
- `nginx`: 反向代理（可选）

### 目录结构
```
AT-MMBS/
├── uploads/        # 上传文件目录
├── logs/           # 日志目录
├── ssl/            # SSL证书目录
└── mysql_data/     # MySQL数据目录（Docker卷）
```

## 常见问题

### Q: 端口被占用怎么办？
修改`.env`文件中的端口配置：
```bash
APP_PORT=8081
MYSQL_PORT=3307
```

### Q: 如何修改默认管理员密码？
1. 登录系统：http://localhost:8080/admin
2. 使用默认账号：admin/admin123
3. 进入个人设置修改密码

### Q: 如何配置HTTPS？
1. 准备SSL证书文件
2. 放置到`ssl/`目录
3. 修改`nginx.conf`启用HTTPS配置
4. 重启nginx服务

### Q: 数据库连接失败？
检查以下内容：
1. MySQL服务是否正常：`docker-compose ps mysql`
2. 环境变量是否正确
3. 防火墙是否阻止连接

## 维护指南

### 日常维护命令
```bash
# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f [service_name]

# 重启服务
docker-compose restart [service_name]

# 停止服务
docker-compose down

# 停止并删除数据
docker-compose down -v
```

### 备份数据库
```bash
# 备份
docker-compose exec mysql mysqldump -u root -p golang_demo > backup.sql

# 恢复
docker-compose exec -T mysql mysql -u root -p golang_demo < backup.sql
```

### 更新应用
```bash
# 拉取最新代码
git pull

# 重新构建并启动
docker-compose build
docker-compose up -d
```

### 查看应用日志
```bash
# 实时查看
docker-compose logs -f app

# 查看最近100行
docker-compose logs --tail=100 app
```

### 性能监控
```bash
# 查看资源使用
docker stats

# 查看磁盘使用
df -h
du -sh uploads/
```

## 生产环境建议

1. **安全性**
   - 修改所有默认密码
   - 使用强密码策略
   - 配置防火墙规则
   - 启用HTTPS
   - 定期更新系统

2. **性能优化**
   - 调整MySQL配置
   - 配置Redis缓存
   - 使用CDN加速静态资源
   - 启用Gzip压缩

3. **高可用性**
   - 配置数据库主从复制
   - 使用负载均衡
   - 定期备份数据
   - 监控服务状态

4. **日志管理**
   - 配置日志轮转
   - 集中日志管理
   - 设置日志级别

## 联系支持

如遇到问题，请通过以下方式获取帮助：
- GitHub Issues: https://github.com/yourusername/AT-MMBS/issues
- 邮箱: support@example.com