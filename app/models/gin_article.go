package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type GinArticle struct {
	gorm.Model
	UUID    uuid.UUID `json:"uuid"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Status  int       `json:"status"`
}
