-- scripts/init.sql - 完整的数据库初始化脚本

-- ============================================
-- 1. 创建必要的扩展
-- ============================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ============================================
-- 2. 插入默认等级数据
-- ============================================
INSERT INTO levels (id, name, created_at) VALUES 
    (gen_random_uuid(), '初级玩家', NOW()),
    (gen_random_uuid(), '中级玩家', NOW()),
    (gen_random_uuid(), '高级玩家', NOW()),
    (gen_random_uuid(), '大师级玩家', NOW()),
    (gen_random_uuid(), '传奇玩家', NOW())
ON CONFLICT (name) DO NOTHING;

-- ============================================
-- 3. 创建示例玩家
-- ============================================
DO $$
DECLARE
    beginner_level_id UUID;
    intermediate_level_id UUID;
    advanced_level_id UUID;
BEGIN
    -- 获取等级ID
    SELECT id INTO beginner_level_id FROM levels WHERE name = '初级玩家';
    SELECT id INTO intermediate_level_id FROM levels WHERE name = '中级玩家';
    SELECT id INTO advanced_level_id FROM levels WHERE name = '高级玩家';
    
    -- 插入示例玩家
    INSERT INTO players (id, name, level_id, balance, created_at, updated_at) VALUES
        (gen_random_uuid(), '张三', beginner_level_id, 100.0, NOW(), NOW()),
        (gen_random_uuid(), '李四', intermediate_level_id, 200.0, NOW(), NOW()),
        (gen_random_uuid(), '王五', advanced_level_id, 300.0, NOW(), NOW()),
        (gen_random_uuid(), '赵六', beginner_level_id, 150.0, NOW(), NOW()),
        (gen_random_uuid(), '钱七', intermediate_level_id, 250.0, NOW(), NOW())
    ON CONFLICT (name) DO NOTHING;
END $$;

-- ============================================
-- 4. 创建示例房间
-- ============================================
INSERT INTO rooms (id, name, description, status, max_players, created_at, updated_at) VALUES
    (gen_random_uuid(), '游戏室A', '适合初学者的房间', 'available', 4, NOW(), NOW()),
    (gen_random_uuid(), '游戏室B', '中级玩家专用房间', 'available', 6, NOW(), NOW()),
    (gen_random_uuid(), '游戏室C', '高级玩家竞技房间', 'available', 8, NOW(), NOW()),
    (gen_random_uuid(), 'VIP包厢', '私人定制房间', 'available', 2, NOW(), NOW()),
    (gen_random_uuid(), '训练室', '新手练习专用', 'maintenance', 10, NOW(), NOW())
ON CONFLICT (name) DO NOTHING;

-- ============================================
-- 5. 初始化奖池
-- ============================================
INSERT INTO prize_pools (id, current_amount, last_win_amount, updated_at) VALUES 
    (gen_random_uuid(), 0.0, 0.0, NOW())
ON CONFLICT DO NOTHING;

-- ============================================
-- 6. 创建一些示例预约（可选）
-- ============================================
DO $$
DECLARE
    room_a_id UUID;
    room_b_id UUID;
    player1_id UUID;
    player2_id UUID;
BEGIN
    -- 获取房间和玩家ID
    SELECT id INTO room_a_id FROM rooms WHERE name = '游戏室A' LIMIT 1;
    SELECT id INTO room_b_id FROM rooms WHERE name = '游戏室B' LIMIT 1;
    SELECT id INTO player1_id FROM players WHERE name = '张三' LIMIT 1;
    SELECT id INTO player2_id FROM players WHERE name = '李四' LIMIT 1;
    
    -- 创建示例预约
    IF room_a_id IS NOT NULL AND player1_id IS NOT NULL THEN
        INSERT INTO reservations (id, room_id, player_id, reservation_date, start_time, end_time, status, created_at, updated_at) VALUES
            (gen_random_uuid(), room_a_id, player1_id, CURRENT_DATE + INTERVAL '1 day', '14:00', '16:00', 'active', NOW(), NOW()),
            (gen_random_uuid(), room_b_id, player2_id, CURRENT_DATE + INTERVAL '2 days', '10:00', '12:00', 'active', NOW(), NOW());
    END IF;
END $$;

-- ============================================
-- 7. 插入一些示例日志
-- ============================================
DO $$
DECLARE
    player1_id UUID;
    player2_id UUID;
BEGIN
    SELECT id INTO player1_id FROM players WHERE name = '张三' LIMIT 1;
    SELECT id INTO player2_id FROM players WHERE name = '李四' LIMIT 1;
    
    IF player1_id IS NOT NULL THEN
        INSERT INTO game_logs (id, player_id, action_type, details, ip, user_agent, created_at, updated_at) VALUES
            (gen_random_uuid(), player1_id, 'register', '{"registration_method": "username"}', '192.168.1.100', 'Mozilla/5.0', NOW(), NOW()),
            (gen_random_uuid(), player1_id, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.100', 'Mozilla/5.0', NOW(), NOW());
    END IF;
    
    IF player2_id IS NOT NULL THEN
        INSERT INTO game_logs (id, player_id, action_type, details, ip, user_agent, created_at, updated_at) VALUES
            (gen_random_uuid(), player2_id, 'register', '{"registration_method": "email"}', '192.168.1.101', 'Chrome/91.0', NOW(), NOW()),
            (gen_random_uuid(), player2_id, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.101', 'Chrome/91.0', NOW(), NOW());
    END IF;
END $$;

-- ============================================
-- 8. 输出初始化完成信息
-- ============================================
DO $$
BEGIN
    RAISE NOTICE '===========================================';
    RAISE NOTICE '数据库初始化完成！';
    RAISE NOTICE '已创建：';
    RAISE NOTICE '- % 个等级', (SELECT COUNT(*) FROM levels);
    RAISE NOTICE '- % 个玩家', (SELECT COUNT(*) FROM players);
    RAISE NOTICE '- % 个房间', (SELECT COUNT(*) FROM rooms);
    RAISE NOTICE '- % 个预约', (SELECT COUNT(*) FROM reservations);
    RAISE NOTICE '- % 条日志', (SELECT COUNT(*) FROM game_logs);
    RAISE NOTICE '===========================================';
END $$;