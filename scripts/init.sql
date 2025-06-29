-- ============================================
-- 完整的数据库初始化脚本（修复版）
-- ============================================

-- 创建必要的扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- 1. 创建所有表结构
-- ============================================

-- 创建等级表
CREATE TABLE IF NOT EXISTS levels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 创建玩家表（注意：name字段添加UNIQUE约束）
CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,  -- 添加UNIQUE约束
    level_id UUID NOT NULL REFERENCES levels(id),
    balance DECIMAL(10,2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建房间表（name字段添加UNIQUE约束）
CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,  -- 添加UNIQUE约束
    description TEXT,
    status VARCHAR(20) DEFAULT 'available',
    max_players INTEGER DEFAULT 4,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建预约表
CREATE TABLE IF NOT EXISTS reservations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL REFERENCES rooms(id),
    player_id UUID NOT NULL REFERENCES players(id),
    reservation_date DATE NOT NULL,
    start_time VARCHAR(5) NOT NULL,
    end_time VARCHAR(5) NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建挑战表
CREATE TABLE IF NOT EXISTS challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(id),
    amount DECIMAL(10,2) DEFAULT 20.01,
    is_winner BOOLEAN DEFAULT FALSE,
    prize_amount DECIMAL(10,2) DEFAULT 0,
    started_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建奖池表
CREATE TABLE IF NOT EXISTS prize_pools (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    current_amount DECIMAL(10,2) DEFAULT 0,
    last_winner_id UUID,
    last_win_amount DECIMAL(10,2) DEFAULT 0,
    last_win_time TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 创建游戏日志表
CREATE TABLE IF NOT EXISTS game_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID REFERENCES players(id),
    action_type VARCHAR(50) NOT NULL,
    details JSONB,
    ip VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建支付表
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(id),
    payment_method VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) DEFAULT 'CNY',
    status VARCHAR(20) DEFAULT 'pending',
    transaction_id VARCHAR(50) UNIQUE,
    payment_details JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- ============================================
-- 2. 创建索引
-- ============================================
CREATE INDEX IF NOT EXISTS idx_players_deleted_at ON players(deleted_at);
CREATE INDEX IF NOT EXISTS idx_rooms_deleted_at ON rooms(deleted_at);
CREATE INDEX IF NOT EXISTS idx_reservations_deleted_at ON reservations(deleted_at);
CREATE INDEX IF NOT EXISTS idx_challenges_deleted_at ON challenges(deleted_at);
CREATE INDEX IF NOT EXISTS idx_game_logs_deleted_at ON game_logs(deleted_at);
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at);

CREATE INDEX IF NOT EXISTS idx_reservations_room_date ON reservations(room_id, reservation_date);
CREATE INDEX IF NOT EXISTS idx_game_logs_player_action ON game_logs(player_id, action_type);
CREATE INDEX IF NOT EXISTS idx_challenges_player_created ON challenges(player_id, created_at);

-- ============================================
-- 3. 插入默认数据
-- ============================================

-- 插入等级数据
INSERT INTO levels (id, name, created_at) VALUES 
    (gen_random_uuid(), '初级玩家', NOW()),
    (gen_random_uuid(), '中级玩家', NOW()),
    (gen_random_uuid(), '高级玩家', NOW()),
    (gen_random_uuid(), '大师级玩家', NOW()),
    (gen_random_uuid(), '传奇玩家', NOW())
ON CONFLICT (name) DO NOTHING;

-- 插入示例玩家（简化版本，去掉可能导致问题的复杂逻辑）
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
    
    -- 检查是否获取到了等级ID
    IF beginner_level_id IS NOT NULL AND intermediate_level_id IS NOT NULL AND advanced_level_id IS NOT NULL THEN
        -- 逐个插入玩家，避免批量插入的问题
        INSERT INTO players (name, level_id, balance) VALUES ('张三', beginner_level_id, 100.0) ON CONFLICT (name) DO NOTHING;
        INSERT INTO players (name, level_id, balance) VALUES ('李四', intermediate_level_id, 200.0) ON CONFLICT (name) DO NOTHING;
        INSERT INTO players (name, level_id, balance) VALUES ('王五', advanced_level_id, 300.0) ON CONFLICT (name) DO NOTHING;
        INSERT INTO players (name, level_id, balance) VALUES ('赵六', beginner_level_id, 150.0) ON CONFLICT (name) DO NOTHING;
        INSERT INTO players (name, level_id, balance) VALUES ('钱七', intermediate_level_id, 250.0) ON CONFLICT (name) DO NOTHING;
        
        RAISE NOTICE '成功插入玩家数据';
    ELSE
        RAISE NOTICE '未能获取等级ID，跳过玩家插入';
    END IF;
