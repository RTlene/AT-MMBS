# AT-MMBS 测试指南

## 可用的测试工具

### 1. 综合测试面板 (comprehensive_test.html)
**用途**：手动交互式测试所有功能  
**使用方法**：
```bash
# 在浏览器中打开
open comprehensive_test.html
```
**特点**：
- 可视化测试界面
- 分模块测试功能
- 实时查看测试日志
- 模拟各种场景

### 2. 自动化测试脚本 (automated_test.js)
**用途**：Node.js环境下的自动化测试  
**使用方法**：
```bash
node automated_test.js
```
**输出**：
- 控制台显示测试进度
- 生成JSON格式测试报告
- 显示成功率统计

### 3. 浏览器模拟测试 (browser_test_simulator.html)
**用途**：模拟真实浏览器环境测试  
**使用方法**：
```bash
# 在浏览器中打开
open browser_test_simulator.html
```
**特点**：
- 自动运行测试序列
- 捕获console输出
- 模拟API调用

## 测试覆盖范围

- ✅ 全局函数（showMessage, showModal等）
- ✅ 用户管理（CRUD操作）
- ✅ 商品管理（包括图片上传）
- ✅ 分类管理（层级结构）
- ✅ 会员管理
- ✅ 订单和支付查看
- ✅ 认证和授权
- ✅ 错误处理
- ✅ XSS防护

## 快速测试命令

```bash
# 运行自动化测试
node automated_test.js

# 查看最新测试报告
cat test-report-*.json | jq '.'
```

## 测试报告

详细的测试报告请查看：`COMPREHENSIVE_TEST_REPORT.md`