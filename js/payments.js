// 支付管理JavaScript文件

// 加载支付数据
async function loadPaymentsData() {
    try {
        const response = await fetch(`/api/payments?page=${currentPage}&size=${pageSize}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayPayments(result.data.payments);
                updatePaymentPagination(result.data.total, result.data.page, result.data.size);
            } else {
                showMessage('加载支付数据失败', 'danger');
            }
        } else {
            showMessage('加载支付数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载支付数据失败:', error);
        showMessage('加载支付数据失败', 'danger');
    }
}

// 显示支付列表
function displayPayments(payments) {
    const tbody = document.getElementById('paymentsTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    payments.forEach(payment => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${payment.payment_no || ''}</td>
            <td>${payment.order_id || ''}</td>
            <td>¥${payment.amount || 0}</td>
            <td>${payment.payment_method || ''}</td>
            <td>
                <span class="badge ${getPaymentStatusBadge(payment.status)}">
                    ${getPaymentStatusText(payment.status)}
                </span>
            </td>
            <td>${payment.create_time ? new Date(payment.create_time).toLocaleDateString() : '未知'}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="viewPayment('${payment.id}')">查看</button>
                <button class="btn btn-sm btn-warning" onclick="updatePaymentStatus('${payment.id}')">更新状态</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 获取支付状态徽章样式
function getPaymentStatusBadge(status) {
    switch (status) {
        case 'pending': return 'bg-warning';
        case 'success': return 'bg-success';
        case 'failed': return 'bg-danger';
        case 'refunded': return 'bg-info';
        default: return 'bg-secondary';
    }
}

// 获取支付状态文本
function getPaymentStatusText(status) {
    switch (status) {
        case 'pending': return '待支付';
        case 'success': return '支付成功';
        case 'failed': return '支付失败';
        case 'refunded': return '已退款';
        default: return '未知';
    }
}

// 更新支付分页
function updatePaymentPagination(total, currentPage, pageSize) {
    const pagination = document.getElementById('paymentsPagination');
    if (!pagination) return;
    
    const totalPages = Math.ceil(total / pageSize);
    let paginationHTML = '';
    
    if (totalPages > 1) {
        paginationHTML += `
            <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changePaymentPage(${currentPage - 1})">上一页</a>
            </li>
        `;
        
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `
                <li class="page-item ${i === currentPage ? 'active' : ''}">
                    <a class="page-link" href="#" onclick="changePaymentPage(${i})">${i}</a>
                </li>
            `;
        }
        
        paginationHTML += `
            <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changePaymentPage(${currentPage + 1})">下一页</a>
            </li>
        `;
    }
    
    pagination.innerHTML = paginationHTML;
}

// 切换支付页面
function changePaymentPage(page) {
    if (page < 1) return;
    currentPage = page;
    loadPaymentsData();
}

// 查看支付详情
async function viewPayment(paymentId) {
    try {
        const response = await fetch(`/api/payments/${paymentId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillPaymentDetailForm(result.data);
                document.getElementById('paymentDetailModal').style.display = 'block';
            } else {
                showMessage('获取支付信息失败', 'danger');
            }
        } else {
            showMessage('获取支付信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取支付信息失败:', error);
        showMessage('获取支付信息失败', 'danger');
    }
}

// 填充支付详情表单
function fillPaymentDetailForm(payment) {
    document.getElementById('paymentDetailId').value = payment.id;
    document.getElementById('paymentDetailNo').value = payment.payment_no || '';
    document.getElementById('paymentDetailOrderId').value = payment.order_id || '';
    document.getElementById('paymentDetailAmount').value = payment.amount || '';
    document.getElementById('paymentDetailMethod').value = payment.payment_method || '';
    document.getElementById('paymentDetailStatus').value = payment.status || '';
    document.getElementById('paymentDetailCreateTime').value = payment.create_time || '';
}

// 更新支付状态
async function updatePaymentStatus(paymentId) {
    const newStatus = prompt('请输入新状态 (pending/success/failed/refunded):');
    if (!newStatus) return;
    
    try {
        const response = await fetch(`/api/payments/status`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify({ id: paymentId, status: newStatus })
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('支付状态更新成功', 'success');
                loadPaymentsData();
            } else {
                showMessage(result.message || '支付状态更新失败', 'danger');
            }
        } else {
            showMessage('支付状态更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新支付状态失败:', error);
        showMessage('更新支付状态失败', 'danger');
    }
}

// 关闭支付详情模态框
function closePaymentDetailModal() {
    document.getElementById('paymentDetailModal').style.display = 'none';
}
