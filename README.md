docker compose up -d 启动

测试接口：
curl http://localhost:8080/health


API文档：
Base URL: http://localhost:8080/api/v1

🎮 玩家管理 API
1. 获取玩家列表

http
GET /api/v1/players

查询参数：
参数类型必需默认值描述pageinteger否1页码，从1开始page_sizeinteger否20每页数量，1-100

请求示例：
bash
curl -X GET "http://localhost:8080/api/v1/players?page=1&page_size=5"

响应示例：
json{
  "success": true,
  "data": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "name": "张三",
      "level_id": "11111111-2222-3333-4444-555555555555",
      "level": {
        "id": "11111111-2222-3333-4444-555555555555",
        "name": "初级玩家",
        "created_at": "2025-06-30T10:00:00+08:00"
      },
      "balance": 100.0,
      "created_at": "2025-06-30T10:00:00+08:00",
      "updated_at": "2025-06-30T10:00:00+08:00"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 5,
    "total": 25,
    "total_pages": 5
  }
}

状态码：
200 OK - 成功获取数据
500 Internal Server Error - 服务器错误


2. 创建玩家

http
POST /api/v1/players

请求体：
json{
  "name": "新玩家",
  "level_id": "11111111-2222-3333-4444-555555555555"
}
字段说明：
字段类型必需限制描述namestring是2-50字符玩家姓名，必须唯一level_idUUID是有效UUID等级ID，必须存在
请求示例：
bashcurl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试玩家",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c"
  }'
成功响应（201 Created）：
json{
  "success": true,
  "data": {
    "id": "new-uuid-generated",
    "name": "测试玩家",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
    "level": {
      "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "name": "初级玩家"
    },
    "balance": 0,
    "created_at": "2025-06-30T10:00:00+08:00",
    "updated_at": "2025-06-30T10:00:00+08:00"
  },
  "message": "Player created successfully"
}
错误响应：
400 Bad Request - 请求格式错误：
json{
  "error": "Invalid request format",
  "details": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
409 Conflict - 玩家名称重复：
json{
  "error": "Player name already exists",
  "code": "DUPLICATE_NAME"
}
422 Unprocessable Entity - 等级ID无效：
json{
  "error": "Invalid level ID provided",
  "code": "INVALID_LEVEL"
}

3. 获取单个玩家
httpGET /api/v1/players/{id}
路径参数：
参数类型描述idUUID玩家唯一标识符
请求示例：
bashcurl -X GET "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
成功响应（200 OK）：
json{
  "success": true,
  "data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "张三",
    "level_id": "11111111-2222-3333-4444-555555555555",
    "level": {
      "id": "11111111-2222-3333-4444-555555555555",
      "name": "初级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    "balance": 100.0,
    "created_at": "2025-06-30T10:00:00+08:00",
    "updated_at": "2025-06-30T10:00:00+08:00"
  }
}
错误响应：
400 Bad Request - ID格式错误：
json{
  "error": "Invalid player ID format",
  "code": "INVALID_UUID"
}
404 Not Found - 玩家不存在：
json{
  "error": "Player not found",
  "code": "PLAYER_NOT_FOUND"
}

4. 更新玩家信息
httpPUT /api/v1/players/{id}
请求体（部分更新）：
json{
  "name": "新名字",
  "level_id": "new-level-uuid"
}
字段说明：
字段类型必需描述namestring否新的玩家姓名level_idUUID否新的等级ID
请求示例：
bashcurl -X PUT "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三改名",
    "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375"
  }'
成功响应（200 OK）：
json{
  "success": true,
  "data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "张三改名",
    "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
    "level": {
      "id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
      "name": "中级玩家"
    },
    "balance": 100.0,
    "created_at": "2025-06-30T10:00:00+08:00",
    "updated_at": "2025-06-30T11:00:00+08:00"
  }
}
状态码：

200 OK - 更新成功
400 Bad Request - 请求格式错误
404 Not Found - 玩家不存在
422 Unprocessable Entity - 业务逻辑错误


