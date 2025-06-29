docker compose up -d 启动


测试接口：
curl http://localhost:8080/health

	•	列出所有等级（GET /levels）
curl http://localhost:8080/api/v1/levels