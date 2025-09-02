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

## 额外修复（第二次）

问题仍然存在的原因是admin.html中有独立的JavaScript代码，这些代码也会在页面加载时执行：

1. **admin.html中的DOMContentLoaded事件**
   - 会调用`loadCategoriesData()`和`showModule('dashboard')`
   - 这些调用发生在admin.js的初始化之前，导致API请求没有认证头

2. **admin.html中的API调用缺少认证头**
   - loadCategoriesData、loadProductsData等函数中的fetch调用没有使用addAuthHeader()

3. **checkAuth函数的路径错误**
   - 使用了'/login'而不是'/login.html'

### 修复方案
- 移除admin.html中的loadCategoriesData()和showModule()调用
- 为admin.html中所有的API调用添加认证头
- 修正登录页面的路径

## 注意事项

- 这是一个临时解决方案
- admin.html中不应该有重复的初始化逻辑，所有初始化应该统一在admin.js中处理
- 长期来看，建议：
  1. 将admin.html中的JavaScript代码移到单独的文件中
  2. 后端实现真正的无状态JWT认证
  3. 或者实现 `/api/users/profile` 端点来获取用户信息