5. 删除玩家
httpDELETE /api/v1/players/{id}
请求示例：
bashcurl -X DELETE "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890"
成功响应（204 No Content）：
无响应体
错误响应：
400 Bad Request - ID格式错误：
json{
  "error": "Invalid player ID format",
  "code": "INVALID_UUID"
}
404 Not Found - 玩家不存在：
json{
  "error": "Player not found",
  "code": "PLAYER_NOT_FOUND"
}

🏆 等级管理 API
1. 获取等级列表
httpGET /api/v1/levels
请求示例：
bashcurl -X GET "http://localhost:8080/api/v1/levels"
成功响应（200 OK）：
json{
  "success": true,
  "data": [
    {
      "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "name": "初级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
      "name": "中级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "3bf1dc65-5312-4b76-b513-fdc4b541086a",
      "name": "高级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "6fa4b5d5-0bbe-4b1b-9481-a35f1257bdba",
      "name": "大师级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "5147c297-5faa-488d-b3ad-bc0600af620a",
      "name": "传奇玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    }
  ]
}

2. 创建等级
httpPOST /api/v1/levels
请求体：
json{
  "name": "超级玩家"
}
字段说明：
字段类型必需限制描述namestring是2-30字符等级名称，必须唯一
请求示例：
bashcurl -X POST "http://localhost:8080/api/v1/levels" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "超级玩家"
  }'
成功响应（201 Created）：
json{
  "success": true,
  "data": {
    "id": "new-generated-uuid",
    "name": "超级玩家",
    "created_at": "2025-06-30T12:00:00+08:00"
  }
}
错误响应：
400 Bad Request - 请求格式错误：
json{
  "error": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
409 Conflict - 等级名称重复：
json{
  "error": "Level name already exists",
  "code": "DUPLICATE_LEVEL_NAME"
}

🧪 API测试示例
完整的测试流程
1. 获取所有等级：
bashcurl http://localhost:8080/api/v1/levels
2. 创建新玩家：
bashcurl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试玩家001",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c"
  }'
3. 获取玩家列表：
bashcurl "http://localhost:8080/api/v1/players?page=1&page_size=10"
4. 获取特定玩家（使用第2步返回的ID）：
bashcurl "http://localhost:8080/api/v1/players/[PLAYER_ID]"
5. 更新玩家信息：
bashcurl -X PUT "http://localhost:8080/api/v1/players/[PLAYER_ID]" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试玩家001改名"
  }'
6. 删除玩家：
bashcurl -X DELETE "http://localhost:8080/api/v1/players/[PLAYER_ID]"

📋 HTTP状态码总结
状态码含义使用场景200 OK成功获取数据、更新成功201 Created创建成功新建玩家、等级204 No Content成功但无内容删除操作成功400 Bad Request请求错误参数格式错误、JSON格式错误404 Not Found资源不存在玩家/等级不存在409 Conflict资源冲突名称重复422 Unprocessable Entity业务逻辑错误等级ID无效、余额不足500 Internal Server Error服务器错误数据库连接失败等

🔧 改进建议实施
需要修改的文件：

api/player.go - 添加更精确的状态码处理
services/player_service.go - 返回更具体的错误类型
新增错误处理工具函数 - 统一错误响应格式

示例改进代码：
go// 错误处理工具函数
func handleServiceError(c *gin.Context, err error, operation string) {
    if strings.Contains(err.Error(), "not found") {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("%s not found", operation),
            "code": "RESOURCE_NOT_FOUND",
        })
    } else if strings.Contains(err.Error(), "duplicate") {
        c.JSON(http.StatusConflict, gin.H{
            "error": fmt.Sprintf("%s already exists", operation),
            "code": "RESOURCE_CONFLICT",
        })
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to %s", operation),
        })
    }
}


1.


玩家和等级管理 API 文档
基础信息

Base URL: http://localhost:8080/api/v1
Content-Type: application/json
时区: Asia/Shanghai (UTC+8)


🎮 玩家管理 API
1. 获取玩家列表
httpGET /api/v1/players
查询参数：
参数类型必需默认值描述pageinteger否1页码，从1开始page_sizeinteger否20每页数量，范围1-100
请求示例：
bash# 获取第一页玩家（默认20个）
curl http://localhost:8080/api/v1/players

