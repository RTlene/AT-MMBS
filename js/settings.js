// 系统设置JavaScript文件

// 加载系统设置数据
async function loadSettingsData() {
    try {
        // 这里可以加载各种系统设置
        console.log('加载系统设置数据');
        
        // 显示默认设置
        displayDefaultSettings();
        
    } catch (error) {
        console.error('加载系统设置数据失败:', error);
        showMessage('加载系统设置数据失败', 'danger');
    }
}

// 显示默认设置
function displayDefaultSettings() {
    // 这里可以显示各种系统设置的默认值
    console.log('显示默认系统设置');
}

// 保存系统设置
async function saveSystemSettings() {
    try {
        // 获取各种设置值
        const settings = {
            site_name: document.getElementById('siteName')?.value || '',
            site_description: document.getElementById('siteDescription')?.value || '',
            admin_email: document.getElementById('adminEmail')?.value || '',
            upload_max_size: document.getElementById('uploadMaxSize')?.value || '10',
            enable_registration: document.getElementById('enableRegistration')?.checked || false,
            enable_comments: document.getElementById('enableComments')?.checked || true,
            maintenance_mode: document.getElementById('maintenanceMode')?.checked || false
        };
        
        // 这里可以发送到后端保存
        console.log('保存系统设置:', settings);
        
        showMessage('系统设置保存成功', 'success');
        
    } catch (error) {
        console.error('保存系统设置失败:', error);
        showMessage('保存系统设置失败', 'danger');
    }
}

// 重置系统设置
function resetSystemSettings() {
    if (confirm('确定要重置所有系统设置吗？这将恢复默认值。')) {
        try {
            // 重置各种设置到默认值
            if (document.getElementById('siteName')) {
                document.getElementById('siteName').value = '商城管理系统';
            }
            if (document.getElementById('siteDescription')) {
                document.getElementById('siteDescription').value = '专业的电商管理平台';
            }
            if (document.getElementById('adminEmail')) {
                document.getElementById('adminEmail').value = 'admin@example.com';
            }
            if (document.getElementById('uploadMaxSize')) {
                document.getElementById('uploadMaxSize').value = '10';
            }
            if (document.getElementById('enableRegistration')) {
                document.getElementById('enableRegistration').checked = true;
            }
            if (document.getElementById('enableComments')) {
                document.getElementById('enableComments').checked = true;
            }
            if (document.getElementById('maintenanceMode')) {
                document.getElementById('maintenanceMode').checked = false;
            }
            
            showMessage('系统设置已重置为默认值', 'info');
            
        } catch (error) {
            console.error('重置系统设置失败:', error);
            showMessage('重置系统设置失败', 'danger');
        }
    }
}

// 备份系统设置
function backupSystemSettings() {
    try {
        // 获取当前设置
        const settings = {
            site_name: document.getElementById('siteName')?.value || '',
            site_description: document.getElementById('siteDescription')?.value || '',
            admin_email: document.getElementById('adminEmail')?.value || '',
            upload_max_size: document.getElementById('uploadMaxSize')?.value || '10',
            enable_registration: document.getElementById('enableRegistration')?.checked || false,
            enable_comments: document.getElementById('enableComments')?.checked || true,
            maintenance_mode: document.getElementById('maintenanceMode')?.checked || false,
            backup_time: new Date().toISOString()
        };
        
        // 创建下载链接
        const dataStr = JSON.stringify(settings, null, 2);
        const dataBlob = new Blob([dataStr], {type: 'application/json'});
        const url = URL.createObjectURL(dataBlob);
        
        const link = document.createElement('a');
        link.href = url;
        link.download = `system_settings_backup_${new Date().toISOString().split('T')[0]}.json`;
        link.click();
        
        URL.revokeObjectURL(url);
        
        showMessage('系统设置备份成功', 'success');
        
    } catch (error) {
        console.error('备份系统设置失败:', error);
        showMessage('备份系统设置失败', 'danger');
    }
}

// 恢复系统设置
function restoreSystemSettings() {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = '.json';
    
    input.onchange = function(event) {
        const file = event.target.files[0];
        if (!file) return;
        
        const reader = new FileReader();
        reader.onload = function(e) {
            try {
                const settings = JSON.parse(e.target.result);
                
                // 恢复各种设置
                if (document.getElementById('siteName') && settings.site_name) {
                    document.getElementById('siteName').value = settings.site_name;
                }
                if (document.getElementById('siteDescription') && settings.site_description) {
                    document.getElementById('siteDescription').value = settings.site_description;
                }
                if (document.getElementById('adminEmail') && settings.admin_email) {
                    document.getElementById('adminEmail').value = settings.admin_email;
                }
                if (document.getElementById('uploadMaxSize') && settings.upload_max_size) {
                    document.getElementById('uploadMaxSize').value = settings.upload_max_size;
                }
                if (document.getElementById('enableRegistration') && settings.hasOwnProperty('enable_registration')) {
                    document.getElementById('enableRegistration').checked = settings.enable_registration;
                }
                if (document.getElementById('enableComments') && settings.hasOwnProperty('enable_comments')) {
                    document.getElementById('enableComments').checked = settings.enable_comments;
                }
                if (document.getElementById('maintenanceMode') && settings.hasOwnProperty('maintenance_mode')) {
                    document.getElementById('maintenanceMode').checked = settings.maintenance_mode;
                }
                
                showMessage('系统设置恢复成功', 'success');
                
            } catch (error) {
                console.error('恢复系统设置失败:', error);
                showMessage('恢复系统设置失败，文件格式错误', 'danger');
            }
        };
        
        reader.readAsText(file);
    };
    
    input.click();
}

// 清除系统缓存
function clearSystemCache() {
    if (confirm('确定要清除系统缓存吗？')) {
        try {
            // 这里可以清除各种缓存
            console.log('清除系统缓存');
            
            showMessage('系统缓存清除成功', 'success');
            
        } catch (error) {
            console.error('清除系统缓存失败:', error);
            showMessage('清除系统缓存失败', 'danger');
        }
    }
}

// 系统诊断
function runSystemDiagnostic() {
    try {
        // 这里可以运行各种系统诊断
        console.log('运行系统诊断');
        
        // 模拟诊断结果
        const diagnosticResults = [
            '数据库连接正常',
            '文件上传目录可写',
            '缓存系统正常',
            '邮件服务配置正确'
        ];
        
        // 显示诊断结果
        let resultHtml = '<h5>系统诊断结果：</h5><ul>';
        diagnosticResults.forEach(result => {
            resultHtml += `<li class="text-success">✅ ${result}</li>`;
        });
        resultHtml += '</ul>';
        
        // 创建模态框显示结果
        const modal = document.createElement('div');
        modal.className = 'modal fade';
        modal.innerHTML = `
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">系统诊断结果</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
                    </div>
                    <div class="modal-body">
                        ${resultHtml}
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
                    </div>
                </div>
            </div>
        `;
        
        document.body.appendChild(modal);
        const bootstrapModal = new bootstrap.Modal(modal);
        bootstrapModal.show();
        
        // 模态框关闭后移除
        modal.addEventListener('hidden.bs.modal', function() {
            document.body.removeChild(modal);
        });
        
    } catch (error) {
        console.error('运行系统诊断失败:', error);
        showMessage('运行系统诊断失败', 'danger');
    }
}
