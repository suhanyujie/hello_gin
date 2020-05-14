package services

import (
	"hello_gin/app/global"
	"hello_gin/app/models"

	uuid "github.com/satori/go.uuid"
)

func AddArticle(title, content string) (error, models.GinArticle) {
	var article = &models.GinArticle{
		UUID:    uuid.NewV4(),
		Title:   title,
		Content: content,
		Status:  1,
	}
	global.GVA_LOG.Info(article)
	// err := global.GVA_DB.Create(article).Error
	return nil, *article
}