# 分页获取玩家列表
curl "http://localhost:8080/api/v1/players?page=1&page_size=5"

# 获取第二页
curl "http://localhost:8080/api/v1/players?page=2&page_size=10"
成功响应（200 OK）：
json{
  "success": true,
  "data": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "name": "张三",
      "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "level": {
        "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
        "name": "初级玩家",
        "created_at": "2025-06-30T10:00:00+08:00"
      },
      "balance": 100.0,
      "created_at": "2025-06-30T10:00:00+08:00",
      "updated_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "b2c3d4e5-f6g7-8901-bcde-f234567890ab",
      "name": "李四",
      "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
      "level": {
        "id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
        "name": "中级玩家",
        "created_at": "2025-06-30T10:00:00+08:00"
      },
      "balance": 200.0,
      "created_at": "2025-06-30T10:00:00+08:00",
      "updated_at": "2025-06-30T10:00:00+08:00"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 5,
    "total": 25,
    "total_pages": 5
  }
}
错误响应：
json{
  "error": "Internal server error"
}

2. 创建新玩家
httpPOST /api/v1/players
请求体：
json{
  "name": "新玩家姓名",
  "level_id": "等级UUID"
}
字段说明：
字段类型必需限制描述namestring是2-50字符玩家姓名，必须唯一level_idUUID是有效UUID等级ID，必须在系统中存在
请求示例：
bash# 创建初级玩家
curl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新手玩家001",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c"
  }'

# 创建中级玩家
curl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "进阶玩家002",
    "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375"
  }'
成功响应（201 Created）：
json{
  "success": true,
  "data": {
    "id": "f9e8d7c6-b5a4-9384-7162-50394857263b",
    "name": "新手玩家001",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
    "level": {
      "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "name": "初级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    "balance": 0,
    "created_at": "2025-06-30T12:00:00+08:00",
    "updated_at": "2025-06-30T12:00:00+08:00"
  },
  "message": "Player created successfully"
}
错误响应：
400 Bad Request - 请求格式错误：
json{
  "error": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
400 Bad Request - 字段格式错误：
json{
  "error": "Invalid UUID format for level_id"
}

3. 获取单个玩家详情
httpGET /api/v1/players/{id}
路径参数：
参数类型描述idUUID玩家的唯一标识符
请求示例：
bash# 获取特定玩家信息
curl http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890

# 获取不存在的玩家（测试404）
curl http://localhost:8080/api/v1/players/99999999-9999-9999-9999-999999999999
成功响应（200 OK）：
json{
  "success": true,
  "data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "张三",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
    "level": {
      "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "name": "初级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    "balance": 100.0,
    "created_at": "2025-06-30T10:00:00+08:00",
    "updated_at": "2025-06-30T10:00:00+08:00"
  }
}
错误响应：
400 Bad Request - ID格式错误：
json{
  "error": "Invalid UUID format"
}
404 Not Found - 玩家不存在：
json{
  "error": "Player not found"
}

4. 更新玩家信息
httpPUT /api/v1/players/{id}
请求体（支持部分更新）：
json{
  "name": "新的玩家姓名",
  "level_id": "新的等级UUID"
}
字段说明：
字段类型必需描述namestring否新的玩家姓名（如果提供）level_idUUID否新的等级ID（如果提供）
请求示例：
bash# 只更新姓名
curl -X PUT "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三改名"
  }'

# 只更新等级
curl -X PUT "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375"
  }'

# 同时更新姓名和等级
curl -X PUT "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三升级版",
    "level_id": "3bf1dc65-5312-4b76-b513-fdc4b541086a"
  }'
成功响应（200 OK）：
json{
  "success": true,
  "data": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "张三升级版",
    "level_id": "3bf1dc65-5312-4b76-b513-fdc4b541086a",
    "level": {
      "id": "3bf1dc65-5312-4b76-b513-fdc4b541086a",
      "name": "高级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    "balance": 100.0,
    "created_at": "2025-06-30T10:00:00+08:00",
    "updated_at": "2025-06-30T13:00:00+08:00"
  }
}
错误响应：
400 Bad Request - 请求格式错误：
json{
  "error": "Invalid request body"
}
404 Not Found - 玩家不存在：
json{
  "error": "Player not found"
}

