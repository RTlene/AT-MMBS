# 端口冲突快速解决方案

## 问题描述
MySQL无法启动，因为端口3306已被占用（可能是您本地安装的MySQL服务）。

## 最快解决方案

### 方案1：一键清理部署（推荐）
```bash
clean-deploy.bat
```
这个脚本会：
- 自动检测端口占用
- 如果3306被占用，自动使用3307端口
- 完成全部部署流程

### 方案2：停止本地MySQL后部署
```bash
# 第一步：停止本地MySQL
stop-local-mysql.bat

# 第二步：正常部署
safe-deploy.bat
```

### 方案3：直接使用备用端口部署
```bash
deploy-alt-port.bat
```
MySQL将运行在3307端口，不影响您的本地MySQL。

## 部署后访问

### 如果使用标准端口（3306）：
- 应用地址：http://localhost
- 管理后台：http://localhost/admin

### 如果使用备用端口（3307）：
- 应用地址：http://localhost:8080
- 管理后台：http://localhost:8080/admin
- MySQL端口：3307（从主机访问）

### 默认登录账号：
- 用户名：admin
- 密码：admin123

## 如果还有问题

1. **查看占用端口的进程**：
   ```bash
   fix-port-conflict.bat
   ```

2. **检查部署状态**：
   ```bash
   docker ps
   docker logs at-mmbs-mysql
   ```

3. **完全重置**：
   ```bash
   docker stop $(docker ps -aq)
   docker rm $(docker ps -aq)
   docker volume prune -f
   clean-deploy.bat
   ```