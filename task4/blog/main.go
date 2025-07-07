package main

import (
	"blog/config"
	"blog/module/comment/controller"
	postController "blog/module/post/controller"
	userController "blog/module/user/controller"
	"blog/pkg/database"
	"blog/pkg/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatal("配置文件加载失败:", err)
	}

	// 初始化数据库
	if err := database.InitDatabase(); err != nil {
		log.Fatal("数据库初始化失败:", err)
	}

	// 创建Gin实例
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.ErrorHandlingMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// 创建控制器实例
	authCtrl := &userController.AuthController{}
	postCtrl := &postController.PostController{}
	commentCtrl := &controller.CommentController{}

	router.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s", firstName)
	})

	// API路由组
	api := router.Group("/api")
	{
		// 用户认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", authCtrl.Register)
			auth.POST("/login", authCtrl.Login)
		}

		// 文章相关路由
		posts := api.Group("/posts")
		{
			// 公开接口（无需认证）
			posts.GET("", postCtrl.GetPosts)    // 获取文章列表
			posts.GET("/:id", postCtrl.GetPost) // 获取单个文章

			// 需要认证的接口
			postsAuth := posts.Group("")
			postsAuth.Use(middleware.AuthMiddleware())
			postsAuth.POST("", postCtrl.CreatePost)       // 创建文章
			postsAuth.PUT("/:id", postCtrl.UpdatePost)    // 更新文章
			postsAuth.DELETE("/:id", postCtrl.DeletePost) // 删除文章
		}

		// 评论相关路由
		comments := api.Group("/comments")
		{
			// 公开接口（无需认证）
			comments.GET("/post/:postId", commentCtrl.GetComments) // 获取文章评论

			// 需要认证的接口
			commentsAuth := comments.Group("")
			commentsAuth.Use(middleware.AuthMiddleware())
			commentsAuth.POST("", commentCtrl.CreateComment) // 创建评论
		}
	}

	// 启动服务器
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("服务器启动在端口 %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
