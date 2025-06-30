# 玩家管理系统 (Player Management System) - 测试报告

## 1. API端点测试结果

### `/players` 端点

#### `GET`: 列出所有玩家
**Test case:**
```bash
curl http://localhost:8080/api/v1/players
```
**Result:**
```json
{
  "data": [
    {
      "id": 1,
      "name": "张三",
      "level_id": 1,
      "balance": 100,
      "created_at": "2025-06-30T08:16:53.541858+08:00",
      "updated_at": "2025-06-30T08:16:53.541858+08:00",
      "level": {
        "id": 1,
        "name": "初级玩家",
        "created_at": "2025-06-30T08:16:53.541326+08:00"
      }
    },
    {
      "id": 2,
      "name": "李四",
      "level_id": 2,
      "balance": 200,
      "created_at": "2025-06-30T08:16:53.541858+08:00",
      "updated_at": "2025-06-30T08:16:53.541858+08:00",
      "level": {
        "id": 2,
        "name": "中级玩家",
        "created_at": "2025-06-30T08:16:53.541326+08:00"
      }
    },
    {
      "id": 3,
      "name": "王五",
      "level_id": 3,
      "balance": 300,
      "created_at": "2025-06-30T08:16:53.541858+08:00",
      "updated_at": "2025-06-30T08:16:53.541858+08:00",
      "level": {
        "id": 3,
        "name": "高级玩家",
        "created_at": "2025-06-30T08:16:53.541326+08:00"
      }
    },
    {
      "id": 4,
      "name": "赵六",
      "level_id": 1,
      "balance": 150,
      "created_at": "2025-06-30T08:16:53.541858+08:00",
      "updated_at": "2025-06-30T08:16:53.541858+08:00",
      "level": {
        "id": 1,
        "name": "初级玩家",
        "created_at": "2025-06-30T08:16:53.541326+08:00"
      }
    },
    {
      "id": 5,
      "name": "钱七",
      "level_id": 2,
      "balance": 250,
      "created_at": "2025-06-30T08:16:53.541858+08:00",
      "updated_at": "2025-06-30T08:16:53.541858+08:00",
      "level": {
        "id": 2,
        "name": "中级玩家",
        "created_at": "2025-06-30T08:16:53.541326+08:00"
      }
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 20,
    "total": 5,
    "total_pages": 1
  },
  "success": true
}
```

#### `POST`: 註冊一個新玩家
**Test case:**
```bash
curl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新玩家001",
    "level_id": 1
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 6,
    "name": "新玩家001",
    "level_id": 1,
    "balance": 0,
    "created_at": "2025-06-30T08:20:15.123456+08:00",
    "updated_at": "2025-06-30T08:20:15.123456+08:00",
    "level": {
      "id": 1,
      "name": "初级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    }
  }
}
```

### `/players/{id}` 端点

#### `GET`: 獲取特定 ID 的玩家詳細資訊
**Test case:**
```bash
curl http://localhost:8080/api/v1/players/1
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "张三",
    "level_id": 1,
    "balance": 100,
    "created_at": "2025-06-30T08:16:53.541858+08:00",
    "updated_at": "2025-06-30T08:16:53.541858+08:00",
    "level": {
      "id": 1,
      "name": "初级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    }
  }
}
```

#### `PUT`: 更新特定 ID 的玩家資訊
**Test case:**
```bash
curl -X PUT "http://localhost:8080/api/v1/players/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三改名",
    "level_id": 2
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "张三改名",
    "level_id": 2,
    "balance": 100,
    "created_at": "2025-06-30T08:16:53.541858+08:00",
    "updated_at": "2025-06-30T08:21:30.654321+08:00",
    "level": {
      "id": 2,
      "name": "中级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    }
  }
}
```

#### `DELETE`: 刪除特定 ID 的玩家
**Test case:**
```bash
curl -X DELETE "http://localhost:8080/api/v1/players/1"
```

**Result:**
```json
{
  "success": true,
  "message": "Player deleted successfully"
}
```

### `/levels` 端点

#### `GET`: 列出所有等級
**Test case:**
```bash
curl http://localhost:8080/api/v1/levels
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "初级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    },
    {
      "id": 2,
      "name": "中级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    },
    {
      "id": 3,
      "name": "高级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    },
    {
      "id": 4,
      "name": "大师级玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    },
    {
      "id": 5,
      "name": "传奇玩家",
      "created_at": "2025-06-30T08:16:53.541326+08:00"
    }
  ]
}
```

#### `POST`: 新增一個等級
**Test case:**
```bash
curl -X POST "http://localhost:8080/api/v1/levels" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "超级玩家"
  }'
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 6,
    "name": "超级玩家",
    "created_at": "2025-06-30T08:22:45.789012+08:00"
  }
}
```

## 2. HTTP 状态码使用

### 成功状态码：
- **200 OK** - 获取数据成功、更新成功、删除成功
- **201 Created** - 创建玩家成功、创建等级成功

### 客户端错误状态码：
- **400 Bad Request** - 请求格式错误、参数验证失败
- **404 Not Found** - 玩家不存在、等级不存在
- **409 Conflict** - 玩家名称重复、等级名称重复

### 服务器错误状态码：
- **500 Internal Server Error** - 数据库连接失败、服务器内部错误
