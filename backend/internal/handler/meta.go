package handler

import (
	"net/http"

	"github.com/boyosoft/gangguanbao/backend/internal/model"
	// "github.com/boyosoft/gangguanbao/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// ListMaterials 材质列表
func (h *Handler) ListMaterials(c *gin.Context) {
	list, err := h.Repo.ListMaterials()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"list": list})
}

// GetMaterial 获取单条材质
func (h *Handler) GetMaterial(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	m, err := h.Repo.GetMaterial(req.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}

// CreateMaterial 新增材质
func (h *Handler) CreateMaterial(c *gin.Context) {
	var m model.Material
	if err := c.ShouldBindJSON(&m); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if m.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "名称为必填"})
		return
	}
	if err := h.Repo.CreateMaterial(&m); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": m.ID})
}

// UpdateMaterial 更新材质
func (h *Handler) UpdateMaterial(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	var m model.Material
	if err := c.ShouldBindJSON(&m); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m.ID = req.ID
	if err := h.Repo.UpdateMaterial(&m); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// DeleteMaterial 删除材质
func (h *Handler) DeleteMaterial(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	if err := h.Repo.DeleteMaterial(req.ID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// ListSpecs 规格列表
func (h *Handler) ListSpecs(c *gin.Context) {
	list, err := h.Repo.ListSpecs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"list": list})
}

// GetSpec 获取单条规格
func (h *Handler) GetSpec(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	s, err := h.Repo.GetSpec(req.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}
	c.IndentedJSON(http.StatusOK, s)
}

// CreateSpec 新增规格
func (h *Handler) CreateSpec(c *gin.Context) {
	var s model.Spec
	if err := c.ShouldBindJSON(&s); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if s.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "规格名称为必填"})
		return
	}
	if s.Unit == "" {
		s.Unit = "件"
	}
	if err := h.Repo.CreateSpec(&s); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": s.ID})
}

// UpdateSpec 更新规格
func (h *Handler) UpdateSpec(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	var s model.Spec
	if err := c.ShouldBindJSON(&s); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.ID = req.ID
	if err := h.Repo.UpdateSpec(&s); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// DeleteSpec 删除规格
func (h *Handler) DeleteSpec(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	if err := h.Repo.DeleteSpec(req.ID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// ListVendors 厂家列表
func (h *Handler) ListVendors(c *gin.Context) {
	list, err := h.Repo.ListVendors()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"list": list})
}

// GetVendor 获取单条厂家
func (h *Handler) GetVendor(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	v, err := h.Repo.GetVendor(req.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}
	c.IndentedJSON(http.StatusOK, v)
}

// CreateVendor 新增厂家
func (h *Handler) CreateVendor(c *gin.Context) {
	var v model.Vendor
	if err := c.ShouldBindJSON(&v); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if v.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "厂家名称为必填"})
		return
	}
	if err := h.Repo.CreateVendor(&v); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": v.ID})
}

// UpdateVendor 更新厂家
func (h *Handler) UpdateVendor(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	var v model.Vendor
	if err := c.ShouldBindJSON(&v); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v.ID = req.ID
	if err := h.Repo.UpdateVendor(&v); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// DeleteVendor 删除厂家
func (h *Handler) DeleteVendor(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	if err := h.Repo.DeleteVendor(req.ID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}
