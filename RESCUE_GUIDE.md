# 紧急救援指南 - AT-MMBS部署

## 当前问题
1. MySQL无法启动，因为3306端口被占用
2. 应用容器显示unhealthy状态
3. 无法访问应用

## 快速解决方案

### 方案1：终极修复脚本（推荐）
```bash
ultimate-fix.bat
```
这个脚本会：
- 停止所有MySQL服务
- 强制释放端口
- 使用3307端口部署MySQL
- 在8080端口部署应用

### 方案2：紧急修复脚本
```bash
emergency-fix.bat
```

### 方案3：直接使用3307端口部署
```bash
deploy-3307.bat
```

## 手动解决步骤

如果脚本不工作，手动执行：

### 1. 完全停止MySQL
```bash
# 停止服务
net stop MySQL
net stop MySQL57
net stop MySQL80

# 强制终止进程
taskkill /F /IM mysqld.exe
```

### 2. 清理Docker
```bash
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker volume prune -f
```

### 3. 检查端口
```bash
netstat -ano | findstr :3306
netstat -ano | findstr :80
```

### 4. 部署
```bash
docker-compose -f docker-compose-3307.yml up -d
```

### 5. 等待并检查
```bash
# 等待40秒
timeout /t 40

# 检查状态
docker ps

# 查看日志
docker logs at-mmbs-app
docker logs at-mmbs-mysql
```

## 访问应用

部署成功后：
- **地址**: http://localhost:8080
- **管理后台**: http://localhost:8080/admin
- **账号**: admin / admin123

## 常见问题解决

### Q: 还是看到端口占用错误？
```bash
# 查找占用3306端口的进程
netstat -ano | findstr :3306
# 记下PID（最后一列数字）
taskkill /PID [PID号] /F
```

### Q: 应用显示unhealthy？
```bash
# 查看具体错误
docker logs at-mmbs-app --tail=100
```

### Q: 无法连接MySQL？
```bash
# 测试MySQL连接
docker exec at-mmbs-mysql mysql -u root -proot123456 -e "SELECT 1;"
```

### Q: 浏览器显示无法访问？
1. 等待60秒让服务完全启动
2. 清除浏览器缓存
3. 检查防火墙设置
4. 尝试 http://127.0.0.1:8080

## 最后的办法

如果所有方法都失败：

1. **重启电脑**（释放所有端口）
2. **运行**：
   ```bash
   ultimate-fix.bat
   ```

## 获取帮助

如果仍然有问题，运行以下命令收集信息：
```bash
# 收集诊断信息
echo "=== Docker状态 ===" > debug.txt
docker ps -a >> debug.txt
echo "=== 端口状态 ===" >> debug.txt
netstat -ano | findstr "3306\|3307\|8080\|80" >> debug.txt
echo "=== 应用日志 ===" >> debug.txt
docker logs at-mmbs-app --tail=50 >> debug.txt
echo "=== MySQL日志 ===" >> debug.txt  
docker logs at-mmbs-mysql --tail=50 >> debug.txt
```

然后查看debug.txt文件内容。