package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
	"strconv"
)

func GetAllPosts() ([]models.Post, error) {
	return dao.GetAllPosts()
}

func GetPostByID(id uint) (models.Post, error) {
	post, err := dao.GetPostByID(id)
	if err != nil {
		return post, errors.New("post not found")
	}
	return post, nil
}

func CreatePost(author, title, content, imagePaths string) error {
	post := models.Post{Author: author, Title: title, Content: content, ImagePath: imagePaths}
	return dao.CreatePost(post)
}

func SearchPosts(query string) ([]models.Post, error) {
	if query == "" {
		return nil, errors.New("query cannot be empty")
	}
	id, _ := strconv.Atoi(query)
	return dao.SearchPosts(uint(id), query)
}

func GetPostByAuthor(author string) ([]models.Post, error) {
	posts, err := dao.GetPostByAuthor(author)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
