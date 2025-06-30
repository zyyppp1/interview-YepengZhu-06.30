
---

# OXO Game API

基于 Go + Gin + PostgreSQL 的游戏管理系统 API，提供玩家管理、房间预约、无尽挑战、支付处理和日志收集等功能。

## 目录

- [快速开始](#快速开始)
- [项目结构](#项目结构)
- [API 文档](#api-文档)
- [开发指南](#开发指南)
- [配置说明](#配置说明)
- [故障排除](#故障排除)

## 快速开始

### 前置要求

- Docker 和 Docker Compose
- Git

### 部署步骤

1. **克隆项目**
```bash
git clone <repository-url>
cd interview-YepengZhu-06.30
```

2. **使用 Docker Compose 启动（推荐）**
```bash
# 一键启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f api
```

3. **手动构建 Docker 镜像**
```bash
# 构建镜像
docker build -t oxo-game-api .

# 运行容器（需要先启动 PostgreSQL）
docker run -d \
  --name oxo-api \
  -p 8080:8080 \
  --env-file .env \
  oxo-game-api
```

4. **验证部署**
```bash
# 健康检查
curl http://localhost:8080/health

# 应返回：
# {"status":"healthy","service":"OXO Game API"}
```

### 停止服务

```bash
# 停止所有服务
docker-compose down

# 停止并清除数据
docker-compose down -v
```

## 项目结构

```
interview-YepengZhu-06.30/
├── api/                    # API 处理器层
│   ├── challenge.go       # 挑战系统 API 处理器
│   ├── log.go            # 日志系统 API 处理器
│   ├── payment.go        # 支付系统 API 处理器
│   ├── player.go         # 玩家管理 API 处理器
│   └── room.go           # 房间管理 API 处理器
├── cmd/                   # 应用程序入口
│   └── main.go           # 主程序文件
├── config/               # 配置管理
│   └── config.go         # 配置加载和管理
├── db/                   # 数据库相关
│   └── db.go            # 数据库连接初始化和迁移
├── doc/                  # 文档目录
│   ├── 支付處理系統 (Payment Processing System)- 测试报告.md
│   ├── 無盡挑戰系統 (Endless Challenge System) - 测试报告.md
│   ├── 玩家管理系统 (Player Management System) - 测试报告.md
│   ├── 遊戲房間管理系統 (Game Room Management System) - 测试报告.md
│   └── 遊戲日誌收集器 (Game Log Collector) - 测试报告.md
├── middleware/           # 中间件
│   └── logger.go        # 日志记录和 CORS 中间件
├── models/              # 数据模型定义
│   ├── challenge.go     # 挑战和奖池模型
│   ├── log.go          # 游戏日志模型（含 JSONB 类型）
│   ├── payment.go      # 支付记录模型
│   ├── player.go       # 玩家和等级模型
│   └── room.go         # 房间和预约模型
├── Postman Test/        # Postman 测试集合
│   └── PostMan 测试.postman_collection.json
├── routes/             # 路由配置
│   └── router.go       # API 路由注册和中间件配置
├── scripts/            # 脚本文件
│   └── init.sql       # 数据库初始化脚本
├── services/           # 业务逻辑层
│   ├── challenge_service.go  # 挑战业务逻辑
│   ├── init.go              # 服务初始化
│   ├── payment_service.go   # 支付和日志业务逻辑
│   ├── player_service.go    # 玩家和等级业务逻辑
│   └── room_service.go      # 房间和预约业务逻辑
├── test/              # 测试文件
│   └── player_test.go # 玩家单元测试（待实现）
├── .env               # 环境变量配置
├── .gitignore        # Git 忽略文件
├── Dockerfile        # Docker 构建文件
├── docker-compose.yml # Docker Compose 编排文件
├── go.mod            # Go 模块定义（依赖管理）
├── go.sum            # Go 依赖校验和（确保依赖完整性）
├── Interview2025.md  # 面试题目要求说明
└── README.md         # 项目说明文档
```

### 各目录详细说明

#### `/api` - API 处理器层
负责处理 HTTP 请求和响应：
- 请求参数验证和绑定
- 调用服务层处理业务逻辑
- 格式化响应数据
- 统一错误处理

#### `/services` - 业务逻辑层
包含所有业务逻辑实现：
- 数据验证和业务规则
- 事务管理
- 调用数据层进行持久化
- 复杂业务流程编排

#### `/models` - 数据模型层
定义数据结构和数据库映射：
- GORM 模型定义
- 表关联关系
- 自定义数据类型（如 JSONB）
- 数据验证规则

#### `/doc` - 文档目录
包含所有 API 测试报告和文档：
- 详细的 API 测试用例
- 请求/响应示例
- 错误处理说明
- HTTP 状态码使用规范

## API 文档

基础 URL: `http://localhost:8080/api/v1`

### 📚 完整 API 测试报告

| 功能模块 | 测试报告 | 说明 |
|---------|---------|------|
| 玩家管理 | [查看测试报告](doc/玩家管理系统%20(Player%20Management%20System)%20-%20测试报告.md) | 玩家 CRUD、等级管理 |
| 房间管理 | [查看测试报告](doc/遊戲房間管理系統%20(Game%20Room%20Management%20System)%20-%20测试报告.md) | 房间 CRUD、预约管理 |
| 挑战系统 | [查看测试报告](doc/無盡挑戰系統%20(Endless%20Challenge%20System)%20-%20测试报告.md) | 挑战参与、结果查询 |
| 日志系统 | [查看测试报告](doc/遊戲日誌收集器%20(Game%20Log%20Collector)%20-%20测试报告.md) | 操作日志记录、查询 |
| 支付系统 | [查看测试报告](doc/支付處理系統%20(Payment%20Processing%20System)-%20测试报告.md) | 多种支付方式处理 |

### 🎮 API 端点概览

#### 玩家管理
- `GET /players` - 获取玩家列表（分页）
- `POST /players` - 创建新玩家
- `GET /players/{id}` - 获取玩家详情
- `PUT /players/{id}` - 更新玩家信息
- `DELETE /players/{id}` - 删除玩家
- `GET /levels` - 获取等级列表
- `POST /levels` - 创建新等级

#### 房间管理
- `GET /rooms` - 获取房间列表
- `POST /rooms` - 创建新房间
- `GET /rooms/{id}` - 获取房间详情
- `PUT /rooms/{id}` - 更新房间信息
- `DELETE /rooms/{id}` - 删除房间
- `GET /reservations` - 查询预约列表
- `POST /reservations` - 创建新预约

#### 挑战系统
- `POST /challenges` - 参加挑战
- `GET /challenges/results` - 获取挑战结果

#### 日志系统
- `GET /logs` - 查询游戏日志
- `POST /logs` - 创建日志记录

#### 支付系统
- `POST /payments` - 处理支付
- `GET /payments/{id}` - 获取支付详情

### 💡 快速测试

```bash
# 测试健康检查
curl http://localhost:8080/health

# 获取所有等级
curl http://localhost:8080/api/v1/levels

# 创建新玩家
curl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试玩家",
    "level_id": 1
  }'
```

更多详细测试用例请查看 [doc/](doc/) 目录下的测试报告。

## 开发指南

### 本地开发环境设置

1. **安装 Go 1.21+**
```bash
# 验证 Go 版本
go version
```

2. **安装依赖**
```bash
# 下载所有依赖包
go mod download

# 整理依赖
go mod tidy
```

3. **设置环境变量**
```bash
# 复制环境变量模板
cp .env.example .env

# 编辑 .env 文件配置数据库连接
```

4. **本地数据库设置**
```bash
# 启动 PostgreSQL
docker run -d \
  --name postgres-dev \
  -e POSTGRES_USER=oxogame \
  -e POSTGRES_PASSWORD=oxogame123 \
  -e POSTGRES_DB=oxogame_db \
  -p 5432:5432 \
  postgres:15-alpine

# 运行数据库迁移
psql -U oxogame -d oxogame_db -f scripts/init.sql
```

5. **启动开发服务器**
```bash
go run cmd/main.go
```

### 代码规范

- 使用 `gofmt` 格式化代码
- 遵循 Go 官方编码规范
- 添加必要的注释，特别是导出的函数和类型
- 错误处理要明确，避免忽略错误

### 测试

#### 运行 Postman 测试
1. 导入 `Postman Test/PostMan 测试.postman_collection.json`
2. 设置环境变量 `base_url` 为 `http://localhost:8080/api/v1`
3. 运行测试集合

#### 性能测试
```bash
# 使用 Apache Bench
ab -n 1000 -c 10 http://localhost:8080/api/v1/players

# 使用 wrk
wrk -t12 -c400 -d30s http://localhost:8080/api/v1/players
```

## 配置说明

### 环境变量 (`.env`)

```env
# 应用配置
APP_NAME=OXO Game API
APP_ENV=development      # development, production

# 服务器配置
PORT=8080
GIN_MODE=debug          # debug, release

# 数据库配置
DB_HOST=localhost       # Docker 环境下使用 postgres
DB_PORT=5432
DB_USER=oxogame
DB_PASSWORD=oxogame123
DB_NAME=oxogame_db
```

### Docker Compose 配置

`docker-compose.yml` 包含两个服务：
- **api**: Go 应用服务
  - 自动构建镜像
  - 依赖 postgres 服务健康检查
  - 挂载环境变量
- **postgres**: PostgreSQL 数据库
  - 使用 Alpine 版本（更小的镜像）
  - 数据持久化到 volume
  - 自动运行初始化脚本

### 数据库初始化

`scripts/init.sql` 自动完成：
- 创建所有必需的表结构
- 设置索引优化查询性能
- 插入默认等级数据（5个等级）
- 插入示例玩家和房间数据
- 初始化奖池

## 故障排除

### 常见问题

1. **数据库连接失败**
   ```bash
   # 检查 PostgreSQL 容器状态
   docker-compose ps postgres
   
   # 查看数据库日志
   docker-compose logs postgres
   ```

2. **端口被占用**
   ```bash
   # 查找占用端口的进程
   lsof -i :8080
   
   # 或修改 docker-compose.yml 中的端口映射
   ```

3. **依赖包下载失败**
   ```bash
   # 设置 Go 代理
   go env -w GOPROXY=https://goproxy.cn,direct
   
   # 清理模块缓存
   go clean -modcache
   ```

4. **Docker 构建失败**
   ```bash
   # 清理 Docker 缓存
   docker system prune -a
   
   # 查看构建日志
   docker-compose build --no-cache
   ```

## API 优化方向

1. **性能优化**
   - 实现数据库连接池配置
   - 添加 Redis 缓存层
   - 使用消息队列处理异步任务

2. **安全增强**
   - 实现 JWT 认证
   - 添加请求限流
   - SQL 注入防护

3. **功能扩展**
   - WebSocket 实时通信
   - 文件上传支持
   - API 版本管理

4. **监控和日志**
   - 集成 Prometheus 监控
   - 结构化日志输出
   - 链路追踪

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 许可证

本项目仅供面试评估使用。

## 联系方式

- 作者：Yepeng Zhu
- 日期：2025-06-30
