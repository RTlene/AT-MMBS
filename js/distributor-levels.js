// 分销等级管理JavaScript文件

// 加载分销等级数据
async function loadDistributorLevelsData() {
    try {
        const response = await fetch('/api/distributor-levels', {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayDistributorLevels(result.data);
            } else {
                showMessage('加载分销等级数据失败', 'danger');
            }
        } else {
            showMessage('加载分销等级数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载分销等级数据失败:', error);
        showMessage('加载分销等级数据失败', 'danger');
    }
}

// 显示分销等级列表
function displayDistributorLevels(levels) {
    const tbody = document.getElementById('distributorLevelsTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    levels.forEach(level => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${level.name || ''}</td>
            <td>${level.description || ''}</td>
            <td>${level.commission_rate || 0}%</td>
            <td>${level.requirements || ''}</td>
            <td>${level.sort || 0}</td>
            <td>
                <span class="badge ${level.status === 1 ? 'bg-success' : 'bg-secondary'}">
                    ${level.status === 1 ? '启用' : '禁用'}
                </span>
            </td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="editDistributorLevel(${level.id})">编辑</button>
                <button class="btn btn-sm btn-danger" onclick="deleteDistributorLevel(${level.id})">删除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 保存分销等级
async function saveDistributorLevel() {
    const name = document.getElementById('distributorLevelName').value.trim();
    const description = document.getElementById('distributorLevelDescription').value.trim();
    const commissionRate = parseFloat(document.getElementById('distributorLevelCommissionRate').value);
    const requirements = document.getElementById('distributorLevelRequirements').value.trim();
    const sort = parseInt(document.getElementById('distributorLevelSort').value) || 0;
    const status = document.getElementById('distributorLevelStatus').checked ? 1 : 0;
    
    if (!name) {
        showMessage('请填写等级名称', 'warning');
        return;
    }
    
    try {
        const distributorLevelData = {
            name,
            description,
            commission_rate: commissionRate,
            requirements,
            sort,
            status
        };
        
        const response = await fetch('/api/distributor-levels', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(distributorLevelData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分销等级添加成功', 'success');
                clearDistributorLevelForm();
                document.getElementById('addDistributorLevelModal').style.display = 'none';
                loadDistributorLevelsData();
            } else {
                showMessage(result.message || '分销等级添加失败', 'danger');
            }
        } else {
            showMessage('分销等级添加失败', 'danger');
        }
    } catch (error) {
        console.error('添加分销等级失败:', error);
        showMessage('添加分销等级失败', 'danger');
    }
}

// 编辑分销等级
async function editDistributorLevel(levelId) {
    try {
        const response = await fetch(`/api/distributor-levels/${levelId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditDistributorLevelForm(result.data);
                document.getElementById('editDistributorLevelModal').style.display = 'block';
            } else {
                showMessage('获取分销等级信息失败', 'danger');
            }
        } else {
            showMessage('获取分销等级信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取分销等级信息失败:', error);
        showMessage('获取分销等级信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditDistributorLevelForm(level) {
    document.getElementById('editDistributorLevelId').value = level.id;
    document.getElementById('editDistributorLevelName').value = level.name || '';
    document.getElementById('editDistributorLevelDescription').value = level.description || '';
    document.getElementById('editDistributorLevelCommissionRate').value = level.commission_rate || 0;
    document.getElementById('editDistributorLevelRequirements').value = level.requirements || '';
    document.getElementById('editDistributorLevelSort').value = level.sort || 0;
    document.getElementById('editDistributorLevelStatus').checked = level.status === 1;
}

// 更新分销等级
async function updateDistributorLevel() {
    const id = document.getElementById('editDistributorLevelId').value;
    const name = document.getElementById('editDistributorLevelName').value.trim();
    const description = document.getElementById('editDistributorLevelDescription').value.trim();
    const commissionRate = parseFloat(document.getElementById('editDistributorLevelCommissionRate').value);
    const requirements = document.getElementById('editDistributorLevelRequirements').value.trim();
    const sort = parseInt(document.getElementById('editDistributorLevelSort').value) || 0;
    const status = document.getElementById('editDistributorLevelStatus').checked ? 1 : 0;
    
    if (!name) {
        showMessage('请填写等级名称', 'warning');
        return;
    }
    
    try {
        const distributorLevelData = {
            name,
            description,
            commission_rate: commissionRate,
            requirements,
            sort,
            status
        };
        
        const response = await fetch(`/api/distributor-levels/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(distributorLevelData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分销等级更新成功', 'success');
                document.getElementById('editDistributorLevelModal').style.display = 'none';
                loadDistributorLevelsData();
            } else {
                showMessage(result.message || '分销等级更新失败', 'danger');
            }
        } else {
            showMessage('分销等级更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新分销等级失败:', error);
        showMessage('更新分销等级失败', 'danger');
    }
}

// 删除分销等级
async function deleteDistributorLevel(levelId) {
    if (!confirm('确定要删除这个分销等级吗？')) {
        return;
    }
    
    try {
        const response = await fetch(`/api/distributor-levels/${levelId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分销等级删除成功', 'success');
                loadDistributorLevelsData();
            } else {
                showMessage(result.message || '分销等级删除失败', 'danger');
            }
        } else {
            showMessage('分销等级删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除分销等级失败:', error);
        showMessage('删除分销等级失败', 'danger');
    }
}

// 清空分销等级表单
function clearDistributorLevelForm() {
    document.getElementById('distributorLevelName').value = '';
    document.getElementById('distributorLevelDescription').value = '';
    document.getElementById('distributorLevelCommissionRate').value = '';
    document.getElementById('distributorLevelRequirements').value = '';
    document.getElementById('distributorLevelSort').value = '';
    document.getElementById('distributorLevelStatus').checked = true;
}

// 打开添加分销等级模态框
function openAddDistributorLevelModal() {
    clearDistributorLevelForm();
    document.getElementById('addDistributorLevelModal').style.display = 'block';
}

// 关闭添加分销等级模态框
function closeAddDistributorLevelModal() {
    document.getElementById('addDistributorLevelModal').style.display = 'none';
    clearDistributorLevelForm();
}

// 关闭编辑分销等级模态框
function closeEditDistributorLevelModal() {
    document.getElementById('editDistributorLevelModal').style.display = 'none';
}
