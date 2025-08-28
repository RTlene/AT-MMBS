// 仪表盘JavaScript文件

// 加载仪表盘数据
async function loadDashboardData() {
    try {
        // 加载统计数据
        await loadStatistics();
        
        // 加载图表数据
        await loadChartData();
        
        // 加载最近活动
        await loadRecentActivities();
        
    } catch (error) {
        console.error('加载仪表盘数据失败:', error);
        showMessage('加载仪表盘数据失败', 'danger');
    }
}

// 加载统计数据
async function loadStatistics() {
    try {
        const response = await fetch('/api/count', {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayStatistics(result.data);
            }
        }
    } catch (error) {
        console.error('加载统计数据失败:', error);
    }
}

// 显示统计数据
function displayStatistics(data) {
    // 用户统计
    const userCountElement = document.getElementById('userCount');
    if (userCountElement) {
        userCountElement.textContent = data.user_count || 0;
    }
    
    // 商品统计
    const productCountElement = document.getElementById('productCount');
    if (productCountElement) {
        productCountElement.textContent = data.product_count || 0;
    }
    
    // 订单统计
    const orderCountElement = document.getElementById('orderCount');
    if (orderCountElement) {
        orderCountElement.textContent = data.order_count || 0;
    }
    
    // 收入统计
    const revenueElement = document.getElementById('totalRevenue');
    if (revenueElement) {
        revenueElement.textContent = `¥${data.total_revenue || 0}`;
    }
}

// 加载图表数据
async function loadChartData() {
    try {
        // 这里可以加载各种图表数据
        // 比如销售趋势、用户增长等
        console.log('加载图表数据');
    } catch (error) {
        console.error('加载图表数据失败:', error);
    }
}

// 加载最近活动
async function loadRecentActivities() {
    try {
        // 这里可以加载最近的活动记录
        // 比如最近的订单、用户注册等
        console.log('加载最近活动');
    } catch (error) {
        console.error('加载最近活动失败:', error);
    }
}

// 刷新仪表盘
function refreshDashboard() {
    loadDashboardData();
}

// 导出数据
function exportDashboardData() {
    // 这里可以实现数据导出功能
    showMessage('数据导出功能开发中', 'info');
}

// 设置仪表盘
function configureDashboard() {
    // 这里可以实现仪表盘配置功能
    showMessage('仪表盘配置功能开发中', 'info');
}
