// 商品管理JavaScript文件

// HTML转义函数
const escapeHtml = (str) => {
    if (!str) return '';
    return str.replace(/[&<>"']/g, function(match) {
        const escape = {
            '&': '&amp;',
            '<': '&lt;',
            '>': '&gt;',
            '"': '&quot;',
            "'": '&#39;'
        };
        return escape[match];
    });
};

// 商品相关变量
let currentProductId = null;
let uploadedImages = [];
let currentEditor = null;

// 加载商品数据
async function loadProductsData() {
    try {
        const response = await fetch(`/api/products?page=${currentPage}&size=${pageSize}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            console.log('商品API响应:', result);
            
            if (result.code === 0 && result.data) {
                displayProducts(result.data.products);
                updatePagination(result.data.total, result.data.page, result.data.size);
            } else {
                showMessage('加载商品数据失败', 'danger');
            }
        } else {
            showMessage('加载商品数据失败', 'danger');
        }
    } catch (error) {
        console.error('加载商品数据失败:', error);
        showMessage('加载商品数据失败', 'danger');
    }
}

// 显示商品列表
function displayProducts(products) {
    const tbody = document.getElementById('productsTableBody');
    if (!tbody) return;
    
    tbody.innerHTML = '';
    
    products.forEach(product => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td style="font-weight: 500;">${product.name || '未命名'}</td>
            <td>
                <span class="badge bg-info" style="padding: 0.3rem 0.6rem;">
                    ${getCategoryName(product.category) || '未知分类'}
                </span>
            </td>
            <td style="font-weight: 500; color: #e74c3c;">¥${(product.price || 0).toFixed(2)}</td>
            <td style="text-align: center;">
                <span style="font-weight: 500; color: ${product.kucun > 10 ? '#27ae60' : '#e74c3c'};">
                    ${product.kucun || 0}
                </span>
            </td>
            <td>
                <span class="badge ${product.status === 1 ? 'bg-success' : 'bg-secondary'}" style="padding: 0.3rem 0.6rem;">
                    ${product.status === 1 ? '上架' : '下架'}
                </span>
            </td>
            <td style="color: #666; font-size: 0.9em;">
                ${product.create_time ? new Date(product.create_time).toLocaleString('zh-CN', { 
                    year: 'numeric', 
                    month: '2-digit', 
                    day: '2-digit',
                    hour: '2-digit',
                    minute: '2-digit'
                }) : '未知'}
            </td>
            <td>
                <div class="btn-group" role="group">
                    <button class="btn btn-sm btn-primary" onclick="editProduct('${product.id}')"
                        style="padding: 0.25rem 0.5rem; font-size: 0.875rem;">
                        编辑
                    </button>
                    <button class="btn btn-sm btn-warning" onclick="toggleProductStatus('${product.id}')"
                        style="padding: 0.25rem 0.5rem; font-size: 0.875rem;">
                        ${product.status === 1 ? '下架' : '上架'}
                    </button>
                    <button class="btn btn-sm btn-danger" onclick="showDeleteProductModal('${product.id}', '${escapeHtml(product.name || product.productName || '未知商品')}')"
                        style="padding: 0.25rem 0.5rem; font-size: 0.875rem;">
                        删除
                    </button>
                </div>
            </td>
        `;
        tbody.appendChild(row);
    });
}

// 获取分类名称
function getCategoryName(categoryId) {
    // 这里应该从分类数据中获取名称，暂时返回ID
    return categoryId;
}

// 更新分页
function updatePagination(total, currentPage, pageSize) {
    const pagination = document.getElementById('productsPagination');
    if (!pagination) return;
    
    const totalPages = Math.ceil(total / pageSize);
    let paginationHTML = '';
    
    if (totalPages > 1) {
        paginationHTML += `
            <li class="page-item ${currentPage === 1 ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changePage(${currentPage - 1})">上一页</a>
            </li>
        `;
        
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `
                <li class="page-item ${i === currentPage ? 'active' : ''}">
                    <a class="page-link" href="#" onclick="changePage(${i})">${i}</a>
                </li>
            `;
        }
        
        paginationHTML += `
            <li class="page-item ${currentPage === totalPages ? 'disabled' : ''}">
                <a class="page-link" href="#" onclick="changePage(${currentPage + 1})">下一页</a>
            </li>
        `;
    }
    
    pagination.innerHTML = paginationHTML;
}

// 切换页面
function changePage(page) {
    if (page < 1) return;
    currentPage = page;
    loadProductsData();
}

