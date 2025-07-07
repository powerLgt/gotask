package middleware

import (
	"blog/module/common/vo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误处理中间件
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// 记录错误日志
		log.Printf("程序发生恐慌: %v", recovered)
		
		// 返回统一的错误响应
		c.JSON(http.StatusInternalServerError, vo.Fail("服务器内部错误"))
		c.Abort()
	})
} 