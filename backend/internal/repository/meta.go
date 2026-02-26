package repository

import (
	"github.com/boyosoft/gangguanbao/backend/internal/model"
)

// ListMaterials 查询所有材质
func (r *Repository) ListMaterials() ([]model.Material, error) {
	var list []model.Material
	err := r.DB.Order("name").Find(&list).Error
	return list, err
}

// GetMaterial 根据ID获取材质
func (r *Repository) GetMaterial(id int) (*model.Material, error) {
	var m model.Material
	err := r.DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// CreateMaterial 新增材质
func (r *Repository) CreateMaterial(m *model.Material) error {
	return r.DB.Create(m).Error
}

// UpdateMaterial 更新材质
func (r *Repository) UpdateMaterial(m *model.Material) error {
	return r.DB.Save(m).Error
}

// DeleteMaterial 删除材质
func (r *Repository) DeleteMaterial(id int) error {
	return r.DB.Delete(&model.Material{}, id).Error
}

// ListSpecs 查询所有规格（含材质名称）
func (r *Repository) ListSpecs() ([]model.Spec, error) {
	var list []model.Spec
	err := r.DB.Table("biz_specs s").
		Select("s.id, s.name, s.material_id, COALESCE(m.name, '') AS material_name, COALESCE(s.unit, '件') AS unit, s.created_at, s.updated_at").
		Joins("LEFT JOIN biz_materials m ON s.material_id = m.id").
		Order("s.name").
		Find(&list).Error
	return list, err
}

// GetSpec 根据ID获取规格（含材质名称）
func (r *Repository) GetSpec(id int) (*model.Spec, error) {
	var s model.Spec
	err := r.DB.Table("biz_specs s").
		Select("s.id, s.name, s.material_id, COALESCE(m.name, '') AS material_name, COALESCE(s.unit, '件') AS unit, s.created_at, s.updated_at").
		Joins("LEFT JOIN biz_materials m ON s.material_id = m.id").
		Where("s.id = ?", id).
		First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// CreateSpec 新增规格
func (r *Repository) CreateSpec(s *model.Spec) error {
	return r.DB.Create(s).Error
}

// UpdateSpec 更新规格
func (r *Repository) UpdateSpec(s *model.Spec) error {
	return r.DB.Save(s).Error
}

// DeleteSpec 删除规格
func (r *Repository) DeleteSpec(id int) error {
	return r.DB.Delete(&model.Spec{}, id).Error
}

// ListVendors 查询所有厂家
func (r *Repository) ListVendors() ([]model.Vendor, error) {
	var list []model.Vendor
	err := r.DB.Order("name").Find(&list).Error
	return list, err
}

// GetVendor 根据ID获取厂家
func (r *Repository) GetVendor(id int) (*model.Vendor, error) {
	var v model.Vendor
	err := r.DB.First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// CreateVendor 新增厂家
func (r *Repository) CreateVendor(v *model.Vendor) error {
	return r.DB.Create(v).Error
}

// UpdateVendor 更新厂家
func (r *Repository) UpdateVendor(v *model.Vendor) error {
	return r.DB.Save(v).Error
}

// DeleteVendor 删除厂家
func (r *Repository) DeleteVendor(id int) error {
	return r.DB.Delete(&model.Vendor{}, id).Error
}