// 清空商品表单
function clearProductForm() {
    document.getElementById('productName').value = '';
    document.getElementById('productCategory').value = '';
    document.getElementById('productPrice').value = '';
    document.getElementById('productKucun').value = '';
    document.getElementById('productContent').innerHTML = '';
    document.getElementById('productImages').innerHTML = '';
    uploadedImages = [];
    currentProductId = null;
    
    // 重置富文本编辑器
    if (currentEditor) {
        currentEditor.innerHTML = '';
    }
}

// 打开添加商品模态框
function openAddProductModal() {
    clearProductForm();
    document.getElementById('addProductModal').style.display = 'block';
    currentEditor = document.getElementById('productContent');
}

// 关闭添加商品模态框
function closeAddProductModal() {
    document.getElementById('addProductModal').style.display = 'none';
    clearProductForm();
}

// 保存商品
async function saveProduct() {
    const name = document.getElementById('productName').value.trim();
    const category = document.getElementById('productCategory').value;
    const price = parseFloat(document.getElementById('productPrice').value);
    const kucun = parseInt(document.getElementById('productKucun').value);
    const content = document.getElementById('productContent').innerHTML;
    
    if (!name || !category || isNaN(price) || isNaN(kucun)) {
        showMessage('请填写完整的商品信息', 'warning');
        return;
    }
    
    try {
        // 先上传图片
        const imagePaths = await uploadProductImages();
        
        const productData = {
            name,
            category,
            price,
            kucun,
            content,
            tp: imagePaths
        };
        
        console.log('发送的商品数据:', productData);
        
        const response = await fetch('/api/products', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(productData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('商品添加成功', 'success');
                closeAddProductModal();
                loadProductsData();
            } else {
                showMessage(result.message || '商品添加失败', 'danger');
            }
        } else {
            showMessage('商品添加失败', 'danger');
        }
    } catch (error) {
        console.error('添加商品失败:', error);
        showMessage('添加商品失败', 'danger');
    }
}

// 编辑商品
async function editProduct(productId) {
    try {
        const response = await fetch(`/api/products/${productId}`, {
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                fillEditProductForm(result.data);
                document.getElementById('editProductModal').style.display = 'block';
            } else {
                showMessage('获取商品信息失败', 'danger');
            }
        } else {
            showMessage('获取商品信息失败', 'danger');
        }
    } catch (error) {
        console.error('获取商品信息失败:', error);
        showMessage('获取商品信息失败', 'danger');
    }
}

// 填充编辑表单
function fillEditProductForm(product) {
    currentProductId = product.id;
    document.getElementById('editProductName').value = product.name || '';
    document.getElementById('editProductCategory').value = product.category || '';
    document.getElementById('editProductPrice').value = product.price || '';
    document.getElementById('editProductKucun').value = product.kucun || '';
    document.getElementById('editProductContent').innerHTML = product.content || '';
    
    // 显示现有图片
    displayEditProductImages(product.tp || []);
    
    currentEditor = document.getElementById('editProductContent');
}

// 显示编辑表单的图片
function displayEditProductImages(imagePaths) {
    const container = document.getElementById('editProductImages');
    if (!container) return;
    
    container.innerHTML = '';
    uploadedImages = [];
    
    if (imagePaths && imagePaths.length > 0) {
        imagePaths.forEach((imagePath, index) => {
            const imageDiv = document.createElement('div');
            imageDiv.className = 'image-item position-relative d-inline-block me-2 mb-2';
            imageDiv.innerHTML = `
                <img src="/uploads/${imagePath}" alt="商品图片" style="width: 100px; height: 100px; object-fit: cover;">
                <button type="button" class="btn btn-sm btn-danger position-absolute top-0 end-0" 
                        onclick="removeEditProductImage(${index})" style="margin: 2px;">×</button>
            `;
            container.appendChild(imageDiv);
            uploadedImages.push(imagePath);
        });
    }
}

// 删除编辑表单中的图片
function removeEditProductImage(index) {
    if (index >= 0 && index < uploadedImages.length) {
        uploadedImages.splice(index, 1);
        displayEditProductImages(uploadedImages);
    }
}

