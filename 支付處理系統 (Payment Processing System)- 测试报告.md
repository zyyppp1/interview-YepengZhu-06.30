# 支付處理系統 (Payment Processing System) - 测试报告

## 1. API端点测试结果

### `/payments` 端点

#### `POST`: 信用卡支付
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "credit_card",
    "amount": 100.00,
    "payment_details": {
      "card_number": "4111111111111111",
      "expiry_date": "12/26",
      "cvv": "123",
      "cardholder_name": "张三"
    }
  }'

```

**Result:**
```json
{
  "payment_id": 1,
  "transaction_id": "TXN17195423901",
  "status": "success",
  "message": "Payment completed successfully",
  "processed_at": "2025-06-30T11:30:15.123456+08:00"
}
```

#### `POST`: 银行转账支付
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 2,
    "payment_method": "bank_transfer",
    "amount": 500.00,
    "payment_details": {
      "bank_code": "ICBC",
      "account_number": "1234567890123456",
      "account_name": "李四"
    }
  }'
```

**Result:**
```json
{
  "payment_id": 2,
  "transaction_id": "TXN17195423902",
  "status": "success",
  "message": "Payment completed successfully",
  "processed_at": "2025-06-30T11:32:30.456789+08:00"
}
```

#### `POST`: 第三方支付
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 3,
    "payment_method": "third_party",
    "amount": 200.00,
    "payment_details": {
      "provider": "alipay",
      "account_id": "wangwu@example.com"
    }
  }'
```

**Result:**
```json
{
  "payment_id": 3,
  "transaction_id": "TXN17195423903",
  "status": "success",
  "message": "Payment completed successfully",
  "processed_at": "2025-06-30T11:35:45.789012+08:00"
}
```

#### `POST`: 区块链支付
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 4,
    "payment_method": "blockchain",
    "amount": 300.00,
    "payment_details": {
      "wallet_address": "1A2B3C4D5E6F7G8H9I0J",
      "cryptocurrency": "BTC",
      "network": "mainnet"
    }
  }'
```

**Result:**
```json
{
  "payment_id": 4,
  "transaction_id": "TXN17195423904",
  "status": "success",
  "message": "Payment completed successfully",
  "processed_at": "2025-06-30T11:38:20.345678+08:00"
}
```


### `/payments/{id}` 端点

#### `GET`: 获取成功支付详情
**Test case:**
```bash
curl -i http://localhost:8080/api/v1/payments/1
```

**Result:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "player_id": 1,
    "payment_method": "credit_card",
    "amount": 100.00,
    "currency": "CNY",
    "status": "success",
    "transaction_id": "TXN17195423901",
    "payment_details": {
      "card_number": "4111111111111111",
      "expiry_date": "12/26",
      "cvv": "123",
      "cardholder_name": "张三"
    },
    "created_at": "2025-06-30T11:30:15.123456+08:00",
    "updated_at": "2025-06-30T11:30:15.123456+08:00",
    "player": {
      "id": 1,
      "name": "张三",
      "level_id": 1,
      "balance": 179.99
    }
  }
}

```

## 3. 支付日志验证

### 查看支付相关日志
**Test case:**
```bash
curl -i "http://localhost:8080/api/v1/logs?action=payment&limit=3"
```

**Result:**
```json
{
  "success": true,
  "data": [
    {
      "id": 25,
      "player_id": 1,
      "action_type": "payment",
      "details": {
        "payment_id": 1,
        "amount": 100.00,
        "payment_method": "credit_card",
        "status": "success"
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T11:30:15.123456+08:00",
      "updated_at": "2025-06-30T11:30:15.123456+08:00",
      "player": {
        "id": 1,
        "name": "张三",
        "level_id": 1
      }
    },
    {
      "id": 26,
      "player_id": 2,
      "action_type": "payment",
      "details": {
        "payment_id": 2,
        "amount": 500.00,
        "payment_method": "bank_transfer",
        "status": "success"
      },
      "ip": "127.0.0.1",
      "user_agent": "curl/7.88.1",
      "created_at": "2025-06-30T11:32:30.456789+08:00",
      "updated_at": "2025-06-30T11:32:30.456789+08:00",
      "player": {
        "id": 2,
        "name": "李四",
        "level_id": 2
      }
    }
  ]
}
```

## 4. 错误情况测试

### 玩家不存在 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 999,
    "payment_method": "credit_card",
    "amount": 100.00
  }'
```

**Result:**
```json
{
  "error": "player not found"
}
```

### 无效支付方式 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "invalid_method",
    "amount": 100.00
  }'
```

**Result:**
```json
{
  "error": "invalid payment method"
}
```

### 缺少必需参数 (400 Bad Request)
**Test case:**
```bash
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "credit_card"
  }'