END $$;

-- 插入示例房间
INSERT INTO rooms (name, description, status, max_players) VALUES
    ('游戏室A', '适合初学者的房间', 'available', 4),
    ('游戏室B', '中级玩家专用房间', 'available', 6),
    ('游戏室C', '高级玩家竞技房间', 'available', 8),
    ('VIP包厢', '私人定制房间', 'available', 2),
    ('训练室', '新手练习专用', 'maintenance', 10)
ON CONFLICT (name) DO NOTHING;

-- 初始化奖池
INSERT INTO prize_pools (current_amount, last_win_amount) VALUES (0.0, 0.0);

-- 创建示例预约
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
    IF room_a_id IS NOT NULL AND room_b_id IS NOT NULL AND player1_id IS NOT NULL AND player2_id IS NOT NULL THEN
        INSERT INTO reservations (room_id, player_id, reservation_date, start_time, end_time, status) VALUES
            (room_a_id, player1_id, CURRENT_DATE + INTERVAL '1 day', '14:00', '16:00', 'active'),
            (room_b_id, player2_id, CURRENT_DATE + INTERVAL '2 days', '10:00', '12:00', 'active');
        
        RAISE NOTICE '成功创建示例预约';
    ELSE
        RAISE NOTICE '未能获取房间或玩家ID，跳过预约创建';
    END IF;
END $$;

-- 插入示例日志
DO $$
DECLARE
    player1_id UUID;
    player2_id UUID;
BEGIN
    SELECT id INTO player1_id FROM players WHERE name = '张三' LIMIT 1;
    SELECT id INTO player2_id FROM players WHERE name = '李四' LIMIT 1;
    
    IF player1_id IS NOT NULL THEN
        INSERT INTO game_logs (player_id, action_type, details, ip, user_agent) VALUES
            (player1_id, 'register', '{"registration_method": "username"}', '192.168.1.100', 'Mozilla/5.0'),
            (player1_id, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.100', 'Mozilla/5.0');
    END IF;
    
    IF player2_id IS NOT NULL THEN
        INSERT INTO game_logs (player_id, action_type, details, ip, user_agent) VALUES
            (player2_id, 'register', '{"registration_method": "email"}', '192.168.1.101', 'Chrome/91.0'),
            (player2_id, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.101', 'Chrome/91.0');
    END IF;
    
    IF player1_id IS NOT NULL OR player2_id IS NOT NULL THEN
        RAISE NOTICE '成功插入示例日志';
    END IF;
END $$;

-- ============================================
-- 4. 输出初始化完成信息
-- ============================================
DO $$
DECLARE
    level_count INTEGER;
    player_count INTEGER;
    room_count INTEGER;
    reservation_count INTEGER;
    log_count INTEGER;
    pool_amount DECIMAL(10,2);
BEGIN
    SELECT COUNT(*) INTO level_count FROM levels;
    SELECT COUNT(*) INTO player_count FROM players;
    SELECT COUNT(*) INTO room_count FROM rooms;
    SELECT COUNT(*) INTO reservation_count FROM reservations;
    SELECT COUNT(*) INTO log_count FROM game_logs;
    SELECT COALESCE(current_amount, 0) INTO pool_amount FROM prize_pools LIMIT 1;
    
    RAISE NOTICE '===========================================';
    RAISE NOTICE '数据库初始化完成！';
    RAISE NOTICE '已创建：';
    RAISE NOTICE '- % 个等级', level_count;
    RAISE NOTICE '- % 个玩家', player_count;
    RAISE NOTICE '- % 个房间', room_count;
    RAISE NOTICE '- % 个预约', reservation_count;
    RAISE NOTICE '- % 条日志', log_count;
    RAISE NOTICE '- 奖池已初始化，当前金额: %', pool_amount;
    RAISE NOTICE '===========================================';
END $$;