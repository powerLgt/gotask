package form

type CreatePostForm struct {
	Title   string `form:"title" binding:"required,min=1,max=200"`
	Content string `form:"content" binding:"required,min=1"`
}

type UpdatePostForm struct {
	Title   string `form:"title" binding:"omitempty,min=1,max=200"`
	Content string `form:"content" binding:"omitempty,min=1"`
}
