package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liyq96/codereview-ai/internal/config"
	"github.com/liyq96/codereview-ai/internal/handler"
	"github.com/liyq96/codereview-ai/internal/route"
	"github.com/liyq96/codereview-ai/internal/service"
	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*config.SystemConfig, error) {
	cfg := &config.SystemConfig{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 错误处理中间件
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "服务器内部错误",
				})
			}
		}()
		c.Next()
	})

	// 创建API分组
	apiV1 := r.Group("/v1")
	{
		// 初始化服务和handler
		services := service.InitServices()
		handlers := handler.InitHandlers(services)

		// 注册路由
		route.RegisterCodeReviewRoutes(apiV1, handlers.CodeReviewHandler)
	}

	return r
}

func startServer(cfg *config.SystemConfig, router *gin.Engine) error {
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("服务器启动在 %s", s.Addr)
	return s.ListenAndServe()
}

func main() {
	// 加载配置
	cfg, err := LoadConfig("config/system_config.yml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 设置gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化路由
	router := setupRouter()

	// 启动服务器
	if err := startServer(cfg, router); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
