# 修复401认证错误和端口冲突问题

## 问题描述

1. **401错误**：登录成功后访问主页返回 `{"code":401,"data":null,"errorMsg":"未提供认证信息"}`
2. **端口冲突**：Nginx无法启动，80端口被占用

## 问题原因

### 401错误原因
- 前端代码不一致：登录页面保存token的key是`authToken`，但admin.js查找的是`adminToken`
- 导致JWT token无法正确传递到后端

### 端口冲突原因
- 本地可能安装了IIS、Apache或其他Web服务器占用80端口
- Nginx容器无法绑定到已占用的端口

## 快速解决方案

### 步骤1：拉取最新修复代码
```bash
git pull origin cursor/explain-background-mode-concept-201c
```

### 步骤2：部署（不使用Nginx，避免80端口冲突）
```bash
deploy-no-nginx.bat
```

这个脚本会：
- ✅ 自动检测3306端口是否被占用
- ✅ 在8080端口运行应用（避开80端口）
- ✅ 不启动Nginx（避免端口冲突）
- ✅ 使用修复后的JWT认证代码

### 步骤3：访问应用
- **应用地址**：http://localhost:8080
- **管理后台**：http://localhost:8080/admin
- **登录账号**：admin / admin123

## 其他解决方案

### 如果想使用80端口

1. **查看占用80端口的进程**：
   ```bash
   netstat -ano | findstr :80
   ```

2. **停止占用80端口的服务**：
   ```bash
   # 停止IIS
   net stop W3SVC
   
   # 停止Apache
   net stop Apache2.4
   ```

3. **使用原始部署**：
   ```bash
   safe-deploy.bat
   ```

### 如果仍有401错误

1. **清除浏览器缓存**：
   - 按 `Ctrl+Shift+Delete`
   - 选择"缓存的图片和文件"
   - 清除

2. **检查浏览器控制台**：
   - 按 `F12` 打开开发者工具
   - 查看Console标签页的错误信息
   - 查看Network标签页的请求头是否包含Authorization

3. **测试JWT认证**：
   ```bash
   test-jwt-fix.bat
   ```

## 技术细节

### JWT认证修复
修改了 `js/admin.js` 中的以下内容：
- `localStorage.getItem('adminToken')` → `localStorage.getItem('authToken')`
- `localStorage.removeItem('adminToken')` → `localStorage.removeItem('authToken')`
- 确保登录和管理页面使用相同的token存储key

### 部署配置
- **docker-compose-no-nginx.yml**：不包含Nginx，应用直接暴露在8080端口
- **deploy-no-nginx.bat**：智能检测端口并部署的脚本

## 验证部署

1. **检查容器状态**：
   ```bash
   docker ps
   ```

2. **查看应用日志**：
   ```bash
   docker logs at-mmbs-app --tail=50
   ```

3. **测试API**：
   ```bash
   curl http://localhost:8080/api/auth/login -X POST -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}"
   ```

## 常见问题

**Q: 为什么不用Nginx？**
A: Nginx主要用于反向代理和静态文件服务。对于开发测试，直接访问应用端口更简单。

**Q: 如何恢复使用Nginx？**
A: 解决80端口冲突后，运行 `safe-deploy.bat` 或 `quick-deploy.bat`。

**Q: 数据会丢失吗？**
A: 不会。MySQL数据保存在Docker卷中，重新部署不会丢失数据。