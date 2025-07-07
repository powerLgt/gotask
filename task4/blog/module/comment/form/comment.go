package form

type CreateCommentForm struct {
	Content string `form:"content" binding:"required,min=1,max=1000"`
	PostID  uint   `form:"postId" binding:"required"`
} 