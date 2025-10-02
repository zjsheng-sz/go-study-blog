# Go Study Blog

一个基于 Go + Gin + Gorm 的博客后端示例项目。

## 运行环境
- Go 1.18 及以上
- MySQL 数据库
- 推荐使用 macOS/Linux/Windows

## 依赖安装
1. 安装 Go 语言环境（https://golang.org/dl/）
2. 安装依赖包（在项目根目录执行）：
   ```sh
   go mod tidy
   ```
3. 配置数据库连接：
   - 编辑 `config/env/dev.yml` 或 `config/env/pro.yml`，填写你的 MySQL 连接信息。

## 启动方式
1. 启动 MySQL 数据库，并确保配置正确。
2. 在项目根目录运行：
   ```sh
   go run main.go
   ```
3. 默认监听端口为 8080，可通过配置文件修改。

## 主要依赖
- [Gin](https://github.com/gin-gonic/gin) — Web 框架
- [Gorm](https://gorm.io/) — ORM 框架
- [Viper](https://github.com/spf13/viper) — 配置管理
- [fsnotify](https://github.com/fsnotify/fsnotify) — 配置热更新

## 常用命令
- 安装依赖：`go mod tidy`
- 启动服务：`go run main.go`
- 构建可执行文件：`go build`

## 目录结构
```
api/           # 路由与控制器
common/        # 公共工具与错误处理
config/        # 配置文件与加载器
logger/        # 日志相关
middleware/    # Gin 中间件
models/        # 数据模型
repositories/  # 数据访问层
services/      # 业务逻辑层
utils/         # 工具函数
http/          # HTTP 测试文件
```

## 其他说明
- 推荐使用 VS Code 或 Goland 进行开发。
- 如需热重载，可结合 [air](https://github.com/cosmtrek/air) 等工具。
- 如遇依赖问题，优先执行 `go mod tidy`。

---
如有问题请提交 Issue 或联系项目维护者。
