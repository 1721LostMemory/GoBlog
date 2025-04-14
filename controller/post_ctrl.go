package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func ShowPosts(c *gin.Context) {
	posts, err := service.GetAllPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"})
		return
	}

	isLogin, _ := c.Get("isLogin")
	user, _ := c.Get("user")
	//fmt.Println("传给前端的user:", user)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "首页",
		"posts":   posts,
		"isLogin": isLogin,
		"user":    user,
	})
}

func ShowPost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := service.GetPostByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	imagePath := post.ImagePath
	files, err := os.ReadDir(imagePath)

	var imageUrls []string
	dir := filepath.ToSlash(imagePath)

	for _, file := range files {
		imageUrls = append(imageUrls, "/"+dir+"/"+file.Name())
		//fmt.Println("/" + dir + "/" + file.Name())
	}

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"post":      post,
		"imageUrls": imageUrls,
	})
}

func NewPostForm(c *gin.Context) {
	isLogin, _ := c.Get("isLogin")

	if !isLogin.(bool) {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "new.html", nil)
}

func CreatePost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	folderName := c.PostForm("folder_name")
	uploadDir := filepath.Join("static", "uploads", folderName)

	//fmt.Println("folder_name:", folderName)

	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["images"]
		for _, file := range files {
			// 创建上传目录
			err := os.MkdirAll(uploadDir, os.ModePerm)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建上传目录"})
				return
			}

			// 获取文件扩展名
			ext := filepath.Ext(file.Filename)

			// 使用时间戳重命名文件（精确到毫秒）
			filename := fmt.Sprintf("%d%s", time.Now().UnixNano()/int64(time.Millisecond), ext)

			dstPath := filepath.Join(uploadDir, filename)

			if err := c.SaveUploadedFile(file, dstPath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "图片保存失败"})
				return
			}

		}
	}

	user, _ := c.Get("user")
	fmt.Println("作者:", user)
	// 调用 service 保存帖子
	if err := service.CreatePost(user.(string), title, content, uploadDir); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建帖子失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func SearchPosts(c *gin.Context) {
	query := c.Query("query")

	results, err := service.SearchPosts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search posts"})
		return
	}

	c.HTML(http.StatusOK, "search.html", gin.H{
		"Title":   "搜索结果",
		"query":   query,
		"results": results})
}
