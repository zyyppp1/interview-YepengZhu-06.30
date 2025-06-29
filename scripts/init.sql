-- scripts/init.sql

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

-- ================================================
-- 新增默认玩家张三、李四、王五 与关联等级
-- ================================================

-- 获取部分等级的 ID，方便 PLAYER 插入
WITH level_ids AS (
  SELECT name, id FROM levels WHERE name IN ('初级玩家','中级玩家','高级玩家')
)
INSERT INTO players (id, name, level_id, balance, created_at, updated_at)
SELECT gen_random_uuid(), u.name, l.id,
       CASE u.name
         WHEN '张三' THEN 100
         WHEN '李四' THEN 200
         WHEN '王五' THEN 300
       END AS balance,
       NOW(), NOW()
FROM (VALUES ('张三'), ('李四'), ('王五')) AS u(name)
JOIN level_ids AS l ON
     (u.name = '张三' AND l.name = '初级玩家')
  OR (u.name = '李四' AND l.name = '中级玩家')
  OR (u.name = '王五' AND l.name = '高级玩家')
ON CONFLICT (name) DO NOTHING;

-- ================================================
-- 初始化房间示例：一个由“张三”创建的房间
-- ================================================

DO $$
DECLARE
  p_id UUID;
  r_id UUID := gen_random_uuid();
BEGIN
  SELECT id INTO p_id FROM players WHERE name = '张三';

  IF p_id IS NOT NULL THEN
    INSERT INTO rooms (id, owner_id, status, created_at, updated_at)
    VALUES (r_id, p_id, 'open', NOW(), NOW())
    ON CONFLICT DO NOTHING;

    -- 可选：将张三自动加入到房间成员表
    INSERT INTO room_members (room_id, player_id, joined_at)
    VALUES (r_id, p_id, NOW())
    ON CONFLICT DO NOTHING;
  END IF;
END;
$$;