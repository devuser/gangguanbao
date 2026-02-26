package handler

import (
	"net/http"

	"github.com/boyosoft/gangguanbao/backend/internal/model"
	"github.com/boyosoft/gangguanbao/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// ListOrders 订单列表（分页、搜索、排序）
func (h *Handler) ListOrders(c *gin.Context) {
	var q repository.OrderQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, total, err := h.Repo.ListOrders(q)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"list":  list,
		"total": total,
	})
}

// GetOrder 获取单条订单
func (h *Handler) GetOrder(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	o, err := h.Repo.GetOrder(req.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}
	c.IndentedJSON(http.StatusOK, o)
}

// CreateOrder 新增订单
func (h *Handler) CreateOrder(c *gin.Context) {
	var o model.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.CreateOrder(&o); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": o.ID})
}

// UpdateOrder 更新订单
func (h *Handler) UpdateOrder(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	var o model.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	o.ID = req.ID
	if err := h.Repo.UpdateOrder(&o); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}

// DeleteOrder 删除订单
func (h *Handler) DeleteOrder(c *gin.Context) {
	var req struct {
		ID int `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	if err := h.Repo.DeleteOrder(req.ID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"ok": true})
}
