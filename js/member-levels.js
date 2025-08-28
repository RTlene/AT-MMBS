// 会员等级管理JavaScript文件

// 加载会员等级数据
async function loadMemberLevelsData() {
    try {
        const response = await fetch('/api/member-levels', {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayMemberLevels(result.data);
            } else {
                showMessage('加载会员等级数据失败', 'danger');
            }
        } else {
            showMessage('加载会员等级数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载会员等级数据失败:', error);
        showMessage('加载会员等级数据失败', 'danger');
    }
}

// 显示会员等级列表
function displayMemberLevels(levels) {
    const tbody = document.getElementById('memberLevelsTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    levels.forEach(level => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${level.name || ''}</td>
            <td>${level.description || ''}</td>
            <td>${level.discount || 0}%</td>
            <td>${level.points_rate || 0}x</td>
            <td>${level.upgrade_conditions || ''}</td>
            <td>
                <span class="badge ${level.auto_upgrade ? 'bg-success' : 'bg-secondary'}">
                    ${level.auto_upgrade ? '是' : '否'}
                </span>
            </td>
            <td>${level.sort || 0}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="editMemberLevel(${level.id})">编辑</button>
                <button class="btn btn-sm btn-danger" onclick="deleteMemberLevel(${level.id})">删除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 保存会员等级
async function saveMemberLevel() {
    const name = document.getElementById('memberLevelName').value.trim();
    const description = document.getElementById('memberLevelDescription').value.trim();
    const discount = parseFloat(document.getElementById('memberLevelDiscount').value);
    const pointsRate = parseFloat(document.getElementById('memberLevelPointsRate').value);
    const upgradeConditions = document.getElementById('memberLevelUpgradeConditions').value.trim();
    const autoUpgrade = document.getElementById('memberLevelAutoUpgrade').checked;
    const sort = parseInt(document.getElementById('memberLevelSort').value);
    
    if (!name) {
        showMessage('请填写等级名称', 'warning');
        return;
    }
    
    try {
        const memberLevelData = {
            name,
            description,
            discount,
            points_rate: pointsRate,
            upgrade_conditions: upgradeConditions,
            auto_upgrade: autoUpgrade,
            sort
        };
        
        const response = await fetch('/api/member-levels', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(memberLevelData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('会员等级添加成功', 'success');
                clearMemberLevelForm();
                document.getElementById('addMemberLevelModal').style.display = 'none';
                loadMemberLevelsData();
            } else {
                showMessage(result.message || '会员等级添加失败', 'danger');
            }
        } else {
            showMessage('会员等级添加失败', 'danger');
        }
    } catch (error) {
        console.error('添加会员等级失败:', error);
        showMessage('添加会员等级失败', 'danger');
    }
}

// 编辑会员等级
async function editMemberLevel(levelId) {
    try {
        const response = await fetch(`/api/member-levels/${levelId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditMemberLevelForm(result.data);
                document.getElementById('editMemberLevelModal').style.display = 'block';
            } else {
                showMessage('获取会员等级信息失败', 'danger');
            }
        } else {
            showMessage('获取会员等级信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取会员等级信息失败:', error);
        showMessage('获取会员等级信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditMemberLevelForm(level) {
    document.getElementById('editMemberLevelId').value = level.id;
    document.getElementById('editMemberLevelName').value = level.name || '';
    document.getElementById('editMemberLevelDescription').value = level.description || '';
    document.getElementById('editMemberLevelDiscount').value = level.discount || 0;
    document.getElementById('editMemberLevelPointsRate').value = level.points_rate || 0;
    document.getElementById('editMemberLevelUpgradeConditions').value = level.upgrade_conditions || '';
    document.getElementById('editMemberLevelAutoUpgrade').checked = level.auto_upgrade || false;
    document.getElementById('editMemberLevelSort').value = level.sort || 0;
}

// 更新会员等级
async function updateMemberLevel() {
    const id = document.getElementById('editMemberLevelId').value;
    const name = document.getElementById('editMemberLevelName').value.trim();
    const description = document.getElementById('editMemberLevelDescription').value.trim();
    const discount = parseFloat(document.getElementById('editMemberLevelDiscount').value);
    const pointsRate = parseFloat(document.getElementById('editMemberLevelPointsRate').value);
    const upgradeConditions = document.getElementById('editMemberLevelUpgradeConditions').value.trim();
    const autoUpgrade = document.getElementById('editMemberLevelAutoUpgrade').checked;
    const sort = parseInt(document.getElementById('editMemberLevelSort').value);
    
    if (!name) {
        showMessage('请填写等级名称', 'warning');
        return;
    }
    
    try {
        const memberLevelData = {
            name,
            description,
            discount,
            points_rate: pointsRate,
            upgrade_conditions: upgradeConditions,
            auto_upgrade: autoUpgrade,
            sort
        };
        
        const response = await fetch(`/api/member-levels/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(memberLevelData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('会员等级更新成功', 'success');
                document.getElementById('editMemberLevelModal').style.display = 'none';
                loadMemberLevelsData();
            } else {
                showMessage(result.message || '会员等级更新失败', 'danger');
            }
        } else {
            showMessage('会员等级更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新会员等级失败:', error);
        showMessage('更新会员等级失败', 'danger');
    }
}

// 删除会员等级
async function deleteMemberLevel(levelId) {
    if (!confirm('确定要删除这个会员等级吗？')) {
        return;
    }
    
    try {
        const response = await fetch(`/api/member-levels/${levelId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('会员等级删除成功', 'success');
                loadMemberLevelsData();
            } else {
                showMessage(result.message || '会员等级删除失败', 'danger');
            }
        } else {
            showMessage('会员等级删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除会员等级失败:', error);
        showMessage('删除会员等级失败', 'danger');
    }
}

// 清空会员等级表单
function clearMemberLevelForm() {
    document.getElementById('memberLevelName').value = '';
    document.getElementById('memberLevelDescription').value = '';
    document.getElementById('memberLevelDiscount').value = '';
    document.getElementById('memberLevelPointsRate').value = '';
    document.getElementById('memberLevelUpgradeConditions').value = '';
    document.getElementById('memberLevelAutoUpgrade').checked = false;
    document.getElementById('memberLevelSort').value = '';
}

// 打开添加会员等级模态框
function openAddMemberLevelModal() {
    clearMemberLevelForm();
    document.getElementById('addMemberLevelModal').style.display = 'block';
}

// 关闭添加会员等级模态框
function closeAddMemberLevelModal() {
    document.getElementById('addMemberLevelModal').style.display = 'none';
    clearMemberLevelForm();
}

// 关闭编辑会员等级模态框
function closeEditMemberLevelModal() {
    document.getElementById('editMemberLevelModal').style.display = 'none';
}
