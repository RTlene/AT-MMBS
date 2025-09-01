# admin.html 第1536行错误修复指南

## 问题描述
admin.html 第1536行存在一个JavaScript转义问题。当商品名称包含单引号（'）时，会导致JavaScript语法错误。

## 错误代码（第1536行）
```javascript
onclick="deleteProduct('${product.id}', '${product.name || product.productName || '未知商品'}')"
```

## 修复方案

### 方案1：手动修改（推荐）
将第1536行的代码：
```javascript
<button class="btn btn-danger" style="padding: 0.2rem 0.5rem; font-size: 0.8rem;" onclick="deleteProduct('${product.id}', '${product.name || product.productName || '未知商品'}')">删除</button>
```

修改为：
```javascript
<button class="btn btn-danger" style="padding: 0.2rem 0.5rem; font-size: 0.8rem;" onclick="deleteProduct('${product.id}', '${(product.name || product.productName || '未知商品').replace(/'/g, "\\'").replace(/"/g, '\\"')}')">删除</button>
```

### 方案2：使用提供的修复函数
1. 在admin.html的`<script>`部分添加以下函数：
```javascript
function safeDeleteProduct(productId, productName) {
    // 转义产品名称中的特殊字符
    const safeName = productName.replace(/'/g, "\\'").replace(/"/g, '\\"');
    // 调用原始的deleteProduct函数
    deleteProduct(productId, safeName);
}
```

2. 然后将第1536行修改为：
```javascript
<button class="btn btn-danger" style="padding: 0.2rem 0.5rem; font-size: 0.8rem;" onclick="safeDeleteProduct('${product.id}', '${product.name || product.productName || '未知商品'}')">删除</button>
```

## 为什么会出现这个问题？
当商品名称包含特殊字符（如单引号）时，会破坏JavaScript字符串的完整性。例如：
- 商品名称：`John's Product`
- 生成的代码：`deleteProduct('123', 'John's Product')` <- 这里的单引号会导致语法错误

## 其他需要检查的地方
建议搜索并检查admin.html中所有类似的动态生成JavaScript代码的地方，确保都进行了适当的转义处理。