package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Account   string `gorm:"column:account" json:"account"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
	Email     string `gorm:"column:email" json:"email"`
	ID        int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	LastIP    string `gorm:"column:last_ip" json:"last_ip"`
	LastTime  int64  `gorm:"column:last_time" json:"last_time"`
	Password  string `gorm:"column:password" json:"password"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`
	Username  string `gorm:"column:username" json:"username"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

type Post struct {
	CategoryID int64  `gorm:"column:category_id" json:"category_id"`
	Content    string `gorm:"column:content" json:"content"`
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	ImgURL     string `gorm:"column:img_url" json:"img_url"`
	IsIndex    int64  `gorm:"column:is_index" json:"is_index"`
	IsTop      int64  `gorm:"column:is_top" json:"is_top"`
	Pv         int64  `gorm:"column:pv" json:"pv"`
	Status     int64  `gorm:"column:status" json:"status"`
	Title      string `gorm:"column:title" json:"title"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
	Uv         int64  `gorm:"column:uv" json:"uv"`
}

func (p *Post) TableName() string {
	return "post"
}

func (p *Post) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

type Catrgory struct {
	CreateTime time.Time      `gorm:"column:create_time"`
	ID         int            `gorm:"column:id;primary_key"`
	IsDelete   sql.NullInt64  `gorm:"column:is_delete"`
	Name       sql.NullString `gorm:"column:name"`
	UpdateTime time.Time      `gorm:"column:update_time"`
}

func (c *Catrgory) TableName() string {
	return "category"
}

func (c *Catrgory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}
