package handler

import (
	"github.com/boyosoft/gangguanbao/backend/internal/repository"
)

// Handler HTTP 处理器
type Handler struct {
	Repo *repository.Repository
}

// New 创建 Handler
func New(repo *repository.Repository) *Handler {
	return &Handler{Repo: repo}
}
