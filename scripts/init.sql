-- 创建 UUID 扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 插入默认等级数据
INSERT INTO levels (id, name, created_at) VALUES 
    (gen_random_uuid(), '初级玩家', NOW()),
    (gen_random_uuid(), '中级玩家', NOW()),
    (gen_random_uuid(), '高级玩家', NOW()),
    (gen_random_uuid(), '大师级玩家', NOW()),
    (gen_random_uuid(), '传奇玩家', NOW())
ON CONFLICT (name) DO NOTHING;

-- 初始化奖池
INSERT INTO prize_pools (id, current_amount, updated_at) VALUES 
    (gen_random_uuid(), 0, NOW())
ON CONFLICT DO NOTHING;