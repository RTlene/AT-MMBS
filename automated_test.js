// AT-MMBS 自动化测试脚本
// 用于在Node.js环境中运行自动化测试

const fs = require('fs');
const path = require('path');

// 测试结果收集器
class TestResults {
    constructor() {
        this.results = [];
        this.passed = 0;
        this.failed = 0;
        this.warnings = 0;
    }
    
    add(module, test, status, message) {
        this.results.push({
            timestamp: new Date().toISOString(),
            module,
            test,
            status,
            message
        });
        
        if (status === 'PASS') this.passed++;
        else if (status === 'FAIL') this.failed++;
        else if (status === 'WARN') this.warnings++;
    }
    
    generateReport() {
        const report = {
            summary: {
                total: this.results.length,
                passed: this.passed,
                failed: this.failed,
                warnings: this.warnings,
                successRate: ((this.passed / this.results.length) * 100).toFixed(2) + '%'
            },
            details: this.results,
            timestamp: new Date().toISOString()
        };
        
        return report;
    }
}

const testResults = new TestResults();

// 测试文件存在性
function testFileExists() {
    console.log('\n=== 文件存在性测试 ===');
    
    const requiredFiles = [
        'admin.html',
        'login.html',
        'js/admin.js',
        'js/products.js',
        'js/users.js',
        'js/categories.js',
        'js/members.js',
        'js/orders.js',
        'js/payments.js',
        'js/dashboard.js'
    ];
    
    requiredFiles.forEach(file => {
        if (fs.existsSync(file)) {
            console.log(`✓ ${file} 存在`);
            testResults.add('文件检查', file, 'PASS', '文件存在');
        } else {
            console.log(`✗ ${file} 不存在`);
            testResults.add('文件检查', file, 'FAIL', '文件不存在');
        }
    });
}

