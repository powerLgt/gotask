package controller

import (
	"blog/module/comment/form"
	"blog/module/comment/service"
	"blog/module/common/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

var commentService = &service.CommentService{}

// 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
	var commentForm form.CreateCommentForm

	// 绑定JSON数据
	if err := c.ShouldBindJSON(&commentForm); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误: "+err.Error()))
		return
	}

	// 从上下文中获取用户ID
	userID, _ := c.Get("user_id")

	// 创建评论
	comment, err := commentService.CreateComment(userID.(uint), commentForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, vo.Success("评论创建成功", comment))
}

// 获取文章的评论列表
func (cc *CommentController) GetComments(c *gin.Context) {
	// 获取文章ID
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("无效的文章ID"))
		return
	}

	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	comments, total, err := commentService.GetCommentsByPostID(uint(postID), page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	data := map[string]interface{}{
		"comments": comments,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	c.JSON(http.StatusOK, vo.Success("获取评论列表成功", data))
}
