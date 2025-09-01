// 这是一个临时的修复方案
// 在admin.html中添加这个函数来安全处理商品删除

function safeDeleteProduct(productId, productName) {
    // 转义产品名称中的特殊字符
    const safeName = productName.replace(/'/g, "\\'").replace(/"/g, '\\"');
    // 调用原始的deleteProduct函数
    deleteProduct(productId, safeName);
}