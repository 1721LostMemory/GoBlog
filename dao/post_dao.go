package dao

import (
	"errors"
	"goblog/config"
	"goblog/models"
)

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := config.MysqlDB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	result := config.MysqlDB.First(&post, id)
	if result.Error != nil {
		return post, errors.New("post not found")
	}
	return post, nil
}

func CreatePost(post models.Post) error {
	result := config.MysqlDB.Create(&post)
	return result.Error
}

func SearchPosts(id uint, query string) ([]models.Post, error) {
	posts := make([]models.Post, 0)
	err := config.MysqlDB.Where("id = ? or author = ? or title LIKE ? or content LIKE ?", id, query, "%"+query+"%", "%"+query+"%").
		Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostByAuthor(author string) ([]models.Post, error) {
	posts := make([]models.Post, 0)
	err := config.MysqlDB.Where("author = ?", author).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
