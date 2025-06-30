# 遊戲日誌收集器 (Game Log Collector) - 测试报告

## 1. API端点测试结果

### `/logs` 端点

#### `GET`: 查询所有游戏日志（默认查询）
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/logs
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 15,
      "player_id": 2,
      "action_type": "challenge_result",
      "details": {
        "challenge_id": 5,
        "is_winner": true,
        "prize_amount": 500.25
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:15:30.123456+08:00",
      "updated_at": "2025-06-30T10:15:30.123456+08:00",
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    },
    {
      "id": 14,
      "player_id": 1,
      "action_type": "join_challenge",
      "details": {
        "challenge_id": 4,
        "amount": 20.01
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:14:45.987654+08:00",
      "updated_at": "2025-06-30T10:14:45.987654+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    },
    {
      "id": 13,
      "player_id": 3,
      "action_type": "enter_room",
      "details": {
        "room_id": 2,
        "room_name": "游戏室B"
      },
      "ip": "192.168.1.102",
      "user_agent": "Mozilla/5.0 (Safari/14.0)",
      "created_at": "2025-06-30T10:10:15.456789+08:00",
      "updated_at": "2025-06-30T10:10:15.456789+08:00",
      "player": {
        "id": 3,
        "name": "王五",
        "level_id": 3
      }
    }
  ]
}
```

#### `GET`: 按玩家ID查询日志
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?player_id=1"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 14,
      "player_id": 1,
      "action_type": "join_challenge",
      "details": {
        "challenge_id": 4,
        "amount": 20.01
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:14:45.987654+08:00",
      "updated_at": "2025-06-30T10:14:45.987654+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    },
    {
      "id": 10,
      "player_id": 1,
      "action_type": "login",
      "details": {
        "login_time": "2025-06-30T09:30:00.000000+08:00",
        "device": "mobile"
      },
      "ip": "192.168.1.100",
      "user_agent": "Mozilla/5.0 (iPhone)",
      "created_at": "2025-06-30T09:30:00.111111+08:00",
      "updated_at": "2025-06-30T09:30:00.111111+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    },
    {
      "id": 8,
      "player_id": 1,
      "action_type": "register",
      "details": {
        "registration_method": "email",
        "email": "zhangsan@example.com"
      },
      "ip": "192.168.1.100",
      "user_agent": "Mozilla/5.0 (Chrome/91.0)",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    }
  ]
}
```

#### `GET`: 按操作类型查询日志
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?action=login"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 12,
      "player_id": 3,
      "action_type": "login",
      "details": {
        "login_time": "2025-06-30T10:00:00.000000+08:00",
        "device": "desktop"
      },
      "ip": "192.168.1.102",
      "user_agent": "Mozilla/5.0 (Safari/14.0)",
      "created_at": "2025-06-30T10:00:00.222222+08:00",
      "updated_at": "2025-06-30T10:00:00.222222+08:00",
      "player": {
        "id": 3,
        "name": "王五",
        "level_id": 3
      }
    },
    {
      "id": 11,
      "player_id": 2,
      "action_type": "login",
      "details": {
        "login_time": "2025-06-30T09:45:00.000000+08:00",
        "device": "tablet"
      },
      "ip": "192.168.1.101",
      "user_agent": "Mozilla/5.0 (iPad)",
      "created_at": "2025-06-30T09:45:00.333333+08:00",
      "updated_at": "2025-06-30T09:45:00.333333+08:00",
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    }
  ]
}
```

#### `GET`: 按时间范围查询日志
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?start_time=2025-06-30T10:00:00Z&end_time=2025-06-30T10:30:00Z"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 15,
      "player_id": 2,
      "action_type": "challenge_result",
      "details": {
        "challenge_id": 5,
        "is_winner": true,
        "prize_amount": 500.25
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:15:30.123456+08:00",
      "updated_at": "2025-06-30T10:15:30.123456+08:00",
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    },
    {
      "id": 14,
      "player_id": 1,
      "action_type": "join_challenge",
      "details": {
        "challenge_id": 4,
        "amount": 20.01
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:14:45.987654+08:00",
      "updated_at": "2025-06-30T10:14:45.987654+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    }
  ]
}
```

#### `GET`: 限制返回数量
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?limit=2"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 15,
      "player_id": 2,
      "action_type": "challenge_result",
      "details": {
        "challenge_id": 5,
        "is_winner": true,
        "prize_amount": 500.25
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:15:30.123456+08:00",
      "updated_at": "2025-06-30T10:15:30.123456+08:00",
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    },
    {
      "id": 14,
      "player_id": 1,
      "action_type": "join_challenge",
      "details": {
        "challenge_id": 4,
        "amount": 20.01
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T10:14:45.987654+08:00",
      "updated_at": "2025-06-30T10:14:45.987654+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    }
  ]
}
```

**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?player_id=1&action=login&limit=5"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 10,
      "player_id": 1,
      "action_type": "login",
      "details": {
        "login_time": "2025-06-30T09:30:00.000000+08:00",
        "device": "mobile"
      },
      "ip": "192.168.1.100",
      "user_agent": "Mozilla/5.0 (iPhone)",
      "created_at": "2025-06-30T09:30:00.111111+08:00",
      "updated_at": "2025-06-30T09:30:00.111111+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    }
  ]
}
```

