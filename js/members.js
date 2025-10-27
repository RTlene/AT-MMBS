// 会员管理JavaScript文件

// 加载会员数据
async function loadMembersData() {
    try {
        const response = await fetch(`/api/members?page=${currentPage}&size=${pageSize}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayMembers(result.data.members);
                updateMemberPagination(result.data.total, result.data.page, result.data.size);
            } else {
                showMessage('加载会员数据失败', 'danger');
            }
        } else {
            showMessage('加载会员数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载会员数据失败:', error);
        showMessage('加载会员数据失败', 'danger');
    }
}

// 显示会员列表
function displayMembers(members) {
    const tbody = document.getElementById('membersTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    members.forEach(member => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${member.username || ''}</td>
            <td>${member.email || ''}</td>
            <td>${member.phone || ''}</td>
            <td>${member.level_name || '普通会员'}</td>
            <td>${member.points || 0}</td>
            <td>
                <span class="badge ${member.status === 1 ? 'bg-success' : 'bg-secondary'}">
                    ${member.status === 1 ? '正常' : '禁用'}
                </span>
            </td>
            <td>${member.create_time ? new Date(member.create_time).toLocaleDateString() : '未知'}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="viewMember('${member.id}')">查看</button>
                <button class="btn btn-sm btn-warning" onclick="editMember('${member.id}')">编辑</button>
                <button class="btn btn-sm btn-danger" onclick="deleteMember('${member.id}')">删除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 更新会员分页
function updateMemberPagination(total, currentPage, pageSize) {
    const pagination = document.getElementById('membersPagination');
    if (!pagination) return;
    
    const totalPages = Math.ceil(total / pageSize);
    let paginationHTML = '';
    
    if (totalPages > 1) {
        paginationHTML += `
            <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeMemberPage(${currentPage - 1})">上一页</a>
            </li>
        `;
        
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `
                <li class="page-item ${i === currentPage ? 'active' : ''}">
                    <a class="page-link" href="#" onclick="changeMemberPage(${i})">${i}</a>
                </li>
            `;
        }
        
        paginationHTML += `
            <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changeMemberPage(${currentPage + 1})">下一页</a>
            </li>
        `;
    }
    
    pagination.innerHTML = paginationHTML;
}

// 切换会员页面
function changeMemberPage(page) {
    if (page < 1) return;
    currentPage = page;
    loadMembersData();
}

// 查看会员详情
async function viewMember(memberId) {
    try {
        const response = await fetch(`/api/members/${memberId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillMemberDetailForm(result.data);
                document.getElementById('memberDetailModal').style.display = 'block';
            } else {
                showMessage('获取会员信息失败', 'danger');
            }
        } else {
            showMessage('获取会员信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取会员信息失败:', error);
        showMessage('获取会员信息失败', 'danger');
    }
}

// 填充会员详情表单
function fillMemberDetailForm(member) {
    document.getElementById('memberDetailId').value = member.id;
    document.getElementById('memberDetailUsername').value = member.username || '';
    document.getElementById('memberDetailEmail').value = member.email || '';
    document.getElementById('memberDetailPhone').value = member.phone || '';
    document.getElementById('memberDetailLevel').value = member.level_name || '';
    document.getElementById('memberDetailPoints').value = member.points || 0;
    document.getElementById('memberDetailStatus').value = member.status || 0;
}

// 编辑会员
async function editMember(memberId) {
    try {
        const response = await fetch(`/api/members/${memberId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditMemberForm(result.data);
                document.getElementById('editMemberModal').style.display = 'block';
            } else {
                showMessage('获取会员信息失败', 'danger');
            }
        } else {
            showMessage('获取会员信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取会员信息失败:', error);
        showMessage('获取会员信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditMemberForm(member) {
    document.getElementById('editMemberId').value = member.id;
    document.getElementById('editMemberUsername').value = member.username || '';
    document.getElementById('editMemberEmail').value = member.email || '';
    document.getElementById('editMemberPhone').value = member.phone || '';
    document.getElementById('editMemberLevel').value = member.level_id || '';
    document.getElementById('editMemberPoints').value = member.points || 0;
    document.getElementById('editMemberStatus').checked = member.status === 1;
}

// 更新会员
async function updateMember() {
    const id = document.getElementById('editMemberId').value;
    const username = document.getElementById('editMemberUsername').value.trim();
    const email = document.getElementById('editMemberEmail').value.trim();
    const phone = document.getElementById('editMemberPhone').value.trim();
    const levelId = document.getElementById('editMemberLevel').value;
    const points = parseInt(document.getElementById('editMemberPoints').value) || 0;
    const status = document.getElementById('editMemberStatus').checked ? 1 : 0;
    
    if (!username) {
        showMessage('请填写用户名', 'warning');
        return;
    }
    
    try {
        const memberData = {
            username,
            email,
            phone,
            level_id: levelId,
            points,
            status
        };
        
        const response = await fetch(`/api/members/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(memberData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('会员更新成功', 'success');
                document.getElementById('editMemberModal').style.display = 'none';
                loadMembersData();
            } else {
                showMessage(result.message || '会员更新失败', 'danger');
            }
        } else {
            showMessage('会员更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新会员失败:', error);
        showMessage('更新会员失败', 'danger');
    }
}

// 删除会员
async function deleteMember(memberId) {
    if (!confirm('确定要删除这个会员吗？')) {
        return;
    }
    
    try {
        const response = await fetch(`/api/members/${memberId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('会员删除成功', 'success');
                loadMembersData();
            } else {
                showMessage(result.message || '会员删除失败', 'danger');
            }
        } else {
            showMessage('会员删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除会员失败:', error);
        showMessage('删除会员失败', 'danger');
    }
}

// 关闭会员详情模态框
function closeMemberDetailModal() {
    document.getElementById('memberDetailModal').style.display = 'none';
}

// 关闭编辑会员模态框
function closeEditMemberModal() {
    document.getElementById('editMemberModal').style.display = 'none';
}
