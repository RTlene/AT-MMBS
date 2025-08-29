# AT-MMBS 项目必要文件清单

本清单列出了项目运行所需的必要文件。不在此列表中的文件可以安全删除。

## 核心源代码文件

### 根目录必要文件
- `main.go` - 主程序入口
- `go.mod` - Go模块定义
- `go.sum` - Go依赖锁定文件
- `.dockerignore` - Docker忽略配置
- `Dockerfile` - Docker镜像构建文件
- `docker-compose.yml` - Docker Compose配置
- `nginx.conf` - Nginx配置文件
- `init.sql` - 数据库初始化脚本

### 配置目录 (config/)
- `config/config.go` - 配置管理

### 数据库目录 (db/)
- `db/init.go` - 数据库初始化
- `db/dao/interface.go` - DAO接口定义
- `db/dao/dao.go` - DAO基础实现
- `db/dao/user_dao.go` - 用户DAO
- `db/dao/product_dao.go` - 产品DAO  
- `db/dao/category_dao.go` - 分类DAO
- `db/dao/order_dao.go` - 订单DAO
- `db/dao/payment_dao.go` - 支付DAO
- `db/dao/member_dao.go` - 会员DAO
- `db/dao/member_level_dao.go` - 会员等级DAO
- `db/dao/distributor_level_dao.go` - 分销商等级DAO
- `db/dao/counter_dao.go` - 计数器DAO
- `db/model/user.go` - 用户模型
- `db/model/product.go` - 产品模型
- `db/model/category.go` - 分类模型
- `db/model/order.go` - 订单模型
- `db/model/payment.go` - 支付模型
- `db/model/member.go` - 会员模型
- `db/model/member_level.go` - 会员等级模型
- `db/model/distributor_level.go` - 分销商等级模型
- `db/model/counter.go` - 计数器模型

### 中间件目录 (middleware/)
- `middleware/auth.go` - 认证中间件
- `middleware/logger.go` - 日志中间件

### 路由目录 (router/)
- `router/router.go` - 路由配置

### 服务目录 (service/)
- `service/common.go` - 通用服务功能
- `service/auth_service.go` - 认证服务
- `service/user_service.go` - 用户服务
- `service/product_service.go` - 产品服务
- `service/category_service.go` - 分类服务
- `service/order_service.go` - 订单服务
- `service/payment_service.go` - 支付服务
- `service/member_service.go` - 会员服务
- `service/member_level_service.go` - 会员等级服务
- `service/distributor_level_service.go` - 分销商等级服务
- `service/counter_service.go` - 计数器服务
- `service/file_upload_service.go` - 文件上传服务
- `service/rest_handlers.go` - REST API处理器
- `service/static_handler.go` - 静态文件处理器

### 前端文件

#### HTML文件
- `index.html` - 首页
- `login.html` - 登录页
- `admin.html` - 管理后台主页

#### JavaScript目录 (js/)
- `js/admin.js` - 管理后台主脚本
- `js/dashboard.js` - 仪表盘功能
- `js/users.js` - 用户管理
- `js/products.js` - 产品管理
- `js/categories.js` - 分类管理
- `js/orders.js` - 订单管理
- `js/payments.js` - 支付管理
- `js/members.js` - 会员管理
- `js/member-levels.js` - 会员等级管理
- `js/distributor-levels.js` - 分销商等级管理
- `js/settings.js` - 设置管理

## 部署和文档文件

### 部署脚本
- `deploy.bat` - Windows部署脚本
- `deploy.sh` - Linux部署脚本
- `start.bat` - Windows启动脚本
- `stop.bat` - Windows停止脚本
- `test-docker-build.bat` - Docker构建测试脚本

### 文档文件
- `README.md` - 项目说明
- `LICENSE` - 许可证
- `DOCKER_DEPLOY.md` - Docker部署文档
- `DEPLOYMENT.md` - 部署文档
- `开发要求.md` - 开发需求文档
- `项目总结.md` - 项目总结文档

## 可以安全删除的文件

以下文件通常可以安全删除：

### 测试和临时文件
- `test.html`
- `test_*.html` (所有test_开头的HTML文件)
- `test.db` - 测试数据库
- `*.log` - 日志文件
- `comprehensive_test.py` - Python测试脚本
- `test_report.md` - 测试报告
- `simple_test.sh` - 简单测试脚本
- `test-deployment.sh` - 测试部署脚本
- `test-modular.sh` - 模块测试脚本

### 构建产物和二进制文件
- `at-mmbs-test` - 测试构建的二进制文件
- `mmbs-modular` - 模块化构建的二进制文件
- 任何 `.exe` 文件（Windows构建产物）

### 其他临时文件
- `deployment_status.txt` - 部署状态文件
- `improvement_suggestions.md` - 改进建议文档
- `WORK_COMPLETED.md` - 工作完成记录
- `fix_database*.sql` - 数据库修复脚本
- `container.config.json` - 容器配置（如果使用docker-compose.yml）
- `main_old.go.bak` - 备份文件

### IDE和系统文件
- `.DS_Store` (macOS)
- `Thumbs.db` (Windows)
- `.idea/` - IntelliJ IDEA配置
- `.vscode/` - VS Code配置（除非团队共享）

## 注意事项

1. **uploads/** 目录：如果有用户上传的文件，请根据需要保留
2. **logs/** 目录：可以清空但建议保留目录结构
3. **ssl/** 目录：如果配置了HTTPS，需要保留证书文件
4. 任何自定义的配置文件请根据实际情况保留

## 使用建议

1. 在删除文件前，建议先备份整个项目
2. 可以使用版本控制系统（Git）的 clean 命令清理未跟踪的文件
3. 删除文件后运行 `test-docker-build.bat` 确保项目仍能正常构建