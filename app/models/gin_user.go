package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type GinUser struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"userName"`
	Password    string    `json:"-"`
	NickName    string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string    `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string    `json:"authorityId" gorm:"default:888"`
}
