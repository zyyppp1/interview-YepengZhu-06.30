# 游戏房间管理系统 (Game Room Management System) - 测试报告

## 1. API端点测试结果

### `/rooms` 端点

#### `GET`: 列出所有游戏房间
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/rooms
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "游戏室A",
      "description": "适合初学者的房间",
      "status": "available",
      "max_players": 4,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    },
    {
      "id": 2,
      "name": "游戏室B",
      "description": "中级玩家专用房间",
      "status": "available",
      "max_players": 6,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    },
    {
      "id": 3,
      "name": "游戏室C",
      "description": "高级玩家竞技房间",
      "status": "available",
      "max_players": 8,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    },
    {
      "id": 4,
      "name": "VIP包厢",
      "description": "私人定制房间",
      "status": "available",
      "max_players": 2,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    },
    {
      "id": 5,
      "name": "训练室",
      "description": "新手练习专用",
      "status": "maintenance",
      "max_players": 10,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    }
  ]
}
```


#### `GET`: 按状态筛选房间
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/rooms?status=available"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "游戏室A",
      "description": "适合初学者的房间",
      "status": "available",
      "max_players": 4,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    },
    {
      "id": 2,
      "name": "游戏室B",
      "description": "中级玩家专用房间",
      "status": "available",
      "max_players": 6,
      "created_at": "2025-06-30T08:16:53.542614+08:00",
      "updated_at": "2025-06-30T08:16:53.542614+08:00"
    }
  ]
}
```

#### `POST`: 新增游戏房间
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/rooms" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "竞技场",
    "description": "高端竞技专用房间",
    "max_players": 8
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 6,
    "name": "竞技场",
    "description": "高端竞技专用房间",
    "status": "available",
    "max_players": 8,
    "created_at": "2025-06-30T09:15:30.123456+08:00",
    "updated_at": "2025-06-30T09:15:30.123456+08:00"
  }
}
```

### `/rooms/{id}` 端点

#### `GET`: 获取特定房间详细信息
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/rooms/1
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "游戏室A",
    "description": "适合初学者的房间",
    "status": "available",
    "max_players": 4,
    "created_at": "2025-06-30T08:16:53.542614+08:00",
    "updated_at": "2025-06-30T08:16:53.542614+08:00"
  }
}
```

#### `PUT`: 更新房间信息
**Test case:**
```bash
curl -i -X PUT "http://localhost:8080/api/v1/rooms/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "游戏室A升级版",
    "description": "升级后的初学者房间",
    "status": "occupied",
    "max_players": 6
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "游戏室A升级版",
    "description": "升级后的初学者房间",
    "status": "occupied",
    "max_players": 6,
    "created_at": "2025-06-30T08:16:53.542614+08:00",
    "updated_at": "2025-06-30T09:20:45.987654+08:00"
  }
}
```

#### `DELETE`: 删除房间
**Test case:**
```bash
curl -i -X DELETE "http://localhost:8080/api/v1/rooms/6"
```

**Result:**
```json
{
  "success": true,
  "message": "Room deleted successfully"
}
```

### `/reservations` 端点

#### `GET`: 查询所有预约
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/reservations
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "player_id": 1,
      "reservation_date": "2025-07-01",
      "start_time": "14:00",
      "end_time": "16:00",
      "status": "active",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "room": {
        "id": 1,
        "name": "游戏室A升级版",
        "description": "升级后的初学者房间",
        "status": "occupied",
        "max_players": 6
      },
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    },
    {
      "id": 2,
      "room_id": 2,
      "player_id": 2,
      "reservation_date": "2025-07-02",
      "start_time": "10:00",
      "end_time": "12:00",
      "status": "active",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "room": {
        "id": 2,
        "name": "游戏室B",
        "description": "中级玩家专用房间",
        "status": "available",
        "max_players": 6
      },
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    }
  ]
}
```

#### `GET`: 按房间ID查询预约
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/reservations?room_id=1"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "player_id": 1,
      "reservation_date": "2025-07-01",
      "start_time": "14:00",
      "end_time": "16:00",
      "status": "active",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "room": {
        "id": 1,
        "name": "游戏室A升级版"
      },
      "player": {
        "id": 1,
        "name": "张三"
      }
    }
  ]
}
```

#### `GET`: 按日期查询预约
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/reservations?date=2025-07-01"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "player_id": 1,
      "reservation_date": "2025-07-01",
      "start_time": "14:00",
      "end_time": "16:00",
      "status": "active",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "room": {
        "id": 1,
        "name": "游戏室A升级版"
      },
      "player": {
        "id": 1,
        "name": "张三"
      }
    }
  ]
}
```

#### `GET`: 限制返回数量
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/reservations?limit=1"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "player_id": 1,
      "reservation_date": "2025-07-01",
      "start_time": "14:00",
      "end_time": "16:00",
      "status": "active",
      "created_at": "2025-06-30T08:16:53.543123+08:00",
      "updated_at": "2025-06-30T08:16:53.543123+08:00",
      "room": {
        "id": 1,
        "name": "游戏室A升级版"
      },
      "player": {
        "id": 1,
        "name": "张三"
      }
    }
  ]
}
```

