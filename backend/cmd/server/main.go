package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/boyosoft/gangguanbao/backend/internal/config"
	"github.com/boyosoft/gangguanbao/backend/internal/handler"
	"github.com/boyosoft/gangguanbao/backend/internal/repository"
)

func main() {
	cfgPath := "config.yaml"
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	}
	// 若为相对路径，尝试在可执行文件同目录查找
	if !filepath.IsAbs(cfgPath) {
		if exe, err := os.Executable(); err == nil {
			dir := filepath.Dir(exe)
			if _, err := os.Stat(filepath.Join(dir, cfgPath)); err == nil {
				cfgPath = filepath.Join(dir, cfgPath)
			}
		}
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	repo, err := repository.New(&cfg.Database)
	if err != nil {
		log.Fatalf("init db: %v", err)
	}
	if sqlDB, err := repo.DB.DB(); err == nil {
		defer sqlDB.Close()
	}

	h := handler.New(repo)

	gin.SetMode(cfg.Server.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	if cfg.Server.Mode != "release" {
		r.Use(gin.Logger())
	}
	r.SetTrustedProxies(nil)

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/materials", h.ListMaterials)
		api.POST("/materials", h.CreateMaterial)
		api.GET("/materials/export", h.ExportMaterials)
		api.GET("/materials/:id", h.GetMaterial)
		api.PUT("/materials/:id", h.UpdateMaterial)
		api.DELETE("/materials/:id", h.DeleteMaterial)

		api.GET("/specs", h.ListSpecs)
		api.POST("/specs", h.CreateSpec)
		api.GET("/specs/export", h.ExportSpecs)
		api.GET("/specs/:id", h.GetSpec)
		api.PUT("/specs/:id", h.UpdateSpec)
		api.DELETE("/specs/:id", h.DeleteSpec)

		api.GET("/vendors", h.ListVendors)
		api.POST("/vendors", h.CreateVendor)
		api.GET("/vendors/export", h.ExportVendors)
		api.GET("/vendors/:id", h.GetVendor)
		api.PUT("/vendors/:id", h.UpdateVendor)
		api.DELETE("/vendors/:id", h.DeleteVendor)

		api.GET("/orders", h.ListOrders)
		api.POST("/orders", h.CreateOrder)
		api.GET("/orders/export", h.ExportOrders)
		api.GET("/orders/:id", h.GetOrder)
		api.PUT("/orders/:id", h.UpdateOrder)
		api.DELETE("/orders/:id", h.DeleteOrder)

		api.GET("/stats", h.GetStats)
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("server listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("run: %v", err)
	}
}
