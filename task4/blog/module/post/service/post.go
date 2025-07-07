package service

import (
	"blog/module/common/model"
	"blog/module/post/form"
	"blog/pkg/database"
	"errors"
)

type PostService struct{}

// 创建文章
func (s *PostService) CreatePost(userID uint, form form.CreatePostForm) (post model.Post, err error) {
	post = model.Post{
		Title:   form.Title,
		Content: form.Content,
		UserID:  userID,
	}

	err = database.DB.Create(&post).Error
	if err != nil {
		err = errors.New("文章创建失败")
		return
	}

	// 预加载用户信息
	database.DB.Preload("User").First(&post, post.ID)
	return
}

// 获取所有文章列表
func (s *PostService) GetAllPosts(page, pageSize int) (posts []model.Post, total int64, err error) {
	offset := (page - 1) * pageSize

	// 获取总数
	database.DB.Model(&model.Post{}).Count(&total)

	// 获取文章列表（预加载用户信息）
	err = database.DB.Preload("User").
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&posts).Error

	return
}

// 根据ID获取单个文章
func (s *PostService) GetPostByID(id uint) (post model.Post, err error) {
	err = database.DB.Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		First(&post, id).Error
	
	if err != nil {
		err = errors.New("文章不存在")
		return
	}

	return
}

// 更新文章
func (s *PostService) UpdatePost(id, userID uint, form form.UpdatePostForm) (post model.Post, err error) {
	// 先检查文章是否存在且属于当前用户
	err = database.DB.Where("id = ? AND user_id = ?", id, userID).First(&post).Error
	if err != nil {
		err = errors.New("文章不存在或无权限修改")
		return
	}

	// 更新字段
	updates := make(map[string]interface{})
	if form.Title != "" {
		updates["title"] = form.Title
	}
	if form.Content != "" {
		updates["content"] = form.Content
	}

	err = database.DB.Model(&post).Updates(updates).Error
	if err != nil {
		err = errors.New("文章更新失败")
		return
	}

	// 重新获取更新后的文章
	database.DB.Preload("User").First(&post, post.ID)
	return
}

// 删除文章
func (s *PostService) DeletePost(id, userID uint) error {
	// 先检查文章是否存在且属于当前用户
	var post model.Post
	err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&post).Error
	if err != nil {
		return errors.New("文章不存在或无权限删除")
	}

	// 删除文章（软删除）
	err = database.DB.Delete(&post).Error
	if err != nil {
		return errors.New("文章删除失败")
	}

	return nil
} 