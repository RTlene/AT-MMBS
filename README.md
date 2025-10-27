# 微信小程序商城后台管理系统

这是一个基于Go语言开发的微信小程序商城后台管理系统，采用模块化架构设计，提供完整的商城管理功能。

## 技术架构

- **后端**: Go + GORM + MySQL
- **前端**: HTML + CSS + JavaScript + Bootstrap 5
- **数据库**: MySQL
- **部署**: Docker

## 功能模块

### 1. 用户管理模块
- 用户实体的增删改查功能
- 用户认证和权限管理

### 2. 商品管理模块
- 商品实体的增删改查功能
- 商品上下架管理
- 商品分类管理
- 商品库存管理

### 3. 商品分类管理模块
- 分类的增删改查功能
- 支持父子分类结构
- 分类状态管理

### 4. 会员管理模块
- 会员实体的增删改查功能
- 会员关系网管理
- 会员分销信息管理
- 会员消费记录管理

### 5. 会员等级管理模块
- 会员等级的增删改查功能
- 等级折扣和积分倍率设置
- 自动升级规则配置

### 6. 分销等级管理模块
- 分销等级的增删改查功能
- 佣金比例设置
- 分销模式配置

### 7. 订单管理模块
- 订单的增删改查功能
- 订单状态管理
- 订单修改限制控制

### 8. 支付管理模块
- 支付记录的增删改查功能
- 支付状态管理
- 退款处理

## 项目结构

```
AT-MMBS/
├── main.go                 # 主程序入口
├── go.mod                  # Go模块文件
├── go.sum                  # Go依赖校验文件
├── Dockerfile              # Docker构建文件
├── container.config.json   # 容器配置文件
├── index.html              # 首页
├── admin.html              # 管理后台页面
├── README.md               # 项目说明文档
├── 开发要求.md             # 开发需求文档
├── db/                     # 数据库相关
│   ├── init.go            # 数据库初始化
│   ├── dao/               # 数据访问层
│   │   ├── interface.go   # DAO接口定义
│   │   ├── dao.go         # 计数器DAO实现
│   │   ├── user_dao.go    # 用户DAO实现
│   │   ├── product_dao.go # 商品DAO实现
│   │   ├── category_dao.go # 分类DAO实现
│   │   ├── member_dao.go  # 会员DAO实现
│   │   ├── member_level_dao.go # 会员等级DAO实现
│   │   ├── distributor_level_dao.go # 分销等级DAO实现
│   │   ├── order_dao.go   # 订单DAO实现
│   │   └── payment_dao.go # 支付DAO实现
│   └── model/             # 数据模型
│       ├── counter.go     # 计数器模型
│       ├── user.go        # 用户模型
│       ├── product.go     # 商品模型
│       ├── category.go    # 分类模型
│       ├── member.go      # 会员模型
│       ├── member_level.go # 会员等级模型
│       ├── distributor_level.go # 分销等级模型
│       ├── order.go       # 订单模型
│       └── payment.go     # 支付模型
└── service/               # 业务逻辑层
    ├── counter_service.go # 计数器服务
    ├── user_service.go    # 用户管理服务
    ├── product_service.go # 商品管理服务
    ├── category_service.go # 分类管理服务
    ├── member_service.go  # 会员管理服务
    ├── member_level_service.go # 会员等级管理服务
    ├── distributor_level_service.go # 分销等级管理服务
    ├── order_service.go   # 订单管理服务
    └── payment_service.go # 支付管理服务
```

## 环境要求

- Go 1.16+
- MySQL 5.7+
- Docker (可选)

## 安装和运行

### 1. 本地开发环境

#### 设置环境变量
```bash
export MYSQL_USERNAME=your_username
export MYSQL_PASSWORD=your_password
export MYSQL_ADDRESS=localhost:3306
export MYSQL_DATABASE=at_mmbs
```

#### 安装依赖
```bash
go mod tidy
```

#### 运行项目
```bash
go run main.go
```

### 2. Docker部署

#### 构建镜像
```bash
docker build -t at-mmbs .
```

#### 运行容器
```bash
docker run -d \
  -p 80:80 \
  -e MYSQL_USERNAME=your_username \
  -e MYSQL_PASSWORD=your_password \
  -e MYSQL_ADDRESS=your_mysql_host:3306 \
  -e MYSQL_DATABASE=at_mmbs \
  at-mmbs
```

## API接口

### 用户管理
- `GET /api/users` - 获取用户列表
- `GET /api/users?id={id}` - 获取单个用户
- `POST /api/users` - 创建用户
- `PUT /api/users` - 更新用户
- `DELETE /api/users?id={id}` - 删除用户

### 商品管理
- `GET /api/products` - 获取商品列表
- `GET /api/products?id={id}` - 获取单个商品
- `POST /api/products` - 创建商品
- `PUT /api/products` - 更新商品
- `DELETE /api/products?id={id}` - 删除商品
- `PUT /api/products/status` - 更新商品状态