5. 删除玩家
httpDELETE /api/v1/players/{id}
请求示例：
bash# 删除指定玩家
curl -X DELETE "http://localhost:8080/api/v1/players/a1b2c3d4-e5f6-7890-abcd-ef1234567890"

# 尝试删除不存在的玩家（测试404）
curl -X DELETE "http://localhost:8080/api/v1/players/99999999-9999-9999-9999-999999999999"
成功响应（200 OK）：
json{
  "success": true,
  "message": "Player deleted successfully"
}
错误响应：
400 Bad Request - ID格式错误：
json{
  "error": "Invalid UUID format"
}
404 Not Found - 玩家不存在：
json{
  "error": "Player not found"
}

🏆 等级管理 API
1. 获取等级列表
httpGET /api/v1/levels
请求示例：
bash# 获取所有等级
curl http://localhost:8080/api/v1/levels
成功响应（200 OK）：
json{
  "success": true,
  "data": [
    {
      "id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c",
      "name": "初级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "abe83cc8-362d-4ea8-9ad7-673f06f10375",
      "name": "中级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "3bf1dc65-5312-4b76-b513-fdc4b541086a",
      "name": "高级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "6fa4b5d5-0bbe-4b1b-9481-a35f1257bdba",
      "name": "大师级玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    },
    {
      "id": "5147c297-5faa-488d-b3ad-bc0600af620a",
      "name": "传奇玩家",
      "created_at": "2025-06-30T10:00:00+08:00"
    }
  ]
}
错误响应：
json{
  "error": "Internal server error"
}

2. 创建新等级
httpPOST /api/v1/levels
请求体：
json{
  "name": "等级名称"
}
字段说明：
字段类型必需限制描述namestring是2-30字符等级名称，必须唯一
请求示例：
bash# 创建新等级
curl -X POST "http://localhost:8080/api/v1/levels" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "超级玩家"
  }'

# 创建另一个等级
curl -X POST "http://localhost:8080/api/v1/levels" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "终极玩家"
  }'

# 测试重复名称（会失败）
curl -X POST "http://localhost:8080/api/v1/levels" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "初级玩家"
  }'
成功响应（201 Created）：
json{
  "success": true,
  "data": {
    "id": "d4c3b2a1-9876-5432-10fe-dcba98765432",
    "name": "超级玩家",
    "created_at": "2025-06-30T14:00:00+08:00"
  }
}
错误响应：
400 Bad Request - 请求格式错误：
json{
  "error": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
400 Bad Request - 等级名称重复：
json{
  "error": "Level name already exists"
}

🧪 完整测试流程示例
场景：创建一个新玩家的完整流程
步骤1：查看可用等级
bashcurl http://localhost:8080/api/v1/levels
步骤2：创建新玩家（使用步骤1获得的等级ID）
bashcurl -X POST "http://localhost:8080/api/v1/players" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试玩家123",
    "level_id": "579cbe11-65ac-4b0a-a7bb-ca2d5ac9889c"
  }'
步骤3：验证玩家已创建
bashcurl http://localhost:8080/api/v1/players
步骤4：获取特定玩家详情（使用步骤2返回的玩家ID）
bashcurl http://localhost:8080/api/v1/players/[PLAYER_ID]
步骤5：更新玩家等级
bashcurl -X PUT "http://localhost:8080/api/v1/players/[PLAYER_ID]" \
  -H "Content-Type: application/json" \
  -d '{
    "level_id": "abe83cc8-362d-4ea8-9ad7-673f06f10375"
  }'
步骤6：删除玩家
bashcurl -X DELETE "http://localhost:8080/api/v1/players/[PLAYER_ID]"
步骤7：验证玩家已删除
bashcurl http://localhost:8080/api/v1/players/[PLAYER_ID]
# 应该返回404 Not Found

📋 状态码总结
状态码含义使用场景200 OK成功获取数据、更新成功、删除成功201 Created创建成功新建玩家、新建等级400 Bad Request请求错误参数格式错误、JSON格式错误404 Not Found资源不存在玩家不存在、等级不存在500 Internal Server Error服务器错误数据库连接失败等