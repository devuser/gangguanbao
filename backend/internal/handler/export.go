package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/boyosoft/gangguanbao/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// ExportOrders 导出订单为 Excel
func (h *Handler) ExportOrders(c *gin.Context) {
	var q repository.OrderQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	q.SortBy = "orderDate"
	q.SortOrder = "desc"

	list, err := h.Repo.ListOrdersForExport(q)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"订单号", "订单日期", "材质", "规格", "厂家", "数量", "单价", "金额", "备注"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, o := range list {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), o.OrderNo)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), strPtr(o.OrderDate))
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), o.MaterialName)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), o.SpecName)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), o.VendorName)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), o.Quantity)
		f.SetCellValue(sheet, "G"+strconv.Itoa(row), floatPtr(o.UnitPrice))
		f.SetCellValue(sheet, "H"+strconv.Itoa(row), floatPtr(o.Amount))
		f.SetCellValue(sheet, "I"+strconv.Itoa(row), o.Remark)
	}

	filename := fmt.Sprintf("订单导出_%s.xlsx", time.Now().Format("20060102_150405"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", mustMarshal(f))
}

// ExportMaterials 导出材质为 Excel
func (h *Handler) ExportMaterials(c *gin.Context) {
	list, err := h.Repo.ListMaterials()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	f := excelize.NewFile()
	sheet := "Sheet1"
	for i, h := range []string{"ID", "名称", "编码", "创建时间"} {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, o := range list {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), o.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), o.Name)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), o.Code)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), timePtr(o.CreatedAt))
	}
	writeExcelResponse(c, "材质", f)
}

// ExportSpecs 导出规格为 Excel
func (h *Handler) ExportSpecs(c *gin.Context) {
	list, err := h.Repo.ListSpecs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	f := excelize.NewFile()
	sheet := "Sheet1"
	for i, h := range []string{"ID", "规格名称", "材质", "单位", "创建时间"} {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, o := range list {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), o.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), o.Name)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), o.MaterialName)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), o.Unit)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), timePtr(o.CreatedAt))
	}
	writeExcelResponse(c, "规格", f)
}

// ExportVendors 导出厂家为 Excel
func (h *Handler) ExportVendors(c *gin.Context) {
	list, err := h.Repo.ListVendors()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	f := excelize.NewFile()
	sheet := "Sheet1"
	for i, h := range []string{"ID", "厂家名称", "编码", "联系人", "电话", "创建时间"} {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, o := range list {
		row := i + 2
		f.SetCellValue(sheet, "A"+strconv.Itoa(row), o.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(row), o.Name)
		f.SetCellValue(sheet, "C"+strconv.Itoa(row), o.Code)
		f.SetCellValue(sheet, "D"+strconv.Itoa(row), o.Contact)
		f.SetCellValue(sheet, "E"+strconv.Itoa(row), o.Phone)
		f.SetCellValue(sheet, "F"+strconv.Itoa(row), timePtr(o.CreatedAt))
	}
	writeExcelResponse(c, "厂家", f)
}

func writeExcelResponse(c *gin.Context, name string, f *excelize.File) {
	filename := fmt.Sprintf("%s导出_%s.xlsx", name, time.Now().Format("20060102_150405"))
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", mustMarshal(f))
}

func mustMarshal(f *excelize.File) []byte {
	buf, err := f.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func strPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func floatPtr(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func timePtr(t *time.Time) interface{} {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
