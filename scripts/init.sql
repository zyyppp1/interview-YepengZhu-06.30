-- ============================================
-- 完整的数据库初始化脚本（递增ID版本）
-- ============================================

-- 创建必要的扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- 1. 创建所有表结构
-- ============================================

-- 创建等级表（使用递增ID）
CREATE TABLE IF NOT EXISTS levels (
    id SERIAL PRIMARY KEY,  -- 从1开始自动递增
    name VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 创建玩家表（使用递增ID）
CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,  -- 从1开始自动递增
    name VARCHAR(50) UNIQUE NOT NULL,
    level_id INTEGER NOT NULL REFERENCES levels(id),
    balance DECIMAL(10,2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建房间表（使用递增ID）
CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,  -- 从1开始自动递增
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    status VARCHAR(20) DEFAULT 'available',
    max_players INTEGER DEFAULT 4,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建预约表（主键使用递增ID，外键使用INTEGER）
CREATE TABLE IF NOT EXISTS reservations (
    id SERIAL PRIMARY KEY,
    room_id INTEGER NOT NULL REFERENCES rooms(id),
    player_id INTEGER NOT NULL REFERENCES players(id),
    reservation_date DATE NOT NULL,
    start_time VARCHAR(5) NOT NULL,
    end_time VARCHAR(5) NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建挑战表（主键使用递增ID，外键使用INTEGER）
CREATE TABLE IF NOT EXISTS challenges (
    id SERIAL PRIMARY KEY,
    player_id INTEGER NOT NULL REFERENCES players(id),
    amount DECIMAL(10,2) DEFAULT 20.01,
    is_winner BOOLEAN DEFAULT FALSE,
    prize_amount DECIMAL(10,2) DEFAULT 0,
    started_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建奖池表（使用递增ID）
CREATE TABLE IF NOT EXISTS prize_pools (
    id SERIAL PRIMARY KEY,
    current_amount DECIMAL(10,2) DEFAULT 0,
    last_winner_id INTEGER,  -- 引用players.id
    last_win_amount DECIMAL(10,2) DEFAULT 0,
    last_win_time TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 创建游戏日志表（主键使用递增ID，外键使用INTEGER）
CREATE TABLE IF NOT EXISTS game_logs (
    id SERIAL PRIMARY KEY,
    player_id INTEGER REFERENCES players(id),
    action_type VARCHAR(50) NOT NULL,
    details JSONB,
    ip VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 创建支付表（主键使用递增ID，外键使用INTEGER）
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    player_id INTEGER NOT NULL REFERENCES players(id),
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

-- 插入等级数据（ID将自动从1开始递增）
INSERT INTO levels (name, created_at) VALUES 
    ('初级玩家', NOW()),        -- ID: 1
    ('中级玩家', NOW()),        -- ID: 2
    ('高级玩家', NOW()),        -- ID: 3
    ('大师级玩家', NOW()),      -- ID: 4
    ('传奇玩家', NOW())         -- ID: 5
ON CONFLICT (name) DO NOTHING;

-- 插入示例玩家（ID将自动从1开始递增）
INSERT INTO players (name, level_id, balance) VALUES 
    ('张三', 1, 100.0),        -- 玩家ID: 1, 等级ID: 1(初级玩家)
    ('李四', 2, 200.0),        -- 玩家ID: 2, 等级ID: 2(中级玩家)
    ('王五', 3, 300.0),        -- 玩家ID: 3, 等级ID: 3(高级玩家)
    ('赵六', 1, 150.0),        -- 玩家ID: 4, 等级ID: 1(初级玩家)
    ('钱七', 2, 250.0)         -- 玩家ID: 5, 等级ID: 2(中级玩家)
ON CONFLICT (name) DO NOTHING;

-- 插入示例房间（ID将自动从1开始递增）
INSERT INTO rooms (name, description, status, max_players) VALUES
    ('游戏室A', '适合初学者的房间', 'available', 4),     -- 房间ID: 1
    ('游戏室B', '中级玩家专用房间', 'available', 6),     -- 房间ID: 2
    ('游戏室C', '高级玩家竞技房间', 'available', 8),     -- 房间ID: 3
    ('VIP包厢', '私人定制房间', 'available', 2),        -- 房间ID: 4
    ('训练室', '新手练习专用', 'maintenance', 10)        -- 房间ID: 5
ON CONFLICT (name) DO NOTHING;

-- 初始化奖池
INSERT INTO prize_pools (current_amount, last_win_amount) VALUES (0.0, 0.0);

-- 创建示例预约（使用具体的数字ID）
INSERT INTO reservations (room_id, player_id, reservation_date, start_time, end_time, status) VALUES
    (1, 1, CURRENT_DATE + INTERVAL '1 day', '14:00', '16:00', 'active'),    -- 张三预约游戏室A
    (2, 2, CURRENT_DATE + INTERVAL '2 days', '10:00', '12:00', 'active'),   -- 李四预约游戏室B
    (3, 3, CURRENT_DATE + INTERVAL '1 day', '16:00', '18:00', 'active'),    -- 王五预约游戏室C
    (4, 4, CURRENT_DATE + INTERVAL '3 days', '19:00', '21:00', 'active')    -- 赵六预约VIP包厢
ON CONFLICT DO NOTHING;

-- 插入示例日志（使用具体的数字ID）
INSERT INTO game_logs (player_id, action_type, details, ip, user_agent) VALUES
    (1, 'register', '{"registration_method": "username"}', '192.168.1.100', 'Mozilla/5.0'),
    (1, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.100', 'Mozilla/5.0'),
    (2, 'register', '{"registration_method": "email"}', '192.168.1.101', 'Chrome/91.0'),
    (2, 'login', '{"login_time": "' || NOW() || '"}', '192.168.1.101', 'Chrome/91.0'),
    (3, 'register', '{"registration_method": "phone"}', '192.168.1.102', 'Safari/14.0'),
    (3, 'enter_room', '{"room_id": 3, "room_name": "游戏室C"}', '192.168.1.102', 'Safari/14.0');

-- ============================================
-- 4. 验证并输出初始化完成信息
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
    RAISE NOTICE '- % 个等级 (ID: 1-%)', level_count, level_count;
    RAISE NOTICE '- % 个玩家 (ID: 1-%)', player_count, player_count;
    RAISE NOTICE '- % 个房间 (ID: 1-%)', room_count, room_count;
    RAISE NOTICE '- % 个预约', reservation_count;
    RAISE NOTICE '- % 条日志', log_count;
    RAISE NOTICE '- 奖池已初始化，当前金额: %', pool_amount;
    RAISE NOTICE '===========================================';
    
    -- 显示具体的ID映射
    RAISE NOTICE '等级ID映射：';
    RAISE NOTICE '1: 初级玩家, 2: 中级玩家, 3: 高级玩家, 4: 大师级玩家, 5: 传奇玩家';
    RAISE NOTICE '玩家ID映射：';
    RAISE NOTICE '1: 张三(初级), 2: 李四(中级), 3: 王五(高级), 4: 赵六(初级), 5: 钱七(中级)';
    RAISE NOTICE '房间ID映射：';
    RAISE NOTICE '1: 游戏室A, 2: 游戏室B, 3: 游戏室C, 4: VIP包厢, 5: 训练室';
    RAISE NOTICE '===========================================';
END $$;