#### `POST`: 新增预约
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/reservations" \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 2,
    "player_id": 3,
    "reservation_date": "2025-07-03T00:00:00+08:00",
    "start_time": "18:00",
    "end_time": "20:00"
  }'
```

**Result:**
```json

{
    "data": {
        "id": 34,
        "room_id": 2,
        "player_id": 3,
        "reservation_date": "2025-07-03T00:00:00Z",
        "start_time": "18:00",
        "end_time": "20:00",
        "status": "active",
        "created_at": "2025-06-30T09:48:51.437722+08:00",
        "updated_at": "2025-06-30T09:48:51.437722+08:00",
        "room": {
            "id": 2,
            "name": "游戏室B",
            "description": "中级玩家专用房间",
            "status": "available",
            "max_players": 6,
            "created_at": "2025-06-30T08:16:53.542614+08:00",
            "updated_at": "2025-06-30T08:16:53.542614+08:00"
        },
        "player": {
            "id": 3,
            "name": "王五",
            "level_id": 3,
            "balance": 300,
            "created_at": "2025-06-30T08:16:53.541858+08:00",
            "updated_at": "2025-06-30T08:16:53.541858+08:00"
        }
    },
    "success": true
}
```

## 2. 错误情况测试

### 房间不存在 (404 Not Found)
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/rooms/999
```

**Result:**
```json
{
  "error": "Room not found"
}
```

### 无效的房间数据 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/rooms" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "max_players": -1
  }'
```

**Result:**
```json
{
  "error": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
```

### 无效的房间状态 (400 Bad Request)
**Test case:**
```bash
curl -i -X PUT "http://localhost:8080/api/v1/rooms/1" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "invalid_status"
  }'
```

**Result:**
```json
{
  "error": "invalid status value"
}
```

### 时间冲突的预约 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/reservations" \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 1,
    "player_id": 2,
    "reservation_date": "2025-07-01",
    "start_time": "15:00",
    "end_time": "17:00"
  }'
```

**Result:**
```json
{
  "error": "time slot already reserved"
}
```

### 无效的时间格式 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/reservations" \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 1,
    "player_id": 2,
    "reservation_date": "2025-07-01",
    "start_time": "25:00",
    "end_time": "26:00"
  }'
```

**Result:**
```json
{
  "error": "Invalid time format, use HH:MM"
}
```

### 结束时间早于开始时间 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/reservations" \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 1,
    "player_id": 2,
    "reservation_date": "2025-07-01",
    "start_time": "18:00",
    "end_time": "16:00"
  }'
```

**Result:**
```json
{
  "error": "End time must be after start time"
}
```

## 3. HTTP 状态码总结

| 状态码 | 含义 | 使用场景 |
|--------|------|----------|
| 200 OK | 成功 | 获取房间列表、获取单个房间、更新房间、删除房间、获取预约列表 |
| 201 Created | 创建成功 | 创建新房间、创建新预约 |
| 400 Bad Request | 请求错误 | 参数格式错误、时间格式错误、时间冲突、无效状态值 |
| 404 Not Found | 资源不存在 | 房间不存在、玩家不存在 |
| 500 Internal Server Error | 服务器错误 | 数据库连接失败等 |


## 4. 完整测试命令集

```bash
# 房间管理测试
curl -i http://localhost:8080/api/v1/rooms
curl -i "http://localhost:8080/api/v1/rooms?status=available"
curl -i -X POST "http://localhost:8080/api/v1/rooms" -H "Content-Type: application/json" -d '{"name": "测试房间", "description": "测试用途", "max_players": 4}'
curl -i http://localhost:8080/api/v1/rooms/1
curl -i -X PUT "http://localhost:8080/api/v1/rooms/1" -H "Content-Type: application/json" -d '{"status": "occupied"}'
curl -i -X DELETE "http://localhost:8080/api/v1/rooms/6"

# 预约管理测试
curl -i http://localhost:8080/api/v1/reservations
curl -i "http://localhost:8080/api/v1/reservations?room_id=1"
curl -i "http://localhost:8080/api/v1/reservations?date=2025-07-01"
curl -i "http://localhost:8080/api/v1/reservations?limit=2"
curl -i -X POST "http://localhost:8080/api/v1/reservations" \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 2,
    "player_id": 3,
    "reservation_date": "2025-07-03T00:00:00+08:00",
    "start_time": "18:00",
    "end_time": "20:00"
  }'
# 错误情况测试
curl -i http://localhost:8080/api/v1/rooms/999
curl -i -X POST "http://localhost:8080/api/v1/rooms" -H "Content-Type: application/json" -d '{"name": ""}'
curl -i -X POST "http://localhost:8080/api/v1/reservations" -H "Content-Type: application/json" -d '{"room_id": 1, "player_id": 2, "reservation_date": "2025-07-01", "start_time": "25:00", "end_time": "26:00"}'
```