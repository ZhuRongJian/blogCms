package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Account    string         `gorm:"column:account" `
	CreateTime time.Time      `gorm:"column:create_time"`
	Email      sql.NullString `gorm:"column:email"`
	ID         int            `gorm:"column:id;primary_key"`
	LastIP     sql.NullString `gorm:"column:last_ip"`
	LastTime   time.Time      `gorm:"column:last_time"`
	Password   string         `gorm:"column:password"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	Username   string         `gorm:"column:username"`
}

func (u *User) TableName() string {
	return "post"
}
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateTime", time.Now())
	scope.SetColumn("CreateTime", time.Now())
	return nil
}

type Post struct {
	CategoryID int            `gorm:"column:category_id"`
	Content    string         `gorm:"column:content"`
	CreateTime time.Time      `gorm:"column:create_time"`
	ID         int            `gorm:"column:id;primary_key"`
	ImgURL     sql.NullString `gorm:"column:img_url"`
	IsIndex    sql.NullInt64  `gorm:"column:is_index"`
	IsTop      sql.NullInt64  `gorm:"column:is_top"`
	Status     sql.NullInt64  `gorm:"column:status"`
	Title      string         `gorm:"column:title"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	UserID     int            `gorm:"column:user_id"`
}

func (p *Post) TableName() string {
	return "post"
}

func (p *Post) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateTime", time.Now())
	scope.SetColumn("CreateTime", time.Now())
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
	scope.SetColumn("UpdateTime", time.Now())
	scope.SetColumn("CreateTime", time.Now())
	return nil
}
