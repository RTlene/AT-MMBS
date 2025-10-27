// 管理后台主要JavaScript文件
// 全局变量
let currentUser = null;
let currentPage = 1;
let pageSize = 10;

// 工具函数
function addAuthHeader(headers = {}) {
    // 优先使用localStorage中的token
    const token = localStorage.getItem('authToken');
    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    } else if (currentUser && currentUser.token) {
        headers['Authorization'] = `Bearer ${currentUser.token}`;
    }
    return headers;
}

function showMessage(message, type = 'info') {
    // 创建消息元素
    const alertDiv = document.createElement('div');
    alertDiv.style.cssText = `
        padding: 15px;
        margin-bottom: 10px;
        border-radius: 5px;
        position: relative;
        animation: slideIn 0.3s ease;
    `;
    
    // 根据类型设置颜色
    const colors = {
        'info': { bg: '#d1ecf1', border: '#bee5eb', text: '#0c5460' },
        'success': { bg: '#d4edda', border: '#c3e6cb', text: '#155724' },
        'warning': { bg: '#fff3cd', border: '#ffeeba', text: '#856404' },
        'danger': { bg: '#f8d7da', border: '#f5c6cb', text: '#721c24' }
    };
    
    const color = colors[type] || colors['info'];
    alertDiv.style.backgroundColor = color.bg;
    alertDiv.style.border = `1px solid ${color.border}`;
    alertDiv.style.color = color.text;
    
    alertDiv.innerHTML = `
        ${message}
        <button onclick="this.parentElement.remove()" style="position: absolute; right: 10px; top: 10px; background: none; border: none; font-size: 20px; cursor: pointer; color: ${color.text};">&times;</button>
    `;
    
    // 添加到消息容器
    const container = document.getElementById('messageContainer');
    if (container) {
        container.appendChild(alertDiv);
        
        // 5秒后自动消失
        setTimeout(() => {
            if (alertDiv.parentNode) {
                alertDiv.style.animation = 'slideOut 0.3s ease';
                setTimeout(() => alertDiv.remove(), 300);
            }
        }, 5000);
    }
}

// 模块显示管理
function showModule(moduleId, event) {
    // 隐藏所有模块
    const modules = document.querySelectorAll('.module-content');
    modules.forEach(module => {
        module.style.display = 'none';
    });

    // 显示选中的模块
    const selectedModule = document.getElementById(moduleId);
    if (selectedModule) {
        selectedModule.style.display = 'block';
    }

    // 更新侧边栏活动状态
    const menuItems = document.querySelectorAll('.sidebar-menu a');
    menuItems.forEach(item => {
        item.classList.remove('active');
    });
    
    // 添加空值检查
    if (event && event.target) {
        event.target.classList.add('active');
    }

    // 加载模块数据
    loadModuleData(moduleId);
}

// 模块数据加载
function loadModuleData(moduleId) {
    switch (moduleId) {
        case 'dashboard':
            loadDashboardData();
            break;
        case 'users':
            loadUsersData();
            break;
        case 'products':
            loadProductsData();
            break;
        case 'categories':
            loadCategoriesData();
            break;
        case 'members':
            loadMembersData();
            break;
        case 'member-levels':
            loadMemberLevelsData();
            break;
        case 'distributor-levels':
            loadDistributorLevelsData();
            break;
        case 'orders':
            loadOrdersData();
            break;
        case 'payments':
            loadPaymentsData();
            break;
        case 'settings':
            loadSettingsData();
            break;
    }
}

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function() {
    // 检查登录状态
    checkLoginStatus();
    
    // 绑定事件监听器
    bindEventListeners();
    
    // 默认显示仪表盘
    showModule('dashboard');
});

// 绑定事件监听器
function bindEventListeners() {
    // 登录表单提交
    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        loginForm.addEventListener('submit', handleLogin);
    }
    
    // 登出按钮
    const logoutBtn = document.querySelector('.logout-btn');
    if (logoutBtn) {
        logoutBtn.addEventListener('click', handleLogout);
    }
}

// 登录状态检查
async function checkLoginStatus() {
    const token = localStorage.getItem('authToken');
    const userInfo = localStorage.getItem('userInfo');
    
    if (!token || !userInfo) {
        window.location.href = '/login.html';
        return;
    }
    
    try {
        // 使用localStorage中保存的用户信息
        const user = JSON.parse(userInfo);
        currentUser = {
            token: token,
            username: user.username,
            ...user
        };
        
        updateUserInfo();
        
        // 可选：验证token是否仍然有效
        // 通过调用一个简单的API来测试
        const testResponse = await fetch('/api/users', {
            headers: addAuthHeader()
        });
        
        if (!testResponse.ok && testResponse.status === 401) {
            console.log('Token验证失败，可能已过期');
            localStorage.removeItem('authToken');
            localStorage.removeItem('userInfo');
            window.location.href = '/login.html';
        }
    } catch (error) {
        console.error('解析用户信息失败:', error);
        localStorage.removeItem('authToken');
        localStorage.removeItem('userInfo');
        window.location.href = '/login.html';
    }
}

// 处理登录
async function handleLogin(event) {
    event.preventDefault();
    
    const formData = new FormData(event.target);
    const username = formData.get('username');
    const password = formData.get('password');
    
    try {
        const response = await fetch('/api/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        });
        
        const result = await response.json();
        if (result.code === 0) {
            localStorage.setItem('authToken', result.data.token);
            currentUser = result.data;
            window.location.href = '/admin.html';
        } else {
            showMessage(result.message || '登录失败', 'danger');
        }
    } catch (error) {
        console.error('登录失败:', error);
        showMessage('登录失败，请检查网络连接', 'danger');
    }
}

// 处理登出
function handleLogout() {
    localStorage.removeItem('authToken');
    localStorage.removeItem('userInfo');
    currentUser = null;
    window.location.href = '/login.html';
}

// 更新用户信息显示
function updateUserInfo() {
    if (currentUser) {
        const userInfoElement = document.querySelector('.user-info');
        if (userInfoElement) {
            userInfoElement.textContent = `欢迎，${currentUser.username}`;
        }
    }
}