### 分类管理
- `GET /api/categories` - 获取分类列表
- `GET /api/categories?id={id}` - 获取单个分类
- `POST /api/categories` - 创建分类
- `PUT /api/categories` - 更新分类
- `DELETE /api/categories?id={id}` - 删除分类
- `GET /api/categories/parent` - 获取父分类

### 会员管理
- `GET /api/members` - 获取会员列表
- `GET /api/members?id={id}` - 获取单个会员
- `POST /api/members` - 创建会员
- `PUT /api/members` - 更新会员
- `DELETE /api/members?id={id}` - 删除会员
- `GET /api/members/network` - 获取会员关系网
- `GET /api/members/distribution` - 获取会员分销信息
- `GET /api/members/consumption` - 获取会员消费记录

### 会员等级管理
- `GET /api/member-levels` - 获取会员等级列表
- `GET /api/member-levels?id={id}` - 获取单个会员等级
- `POST /api/member-levels` - 创建会员等级
- `PUT /api/member-levels` - 更新会员等级
- `DELETE /api/member-levels?id={id}` - 删除会员等级

### 分销等级管理
- `GET /api/distributor-levels` - 获取分销等级列表
- `GET /api/distributor-levels?id={id}` - 获取单个分销等级
- `POST /api/distributor-levels` - 创建分销等级
- `PUT /api/distributor-levels` - 更新分销等级
- `DELETE /api/distributor-levels?id={id}` - 删除分销等级

### 订单管理
- `GET /api/orders` - 获取订单列表
- `GET /api/orders?id={id}` - 获取单个订单
- `GET /api/orders?order_no={order_no}` - 根据订单号获取订单
- `POST /api/orders` - 创建订单
- `PUT /api/orders` - 更新订单
- `DELETE /api/orders?id={id}` - 废除订单
- `PUT /api/orders/status` - 更新订单状态

### 支付管理
- `GET /api/payments` - 获取支付记录列表
- `GET /api/payments?id={id}` - 获取单个支付记录
- `GET /api/payments?order_id={order_id}` - 根据订单ID获取支付记录
- `POST /api/payments` - 创建支付记录
- `PUT /api/payments` - 更新支付记录
- `PUT /api/payments/status` - 更新支付状态

## 访问地址

- 首页: `http://localhost:80/`
- 管理后台: `http://localhost:80/admin`

## 开发说明

### 模块化设计
项目采用模块化开发方式，每个功能模块都有独立的：
- 数据模型 (Model)
- 数据访问层 (DAO)
- 业务逻辑层 (Service)
- API接口

### 数据库设计
- 使用GORM作为ORM框架
- 支持自动迁移数据库表结构
- 使用MySQL作为数据库

### 前端界面
- 采用Bootstrap 5框架
- 响应式设计，支持移动端
- 现代化的UI设计风格

## 注意事项

1. 首次运行时会自动创建数据库表
2. 请确保MySQL服务正常运行
3. 环境变量配置正确
4. 建议在生产环境中使用HTTPS

## 许可证

本项目采用MIT许可证，详见LICENSE文件。

## 贡献

欢迎提交Issue和Pull Request来改进这个项目。

## 联系方式

如有问题或建议，请通过GitHub Issues联系我们。

## 问题修复记录

### 2024年问题修复

#### 1. 商品编辑和状态切换问题
- **问题描述**: 商品编辑无法更新，状态切换失败，报错 `Error 1292: Incorrect datetime value: '0000-00-00' for column 'create_time'`
- **原因分析**: 数据库中存在无效的日期值 `0000-00-00`，这在MySQL中是不合法的
- **解决方案**: 
  - 修复了DAO层的更新逻辑，避免覆盖 `create_time` 字段
  - 修复了状态切换时的字段更新逻辑
  - 创建了数据库修复脚本 `fix_database.sql` 来修复现有数据
- **修复文件**: 
  - `db/dao/product_dao.go`
  - `service/product_service.go`
  - `fix_database.sql`

#### 2. 富文本编辑器功能增强
- **问题描述**: 商品编辑的富文本编辑器只能通过链接插入图片，缺少本地上传功能
- **解决方案**: 
  - 添加了图片上传按钮
  - 实现了本地上传图片到富文本编辑器的功能
  - 支持在光标位置插入图片
- **修复文件**: `admin.html`

### 使用说明

#### 数据库修复
如果遇到日期相关的数据库错误，请执行以下步骤：

1. 运行数据库修复脚本：
```sql
mysql -u your_username -p < fix_database.sql
```

2. 或者手动执行修复SQL：
```sql
USE at_mmbs;
UPDATE products SET create_time = CURRENT_TIMESTAMP WHERE create_time = '0000-00-00' OR create_time IS NULL;
UPDATE products SET update_time = CURRENT_TIMESTAMP WHERE update_time = '0000-00-00' OR update_time IS NULL;
```

#### 富文本编辑器使用
- 点击"图片"按钮可以选择本地图片文件上传
- 上传成功后图片会自动插入到编辑器的光标位置
- 支持调整图片大小，图片会自动适应编辑器宽度
