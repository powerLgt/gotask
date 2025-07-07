package controller

import (
	"blog/module/common/vo"
	"blog/module/post/form"
	"blog/module/post/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

var postService = &service.PostService{}

// 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
	var createForm form.CreatePostForm

	// 绑定JSON数据
	if err := c.ShouldBindJSON(&createForm); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误: "+err.Error()))
		return
	}

	// 从上下文中获取用户ID
	userID, _ := c.Get("user_id")

	// 创建文章
	post, err := postService.CreatePost(userID.(uint), createForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, vo.Success("文章创建成功", post))
}

// 获取文章列表
func (pc *PostController) GetPosts(c *gin.Context) {
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

	posts, total, err := postService.GetAllPosts(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.Fail("获取文章列表失败"))
		return
	}

	data := map[string]interface{}{
		"posts":    posts,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	c.JSON(http.StatusOK, vo.Success("获取文章列表成功", data))
}

// 获取单个文章
func (pc *PostController) GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("无效的文章ID"))
		return
	}

	post, err := postService.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("获取文章成功", post))
}

// 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("无效的文章ID"))
		return
	}

	var updateForm form.UpdatePostForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误: "+err.Error()))
		return
	}

	// 从上下文中获取用户ID
	userID, _ := c.Get("user_id")

	post, err := postService.UpdatePost(uint(id), userID.(uint), updateForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("文章更新成功", post))
}

// 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail("无效的文章ID"))
		return
	}

	// 从上下文中获取用户ID
	userID, _ := c.Get("user_id")

	err = postService.DeletePost(uint(id), userID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, vo.Success("文章删除成功", nil))
} 