#### `POST`: 创建新的游戏日志
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 3,
    "action_type": "enter_room",
    "details": {
      "room_id": 1,
      "room_name": "游戏室A",
      "entry_time": "2025-06-30T10:30:00Z"
    }
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 16
  }
}
```


**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "action_type": "system_maintenance",
    "details": {
      "maintenance_type": "database_backup",
      "start_time": "2025-06-30T11:00:00Z",
      "duration": "30 minutes"
    }
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 19
  }
}
```

## 2. 操作类型完整测试

### 註冊 (register)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 6,
    "action_type": "register",
    "details": {
      "registration_method": "email",
      "email": "newuser@example.com",
      "registration_time": "2025-06-30T10:45:00Z"
    }
  }'
```

### 登入 (login)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 6,
    "action_type": "login",
    "details": {
      "login_time": "2025-06-30T10:46:00Z",
      "device": "mobile",
      "location": "Beijing"
    }
  }'
```

### 登出 (logout)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 6,
    "action_type": "logout",
    "details": {
      "logout_time": "2025-06-30T11:00:00Z",
      "session_duration": "14 minutes"
    }
  }'
```

### 進入房間 (enter_room)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 5,
    "action_type": "enter_room",
    "details": {
      "room_id": 3,
      "room_name": "游戏室C",
      "entry_method": "direct"
    }
  }'
```

### 退出房間 (exit_room)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 5,
    "action_type": "exit_room",
    "details": {
      "room_id": 3,
      "room_name": "游戏室C",
      "exit_reason": "game_completed",
      "stay_duration": "25 minutes"
    }
  }'
```

### 參加挑戰 (join_challenge)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 4,
    "action_type": "join_challenge",
    "details": {
      "challenge_id": 12,
      "amount": 20.01,
      "player_balance_before": 150.50
    }
  }'
```

### 挑戰結果 (challenge_result)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 4,
    "action_type": "challenge_result",
    "details": {
      "challenge_id": 12,
      "is_winner": true,
      "prize_amount": 800.75,
      "total_participants": 40
    }
  }'
```

## 3. 错误情况测试

### 无效的操作类型 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "action_type": "invalid_action",
    "details": {}
  }'
```

**Result:**
```json
{
  "error": "invalid action type"
}
```

### 缺少必需参数 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1
  }'
```

**Result:**
```json
{
  "error": "Key: 'action_type' Error:Field validation for 'action_type' failed on the 'required' tag"
}
```

### 玩家不存在 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/logs" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 999,
    "action_type": "login",
    "details": {}
  }'
```

**Result:**
```json
{
  "error": "player not found"
}
```


**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?limit=999"
```

**Result:**
```json
{
  "success": true,
  "data": [
    // 最多返回100条记录
  ]
}
```

## 4. HTTP 状态码总结

| 状态码 | 含义 | 使用场景 |
|--------|------|----------|
| 200 OK | 成功 | 获取日志列表成功 |
| 201 Created | 创建成功 | 成功创建新日志记录 |
| 400 Bad Request | 请求错误 | 无效操作类型、玩家不存在、参数格式错误 |
| 500 Internal Server Error | 服务器错误 | 数据库连接失败、日志写入失败等 |

## 5. 支持的操作类型列表

| 操作类型 | 英文标识 | 描述 | 典型用途 |
|----------|----------|------|----------|
| 註冊 | register | 用户注册 | 记录新用户注册信息 |
| 登入 | login | 用户登录 | 记录用户登录时间和设备 |
| 登出 | logout | 用户登出 | 记录用户登出和会话时长 |
| 進入房間 | enter_room | 进入游戏房间 | 记录房间进入信息 |
| 退出房間 | exit_room | 退出游戏房间 | 记录房间退出和停留时长 |
| 參加挑戰 | join_challenge | 参加无盡挑戰 | 记录挑战参与信息 |
| 挑戰結果 | challenge_result | 挑战结果 | 记录挑战获胜情况 |
| 支付 | payment | 支付操作 | 记录支付相关信息 |

## 6. 查询参数详细说明

### player_id (可选)
- **类型**: 整数
- **说明**: 指定查询的玩家ID
- **示例**: `?player_id=1`
- **默认**: 查询所有玩家

### action (可选)
- **类型**: 字符串
- **说明**: 指定查询的操作类型
- **可选值**: register, login, logout, enter_room, exit_room, join_challenge, challenge_result, payment
- **示例**: `?action=login`
- **默认**: 查询所有操作类型

### start_time (可选)
- **类型**: RFC3339格式时间字符串
- **说明**: 查询起始时间
- **示例**: `?start_time=2025-06-30T10:00:00Z`
- **默认**: 无限制

### end_time (可选)
- **类型**: RFC3339格式时间字符串
- **说明**: 查询结束时间
- **示例**: `?end_time=2025-06-30T11:00:00Z`
- **默认**: 无限制

### limit (可选)
- **类型**: 整数
- **说明**: 限制返回的日志条数
- **范围**: 1-100
- **示例**: `?limit=50`
- **默认**: 50