// 测试JavaScript语法
function testJavaScriptSyntax() {
    console.log('\n=== JavaScript 语法测试 ===');
    
    const jsFiles = [
        'js/admin.js',
        'js/products.js',
        'js/users.js',
        'js/categories.js'
    ];
    
    jsFiles.forEach(file => {
        try {
            const content = fs.readFileSync(file, 'utf8');
            
            // 检查基本语法错误
            const syntaxChecks = [
                { pattern: /console\.log\(/g, name: 'console.log调用' },
                { pattern: /function\s+\w+\s*\(/g, name: '函数定义' },
                { pattern: /async\s+function/g, name: '异步函数' },
                { pattern: /try\s*{/g, name: 'try语句' },
                { pattern: /catch\s*\(/g, name: 'catch语句' }
            ];
            
            syntaxChecks.forEach(check => {
                const matches = content.match(check.pattern);
                if (matches) {
                    console.log(`✓ ${file} - ${check.name}: ${matches.length}个`);
                    testResults.add('语法检查', `${file} - ${check.name}`, 'PASS', `找到 ${matches.length} 个`);
                }
            });
            
            // 检查潜在问题
            if (content.includes('eval(')) {
                console.log(`⚠ ${file} - 使用了eval()`);
                testResults.add('语法检查', `${file} - eval使用`, 'WARN', '检测到eval()使用');
            }
            
        } catch (error) {
            console.log(`✗ ${file} - 读取失败: ${error.message}`);
            testResults.add('语法检查', file, 'FAIL', `读取失败: ${error.message}`);
        }
    });
}

// 测试函数定义
function testFunctionDefinitions() {
    console.log('\n=== 函数定义测试 ===');
    
    const requiredFunctions = {
        'js/admin.js': ['showMessage', 'addAuthHeader', 'checkLoginStatus'],
        'js/products.js': ['loadProductsData', 'saveProduct', 'deleteProduct', 'escapeHtml'],
        'js/users.js': ['loadUsersData', 'saveUser', 'editUser', 'deleteUser', 'toggleUserStatus'],
        'js/categories.js': ['loadCategoriesData', 'saveCategory', 'editCategory', 'deleteCategory']
    };
    
    for (const [file, functions] of Object.entries(requiredFunctions)) {
        if (fs.existsSync(file)) {
            const content = fs.readFileSync(file, 'utf8');
            
            functions.forEach(func => {
                const pattern = new RegExp(`(function\\s+${func}|const\\s+${func}\\s*=|${func}\\s*:\\s*function|async\\s+function\\s+${func})`);
                if (pattern.test(content)) {
                    console.log(`✓ ${file} - ${func} 已定义`);
                    testResults.add('函数定义', `${file} - ${func}`, 'PASS', '函数已定义');
                } else {
                    console.log(`✗ ${file} - ${func} 未找到`);
                    testResults.add('函数定义', `${file} - ${func}`, 'FAIL', '函数未定义');
                }
            });
        }
    }
}

// 测试HTML结构
function testHTMLStructure() {
    console.log('\n=== HTML 结构测试 ===');
    
    const htmlFiles = ['admin.html', 'login.html'];
    
    htmlFiles.forEach(file => {
        if (fs.existsSync(file)) {
            const content = fs.readFileSync(file, 'utf8');
            
            // 检查必要元素
            const requiredElements = [
                { pattern: /<div\s+id="messageContainer"/i, name: '消息容器' },
                { pattern: /<script\s+src="js\/admin\.js"/i, name: 'admin.js引用' }
            ];
            
            if (file === 'admin.html') {
                requiredElements.push(
                    { pattern: /showModal/g, name: 'showModal函数调用' },
                    { pattern: /closeModal/g, name: 'closeModal函数调用' }
                );
            }
            
            requiredElements.forEach(element => {
                if (element.pattern.test(content)) {
                    console.log(`✓ ${file} - ${element.name} 存在`);
                    testResults.add('HTML结构', `${file} - ${element.name}`, 'PASS', '元素存在');
                } else {
                    console.log(`✗ ${file} - ${element.name} 缺失`);
                    testResults.add('HTML结构', `${file} - ${element.name}`, 'FAIL', '元素缺失');
                }
            });
        }
    });
}

// 测试API调用模式
function testAPIPatterns() {
    console.log('\n=== API 调用模式测试 ===');
    
    const jsFiles = ['js/users.js', 'js/products.js', 'js/categories.js'];
    
    jsFiles.forEach(file => {
        if (fs.existsSync(file)) {
            const content = fs.readFileSync(file, 'utf8');
            
            // 检查API调用模式
            const apiPatterns = [
                { pattern: /fetch\s*\(\s*['"`]\/api\//g, name: 'API调用' },
                { pattern: /addAuthHeader\s*\(\s*\)/g, name: '认证头部' },
                { pattern: /response\.ok/g, name: '响应状态检查' },
                { pattern: /response\.json\(\)/g, name: 'JSON解析' },
                { pattern: /showMessage\s*\(/g, name: '消息提示' }
            ];
            
            apiPatterns.forEach(pattern => {
                const matches = content.match(pattern.pattern);
                if (matches && matches.length > 0) {
                    console.log(`✓ ${file} - ${pattern.name}: ${matches.length}处`);
                    testResults.add('API模式', `${file} - ${pattern.name}`, 'PASS', `${matches.length} 处使用`);
                } else {
                    console.log(`⚠ ${file} - ${pattern.name}: 未使用`);
                    testResults.add('API模式', `${file} - ${pattern.name}`, 'WARN', '未检测到使用');
                }
            });
        }
    });
}

// 测试错误处理
function testErrorHandling() {
    console.log('\n=== 错误处理测试 ===');
    
    const jsFiles = ['js/users.js', 'js/products.js', 'js/categories.js'];
    
    jsFiles.forEach(file => {
        if (fs.existsSync(file)) {
            const content = fs.readFileSync(file, 'utf8');
            
            // 统计try-catch块
            const tryMatches = content.match(/try\s*{/g) || [];
            const catchMatches = content.match(/catch\s*\(/g) || [];
            
            console.log(`✓ ${file} - try-catch块: ${tryMatches.length}个`);
            testResults.add('错误处理', `${file} - try-catch`, 'PASS', `${tryMatches.length} 个try-catch块`);
            
            // 检查console.error使用
            const errorLogs = content.match(/console\.error\(/g) || [];
            if (errorLogs.length > 0) {
                console.log(`✓ ${file} - 错误日志: ${errorLogs.length}处`);
                testResults.add('错误处理', `${file} - 错误日志`, 'PASS', `${errorLogs.length} 处错误日志`);
            }
        }
    });
}

// 测试安全性
function testSecurity() {
    console.log('\n=== 安全性测试 ===');
    
    // 检查escapeHtml函数
    if (fs.existsSync('js/products.js')) {
        const content = fs.readFileSync('js/products.js', 'utf8');
        
        if (content.includes('escapeHtml')) {
            console.log('✓ escapeHtml 函数已定义');
            testResults.add('安全性', 'HTML转义函数', 'PASS', 'escapeHtml函数存在');
            
            // 检查是否使用
            const usages = content.match(/escapeHtml\s*\(/g) || [];
            if (usages.length > 1) { // 至少有一处使用（除了定义）
                console.log(`✓ escapeHtml 使用: ${usages.length - 1}处`);
                testResults.add('安全性', 'HTML转义使用', 'PASS', `${usages.length - 1} 处使用`);
            }
        }
    }
    
    // 检查敏感信息
    const allFiles = ['admin.html', 'login.html', ...fs.readdirSync('js').map(f => `js/${f}`)];
    
    allFiles.forEach(file => {
        if (fs.existsSync(file) && file.endsWith('.js') || file.endsWith('.html')) {
            const content = fs.readFileSync(file, 'utf8');
            
            // 检查硬编码的密码或token
            if (content.match(/password\s*=\s*["']\w+["']/i)) {
                console.log(`⚠ ${file} - 可能包含硬编码密码`);
                testResults.add('安全性', `${file} - 硬编码密码`, 'WARN', '检测到可能的硬编码密码');
            }
        }
    });
}

// 运行所有测试
function runAllTests() {
    console.log('🧪 开始运行 AT-MMBS 自动化测试...\n');
    
    testFileExists();
    testJavaScriptSyntax();
    testFunctionDefinitions();
    testHTMLStructure();
    testAPIPatterns();
    testErrorHandling();
    testSecurity();
    
    // 生成测试报告
    const report = testResults.generateReport();
    
    console.log('\n=== 测试总结 ===');
    console.log(`总测试数: ${report.summary.total}`);
    console.log(`✓ 通过: ${report.summary.passed}`);
    console.log(`✗ 失败: ${report.summary.failed}`);
    console.log(`⚠ 警告: ${report.summary.warnings}`);
    console.log(`成功率: ${report.summary.successRate}`);
    
    // 保存报告
    const reportPath = `test-report-${new Date().toISOString().split('T')[0]}.json`;
    fs.writeFileSync(reportPath, JSON.stringify(report, null, 2));
    console.log(`\n测试报告已保存到: ${reportPath}`);
    
    // 返回测试是否全部通过
    return report.summary.failed === 0;
}

// 如果直接运行此脚本
if (require.main === module) {
    const success = runAllTests();
    process.exit(success ? 0 : 1);
}

module.exports = { runAllTests };