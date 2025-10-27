// 管理后台主要JavaScript文件
// 全局变量
let currentUser = null;
let currentPage = 1;
let pageSize = 10;

// 工具函数
function addAuthHeader(headers = {}) {
    // 优先使用currentUser的token
    if (currentUser && currentUser.token) {
        headers['Authorization'] = `Bearer ${currentUser.token}`;
    } else {
        // 如果currentUser未设置，从localStorage获取
        const token = localStorage.getItem('authToken');
        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }
    }
    return headers;
}

function showMessage(message, type = 'info') {
    const alertDiv = document.createElement('div');
    alertDiv.className = `alert alert-${type} alert-dismissible fade show`;
    alertDiv.innerHTML = `
        ${message}
        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
    `;
    
    const container = document.querySelector('.container-fluid');
    container.insertBefore(alertDiv, container.firstChild);
    
    setTimeout(() => {
        if (alertDiv.parentNode) {
            alertDiv.remove();
        }
    }, 5000);
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
            // 先加载分类数据，然后加载商品数据
            loadCategories().then(() => {
                loadProductsData();
            });
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
document.addEventListener('DOMContentLoaded', async function() {
    // 先检查登录状态，等待完成
    await checkLoginStatus();
    
    // 只有登录验证通过后才执行以下操作
    if (currentUser && currentUser.token) {
        // 绑定事件监听器
        bindEventListeners();
        
        // 加载分类数据（供其他模块使用）
        loadCategories();
        
        // 默认显示仪表盘
        showModule('dashboard');
    }
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
    if (!token) {
        window.location.href = '/login.html';
        return;
    }
    
    // 设置currentUser对象以便addAuthHeader可以使用
    currentUser = {
        token: token,
        info: JSON.parse(localStorage.getItem('userInfo') || '{}')
    };
    
    // 直接使用本地存储的用户信息
    try {
        // 简单验证token格式（JWT应该有三部分）
        const tokenParts = token.split('.');
        if (tokenParts.length !== 3) {
            throw new Error('Invalid token format');
        }
        
        // 更新用户信息显示
        updateUserInfo();
        
        // 暂时跳过token验证，因为会导致循环问题
        // 如果token真的无效，后续的API调用会返回401
        console.log('User logged in with token:', token.substring(0, 20) + '...');
        
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
    if (currentUser && currentUser.info) {
        const userInfoElement = document.querySelector('.user-info');
        if (userInfoElement) {
            const username = currentUser.info.username || currentUser.info.name || '用户';
            userInfoElement.textContent = `欢迎，${username}`;
        }
    }
}