// 更新商品
async function updateProduct() {
    const name = document.getElementById('editProductName').value.trim();
    const category = document.getElementById('editProductCategory').value;
    const price = parseFloat(document.getElementById('editProductPrice').value);
    const kucun = parseInt(document.getElementById('editProductKucun').value);
    const content = document.getElementById('editProductContent').innerHTML;
    
    if (!name || !category || isNaN(price) || isNaN(kucun)) {
        showMessage('请填写完整的商品信息', 'warning');
        return;
    }
    
    try {
        const productData = {
            name,
            category,
            price,
            kucun,
            content,
            tp: uploadedImages
        };
        
        console.log('更新的商品数据:', productData);
        
        const response = await fetch(`/api/products/${currentProductId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(productData)
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('商品更新成功', 'success');
                document.getElementById('editProductModal').style.display = 'none';
                loadProductsData();
            } else {
                showMessage(result.message || '商品更新失败', 'danger');
            }
        } else {
            showMessage('商品更新失败', 'danger');
        }
    } catch (error) {
        console.error('更新商品失败:', error);
        showMessage('更新商品失败', 'danger');
    }
}

// 删除商品
async function deleteProduct(productId) {
    try {
        const response = await fetch(`/api/products/${productId}`, {
            method: 'DELETE',
            headers: addAuthHeader()
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                showMessage('商品删除成功', 'success');
                loadProductsData();
            } else {
                showMessage(result.message || '商品删除失败', 'danger');
            }
        } else {
            showMessage('商品删除失败', 'danger');
        }
    } catch (error) {
        console.error('删除商品失败:', error);
        showMessage('删除商品失败', 'danger');
    }
}

// 切换商品状态
async function toggleProductStatus(productId) {
    console.log('开始切换商品状态，ID:', productId);
    
    try {
        // 先获取当前商品信息
        const getResponse = await fetch(`/api/products/${productId}`, {
            headers: addAuthHeader()
        });
        
        if (!getResponse.ok) {
            showMessage('获取商品信息失败', 'danger');
            return;
        }
        
        const getResult = await getResponse.json();
        console.log('获取商品信息结果:', getResult);
        
        if (getResult.code !== 0) {
            showMessage('获取商品信息失败', 'danger');
            return;
        }
        
        const currentStatus = getResult.data.status;
        const newStatus = currentStatus === 1 ? 0 : 1;
        console.log('当前状态:', currentStatus, '新状态:', newStatus);
        
        // 发送状态更新请求
        const updateData = { id: productId, status: newStatus };
        console.log('发送状态更新请求:', updateData);
        
        const updateResponse = await fetch('/api/products/status', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                ...addAuthHeader()
            },
            body: JSON.stringify(updateData)
        });
        
        console.log('状态更新响应状态:', updateResponse.status);
        
        if (updateResponse.ok) {
            const updateResult = await updateResponse.json();
            console.log('状态更新API响应:', updateResult);
            
            if (updateResult.code === 0) {
                showMessage('商品状态更新成功', 'success');
                loadProductsData();
            } else {
                showMessage(updateResult.message || '商品状态更新失败', 'danger');
            }
        } else {
            showMessage('商品状态更新失败', 'danger');
        }
    } catch (error) {
        console.error('切换商品状态失败:', error);
        showMessage('切换商品状态失败', 'danger');
    }
}

// 关闭编辑商品模态框
function closeEditProductModal() {
    document.getElementById('editProductModal').style.display = 'none';
    clearProductForm();
}

// 图片上传相关函数
async function uploadProductImages() {
    const imageInput = document.getElementById('productImageInput');
    if (!imageInput || !imageInput.files || imageInput.files.length === 0) {
        return uploadedImages;
    }
    
    const files = Array.from(imageInput.files);
    const uploadPromises = files.map(file => uploadImage(file));
    
    try {
        const results = await Promise.all(uploadPromises);
        const newImages = results.filter(result => result.success).map(result => result.path);
        uploadedImages = [...uploadedImages, ...newImages];
        return uploadedImages;
    } catch (error) {
        console.error('图片上传失败:', error);
        showMessage('图片上传失败', 'danger');
        return uploadedImages;
    }
}

// 上传单张图片
async function uploadImage(file) {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('type', 'products');
    
    try {
        const response = await fetch('/api/upload', {
            method: 'POST',
            body: formData
        });
        
        if (response.ok) {
            const result = await response.json();
            if (result.code === 0) {
                return { success: true, path: result.data.path };
            } else {
                return { success: false, error: result.message };
            }
        } else {
            return { success: false, error: '上传失败' };
        }
    } catch (error) {
        console.error('图片上传失败:', error);
        return { success: false, error: error.message };
    }
}

// 富文本编辑器图片插入
function insertImageToEditor(imageUrl) {
    if (currentEditor) {
        const img = document.createElement('img');
        img.src = imageUrl;
        img.style.cssText = 'max-width: 100%; height: auto; margin: 5px; border: 1px solid #ddd; border-radius: 4px;';
        currentEditor.appendChild(img);
    }
}

// 处理图片文件选择
function handleImageFileSelect(event) {
    const files = event.target.files;
    if (files && files.length > 0) {
        Array.from(files).forEach(file => {
            if (file.type.startsWith('image/')) {
                const reader = new FileReader();
                reader.onload = function(e) {
                    insertImageToEditor(e.target.result);
                };
                reader.readAsDataURL(file);
            }
        });
    }
}
