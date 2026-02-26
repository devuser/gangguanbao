-- 将业务表统一添加 biz_ 前缀
-- 执行前请备份数据库
-- 兼容 MySQL 5.0+

USE gangguanbao_production;

-- 1. 删除外键约束（按依赖顺序）
ALTER TABLE orders DROP FOREIGN KEY fk_order_material;
ALTER TABLE orders DROP FOREIGN KEY fk_order_spec;
ALTER TABLE orders DROP FOREIGN KEY fk_order_vendor;
ALTER TABLE specs DROP FOREIGN KEY fk_spec_material;

-- 2. 重命名表（无外键依赖的表先重命名）
RENAME TABLE materials TO biz_materials;
RENAME TABLE vendors TO biz_vendors;
RENAME TABLE specs TO biz_specs;
RENAME TABLE orders TO biz_orders;

-- 3. 重新添加外键约束
ALTER TABLE biz_specs
    ADD CONSTRAINT fk_spec_material FOREIGN KEY (material_id) REFERENCES biz_materials(id) ON DELETE SET NULL;

ALTER TABLE biz_orders
    ADD CONSTRAINT fk_order_material FOREIGN KEY (material_id) REFERENCES biz_materials(id) ON DELETE SET NULL,
    ADD CONSTRAINT fk_order_spec FOREIGN KEY (spec_id) REFERENCES biz_specs(id) ON DELETE SET NULL,
    ADD CONSTRAINT fk_order_vendor FOREIGN KEY (vendor_id) REFERENCES biz_vendors(id) ON DELETE SET NULL;
