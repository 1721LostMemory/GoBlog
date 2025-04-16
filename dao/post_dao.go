package dao

import (
	"context"
	"errors"
	"fmt"
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
	if result.Error != nil {
		return result.Error
	}
	key := "rank:user:post"
	if err := config.RedisDB.ZIncrBy(context.Background(), key, 1, post.Author).Err(); err != nil {
		// Redis 写失败：回滚数据库写入
		rollbackErr := config.MysqlDB.Delete(&post).Error
		if rollbackErr != nil {
			// 两边都失败，输出详细日志
			return fmt.Errorf("redis 写失败: %v，数据库回滚失败: %v", err, rollbackErr)
		}
		return fmt.Errorf("redis 写失败，已回滚数据库: %v", err)
	}

	return nil
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
