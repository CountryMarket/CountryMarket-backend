package model

import (
	"gorm.io/gorm"
)

type UserComment struct {
	UserId    int
	ProductId int
	Comment   string
	Visible   bool
	gorm.Model
}

func (m *model) CommentGetProductComment(productId int) ([]UserComment, error) {
	var resComments []UserComment
	err := m.db.Model(&UserComment{}).Where("product_id = ? AND visible = 1", productId).Scan(&resComments).Error
	return resComments, err
}
func (m *model) CommentAddComment(userId, productId int, comment string) error {
	return m.db.Model(&UserComment{}).Create(&UserComment{
		UserId:    userId,
		ProductId: productId,
		Comment:   comment,
		Visible:   true,
	}).Error
}
func (m *model) CommentAddComments(userId int, userComments []UserComment) error {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range userComments {
			err := m.CommentAddComment(userId, v.ProductId, v.Comment)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
func (m *model) CommentDeleteComment(commentId int) error {
	return m.db.Model(&UserComment{}).Where("id = ?", commentId).Update("visible", false).Error
}
