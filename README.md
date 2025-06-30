# OXO Game API

基于 Go + Gin + PostgreSQL 的游戏管理系统 API，提供玩家管理、房间预约、无尽挑战、支付处理和日志收集等功能。

## 目录

- [快速开始](#快速开始)
- [项目结构](#项目结构)
- [API 文档](#api-文档)
- [开发指南](#开发指南)
- [配置说明](#配置说明)

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
或
docker compose up -d

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
│   ├── challenge.go       # 挑战系统 API
│   ├── log.go            # 日志系统 API
│   ├── payment.go        # 支付系统 API
│   ├── player.go         # 玩家管理 API
│   └── room.go           # 房间管理 API
├── cmd/                   # 应用程序入口
│   └── main.go           # 主程序文件
├── config/               # 配置管理
│   └── config.go         # 配置加载器
├── db/                   # 数据库相关
│   └── db.go            # 数据库连接和迁移
├── middleware/           # 中间件
│   └── logger.go        # 日志和 CORS 中间件
├── models/              # 数据模型
│   ├── challenge.go     # 挑战和奖池模型
│   ├── log.go          # 游戏日志模型
│   ├── payment.go      # 支付记录模型
│   ├── player.go       # 玩家和等级模型
│   └── room.go         # 房间和预约模型
├── routes/             # 路由配置
│   └── router.go       # 路由注册
├── services/           # 业务逻辑层
│   ├── challenge_service.go  # 挑战业务逻辑
│   ├── init.go              # 服务初始化
│   ├── payment_service.go   # 支付和日志业务逻辑
│   ├── player_service.go    # 玩家业务逻辑
│   └── room_service.go      # 房间业务逻辑
├── scripts/            # 脚本文件
│   └── init.sql       # 数据库初始化脚本
├── test/              # 测试文件
│   └── player_test.go # 玩家测试（待实现）
├── .env               # 环境变量配置
├── .gitignore        # Git 忽略文件
├── Dockerfile        # Docker 构建文件
├── docker-compose.yml # Docker Compose 配置
├── go.mod            # Go 模块定义
├── go.sum            # Go 依赖版本锁定
└── Interview2025.md  # 面试题目说明
```

## API 文档

基础 URL: `http://localhost:8080/api/v1`

### 📚 完整 API 测试报告

| 功能模块 | 测试报告 | 说明 |
|---------|---------|------|
| 玩家管理 | [玩家管理系统测试报告](玩家管理系统%20(Player%20Management%20System)%20-%20测试报告.md) | 玩家 CRUD、等级管理 |
| 房间管理 | [游戏房间管理系统测试报告](遊戲房間管理系統%20(Game%20Room%20Management%20System)%20-%20测试报告.md) | 房间 CRUD、预约管理 |
| 挑战系统 | [无尽挑战系统测试报告](無盡挑戰系統%20(Endless%20Challenge%20System)%20-%20测试报告.md) | 挑战参与、结果查询 |
| 日志系统 | [游戏日志收集器测试报告](遊戲日誌收集器%20(Game%20Log%20Collector)%20-%20测试报告.md) | 操作日志记录、查询 |
| 支付系统 | [支付处理系统测试报告](支付處理系統%20(Payment%20Processing%20System)-%20测试报告.md) | 多种支付方式处理 |

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

### 💡 API 使用示例

查看 [README.md](README.md) 文件获取详细的 API 使用示例和参数说明。

## 开发指南

### 本地开发环境设置

1. **安装 Go 1.24+**
```bash
# 验证 Go 版本
go version
```

2. **安装依赖**
```bash
go mod download
```

3. **设置环境变量**
```bash
cp .env.example .env
# 编辑 .env 文件配置数据库连接
```

4. **运行数据库迁移**
```bash
# 确保 PostgreSQL 已启动
psql -U oxogame -d oxogame_db -f scripts/init.sql
```

5. **启动开发服务器**
```bash
go run cmd/main.go
```

### 代码结构说明

#### API 层 (`/api`)
负责处理 HTTP 请求和响应，主要功能：
- 参数验证
- 调用服务层
- 格式化响应

#### 服务层 (`/services`)
包含业务逻辑，主要功能：
- 数据验证
- 业务规则实现
- 事务管理
- 调用数据层

#### 模型层 (`/models`)
定义数据结构和数据库映射：
- GORM 模型定义
- 关联关系
- 数据验证规则

#### 中间件 (`/middleware`)
- 请求日志记录
- CORS 处理
- 错误恢复

### 添加新功能

1. 在 `models/` 中定义数据模型
2. 在 `services/` 中实现业务逻辑
3. 在 `api/` 中创建 API 处理器
4. 在 `routes/router.go` 中注册路由
5. 更新数据库迁移脚本

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
DB_HOST=localhost
DB_PORT=5432
DB_USER=oxogame
DB_PASSWORD=oxogame123
DB_NAME=oxogame_db
```

### Docker Compose 配置

`docker-compose.yml` 包含两个服务：
- **api**: Go 应用服务
- **postgres**: PostgreSQL 数据库

### 数据库初始化

`scripts/init.sql` 包含：
- 表结构创建
- 索引创建
- 默认数据插入
- 初始等级和示例玩家

## 测试

### 运行集成测试

```bash
# 使用 Postman 集合
导入 "PostMan 测试.postman_collection.json"

# 或使用 curl 命令
查看各个测试报告文档中的完整测试命令
```

### 性能测试

```bash
# 使用 Apache Bench
ab -n 1000 -c 10 http://localhost:8080/api/v1/players

# 使用 wrk
wrk -t12 -c400 -d30s http://localhost:8080/api/v1/players
```

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查 PostgreSQL 是否运行
   - 验证数据库凭据
   - 确认网络连接

2. **端口被占用**
   ```bash
   # 查找占用端口的进程
   lsof -i :8080
   # 或更改 .env 中的 PORT
   ```

3. **Docker 构建失败**
   - 确保 Docker 守护进程运行中
   - 检查 Dockerfile 语法
   - 清理 Docker 缓存：`docker system prune`

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