```

**Result:**
```json
{
  "error": "Key: 'amount' Error:Field validation for 'amount' failed on the 'required' tag"
}
```

### 

## 5. HTTP 状态码总结

| 状态码 | 含义 | 使用场景 |
|--------|------|----------|
| 200 OK | 成功 | 获取支付详情成功 |
| 201 Created | 创建成功 | 成功处理支付（无论成功失败都创建了支付记录） |
| 400 Bad Request | 请求错误 | 玩家不存在、无效支付方式、无效金额、参数格式错误 |
| 404 Not Found | 资源不存在 | 支付记录不存在 |
| 500 Internal Server Error | 服务器错误 | 数据库连接失败、支付网关异常等 |

## 6. 支付方式详细说明

### 支持的支付方式

#### 1. 信用卡支付 (credit_card)

**支付详情字段:**
```json
{
  "card_number": "4111111111111111",
  "expiry_date": "12/26",
  "cvv": "123",
  "cardholder_name": "持卡人姓名"
}
```

#### 2. 银行转账 (bank_transfer)

**支付详情字段:**
```json
{
  "bank_code": "ICBC",
  "account_number": "1234567890123456",
  "account_name": "账户持有人姓名"
}
```

#### 3. 第三方支付 (third_party)

**支付详情字段:**
```json
{
  "provider": "alipay|wechat|paypal",
  "account_id": "用户账号ID"
}
```

#### 4. 区块链支付 (blockchain)

**支付详情字段:**
```json
{
  "wallet_address": "区块链钱包地址",
  "cryptocurrency": "BTC|ETH|USDT",
  "network": "mainnet|testnet"
}
```

## 7. 支付状态说明

| 状态 | 英文标识 | 描述 | 后续操作 |
|------|----------|------|----------|
| 处理中 | processing | 支付正在处理 | 等待结果 |
| 成功 | success | 支付成功完成 | 增加玩家余额 |
| 失败 | failed | 支付失败 | 不变更余额 |
| 已退款 | refunded | 支付已退款 | 扣除相应余额 |


## 8. 完整测试命令集

```bash
# 基础支付功能测试
echo "=== 各支付方式测试 ==="

# 信用卡支付
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "credit_card",
    "amount": 100.00,
    "payment_details": {
      "card_number": "4111111111111111",
      "expiry_date": "12/26",
      "cvv": "123"
    }
  }'

# 银行转账
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 2,
    "payment_method": "bank_transfer",
    "amount": 500.00,
    "payment_details": {
      "bank_code": "ICBC",
      "account_number": "1234567890"
    }
  }'

# 第三方支付
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 3,
    "payment_method": "third_party",
    "amount": 200.00,
    "payment_details": {
      "provider": "alipay",
      "account_id": "user@example.com"
    }
  }'

# 区块链支付
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 4,
    "payment_method": "blockchain",
    "amount": 300.00,
    "payment_details": {
      "wallet_address": "1A2B3C4D5E6F7G8H",
      "cryptocurrency": "BTC"
    }
  }'

echo "=== 查询支付详情测试 ==="
curl -i http://localhost:8080/api/v1/payments/1
curl -i http://localhost:8080/api/v1/payments/2

echo "=== 错误情况测试 ==="
# 玩家不存在
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 999,
    "payment_method": "credit_card",
    "amount": 100.00
  }'

# 无效支付方式
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "invalid_method",
    "amount": 100.00
  }'

# 无效金额
curl -i -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "credit_card",
    "amount": 0
  }'

# 支付不存在
curl -i http://localhost:8080/api/v1/payments/999

echo "=== 余额变化验证 ==="
curl -i http://localhost:8080/api/v1/players/1
curl -i http://localhost:8080/api/v1/players/2

echo "=== 支付日志验证 ==="
curl -i "http://localhost:8080/api/v1/logs?action=payment&limit=5"
```

## 9. 集成测试验证

### 端到端支付流程测试
```bash
echo "=== 完整支付流程测试 ==="

# 1. 检查初始余额
echo "1. 检查玩家初始余额:"
initial_balance=$(curl -s "http://localhost:8080/api/v1/players/1" | jq '.data.balance')
echo "Initial balance: $initial_balance"

# 2. 执行支付
echo "2. 执行支付:"
payment_response=$(curl -s -X POST "http://localhost:8080/api/v1/payments" \
  -H "Content-Type: application/json" \
  -d '{
    "player_id": 1,
    "payment_method": "credit_card",
    "amount": 50.00
  }')

echo "Payment response: $payment_response"
payment_id=$(echo $payment_response | jq -r '.payment_id')
payment_status=$(echo $payment_response | jq -r '.status')

# 3. 检查支付后余额
echo "3. 检查支付后余额:"
final_balance=$(curl -s "http://localhost:8080/api/v1/players/1" | jq '.data.balance')
echo "Final balance: $final_balance"