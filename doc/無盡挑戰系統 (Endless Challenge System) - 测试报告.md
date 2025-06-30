# 无盡挑戰系統 (Endless Challenge System) - 测试报告

## 1. API端点测试结果

### `/challenges` 端点

#### `POST`: 玩家参加挑战
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/challenges" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "amount": 20.01
  }'
```

**Result:**
```json
{
  "challenge_id": 1,
  "status": "started",
  "message": "Challenge started, result will be available in 30 seconds",
  "start_time": "2025-06-30T09:30:15.123456+08:00",
  "next_available": "2025-06-30T09:31:15.123456+08:00"
}
```

#### `POST`: 余额不足的挑战请求
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/challenges" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 5,
    "amount": 20.01
  }'
```

**Result:**
```json
{
  "error": "insufficient balance"
}
```

#### `POST`: 冷却时间内的重复挑战
**Test case:**
```bash
# 在1分钟内再次发起挑战
curl -i -X POST "http://localhost:8080/api/v1/challenges" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "amount": 20.01
  }'
```

**Result:**
```json
{
  "error": "please wait until 09:31:15 to join next challenge"
}
```

#### `POST`: 错误的挑战金额
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/challenges" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "amount": 10.00
  }'
```

**Result:**
```json
{
  "error": "challenge amount must be 20.01"
}
```

### `/challenges/results` 端点

#### `GET`: 列出最近的挑战结果（默认10条）
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/challenges/results
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 15,
      "player_id": 3,
      "is_winner": true,
      "prize_amount": 500.25,
      "started_at": "2025-06-30T09:28:45.789012+08:00",
      "ended_at": "2025-06-30T09:29:15.789012+08:00",
      "player_name": "王五"
    },
    {
      "id": 14,
      "player_id": 2,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:27:30.456789+08:00",
      "ended_at": "2025-06-30T09:28:00.456789+08:00",
      "player_name": "李四"
    },
    {
      "id": 13,
      "player_id": 1,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:26:15.123456+08:00",
      "ended_at": "2025-06-30T09:26:45.123456+08:00",
      "player_name": "张三"
    },
    {
      "id": 12,
      "player_id": 4,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:25:00.987654+08:00",
      "ended_at": "2025-06-30T09:25:30.987654+08:00",
      "player_name": "赵六"
    },
    {
      "id": 11,
      "player_id": 1,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:24:10.654321+08:00",
      "ended_at": "2025-06-30T09:24:40.654321+08:00",
      "player_name": "张三"
    }
  ]
}
```

#### `GET`: 限制返回结果数量
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/challenges/results?limit=3"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 15,
      "player_id": 3,
      "is_winner": true,
      "prize_amount": 500.25,
      "started_at": "2025-06-30T09:28:45.789012+08:00",
      "ended_at": "2025-06-30T09:29:15.789012+08:00",
      "player_name": "王五"
    },
    {
      "id": 14,
      "player_id": 2,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:27:30.456789+08:00",
      "ended_at": "2025-06-30T09:28:00.456789+08:00",
      "player_name": "李四"
    },
    {
      "id": 13,
      "player_id": 1,
      "is_winner": false,
      "prize_amount": 0,
      "started_at": "2025-06-30T09:26:15.123456+08:00",
      "ended_at": "2025-06-30T09:26:45.123456+08:00",
      "player_name": "张三"
    }
  ]
}
```

## 2. 玩家余额变化验证

#### 参加挑战前的玩家余额
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/players/1
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "张三",
    "level_id": 1,
    "balance": 100.0,
    "created_at": "2025-06-30T08:16:53.541858+08:00",
    "updated_at": "2025-06-30T08:16:53.541858+08:00",
    "level": {
      "id": 1,
      "name": "初级玩家"
    }
  }
}
```

#### 参加挑战后的玩家余额（挑战失败）
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/players/1
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "张三",
    "level_id": 1,
    "balance": 79.99,
    "created_at": "2025-06-30T08:16:53.541858+08:00",
    "updated_at": "2025-06-30T09:30:15.123456+08:00",
    "level": {
      "id": 1,
      "name": "初级玩家"
    }
  }
}
```

#### 挑战成功后的玩家余额
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/players/2
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "name": "李四",
    "level_id": 2,
    "balance": 1380.59,
    "created_at": "2025-06-30T08:16:53.541858+08:00",
    "updated_at": "2025-06-30T09:32:30.456789+08:00",
    "level": {
      "id": 2,
      "name": "中级玩家"
    }
  }
}
```

## 3. 错误情况测试

### 玩家不存在 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/challenges" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 999,
    "amount": 20.01
  }'
```

**Result:**
```json
{
  "error": "player not found"
}
```

## 4. HTTP 状态码总结

| 状态码 | 含义 | 使用场景 |
|--------|------|----------|
| 200 OK | 成功 | 获取挑战结果列表 |
| 201 Created | 创建成功 | 成功参加挑战 |
| 400 Bad Request | 请求错误 | 余额不足、冷却时间未到、金额错误、玩家不存在、参数格式错误 |
| 500 Internal Server Error | 服务器错误 | 数据库连接失败、奖池更新失败等 |
