package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/boyosoft/gangguanbao/backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository 数据访问层
type Repository struct {
	DB *gorm.DB
}

// New 创建 Repository
func New(cfg *config.DatabaseConfig) (*Repository, error) {
	port := cfg.Port
	if port <= 0 {
		port = 3306
	}
	charset := cfg.Charset
	if charset == "" {
		charset = "utf8"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		cfg.User, cfg.Password, cfg.Host, port, cfg.DBName, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}
