package service

import (
	"blog/module/comment/form"
	"blog/module/common/model"
	"blog/pkg/database"
	"errors"
)

type CommentService struct{}

// 创建评论
func (cs *CommentService) CreateComment(userID uint, commentForm form.CreateCommentForm) (comment model.Comment, err error) {
	// 先检查文章是否存在
	var post model.Post
	if err = database.DB.First(&post, commentForm.PostID).Error; err != nil {
		err = errors.New("文章不存在")
		return
	}

	// 创建评论
	comment = model.Comment{
		Content: commentForm.Content,
		UserID:  userID,
		PostID:  commentForm.PostID,
	}

	err = database.DB.Create(&comment).Error
	if err != nil {
		err = errors.New("评论创建失败")
		return
	}

	// 预加载用户信息
	database.DB.Preload("User").First(&comment, comment.ID)
	return
}

// 获取文章的所有评论
func (cs *CommentService) GetCommentsByPostID(postID uint, page, pageSize int) (comments []model.Comment, total int64, err error) {
	offset := (page - 1) * pageSize

	// 先检查文章是否存在
	var post model.Post
	if err = database.DB.First(&post, postID).Error; err != nil {
		err = errors.New("文章不存在")
		return
	}

	// 获取评论总数
	database.DB.Model(&model.Comment{}).Where("post_id = ?", postID).Count(&total)

	// 获取评论列表
	err = database.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&comments).Error

	return
} 