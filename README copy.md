### 1. 玩家管理系統 (Player Management System)

您的任務是設計並實現一個 RESTful API，用於管理 OXO 系列遊戲中的玩家和等級。

**要求**：

1. Endpoint
    * `/players`:
        * `GET`: 列出所有玩家，返回 JSON 格式的清單，包含每個玩家的 ID、名稱、等級等資訊。
        Test case:
        Result:
        * `POST`: 註冊一個新玩家，接收 JSON 格式的請求，包含玩家的名稱和等級。返回新玩家的 ID。
        Test case:
        Result:
    * `/players/{id}`:
        * `GET`: 獲取特定 ID 的玩家詳細資訊。
        :
        * `PUT`: 更新特定 ID 的玩家資訊。
        :
        * `DELETE`: 刪除特定 ID 的玩家。
        :
    * `/levels`:
        * `GET`: 列出所有等級，返回 JSON 格式的清單，包含每個等級的 ID 和名稱。
        :
        * `POST`: 新增一個等級，接收 JSON 格式的請求，包含等級名稱。返回新等級的 ID。
        :

**答題提示**：

使用合適的 HTTP 狀態碼來表示操作結果:
状态码：

提供詳細的 API 文檔，包括示例請求和回應：



