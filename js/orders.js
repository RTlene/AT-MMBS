// 订单管理JavaScript文件

// 加载订单数据
async function loadOrdersData() {
    try {
        const response = await fetch(`/api/orders?page=${currentPage}&size=${pageSize}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayOrders(result.data.orders);
                updateOrderPagination(result.data.total, result.data.page, result.data.size);
            } else {
                showMessage('加载订单数据失败', 'danger');
            }
        } else {
            showMessage('加载订单数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载订单数据失败:', error);
        showMessage('加载订单数据失败', 'danger');
    }
}

// 显示订单列表
function displayOrders(orders) {
    const tbody = document.getElementById('ordersTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    orders.forEach(order => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${order.order_no || ''}</td>
            <td>${order.user_id || ''}</td>
            <td>¥${order.total_amount || 0}</td>
            <td>
                <span class="badge ${getOrderStatusBadge(order.status)}">
                    ${getOrderStatusText(order.status)}
                </span>
            </td>
            <td>${order.create_time ? new Date(order.create_time).toLocaleDateString() : '未知'}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="viewOrder('${order.id}')">查看</button>
                <button class="btn btn-sm btn-warning" onclick="updateOrderStatus('${order.id}')">更新状态</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 获取订单状态徽章样式
function getOrderStatusBadge(status) {
    switch (status) {
        case 'pending': return 'bg-warning';
        case 'paid': return 'bg-success';
        case 'shipped': return 'bg-info';
        case 'delivered': return 'bg-primary';
        case 'cancelled': return 'bg-danger';
        default: return 'bg-secondary';
    }
}

// 获取订单状态文本
function getOrderStatusText(status) {
    switch (status) {
        case 'pending': return '待付款';
        case 'paid': return '已付款';
        case 'shipped': return '已发货';
        case 'delivered': return '已送达';
        case 'cancelled': return '已取消';
        default: return '未知';
    }
}

// 更新订单分页
function updateOrderPagination(total, currentPage, pageSize) {
    const pagination = document.getElementById('ordersPagination');
    if (!pagination) return;
    
    const totalPages = Math.ceil(total / pageSize);
    let paginationHTML = '';
    
    if (totalPages > 1) {
        paginationHTML += `
            <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeOrderPage(${currentPage - 1})">上一页</a>
            </li>
        `;
        
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `
                <li class="page-item ${i === currentPage ? 'active' : ''}">
                    <a class="page-link" href="#" onclick="changeOrderPage(${i})">${i}</a>
                </li>
            `;
        }
        
        paginationHTML += `
            <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeOrderPage(${currentPage + 1})">下一页</a>
            </li>
        `;
    }
    
    pagination.innerHTML = paginationHTML;
}

// 切换订单页面
function changeOrderPage(page) {
    if (page < 1) return;
    currentPage = page;
    loadOrdersData();
}

// 查看订单详情
async function viewOrder(orderId) {
    try {
        const response = await fetch(`/api/orders/${orderId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillOrderDetailForm(result.data);
                document.getElementById('orderDetailModal').style.display = 'block';
            } else {
                showMessage('获取订单信息失败', 'danger');
            }
        } else {
            showMessage('获取订单信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取订单信息失败:', error);
        showMessage('获取订单信息失败', 'danger');
    }
}

// 填充订单详情表单
function fillOrderDetailForm(order) {
    document.getElementById('orderDetailId').value = order.id;
    document.getElementById('orderDetailNo').value = order.order_no || '';
    document.getElementById('orderDetailUserId').value = order.user_id || '';
    document.getElementById('orderDetailAmount').value = order.total_amount || '';
    document.getElementById('orderDetailStatus').value = order.status || '';
    document.getElementById('orderDetailCreateTime').value = order.create_time || '';
}

// 更新订单状态
async function updateOrderStatus(orderId) {
    const newStatus = prompt('请输入新状态 (pending/paid/shipped/delivered/cancelled):');
    if (!newStatus) return;
    
    try {
        const response = await fetch(`/api/orders/status`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify({ id: orderId, status: newStatus })
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('订单状态更新成功', 'success');
                loadOrdersData();
            } else {
                showMessage(result.message || '订单状态更新失败', 'danger');
            }
        } else {
            showMessage('订单状态更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新订单状态失败:', error);
        showMessage('更新订单状态失败', 'danger');
    }
}

// 关闭订单详情模态框
function closeOrderDetailModal() {
    document.getElementById('orderDetailModal').style.display = 'none';
}
