-- 规格及厂家查询数据库 schema
-- 用于统计本公司往年订单数据
-- 兼容 MySQL 5.0

CREATE DATABASE IF NOT EXISTS gangguanbao_production DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE gangguanbao_production;

-- 材质表
CREATE TABLE IF NOT EXISTS biz_materials (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '材质名称',
    code VARCHAR(50) COMMENT '材质编码',
    created_at DATETIME NULL COMMENT '创建时间',
    updated_at DATETIME NULL COMMENT '更新时间',
    UNIQUE KEY uk_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='材质';

-- 规格表
CREATE TABLE IF NOT EXISTS biz_specs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '规格名称',
    material_id INT COMMENT '关联材质',
    unit VARCHAR(20) DEFAULT '件' COMMENT '单位',
    created_at DATETIME NULL COMMENT '创建时间',
    updated_at DATETIME NULL COMMENT '更新时间',
    KEY idx_material (material_id),
    CONSTRAINT fk_spec_material FOREIGN KEY (material_id) REFERENCES biz_materials(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='规格';

-- 厂家/供应商表
CREATE TABLE IF NOT EXISTS biz_vendors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL COMMENT '厂家名称',
    code VARCHAR(50) COMMENT '厂家编码',
    contact VARCHAR(100) COMMENT '联系人',
    phone VARCHAR(50) COMMENT '电话',
    created_at DATETIME NULL COMMENT '创建时间',
    updated_at DATETIME NULL COMMENT '更新时间',
    KEY idx_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='厂家';

-- 订单表（往年数据）
CREATE TABLE IF NOT EXISTS biz_orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_no VARCHAR(50) COMMENT '订单号',
    material_id INT COMMENT '材质ID',
    spec_id INT COMMENT '规格ID',
    vendor_id INT COMMENT '厂家ID',
    order_date DATE COMMENT '订单日期',
    quantity DECIMAL(18,4) DEFAULT 0 COMMENT '数量',
    unit_price DECIMAL(18,4) COMMENT '单价',
    amount DECIMAL(18,2) COMMENT '金额',
    remark VARCHAR(500) COMMENT '备注',
    created_at DATETIME NULL COMMENT '创建时间',
    updated_at DATETIME NULL COMMENT '更新时间',
    KEY idx_order_date (order_date),
    KEY idx_material (material_id),
    KEY idx_spec (spec_id),
    KEY idx_vendor (vendor_id),
    KEY idx_order_no (order_no),
    CONSTRAINT fk_order_material FOREIGN KEY (material_id) REFERENCES biz_materials(id) ON DELETE SET NULL,
    CONSTRAINT fk_order_spec FOREIGN KEY (spec_id) REFERENCES biz_specs(id) ON DELETE SET NULL,
    CONSTRAINT fk_order_vendor FOREIGN KEY (vendor_id) REFERENCES biz_vendors(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单';

-- 示例数据（可选）
INSERT INTO biz_materials (name, code) VALUES ('不锈钢', 'SS'), ('碳钢', 'CS'), ('合金钢', 'AS') ON DUPLICATE KEY UPDATE name=name;
INSERT INTO biz_vendors (name, code) VALUES ('某某钢铁厂', 'V001'), ('某某管业', 'V002') ON DUPLICATE KEY UPDATE name=name;

-- 规格 10 条
INSERT INTO biz_specs (name, material_id, unit) VALUES
('DN50', 1, '米'),
('DN80', 1, '米'),
('DN100', 1, '米'),
('Φ32*3', 2, '吨'),
('Φ57*3.5', 2, '吨'),
('Φ89*4', 2, '吨'),
('316L-20*2', 3, '公斤'),
('304-25*1.5', 3, '公斤'),
('219*6', 2, '米'),
('273*8', 1, '米');

-- 订单 30 条
INSERT INTO biz_orders (order_no, material_id, spec_id, vendor_id, order_date, quantity, unit_price, amount, remark) VALUES
('ORD202301001', 1, 1, 1, '2023-01-15', 100, 25.50, 2550.00, ''),
('ORD202301002', 1, 2, 1, '2023-01-20', 80, 32.00, 2560.00, ''),
('ORD202301003', 2, 4, 2, '2023-02-10', 5.5, 5800.00, 31900.00, ''),
('ORD202302001', 1, 3, 1, '2023-02-25', 120, 38.00, 4560.00, ''),
('ORD202302002', 2, 5, 2, '2023-03-05', 3.2, 6200.00, 19840.00, ''),
('ORD202303001', 3, 7, 1, '2023-03-18', 500, 45.00, 22500.00, ''),
('ORD202303002', 3, 8, 2, '2023-03-22', 350, 42.00, 14700.00, ''),
('ORD202304001', 1, 1, 1, '2023-04-08', 90, 26.00, 2340.00, ''),
('ORD202304002', 2, 9, 2, '2023-04-15', 200, 28.50, 5700.00, ''),
('ORD202304003', 1, 10, 1, '2023-04-20', 60, 55.00, 3300.00, ''),
('ORD202305001', 2, 4, 2, '2023-05-12', 4.8, 5900.00, 28320.00, ''),
('ORD202305002', 3, 7, 1, '2023-05-25', 420, 44.00, 18480.00, ''),
('ORD202306001', 1, 2, 1, '2023-06-08', 75, 31.50, 2362.50, ''),
('ORD202306002', 1, 3, 2, '2023-06-18', 110, 37.00, 4070.00, ''),
('ORD202307001', 2, 5, 2, '2023-07-05', 6.0, 6000.00, 36000.00, ''),
('ORD202307002', 3, 8, 1, '2023-07-20', 280, 43.00, 12040.00, ''),
('ORD202308001', 1, 1, 1, '2023-08-10', 95, 25.00, 2375.00, ''),
('ORD202308002', 2, 9, 2, '2023-08-22', 180, 29.00, 5220.00, ''),
('ORD202309001', 1, 10, 1, '2023-09-05', 70, 54.00, 3780.00, ''),
('ORD202309002', 2, 4, 2, '2023-09-18', 5.2, 5850.00, 30420.00, ''),
('ORD202310001', 3, 7, 1, '2023-10-08', 450, 46.00, 20700.00, ''),
('ORD202310002', 1, 2, 2, '2023-10-20', 85, 32.50, 2762.50, ''),
('ORD202311001', 1, 3, 1, '2023-11-12', 130, 36.00, 4680.00, ''),
('ORD202311002', 2, 5, 2, '2023-11-25', 4.0, 6100.00, 24400.00, ''),
('ORD202312001', 3, 8, 1, '2023-12-05', 320, 41.00, 13120.00, ''),
('ORD202312002', 1, 1, 2, '2023-12-18', 105, 24.50, 2572.50, ''),
('ORD202401001', 2, 9, 2, '2024-01-10', 220, 27.00, 5940.00, ''),
('ORD202401002', 1, 10, 1, '2024-01-22', 65, 56.00, 3640.00, ''),
('ORD202402001', 3, 7, 1, '2024-02-08', 380, 45.50, 17290.00, ''),
('ORD202402002', 2, 4, 2, '2024-02-20', 5.8, 5750.00, 33350.00, '');
