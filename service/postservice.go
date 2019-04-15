package service

import (
	"blog-cms/conf"
	"blog-cms/models"
)

type PostService struct {
}

func (p *PostService) List() {

}

func (p *PostService) Add(post *models.Post) (*models.Post, error) {
	result := conf.DB.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

func (p *PostService) Update() {

}
