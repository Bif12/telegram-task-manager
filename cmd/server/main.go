package main

import (
	"log"

	"github.com/bif12/telegram-task-manager/internal/config"
	"github.com/bif12/telegram-task-manager/internal/database"
	"github.com/bif12/telegram-task-manager/internal/handlers"
	"github.com/bif12/telegram-task-manager/internal/middleware"
	"github.com/bif12/telegram-task-manager/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Load()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Start Scheduler
	scheduler := services.NewScheduler()
	scheduler.Start()

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// GZIP compression
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Set Permissions-Policy header
	r.Use(func(c *gin.Context) {
		c.Header("Permissions-Policy", "interest-cohort=(), attribution-reporting=(), run-ad-auction=(), join-ad-interest-group=(), browsing-topics=()")
		c.Next()
	})

	// Health check endpoint
	r.GET("/health", handlers.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Public routes
		v1.GET("/health", handlers.HealthCheck)

		// Handlers
		authHandler := handlers.NewAuthHandler()
		telegramHandler := handlers.NewTelegramHandler()
		projectHandler := handlers.NewProjectHandler()
		taskHandler := handlers.NewTaskHandler()
		goalHandler := handlers.NewGoalHandler()
		habitHandler := handlers.NewHabitHandler()
		activityHandler := handlers.NewActivityHandler()

		// Start Telegram Polling in background
		go telegramHandler.StartPolling()

		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/telegram/verify", telegramHandler.VerifyTelegramWebApp)
			auth.POST("/telegram/widget", telegramHandler.VerifyTelegramWidget)

		}

		// Telegram bot webhook
		v1.POST("/telegram/webhook", telegramHandler.HandleWebhook)

		// Protected routes (require authentication)
		protected := v1.Group("")
		protected.Use(middleware.AuthRequired())
		{
			// User
			protected.GET("/user/me", authHandler.GetMe)

			// Projects
			protected.GET("/projects", projectHandler.List)
			protected.POST("/projects", projectHandler.Create)
			protected.GET("/projects/:id", projectHandler.Get)
			protected.PUT("/projects/:id", projectHandler.Update)
			protected.DELETE("/projects/:id", projectHandler.Delete)

			// Tasks
			protected.GET("/tasks", taskHandler.List)
			protected.POST("/tasks", taskHandler.Create)
			protected.GET("/tasks/:id", taskHandler.Get)
			protected.PUT("/tasks/:id", taskHandler.Update)
			protected.DELETE("/tasks/:id", taskHandler.Delete)

			// Goals
			protected.GET("/goals", goalHandler.List)
			protected.GET("/goals/:id", goalHandler.Get)
			protected.POST("/goals", goalHandler.Create)
			protected.PUT("/goals/:id", goalHandler.Update)
			protected.DELETE("/goals/:id", goalHandler.Delete)

			// Milestones
			milestoneHandler := handlers.NewMilestoneHandler()
			protected.POST("/milestones", milestoneHandler.Create)
			protected.PUT("/milestones/:id", milestoneHandler.Update)
			protected.DELETE("/milestones/:id", milestoneHandler.Delete)

			// Habits
			protected.GET("/habits", habitHandler.List)
			protected.POST("/habits", habitHandler.Create)
			protected.PUT("/habits/:id", habitHandler.Update)
			protected.DELETE("/habits/:id", habitHandler.Delete)
			protected.POST("/habits/:id/log", habitHandler.Log)

			// Activity
			protected.GET("/activities", activityHandler.List)
		}
	}

	// Start server
	port := ":" + config.AppConfig.Port
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
