package handler

import (
	"net/http"

	"github.com/boyosoft/gangguanbao/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

// StatsResponse 统计分析响应
type StatsResponse struct {
	Counts           repository.StatsCounts   `json:"counts"`
	OrdersByYear     []repository.YearCount   `json:"ordersByYear"`
	OrdersByMonthCur []repository.MonthCount  `json:"ordersByMonthCur"`
	OrdersByMonthLast []repository.MonthCount `json:"ordersByMonthLast"`
	OrdersByMaterial []repository.NameCount  `json:"ordersByMaterial"`
	OrdersBySpec     []repository.NameCount  `json:"ordersBySpec"`
}

// GetStats 获取统计分析数据
func (h *Handler) GetStats(c *gin.Context) {
	curYear := repository.GetCurrentYear()
	lastYear := curYear - 1

	counts, err := h.Repo.GetStatsCounts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byYear, err := h.Repo.GetOrdersByYear()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byMonthCur, err := h.Repo.GetOrdersByMonth(curYear)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byMonthLast, err := h.Repo.GetOrdersByMonth(lastYear)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byMaterial, err := h.Repo.GetOrdersByMaterial(curYear)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bySpec, err := h.Repo.GetOrdersBySpec(curYear)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, StatsResponse{
		Counts:            counts,
		OrdersByYear:      byYear,
		OrdersByMonthCur:  byMonthCur,
		OrdersByMonthLast: byMonthLast,
		OrdersByMaterial:  byMaterial,
		OrdersBySpec:      bySpec,
	})
}
