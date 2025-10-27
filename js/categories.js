// 分类管理JavaScript文件

// 加载分类数据
async function loadCategoriesData() {
    try {
        const response = await fetch('/api/categories', {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0 && result.data) {
                displayCategories(result.data);
            } else {
                showMessage('加载分类数据失败', 'danger');
            }
        } else {
            showMessage('加载分类数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载分类数据失败:', error);
        showMessage('加载分类数据失败', 'danger');
    }
}

// 显示分类列表
function displayCategories(categories) {
    const tbody = document.getElementById('categoriesTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    categories.forEach(category => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${category.name || ''}</td>
            <td>${category.description || ''}</td>
            <td>${category.parent_id || '无'}</td>
            <td>${category.sort || 0}</td>
            <td>
                <span class="badge ${category.status === 1 ? 'bg-success' : 'bg-secondary'}">
                    ${category.status === 1 ? '启用' : '禁用'}
                </span>
            </td>
            <td>${category.create_time ? new Date(category.create_time).toLocaleDateString() : '未知'}</td>
            <td>
                <button class="btn btn-sm btn-primary" onclick="editCategory('${category.id}')">编辑</button>
                <button class="btn btn-sm btn-danger" onclick="deleteCategory('${category.id}')">删除</button>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 保存分类
async function saveCategory() {
    const name = document.getElementById('categoryName').value.trim();
    const description = document.getElementById('categoryDescription').value.trim();
    const parentId = document.getElementById('categoryParent').value;
    const sort = parseInt(document.getElementById('categorySort').value) || 0;
    const status = document.getElementById('categoryStatus').checked ? 1 : 0;
    
    if (!name) {
        showMessage('请填写分类名称', 'warning');
        return;
    }
    
    try {
        const categoryData = {
            name,
            description,
            parent_id: parentId || null,
            sort,
            status
        };
        
        const response = await fetch('/api/categories', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(categoryData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分类添加成功', 'success');
                clearCategoryForm();
                document.getElementById('addCategoryModal').style.display = 'none';
                loadCategoriesData();
            } else {
                showMessage(result.message || '分类添加失败', 'danger');
            }
        } else {
            showMessage('分类添加失败', 'danger');
        }
    } catch (error) {
        console.error('添加分类失败:', error);
        showMessage('添加分类失败', 'danger');
    }
}

// 编辑分类
async function editCategory(categoryId) {
    try {
        const response = await fetch(`/api/categories/${categoryId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditCategoryForm(result.data);
                document.getElementById('editCategoryModal').style.display = 'block';
            } else {
                showMessage('获取分类信息失败', 'danger');
            }
        } else {
            showMessage('获取分类信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取分类信息失败:', error);
        showMessage('获取分类信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditCategoryForm(category) {
    document.getElementById('editCategoryId').value = category.id;
    document.getElementById('editCategoryName').value = category.name || '';
    document.getElementById('editCategoryDescription').value = category.description || '';
    document.getElementById('editCategoryParent').value = category.parent_id || '';
    document.getElementById('editCategorySort').value = category.sort || 0;
    document.getElementById('editCategoryStatus').checked = category.status === 1;
}

// 更新分类
async function updateCategory() {
    const id = document.getElementById('editCategoryId').value;
    const name = document.getElementById('editCategoryName').value.trim();
    const description = document.getElementById('editCategoryDescription').value.trim();
    const parentId = document.getElementById('editCategoryParent').value;
    const sort = parseInt(document.getElementById('editCategorySort').value) || 0;
    const status = document.getElementById('editCategoryStatus').checked ? 1 : 0;
    
    if (!name) {
        showMessage('请填写分类名称', 'warning');
        return;
    }
    
    try {
        const categoryData = {
            name,
            description,
            parent_id: parentId || null,
            sort,
            status
        };
        
        const response = await fetch(`/api/categories/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(categoryData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分类更新成功', 'success');
                document.getElementById('editCategoryModal').style.display = 'none';
                loadCategoriesData();
            } else {
                showMessage(result.message || '分类更新失败', 'danger');
            }
        } else {
            showMessage('分类更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新分类失败:', error);
        showMessage('更新分类失败', 'danger');
    }
}

// 删除分类
async function deleteCategory(categoryId) {
    if (!confirm('确定要删除这个分类吗？')) {
        return;
    }
    
    try {
        const response = await fetch(`/api/categories/${categoryId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('分类删除成功', 'success');
                loadCategoriesData();
            } else {
                showMessage(result.message || '分类删除失败', 'danger');
            }
        } else {
            showMessage('分类删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除分类失败:', error);
        showMessage('删除分类失败', 'danger');
    }
}

// 清空分类表单
function clearCategoryForm() {
    document.getElementById('categoryName').value = '';
    document.getElementById('categoryDescription').value = '';
    document.getElementById('categoryParent').value = '';
    document.getElementById('categorySort').value = '';
    document.getElementById('categoryStatus').checked = true;
}

// 打开添加分类模态框
function openAddCategoryModal() {
    clearCategoryForm();
    document.getElementById('addCategoryModal').style.display = 'block';
}

// 关闭添加分类模态框
function closeAddCategoryModal() {
    document.getElementById('addCategoryModal').style.display = 'none';
    clearCategoryForm();
}

// 关闭编辑分类模态框
function closeEditCategoryModal() {
    document.getElementById('editCategoryModal').style.display = 'none';
}
