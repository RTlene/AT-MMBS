# 登录跳转问题修复说明

## 问题描述
用户登录后访问主页（admin.html）时，会被重定向回登录页面。

## 问题原因

1. **异步执行顺序问题**
   - 页面初始化时，`checkLoginStatus()` 是异步执行的
   - 但 `loadCategories()` 等函数在 `currentUser` 设置之前就被调用
   - 导致API请求没有携带认证头，返回401错误

2. **不存在的API端点**
   - `checkLoginStatus` 函数尝试调用 `/api/users/profile` 端点
   - 但后端并没有实现这个端点，导致404错误
   - 错误处理逻辑会清除token并跳转到登录页

3. **认证头添加逻辑问题**
   - `addAuthHeader` 函数依赖 `currentUser` 对象
   - 但在某些情况下，`currentUser` 还未初始化就被调用

## 解决方案

### 1. 修改页面初始化逻辑
```javascript
// 改为async函数，等待登录验证完成
document.addEventListener('DOMContentLoaded', async function() {
    await checkLoginStatus();
    
    // 只有登录验证通过后才执行其他操作
    if (currentUser && currentUser.token) {
        bindEventListeners();
        loadCategories();
        showModule('dashboard');
    }
});
```

### 2. 改进 addAuthHeader 函数
```javascript
function addAuthHeader(headers = {}) {
    // 优先使用currentUser的token
    if (currentUser && currentUser.token) {
        headers['Authorization'] = `Bearer ${currentUser.token}`;
    } else {
        // fallback: 从localStorage获取
        const token = localStorage.getItem('authToken');
        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }
    }
    return headers;
}
```

### 3. 修改 checkLoginStatus 函数
- 移除对不存在的 `/api/users/profile` 端点的调用
- 使用本地存储的用户信息
- 通过调用其他需要认证的API（如 `/api/users`）来验证token有效性

### 4. 修复 updateUserInfo 函数
- 正确访问 `currentUser.info` 中的用户名信息

## 测试建议

1. 清除浏览器缓存和localStorage
2. 重新登录系统
3. 刷新页面，确认不会被重定向到登录页
4. 检查网络请求，确认API调用都带有正确的Authorization头

## 注意事项

- 这是一个临时解决方案
- 长期来看，建议后端实现真正的无状态JWT认证
- 或者实现 `/api/users/profile` 端点来获取用户信息