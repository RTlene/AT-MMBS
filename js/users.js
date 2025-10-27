// 用户管理JavaScript文件

// 加载用户数据
async function loadUsersData() {
    try {
        const response = await fetch(`/api/users?page=${currentPage}&size=${pageSize}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayUsers(result.data.users);
                updateUserPagination(result.data.total, result.data.page, result.data.size);
            } else {
                showMessage('加载用户数据失败', 'danger');
            }
        } else {
            showMessage('加载用户数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载用户数据失败:', error);
        showMessage('加载用户数据失败', 'danger');
    }
}

// 显示用户列表
function displayUsers(users) {
    const tbody = document.getElementById('usersTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    users.forEach(user => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${user.username || ''}</td>
            <td>${user.email || ''}</td>
            <td>${user.phone || ''}</td>
            <td>${user.role || 'user'}</td>
            <td>
                <span class="badge ${user.status === 1 ? 'bg-success' : 'bg-secondary'}">
                    ${user.status === 1 ? '启用' : '禁用'}
                </span>
            </td>
            <td>${user.create_time ? new Date(user.create_time).toLocaleDateString() : '未知'}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="editUser('${user.id}')">编辑</button>
                <button class="btn btn-sm btn-warning" onclick="toggleUserStatus('${user.id}')">
                    ${user.status === 1 ? '禁用' : '启用'}
                </button>
                <button class="btn btn-sm btn-danger" onclick="deleteUser('${user.id}')">删除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 更新用户分页
function updateUserPagination(total, currentPage, pageSize) {
    const pagination = document.getElementById('usersPagination');
    if (!pagination) return;
    
    const totalPages = Math.ceil(total / pageSize);
    let paginationHTML = '';
    
    if (totalPages > 1) {
        paginationHTML += `
            <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeUserPage(${currentPage - 1})">上一页</a>
            </li>
        `;
        
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `
                <li class="page-item ${i === currentPage ? 'active' : ''}">
                    <a class="page-link" href="#" onclick="changeUserPage(${i})">${i}</a>
                </li>
            `;
        }
        
        paginationHTML += `
            <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeUserPage(${currentPage + 1})">下一页</a>
            </li>
        `;
    }
    
    pagination.innerHTML = paginationHTML;
}

// 切换用户页面
function changeUserPage(page) {
    if (page < 1) return;
    currentPage = page;
    loadUsersData();
}

// 保存用户
async function saveUser() {
    const username = document.getElementById('userUsername').value.trim();
    const email = document.getElementById('userEmail').value.trim();
    const phone = document.getElementById('userPhone').value.trim();
    const password = document.getElementById('userPassword').value;
    const role = document.getElementById('userRole').value;
    const status = document.getElementById('userStatus').checked ? 1 : 0;
    
    if (!username || !password) {
        showMessage('请填写用户名和密码', 'warning');
        return;
    }
    
    try {
        const userData = {
            username,
            email,
            phone,
            password,
            role,
            status
        };
        
        const response = await fetch('/api/users', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(userData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('用户添加成功', 'success');
                clearUserForm();
                document.getElementById('addUserModal').style.display = 'none';
                loadUsersData();
            } else {
                showMessage(result.message || '用户添加失败', 'danger');
            }
        } else {
            showMessage('用户添加失败', 'danger');
        }
    } catch (error) {
        console.error('添加用户失败:', error);
        showMessage('添加用户失败', 'danger');
    }
}

// 编辑用户
async function editUser(userId) {
    try {
        const response = await fetch(`/api/users/${userId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditUserForm(result.data);
                document.getElementById('editUserModal').style.display = 'block';
            } else {
                showMessage('获取用户信息失败', 'danger');
            }
        } else {
            showMessage('获取用户信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取用户信息失败:', error);
        showMessage('获取用户信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditUserForm(user) {
    document.getElementById('editUserId').value = user.id;
    document.getElementById('editUserUsername').value = user.username || '';
    document.getElementById('editUserEmail').value = user.email || '';
    document.getElementById('editUserPhone').value = user.phone || '';
    document.getElementById('editUserRole').value = user.role || 'user';
    document.getElementById('editUserStatus').checked = user.status === 1;
}

// 更新用户
async function updateUser() {
    const id = document.getElementById('editUserId').value;
    const username = document.getElementById('editUserUsername').value.trim();
    const email = document.getElementById('editUserEmail').value.trim();
    const phone = document.getElementById('editUserPhone').value.trim();
    const role = document.getElementById('editUserRole').value;
    const status = document.getElementById('editUserStatus').checked ? 1 : 0;
    
    if (!username) {
        showMessage('请填写用户名', 'warning');
        return;
    }
    
    try {
        const userData = {
            username,
            email,
            phone,
            role,
            status
        };
        
        const response = await fetch(`/api/users/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(userData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('用户更新成功', 'success');
                document.getElementById('editUserModal').style.display = 'none';
                loadUsersData();
            } else {
                showMessage(result.message || '用户更新失败', 'danger');
            }
        } else {
            showMessage('用户更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新用户失败:', error);
        showMessage('更新用户失败', 'danger');
    }
}

// 删除用户
async function deleteUser(userId) {
    if (!confirm('确定要删除这个用户吗？')) {
        return;
    }
    
    try {
        const response = await fetch(`/api/users/${userId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('用户删除成功', 'success');
                loadUsersData();
            } else {
                showMessage(result.message || '用户删除失败', 'danger');
            }
        } else {
            showMessage('用户删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除用户失败:', error);
        showMessage('删除用户失败', 'danger');
    }
}

// 切换用户状态
async function toggleUserStatus(userId) {
    try {
        const response = await fetch(`/api/users/${userId}/status`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify({ status: 'toggle' })
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('用户状态更新成功', 'success');
                loadUsersData();
            } else {
                showMessage(result.message || '用户状态更新失败', 'danger');
            }
        } else {
            showMessage('用户状态更新失败', 'danger');
        }
    } catch (error) {
        console.error('切换用户状态失败:', error);
        showMessage('切换用户状态失败', 'danger');
    }
}

// 清空用户表单
function clearUserForm() {
    document.getElementById('userUsername').value = '';
    document.getElementById('userEmail').value = '';
    document.getElementById('userPhone').value = '';
    document.getElementById('userPassword').value = '';
    document.getElementById('userRole').value = 'user';
    document.getElementById('userStatus').checked = true;
}

// 打开添加用户模态框
function openAddUserModal() {
    clearUserForm();
    document.getElementById('addUserModal').style.display = 'block';
}

// 关闭添加用户模态框
function closeAddUserModal() {
    document.getElementById('addUserModal').style.display = 'none';
    clearUserForm();
}

// 关闭编辑用户模态框
function closeEditUserModal() {
    document.getElementById('editUserModal').style.display = 'none';
}
