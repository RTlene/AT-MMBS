# 401错误快速解决

## 问题
登录成功但访问页面返回401错误

## 原因
**Docker容器重启后，内存中的登录信息丢失了**

## 立即解决（30秒）

### 方法1：运行快速重登录脚本
```bash
quick-relogin.bat
```

### 方法2：手动重新登录
1. 清除浏览器缓存：`Ctrl+Shift+Delete`
2. 访问：http://localhost:8080/login.html
3. 登录：admin / admin123

## 为什么会这样？

当前系统的认证机制：
- ❌ Token保存在**内存**中
- ❌ Docker容器重启 = 内存清空
- ❌ 之前的登录失效

## 永久解决方案

需要修改代码，使用：
- ✅ JWT Token（无状态）
- ✅ Redis存储Token
- ✅ 数据库存储Token

详见 `AUTH_ISSUE_SOLUTION.md`

## 记住

**每次重启容器后都需要重新登录！**