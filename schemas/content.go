package schemas

import (
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Title string
	Desc string
	Content_Type string
	Genre string
}
