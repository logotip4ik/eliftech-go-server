package controllers

import (
	"fmt"
	"home/work/models"
	"home/work/utils/token"

	"github.com/gin-gonic/gin"
)

func FindManyPosts(c *gin.Context) {
	var Posts []models.Post

	if err := models.DB.Select("title", "content", "user_id", "created_at", "updated_at").Find(&Posts).Error; err != nil {
		fmt.Printf("error ocurred when fetching books: %s", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{ "data": Posts })
}

func FindOnePost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).Select("title", "content", "user_id", "created_at", "updated_at").First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": fmt.Sprintf("post with id %s not found", c.Param("id"))})
		return
	}

	c.JSON(200, gin.H{"data": post})
}

func UpdateOnePost(c *gin.Context) {
  // Get model if exist
  var post models.Post
  if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(404, gin.H{"error": err})
    return
  }

  // Validate input
  var input models.UpdatePostInput
  if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
    return
  }

	post.Title = input.Title
	post.Content = input.Content

  if err := models.DB.Model(&post).Updates(post).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	models.DB.Table("posts").Select("title", "content", "user_id", "updated_at", "created_at").First(&post, post.ID)
	
  c.JSON(200, gin.H{"data": post})
}

func CreateOnePost(c *gin.Context) {
	var input models.CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error:": err.Error()})
		return
	}

	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	post := models.Post{Title: input.Title, Content: input.Content, UserID: userId}
	if err := models.DB.Select("title", "content", "created_at", "updated_at").Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{"data": post})
}

func DeleteOneBook(c *gin.Context) {
	var post models.Post
  if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
    c.JSON(404, gin.H{"error": err})
    return
  }

  if err := models.DB.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": err});
		return
	}

  c.JSON(200, gin.H{"data": true})
}